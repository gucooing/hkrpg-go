package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

// 队伍更新通知
func (g *GamePlayer) SyncLineupNotify(index uint32, isBattleLine bool) {
	rsq := new(proto.SyncLineupNotify)
	if isBattleLine {
		rsq.Lineup = g.GetBattleLineUpPb(index)
	} else {
		rsq.Lineup = g.GetLineUpPb(index)
	}
	g.Send(cmd.SyncLineupNotify, rsq)
}

func (g *GamePlayer) SceneGroupRefreshScNotify(index uint32) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshInfo: g.GetSceneGroupRefreshInfoByLineUP(g.GetLineUpById(index), g.GetPosPb(), g.GetRotPb()),
	}
	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *GamePlayer) HandleGetAllLineupDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAllLineupDataScRsp)
	rsp.LineupList = make([]*proto.LineupInfo, 0)
	db := g.GetLineUp()
	rsp.CurIndex = db.MainLineUp

	// 添加普通队伍
	for i := 0; i < MaxLineupList; i++ {
		lineupList := g.GetLineUpPb(uint32(i))
		rsp.LineupList = append(rsp.LineupList, lineupList)
	}

	g.Send(cmd.GetAllLineupDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetCurLineupDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurLineupDataScRsp)
	rsp.Lineup = g.GetLineUpPb(g.GetLineUp().MainLineUp)

	g.Send(cmd.GetCurLineupDataScRsp, rsp)
}

func (g *GamePlayer) HandleJoinLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.JoinLineupCsReq, payloadMsg)
	req := msg.(*proto.JoinLineupCsReq)

	g.UnDbLineUp(req.Index, req.Slot, req.BaseAvatarId)

	// 队伍更新通知
	g.SyncLineupNotify(req.Index, false)
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := new(proto.LineupAvatar)
	g.Send(cmd.JoinLineupScRsp, rsp)
}

func (g *GamePlayer) HandleSwitchLineupIndexCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SwitchLineupIndexCsReq, payloadMsg)
	req := msg.(*proto.SwitchLineupIndexCsReq)

	lineUpDb := g.GetLineUp()
	lineUpDb.MainLineUp = req.Index
	// 队伍更新通知
	g.SyncLineupNotify(req.Index, false)
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := &proto.SwitchLineupIndexScRsp{Index: req.Index}

	g.Send(cmd.SwitchLineupIndexScRsp, rsp)
}

func (g *GamePlayer) HandleSwapLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SwapLineupCsReq, payloadMsg)
	req := msg.(*proto.SwapLineupCsReq)

	// 交换角色
	g.SwapLineup(req.Index, req.SrcSlot, req.DstSlot)

	// 队伍更新通知
	g.SyncLineupNotify(req.Index, false)
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := &proto.SwapLineupCsReq{}

	g.Send(cmd.SwapLineupScRsp, rsp)
}

func (g *GamePlayer) SetLineupNameCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetLineupNameCsReq, payloadMsg)
	req := msg.(*proto.SetLineupNameCsReq)
	db := g.GetLineUpById(req.Index)
	db.Name = req.Name

	// 队伍更新通知
	g.SyncLineupNotify(req.Index, false)
	g.SceneGroupRefreshScNotify(req.Index)

	rsp := &proto.SetLineupNameScRsp{
		Index: req.Index,
		Name:  req.Name,
	}

	g.Send(cmd.SetLineupNameScRsp, rsp)
}

func (g *GamePlayer) ReplaceLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ReplaceLineupCsReq, payloadMsg)
	req := msg.(*proto.ReplaceLineupCsReq)

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
	}
	db.LeaderSlot = 0
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	for _, avatarList := range req.Slots {
		db.AvatarIdList[avatarList.Slot] = &spb.LineAvatarList{AvatarId: avatarList.Id, Slot: avatarList.Slot}
	}

	// 队伍更新通知
	g.SyncLineupNotify(index, isBattleLine)
	if isBattleLine {
		// 将角色属性拷贝出来
		for _, avatar := range req.Slots {
			avatarBin := g.GetAvatarBinById(avatar.Id)
			g.CopyBattleAvatar(avatarBin)
		}
	} else {
		g.SceneGroupRefreshScNotify(req.Index)
	}

	g.Send(cmd.ReplaceLineupScRsp, nil)
}

func (g *GamePlayer) ChangeLineupLeaderCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ChangeLineupLeaderCsReq, payloadMsg)
	req := msg.(*proto.ChangeLineupLeaderCsReq)

	rsp := &proto.ChangeLineupLeaderScRsp{Slot: req.Slot}

	db := g.GetCurLineUp()
	db.LeaderSlot = req.Slot

	g.Send(cmd.ChangeLineupLeaderScRsp, rsp)
}

func (g *GamePlayer) QuitLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.QuitLineupCsReq, payloadMsg)
	req := msg.(*proto.QuitLineupCsReq)
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
	g.SyncLineupNotify(req.Index, false)
	g.SceneGroupRefreshScNotify(req.Index)

	g.Send(cmd.QuitLineupScRsp, nil)
}
