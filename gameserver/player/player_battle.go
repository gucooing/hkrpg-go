package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/***********************************攻击事件处理***********************************/

func (g *GamePlayer) SceneCastSkillCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneCastSkillCsReq, payloadMsg)
	req := msg.(*proto.SceneCastSkillCsReq)
	rsp := &proto.SceneCastSkillScRsp{
		AttackedGroupId: req.AttackedGroupId,
	}
	// 根据各种情况进行处理
	if req.SkillIndex != 0 { // 这里的情况是角色释放技能
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	var mpem *MPEM
	mpem = g.GetMem(req.HitTargetEntityIdList)
	if len(mpem.EntityId) == 0 { // 这里的情况是，是怪物主动攻击发生的战斗
		mpem = g.GetMem([]uint32{req.CasterId})
	}
	if len(mpem.EntityId) == 0 { // 这里的情况是角色普通攻击并没有命中怪物
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	if !mpem.IsBattle { // 不是战斗就要去处理物品效果了
		g.SceneCastSkillProp(mpem)
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	battleInfo, battleBackup := g.GetSceneBattleInfo(mpem.MPid, g.GetBattleLineUp())
	// 记录战斗
	battleBackup.monsterEntity = mpem.EntityId
	g.AddBattleBackup(battleBackup)
	// 回复
	rsp.BattleInfo = battleInfo
	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

/***********************************战斗结算***********************************/

func (g *GamePlayer) PVEBattleResultCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PVEBattleResultCsReq, payloadMsg)
	req := msg.(*proto.PVEBattleResultCsReq)
	battleBin := g.GetBattleBackupById(req.BattleId)
	if battleBin == nil {
		return
	}
	var teleportToAnchor = false
	rsp := &proto.PVEBattleResultScRsp{
		BattleAvatarList: make([]*proto.BattleAvatar, 0),
		BattleId:         req.BattleId,
		StageId:          req.StageId,
		EndStatus:        req.EndStatus, // 战斗结算状态
		CheckIdentical:   true,          // 反作弊验证
		BinVersion:       "",
		ResVersion:       strconv.Itoa(int(req.ClientResVersion)), // 版本验证
	}
	// 更新角色状态
	g.BattleUpAvatar(req.Stt.GetBattleAvatarList(), req.GetEndStatus())

	// 根据不同结算状态处理
	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_WIN: // 胜利
		// 删除怪物实体
		g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
			GroupRefreshInfo: g.GetDelSceneGroupRefreshInfo(battleBin.monsterEntity),
		})
		// 账号状态改变通知
		g.PlayerPlayerSyncScNotify()
		// 体力改变通知
		g.StaminaInfoScNotify()
	case proto.BattleEndStatus_BATTLE_END_LOSE: // 失败
		teleportToAnchor = true
	case proto.BattleEndStatus_BATTLE_END_QUIT:
		teleportToAnchor = true
	}

	// 是否传送到最近锚点
	if teleportToAnchor {
		// 当前坐标通知(移动到最近锚点)
		g.EnterSceneByServerScNotify(g.GetCurEntryId(), 0)
	}

	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		g.ChallengePVEBattleResultCsReq(req)
	case spb.BattleType_Battle_CHALLENGE_Story:
		g.ChallengePVEBattleResultCsReq(req)
	}

	g.DelBattleBackupById(req.BattleId)

	g.Send(cmd.PVEBattleResultScRsp, rsp)
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
	rsp.BattleInfo.WorldLevel = g.BasicBin.WorldLevel

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
			WorldLevel:    g.BasicBin.WorldLevel,
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
	g.OnlineData.Battle = make(map[uint32]*Battle)
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
	for _, displayItem := range gdconf.GetMappingInfoById(req.CocoonId, g.BasicBin.WorldLevel).DisplayItemList {
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
	g.OnlineData.Battle[rsp.BattleInfo.BattleId] = battle

	g.Send(cmd.StartCocoonStageScRsp, rsp)
}

/***********************************物品破坏处理***********************************/

func (g *GamePlayer) SceneCastSkillProp(pem *MPEM) {
	for _, propId := range pem.MPid {
		conf := gdconf.GetMazePropId(propId)
		if conf == nil {
			continue
		}
		if conf.RecoverMp {
			g.AddLineUpMp(2) // 如果涉及到更新战斗中的队伍状态，这部分需要改
		}
		if conf.RecoverHp {

		}
	}
}
