package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetRaid() *spb.Raid {
	db := g.GetBattle()
	if db.Rain == nil {
		db.Rain = &spb.Raid{}
	}
	return db.Rain
}

func (g *GamePlayer) GetRaidMap() map[uint32]*spb.RaidInfo {
	db := g.GetRaid()
	if db.RaidMap == nil {
		db.RaidMap = make(map[uint32]*spb.RaidInfo)
	}
	return db.RaidMap
}

func (g *GamePlayer) GetFinishRaidMap() map[uint32]*spb.RaidInfo {
	db := g.GetRaid()
	if db.FinishRaidMap == nil {
		db.FinishRaidMap = make(map[uint32]*spb.RaidInfo)
	}
	return db.FinishRaidMap
}

func (g *GamePlayer) GetFinishRaidInfo(raid uint32) *spb.RaidInfo {
	db := g.GetFinishRaidMap()
	return db[raid]
}

func (g *GamePlayer) NewRaidInfo(raid uint32) {
	db := g.GetRaid()
	if db.CurRaidId == raid {
		db.CurRaidId = 0
	}
	delete(g.GetRaidMap(), raid)
}

func (g *GamePlayer) GetRaidInfo(raid uint32) *spb.RaidInfo {
	db := g.GetRaidMap()
	return db[raid]
}

func (g *GamePlayer) GetCurRaidInfo() *spb.RaidInfo {
	db := g.GetRaid()
	if db.RaidMap == nil {
		db.RaidMap = make(map[uint32]*spb.RaidInfo)
	}
	return db.RaidMap[db.CurRaidId]
}

func (g *GamePlayer) newRaidInfo(req *proto.StartRaidCsReq) proto.Retcode {
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
	// 调整队伍
	g.SetRaidLineUp(req, conf)
	// 获取任务
	mainMission := conf.MainMissionIDList[0]
	missionConf := gdconf.GetGoppMainMissionById(mainMission)
	if missionConf == nil {
		return proto.Retcode_RET_RECENT_ELEMENT_STAGE_NOT_MATCH
	}
	g.DelMainMission(conf.MainMissionIDList) // 删除任务
	g.AddMainMission([]uint32{mainMission})  // 添加任务
	g.InspectMission(nil)                    // 检查任务
	// 获取场景
	subMissionConf := gdconf.GetSubMainMissionById(missionConf.StartSubMissionList[0])
	if subMissionConf == nil {
		return proto.Retcode_RET_RECENT_ELEMENT_STAGE_NOT_MATCH
	}
	raidInfo.EntryId = floorTentry(subMissionConf.LevelFloorID)
	// 重置场景状态
	blockBin := g.GetBlock(raidInfo.EntryId)
	blockBin.BlockList = make(map[uint32]*spb.BlockList)
	raidInfo.Status = spb.RaidStatus_RAID_STATUS_DOING
	return proto.Retcode_RET_SUCC
}

// 设置队伍
func (g *GamePlayer) SetRaidLineUp(req *proto.StartRaidCsReq, conf *gdconf.RaidConfig) {
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
	g.SetBattleLineUp(Raid, avatarList)
}

func (g *GamePlayer) CheckRaid() {
	db := g.GetCurRaidInfo()
	if db == nil {
		return
	}
	conf := gdconf.GetRaidConfig(db.RaidId, db.HardLevel)
	if conf == nil {
		return
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
		g.RaidInfoNotify(db.RaidId)
	}
}

func (g *GamePlayer) RaidInfoNotify(raidID uint32) {
	db := g.GetRaidInfo(raidID)
	if db == nil {
		return
	}
	notify := &proto.RaidInfoNotify{
		ItemList:   &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		RaidId:     db.RaidId,
		Status:     proto.RaidStatus(db.Status),
		WorldLevel: db.HardLevel,
		FinishTime: db.FinishTime,
	}
	// TODO 有重复领取的问题，db加个字段就行了
	if db.Status == spb.RaidStatus_RAID_STATUS_FINISH {
		notify.ItemList.ItemList = g.RaidReward(db.RaidId, db.HardLevel)
	}

	g.Send(cmd.RaidInfoNotify, notify)
}

