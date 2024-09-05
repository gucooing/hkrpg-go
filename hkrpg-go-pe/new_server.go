package hkrpg_go_pe

import (
	"context"
	"log"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/dbconf"
	"github.com/gucooing/hkrpg-go/dispatch"
	"github.com/gucooing/hkrpg-go/gameserver"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/gateserver"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen = 343 * 1024 // 最大应用层包长度
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数
var QPS int64

type HkRpgGoServer struct {
	config           *Config
	db               *database.DisaptchStore
	apiRouter        *gin.Engine // http api
	Dispatch         *dispatch.Server
	kcpListener      *kcp.Listener
	sessionIdCounter uint32
	kcpEventChan     chan *gateserver.KcpEvent
	playerMap        map[uint32]*PlayerGame
	playerMapLock    sync.Mutex // 玩家列表互斥锁
	// 下面是定时器
	everyDay4        *time.Ticker
	autoUpDataPlayer *time.Ticker
	CmdRouteManager  *CmdRouteManager
}

func newStorePE(cfg *Config) *database.DisaptchStore {
	s := new(database.DisaptchStore)
	s.AccountMysql = database.NewSqlite(cfg.SqlPath)
	s.AccountMysql.AutoMigrate(
		&constant.Account{},      // sdk账户
		&constant.PlayerUid{},    // 映射表
		&constant.PlayerData{},   // 玩家数据
		&constant.BlockData{},    // 地图数据
		&constant.RogueConf{},    // 模拟宇宙配置
		&constant.ScheduleConf{}, // 忘记了
		&constant.PlayerBasic{},  // 好友简要信息
		&constant.Mail{},         // 邮件配置
		&constant.PlayerMail{},   // 玩家邮件配置
		&constant.RegionConf{},   // 区服配置
	)

	logger.Info("数据库连接成功")
	return s
}

// 初始化数据库步骤
func NewServer(cfg *Config) *HkRpgGoServer {
	s := new(HkRpgGoServer)
	s.config = cfg
	// 加载res
	gdconf.InitGameDataConfig(cfg.GameDataConfigPath)
	// 初始化数据库
	s.db = newStorePE(cfg)
	dbconf.GameServer(s.db.AccountMysql)
	database.GSS = &database.GameStore{PeMysql: s.db.AccountMysql}
	// 初始化dispatch
	gin.SetMode(gin.ReleaseMode) // 初始化gin
	dispatchList := make([]dispatch.Dispatch, 0)
	for _, d := range cfg.Dispatch.DispatchList {
		dispatchList = append(dispatchList, dispatch.Dispatch{
			Name:        d.Name,
			Title:       d.Title,
			Type:        d.Type,
			DispatchUrl: d.DispatchUrl,
		})
	}
	s.Dispatch = &dispatch.Server{
		Router:       gin.New(),
		Ec2b:         alg.GetEc2b(),
		IsAutoCreate: cfg.Dispatch.AutoCreate,
		Store:        s.db,
		InnerAddr:    cfg.Dispatch.Addr,
		Port:         cfg.Dispatch.Port,
		OuterAddr:    cfg.Dispatch.OuterAddr,
		DispatchList: dispatchList,
		KcpPort:      alg.S2U32(cfg.GameServer.OuterPort),
		KcpIp:        cfg.GameServer.OuterAddr,
		IsPe:         true,
	}
	s.Dispatch.Router.Use(gin.Recovery())
	// 启动kcp
	addr := cfg.GameServer.InnerAddr + ":" + cfg.GameServer.InnerPort
	logger.Info("KCP监听地址:%s", addr)
	logger.Info("KCP对外地址:%s", cfg.GameServer.OuterAddr+":"+cfg.GameServer.OuterPort)
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		log.Printf("listen kcp err: %v\n", err)
		os.Exit(0)
	}
	s.kcpListener = kcpListener
	s.kcpListener.EnetHandle()
	go kcpNetInfo()
	s.playerMap = make(map[uint32]*PlayerGame)
	s.CmdRouteManager = NewCmdRouteManager()
	// 启动http api
	go s.newHttpApi()
	// 开启game定时器
	s.autoUpDataPlayer = time.NewTicker(gameserver.AutoUpDataPlayerTicker * time.Second)
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	s.everyDay4 = time.NewTicker(everyDay4)
	go s.gameTicker()
	return s
}

