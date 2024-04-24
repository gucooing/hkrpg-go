package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/***********************************忘却之庭***********************************/

// 获取状态

func (g *GamePlayer) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurChallengeScRsp)

	challengeState := g.GetChallengeState()

	rsp.ChallengeInfo = &proto.ChallengeInfo{
		ChallengeId:     challengeState.ChallengeId,
		Status:          challengeState.Status,
		RoundCount:      challengeState.RoundCount,
		ExtraLineupType: challengeState.ExtraLineupType,
		Score:           challengeState.ScoreOne,
		ScoreTwo:        challengeState.ScoreTwo,
		StoryInfo:       &proto.ChallengeStoryInfo{StoryBuffs: &proto.ChallengeStoryInfo_CurStoryBuffs{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: make([]uint32, 0)}}},
	}
	if challengeState.ChallengeCount == 1 {
		rsp.ChallengeInfo.StoryInfo.StoryBuffs = &proto.ChallengeStoryInfo_CurStoryBuffs{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: []uint32{challengeState.StoryBuffOne}}}
		// rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffOne)
	} else {
		rsp.ChallengeInfo.StoryInfo.StoryBuffs = &proto.ChallengeStoryInfo_CurStoryBuffs{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: []uint32{challengeState.StoryBuffTwo}}}
		// rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffTwo)
	}

	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

// 进入忘却之庭

