/*
好用的kcp
爱 来自 hk4e-go
*/
package Net

import (
	"encoding/binary"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/internal/Game"
	"github.com/gucooing/hkrpg-go/internal/SDK"
	"github.com/gucooing/hkrpg-go/pkg/config"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen            = 343 * 1024 // 最大应用层包长度
	KcpConnAddrChangeNotify = "KcpConnAddrChangeNotify"
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数

type KcpConnManager struct {
	kcpListener *kcp.Listener
	// 会话
	sessionIdCounter uint32
	// 事件
	kcpEventChan chan *KcpEvent
}

type KcpEvent struct {
	SessionId    uint32
	EventId      string
	EventMessage any
}

type RemoteKick struct {
	regFinishNotifyChan  chan bool
	userId               uint32
	kickFinishNotifyChan chan bool
}

func Run(s *SDK.Server) error {
	k := KcpConnManager{}

	addr := "0.0.0.0:" + strconv.Itoa(int(config.GetConfig().Game.Port))
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		logger.Error("listen kcp err: %v", err)
		return err
	}
	logger.Info("kcp服务在 %s 上启动", addr)
	k.kcpListener = kcpListener
	go k.kcpNetInfo()
	go k.kcpEnetHandle(kcpListener)

	for {
		kcpConn, err := kcpListener.AcceptKCP()
		if err != nil {
			logger.Error("accept kcp err: %v", err)
			continue
		}

		kcpConn.SetACKNoDelay(true)
		kcpConn.SetWriteDelay(false)
		kcpConn.SetWindowSize(256, 256)
		kcpConn.SetMtu(1200)
		g := NewGame(kcpConn)
		g.Db = s.Store
		CLIENT_CONN_NUM++
		go recvHandle(g)
		go sendNet(g)
	}
}

func NewGame(kcpConn *kcp.UDPSession) *Game.Game {
	g := new(Game.Game)
	g.KcpConn = kcpConn
	g.ServerCmdProtoMap = cmd.NewCmdProtoMap()
	g.NetMsgInput = make(chan *Game.NetMsg, 1000)
	return g
}

func recvHandle(g *Game.Game) {
	payload := make([]byte, PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := g.KcpConn.Read(payload)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = payload[:recvLen]
		kcpMsgList := make([]*KcpMsg, 0)
		DecodeBinToPayload(bin, &kcpMsgList)
		for _, v := range kcpMsgList {
			// name := g.ServerCmdProtoMap.GetCmdNameByCmdId(v.CmdId)
			// logger.Error("C --> S: %v", v.CmdId)
			// payloadMsg := DecodePayloadToProto(g, v) TODO 由于 req 大部分缺失，所以不预处理数据
			g.RegisterMessage(v.CmdId, v.ProtoData)
		}
	}
}

// kcp连接事件处理函数
func (k *KcpConnManager) kcpEnetHandle(listener *kcp.Listener) {
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
			sessionId := atomic.AddUint32(&k.sessionIdCounter, 1)
			listener.SendEnetNotifyToPeer(&kcp.Enet{
				Addr:      enetNotify.Addr,
				SessionId: sessionId,
				Conv:      binary.BigEndian.Uint32(random.GetRandomByte(4)),
				ConnType:  kcp.ConnEnetEst,
				EnetType:  enetNotify.EnetType,
			})
		case kcp.ConnEnetAddrChange:
			// 连接地址改变通知
			k.kcpEventChan <- &KcpEvent{
				SessionId:    enetNotify.SessionId,
				EventId:      KcpConnAddrChangeNotify,
				EventMessage: enetNotify.Addr,
			}
		default:
		}
	}
}

func sendNet(g *Game.Game) {
	for {
		netMsg := <-g.NetMsgInput
		SendHandle(netMsg.G, netMsg.CmdId, netMsg.PlayerMsg)
	}
}

// 发送事件处理
func SendHandle(g *Game.Game, cmdid uint16, playerMsg pb.Message) {
	rspMsg := new(ProtoMsg)
	rspMsg.CmdId = cmdid
	rspMsg.PayloadMessage = playerMsg
	kcpMsg := EncodeProtoToPayload(rspMsg)
	if kcpMsg.CmdId == 0 {
		logger.Error("cmdid error")
	}
	// logger.Debug("S --> C: %v", kcpMsg.CmdId)
	binMsg := EncodePayloadToBin(kcpMsg, nil)
	_, err := g.KcpConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

func (k *KcpConnManager) kcpNetInfo() {
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
