package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 队伍更新通知
func (g *GamePlayer) SyncLineupNotify(db *spb.Line) {
	rsq := &proto.SyncLineupNotify{
		ReasonList: make([]proto.SyncLineupReason, 0),
		Lineup:     g.GetLineUpPb(db),
	}
	g.Send(cmd.SyncLineupNotify, rsq)
}

func (g *GamePlayer) SceneGroupRefreshScNotify(index uint32) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: g.GetSceneGroupRefreshInfoByLineUP(g.GetLineUpById(index), g.GetPosPb(), g.GetRotPb()),
	}
	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *GamePlayer) HandleGetAllLineupDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetAllLineupDataScRsp)
	rsp.LineupList = make([]*proto.LineupInfo, 0)
	db := g.GetLineUp()
	rsp.CurIndex = db.MainLineUp

	// 添加普通队伍
	for i := 0; i < MaxLineupList; i++ {
		lineupList := g.GetLineUpPb(g.GetLineUpById(uint32(i)))
		rsp.LineupList = append(rsp.LineupList, lineupList)
	}

	g.Send(cmd.GetAllLineupDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetCurLineupDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetCurLineupDataScRsp)
	rsp.Lineup = g.GetLineUpPb(g.GetCurLineUp())

	g.Send(cmd.GetCurLineupDataScRsp, rsp)
}

func (g *GamePlayer) GetLineupAvatarDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetLineupAvatarDataScRsp)
	rsp.AvatarDataList = g.GetLineupAvatarDataList(g.GetCurLineUp())

	g.Send(cmd.GetLineupAvatarDataScRsp, rsp)
}

func (g *GamePlayer) HandleJoinLineupCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.JoinLineupCsReq)

	g.UnDbLineUp(req.Index, req.Slot, req.BaseAvatarId)

	// 队伍更新通知
	g.SyncLineupNotify(g.GetLineUpById(req.Index))
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := new(proto.LineupAvatar)
	g.Send(cmd.JoinLineupScRsp, rsp)
}

func (g *GamePlayer) HandleSwitchLineupIndexCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SwitchLineupIndexCsReq)

	lineUpDb := g.GetLineUp()
	lineUpDb.MainLineUp = req.Index
	// 队伍更新通知
	g.SyncLineupNotify(g.GetCurLineUp())
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := &proto.SwitchLineupIndexScRsp{Index: req.Index}

	g.Send(cmd.SwitchLineupIndexScRsp, rsp)
}

func (g *GamePlayer) HandleSwapLineupCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SwapLineupCsReq)

	// 交换角色
	g.SwapLineup(req.Index, req.SrcSlot, req.DstSlot)

	// 队伍更新通知
	g.SyncLineupNotify(g.GetLineUpById(req.Index))
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := &proto.SwapLineupCsReq{}

	g.Send(cmd.SwapLineupScRsp, rsp)
}

func (g *GamePlayer) SetLineupNameCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetLineupNameCsReq)
	db := g.GetLineUpById(req.Index)
	db.Name = req.Name

	// 队伍更新通知
	g.SyncLineupNotify(g.GetLineUpById(req.Index))
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := &proto.SetLineupNameScRsp{
		Index: req.Index,
		Name:  req.Name,
	}

	g.Send(cmd.SetLineupNameScRsp, rsp)
}

func (g *GamePlayer) ReplaceLineupCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ReplaceLineupCsReq)

	index := req.Index
	isBattleLine := false
	var db *spb.Line

	switch req.ExtraLineupType {
	case proto.ExtraLineupType_LINEUP_NONE:
		db = g.GetLineUpById(index)
	case proto.ExtraLineupType_LINEUP_CHALLENGE:
		index = Challenge_1
		db = g.GetBattleLineUpById(index)
		isBattleLine = true
	case proto.ExtraLineupType_LINEUP_CHALLENGE_2:
		index = Challenge_2
		db = g.GetBattleLineUpById(index)
		isBattleLine = true
	case proto.ExtraLineupType_LINEUP_ROGUE:
		index = Rogue
		db = g.GetBattleLineUpById(index)
		isBattleLine = true
	}
	db.LeaderSlot = req.LeaderSlot
	db.LineType = spb.ExtraLineupType(req.ExtraLineupType)
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	for _, avatarList := range req.LineupSlotList {
		db.AvatarIdList[avatarList.Slot] = &spb.LineAvatarList{AvatarId: avatarList.Id, Slot: avatarList.Slot}
	}

	// 队伍更新通知
	g.SyncLineupNotify(db)
	if isBattleLine {
		// 将角色属性拷贝出来
		for _, avatar := range req.LineupSlotList {
			avatarBin := g.GetAvatarBinById(avatar.Id)
			g.CopyBattleAvatar(avatarBin)
		}
	} else {
		g.SceneGroupRefreshScNotify(req.Index)
	}

	g.Send(cmd.ReplaceLineupScRsp, &proto.ReplaceLineupScRsp{})
}

func (g *GamePlayer) ChangeLineupLeaderCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ChangeLineupLeaderCsReq)

	rsp := &proto.ChangeLineupLeaderScRsp{Slot: req.Slot}

	db := g.GetCurLineUp()
	db.LeaderSlot = req.Slot

	g.Send(cmd.ChangeLineupLeaderScRsp, rsp)
}

func (g *GamePlayer) QuitLineupCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.QuitLineupCsReq)
	db := g.GetCurLineUp()

	for id, lineAvatar := range db.AvatarIdList {
		if lineAvatar.AvatarId == req.BaseAvatarId {
			if db.LeaderSlot == id {
				db.LeaderSlot = 0
			}
			delete(db.AvatarIdList, id)
		}
	}
	// 队伍更新通知
	g.SyncLineupNotify(db)
	g.SceneGroupRefreshScNotify(req.Index)

	g.Send(cmd.QuitLineupScRsp, nil)
}
