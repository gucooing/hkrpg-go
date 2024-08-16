package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) SetTurnFoodSwitchCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetTurnFoodSwitchCsReq)
	rsp := &proto.SetTurnFoodSwitchScRsp{
		Retcode:     0,
		DGLLJFNEMOK: req.DGLLJFNEMOK,
		LNEKBEGKACP: req.LNEKBEGKACP,
	}
	g.Send(cmd.SetTurnFoodSwitchScRsp, rsp)
}

func (g *GamePlayer) SceneCastSkillCostMpCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneCastSkillCostMpCsReq)
	rsp := &proto.SceneCastSkillCostMpScRsp{
		CastEntityId: req.CastEntityId,
		Retcode:      0,
	}
	g.Send(cmd.SceneCastSkillCostMpScRsp, rsp)
}

func (g *GamePlayer) SceneEnterStageCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneEnterStageCsReq)
	rsp := &proto.SceneEnterStageScRsp{
		Retcode: 0,
	}
	// 获取战斗角色
	avatarMap := make(map[uint32]*BattleAvatar, 0)
	stageConfig := gdconf.GetStageConfigById(req.EventId)
	if stageConfig == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_BATTLE_STAGE_NOT_MATCH)
		g.Send(cmd.SceneEnterStageScRsp, rsp)
		return
	}
	if stageConfig.TrialAvatarList != nil {
		for _, trial := range stageConfig.TrialAvatarList {
			avatarMap[trial] = &BattleAvatar{
				AvatarId:   trial,
				AvatarType: spb.LineAvatarType_LineAvatarType_TRIAL,
				AssistUid:  0,
			}
		}
	}
	if len(avatarMap) == 0 {
		lineUp := g.GetBattleLineUp()
		for index, avatar := range lineUp.AvatarIdList {
			avatarMap[avatar.AvatarId] = &BattleAvatar{
				AvatarId:   avatar.AvatarId,
				AvatarType: avatar.LineAvatarType,
				AssistUid:  0,
				Index:      index,
			}
		}
	}
	battleInfo, battleBackup := g.GetSceneBattleInfo([]uint32{req.EventId}, nil, avatarMap, g.GetWorldLevel(), 0)
	rsp.BattleInfo = battleInfo
	// 记录战斗
	battleBackup.EventId = req.EventId
	g.AddBattleBackup(battleBackup)
	g.Send(cmd.SceneEnterStageScRsp, rsp)
}

/***********************************攻击事件处理***********************************/

