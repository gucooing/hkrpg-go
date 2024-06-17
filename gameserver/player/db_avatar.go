package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func NewAvatar() *spb.Avatar {
	return &spb.Avatar{
		AvatarList:        make(map[uint32]*spb.AvatarBin),
		Gender:            spb.Gender_GenderMan,
		CurMainAvatar:     spb.HeroBasicType_BoyWarrior,
		HeroBasicTypeInfo: make(map[uint32]*spb.HeroBasicTypeInfo),
		BattleAvatarList:  make(map[uint32]*spb.AvatarBin),
	}
}

func (g *GamePlayer) GetAvatar() *spb.Avatar {
	db := g.GetBasicBin()
	if db.Avatar == nil {
		db.Avatar = NewAvatar()
	}
	return db.Avatar
}

func (g *GamePlayer) GetAvatarList() map[uint32]*spb.AvatarBin {
	db := g.GetAvatar()
	if db.AvatarList == nil {
		db.AvatarList = make(map[uint32]*spb.AvatarBin)
	}
	return db.AvatarList
}

func (g *GamePlayer) GetBattleAvatarList() map[uint32]*spb.AvatarBin {
	db := g.GetAvatar()
	if db.BattleAvatarList == nil {
		db.BattleAvatarList = make(map[uint32]*spb.AvatarBin)
	}
	return db.BattleAvatarList
}

func (g *GamePlayer) GetAvatarBinById(avatarId uint32) *spb.AvatarBin {
	bin := g.GetAvatarList()
	return bin[avatarId]
}

func (g *GamePlayer) GetBattleAvatarBinById(avatarId uint32) *spb.AvatarBin {
	bin := g.GetBattleAvatarList()
	return bin[avatarId]
}

func (g *GamePlayer) GetAvatarById(avatarId uint32) *spb.AvatarBin {
	var bin map[uint32]*spb.AvatarBin
	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_NONE:
		bin = g.GetAvatarList()
	default:
		bin = g.GetBattleAvatarList()
	}
	return bin[avatarId]
}

func (g *GamePlayer) GetCurAvatar() *spb.AvatarBin {
	db := g.GetSceneAvatarId()
	return g.GetAvatarBinById(db)
}

func (g *GamePlayer) GetHeroBasicTypeInfo() map[uint32]*spb.HeroBasicTypeInfo {
	db := g.GetAvatar()
	if db.HeroBasicTypeInfo == nil {
		db.HeroBasicTypeInfo = make(map[uint32]*spb.HeroBasicTypeInfo)
	}
	return db.HeroBasicTypeInfo
}

func (g *GamePlayer) GetHeroBasicTypeInfoBy(basicType spb.HeroBasicType) *spb.HeroBasicTypeInfo {
	db := g.GetHeroBasicTypeInfo()
	if db[uint32(basicType)] == nil {
		g.AddHeroBasicTypeInfo(basicType)
	}
	return db[uint32(basicType)]
}

func (g *GamePlayer) AddHeroBasicTypeInfo(basicType spb.HeroBasicType) {
	db := g.GetHeroBasicTypeInfo()
	if db[uint32(basicType)] != nil {
		return
	}
	db[uint32(basicType)] = &spb.HeroBasicTypeInfo{
		Rank:          0,
		BasicType:     basicType,
		SkillTreeList: g.GetBasicTypeSkillTreeList(uint32(basicType)),
	}
}