func (g *GamePlayer) StartChallengeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartChallengeCsReq, payloadMsg)
	req := msg.(*proto.StartChallengeCsReq)
	battleState := g.GetBattleState()
	challengeState := g.NewChallengeState()

	// 设置战斗类型
	if req.StoryInfo != nil {
		battleState.BattleType = spb.BattleType_Battle_CHALLENGE_Story
		// 缓存buff
		// challengeState.StoryBuffOne = req.StoryInfo.StoryBuffInfo.StoryBuffOne
		// challengeState.StoryBuffTwo = req.StoryInfo.StoryBuffInfo.StoryBuffTwo
	} else {
		battleState.BattleType = spb.BattleType_Battle_CHALLENGE
	}

	// 从表中获取数据
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(strconv.Itoa(int(req.ChallengeId)))
	if challengeMazeConfig == nil {
		rsp := &proto.StartChallengeScRsp{
			Retcode: 2,
		}
		g.Send(cmd.StartChallengeScRsp, rsp)
		return
	}
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(challengeMazeConfig.MapEntranceID)))
	sceneGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, challengeMazeConfig.MazeGroupID1)
	if sceneGroup.AnchorList == nil {
		rsp := &proto.StartChallengeScRsp{
			Retcode: 2,
		}
		g.Send(cmd.StartChallengeScRsp, rsp)
		return
	}
	// 设置公共数据
	challengeState.ChallengeId = req.ChallengeId
	challengeState.Status = proto.ChallengeStatus_CHALLENGE_DOING
	challengeState.RoundCount = 0
	challengeState.ExtraLineupType = proto.ExtraLineupType_LINEUP_CHALLENGE
	// 获取坐标信息
	challengeState.Pos = &spb.VectorBin{
		X: int32(sceneGroup.AnchorList[0].PosX * 1000),
		Y: int32(sceneGroup.AnchorList[0].PosY * 1000),
		Z: int32(sceneGroup.AnchorList[0].PosZ * 1000),
	}
	challengeState.Rot = &spb.VectorBin{
		X: int32(sceneGroup.AnchorList[0].RotX * 1000),
		Y: int32(sceneGroup.AnchorList[0].RotY * 1000),
		Z: int32(sceneGroup.AnchorList[0].RotZ * 1000),
	}
	challengeState.NPCMonsterPos = &spb.VectorBin{
		X: int32(sceneGroup.MonsterList[0].PosX * 1000),
		Y: int32(sceneGroup.MonsterList[0].PosY * 1000),
		Z: int32(sceneGroup.MonsterList[0].PosZ * 1000),
	}
	challengeState.NPCMonsterRot = &spb.VectorBin{
		X: int32(sceneGroup.MonsterList[0].RotX * 1000),
		Y: int32(sceneGroup.MonsterList[0].RotY * 1000),
		Z: int32(sceneGroup.MonsterList[0].RotZ * 1000),
	}
	challengeState.PlaneID = mapEntrance.PlaneID
	challengeState.FloorID = mapEntrance.FloorID
	challengeState.EntranceID = challengeMazeConfig.MapEntranceID
	challengeState.ChallengeCount = challengeMazeConfig.StageNum
	challengeState.CurChallengeCount = 1
	// 添加场景buff到buff列表
	challengeState.SceneBuffList = append(challengeState.SceneBuffList, challengeMazeConfig.MazeBuffID)
	switch battleState.BattleType {
	case spb.BattleType_Battle_CHALLENGE:
		challengeState.ChallengeTargetID = challengeMazeConfig.ChallengeTargetID
		challengeState.ChallengeCountDown = challengeMazeConfig.ChallengeCountDown
	case spb.BattleType_Battle_CHALLENGE_Story:
		storyMazeExtra := gdconf.GetChallengeStoryMazeExtraById(challengeState.ChallengeId)
		challengeState.ChallengeCountDown = storyMazeExtra.TurnLimit
		challengeState.ChallengeTargetID = challengeMazeConfig.ChallengeTargetID
	}

	// 添加波次
	challengeState.CurChallengeBattle = make(map[uint32]*CurChallengeBattle)
	for id, challengeRoom := range challengeMazeConfig.ChallengeState {
		curChallengeBattle := &CurChallengeBattle{
			NPCMonsterID: challengeRoom.NPCMonsterID,
			EventID:      challengeRoom.EventID,
			GroupID:      challengeRoom.GroupID,
			ConfigID:     challengeRoom.ConfigID,
		}
		challengeState.CurChallengeBattle[id] = curChallengeBattle
	}

	// 下面是设置回包

	// 获取关卡信息
	challengeInfo := &proto.ChallengeInfo{
		ChallengeId:     challengeState.ChallengeId,
		Status:          challengeState.Status,
		RoundCount:      challengeState.RoundCount,
		ExtraLineupType: challengeState.ExtraLineupType,
		Score:           challengeState.ChallengeScore,
		StoryInfo:       &proto.ChallengeStoryInfo{StoryBuffs: &proto.ChallengeStoryInfo_CurStoryBuffs{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: make([]uint32, 0)}}},
	}
	challengeInfo.StoryInfo.StoryBuffs = &proto.ChallengeStoryInfo_CurStoryBuffs{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: []uint32{challengeState.StoryBuffOne}}}
	// challengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(challengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffOne)

	// 获取世界
	scene := g.GetChallengeScene()

	// 获取队伍
	lineup := g.GetLineUpPb(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE))

	rsp := &proto.StartChallengeScRsp{
		ChallengeInfo: challengeInfo,
		Scene:         scene,
		Lineup:        lineup,
	}

	g.Send(cmd.StartChallengeScRsp, rsp)
}

