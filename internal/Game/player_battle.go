package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) SceneCastSkillCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneCastSkillCsReq, payloadMsg)
	req := msg.(*proto.SceneCastSkillCsReq)

	var targetIndex uint32 = 0
	var stageConfig *gdconf.StageConfig
	var stageID *gdconf.PlaneEvent

	if req.SkillIndex == 1 {
		avatarId := g.Player.EntityList[req.CasterId]
		g.Player.DbAvatar.Avatar[avatarId].BuffList = (avatarId * 100) + req.SkillIndex
		// 此处buff获取应该更改，不应这样获取，有的技能并不会增加buff而是回复生命等功能
	}
	if len(req.HitTargetIdList) == 0 {
		rsp := &proto.SceneCastSkillScRsp{
			AttackedGroupId: req.AttackedGroupId,
		}
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}

	if g.Player.EntityList[req.HitTargetIdList[0]] == 0 {
		rsp := &proto.SceneCastSkillScRsp{
			AttackedGroupId: req.AttackedGroupId,
		}
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	eventID := g.Player.EntityList[req.HitTargetIdList[0]]
	stageID = gdconf.GetPlaneEventById(eventID, g.Player.WorldLevel)
	if stageID == nil {
		event := g.Player.EntityList[req.CasterId]
		stageID = gdconf.GetPlaneEventById(event, g.Player.WorldLevel)
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
				MaxSp: avatar.SpBar.CurSp,
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
	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

func (g *Game) PVEBattleResultCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PVEBattleResultCsReq, payloadMsg)
	req := msg.(*proto.PVEBattleResultCsReq)
	/*
		SceneGroupRefreshScNotify
		{
		  "groupRefreshInfo": [{
		    "groupId": 16,
		    "refreshEntity": [{
		      "delEntity": 22
		    }]
		  }]
		}
	*/
	// 更新角色状态
	for _, avatar := range req.Stt.BattleAvatarList {
		g.Player.DbAvatar.Avatar[avatar.Id].Type = avatar.AvatarType
		g.Player.DbAvatar.Avatar[avatar.Id].SpBar.CurSp = uint32((avatar.AvatarStatus.LeftSp / avatar.AvatarStatus.MaxSp) * 10000)
		if avatar.AvatarStatus.LeftHp == 0 {
			g.Player.DbAvatar.Avatar[avatar.Id].Hp = 1000
			g.Player.DbAvatar.Avatar[avatar.Id].Type = proto.AvatarType_AVATAR_FORMAL_TYPE
		} else {
			g.Player.DbAvatar.Avatar[avatar.Id].Hp = uint32((avatar.AvatarStatus.LeftHp / avatar.AvatarStatus.MaxHp) * 10000)
		}
	}
	// 账号状态改变通知
	g.PlayerPlayerSyncScNotify()
	// 更新队伍状态
	g.SyncLineupNotify(g.Player.DbLineUp.MainLineUp)

	rsp := &proto.PVEBattleResultScRsp{
		BattleId:       req.BattleId,
		StageId:        req.StageId,
		EndStatus:      req.EndStatus, // 战斗结算状态
		CheckIdentical: true,          // 反作弊验证
		DropData: &proto.ItemList{ItemList: []*proto.Item{{
			ItemId:      0,
			Level:       0,
			Num:         0,
			MainAffixId: 0,
			Rank:        0,
			Promotion:   0,
			UniqueId:    0,
		}}},
		ResVersion: strconv.Itoa(int(req.ClientResVersion)),
	}
	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

func (g *Game) StartRogueCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartRogueCsReq, payloadMsg)
	req := msg.(*proto.StartRogueCsReq)

	rsp := &proto.StartRogueScRsp{
		Scene: nil,
		Lineup: &proto.LineupInfo{
			AvatarList:      make([]*proto.LineupAvatar, 0),
			ExtraLineupType: proto.ExtraLineupType_LINEUP_ROGUE,
			MaxMp:           5,
			Mp:              5,
		},
		RogueInfo: &proto.RogueInfo{
			RogueData:            nil,
			RogueVirtualItemInfo: nil,
			RogueScoreInfo:       nil,
			RoomMap:              nil,
			RogueAreaList:        nil,
			Status:               0,
			SeasonId:             0,
			RogueProgress:        nil,
			RogueAeonInfo:        nil,
			EndTime:              0,
			BaseAvatarIdList:     nil,
			BeginTime:            0,
			RogueCoin:            0,
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
				MaxSp: avatar.SpBar.CurSp,
			},
		}
		rsp.Lineup.AvatarList = append(rsp.Lineup.AvatarList, lineupAvatar)
	}

	g.Send(cmd.StartRogueScRsp, rsp)
}

func (g *Game) StartCocoonStageCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartCocoonStageCsReq, payloadMsg)
	req := msg.(*proto.StartCocoonStageCsReq)
	var targetIndex uint32 = 0

	rsp := new(proto.StartCocoonStageScRsp)
	rsp.PropEntityId = req.PropEntityId
	rsp.CocoonId = req.CocoonId
	rsp.Wave = req.Wave

	cocoonConfig := gdconf.GetCocoonConfigById(req.CocoonId, req.WorldLevel)
	stageConfig := gdconf.GetStageConfigById(cocoonConfig.StageID)

	rsp.BattleInfo = new(proto.SceneBattleInfo)
	rsp.BattleInfo.LogicRandomSeed = gdconf.GetLoadingDesc()
	rsp.BattleInfo.BattleId = g.GetBattleIdGuid() // 战斗Id
	rsp.BattleInfo.StageId = cocoonConfig.StageID

	// 怪物波列表
	for id, monsterListMap := range stageConfig.MonsterList {
		monsterWaveList := &proto.SceneMonsterWave{
			StageId: cocoonConfig.StageID,
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
				MaxSp: avatar.SpBar.CurSp,
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

	g.Send(cmd.StartCocoonStageScRsp, rsp)
}