func (g *GamePlayer) GetBasicTypeSkillTreeList(avatarId uint32) []*spb.AvatarSkillBin {
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

func (g *GamePlayer) GetSkillTreeList(avatarId uint32) []*spb.AvatarSkillBin {
	skilltreeList := make([]*spb.AvatarSkillBin, 0)
	if avatarId/1000 == 8 {
		avatarId = 8001
	}
	avatarBin := g.GetAvatarBinById(avatarId)
	if avatarId == 8001 {
		basicInfo := g.GetHeroBasicTypeInfoBy(g.GetAvatar().CurMainAvatar)
		if basicInfo == nil {
			g.AddHeroBasicTypeInfo(g.GetAvatar().CurMainAvatar)
			basicInfo = g.GetHeroBasicTypeInfoBy(g.GetAvatar().CurMainAvatar)
		}
		for _, info := range basicInfo.SkillTreeList {
			if info.Level == 0 {
				continue
			}
			avatarSkillBin := &spb.AvatarSkillBin{
				PointId: info.PointId,
				Level:   info.Level,
			}
			skilltreeList = append(skilltreeList, avatarSkillBin)
		}
		return skilltreeList
	}
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
	if avatarBin.SkilltreeList == nil {
		for id, level := range gdconf.GetAvatarSkilltreeListById(avatarId) {
			avatarSkillBin := &spb.AvatarSkillBin{
				PointId: id,
				Level:   level,
			}
			skilltreeList = append(skilltreeList, avatarSkillBin)
		}
		avatarBin.SkilltreeList = skilltreeList
	}
	return avatarBin.SkilltreeList
}

func (g *GamePlayer) SetAvatarMakSkillByAvatarId(avatarId uint32) {
	db := g.GetSkillTreeList(avatarId)
	if db == nil {
		return
	}
	for _, avatarSkill := range db {
		conf := gdconf.GetAvatarSkilltreeBySkillId(avatarSkill.PointId, 1)
		if conf == nil {
			continue
		}
		avatarSkill.Level = conf.MaxLevel
	}
}

func (g *GamePlayer) AddAvatar(avatarId uint32, src proto.AddAvatarSrcState) {
	if gdconf.GetAvatarDataById(avatarId) == nil {
		return // 过滤没有的角色
	}
	db := g.GetAvatarList()
	if db[avatarId] != nil {
		var pileItem []*Material
		pileItem = append(pileItem, &Material{
			Tid: avatarId + 10000,
			Num: 1,
		})
		g.AddMaterial(pileItem)
		return
	}
	db[avatarId] = &spb.AvatarBin{
		AvatarId:          avatarId,
		Exp:               0,
		Level:             1,
		AvatarType:        uint32(spb.AvatarType_AVATAR_FORMAL_TYPE),
		FirstMetTimeStamp: uint64(time.Now().Unix()),
		PromoteLevel:      0,
		Rank:              0,
		Hp:                10000,
		SpBar: &spb.AvatarSpBarInfo{
			CurSp: 10000,
			MaxSp: 10000,
		},
		SkilltreeList:     g.GetSkillTreeList(avatarId),
		EquipmentUniqueId: 0,
		EquipRelic:        make(map[uint32]uint32),
		TakenRewards:      make([]uint32, 0),
		BuffList:          0,
	}

	g.AvatarPlayerSyncScNotify(avatarId)
	g.Send(cmd.AddAvatarScNotify, &proto.AddAvatarScNotify{
		Reward:       nil,
		BaseAvatarId: avatarId,
		Src:          src,
		IsNew:        true,
	})
}

func (g *GamePlayer) CopyBattleAvatar(avatarBin *spb.AvatarBin) {
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
		Rank:              avatarBin.Rank,
		Hp:                12000,
		SpBar: &spb.AvatarSpBarInfo{
			CurSp: 6000,
			MaxSp: 12000,
		},
		SkilltreeList:     avatarBin.SkilltreeList,
		EquipmentUniqueId: avatarBin.EquipmentUniqueId,
		EquipRelic:        avatarBin.EquipRelic,
		TakenRewards:      avatarBin.TakenRewards,
		BuffList:          avatarBin.BuffList,
	}
}

func (g *GamePlayer) AddAvatarRank(rank uint32, db *spb.AvatarBin) {
	if db == nil {
		return
	}
	db.Rank += rank
	if db.Rank > 6 || db.Rank < 0 {
		db.Rank = 6
	}
}

func (g *GamePlayer) BattleUpAvatar(abi []*proto.AvatarBattleInfo, bt proto.BattleEndStatus) {
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
		if avatarStt.Id/1000 == 8 {
			avatarId = 8001
		}
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

	g.SyncLineupNotifyByLineBin(g.GetBattleLineUp())
}

func (g *GamePlayer) SetAvatarEquipRelic(avatarId, slot, relicId uint32) {
	db := g.GetAvatarBinById(avatarId)
	if db == nil {
		return
	}
	if db.EquipRelic == nil {
		db.EquipRelic = make(map[uint32]uint32)
	}
	db.EquipRelic[slot] = relicId
}