func (g *GamePlayer) GetChallengeScene() *proto.SceneInfo {
	challengeState := g.GetChallengeState()

	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	// 获取映射信息

	anchorPos := challengeState.Pos
	anchorRot := challengeState.Rot
	curChallengeBattle := challengeState.CurChallengeBattle[1]
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            challengeState.PlaneID,
		FloorId:            challengeState.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(challengeState.PlaneID))).WorldID,
		EntryId:            challengeState.EntranceID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(challengeState.PlaneID))).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}

	monsterEntity := make(map[uint32]*MonsterEntity, 0)
	avatarEntity := make(map[uint32]*AvatarEntity, 0)
	npcEntity := make(map[uint32]*NpcEntity, 0)
	// 将进入场景的角色添加到实体列表里
	entityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    0,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	curLineUp := g.GetBattleLineUpById(6)
	for id, lineAvatar := range curLineUp.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType(avatarBin.AvatarType), // TODO
				BaseAvatarId: avatarBin.AvatarId,
			},
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: anchorPos.X,
					Y: anchorPos.Y,
					Z: anchorPos.Z,
				},
				Rot: &proto.Vector{
					X: anchorRot.X,
					Y: anchorRot.Y,
					Z: anchorRot.Z,
				},
			},
		}
		// 为进入场景的角色设置与上面相同的实体id
		if id == curLineUp.LeaderSlot {
			entityList.EntityId = leaderEntityId
			avatarEntity[leaderEntityId] = &AvatarEntity{
				AvatarId: lineAvatar.AvatarId,
				GroupId:  0,
			}
		} else {
			avatarEntity[entityId] = &AvatarEntity{
				AvatarId: lineAvatar.AvatarId,
				GroupId:  0,
			}
			entityList.EntityId = entityId
		}
		entityGroup.EntityList = append(entityGroup.EntityList, entityList)
	}
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)

	// 添加怪物实体
	entityGroupNPCMonster := &proto.SceneEntityGroupInfo{
		GroupId:    curChallengeBattle.GroupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	entityId := uint32(g.GetNextGameObjectGuid())
	entityList := &proto.SceneEntityInfo{
		GroupId:  curChallengeBattle.GroupID,
		InstId:   curChallengeBattle.ConfigID,
		EntityId: entityId,
		Motion: &proto.MotionInfo{
			Pos: &proto.Vector{
				X: challengeState.NPCMonsterPos.X,
				Y: challengeState.NPCMonsterPos.Y,
				Z: challengeState.NPCMonsterPos.Z,
			},
			Rot: &proto.Vector{
				X: 0,
				Y: challengeState.NPCMonsterRot.Y,
				Z: 0,
			},
		},
		NpcMonster: &proto.SceneNpcMonsterInfo{
			WorldLevel: g.PlayerPb.WorldLevel,
			MonsterId:  curChallengeBattle.NPCMonsterID,
			EventId:    curChallengeBattle.EventID,
		},
	}
	monsterEntity[entityId] = &MonsterEntity{
		MonsterEId: curChallengeBattle.EventID,
		GroupId:    curChallengeBattle.GroupID,
	}
	entityGroupNPCMonster.EntityList = append(entityGroupNPCMonster.EntityList, entityList)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupNPCMonster)
	g.GetSceneEntity().MonsterEntity = monsterEntity
	g.GetSceneEntity().AvatarEntity = avatarEntity
	g.GetSceneEntity().NpcEntity = npcEntity

	return scene
}

// 忘却之庭战斗退出/结束

func (g *GamePlayer) LeaveChallengeCsReq(payloadMsg []byte) {
	if g.GetBattleState().ChallengeState.Status == proto.ChallengeStatus_CHALLENGE_DOING {
		g.Send(cmd.QuitBattleScNotify, nil)
	}
	g.Send(cmd.LeaveChallengeScRsp, nil)

	g.EnterSceneByServerScNotify(g.GetScene().EntryId, 0)
	g.GetBattleState().BattleType = spb.BattleType_Battle_NONE
	g.GetBattleState().BuffList = make([]uint32, 0)
}

// 忘却之庭世界发生攻击事件

