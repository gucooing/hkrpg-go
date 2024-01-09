package Game

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/***********************************大世界攻击事件处理***********************************/

func (g *Game) SceneCastSkillCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneCastSkillCsReq, payloadMsg)
	req := msg.(*proto.SceneCastSkillCsReq)

	var targetIndex uint32 = 0
	var stageConfig *gdconf.StageConfig
	var stageID *gdconf.PlaneEvent
	lineUp := g.GetLineUp().MainLineUp
	battleState := g.GetBattleState()

	if req.SkillIndex == 1 {
		avatarId := g.Player.EntityList[req.CasterId].Entity
		skillId := (avatarId * 100) + req.SkillIndex
		if gdconf.GetMazeBuffById(skillId, req.SkillIndex) != nil {
			g.PlayerPb.Avatar.Avatar[avatarId].BuffList = skillId
		} else {
			// 技能处理，有的技能并不会增加buff而是回复生命等功能
		}
	}
	if len(req.HitTargetEntityIdList) == 0 {
		rsp := &proto.SceneCastSkillScRsp{
			AttackedGroupId: req.AttackedGroupId,
		}
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}

	if g.Player.EntityList[req.HitTargetEntityIdList[0]] == nil {
		rsp := &proto.SceneCastSkillScRsp{
			AttackedGroupId: req.AttackedGroupId,
		}
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	entity := g.Player.EntityList[req.HitTargetEntityIdList[0]]
	stageID = gdconf.GetPlaneEventById(entity.Entity, g.PlayerPb.WorldLevel)
	if stageID == nil {
		newEntity := g.Player.EntityList[req.CasterId]
		stageID = gdconf.GetPlaneEventById(newEntity.Entity, g.PlayerPb.WorldLevel)
		stageConfig = gdconf.GetStageConfigById(stageID.StageID)
	} else {
		stageConfig = gdconf.GetStageConfigById(stageID.StageID)
	}

	// 构造回复包
	rsp := &proto.SceneCastSkillScRsp{
		AttackedGroupId: req.AttackedGroupId,
		BattleInfo: &proto.SceneBattleInfo{
			BuffList:        make([]*proto.BattleBuff, 0), // Buff列表
			LogicRandomSeed: gdconf.GetLoadingDesc(),      // 逻辑随机种子
			StageId:         stageID.StageID,              // 阶段id
			// TurnSnapshotList: nil,                          // 打开快照列表？
			WorldLevel:       g.PlayerPb.WorldLevel,
			RoundsLimit:      0,                              // 回合限制
			BattleId:         g.GetBattleIdGuid(),            // 战斗Id
			BattleAvatarList: make([]*proto.BattleAvatar, 0), // 战斗角色列表
		},
	}
	switch battleState.BattleType {
	case spb.BattleType_Battle_NONE:

	case spb.BattleType_Battle_ROGUE:
		logger.Info("正在进行模拟宇宙")
	case spb.BattleType_Battle_CHALLENGE:
		if battleState.ChallengeState.ExtraLineupType == proto.ExtraLineupType_LINEUP_CHALLENGE {
			lineUp = 6
		} else {
			lineUp = 7
		}
		rsp.BattleInfo.RoundsLimit = battleState.ChallengeState.ChallengeCountDown
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
	for id, avatarId := range g.GetLineUpById(lineUp).AvatarIdList {
		if avatarId == 0 {
			continue
		}
		avatar := g.GetAvatar().Avatar[avatarId]

		battleAvatar := &proto.BattleAvatar{
			AvatarType:    proto.AvatarType_AVATAR_FORMAL_TYPE,
			Id:            avatarId,
			Level:         avatar.Level,
			Rank:          avatar.Rank,
			Index:         uint32(id),
			SkilltreeList: make([]*proto.AvatarSkillTree, 0),
			Hp:            avatar.Hp,
			Promotion:     avatar.PromoteLevel,
			RelicList:     make([]*proto.BattleRelic, 0),
			WorldLevel:    g.PlayerPb.WorldLevel,
			SpBar: &proto.SpBarInfo{
				CurSp: avatar.SpBar.CurSp,
				MaxSp: avatar.SpBar.MaxSp,
			},
		}
		for _, skill := range g.GetSkillTreeList(avatar.AvatarId) {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			battleAvatar.SkilltreeList = append(battleAvatar.SkilltreeList, avatarSkillTree)
		}
		for _, relic := range avatar.EquipRelic {
			relicdb := g.GetRelicById(relic)
			equipRelic := &proto.BattleRelic{
				Id:           relicdb.Tid,
				Level:        relicdb.Level,
				MainAffixId:  relicdb.MainAffixId,
				SubAffixList: make([]*proto.RelicAffix, 0),
				UniqueId:     relicdb.UniqueId,
			}
			for _, subAddix := range relicdb.SubAffixList {
				relicAffix := &proto.RelicAffix{
					AffixId: subAddix.AffixId,
					Cnt:     subAddix.Cnt,
					Step:    subAddix.Step,
				}
				equipRelic.SubAffixList = append(equipRelic.SubAffixList, relicAffix)
			}
			battleAvatar.RelicList = append(battleAvatar.RelicList, equipRelic)
		}
		// 获取角色装备的光锥
		if avatar.EquipmentUniqueId != 0 {
			equipment := g.GetEquipment(avatar.EquipmentUniqueId)
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
			g.PlayerPb.Avatar.Avatar[avatarId].BuffList = 0
		}
	}

	// 检查场景buff

	for _, buffId := range g.Player.BattleState.BuffList {
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

	// 储存战斗信息
	g.Player.Battle = make(map[uint32]*Battle)
	battle := &Battle{
		BattleId:         rsp.BattleInfo.BattleId,
		EventID:          req.HitTargetEntityIdList[0],
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
	pos := g.GetPos()
	rot := g.GetRot()
	lineUp := g.GetLineUp().MainLineUp
	battleState := g.GetBattleState()

	// 要记得扣体力,和回复奖励
	/*
		主逻辑处理需要处理提前获取
		1.队伍
		2.坐标
	*/

	rsp.BattleId = req.BattleId
	rsp.StageId = req.StageId
	rsp.EndStatus = req.EndStatus // 战斗结算状态
	rsp.CheckIdentical = true     // 反作弊验证
	rsp.BinVersion = ""
	rsp.ResVersion = strconv.Itoa(int(req.ClientResVersion)) // 版本验证

	// 更新角色状态
	for _, avatarStt := range req.Stt.BattleAvatarList {
		avatar := g.GetAvatar().Avatar[avatarStt.Id]
		avatar.AvatarType = uint32(avatarStt.AvatarType)
		avatar.SpBar.CurSp = uint32((avatarStt.AvatarStatus.LeftSp / avatarStt.AvatarStatus.MaxSp) * 10000)
		if avatarStt.AvatarStatus.LeftHp == float64(0) {
			avatar.Hp = 2000
			avatar.AvatarType = uint32(proto.AvatarType_AVATAR_FORMAL_TYPE)
		} else {
			avatar.Hp = uint32((avatarStt.AvatarStatus.LeftHp / avatarStt.AvatarStatus.MaxHp) * 10000)
		}
	}

	switch battleState.BattleType {
	case spb.BattleType_Battle_NONE:
		// 撤退
		if req.EndStatus == proto.BattleEndStatus_BATTLE_END_QUIT {
			// 删除储存的战斗信息
			delete(g.Player.Battle, req.BattleId)
			g.Send(cmd.PVEBattleResultScRsp, rsp)
			return
		}
		// 更新队伍状态
		g.BattleSyncLineupNotify(lineUp)
	case spb.BattleType_Battle_ROGUE:
		logger.Info("正在进行模拟宇宙")
	case spb.BattleType_Battle_CHALLENGE:
		// 战斗胜利
		if req.EndStatus == proto.BattleEndStatus_BATTLE_END_WIN {
			battleState.ChallengeState.RoundCount += req.Stt.CocoonDeadWave
			if battleState.ChallengeState.CurChallengeCount == battleState.ChallengeState.ChallengeCount {
				// 战斗正常结束进入结算
				challengeDb := g.GetChallenge()
				var stage uint32 = 0
				for _, challengeTargetID := range battleState.ChallengeState.ChallengeTargetID {
					challengeTargetConfig := gdconf.GetChallengeTargetConfigById(challengeTargetID)
					if challengeTargetConfig.ChallengeTargetType == "DEAD_AVATAR" {
						// 是否有角色死亡
						stage += 3
					} else {
						if (battleState.ChallengeState.ChallengeCountDown - battleState.ChallengeState.RoundCount) >= challengeTargetConfig.ChallengeTargetParam1 {
							stage += 2
						}
					}
				}

				if challengeDb.ChallengeList[battleState.ChallengeState.ChallengeId] < stage {
					challengeDb.ChallengeList[battleState.ChallengeState.ChallengeId] = stage
				}
				challengeSettleNotify := &proto.ChallengeSettleNotify{
					Stars:       stage,
					Reward:      nil,
					ChallengeId: battleState.ChallengeState.ChallengeId,
					IsWin:       true,
				}
				g.Send(cmd.ChallengeSettleNotify, challengeSettleNotify)
				battleState.BattleType = spb.BattleType_Battle_NONE
				battleState.ChallengeState.Status = proto.ChallengeStatus_CHALLENGE_FINISH
				battleState.BuffList = make([]uint32, 0)
			} else {
				// 还差一波
				pos = battleState.ChallengeState.Pos
				rot = battleState.ChallengeState.Rot
				battleState.ChallengeState.CurChallengeCount++
				battleState.ChallengeState.ExtraLineupType = proto.ExtraLineupType_LINEUP_CHALLENGE_2
				g.HandleBattleChallenge()

				challengeLineupNotify := &proto.ChallengeLineupNotify{
					ExtraLineupType: proto.ExtraLineupType_LINEUP_CHALLENGE_2,
				}
				g.Send(cmd.ChallengeLineupNotify, challengeLineupNotify)
			}
		}
	}

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
							DelEntity: battle.EventID,
						},
					},
				},
			}
			g.Send(cmd.SceneGroupRefreshScNotify, nitify)
			delete(g.Player.EntityList, battle.EventID)
		}

		g.GetItem().MaterialMap[11] -= battle.StaminaCost * battle.Wave // 扣除体力

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
	g.SceneEntityMoveScNotify(pos, rot)

	// 体力改变通知
	g.StaminaInfoScNotify()

	// 删除储存的战斗信息
	delete(g.Player.Battle, req.BattleId)

	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