func (s *HkRpgGoServer) gameTicker() {
	for {
		select {
		case <-s.autoUpDataPlayer.C:
			s.AutoUpDataPlayer()
		case <-s.everyDay4.C: // 4点事件
			s.GlobalRotationEvent4h()
		}
	}
}

func (s *HkRpgGoServer) AutoUpDataPlayer() {
	logger.Info("开始自动保存玩家数据")
	timestamp := time.Now().Unix()
	playerList := s.getAllPlayer()
	var num int
	for _, g := range playerList {
		if g.Uid == 0 {
			continue
		}
		if g.LastActiveTime+50 < timestamp {
			logger.Info("[UID:%v]玩家长时间无响应离线", g.Uid)
			s.killPlayer(g)
			continue
		}
		lastUpDataTime := g.GamePlayer.LastUpDataTime
		if timestamp-lastUpDataTime >= 180 {
			logger.Debug("[UID:%v]玩家数据自动保存", g.Uid)
			g.GamePlayer.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_ONLINE)
			g.GamePlayer.LastUpDataTime = timestamp + rand.Int63n(120)
			num++
		}
	}
	logger.Info("保存玩家数据结束,保存玩家数量:%v", num)
}

func (s *HkRpgGoServer) GlobalRotationEvent4h() {
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	s.everyDay4 = time.NewTicker(everyDay4)
}

func (s *HkRpgGoServer) RunGameServer() error {
	for {
		kcpConn, err := s.kcpListener.AcceptKCP()
		if err != nil {
			logger.Error("accept kcp err: %v", err)
			return err
		}
		go func() {
			CLIENT_CONN_NUM++
			kcpConn.SetACKNoDelay(true)
			kcpConn.SetWriteDelay(false)
			kcpConn.SetWindowSize(256, 256)
			kcpConn.SetMtu(1200)
			// 读取密钥相关文件
			g := s.NewGame(kcpConn)
			s.recvHandle(g)
		}()
	}
}

// kcp统计
func kcpNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	kcpErrorCount := uint64(0)
	for {
		<-ticker.C
		snmp := kcp.DefaultSnmp.Copy()
		kcpErrorCount += snmp.KCPInErrors
		logger.Debug("kcp send: %v B/s, kcp recv: %v B/s", snmp.BytesSent/60, snmp.BytesReceived/60)
		logger.Debug("udp send: %v B/s, udp recv: %v B/s", snmp.OutBytes/60, snmp.InBytes/60)
		logger.Debug("udp send: %v pps, udp recv: %v pps", snmp.OutPkts/60, snmp.InPkts/60)
		clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
		logger.Debug("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		logger.Debug("QPS: %v /s", QPS/60)
		QPS = 0
		kcp.DefaultSnmp.Reset()
	}
}

func (s *HkRpgGoServer) recvHandle(p *PlayerGame) {
	payload := make([]byte, PacketMaxLen)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker player uid: %v", p.Uid)
			CLIENT_CONN_NUM--
			s.killPlayer(p)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := p.KcpConn.Read(payload)
		QPS++
		if err != nil {
			CLIENT_CONN_NUM--
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = payload[:recvLen]
		kcpMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &kcpMsgList, p.XorKey)
		for _, msg := range kcpMsgList {
			playerMsg := cmd.DecodePayloadToProto(msg)
			if playerMsg == nil {
				logger.Warn("[UID:%v]DecodePayloadToProto error", p.Uid)
				continue
			}
			if msg.CmdId == cmd.PlayerLogoutCsReq {
				logger.Info("[UID:%v]玩家主动离线", p.Uid)
				s.killPlayer(p)
				return
			}
			if p.IsLogin {
				s.RegisterMessage(p, msg.CmdId, playerMsg)
			} else {
				if msg.CmdId == cmd.PlayerGetTokenCsReq {
					s.PlayerGetTokenCsReq(p, playerMsg)
				} else {
					return
				}
			}
		}
	}
}

