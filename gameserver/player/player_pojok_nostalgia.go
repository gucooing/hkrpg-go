package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

/******************************以太战线******************************/

func (g *GamePlayer) GetAetherDivideInfoCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetAetherDivideInfoScRsp{
		AvatarList:      make([]*proto.AetherDivideSpiritInfo, 0),
		LCNFCGHKACO:     1,
		OMKNCCGDKNP:     1,
		Retcode:         0,
		AetherSkillList: make([]*proto.AetherSkillInfo, 0),
		LineupList:      make([]*proto.AetherDivideLineupInfo, 0),
	}
	// add avatar
	for _, db := range g.GetPd().GetAetherDivideAvatar() {
		rsp.AvatarList = append(rsp.AvatarList,
			g.GetPd().GetAetherDivideSpiritInfo(db.AvatarId))
	}
	// add skill
	for _, db := range g.GetPd().GetAetherSkill() {
		rsp.AetherSkillList = append(rsp.AetherSkillList,
			g.GetPd().GetAetherSkillInfo(db.ItemId))
	}
	// add lineup
	for _, db := range g.GetPd().GetAetherDivideLineup() {
		rsp.LineupList = append(rsp.LineupList,
			g.GetPd().GetAetherDivideLineupInfo(db.Index))
	}

	g.Send(cmd.GetAetherDivideInfoScRsp, rsp)
}

func (g *GamePlayer) GetAetherDivideChallengeInfoCsReq(payloadMsg pb.Message) {
	x := make([]uint32, 0)
	for i := 1; i < 21; i++ {
		x = append(x, uint32(i))
	}
	rsp := &proto.GetAetherDivideChallengeInfoScRsp{
		// ECDFJJCPFJA:         1,
		// Retcode:             0,
		ANLOIAIEKHB: x,
		DEOOAOCGIIF: x,
	}

	g.Send(cmd.GetAetherDivideChallengeInfoScRsp, rsp)
}

func (g *GamePlayer) SetAetherDivideLineUpCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetAetherDivideLineUpCsReq)

	db := g.GetPd().GetAetherDivide()
	db.Lineup[req.Lineup.Slot] = &spb.AetherDivideLineup{
		Index:      req.Lineup.Slot,
		AvatarList: req.Lineup.AvatarList,
	}

	rsp := &proto.SetAetherDivideLineUpScRsp{
		Retcode: 0,
		Lineup:  req.Lineup,
	}

	g.Send(cmd.SetAetherDivideLineUpScRsp, rsp)
}

func (g *GamePlayer) EquipAetherDividePassiveSkillCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.EquipAetherDividePassiveSkillCsReq)

	rsp := &proto.EquipAetherDividePassiveSkillScRsp{
		Retcode: 0,
	}

	avatarDb := g.GetPd().GetAetherDivideAvatarInfoById(req.AvatarId)
	if avatarDb == nil {
		g.Send(cmd.EquipAetherDividePassiveSkillScRsp, rsp)
		return
	}
	if avatarDb.PassiveSkill == nil {
		avatarDb.PassiveSkill = make(map[uint32]uint32)
	}
	avatarDb.PassiveSkill[req.Slot] = req.ItemId
	skillDb := g.GetPd().GetAetherSkillById(req.ItemId)
	if skillDb == nil {
		g.Send(cmd.EquipAetherDividePassiveSkillScRsp, rsp)
		return
	}
	skillDb.DressAvatarId = req.AvatarId

	rsp.AvatarInfo = g.GetPd().GetAetherDivideSpiritInfo(req.AvatarId)
	rsp.AetherSkillInfo = g.GetPd().GetAetherSkillInfo(req.ItemId)

	g.Send(cmd.EquipAetherDividePassiveSkillScRsp, rsp)
}

func (g *GamePlayer) ClearAetherDividePassiveSkillCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ClearAetherDividePassiveSkillCsReq)

	rsp := &proto.ClearAetherDividePassiveSkillScRsp{
		Retcode: 0,
	}

	avatarDb := g.GetPd().GetAetherDivideAvatarInfoById(req.AvatarId)
	if avatarDb == nil {
		g.Send(cmd.EquipAetherDividePassiveSkillScRsp, rsp)
		return
	}
	if avatarDb.PassiveSkill == nil {
		avatarDb.PassiveSkill = make(map[uint32]uint32)
	}
	oldItemId := avatarDb.PassiveSkill[req.Slot]
	avatarDb.PassiveSkill[req.Slot] = 0

	skillDb := g.GetPd().GetAetherSkillById(oldItemId)
	if skillDb == nil {
		g.Send(cmd.EquipAetherDividePassiveSkillScRsp, rsp)
		return
	}
	skillDb.DressAvatarId = 0

	rsp.AvatarInfo = g.GetPd().GetAetherDivideSpiritInfo(req.AvatarId)
	rsp.AetherSkillInfo = g.GetPd().GetAetherSkillInfo(oldItemId)

	g.Send(cmd.ClearAetherDividePassiveSkillScRsp, rsp)
}

func (g *GamePlayer) AetherDivideTakeChallengeRewardCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.AetherDivideTakeChallengeRewardCsReq)

	rsp := &proto.AetherDivideTakeChallengeRewardScRsp{
		Retcode:     0,
		ChallengeId: req.ChallengeId,
		Reward:      &proto.ItemList{},
	}

	g.Send(cmd.AetherDivideTakeChallengeRewardScRsp, rsp)
}

