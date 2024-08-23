package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

/******************************以太战线******************************/

func (g *GamePlayer) GetAetherDivideInfoCsReq(payloadMsg pb.Message) {
	db := g.GetPd().GetAetherDivide()
	rsp := &proto.GetAetherDivideInfoScRsp{
		AvatarList:  make([]*proto.AetherDivideSpiritInfo, 0),
		NEOEJNKCKDM: 0,
		// KFBJEFGKIPH:     5,
		// CHKKOOPNPKG:     1,
		MMIJLNONOOI:     0,
		Retcode:         0,
		AetherSkillList: make([]*proto.AetherSkillInfo, 0),
		LineupList:      make([]*proto.AetherDivideLineupInfo, 0),
		// OHIMLBKKODO:     8014143,
	}
	// add avatar
	for _, avatar := range db.AvatarList {
		rsp.AvatarList = append(rsp.AvatarList, g.GetPd().GetAetherDivideSpiritInfo(avatar.AvatarId))
	}
	// add skill
	for _, skill := range db.AetherSkillList {
		rsp.AetherSkillList = append(rsp.AetherSkillList, g.GetPd().GetAetherSkillInfo(skill.ItemId))
	}
	// add lineup
	for _, lineup := range db.Lineup {
		rsp.LineupList = append(rsp.LineupList, g.GetPd().GetAetherDivideLineupInfo(lineup.Index))
	}

	g.Send(cmd.GetAetherDivideInfoScRsp, rsp)
}

func (g *GamePlayer) GetAetherDivideChallengeInfoCsReq(payloadMsg pb.Message) {
	x := make([]uint32, 0)
	for i := 1; i < 21; i++ {
		x = append(x, uint32(i))
	}
	rsp := &proto.GetAetherDivideChallengeInfoScRsp{
		// FIBDNGJJDFD:         1,
		// Retcode:             0,
		EELDBMONEIM:         x,
		FinishChallengeList: x,
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
	avatar := make([]*model.AetherAvatar, 0)
	if stageConfig.TrialAvatarList != nil {
		for _, trial := range stageConfig.TrialAvatarList {
			avatarConf := gdconf.GetAetherDivideSpiritTrial(trial)
			if avatarConf == nil {
				continue
			}
			avatar = append(avatar, &model.AetherAvatar{
				AvatarId: trial,
				Type:     spb.AvatarType_AVATAR_TRIAL_TYPE,
			})
		}
	}
	if len(avatar) == 0 && req.LineupIndex != 0 {
		lineup := g.GetPd().GetAetherDivideLineupInfo(req.LineupIndex)
		if lineup != nil {
			for _, avatarId := range lineup.AvatarList {
				avatar = append(avatar, &model.AetherAvatar{
					AvatarId: avatarId,
					Type:     spb.AvatarType_AVATAR_TYPE_NONE,
				})
			}
		}
	}
	monsterWaveList, _ := g.GetPd().GetSceneMonsterWave([]uint32{conf.EventID}, battleBackup.WorldLevel, battleBackup)
	info := &proto.AetherDivideBattleInfo{
		BattleAvatarList: g.GetPd().GetAetherAvatarInfoList(avatar),
		MonsterWaveList:  monsterWaveList,
		BattleId:         battleBackup.BattleId,
		BuffList:         make([]*proto.BattleBuff, 0),
		LogicRandomSeed:  gdconf.GetLoadingDesc(),
		StageId:          battleBackup.StageID,
	}
	rsp.BattleInfo = info
	// 记录战斗
	g.GetPd().AddBattleBackup(battleBackup)

	g.Send(cmd.StartAetherDivideChallengeBattleScRsp, rsp)
}

/******************************分割线******************************/
