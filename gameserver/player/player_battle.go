package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) SceneCastSkillCostMpCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneCastSkillCostMpCsReq, payloadMsg)
	req := msg.(*proto.SceneCastSkillCostMpCsReq)
	rsp := &proto.SceneCastSkillCostMpScRsp{
		CastEntityId: req.CastEntityId,
		Retcode:      0,
	}
	g.Send(cmd.SceneCastSkillMpUpdateScNotify, &proto.SceneCastSkillMpUpdateScNotify{CastEntityId: req.CastEntityId})
	g.Send(cmd.SceneCastSkillCostMpScRsp, rsp)
}

func (g *GamePlayer) SceneEnterStageCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneEnterStageCsReq, payloadMsg)
	req := msg.(*proto.SceneEnterStageCsReq)
	battleInfo, battleBackup := g.GetSceneBattleInfo([]uint32{req.EventId}, g.GetBattleLineUp())
	rsp := &proto.SceneEnterStageScRsp{
		Retcode:    0,
		BattleInfo: battleInfo,
	}
	// 记录战斗
	battleBackup.EventId = req.EventId
	g.AddBattleBackup(battleBackup)
	g.Send(cmd.SceneEnterStageScRsp, rsp)
}

/***********************************攻击事件处理***********************************/

func (g *GamePlayer) SceneCastSkillCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneCastSkillCsReq, payloadMsg)
	req := msg.(*proto.SceneCastSkillCsReq)
	rsp := &proto.SceneCastSkillScRsp{
		CastEntityId: req.CastEntityId, // 攻击唯一id
	}
	// 根据各种情况进行处理
	if req.SkillIndex != 0 {
		// 这里的情况是角色释放技能
		g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
			GroupRefreshList: g.GetAddBuffSceneEntityRefreshInfo(req.AttackedByEntityId, g.GetRotPb(), g.GetPosPb()),
		})
		g.DelLineUpMp(1)
		g.HandleAvatarSkill(req.AttackedByEntityId, req.SkillIndex)
	}
	// 添加参与此次攻击的实体
	mpem := &MPEM{
		IsAvatar:        false,
		MonsterEntityId: make([]uint32, 0),
		MonsterId:       make([]uint32, 0),
		PropEntityId:    make([]uint32, 0),
		PropId:          make([]uint32, 0),
	}
	// 怪物协助列表
	g.GetMem([]uint32{req.AttackedByEntityId}, mpem)
	if req.AssistMonsterEntityInfo != nil {
		for _, info := range req.AssistMonsterEntityInfo {
			g.GetMem(info.EntityIdList, mpem)
		}
	} else {
		if req.AssistMonsterEntityIdList != nil {
			g.GetMem(req.AssistMonsterEntityIdList, mpem)
		}
	}
	if mpem.PropId != nil { // 物品效果
		g.SceneCastSkillProp(mpem)
	}
	g.SyncLineupNotify(g.GetBattleLineUp())               // 队伍同步
	if !mpem.IsAvatar || len(mpem.MonsterEntityId) == 0 { // 是否满足战斗条件
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	battleInfo, battleBackup := g.GetSceneBattleInfo(mpem.MonsterId, g.GetBattleLineUp())
	// 记录战斗
	battleBackup.monsterEntity = mpem.MonsterEntityId
	battleBackup.AttackedByEntityId = req.AttackedByEntityId
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
		AvatarBattleList: make([]*proto.BattleAvatar, 0),
		BattleId:         req.BattleId,
		StageId:          req.StageId,
		EndStatus:        req.EndStatus, // 战斗结算状态
		CheckIdentical:   true,          // 反作弊验证
		BinVersion:       "",
		ResVersion:       strconv.Itoa(int(req.ClientVersion)), // 版本验证
	}
	// 更新角色状态
	g.BattleUpAvatar(req.Stt.GetAvatarBattleList(), req.GetEndStatus())

	// 根据不同结算状态处理
	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_WIN: // 胜利
		// 删除怪物实体
		g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
			GroupRefreshList: g.GetDelSceneGroupRefreshInfo(battleBin.monsterEntity),
		})
		// 账号状态改变通知
		g.PlayerPlayerSyncScNotify()
		// 体力改变通知
		g.StaminaInfoScNotify()
		// 任务判断
		if battleBin.EventId != 0 {
			rsp.EventId = battleBin.EventId
			g.UpBattleSubMission(req.BattleId)
		}
		if battleBin.CocoonId != 0 { // 副本处理
			g.CocoonBattle(battleBin.CocoonId, battleBin.WorldLevel)
			g.FinishCocoon(battleBin.CocoonId)
		}
	case proto.BattleEndStatus_BATTLE_END_LOSE: // 失败
		teleportToAnchor = true
	case proto.BattleEndStatus_BATTLE_END_QUIT:
		teleportToAnchor = true
	}

	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		g.ChallengePVEBattleResultCsReq(req)
	case spb.BattleType_Battle_CHALLENGE_Story:
		g.ChallengePVEBattleResultCsReq(req)
	case spb.BattleType_Battle_ROGUE:
		teleportToAnchor = false
		g.RoguePVEBattleResultCsReq(req, len(battleBin.monsterEntity))
	case spb.BattleType_Battle_TrialActivity: // 角色试用
		g.TrialActivityPVEBattleResultScRsp(req)
		teleportToAnchor = true
	}

	// 是否传送到最近锚点
	if teleportToAnchor {
		// 当前坐标通知(移动到最近锚点)
		g.EnterSceneByServerScNotify(g.GetCurEntryId(), 0, 0, 0)
	}

	g.DelBattleBackupById(req.BattleId)
	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

