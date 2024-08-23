package model

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func newPojokNostalgia() *spb.PojokNostalgia {
	return &spb.PojokNostalgia{}
}

func (g *PlayerData) GetPojokNostalgia() *spb.PojokNostalgia {
	db := g.GetBasicBin()
	if db.PojokNostalgia == nil {
		db.PojokNostalgia = newPojokNostalgia()
	}
	return db.PojokNostalgia
}

/******************************以太战线******************************/

// 获取以太战线
func (g *PlayerData) GetAetherDivide() *spb.AetherDivide {
	db := g.GetPojokNostalgia()
	if db.AetherDivide == nil {
		db.AetherDivide = newAetherDivide()
	}
	return db.AetherDivide
}

func newAetherDivide() *spb.AetherDivide {
	db := &spb.AetherDivide{
		AvatarList:      make(map[uint32]*spb.AetherDivideAvatarInfo),
		Lineup:          make(map[uint32]*spb.AetherDivideLineup),
		AetherSkillList: make(map[uint32]*spb.AetherSkill),
	}

	// add passive skill
	for _, conf := range gdconf.GetAetherDividePassiveSkillMap() {
		db.AetherSkillList[conf.ItemID] = &spb.AetherSkill{
			ItemId:        conf.ItemID,
			Num:           1,
			DressAvatarId: 0,
		}
	}
	// add avatar
	for _, conf := range gdconf.GetAetherDivideSpiritMap() {
		db.AvatarList[conf.AvatarID] = &spb.AetherDivideAvatarInfo{
			AvatarId:     conf.AvatarID,
			Promotion:    conf.MaxPromotion, // max
			PassiveSkill: make(map[uint32]uint32),
			CurSp:        1000,
			MaxSp:        uint32(conf.SPMax.Value) * 1000,
		}
	}

	// add lineup
	db.Lineup[1] = &spb.AetherDivideLineup{
		Index:      1,
		AvatarList: []uint32{6001},
	}

	return db
}

func (g *PlayerData) GetAetherDivideAvatarInfoById(id uint32) *spb.AetherDivideAvatarInfo {
	db := g.GetAetherDivide()
	return db.AvatarList[id]
}

func (g *PlayerData) GetAetherSkillById(id uint32) *spb.AetherSkill {
	db := g.GetAetherDivide()
	return db.AetherSkillList[id]
}

func (g *PlayerData) GetAetherDivideLineupById(id uint32) *spb.AetherDivideLineup {
	db := g.GetAetherDivide()
	return db.Lineup[id]
}

func (g *PlayerData) GetAetherDivideSpiritInfo(avatarId uint32) *proto.AetherDivideSpiritInfo {
	db := g.GetAetherDivideAvatarInfoById(avatarId)
	if db == nil {
		return nil
	}
	info := &proto.AetherDivideSpiritInfo{
		PassiveSkill: make([]*proto.PassiveSkillItem, 0),
		AvatarId:     db.AvatarId,
		Promotion:    db.Promotion,
		SpBar: &proto.SpBarInfo{
			CurSp: db.CurSp,
			MaxSp: db.MaxSp,
		},
	}
	for slot, itemId := range db.PassiveSkill {
		info.PassiveSkill = append(info.PassiveSkill, &proto.PassiveSkillItem{
			Slot:   slot,
			ItemId: itemId,
		})
	}
	return info
}

func (g *PlayerData) GetAetherSkillInfo(itemId uint32) *proto.AetherSkillInfo {
	db := g.GetAetherSkillById(itemId)
	if db == nil {
		return nil
	}
	info := &proto.AetherSkillInfo{
		ItemId:        db.ItemId,
		DressAvatarId: db.DressAvatarId,
		Num:           db.Num,
	}
	return info
}

func (g *PlayerData) GetAetherDivideLineupInfo(index uint32) *proto.AetherDivideLineupInfo {
	db := g.GetAetherDivideLineupById(index)
	if db == nil {
		return nil
	}
	info := &proto.AetherDivideLineupInfo{
		AvatarList: db.AvatarList,
		Slot:       db.Index,
	}
	return info
}

type AetherAvatar struct {
	AvatarId uint32
	Type     spb.AvatarType
}

func (g *PlayerData) GetAetherAvatarInfoList(avatarInfo []*AetherAvatar) []*proto.AetherAvatarInfo {
	infoList := make([]*proto.AetherAvatarInfo, 0)
	for index, avatar := range avatarInfo {
		if avatar.Type == spb.AvatarType_AVATAR_TRIAL_TYPE {
			avatarConf := gdconf.GetAetherDivideSpiritTrial(avatar.AvatarId)
			if avatarConf == nil {
				continue
			}
			infoList = append(infoList, &proto.AetherAvatarInfo{
				Id:        avatarConf.SpiritID,
				Index:     uint32(index),
				Promotion: avatarConf.Promotion,
				Type:      proto.AetherdivideSpiritLineupType_AETHERDIVIDE_SPIRIT_LINEUP_TRIAL,
				SpBar: &proto.SpBarInfo{
					CurSp: 3000,
					MaxSp: 3000,
				},
				PassiveSkill: make([]uint32, 0),
			})
			continue
		}
		if db := g.GetAetherDivideAvatarInfoById(avatar.AvatarId); db != nil {
			info := &proto.AetherAvatarInfo{
				Id:        db.AvatarId,
				Index:     uint32(index),
				Promotion: db.Promotion,
				Type:      proto.AetherdivideSpiritLineupType_AETHERDIVIDE_SPIRIT_LINEUP_NORMAL,
				SpBar: &proto.SpBarInfo{
					CurSp: db.CurSp,
					MaxSp: db.MaxSp,
				},
				PassiveSkill: make([]uint32, 0),
			}
			for _, skill := range db.PassiveSkill {
				info.PassiveSkill = append(info.PassiveSkill, skill)
			}
			infoList = append(infoList, info)
		}
	}
	return infoList
}

/******************************分割线******************************/
