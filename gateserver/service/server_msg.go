package service

import (
	"github.com/gucooing/hkrpg-go/pkg/mq"
	smd "github.com/gucooing/hkrpg-go/protocol/server"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

type HandlerFunc func(g *GateServer, payloadMsg pb.Message)

var handlerFuncRouteMap map[uint16]HandlerFunc

func init() {
	handlerFuncRouteMap = map[uint16]HandlerFunc{
		smd.PlayerLogoutReq: playerLogoutReq,
	}
}

func playerLogoutReq(g *GateServer, payloadMsg pb.Message) {
	req := payloadMsg.(*spb.PlayerLogoutReq)
	rsp := &spb.PlayerLogoutRsp{
		Uid:       req.Uid,
		GateAppId: req.GateAppId,
		Status:    req.Status,
	}
	s := g.GetSession(req.Uid)
	if s != nil {
		PlayerLogoutCsReq(g, s, nil)
	}
	g.MessageQueue.SendToNode(&mq.NetMsg{
		MsgType:      mq.ServerMsg,
		Uid:          req.Uid,
		CmdId:        smd.PlayerLogoutRsp,
		ServiceMsgPb: rsp,
	})
}
