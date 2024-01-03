package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) GetFirstTalkNpcCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.GetFirstTalkNpcScRsp, rsp)
}

func (g *Game) GetNpcTakenRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetNpcTakenRewardCsReq, payloadMsg)
	req := msg.(*proto.GetNpcTakenRewardCsReq)

	rsp := &proto.GetNpcTakenRewardScRsp{NpcId: req.NpcId}
	g.Send(cmd.GetNpcTakenRewardScRsp, rsp)
}
