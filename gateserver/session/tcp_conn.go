package session

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/endec"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
)

var TcpSessionId uint32

type TcpListener struct {
	*Listener
	tcpListener net.Listener
}

func (t *TcpListener) Null() {}

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
		if s.SessionState == SessionClose {
			return
		}
		x := make([]byte, recvLen)
		copy(x, payload[:recvLen])
		bin = append(bin, x...)
		if len(bin) > 16 {
			// 头部幻数错误
			if binary.BigEndian.Uint32(bin[:4]) != 0x9d74c714 {
				logger.Error("packet head magic 0x9d74c714 error")
				return
			}
			// 协议号
			cmdId := binary.BigEndian.Uint16(bin[4:6])
			// 头部长度
			headLen := binary.BigEndian.Uint16(bin[6:8])
			// proto长度
			protoLen := binary.BigEndian.Uint32(bin[8:12])
			// 检查长度
			packetLen := int(headLen) + int(protoLen) + 16
			if packetLen > alg.PacketMaxLen {
				logger.Error("packet len too long")
				return
			}
			if len(bin) < packetLen {
				continue
			}
			// 尾部幻数错误
			if binary.BigEndian.Uint32(bin[len(bin)-4:]) != 0xd7a152c8 {
				logger.Error("packet tail magic 0xd7a152c8 error")
				return
			}
			data := bin[12 : 12+int(headLen)+int(protoLen)]
			if s.XorKey != nil {
				endec.Xor(data, s.XorKey)
			}
			// 头部数据
			headData := data[int(headLen):]
			// proto数据
			protoData := data[int(headLen) : int(headLen)+int(protoLen)]
			// 返回数据
			kcpMsg := new(alg.PackMsg)
			kcpMsg.CmdId = cmdId
			kcpMsg.HeadData = make([]byte, headLen)
			kcpMsg.ProtoData = make([]byte, protoLen)
			copy(kcpMsg.HeadData, headData)
			copy(kcpMsg.ProtoData, protoData)
			bin = bin[16+int(headLen)+int(protoLen):]
			QPS++
			s.RecvChan <- kcpMsg
		}
	}
}

func (s *TcpSession) sendHandle() {
	defer s.Close()
	for {
		packMsg, ok := <-s.SendChan
		if !ok {
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
		_, err := s.tcpConn.Write(binMsg)
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
		logger.Debug("conn num: %v", clientConnNum)
		logger.Debug("QPS: %v /s", QPS/60)
		QPS = 0
		kcp.DefaultSnmp.Reset()
	}
}

func (s *TcpSession) Close() {
	if s.SessionState == SessionClose {
		return
	}
	s.SessionState = SessionClose
	// 等待所有待发送的消息发送完毕
	for {
		if len(s.SendChan) == 0 {
			time.Sleep(time.Millisecond * 100)
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	// 断开kcp
	s.tcpConn.Close()
	// 断开通道
	close(s.RecvChan)
	close(s.SendChan)

	logger.Info("[UID:%v]玩家下线GATE", s.Uid)
	atomic.AddInt64(&CLIENT_CONN_NUM, -1)
}

func (t *TcpListener) Close() {

}
