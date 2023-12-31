package Game

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

/***********************************大世界攻击事件处理***********************************/

func (g *Game) SceneCastSkillCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneCastSkillCsReq, payloadMsg)
	req := msg.(*proto.SceneCastSkillCsReq)

	var targetIndex uint32 = 0
	var stageConfig *gdconf.StageConfig
	var stageID *gdconf.PlaneEvent

	if req.SkillIndex == 1 {
		avatarId := g.Player.EntityList[req.CasterId].Entity
		skillId := (avatarId * 100) + req.SkillIndex
		if gdconf.GetMazeBuffById(skillId, req.SkillIndex) != nil {
			g.Player.DbAvatar.Avatar[avatarId].BuffList = skillId
		} else {
			// 技能处理，有的技能并不会增加buff而是回复生命等功能
		}
	}
	if len(req.HitTargetIdList) == 0 {
		rsp := &proto.SceneCastSkillScRsp{
			AttackedGroupId: req.AttackedGroupId,
		}
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}

	if g.Player.EntityList[req.HitTargetIdList[0]] == nil {
		rsp := &proto.SceneCastSkillScRsp{
			AttackedGroupId: req.AttackedGroupId,
		}
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	entity := g.Player.EntityList[req.HitTargetIdList[0]]
	stageID = gdconf.GetPlaneEventById(entity.Entity, g.Player.WorldLevel)
	if stageID == nil {
		newEntity := g.Player.EntityList[req.CasterId]
		stageID = gdconf.GetPlaneEventById(newEntity.Entity, g.Player.WorldLevel)
		stageConfig = gdconf.GetStageConfigById(stageID.StageID)
	} else {
		stageConfig = gdconf.GetStageConfigById(stageID.StageID)
	}

	// 构造回复包
	rsp := &proto.SceneCastSkillScRsp{
		AttackedGroupId: req.AttackedGroupId,
		BattleInfo: &proto.SceneBattleInfo{
			BuffList:         make([]*proto.BattleBuff, 0), // Buff列表
			LogicRandomSeed:  gdconf.GetLoadingDesc(),      // 逻辑随机种子
			StageId:          stageID.StageID,              // 阶段id
			TurnSnapshotList: nil,                          // 打开快照列表？
			WorldLevel:       g.Player.WorldLevel,
			RoundsLimit:      0,                              // 回合限制
			BattleId:         g.GetBattleIdGuid(),            // 战斗Id
			BattleAvatarList: make([]*proto.BattleAvatar, 0), // 战斗角色列表
		},
	}

	// 怪物波列表
	for id, monsterListMap := range stageConfig.MonsterList {
		monsterWaveList := &proto.SceneMonsterWave{
			StageId: stageID.StageID,
			WaveId:  uint32(id + 1),
		}
		for _, monsterList := range monsterListMap {
			sceneMonster := &proto.SceneMonster{
				MonsterId: monsterList,
			}
			monsterWaveList.MonsterList = append(monsterWaveList.MonsterList, sceneMonster)
		}
		rsp.BattleInfo.MonsterWaveList = append(rsp.BattleInfo.MonsterWaveList, monsterWaveList)
	}
	// 添加角色
	for id, Lineup := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
		if Lineup == 0 {
			continue
		}
		avatar := g.Player.DbAvatar.Avatar[Lineup]

		battleAvatar := &proto.BattleAvatar{
			AvatarType:    proto.AvatarType_AVATAR_FORMAL_TYPE,
			Id:            Lineup,
			Level:         avatar.Level,
			Rank:          avatar.Rank,
			Index:         uint32(id),
			SkilltreeList: g.GetSkilltree(avatar.AvatarId),
			Hp:            avatar.Hp,
			Promotion:     avatar.Promotion,
			RelicList:     make([]*proto.BattleRelic, 0),
			WorldLevel:    g.Player.WorldLevel,
			SpBar: &proto.SpBarInfo{
				CurSp: avatar.SpBar.CurSp,
				MaxSp: avatar.SpBar.MaxSp,
			},
		}
		// 获取角色装备的光锥
		if avatar.EquipmentUniqueId != 0 {
			equipment := g.Player.DbItem.EquipmentMap[avatar.EquipmentUniqueId]
			equipmentList := &proto.BattleEquipment{
				Id:        equipment.Tid,
				Level:     equipment.Level,
				Promotion: equipment.Promotion,
				Rank:      equipment.Rank,
			}
			battleAvatar.EquipmentList = append(battleAvatar.EquipmentList, equipmentList)
		}
		rsp.BattleInfo.BattleAvatarList = append(rsp.BattleInfo.BattleAvatarList, battleAvatar)
		// 检查是否有提前释放的技能，添加到buff里
		if avatar.BuffList != 0 {
			buffList := &proto.BattleBuff{
				Id:              avatar.BuffList,
				Level:           1,
				OwnerId:         targetIndex,
				TargetIndexList: []uint32{targetIndex},
				WaveFlag:        4294967295, // 失效时间
			}
			rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
			targetIndex++
			g.Player.DbAvatar.Avatar[Lineup].BuffList = 0
		}
	}

	// 储存战斗信息
	g.Player.Battle = make(map[uint32]*Battle)
	battle := &Battle{
		BattleId:         rsp.BattleInfo.BattleId,
		EventID:          req.HitTargetIdList[0],
		LogicRandomSeed:  rsp.BattleInfo.LogicRandomSeed,
		RoundsLimit:      rsp.BattleInfo.RoundsLimit,
		BuffList:         rsp.BattleInfo.BuffList,
		BattleAvatarList: rsp.BattleInfo.BattleAvatarList,
	}
	g.Player.Battle[rsp.BattleInfo.BattleId] = battle

	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

