package session

import (
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

type SessionState int

var Ec2b *random.Ec2b

const (
	SessionLogin    SessionState = 0 // 尚未登录
	SessionActivity SessionState = 1 // 正常活动
	SessionFreeze   SessionState = 2 // 冻结连接
	SessionClose    SessionState = 3 // 关闭连接
)

type SessionAll interface {
	NewSession()
	GetSession() *Session
	recvHandle() // 监听收包
	sendHandle() // 监听发包
	Close()
}

type Session struct {
	SessionId            uint32
	Uid                  uint32            // 玩家游戏id
	Seed                 uint64            // 随机种子
	XorKey               []byte            // 密钥
	SessionState         SessionState      // 状态
	GameAppId            uint32            // 所属gs服务器(仅分布式
	recvChan             chan *alg.PackMsg // 网络接收通道
	recvSync             sync.Mutex        // 发送锁
	isCloseRecv          bool              // 是否关闭发送通道
	sendChan             chan *alg.PackMsg // 网络发送通道
	sendSync             sync.Mutex        // 发送锁
	isCloseSend          bool              // 是否关闭发送通道
	KickFinishNotifyChan chan bool         // 等待离线通道
	LastActiveTime       int64             // 活跃时间
}

type KcpSession struct {
	*Session
	kcpConn *kcp.UDPSession
}

type TcpSession struct {
	*Session
	tcpConn net.Conn
}

func (s *KcpSession) NewSession() {
	s.XorKey = Ec2b.XorKey()
	s.SessionId = s.kcpConn.GetSessionId()
	s.SessionState = SessionLogin
	s.recvChan = make(chan *alg.PackMsg, 100)
	s.sendChan = make(chan *alg.PackMsg, 100)
	s.KickFinishNotifyChan = make(chan bool, 1)
	s.LastActiveTime = time.Now().Unix()
}

func (s *TcpSession) NewSession() {
	s.XorKey = Ec2b.XorKey()
	s.SessionId = atomic.LoadUint32(&TcpSessionId)
	s.SessionState = SessionLogin
	s.recvChan = make(chan *alg.PackMsg, 100)
	s.sendChan = make(chan *alg.PackMsg, 100)
	s.KickFinishNotifyChan = make(chan bool, 1)
	s.LastActiveTime = time.Now().Unix()
}

func (s *KcpSession) GetSession() *Session {
	return s.Session
}

func (s *TcpSession) GetSession() *Session {
	return s.Session
}

// 将消息发送到客户端
func (s *Session) SendClient(msg *alg.PackMsg) {
	s.sendSync.Lock()
	defer s.sendSync.Unlock()
	if s.SessionState == SessionClose {
		if !s.isCloseSend {
			close(s.sendChan)
			s.isCloseSend = true
		}
		return
	}
	s.sendChan <- msg
}

// 接收发送到客户端的消息
func (s *Session) recvClient() (*alg.PackMsg, error) {
	if s.SessionState == SessionClose || s.isCloseSend {
		return nil, io.EOF
	}
	msg, ok := <-s.sendChan
	if !ok {
		return nil, io.EOF
	}
	return msg, nil
}

// 发送消息到服务端
func (s *Session) sendServer(msg *alg.PackMsg) {
	s.recvSync.Lock()
	defer s.recvSync.Unlock()
	if s.SessionState == SessionClose {
		if !s.isCloseRecv {
			close(s.recvChan)
			s.isCloseRecv = true
		}
		return
	}
	s.LastActiveTime = time.Now().Unix()
	s.recvChan <- msg
}

// 接收发送到服务端的消息
func (s *Session) RecvServer() (*alg.PackMsg, error) {
	if s.SessionState == SessionClose ||
		s.isCloseRecv {
		return nil, io.EOF
	}
	msg, ok := <-s.recvChan
	if !ok {
		return nil, io.EOF
	}
	return msg, nil
}
