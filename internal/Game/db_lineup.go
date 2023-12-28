package Game

type DbLineUp struct {
	LineUpList   []*LineUp
	MainLineUp   uint32 // 当前上场队伍
	MainAvatarId uint32 // 当前上场角色
}

type LineUp struct {
	Name         string
	AvatarIdList []uint32
}

func NewDbLineUp(data *PlayerData) *PlayerData {
	data.DbLineUp = new(DbLineUp)
	data.DbLineUp.LineUpList = []*LineUp{
		{Name: "Team 1", AvatarIdList: make([]uint32, 4)},
		{Name: "Team 2", AvatarIdList: make([]uint32, 4)},
		{Name: "Team 3", AvatarIdList: make([]uint32, 4)},
		{Name: "Team 4", AvatarIdList: make([]uint32, 4)},
		{Name: "Team 5", AvatarIdList: make([]uint32, 4)},
		{Name: "Team 6", AvatarIdList: make([]uint32, 4)},
	}
	data.DbLineUp.MainLineUp = 0
	data.DbLineUp.MainAvatarId = 0
	return data
}

// 队伍更新
func (g *Game) UnDbLineUp(index uint32, Slot uint32, avatarId uint32) {
	g.Player.DbLineUp.LineUpList[index].AvatarIdList[Slot] = avatarId
}

// 交换角色
func (g *Game) SwapLineup(index, src_slot, dst_slot uint32) {
	lineUpList := g.Player.DbLineUp.LineUpList[index]
	lineUpList.AvatarIdList[src_slot], lineUpList.AvatarIdList[dst_slot] = lineUpList.AvatarIdList[dst_slot], lineUpList.AvatarIdList[src_slot]
}

func (g *Game) GetSceneAvatarId() uint32 {
	return g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList[g.Player.DbLineUp.MainAvatarId]
}
