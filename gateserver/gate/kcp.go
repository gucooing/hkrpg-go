package gate

import (
	"encoding/binary"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

var QPS int64

func (s *GateServer) RunKcp() error {
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
			sessionId := kcpConn.GetSessionId()
			logger.Info("sessionId:%v", sessionId)
			// 读取密钥相关文件
			g := s.NewGame(kcpConn)
			go s.recvHandle(g)
		}()
	}
}

// kcp接收
func (s *GateServer) recvHandle(p *PlayerGame) {
	payload := make([]byte, PacketMaxLen)

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker player uid: %v", p.Uid)
			CLIENT_CONN_NUM--
			s.passPlayerKill(p, spb.Retcode_RET_PLAYER_SYSTEM_ERROR)
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
		kcpMsgList := make([]*alg.PackMsg, 0)
		bin = payload[:recvLen]
		alg.DecodeBinToPayload(bin, &kcpMsgList, p.XorKey)
		for _, msg := range kcpMsgList {
			// playerMsg := alg.DecodePayloadToProto(msg)
			switch p.Status {
			case spb.PlayerStatus_PlayerStatus_PreLogin:
				if msg.CmdId == cmd.PlayerGetTokenCsReq {
					p.Status = spb.PlayerStatus_PlayerStatus_LoggingIn
					s.PlayerGetTokenCsReq(p, msg.ProtoData)
				} else {
					return
				}
			case spb.PlayerStatus_PlayerStatus_LoggingIn:
				continue
			case spb.PlayerStatus_PlayerStatus_PostLogin:
				p.PlayerRegisterMessage(msg.CmdId, msg)
			default:
				return
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
		p.XorKey = random.CreateXorPad(p.Seed, false)
		logger.Info("uid:%v,seed:%v,密钥交换成功", p.Uid, p.Seed)
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
