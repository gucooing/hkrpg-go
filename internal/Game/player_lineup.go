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
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: proto.AvatarType_AVATAR_FORMAL_TYPE,
			Slot:       uint32(slot),
			Satiety:    0,
			Hp:         10000,
			Id:         avatarId,
			SpBar:      &proto.SpBarInfo{CurSp: 0, MaxSp: 10000},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsq.Lineup = lineupList

	// 更新数据库
	g.UpDataPlayer()

	g.send(cmd.SyncLineupNotify, rsq)
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
			lineupAvatar := &proto.LineupAvatar{
				AvatarType: proto.AvatarType_AVATAR_FORMAL_TYPE,
				Slot:       uint32(slot),
				Satiety:    0,
				Hp:         10000,
				Id:         avatarId,
				SpBar:      &proto.SpBarInfo{CurSp: 0, MaxSp: 10000},
			}
			lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
		}
		rsp.LineupList = append(rsp.LineupList, lineupList)
	}

	g.send(cmd.GetAllLineupDataScRsp, rsp)
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
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: proto.AvatarType_AVATAR_FORMAL_TYPE,
			Slot:       uint32(slot),
			Satiety:    0,
			Hp:         10000,
			Id:         avatarId,
			SpBar:      &proto.SpBarInfo{CurSp: 0, MaxSp: 10000},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsp.Lineup = lineupList

	g.send(cmd.GetCurLineupDataScRsp, rsp)
}

func (g *Game) HandleJoinLineupCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.JoinLineupCsReq, payloadMsg)
	req := msg.(*proto.JoinLineupCsReq)

	g.UnDbLineUp(req.Index, req.Slot, req.BaseAvatarId)

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := new(proto.LineupAvatar)
	g.send(cmd.JoinLineupScRsp, rsp)
}

func (g *Game) HandleSwitchLineupIndexCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SwitchLineupIndexCsReq, payloadMsg)
	req := msg.(*proto.SwitchLineupIndexCsReq)

	g.Player.DbLineUp.MainLineUp = req.Index
	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SwitchLineupIndexScRsp{Index: req.Index}

	g.send(cmd.SwitchLineupIndexScRsp, rsp)
}

func (g *Game) HandleSwapLineupCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SwapLineupCsReq, payloadMsg)
	req := msg.(*proto.SwapLineupCsReq)

	// 交换角色
	g.SwapLineup(req.Index, req.SrcSlot, req.DstSlot)

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SwapLineupCsReq{}

	g.send(cmd.SwapLineupScRsp, rsp)
}

func (g *Game) SetLineupNameCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SetLineupNameCsReq, payloadMsg)
	req := msg.(*proto.SetLineupNameCsReq)
	g.Player.DbLineUp.LineUpList[req.Index].Name = req.Name

	// 队伍更新通知
	g.SyncLineupNotify(req.Index)

	rsp := &proto.SetLineupNameScRsp{
		Index: req.Index,
		Name:  req.Name,
	}

	g.send(cmd.SetLineupNameScRsp, rsp)
}
