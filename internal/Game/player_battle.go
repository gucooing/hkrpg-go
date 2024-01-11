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
	var buffLists []uint32
	lineUp := g.GetLineUp().MainLineUp
	battleState := g.GetBattleState()

	// 添加buff
	switch battleState.BattleType {
	case spb.BattleType_Battle_NONE:
		if req.SkillIndex == 1 {
			avatarId := g.Player.EntityList[req.CasterId].Entity
			skillId := (avatarId * 100) + req.SkillIndex
			if gdconf.GetMazeBuffById(skillId, req.SkillIndex) != nil {
				g.PlayerPb.Avatar.Avatar[avatarId].BuffList = skillId
			} else {
				// 技能处理，有的技能并不会增加buff而是回复生命等功能
			}
		}
	case spb.BattleType_Battle_ROGUE:
	case spb.BattleType_Battle_CHALLENGE:
		if req.SkillIndex == 1 {
			avatarId := g.Player.EntityList[req.CasterId].Entity
			skillId := (avatarId * 100) + req.SkillIndex
			if gdconf.GetMazeBuffById(skillId, req.SkillIndex) != nil {
				battleState.ChallengeState.AvatarBuffList = append(battleState.ChallengeState.AvatarBuffList, skillId)
			} else {
				// 技能处理，有的技能并不会增加buff而是回复生命等功能
			}
		}
	case spb.BattleType_Battle_CHALLENGE_Story:
		if req.SkillIndex == 1 {
			avatarId := g.Player.EntityList[req.CasterId].Entity
			skillId := (avatarId * 100) + req.SkillIndex
			if gdconf.GetMazeBuffById(skillId, req.SkillIndex) != nil {
				battleState.ChallengeState.AvatarBuffList = append(battleState.ChallengeState.AvatarBuffList, skillId)
			} else {
				// 技能处理，有的技能并不会增加buff而是回复生命等功能
			}
		}
	case spb.BattleType_Battle_TrialActivity:
		if req.SkillIndex == 1 {
			avatarId := g.Player.EntityList[req.CasterId].Entity
			skillId := (avatarId * 100) + req.SkillIndex
			if gdconf.GetMazeBuffById(skillId, req.SkillIndex) != nil {
				battleState.TrialActivityState.AvatarBuffList = append(battleState.TrialActivityState.AvatarBuffList, skillId)
			} else {
				// 技能处理，有的技能并不会增加buff而是回复生命等功能
			}
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
			BattleTargetInfo: make(map[uint32]*proto.BattleTargetList),
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
	switch battleState.BattleType {
	case spb.BattleType_Battle_NONE:

	case spb.BattleType_Battle_ROGUE:
		logger.Info("正在进行模拟宇宙")
	case spb.BattleType_Battle_CHALLENGE:
		// 缓存当前战斗实体
		battleState.ChallengeState.EventID = req.HitTargetEntityIdList[0]
		g.ChallengeSceneCastSkillCsReq(rsp)
		return
	case spb.BattleType_Battle_CHALLENGE_Story:
		// 缓存当前战斗实体
		battleState.ChallengeState.EventID = req.HitTargetEntityIdList[0]
		g.ChallengeStorySceneCastSkillCsReq(rsp)
		return
	case spb.BattleType_Battle_TrialActivity:
		g.TrialActivitySceneCastSkillScRsp(rsp)
		return
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
				DynamicValues:   make(map[string]float32),
			}
			buffList.DynamicValues["SkillIndex"] = 1
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
			DynamicValues:   make(map[string]float32),
		}
		buffList.DynamicValues["SkillIndex"] = 1
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
		targetIndex++
	}
	// 添加临时buff
	for _, buffId := range buffLists {
		if buffId == 0 {
			continue
		}
		buffList := &proto.BattleBuff{
			Id:              buffId,
			Level:           1,
			OwnerId:         targetIndex,
			TargetIndexList: []uint32{targetIndex},
			WaveFlag:        4294967295, // 失效时间
			DynamicValues:   make(map[string]float32),
		}
		buffList.DynamicValues["SkillIndex"] = 1
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

	rsp := &proto.PVEBattleResultScRsp{
		BattleId:       req.BattleId,
		StageId:        req.StageId,
		EndStatus:      req.EndStatus, // 战斗结算状态
		CheckIdentical: true,          // 反作弊验证
		BinVersion:     "",
		ResVersion:     strconv.Itoa(int(req.ClientResVersion)), // 版本验证
	}
	pos := g.GetPos()
	rot := g.GetRot()
	lineUp := g.GetLineUp().MainLineUp
	battleState := g.GetBattleState()

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
		g.ChallengePVEBattleResultCsReq(req)
		g.Send(cmd.PVEBattleResultScRsp, rsp)
		return
	case spb.BattleType_Battle_CHALLENGE_Story:
		// g.ChallengeStoryPVEBattleResultCsReq(req)
		g.ChallengeStoryPVEBattleResultCsReq(req)
		g.Send(cmd.PVEBattleResultScRsp, rsp)
		return
	case spb.BattleType_Battle_TrialActivity:
		g.TrialActivityPVEBattleResultScRsp(rsp)
		return

	}

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
	g.SceneEntityMoveScNotify(pos, rot, g.GetScene().EntryId)

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
func (g *Game) SceneEntityMoveScNotify(pos, rot *spb.VectorBin, entryId uint32) {
	if pos == nil {
		pos = g.GetPos()
	}
	if rot == nil {
		rot = g.GetRot()
	}

	notify := &proto.SceneEntityMoveScNotify{
		EntryId:          entryId,
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

/***********************************模拟宇宙***********************************/

func (g *Game) GetRogueInfoCsReq(payloadMsg []byte) {
	beginTime := time.Now().AddDate(0, 0, -1).Unix()
	endTime := beginTime + int64(time.Hour.Seconds()*24*8)
	rsp := new(proto.GetRogueInfoScRsp)
	rogueInfo := &proto.RogueInfo{
		RogueGetInfo: &proto.RogueGetInfo{
			RogueSeasonInfo: &proto.RogueSeasonInfo{
				BeginTime: 1704052800,
				EndTime:   4070894399,
				SeasonId:  78,
			},
			RogueScoreRewardInfo: &proto.RogueScoreRewardInfo{
				BeginTime:            beginTime,
				EndTime:              endTime,
				PoolId:               20 + g.PlayerPb.WorldLevel,
				HasTakenInitialScore: true,
				PoolRefreshed:        true,
			},
			RogueAreaList: &proto.RogueAreaList{RogueArea: make([]*proto.RogueArea, 0)},
			RogueData: &proto.RogueInfoData{
				RogueSeasonInfo: &proto.RogueSeasonInfo{
					BeginTime: beginTime,
					SeasonId:  78,
					EndTime:   endTime,
				},
				RogueScoreInfo: &proto.RogueScoreRewardInfo{
					PoolId:               20 + g.PlayerPb.WorldLevel,
					HasTakenInitialScore: true,
					PoolRefreshed:        true,
				},
			},
			RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
				RogueAbilityPoint: 0,
			},
			RoomMap:          nil,
			Status:           0,
			RogueProgress:    nil,
			RogueAeonInfo:    nil,
			BaseAvatarIdList: nil,
			RogueCoin:        5,
		},
	}
	for _, rogueArea := range gdconf.GetRogueAreaMap() {
		RogueArea := &proto.RogueArea{
			AreaId:          rogueArea.RogueAreaID,
			RogueAreaStatus: proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS,
		}
		rogueInfo.RogueGetInfo.RogueAreaList.RogueArea = append(rogueInfo.RogueGetInfo.RogueAreaList.RogueArea, RogueArea)
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
	/*
		rogueScoreInfo := &proto.RogueScoreRewardInfo{
			HasTakenInitialScore: true, // 已取得初始积分？
			Score:                0,
			PoolRefreshed:        true, // 刷新？
			TakenScoreRewardList: nil,
			PoolId:               20 + g.PlayerPb.WorldLevel,
		}
	*/
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
			RogueGetInfo: &proto.RogueGetInfo{
				RogueSeasonInfo: &proto.RogueSeasonInfo{
					BeginTime: 1704052800,
					EndTime:   4070894399,
					SeasonId:  78,
				},
				RogueScoreRewardInfo: &proto.RogueScoreRewardInfo{
					BeginTime:            beginTime,
					EndTime:              endTime,
					PoolId:               20 + g.PlayerPb.WorldLevel,
					HasTakenInitialScore: true,
					PoolRefreshed:        true,
				},
				RogueAreaList: &proto.RogueAreaList{RogueArea: make([]*proto.RogueArea, 0)},
				RogueData: &proto.RogueInfoData{
					RogueSeasonInfo: &proto.RogueSeasonInfo{
						BeginTime: 1704052800,
						EndTime:   4070894399,
						SeasonId:  78,
					},
					RogueScoreInfo: &proto.RogueScoreRewardInfo{
						BeginTime:            beginTime,
						EndTime:              endTime,
						PoolId:               20 + g.PlayerPb.WorldLevel,
						HasTakenInitialScore: true,
						PoolRefreshed:        true,
					},
				},
				RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
					RogueAbilityPoint: 0,
				},
				RoomMap:          nil,
				Status:           0,
				RogueProgress:    nil,
				RogueAeonInfo:    nil,
				BaseAvatarIdList: nil,
				RogueCoin:        0,
			},
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
