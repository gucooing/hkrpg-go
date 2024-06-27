package player

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
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

func (g *GamePlayer) GetRogueTourn() *spb.RogueTourn {
	db := g.GetBattle()
	if db.RogueTourn == nil {
		db.RogueTourn = newRogueTourn()
	}
	return db.RogueTourn
}

func (g *GamePlayer) GetCurRogueTourn() *spb.CurRogueTourn {
	db := g.GetRogueTourn()
	return db.CurRogueTourn
}

func (g *GamePlayer) NewCurRogueTourn(areaId uint32) {
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
		CurRoomIndex:       0,
	}
	curRogueTourn.CurLayerList[3] = &spb.LayerInfo{
		LayerId:            1301,
		Status:             0,
		LayerIndex:         3,
		RogueTournRoomList: newRogueTournRoomInfoList(4),
		CurRoomIndex:       0,
	}
	// 选择第一关的第一个房间
	curRoom := g.GetCurRogueTournRoom()
	curRoom.RoomId = 11036011
	curRoom.Status = spb.RogueTournRoomStatus_ROGUE_TOURN_ROOM_STATUS_PROCESSING
}

func newRogueTournRoomInfoList(num int) map[uint32]*spb.RogueTournRoomInfo {
	list := make(map[uint32]*spb.RogueTournRoomInfo)
	for i := 0; i < num; i++ {
		list[uint32(i)] = &spb.RogueTournRoomInfo{RoomIndex: uint32(i)}
	}
	return list
}

func (g *GamePlayer) GetCurLayerInfo() *spb.LayerInfo {
	db := g.GetCurRogueTourn()
	if db == nil {
		return nil
	}
	return db.CurLayerList[db.CurLayerIndex]
}

func (g *GamePlayer) GetCurRogueTournRoom() *spb.RogueTournRoomInfo {
	db := g.GetCurLayerInfo()
	if db == nil {
		return nil
	}
	return db.RogueTournRoomList[db.CurRoomIndex]
}

func (g *GamePlayer) GetCurRogueTournFormula() []uint32 {
	db := g.GetCurRogueTourn()
	if db == nil {
		return make([]uint32, 0)
	}
	if db.FormulaList == nil {
		db.FormulaList = make([]uint32, 0)
	}
	return db.FormulaList
}

