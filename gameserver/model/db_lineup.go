package model

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	MaxLineupList = 20 // 设置最大普通队伍数
	MaxMp         = 5  // 设置最大队伍能量
	Challenge_1   = 1  // 第一个忘却之庭队伍
	Challenge_2   = 2  // 第二个忘却之庭队伍
	Rogue         = 3  // 第一个模拟宇宙队伍
	RogueTourn    = 4  // 差分宇宙
	Activity      = 5  // 角色试用队伍
	Raid          = 6  // 副本队伍
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

func (g *PlayerData) GetLineUpMp() uint32 {
	db := g.GetBattleLineUp()
	if db.LineType != spb.ExtraLineupType_LINEUP_NONE {
		return db.Mp
	}
	return g.GetLineUp().Mp
}

func (g *PlayerData) AddLineUpMp(mp uint32) {
	db := g.GetBattleLineUp()
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

func (g *PlayerData) DelLineUpMp(mp uint32) {
	db := g.GetBattleLineUp()
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

func (g *PlayerData) GetLineUpById(index uint32) *spb.Line {
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

func (g *PlayerData) GetStoryLine() map[uint32]*spb.Line {
	db := g.GetLineUp()
	if db.StoryLineList == nil {
		db.StoryLineList = make(map[uint32]*spb.Line)
	}
	return db.StoryLineList
}

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
		LineAvatarType: spb.LineAvatarType_LineAvatarType_TRIAL,
	}
	db.StoryLineList[storyLineID] = storyLine
}

func (g *PlayerData) NewTrialLine(trialList []uint32) {
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
		if g.GetAvatarById(id) != nil {
			db.AvatarIdList[uint32(slot)] = &spb.LineAvatarList{
				AvatarId:       id,
				Slot:           uint32(slot),
				LineAvatarType: spb.LineAvatarType_LineAvatarType_MI,
			}
		}
		if gdconf.GetSpecialAvatarById(id) != nil {
			db.AvatarIdList[uint32(slot)] = &spb.LineAvatarList{
				AvatarId:       id,
				Slot:           uint32(slot),
				LineAvatarType: spb.LineAvatarType_LineAvatarType_TRIAL,
			}
		}
		if id == trialList[4] {
			db.LeaderSlot = uint32(slot)
		}
	}
}

