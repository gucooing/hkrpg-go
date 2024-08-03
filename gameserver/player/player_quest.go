package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetQuestDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetQuestDataScRsp)
	rsp.QuestList = make([]*proto.Quest, 0)
	for _, questList := range gdconf.GetQuestDataMap() {
		quest := &proto.Quest{
			Progress:   0,
			Status:     proto.QuestStatus_QUEST_DOING,
			Id:         questList.QuestID,
			FinishTime: 1699688465,
		}
		rsp.QuestList = append(rsp.QuestList, quest)
	}

	g.Send(cmd.GetQuestDataScRsp, rsp)
}
