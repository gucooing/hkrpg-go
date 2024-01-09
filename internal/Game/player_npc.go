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
	rsp := new(proto.GetNpcTakenRewardScRsp)
	rsp.NpcId = req.NpcId

	g.Send(cmd.GetNpcTakenRewardScRsp, rsp)
}

func (g *Game) GetFirstTalkByPerformanceNpcCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetFirstTalkByPerformanceNpcCsReq, payloadMsg)
	req := msg.(*proto.GetFirstTalkByPerformanceNpcCsReq)
	rsp := &proto.GetFirstTalkByPerformanceNpcScRsp{
		NpcTalkInfoList: make([]*proto.NpcTalkInfo, 0),
	}
	for _, getNpcList := range req.NpcTalkList {
		npcTalkInfo := &proto.NpcTalkInfo{NpcTalkId: getNpcList}
		rsp.NpcTalkInfoList = append(rsp.NpcTalkInfoList, npcTalkInfo)
	}
	g.Send(cmd.GetFirstTalkByPerformanceNpcScRsp, rsp)
}
