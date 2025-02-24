package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func GetFirstTalkNpcCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetFirstTalkNpcCsReq)
	rsp := &proto.GetFirstTalkNpcScRsp{
		Retcode:           0,
		NpcMeetStatusList: make([]*proto.FirstNpcTalkInfo, 0),
	}
	for _, seriesId := range req.NpcIdList {
		rsp.NpcMeetStatusList = append(rsp.NpcMeetStatusList, &proto.FirstNpcTalkInfo{
			IsMeet: false,
			NpcId:  seriesId,
		})
	}
	g.Send(cmd.GetFirstTalkNpcScRsp, rsp)
}

func GetNpcTakenRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetNpcTakenRewardCsReq)
	rsp := new(proto.GetNpcTakenRewardScRsp)
	rsp.NpcId = req.NpcId

	g.Send(cmd.GetNpcTakenRewardScRsp, rsp)
}

func GetFirstTalkByPerformanceNpcCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetFirstTalkByPerformanceNpcCsReq)
	rsp := &proto.GetFirstTalkByPerformanceNpcScRsp{
		NpcMeetStatusList: make([]*proto.NpcMeetByPerformanceStatus, 0),
	}
	for _, getNpcList := range req.PerformanceIdList {
		npcTalkInfo := &proto.NpcMeetByPerformanceStatus{PerformanceId: getNpcList}
		rsp.NpcMeetStatusList = append(rsp.NpcMeetStatusList, npcTalkInfo)
	}
	g.Send(cmd.GetFirstTalkByPerformanceNpcScRsp, rsp)
}

func GetNpcMessageGroupCsReq(g *GamePlayer, payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.GetNpcMessageGroupCsReq)
	rsp := &proto.GetNpcMessageGroupScRsp{
		MessageGroupList: make([]*proto.MessageGroup, 0),
		Retcode:          0,
	}

	for _, db := range g.GetPd().GetMessageGroup() {
		messageGroup := &proto.MessageGroup{
			Id:                 db.Id,
			RefreshTime:        db.RefreshTime,
			Status:             proto.MessageGroupStatus(db.Status),
			MessageSectionList: make([]*proto.MessageSection, 0),
			MessageSectionId:   db.CurMessageSectionId,
		}
		for _, msgSection := range db.MessageSectionList {
			info := &proto.MessageSection{
				Status:          proto.MessageSectionStatus(msgSection.Status),
				Id:              msgSection.Id,
				MessageItemList: make([]uint32, 0),
				FrozenItemId:    0,
				ItemList:        make([]*proto.MessageItem, 0),
			}
			for itemId, _ := range msgSection.MessageItemList {
				info.ItemList = append(info.ItemList, &proto.MessageItem{
					TextId: 0,
					ItemId: itemId,
				})
			}
			messageGroup.MessageSectionList = append(messageGroup.MessageSectionList, info)
		}
		rsp.MessageGroupList = append(rsp.MessageGroupList, messageGroup)
	}

	g.Send(cmd.GetNpcMessageGroupScRsp, rsp)
}

func FinishPerformSectionIdCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishPerformSectionIdCsReq)

	addItem := model.NewAddItem(nil)

	contactId := g.GetPd().FinishMessageGroup(req)

	if contactId != 0 {
		addItem.AllSync.MessageGroupList = append(addItem.AllSync.MessageGroupList, contactId)
		addItem.AllSync.MessageSectionList = append(addItem.AllSync.MessageSectionList, req.SectionId)
	}

	// 任务检查
	finishSubMission := g.GetPd().MessagePerformSectionFinish(req.SectionId)
	if len(finishSubMission) > 0 {
			g.InspectMission(finishSubMission...)
		}

	rsp := &proto.FinishPerformSectionIdScRsp{
		Reward:    &proto.ItemList{},
		Retcode:   0,
		SectionId: req.SectionId,
		ItemList:  req.ItemList,
	}

	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.Send(cmd.FinishPerformSectionIdScRsp, rsp)
}

func (g *GamePlayer) MessageGroupPlayerSyncScNotify(contactId uint32) {
	db := g.GetPd().GetMessageGroupByContactId(contactId)
	if db == nil {
		return
	}
	syncStatus := &proto.SyncStatus{
		MessageGroupStatus: make([]*proto.GroupStatus, 0),
		SectionStatus:      make([]*proto.SectionStatus, 0),
	}
	notify := &proto.PlayerSyncScNotify{
		SyncStatus: syncStatus,
	}

	syncStatus.MessageGroupStatus = append(syncStatus.MessageGroupStatus, &proto.GroupStatus{
		RefreshTime: db.RefreshTime,
		GroupId:     db.Id,
		GroupStatus: proto.MessageGroupStatus(db.Status),
	})

	for _, msgSection := range db.MessageSectionList {
		syncStatus.SectionStatus = append(syncStatus.SectionStatus, &proto.SectionStatus{
			SectionId:     msgSection.Id,
			SectionStatus: proto.MessageSectionStatus(msgSection.Status),
		})
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func GetNpcStatusCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetNpcStatusScRsp{
		NpcStatusList: make([]*proto.NpcStatus, 0),
		Retcode:       0,
	}
	db := g.GetPd().GetMessageGroup()
	if db != nil {
		for _, info := range db {
			isFinish := false
			if info.Status == spb.MessageGroupStatus_MESSAGE_GROUP_FINISH {
				isFinish = true
			}
			rsp.NpcStatusList = append(rsp.NpcStatusList, &proto.NpcStatus{
				IsFinish: isFinish,
				NpcId:    info.ContactId,
			})
		}
	}

	g.Send(cmd.GetNpcStatusScRsp, rsp)
}

func FinishItemIdCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishItemIdCsReq)
	rsp := &proto.FinishItemIdScRsp{
		TextId:  req.TextId,
		ItemId:  req.ItemId,
		Retcode: 0,
	}
	g.Send(cmd.FinishItemIdScRsp, rsp)
}

func FinishSectionIdCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishSectionIdCsReq)
	rsp := &proto.FinishSectionIdScRsp{
		Retcode:   0,
		SectionId: req.SectionId,
		Reward:    nil,
	}
	g.Send(cmd.FinishSectionIdScRsp, rsp)
}

func FinishFirstTalkByPerformanceNpcCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishFirstTalkByPerformanceNpcCsReq)
	rsp := &proto.FinishFirstTalkByPerformanceNpcScRsp{
		PerformanceId: req.PerformanceId,
		Reward:        nil,
		Retcode:       0,
	}
	g.Send(cmd.FinishFirstTalkByPerformanceNpcScRsp, rsp)
}
