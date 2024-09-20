package hkrpg_go_pe

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/dbconf"
	"github.com/gucooing/hkrpg-go/dispatch/sdk"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/gateserver/session"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/muipserver/api"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

const (
	Ticker                 = 5   // 定时器间隔时间 / s
	AutoUpDataPlayerTicker = 120 // 定时执行玩家数据保存间隔时间 / s
)

type HkRpgGoServer struct {
	config        *Config
	ec2b          *random.Ec2b
	Sdk           *sdk.Server
	KcpConn       *session.KcpConn
	playerMap     map[uint32]*PlayerGame
	playerMapLock *sync.RWMutex // 读写锁

	comm *api.ApiServer

	// 下面是定时器
	everyDay4        *time.Ticker
	autoUpDataPlayer *time.Ticker
}

type PlayerGame struct {
	GamePlayer     *player.GamePlayer // 玩家内存
	S              *session.Session   // 会话
	LastActiveTime int64              // 最近一次的活跃时间
}

// 初始化数据库步骤
func NewServer(cfg *Config) *HkRpgGoServer {
	s := new(HkRpgGoServer)
	s.config = cfg
	s.ec2b = alg.GetEc2b()
	s.playerMapLock = new(sync.RWMutex)
	// 加载res
	gdconf.InitGameDataConfig(cfg.GameDataConfigPath)
	// 初始化数据库
	database.NewPE(cfg.SqlPath)
	dbconf.GameServer(database.PE)
	database.GATE = &database.GateStore{PlayerUidMysql: database.PE}
	database.DISPATCH = &database.DisaptchStore{AccountMysql: database.PE}
	database.GSS = &database.GameStore{PlayerDataMysql: database.PE, ServerConf: database.PE}
	// 初始化dispatch
	s.Sdk = &sdk.Server{
		IsAutoCreate:       cfg.Dispatch.AutoCreate,
		OuterAddr:          fmt.Sprintf("http://%s:%s", cfg.Dispatch.AppNet.OuterAddr, cfg.Dispatch.AppNet.OuterPort),
		RegionInfo:         make(map[string]*sdk.RegionInfo),
		UpstreamServerList: cfg.UpstreamServerList,
		UpstreamServerLock: new(sync.RWMutex),
	}
	for _, d := range cfg.Dispatch.DispatchList {
		s.Sdk.RegionInfo[d.Name] = &sdk.RegionInfo{
			Name:        d.Name,
			Title:       d.Title,
			Type:        alg.S2U32(d.Type),
			Ec2b:        s.ec2b,
			MinGateAddr: cfg.GameServer.AppNet.OuterAddr,
			MinGatePort: alg.S2U32(cfg.GameServer.AppNet.OuterPort),
		}
	}
	gin.SetMode(gin.ReleaseMode) // 初始化gin
	sdkRouter := gin.New()
	sdkRouter.Use(gin.Recovery())
	s.Sdk.GetSdkRouter(sdkRouter) // 初始化路由
	go s.Sdk.UpUpstreamServer()
	go func() {
		err := alg.NewHttp(cfg.Dispatch.AppNet, sdkRouter)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	}()
	// new Kcp
	k, err := session.NewKcpConn(cfg.GameServer.AppNet)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	k.Ec2b = s.ec2b
	s.KcpConn = k
	go func() {
		err = k.RunKcp()
		if err != nil {
			logger.Error(err.Error())
			return
		}
	}()
	session.MAX_CLIENT__CONN_NUM = cfg.MaxPlayer
	go s.loginSessionManagement()
	// new game
	player.ISPE = true
	s.playerMap = make(map[uint32]*PlayerGame)
	// 启动http api
	s.comm = api.NewApiServer(cfg.Gm.SignKey, sdkRouter)
	go s.newHttpApi()
	// 开启game定时器
	s.autoUpDataPlayer = time.NewTicker(AutoUpDataPlayerTicker * time.Second)
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	s.everyDay4 = time.NewTicker(everyDay4)
	go s.gameTicker()
	return s
}

