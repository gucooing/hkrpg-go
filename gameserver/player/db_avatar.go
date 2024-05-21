package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) NewAvatar() *spb.Avatar {
	return &spb.Avatar{
		AvatarList:        make(map[uint32]*spb.AvatarBin),
		Gender:            spb.Gender_GenderMan,
		CurMainAvatar:     spb.HeroBasicType_BoyWarrior,
		HeroBasicTypeInfo: g.GetHeroBasicTypeInfo(),
	}
}

func (g *GamePlayer) GetAvatar() *spb.Avatar {
	db := g.GetBasicBin()
	if db.Avatar == nil {
		db.Avatar = g.NewAvatar()
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

func (g *GamePlayer) GetAvatarBinById(avatarId uint32) *spb.AvatarBin {
	bin := g.GetAvatarList()
	return bin[avatarId]
}

func (g *GamePlayer) GetCurAvatar() *spb.AvatarBin {
	db := g.GetSceneAvatarId()
	return g.GetAvatarBinById(db)
}

func (g *GamePlayer) GetHeroBasicTypeInfo() []*spb.HeroBasicTypeInfo {
	heroBasic := make([]*spb.HeroBasicTypeInfo, 0)
	heroBasicTypeInfo := &spb.HeroBasicTypeInfo{
		Rank:          0,
		BasicType:     spb.HeroBasicType_BoyWarrior,
		SkillTreeList: g.GetBasicTypeSkillTreeList(uint32(spb.HeroBasicType_BoyWarrior)),
	}
	heroBasic = append(heroBasic, heroBasicTypeInfo)

	heroBasicTypeInfo = &spb.HeroBasicTypeInfo{
		Rank:          0,
		BasicType:     spb.HeroBasicType_BoyKnight,
		SkillTreeList: g.GetBasicTypeSkillTreeList(uint32(spb.HeroBasicType_BoyKnight)),
	}
	heroBasic = append(heroBasic, heroBasicTypeInfo)
	return heroBasic
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
	avatarBin := g.GetAvatarBinById(avatarId)
	if avatarId/1000 == 8 || avatarBin == nil {
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

func (g *GamePlayer) AddAvatar(avatarId uint32) {
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
}

func (g *GamePlayer) BattleUpAvatar(abi []*proto.AvatarBattleInfo, bt proto.BattleEndStatus) {
	var deadAatarNum uint32 = 0
	for _, avatarStt := range abi {
		avatarBin := g.GetAvatarBinById(avatarStt.Id)
		if avatarBin == nil {
			continue
		}
		sp := uint32((avatarStt.AvatarStatus.LeftSp / avatarStt.AvatarStatus.MaxSp) * 10000)
		hp := uint32((avatarStt.AvatarStatus.LeftHp / avatarStt.AvatarStatus.MaxHp) * 10000)
		if hp == 0 {
			deadAatarNum++
			hp = 2000
		}
		switch bt {
		case proto.BattleEndStatus_BATTLE_END_NONE:
		case proto.BattleEndStatus_BATTLE_END_WIN: // 胜利
			avatarBin.Hp = hp
			avatarBin.SpBar.CurSp = sp
		case proto.BattleEndStatus_BATTLE_END_LOSE: // 失败
			avatarBin.Hp = hp
			avatarBin.SpBar.CurSp = sp
		case proto.BattleEndStatus_BATTLE_END_QUIT: // 撤退

		}
	}

	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE: // 忘却之庭不需要更新状态
		g.AddChallengeDeadAvatar(deadAatarNum)
	case spb.BattleType_Battle_CHALLENGE_Story:
		g.AddChallengeDeadAvatar(deadAatarNum)
	}
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetProtoAvatarById(avatarId uint32) *proto.Avatar {
	avatardb := g.GetAvatarBinById(avatarId)
	if avatardb == nil {
		return nil
	}
	avatar := &proto.Avatar{
		SkilltreeList:     make([]*proto.AvatarSkillTree, 0),
		Exp:               avatardb.Exp,
		BaseAvatarId:      avatardb.AvatarId,
		Rank:              avatardb.Rank,
		EquipmentUniqueId: avatardb.EquipmentUniqueId,
		EquipRelicList:    make([]*proto.EquipRelic, 0),
		TakenRewards:      avatardb.TakenRewards,
		FirstMetTimestamp: avatardb.FirstMetTimeStamp,
		Promotion:         avatardb.PromoteLevel,
		Level:             avatardb.Level,
	}
	for _, skill := range g.GetSkillTreeList(avatarId) {
		if avatarId/1000 == 8 {
			break
		}
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
			Slot:          id,
			RelicUniqueId: relic,
		}
		avatar.EquipRelicList = append(avatar.EquipRelicList, equipRelic)
	}

	return avatar
}

type BattleAvatar struct {
	AvatarId  uint32 // 角色id
	IsAssist  bool   // 是否助战
	AssistUid uint32 // 助战uid
}

func (g *GamePlayer) GetProtoBattleAvatar(bAList map[uint32]*BattleAvatar) []*proto.BattleAvatar {
	battleAvatarList := make([]*proto.BattleAvatar, 0)
	for id, bA := range bAList {
		var avatarBin *spb.AvatarBin
		if bA.IsAssist {
			// TODO 助战情况
		} else {
			avatarBin = g.GetAvatarBinById(bA.AvatarId)
		}
		if avatarBin == nil {
			continue
		}
		battleAvatar := &proto.BattleAvatar{
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
		// 获取技能
		for _, skill := range avatarBin.GetSkilltreeList() {
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
		battleAvatarList = append(battleAvatarList, battleAvatar)
	}
	return battleAvatarList
}