/***********************************战斗结算***********************************/

func (g *Game) PVEBattleResultCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PVEBattleResultCsReq, payloadMsg)
	req := msg.(*proto.PVEBattleResultCsReq)

	rsp := new(proto.PVEBattleResultScRsp)

	// 要记得扣体力,和回复奖励

	rsp.BattleId = req.BattleId
	rsp.StageId = req.StageId
	rsp.EndStatus = req.EndStatus // 战斗结算状态
	rsp.CheckIdentical = true     // 反作弊验证
	rsp.BinVersion = ""
	rsp.ResVersion = strconv.Itoa(int(req.ClientResVersion)) // 版本验证

	// 撤退
	if req.EndStatus == proto.BattleEndStatus_BATTLE_END_QUIT {
		// 删除储存的战斗信息
		delete(g.Player.Battle, req.BattleId)
		g.Send(cmd.PVEBattleResultScRsp, rsp)
		return
	}

	// 更新角色状态
	for _, avatar := range req.Stt.BattleAvatarList {
		g.Player.DbAvatar.Avatar[avatar.Id].Type = avatar.AvatarType
		g.Player.DbAvatar.Avatar[avatar.Id].SpBar.CurSp = uint32((avatar.AvatarStatus.LeftSp / avatar.AvatarStatus.MaxSp) * 10000)
		if avatar.AvatarStatus.LeftHp == float64(0) {
			g.Player.DbAvatar.Avatar[avatar.Id].Hp = 2000
			g.Player.DbAvatar.Avatar[avatar.Id].Type = proto.AvatarType_AVATAR_FORMAL_TYPE
		} else {
			g.Player.DbAvatar.Avatar[avatar.Id].Hp = uint32((avatar.AvatarStatus.LeftHp / avatar.AvatarStatus.MaxHp) * 10000)
		}
	}
	// 更新队伍状态
	g.BattleSyncLineupNotify(g.Player.DbLineUp.MainLineUp)

	// 账号状态改变通知
	g.PlayerPlayerSyncScNotify()

	// 胜利时获取奖励
	if req.EndStatus == proto.BattleEndStatus_BATTLE_END_WIN {
		battle := g.Player.Battle[req.BattleId]
		entity := g.Player.EntityList[battle.EventID]

		if entity != nil {
			// 删除实体
			nitify := new(proto.SceneGroupRefreshScNotify)
			nitify.GroupRefreshInfo = []*proto.SceneGroupRefreshInfo{
				{
					GroupId: entity.GroupId,
					RefreshEntity: []*proto.SceneEntityRefreshInfo{
						{
							UpdateType: &proto.SceneEntityRefreshInfo_DelEntity{
								DelEntity: battle.EventID,
							},
						},
					},
				},
			}
			g.Send(cmd.SceneGroupRefreshScNotify, nitify)
			delete(g.Player.EntityList, battle.EventID)
		}

		g.Player.DbItem.MaterialMap[11].Num -= battle.StaminaCost * battle.Wave // 扣除体力

		// 获取奖励
		rsp.DropData = &proto.ItemList{ItemList: make([]*proto.Item, 0)}
		for _, drop := range battle.DisplayItemList {
			item := &proto.Item{
				ItemId:      drop.Tid,
				Level:       0,
				Num:         drop.Num * battle.Wave,
				MainAffixId: 0,
				Rank:        0,
				Promotion:   0,
				UniqueId:    0,
			}
			rsp.DropData.ItemList = append(rsp.DropData.ItemList, item)

			g.AddMaterial(drop.Tid, drop.Num*battle.Wave)
		}
	}

	// 当前坐标通知(失败情况应该是移动到最近锚点)
	g.SceneEntityMoveScNotify()

	// 体力改变通知
	g.StaminaInfoScNotify()

	// 删除储存的战斗信息
	delete(g.Player.Battle, req.BattleId)

	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

