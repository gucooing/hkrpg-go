package model

import (
	"math/rand"
	"strings"

	"github.com/gucooing/hkrpg-go/dbconf"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func newRogueTourn() *spb.RogueTourn {
	db := &spb.RogueTourn{
		Exp:                    0,
		TakenLevelRewards:      make([]uint32, 0),
		InspirationCircuitInfo: make(map[uint32]*spb.InspirationCircuitInfo),
		UnlockDifficultyInfo:   make([]uint32, 0),
	}
	db.InspirationCircuitInfo[100] = &spb.InspirationCircuitInfo{
		InspirationCircuitId: 100,
		Status:               spb.RogueTalentStatus_ROGUE_TALENT_STATUS_UNLOCK,
	}
	for v := range gdconf.GetRogueTournDifficultyCompMap() {
		db.UnlockDifficultyInfo = append(db.UnlockDifficultyInfo, v)
	}
	return db
}

func (g *PlayerData) GetRogueTourn() *spb.RogueTourn {
	db := g.GetBattle()
	if db.RogueTourn == nil {
		db.RogueTourn = newRogueTourn()
	}
	return db.RogueTourn
}

func (g *PlayerData) GetCurRogueTourn() *spb.CurRogueTourn {
	db := g.GetRogueTourn()
	return db.CurRogueTourn
}

func (g *PlayerData) NewCurRogueTourn(areaId uint32) {
	db := g.GetRogueTourn()
	curRogueTourn := new(spb.CurRogueTourn)
	db.CurRogueTourn = curRogueTourn
	curRogueTourn.AreaId = areaId
	curRogueTourn.CurLayerIndex = 1
	curRogueTourn.CurLayerList = make(map[uint32]*spb.LayerInfo)
	curRogueTourn.FormulaList = make([]uint32, 0)
	// 生成新关卡
	curRogueTourn.CurLayerList[1] = &spb.LayerInfo{
		LayerId:            1101,
		Status:             spb.RogueTournLayerStatus_ROGUE_TOURN_LAYER_STATUS_PROCESSING,
		LayerIndex:         1,
		RogueTournRoomList: newRogueTournRoomInfoList(4),
		CurRoomIndex:       1,
	}
	curRogueTourn.CurLayerList[2] = &spb.LayerInfo{
		LayerId:            1201,
		Status:             0,
		LayerIndex:         2,
		RogueTournRoomList: newRogueTournRoomInfoList(5),
		CurRoomIndex:       1,
	}
	curRogueTourn.CurLayerList[3] = &spb.LayerInfo{
		LayerId:            1301,
		Status:             0,
		LayerIndex:         3,
		RogueTournRoomList: newRogueTournRoomInfoList(4),
		CurRoomIndex:       1,
	}
	// 选择第一关的第一个房间
	curRoom := g.GetCurRogueTournRoom()
	curRoom.RoomId = gdconf.GetRogueTournRoomGenaByType(3).RogueRoomID
	curRoom.Status = spb.RogueTournRoomStatus_ROGUE_TOURN_ROOM_STATUS_PROCESSING
}

func newRogueTournRoomInfoList(num int) map[uint32]*spb.RogueTournRoomInfo {
	list := make(map[uint32]*spb.RogueTournRoomInfo)
	for i := 1; i < num+1; i++ {
		list[uint32(i)] = &spb.RogueTournRoomInfo{RoomIndex: uint32(i)}
	}
	return list
}

func (g *PlayerData) GetCurLayerInfo() *spb.LayerInfo {
	db := g.GetCurRogueTourn()
	if db == nil {
		return nil
	}
	return db.CurLayerList[db.CurLayerIndex]
}

func (g *PlayerData) GetCurRogueTournRoom() *spb.RogueTournRoomInfo {
	db := g.GetCurLayerInfo()
	if db == nil {
		return nil
	}
	return db.RogueTournRoomList[db.CurRoomIndex]
}

func (g *PlayerData) GetCurRogueTournFormula() []uint32 {
	db := g.GetCurRogueTourn()
	if db == nil {
		return make([]uint32, 0)
	}
	if db.FormulaList == nil {
		db.FormulaList = make([]uint32, 0)
	}
	return db.FormulaList
}

func (g *PlayerData) AddCurRogueTournFormula(id uint32) {
	db := g.GetCurRogueTourn()
	if db == nil {
		return
	}
	if db.FormulaList == nil {
		db.FormulaList = make([]uint32, 0)
	}
	db.FormulaList = append(db.FormulaList, id)
}

func (g *PlayerData) UpdateRogueTournEnterRoom(curRoomIndex, nextTypeId uint32) (uint32, uint32) {
	db := g.GetCurRogueTourn()
	curLayer := g.GetCurLayerInfo()
	if db == nil {
		return 0, 0
	}
	curLayer.RogueTournRoomList[curLayer.CurRoomIndex].Status = spb.RogueTournRoomStatus_ROGUE_TOURN_ROOM_STATUS_FINISH
	if curLayer.LayerId == 1101 && curRoomIndex == 4 || // 这两个情况是进入下一关
		curLayer.LayerId == 1201 && curRoomIndex == 5 {
		curLayer.Status = spb.RogueTournLayerStatus_ROGUE_TOURN_LAYER_STATUS_FINISH
		db.CurLayerIndex++
	} else { // 不进入下一关
		curLayer.CurRoomIndex++
	}
	newCurLayer := g.GetCurLayerInfo()
	newCurLayer.Status = spb.RogueTournLayerStatus_ROGUE_TOURN_LAYER_STATUS_PROCESSING
	newCurLayer.RogueTournRoomList[newCurLayer.CurRoomIndex].RoomId = gdconf.GetRogueTournRoomGenaByType(nextTypeId).RogueRoomID
	newCurLayer.RogueTournRoomList[newCurLayer.CurRoomIndex].Status = spb.RogueTournRoomStatus_ROGUE_TOURN_ROOM_STATUS_PROCESSING

	return newCurLayer.LayerIndex, newCurLayer.CurRoomIndex
}

func (g *PlayerData) GetNextRogueTournRoomType(curRoomIndex, curTypeId uint32) uint32 {
	switch curTypeId {
	case 1101:
		if curRoomIndex == 3 {
			return 1
		}
	case 1201:
		if curRoomIndex == 4 {
			return 1
		}
	case 1301:
		if curRoomIndex == 2 {
			return 10
		} else if curRoomIndex == 3 {
			return 1
		} else if curRoomIndex == 4 {
			return 0
		}
	}
	x := []uint32{3, 4, 5, 7, 8}
	return x[rand.Intn(len(x))]
}

/****************************************************功能***************************************************/

func (g *PlayerData) GetRogueTournSeasonInfo() *proto.RogueTournSeasonInfo {
	info := &proto.RogueTournSeasonInfo{
		SubTournId:  1,
		MainTournId: 1,
	}
	return info
}

func (g *PlayerData) GetInspirationCircuitInfo() *proto.RogueTournPermanentTalentInfo {
	info := &proto.RogueTournPermanentTalentInfo{
		TalentInfoList:     &proto.RogueTalentInfoList{TalentInfo: make([]*proto.RogueTalentInfo, 0)},
		TournTalentCoinNum: g.GetMaterialById(Inspiration),
	}
	for v, k := range gdconf.GetRogueTournPermanentTalentMap() {
		status := proto.RogueTalentStatus_ROGUE_TALENT_STATUS_LOCK
		if k.IsImportant {
			status = proto.RogueTalentStatus_ROGUE_TALENT_STATUS_UNLOCK
		}
		status = proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE
		info.TalentInfoList.TalentInfo = append(info.TalentInfoList.TalentInfo, &proto.RogueTalentInfo{
			Status:   status,
			TalentId: v,
		})
	}
	return info
}

func (g *PlayerData) GetExtraScoreInfo() *proto.ExtraScoreInfo {
	conf := dbconf.GetCurRogue()
	if conf == nil {
		return nil
	}
	info := &proto.ExtraScoreInfo{
		Week:     2,
		IsFinish: true,
		EndTime:  conf.EndTime.Time.Unix(),
		// JLDGFGMEMJH: 1000,
	}
	return info
}

func (g *PlayerData) GetRogueTournExpInfo() *proto.RogueTournExpInfo {
	db := g.GetRogueTourn()
	info := &proto.RogueTournExpInfo{
		TakenLevelRewards: db.TakenLevelRewards,
		Exp:               db.Exp,
	}

	return info
}

func (g *PlayerData) GetRogueTournHandbookInfo() *proto.RogueTournHandbookInfo {
	info := &proto.RogueTournHandbookInfo{
		// HandbookFormulaList: make([]uint32, 0),
		HandbookBuffList: make([]uint32, 0),
		// DFPKGDJNMHA:            1,
		// HandbookEventList:      make([]uint32, 0),
		// HandbookAvatarBaseList: make([]uint32, 0),
		// HandbookMiracleList:    make([]uint32, 0),
		// TakeHandbookRewardList: make([]uint32, 0),
	}
	// 添加方程
	// for id := range gdconf.GetRogueTournFormulaMap() {
	// 	info.HandbookFormulaList = append(info.HandbookFormulaList, id)
	// }
	return info
}

func (g *PlayerData) GetRogueTournDifficultyInfo() []*proto.RogueTournDifficultyInfo {
	info := make([]*proto.RogueTournDifficultyInfo, 0)
	for _, id := range []uint32{10101, 10102, 10103, 10104, 10105, 10106} {
		info = append(info, &proto.RogueTournDifficultyInfo{
			DifficultyId: id,
			IsUnlocked:   true,
		})
	}

	return info
}

func (g *PlayerData) GetRogueTournAreaInfo() []*proto.RogueTournAreaInfo {
	info := make([]*proto.RogueTournAreaInfo, 0)
	for _, id := range []uint32{101, 201, 202, 203, 204, 205, 1011, 1012, 1013, 1014, 1015} {
		info = append(info, &proto.RogueTournAreaInfo{
			AreaId:                      id,
			IsTournFinish:               true,
			IsTakenReward:               false,
			IsUnlocked:                  true,
			UnlockedTournDifficultyList: make([]uint32, 0),
		})
	}
	return info
}

func (g *PlayerData) GetRogueTournCurInfo() *proto.RogueTournCurInfo {
	curRogueTourn := g.GetCurRogueTourn()
	if curRogueTourn == nil {
		return nil
	}
	info := &proto.RogueTournCurInfo{
		RogueTournCurAreaInfo: g.GetRogueTournCurAreaInfo(),

		RogueTournCurGameInfo: &proto.RogueTournCurGameInfo{
			RogueTournGameAreaInfo: &proto.RogueTournGameAreaInfo{
				GameAreaId: curRogueTourn.AreaId,
			},
			MiracleInfo:      g.GetChessRogueMiracleInfo(),
			TournFormulaInfo: g.GetRogueTournFormulaInfo(),
			ItemValue: &proto.RogueGameItemValue{
				VirtualItem: map[uint32]uint32{Cf: g.GetMaterialById(Cf)},
			},
			Level: g.GetRogueTournLayerInfo(),
			Lineup: &proto.RogueTournLineupInfo{
				RogueReviveCost: &proto.ItemCostData{ItemList: []*proto.ItemCost{
					{
						PileItem: &proto.PileItem{
							ItemId:  Cf,
							ItemNum: g.GetMaterialById(Cf),
						},
					},
				}},
			},
			Buff:        g.GetRogueTournBuffInfo(),
			UnlockValue: g.GetKeywordUnlockInfo(),
		},
	}
	return info
}

func (g *PlayerData) GetRogueTournCurAreaInfo() *proto.RogueTournCurAreaInfo {
	curRogueTourn := g.GetCurRogueTourn()
	if curRogueTourn == nil {
		return nil
	}
	info := &proto.RogueTournCurAreaInfo{
		RogueSubMode: 301,
		SubAreaId:    curRogueTourn.AreaId,
	}
	return info
}

func (g *PlayerData) GetKeywordUnlockInfo() *proto.KeywordUnlockValue {
	info := &proto.KeywordUnlockValue{KeywordUnlockMap: map[uint32]bool{
		1615010: false,
		1615110: false,
		1615210: false,
		1615310: false,
	}}
	return info
}

func (g *PlayerData) GetChessRogueMiracleInfo() *proto.ChessRogueMiracleInfo {
	info := &proto.ChessRogueMiracleInfo{
		ChessRogueMiracleInfo: &proto.ChessRogueMiracle{
			MiracleList: make([]*proto.GameRogueMiracle, 0),
		},
	}

	return info
}

func (g *PlayerData) GetRogueTournLayerInfo() *proto.RogueTournLevelInfo {
	curRogueTourn := g.GetCurRogueTourn()
	curLayer := g.GetCurLayerInfo()
	if curRogueTourn == nil || curLayer == nil {
		return nil
	}
	info := &proto.RogueTournLevelInfo{
		Status:        proto.RogueTournLevelStatus(curLayer.Status),
		LevelInfoList: make([]*proto.RogueTournLevel, 0),
		Reason:        0,
		CurLevelIndex: curRogueTourn.CurLayerIndex,
	}
	for _, layerInfo := range curRogueTourn.CurLayerList {
		layerInfoList := &proto.RogueTournLevel{
			LayerId:       layerInfo.LayerId,
			Status:        proto.RogueTournLayerStatus(layerInfo.Status),
			CurRoomIndex:  layerInfo.CurRoomIndex,
			TournRoomList: make([]*proto.RogueTournRoomList, 0),
			LevelIndex:    layerInfo.LayerIndex,
		}
		for _, roomInfo := range layerInfo.RogueTournRoomList {
			layerInfoList.TournRoomList = append(layerInfoList.TournRoomList, &proto.RogueTournRoomList{
				Status:    proto.RogueTournRoomStatus(roomInfo.Status),
				RoomIndex: roomInfo.RoomIndex,
				RoomId:    roomInfo.RoomId,
			})
		}
		info.LevelInfoList = append(info.LevelInfoList, layerInfoList)
	}

	return info
}

func (g *PlayerData) GetRogueTournBuffInfo() *proto.ChessRogueBuffInfo {
	info := &proto.ChessRogueBuffInfo{
		ChessRogueBuffInfo: &proto.ChessRogueBuff{
			BuffList: make([]*proto.RogueCommonBuff, 0),
		},
	}
	return info
}

func (g *PlayerData) GetRogueTournFormulaInfo() *proto.RogueTournFormulaInfo {
	db := g.GetCurRogueTournFormula()
	info := &proto.RogueTournFormulaInfo{
		FormulaTypeValue: &proto.FormulaTypeValue{
			FormulaTypeMap: make(map[uint32]int32),
		},
	}
	for _, id := range db {
		conf := gdconf.GetRogueTournFormulaById(id)
		if conf == nil {
			continue
		}
		formulaInfo := &proto.FormulaInfo{
			IsExpand:  false,
			FormulaId: id,
			FormulaBuffTypeList: []*proto.FormulaBuffTypeInfo{
				{
					FormulaBuffNum: conf.MainBuffNum,
					Key:            conf.MainBuffTypeID,
				},
				{
					FormulaBuffNum: conf.SubBuffNum,
					Key:            conf.SubBuffTypeID,
				},
			},
		}
		info.GameFormulaInfo = append(info.GameFormulaInfo, formulaInfo)
	}
	return info
}

func (g *PlayerData) GetRogueMapRotateInfo(roomId uint32) *proto.RogueMapRotateInfo {
	info := &proto.RogueMapRotateInfo{
		RotaterDataList: make([]*proto.RotaterData, 0),
		ChargerInfo:     make([]*proto.ChargerInfo, 0),
		EnergyInfo:      &proto.RotaterEnergyInfo{},
	}
	roomConf := gdconf.GetRogueTournRoomGenById(roomId)
	if roomConf == nil {
		return info
	}
	if roomConf.RotateInfo.IsRotate {
		// info.IsRotate = roomConf.RotateInfo.IsRotate
		// info.HMCAFEJAPJK = roomConf.RotateInfo.RotateNum
		// info.MapInfo = &proto.IJJHKDNFKMD{
		// 	BIKIIIKJIIG: &proto.Vector{},
		// 	EFGOCIAIKMN: &proto.AEKLIMBAKCL{
		// 		Z:           -0.70710677,
		// 		Y:           0,
		// 		X:           0,
		// 		EAGOBFLBPFN: 0.70710677,
		// 	},
		// }
	}

	return info
}

func (g *PlayerData) GetRogueTournScene(roomId uint32) *proto.SceneInfo {
	roomConf := gdconf.GetRogueTournRoomGenById(roomId)
	if roomConf == nil {
		return nil
	}
	mapEntrance := gdconf.GetMapEntranceById(roomConf.MapEntrance)
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
		EntryId:            roomConf.MapEntrance,
		GameModeType:       17, // gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}
	// 获取场景实体
	var avatarGroupID uint32 = 0
	for _, groupID := range roomConf.GroupWithContent {
		sceneGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, groupID)
		if sceneGroup == nil {
			continue
		}
		if len(sceneGroup.AnchorList) > 0 {
			avatarGroupID = groupID
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
		g.GetRogueTournPropByID(entityGroupLists, sceneGroup)
		// 添加怪物实体
		g.GetRogueTournNPCMonsterByID(entityGroupLists, sceneGroup, roomConf.NpcMonster[groupID])
		// 添加NPC实体
		g.GetNPCByID(entityGroupLists, sceneGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}
	// 添加队伍实体
	entityGroupList := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	startGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, avatarGroupID)
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
	lineUp := g.GetBattleLineUpById(RogueTourn)

	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroupList, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)

	return scene
}

