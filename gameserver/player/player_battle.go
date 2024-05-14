package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/***********************************大世界攻击事件处理***********************************/

func (g *GamePlayer) SceneCastSkillCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneCastSkillCsReq, payloadMsg)
	req := msg.(*proto.SceneCastSkillCsReq)
	var mem []uint32 // 目标怪物列表
	// 根据攻击发起者的信息确定怪物实体
	entity := g.GetEntityById(req.CasterId)
	switch entity.(type) {
	case *MonsterEntity:
		mem = append(mem, req.CasterId)
	case *AvatarEntity:
		mem = g.GetMem(req.HitTargetEntityIdList)
	}
	if len(mem) == 0 {
		rsp := &proto.SceneCastSkillScRsp{
			AttackedGroupId: req.AttackedGroupId,
		}
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}

	// 构造回复包
	rsp := &proto.SceneCastSkillScRsp{
		AttackedGroupId: req.AttackedGroupId,
		BattleInfo:      g.GetSceneBattleInfo(mem, g.GetBattleLineUp()),
	}
	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

/***********************************战斗结算***********************************/

func (g *GamePlayer) PVEBattleResultCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PVEBattleResultCsReq, payloadMsg)
	req := msg.(*proto.PVEBattleResultCsReq)
	var pileItem []*Material

	rsp := &proto.PVEBattleResultScRsp{
		BattleAvatarList: make([]*proto.BattleAvatar, 0),
		BattleId:         req.BattleId,
		StageId:          req.StageId,
		EndStatus:        req.EndStatus, // 战斗结算状态
		CheckIdentical:   true,          // 反作弊验证
		BinVersion:       "",
		ResVersion:       strconv.Itoa(int(req.ClientResVersion)), // 版本验证
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
		g.RoguePVEBattleResultCsReq(req, rsp)
		return
	case spb.BattleType_Battle_CHALLENGE:
		g.ChallengePVEBattleResultCsReq(req)
		g.Send(cmd.PVEBattleResultScRsp, rsp)
		return
	// case spb.BattleType_Battle_CHALLENGE_Story:
	// 	g.ChallengeStoryPVEBattleResultCsReq(req)
	// 	g.Send(cmd.PVEBattleResultScRsp, rsp)
	// 	return
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

		// 删除实体
		nitify := &proto.SceneGroupRefreshScNotify{
			GroupRefreshInfo: make([]*proto.SceneGroupRefreshInfo, 0),
		}
		// for _, eventId := range battle.EventIDList {
		// 	entity := g.GetSceneEntity().MonsterEntity[eventId]
		// 	if entity != nil {
		// 		groupRefreshInfo := &proto.SceneGroupRefreshInfo{
		// 			GroupId: entity.GroupId,
		// 			RefreshEntity: []*proto.SceneEntityRefreshInfo{
		// 				{
		// 					DelEntity: eventId,
		// 				},
		// 			},
		// 		}
		// 		nitify.GroupRefreshInfo = append(nitify.GroupRefreshInfo, groupRefreshInfo)
		// 		delete(g.GetSceneEntity().MonsterEntity, eventId)
		// 	}
		// }
		g.Send(cmd.SceneGroupRefreshScNotify, nitify)

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

			pileItem = append(pileItem, &Material{
				Tid: drop.Tid,
				Num: drop.Num * battle.Wave,
			})
		}
	}

	g.AddMaterial(pileItem)

	// 当前坐标通知(失败情况应该是移动到最近锚点)
	g.SceneEntityMoveScNotify(pos, rot, g.GetScene().EntryId)

	// 体力改变通知
	g.StaminaInfoScNotify()

	// 删除储存的战斗信息
	delete(g.Player.Battle, req.BattleId)

	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

