package model

import (
	"math/rand"

	"github.com/gucooing/hkrpg-go/dbconf"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	QuestRogue      = 0       // 模拟宇宙
	RogueDlc        = 6000302 // 模拟宇宙：寰宇蝗灾
	RogueNous       = 6000901 // 模拟宇宙：黄金与机械
	QuestRogueTourn = 301     // 差分宇宙
)

// Default Probability
const (
	RogueBuffType      = 900  // 各属性
	AddRogueBuffType   = 1900 // 属性增加概率
	RogueBuffRarityOne = 6000 // 白
	RogueBuffRarityTwo = 3000 // 蓝
	AddRogueBuffRarity = 1000 // 品质增加概率
)

type RogueInfoOnline struct { // 模拟宇宙临时数据
	RogueBuffByType    map[uint32]*RogueBuffByType
	RogueBuffRarityOne int32 // 白的概率
	RogueBuffRarityTwo int32 // 蓝的概率
}

type RogueBuffByType struct {
	Weight          int32                       // 权重
	RogueBuffRarity map[uint32]*RogueBuffRarity // 稀有度
}

type RogueBuffRarity struct {
	Rarity   uint32   // 稀有度
	BuffList []uint32 // buff列表
}

func (g *PlayerData) GetDbRogue() *spb.Rogue {
	db := g.GetBattle()
	if db.Rogue == nil {
		db.Rogue = &spb.Rogue{
			RogueArea: make(map[uint32]*spb.RogueArea),
		}
	}
	return db.Rogue
}

func (g *PlayerData) GetRogueHistory() map[uint32]*spb.RogueHistory {
	db := g.GetDbRogue()
	if db.RogueHistoryList == nil {
		db.RogueHistoryList = make(map[uint32]*spb.RogueHistory)
	}
	return db.RogueHistoryList
}

