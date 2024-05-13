package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) StartTrialActivityCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartTrialActivityCsReq, payloadMsg)
	req := msg.(*proto.StartTrialActivityCsReq)
	if req.TrialActivityId == 0 {
		return
	}

	avatarDemo := gdconf.GetAvatarDemoConfigById(req.TrialActivityId)
	if avatarDemo == nil {
		return
	}
	lineup := g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ACTIVITY))

	trialActivityState := g.GetTrialActivityState()
	trialActivityState.AvatarDemoId = req.TrialActivityId
	// 记录场景
	trialActivityState.PlaneID = avatarDemo.PlaneID
	trialActivityState.FloorID = avatarDemo.FloorID
	trialActivityState.EntranceID = avatarDemo.MapEntranceID
	// 记录怪物信息
	trialActivityState.EventID = avatarDemo.EventIDList1[0]
	trialActivityState.ConfigID = avatarDemo.ConfigList1[0]
	trialActivityState.NPCMonsterID = avatarDemo.NpcMonsterIDList1[0]
	trialActivityState.GroupID = avatarDemo.MazeGroupID1

	for id, avatarId := range avatarDemo.TrialAvatarList {
		lineup.AvatarIdList[uint32(id)] = &spb.LineAvatarList{
			Slot:     uint32(id),
			AvatarId: avatarId,
		}
	}

	g.Send(cmd.ExtraLineupDestroyNotify, &proto.ExtraLineupDestroyNotify{ExtraLineupType: proto.ExtraLineupType_LINEUP_STAGE_TRIAL})
	g.Send(cmd.SyncServerSceneChangeNotify, nil)

	g.SyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_ACTIVITY), true)

	g.StartTrialEnterSceneByServerScNotify()

	g.GetBattleState().BattleType = spb.BattleType_Battle_TrialActivity

	rsp := &proto.StartTrialActivityScRsp{TrialActivityId: req.TrialActivityId}
	g.Send(cmd.StartTrialActivityScRsp, rsp)
}

