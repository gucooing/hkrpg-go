package model

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewAvatar() *spb.Avatar {
	return &spb.Avatar{
		AvatarList:       make(map[uint32]*spb.AvatarBin),
		Gender:           spb.Gender_GenderMan,
		BattleAvatarList: make(map[uint32]*spb.AvatarBin),
	}
}

func (g *PlayerData) GetAvatar() *spb.Avatar {
	db := g.GetBasicBin()
	if db.Avatar == nil {
		db.Avatar = NewAvatar()
	}
	return db.Avatar
}

func (g *PlayerData) GetAvatarList() map[uint32]*spb.AvatarBin {
	db := g.GetAvatar()
	if db.AvatarList == nil {
		db.AvatarList = make(map[uint32]*spb.AvatarBin)
	}
	return db.AvatarList
}

func (g *PlayerData) GetBattleAvatarList() map[uint32]*spb.AvatarBin {
	db := g.GetAvatar()
	if db.BattleAvatarList == nil {
		db.BattleAvatarList = make(map[uint32]*spb.AvatarBin)
	}
	return db.BattleAvatarList
}

func (g *PlayerData) GetAvatarBinById(avatarId uint32) *spb.AvatarBin {
	bin := g.GetAvatarList()
	patchConf := gdconf.GetMultiplePathAvatarConfig(avatarId)
	if patchConf != nil {
		avatarId = patchConf.BaseAvatarID
	}
	return bin[avatarId]
}

func (g *PlayerData) GetBattleAvatarBinById(avatarId uint32) *spb.AvatarBin {
	bin := g.GetBattleAvatarList()
	return bin[avatarId]
}

func (g *PlayerData) GetAvatarById(avatarId uint32) *spb.AvatarBin {
	// var bin map[uint32]*spb.AvatarBin
	// switch g.GetBattleStatus() {
	// case spb.BattleType_Battle_NONE:
	// 	bin = g.GetAvatarList()
	// default:
	// 	bin = g.GetBattleAvatarList()
	// }
	bin := g.GetAvatarList()
	patchConf := gdconf.GetMultiplePathAvatarConfig(avatarId)
	if patchConf != nil {
		avatarId = patchConf.BaseAvatarID
	}
	return bin[avatarId]
}

func (g *PlayerData) GetCurAvatar() *spb.AvatarBin {
	db := g.GetSceneAvatarId()
	return g.GetAvatarBinById(db)
}

// 8001,8002,8003,8004,8005,8006 -> 8001
// 1001,1224 -> 1001
func (g *PlayerData) AddAvatar(avatarId uint32) {
	if gdconf.GetAvatarDataById(avatarId) == nil {
		return // 过滤没有的角色
	}
	db := g.GetAvatarList()
	// 重复角色判断
	if db[avatarId] != nil {
		var pileItem []*Material
		pileItem = append(pileItem, &Material{
			Tid: avatarId + 10000,
			Num: 1,
		})
		g.AddMaterial(pileItem)
		return
	}
	// 多命途判断
	patchConf := gdconf.GetMultiplePathAvatarConfig(avatarId)
	if patchConf != nil && patchConf.BaseAvatarID != patchConf.AvatarID {
		if db[patchConf.BaseAvatarID] == nil {
			return
		}
		g.AddMultiPathAvatar(avatarId)
		return
	}
	db[avatarId] = &spb.AvatarBin{
		AvatarId:          avatarId,
		Exp:               0,
		Level:             1,
		AvatarType:        uint32(spb.AvatarType_AVATAR_FORMAL_TYPE),
		FirstMetTimeStamp: uint64(time.Now().Unix()),
		PromoteLevel:      0,
		TakenRewards:      make([]uint32, 0),
		Hp:                10000,
		SpBar: &spb.AvatarSpBarInfo{
			CurSp: 10000,
			MaxSp: 10000,
		},
		IsMultiPath: false,
		CurPath:     avatarId,
		MultiPathAvatarInfoList: map[uint32]*spb.MultiPathAvatarInfo{
			avatarId: {
				AvatarId:          avatarId,
				Rank:              0,
				SkilltreeList:     g.newSkillTreeList(avatarId),
				EquipmentUniqueId: 0,
				EquipRelic:        make(map[uint32]uint32),
			},
		},
	}
}