// Session消息队列
func (h *HkRpgGoServer) loginSessionManagement() {
	loginSessionMap := make(map[uint32]*session.Session) // 登录列表
	for {
		select {
		case s := <-h.KcpConn.LoginSessionChan: // 添加登录会话
			loginSessionMap[s.SessionId] = s
			go h.sessionLogin(s)
		case s := <-h.KcpConn.DelLoginSessionChan: // 删除登录会话
			delete(loginSessionMap, s.SessionId)
		}
	}
}

func (h *HkRpgGoServer) sessionMsg(p *PlayerGame) {
	for {
		packMsg, ok := <-p.S.RecvChan
		p.LastActiveTime = time.Now().Unix()
		if !ok || p.S.SessionState == session.SessionClose {
			return
		}
		switch p.S.SessionState {
		case session.SessionLogin:
		case session.SessionActivity:
			protoMsg := cmd.DecodePayloadToProto(packMsg)
			h.packetCapture(p, packMsg.CmdId, protoMsg)
		case session.SessionFreeze:
			continue
		case session.SessionClose:
			return
		}
	}
}

// 接收game传来的消息
func (g *PlayerGame) recvGameMsg() {
	for {
		bin, ok := <-g.GamePlayer.SendChan
		if !ok {
			return
		}
		if g.S.SessionState == session.SessionClose {
			return
		}
		switch bin.MsgType {
		case player.Server:
			protoData, err := pb.Marshal(bin.PlayerMsg)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			g.S.SendChan <- &alg.PackMsg{
				CmdId:     bin.CmdId,
				HeadData:  nil,
				ProtoData: protoData,
			}
		}
	}
}

// 将消息发送给game
func (g *PlayerGame) sendGameMsg(msgType player.MsgType, cmdId uint16, playerMsg pb.Message) {
	if g.S.SessionState == session.SessionClose {
		return
	}
	g.GamePlayer.RecvChan <- player.Msg{
		CmdId:       cmdId,
		MsgType:     msgType,
		PlayerMsg:   playerMsg,
		CommandList: nil,
		CommandId:   0,
		CommandRsp:  "",
	}
}

func (h *HkRpgGoServer) AddPlayer(s *session.Session) *PlayerGame {
	h.playerMapLock.Lock()
	defer h.playerMapLock.Unlock()
	g := &PlayerGame{
		GamePlayer:     player.NewPlayer(s.Uid),
		S:              s,
		LastActiveTime: time.Now().Unix(),
	}
	h.playerMap[s.Uid] = g
	return h.playerMap[s.Uid]
}

func (h *HkRpgGoServer) GetPlayer(uid uint32) *PlayerGame {
	h.playerMapLock.RLock()
	defer h.playerMapLock.RUnlock()
	return h.playerMap[uid]
}

func (h *HkRpgGoServer) GetAllPlayer() map[uint32]*PlayerGame {
	playerMap := make(map[uint32]*PlayerGame)
	h.playerMapLock.RLock()
	for k, v := range h.playerMap {
		playerMap[k] = v
	}
	h.playerMapLock.RUnlock()
	return playerMap
}

func (h *HkRpgGoServer) DelPlayer(uid uint32) {
	h.playerMapLock.Lock()
	p := h.playerMap[uid]
	delete(h.playerMap, uid)
	h.playerMapLock.Unlock()
	if p != nil {
		p.Close()
	}
}

func (h *HkRpgGoServer) sessionLogin(s *session.Session) {
	timeout := time.After(5 * time.Second)
	select {
	case packMsg, ok := <-s.RecvChan:
		if !ok {
			return
		}
		if packMsg.CmdId != cmd.PlayerGetTokenCsReq {
			return
		}
		rsp := h.playerLogin(s, packMsg.ProtoData)
		protoData, err := pb.Marshal(rsp)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		h.KcpConn.DelLoginSessionChan <- s
		if rsp.Retcode == 0 {
			p := h.AddPlayer(s)
			s.SessionState = session.SessionActivity
			atomic.AddInt64(&session.CLIENT_CONN_NUM, 1)
			go h.sessionMsg(p)
			go p.recvGameMsg()
			go p.GamePlayer.RecvMsg()
		}
		s.SendChan <- &alg.PackMsg{
			CmdId:     cmd.PlayerGetTokenScRsp,
			HeadData:  nil,
			ProtoData: protoData,
		}
	case <-timeout:
		logger.Warn("Session login timed out")
		return
	}
}