func (g *GamePlayer) GetAvatarEquipRelic(avatarId, slot uint32) *spb.Relic {
	db := g.GetAvatarBinById(avatarId)
	if db == nil {
		return nil
	}
	if db.EquipRelic == nil {
		db.EquipRelic = make(map[uint32]uint32)
	}
	return g.GetRelicById(db.EquipRelic[slot])
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetProtoAvatarById(avatarId uint32) *proto.Avatar {
	if avatarId/1000 == 8 {
		avatarId = 8001
	}
	avatardb := g.GetAvatarBinById(avatarId)
	if avatardb == nil {
		return nil
	}
	avatar := &proto.Avatar{
		SkilltreeList:               make([]*proto.AvatarSkillTree, 0),
		Exp:                         avatardb.Exp,
		BaseAvatarId:                avatardb.AvatarId,
		Rank:                        avatardb.Rank,
		EquipmentUniqueId:           avatardb.EquipmentUniqueId,
		EquipRelicList:              make([]*proto.EquipRelic, 0),
		HasTakenPromotionRewardList: avatardb.TakenRewards,
		FirstMetTimeStamp:           avatardb.FirstMetTimeStamp,
		Promotion:                   avatardb.PromoteLevel,
		Level:                       avatardb.Level,
	}
	for _, skill := range g.GetSkillTreeList(avatarId) {
		if skill.Level == 0 {
			continue
		}
		avatarSkillTree := &proto.AvatarSkillTree{
			PointId: skill.PointId,
			Level:   skill.Level,
		}
		avatar.SkilltreeList = append(avatar.SkilltreeList, avatarSkillTree)
	}
	for id, relic := range avatardb.EquipRelic {
		if relic == 0 {
			continue
		}
		equipRelic := &proto.EquipRelic{
			Type:          id,
			RelicUniqueId: relic,
		}
		avatar.EquipRelicList = append(avatar.EquipRelicList, equipRelic)
	}
	if avatarId == 8001 {
		basic := g.GetHeroBasicTypeInfoBy(g.GetAvatar().CurMainAvatar)
		avatar.SkilltreeList = make([]*proto.AvatarSkillTree, 0)
		avatar.Rank = basic.Rank
	}

	return avatar
}

type BattleAvatar struct {
	AvatarId   uint32             // 角色id
	AvatarType spb.LineAvatarType // 角色类型
	AssistUid  uint32             // 助战uid
}

func (g *GamePlayer) GetProtoBattleAvatar(bAList map[uint32]*BattleAvatar) ([]*proto.BattleAvatar, []*proto.BattleBuff) {
	battleAvatarList := make([]*proto.BattleAvatar, 0)
	buffList := make([]*proto.BattleBuff, 0)
	for id, bA := range bAList {
		if bA.AvatarId == 0 {
			continue
		}
		battleAvatar := new(proto.BattleAvatar)
		switch bA.AvatarType {
		case spb.LineAvatarType_LineAvatarType_MI:
			avatarBin := g.GetAvatarById(bA.AvatarId)
			if avatarBin == nil {
				continue
			}
			battleAvatar = &proto.BattleAvatar{
				AvatarType:    proto.AvatarType(avatarBin.AvatarType),
				Id:            avatarBin.AvatarId,
				Level:         avatarBin.Level,
				Rank:          avatarBin.Rank,
				Index:         id,
				SkilltreeList: make([]*proto.AvatarSkillTree, 0),
				EquipmentList: make([]*proto.BattleEquipment, 0),
				Hp:            avatarBin.Hp,
				Promotion:     avatarBin.PromoteLevel,
				RelicList:     make([]*proto.BattleRelic, 0),
				WorldLevel:    g.GetWorldLevel(),
				AssistUid:     bA.AssistUid,
				SpBar: &proto.SpBarInfo{
					CurSp: avatarBin.SpBar.CurSp,
					MaxSp: avatarBin.SpBar.MaxSp,
				},
			}
			if bA.AvatarId == 8001 {
				battleAvatar.Id = uint32(g.GetAvatar().CurMainAvatar)
			}
			// 获取技能
			for _, skill := range g.GetSkillTreeList(bA.AvatarId) {
				if skill.Level == 0 {
					continue
				}
				avatarSkillTree := &proto.AvatarSkillTree{
					PointId: skill.PointId,
					Level:   skill.Level,
				}
				battleAvatar.SkilltreeList = append(battleAvatar.SkilltreeList, avatarSkillTree)
			}
			// 获取装备
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
				equipmentList := &proto.BattleEquipment{
					Id:        equipment.Tid,
					Level:     equipment.Level,
					Promotion: equipment.Promotion,
					Rank:      equipment.Rank,
				}
				battleAvatar.EquipmentList = append(battleAvatar.EquipmentList, equipmentList)
			}
		case spb.LineAvatarType_LineAvatarType_TRIAL:
			avatarBin := gdconf.GetSpecialAvatarById(bA.AvatarId)
			if avatarBin == nil {
				continue
			}
			battleAvatar = &proto.BattleAvatar{
				AvatarType:    proto.AvatarType_AVATAR_TRIAL_TYPE,
				Id:            bA.AvatarId,
				Level:         avatarBin.Level,
				Rank:          0,
				Index:         id,
				SkilltreeList: make([]*proto.AvatarSkillTree, 0),
				EquipmentList: make([]*proto.BattleEquipment, 0),
				Hp:            10000,
				Promotion:     avatarBin.Promotion,
				RelicList:     make([]*proto.BattleRelic, 0),
				WorldLevel:    g.GetWorldLevel(),
				AssistUid:     bA.AssistUid,
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
				battleAvatar.SkilltreeList = append(battleAvatar.SkilltreeList, avatarSkillTree)
			}
			// 获取角色装备的光锥
			if avatarBin.EquipmentID != 0 {
				equipmentList := &proto.BattleEquipment{
					Id:        avatarBin.EquipmentID,
					Level:     avatarBin.EquipmentLevel,
					Promotion: avatarBin.Promotion,
					Rank:      avatarBin.EquipmentRank,
				}
				battleAvatar.EquipmentList = append(battleAvatar.EquipmentList, equipmentList)
			}
		default:
			continue
		}
		battleAvatarList = append(battleAvatarList, battleAvatar)
		// 添加该角色的buff
		info := g.GetOnLineAvatarBuffById(bA.AvatarId)
		if info != nil {
			buffList = append(buffList, &proto.BattleBuff{
				Id:              info.BuffId,
				Level:           1,
				OwnerIndex:      id,
				WaveFlag:        4294967295,
				TargetIndexList: []uint32{1},
				DynamicValues:   make(map[string]float32),
			})
			g.DelOnLineAvatarBuff(info.AvatarId, info.BuffId)
		}
	}
	return battleAvatarList, buffList
}