// AddMultiPathAvatar 添加命途
func (g *PlayerData) AddMultiPathAvatar(avatarId uint32) {
	patchConf := gdconf.GetMultiplePathAvatarConfig(avatarId)
	if patchConf == nil {
		return
	}
	db := g.GetAvatarById(patchConf.BaseAvatarID)
	if db == nil {
		return
	}
	db.IsMultiPath = true
	if db.MultiPathAvatarInfoList[avatarId] == nil {
		db.MultiPathAvatarInfoList[avatarId] = &spb.MultiPathAvatarInfo{
			AvatarId:          avatarId,
			Rank:              0,
			SkilltreeList:     g.newSkillTreeList(avatarId),
			EquipmentUniqueId: 0,
			EquipRelic:        make(map[uint32]uint32),
		}
	}
}

// 获取命途
func (g *PlayerData) GetMultiPathAvatar(avatarId uint32) *spb.MultiPathAvatarInfo {
	patchConf := gdconf.GetMultiplePathAvatarConfig(avatarId)
	if patchConf == nil {
		return nil
	}
	db := g.GetAvatarById(patchConf.BaseAvatarID)
	if db == nil {
		return nil
	}
	return db.MultiPathAvatarInfoList[avatarId]
}

func (g *PlayerData) GetCurMultiPathAvatar(avatarId uint32) *spb.MultiPathAvatarInfo {
	db := g.GetAvatarById(avatarId)
	if db == nil {
		return nil
	}
	return db.MultiPathAvatarInfoList[avatarId]
}

// 添加技能
func (g *PlayerData) newSkillTreeList(avatarId uint32) []*spb.AvatarSkillBin {
	skilltreeList := make([]*spb.AvatarSkillBin, 0)
	for id, level := range gdconf.GetAvatarSkilltreeListById(avatarId) {
		avatarSkillBin := &spb.AvatarSkillBin{
			PointId: id,
			Level:   level,
		}
		skilltreeList = append(skilltreeList, avatarSkillBin)
	}
	return skilltreeList
}

// 获取技能
func (g *PlayerData) GetSkillTreeList(avatarId uint32) []*spb.AvatarSkillBin {
	skilltreeList := make([]*spb.AvatarSkillBin, 0)
	avatarBin := g.GetAvatarBinById(avatarId)
	if avatarBin == nil {
		for id, level := range gdconf.GetAvatarSkilltreeListById(avatarId) {
			avatarSkillBin := &spb.AvatarSkillBin{
				PointId: id,
				Level:   level,
			}
			skilltreeList = append(skilltreeList, avatarSkillBin)
		}
		return skilltreeList
	}
	// 根据当前命途获取技能
	curPath := avatarBin.MultiPathAvatarInfoList[avatarBin.CurPath]
	if curPath == nil {
		return skilltreeList
	}
	return curPath.SkilltreeList
}

func (g *PlayerData) CopyBattleAvatar(avatarBin *spb.AvatarBin) {
	db := g.GetBattleAvatarList()
	if avatarBin == nil {
		return
	}
	db[avatarBin.AvatarId] = &spb.AvatarBin{
		AvatarId:          avatarBin.AvatarId,
		Exp:               avatarBin.Exp,
		Level:             avatarBin.Level,
		AvatarType:        uint32(spb.AvatarType_AVATAR_FORMAL_TYPE),
		FirstMetTimeStamp: avatarBin.FirstMetTimeStamp,
		PromoteLevel:      avatarBin.PromoteLevel,
		TakenRewards:      avatarBin.TakenRewards,
		Hp:                12000,
		SpBar: &spb.AvatarSpBarInfo{
			CurSp: 6000,
			MaxSp: 12000,
		},
		IsMultiPath:             avatarBin.IsMultiPath,
		CurPath:                 avatarBin.CurPath,
		MultiPathAvatarInfoList: avatarBin.MultiPathAvatarInfoList,
	}
}

