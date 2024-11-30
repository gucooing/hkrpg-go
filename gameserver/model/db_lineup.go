package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	MaxLineupList = 20 // 设置最大普通队伍数
	MaxMp         = 5  // 设置最大队伍能量
	Challenge_1   = 1  // 第一个忘却之庭队伍
	Challenge_2   = 2  // 第二个忘却之庭队伍
	Rogue         = 3  // 第一个模拟宇宙队伍
	RogueTourn    = 4  // 差分宇宙
	Activity      = 5  // 角色试用队伍
	Raid          = 6  // 故事性/副本
)

func NewLineUp() *spb.LineUp {
	return &spb.LineUp{
		MainLineUp:     0,
		Mp:             MaxMp,
		LineUpList:     nil,
		BattleLineList: nil,
		StoryLineList:  nil,
	}
}

func (g *PlayerData) GetLineUp() *spb.LineUp {
	db := g.GetBasicBin()
	if db.LineUp == nil {
		db.LineUp = NewLineUp()
	}
	return db.LineUp
}

// 获取常规队伍的能量
func (g *PlayerData) GetLineUpMp() uint32 {
	db := g.GetCurLineUp()
	if db.LineType != spb.ExtraLineupType_LINEUP_NONE {
		return db.Mp
	}
	return g.GetLineUp().Mp
}

// 获取指定队伍的能量
func (g *PlayerData) GetLineupDbMp(db *spb.Line) uint32 {
	if db.LineType != spb.ExtraLineupType_LINEUP_NONE {
		return db.Mp
	}
	return g.GetLineUp().Mp
}

// 添加当前队伍能量
func (g *PlayerData) AddLineUpMp(mp uint32) {
	db := g.GetCurLineUp()
	if db.LineType != spb.ExtraLineupType_LINEUP_NONE {
		if db.Mp += mp; db.Mp > 5 {
			db.Mp = 5
		}
	} else {
		lineUp := g.GetLineUp()
		if lineUp.Mp += mp; lineUp.Mp > 5 {
			lineUp.Mp = 5
		}
	}
}

// 删除当前队伍能量
func (g *PlayerData) DelLineUpMp(mp uint32) {
	db := g.GetCurLineUp()
	if db.LineType != spb.ExtraLineupType_LINEUP_NONE {
		if db.Mp -= mp; db.Mp < 0 {
			db.Mp = 0
		}
	} else {
		lineUp := g.GetLineUp()
		if lineUp.Mp -= mp; lineUp.Mp < 0 {
			lineUp.Mp = 0
		}
	}
}

// 通过id获取常规队伍
func (g *PlayerData) GetLineUpById(index uint32) *spb.Line {
	if index > MaxLineupList {
		return nil
	}
	db := g.GetLineUp()
	if db.LineUpList == nil {
		db.LineUpList = make(map[uint32]*spb.Line)
		db.LineUpList[0] = &spb.Line{
			Name:         "hkrpg",
			AvatarIdList: make(map[uint32]*spb.LineAvatarList, 4),
			LeaderSlot:   0,
			Index:        0,
		}
		db.LineUpList[0].AvatarIdList[0] = &spb.LineAvatarList{AvatarId: 8001, Slot: 0}
		db.LineUpList[0].AvatarIdList[1] = &spb.LineAvatarList{AvatarId: 1001, Slot: 1}
	}
	if db.LineUpList[index] == nil {
		db.LineUpList[index] = &spb.Line{
			Name:         "",
			AvatarIdList: make(map[uint32]*spb.LineAvatarList, 4),
			LeaderSlot:   0,
			Index:        index,
		}
	}
	return db.LineUpList[index]
}

// 通过id获取战斗队伍
func (g *PlayerData) GetBattleLineUpById(index uint32) *spb.Line {
	db := g.GetLineUp()
	if db.BattleLineList == nil {
		db.BattleLineList = make(map[uint32]*spb.Line)
	}
	if db.BattleLineList[index] == nil {
		db.BattleLineList[index] = &spb.Line{
			AvatarIdList: make(map[uint32]*spb.LineAvatarList, 4),
			LeaderSlot:   0,
		}
	}
	return db.BattleLineList[index]
}

// 获取当前普通队伍
func (g *PlayerData) GetCurNontLine() *spb.Line {
	return g.GetLineUpById(g.GetLineUp().MainLineUp)
}

