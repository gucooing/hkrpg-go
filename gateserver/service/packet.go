package service

import (
	"github.com/gucooing/hkrpg-go/gateserver/session"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type packetFunc func(g *GateServer, s *session.Session, payloadMsg pb.Message)

var packetCaptureMap = map[uint16]packetFunc{ // 抽包
	cmd.PlayerLogoutCsReq: PlayerLogoutCsReq, // 下线请求
	cmd.SendMsgCsReq:      SendMsgCsReq,      // 发送聊天消息
}

func (g *GateServer) packetCapture(s *session.Session, packMsg *alg.PackMsg) {
	handelFunc, ok := packetCaptureMap[packMsg.CmdId]
	if !ok {
		netMsg := &mq.NetMsg{
			MsgType:        mq.GameServer,
			Uid:            s.Uid,
			CmdId:          packMsg.CmdId,
			ServiceMsgByte: packMsg.ProtoData,
		}
		g.MessageQueue.SendToGame(s.GameAppId, netMsg)
		return
	}
	payloadMsg := cmd.DecodePayloadToProto(packMsg)
	handelFunc(g, s, payloadMsg)
}

func PlayerLogoutCsReq(g *GateServer, s *session.Session, payloadMsg pb.Message) {
	g.MessageQueue.SendToGame(s.GameAppId, &mq.NetMsg{ // 通知GS下线
		MsgType: mq.PlayerLogout,
		Uid:     s.Uid,
	})
	g.DelSession(s)
}

func SendMsgCsReq(g *GateServer, s *session.Session, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SendMsgCsReq)
	notify := &proto.RevcMsgScNotify{
		SourceUid:   s.Uid,
		MessageText: req.MessageText,
		ExtraId:     req.ExtraId,
		MessageType: req.MessageType,
		TargetUid:   req.TargetList[0],
		IGNEAJDPAPE: req.IGNEAJDPAPE,
		ChatType:    req.ChatType,
	}
	protoData, err := pb.Marshal(notify)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	s.SendChan <- &alg.PackMsg{
		CmdId:     cmd.RevcMsgScNotify,
		HeadData:  nil,
		ProtoData: protoData,
	}
}