func (g *GamePlayer) StartTrialEnterSceneByServerScNotify() {
	rsp := new(proto.EnterSceneByServerScNotify)
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	trialActivityState := g.GetTrialActivityState()

	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(trialActivityState.EntranceID)))
	if mapEntrance == nil {
		return
	}
	foorMap := gdconf.GetFloorById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return
	}
	var anchorID = mapEntrance.StartAnchorID

	anchorID = foorMap.StartAnchorID

	// 获取队伍
	lineup := g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ACTIVITY))
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		ExtraLineupType: proto.ExtraLineupType_LINEUP_ACTIVITY,
		MaxMp:           5,
		Mp:              5,
		PlaneId:         0,
	}
	for slot, lineAvatar := range lineup.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: proto.AvatarType_AVATAR_TRIAL_TYPE,
			Slot:       slot,
			Satiety:    0,
			Hp:         10000,
			Id:         lineAvatar.AvatarId,
			SpBar: &proto.SpBarInfo{
				CurSp: 5000,
				MaxSp: 10000,
			},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}

	rsp.Lineup = lineupList

	rsp.Scene = &proto.SceneInfo{
		ClientPosVersion:   5,
		EntryId:            trialActivityState.EntranceID,
		FloorId:            trialActivityState.FloorID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(trialActivityState.PlaneID))).PlaneType),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		LeaderEntityId:     leaderEntityId,
		LightenSectionList: make([]uint32, 0),
		PlaneId:            trialActivityState.PlaneID,
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(trialActivityState.PlaneID))).WorldID,
	}

	for i := uint32(0); i < 100; i++ {
		rsp.Scene.LightenSectionList = append(rsp.Scene.LightenSectionList, i)
	}

	monsterEntity := make(map[uint32]*MonsterEntity, 0)
	avatarEntity := make(map[uint32]*AvatarEntity, 0)
	npcEntity := make(map[uint32]*NpcEntity, 0)
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 添加队伍角色进实体列表，并设置坐标
	if foorMap.Groups[trialActivityState.GroupID] == nil {
		return
	}
	for _, anchor := range foorMap.Groups[trialActivityState.GroupID].AnchorList {
		if anchor.ID == anchorID {
			lineUpBin := g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ACTIVITY))
			for id, lineAvatar := range lineUpBin.AvatarIdList {
				if lineAvatar == nil || lineAvatar.AvatarId == 0 {
					continue
				}
				avatarid := gdconf.GetSpecialAvatarById(lineAvatar.AvatarId).AvatarID
				entityId := uint32(g.GetNextGameObjectGuid())
				entityList := &proto.SceneEntityInfo{
					Actor: &proto.SceneActorInfo{
						AvatarType:   proto.AvatarType_AVATAR_TRIAL_TYPE,
						BaseAvatarId: avatarid,
					},
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
					avatarEntity[leaderEntityId] = &AvatarEntity{
						AvatarId: avatarid,
					}
				} else {
					entityList.EntityId = entityId
					avatarEntity[entityId] = &AvatarEntity{
						AvatarId: avatarid,
					}
				}
				entityGroup.EntityList = append(entityGroup.EntityList, entityList)
			}
			break
		}
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	// 获取场景实体
	for _, levelGroup := range foorMap.Groups {
		if levelGroup.GroupId == 0 {
			continue
		}
		if len(levelGroup.PropList) == 0 && len(levelGroup.NPCList) == 0 && len(levelGroup.MonsterList) == 0 {
			continue
		}
		rsp.Scene.GroupIdList = append(rsp.Scene.GroupIdList, levelGroup.GroupId)

		// 添加物品实体
		propList := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		g.GetPropByID(propList, levelGroup, levelGroup.GroupId)
		if len(propList.EntityList) != 0 {
			rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, propList)
		}
	}

	// 添加怪物实体
	// [实体id]怪物群id
	entityGroupLists := &proto.SceneEntityGroupInfo{
		GroupId:    trialActivityState.GroupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for _, monsterList := range foorMap.Groups[trialActivityState.GroupID].MonsterList {
		if monsterList.ID == trialActivityState.ConfigID {
			entityId := uint32(g.GetNextGameObjectGuid())
			entityList := &proto.SceneEntityInfo{
				GroupId:  trialActivityState.GroupID,
				InstId:   trialActivityState.ConfigID,
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
				NpcMonster: &proto.SceneNpcMonsterInfo{
					WorldLevel: g.PlayerPb.WorldLevel,
					MonsterId:  trialActivityState.NPCMonsterID,
					EventId:    trialActivityState.EventID,
				},
			}
			// 添加实体
			monsterEntity[entityId] = &MonsterEntity{
				MonsterEId: trialActivityState.EventID,
				GroupId:    trialActivityState.GroupID,
				Pos: &Vector{
					X: int32(monsterList.PosX * 1000),
					Y: int32(monsterList.PosY * 1000),
					Z: int32(monsterList.PosZ * 1000),
				},
				Rot: &Vector{
					X: 0,
					Y: int32(monsterList.RotY * 1000),
					Z: 0,
				},
			}
			entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
			break
		} else {
			continue
		}
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroupLists)
	g.GetSceneEntity().MonsterEntity = monsterEntity
	g.GetSceneEntity().AvatarEntity = avatarEntity
	g.GetSceneEntity().NpcEntity = npcEntity

	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *GamePlayer) TrialActivitySceneCastSkillScRsp(rsp *proto.SceneCastSkillScRsp) {
	// var targetIndex uint32 = 0
	// trialActivityState := g.GetTrialActivityState()
	// 添加角色
	rsp.BattleInfo.BattleAvatarList = g.TrialActivityGetBattleAvatarList()
	// 添加角色buff
	/*
		for _, buffId := range trialActivityState.AvatarBuffList {
			buffList := &proto.BattleBuff{
				Id:              buffId,
				Level:           1,
				OwnerId:         targetIndex,
				TargetIndexList: []uint32{targetIndex},
				WaveFlag:        4294967295, // 失效时间
			}
			rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
			targetIndex++
		}
	*/

	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

func (g *GamePlayer) TrialActivityGetBattleAvatarList() []*proto.BattleAvatar {
	battleAvatarList := make([]*proto.BattleAvatar, 0)
	lineupBin := g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ACTIVITY))
	for id, lineAvatar := range lineupBin.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		avatar := gdconf.GetSpecialAvatarById(lineAvatar.AvatarId)

		battleAvatar := &proto.BattleAvatar{
			AvatarType:    proto.AvatarType_AVATAR_TRIAL_TYPE,
			Id:            lineAvatar.AvatarId,
			Level:         avatar.Level,
			Rank:          0,
			Index:         id,
			SkilltreeList: make([]*proto.AvatarSkillTree, 0),
			Hp:            10000,
			Promotion:     avatar.Promotion,
			RelicList:     make([]*proto.BattleRelic, 0),
			WorldLevel:    g.PlayerPb.WorldLevel,
			SpBar: &proto.SpBarInfo{
				CurSp: 6000,
				MaxSp: 10000,
			},
		}
		for _, skill := range g.TrialActivityGetSkillTreeList(avatar.AvatarID) {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			battleAvatar.SkilltreeList = append(battleAvatar.SkilltreeList, avatarSkillTree)
		}
		// 获取角色装备的光锥
		if avatar.EquipmentID != 0 {
			equipmentList := &proto.BattleEquipment{
				Id:        avatar.EquipmentID,
				Level:     avatar.EquipmentLevel,
				Promotion: avatar.EquipmentPromotion,
				Rank:      avatar.EquipmentRank,
			}
			battleAvatar.EquipmentList = append(battleAvatar.EquipmentList, equipmentList)
		}
		battleAvatarList = append(battleAvatarList, battleAvatar)
	}
	return battleAvatarList
}