// 玩家登录
func (h *HkRpgGoServer) playerLogin(s *session.Session, protoData []byte) *proto.PlayerGetTokenScRsp {
	rsp := &proto.PlayerGetTokenScRsp{
		BlackInfo: &proto.BlackInfo{},
	}
	req := new(proto.PlayerGetTokenCsReq)
	err := pb.Unmarshal(protoData, req)
	if err != nil {
		logger.Error("pb Unmarshal PlayerGetTokenCsReq err: %v", err)
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		return rsp
	}
	// 人数验证
	if session.MAX_CLIENT__CONN_NUM != -1 &&
		session.MAX_CLIENT__CONN_NUM <= session.CLIENT_CONN_NUM {
		logger.Info("client conn max")
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		return rsp
	}

	account := database.GetPlayerUidByAccountId(database.GATE.PlayerUidMysql, alg.S2U32(req.AccountUid))

	// token 验证
	if req.Token != account.ComboToken {
		rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_VERIFY_ERROR)
		return rsp
	}

	// ban 验证
	if account.IsBan && account.BanEndTime >= time.Now().Unix() {
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		return rsp
	}
	// 重复登录验证
	if old := h.GetPlayer(account.Uid); old != nil {
		// 通知客户端下线
		bin, _ := pb.Marshal(&proto.PlayerKickOutScNotify{
			BlackInfo: &proto.BlackInfo{},
		})
		old.S.SendChan <- &alg.PackMsg{
			CmdId:     cmd.PlayerKickOutScNotify,
			HeadData:  nil,
			ProtoData: bin,
		}
		h.DelPlayer(account.Uid)
	}

	// 回包
	s.Uid = account.Uid
	if s.XorKey != nil {
		s.Seed = random.GetTimeRand().Uint64()
	}
	rsp.Uid = s.Uid
	rsp.SecretKeySeed = s.Seed

	return rsp
}

func (h *HkRpgGoServer) gameTicker() {
	for {
		select {
		case <-h.autoUpDataPlayer.C:
			h.AutoUpDataPlayer()
		case <-h.everyDay4.C: // 4点事件
			h.GlobalRotationEvent4h()
		}
	}
}

func (h *HkRpgGoServer) AutoUpDataPlayer() {
	playerList := h.GetAllPlayer()
	if len(playerList) == 0 {
		return
	}
	timestamp := time.Now().Unix()
	logger.Info("开始自动保存玩家数据")
	var num int
	for _, g := range playerList {
		if g.LastActiveTime+50 < timestamp {
			logger.Info("[UID:%v]玩家长时间无响应离线", g.S.Uid)
			h.DelPlayer(g.S.Uid)
			continue
		}
		lastUpDataTime := g.GamePlayer.LastUpDataTime
		if timestamp-lastUpDataTime >= 180 {
			logger.Debug("[UID:%v]玩家数据自动保存", g.S.Uid)
			g.GamePlayer.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_ONLINE)
			g.GamePlayer.LastUpDataTime = timestamp + rand.Int63n(120)
			num++
		}
	}
	logger.Info("保存玩家数据结束,保存玩家数量:%v", num)
}

func (h *HkRpgGoServer) GlobalRotationEvent4h() {
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	h.everyDay4 = time.NewTicker(everyDay4)
}

func (g *PlayerGame) Close() {
	if g.S.SessionState == session.SessionClose {
		return
	}
	// 下线GATE
	g.S.Close()
	// 下线GS
	g.GamePlayer.Close()
}

func (h *HkRpgGoServer) Close() {
	h.killAutoUpDataPlayer() // 保存玩家数据
}

func (h *HkRpgGoServer) killAutoUpDataPlayer() {
	logger.Info("开始保存玩家数据")
	var num int
	playerList := h.GetAllPlayer()
	for _, g := range playerList {
		g.Close()
		num++
	}
	logger.Info("保存玩家数据结束,保存玩家数量:%v", num)
}
