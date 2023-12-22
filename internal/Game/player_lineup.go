package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 队伍更新通知
func (g *Game) SyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	lineUp := g.Player.DbLineUp.LineUpList[index]
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           index,
		ExtraLineupType: proto.ExtraLineupType_LINEUP_NONE,
		MaxMp:           5,
		Mp:              5,
		Name:            lineUp.Name,
		PlaneId:         0,
	}
	for slot, avatarId := range lineUp.AvatarIdList {
		if avatarId == 0 {
			continue
		}
		avatar := g.Player.DbAvatar.Avatar[avatarId]
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: avatar.Type,
			Slot:       uint32(slot),
			Satiety:    0,
			Hp:         avatar.Hp,
			Id:         avatarId,
			SpBar:      avatar.SpBar,
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsq.Lineup = lineupList

	g.SceneGroupRefreshScNotify()

	g.Send(cmd.SyncLineupNotify, rsq)
}

func (g *Game) SceneGroupRefreshScNotify() {
	notify := new(proto.SceneGroupRefreshScNotify)
	notify.GroupRefreshInfo = make([]*proto.SceneGroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.SceneGroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}
	pos := g.Player.Pos
	rot := g.Player.Rot
	for _, lineup := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
		if lineup == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
			UpdateType: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
						AvatarType:   g.Player.DbAvatar.Avatar[lineup].Type,
						BaseAvatarId: lineup,
					}},
					Motion: &proto.MotionInfo{
						Pos: &proto.Vector{
							X: int32(pos.X),
							Y: int32(pos.Y),
							Z: int32(pos.Z),
						},
						Rot: &proto.Vector{
							X: int32(rot.X),
							Y: int32(rot.Y),
							Z: int32(rot.Z),
						},
					},
					EntityId: entityId,
				},
			},
		}
		g.Player.EntityList[entityId] = lineup
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	}
	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *Game) HandleGetAllLineupDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAllLineupDataScRsp)
	rsp.LineupList = make([]*proto.LineupInfo, 0)
	rsp.CurIndex = 0

	for index, lineUp := range g.Player.DbLineUp.LineUpList {
		lineupList := &proto.LineupInfo{
			IsVirtual:       false,
			LeaderSlot:      0,
			AvatarList:      make([]*proto.LineupAvatar, 0),
			Index:           uint32(index),
			ExtraLineupType: proto.ExtraLineupType_LINEUP_NONE,
			MaxMp:           5,
			Mp:              5,
			Name:            lineUp.Name,
			PlaneId:         0,
		}
		for slot, avatarId := range lineUp.AvatarIdList {
			if avatarId == 0 {
				continue
			}
			avatar := g.Player.DbAvatar.Avatar[avatarId]
			lineupAvatar := &proto.LineupAvatar{
				AvatarType: avatar.Type,
				Slot:       uint32(slot),
				Satiety:    0,
				Hp:         avatar.Hp,
				Id:         avatarId,
				SpBar:      avatar.SpBar,
			}
			lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
		}
		rsp.LineupList = append(rsp.LineupList, lineupList)
	}

	g.Send(cmd.GetAllLineupDataScRsp, rsp)
}

func (g *Game) HandleGetCurLineupDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurLineupDataScRsp)
	lineUp := g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp]
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           g.Player.DbLineUp.MainLineUp,
		ExtraLineupType: proto.ExtraLineupType_LINEUP_NONE,
		MaxMp:           5,
		Mp:              5,
		Name:            lineUp.Name,
		PlaneId:         0,
	}
	for slot, avatarId := range lineUp.AvatarIdList {
		if avatarId == 0 {
			continue
		}
		avatar := g.Player.DbAvatar.Avatar[avatarId]
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: avatar.Type,
			Slot:       uint32(slot),
			Satiety:    0,
			Hp:         avatar.Hp,
			Id:         avatarId,
			SpBar:      avatar.SpBar,
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsp.Lineup = lineupList

	g.Send(cmd.GetCurLineupDataScRsp, rsp)
}

func (g *Game) HandleJoinLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.JoinLineupCsReq, payloadMsg)
	req := msg.(*proto.JoinLineupCsReq)

	g.UnDbLineUp(req.Index, req.Slot, req.BaseAvatarId)

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.LineupAvatar)
	g.Send(cmd.JoinLineupScRsp, rsp)
}

func (g *Game) HandleSwitchLineupIndexCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SwitchLineupIndexCsReq, payloadMsg)
	req := msg.(*proto.SwitchLineupIndexCsReq)

	g.Player.DbLineUp.MainLineUp = req.Index
	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SwitchLineupIndexScRsp{Index: req.Index}

	g.Send(cmd.SwitchLineupIndexScRsp, rsp)
}

func (g *Game) HandleSwapLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SwapLineupCsReq, payloadMsg)
	req := msg.(*proto.SwapLineupCsReq)

	// 交换角色
	g.SwapLineup(req.Index, req.SrcSlot, req.DstSlot)

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SwapLineupCsReq{}

	g.Send(cmd.SwapLineupScRsp, rsp)
}

func (g *Game) SetLineupNameCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetLineupNameCsReq, payloadMsg)
	req := msg.(*proto.SetLineupNameCsReq)
	g.Player.DbLineUp.LineUpList[req.Index].Name = req.Name

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SetLineupNameScRsp{
		Index: req.Index,
		Name:  req.Name,
	}

	g.Send(cmd.SetLineupNameScRsp, rsp)
}

func (g *Game) ReplaceLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ReplaceLineupCsReq, payloadMsg)
	req := msg.(*proto.ReplaceLineupCsReq)
	g.Player.DbLineUp.LineUpList[req.Index].AvatarIdList = []uint32{0, 0, 0, 0}
	for _, avatarid := range req.Slots {
		g.Player.DbLineUp.LineUpList[req.Index].AvatarIdList[avatarid.Slot] = avatarid.Id
	}

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.ReplaceLineupScRsp, rsp)
}

func (g *Game) ChangeLineupLeaderCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ChangeLineupLeaderCsReq, payloadMsg)
	req := msg.(*proto.ChangeLineupLeaderCsReq)

	rsp := &proto.ChangeLineupLeaderScRsp{Slot: req.Slot}

	g.Send(cmd.ChangeLineupLeaderScRsp, rsp)
}

func (g *Game) QuitLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.QuitLineupCsReq, payloadMsg)
	req := msg.(*proto.QuitLineupCsReq)

	for id, avatarId := range g.Player.DbLineUp.LineUpList[req.Index].AvatarIdList {
		if avatarId == req.BaseAvatarId {
			g.Player.DbLineUp.LineUpList[req.Index].AvatarIdList[id] = 0
		}
	}

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.QuitLineupScRsp, rsp)
}
