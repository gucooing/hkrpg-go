/*
好用的kcp
爱 来自 hk4e-go
*/
package gate

import (
	"encoding/binary"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen            = 343 * 1024 // 最大应用层包长度
	KcpConnAddrChangeNotify = "KcpConnAddrChangeNotify"
)

var syncPl sync.Mutex
var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数
var GATESERVER *GateServer

type GateServer struct {
	AppId            string
	WorkerId         int64
	Port             string
	Config           *config.Config
	Store            *Store
	snowflake        *alg.SnowflakeWorker // 雪花唯一id生成器
	kcpListener      *kcp.Listener
	nodeConn         net.Conn
	kcpFin           bool
	sessionIdCounter uint32
	sessionMap       map[uint32]*PlayerGame
	waitingLoginMap  map[uint32]*PlayerGame
	playerMap        map[int64]*PlayerGame // 玩家内存
	kcpEventChan     chan *KcpEvent
	gameAppId        string                  // 最优appid
	gameAll          map[string]*serviceGame // 从node拉取的game列表
	errGameAppId     []string
	Ec2b             *random.Ec2b

	RecvCh chan *TcpNodeMsg
	Ticker *time.Ticker
	Stop   chan struct{}
}

type PlayerGame struct {
	GameAppId string
	// IsToken             bool // 是否通过token验证
	Status         spb.PlayerStatus
	Uid            uint32 // uid
	AccountId      uint32
	Uuid           int64 // 唯一临时uuid
	Seed           uint64
	XorKey         []byte // 密钥
	KcpConn        *kcp.UDPSession
	GameConn       net.Conn
	LastActiveTime int64 // 最近一次的活跃时间
	ticker         *time.Timer
}

type KcpEvent struct {
	SessionId    uint32
	EventId      string
	EventMessage any
}

type serviceGame struct {
	addr  string
	num   uint64
	appId string
	port  string
}

type TcpNodeMsg struct {
	cmdId      uint16
	serviceMsg pb.Message
}

func NewGate(cfg *config.Config) *GateServer {
	s := new(GateServer)
	// TODO
	GATESERVER = s

	s.Ec2b = alg.GetEc2b()
	s.Config = cfg
	s.Store = NewStore(s.Config) // 初始化数据库连接
	s.sessionMap = make(map[uint32]*PlayerGame)
	s.waitingLoginMap = make(map[uint32]*PlayerGame)
	s.playerMap = make(map[int64]*PlayerGame)
	s.AppId = alg.GetAppId()
	s.WorkerId = 1
	s.snowflake = alg.NewSnowflakeWorker(s.WorkerId)
	logger.Info("GateServer AppId:%s", s.AppId)
	port := s.Config.AppList[s.AppId].App["port_player"].Port
	if port == "" {
		log.Println("GateServer Port error")
		os.Exit(0)
	}
	s.Port = port

	addr := "0.0.0.0:" + s.Port
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		log.Printf("listen kcp err: %v\n", err)
		os.Exit(0)
	}
	s.kcpListener = kcpListener

	s.RecvCh = make(chan *TcpNodeMsg)
	s.Ticker = time.NewTicker(5 * time.Second)
	s.Stop = make(chan struct{})
	s.ServiceStart()

	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		log.Printf("nodeserver error:%s\n", err.Error())
		os.Exit(0)
	}
	s.nodeConn = tcpConn
	s.gameAll = make(map[string]*serviceGame)
	s.errGameAppId = make([]string, 0)

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATESERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			Close()
			os.Exit(0)
		}
	}()

	go s.recvNode()
	go s.kcpNetInfo()
	go s.kcpEnetHandle(kcpListener)
	go s.AutoUpDataPlayer()
	// 向node注册
	s.Connection()

	return s
}

func (s *GateServer) Run() error {
	for {
		kcpConn, err := s.kcpListener.AcceptKCP()
		if s.kcpFin {
			logger.Info("kcp error")
			break
		}
		if err != nil {
			logger.Error("accept kcp err: %v", err)
			continue
		}

		kcpConn.SetACKNoDelay(true)
		kcpConn.SetWriteDelay(false)
		kcpConn.SetWindowSize(256, 256)
		kcpConn.SetMtu(1200)
		sessionId := kcpConn.GetSessionId()
		logger.Info("sessionId:%v", sessionId)
		// 读取密钥相关文件
		g := s.NewGame(kcpConn)
		go s.recvHandle(g)
	}
	return nil
}

func (s *GateServer) NewGame(kcpConn *kcp.UDPSession) *PlayerGame {
	g := new(PlayerGame)
	g.KcpConn = kcpConn
	g.XorKey = s.Ec2b.XorKey()
	g.LastActiveTime = time.Now().Unix()

	return g
}

