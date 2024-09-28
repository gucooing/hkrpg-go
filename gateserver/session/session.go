package session

import (
	"net"
	"sync/atomic"

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
	RecvChan             chan *alg.PackMsg // 网络接收通道
	SendChan             chan *alg.PackMsg // 网络发送通道
	KickFinishNotifyChan chan bool         // 等待离线通道
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
	s.RecvChan = make(chan *alg.PackMsg, 100)
	s.SendChan = make(chan *alg.PackMsg, 100)
	s.KickFinishNotifyChan = make(chan bool, 1)
}

func (s *TcpSession) NewSession() {
	s.XorKey = Ec2b.XorKey()
	s.SessionId = atomic.LoadUint32(&TcpSessionId)
	s.SessionState = SessionLogin
	s.RecvChan = make(chan *alg.PackMsg, 100)
	s.SendChan = make(chan *alg.PackMsg, 100)
	s.KickFinishNotifyChan = make(chan bool, 1)
}

func (s *KcpSession) GetSession() *Session {
	return s.Session
}

func (s *TcpSession) GetSession() *Session {
	return s.Session
}