// 获取故事线队伍
func (g *PlayerData) GetStoryLine() map[uint32]*spb.Line {
	db := g.GetLineUp()
	if db.StoryLineList == nil {
		db.StoryLineList = make(map[uint32]*spb.Line)
	}
	return db.StoryLineList
}

// 新建故事线队伍
func (g *PlayerData) NewStoryLine(storyLineID uint32) {
	curLineup := g.GetCurLineUp()
	db := g.GetLineUp()
	if db.StoryLineList == nil {
		db.StoryLineList = make(map[uint32]*spb.Line)
	}
	conf := gdconf.GetStroyLineTrialAvatarData(storyLineID)
	if conf == nil {
		return
	}
	storyLine := new(spb.Line)
	storyLine.LeaderSlot = 0
	storyLine.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	for slot, lineAvatar := range curLineup.AvatarIdList {
		storyLine.AvatarIdList[slot] = &spb.LineAvatarList{
			Slot:           slot,
			AvatarId:       lineAvatar.AvatarId,
			LineAvatarType: 0,
		}
	}
	storyLine.AvatarIdList[0] = &spb.LineAvatarList{
		Slot:           0,
		AvatarId:       conf.CaptainAvatarID,
		LineAvatarType: spb.AvatarType_AVATAR_TRIAL_TYPE,
	}
	db.StoryLineList[storyLineID] = storyLine
}

// 通过id获取故事线队伍
func (g *PlayerData) GetStoryLineById(index uint32) *spb.Line {
	db := g.GetLineUp()
	if db.StoryLineList == nil {
		db.StoryLineList = make(map[uint32]*spb.Line)
	}
	if db.StoryLineList[index] == nil {
		db.StoryLineList[index] = &spb.Line{
			AvatarIdList: make(map[uint32]*spb.LineAvatarList, 4),
			LineType:     spb.ExtraLineupType_LINEUP_STAGE_TRIAL,
			LeaderSlot:   0,
		}
	}
	return db.StoryLineList[index]
}

// 获取当前故事线队伍
func (g *PlayerData) GetCurChangeStory() *spb.Line {
	if db := g.GetChangeStory(); db == nil || !g.IsChangeStory() {
		return g.GetCurLineUp()
	} else {
		return g.GetStoryLineById(db.CurChangeStory)
	}
}

// 新建队伍
func (g *PlayerData) NewLineByAvatarList(trialList []uint32) {
	db := g.GetCurLineUp()
	if len(trialList) < 5 {
		return
	}
	db.LeaderSlot = 0
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList, 4)
	for slot, id := range trialList {
		if slot == 4 {
			continue
		}
		if g.GetAvatarBinById(id) != nil {
			db.AvatarIdList[uint32(slot)] = &spb.LineAvatarList{
				AvatarId:       id,
				Slot:           uint32(slot),
				LineAvatarType: spb.AvatarType_AVATAR_FORMAL_TYPE,
			}
		}
		if gdconf.GetSpecialAvatarById(id) != nil {
			db.AvatarIdList[uint32(slot)] = &spb.LineAvatarList{
				AvatarId:       id,
				Slot:           uint32(slot),
				LineAvatarType: spb.AvatarType_AVATAR_TRIAL_TYPE,
			}
		}
		if id == trialList[4] {
			db.LeaderSlot = uint32(slot)
		}
	}
}

// 向当前队伍中添加一个角色
func (g *PlayerData) AddCurLineUpAvatar(avatarId uint32) *spb.LineAvatarList {
	db := g.GetCurLineUp()
	var lineAvatarType spb.AvatarType
	if g.GetAvatarBinById(avatarId) != nil {
		lineAvatarType = spb.AvatarType_AVATAR_FORMAL_TYPE
	}
	if gdconf.GetSpecialAvatarById(avatarId) != nil {
		lineAvatarType = spb.AvatarType_AVATAR_TRIAL_TYPE
	}
	lineAvatar := &spb.LineAvatarList{
		AvatarId:       avatarId,
		LineAvatarType: lineAvatarType,
	}
	isUp := false
	for i := 0; i < 4; i++ {
		if db.AvatarIdList[uint32(i)] == nil {
			lineAvatar.Slot = uint32(i)
			db.AvatarIdList[uint32(i)] = lineAvatar
			isUp = true
			break
		} else {
			if db.AvatarIdList[uint32(i)].AvatarId == 0 {
				lineAvatar.Slot = uint32(i)
				db.AvatarIdList[uint32(i)] = lineAvatar
				isUp = true
				break
			}
		}
	}
	if !isUp {
		lineAvatar.Slot = 3
		db.AvatarIdList[3] = lineAvatar
	}
	return lineAvatar
}