func (g *Game) HandleBattleNone() {

}

func (g *Game) HandleBattleRogue() {

}

func (g *Game) HandleBattleChallenge() {
	/*
		*1.删除原来的实体 1460 delEntity
		*2.发送奖励物品 652
		*3.通知账号信息
		*4.更新原来的队伍 709
		5.添加下一阶段的实体 1460 addEntity
		6.添加下一阶段队伍实体 1460 addEntity
		7.通知下一阶段队伍 709
		8.通知队伍变更 1770
		9.客户端位置更新通知 1409
		*10.正常回复pve包
	*/
	// 更新新的队伍
	g.BattleSyncLineupNotify(7)
	g.ChallengeAddSceneGroupRefreshScNotify()
}

// 队伍更新通知
func (g *Game) BattleSyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	lineUp := g.GetLineUpById(index)
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           index,
		ExtraLineupType: proto.ExtraLineupType(lineUp.ExtraLineupType),
		MaxMp:           5,
		Mp:              5,
		Name:            lineUp.Name,
		PlaneId:         0,
	}
	for slot, avatarId := range lineUp.AvatarIdList {
		if avatarId == 0 {
			continue
		}
		avatar := g.PlayerPb.Avatar.Avatar[avatarId]
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: proto.AvatarType(avatar.AvatarType),
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

// 当前坐标通知
func (g *Game) SceneEntityMoveScNotify(pos, rot *spb.VectorBin) {
	if pos == nil {
		pos = g.GetPos()
	}
	if rot == nil {
		rot = g.GetRot()
	}
	entryId := g.GetScene()

	notify := &proto.SceneEntityMoveScNotify{
		EntryId:          entryId.EntryId,
		ClientPosVersion: 0,
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
	}

	g.Send(cmd.SceneEntityMoveScNotify, notify)
}

func (g *Game) ChallengeAddSceneGroupRefreshScNotify() {
	challengeState := g.GetChallengeState()
	pos := challengeState.Pos
	rot := challengeState.Rot
	nPCMonsterPos := challengeState.NPCMonsterPos
	nPCMonsterRot := challengeState.NPCMonsterRot

	curChallengeBattle := challengeState.CurChallengeBattle[challengeState.CurChallengeCount]

	notify := new(proto.SceneGroupRefreshScNotify)
	notify.GroupRefreshInfo = make([]*proto.SceneGroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.SceneGroupRefreshInfo{
		GroupId:       curChallengeBattle.GroupID,
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}

	for _, avatarId := range g.GetLineUp().LineUpList[7].AvatarIdList {
		if avatarId == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				Actor: &proto.SceneActorInfo{
					AvatarType:   proto.AvatarType(g.GetAvatar().Avatar[avatarId].AvatarType),
					BaseAvatarId: avatarId,
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
		g.Player.EntityList[entityId] = &EntityList{
			Entity:  avatarId,
			GroupId: 0,
		}
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
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

	g.Player.EntityList[entityId] = &EntityList{
		Entity:  curChallengeBattle.EventID,
		GroupId: curChallengeBattle.GroupID,
	}
	sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)

	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
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
			PoolId:               20 + g.PlayerPb.WorldLevel,
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
				PoolId:               20 + g.PlayerPb.WorldLevel,
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
		req.BaseAvatarIdList = g.GetLineUpById(g.PlayerPb.LineUp.MainLineUp).AvatarIdList
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
		Score:                0,
		PoolRefreshed:        true, // 刷新？
		TakenScoreRewardList: nil,
		PoolId:               20 + g.PlayerPb.WorldLevel,
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
			Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
				BaseAvatarId: avatarId,
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

		// 添加物品实体
		propList := g.GetPropByID(sceneGroup, stou32(groupID))
		if len(propList.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, propList)
		}
		// 添加怪物实体
		nPCMonsterList, x := g.GetNPCMonsterByID(sceneGroup, stou32(groupID), entityMap)
		entityMap = x
		if len(nPCMonsterList.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, nPCMonsterList)
		}
		// 添加NPC实体
		nPCList := g.GetNPCByID(sceneGroup, stou32(groupID))
		if len(nPCList.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, nPCList)
		}
	}

	// 先更新队伍
	for id, avatarid := range req.BaseAvatarIdList {
		g.GetLineUpById(g.PlayerPb.LineUp.MainLineUp).AvatarIdList[id] = avatarid
	}
	g.PlayerPb.LineUp.MainAvatarId = 0
	// 队伍更新通知
	g.SyncLineupNotify(g.PlayerPb.LineUp.MainLineUp)

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
		avatar := g.PlayerPb.Avatar.Avatar[avatarId]
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: proto.AvatarType(avatar.AvatarType),
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
	g.GetBattleState().BattleType = spb.BattleType_Battle_ROGUE
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
	rsp.BattleInfo.WorldLevel = g.PlayerPb.WorldLevel

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
	for id, Lineup := range g.GetLineUpById(g.PlayerPb.LineUp.MainLineUp).AvatarIdList {
		if Lineup == 0 {
			continue
		}
		avatar := g.PlayerPb.Avatar.Avatar[Lineup]

		battleAvatar := &proto.BattleAvatar{
			AvatarType:    proto.AvatarType_AVATAR_FORMAL_TYPE,
			Id:            Lineup,
			Level:         avatar.Level,
			Rank:          avatar.Rank,
			Index:         uint32(id),
			SkilltreeList: make([]*proto.AvatarSkillTree, 0),
			Hp:            avatar.Hp,
			Promotion:     avatar.PromoteLevel,
			RelicList:     make([]*proto.BattleRelic, 0),
			WorldLevel:    g.PlayerPb.WorldLevel,
			SpBar: &proto.SpBarInfo{
				CurSp: avatar.SpBar.CurSp,
				MaxSp: avatar.SpBar.MaxSp,
			},
		}
		for _, skill := range g.GetSkillTreeList(avatar.AvatarId) {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			battleAvatar.SkilltreeList = append(battleAvatar.SkilltreeList, avatarSkillTree)
		}
		for _, relic := range avatar.EquipRelic {
			relicdb := g.GetRelicById(relic)
			equipRelic := &proto.BattleRelic{
				Id:           relicdb.Tid,
				Level:        relicdb.Level,
				MainAffixId:  relicdb.MainAffixId,
				SubAffixList: make([]*proto.RelicAffix, 0),
				UniqueId:     relicdb.UniqueId,
			}
			for _, subAddix := range relicdb.SubAffixList {
				relicAffix := &proto.RelicAffix{
					AffixId: subAddix.AffixId,
					Cnt:     subAddix.Cnt,
					Step:    subAddix.Step,
				}
				equipRelic.SubAffixList = append(equipRelic.SubAffixList, relicAffix)
			}
			battleAvatar.RelicList = append(battleAvatar.RelicList, equipRelic)
		}
		// 获取角色装备的光锥
		if avatar.EquipmentUniqueId != 0 {
			equipment := g.GetEquipment(avatar.EquipmentUniqueId)
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
			g.PlayerPb.Avatar.Avatar[Lineup].BuffList = 0
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
	for _, displayItem := range gdconf.GetMappingInfoById(req.CocoonId, g.PlayerPb.WorldLevel).DisplayItemList {
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

/***********************************忘却之庭***********************************/

func (g *Game) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurChallengeScRsp)

	challengeState := g.GetChallengeState()

	if challengeState != nil {
		rsp.ChallengeInfo = &proto.ChallengeInfo{
			ChallengeId:     challengeState.ChallengeId,
			Status:          challengeState.Status,
			RoundCount:      challengeState.RoundCount,
			ExtraLineupType: challengeState.ExtraLineupType,
		}
	}

	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

func (g *Game) StartChallengeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartChallengeCsReq, payloadMsg)
	req := msg.(*proto.StartChallengeCsReq)
	battleState := g.GetBattleState()
	challengeState := g.GetChallengeState()

	// 如果是新战斗就添加
	if battleState.BattleType != spb.BattleType_Battle_CHALLENGE {
		challengeState.ChallengeId = req.ChallengeId
		challengeState.Status = proto.ChallengeStatus_CHALLENGE_DOING
		challengeState.RoundCount = 0
		challengeState.ExtraLineupType = proto.ExtraLineupType_LINEUP_CHALLENGE
	}

	challengeInfo := &proto.ChallengeInfo{
		ChallengeId:     challengeState.ChallengeId,
		Status:          challengeState.Status,
		RoundCount:      challengeState.RoundCount,
		ExtraLineupType: challengeState.ExtraLineupType,
	}
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(strconv.Itoa(int(req.ChallengeId)))
	if challengeInfo == nil {
		rsp := &proto.StartChallengeScRsp{
			Retcode: 2,
		}
		g.Send(cmd.StartChallengeScRsp, rsp)
		return
	}
	// 获取映射信息
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(challengeMazeConfig.MapEntranceID)))
	sceneGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, challengeMazeConfig.MazeGroupID1)
	if sceneGroup.AnchorList == nil {
		rsp := &proto.StartChallengeScRsp{
			Retcode: 2,
		}
		g.Send(cmd.StartChallengeScRsp, rsp)
		return
	}
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
	challengeState.ChallengeTargetID = challengeMazeConfig.ChallengeTargetID
	challengeState.ChallengeCountDown = challengeMazeConfig.ChallengeCountDown
	battleState.BuffList = append(battleState.BuffList, challengeMazeConfig.MazeBuffID)
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

	scene := g.GetChallengeScene()

	lineup := g.GetLineUpPb(6)

	rsp := &proto.StartChallengeScRsp{
		ChallengeInfo: challengeInfo,
		Scene:         scene,
		Lineup:        lineup,
	}

	battleState.BattleType = spb.BattleType_Battle_CHALLENGE
	g.Send(cmd.StartChallengeScRsp, rsp)
}

