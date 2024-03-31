package node

import (
	"bufio"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) gameRecvHandle() {
	payload := make([]byte, PacketMaxLen)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GAME SERVICE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.killService()
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := bufio.NewReader(s.Conn).Read(payload)
		if err != nil {
			s.killService()
			break
		}
		bin = payload[:recvLen]
		msgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &msgList, nil)
		for _, msg := range msgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.gameRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) gameRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLogoutNotify: // 玩家下线成功通知
		s.gamePlayerLogoutNotify(serviceMsg)
	case cmd.GetAllServiceReq: // 获取目标服务所有
		s.GetAllServiceReq(serviceMsg)
	default:
		logger.Info("game -> node error cmdid:%v", cmdId)
	}
}

func (s *Service) gamePlayerLogoutNotify(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutNotify)
	s.PlayerNum--
	logger.Info("[UID:%v]node game离线成功", req.Uid)
}
