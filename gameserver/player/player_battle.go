package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

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
	if req.SkillIndex != 0 { // 这里的情况是角色释放技能
		g.Send(cmd.SceneCastSkillScRsp, rsp)
	}
	if len(req.HitTargetEntityIdList) == 0 {
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	var mpem *MPEM
	mpem = g.GetMem(req.HitTargetEntityIdList) // 被攻击者
	if len(mpem.EntityId) == 0 {               // 这里的情况是，是怪物主动攻击发生的战斗
		mpem = g.GetMem([]uint32{req.AttackedByEntityId}) // 发起攻击者
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
	// 任务判断
	if battleBin.EventId != 0 {
		g.UpBattleSubMission(req)
	}
	// 副本处理
	g.CocoonBattle(battleBin.CocoonId, battleBin.WorldLevel)
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
