package player

import (
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

func (g *GamePlayer) GetChangeStoryInfo(id uint32) *spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	return db.ChangeStoryInfo[id]
}

func (g *GamePlayer) GetCurChangeStoryInfo() *spb.ChangeStoryInfo {
	db := g.GetChangeStory()
	if db.IsChangeStory {
		return nil
	}
	return g.GetChangeStoryInfo(db.CurChangeStory)
}

func (g *GamePlayer) GetCurChangeStoryLineup() *spb.Line {
	db := g.GetChangeStory()
	return g.GetStoryLineById(db.CurChangeStory)
}
