package node

import (
	"bufio"
	"time"

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
			s.n.killService(s)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := bufio.NewReader(s.Conn).Read(payload)
		if err != nil {
			s.n.killService(s)
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
	case cmd.GameToNodePingReq: // 获取目标服务所有
		s.GameToNodePingReq(serviceMsg)
	default:
		logger.Info("game -> node error cmdid:%v", cmdId)
	}
}

func (s *Service) GameToNodePingReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GameToNodePingReq)
	if req.GameServerId != s.AppId {
		return
	}
	s.PlayerNum = req.PlayerNum
	rsp := &spb.GameToNodePingRsp{
		GameServerTime: req.GameServerTime,
		NodeTime:       time.Now().UnixNano() / 1e6,
	}
	s.sendHandle(cmd.GameToNodePingRsp, rsp)
}
