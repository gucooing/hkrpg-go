package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 队伍更新通知
func (g *GamePlayer) SyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	rsq.Lineup = g.GetLineUpPb(index)

	g.SceneGroupRefreshScNotify()

	g.Send(cmd.SyncLineupNotify, rsq)
}

func (g *GamePlayer) SceneGroupRefreshScNotify() {
	notify := new(proto.SceneGroupRefreshScNotify)
	notify.GroupRefreshInfo = make([]*proto.SceneGroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.SceneGroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}
	pos := g.GetPos()
	rot := g.GetRot()
	for _, lineup := range g.PlayerPb.LineUp.LineUpList[g.PlayerPb.LineUp.MainLineUp].AvatarIdList {
		if lineup == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				Actor: &proto.SceneActorInfo{
					AvatarType:   proto.AvatarType(g.PlayerPb.Avatar.Avatar[lineup].AvatarType),
					BaseAvatarId: lineup,
				},
				Motion: &proto.MotionInfo{
					Pos: &proto.Vector{
						X: pos.X,
						Y: pos.Y,
						Z: pos.Z,
					},
					Rot: &proto.Vector{
						X: rot.X,
						Y: rot.Y,
						Z: rot.Z,
					},
				},
				EntityId: entityId,
			},
		}
		g.GetSceneEntity().AvatarEntity[entityId] = &AvatarEntity{
			AvatarId: lineup,
			GroupId:  0,
		}
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	}
	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *GamePlayer) HandleGetAllLineupDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAllLineupDataScRsp)
	rsp.LineupList = make([]*proto.LineupInfo, 0)
	rsp.CurIndex = 0

	for id, _ := range g.GetLineUp().GetLineUpList() {
		lineupList := g.GetLineUpPb(id)
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
	g.GetLineUp().MainAvatarId = 0

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.LineupAvatar)
	g.Send(cmd.JoinLineupScRsp, rsp)
}

func (g *GamePlayer) HandleSwitchLineupIndexCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SwitchLineupIndexCsReq, payloadMsg)
	req := msg.(*proto.SwitchLineupIndexCsReq)

	lineUpDb := g.GetLineUp()
	lineUpDb.MainLineUp = req.Index
	lineUpDb.MainAvatarId = 0
	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SwitchLineupIndexScRsp{Index: req.Index}

	g.Send(cmd.SwitchLineupIndexScRsp, rsp)
}

func (g *GamePlayer) HandleSwapLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SwapLineupCsReq, payloadMsg)
	req := msg.(*proto.SwapLineupCsReq)

	// 交换角色
	g.SwapLineup(req.Index, req.SrcSlot, req.DstSlot)
	g.GetLineUp().MainAvatarId = 0

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SwapLineupCsReq{}

	g.Send(cmd.SwapLineupScRsp, rsp)
}

func (g *GamePlayer) SetLineupNameCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetLineupNameCsReq, payloadMsg)
	req := msg.(*proto.SetLineupNameCsReq)
	g.GetLineUp().LineUpList[req.Index].Name = req.Name

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SetLineupNameScRsp{
		Index: req.Index,
		Name:  req.Name,
	}

	g.Send(cmd.SetLineupNameScRsp, rsp)
}

func (g *GamePlayer) ReplaceLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ReplaceLineupCsReq, payloadMsg)
	req := msg.(*proto.ReplaceLineupCsReq)
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	index := req.Index
	lineUpDb := g.GetLineUp()

	switch req.ExtraLineupType {
	case proto.ExtraLineupType_LINEUP_CHALLENGE:
		index = 6
	case proto.ExtraLineupType_LINEUP_CHALLENGE_2:
		index = 7
	case proto.ExtraLineupType_LINEUP_CHALLENGE_3:
		index = 8
	case proto.ExtraLineupType_LINEUP_ROGUE:
		index = 9
	case proto.ExtraLineupType_LINEUP_STAGE_TRIAL:
		index = 10
	}

	lineUpDb.LineUpList[index].AvatarIdList = []uint32{0, 0, 0, 0}
	for _, avatarId := range req.Slots {
		lineUpDb.LineUpList[index].AvatarIdList[avatarId.Slot] = avatarId.Id
	}

	lineUpDb.MainAvatarId = 0

	// 队伍更新通知
	g.SyncLineupNotify(index)

	g.Send(cmd.ReplaceLineupScRsp, rsp)
}

func (g *GamePlayer) ChangeLineupLeaderCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ChangeLineupLeaderCsReq, payloadMsg)
	req := msg.(*proto.ChangeLineupLeaderCsReq)

	rsp := &proto.ChangeLineupLeaderScRsp{Slot: req.Slot}

	g.GetLineUp().MainAvatarId = req.Slot

	g.Send(cmd.ChangeLineupLeaderScRsp, rsp)
}

func (g *GamePlayer) QuitLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.QuitLineupCsReq, payloadMsg)
	req := msg.(*proto.QuitLineupCsReq)
	lineUpDb := g.GetLineUp()

	for id, avatarId := range lineUpDb.LineUpList[req.Index].AvatarIdList {
		if avatarId == req.BaseAvatarId {
			lineUpDb.LineUpList[req.Index].AvatarIdList[id] = 0
		}
	}

	lineUpDb.MainAvatarId = 0
	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.QuitLineupScRsp, rsp)
}