func (g *GamePlayer) SceneCastSkillCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneCastSkillCsReq)
	rsp := &proto.SceneCastSkillScRsp{
		CastEntityId: req.CastEntityId, // 攻击唯一id
	}
	isBattle := false
	// 添加参与此次攻击的实体
	sce := &SceneCastEntity{
		IsAvatar:            false,
		MonsterEntityIdList: make([]uint32, 0),
		MonsterIdList:       make([]uint32, 0),
		PropEntityIdList:    make([]uint32, 0),
		PropIdList:          make([]uint32, 0),
	}
	// HitTargetEntityIdList 被攻击目标
	// AssistMonsterEntityIdList 击中列表
	// AssistMonsterEntityInfo 击中敌人群
	// AttackedByEntityId 攻击发起者

	// 添加攻击发起者
	g.GetMem([]uint32{req.AttackedByEntityId}, sce)
	// 添加被攻击者
	if req.AssistMonsterEntityInfo != nil {
		for _, info := range req.AssistMonsterEntityInfo {
			g.GetMem(info.EntityIdList, sce)
		}
	} else {
		g.GetMem(req.AssistMonsterEntityIdList, sce)
	}

	g.SceneCastSkillProp(sce) // 物品效果
	var skill *gdconf.GoppMazeSkill
	if req.SkillIndex != 0 {
		skill = gdconf.GetGoppMazeSkill(sce.AvatarId, 2)
		g.DelMp(sce.AvatarId, req.CastEntityId)
	} else {
		skill = gdconf.GetGoppMazeSkill(sce.AvatarId, 1)
	}
	if skill == nil {
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	isBattle = g.sceneCastSkill(sce, skill, req)
	if len(sce.MonsterIdList) == 0 || !isBattle { // 是否满足战斗条件
		g.Send(cmd.SceneCastSkillScRsp, rsp)
		return
	}
	// 获取战斗角色
	avatarMap := make(map[uint32]*BattleAvatar, 0)
	lineUp := g.GetBattleLineUp()
	for index, avatar := range lineUp.AvatarIdList {
		avatarMap[avatar.AvatarId] = &BattleAvatar{
			AvatarId:   avatar.AvatarId,
			AvatarType: avatar.LineAvatarType,
			AssistUid:  0,
			Index:      index,
		}
	}
	battleInfo, battleBackup := g.GetSceneBattleInfo(sce.MonsterIdList, nil, avatarMap, g.GetWorldLevel(), 0)
	// 记录战斗
	battleBackup.monsterEntity = sce.MonsterEntityIdList
	battleBackup.AttackedByEntityId = req.AttackedByEntityId
	g.AddBattleBackup(battleBackup)
	// 回复
	rsp.BattleInfo = battleInfo
	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

/***********************************战斗结算***********************************/

func (g *GamePlayer) PVEBattleResultCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PVEBattleResultCsReq)
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
		DropData:         &proto.ItemList{ItemList: make([]*proto.Item, 0)},
	}
	// 更新角色状态
	g.BattleUpAvatar(req.Stt.GetBattleAvatarList(), req.GetEndStatus())

	allSync := &AllPlayerSync{
		IsBasic:       true,
		MaterialList:  make([]uint32, 0),
		EquipmentList: nil,
		RelicList:     nil,
	}
	addPileItem := make([]*Material, 0)
	delPileItem := make([]*Material, 0)

	// 根据不同结算状态处理
	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_WIN: // 胜利
		// 删除怪物实体
		g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
			GroupRefreshList: g.GetDelSceneGroupRefreshInfo(battleBin.monsterEntity),
		})
		// 任务判断
		if battleBin.EventId != 0 {
			rsp.EventId = battleBin.EventId
			g.UpBattleSubMission(req.BattleId)
			if req.Stt.CustomValues != nil {
				g.BattleCustomValues(req.Stt.CustomValues, battleBin.EventId)
			}
		}
		// 参战角色经验添加
		for _, avatar := range req.Stt.GetBattleAvatarList() {
			if _, ok := g.AvatarAddExp(avatar.Id, battleBin.AvatarExpReward); ok {
				allSync.AvatarList = append(allSync.AvatarList, avatar.Id)
			}
		}
		// 获取奖励
		addPileItem = append(addPileItem, battleBin.DisplayItemList...)
		for _, displayItem := range battleBin.DisplayItemList {
			allSync.MaterialList = append(allSync.MaterialList, displayItem.Tid)
			rsp.DropData.ItemList = append(rsp.DropData.ItemList, &proto.Item{
				ItemId: displayItem.Tid,
				Num:    displayItem.Num,
			})
		}
		if conf := gdconf.GetCocoonConfigById(battleBin.CocoonId, battleBin.WorldLevel); conf != nil { // 副本处理
			rsp.DropData.ItemList = append(rsp.DropData.ItemList,
				g.getBattleDropData(conf.MappingInfoID, allSync, addPileItem, battleBin.WorldLevel)...)
			g.FinishCocoon(battleBin.CocoonId)
			delPileItem = append(delPileItem, &Material{
				Tid: Stamina,
				Num: conf.StaminaCost,
			})
		}
		if conf := gdconf.GetFarmElementConfig(req.StageId); conf != nil {
			rsp.DropData.ItemList = append(rsp.DropData.ItemList,
				g.getBattleDropData(conf.MappingInfoID, allSync, addPileItem, g.GetWorldLevel())...)
			delPileItem = append(delPileItem, &Material{
				Tid: Stamina,
				Num: conf.StaminaCost,
			})
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

	g.DelMaterial(delPileItem)
	g.AddItem(addPileItem)
	g.StaminaInfoScNotify()
	g.AllPlayerSyncScNotify(allSync)

	g.DelBattleBackupById(req.BattleId)
	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

/***********************************关卡/副本***********************************/

func (g *GamePlayer) StartCocoonStageCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartCocoonStageCsReq)
	rsp := &proto.StartCocoonStageScRsp{
		PropEntityId: req.PropEntityId,
		CocoonId:     req.CocoonId, // 关卡id
		Retcode:      0,
		Wave:         req.Wave,
	}
	cocoonConfig := gdconf.GetCocoonConfigById(req.CocoonId, req.WorldLevel)
	if cocoonConfig == nil {
		logger.Warn("No Cocoon like this can be found,cocoonId:%v,worldLevel:%v", req.CocoonId, req.WorldLevel)
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
		g.Send(cmd.StartCocoonStageScRsp, rsp)
		return
	}
	// 获取角色
	avatarMap := make(map[uint32]*BattleAvatar, 0)
	lineUp := g.GetBattleLineUp()
	for index, avatar := range lineUp.AvatarIdList {
		avatarMap[avatar.AvatarId] = &BattleAvatar{
			AvatarId:   avatar.AvatarId,
			AvatarType: avatar.LineAvatarType,
			AssistUid:  0,
			Index:      index,
		}
	}
	g.SetBattleStatus(spb.BattleType_Battle_NONE) // 设置战斗状态
	battleInfo, battleBackup := g.GetSceneBattleInfo(nil, cocoonConfig.StageIDList, avatarMap, req.WorldLevel, cocoonConfig.StageID)
	if battleInfo == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
		g.Send(cmd.StartCocoonStageScRsp, rsp)
		return
	}
	rsp.BattleInfo = battleInfo
	// 储存战斗信息
	battleBackup.CocoonId = req.CocoonId
	battleBackup.WorldLevel = req.WorldLevel
	g.AddBattleBackup(battleBackup)
	g.Send(cmd.StartCocoonStageScRsp, rsp)
}

func (g *GamePlayer) ActivateFarmElementCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ActivateFarmElementCsReq)

	rsp := &proto.ActivateFarmElementScRsp{
		WorldLevel: req.WorldLevel,
		EntityId:   req.EntityId,
	}
	g.Send(cmd.ActivateFarmElementScRsp, rsp)
}

/***********************************物品破坏处理***********************************/

func (g *GamePlayer) SceneCastSkillProp(sce *SceneCastEntity) {
	var addMPCost uint32 = 0
	allSync := &AllPlayerSync{AvatarList: make([]uint32, 0)}
	for _, propId := range sce.PropIdList {
		conf := gdconf.GetMazePropId(propId)
		if conf == nil {
			continue
		}
		if conf.RecoverMp {
			addMPCost += 2
		}
		if conf.RecoverHp {
			g.AvatarRecoverPercent(sce.AvatarId, 0.3, 0)
			allSync.AvatarList = append(allSync.AvatarList, sce.AvatarId)
		}
	}
	if addMPCost > 0 {
		g.AddLineUpMp(2) // 如果涉及到更新战斗中的队伍状态，这部分需要改
		g.SyncLineupNotify(g.GetBattleLineUp())
	}
	g.AllPlayerSyncScNotify(allSync)
}
