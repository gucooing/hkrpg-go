package model

import (
	"time"

	"github.com/gucooing/hkrpg-go/dbconf"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func (g *PlayerData) NewQuestRogue(req *proto.StartRogueCsReq) (rogue *spb.CurRogue, err proto.Retcode) {
	db := g.GetBattle()
	conf := gdconf.GetRogueAreaConfigById(req.AreaId)
	if conf == nil {
		err = proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN
		return
	}
	mapId := conf.AreaProgress*100 + conf.Difficulty
	rogueMap := gdconf.GetRogueMapById(mapId)
	if rogueMap == nil {
		err = proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN
		return
	}
	rogue = &spb.CurRogue{
		CurAreaId:     req.AreaId,
		QueuePosition: 0,
		AeonId:        req.AeonId,
		BuffList:      make(map[uint32]*spb.RogueBuff),
		Status:        spb.RogueStatus_ROGUE_STATUS_DOING,
	}
	questRogue := &spb.CurQuestRogue{
		RogueMapId:   mapId,
		CurSiteId:    rogueMap.StartId,
		RogueRoomMap: make(map[uint32]*spb.RogueRoomInfo),
	}
	switch conf.AreaProgress {
	case 0:
		questRogue.RogueRoomMap[1] = &spb.RogueRoomInfo{
			RoomId:         100,
			RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY,
			NextSiteIdList: make([]uint32, 0),
		}
	case 1:
		for id, site := range rogueMap.SiteList {
			questRogue.RogueRoomMap[id] = &spb.RogueRoomInfo{
				RoomId:         gdconf.GetRogueRoomTypeBy100(id),
				RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
				NextSiteIdList: site.NextSiteIDList,
			}
		}
	default:
		for id, site := range rogueMap.SiteList {
			questRogue.RogueRoomMap[id] = &spb.RogueRoomInfo{
				RoomId:         gdconf.GetRogueRoomTypeBySiteID(id),
				RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
				NextSiteIdList: site.NextSiteIDList,
			}
		}
	}
	questRogue.RogueRoomMap[rogueMap.StartId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY

	rogue.RogueInfo = &spb.CurRogue_QuestRogue{QuestRogue: questRogue}
	g.SetMaterialById(uint32(proto.VirtualItemType_VIRTUAL_ITEM_ROGUE_COIN), 100) // 将宇宙碎片重置成100个
	g.SetBattleStatus(spb.BattleType_Battle_QUSET_ROGUE)
	db.CurRogue = rogue
	return
}

func (g *PlayerData) GetQuestRogue() *spb.QuestRogue {
	db := g.GetBattle()
	if db.QuestRogue == nil {
		db.QuestRogue = &spb.QuestRogue{
			RogueArea: make(map[uint32]*spb.QuestRogueArea),
		}
	}
	return db.QuestRogue
}

func (g *PlayerData) GetQuestRogueArea(areaId uint32) *spb.QuestRogueArea {
	rogue := g.GetQuestRogue()
	if rogue.RogueArea == nil {
		rogue.RogueArea = make(map[uint32]*spb.QuestRogueArea)
	}
	if rogue.RogueArea[areaId] == nil {
		rogue.RogueArea[areaId] = &spb.QuestRogueArea{
			AreaId:          areaId,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_UNLOCK,
		}
	}
	return rogue.RogueArea[areaId]
}

func (g *PlayerData) GetQuestRogueRoom() map[uint32]*spb.RogueRoomInfo {
	db := g.GetCurRogue()
	if db == nil {
		return make(map[uint32]*spb.RogueRoomInfo)
	}
	switch x := db.RogueInfo.(type) {
	case *spb.CurRogue_QuestRogue:
		return x.QuestRogue.RogueRoomMap
	}
	return make(map[uint32]*spb.RogueRoomInfo)
}

func (g *PlayerData) GetCurQuestRogueRoom() *spb.RogueRoomInfo {
	db := g.GetCurRogue()
	if db == nil {
		return new(spb.RogueRoomInfo)
	}
	switch x := db.RogueInfo.(type) {
	case *spb.CurRogue_QuestRogue:
		return x.QuestRogue.RogueRoomMap[x.QuestRogue.CurSiteId]
	}
	return new(spb.RogueRoomInfo)
}

func (g *PlayerData) GetCurQuestRogueRoomId() uint32 {
	db := g.GetCurQuestRogueRoom()
	if db == nil {
		return 0
	}
	return db.RoomId
}
func (g *PlayerData) GetQuestRoomBySiteId(siteId uint32) *spb.RogueRoomInfo {
	db := g.GetCurRogue()
	if db == nil {
		return new(spb.RogueRoomInfo)
	}
	switch x := db.RogueInfo.(type) {
	case *spb.CurRogue_QuestRogue:
		return x.QuestRogue.RogueRoomMap[siteId]
	}
	return new(spb.RogueRoomInfo)
}

func (g *PlayerData) SetQuestRogueRoomStatus(Status spb.RoomStatus, siteId uint32) {
	room := g.GetQuestRoomBySiteId(siteId)
	if room == nil {
		return
	}
	room.RoomStatus = Status
}

func (g *PlayerData) GetQuestRogueHistory() map[uint32]*spb.QuestRogueHistory {
	db := g.GetQuestRogue()
	if db.QuestRogueHistoryList == nil {
		db.QuestRogueHistoryList = make(map[uint32]*spb.QuestRogueHistory)
	}
	return db.QuestRogueHistoryList
}

func (g *PlayerData) GetQuestRogueHistoryById(weekId uint32) *spb.QuestRogueHistory {
	db := g.GetQuestRogueHistory()
	if db[weekId] == nil {
		db[weekId] = &spb.QuestRogueHistory{
			WeekId:  weekId,
			Score:   0,
			RowInfo: make(map[uint32]bool),
		}
	}
	return db[weekId]
}

/***************************接口**************************/

func (g *PlayerData) GetQuestRogueInfo() *proto.RogueInfo {
	rogueInfo := &proto.RogueInfo{
		RogueGetInfo: &proto.RogueGetInfo{
			RogueSeasonInfo:      g.GetQuestRogueSeasonInfo(),
			RogueScoreRewardInfo: g.GetQuestRogueScoreRewardInfo(),
			RogueAreaInfo:        g.GetQuestRogueAreaInfo(),
			RogueAeonInfo:        g.GetQuestRogueAeonInfo(),
			RogueVirtualItemInfo: &proto.RogueGetVirtualItemInfo{
				// APAINNEEPIF: 8,
			},
		},
		RogueCurrentInfo: g.GetQuestRogueCurrentInfo(),
	}
	return rogueInfo
}

func (g *PlayerData) GetQuestRogueSeasonInfo() *proto.RogueSeasonInfo {
	conf := dbconf.GetCurRogue()
	if conf == nil {
		return nil
	}
	info := &proto.RogueSeasonInfo{
		EndTime:   4070894399,
		BeginTime: 1711310400,
		Season:    conf.SeasonId,
	}
	return info
}

func (g *PlayerData) GetQuestRogueScoreRewardInfo() *proto.RogueScoreRewardInfo {
	conf := dbconf.GetCurRogue()
	if conf == nil {
		return nil
	}
	year, week := time.Now().ISOWeek()
	db := g.GetQuestRogueHistoryById(uint32((year%10+(year/10)%10*10)*100 + week))
	info := &proto.RogueScoreRewardInfo{
		PoolId:                 20 + g.GetWorldLevel(),
		RewardEndTime:          conf.EndTime.Time.Unix(),
		RewardBeginTime:        conf.BeginTime.Time.Unix(),
		PoolRefreshed:          true,     // 是否刷新
		HasTakenInitialScore:   true,     // 是否已取得初始分数
		ExploreScore:           db.Score, // 本期分数
		TakenNormalFreeRowList: make([]uint32, 0),
	}
	for rowId, row := range db.RowInfo {
		if row {
			info.TakenNormalFreeRowList = append(info.TakenNormalFreeRowList, rowId)
		}
	}
	return info
}

func (g *PlayerData) GetQuestRogueAreaInfo() *proto.RogueAreaInfo {
	info := &proto.RogueAreaInfo{
		RogueAreaList: make([]*proto.RogueArea, 0),
	}
	conf := dbconf.GetCurRogue()
	if conf == nil {
		return info
	}
	cfRogueManager := gdconf.GetRogueManagerById(conf.SeasonId)
	if cfRogueManager == nil {
		return info
	}
	for _, rogueArea := range cfRogueManager.RogueAreaIDList {
		dbRogueArea := g.GetQuestRogueArea(rogueArea)
		RogueArea := &proto.RogueArea{
			AreaId:         dbRogueArea.AreaId,
			AreaStatus:     proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS, // proto.RogueAreaStatus(dbRogueArea.RogueAreaStatus),
			MapId:          0,
			HasTakenReward: true,
			RogueStatus:    0,
			// CurReachRoomNum: 0,
		}
		info.RogueAreaList = append(info.RogueAreaList, RogueArea)
	}

	return info
}

func (g *PlayerData) GetQuestRogueAeonInfo() *proto.RogueAeonInfo {
	info := &proto.RogueAeonInfo{
		IsUnlocked:             true,
		UnlockedAeonEnhanceNum: 3,
		AeonIdList:             []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		UnlockedAeonNum:        9,
	}

	return info
}

func (g *PlayerData) GetQuestRogueCurrentInfo() *proto.RogueCurrentInfo {
	db := g.GetCurRogue()
	if db == nil {
		return nil
	}
	info := &proto.RogueCurrentInfo{
		RogueAeonInfo:   g.GetGameAeonInfo(),
		GameMiracleInfo: g.GetGameMiracleInfo(),
		RogueLineupInfo: g.GetRogueLineupInfo(),
		Status:          proto.RogueStatus(db.Status),
		RogueMap:        g.GetRogueMap(),
		PendingAction:   g.GetFirstRogueAction(),
		// IsExploreWin:    db.IsWin,
		ModuleInfo:      &proto.RogueModuleInfo{ModuleIdList: make([]uint32, 0)},
		VirtualItemInfo: g.GetRogueVirtualItem(),
		RogueBuffInfo:   g.GetRogueBuffInfo(),
	}

	return info
}

func (g *PlayerData) GetRogueScene(roomId uint32) *proto.SceneInfo {
	rogueRoom := gdconf.GetRogueRoomById(roomId)
	if rogueRoom == nil {
		return nil
	}
	mapEntrance := gdconf.GetMapEntranceById(rogueRoom.MapEntrance)
	if mapEntrance == nil {
		return nil
	}
	startGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, rogueRoom.GroupID)
	var pos *proto.Vector
	var rot *proto.Vector
	for _, anchor := range startGroup.AnchorList {
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
		break
	}

	scene := g.GetBattleLoadScene(rogueRoom.MapEntrance, pos, rot, g.GetBattleLineUpById(Rogue))
	for groupID, content := range rogueRoom.GroupWithContent {
		sceneGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, groupID)
		if sceneGroup == nil {
			continue
		}
		// scene.GroupIdList = append(scene.GroupIdList, groupID)
		sceneGroupState := &proto.SceneGroupState{
			GroupId:   groupID,
			IsDefault: true,
		}
		scene.GroupStateList = append(scene.GroupStateList, sceneGroupState)

		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    groupID,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加物品实体
		g.GetRoguePropByID(entityGroupLists, sceneGroup)
		// 添加怪物实体
		g.GetRogueNPCMonsterByID(entityGroupLists, sceneGroup, content)
		// 添加NPC实体
		g.GetSceneNPCByConf(entityGroupLists, sceneGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}

	return scene
}

func (g *PlayerData) GetRogueNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, ida uint32) {
	for _, monsterList := range sceneGroup.MonsterList {
		entityId := g.GetNextGameObjectGuid()
		rogueMonsterID := gdconf.GetRogueMonsterGroupByGroupID(ida)
		rogueMonster := gdconf.GetRogueMonsterByRogueMonsterID(rogueMonsterID)
		pos := &proto.Vector{
			X: int32(monsterList.PosX * 1000),
			Y: int32(monsterList.PosY * 1000),
			Z: int32(monsterList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(monsterList.RotX * 1000),
			Y: int32(monsterList.RotY * 1000),
			Z: int32(monsterList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			EntityOneofCase: &proto.SceneEntityInfo_NpcMonster{
				NpcMonster: &proto.SceneNpcMonsterInfo{
					WorldLevel: g.GetWorldLevel(),
					MonsterId:  rogueMonster.NpcMonsterID,
					EventId:    rogueMonster.EventID,
				},
			},
		}
		// 添加实体
		g.AddEntity(sceneGroup.GroupId, &MonsterEntity{
			Entity: Entity{
				InstId:   monsterList.ID,
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			EventID: rogueMonster.EventID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}

func (g *PlayerData) GetRoguePropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) {
	for _, propList := range sceneGroup.PropList {
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(propList.PosX * 1000),
			Y: int32(propList.PosY * 1000),
			Z: int32(propList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(propList.RotX * 1000),
			Y: int32(propList.RotY * 1000),
			Z: int32(propList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId, // 文件名后那个G
			InstId:   propList.ID,        // ID
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
		}
		prop := &proto.ScenePropInfo{
			PropId:    propList.PropID, // PropID
			PropState: gdconf.GetStateValue(propList.State),
		}
		if propList.PropID == 1000 || propList.PropID == 1021 || propList.PropID == 1022 || propList.PropID == 1023 {
			index := 0
			if propList.Name == "Door2" {
				index = 1
			}
			room := g.GetCurQuestRogueRoom()
			if len(room.NextSiteIdList) == 1 {
				index = 0
			}
			if len(room.NextSiteIdList) > 0 {
				siteId := room.NextSiteIdList[index]
				nextRoom := g.GetQuestRoomBySiteId(siteId)
				exceRoom := gdconf.GetRogueRoomById(nextRoom.RoomId)

				switch exceRoom.RogueRoomType {
				case 3, 8:
					prop.PropId = 1022
				case 5:
					prop.PropId = 1023
				default:
					prop.PropId = 1021
				}
				prop.ExtraInfo = &proto.PropExtraInfo{
					RogueInfo: &proto.PropRogueInfo{
						RoomId: nextRoom.RoomId,
						SiteId: siteId,
					},
				}
			} else {
				prop.ExtraInfo = &proto.PropExtraInfo{}
				prop.PropId = 1000
			}
			prop.PropState = 1
		}
		entityList.EntityOneofCase = &proto.SceneEntityInfo_Prop{Prop: prop}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}
