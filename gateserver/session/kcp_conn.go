package session

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/suppl/pushc"
)

type KcpListener struct {
	*Listener
	kcpListener *kcp.Listener
}

func (k *KcpListener) initListener() error {
	addr := fmt.Sprintf("%s:%s", k.netInfo.InnerAddr, k.netInfo.InnerPort)
	logger.Info("kcp监听地址:%s", addr)
	logger.Info("kcp对外地址:%s", fmt.Sprintf("%s:%s", k.netInfo.OuterAddr, k.netInfo.OuterPort))
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		return fmt.Errorf("listen kcp err: %v\n", err)
	}
	k.kcpListener = kcpListener
	k.kcpListener.EnetHandle()
	go kcpNetInfo()
	pushc.PushServer(&constant.LogPush{
		PushMessage: constant.PushMessage{
			Tag: "gateway",
		},
		LogMsg:   "网关模式为KCP",
		LogLevel: constant.INFO,
	})
	return nil
}

func (k *KcpListener) GetListener() *Listener {
	return k.Listener
}

func (k *KcpListener) Run() error {
	defer k.Close()
	for {
		kcpConn, err := k.kcpListener.AcceptKCP()
		if err != nil {
			return fmt.Errorf("accept kcp err: %v", err)
		}
		go func() {
			kcpConn.SetACKNoDelay(true)
			kcpConn.SetWriteDelay(false)
			kcpConn.SetWindowSize(256, 256)
			kcpConn.SetMtu(1500)
			// new Session
			s := &KcpSession{
				kcpConn: kcpConn,
				Session: new(Session),
			}
			s.NewSession()
			go s.recvHandle()
			go s.sendHandle()
			k.LoginSessionChan <- s
		}()
	}
}

func (s *KcpSession) recvHandle() {
	defer s.Close()
	payload := make([]byte, alg.PacketMaxLen)
	for {
		recvLen, err := s.kcpConn.Read(payload)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		QPS++
		bin := payload[:recvLen]
		kcpMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &kcpMsgList, s.XorKey)
		for _, v := range kcpMsgList {
			s.sendServer(v)
		}
	}
}

func (s *KcpSession) sendHandle() {
	defer s.Close()
	for {
		packMsg, err := s.recvClient()
		if err != nil {
			logger.Debug("exit send loop, send chan close, sessionId: %v", s.SessionId)
			return
		}
		if s.SessionState == SessionFreeze {
			continue
		}
		binMsg := alg.EncodePayloadToBin(packMsg, s.XorKey)
		if packMsg.CmdId == cmd.PlayerGetTokenScRsp &&
			s.XorKey != nil &&
			s.SessionState == SessionActivity {
			s.XorKey = random.CreateXorPad(s.Seed, false)
			logger.Info("uid:%v,seed:%v,密钥交换成功", s.Uid, s.Seed)
		}
		_, err = s.kcpConn.Write(binMsg)
		if err != nil {
			logger.Debug("exit send loop, conn write err: %v", err)
			return
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
		logger.Info("kcp send: %v B/s, kcp recv: %v B/s", snmp.BytesSent/60, snmp.BytesReceived/60)
		logger.Info("udp send: %v B/s, udp recv: %v B/s", snmp.OutBytes/60, snmp.InBytes/60)
		logger.Info("udp send: %v pps, udp recv: %v pps", snmp.OutPkts/60, snmp.InPkts/60)
		clientConnNum := atomic.LoadInt64(&CLIENT_CONN_NUM)
		logger.Info("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		logger.Info("QPS: %v /s", QPS/60)
		QPS = 0
		kcp.DefaultSnmp.Reset()
	}
}

func (s *KcpSession) Close() {
	if s.SessionState == SessionClose {
		return
	}
	s.SessionState = SessionClose
	// 断开kcp
	s.kcpConn.Close()
	s.SendClient(nil) // 主动调用关闭
	s.sendServer(nil) // 主动调用关闭
	logger.Info("[UID:%v]玩家下线GATE", s.Uid)
}

func (k *KcpListener) Close() {

}