// 在当前队伍中删除目标角色
func (g *PlayerData) DelCurLineUpAvatar(trialAvatarId uint32) {
	db := g.GetCurLineUp()
	isJh := false
	for id, info := range db.AvatarIdList {
		if info.AvatarId == trialAvatarId {
			if id == db.LeaderSlot {
				isJh = true
			}
			delete(db.AvatarIdList, id)
		}
	}
	isDelTrial := true
	for id, info := range db.AvatarIdList {
		if info.AvatarId != 0 {
			isDelTrial = false
			if isJh {
				db.LeaderSlot = id
			}
		}
	}
	if isDelTrial {
		db = g.GetCurLineUp() // 更改当前队伍
		db.LeaderSlot = 0
		db.AvatarIdList[0] = &spb.LineAvatarList{
			Slot:           0,
			AvatarId:       8001,
			LineAvatarType: 0,
		}
		db.AvatarIdList[1] = &spb.LineAvatarList{
			Slot:           1,
			AvatarId:       1001,
			LineAvatarType: 0,
		}
	}
}

// 获取当前上场角色id
func (g *PlayerData) GetSceneAvatarId() uint32 {
	db := g.GetCurLineUp()
	if db.AvatarIdList[db.LeaderSlot] == nil {
		logger.Error("获取的角色不存在于队伍中")
		return 8001
	} else {
		return db.AvatarIdList[db.LeaderSlot].AvatarId
	}
}

// 队伍更新
func (g *PlayerData) UnDbLineUp(index uint32, Slot uint32, avatarId uint32) {
	db := g.GetLineUpById(index)
	db.AvatarIdList[Slot] = &spb.LineAvatarList{AvatarId: avatarId, Slot: Slot}
	db.LeaderSlot = Slot
}

// // 队伍更新
// func (g *PlayerData) UnDbAvatarLineUp(db *spb.Line, avatarId, newAvatarId uint32) {
// 	for _, info := range db.AvatarIdList {
// 		if info.AvatarId == avatarId {
// 			info.AvatarId = newAvatarId
// 		}
// 	}
// 	g.SyncLineupNotify(db)
// }

// 获取当前队伍（普通/战斗
func (g *PlayerData) GetCurLineUp() *spb.Line {
	if g.IsChangeStory() {
		return g.GetCurChangeStory()
	}
	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_NONE:
		return g.GetCurNontLine()
	case spb.BattleType_Battle_CHALLENGE:
		return g.GetChallengeLineUp()
	case spb.BattleType_Battle_CHALLENGE_Story:
		return g.GetChallengeLineUp()
	case spb.BattleType_Battle_QUSET_ROGUE:
		return g.GetBattleLineUpById(Rogue)
	case spb.BattleType_Battle_ROGUE_TOURN:
		return g.GetBattleLineUpById(RogueTourn)
	case spb.BattleType_Battle_TrialActivity:
		return g.GetBattleLineUpById(Activity)
	case spb.BattleType_Battle_RAID:
		return g.GetBattleLineUpById(Raid)
	default:
		return g.GetCurNontLine()
	}
}

// 获取忘却之庭队伍
func (g *PlayerData) GetChallengeLineUp() *spb.Line {
	challengeStatus := g.GetCurChallenge()
	switch challengeStatus.CurStage {
	case 1:
		return g.GetBattleLineUpById(Challenge_1)
	case 2:
		return g.GetBattleLineUpById(Challenge_2)
	default:
		logger.Debug("[UID:%v]没有此忘却之庭队伍id:%v", g.GetBasicBin().Uid, challengeStatus.CurStage)
		return g.GetCurNontLine()
	}
}

