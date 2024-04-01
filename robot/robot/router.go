package robot

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
)

func (r *RoBot) RegisterMessage(cmdId uint16, payloadMsg []byte) {
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
