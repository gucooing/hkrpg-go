package Game

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetMissionStatusCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.GetMissionStatusCsReq, payloadMsg)
	req := msg.(*proto.GetMissionStatusCsReq)

	rsp := new(proto.GetMissionStatusScRsp)
	rsp.FinishedMainMissionIdList = []uint32{}
	rsp.SubMissionStatusList = make([]*proto.Mission, 0)
	for _, id := range req.MainMissionIdList {
		rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, id)
	}
	for _, id := range req.SubMissionIdList {
		rsp.SubMissionStatusList = append(rsp.SubMissionStatusList, &proto.Mission{
			Id:     id,
			Status: proto.MissionStatus_MISSION_FINISH,
		})
	}

	g.send(cmd.GetMissionStatusScRsp, rsp)
}

func (g *Game) GetQuestDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetQuestDataScRsp)
	rsp.QuestList = make([]*proto.Quest, 0)
	for _, questList := range gdconf.GetQuestDataMap() {
		quest := &proto.Quest{
			FinishTime: 10000,
			Progress:   questList.QuestType,
			Status:     proto.Quest_QUEST_CLOSE,
			Id:         questList.QuestID,
		}
		rsp.QuestList = append(rsp.QuestList, quest)
	}

	g.send(cmd.GetQuestDataScRsp, rsp)
}
