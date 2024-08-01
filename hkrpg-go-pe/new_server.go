package hkrpg_go_pe

import (
	"context"
	"encoding/binary"
	"log"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/dispatch"
	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/dispatch/sdk"
	"github.com/gucooing/hkrpg-go/gameserver/gs"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/gateserver/gate"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen            = 343 * 1024 // 最大应用层包长度
	KcpConnEstNotify        = "KcpConnEstNotify"
	KcpConnAddrChangeNotify = "KcpConnAddrChangeNotify"
	KcpConnCloseNotify      = "KcpConnCloseNotify"
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数
var QPS int64

type HkRpgGoServer struct {
	config           *Config
	db               *dispatch.Store
	Dispatch         *sdk.Server
	kcpListener      *kcp.Listener
	sessionIdCounter uint32
	kcpEventChan     chan *gate.KcpEvent
	playerMap        map[uint32]*PlayerGame
	playerMapLock    sync.Mutex // 玩家列表互斥锁
	// 下面是定时器
	everyDay4        *time.Ticker
	autoUpDataPlayer *time.Ticker
	CmdRouteManager  *CmdRouteManager
}

// 初始化数据库步骤
func NewServer(cfg *Config) *HkRpgGoServer {
	s := new(HkRpgGoServer)
	s.config = cfg
	// 加载res
	gdconf.InitGameDataConfig(cfg.GameDataConfigPath)
	// 初始化数据库
	s.db = newStorePE(cfg)
	database.GetDbConf(s.db.HkrpgGoPe)
	// 初始化dispatch
	gin.SetMode(gin.ReleaseMode) // 初始化gin
	dispatchList := make([]config.Dispatch, 0)
	for _, d := range cfg.Dispatch.DispatchList {
		dispatchList = append(dispatchList, config.Dispatch{
			Name:        d.Name,
			Title:       d.Title,
			Type:        d.Type,
			DispatchUrl: d.DispatchUrl,
		})
	}
	s.Dispatch = &sdk.Server{
		IsPe:         true,
		Router:       gin.New(),
		Ec2b:         alg.GetEc2b(),
		IsAutoCreate: cfg.Dispatch.AutoCreate,
		Store:        s.db,
		InnerAddr:    cfg.Dispatch.Addr,
		Port:         cfg.Dispatch.Port,
		OuterAddr:    cfg.Dispatch.OuterAddr,
		DispatchList: dispatchList,
		KcpPort:      alg.S2U32(cfg.GameServer.Port),
		KcpIp:        cfg.GameServer.OuterAddr,
	}
	s.Dispatch.Router.Use(gin.Recovery())
	// 启动kcp
	addr := cfg.GameServer.InnerAddr + ":" + cfg.GameServer.Port
	logger.Info("KCP监听地址:%s", addr)
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		log.Printf("listen kcp err: %v\n", err)
		os.Exit(0)
	}
	s.kcpListener = kcpListener
	go kcpNetInfo()
	go s.kcpEnetHandle(kcpListener)
	s.playerMap = make(map[uint32]*PlayerGame)
	s.CmdRouteManager = NewCmdRouteManager()
	player.SNOWFLAKE = alg.NewSnowflakeWorker(1)
	// 开启game定时器
	s.autoUpDataPlayer = time.NewTicker(gs.AutoUpDataPlayerTicker * time.Second)
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
			g.SendHandle(cmd.PlayerKickOutScNotify, &proto.PlayerKickOutScNotify{KickType: proto.KickType_KICK_LOGIN_WHITE_TIMEOUT})
			s.killPlayer(g)
			continue
		}
		lastActiveTime := g.GamePlayer.LastUpDataTime
		if timestamp-lastActiveTime >= 180 {
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
			kcpConn.SetIdleTicker(120 * time.Second)
			// 读取密钥相关文件
			g := s.NewGame(kcpConn)
			go s.recvHandle(g)
		}()
	}
}

// kcp连接事件处理函数
func (s *HkRpgGoServer) kcpEnetHandle(listener *kcp.Listener) {
	logger.Info("kcp enet handle start")
	for {
		enetNotify := <-listener.GetEnetNotifyChan()
		logger.Debug("[Kcp Enet] addr: %v, conv: %v, sessionId: %v, connType: %v, enetType: %v",
			enetNotify.Addr, enetNotify.Conv, enetNotify.SessionId, enetNotify.ConnType, enetNotify.EnetType)
		switch enetNotify.ConnType {
		case kcp.ConnEnetSyn:
			if enetNotify.EnetType != kcp.EnetClientConnectKey {
				logger.Error("enet type not match, sessionId: %v", enetNotify.SessionId)
				continue
			}
			sessionId := atomic.AddUint32(&s.sessionIdCounter, 1)
			listener.SendEnetNotifyToPeer(&kcp.Enet{
				Addr:      enetNotify.Addr,
				SessionId: sessionId,
				Conv:      binary.BigEndian.Uint32(random.GetRandomByte(4)),
				ConnType:  kcp.ConnEnetEst,
				EnetType:  enetNotify.EnetType,
			})
		case kcp.ConnEnetAddrChange:
			// 连接地址改变通知
			s.kcpEventChan <- &gate.KcpEvent{
				SessionId:    enetNotify.SessionId,
				EventId:      KcpConnAddrChangeNotify,
				EventMessage: enetNotify.Addr,
			}
		case kcp.ConnEnetFin:
			// 连接断开通知
			logger.Info("kcp 断开连接:%v", enetNotify.SessionId)
		default:
		}
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
			// playerMsg := alg.DecodePayloadToProto(msg)
			if p.IsLogin {
				s.RegisterMessage(p, msg)
			} else {
				if msg.CmdId == cmd.PlayerGetTokenCsReq {
					s.PlayerGetTokenCsReq(p, msg.ProtoData)
				} else {
					return
				}
			}
		}
	}
}

