package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 队伍更新通知
func (g *Game) SyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	lineUp := g.GetLineUpById(index)
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
		avatar := g.PlayerPb.Avatar.Avatar[avatarId]
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
	pos := g.GetPos()
	rot := g.GetRot()
	for _, lineup := range g.PlayerPb.LineUp.LineUpList[g.PlayerPb.LineUp.MainLineUp].AvatarIdList {
		if lineup == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
			UpdateType: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
						AvatarType:   proto.AvatarType(g.PlayerPb.Avatar.Avatar[lineup].AvatarType),
						BaseAvatarId: lineup,
					}},
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
			},
		}
		g.Player.EntityList[entityId] = &EntityList{
			Entity:  lineup,
			GroupId: 0,
		}
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	}
	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *Game) HandleGetAllLineupDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAllLineupDataScRsp)
	rsp.LineupList = make([]*proto.LineupInfo, 0)
	rsp.CurIndex = 0

	for i := 0; i < 6; i++ {
		lineUp := g.GetLineUpById(uint32(i))
		lineupList := &proto.LineupInfo{
			IsVirtual:       false,
			LeaderSlot:      0,
			AvatarList:      make([]*proto.LineupAvatar, 0),
			Index:           uint32(i),
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
			avatar := g.PlayerPb.Avatar.Avatar[avatarId]
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
		rsp.LineupList = append(rsp.LineupList, lineupList)
	}

	g.Send(cmd.GetAllLineupDataScRsp, rsp)
}

func (g *Game) HandleGetCurLineupDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurLineupDataScRsp)
	lineUp := g.GetLineUpById(g.PlayerPb.LineUp.MainLineUp)
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           g.PlayerPb.LineUp.MainLineUp,
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
		avatar := g.PlayerPb.Avatar.Avatar[avatarId]
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
	rsp.Lineup = lineupList

	g.Send(cmd.GetCurLineupDataScRsp, rsp)
}

func (g *Game) HandleJoinLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.JoinLineupCsReq, payloadMsg)
	req := msg.(*proto.JoinLineupCsReq)

	g.UnDbLineUp(req.Index, req.Slot, req.BaseAvatarId)
	g.PlayerPb.LineUp.MainAvatarId = 0

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.LineupAvatar)
	g.Send(cmd.JoinLineupScRsp, rsp)
}

func (g *Game) HandleSwitchLineupIndexCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SwitchLineupIndexCsReq, payloadMsg)
	req := msg.(*proto.SwitchLineupIndexCsReq)

	g.PlayerPb.LineUp.MainLineUp = req.Index
	g.PlayerPb.LineUp.MainAvatarId = 0
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
	g.PlayerPb.LineUp.MainAvatarId = 0

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SwapLineupCsReq{}

	g.Send(cmd.SwapLineupScRsp, rsp)
}

func (g *Game) SetLineupNameCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetLineupNameCsReq, payloadMsg)
	req := msg.(*proto.SetLineupNameCsReq)
	g.PlayerPb.LineUp.LineUpList[req.Index].Name = req.Name

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
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因

	switch req.ExtraLineupType {
	case proto.ExtraLineupType_LINEUP_CHALLENGE:
		g.NewChallengeLineUp(req)
		g.Send(cmd.ReplaceLineupScRsp, rsp)
		return
	case proto.ExtraLineupType_LINEUP_CHALLENGE_2:
		g.NewChallengeLineUp(req)
		g.Send(cmd.ReplaceLineupScRsp, rsp)
		return
	}

	g.PlayerPb.LineUp.LineUpList[req.Index].AvatarIdList = []uint32{0, 0, 0, 0}
	for _, avatarid := range req.Slots {
		g.PlayerPb.LineUp.LineUpList[req.Index].AvatarIdList[avatarid.Slot] = avatarid.Id
	}

	g.PlayerPb.LineUp.MainAvatarId = 0

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	g.Send(cmd.ReplaceLineupScRsp, rsp)
}

func (g *Game) ChangeLineupLeaderCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ChangeLineupLeaderCsReq, payloadMsg)
	req := msg.(*proto.ChangeLineupLeaderCsReq)

	rsp := &proto.ChangeLineupLeaderScRsp{Slot: req.Slot}

	g.PlayerPb.LineUp.MainAvatarId = req.Slot

	g.Send(cmd.ChangeLineupLeaderScRsp, rsp)
}

func (g *Game) QuitLineupCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.QuitLineupCsReq, payloadMsg)
	req := msg.(*proto.QuitLineupCsReq)

	for id, avatarId := range g.PlayerPb.LineUp.LineUpList[req.Index].AvatarIdList {
		if avatarId == req.BaseAvatarId {
			g.PlayerPb.LineUp.LineUpList[req.Index].AvatarIdList[id] = 0
		}
	}

	g.PlayerPb.LineUp.MainAvatarId = 0
	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.QuitLineupScRsp, rsp)
}
