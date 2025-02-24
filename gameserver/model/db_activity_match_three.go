package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func newMatchThree() *spb.MatchThree {
	return &spb.MatchThree{}
}

func newBirdInfoMap() map[uint32]*spb.MatchThreeBirdInfo {
	list := make(map[uint32]*spb.MatchThreeBirdInfo)
	for _, conf := range gdconf.GetMatchThreeBirdMap() {
		if conf.UnlockLevel == 0 {
			list[conf.BirdID] = &spb.MatchThreeBirdInfo{
				BirdId: conf.BirdID,
			}
		}
	}
	return list
}

func (g *PlayerData) GetMatchThree() *spb.MatchThree {
	db := g.GetActivity()
	if db.MatchThree == nil {
		db.MatchThree = newMatchThree()
	}

	return db.MatchThree
}

func (g *PlayerData) GetMatchThreeLevelList() map[uint32]*spb.MatchThreeLevelMap {
	db := g.GetMatchThree()
	if db.LevelList == nil {
		db.LevelList = make(map[uint32]*spb.MatchThreeLevelMap)
	}
	return db.LevelList
}

func (g *PlayerData) GetMatchThreeLevelMap(levelId uint32) *spb.MatchThreeLevelMap {
	db := g.GetMatchThreeLevelList()
	if db == nil {
		db = make(map[uint32]*spb.MatchThreeLevelMap)
	}
	if db[levelId] == nil {
		db[levelId] = &spb.MatchThreeLevelMap{
			Level: map[uint32]*spb.MatchThreeLevel{},
		}
	}
	return db[levelId]
}

func (g *PlayerData) GetMatchThreeLevel(levelId, mode uint32) *spb.MatchThreeLevel {
	db := g.GetMatchThreeLevelMap(levelId)
	if db.Level == nil {
		db.Level = make(map[uint32]*spb.MatchThreeLevel)
	}
	return db.Level[mode]
}

func (g *PlayerData) AddMatchThreeLevel(level *spb.MatchThreeLevel) {
	db := g.GetMatchThreeLevelMap(level.LevelId)
	if db.Level == nil {
		db.Level = make(map[uint32]*spb.MatchThreeLevel)
	}
	db.Level[level.Mode] = level
}

func (g *PlayerData) GetMatchThreeBirdInfoMap() map[uint32]*spb.MatchThreeBirdInfo {
	db := g.GetMatchThree()
	if db.BirdInfoMap == nil {
		db.BirdInfoMap = newBirdInfoMap()
	}
	return db.BirdInfoMap
}

func (g *PlayerData) GetMatchThreeBirdInfo(birdId uint32) *spb.MatchThreeBirdInfo {
	db := g.GetMatchThreeBirdInfoMap()
	if db[birdId] == nil {
		db[birdId] = &spb.MatchThreeBirdInfo{
			BirdId: birdId,
		}
	}
	return db[birdId]
}

func (g *PlayerData) UpMatchThreeBirdInfo(birdId, score uint32) {
	db := g.GetMatchThreeBirdInfoMap()
	for _, conf := range gdconf.GetMatchThreeBirdMap() {
		if conf.UnlockLevel == 0 ||
			len(g.GetMatchThreeLevelMap(conf.UnlockLevel).Level) != 0 {
			db[conf.BirdID] = &spb.MatchThreeBirdInfo{
				BirdId: conf.BirdID,
			}
		}
	}

	if db[birdId] != nil {
		db[birdId].Count++
		db[birdId].BirdTopScore = alg.MaxUin32(db[birdId].BirdTopScore, score)
	}
}

func (g *PlayerData) SetMatchThreeBirdPos(birdId, pos uint32) bool {
	db := g.GetMatchThreeBirdInfo(birdId)
	if db == nil {
		return false
	}
	db.Pos = pos
	return true
}

/*****************************接口***************************/

func (g *PlayerData) GetMatchThreeData() *proto.MatchThreeData {
	info := &proto.MatchThreeData{
		FinishedLevels:  g.MatchThreeFinishedLevelInfo(),
		BirdRecordInfos: g.GetMatchThreeBirdInfoList(),
	}

	return info
}

func (g *PlayerData) MatchThreeFinishedLevelInfo() []*proto.MatchThreeFinishedLevelInfos {
	list := make([]*proto.MatchThreeFinishedLevelInfos, 0)
	for _, modeList := range g.GetMatchThreeLevelList() {
		if modeList.Level == nil {
			continue
		}
		for _, info := range modeList.Level {
			list = append(list, &proto.MatchThreeFinishedLevelInfos{
				ModeId:  info.Mode,
				LevelId: info.LevelId,
			})
		}
	}
	return list
}

func (g *PlayerData) GetMatchThreeBirdInfoList() []*proto.MatchThreeBirdInfo {
	list := make([]*proto.MatchThreeBirdInfo, 0)
	for _, info := range g.GetMatchThreeBirdInfoMap() {
		list = append(list, &proto.MatchThreeBirdInfo{
			BirdTopScore: info.BirdTopScore,
			Pos:          0,
			BirdId:       info.BirdId,
			Count:        info.Count,
		})
	}
	return list
}
