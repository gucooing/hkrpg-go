package player

import (
	gsdb "github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/*************模拟宇宙*************/

func (g *GamePlayer) GetDbRogue() *spb.Rogue {
	db := g.GetBattle()
	if db.Rogue == nil {
		db.Rogue = &spb.Rogue{
			RogueArea: make(map[uint32]*spb.RogueArea),
		}
	}

	return db.Rogue
}

func (g *GamePlayer) GetCurDbRogue() *spb.CurRogue {
	rogue := g.GetDbRogue()
	if rogue.CurRogue == nil {
		rogue.CurRogue = new(spb.CurRogue)
	}
	return rogue.CurRogue
}

func (g *GamePlayer) GetCurDbRoom() *spb.RogueRoom {
	curRogue := g.GetCurDbRogue()
	return curRogue.RogueSceneMap[curRogue.CurSiteId]
}

func (g *GamePlayer) GetDbRoomBySiteId(siteId uint32) *spb.RogueRoom {
	curRogue := g.GetCurDbRogue()
	return curRogue.RogueSceneMap[siteId]
}

func (g *GamePlayer) GetDbRogueArea(areaId uint32) *spb.RogueArea {
	rogue := g.GetDbRogue()
	if rogue.RogueArea == nil {
		rogue.RogueArea = make(map[uint32]*spb.RogueArea)
		rogue.RogueArea[100] = &spb.RogueArea{
			AreaId:          100,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_UNLOCK,
		}
	}
	if rogue.RogueArea[areaId] == nil {
		rogue.RogueArea[areaId] = &spb.RogueArea{
			AreaId:          areaId,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_LOCK,
		}
	}

	return rogue.RogueArea[areaId]
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetRogueScoreRewardInfo() *proto.RogueScoreRewardInfo {
	conf := gsdb.GetCurRogue()
	if conf == nil {
		return nil
	}
	info := &proto.RogueScoreRewardInfo{
		PoolId:                 20 + g.GetWorldLevel(),
		EndTime:                conf.EndTime.Time.Unix(),
		BeginTime:              conf.EndTime.Time.Unix(),
		PoolRefreshed:          true,  // 是否刷新
		HasTakenInitialScore:   false, // 是否已取得初始分数
		ExploreScore:           0,     // 本期分数
		TakenNormalFreeRowList: make([]uint32, 0),
	}
	return info
}

func (g *GamePlayer) GetRogueSeasonInfo() *proto.RogueSeasonInfo {
	conf := gsdb.GetCurRogue()
	if conf == nil {
		return nil
	}
	info := &proto.RogueSeasonInfo{
		EndTime:   conf.EndTime.Time.Unix(),
		BeginTime: conf.EndTime.Time.Unix(),
		Season:    conf.SeasonId,
	}
	return info
}
