package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) GetFirstTalkNpcCsReq(payloadMsg []byte) {
	g.Send(cmd.GetFirstTalkNpcScRsp, nil)
}

func (g *GamePlayer) GetNpcTakenRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetNpcTakenRewardCsReq, payloadMsg)
	req := msg.(*proto.GetNpcTakenRewardCsReq)
	rsp := new(proto.GetNpcTakenRewardScRsp)
	rsp.NpcId = req.NpcId

	g.Send(cmd.GetNpcTakenRewardScRsp, rsp)
}

func (g *GamePlayer) GetFirstTalkByPerformanceNpcCsReq(payloadMsg []byte) {
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