func (g *GamePlayer) ChallengeSceneCastSkillCsReq(rsp *proto.SceneCastSkillScRsp) {
	// battleState := g.GetBattleState()
	challengeState := g.GetChallengeState()
	var lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE)
	var targetIndex uint32 = 0

	// 通过波次获取队伍
	if challengeState.ExtraLineupType == proto.ExtraLineupType_LINEUP_CHALLENGE {
		lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE)
	} else {
		lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2)
	}

	// 添加角色
	rsp.BattleInfo.BattleAvatarList = g.GetBattleAvatarList(lineUpId)
	// 添加回合限制
	rsp.BattleInfo.RoundsLimit = challengeState.ChallengeCountDown

	// 添加场景buff
	for _, buffId := range challengeState.SceneBuffList {
		buffList := &proto.BattleBuff{
			Id:       buffId,
			Level:    1,
			OwnerId:  4294967295,
			WaveFlag: 4294967295, // 失效时间
		}
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
	}
	// 添加角色buff
	for _, buffId := range challengeState.AvatarBuffList {
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

	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

func (g *GamePlayer) GetBattleAvatarList(lineUpId uint32) []*proto.BattleAvatar {
	battleAvatarList := make([]*proto.BattleAvatar, 0)
	for id, lineAvatar := range g.GetLineUpById(lineUpId).AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)

		battleAvatar := &proto.BattleAvatar{
			AvatarType:    proto.AvatarType_AVATAR_FORMAL_TYPE,
			Id:            avatarBin.AvatarId,
			Level:         avatarBin.Level,
			Rank:          avatarBin.Rank,
			Index:         id,
			SkilltreeList: make([]*proto.AvatarSkillTree, 0),
			Hp:            10000,
			Promotion:     avatarBin.PromoteLevel,
			RelicList:     make([]*proto.BattleRelic, 0),
			WorldLevel:    g.PlayerPb.WorldLevel,
			SpBar: &proto.SpBarInfo{
				CurSp: 10000,
				MaxSp: 10000,
			},
		}
		for _, skill := range g.GetSkillTreeList(avatarBin.AvatarId) {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			battleAvatar.SkilltreeList = append(battleAvatar.SkilltreeList, avatarSkillTree)
		}
		for _, relic := range avatarBin.EquipRelic {
			equipRelic := g.GetProtoBattleRelicById(relic)
			if equipRelic == nil {
				delete(avatarBin.EquipRelic, relic)
				continue
			}
			battleAvatar.RelicList = append(battleAvatar.RelicList, equipRelic)
		}
		// 获取角色装备的光锥
		if avatarBin.EquipmentUniqueId != 0 {
			equipment := g.GetEquipment(avatarBin.EquipmentUniqueId)
			equipmentList := &proto.BattleEquipment{
				Id:        equipment.Tid,
				Level:     equipment.Level,
				Promotion: equipment.Promotion,
				Rank:      equipment.Rank,
			}
			battleAvatar.EquipmentList = append(battleAvatar.EquipmentList, equipmentList)
		}
		battleAvatarList = append(battleAvatarList, battleAvatar)
	}
	return battleAvatarList
}

// 忘却之庭世界战斗结算事件

func (g *GamePlayer) ChallengePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq) {
	battleState := g.GetBattleState()
	challengeState := g.GetChallengeState()
	pos := challengeState.Pos
	rot := challengeState.Rot

	if challengeState.ExtraLineupType == proto.ExtraLineupType_LINEUP_CHALLENGE {
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE))
	} else {
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2))
	}

	// 战斗失败
	if req.EndStatus == proto.BattleEndStatus_BATTLE_END_LOSE {
		// 发送战斗状态
		challengeSettleNotify := &proto.ChallengeSettleNotify{
			ScoreTwo:       0,
			Stars:          0,
			Reward:         nil, // TODO 记得发奖励
			ChallengeId:    challengeState.ChallengeId,
			ChallengeScore: challengeState.ChallengeScore,
			IsWin:          false,
		}
		g.Send(cmd.ChallengeSettleNotify, challengeSettleNotify)
		return
	}

	// 删除实体
	nitify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshInfo: make([]*proto.SceneGroupRefreshInfo, 0),
	}
	for _, eventId := range challengeState.MonsterEntityMap {
		entity := g.GetSceneEntity().MonsterEntity[eventId]
		if entity != nil {
			groupRefreshInfo := &proto.SceneGroupRefreshInfo{
				GroupId: entity.GroupId,
				RefreshEntity: []*proto.SceneEntityRefreshInfo{
					{
						DelEntity: eventId,
					},
				},
			}
			nitify.GroupRefreshInfo = append(nitify.GroupRefreshInfo, groupRefreshInfo)
			delete(g.GetSceneEntity().MonsterEntity, eventId)
		}
	}
	g.Send(cmd.SceneGroupRefreshScNotify, nitify)

	// 获取已使用回合数
	challengeState.RoundCount += req.Stt.RoundCnt
	// 通过波次数判断是否还有一关
	if challengeState.CurChallengeCount == challengeState.ChallengeCount {
		// 战斗正常结束进入结算

		// 计算分数
		var stage uint32 = 0
		for _, challengeTargetID := range challengeState.ChallengeTargetID {
			challengeTargetConfig := gdconf.GetChallengeTargetConfigById(challengeTargetID)
			if challengeTargetConfig.ChallengeTargetType == "DEAD_AVATAR" {
				// 是否有角色死亡
				stage += 3
			} else {
				if (challengeState.ChallengeCountDown - challengeState.RoundCount) >= challengeTargetConfig.ChallengeTargetParam1 {
					stage += 2
				}
			}
		}

		// 将战斗结果储存到数据库
		challengeDb := g.GetChallengeById(challengeState.ChallengeId)
		if challengeDb.Stars < stage {
			challengeDb.Stars = stage
		}
		// 发送战斗胜利通知
		challengeSettleNotify := &proto.ChallengeSettleNotify{
			Stars:       stage,
			Reward:      nil, // TODO 记得发奖励
			ChallengeId: challengeState.ChallengeId,
			IsWin:       true,
		}
		g.Send(cmd.ChallengeSettleNotify, challengeSettleNotify)
		// 战斗正式结束，还原战斗信息
		battleState.BattleType = spb.BattleType_Battle_NONE
		challengeState.Status = proto.ChallengeStatus_CHALLENGE_FINISH
	} else {
		// 还差一波
		challengeState.CurChallengeCount++
		challengeState.ExtraLineupType = proto.ExtraLineupType_LINEUP_CHALLENGE_2
		challengeState.AvatarBuffList = make([]uint32, 0)
		// 添加怪物
		g.ChallengeAddSceneGroupRefreshScNotify()
		// 添加角色
		g.ChallengeAddAvatarSceneGroupRefreshScNotify()
		// 更新新的队伍
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2))
		// 通知当前战斗的队伍
		challengeLineupNotify := &proto.ChallengeLineupNotify{
			ExtraLineupType: challengeState.ExtraLineupType,
		}
		g.Send(cmd.ChallengeLineupNotify, challengeLineupNotify)
		// 通知坐标
		g.SceneEntityMoveScNotify(pos, rot, challengeState.EntranceID)
	}
}