// 队伍更新通知
func (g *Game) BattleSyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	lineUp := g.Player.DbLineUp.LineUpList[index]
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           index,
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
	rsq.Lineup = lineupList

	g.Send(cmd.SyncLineupNotify, rsq)
}

/***********************************模拟宇宙***********************************/

func (g *Game) GetRogueInfoCsReq(payloadMsg []byte) {
	beginTime := time.Now().AddDate(0, 0, -1).Unix()
	endTime := beginTime + int64(time.Hour.Seconds()*24*8)
	rsp := new(proto.GetRogueInfoScRsp)
	rogueInfo := &proto.RogueInfo{
		BeginTime: beginTime,
		EndTime:   endTime,
		SeasonId:  77,
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
			RogueAbilityPoint: 0,
		},
		RogueScoreInfo: &proto.RogueScoreRewardInfo{
			PoolId:               20 + g.Player.WorldLevel,
			HasTakenInitialScore: true,
			PoolRefreshed:        true,
		},
		RogueData: &proto.RogueInfoData{
			RogueSeasonInfo: &proto.RogueSeasonInfo{
				BeginTime: beginTime,
				SeasonId:  77,
				EndTime:   endTime,
			},
			RogueScoreInfo: &proto.RogueScoreRewardInfo{
				PoolId:               20 + g.Player.WorldLevel,
				HasTakenInitialScore: true,
				PoolRefreshed:        true,
			},
		},
		RogueAreaList: make([]*proto.RogueArea, 0),
	}
	for _, rogueArea := range gdconf.GetRogueAreaMap() {
		RogueArea := &proto.RogueArea{
			AreaId:          rogueArea.RogueAreaID,
			RogueAreaStatus: proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS,
		}
		rogueInfo.RogueAreaList = append(rogueInfo.RogueAreaList, RogueArea)
	}
	rsp.RogueInfo = rogueInfo

	g.Send(cmd.GetRogueInfoScRsp, rsp)
}

