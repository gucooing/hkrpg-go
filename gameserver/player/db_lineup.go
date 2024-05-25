package player

import (
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
)

func (g *GamePlayer) NewLineUp() *spb.LineUp {
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
		db.LineUp = &spb.LineUp{
			MainLineUp:     0,
			Mp:             MaxMp,
			LineUpList:     make(map[uint32]*spb.Line),
			BattleLineList: make(map[uint32]*spb.Line),
		}
		db.LineUp.LineUpList[0] = &spb.Line{
			Name:         "hkrpg",
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:   0,
		}
		db.LineUp.LineUpList[0].AvatarIdList[0] = &spb.LineAvatarList{AvatarId: uint32(g.GetAvatar().CurMainAvatar), Slot: 0}
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
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:   0,
		}
		db.LineUpList[0].AvatarIdList[0] = &spb.LineAvatarList{AvatarId: uint32(g.GetAvatar().CurMainAvatar), Slot: 0}
	}
	if db.LineUpList[index] == nil {
		db.LineUpList[index] = &spb.Line{
			Name:         "",
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:   0,
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
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:   0,
		}
	}
	return db.BattleLineList[index]
}

func (g *GamePlayer) GetCurLineUp() *spb.Line {
	db := g.GetLineUp()
	return db.LineUpList[db.MainLineUp]
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
		return nil
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

/*****************************************功能方法****************************/

func (g *GamePlayer) GetLineUpPb(id uint32) *proto.LineupInfo {
	db := g.GetLineUpById(id)
	var wtmLeaderSlot = false
	if db.AvatarIdList[db.LeaderSlot] == nil {
		wtmLeaderSlot = true
	}
	avatarList := make([]*proto.LineupAvatar, 0)
	for slot, lineAvatar := range db.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
		if wtmLeaderSlot {
			db.LeaderSlot = slot
		}
		if avatarBin == nil {
			continue
		}
		lineupAvatar := &proto.LineupAvatar{
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
		avatarList = append(avatarList, lineupAvatar)
	}
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      db.LeaderSlot,
		AvatarList:      avatarList,
		ExtraLineupType: proto.ExtraLineupType_LINEUP_NONE,
		Index:           id,
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
	}
	var wtmLeaderSlot = false
	if db.AvatarIdList[db.LeaderSlot] == nil {
		wtmLeaderSlot = true
	}
	avatarList := make([]*proto.LineupAvatar, 0)
	for slot, lineAvatar := range db.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		if wtmLeaderSlot {
			db.LeaderSlot = slot
		}
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
		if avatarBin == nil {
			lineupAvatar := &proto.LineupAvatar{
				AvatarType: proto.AvatarType_AVATAR_FORMAL_TYPE,
				Slot:       slot,
				Satiety:    0,
				Hp:         10000,
				Id:         lineAvatar.AvatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: 6000,
					MaxSp: 10000,
				},
			}
			avatarList = append(avatarList, lineupAvatar)
		} else {
			lineupAvatar := &proto.LineupAvatar{
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
			avatarList = append(avatarList, lineupAvatar)
		}
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
