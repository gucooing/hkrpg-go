package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

// 获取以太战线
func (g *PlayerData) GetAetherDivide() *spb.AetherDivide {
	db := g.GetActivity()
	if db.AetherDivide == nil {
		db.AetherDivide = &spb.AetherDivide{
			AvatarList:      newAetherDivideAvatar(),
			Lineup:          newAetherDivideLineup(),
			AetherSkillList: newAetherSkill(),
		}
	}
	return db.AetherDivide
}

func newAetherDivideLineup() map[uint32]*spb.AetherDivideLineup {
	db := make(map[uint32]*spb.AetherDivideLineup)
	db[1] = &spb.AetherDivideLineup{
		Index:      1,
		AvatarList: []uint32{6010, 6006},
	}
	return db
}

func newAetherSkill() map[uint32]*spb.AetherSkill {
	db := make(map[uint32]*spb.AetherSkill)
	return db
}

func newAetherDivideAvatar() map[uint32]*spb.AetherDivideAvatarInfo {
	db := make(map[uint32]*spb.AetherDivideAvatarInfo)
	for _, id := range []uint32{6010, 6006} {
		if conf := gdconf.GetAetherDivideSpirit(id); conf != nil {
			db[conf.AvatarID] = &spb.AetherDivideAvatarInfo{
				AvatarId:     conf.AvatarID,
				Promotion:    1,
				PassiveSkill: make(map[uint32]uint32),
				CurSp:        0,
				MaxSp:        uint32(conf.SPMax.Value) * 1000,
			}
		}
	}

	return db
}

func (g *PlayerData) GetAetherDivideAvatar() map[uint32]*spb.AetherDivideAvatarInfo {
	db := g.GetAetherDivide()
	if db.AvatarList == nil {
		db.AvatarList = newAetherDivideAvatar()
	}
	return db.AvatarList
}

func (g *PlayerData) AddAetherDivideAvatar(avatarId uint32) {
	conf := gdconf.GetAetherDivideSpirit(avatarId)
	if conf == nil {
		logger.Debug("add AetherDivideAvatar error:%v", avatarId)
		return
	}
	db := g.GetAetherDivideAvatar()
	if db[avatarId] != nil {
		return
	}
}

func (g *PlayerData) GetAetherDivideLineup() map[uint32]*spb.AetherDivideLineup {
	db := g.GetAetherDivide()
	if db.Lineup == nil {
		db.Lineup = newAetherDivideLineup()
	}
	return db.Lineup
}

func (g *PlayerData) GetAetherSkill() map[uint32]*spb.AetherSkill {
	db := g.GetAetherDivide()
	if db.AetherSkillList == nil {
		db.AetherSkillList = newAetherSkill()
	}
	return db.AetherSkillList
}

func (g *PlayerData) GetAetherDivideAvatarInfoById(id uint32) *spb.AetherDivideAvatarInfo {
	db := g.GetAetherDivideAvatar()
	return db[id]
}

func (g *PlayerData) GetAetherSkillById(id uint32) *spb.AetherSkill {
	db := g.GetAetherSkill()
	return db[id]
}

func (g *PlayerData) GetAetherDivideLineupById(id uint32) *spb.AetherDivideLineup {
	db := g.GetAetherDivideLineup()
	return db[id]
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

func (g *PlayerData) GetAetherAvatarrMap(avatarList []uint32) []*AetherAvatar {
	avatarMap := make([]*AetherAvatar, 0)
	if len(avatarList) != 0 {
		for _, id := range avatarList {
			db := g.GetAetherDivideAvatarInfoById(id)
			if db == nil {
				avatarConf := gdconf.GetAetherDivideSpiritTrial(id)
				if avatarConf != nil {
					avatarMap = append(avatarMap, &AetherAvatar{
						AvatarId: id,
						Type:     spb.AvatarType_AVATAR_TRIAL_TYPE,
					})
				}
			} else {
				avatarMap = append(avatarMap, &AetherAvatar{
					AvatarId: id,
					Type:     spb.AvatarType_AVATAR_TYPE_NONE,
				})
			}
		}
	} else {
		lineup := g.GetAetherDivideLineupById(1)
		if lineup != nil {
			for _, avatarId := range lineup.AvatarList {
				avatarMap = append(avatarMap, &AetherAvatar{
					AvatarId: avatarId,
					Type:     spb.AvatarType_AVATAR_TYPE_NONE,
				})
			}
		}
	}

	return avatarMap
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

func (g *PlayerData) GetAetherDivideBattleInfo(battleBackup *BattleBackup) *proto.AetherDivideBattleInfo {
	if ((battleBackup.Sce == nil || len(battleBackup.Sce.EvenIdList) == 0) && len(battleBackup.StageIDList) == 0) ||
		len(battleBackup.AetherAvatarList) == 0 {
		logger.Warn("异常战斗请求")
		return nil
	}
	monsterWaveList, stageID := g.GetSceneMonsterWave(battleBackup.Sce.EvenIdList, battleBackup.WorldLevel, battleBackup)
	if battleBackup.StageID == 0 {
		battleBackup.StageID = stageID
	}
	info := &proto.AetherDivideBattleInfo{
		BattleAvatarList: g.GetAetherAvatarInfoList(battleBackup.AetherAvatarList),
		MonsterWaveList:  monsterWaveList,
		BattleId:         battleBackup.BattleId,
		BuffList:         make([]*proto.BattleBuff, 0),
		LogicRandomSeed:  gdconf.GetLoadingDesc(),
		StageId:          battleBackup.StageID,
	}
	return info
}
