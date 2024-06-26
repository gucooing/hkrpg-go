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

func (g *GamePlayer) Get() {

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
		Week:        2,
		IsFinish:    true,
		EndTime:     conf.EndTime.Time.Unix(),
		JLDGFGMEMJH: 1000,
	}
	return info
}

func (g *GamePlayer) GetRogueTournExpInfo() *proto.RogueTournExpInfo {
	info := &proto.RogueTournExpInfo{
		TakenLevelRewards: make([]uint32, 0),
		Exp:               800,
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
	info := &proto.RogueTournCurInfo{
		RogueTournCurAreaInfo: &proto.RogueTournCurAreaInfo{
			FMCJCLEJCEJ: 84,
			PNKJCLDGFFP: 3,
			PendingAction: &proto.RogueCommonPendingAction{
				QueuePosition: 5,
				RogueAction: &proto.RogueAction{
					Action: &proto.RogueAction_RogueFormulaSelectInfo{
						RogueFormulaSelectInfo: &proto.RogueFormulaSelectInfo{
							SelectFormulaIdListFieldNumber: make([]uint32, 0), // []uint32{130906, 130809, 130408},
						},
					},
				},
			},
			RogueSubMode: 301,
			AreaId:       201,
		},
		RogueTournCurGameInfo: &proto.RogueTournCurGameInfo{
			RogueTournGameAreaInfo: &proto.RogueTournGameAreaInfo{
				AreaId: 201,
			},
			RogueTournMiracleInfo: g.GetChessRogueMiracleInfo(),
			RogueTournFormulaInfo: g.GetRogueTournFormulaInfo(),
			RogueTournValuesItem: &proto.RogueTournValuesItem{
				VirtualItem: map[uint32]uint32{31: 100},
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
			KeywordUnlockInfo: &proto.KeywordUnlockInfo{KeywordUnlockMap: map[uint32]bool{
				1615010: false,
				1615110: false,
				1615210: false,
				1615310: false,
			}},
			DLCLNIJBHBD: nil,
			ENGCMKFPKLH: nil,
		},
	}
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
	info := &proto.RogueTournLayerInfo{
		Status:        proto.RogueTournLevelStatus_ROGUE_TOURN_LEVEL_STATUS_PROCESSING,
		LayerInfoList: make([]*proto.LayerInfoList, 0),
		Reason:        0,
		CurLayerIndex: 1,
	}
	info.LayerInfoList = append(info.LayerInfoList, &proto.LayerInfoList{
		LayerId:               1101,
		RogueTournLayerStatus: proto.RogueTournLayerStatus_ROGUE_TOURN_LAYER_STATUS_PROCESSING,
		CurRoomIndex:          1,
		RogueTournRoomList: []*proto.RogueTournRoomList{
			{
				RogueTournRoomStatus: proto.RogueTournRoomStatus_ROGUE_TOURN_ROOM_STATUS_PROCESSING,
				RoomIndex:            1,
				RoomId:               21037011,
			},
		},
		LayerIndex: 1,
	})
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
	info := &proto.RogueTournFormulaInfo{
		GameFormulaInfo: []*proto.FormulaInfo{
			{
				IsExpand:  false,
				FormulaId: 120102,
				FormulaBuffTypeList: []*proto.FormulaBuffTypeInfo{
					{
						Num:        3,
						BuffTypeId: 120,
					},
					{
						Num:        2,
						BuffTypeId: 121,
					},
				},
			},
		},
		FormulaTypeInfo: &proto.FormulaTypeInfo{
			FormulaTypeMap: make(map[uint32]int32),
		},
	}
	return info
}

func (g *GamePlayer) GetRogueTournScene(entryId uint32) *proto.SceneInfo {
	mapEntrance := gdconf.GetMapEntranceById(entryId)
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
		EntryId:            entryId,
		GameModeType:       17, // gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
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
	startGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, 32)
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

	for _, groupID := range []uint32{31, 32, 37, 38, 39} {
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
		g.GetRogueTournPropByID(entityGroupLists, sceneGroup)
		// 添加怪物实体
		g.GetRogueTournNPCMonsterByID(entityGroupLists, sceneGroup)
		// 添加NPC实体
		g.GetNPCByID(entityGroupLists, sceneGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}

	return scene
}

func (g *GamePlayer) GetRogueTournNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) {
	for _, monsterList := range sceneGroup.MonsterList {
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
				MonsterId: 2002010,
				EventId:   83000022,
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
			EventID: 83000022,
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
