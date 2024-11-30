package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewHandbook() *spb.Handbook {
	return &spb.Handbook{}
}

func (g *PlayerData) GetHandbook() *spb.Handbook {
	db := g.GetBasicBin()
	if db.Handbook == nil {
		db.Handbook = NewHandbook()
	}
	return db.Handbook
}

func (g *PlayerData) GetRogueHandbook() *spb.RogueHandbook {
	db := g.GetHandbook()
	if db.RogueHandbook == nil {
		db.RogueHandbook = &spb.RogueHandbook{}
	}
	return db.RogueHandbook
}

func (g *PlayerData) GetRogueHandbookMiracleMap() map[uint32]*spb.RogueHandbookMiracleInfo {
	db := g.GetRogueHandbook()
	if db.MiracleInfo == nil {
		db.MiracleInfo = make(map[uint32]*spb.RogueHandbookMiracleInfo)
	}
	return db.MiracleInfo
}

func (g *PlayerData) GetRogueHandbookMiracle(miracleId uint32) *spb.RogueHandbookMiracleInfo {
	db := g.GetRogueHandbookMiracleMap()
	return db[miracleId]
}

func (g *PlayerData) GetRogueHandbookEventMap() map[uint32]*spb.RogueHandbookEventInfo {
	db := g.GetRogueHandbook()
	if db.EventInfo == nil {
		db.EventInfo = make(map[uint32]*spb.RogueHandbookEventInfo)
	}
	return db.EventInfo
}

func (g *PlayerData) GetRogueHandbookEvent(eventId uint32) *spb.RogueHandbookEventInfo {
	db := g.GetRogueHandbookEventMap()
	return db[eventId]
}

func (g *PlayerData) GetRogueHandbookMazeBuffMap() map[uint32]bool {
	db := g.GetRogueHandbook()
	if db.BuffList == nil {
		db.BuffList = make(map[uint32]bool)
	}
	return db.BuffList
}

func (g *PlayerData) GetRogueHandbookAeonMap() map[uint32]*spb.RogueHandbookAeonInfo {
	db := g.GetRogueHandbook()
	if db.AeonInfo == nil {
		db.AeonInfo = make(map[uint32]*spb.RogueHandbookAeonInfo)
	}
	return db.AeonInfo
}

func (g *PlayerData) UnlockRogueHandbook() {
	miracleMap := g.GetRogueHandbookMiracleMap()
	for _, conf := range gdconf.GetRogueHandbookMiracleMap() {
		if miracleMap[conf.MiracleHandbookID] == nil {
			miracleMap[conf.MiracleHandbookID] = &spb.RogueHandbookMiracleInfo{
				MiracleHandbookId: conf.MiracleHandbookID,
				IsTakenReward:     false,
			}
		}
	}

	eventMap := g.GetRogueHandbookEventMap()
	for _, conf := range gdconf.GetRogueHandBookEventMap() {
		if eventMap[conf.EventHandbookID] == nil {
			eventMap[conf.EventHandbookID] = &spb.RogueHandbookEventInfo{
				EventHandbookId: conf.EventHandbookID,
				IsTakenReward:   false,
			}
		}
	}

	buffMap := g.GetRogueHandbookMazeBuffMap()
	for _, id := range gdconf.GetAllBuff() {
		if !buffMap[id] {
			buffMap[id] = true
		}
	}

	aeonMap := g.GetRogueHandbookAeonMap()
	for rogueAeonID, conf := range gdconf.GetRogueAeonStoryConfigMap() {
		if aeonMap[rogueAeonID] == nil {
			aeonMap[rogueAeonID] = &spb.RogueHandbookAeonInfo{
				RogueAeonId:        rogueAeonID,
				UnlockAeonStoryMap: make(map[uint32]bool),
			}
		}
		info := aeonMap[rogueAeonID]
		for _, aeonStoryConf := range conf {
			if !info.UnlockAeonStoryMap[aeonStoryConf.AeonStoryID] {
				info.UnlockAeonStoryMap[aeonStoryConf.AeonStoryID] = true
			}
		}
	}
}

/*********************接口方法*********************/

func (g *PlayerData) GetRogueHandbookMiracleInfoList() []*proto.RogueHandbookMiracleInfo {
	list := make([]*proto.RogueHandbookMiracleInfo, 0)

	for _, info := range g.GetRogueHandbookMiracleMap() {
		list = append(list, &proto.RogueHandbookMiracleInfo{
			HasTakenReward:    info.IsTakenReward,
			MiracleHandbookId: info.MiracleHandbookId,
		})
	}

	return list
}

func (g *PlayerData) GetRogueHandbookEventInfoList() []*proto.RogueHandbookEventInfo {
	list := make([]*proto.RogueHandbookEventInfo, 0)

	for _, info := range g.GetRogueHandbookEventMap() {
		list = append(list, &proto.RogueHandbookEventInfo{
			HasTakenReward:  info.IsTakenReward,
			EventHandbookId: info.EventHandbookId,
		})
	}

	return list
}

func (g *PlayerData) GetRogueHandbookMazeBuffList() []*proto.RogueHandbookMazeBuff {
	list := make([]*proto.RogueHandbookMazeBuff, 0)

	for id := range g.GetRogueHandbookMazeBuffMap() {
		list = append(list, &proto.RogueHandbookMazeBuff{
			MazeBuffId: id,
		})
	}

	return list
}

func (g *PlayerData) GetRogueHandbookAeonInfoList() []*proto.RogueHandbookAeonInfo {
	list := make([]*proto.RogueHandbookAeonInfo, 0)

	for _, db := range g.GetRogueHandbookAeonMap() {
		info := &proto.RogueHandbookAeonInfo{
			AeonId:              db.RogueAeonId,
			UnlockAeonStoryList: make([]uint32, 0),
		}
		for id := range db.UnlockAeonStoryMap {
			info.UnlockAeonStoryList = append(info.UnlockAeonStoryList, id)
		}
		list = append(list, info)
	}

	return list
}