func (g *GamePlayer) RaidReward(raidID, hardLevel uint32) []*proto.Item {
	conf := gdconf.GetRaidConfig(raidID, hardLevel)
	itemList := make([]*proto.Item, 0)
	if conf == nil {
		return itemList
	}
	allSync := &AllPlayerSync{IsBasic: true}
	pileItem := make([]*Material, 0)
	switch conf.Type {
	case constant.RaidConfigTypeEquilibriumTrial:
		g.AddWorldLevel(1)
	default:
		if !conf.SkipRewardOnFinish && conf.RewardList != nil {
			for _, reward := range conf.RewardList {
				rewardConf := gdconf.GetRewardDataById(reward)
				if rewardConf == nil {
					continue
				}
				allSync.MaterialList = append(allSync.MaterialList, Hcoin)
				pileItem = append(pileItem, &Material{
					Tid: Hcoin,
					Num: rewardConf.Hcoin,
				})
				itemList = append(itemList, &proto.Item{
					ItemId: Hcoin,
					Num:    rewardConf.Hcoin,
				})
				for _, info := range rewardConf.Items {
					allSync.MaterialList = append(allSync.MaterialList, info.ItemID)
					pileItem = append(pileItem, &Material{
						Tid: info.ItemID,
						Num: info.Count,
					})
					itemList = append(itemList, &proto.Item{
						ItemId: info.ItemID,
						Num:    info.Count,
					})
				}
			}
		}
	}
	g.AllPlayerSyncScNotify(allSync)
	return itemList
}

func (g *GamePlayer) RaidEnterSceneByServerScNotify(entryId uint32) {
	rsp := new(proto.EnterSceneByServerScNotify)
	mapEntrance := gdconf.GetMapEntranceById(entryId)
	if mapEntrance == nil {
		return
	}
	teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if teleportsMap == nil {
		return
	}
	anchorID := mapEntrance.StartAnchorID
	groupID := mapEntrance.StartGroupID

	var pos *proto.Vector
	var rot *proto.Vector

	// 获取队伍
	lineDb := g.GetBattleLineUpById(Raid)
	rsp.Lineup = g.GetLineUpPb(lineDb)
	// 获取坐标
	var anchor *gdconf.AnchorList
	if anchorID == 0 || groupID == 0 {
		anchor = gdconf.GetAnchorByIndex(mapEntrance.PlaneID, mapEntrance.FloorID)
	} else {
		anchor = gdconf.GetAnchor(mapEntrance.PlaneID, mapEntrance.FloorID, groupID, anchorID)
	}

	if anchor != nil {
		pos = &proto.Vector{
			X: int32(anchor.PosX * 1000),
			Y: int32(anchor.PosY * 1000),
			Z: int32(anchor.PosZ * 1000),
		}
		rot = &proto.Vector{
			X: int32(anchor.RotX * 1000),
			Y: int32(anchor.RotY * 1000),
			Z: int32(anchor.RotZ * 1000),
		}
	}

	if pos == nil {
		// 这都没有那就不要传送了
		logger.Error("raid entryId:%v error", entryId)
		g.Send(cmd.EnterSceneByServerScNotify, rsp)
		return
	}
	rsp.Scene = g.GetRaidSceneInfo(entryId, pos, rot, lineDb)
	// g.SetCurEntryId(entryId)
	g.EnterMapByEntrance(entryId) // 任务检查
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *GamePlayer) GetRaidSceneInfo(entryId uint32, pos, rot *proto.Vector, lineUp *spb.Line) *proto.SceneInfo {
	leaderEntityId := g.GetNextGameObjectGuid()
	mapEntrance := gdconf.GetMapEntranceById(entryId)
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetServerGroup(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return nil
	}
	scene := &proto.SceneInfo{
		WorldId:            gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID,
		LeaderEntityId:     leaderEntityId,
		FloorId:            mapEntrance.FloorID,
		GameModeType:       8,
		PlaneId:            mapEntrance.PlaneID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		SceneMissionInfo:   g.GetMissionStatusBySceneInfo(gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)),
	}
	for i := uint32(0); i < 100; i++ {
		scene.LightenSectionList = append(scene.LightenSectionList, i)
	}
	// 获取场景实体
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 清理老实体列表
	g.UpSceneMap()
	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
	blockBin := g.GetBlock(entryId)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)
	for _, levelGroup := range foorMap {
		if !g.IfLoadMap(levelGroup) {
			g.AddNoLoadedGroup(entryId, mapEntrance.PlaneID, mapEntrance.FloorID, levelGroup.GroupId)
			continue
		} else {
			g.AddLoadedGroup(entryId, mapEntrance.PlaneID, mapEntrance.FloorID, levelGroup.GroupId)
		}
		scene.GroupIdList = append(scene.GroupIdList, levelGroup.GroupId)
		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加物品实体
		g.GetPropByID(entityGroupLists, levelGroup, blockBin, entryId)
		// 添加怪物实体
		g.GetNPCMonsterByID(entityGroupLists, levelGroup)
		// 添加NPC实体
		g.GetNPCByID(entityGroupLists, levelGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}
	return scene
}