/***********************************关卡/副本***********************************/

func (g *GamePlayer) StartCocoonStageCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartCocoonStageCsReq, payloadMsg)
	req := msg.(*proto.StartCocoonStageCsReq)
	g.SetBattleStatus(spb.BattleType_Battle_NONE) // 设置战斗状态
	battleInfo, battleBackup := g.GetCocoonBattleInfo(g.GetCurLineUp(), req)
	if battleInfo == nil {
		g.Send(cmd.StartCocoonStageScRsp, &proto.StartCocoonStageScRsp{Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)})
		return
	}
	rsp := &proto.StartCocoonStageScRsp{
		PropEntityId: req.PropEntityId,
		CocoonId:     req.CocoonId, // 关卡id
		Retcode:      0,
		Wave:         req.Wave,
		BattleInfo:   battleInfo,
	}
	// 储存战斗信息
	g.AddBattleBackup(battleBackup)
	g.Send(cmd.StartCocoonStageScRsp, rsp)
}

func (g *GamePlayer) ActivateFarmElementCsReq(payloadMsg []byte) {
	// msg := g.DecodePayloadToProto(cmd.ActivateFarmElementCsReq, payloadMsg)
	// req := msg.(*proto.ActivateFarmElementCsReq)
}

/***********************************物品破坏处理***********************************/

func (g *GamePlayer) SceneCastSkillProp(pem *MPEM) {
	for _, propId := range pem.PropId {
		conf := gdconf.GetMazePropId(propId)
		if conf == nil {
			continue
		}
		if conf.RecoverMp {
			g.AddLineUpMp(2) // 如果涉及到更新战斗中的队伍状态，这部分需要改
		}
		if conf.RecoverHp {
			g.AvatarRecoverPercent(pem.AvatarId, 0.3, 0)
		}
	}
}

func (g *GamePlayer) SyncEntityBuffChangeListScNotify(entityId uint32) {
	notify := &proto.SyncEntityBuffChangeListScNotify{
		EntityBuffChangeList: make([]*proto.EntityBuffChange, 0),
	}
	notify.EntityBuffChangeList = append(notify.EntityBuffChangeList, &proto.EntityBuffChange{
		EntityId: entityId,
	})
	g.Send(cmd.SyncEntityBuffChangeListScNotify, notify)
}
