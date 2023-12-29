package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 当前坐标通知
func (g *Game) SceneEntityMoveScNotify() {
	pos := g.Player.Pos
	rot := g.Player.Rot

	notify := &proto.SceneEntityMoveScNotify{
		EntryId:          g.Player.DbScene.EntryId,
		ClientPosVersion: 0,
		Motion: &proto.MotionInfo{
			Pos: &proto.Vector{
				X: int32(pos.X),
				Y: int32(pos.Y),
				Z: int32(pos.Z),
			},
			Rot: &proto.Vector{
				X: int32(rot.X),
				Y: int32(rot.Y),
				Z: int32(rot.Z),
			},
		},
	}

	g.Send(cmd.SceneEntityMoveScNotify, notify)
}

// 通知客户端进入场景
func (g *Game) EnterSceneByServerScNotify(entryId, teleportId uint32) {
	rsp := new(proto.EnterSceneByServerScNotify)
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))
	if mapEntrance == nil {
		return
	}
	foorMap := gdconf.GetFloorById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return
	}

	var groupID = mapEntrance.StartGroupID
	var anchorID = mapEntrance.StartAnchorID
	entityMap := make(map[uint32]*EntityList) // [实体id]怪物群id

	if teleportId != 0 {
		groupID = foorMap.Teleports[teleportId].AnchorGroupID
		anchorID = foorMap.Teleports[teleportId].AnchorID
	} else if anchorID == 0 {
		groupID = foorMap.StartGroupID
		anchorID = foorMap.StartAnchorID
	}

	// 获取队伍
	lineUp := g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp]
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           g.Player.DbLineUp.MainLineUp,
		ExtraLineupType: proto.ExtraLineupType_LINEUP_NONE,
		MaxMp:           5,
		Mp:              5,
		Name:            lineUp.Name,
		PlaneId:         0,
	}
	for slot, avatarId := range lineUp.AvatarIdList {
		if avatarId == 0 {
			continue
		}
		avatar := g.Player.DbAvatar.Avatar[avatarId]
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: avatar.Type,
			Slot:       uint32(slot),
			Satiety:    0,
			Hp:         avatar.Hp,
			Id:         avatarId,
			SpBar: &proto.SpBarInfo{
				CurSp: avatar.SpBar.CurSp,
				MaxSp: avatar.SpBar.MaxSp,
			},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsp.Lineup = lineupList

	rsp.Scene = &proto.SceneInfo{
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).WorldID,
		LeaderEntityId:     leaderEntityId,
		FloorId:            mapEntrance.FloorID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).PlaneType),
		PlaneId:            mapEntrance.PlaneID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
	}

	for i := uint32(0); i < 100; i++ {
		rsp.Scene.LightenSectionList = append(rsp.Scene.LightenSectionList, i)
	}

	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 添加队伍角色进实体列表，并设置坐标
	for _, anchor := range foorMap.Groups[groupID].AnchorList {
		if anchor.ID == anchorID {
			for id, avatarid := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
				if avatarid == 0 {
					continue
				}
				entityId := uint32(g.GetNextGameObjectGuid())
				entityList := &proto.SceneEntityInfo{
					EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
						AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
						BaseAvatarId: avatarid,
					}},
					Motion: &proto.MotionInfo{
						Pos: &proto.Vector{
							X: int32(anchor.PosX * 1000),
							Y: int32(anchor.PosY * 1000),
							Z: int32(anchor.PosZ * 1000),
						},
						Rot: &proto.Vector{
							X: int32(anchor.RotX * 1000),
							Y: int32(anchor.RotY * 1000),
							Z: int32(anchor.RotZ * 1000),
						},
					},
				}
				// 为进入场景的角色设置与上面相同的实体id
				if id == 0 {
					entityList.EntityId = leaderEntityId
					entityMap[leaderEntityId] = &EntityList{
						Entity:  avatarid,
						GroupId: groupID,
					}
				} else {
					entityList.EntityId = entityId
					entityMap[entityId] = &EntityList{
						Entity:  avatarid,
						GroupId: groupID,
					}
				}
				entityGroup.EntityList = append(entityGroup.EntityList, entityList)
			}

			// TODO 数据保存问题？
			g.Player.Pos = &Vector{
				X: int(anchor.PosX * 1000),
				Y: int(anchor.PosY * 1000),
				Z: int(anchor.PosZ * 1000),
			}
			g.Player.Rot = &Vector{
				X: int(anchor.RotX * 1000),
				Y: int(anchor.RotY * 1000),
				Z: int(anchor.RotZ * 1000),
			}
			g.Player.DbScene.EntryId = entryId
			break
		}
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	// 获取场景实体
	for _, levelGroup := range foorMap.Groups {

		if levelGroup.GroupName == "RogueBase_Common" {
			continue
		}

		rsp.Scene.GroupIdList = append(rsp.Scene.GroupIdList, levelGroup.GroupId)

		sceneGroupState := &proto.SceneGroupState{
			GroupId:   levelGroup.GroupId,
			IsDefault: true,
		}

		rsp.Scene.GroupStateList = append(rsp.Scene.GroupStateList, sceneGroupState)

		entityGroupList := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加物品实体
		for _, propList := range levelGroup.PropList {
			entityList := &proto.SceneEntityInfo{
				GroupId:  levelGroup.GroupId, // 文件名后那个G
				InstId:   propList.ID,        // ID
				EntityId: uint32(g.GetNextGameObjectGuid()),
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: int32(propList.PosX * 1000),
						Y: int32(propList.PosY * 1000),
						Z: int32(propList.PosZ * 1000),
					},
					Rot: &proto.Vector{
						X: 0,
						Y: int32(propList.RotY * 1000),
						Z: 0,
					},
				},
				EntityCase: &proto.SceneEntityInfo_Prop{Prop: &proto.ScenePropInfo{
					PropId:    propList.PropID, // PropID
					PropState: gdconf.GetStateValue(propList.State),
				}},
			}
			entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
		}
		// 添加怪物实体
		for _, monsterList := range levelGroup.MonsterList {
			entityId := uint32(g.GetNextGameObjectGuid())
			entityList := &proto.SceneEntityInfo{
				GroupId:  levelGroup.GroupId,
				InstId:   monsterList.ID,
				EntityId: entityId,
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: int32(monsterList.PosX * 1000),
						Y: int32(monsterList.PosY * 1000),
						Z: int32(monsterList.PosZ * 1000),
					},
					Rot: &proto.Vector{
						X: 0,
						Y: int32(monsterList.RotY * 1000),
						Z: 0,
					},
				},
				EntityCase: &proto.SceneEntityInfo_NpcMonster{NpcMonster: &proto.SceneNpcMonsterInfo{
					WorldLevel: g.Player.WorldLevel,
					MonsterId:  monsterList.NPCMonsterID,
					EventId:    monsterList.EventID,
				}},
			}
			entityMap[entityId] = &EntityList{
				Entity:  monsterList.EventID,
				GroupId: levelGroup.GroupId,
			}
			entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
		}
		/*
			// 添加NPC实体
			for _, npcList := range levelGroup.NPCList {
				entityList := &proto.SceneEntityInfo{
					GroupId:  levelGroup.GroupId,
					InstId:   npcList.ID,
					EntityId: uint32(g.GetNextGameObjectGuid()),
					Motion: &proto.MotionInfo{
						Pos: &proto.Vector{
							X: int32(npcList.PosX * 1000),
							Y: int32(npcList.PosY * 1000),
							Z: int32(npcList.PosZ * 1000),
						},
						Rot: &proto.Vector{
							X: 0,
							Y: int32(npcList.RotY * 1000),
							Z: 0,
						},
					},
					EntityCase: &proto.SceneEntityInfo_Npc{Npc: &proto.SceneNpcInfo{
						NpcId: npcList.NPCID,
					}},
				}
				entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
			}
		*/
		if len(entityGroupList.EntityList) == 0 {
			continue
		}
		rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroupList)
	}

	g.Player.EntityList = entityMap
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *Game) GetRogueScoreRewardInfoCsReq() {
	rsp := new(proto.GetRogueScoreRewardInfoScRsp)
	rsp.ScoreRewardInfo = &proto.RogueScoreRewardInfo{
		HasTakenInitialScore: true,
		PoolRefreshed:        true,
		PoolId:               22,
	}

	g.Send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func (g *Game) HandleGetEnteredSceneCsReq(payloadMsg []byte) {
	rsp := new(proto.GetEnteredSceneScRsp)
	enteredSceneInfo := &proto.EnteredSceneInfo{
		FloorId: 20001001,
		PlaneId: 20001,
	}
	rsp.EnteredSceneInfo = []*proto.EnteredSceneInfo{enteredSceneInfo}

	g.Send(cmd.GetEnteredSceneScRsp, rsp)
}

// 客户端登录需要的包，不是传送的通知包
func (g *Game) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(g.Player.DbScene.EntryId)))
	entityMap := make(map[uint32]*EntityList)

	rsp := new(proto.GetCurSceneInfoScRsp)
	pos := g.Player.Pos
	rot := g.Player.Rot
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	rsp.Scene = &proto.SceneInfo{
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).WorldID,                        // 世界id
		LeaderEntityId:     leaderEntityId,                                                                                 // 进入场景的角色实体id
		FloorId:            mapEntrance.FloorID,                                                                            // 上面表查询到的，对应文件名中F开头后面的数字
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).PlaneType), // 未知
		PlaneId:            mapEntrance.PlaneID,                                                                            // 上面表查询到的，对应文件名中P开头后面的数字
		EntryId:            g.Player.DbScene.EntryId,                                                                       // 应该是具体的场景条目
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),                                                         // 实体列表
		GroupIdList:        make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
	}

	for i := uint32(0); i < 100; i++ {
		rsp.Scene.LightenSectionList = append(rsp.Scene.LightenSectionList, i)
	}

	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 将进入场景的角色添加到实体列表里
	for id, avatarid := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
		if avatarid == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
				AvatarType:   g.Player.DbAvatar.Avatar[avatarid].Type,
				BaseAvatarId: avatarid,
			}},
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(pos.X),
					Y: int32(pos.Y),
					Z: int32(pos.Z),
				},
				Rot: &proto.Vector{
					X: int32(rot.X),
					Y: int32(rot.Y),
					Z: int32(rot.Z),
				},
			},
		}
		// 为进入场景的角色设置与上面相同的实体id
		if id == 0 {
			entityList.EntityId = leaderEntityId
			entityMap[leaderEntityId] = &EntityList{
				Entity:  avatarid,
				GroupId: 0,
			}
		} else {
			entityMap[entityId] = &EntityList{
				Entity:  avatarid,
				GroupId: 0,
			}
			entityList.EntityId = entityId
		}
		entityGroup.EntityList = append(entityGroup.EntityList, entityList)
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	// 获取场景实体
	for _, levelGroup := range gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID) {
		rsp.Scene.GroupIdList = append(rsp.Scene.GroupIdList, levelGroup.GroupId)

		sceneGroupState := &proto.SceneGroupState{
			GroupId:   levelGroup.GroupId,
			IsDefault: true,
		}

		rsp.Scene.GroupStateList = append(rsp.Scene.GroupStateList, sceneGroupState)

		entityGroupList := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加物品实体
		for _, propList := range levelGroup.PropList {
			entityList := &proto.SceneEntityInfo{
				GroupId:  levelGroup.GroupId, // 文件名后那个G
				InstId:   propList.ID,        // ID
				EntityId: uint32(g.GetNextGameObjectGuid()),
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: int32(propList.PosX * 1000),
						Y: int32(propList.PosY * 1000),
						Z: int32(propList.PosZ * 1000),
					},
					Rot: &proto.Vector{
						X: 0,
						Y: int32(propList.RotY * 1000),
						Z: 0,
					},
				},
				EntityCase: &proto.SceneEntityInfo_Prop{Prop: &proto.ScenePropInfo{
					PropId:    propList.PropID, // PropID
					PropState: gdconf.GetStateValue(propList.State),
				}},
			}
			entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
		}
		// 添加怪物实体
		for _, monsterList := range levelGroup.MonsterList {
			entityId := uint32(g.GetNextGameObjectGuid())
			entityList := &proto.SceneEntityInfo{
				GroupId:  levelGroup.GroupId,
				InstId:   monsterList.ID,
				EntityId: entityId,
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: int32(monsterList.PosX * 1000),
						Y: int32(monsterList.PosY * 1000),
						Z: int32(monsterList.PosZ * 1000),
					},
					Rot: &proto.Vector{
						X: 0,
						Y: int32(monsterList.RotY * 1000),
						Z: 0,
					},
				},
				EntityCase: &proto.SceneEntityInfo_NpcMonster{NpcMonster: &proto.SceneNpcMonsterInfo{
					WorldLevel: g.Player.WorldLevel,
					MonsterId:  monsterList.NPCMonsterID,
					EventId:    monsterList.EventID,
				}},
			}
			entityMap[entityId] = &EntityList{
				Entity:  monsterList.EventID,
				GroupId: levelGroup.GroupId,
			}
			entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
		}
		/*
			// 添加NPC实体
			for _, npcList := range levelGroup.NPCList {
				entityList := &proto.SceneEntityInfo{
					GroupId:  levelGroup.GroupId,
					InstId:   npcList.ID,
					EntityId: uint32(g.GetNextGameObjectGuid()),
					Motion: &proto.MotionInfo{
						Pos: &proto.Vector{
							X: int32(npcList.PosX * 1000),
							Y: int32(npcList.PosY * 1000),
							Z: int32(npcList.PosZ * 1000),
						},
						Rot: &proto.Vector{
							X: 0,
							Y: int32(npcList.RotY * 1000),
							Z: 0,
						},
					},
					EntityCase: &proto.SceneEntityInfo_Npc{Npc: &proto.SceneNpcInfo{
						NpcId: npcList.NPCID,
					}},
				}
				entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
			}
		*/
		rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroupList)
	}

	g.Player.EntityList = entityMap
	g.Send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *Game) HanldeGetSceneMapInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetSceneMapInfoCsReq, payloadMsg)
	req := msg.(*proto.GetSceneMapInfoCsReq)

	entryId := req.EntryIdList[0]

	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))

	rsp := new(proto.GetSceneMapInfoScRsp)

	if mapEntrance != nil {
		groupMap := gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)
		if groupMap != nil {
			if int(mapEntrance.StartGroupID) > len(groupMap) {
				return
			}

			mapList := &proto.MazeMapData{
				LightenSectionList: make([]uint32, 0),
				UnlockedChestList: []*proto.MazeChest{
					{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_NORMAL},
					{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_PUZZLE},
					{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_CHALLENGE},
				},
				UnlockedTeleportList: make([]uint32, 0),
			}

			mapList.EntryId = entryId

			for i := uint32(0); i < 100; i++ {
				mapList.LightenSectionList = append(mapList.LightenSectionList, i)
			}

			for _, groupInfo := range groupMap {
				mazeGroup := &proto.MazeGroup{GroupId: groupInfo.GroupId}
				mapList.MazeGroupList = append(mapList.MazeGroupList, mazeGroup)
			}

			for _, groupMapList := range groupMap {
				for _, propList := range groupMapList.PropList {
					if propList.State != "CheckPointDisable" && propList.State != "CheckPointEnable" {
						continue
					}
					mazeProp := &proto.MazeProp{
						State:    gdconf.GetStateValue(propList.State),
						GroupId:  groupMapList.GroupId,
						ConfigId: propList.ID,
					}
					mapList.MazePropList = append(mapList.MazePropList, mazeProp)
					mapList.UnlockedTeleportList = append(mapList.UnlockedTeleportList, propList.MappingInfoID)
				}
			}
			rsp.MapList = append(rsp.MapList, mapList)
		}
	}
	g.Send(cmd.GetSceneMapInfoScRsp, rsp)
}

func (g *Game) EnterSceneCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.EnterSceneCsReq, payloadMsg)
	req := msg.(*proto.EnterSceneCsReq)
	rsp := &proto.GetEnteredSceneScRsp{}

	g.EnterSceneByServerScNotify(req.EntryId, req.TeleportId)

	g.Send(cmd.EnterSceneScRsp, rsp)
	g.Send(cmd.SceneUpdatePositionVersionNotify, rsp)
}