// 发送事件处理
func (p *PlayerGame) SendHandle(cmdId uint16, playerMsg pb.Message) {
	if p.KcpConn == nil {
		return
	}
	rspMsg := new(cmd.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	kcpMsg := cmd.EncodeProtoToPayload(rspMsg)
	binMsg := alg.EncodePayloadToBin(kcpMsg, p.XorKey)
	// 密钥交换
	if kcpMsg.CmdId == cmd.PlayerGetTokenScRsp {
		if p.Seed == 0 {
			return
		}
		p.XorKey = random.CreateXorPad(p.Seed, false)
		logger.Info("[UID:%v][SEED:%v]玩家登录成功", p.Uid, p.Seed)
	}
	_, err := p.KcpConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

type PlayerGame struct {
	IsLogin        bool
	Seed           uint64
	Uid            uint32
	XorKey         []byte // 密钥
	KcpConn        *kcp.UDPSession
	GamePlayer     *player.GamePlayer // 玩家内存
	LastActiveTime int64              // 最近一次的活跃时间
}

func (s *HkRpgGoServer) NewGame(kcpConn *kcp.UDPSession) *PlayerGame {
	pg := &PlayerGame{
		IsLogin: false,
		Seed:    0,
		XorKey:  s.Dispatch.Ec2b.XorKey(),
		KcpConn: kcpConn,
	}
	return pg
}

func (s *HkRpgGoServer) PlayerGetTokenCsReq(p *PlayerGame, playerMsg pb.Message) {
	req := playerMsg.(*proto.PlayerGetTokenCsReq)
	rsp := new(proto.PlayerGetTokenScRsp)
	if req.Token == "" || req.AccountUid == "" {
		logger.Info("异常的登录请求 Token:%v AccountUid:%s", req.Token, req.AccountUid)
		return
	}
	// 人数验证
	if s.config.GameServer.MaxPlayer != -1 {
		if CLIENT_CONN_NUM >= GetConfig().GameServer.MaxPlayer {
			rsp.Uid = 0
			rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
			rsp.Msg = "当前服务器人数过多，请稍后再试。"
			p.SendHandle(cmd.PlayerGetTokenScRsp, rsp)
			return
		}
	}
	accountUid := alg.S2U32(req.AccountUid)
	playerUid := database.GetPlayerUidByAccountId(s.Dispatch.Store.AccountMysql, accountUid)
	// token验证
	if playerUid.ComboToken != req.Token {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "token验证失败"
		p.SendHandle(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}
	// 封禁验证
	if playerUid.BanEndTime >= time.Now().Unix() {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = playerUid.BanMsg
		rsp.BlackInfo = &proto.BlackInfo{
			BeginTime:  playerUid.BanBeginTime,
			EndTime:    playerUid.BanEndTime,
			LimitLevel: 0,
			BanType:    第三方辅助,
		}
		p.SendHandle(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁,原因:%s", accountUid, playerUid.BanMsg)
		return
	}
	// 重复登录验证
	if old := s.GetPlayer(playerUid.Uid); old != nil {
		s.takeKillPlayer(old, spb.PlayerOfflineReason_OFFLINE_REPEAT_LOGIN)
		logger.Info("[UID:%v]重复登录", playerUid.Uid)
	}
	p.Uid = playerUid.Uid
	// 拉取玩家数据
	p.GamePlayer = s.NewPlayer(playerUid.Uid, accountUid)
	p.LastActiveTime = time.Now().Unix()
	go p.recvPlayer()
	go p.GamePlayer.RecvMsg()
	s.addPlayer(p) // 添加角色到列表中
	p.GamePlayer.GetPlayerDateByDb()
	// 生成seed
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	p.Seed = serverSeedUint64
	rsp.SecretKeySeed = p.Seed
	rsp.BlackInfo = &proto.BlackInfo{}
	rsp.Uid = p.Uid
	p.SendHandle(cmd.PlayerGetTokenScRsp, rsp)
}

func (s *HkRpgGoServer) NewPlayer(uid, accountId uint32) *player.GamePlayer {
	g := new(player.GamePlayer)
	g.Uid = uid
	g.AccountId = accountId
	g.RecvChan = make(chan player.Msg)
	g.SendChan = make(chan player.Msg)
	g.RecvCtx, g.RecvCal = context.WithCancel(context.Background())
	g.SendCtx, g.SendCal = context.WithCancel(context.Background())
	g.IsJumpMission = s.config.GameServer.IsJumpMission
	g.Store = database.GSS // TODO
	g.IsPE = true
	g.RouteManager = player.NewRouteManager(g)
	g.LastUpDataTime = time.Now().Unix()

	return g
}

func (s *HkRpgGoServer) RegisterMessage(p *PlayerGame, cmdId uint16, payloadMsg pb.Message) {
	p.LastActiveTime = time.Now().Unix()
	if p.GamePlayer.RecvChan == nil {
		return
	}
	timeout := time.After(2 * time.Second)
	select {
	case p.GamePlayer.RecvChan <- player.Msg{
		CmdId:     cmdId,
		MsgType:   player.Client,
		PlayerMsg: payloadMsg,
	}:
		if p.GamePlayer.IsClosed {
			close(p.GamePlayer.RecvChan)
		}
	case <-timeout:
		return
	}
}

func (p *PlayerGame) recvPlayer() {
	for {
		select {
		case bin, ok := <-p.GamePlayer.SendChan:
			if !ok {
				return
			}
			switch bin.MsgType {
			case player.Server:
				p.SendHandle(bin.CmdId, bin.PlayerMsg)
			}
		case <-p.GamePlayer.SendCtx.Done():
			p.GamePlayer.IsClosed = true
			return
		}
	}
}

func (s *HkRpgGoServer) addPlayer(p *PlayerGame) {
	s.playerMapLock.Lock()
	p.IsLogin = true
	s.playerMap[p.Uid] = p
	s.playerMapLock.Unlock()
}

func (s *HkRpgGoServer) GetPlayer(uid uint32) *PlayerGame {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	return s.playerMap[uid]
}

func (s *HkRpgGoServer) getAllPlayer() map[uint32]*PlayerGame {
	playerMap := make(map[uint32]*PlayerGame)
	s.playerMapLock.Lock()
	for k, v := range s.playerMap {
		playerMap[k] = v
	}
	s.playerMapLock.Unlock()
	return playerMap
}

func (s *HkRpgGoServer) killPlayer(p *PlayerGame) {
	p.KcpConn.CloseConn(0) // 断开kcp连接
	if p.GamePlayer.SendCal != nil {
		p.GamePlayer.SendCal()
	}
	if p.GamePlayer.RecvCal != nil {
		p.GamePlayer.RecvCal()
	}
	p.GamePlayer.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_OFFLINE) // 保存数据
	delete(s.playerMap, p.Uid)
}

func (s *HkRpgGoServer) takeKillPlayer(p *PlayerGame, status spb.PlayerOfflineReason) {
	var kickType proto.KickType
	switch status {
	case spb.PlayerOfflineReason_OFFLINE_GAME_ERROR:
		kickType = proto.KickType_KICK_BY_GM
	case spb.PlayerOfflineReason_OFFLINE_REPEAT_LOGIN:
		kickType = proto.KickType_KICK_BLACK // KickType_KICK_SQUEEZED
	}
	s.PlayerKickOutScNotify(p, kickType, 使用外挂)
	s.killPlayer(p)
}

const (
	违反用户协议 = 0
	使用外挂   = 1
	第三方辅助  = 2
	发布违规信息 = 3
	登录存在异常 = 4
	账号异常   = 5
)

func (s *HkRpgGoServer) PlayerKickOutScNotify(p *PlayerGame, kickType proto.KickType, banType uint32) {
	p.SendHandle(cmd.PlayerKickOutScNotify, &proto.PlayerKickOutScNotify{
		BlackInfo: &proto.BlackInfo{
			BeginTime:  time.Now().Unix(),
			EndTime:    4294967295,
			LimitLevel: 0,
			BanType:    banType,
		},
		KickType: kickType,
	})
}

func (s *HkRpgGoServer) Close() {
	s.killAutoUpDataPlayer() // 保存玩家数据
}

func (s *HkRpgGoServer) killAutoUpDataPlayer() {
	logger.Info("开始保存玩家数据")
	var num int
	playerList := s.getAllPlayer()
	for _, g := range playerList {
		s.takeKillPlayer(g, spb.PlayerOfflineReason_OFFLINE_GAME_ERROR)
		num++
	}
	logger.Info("保存玩家数据结束,保存玩家数量:%v", num)
}