func (g *PlayerData) GetTrialAvatar(trialAvatarId uint32) *spb.LineAvatarList {
	db := g.GetBattleLineUp()
	var lineAvatarType spb.LineAvatarType
	if g.GetAvatarById(trialAvatarId) != nil {
		lineAvatarType = spb.LineAvatarType_LineAvatarType_MI
	}
	if gdconf.GetSpecialAvatarById(trialAvatarId) != nil {
		lineAvatarType = spb.LineAvatarType_LineAvatarType_TRIAL
	}
	lineAvatar := &spb.LineAvatarList{
		AvatarId:       trialAvatarId,
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

func (g *PlayerData) DelTrialAvatar(trialAvatarId uint32) {
	db := g.GetBattleLineUp()
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

func (g *PlayerData) GetCurLineUp() *spb.Line {
	if g.IsChangeStory() {
		db := g.GetChangeStory()
		return g.GetStoryLineById(db.CurChangeStory)
	}
	db := g.GetLineUpById(g.GetLineUp().MainLineUp)
	return db
}

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

// 交换角色
func (g *PlayerData) SwapLineup(index, src_slot, dst_slot uint32) {
	db := g.GetLineUpById(index)
	src := db.AvatarIdList[src_slot]
	dst := db.AvatarIdList[dst_slot]
	src.Slot = dst.Slot
	dst.Slot = src.Slot
	db.AvatarIdList[src_slot] = dst
	db.AvatarIdList[dst_slot] = src
}

func (g *PlayerData) GetBattleLineUp() *spb.Line {
	if g.IsChangeStory() {
		return g.GetCurLineUp()
	}
	status := g.GetBattleStatus()
	switch status {
	case spb.BattleType_Battle_NONE:
		return g.GetCurLineUp()
	case spb.BattleType_Battle_CHALLENGE:
		return g.GetChallengeLineUp()
	case spb.BattleType_Battle_CHALLENGE_Story:
		return g.GetChallengeLineUp()
	case spb.BattleType_Battle_ROGUE:
		return g.GetBattleLineUpById(Rogue)
	case spb.BattleType_Battle_ROGUE_TOURN:
		return g.GetBattleLineUpById(RogueTourn)
	case spb.BattleType_Battle_TrialActivity:
		return g.GetBattleLineUpById(Activity)
	case spb.BattleType_Battle_RAID:
		return g.GetBattleLineUpById(Raid)
	default:
		return g.GetCurLineUp()
	}
	return g.GetCurLineUp() // 多余了
}

func (g *PlayerData) GetChallengeLineUp() *spb.Line {
	challengeStatus := g.GetCurChallenge()
	switch challengeStatus.CurStage {
	case 1:
		return g.GetBattleLineUpById(Challenge_1)
	case 2:
		return g.GetBattleLineUpById(Challenge_2)
	default:
		logger.Debug("[UID:%v]没有此忘却之庭队伍id:%v", g.GetBasicBin().Uid, challengeStatus.CurStage)
		return g.GetCurLineUp()
	}
}

func (g *PlayerData) SpecialMainAvatar(trialAvatarId uint32) (bool, spb.LineAvatarType) {
	conf := gdconf.GetSpecialAvatarById(trialAvatarId)
	if conf == nil {
		return true, spb.LineAvatarType_LineAvatarType_MI
	}
	avatarDb := g.GetAvatar()
	if conf.AvatarID/1000 == 8 {
		switch avatarDb.Gender {
		case spb.Gender_GenderMan:
			if conf.AvatarID%2 == 0 {
				return false, spb.LineAvatarType_LineAvatarType_TRIAL
			}
		case spb.Gender_GenderWoman:
			if conf.AvatarID%2 != 0 {
				return false, spb.LineAvatarType_LineAvatarType_TRIAL
			}
		}
	}
	return true, spb.LineAvatarType_LineAvatarType_TRIAL
}

/*****************************************功能方法****************************/

func (g *PlayerData) GetLineUpPb(db *spb.Line) *proto.LineupInfo {
	if db == nil {
		return nil
	}
	avatarList := make([]*proto.LineupAvatar, 0)
	// if db.AvatarIdList[db.LeaderSlot] == nil {
	//
	// }
	for slot, lineAvatar := range db.AvatarIdList {
		switch lineAvatar.LineAvatarType {
		case spb.LineAvatarType_LineAvatarType_MI:
			x := g.GetLineupAvatar(lineAvatar.AvatarId, slot)
			if x == nil {
				delete(db.AvatarIdList, slot)
				continue
			}
			avatarList = append(avatarList, x)
		case spb.LineAvatarType_LineAvatarType_TRIAL:
			avatarList = append(avatarList, g.GetTrialLineupAvatar(lineAvatar.AvatarId, slot))
			continue
		default:
			continue
		}
	}

	lineupList := &proto.LineupInfo{
		AvatarList:            avatarList,
		Mp:                    g.GetLineUpMp(),
		IsVirtual:             false,
		Index:                 db.Index,
		PlaneId:               0,
		Name:                  db.Name,
		ExtraLineupType:       proto.ExtraLineupType(db.LineType),
		MaxMp:                 MaxMp,
		LeaderSlot:            db.LeaderSlot,
		GameStoryLineId:       0,
		StoryLineAvatarIdList: make([]uint32, 0),
		Sus:                   make([]uint32, 0),
	}
	if changeStory := g.GetCurChangeStoryInfo(); changeStory != nil {
		lineupList.GameStoryLineId = changeStory.ChangeStoryId
		// lineupList.Sus = []uint32{15}
	}
	return lineupList
}

func (g *PlayerData) GetLineupAvatar(avatarId, index uint32) *proto.LineupAvatar {
	db := g.GetAvatarBinById(avatarId)
	if db == nil {
		return nil
	}
	info := &proto.LineupAvatar{
		Slot:    index,
		Satiety: 0,
		Hp:      db.Hp,
		SpBar: &proto.SpBarInfo{
			CurSp: db.SpBar.CurSp,
			MaxSp: db.SpBar.MaxSp,
		},
		Id:         avatarId,
		AvatarType: proto.AvatarType(db.AvatarType),
	}
	return info
}

func (g *PlayerData) GetTrialLineupAvatar(avatarId, index uint32) *proto.LineupAvatar {
	info := &proto.LineupAvatar{
		Slot:    index,
		Satiety: 0,
		Hp:      10000,
		SpBar: &proto.SpBarInfo{
			CurSp: 6000,
			MaxSp: 10000,
		},
		Id:         avatarId,
		AvatarType: proto.AvatarType_AVATAR_TRIAL_TYPE,
	}
	return info
}

func (g *PlayerData) GetLineupAvatarDataList(db *spb.Line) []*proto.LineupAvatarData {
	lADList := make([]*proto.LineupAvatarData, 0)
	if db == nil {
		return lADList
	}
	for slot, lineAvatar := range db.AvatarIdList {
		switch lineAvatar.LineAvatarType {
		case spb.LineAvatarType_LineAvatarType_MI:
			lADList = append(lADList, g.GetLineupAvatarData(lineAvatar.AvatarId, slot))
		case spb.LineAvatarType_LineAvatarType_TRIAL:
			lADList = append(lADList, g.GetTrialLineupAvatarData(lineAvatar.AvatarId, slot))
		default:
			continue
		}
	}

	return lADList
}

func (g *PlayerData) GetLineupAvatarData(avatarId, index uint32) *proto.LineupAvatarData {
	db := g.GetAvatarBinById(avatarId)
	if db == nil {
		return nil
	}
	info := &proto.LineupAvatarData{
		Hp:         db.Hp,
		Id:         avatarId,
		AvatarType: proto.AvatarType(db.AvatarType),
	}
	return info
}

func (g *PlayerData) GetTrialLineupAvatarData(avatarId, index uint32) *proto.LineupAvatarData {
	info := &proto.LineupAvatarData{
		Hp:         10000,
		Id:         avatarId,
		AvatarType: proto.AvatarType_AVATAR_TRIAL_TYPE,
	}
	return info
}