func (g *PlayerData) AddAvatarRank(rank uint32, db *spb.AvatarBin) {
	if db == nil {
		return
	}
	if c := db.MultiPathAvatarInfoList[db.CurPath]; c != nil {
		c.Rank += rank
	}
}

// 战斗结束后更新角色状态
func (g *PlayerData) BattleUpAvatar(abi []*proto.AvatarBattleInfo, bt proto.BattleEndStatus) {
	var deadAatarNum uint32 = 0
re:
	for _, avatarStt := range abi {
		switch bt {
		case proto.BattleEndStatus_BATTLE_END_NONE:
			break re
		case proto.BattleEndStatus_BATTLE_END_WIN: // 胜利
		case proto.BattleEndStatus_BATTLE_END_LOSE: // 失败
		case proto.BattleEndStatus_BATTLE_END_QUIT: // 撤退
			break re
		}
		avatarId := avatarStt.Id
		avatarBin := g.GetAvatarById(avatarId)
		if avatarBin == nil {
			continue
		}
		sp := uint32((avatarStt.AvatarStatus.LeftSp / avatarStt.AvatarStatus.MaxSp) * 10000)
		hp := uint32((avatarStt.AvatarStatus.LeftHp / avatarStt.AvatarStatus.MaxHp) * 10000)
		if hp == 0 {
			deadAatarNum++
			hp = 2000
		}
		avatarBin.Hp = hp
		avatarBin.SpBar.CurSp = sp
	}

	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		g.AddChallengeDeadAvatar(deadAatarNum)
	case spb.BattleType_Battle_CHALLENGE_Story:
		g.AddChallengeDeadAvatar(deadAatarNum)
	}
}

func (g *PlayerData) SetAvatarEquipRelic(avatarId, slot, relicId uint32) {
	db := g.GetAvatarBinById(avatarId)
	if db == nil {
		return
	}
	if v := db.MultiPathAvatarInfoList[db.CurPath]; v != nil {
		if v.EquipRelic == nil {
			v.EquipRelic = make(map[uint32]uint32)
		}
		v.EquipRelic[slot] = relicId
	}
}

func (g *PlayerData) GetAvatarEquipRelic(avatarId, slot uint32) *spb.Relic {
	db := g.GetAvatarBinById(avatarId)
	if db == nil {
		return nil
	}
	if v := db.MultiPathAvatarInfoList[db.CurPath]; v != nil {
		if v.EquipRelic == nil {
			v.EquipRelic = make(map[uint32]uint32)
		}
		return g.GetRelicById(v.EquipRelic[slot])
	}
	return nil
}

func (g *PlayerData) AvatarAddExp(avatarId, exp uint32) (uint32, bool) {
	conf := gdconf.GetAvatarDataById(avatarId)
	dbAvatar := g.GetAvatarById(avatarId)
	if conf == nil || dbAvatar == nil {
		return 0, false
	}
	level, exp, newExp := gdconf.GetExpTypeByLevel(conf.ExpGroup, exp, dbAvatar.Level, dbAvatar.PromoteLevel, dbAvatar.AvatarId)
	if level == 0 && exp == 0 {
		return exp, false
	}
	dbAvatar.Exp = exp
	dbAvatar.Level = level
	return newExp, true
}

func (g *PlayerData) AvatarRecover(avatarId uint32) {
	db := g.GetAvatarById(avatarId)
	if db != nil {
		db.Hp = 10000
	}
}

