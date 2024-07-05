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
	Activity      = 5  // 角色试用队伍
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
	isTrial := false
	if g.GetAvatarById(trialAvatarId) != nil {
		lineAvatarType = spb.LineAvatarType_LineAvatarType_MI
	}
	if gdconf.GetSpecialAvatarById(trialAvatarId) != nil {
		isTrial = true
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
	g.AddAvatarSceneGroupRefreshScNotify(trialAvatarId, isTrial, g.GetPosPb(), g.GetRotPb())
	g.SyncLineupNotify(db)
}

func (g *GamePlayer) DelTrialAvatar(trialAvatarId uint32) {
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
	// g.GetAddAvatarSceneEntityRefreshInfo(db, g.GetPosPb(), g.GetRotPb())
	g.SyncLineupNotify(db)
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
	case spb.BattleType_Battle_ROGUE_TOURN:
		return g.GetBattleLineUpById(RogueTourn)
	case spb.BattleType_Battle_TrialActivity:
		return g.GetBattleLineUpById(Activity)
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
	g.SyncLineupNotify(g.GetCurLineUp())
}

func (g *GamePlayer) SetBattleLineUp(index uint32, avatarList []uint32) {
	if avatarList == nil { // 没有传入角色
		avatarList = make([]uint32, 0)
		for _, info := range g.GetCurLineUp().AvatarIdList {
			avatarList = append(avatarList, info.AvatarId)
		}
	}
	var lineUpType spb.ExtraLineupType
	var avatarType spb.LineAvatarType
	switch index {
	case Challenge_1:
		lineUpType = spb.ExtraLineupType_LINEUP_CHALLENGE
	case Challenge_2:
		lineUpType = spb.ExtraLineupType_LINEUP_CHALLENGE_2
	case Rogue:
		lineUpType = spb.ExtraLineupType_LINEUP_ROGUE
	case RogueTourn:
		lineUpType = spb.ExtraLineupType_LINEUP_TOURN_ROGUE
	case Activity:
		lineUpType = spb.ExtraLineupType_LINEUP_STAGE_TRIAL
		avatarType = spb.LineAvatarType_LineAvatarType_TRIAL
	}
	db := g.GetBattleLineUpById(index)
	db.LeaderSlot = 0
	db.LineType = lineUpType
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	for id, avatarId := range avatarList {
		db.AvatarIdList[uint32(id)] = &spb.LineAvatarList{AvatarId: avatarId, Slot: uint32(id), LineAvatarType: avatarType}
	}
	// 拷贝角色
	for _, avatar := range avatarList {
		avatarBin := g.GetAvatarBinById(avatar)
		g.CopyBattleAvatar(avatarBin)
	}
	g.SyncLineupNotify(db)
}

/*****************************************功能方法****************************/

func (g *GamePlayer) GetLineUpPb(db *spb.Line) *proto.LineupInfo {
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
			avatarList = append(avatarList, g.GetLineupAvatar(lineAvatar.AvatarId, slot))
		case spb.LineAvatarType_LineAvatarType_TRIAL:
			avatarList = append(avatarList, g.GetTrialLineupAvatar(lineAvatar.AvatarId, slot))
		default:
			continue
		}
	}

	lineupList := &proto.LineupInfo{
		AvatarList:      avatarList,
		Mp:              g.GetLineUpMp(),
		IsVirtual:       false,
		Index:           db.Index,
		PlaneId:         0,
		Name:            db.Name,
		ExtraLineupType: proto.ExtraLineupType(db.LineType),
		MaxMp:           MaxMp,
		LeaderSlot:      db.LeaderSlot,
	}
	return lineupList
}

func (g *GamePlayer) GetLineupAvatar(avatarId, index uint32) *proto.LineupAvatar {
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

func (g *GamePlayer) GetTrialLineupAvatar(avatarId, index uint32) *proto.LineupAvatar {
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