func (g *GamePlayer) AddCurRogueTournFormula(id uint32) {
	db := g.GetCurRogueTourn()
	if db == nil {
		return
	}
	if db.FormulaList == nil {
		db.FormulaList = make([]uint32, 0)
	}
	db.FormulaList = append(db.FormulaList, id)
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetRogueTournSeasonInfo() *proto.RogueTournSeasonInfo {
	info := &proto.RogueTournSeasonInfo{
		SubTournId:  1,
		MainTournId: 1,
	}
	return info
}

func (g *GamePlayer) GetInspirationCircuitInfo() *proto.InspirationCircuitInfo {
	info := &proto.InspirationCircuitInfo{
		RogueTalentInfo: &proto.RogueTalentInfo{RogueTalentList: make([]*proto.RogueTalent, 0)},
		InspirationNum:  g.GetMaterialById(Inspiration),
	}
	for v, k := range gdconf.GetRogueTournPermanentTalentMap() {
		status := proto.RogueTalentStatus_ROGUE_TALENT_STATUS_LOCK
		if k.IsImportant {
			status = proto.RogueTalentStatus_ROGUE_TALENT_STATUS_UNLOCK
		}
		status = proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE
		info.RogueTalentInfo.RogueTalentList = append(info.RogueTalentInfo.RogueTalentList, &proto.RogueTalent{
			Status:   status,
			TalentId: v,
		})
	}
	return info
}

func (g *GamePlayer) GetExtraScoreInfo() *proto.ExtraScoreInfo {
	conf := database.GetCurRogue()
	info := &proto.ExtraScoreInfo{
		Week:     2,
		IsFinish: true,
		EndTime:  conf.EndTime.Time.Unix(),
		// JLDGFGMEMJH: 1000,
	}
	return info
}

func (g *GamePlayer) GetRogueTournExpInfo() *proto.RogueTournExpInfo {
	db := g.GetRogueTourn()
	info := &proto.RogueTournExpInfo{
		TakenLevelRewards: db.TakenLevelRewards,
		Exp:               db.Exp,
	}

	return info
}

func (g *GamePlayer) GetRogueTournCollectionInfo() *proto.RogueTournCollectionInfo {
	info := &proto.RogueTournCollectionInfo{
		OBFNIDGAFMN: make([]uint32, 0),
		KJHPDANECOM: make([]uint32, 0),
		EOJECMKIABF: 1,
		MFMLAPAONCM: make([]uint32, 0),
		PHNBGLOFFJM: make([]uint32, 0),
		JOAHOHIPAAG: make([]uint32, 0),
	}
	return info
}

func (g *GamePlayer) GetRogueTournDifficultyInfo() []*proto.RogueTournDifficultyInfo {
	info := make([]*proto.RogueTournDifficultyInfo, 0)
	for _, id := range []uint32{10101, 10102, 10103, 10104, 10105, 10106} {
		info = append(info, &proto.RogueTournDifficultyInfo{
			DifficultyId: id,
			IsUnlock:     true,
		})
	}

	return info
}

func (g *GamePlayer) GetRogueTournAreaInfo() []*proto.RogueTournAreaInfo {
	info := make([]*proto.RogueTournAreaInfo, 0)
	for _, id := range []uint32{101, 201, 202, 203, 204, 205, 1011, 1012, 1013, 1014, 1015} {
		info = append(info, &proto.RogueTournAreaInfo{
			AreaId:                      id,
			IsFinish:                    true,
			IsTakenReward:               false,
			IsUnlock:                    true,
			UnlockedTournDifficultyList: make([]uint32, 0),
		})
	}
	return info
}

func (g *GamePlayer) GetRogueTournCurInfo() *proto.RogueTournCurInfo {
	curRogueTourn := g.GetCurRogueTourn()
	if curRogueTourn == nil {
		return nil
	}
	info := &proto.RogueTournCurInfo{
		RogueTournCurAreaInfo: g.GetRogueTournCurAreaInfo(),
		RogueTournCurGameInfo: &proto.RogueTournCurGameInfo{
			RogueTournGameAreaInfo: &proto.RogueTournGameAreaInfo{
				AreaId: curRogueTourn.AreaId,
			},
			RogueTournMiracleInfo: g.GetChessRogueMiracleInfo(),
			RogueTournFormulaInfo: g.GetRogueTournFormulaInfo(),
			RogueTournValuesItem: &proto.RogueTournValuesItem{
				VirtualItem: map[uint32]uint32{Cf: g.GetMaterialById(Cf)},
			},
			RogueTournLayerInfo: g.GetRogueTournLayerInfo(),
			RogueTournVirtualItem: &proto.RogueTournVirtualItem{
				GameItemInfo: &proto.ItemCostData{ItemList: []*proto.ItemCost{
					{
						ItemOneofCase: &proto.ItemCost_PileItem{
							PileItem: &proto.PileItem{
								ItemId:  Cf,
								ItemNum: g.GetMaterialById(Cf),
							},
						},
					},
				}},
			},
			RogueTournBuffInfo: g.GetRogueTournBuffInfo(),
			KeywordUnlockInfo:  g.GetKeywordUnlockInfo(),
			DLCLNIJBHBD:        nil,
			ENGCMKFPKLH:        nil,
		},
	}
	return info
}

func (g *GamePlayer) GetRogueTournCurAreaInfo() *proto.RogueTournCurAreaInfo {
	curRogueTourn := g.GetCurRogueTourn()
	if curRogueTourn == nil {
		return nil
	}
	info := &proto.RogueTournCurAreaInfo{
		RogueSubMode: 301,
		AreaId:       curRogueTourn.AreaId,
	}
	return info
}

func (g *GamePlayer) GetKeywordUnlockInfo() *proto.KeywordUnlockInfo {
	info := &proto.KeywordUnlockInfo{KeywordUnlockMap: map[uint32]bool{
		1615010: false,
		1615110: false,
		1615210: false,
		1615310: false,
	}}
	return info
}

func (g *GamePlayer) GetChessRogueMiracleInfo() *proto.ChessRogueMiracleInfo {
	info := &proto.ChessRogueMiracleInfo{
		MiracleInfo: &proto.ChessRogueMiracle{
			MiracleList: make([]*proto.GameRogueMiracle, 0),
		},
	}

	return info
}

func (g *GamePlayer) GetRogueTournLayerInfo() *proto.RogueTournLayerInfo {
	curRogueTourn := g.GetCurRogueTourn()
	curLayer := g.GetCurLayerInfo()
	if curRogueTourn == nil || curLayer == nil {
		return nil
	}
	info := &proto.RogueTournLayerInfo{
		Status:        proto.RogueTournLevelStatus(curLayer.Status),
		LayerInfoList: make([]*proto.LayerInfoList, 0),
		Reason:        0,
		CurLayerIndex: curRogueTourn.CurLayerIndex,
	}
	for _, layerInfo := range curRogueTourn.CurLayerList {
		layerInfoList := &proto.LayerInfoList{
			LayerId:               layerInfo.LayerId,
			RogueTournLayerStatus: proto.RogueTournLayerStatus(layerInfo.Status),
			CurRoomIndex:          layerInfo.CurRoomIndex,
			RogueTournRoomList:    make([]*proto.RogueTournRoomList, 0),
			LayerIndex:            layerInfo.LayerIndex,
		}
		for _, roomInfo := range layerInfo.RogueTournRoomList {
			layerInfoList.RogueTournRoomList = append(layerInfoList.RogueTournRoomList, &proto.RogueTournRoomList{
				RogueTournRoomStatus: proto.RogueTournRoomStatus(roomInfo.Status),
				RoomIndex:            roomInfo.RoomIndex,
				RoomId:               roomInfo.RoomId,
			})
		}
		info.LayerInfoList = append(info.LayerInfoList, layerInfoList)
	}

	return info
}

func (g *GamePlayer) GetRogueTournBuffInfo() *proto.RogueTournBuffInfo {
	info := &proto.RogueTournBuffInfo{
		RogueTournMazeBuffInfo: &proto.RogueTournMazeBuffInfo{
			BuffList: make([]*proto.RogueCommonBuff, 0),
		},
	}
	return info
}

func (g *GamePlayer) GetRogueTournFormulaInfo() *proto.RogueTournFormulaInfo {
	db := g.GetCurRogueTournFormula()
	info := &proto.RogueTournFormulaInfo{
		FormulaTypeInfo: &proto.FormulaTypeInfo{
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
					Num:        conf.MainBuffNum,
					BuffTypeId: conf.MainBuffTypeID,
				},
				{
					Num:        conf.SubBuffNum,
					BuffTypeId: conf.SubBuffTypeID,
				},
			},
		}
		info.GameFormulaInfo = append(info.GameFormulaInfo, formulaInfo)
	}
	return info
}