func (g *PlayerData) getAvatarBaseHp(avatarId uint32) float64 {
	avatarDb := g.GetAvatarById(avatarId)
	if avatarDb == nil {
		return 0
	}
	logger.Debug("avatar %v oldHp:%v", avatarId, avatarDb.Hp)
	avatarConf := gdconf.GetAvatarPromotionConfig(avatarId, avatarDb.PromoteLevel)
	if avatarConf == nil {
		return 0
	}
	// 获取装备光锥增加血量
	var equipmentAddHp float64
	equipmentDb := g.GetEquipmentById(avatarDb.MultiPathAvatarInfoList[avatarDb.CurPath].EquipmentUniqueId)
	if equipmentDb != nil {
		equipmentConf := gdconf.GetEquipmentPromotionConfig(equipmentDb.Tid, equipmentDb.Promotion)
		equipmentAddHp = equipmentConf.BaseHP.Value + equipmentConf.BaseHPAdd.Value*float64(equipmentDb.Level-1)
	}
	// 计算白字
	baseHp := avatarConf.HPBase.Value + avatarConf.HPAdd.Value*float64(avatarDb.Level-1) + equipmentAddHp
	logger.Debug("avatar %v base baseHp:%v", avatarId, baseHp)
	return baseHp
}

func (g *PlayerData) getAvatarEquiHp(avatarId uint32, baseHp float64) float64 {
	avatarDb := g.GetAvatarById(avatarId)
	if avatarDb == nil {
		return 0
	}
	path := avatarDb.MultiPathAvatarInfoList[avatarDb.CurPath]
	// 获取装备光锥增加血量
	var equipmentAddHp float64
	equipmentDb := g.GetEquipmentById(path.EquipmentUniqueId)
	if equipmentDb != nil {
		equipmentConf := gdconf.GetEquipmentSkillConfig(equipmentDb.Tid, equipmentDb.Rank)
		if equipmentConf != nil {
			for _, abilityProperty := range equipmentConf.AbilityProperty {
				if abilityProperty.PropertyType == "AttackAddedRatio" {
					equipmentAddHp = baseHp * abilityProperty.Value.Value
					break
				}
			}
		}
	}
	// 获取圣遗物增加血量
	var relicAddHp float64
	for _, relicUniqueId := range path.EquipRelic {
		relicDb := g.GetRelicById(relicUniqueId)
		if relicDb == nil {
			continue
		}
		relicConf := gdconf.GetRelicById(relicDb.Tid)
		if relicConf == nil {
			continue
		}
		// 主属性
		mainAffixConf := gdconf.GetRelicMainAffixConfig(relicConf.MainAffixGroup, relicDb.MainAffixId)
		if mainAffixConf != nil {
			if mainAffixConf.Property == "HPDelta" {
				relicAddHp += mainAffixConf.BaseValue.Value + mainAffixConf.LevelAdd.Value*float64(relicDb.Level)
			}
			if mainAffixConf.Property == "HPAddedRatio" {
				relicAddHp += (mainAffixConf.BaseValue.Value + mainAffixConf.LevelAdd.Value*float64(relicDb.Level)) * baseHp
			}
		}
		// 副属性
		for _, subAffix := range relicDb.RelicAffix {
			subAffixConf := gdconf.GetRelicSubAffixConfig(relicConf.SubAffixGroup, subAffix.AffixId)
			if subAffixConf == nil {
				continue
			}
			if subAffixConf.Property == "HPDelta" {
				relicAddHp += subAffixConf.BaseValue.Value*float64(subAffix.Cnt) + subAffixConf.StepValue.Value*float64(subAffix.Step)
			}
			if subAffixConf.Property == "HPAddedRatio" {
				relicAddHp += (subAffixConf.BaseValue.Value*float64(subAffix.Cnt) + subAffixConf.StepValue.Value*float64(subAffix.Step)) * baseHp
			}
		}
	}
	equiHp := equipmentAddHp + relicAddHp
	logger.Debug("avatar %v equiHp:%v", avatarId, equiHp)

	return equiHp
}