func (g *GamePlayer) ChallengeSyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	lineUp := g.GetBattleLineUpById(index)
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           index,
		ExtraLineupType: proto.ExtraLineupType(lineUp.ExtraLineupType),
		MaxMp:           5,
		Mp:              5,
		PlaneId:         0,
	}
	for slot, lineAvatar := range lineUp.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: proto.AvatarType(avatarBin.AvatarType),
			Slot:       slot,
			Satiety:    0,
			Hp:         10000,
			Id:         avatarBin.AvatarId,
			SpBar: &proto.SpBarInfo{
				CurSp: 10000,
				MaxSp: 10000,
			},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsq.Lineup = lineupList

	g.Send(cmd.SyncLineupNotify, rsq)
}

func (g *GamePlayer) ChallengeAddAvatarSceneGroupRefreshScNotify() {
	challengeState := g.GetChallengeState()
	pos := challengeState.Pos
	rot := challengeState.Rot

	notify := new(proto.SceneGroupRefreshScNotify)
	notify.GroupRefreshInfo = make([]*proto.SceneGroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.SceneGroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}

	for _, lineAvatar := range g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2)).AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
		sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				Actor: &proto.SceneActorInfo{
					AvatarType:   proto.AvatarType(avatarBin.AvatarType),
					BaseAvatarId: avatarBin.AvatarId,
				},
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: pos.X,
						Y: pos.Y,
						Z: pos.Z,
					},
					Rot: &proto.Vector{
						X: rot.X,
						Y: rot.Y,
						Z: rot.Z,
					},
				},
				EntityId: entityId,
			},
		}
		g.GetSceneEntity().AvatarEntity[entityId] = &AvatarEntity{
			AvatarId: avatarBin.AvatarId,
			GroupId:  0,
		}
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	}
	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *GamePlayer) ChallengeAddSceneGroupRefreshScNotify() {
	challengeState := g.GetChallengeState()
	nPCMonsterPos := challengeState.NPCMonsterPos
	nPCMonsterRot := challengeState.NPCMonsterRot

	curChallengeBattle := challengeState.CurChallengeBattle[challengeState.CurChallengeCount]

	notify := new(proto.SceneGroupRefreshScNotify)
	notify.GroupRefreshInfo = make([]*proto.SceneGroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.SceneGroupRefreshInfo{
		GroupId:       curChallengeBattle.GroupID,
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}

	// 添加怪物实体
	entityId := uint32(g.GetNextGameObjectGuid())
	sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
		AddEntity: &proto.SceneEntityInfo{
			GroupId:  curChallengeBattle.GroupID,
			InstId:   curChallengeBattle.ConfigID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: nPCMonsterPos.X,
					Y: nPCMonsterPos.Y,
					Z: nPCMonsterPos.Z,
				},
				Rot: &proto.Vector{
					X: 0,
					Y: nPCMonsterRot.Y,
					Z: 0,
				},
			},
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.PlayerPb.WorldLevel,
				MonsterId:  curChallengeBattle.NPCMonsterID,
				EventId:    curChallengeBattle.EventID,
			},
		},
	}

	g.GetSceneEntity().MonsterEntity[entityId] = &MonsterEntity{
		MonsterEId: curChallengeBattle.EventID,
		GroupId:    curChallengeBattle.GroupID,
	}
	sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)

	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