func (g *GamePlayer) GetRogueTournScene(roomId uint32) *proto.SceneInfo {
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
		EntryId:            roomConf.RogueRoomType,
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
	lineUp := g.GetBattleLineUpById(Rogue)

	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroupList, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)

	return scene
}

func (g *GamePlayer) GetRogueTournNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, npcMonster map[uint32]*gdconf.RogueTournMonsterInfo) {
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
			NpcMonster: &proto.SceneNpcMonsterInfo{
				ExtraInfo: &proto.NpcMonsterExtraInfo{RogueInfo: &proto.NpcMonsterRogueInfo{
					RogueMonsterId: info.RogueMonsterID,
				}},
				MonsterId: info.NpcMonsterID,
				EventId:   info.EventID,
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

func (g *GamePlayer) GetRogueTournPropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) {
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
		// 3:战斗 5:事件 8:奖励
		switch propList.PropID {
		case 1033:
			entityList.Prop.PropId = 1034
			entityList.Prop.PropState = 1
			entityList.Prop.ExtraInfo = &proto.PropExtraInfo{
				InfoOneofCase: &proto.PropExtraInfo_RogueTournDoorInfo{
					RogueTournDoorInfo: &proto.RogueTournDoorInfo{
						RogueTournRoomType: 8,
					},
				},
			}

		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}
