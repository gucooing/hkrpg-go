package service

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	smd "github.com/gucooing/hkrpg-go/protocol/server"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

type HandlerFunc func(s *NodeDiscoveryService, payloadMsg pb.Message)

var handlerFuncRouteMap map[uint16]HandlerFunc

func init() {
	handlerFuncRouteMap = map[uint16]HandlerFunc{
		smd.PlayerLogoutRsp: playerLogoutRsp, // 下线回调
	}
}

func playerLogoutRsp(s *NodeDiscoveryService, payloadMsg pb.Message) {
	req := payloadMsg.(*spb.PlayerLogoutRsp)
	if req.Status == spb.LOGOUTSTATUS_OFFLINE_REPEAT_LOGIN {
		s.MessageQueue.SendToGate(req.GateAppId, &mq.NetMsg{
			MsgType: mq.PlayerLoginKill,
			Uid:     req.Uid,
		})
	}
}

func (s *NodeDiscoveryService) messageQueue() {
	for {
		netMsg := <-s.MessageQueue.GetNetMsg()
		switch netMsg.OriginServerType {
		case spb.ServerType_SERVICE_GATE:
			go s.gateMsgHandle(netMsg)
		default:
			logger.Error("unknow server type: %v", netMsg.OriginServerType)
		}
	}
}

func (g *NodeDiscoveryService) gateMsgHandle(netMsg *mq.NetMsg) {
	switch netMsg.MsgType {
	case mq.ServerMsg:
		handle, ok := handlerFuncRouteMap[netMsg.CmdId]
		if !ok {
			logger.Error("server msg error,cmd:%s",
				smd.GetSharedCmdProtoMap().GetCmdNameByCmdId(netMsg.CmdId))
			return
		}
		handle(g, netMsg.ServiceMsgPb)
	}
}