func (g *PlayerData) GetRogueTournNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, npcMonster map[uint32]*gdconf.RogueTournMonsterInfo) {
	for _, monsterList := range sceneGroup.MonsterList {
		info := npcMonster[monsterList.ID]
		if info == nil {
			info = &gdconf.RogueTournMonsterInfo{ // 默认怪物
				RogueMonsterID: 3003221,
				NpcMonsterID:   3003052,
				EventID:        83003221,
			}
		}
		entityId := g.GetNextGameObjectGuid()
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
					ExtraInfo: &proto.NpcMonsterExtraInfo{
						RogueGameInfo: &proto.NpcMonsterRogueInfo{
							RogueMonsterId: info.RogueMonsterID,
						},
					},
					MonsterId: info.NpcMonsterID,
					EventId:   info.EventID,
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
			EventID: info.EventID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}

func (g *PlayerData) GetRogueTournPropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) {
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
		// 3:战斗 5:事件 8:奖励
		db := g.GetCurLayerInfo()
		if strings.Contains(propList.Name, "Door") {
			roomType := g.GetNextRogueTournRoomType(db.CurRoomIndex, db.LayerId)
			var propId uint32 = 0
			switch roomType {
			case 0:
				propId = 1033 // 白色
			case 5:
				propId = 1035 // 蓝色
			case 10:
				propId = 1036 // 绿色
			default:
				propId = 1034 // 红色
			}
			prop.PropState = 1
			prop.PropId = propId // 颜色
			prop.ExtraInfo = &proto.PropExtraInfo{
				RogueTournDoorInfo: &proto.RogueTournDoorInfo{
					RogueDoorNextRoomType: roomType,
				},
			}
		}
		entityList.EntityOneofCase = &proto.SceneEntityInfo_Prop{
			Prop: prop,
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}
