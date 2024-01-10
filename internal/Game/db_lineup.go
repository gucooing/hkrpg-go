package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type DbLineUp struct {
	LineUpList   []*LineUp
	MainLineUp   uint32 // 当前上场队伍
	MainAvatarId uint32 // 当前上场角色
}

type LineUp struct {
	Name string
	// ExtraLineupType proto.ExtraLineupType
	AvatarIdList []uint32
}

func (g *Game) GetLineUp() *spb.LineUp {
	if g.PlayerPb.LineUp == nil {
		g.PlayerPb.LineUp = &spb.LineUp{
			MainLineUp:   0,
			MainAvatarId: 0,
			LineUpList:   make(map[uint32]*spb.Line),
		}
		g.PlayerPb.LineUp.LineUpList[0] = &spb.Line{Name: "hkrpg", AvatarIdList: []uint32{uint32(g.GetAvatar().CurMainAvatar), 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_NONE}
		g.PlayerPb.LineUp.LineUpList[1] = &spb.Line{Name: "hkrpg", AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_NONE}
		g.PlayerPb.LineUp.LineUpList[2] = &spb.Line{Name: "hkrpg", AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_NONE}
		g.PlayerPb.LineUp.LineUpList[3] = &spb.Line{Name: "hkrpg", AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_NONE}
		g.PlayerPb.LineUp.LineUpList[4] = &spb.Line{Name: "hkrpg", AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_NONE}
		g.PlayerPb.LineUp.LineUpList[5] = &spb.Line{Name: "hkrpg", AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_NONE}
		g.PlayerPb.LineUp.LineUpList[6] = &spb.Line{AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_CHALLENGE}
		g.PlayerPb.LineUp.LineUpList[7] = &spb.Line{AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_CHALLENGE_2}
		g.PlayerPb.LineUp.LineUpList[8] = &spb.Line{AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_CHALLENGE_3}
		g.PlayerPb.LineUp.LineUpList[9] = &spb.Line{AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_ROGUE}
		g.PlayerPb.LineUp.LineUpList[10] = &spb.Line{AvatarIdList: []uint32{0, 0, 0, 0}, ExtraLineupType: spb.ExtraLineupType_LINEUP_STAGE_TRIAL}
	}
	return g.PlayerPb.LineUp
}

func (g *Game) GetLineUpById(index uint32) *spb.Line {
	return g.GetLineUp().LineUpList[index]
}

func (g *Game) GetLineUpPb(id uint32) *proto.LineupInfo {
	lineUp := g.GetLineUpById(id)
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           id,
		ExtraLineupType: proto.ExtraLineupType(lineUp.ExtraLineupType),
		MaxMp:           5,
		Mp:              5,
		Name:            lineUp.Name,
		PlaneId:         0,
	}
	for slot, avatarId := range lineUp.AvatarIdList {
		if avatarId == 0 {
			continue
		}
		avatar := g.GetAvatar().Avatar[avatarId]
		if avatar == nil {
			lineupAvatar := &proto.LineupAvatar{
				AvatarType: proto.AvatarType_AVATAR_TRIAL_TYPE,
				Slot:       uint32(slot),
				Satiety:    0,
				Hp:         10000,
				Id:         avatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: 5000,
					MaxSp: 10000,
				},
			}
			lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
		} else {
			lineupAvatar := &proto.LineupAvatar{
				AvatarType: proto.AvatarType(avatar.AvatarType),
				Slot:       uint32(slot),
				Satiety:    0,
				Hp:         avatar.Hp,
				Id:         avatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: avatar.SpBar.CurSp,
					MaxSp: avatar.SpBar.MaxSp,
				},
			}
			lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
		}
	}
	return lineupList
}

// 队伍更新
func (g *Game) UnDbLineUp(index uint32, Slot uint32, avatarId uint32) {
	g.PlayerPb.LineUp.LineUpList[index].AvatarIdList[Slot] = avatarId
}

// 交换角色
func (g *Game) SwapLineup(index, src_slot, dst_slot uint32) {
	lineUpList := g.PlayerPb.LineUp.LineUpList[index]
	lineUpList.AvatarIdList[src_slot], lineUpList.AvatarIdList[dst_slot] = lineUpList.AvatarIdList[dst_slot], lineUpList.AvatarIdList[src_slot]
}

func (g *Game) GetSceneAvatarId() uint32 {
	return g.PlayerPb.LineUp.LineUpList[g.PlayerPb.LineUp.MainLineUp].AvatarIdList[g.PlayerPb.LineUp.MainAvatarId]
}