func (g *Game) GetChallengeScene() *proto.SceneInfo {
	challengeState := g.GetChallengeState()

	entityMap := make(map[uint32]*EntityList) // [实体id]怪物群id

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

	// 将进入场景的角色添加到实体列表里
	entityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    0,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for id, slots := range g.GetLineUpPb(6).AvatarList {
		if slots == nil {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			Actor: &proto.SceneActorInfo{
				AvatarType:   slots.AvatarType, // TODO
				BaseAvatarId: slots.Id,
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
		if id == 0 {
			entityList.EntityId = leaderEntityId
			entityMap[leaderEntityId] = &EntityList{
				Entity:  slots.Slot,
				GroupId: 0,
			}
		} else {
			entityMap[entityId] = &EntityList{
				Entity:  slots.Slot,
				GroupId: 0,
			}
			entityList.EntityId = entityId
		}
		entityGroup.EntityList = append(entityGroup.EntityList, entityList)
	}
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)

	// 添加怪物实体
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
	entityMap[entityId] = &EntityList{
		Entity:  curChallengeBattle.EventID,
		GroupId: curChallengeBattle.GroupID,
	}
	entityGroup.EntityList = append(entityGroup.EntityList, entityList)

	g.Player.EntityList = entityMap
	return scene
}

func (g *Game) LeaveChallengeCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	if g.GetBattleState().ChallengeState.Status == proto.ChallengeStatus_CHALLENGE_DOING {
		g.Send(cmd.QuitBattleScNotify, rsp)
	}
	g.Send(cmd.LeaveChallengeScRsp, rsp)

	g.EnterSceneByServerScNotify(g.GetScene().EntryId, 0)
	g.GetBattleState().BattleType = spb.BattleType_Battle_NONE
	g.GetBattleState().BuffList = make([]uint32, 0)
}
