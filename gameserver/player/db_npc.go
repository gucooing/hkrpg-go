package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func NewMessageGroup() map[uint32]*spb.MessageGroup {
	return make(map[uint32]*spb.MessageGroup)
}

func (g *GamePlayer) GetMessageGroup() map[uint32]*spb.MessageGroup {
	db := g.GetBasicBin()
	if db.MessageGroupList == nil {
		db.MessageGroupList = NewMessageGroup()
	}
	return db.MessageGroupList
}

func (g *GamePlayer) GetMessageGroupByContactId(contactId uint32) *spb.MessageGroup {
	db := g.GetMessageGroup()
	return db[contactId]
}

func (g *GamePlayer) AddMessageGroup(sectionId uint32) {
	db := g.GetMessageGroup()
	confMg := gdconf.GetMessageGroupConfigBySectionID(sectionId)
	confMs := gdconf.GetMessageSectionConfig(sectionId)
	if confMg == nil || confMs == nil {
		return
	}
	contactId := confMg.MessageContactsID
	db[contactId] = &spb.MessageGroup{
		ContactId:          confMg.MessageContactsID,
		Id:                 confMg.ID,
		MessageSectionList: make(map[uint32]*spb.MessageSection),
		RefreshTime:        time.Now().Unix(),
		Status:             spb.MessageGroupStatus_MESSAGE_GROUP_DOING,
	}
	for _, confsectionId := range confMg.MessageSectionIDList {
		db[contactId].MessageSectionList[confsectionId] = &spb.MessageSection{
			Id:     confsectionId,
			Status: spb.MessageSectionStatus_MESSAGE_SECTION_DOING,
		}
	}

	g.MessageGroupPlayerSyncScNotify(contactId)
}

func (g *GamePlayer) FinishMessageGroup(sectionId uint32) {
	conf := gdconf.GetMessageGroupConfigBySectionID(sectionId)
	if conf == nil {
		return
	}
	db := g.GetMessageGroupByContactId(conf.MessageContactsID)
	if db == nil {
		return
	}
	if db.MessageSectionList[sectionId] != nil {
		db.MessageSectionList[sectionId].Status = spb.MessageSectionStatus_MESSAGE_SECTION_FINISH
	}
	isFinish := true
	for _, messageSection := range db.MessageSectionList {
		if messageSection.Status != spb.MessageSectionStatus_MESSAGE_SECTION_FINISH {
			isFinish = false
		}
	}
	if isFinish {
		db.Status = spb.MessageGroupStatus_MESSAGE_GROUP_FINISH
	}
	g.MessagePerformSectionFinish(sectionId)
	g.MessageGroupPlayerSyncScNotify(db.ContactId)
}