func (g *GamePlayer) StartAetherDivideChallengeBattleCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartAetherDivideChallengeBattleCsReq)
	rsp := &proto.StartAetherDivideChallengeBattleScRsp{
		Retcode: 0,
	}
	conf := gdconf.GetAetherDivideChallengeList(req.ChallengeId)

	battleBackup := &model.BattleBackup{
		IsBattle:       true,
		EventId:        conf.EventID,
		WorldLevel:     g.GetPd().GetWorldLevel(),
		AetherDivideId: req.ChallengeId,
		Sce:            new(model.SceneCastEntity),
		BattleId:       g.GetPd().GetBattleIdGuid(),
	}
	planeEvent := gdconf.GetPlaneEventById(conf.EventID, battleBackup.WorldLevel)
	if planeEvent == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_BATTLE_STAGE_NOT_MATCH)
		g.Send(cmd.StartAetherDivideChallengeBattleScRsp, rsp)
		return
	}
	stageConfig := gdconf.GetStageConfigById(planeEvent.StageID)
	if stageConfig == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_BATTLE_STAGE_NOT_MATCH)
		g.Send(cmd.StartAetherDivideChallengeBattleScRsp, rsp)
		return
	}
	battleBackup.StageID = planeEvent.StageID
	battleBackup.Sce.EvenIdList = []uint32{conf.EventID}
	// avatar := make([]*model.AetherAvatar, 0)
	// if stageConfig.TrialAvatarList != nil {
	// 	for _, trial := range stageConfig.TrialAvatarList {
	// 		avatarConf := gdconf.GetAetherDivideSpiritTrial(trial)
	// 		if avatarConf == nil {
	// 			continue
	// 		}
	// 		avatar = append(avatar, &model.AetherAvatar{
	// 			AvatarId: trial,
	// 			Type:     spb.AvatarType_AVATAR_TRIAL_TYPE,
	// 		})
	// 	}
	// }
	// if len(avatar) == 0 && req.LineupIndex != 0 {
	//
	// }
	battleBackup.AetherAvatarList = g.GetPd().GetAetherAvatarrMap(stageConfig.TrialAvatarList)
	rsp.BattleInfo = g.GetPd().GetAetherDivideBattleInfo(battleBackup)
	// 记录战斗
	g.GetPd().AddBattleBackup(battleBackup)

	g.Send(cmd.StartAetherDivideChallengeBattleScRsp, rsp)
}

func (g *GamePlayer) StartAetherDivideSceneBattleCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartAetherDivideSceneBattleCsReq)
	battleBackup := &model.BattleBackup{
		Sce:        new(model.SceneCastEntity), // 添加参与此次攻击的实体
		WorldLevel: g.GetPd().GetWorldLevel(),
	}
	rsp := &proto.StartAetherDivideSceneBattleScRsp{
		CastEntityId: req.CastEntityId,
		BattleInfo:   nil,
		Retcode:      0,
	}
	// 添加攻击发起者
	g.GetPd().GetMem([]uint32{req.AttackedByEntityId}, battleBackup.Sce)
	// 添加被攻击者
	g.GetPd().GetMem(req.AssistMonsterEntityIdList, battleBackup.Sce)
	if len(battleBackup.Sce.EvenIdList) == 0 || !battleBackup.Sce.IsAvatar { // 是否满足战斗条件
		g.Send(cmd.StartAetherDivideSceneBattleScRsp, rsp)
		return
	}
	battleBackup.BattleId = g.GetPd().GetBattleIdGuid()
	battleBackup.AetherAvatarList = g.GetPd().GetAetherAvatarrMap(nil)
	rsp.BattleInfo = g.GetPd().GetAetherDivideBattleInfo(battleBackup)
	// 记录战斗
	g.GetPd().AddBattleBackup(battleBackup)

	g.Send(cmd.StartAetherDivideSceneBattleScRsp, rsp)
}

func (g *GamePlayer) StartAetherDivideStageBattleCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartAetherDivideStageBattleCsReq)
	battleBackup := &model.BattleBackup{
		Sce:        new(model.SceneCastEntity),
		WorldLevel: g.GetPd().GetWorldLevel(),
		BattleId:   g.GetPd().GetBattleIdGuid(),
		EventId:    req.EventId,
	}
	rsp := &proto.StartAetherDivideStageBattleScRsp{
		BattleInfo: nil,
		Retcode:    0,
	}
	planeEvent := gdconf.GetPlaneEventById(req.EventId, battleBackup.WorldLevel)
	if planeEvent == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_BATTLE_STAGE_NOT_MATCH)
		g.Send(cmd.StartAetherDivideStageBattleScRsp, rsp)
		return
	}
	stageConfig := gdconf.GetStageConfigById(planeEvent.StageID)
	if stageConfig == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_BATTLE_STAGE_NOT_MATCH)
		g.Send(cmd.StartAetherDivideStageBattleScRsp, rsp)
		return
	}
	battleBackup.StageID = planeEvent.StageID
	battleBackup.Sce.EvenIdList = []uint32{req.EventId}
	battleBackup.AetherAvatarList = g.GetPd().GetAetherAvatarrMap(stageConfig.TrialAvatarList)

	rsp.BattleInfo = g.GetPd().GetAetherDivideBattleInfo(battleBackup)
	// 记录战斗
	g.GetPd().AddBattleBackup(battleBackup)

	g.Send(cmd.StartAetherDivideStageBattleScRsp, rsp)
}

func (g *GamePlayer) LeaveAetherDivideSceneCsReq(payloadMsg pb.Message) {
	g.EnterSceneByServerScNotify(2013601, 0, 0, 0)
	g.Send(cmd.LeaveAetherDivideSceneScRsp, &proto.LeaveAetherDivideSceneScRsp{})
}

/******************************分割线******************************/