func (g *PlayerData) GetRogueHistoryById(id uint32) (*spb.RogueHistory, bool) {
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

func (g *PlayerData) NewCurRogue() {
	db := g.GetDbRogue()
	db.CurRogue = new(spb.CurRogue)
}

func (g *PlayerData) GetCurRogue() *spb.CurRogue {
	db := g.GetDbRogue()
	if db.CurRogue == nil {
		db.CurRogue = new(spb.CurRogue)
	}
	return db.CurRogue
}

func (g *PlayerData) GetRogueRoom() map[uint32]*spb.RogueRoom {
	db := g.GetCurRogue()
	if db.RogueRoomMap == nil {
		db.RogueRoomMap = make(map[uint32]*spb.RogueRoom)
	}
	return db.RogueRoomMap
}

func (g *PlayerData) GetCurRogueRoom() *spb.RogueRoom {
	db := g.GetCurRogue()
	if db.RogueRoomMap == nil {
		db.RogueRoomMap = make(map[uint32]*spb.RogueRoom)
	}
	return db.RogueRoomMap[db.CurSiteId]
}

func (g *PlayerData) GetCurRogueRoomId() uint32 {
	db := g.GetCurRogueRoom()
	if db == nil {
		return 0
	}
	return db.RoomId
}
func (g *PlayerData) GetRoomBySiteId(siteId uint32) *spb.RogueRoom {
	db := g.GetCurRogue()
	if db.RogueRoomMap == nil {
		db.RogueRoomMap = make(map[uint32]*spb.RogueRoom)
	}
	return db.RogueRoomMap[siteId]
}

func (g *PlayerData) FinishRogueRoom(siteId uint32) {
	room := g.GetRoomBySiteId(siteId)
	if room == nil {
		return
	}
	room.RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_FINISH
}

func (g *PlayerData) UpCurRogueRoom(siteId uint32) {
	db := g.GetCurRogue()
	if db.RogueRoomMap[siteId] == nil {
		return
	}
	db.RogueRoomMap[siteId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY
	db.CurSiteId = siteId
}

func (g *PlayerData) GetDbRogueArea(areaId uint32) *spb.RogueArea {
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

func (g *PlayerData) GetRogueBuffNum() uint32 {
	db := g.GetCurRogue()
	if db == nil {
		return 0
	}
	return db.BuffNum
}

func (g *PlayerData) AddRogueBuffNum() {
	db := g.GetCurRogue()
	if db != nil {
		db.BuffNum++
	}
}

func (g *PlayerData) GetRogueBuffList() map[uint32]*spb.RogueBuff {
	db := g.GetCurRogue()
	if db.BuffList == nil {
		db.BuffList = make(map[uint32]*spb.RogueBuff)
	}
	return db.BuffList
}

func (g *PlayerData) GetRogueBuffById(id uint32) *spb.RogueBuff {
	db := g.GetRogueBuffList()
	return db[id]
}

func (g *PlayerData) AddRogueBuff(id uint32) {
	db := g.GetRogueBuffList()
	conf := gdconf.GetBuffById(id)
	if db[id] != nil {
		newLevel := db[id].BuffLevel
		if conf[newLevel+1] != nil {
			db[id].BuffLevel++
		}
	} else {
		db[id] = &spb.RogueBuff{
			BuffId:    id,
			BuffLevel: 1,
		}
	}
}

/**************************************************Buff获取概率计算*******************************************/

func (g *PlayerData) GetRogueInfoOnline() *RogueInfoOnline {
	db := g.GetCurBattle()
	if db.RogueInfoOnline == nil {
		db.RogueInfoOnline = &RogueInfoOnline{}
	}
	return db.RogueInfoOnline
}

func (g *PlayerData) NewGetRogueBuffByType() {
	db := g.GetRogueInfoOnline()
	db.RogueBuffRarityOne = RogueBuffRarityOne
	db.RogueBuffRarityTwo = RogueBuffRarityTwo
	rogueBuffByTypeList := make(map[uint32]*RogueBuffByType, 0)
	buffTypeList := gdconf.GetRogueBuffByType()
	if buffTypeList != nil {
		for typeId, rogueBuffByType := range buffTypeList {
			if typeId == 100 { // 过滤基础
				continue
			}
			rogueBuffRarityList := make(map[uint32]*RogueBuffRarity)
			for rarityId, buffListConf := range rogueBuffByType {
				buffList := make([]uint32, 0)
				for _, buff := range buffListConf {
					// 此处加个判断特殊祝福就行了
					conf := gdconf.GetBuffByIdAndLevel(buff, 1)
					if conf.ActivityModuleID != 0 && conf.ActivityModuleID != g.GetCurRogue().RogueActivityModuleID {
						continue
					}
					buffList = append(buffList, buff)
				}
				rogueBuffRarityList[rarityId] = &RogueBuffRarity{
					Rarity:   rarityId,
					BuffList: buffList,
				}
			}
			rogueBuffByTypeList[typeId] = &RogueBuffByType{
				Weight:          RogueBuffType,
				RogueBuffRarity: rogueBuffRarityList,
			}
		}
	}

	db.RogueBuffByType = rogueBuffByTypeList
}

func (g *PlayerData) GetRogueBuffByType() map[uint32]*RogueBuffByType {
	db := g.GetRogueInfoOnline()
	if db.RogueBuffByType == nil {
		g.NewGetRogueBuffByType()
	}
	return db.RogueBuffByType
}

func (g *PlayerData) GetRogueBuff() uint32 {
	db := g.GetRogueInfoOnline()
	rogueBuffByTypeList := db.RogueBuffByType
	var totalWeight int32 = 0
	for id, rogueBuffByType := range rogueBuffByTypeList {
		if rogueBuffByType.RogueBuffRarity == nil || len(rogueBuffByType.RogueBuffRarity) == 0 {
			continue
		}
		if id == 0 {
			rogueBuffByType.Weight += AddRogueBuffType
		}
		totalWeight += rogueBuffByType.Weight
	}
	if totalWeight == 0 {
		return 600000
	}
	randomWeight := rand.Int31n(totalWeight)
	for _, rogueBuffByType := range rogueBuffByTypeList {
		if rogueBuffByType.RogueBuffRarity == nil || len(rogueBuffByType.RogueBuffRarity) == 0 {
			continue
		}
		if randomWeight <= rogueBuffByType.Weight {
			// 已选定命途属性
			var rarityTotalWeight int32 = 0
			for _, rogueBuffRarity := range rogueBuffByType.RogueBuffRarity {
				var weight int32 = 0
				switch rogueBuffRarity.Rarity {
				case 1:
					weight = db.RogueBuffRarityOne
				case 2:
					weight = db.RogueBuffRarityTwo
				default:
					continue
				}
				rarityTotalWeight += weight
			}
			if rarityTotalWeight == 0 {
				return 600000
			}
			rarityRandomWeight := rand.Int31n(rarityTotalWeight)
			for _, rogueBuffRarity := range rogueBuffByType.RogueBuffRarity {
				if rogueBuffRarity.BuffList == nil || len(rogueBuffRarity.BuffList) == 0 {
					continue
				}
				var weight int32 = 0
				switch rogueBuffRarity.Rarity {
				case 1:
					weight = db.RogueBuffRarityOne
				case 2:
					weight = db.RogueBuffRarityTwo
				default:
					continue
				}
				if rarityRandomWeight <= weight {
					// 已选定稀有属性
					idIndex := rand.Intn(len(rogueBuffRarity.BuffList))
					return rogueBuffRarity.BuffList[idIndex]
				}
				randomWeight -= weight
			}
		}
		randomWeight -= rogueBuffByType.Weight
	}
	return 600000
}

/****************************************************功能***************************************************/

func (g *PlayerData) GetRogueInfo() *proto.RogueInfo {
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

func (g *PlayerData) GetRogueCurrentInfo() *proto.RogueCurrentInfo {
	db := g.GetCurRogue()
	info := &proto.RogueCurrentInfo{
		RogueAeonInfo:   g.GetGameAeonInfo(),
		GameMiracleInfo: g.GetGameMiracleInfo(),
		RogueLineupInfo: g.GetRogueLineupInfo(),
		Status:          proto.RogueStatus(db.Status),
		RoomMap:         g.GetRogueMap(),
		PendingAction:   g.GetRogueCommonPendingAction(),
		IsExploreWin:    db.IsWin,
		ModuleInfo:      &proto.RogueModuleInfo{ModuleIdList: make([]uint32, 0)},
		VirtualItemInfo: g.GetRogueVirtualItem(),
		RogueBuffInfo:   g.GetRogueBuffInfo(),
	}

	return info
}

func (g *PlayerData) GetRogueScoreRewardInfo() *proto.RogueScoreRewardInfo {
	conf := dbconf.GetCurRogue()
	if conf == nil {
		return nil
	}
	db, poolRefreshed := g.GetRogueHistoryById(conf.SeasonId)
	info := &proto.RogueScoreRewardInfo{
		PoolId:                 20 + g.GetWorldLevel(),
		RewardEndTime:          conf.EndTime.Time.Unix(),
		RewardBeginTime:        conf.BeginTime.Time.Unix(),
		PoolRefreshed:          poolRefreshed, // 是否刷新
		HasTakenInitialScore:   false,         // 是否已取得初始分数
		ExploreScore:           db.Score,      // 本期分数
		TakenNormalFreeRowList: make([]uint32, 0),
	}
	return info
}

func (g *PlayerData) GetRogueSeasonInfo() *proto.RogueSeasonInfo {
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

func (g *PlayerData) GetRogueAreaInfo() *proto.RogueAreaInfo {
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
		dbRogueArea := g.GetDbRogueArea(rogueArea)
		RogueArea := &proto.RogueArea{
			AreaId:     dbRogueArea.AreaId,
			AreaStatus: proto.RogueAreaStatus(dbRogueArea.RogueAreaStatus),
			// MapId:           0,
			HasTakenReward: false,
			RogueStatus:    0,
			// CurReachRoomNum: 0,
		}
		info.RogueAreaList = append(info.RogueAreaList, RogueArea)
	}

	return info
}

func (g *PlayerData) GetRogueAeonInfo() *proto.RogueAeonInfo {
	info := &proto.RogueAeonInfo{
		IsUnlocked:             true,
		UnlockedAeonEnhanceNum: 3,
		AeonIdList:             []uint32{1, 2, 3, 4, 5, 6, 7},
		UnlockedAeonNum:        9,
	}

	return info
}

func (g *PlayerData) GetGameAeonInfo() *proto.GameAeonInfo {
	rogue := g.GetCurRogue()
	info := &proto.GameAeonInfo{
		IsUnlocked:             true,
		UnlockedAeonEnhanceNum: 3,
		GameAeonId:             rogue.AeonId,
	}
	return info
}

func (g *PlayerData) GetRogueMap() *proto.RogueMapInfo {
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

func (g *PlayerData) GetRogueLineupInfo() *proto.RogueLineupInfo {
	info := &proto.RogueLineupInfo{
		BaseAvatarIdList: make([]uint32, 0),
		ReviveInfo:       nil,
	}

	lineup := g.GetBattleLineUpById(Rogue)
	if lineup.AvatarIdList != nil {
		for _, avatar := range lineup.AvatarIdList {
			if avatar.AvatarId == 0 {
				continue
			}
			info.BaseAvatarIdList = append(info.BaseAvatarIdList, avatar.AvatarId)
		}
	}

	return info
}

func (g *PlayerData) GetRogueBuffInfo() *proto.RogueBuffInfo {
	info := &proto.RogueBuffInfo{
		MazeBuffList: make([]*proto.RogueBuff, 0),
	}
	for _, in := range g.GetRogueBuffList() {
		info.MazeBuffList = append(info.MazeBuffList, &proto.RogueBuff{
			Level:  in.BuffLevel,
			BuffId: in.BuffId,
		})
	}
	return info
}

func (g *PlayerData) GetRogueVirtualItem() *proto.RogueVirtualItem {
	info := &proto.RogueVirtualItem{
		// Sus:        0,
		// RogueMoney: g.GetMaterialById(Cf),
	}

	return info
}

func (g *PlayerData) GetGameMiracleInfo() *proto.GameMiracleInfo {
	info := &proto.GameMiracleInfo{
		GameMiracleInfo: &proto.RogueMiracleInfo{
			MiracleList: make([]*proto.RogueMiracle, 0),
		},
	}

	return info
}

func (g *PlayerData) GetRogueCommonPendingAction() *proto.RogueCommonPendingAction {
	info := &proto.RogueCommonPendingAction{
		QueuePosition: 0,
		RogueAction:   &proto.RogueAction{},
	}

	return info
}

func (g *PlayerData) GetCurRogueBuff() []*proto.BattleBuff {
	buffList := make([]*proto.BattleBuff, 0)
	db := g.GetRogueBuffList()
	for _, buff := range db {
		buffList = append(buffList, &proto.BattleBuff{
			Id:              buff.BuffId,
			Level:           buff.BuffLevel,
			OwnerIndex:      4294967295,
			WaveFlag:        4294967295,
			TargetIndexList: make([]uint32, 0),
			DynamicValues:   make(map[string]float32),
		})
	}

	return buffList
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
	lineUp := g.GetBattleLineUpById(Rogue)

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
			room := g.GetCurRogueRoom()
			if len(room.NextSiteIdList) == 1 {
				index = 0
			}
			if len(room.NextSiteIdList) > 0 {
				siteId := room.NextSiteIdList[index]
				nextRoom := g.GetRoomBySiteId(siteId)
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
