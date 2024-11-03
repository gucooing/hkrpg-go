package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewChangeStory() *spb.ChangeStory {
	return &spb.ChangeStory{
		CurChangeStory:  0,
		ChangeStoryInfo: make(map[uint32]*spb.ChangeStoryInfo),
		IsChangeStory:   false,
	}
}

func (g *PlayerData) GetChangeStory() *spb.ChangeStory {
	db := g.GetBasicBin()
	if db.ChangeStory == nil {
		db.ChangeStory = NewChangeStory()
	}
	if g.GetIsJumpMission() {
		db.ChangeStory.IsChangeStory = false
	}
	return db.ChangeStory
}

func (g *PlayerData) IsChangeStory() bool {
	db := g.GetChangeStory()
	return db.IsChangeStory
}

func (g *PlayerData) GetAllChangeStoryInfo() map[uint32]*spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	if db.ChangeStoryInfo == nil {
		db.ChangeStoryInfo = make(map[uint32]*spb.ChangeStoryInfo)
	}
	return db.ChangeStoryInfo
}

func (g *PlayerData) GetChangeStoryInfo(id uint32) *spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	if db.ChangeStoryInfo == nil {
		db.ChangeStoryInfo = make(map[uint32]*spb.ChangeStoryInfo)
	}
	return db.ChangeStoryInfo[id]
}

func (g *PlayerData) GetCurChangeStoryInfo() *spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	if !db.IsChangeStory {
		return nil
	}
	if conf := gdconf.GetStoryLine(db.CurChangeStory); conf == nil {
		db.IsChangeStory = false
		db.CurChangeStory = 0
		return nil
	}
	return g.GetChangeStoryInfo(db.CurChangeStory)
}

func (g *PlayerData) MissionAddChangeStoryLine(finishActionPara []uint32) (uint32, uint32, uint32, bool) {
	if len(finishActionPara) < 4 {
		return 0, 0, 0, false
	}
	db := g.GetChangeStory()
	storyLineId := finishActionPara[0]
	entryId := finishActionPara[1]
	anchorGroup := finishActionPara[2]
	anchorId := finishActionPara[3]
	if storyLineId == 0 {
		delete(db.ChangeStoryInfo, db.CurChangeStory)
		db.IsChangeStory = false
		db.CurChangeStory = 0
		return g.GetCurEntryId(), 0, 0, true
	}
	conf := gdconf.GetStoryLine(storyLineId)
	if conf != nil {
		if entryId == 0 {
			entryId = conf.InitEntranceID
			anchorGroup = conf.InitGroupID
			anchorId = conf.InitAnchorID
		}
		// 更新队伍
		g.NewStoryLine(storyLineId)
		db.CurChangeStory = storyLineId
		db.ChangeStoryInfo[storyLineId] = &spb.ChangeStoryInfo{
			ChangeStoryId: storyLineId,
			Scene: &spb.Scene{
				EntryId:  entryId,
				GroupId:  anchorGroup,
				AnchorId: anchorId,
				Pos:      nil,
				Rot:      nil,
			},
		}
		db.IsChangeStory = true
	} else {
		return 0, 0, 0, false
	}
	// g.StoryLineInfoScNotify() // 通知一次
	// g.SyncLineupNotify(g.GetCurLineUp())
	// g.ChangeStoryLineFinishScNotify()
	// 传送
	return entryId, anchorGroup, anchorId, true
}

func (g *PlayerData) GetDimensionId() uint32 {
	if db := g.GetCurChangeStoryInfo(); db != nil {
		conf := gdconf.GetStoryLineFloorData(db.ChangeStoryId)
		if conf != nil {
			return conf.DimensionID
		}
	}
	return 0
}

func (g *PlayerData) GameStoryLineId() uint32 {
	if db := g.GetCurChangeStoryInfo(); db != nil {
		return db.ChangeStoryId
	}
	return 0
}
