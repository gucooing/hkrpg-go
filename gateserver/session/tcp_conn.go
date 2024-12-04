package session

import (
	"fmt"
	"net"
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

var TcpSessionId uint32

type TcpListener struct {
	*Listener
	tcpListener net.Listener
}

func (t *TcpListener) initListener() error {
	addr := fmt.Sprintf("%s:%s", t.netInfo.InnerAddr, t.netInfo.InnerPort)
	logger.Info("tcp监听地址:%s", addr)
	logger.Info("tcp对外地址:%s", fmt.Sprintf("%s:%s", t.netInfo.OuterAddr, t.netInfo.OuterPort))
	tcpListener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen tcp err: %v\n", err)
	}
	t.tcpListener = tcpListener
	go tcpNetInfo()
	pushc.PushServer(&constant.LogPush{
		PushMessage: constant.PushMessage{
			Tag: "gateway",
		},
		LogMsg:   "网关模式为TCP",
		LogLevel: constant.INFO,
	})

	return nil
}

func (t *TcpListener) GetListener() *Listener {
	return t.Listener
}

func (t *TcpListener) Run() error {
	defer t.Close()
	for {
		tcpConn, err := t.tcpListener.Accept()
		if err != nil {
			return fmt.Errorf("accept tcp err: %v", err)
		}
		atomic.AddUint32(&TcpSessionId, 1)
		go func() {
			// new Session
			s := &TcpSession{
				tcpConn: tcpConn,
				Session: new(Session),
			}
			s.NewSession()
			go s.recvHandle()
			go s.sendHandle()
			t.LoginSessionChan <- s
		}()
	}
}

func (s *TcpSession) recvHandle() {
	defer s.Close()
	payload := make([]byte, alg.PacketMaxLen)
	bin := make([]byte, 0)
	for {
		recvLen, err := s.tcpConn.Read(payload)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		x := make([]byte, recvLen)
		copy(x, payload[:recvLen])
		bin = append(bin, x...)
		if len(bin) > 16 {
			kcpMsgList := make([]*alg.PackMsg, 0)
			err = alg.TcpDecodeBinToPayload(bin, &kcpMsgList, s.XorKey)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			for _, v := range kcpMsgList {
				bin = bin[v.Length:]
				s.sendServer(v)
			}
		}
	}
}

func (s *TcpSession) sendHandle() {
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
		binMsg := alg.TcpEncodePayloadToBin(packMsg, s.XorKey)
		if packMsg.CmdId == cmd.PlayerGetTokenScRsp &&
			s.XorKey != nil &&
			s.SessionState == SessionActivity {
			s.XorKey = random.CreateXorPad(s.Seed, false)
			logger.Info("uid:%v,seed:%v,密钥交换成功", s.Uid, s.Seed)
		}
		_, err = s.tcpConn.Write(binMsg)
		if err != nil {
			logger.Debug("exit send loop, conn write err: %v", err)
			return
		}
	}
}

// tcp统计
func tcpNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		clientConnNum := atomic.LoadInt64(&CLIENT_CONN_NUM)
		logger.Info("conn num: %v", clientConnNum)
		logger.Info("QPS: %v /s", QPS/60)
		QPS = 0
		kcp.DefaultSnmp.Reset()
	}
}

func (s *TcpSession) Close() {
	if s.SessionState == SessionClose {
		return
	}
	s.SessionState = SessionClose
	// 断开kcp
	s.tcpConn.Close()
	s.sendChan <- nil // 主动调用关闭
	s.recvChan <- nil // 主动调用关闭
	logger.Info("[UID:%v]玩家下线GATE", s.Uid)
}

func (t *TcpListener) Close() {

}