// 恢复角色
func (g *PlayerData) AvatarRecoverPercent(avatarId uint32, Value, percent float64) {
	avatarDb := g.GetAvatarById(avatarId)
	if avatarDb == nil {
		return
	}
	// 计算白字
	baseHp := g.getAvatarBaseHp(avatarId)
	// 计算绿字
	equiHp := g.getAvatarEquiHp(avatarId, baseHp)
	// 正式计算血量
	hp := baseHp + equiHp // 总血量
	// 计算现有血量
	oldHp := float64(avatarDb.Hp) / 10000 * hp
	// 计算增加后的血量
	newHp := oldHp + Value + hp*percent
	// 更新客户端血量
	avatarDb.Hp = uint32(newHp / hp * 10000)
	if avatarDb.Hp < 0 {
		avatarDb.Hp = 2000
	}
	logger.Debug("avatar %v new hp:%v", avatarId, avatarDb.Hp)
}

func (g *PlayerData) CheckUnlockMultiPath(allSync *AllPlayerSync) { // 任务检查发放命途
	finishMainMissionList := g.GetFinishMainMissionList()   // 已完成的主任务
	finishSubMissionList := g.GetFinishSubMainMissionList() // 已完成的子任务
	for _, info := range gdconf.GetMultiplePathAvatarConfigMap() {
		if info.UnlockConditions == nil {
			continue
		}
		db := g.GetAvatarById(info.AvatarID)
		if db == nil {
			continue
		}
		if db.MultiPathAvatarInfoList[info.AvatarID] != nil {
			continue
		}
		var isUnlock = false
		for _, unlockInfo := range info.UnlockConditions {
			switch unlockInfo.Type {
			case "FinishMainMission":
				if finishMainMissionList[alg.S2U32(unlockInfo.Param)] != nil {
					isUnlock = true
				} else {
					isUnlock = false
					break
				}
			case "FinishSubMission":
				if finishSubMissionList[alg.S2U32(unlockInfo.Param)] != nil {
					isUnlock = true
				} else {
					isUnlock = false
					break
				}
			}
		}
		if isUnlock {
			allSync.AvatarList = append(allSync.AvatarList, info.BaseAvatarID)
			g.AddMultiPathAvatar(info.AvatarID)
		}
	}
}

func (g *PlayerData) GetBattleAvatarMap(lineUp *spb.Line) map[uint32]*BattleAvatar {
	avatarMap := make(map[uint32]*BattleAvatar, 0)
	if lineUp == nil || lineUp.AvatarIdList == nil {
		return avatarMap
	}
	for index, avatar := range lineUp.AvatarIdList {
		baseAvatarId := avatar.AvatarId
		if avatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
			conf := gdconf.GetSpecialAvatarById(avatar.AvatarId)
			if conf == nil {
				continue
			}
			baseAvatarId = conf.AvatarID
		}
		info := &BattleAvatar{
			AvatarId:     avatar.AvatarId,
			BaseAvatarId: baseAvatarId,
			AvatarType:   avatar.LineAvatarType,
			AssistUid:    0,
			Index:        index,
		}
		if lineUp.LeaderSlot == avatar.Slot {
			info.IsCur = true
		}
		avatarMap[baseAvatarId] = info
	}
	return avatarMap
}

/****************************************************功能***************************************************/

