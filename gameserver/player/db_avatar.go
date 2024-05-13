package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetAvatar() *spb.Avatar {
	db := g.GetPlayerPb()
	if db.Avatar == nil {
		db.Avatar = &spb.Avatar{
			Avatar:            make(map[uint32]*spb.AvatarBin),
			Gender:            spb.Gender_GenderMan,
			CurMainAvatar:     spb.HeroBasicType_BoyWarrior,
			HeroBasicTypeInfo: g.GetHeroBasicTypeInfo(),
		}
	}
	return db.Avatar
}

func (g *GamePlayer) GetAvatarBinById(avatarId uint32) *spb.AvatarBin {
	bin := g.GetAvatar()
	if bin.Avatar == nil {
		bin.Avatar = make(map[uint32]*spb.AvatarBin)
	}
	return bin.Avatar[avatarId]
}

func (g *GamePlayer) GetHeroBasicTypeInfo() []*spb.HeroBasicTypeInfo {
	heroBasic := make([]*spb.HeroBasicTypeInfo, 0)
	if g.PlayerPb.Avatar == nil || g.PlayerPb.Avatar.HeroBasicTypeInfo == nil {
		heroBasicTypeInfo := &spb.HeroBasicTypeInfo{
			Rank:          0,
			BasicType:     spb.HeroBasicType_BoyWarrior,
			SkillTreeList: g.GetSkillTreeList(uint32(spb.HeroBasicType_BoyWarrior)),
		}
		heroBasic = append(heroBasic, heroBasicTypeInfo)

		heroBasicTypeInfo = &spb.HeroBasicTypeInfo{
			Rank:          0,
			BasicType:     spb.HeroBasicType_BoyKnight,
			SkillTreeList: g.GetSkillTreeList(uint32(spb.HeroBasicType_BoyKnight)),
		}
		heroBasic = append(heroBasic, heroBasicTypeInfo)
	} else {
		g.PlayerPb.Avatar.HeroBasicTypeInfo = heroBasic
	}
	return heroBasic
}

func (g *GamePlayer) GetSkillTreeList(avatarId uint32) []*spb.AvatarSkillBin {
	skilltreeList := make([]*spb.AvatarSkillBin, 0)
	if avatarId/1000 == 8 || g.PlayerPb.Avatar.Avatar[avatarId] == nil {
		for id, level := range gdconf.GetAvatarSkilltreeListById(avatarId) {
			avatarSkillBin := &spb.AvatarSkillBin{
				PointId: id,
				Level:   level,
			}
			skilltreeList = append(skilltreeList, avatarSkillBin)
		}
		return skilltreeList
	}
	if g.PlayerPb.Avatar.Avatar[avatarId].SkilltreeList == nil {
		for id, level := range gdconf.GetAvatarSkilltreeListById(avatarId) {
			avatarSkillBin := &spb.AvatarSkillBin{
				PointId: id,
				Level:   level,
			}
			skilltreeList = append(skilltreeList, avatarSkillBin)
		}
		g.PlayerPb.Avatar.Avatar[avatarId].SkilltreeList = skilltreeList
	}
	return g.PlayerPb.Avatar.Avatar[avatarId].SkilltreeList
}

func (g *GamePlayer) AddAvatar(avatarId uint32) {
	var pileItem []*Material
	if g.PlayerPb.Avatar.Avatar[avatarId] != nil {
		pileItem = append(pileItem, &Material{
			Tid: avatarId + 10000,
			Num: 1,
		})
		g.AddMaterial(pileItem)
		return
	}

	g.PlayerPb.Avatar.Avatar[avatarId] = &spb.AvatarBin{
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

/***************************************功能接口**************************/

func (g *GamePlayer) GetProtoAvatarById(avatarId uint32) *proto.Avatar {
	avatardb := g.GetAvatar().Avatar[avatarId]
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

func (g *GamePlayer) GetProtoBattleAvatar(bAList []BattleAvatar) []*proto.BattleAvatar {
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
			Index:         uint32(id),
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