func (g *Game) StartRogueCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartRogueCsReq, payloadMsg)
	req := msg.(*proto.StartRogueCsReq)

	if req.BaseAvatarIdList == nil {
		req.BaseAvatarIdList = g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList
	}

	entityMap := make(map[uint32]*EntityList) // [实体id]怪物群id
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	// 获取地图
	rogueAreaConfig := gdconf.GetRogueAreaConfigById(strconv.Itoa(int(req.AreaId)))
	if rogueAreaConfig == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	rogueMapID := (rogueAreaConfig.AreaProgress * 100) + rogueAreaConfig.Difficulty
	rogueMapStart := gdconf.GetRogueMapStartById(strconv.Itoa(int(rogueMapID)))
	if rogueMapStart == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	// 获取映射信息
	rogueMap := gdconf.GetRogueRoomIDBySiteID()

	beginTime := time.Now().AddDate(0, 0, -1).Unix()
	endTime := beginTime + int64(time.Hour.Seconds()*24*8)

	// 可独立成单独的方法
	rogueScoreInfo := &proto.RogueScoreRewardInfo{
		HasTakenInitialScore: true, // 已取得初始积分？
		RogueImmersifier:     0,
		Score:                0,
		PoolRefreshed:        true, // 刷新？
		TakenScoreRewardList: nil,
		PoolId:               20 + g.Player.WorldLevel,
	}
	// 可独立成单独的方法
	roomMap := &proto.RogueMapInfo{
		MapId:     rogueMapID,
		AreaId:    req.AreaId,
		CurSiteId: rogueMapStart.SiteID,
		CurRoomId: rogueMap[1],
		RoomList:  make([]*proto.RogueRoom, 0),
	}
	for id, rogue := range rogueMap {
		roomList := &proto.RogueRoom{
			SiteId: id,
			RoomId: rogue,
		}
		if id == rogueMapStart.SiteID {
			roomList.RoomStatus = proto.RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY
		} else {
			roomList.RoomStatus = proto.RogueRoomStatus_ROGUE_ROOM_STATUS_NONE
		}

		roomMap.RoomList = append(roomMap.RoomList, roomList)
	}
	// 可独立成单独的方法
	rogueAreaList := make([]*proto.RogueArea, 0)
	for _, rogueArea := range gdconf.GetRogueAreaMap() {
		RogueArea := &proto.RogueArea{
			AreaId:          rogueArea.RogueAreaID,
			RogueAreaStatus: proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS,
		}
		rogueAreaList = append(rogueAreaList, RogueArea)
	}
	// 可独立成单独的方法
	rogueRoom := gdconf.GetRogueRoomById(strconv.Itoa(int(rogueMap[rogueMapStart.SiteID])))
	if rogueRoom == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(rogueRoom.MapEntrance)))
	if mapEntrance == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            mapEntrance.PlaneID,
		FloorId:            mapEntrance.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).WorldID,
		EntryId:            rogueRoom.MapEntrance,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}

	entityGroupList := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}

	// 添加角色信息
	startGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, rogueRoom.GroupID)
	anchor := startGroup.AnchorList[0]
	for id, avatarId := range req.BaseAvatarIdList {
		if avatarId == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
				BaseAvatarId: avatarId,
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
				Entity:  avatarId,
				GroupId: rogueRoom.GroupID,
			}
		} else {
			entityList.EntityId = entityId
			entityMap[entityId] = &EntityList{
				Entity:  avatarId,
				GroupId: rogueRoom.GroupID,
			}
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}

	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)

	// 获取场景实体
	for groupID, _ := range rogueRoom.GroupWithContent {
		sceneGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, stou32(groupID))
		if sceneGroup == nil {
			continue
		}
		scene.GroupIdList = append(scene.GroupIdList, stou32(groupID))

		sceneGroupState := &proto.SceneGroupState{
			GroupId:   stou32(groupID),
			IsDefault: true,
		}

		scene.GroupStateList = append(scene.GroupStateList, sceneGroupState)

		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    stou32(groupID),
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加物品实体
		for _, propList := range sceneGroup.PropList {
			entityList := &proto.SceneEntityInfo{
				GroupId:  stou32(groupID), // 文件名后那个G
				InstId:   propList.ID,     // ID
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
			entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
		}
		// 添加怪物实体
		for _, monsterList := range sceneGroup.MonsterList {
			entityId := uint32(g.GetNextGameObjectGuid())
			entityList := &proto.SceneEntityInfo{
				GroupId:  stou32(groupID),
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
				GroupId: stou32(groupID),
			}
			entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
		}

		// 添加NPC实体
		for _, npcList := range sceneGroup.NPCList {
			entityList := &proto.SceneEntityInfo{
				GroupId:  stou32(groupID),
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
			entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
		}

		if len(entityGroupLists.EntityList) == 0 {
			continue
		}
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}

	// 先更新队伍
	for id, avatarid := range req.BaseAvatarIdList {
		g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList[id] = avatarid
	}
	g.Player.DbLineUp.MainAvatarId = 0
	// 队伍更新通知
	g.SyncLineupNotify(g.Player.DbLineUp.MainLineUp)

	rsp := &proto.StartRogueScRsp{
		Scene: scene,
		Lineup: &proto.LineupInfo{
			AvatarList:      make([]*proto.LineupAvatar, 0),
			ExtraLineupType: proto.ExtraLineupType_LINEUP_ROGUE,
			MaxMp:           5,
			Mp:              5,
		},
		RogueInfo: &proto.RogueInfo{
			RogueData: &proto.RogueInfoData{
				RogueScoreInfo: rogueScoreInfo,
				RogueSeasonInfo: &proto.RogueSeasonInfo{
					BeginTime: beginTime,
					SeasonId:  77,
					EndTime:   endTime,
				},
			},
			RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
				RoguePumanCoupon:  0,
				RogueCoin:         0,
				RogueImmersifier:  0,
				RogueAbilityPoint: 0,
			},
			RogueScoreInfo: rogueScoreInfo,
			RoomMap:        roomMap,
			RogueAreaList:  rogueAreaList,
			Status:         proto.RogueStatus_ROGUE_STATUS_DOING,
			SeasonId:       77,
			RogueProgress: &proto.RogueCurrentInfo{
				RogueAvatarInfo: &proto.RogueAvatarInfo{
					BaseAvatarIdList: req.BaseAvatarIdList,
				},
				RoomMap: roomMap,
				Status:  proto.RogueStatus_ROGUE_STATUS_DOING,
				RogueBuffInfo: &proto.RogueBuffInfo{
					BuffSelectInfo: nil,
					MazeBuffList:   nil,
				},
				RogueMiracleInfo: &proto.RogueMiracleInfo{
					MiracleSelectInfo: nil,
					RogueMiracleInfo:  nil,
				},
			},
			RogueAeonInfo: &proto.RogueAeonInfo{
				UnlockAeonEnhanceNum: 0,
				UnlockAeonNum:        0,
				SelectedAeonId:       0,
				AeonIdList:           nil,
			},
			EndTime:          endTime,
			BaseAvatarIdList: req.BaseAvatarIdList,
			BeginTime:        beginTime,
			RogueCoin:        0,
		},
	}

	for slot, avatarId := range req.BaseAvatarIdList {
		if avatarId == 0 {
			continue
		}
		avatar := g.Player.DbAvatar.Avatar[avatarId]
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: avatar.Type,
			Slot:       uint32(slot),
			Hp:         avatar.Hp,
			Id:         avatarId,
			SpBar: &proto.SpBarInfo{
				CurSp: avatar.SpBar.CurSp,
				MaxSp: avatar.SpBar.MaxSp,
			},
		}
		rsp.Lineup.AvatarList = append(rsp.Lineup.AvatarList, lineupAvatar)
	}

	g.Player.EntityList = entityMap
	g.Player.IsRogue = true
	g.Send(cmd.StartRogueScRsp, rsp)
}

