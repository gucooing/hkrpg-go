package player

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
)

func NewLineUp() *spb.LineUp {
	return &spb.LineUp{
		MainLineUp:     0,
		Mp:             MaxMp,
		LineUpList:     nil,
		BattleLineList: nil,
	}
}

func (g *GamePlayer) GetLineUp() *spb.LineUp {
	db := g.GetBasicBin()
	if db.LineUp == nil {
		db.LineUp = NewLineUp()
	}
	return db.LineUp
}

func (g *GamePlayer) GetLineUpMp() uint32 {
	db := g.GetLineUp()
	return db.Mp
}

func (g *GamePlayer) GetLineUpById(index uint32) *spb.Line {
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

func (g *GamePlayer) GetBattleLineUpById(index uint32) *spb.Line {
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

func (g *GamePlayer) NewTrialLine(trialList []uint32) {
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
		if id == 0 {
			db.AvatarIdList[uint32(slot)] = &spb.LineAvatarList{
				AvatarId:       id,
				Slot:           uint32(slot),
				LineAvatarType: 0,
			}
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

func (g *GamePlayer) GetTrialAvatar(trialAvatarId uint32) {
	db := g.GetCurLineUp()
	var lineAvatarType spb.LineAvatarType
	if g.GetAvatarById(trialAvatarId) != nil {
		lineAvatarType = spb.LineAvatarType_LineAvatarType_MI
	}
	if gdconf.GetSpecialAvatarById(trialAvatarId) != nil {
		lineAvatarType = spb.LineAvatarType_LineAvatarType_TRIAL
	}
	isUp := false
	for i := 0; i < 4; i++ {
		if db.AvatarIdList[uint32(i)] == nil {
			db.AvatarIdList[uint32(i)] = &spb.LineAvatarList{
				Slot:           uint32(i),
				AvatarId:       trialAvatarId,
				LineAvatarType: lineAvatarType,
			}
			isUp = true
			break
		} else {
			if db.AvatarIdList[uint32(i)].AvatarId == 0 {
				db.AvatarIdList[uint32(i)] = &spb.LineAvatarList{
					Slot:           uint32(i),
					AvatarId:       trialAvatarId,
					LineAvatarType: lineAvatarType,
				}
				isUp = true
				break
			}
		}
	}
	if !isUp {
		db.AvatarIdList[3] = &spb.LineAvatarList{
			Slot:           3,
			AvatarId:       trialAvatarId,
			LineAvatarType: lineAvatarType,
		}
	}
	g.AddAvatarSceneGroupRefreshScNotify(trialAvatarId, false, g.GetPosPb(), g.GetRotPb())
	g.SyncLineupNotifyByLineBin(db)
}

func (g *GamePlayer) DelTrialAvatar(trialAvatarId uint32) {
	db := g.GetCurLineUp()
	for _, id := range db.AvatarIdList {
		if id.AvatarId == trialAvatarId {
			id.AvatarId = 0
			id.LineAvatarType = 0
		}
	}
	isDelTrial := true
	for _, id := range db.AvatarIdList {
		if id.AvatarId != 0 {
			isDelTrial = false
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
	// g.GetAddAvatarSceneEntityRefreshInfo(db, g.GetPosPb(), g.GetRotPb())
	g.SyncLineupNotifyByLineBin(db)
}

func (g *GamePlayer) GetCurLineUp() *spb.Line {
	db := g.GetLineUpById(g.GetLineUp().MainLineUp)
	return db
}

func (g *GamePlayer) GetSceneAvatarId() uint32 {
	db := g.GetCurLineUp()
	if db.AvatarIdList[db.LeaderSlot] == nil {
		logger.Error("获取的角色不存在于队伍中")
		return 8001
	} else {
		return db.AvatarIdList[db.LeaderSlot].AvatarId
	}
}

// 队伍更新
func (g *GamePlayer) UnDbLineUp(index uint32, Slot uint32, avatarId uint32) {
	db := g.GetLineUpById(index)
	db.AvatarIdList[Slot] = &spb.LineAvatarList{AvatarId: avatarId, Slot: Slot}
	db.LeaderSlot = Slot
}

// 交换角色
func (g *GamePlayer) SwapLineup(index, src_slot, dst_slot uint32) {
	db := g.GetLineUpById(index)
	src := db.AvatarIdList[src_slot]
	dst := db.AvatarIdList[dst_slot]
	src.Slot = dst.Slot
	dst.Slot = src.Slot
	db.AvatarIdList[src_slot] = dst
	db.AvatarIdList[dst_slot] = src
}

func (g *GamePlayer) GetBattleLineUp() *spb.Line {
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
	case spb.BattleType_Battle_TrialActivity:
	default:
		return g.GetCurLineUp()
	}
	return g.GetCurLineUp() // 多余了
}

func (g *GamePlayer) GetChallengeLineUp() *spb.Line {
	challengeStatus := g.GetCurChallenge()
	switch challengeStatus.CurStage {
	case 1:
		return g.GetBattleLineUpById(Challenge_1)
	case 2:
		return g.GetBattleLineUpById(Challenge_2)
	default:
		logger.Debug("[UID:%v]没有此忘却之庭队伍id:%v", g.Uid, challengeStatus.CurStage)
		return g.GetCurLineUp()
	}
}

func (g *GamePlayer) AddLineUpMp(mp uint32) {
	db := g.GetLineUp()
	if db.Mp += mp; db.Mp > 5 {
		db.Mp = 5
	}
	// 更新通知
	g.SyncLineupNotify(db.MainLineUp, false)
}

func (g *GamePlayer) SetBattleLineUp(index uint32, avatarList []uint32) {
	db := g.GetBattleLineUpById(index)
	db.LeaderSlot = 0
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	for id, avatarId := range avatarList {
		db.AvatarIdList[uint32(id)] = &spb.LineAvatarList{AvatarId: avatarId, Slot: uint32(id)}
	}
	for _, avatar := range avatarList {
		avatarBin := g.GetAvatarBinById(avatar)
		g.CopyBattleAvatar(avatarBin)
	}
	g.SyncLineupNotify(index, true)
}

/*****************************************功能方法****************************/

func (g *GamePlayer) GetLineUpPb(db *spb.Line) *proto.LineupInfo {
	var wtmLeaderSlot = false
	if db.AvatarIdList[db.LeaderSlot] == nil || db.AvatarIdList[db.LeaderSlot].AvatarId == 0 {
		wtmLeaderSlot = true
	}
	avatarList := make([]*proto.LineupAvatar, 0)
	for slot, lineAvatar := range db.AvatarIdList {
		if wtmLeaderSlot {
			db.LeaderSlot = slot
		}
		lineupAvatar := &proto.LineupAvatar{}
		switch lineAvatar.LineAvatarType {
		case spb.LineAvatarType_LineAvatarType_MI:
			avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
			if avatarBin == nil {
				if slot == db.LeaderSlot {
					wtmLeaderSlot = true
				}
				continue
			}
			lineupAvatar = &proto.LineupAvatar{
				AvatarType: proto.AvatarType(avatarBin.AvatarType),
				Slot:       slot,
				Satiety:    0,
				Hp:         avatarBin.Hp,
				Id:         lineAvatar.AvatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: avatarBin.SpBar.CurSp,
					MaxSp: avatarBin.SpBar.MaxSp,
				},
			}
		case spb.LineAvatarType_LineAvatarType_TRIAL:
			lineupAvatar = &proto.LineupAvatar{
				AvatarType: proto.AvatarType_AVATAR_TRIAL_TYPE,
				Slot:       slot,
				Satiety:    0,
				Hp:         10000,
				Id:         lineAvatar.AvatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: 6000,
					MaxSp: 10000,
				},
			}
		default:
			continue
		}
		avatarList = append(avatarList, lineupAvatar)
	}
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      db.LeaderSlot,
		AvatarList:      avatarList,
		ExtraLineupType: proto.ExtraLineupType_LINEUP_NONE,
		Index:           db.Index,
		MaxMp:           MaxMp,
		Mp:              g.GetLineUpMp(),
		Name:            db.Name,
		PlaneId:         0,
	}
	return lineupList
}

func (g *GamePlayer) GetBattleLineUpPb(id uint32) *proto.LineupInfo {
	db := g.GetBattleLineUpById(id)
	var lineUpType proto.ExtraLineupType
	switch id {
	case Challenge_1:
		lineUpType = proto.ExtraLineupType_LINEUP_CHALLENGE
	case Challenge_2:
		lineUpType = proto.ExtraLineupType_LINEUP_CHALLENGE_2
	case Rogue:
		lineUpType = proto.ExtraLineupType_LINEUP_ROGUE
	case RogueTourn:
		lineUpType = proto.ExtraLineupType_LINEUP_TOURN_ROGUE
	}
	var wtmLeaderSlot = false
	if db.AvatarIdList[db.LeaderSlot] == nil || db.AvatarIdList[db.LeaderSlot].AvatarId == 0 {
		wtmLeaderSlot = true
	}
	avatarList := make([]*proto.LineupAvatar, 0)
	for slot, lineAvatar := range db.AvatarIdList {
		if lineAvatar.AvatarId == 0 {
			continue
		}
		if wtmLeaderSlot {
			db.LeaderSlot = slot
		}
		lineupAvatar := &proto.LineupAvatar{}
		switch lineAvatar.LineAvatarType {
		case spb.LineAvatarType_LineAvatarType_MI:
			avatarBin := g.GetBattleAvatarBinById(lineAvatar.AvatarId)
			if avatarBin == nil {
				continue
			}
			lineupAvatar = &proto.LineupAvatar{
				AvatarType: proto.AvatarType(avatarBin.AvatarType),
				Slot:       slot,
				Satiety:    0,
				Hp:         avatarBin.Hp,
				Id:         lineAvatar.AvatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: avatarBin.SpBar.CurSp,
					MaxSp: avatarBin.SpBar.MaxSp,
				},
			}
		case spb.LineAvatarType_LineAvatarType_TRIAL:
			lineupAvatar = &proto.LineupAvatar{
				AvatarType: proto.AvatarType_AVATAR_TRIAL_TYPE,
				Slot:       slot,
				Satiety:    0,
				Hp:         10000,
				Id:         lineAvatar.AvatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: 10000,
					MaxSp: 10000,
				},
			}
		default:
			continue
		}
		avatarList = append(avatarList, lineupAvatar)
	}
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      db.LeaderSlot,
		AvatarList:      avatarList,
		ExtraLineupType: lineUpType,
		Index:           0,
		MaxMp:           MaxMp,
		Mp:              g.GetLineUpMp(),
		Name:            db.Name,
		PlaneId:         0,
	}
	return lineupList
}
