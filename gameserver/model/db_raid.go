package model

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func (g *PlayerData) GetRaid() *spb.Raid {
	db := g.GetBattle()
	if db.Rain == nil {
		db.Rain = &spb.Raid{}
	}
	return db.Rain
}

func (g *PlayerData) GetRaidMap() map[uint32]*spb.RaidInfo {
	db := g.GetRaid()
	if db.RaidMap == nil {
		db.RaidMap = make(map[uint32]*spb.RaidInfo)
	}
	return db.RaidMap
}

func (g *PlayerData) GetFinishRaidMap() map[uint32]*spb.RaidInfo {
	db := g.GetRaid()
	if db.FinishRaidMap == nil {
		db.FinishRaidMap = make(map[uint32]*spb.RaidInfo)
	}
	return db.FinishRaidMap
}

func (g *PlayerData) GetFinishRaidInfo(raid uint32) *spb.RaidInfo {
	db := g.GetFinishRaidMap()
	return db[raid]
}

func (g *PlayerData) NewRaidInfoDb(raid uint32) {
	db := g.GetRaid()
	if db.CurRaidId == raid {
		db.CurRaidId = 0
	}
	delete(g.GetRaidMap(), raid)
}

func (g *PlayerData) GetRaidInfo(raid uint32) *spb.RaidInfo {
	db := g.GetRaidMap()
	return db[raid]
}

func (g *PlayerData) GetCurRaidInfo() *spb.RaidInfo {
	db := g.GetRaid()
	if db.RaidMap == nil {
		db.RaidMap = make(map[uint32]*spb.RaidInfo)
	}
	return db.RaidMap[db.CurRaidId]
}

func (g *PlayerData) NewRaidInfo(req *proto.StartRaidCsReq) proto.Retcode {
	raidDb := g.GetRaid()
	raidDb.CurRaidId = req.RaidId
	db := g.GetRaidMap()
	db[req.RaidId] = &spb.RaidInfo{}
	raidInfo := db[req.RaidId]
	raidInfo.RaidId = req.RaidId
	raidInfo.HardLevel = req.WorldLevel
	conf := gdconf.GetRaidConfig(req.RaidId, req.WorldLevel)
	if conf == nil {
		return proto.Retcode_RET_RECENT_ELEMENT_STAGE_NOT_MATCH
	}
	// 获取任务
	mainMission := conf.MainMissionIDList[0]
	missionConf := gdconf.GetGoppMainMissionById(mainMission)
	if missionConf == nil {
		return proto.Retcode_RET_RECENT_ELEMENT_STAGE_NOT_MATCH
	}
	g.DelMainMission(conf.MainMissionIDList) // 删除任务
	g.AddMainMission([]uint32{mainMission})  // 添加任务
	// 获取场景
	subMissionConf := gdconf.GetSubMainMissionById(missionConf.StartSubMissionList[0])
	if subMissionConf == nil {
		return proto.Retcode_RET_RECENT_ELEMENT_STAGE_NOT_MATCH
	}
	raidInfo.EntryId = FloorTentry(subMissionConf.LevelFloorID)
	// 重置场景状态
	blockBin := g.GetBlock(raidInfo.EntryId)
	blockBin.BlockList = make(map[uint32]*spb.BlockList)
	raidInfo.Status = spb.RaidStatus_RAID_STATUS_DOING
	return proto.Retcode_RET_SUCC
}

// 设置队伍
func (g *PlayerData) SetRaidLineUp(req *proto.StartRaidCsReq, conf *gdconf.RaidConfig) []uint32 {
	avatarList := make([]uint32, 0)
	switch conf.TeamType {
	case constant.RaidTeamTypePlayer: // 原有
		for _, info := range g.GetCurLineUp().AvatarIdList {
			avatarList = append(avatarList, info.AvatarId)
		}
	case constant.RaidTeamTypeTrial: // 仅试用
		avatarList = conf.TrialAvatarList
	case constant.RaidTeamTypeTrialOnly: // 仅试用
		avatarList = conf.TrialAvatarList
	case constant.RaidTeamTypeTrialAndPlayer: // 原有补位试用
		avatarList = conf.TrialAvatarList
	case constant.RaidTeamTypeTrialOrPlayer: // 选择的角色中必须要有这个试用
		avatarList = req.AvatarList
	}
	return avatarList
}

func (g *PlayerData) CheckRaid() (uint32, bool) {
	db := g.GetCurRaidInfo()
	if db == nil {
		return 0, false
	}
	conf := gdconf.GetRaidConfig(db.RaidId, db.HardLevel)
	if conf == nil {
		return 0, false
	}
	var finish = true
	finishMainMissionList := g.GetFinishMainMissionList()
	for _, id := range conf.MainMissionIDList {
		if finishMainMissionList[id] == nil {
			finish = false
			break
		}
	}
	if finish {
		finishDb := g.GetFinishRaidMap()
		if finishDb[db.RaidId] == nil {
			finishDb[db.RaidId] = db
		}
		db.Status = spb.RaidStatus_RAID_STATUS_FINISH
		db.FinishTime = uint64(time.Now().Unix())
		return db.RaidId, true
	}
	return 0, false
}

func (g *PlayerData) RaidReward(raidID, hardLevel uint32, addItem *AddItem) {
	conf := gdconf.GetRaidConfig(raidID, hardLevel)
	if conf != nil {
		addItem = NewAddItem(addItem)
		switch conf.Type {
		case constant.RaidConfigTypeEquilibriumTrial:
			g.AddWorldLevel(1)
		default:
			if !conf.SkipRewardOnFinish && conf.RewardList != nil {
				for _, reward := range conf.RewardList {
					pile := GetRewardData(reward)
					addItem.PileItem = append(addItem.PileItem, pile...)
				}
			}
		}
		g.AddItem(addItem)
	}
}
