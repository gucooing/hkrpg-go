package service

import (
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/session"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type packetFunc func(g *GateServer, s session.SessionAll, payloadMsg pb.Message)

var packetCaptureMap = map[uint16]packetFunc{ // 抽包
	cmd.PlayerLogoutCsReq:               PlayerLogoutCsReq,               // 下线请求
	cmd.SendMsgCsReq:                    SendMsgCsReq,                    // 发送聊天消息
	cmd.GetFriendRecommendListInfoCsReq: GetFriendRecommendListInfoCsReq, // 获取附近的人
}

func (g *GateServer) packetCapture(s session.SessionAll, packMsg *alg.PackMsg) {
	handelFunc, ok := packetCaptureMap[packMsg.CmdId]
	if !ok {
		netMsg := &mq.NetMsg{
			MsgType:        mq.GameServer,
			Uid:            s.GetSession().Uid,
			CmdId:          packMsg.CmdId,
			ServiceMsgByte: packMsg.ProtoData,
		}
		g.MessageQueue.SendToGame(s.GetSession().GameAppId, netMsg)
		return
	}
	payloadMsg := cmd.DecodePayloadToProto(packMsg)
	handelFunc(g, s, payloadMsg)
}

func PlayerLogoutCsReq(g *GateServer, s session.SessionAll, payloadMsg pb.Message) {
	g.MessageQueue.SendToGame(s.GetSession().GameAppId, &mq.NetMsg{ // 通知GS下线
		MsgType: mq.PlayerLogout,
		Uid:     s.GetSession().Uid,
	})
	g.DelSession(s)
}

func SendMsgCsReq(g *GateServer, sAll session.SessionAll, payloadMsg pb.Message) {
	s := sAll.GetSession()
	req := payloadMsg.(*proto.SendMsgCsReq)

	targetList := req.TargetList
	targetList = append(targetList, s.Uid)
	for _, targetUid := range targetList {
		target := g.GetSession(targetUid)
		if target == nil {
			continue
		}
		notify := &proto.RevcMsgScNotify{
			SourceUid:   s.Uid,
			MessageText: req.MessageText,
			ExtraId:     req.ExtraId,
			MessageType: req.MessageType,
			TargetUid:   targetUid,
			BNABNCCMILM: req.BNABNCCMILM,
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

	rsp := &proto.SendMsgScRsp{
		EndTime: uint64(time.Now().Unix()),
		Retcode: 0,
	}
	rspbin, err := pb.Marshal(rsp)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	s.SendChan <- &alg.PackMsg{
		CmdId:     cmd.SendMsgScRsp,
		HeadData:  nil,
		ProtoData: rspbin,
	}
}

func GetFriendRecommendListInfoCsReq(g *GateServer, sAll session.SessionAll, payloadMsg pb.Message) {
	rsp := &proto.GetFriendRecommendListInfoScRsp{
		PlayerInfoList: make([]*proto.FriendRecommendInfo, 0),
		Retcode:        0,
	}

	rspbin, err := pb.Marshal(rsp)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	s := sAll.GetSession()
	s.SendChan <- &alg.PackMsg{
		CmdId:     cmd.SendMsgScRsp,
		HeadData:  nil,
		ProtoData: rspbin,
	}
}