func (g *GamePlayer) TrialActivityGetSkillTreeList(avatarId uint32) []*spb.AvatarSkillBin {
	skilltreeList := make([]*spb.AvatarSkillBin, 0)
	for id, level := range gdconf.GetAvatarSkilltreeListById(avatarId) {
		avatarSkillBin := &spb.AvatarSkillBin{
			PointId: id,
			Level:   level,
		}
		skilltreeList = append(skilltreeList, avatarSkillBin)
	}
	return skilltreeList

}

func (g *GamePlayer) TrialActivityPVEBattleResultScRsp(rsp *proto.PVEBattleResultScRsp) {
	rsp.BattleAvatarList = g.TrialActivityGetBattleAvatarList()
	if rsp.EndStatus == proto.BattleEndStatus_BATTLE_END_WIN {
		// 传送回原来的场景
		g.SceneByServerScNotify(g.GetScene().EntryId, g.GetPos(), g.GetRot())
		// 储存通关状态
		g.GetActivity().TrialActivity = append(g.GetActivity().TrialActivity, g.GetTrialActivityState().AvatarDemoId)
		// 发送通关通知
		scNotify := &proto.TrialActivityDataChangeScNotify{
			TrialActivityInfo: &proto.TrialActivityInfo{
				TrialActivityId: g.GetTrialActivityState().AvatarDemoId,
				TakenReward:     false,
			},
		}
		g.Send(cmd.TrialActivityDataChangeScNotify, scNotify)
		notify := &proto.CurTrialActivityScNotify{
			TrialActivityId: g.GetTrialActivityState().AvatarDemoId,
			Status:          proto.TrialActivityStatus_TRIAL_ACTIVITY_STATUS_FINISH,
		}
		g.Send(cmd.CurTrialActivityScNotify, notify)
		// 恢复战斗状态为空
		g.GetBattleState().BattleType = spb.BattleType_Battle_NONE
	}
	g.Send(cmd.PVEBattleResultScRsp, rsp)
}