// 下面是活动

func (g *GamePlayer) ChallengeStorySceneCastSkillCsReq(rsp *proto.SceneCastSkillScRsp) {
	challengeState := g.GetChallengeState()
	var lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE)
	var targetIndex uint32 = 0
	storyMazeExtra := gdconf.GetChallengeStoryMazeExtraById(challengeState.ChallengeId)

	// 通过波次获取队伍
	if challengeState.ExtraLineupType == proto.ExtraLineupType_LINEUP_CHALLENGE {
		lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE)
	} else {
		lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2)
	}

	// 添加角色
	rsp.BattleInfo.BattleAvatarList = g.GetBattleAvatarList(lineUpId)
	// 添加回合限制
	rsp.BattleInfo.RoundsLimit = challengeState.ChallengeCountDown

	// 添加关卡buff
	if challengeState.CurChallengeCount == 1 {
		buffListStory := &proto.BattleBuff{
			Id:              challengeState.StoryBuffOne,
			Level:           1,
			OwnerId:         0,
			TargetIndexList: []uint32{0},
			WaveFlag:        4294967295, // 失效时间
			DynamicValues:   make(map[string]float32),
		}
		buffListStory.DynamicValues["SkillIndex"] = 1
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffListStory)
	} else {
		buffListStory := &proto.BattleBuff{
			Id:              challengeState.StoryBuffTwo,
			Level:           1,
			OwnerId:         0,
			TargetIndexList: []uint32{0},
			WaveFlag:        4294967295, // 失效时间
			DynamicValues:   make(map[string]float32),
		}
		buffListStory.DynamicValues["SkillIndex"] = 1
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffListStory)
	}

	// 添加场景buff
	for _, buffId := range challengeState.SceneBuffList {
		buffList := &proto.BattleBuff{
			Id:       buffId,
			Level:    1,
			OwnerId:  4294967295,
			WaveFlag: 4294967295, // 失效时间
		}
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
	}
	// 添加角色buff
	for _, buffId := range challengeState.AvatarBuffList {
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
	rsp.BattleInfo.BattleTargetInfo = make(map[uint32]*proto.BattleTargetList)
	rsp.BattleInfo.BattleTargetInfo[1] = &proto.BattleTargetList{
		BattleTargetList: []*proto.BattleTarget{{
			Id: 10001,
		}},
	}
	battleTargetList := make([]*proto.BattleTarget, 0)
	for _, id := range storyMazeExtra.BattleTargetID {
		battleTarget := &proto.BattleTarget{
			Id:          id,
			Progress:    0,
			CDLKMKKOGLL: 0,
		}
		battleTargetList = append(battleTargetList, battleTarget)
	}
	rsp.BattleInfo.BattleTargetInfo[5] = &proto.BattleTargetList{
		BattleTargetList: battleTargetList,
	}

	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

func (g *GamePlayer) ChallengeStoryPVEBattleResultCsReq(req *proto.PVEBattleResultCsReq) {
	battleState := g.GetBattleState()
	challengeState := g.GetChallengeState()
	pos := challengeState.Pos
	rot := challengeState.Rot

	if challengeState.ExtraLineupType == proto.ExtraLineupType_LINEUP_CHALLENGE {
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE))
	} else {
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2))
	}

	// 删除实体
	nitify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshInfo: make([]*proto.SceneGroupRefreshInfo, 0),
	}
	for _, eventId := range challengeState.MonsterEntityMap {
		entity := g.GetSceneEntity().MonsterEntity[eventId]
		if entity != nil {
			groupRefreshInfo := &proto.SceneGroupRefreshInfo{
				GroupId: entity.GroupId,
				RefreshEntity: []*proto.SceneEntityRefreshInfo{
					{
						DelEntity: eventId,
					},
				},
			}
			nitify.GroupRefreshInfo = append(nitify.GroupRefreshInfo, groupRefreshInfo)
			delete(g.GetSceneEntity().MonsterEntity, eventId)
		}
	}
	g.Send(cmd.SceneGroupRefreshScNotify, nitify)

	// 获取分数
	challengeState.ChallengeScore += req.Stt.ChallengeScore
	if challengeState.CurChallengeCount == 1 {
		challengeState.ScoreOne = req.Stt.ChallengeScore
	} else {
		challengeState.ScoreTwo = req.Stt.ChallengeScore
	}
	// 通过波次数判断是否还有一关
	if challengeState.CurChallengeCount == challengeState.ChallengeCount {
		// 战斗正常结束进入结算

		// 计算分数
		var stage uint32 = 0
		if challengeState.ChallengeScore >= 30000 {
			stage++
		}
		for _, challengeTargetID := range challengeState.ChallengeTargetID {
			challengeTargetConfig := gdconf.GetChallengeTargetConfigById(challengeTargetID)
			if challengeState.ChallengeScore >= challengeTargetConfig.ChallengeTargetParam1 {
				stage += 2
			}
		}

		// 将战斗结果储存到数据库
		challengeDb := g.GetChallengeById(challengeState.ChallengeId)
		if challengeDb.Stars < stage {
			challengeDb.Stars = stage
		}
		if challengeDb.ScoreOne < challengeState.ScoreOne {
			challengeDb.ScoreOne = challengeState.ScoreOne
		}
		if challengeDb.ScoreTwo < challengeState.ScoreTwo {
			challengeDb.ScoreTwo = challengeState.ScoreTwo
		}
		// 发送战斗胜利通知
		challengeSettleNotify := &proto.ChallengeSettleNotify{
			Stars:          stage,
			Reward:         nil, // TODO 记得发奖励
			ChallengeId:    challengeState.ChallengeId,
			IsWin:          true,
			ChallengeScore: challengeState.ScoreOne,
			ScoreTwo:       challengeState.ScoreTwo,
		}
		g.Send(cmd.ChallengeSettleNotify, challengeSettleNotify)
		// 战斗正式结束，还原战斗信息
		battleState.BattleType = spb.BattleType_Battle_NONE
		challengeState.Status = proto.ChallengeStatus_CHALLENGE_FINISH
	} else {
		// 还差一波
		challengeState.CurChallengeCount++
		challengeState.ExtraLineupType = proto.ExtraLineupType_LINEUP_CHALLENGE_2
		challengeState.AvatarBuffList = make([]uint32, 0)
		// 添加怪物
		g.ChallengeAddSceneGroupRefreshScNotify()
		// 添加角色
		g.ChallengeAddAvatarSceneGroupRefreshScNotify()
		// 更新新的队伍
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2))
		// 通知当前战斗的队伍
		challengeLineupNotify := &proto.ChallengeLineupNotify{
			ExtraLineupType: challengeState.ExtraLineupType,
		}
		g.Send(cmd.ChallengeLineupNotify, challengeLineupNotify)
		// 通知坐标
		g.SceneEntityMoveScNotify(pos, rot, challengeState.EntranceID)
	}
}