func (g *Game) GetRogueTalentInfoCsReq() {
	rsp := &proto.GetRogueTalentInfoScRsp{
		TalentInfo: &proto.RogueTalentInfo{
			RogueTalent: make([]*proto.RogueTalent, 0),
		},
	}

	for _, talent := range gdconf.GetTalentIDList() {
		rogueTalent := &proto.RogueTalent{
			Status: proto.RogueTalentStatus_ROGUE_TALENT_STATUS_UNLOCK,
			// UnlockProgressList: nil,
			TalentId: talent,
		}
		rsp.TalentInfo.RogueTalent = append(rsp.TalentInfo.RogueTalent, rogueTalent)
	}

	g.Send(cmd.GetRogueTalentInfoScRsp, rsp)
}

func (g *Game) LeaveRogueCsReq(payloadMsg []byte) {
}

func (g *Game) QuitRogueCsReq(payloadMsg []byte) {
}

/***********************************关卡/副本***********************************/

func (g *Game) StartCocoonStageCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartCocoonStageCsReq, payloadMsg)
	req := msg.(*proto.StartCocoonStageCsReq)
	var targetIndex uint32 = 0

	rsp := new(proto.StartCocoonStageScRsp)
	rsp.PropEntityId = req.PropEntityId
	rsp.CocoonId = req.CocoonId
	rsp.Wave = req.Wave

	cocoonConfig := gdconf.GetCocoonConfigById(req.CocoonId, req.WorldLevel)

	if len(cocoonConfig.DropList) == 0 {
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
		g.Send(cmd.StartCocoonStageScRsp, rsp)
		return
	}

	rsp.BattleInfo = new(proto.SceneBattleInfo)
	rsp.BattleInfo.LogicRandomSeed = gdconf.GetLoadingDesc()
	rsp.BattleInfo.BattleId = g.GetBattleIdGuid() // 战斗Id
	rsp.BattleInfo.StageId = cocoonConfig.StageID
	rsp.BattleInfo.WorldLevel = g.Player.WorldLevel

	if rsp.BattleInfo.StageId == 0 {
		rsp.BattleInfo.StageId = cocoonConfig.StageIDList[0]
	}

	for id, StageList := range cocoonConfig.StageIDList {
		stageConfig := gdconf.GetStageConfigById(StageList)
		// 怪物波列表
		for _, monsterListMap := range stageConfig.MonsterList {
			monsterWaveList := &proto.SceneMonsterWave{
				StageId: StageList,
				WaveId:  uint32(id + 1),
			}
			for _, monsterList := range monsterListMap {
				sceneMonster := &proto.SceneMonster{
					MonsterId: monsterList,
				}
				monsterWaveList.MonsterList = append(monsterWaveList.MonsterList, sceneMonster)
			}
			rsp.BattleInfo.MonsterWaveList = append(rsp.BattleInfo.MonsterWaveList, monsterWaveList)
		}
	}

	// 添加角色
	for id, Lineup := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
		if Lineup == 0 {
			continue
		}
		avatar := g.Player.DbAvatar.Avatar[Lineup]

		battleAvatar := &proto.BattleAvatar{
			AvatarType:    proto.AvatarType_AVATAR_FORMAL_TYPE,
			Id:            Lineup,
			Level:         avatar.Level,
			Rank:          avatar.Rank,
			Index:         uint32(id),
			SkilltreeList: g.GetSkilltree(avatar.AvatarId),
			Hp:            avatar.Hp,
			Promotion:     avatar.Promotion,
			RelicList:     make([]*proto.BattleRelic, 0),
			WorldLevel:    g.Player.WorldLevel,
			SpBar: &proto.SpBarInfo{
				CurSp: avatar.SpBar.CurSp,
				MaxSp: avatar.SpBar.MaxSp,
			},
		}
		// 获取角色装备的光锥
		if avatar.EquipmentUniqueId != 0 {
			equipment := g.Player.DbItem.EquipmentMap[avatar.EquipmentUniqueId]
			equipmentList := &proto.BattleEquipment{
				Id:        equipment.Tid,
				Level:     equipment.Level,
				Promotion: equipment.Promotion,
				Rank:      equipment.Rank,
			}
			battleAvatar.EquipmentList = append(battleAvatar.EquipmentList, equipmentList)
		}
		rsp.BattleInfo.BattleAvatarList = append(rsp.BattleInfo.BattleAvatarList, battleAvatar)
		// 检查是否有提前释放的技能，添加到buff里
		if avatar.BuffList != 0 {
			buffList := &proto.BattleBuff{
				Id:              avatar.BuffList,
				Level:           1,
				OwnerId:         targetIndex,
				TargetIndexList: []uint32{targetIndex},
				WaveFlag:        4294967295, // 失效时间
			}
			rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
			targetIndex++
			g.Player.DbAvatar.Avatar[Lineup].BuffList = 0
		}
	}

	// 储存战斗信息
	g.Player.Battle = make(map[uint32]*Battle)
	battle := &Battle{
		BattleId:         rsp.BattleInfo.BattleId,
		Wave:             req.Wave,
		EventID:          req.CocoonId,
		LogicRandomSeed:  rsp.BattleInfo.LogicRandomSeed,
		RoundsLimit:      rsp.BattleInfo.RoundsLimit,
		StaminaCost:      cocoonConfig.StaminaCost,
		BuffList:         rsp.BattleInfo.BuffList,
		BattleAvatarList: rsp.BattleInfo.BattleAvatarList,
	}
	// 添加奖励
	for _, displayItem := range gdconf.GetMappingInfoById(req.CocoonId, g.Player.WorldLevel).DisplayItemList {
		material := &Material{
			Tid: displayItem.ItemID,
			Num: displayItem.ItemNum,
		}
		if material.Num == 0 {
			material.Num += 1
		}
		battle.DisplayItemList = append(battle.DisplayItemList, material)
	}
	g.Player.Battle[rsp.BattleInfo.BattleId] = battle

	g.Send(cmd.StartCocoonStageScRsp, rsp)
}