func (g *GamePlayer) GetPlayerHeroBasicTypeInfo() []*proto.PlayerHeroBasicTypeInfo {
	basicTypeInfoList := make([]*proto.PlayerHeroBasicTypeInfo, 0)
	avatarDb := g.GetAvatar()
	avatarBin := g.GetAvatarBinById(8001)
	for id, heroBasic := range g.GetHeroBasicTypeInfo() {
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
		basicTypeInfo := &proto.PlayerHeroBasicTypeInfo{
			SkillTreeList:  make([]*proto.AvatarSkillTree, 0),
			BasicType:      proto.HeroBasicType(heroBasic.BasicType),
			EquipRelicList: make([]*proto.EquipRelic, 0),
			Rank:           heroBasic.Rank,
		}
		// 获取装备圣遗物
		for tp, relic := range avatarBin.EquipRelic {
			basicTypeInfo.EquipRelicList = append(basicTypeInfo.EquipRelicList, &proto.EquipRelic{
				Type:          tp,
				RelicUniqueId: relic,
			})
		}
		// 添加技能
		for _, skill := range heroBasic.SkillTreeList {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			basicTypeInfo.SkillTreeList = append(basicTypeInfo.SkillTreeList, avatarSkillTree)
		}
		basicTypeInfoList = append(basicTypeInfoList, basicTypeInfo)
	}
	return basicTypeInfoList
}