// 筛选试用队伍中 男/女主角
func (g *PlayerData) SpecialMainAvatar(trialAvatarId uint32) (bool, spb.AvatarType) {
	conf := gdconf.GetSpecialAvatarById(trialAvatarId)
	if conf == nil {
		return true, spb.AvatarType_AVATAR_FORMAL_TYPE
	}
	avatarDb := g.GetAvatar()
	if conf.AvatarID/1000 == 8 {
		switch avatarDb.Gender {
		case spb.Gender_GenderMan:
			if conf.AvatarID%2 == 0 {
				return false, spb.AvatarType_AVATAR_TRIAL_TYPE
			}
		case spb.Gender_GenderWoman:
			if conf.AvatarID%2 != 0 {
				return false, spb.AvatarType_AVATAR_TRIAL_TYPE
			}
		}
	}
	return true, spb.AvatarType_AVATAR_TRIAL_TYPE
}

/*****************************************功能方法****************************/

func (g *PlayerData) GetLineUpPb(db *spb.Line) *proto.LineupInfo {
	if db == nil {
		return nil
	}
	avatarList := make([]*proto.LineupAvatar, 0)
	for slot, lineAvatar := range db.AvatarIdList {
	tx:
		info := &proto.LineupAvatar{
			Slot:    slot,
			Satiety: 0,
			Hp:      10000,
			SpBar: &proto.SpBarInfo{
				CurSp: 6000,
				MaxSp: 10000,
			},
			Id:         lineAvatar.AvatarId,
			AvatarType: proto.AvatarType(lineAvatar.LineAvatarType),
		}

		switch lineAvatar.LineAvatarType {
		case spb.AvatarType_AVATAR_TYPE_NONE:
			lineAvatar.LineAvatarType = spb.AvatarType_AVATAR_FORMAL_TYPE
			goto tx
		case spb.AvatarType_AVATAR_FORMAL_TYPE:
			avatarInfo := g.GetAvatarBinById(lineAvatar.AvatarId)
			if avatarInfo == nil {
				delete(db.AvatarIdList, slot)
				continue
			}
			info.Hp = avatarInfo.Hp
			info.SpBar.CurSp = avatarInfo.SpBar.CurSp
			info.SpBar.MaxSp = avatarInfo.SpBar.MaxSp

		case spb.AvatarType_AVATAR_TRIAL_TYPE:
		case spb.AvatarType_AVATAR_ASSIST_TYPE:
			// TODO 拉取援助角色信息
		default:
			logger.Error("@LogTag(player_error_%v)@异常的队伍角色Type:%s", g.GetBasicBin().Uid, lineAvatar.LineAvatarType.String())
			continue
		}
		avatarList = append(avatarList, info)
	}

	lineupList := &proto.LineupInfo{
		AvatarList:            avatarList,
		Mp:                    g.GetLineupDbMp(db),
		IsVirtual:             false,
		Index:                 db.Index,
		PlaneId:               0,
		Name:                  db.Name,
		ExtraLineupType:       proto.ExtraLineupType(db.LineType),
		MaxMp:                 MaxMp,
		LeaderSlot:            db.LeaderSlot,
		GameStoryLineId:       0,
		StoryLineAvatarIdList: make([]uint32, 0),
	}
	if changeStory := g.GetCurChangeStoryInfo(); changeStory != nil {
		lineupList.GameStoryLineId = changeStory.ChangeStoryId
	}
	return lineupList
}

func (g *PlayerData) GetLineupAvatarDataList(db *spb.Line) []*proto.LineupAvatarData {
	infoList := make([]*proto.LineupAvatarData, 0)
	if db == nil {
		return infoList
	}
	for _, lineAvatar := range db.AvatarIdList {

		info := &proto.LineupAvatarData{
			Hp:         10000,
			Id:         lineAvatar.AvatarId,
			AvatarType: proto.AvatarType(lineAvatar.LineAvatarType),
		}

		switch lineAvatar.LineAvatarType {
		case spb.AvatarType_AVATAR_FORMAL_TYPE:
			avatarInfo := g.GetAvatarBinById(lineAvatar.AvatarId)
			if avatarInfo == nil {
				continue
			}
			info.Hp = avatarInfo.Hp

		case spb.AvatarType_AVATAR_TRIAL_TYPE:
		case spb.AvatarType_AVATAR_ASSIST_TYPE:
			// TODO 拉取援助角色信息
		default:
			logger.Error("@LogTag(player_error_%v)@异常的队伍角色Type:%s", g.GetBasicBin().Uid, lineAvatar.LineAvatarType.String())
			continue
		}
		infoList = append(infoList, info)
	}

	return infoList
}
