package model

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewMessageGroup() map[uint32]*spb.MessageGroup {
	return make(map[uint32]*spb.MessageGroup)
}

func (g *PlayerData) GetMessageGroup() map[uint32]*spb.MessageGroup {
	db := g.GetBasicBin()
	if db.MessageGroupList == nil {
		db.MessageGroupList = NewMessageGroup()
	}
	return db.MessageGroupList
}

func (g *PlayerData) GetMessageGroupByContactId(contactId uint32) *spb.MessageGroup {
	db := g.GetMessageGroup()
	return db[contactId]
}

func (g *PlayerData) GetMessageGroupBySectionId(sectionId uint32) *spb.MessageGroup {
	conf := gdconf.GetMessageGroupConfigBySectionID(sectionId)
	if conf == nil {
		return nil
	}
	return g.GetMessageGroupByContactId(conf.MessageContactsID)
}

func (g *PlayerData) GetMessageSection(sectionId uint32) *spb.MessageSection {
	conf := gdconf.GetMessageGroupConfigBySectionID(sectionId)
	if conf == nil {
		return nil
	}
	db := g.GetMessageGroupByContactId(conf.MessageContactsID)
	if db == nil {
		return nil
	}

	if db.MessageSectionList == nil {
		db.MessageSectionList = make(map[uint32]*spb.MessageSection)
	}
	if db.MessageSectionList[sectionId] == nil {
		db.MessageSectionList[sectionId] = &spb.MessageSection{
			Id:              sectionId,
			Status:          spb.MessageSectionStatus_MESSAGE_SECTION_DOING,
			ItemList:        make(map[uint32]bool),
			MessageItemList: make(map[uint32]bool),
		}
	}

	return db.MessageSectionList[sectionId]
}

func (g *PlayerData) AddMessageGroup(sectionId uint32) uint32 {
	db := g.GetMessageGroup()
	confMg := gdconf.GetMessageGroupConfigBySectionID(sectionId)
	confMs := gdconf.GetMessageSectionConfig(sectionId)
	if confMg == nil || confMs == nil {
		return 0
	}
	contactId := confMg.MessageContactsID
	info := &spb.MessageGroup{
		ContactId:          confMg.MessageContactsID,
		Id:                 confMg.ID,
		MessageSectionList: make(map[uint32]*spb.MessageSection),
		RefreshTime:        time.Now().Unix(),
		Status:             spb.MessageGroupStatus_MESSAGE_GROUP_DOING,
	}
	for _, confsectionId := range confMg.MessageSectionIDList {
		info.MessageSectionList[confsectionId] = &spb.MessageSection{
			Id:     confsectionId,
			Status: spb.MessageSectionStatus_MESSAGE_SECTION_DOING,
		}
	}
	db[contactId] = info
	return contactId
}

func (g *PlayerData) FinishMessageGroup(req *proto.FinishPerformSectionIdCsReq) uint32 {
	db := g.GetMessageSection(req.SectionId)
	group := g.GetMessageGroupBySectionId(req.SectionId)
	if group == nil || db == nil {
		return 0
	}
	if db.ItemList == nil {
		db.ItemList = make(map[uint32]bool)
	}
	for _, item := range req.ItemList {
		db.ItemList[item.ItemId] = true
	}
	db.Status = spb.MessageSectionStatus_MESSAGE_SECTION_FINISH

	isFinish := true
	for _, info := range group.MessageSectionList {
		if info.Status != spb.MessageSectionStatus_MESSAGE_SECTION_FINISH {
			isFinish = false
		}
	}
	if isFinish {
		group.Status = spb.MessageGroupStatus_MESSAGE_GROUP_FINISH
	}
	return group.ContactId
}
