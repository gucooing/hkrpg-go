/*
好用的kcp
爱 来自 hk4e-go
*/
package gate

import (
	"encoding/binary"
	"net"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	PacketMaxLen            = 343 * 1024 // 最大应用层包长度
	KcpConnAddrChangeNotify = "KcpConnAddrChangeNotify"
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数
var GAMESERVER *GateServer

type GateServer struct {
	AppId            string
	Port             string
	Config           *config.Config
	Store            *Store
	kcpListener      *kcp.Listener
	nodeConn         net.Conn
	kcpFin           bool
	sessionIdCounter uint32
	sessionMap       map[uint32]*PlayerGame
	kcpEventChan     chan *KcpEvent
	gameAppId        string                  // 最优appid
	gameAll          map[string]*serviceGame // 从node拉取的game列表
	errGameAppId     []string
}

type PlayerGame struct {
	GameAppId           string
	IsToken             bool // 是否通过token验证
	Uid                 uint32
	Seed                uint64
	XorKey              []byte // 密钥
	KcpConn             *kcp.UDPSession
	GameConn            net.Conn
	IsConnect           bool
	PlayerOfflineReason spb.PlayerOfflineReason
	LastActiveTime      int64 // 最近一次的活跃时间
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
}

func NewGate(cfg *config.Config) *GateServer {
	s := new(GateServer)
	// TODO
	GAMESERVER = s

	cfg.Ec2b = alg.GetEc2b()
	s.Config = cfg
	s.Store = NewStore(s.Config) // 初始化数据库连接
	s.sessionMap = make(map[uint32]*PlayerGame)
	s.AppId = alg.GetAppId()
	logger.Info("GateServer AppId:%s", s.AppId)
	port := s.Config.AppList[s.AppId].App["port_player"].Port
	if port == "" {
		panic("GateServer Port error")
	}
	s.Port = port

	addr := "0.0.0.0:" + s.Port
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		logger.Error("listen kcp err: %v", err)
		return nil
	}
	s.kcpListener = kcpListener
	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		logger.Error("node error:", err)
		return nil
	}
	s.nodeConn = tcpConn
	s.gameAll = make(map[string]*serviceGame)
	s.errGameAppId = make([]string, 0)
	go s.recvNode()
	go s.kcpNetInfo()
	go s.kcpEnetHandle(kcpListener)
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
		// 读取密钥相关文件
		g := s.NewGame(kcpConn)
		go s.recvHandle(g)
	}
	return nil
}

func (s *GateServer) NewGame(kcpConn *kcp.UDPSession) *PlayerGame {
	g := new(PlayerGame)
	g.KcpConn = kcpConn
	g.XorKey = config.GetConfig().Ec2b.XorKey()

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
			KickPlayer(p)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := p.KcpConn.Read(payload)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = payload[:recvLen]
		kcpMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &kcpMsgList, p.XorKey)
		for _, msg := range kcpMsgList {
			// playerMsg := alg.DecodePayloadToProto(msg)
			if p.IsToken {
				p.PlayerRegisterMessage(msg.CmdId, msg)
			} else {
				if msg.CmdId == cmd.PlayerGetTokenCsReq {
					s.HandlePlayerGetTokenCsReq(p, msg.ProtoData)
				} else {
					p.KcpConn.Close()
					return
				}
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
		if GAMESERVER.sessionMap[p.Uid] == nil {
			GAMESERVER.sessionMap[p.Uid] = p
			CLIENT_CONN_NUM = int32(len(GAMESERVER.sessionMap))
		}
		// 如果不为空则是断线重连
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
	GAMESERVER.kcpFin = true
	for _, player := range GAMESERVER.sessionMap {
		KickPlayer(player)
	}

	return nil
}

func KickPlayer(p *PlayerGame) {
	notify := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	GateToPlayer(p, cmd.PlayerKickOutScNotify, notify)
	p.KcpConn.Close()
	// 发送下线通知到game
	p.GameConn.Close()
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
