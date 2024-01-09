package internal

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (r *RoBot) RegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLoginScRsp:
		r.PlayerLoginScRsp()
		r.PlayerHeartBeatCsReq()
	case cmd.PlayerHeartBeatScRsp:
		r.PlayerHeartbeatScRsp(payloadMsg)
	case cmd.GetCurSceneInfoScRsp:
		r.GetCurSceneInfoScRsp(payloadMsg)
	case cmd.EnterSceneByServerScNotify:
		r.EnterSceneByServerScNotify(payloadMsg)
	default:

	}
}