func (s *GateServer) recvHandle(p *PlayerGame) {
	payload := make([]byte, PacketMaxLen)

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker player uid: %v", p.Uid)
			p.Status = spb.PlayerStatus_PlayerStatus_PassiveOffline
			KickPlayer(p)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := p.KcpConn.Read(payload)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			p.Status = spb.PlayerStatus_PlayerStatus_Offline
			return
		}
		bin = payload[:recvLen]
		kcpMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &kcpMsgList, p.XorKey)
		for _, msg := range kcpMsgList {
			// playerMsg := alg.DecodePayloadToProto(msg)
			switch p.Status {
			case spb.PlayerStatus_PlayerStatus_PreLogin:
				if msg.CmdId == cmd.PlayerGetTokenCsReq {
					p.Status = spb.PlayerStatus_PlayerStatus_LoggingIn
					// s.HandlePlayerGetTokenCsReq(p, msg.ProtoData)
					s.PlayerGetTokenCsReq(p, msg.ProtoData)
				} else {
					p.KcpConn.Close()
					return
				}
			case spb.PlayerStatus_PlayerStatus_LoggingIn:
				continue
			case spb.PlayerStatus_PlayerStatus_PostLogin:
				p.PlayerRegisterMessage(msg.CmdId, msg)
			}
		}
	}
}

// kcp连接事件处理函数
func (s *GateServer) kcpEnetHandle(listener *kcp.Listener) {
	logger.Info("kcp enet handle start")
	for {
		enetNotify := <-listener.GetEnetNotifyChan()
		logger.Info("[Kcp Enet] addr: %v, conv: %v, sessionId: %v, connType: %v, enetType: %v",
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
			s.kcpEventChan <- &KcpEvent{
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

// 发送事件处理
func SendHandle(p *PlayerGame, kcpMsg *alg.PackMsg) {
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
		p.XorKey = createXorPad(p.Seed)
		logger.Info("uid:%v,seed:%v,密钥交换成功", p.Uid, p.Seed)
	}
}

func createXorPad(seed uint64) []byte {
	keyBlock := random.NewKeyBlock(seed, false)
	xorKey := keyBlock.XorKey()
	key := make([]byte, 4096)
	copy(key, xorKey[:])
	return key
}

func Close() error {
	GATESERVER.kcpFin = true
	for _, player := range GATESERVER.sessionMap {
		player.Status = spb.PlayerStatus_PlayerStatus_PassiveOffline
		KickPlayer(player)
	}
	close(GATESERVER.Stop)
	return nil
}

func KickPlayer(p *PlayerGame) {
	logger.Info("[UID:%v]离线目标GameServer:%v", p.Uid, p.GameAppId)
	GateToPlayer(p, cmd.PlayerKickOutScNotify, nil)
	p.KcpConn.Close()
	p.GameConn.Close()
	syncPl.Lock()
	delete(GATESERVER.sessionMap, p.Uid)
	syncPl.Unlock()
	CLIENT_CONN_NUM = int32(len(GATESERVER.sessionMap))
}

func KickwaitingLoginMap(p *PlayerGame) {
	logger.Info("[UID:%v]离线目标GameServer:%v", p.Uid, p.GameAppId)

	p.KcpConn.Close()
	p.GameConn.Close()
	syncPl.Lock()
	delete(GATESERVER.waitingLoginMap, p.Uid)
	syncPl.Unlock()
}

func (s *GateServer) AutoUpDataPlayer() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		for _, g := range GATESERVER.sessionMap {
			lastActiveTime := g.LastActiveTime
			timestamp := time.Now().Unix()
			if timestamp-lastActiveTime >= 60 {
				logger.Debug("玩家超时离线")
				g.Status = spb.PlayerStatus_PlayerStatus_Offline
				KickPlayer(g)
			}
		}
	}
}

func (s *GateServer) kcpNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	kcpErrorCount := uint64(0)
	for {
		<-ticker.C
		snmp := kcp.DefaultSnmp.Copy()
		kcpErrorCount += snmp.KCPInErrors
		logger.Info("kcp send: %v B/s, kcp recv: %v B/s", snmp.BytesSent/60, snmp.BytesReceived/60)
		logger.Info("udp send: %v B/s, udp recv: %v B/s", snmp.OutBytes/60, snmp.InBytes/60)
		logger.Info("udp send: %v pps, udp recv: %v pps", snmp.OutPkts/60, snmp.InPkts/60)
		clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
		logger.Info("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		kcp.DefaultSnmp.Reset()
	}
}