func (g *PlayerData) GetProtoAvatarById(avatarId uint32) *proto.Avatar {
	avatardb := g.GetAvatarBinById(avatarId)
	if avatardb == nil {
		return nil
	}
	patch := avatardb.MultiPathAvatarInfoList[avatardb.CurPath]
	if patch == nil {
		return nil
	}
	avatar := &proto.Avatar{
		SkilltreeList:               make([]*proto.AvatarSkillTree, 0),
		Exp:                         avatardb.Exp,
		BaseAvatarId:                avatardb.AvatarId,
		EquipRelicList:              make([]*proto.EquipRelic, 0),
		HasTakenPromotionRewardList: avatardb.TakenRewards,
		FirstMetTimeStamp:           avatardb.FirstMetTimeStamp,
		Promotion:                   avatardb.PromoteLevel,
		Level:                       avatardb.Level,
	}
	avatar.Rank = patch.Rank
	avatar.EquipmentUniqueId = patch.EquipmentUniqueId
	for _, skill := range patch.SkilltreeList {
		if skill.Level == 0 {
			continue
		}
		avatarSkillTree := &proto.AvatarSkillTree{
			PointId: skill.PointId,
			Level:   skill.Level,
		}
		avatar.SkilltreeList = append(avatar.SkilltreeList, avatarSkillTree)
	}
	for id, relic := range patch.EquipRelic {
		if relic == 0 {
			continue
		}
		equipRelic := &proto.EquipRelic{
			Type:          id,
			RelicUniqueId: relic,
		}
		avatar.EquipRelicList = append(avatar.EquipRelicList, equipRelic)
	}

	return avatar
}

type BattleAvatar struct {
	AvatarId     uint32 // 角色id
	BaseAvatarId uint32
	AvatarType   spb.LineAvatarType // 角色类型
	AssistUid    uint32             // 助战uid
	Index        uint32
	IsCur        bool
}

// 添加战斗角色列表
func (g *PlayerData) GetProtoBattleAvatar(bAList map[uint32]*BattleAvatar) []*proto.BattleAvatar {
	battleAvatarList := make([]*proto.BattleAvatar, 0)
	for _, bA := range bAList {
		if bA.AvatarId == 0 || bA.Index > 3 {
			continue
		}
		switch bA.AvatarType {
		case spb.LineAvatarType_LineAvatarType_MI:
			battleAvatarList = append(battleAvatarList, g.GetBattleAvatar(bA.AvatarId, bA.Index))
		case spb.LineAvatarType_LineAvatarType_TRIAL:
			if ok, _ := g.SpecialMainAvatar(bA.AvatarId); !ok {
				continue
			}
			battleAvatarList = append(battleAvatarList, g.GetTrialBattleAvatar(bA.AvatarId, bA.Index))
		default:
			continue
		}
	}
	return battleAvatarList
}

// 角色
func (g *PlayerData) GetBattleAvatar(avatarId, index uint32) *proto.BattleAvatar {
	db := g.GetAvatarById(avatarId)
	if db == nil {
		return nil
	}
	pathDb := db.MultiPathAvatarInfoList[db.CurPath]
	if pathDb == nil {
		return nil
	}
	info := &proto.BattleAvatar{
		AvatarType:    proto.AvatarType(db.AvatarType),
		Id:            db.CurPath, // 当前命途
		Level:         db.Level,
		Rank:          pathDb.Rank,
		Index:         index,
		SkilltreeList: make([]*proto.AvatarSkillTree, 0),
		EquipmentList: make([]*proto.BattleEquipment, 0),
		Hp:            db.Hp,
		Promotion:     db.PromoteLevel,
		RelicList:     make([]*proto.BattleRelic, 0),
		WorldLevel:    g.GetWorldLevel(),
		SpBar: &proto.SpBarInfo{
			CurSp: db.SpBar.CurSp,
			MaxSp: db.SpBar.MaxSp,
		},
	}
	// 获取技能
	for _, skill := range g.GetSkillTreeList(avatarId) {
		if skill.Level == 0 {
			continue
		}
		avatarSkillTree := &proto.AvatarSkillTree{
			PointId: skill.PointId,
			Level:   skill.Level,
		}
		info.SkilltreeList = append(info.SkilltreeList, avatarSkillTree)
	}
	// 获取装备
	for _, relic := range pathDb.EquipRelic {
		equipRelic := g.GetProtoBattleRelicById(relic)
		if equipRelic == nil {
			delete(pathDb.EquipRelic, relic)
			continue
		}
		info.RelicList = append(info.RelicList, equipRelic)
	}
	// 获取角色装备的光锥
	if pathDb.EquipmentUniqueId != 0 {
		equipment := g.GetEquipment(pathDb.EquipmentUniqueId)
		equipmentList := &proto.BattleEquipment{
			Id:        equipment.Tid,
			Level:     equipment.Level,
			Promotion: equipment.Promotion,
			Rank:      equipment.Rank,
		}
		info.EquipmentList = append(info.EquipmentList, equipmentList)
	}
	return info
}

