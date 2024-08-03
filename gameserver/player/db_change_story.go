package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func NewChangeStory() *spb.ChangeStory {
	return &spb.ChangeStory{
		CurChangeStory:  0,
		ChangeStoryInfo: make(map[uint32]*spb.ChangeStoryInfo),
		IsChangeStory:   false,
	}
}

func (g *GamePlayer) GetChangeStory() *spb.ChangeStory {
	db := g.GetBasicBin()
	if db.ChangeStory == nil {
		db.ChangeStory = NewChangeStory()
	}
	return db.ChangeStory
}

func (g *GamePlayer) IsChangeStory() bool {
	db := g.GetChangeStory()
	return db.IsChangeStory
}

func (g *GamePlayer) GetAllChangeStoryInfo() map[uint32]*spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	if db.ChangeStoryInfo == nil {
		db.ChangeStoryInfo = make(map[uint32]*spb.ChangeStoryInfo)
	}
	return db.ChangeStoryInfo
}

func (g *GamePlayer) GetChangeStoryInfo(id uint32) *spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	if db.ChangeStoryInfo == nil {
		db.ChangeStoryInfo = make(map[uint32]*spb.ChangeStoryInfo)
	}
	return db.ChangeStoryInfo[id]
}

func (g *GamePlayer) GetCurChangeStoryInfo() *spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	if !db.IsChangeStory {
		return nil
	}
	return g.GetChangeStoryInfo(db.CurChangeStory)
}

func (g *GamePlayer) MissionAddChangeStoryLine(finishActionPara []uint32) {
	if len(finishActionPara) < 4 {
		return
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
		return
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
	}
	g.StoryLineInfoScNotify() // 通知一次
	g.SyncLineupNotify(g.GetCurLineUp())
	g.ChangeStoryLineFinishScNotify()
	// 传送
	g.EnterSceneByServerScNotify(entryId, 0, anchorGroup, anchorId)
}

func (g *GamePlayer) GetDimensionId() uint32 {
	if db := g.GetCurChangeStoryInfo(); db != nil {
		return 1
	}
	return 0
}

func (g *GamePlayer) GameStoryLineId() uint32 {
	if db := g.GetCurChangeStoryInfo(); db != nil {
		return db.ChangeStoryId
	}
	return 0
}
