package session

import (
	"sync/atomic"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type SessionState int

const (
	SessionLogin    SessionState = 0 // 尚未登录
	SessionActivity SessionState = 1 // 正常活动
	SessionFreeze   SessionState = 2 // 冻结连接
	SessionClose    SessionState = 3 // 关闭连接
)

type Session struct {
	kcpConn      *kcp.UDPSession
	SessionId    uint32
	Uid          uint32            // 玩家游戏id
	Seed         uint64            // 随机种子
	XorKey       []byte            // 密钥
	SessionState SessionState      // 状态
	GameAppId    uint32            // 所属gs服务器(仅分布式
	RecvChan     chan *alg.PackMsg // kcp接收通道
	SendChan     chan *alg.PackMsg // kcp发送通道
}

func NewSession(kcpConn *kcp.UDPSession, xorKey []byte) *Session {
	s := new(Session)
	s.XorKey = xorKey
	s.kcpConn = kcpConn
	s.SessionId = kcpConn.GetSessionId()
	s.SessionState = SessionLogin
	s.RecvChan = make(chan *alg.PackMsg)
	s.SendChan = make(chan *alg.PackMsg)

	return s
}

// 监听收包
func (s *Session) recvHandle() {
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
			s.RecvChan <- v
		}
	}
}

// 监听发包
func (s *Session) sendHandle() {
	for {
		packMsg, ok := <-s.SendChan
		if !ok {
			logger.Debug("exit send loop, send chan close, sessionId: %v", s.Uid)
			// TODO KILL
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
		_, err := s.kcpConn.Write(binMsg)
		if err != nil {
			logger.Debug("exit send loop, conn write err: %v", err)
			// TODO KILL
			return
		}
	}
}

func (s *Session) Close() {
	if s.SessionState == SessionClose {
		return
	}
	s.SessionState = SessionClose
	// 通知客户端下线
	protoData, err := pb.Marshal(&proto.PlayerKickOutScNotify{
		BlackInfo: &proto.BlackInfo{},
	})
	if err != nil {
		binMsg := alg.EncodePayloadToBin(&alg.PackMsg{
			CmdId:     cmd.PlayerKickOutScNotify,
			HeadData:  nil,
			ProtoData: protoData,
		}, s.XorKey)
		s.kcpConn.Write(binMsg)
	}

	// 断开kcp
	s.kcpConn.Close()
	// 断开通道
	close(s.RecvChan)
	close(s.SendChan)

	atomic.AddInt32(&CLIENT_CONN_NUM, -1)
}
