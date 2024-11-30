package player

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func GetQuestDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetQuestDataScRsp)
	rsp.QuestList = make([]*proto.Quest, 0)
	for _, questInfo := range gdconf.GetQuestDataMap() {
		status := proto.QuestStatus_QUEST_DOING
		var progress uint32 = 0
		if g.GetPd().GetIsJumpMission() {
			status = proto.QuestStatus_QUEST_CLOSE
			progress = 1
		}
		quest := &proto.Quest{
			Progress:   progress,
			Status:     status,
			Id:         questInfo.QuestID,
			FinishTime: 1699688465,
		}
		rsp.QuestList = append(rsp.QuestList, quest)
	}

	g.Send(cmd.GetQuestDataScRsp, rsp)
}