// 试用角色
func (g *PlayerData) GetTrialBattleAvatar(avatarId, index uint32) *proto.BattleAvatar {
	avatarBin := gdconf.GetSpecialAvatarById(avatarId)
	if avatarBin == nil {
		return nil
	}
	info := &proto.BattleAvatar{
		AvatarType:    proto.AvatarType_AVATAR_TRIAL_TYPE,
		Id:            avatarId,
		Level:         avatarBin.Level,
		Rank:          0,
		Index:         index,
		SkilltreeList: make([]*proto.AvatarSkillTree, 0),
		EquipmentList: make([]*proto.BattleEquipment, 0),
		Hp:            10000,
		Promotion:     avatarBin.Promotion,
		RelicList:     make([]*proto.BattleRelic, 0),
		WorldLevel:    g.GetWorldLevel(),
		SpBar: &proto.SpBarInfo{
			CurSp: 6000,
			MaxSp: 10000,
		},
	}
	// 获取技能
	for _, skill := range g.GetSkillTreeList(avatarBin.PlayerID) {
		if skill.Level == 0 {
			continue
		}
		avatarSkillTree := &proto.AvatarSkillTree{
			PointId: skill.PointId,
			Level:   skill.Level,
		}
		info.SkilltreeList = append(info.SkilltreeList, avatarSkillTree)
	}
	// 获取角色装备的光锥
	if avatarBin.EquipmentID != 0 {
		equipmentList := &proto.BattleEquipment{
			Id:        avatarBin.EquipmentID,
			Level:     avatarBin.EquipmentLevel,
			Promotion: avatarBin.Promotion,
			Rank:      avatarBin.EquipmentRank,
		}
		info.EquipmentList = append(info.EquipmentList, equipmentList)
	}

	return info
}

func (g *PlayerData) GetMultiPathAvatarInfo(avatarId uint32) []*proto.MultiPathAvatarInfo {
	basicTypeInfoList := make([]*proto.MultiPathAvatarInfo, 0)
	avatarDb := g.GetAvatar()
	avatarBin := g.GetAvatarBinById(avatarId)
	for id, pathInfo := range avatarBin.MultiPathAvatarInfoList {
		if avatarBin.AvatarId == 8001 {
			switch avatarDb.Gender {
			case spb.Gender_GenderMan:
				if id%2 == 0 {
					continue
				}
			case spb.Gender_GenderWoman:
				if id%2 != 0 {
					continue
				}
			}
		}
		basicTypeInfo := &proto.MultiPathAvatarInfo{
			MultiPathSkillTree: make([]*proto.AvatarSkillTree, 0),
			AvatarId:           proto.MultiPathAvatarType(pathInfo.AvatarId),
			EquipRelicList:     make([]*proto.EquipRelic, 0),
			Rank:               pathInfo.Rank,
			PathEquipmentId:    pathInfo.EquipmentUniqueId,
		}
		// 获取装备圣遗物
		for tp, relic := range pathInfo.EquipRelic {
			basicTypeInfo.EquipRelicList = append(basicTypeInfo.EquipRelicList, &proto.EquipRelic{
				Type:          tp,
				RelicUniqueId: relic,
			})
		}
		// 添加技能
		for _, skill := range pathInfo.SkilltreeList {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			basicTypeInfo.MultiPathSkillTree = append(basicTypeInfo.MultiPathSkillTree, avatarSkillTree)
		}
		basicTypeInfoList = append(basicTypeInfoList, basicTypeInfo)
	}
	return basicTypeInfoList
}
