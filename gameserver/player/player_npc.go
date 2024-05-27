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
		NpcMeetStatusList: make([]*proto.NpcMeetStatusInfo, 0),
	}
	for _, getNpcList := range req.FirstTalkIdList {
		npcTalkInfo := &proto.NpcMeetStatusInfo{MeetId: getNpcList}
		rsp.NpcMeetStatusList = append(rsp.NpcMeetStatusList, npcTalkInfo)
	}
	g.Send(cmd.GetFirstTalkByPerformanceNpcScRsp, rsp)
}
