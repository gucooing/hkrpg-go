package Game

import (
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
	}
	return g.PlayerPb.LineUp
}

func (g *Game) GetLineUpById(index uint32) *spb.Line {
	return g.PlayerPb.LineUp.LineUpList[index]
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
