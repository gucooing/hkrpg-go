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

func (g *GamePlayer) GetNpcMessageGroupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetNpcMessageGroupCsReq, payloadMsg)
	req := msg.(*proto.GetNpcMessageGroupCsReq)
	rsp := &proto.GetNpcMessageGroupScRsp{
		MessageGroupList: make([]*proto.MessageGroup, 0),
		Retcode:          0,
	}

	for _, contactId := range req.ContactIdList {
		db := g.GetMessageGroupByContactId(contactId)
		if db != nil {
			messageGroup := &proto.MessageGroup{
				Id:                 db.Id,
				RefreshTime:        db.RefreshTime,
				Status:             proto.MessageGroupStatus(db.Status),
				MessageSectionList: make([]*proto.MessageSection, 0),
				MessageSectionId:   0,
			}
			for _, msgSection := range db.MessageSectionList {
				messageGroup.MessageSectionList = append(messageGroup.MessageSectionList, &proto.MessageSection{
					Status:         proto.MessageSectionStatus(msgSection.Status),
					Id:             msgSection.Id,
					ToChooseItemId: make([]uint32, 0),
					MessageItemId:  0,
					ItemList:       make([]*proto.MessageItem, 0),
				})
			}
			rsp.MessageGroupList = append(rsp.MessageGroupList, messageGroup)
		}
	}

	g.Send(cmd.GetNpcMessageGroupScRsp, rsp)
}

func (g *GamePlayer) FinishPerformSectionIdCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.FinishPerformSectionIdCsReq, payloadMsg)
	req := msg.(*proto.FinishPerformSectionIdCsReq)

	g.FinishMessageGroup(req.SectionId)

	rsp := &proto.FinishPerformSectionIdScRsp{
		Reward:    &proto.ItemList{},
		Retcode:   0,
		SectionId: req.SectionId,
		ItemList:  make([]*proto.MessageItem, 0),
	}
	g.Send(cmd.FinishPerformSectionIdScRsp, rsp)
}

func (g *GamePlayer) MessageGroupPlayerSyncScNotify(contactId uint32) {
	db := g.GetMessageGroupByContactId(contactId)
	if db == nil {
		return
	}
	notify := &proto.PlayerSyncScNotify{
		MessageGroupStatus: make([]*proto.GroupStatus, 0),
		SectionStatus:      make([]*proto.SectionStatus, 0),
	}

	notify.MessageGroupStatus = append(notify.MessageGroupStatus, &proto.GroupStatus{
		RefreshTime: db.RefreshTime,
		GroupId:     db.Id,
		GroupStatus: proto.MessageGroupStatus(db.Status),
	})

	for _, msgSection := range db.MessageSectionList {
		notify.SectionStatus = append(notify.SectionStatus, &proto.SectionStatus{
			SectionId:     msgSection.Id,
			SectionStatus: proto.MessageSectionStatus(msgSection.Status),
		})
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}