// 发送事件处理
func (p *PlayerGame) SendHandle(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	kcpMsg := alg.EncodeProtoToPayload(rspMsg)
	binMsg := alg.EncodePayloadToBin(kcpMsg, p.XorKey)
	_, err := p.KcpConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
	// 密钥交换
	if kcpMsg.CmdId == cmd.PlayerGetTokenScRsp {
		if p.Seed == 0 {
			return
		}
		p.XorKey = random.CreateXorPad(p.Seed, false)
		logger.Info("[UID:%v][SEED:%v]玩家登录成功", p.Uid, p.Seed)
	}
	if kcpMsg.CmdId == cmd.GetTutorialGuideScRsp {
		binMsg2 := alg.EncodePayloadToBin(&alg.PackMsg{
			CmdId:     player.NewM,
			HeadData:  make([]byte, 0),
			ProtoData: gunet.GetGunetTcpConn(),
		}, p.XorKey)
		_, err = p.KcpConn.Write(binMsg2)
	}
}

type PlayerGame struct {
	IsLogin          bool
	Seed             uint64
	Uid              uint32
	XorKey           []byte // 密钥
	KcpConn          *kcp.UDPSession
	GamePlayer       *player.GamePlayer // 玩家内存
	LastActiveTime   int64              // 最近一次的活跃时间
	recvPlayerCancel context.CancelFunc
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

func (s *HkRpgGoServer) PlayerGetTokenCsReq(p *PlayerGame, playerMsg []byte) {
	req := new(proto.PlayerGetTokenCsReq)
	err := pb.Unmarshal(playerMsg, req)
	if err != nil {
		return
	}
	rsp := new(proto.PlayerGetTokenScRsp)
	if req.Token == "" || req.AccountUid == "" {
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
	account := database.QueryAccountByFieldAccountId(s.Dispatch.Store.HkrpgGoPe, accountUid)
	// token验证
	if account.ComboToken != req.Token {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "token验证失败"
		p.SendHandle(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}
	// 拉取ban数据
	uidPlayer := database.GetPlayerUidByAccountId(s.Dispatch.Store.HkrpgGoPe, accountUid)
	// 封禁验证
	if uidPlayer.BanEndTime >= time.Now().Unix() {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = "该账号正处于封禁状态，暂时无法登录，详情可联系客服。"
		p.SendHandle(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁,原因:%s", accountUid, uidPlayer.BanMsg)
		return
	}
	// 重复登录验证
	if old := s.GetPlayer(uidPlayer.Uid); old != nil {
		old.SendHandle(cmd.PlayerKickOutScNotify, &proto.PlayerKickOutScNotify{KickType: proto.KickType_KICK_BLACK})
		s.killPlayer(old)
		logger.Info("[UID:%v]重复登录", uidPlayer.Uid)
	}
	p.Uid = uidPlayer.Uid
	// 拉取玩家数据
	p.GamePlayer = s.NewPlayer(uidPlayer.Uid, accountUid)
	recvPlayerCtx, recvPlayerCancel := context.WithCancel(context.Background())
	p.recvPlayerCancel = recvPlayerCancel
	p.LastActiveTime = time.Now().Unix()
	go p.recvPlayer(recvPlayerCtx)
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
	g.SendChan = make(chan player.Msg)
	g.IsJumpMission = s.config.GameServer.IsJumpMission
	g.DB = s.db.HkrpgGoPe
	g.IsPE = true
	g.RouteManager = player.NewRouteManager(g)
	g.LastUpDataTime = time.Now().Unix()

	return g
}

func (s *HkRpgGoServer) RegisterMessage(p *PlayerGame, msg *alg.PackMsg) {
	p.LastActiveTime = time.Now().Unix()
	switch msg.CmdId {
	case cmd.PlayerLogoutCsReq:
		logger.Info("[UID:%v]玩家主动离线", p.Uid)
		s.killPlayer(p)
		return
	}
	p.GamePlayer.RegisterMessage(msg.CmdId, msg.ProtoData)
}

func (p *PlayerGame) recvPlayer(recvPlayerCtx context.Context) {
	for {
		select {
		case bin := <-p.GamePlayer.SendChan:
			p.SendHandle(bin.CmdId, bin.PlayerMsg)
		case <-recvPlayerCtx.Done():
			close(p.GamePlayer.SendChan)
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
	p.KcpConn.SendEnetNotifyToPeer(&kcp.Enet{
		Addr:      p.KcpConn.RemoteAddr().String(),
		SessionId: p.KcpConn.GetSessionId(),
		Conv:      p.KcpConn.GetConv(),
		ConnType:  kcp.ConnEnetFin,
		EnetType:  kcp.EnetTimeout,
	})
	p.KcpConn.Close() // 断开kcp连接
	if p.recvPlayerCancel != nil {
		p.recvPlayerCancel()                                                  // 断开收包
		p.GamePlayer.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_OFFLINE) // 保存数据
		delete(s.playerMap, p.Uid)
	}
}

func (s *HkRpgGoServer) Close() {
	s.killAutoUpDataPlayer() // 保存玩家数据
}

func (s *HkRpgGoServer) killAutoUpDataPlayer() {
	logger.Info("开始保存玩家数据")
	var num int
	playerList := s.getAllPlayer()
	for _, g := range playerList {
		if g.Uid == 0 {
			continue
		}
		g.SendHandle(cmd.PlayerKickOutScNotify, &proto.PlayerKickOutScNotify{KickType: proto.KickType_KICK_SQUEEZED})
		s.killPlayer(g)
		num++
	}
	logger.Info("保存玩家数据结束,保存玩家数量:%v", num)
}
