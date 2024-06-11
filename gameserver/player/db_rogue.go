package player

import (
	gsdb "github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
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

func (g *GamePlayer) GetRogueHistory() map[uint32]*spb.RogueHistory {
	db := g.GetDbRogue()
	if db.RogueHistoryList == nil {
		db.RogueHistoryList = make(map[uint32]*spb.RogueHistory)
	}
	return db.RogueHistoryList
}

func (g *GamePlayer) GetRogueHistoryById(id uint32) (*spb.RogueHistory, bool) {
	db := g.GetRogueHistory()
	isPoolRefreshed := false
	if db[id] == nil {
		isPoolRefreshed = true
		db[id] = &spb.RogueHistory{
			SeasonId: id,
			Score:    0,
		}
	}
	return db[id], isPoolRefreshed
}

func (g *GamePlayer) GetCurRogue() *spb.CurRogue {
	db := g.GetDbRogue()
	if db.CurRogue == nil {
		db.CurRogue = new(spb.CurRogue)
	}
	return db.CurRogue
}

func (g *GamePlayer) GetRogueRoom() map[uint32]*spb.RogueRoom {
	db := g.GetCurRogue()
	if db.RogueRoomMap == nil {
		db.RogueRoomMap = make(map[uint32]*spb.RogueRoom)
	}
	return db.RogueRoomMap
}

func (g *GamePlayer) GetCurRogueRoom() *spb.RogueRoom {
	db := g.GetCurRogue()
	if db.RogueRoomMap == nil {
		db.RogueRoomMap = make(map[uint32]*spb.RogueRoom)
	}
	return db.RogueRoomMap[db.CurSiteId]
}

func (g *GamePlayer) GetCurRogueRoomId() uint32 {
	db := g.GetCurRogueRoom()
	if db == nil {
		return 0
	}
	return db.RoomId
}
func (g *GamePlayer) GetRoomBySiteId(siteId uint32) *spb.RogueRoom {
	db := g.GetCurRogue()
	if db.RogueRoomMap == nil {
		db.RogueRoomMap = make(map[uint32]*spb.RogueRoom)
	}
	return db.RogueRoomMap[siteId]
}

func (g *GamePlayer) GetDbRogueArea(areaId uint32) *spb.RogueArea {
	rogue := g.GetDbRogue()
	if rogue.RogueArea == nil {
		rogue.RogueArea = make(map[uint32]*spb.RogueArea)
	}
	if rogue.RogueArea[areaId] == nil {
		rogue.RogueArea[areaId] = &spb.RogueArea{
			AreaId:          areaId,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_UNLOCK,
		}
	}
	return rogue.RogueArea[areaId]
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetRogueInfo() *proto.RogueInfo {
	rogueInfo := &proto.RogueInfo{
		RogueGetInfo: &proto.RogueGetInfo{
			RogueSeasonInfo:      g.GetRogueSeasonInfo(),
			RogueScoreRewardInfo: g.GetRogueScoreRewardInfo(),
			RogueAreaInfo:        g.GetRogueAreaInfo(),
			RogueAeonInfo:        g.GetRogueAeonInfo(),
			RogueVirtualItemInfo: &proto.RogueGetVirtualItemInfo{},
		},
		RogueCurrentInfo: g.GetRogueCurrentInfo(),
	}
	return rogueInfo
}

func (g *GamePlayer) GetRogueCurrentInfo() *proto.RogueCurrentInfo {
	info := &proto.RogueCurrentInfo{
		RogueAeonInfo:    g.GetGameAeonInfo(),
		GameMiracleInfo:  nil,
		RogueLineupInfo:  nil,
		Status:           0,
		MapInfo:          g.GetRogueMap(),
		PendingAction:    nil,
		IsWin:            false,
		ModuleInfo:       nil,
		RogueVirtualItem: nil,
		RogueBuffInfo:    nil,
	}

	return info
}

func (g *GamePlayer) GetRogueScoreRewardInfo() *proto.RogueScoreRewardInfo {
	conf := gsdb.GetCurRogue()
	if conf == nil {
		return nil
	}
	db, poolRefreshed := g.GetRogueHistoryById(conf.SeasonId)
	info := &proto.RogueScoreRewardInfo{
		PoolId:                 20 + g.GetWorldLevel(),
		EndTime:                conf.EndTime.Time.Unix(),
		BeginTime:              conf.EndTime.Time.Unix(),
		PoolRefreshed:          poolRefreshed, // 是否刷新
		HasTakenInitialScore:   false,         // 是否已取得初始分数
		ExploreScore:           db.Score,      // 本期分数
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

func (g *GamePlayer) GetRogueAreaInfo() *proto.RogueAreaInfo {
	info := &proto.RogueAreaInfo{RogueAreaList: make([]*proto.RogueArea, 0)}
	conf := gsdb.GetCurRogue()
	if conf == nil {
		return info
	}
	cfRogueManager := gdconf.GetRogueManagerById(conf.SeasonId)
	if cfRogueManager == nil {
		return info
	}
	for _, rogueArea := range cfRogueManager.RogueAreaIDList {
		dbRogueArea := g.GetDbRogueArea(rogueArea)
		RogueArea := &proto.RogueArea{
			AreaId:     dbRogueArea.AreaId,
			AreaStatus: proto.RogueAreaStatus(dbRogueArea.RogueAreaStatus),

			MapId:           0,
			HasTakenReward:  false,
			RogueStatus:     0,
			CurReachRoomNum: 0,
		}
		info.RogueAreaList = append(info.RogueAreaList, RogueArea)
	}

	return info
}

func (g *GamePlayer) GetRogueAeonInfo() *proto.RogueAeonInfo {
	info := &proto.RogueAeonInfo{
		IsUnlocked:             true,
		UnlockedAeonEnhanceNum: 3,
		AeonIdList:             []uint32{1, 2, 3, 4, 5, 6, 7},
		UnlockedAeonNum:        9,
	}

	return info
}

func (g *GamePlayer) GetGameAeonInfo() *proto.GameAeonInfo {
	rogue := g.GetCurRogue()
	info := &proto.GameAeonInfo{
		IsUnlocked:             true,
		UnlockedAeonEnhanceNum: 3,
		AeonId:                 rogue.AeonId,
	}
	return info
}

func (g *GamePlayer) GetRogueMap() *proto.RogueMapInfo {
	rogue := g.GetCurRogue()
	roomMap := &proto.RogueMapInfo{
		MapId:     rogue.RogueMapId,
		AreaId:    rogue.CurAreaId,
		CurSiteId: rogue.CurSiteId, // 当前id
		CurRoomId: g.GetCurRogueRoomId(),
		RoomList:  make([]*proto.RogueRoom, 0),
	}
	for id, rogueScene := range rogue.RogueRoomMap {
		roomList := &proto.RogueRoom{
			SiteId:    id,
			RoomId:    rogueScene.RoomId,
			CurStatus: proto.RogueRoomStatus(rogueScene.RoomStatus),
		}
		roomMap.RoomList = append(roomMap.RoomList, roomList)
	}

	return roomMap
}

func (g *GamePlayer) GetRogueScene(roomId uint32) *proto.SceneInfo {
	rogueRoom := gdconf.GetRogueRoomById(roomId)
	if rogueRoom == nil {
		return nil
	}
	mapEntrance := gdconf.GetMapEntranceById(rogueRoom.MapEntrance)
	if mapEntrance == nil {
		return nil
	}
	leaderEntityId := g.GetNextGameObjectGuid()
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            mapEntrance.PlaneID,
		FloorId:            mapEntrance.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID,
		EntryId:            rogueRoom.MapEntrance,
		GameModeType:       5, // gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}
	// 获取场景实体
	entityGroupList := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
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
	lineUp := g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ROGUE))

	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroupList, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)

	for groupID, ida := range rogueRoom.GroupWithContent {
		sceneGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, groupID)
		if sceneGroup == nil {
			continue
		}
		scene.GroupIdList = append(scene.GroupIdList, groupID)
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
		g.GetRogueNPCMonsterByID(entityGroupLists, sceneGroup, ida)
		// 添加NPC实体
		g.GetNPCByID(entityGroupLists, sceneGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}

	return scene
}

func (g *GamePlayer) GetRogueNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, ida uint32) {
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
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.GetWorldLevel(),
				MonsterId:  rogueMonster.NpcMonsterID,
				EventId:    rogueMonster.EventID,
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

func (g *GamePlayer) GetRoguePropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) {
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
			Prop: &proto.ScenePropInfo{
				PropId:    propList.PropID, // PropID
				PropState: gdconf.GetStateValue(propList.State),
			},
		}
		// if propList.PropID == 1000 || propList.PropID == 1021 || propList.PropID == 1022 || propList.PropID == 1023 {
		// 	index := 0
		// 	if propList.Name == "Door2" {
		// 		index = 1
		// 	}
		// 	room := g.GetCurDbRoom()
		// 	if propList.Name == "Door1" && len(room.NextSiteIdList) == 1 {
		// 		continue
		// 	}
		// 	if len(room.NextSiteIdList) == 1 {
		// 		index = 0
		// 	}
		// 	if len(room.NextSiteIdList) > 0 {
		// 		siteId := room.NextSiteIdList[index]
		// 		nextRoom := g.GetDbRoomBySiteId(siteId)
		// 		exceRoom := gdconf.GetRogueRoomById(nextRoom.RoomId)
		//
		// 		switch exceRoom.RogueRoomType {
		// 		case 3, 8:
		// 			entityList.Prop.PropId = 1022
		// 		case 5:
		// 			entityList.Prop.PropId = 1023
		// 		default:
		// 			entityList.Prop.PropId = 1021
		// 		}
		// 		entityList.Prop.ExtraInfo = &proto.PropExtraInfo{
		// 			InfoOneofCase: &proto.PropExtraInfo_RogueInfo{
		// 				RogueInfo: &proto.PropRogueInfo{
		// 					RoomId: nextRoom.RoomId,
		// 					SiteId: siteId,
		// 				},
		// 			},
		// 		}
		// 	} else {
		// 		entityList.Prop.PropId = 1000
		// 	}
		// 	entityList.Prop.PropState = 1
		// }
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}