// 队伍更新通知
func (g *GamePlayer) BattleSyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	lineUp := g.GetLineUpById(index)
	lineupList := &proto.LineupInfo{
		IsVirtual:  false,
		LeaderSlot: 0,
		AvatarList: make([]*proto.LineupAvatar, 0),
		Index:      index,
		// ExtraLineupType: proto.ExtraLineupType(lineUp.ExtraLineupType),
		MaxMp:   5,
		Mp:      5,
		Name:    lineUp.Name,
		PlaneId: 0,
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
			Hp:         avatarBin.Hp,
			Id:         lineAvatar.AvatarId,
			SpBar: &proto.SpBarInfo{
				CurSp: avatarBin.SpBar.CurSp,
				MaxSp: avatarBin.SpBar.MaxSp,
			},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsq.Lineup = lineupList

	g.Send(cmd.SyncLineupNotify, rsq)
}

// 当前坐标通知
func (g *GamePlayer) SceneEntityMoveScNotify(pos, rot *spb.VectorBin, entryId uint32) {
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

/***********************************关卡/副本***********************************/

func (g *GamePlayer) StartCocoonStageCsReq(payloadMsg []byte) {
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
	curLineUp := g.GetCurLineUp()
	for id, lineAvatar := range curLineUp.AvatarIdList {
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
			Hp:            avatarBin.Hp,
			Promotion:     avatarBin.PromoteLevel,
			RelicList:     make([]*proto.BattleRelic, 0),
			WorldLevel:    g.PlayerPb.WorldLevel,
			SpBar: &proto.SpBarInfo{
				CurSp: avatarBin.SpBar.CurSp,
				MaxSp: avatarBin.SpBar.MaxSp,
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
			if equipment == nil {
				avatarBin.EquipmentUniqueId = 0
			} else {
				equipmentList := &proto.BattleEquipment{
					Id:        equipment.Tid,
					Level:     equipment.Level,
					Promotion: equipment.Promotion,
					Rank:      equipment.Rank,
				}
				battleAvatar.EquipmentList = append(battleAvatar.EquipmentList, equipmentList)
			}
		}
		rsp.BattleInfo.BattleAvatarList = append(rsp.BattleInfo.BattleAvatarList, battleAvatar)
		// 检查是否有提前释放的技能，添加到buff里
		if avatarBin.BuffList != 0 {
			buffList := &proto.BattleBuff{
				Id:              avatarBin.BuffList,
				Level:           1,
				OwnerId:         targetIndex,
				TargetIndexList: []uint32{targetIndex},
				WaveFlag:        4294967295, // 失效时间
			}
			rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
			targetIndex++
			avatarBin.BuffList = 0
		}
	}

	// 储存战斗信息
	g.Player.Battle = make(map[uint32]*Battle)
	battle := &Battle{
		BattleId: rsp.BattleInfo.BattleId,
		Wave:     req.Wave,
		// EventIDList:          req.CocoonId,
		LogicRandomSeed:  rsp.BattleInfo.LogicRandomSeed,
		RoundsLimit:      rsp.BattleInfo.RoundsLimit,
		StaminaCost:      cocoonConfig.StaminaCost,
		BuffList:         rsp.BattleInfo.BuffList,
		BattleAvatarList: rsp.BattleInfo.BattleAvatarList,
	}
	battle.EventIDList = append(battle.EventIDList, req.CocoonId)
	// 添加奖励 TODO 发送的奖励有问题，所以暂时注释掉（有时间再康
	for _, displayItem := range gdconf.GetMappingInfoById(req.CocoonId, g.PlayerPb.WorldLevel).DisplayItemList {
		itemConf := gdconf.GetItemConfigById(displayItem.ItemID)
		if itemConf == nil {
			continue
		}
		material := &Material{}
		if material.Num != 0 {
			material.Tid = displayItem.ItemID
			material.Num = displayItem.ItemNum
		} else {
			switch itemConf.ItemSubType {
			case "Virtual":
				material.Tid = displayItem.ItemID
				material.Num = 5000
			case "Material":
				material.Tid = displayItem.ItemID
				material.Num = 5
			default:
				continue
			}
		}
		battle.DisplayItemList = append(battle.DisplayItemList, material)
	}
	g.Player.Battle[rsp.BattleInfo.BattleId] = battle

	g.Send(cmd.StartCocoonStageScRsp, rsp)
}
