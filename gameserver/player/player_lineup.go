package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

// 队伍更新通知
func (g *GamePlayer) SyncLineupNotify(db *spb.Line) {
	rsq := &proto.SyncLineupNotify{
		ReasonList: make([]proto.SyncLineupReason, 0),
		Lineup:     g.GetPd().GetLineUpPb(db),
	}
	g.Send(cmd.SyncLineupNotify, rsq)
}

// 场景更新目标队伍
func (g *GamePlayer) SceneGroupRefreshScNotify(line *spb.Line) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: g.GetPd().GetSceneGroupRefreshInfoByLineUP(
			line, g.GetPd().GetPosPb(), g.GetPd().GetRotPb()),
	}
	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

// 获取全部队伍
func HandleGetAllLineupDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetAllLineupDataScRsp)
	rsp.LineupList = make([]*proto.LineupInfo, 0)
	db := g.GetPd().GetLineUp()
	rsp.CurIndex = db.MainLineUp

	// 添加普通队伍
	for i := 0; i < model.MaxLineupList; i++ {
		lineupList := g.GetPd().GetLineUpPb(g.GetPd().GetLineUpById(uint32(i)))
		rsp.LineupList = append(rsp.LineupList, lineupList)
	}

	if changeStory := g.GetPd().GetCurChangeStoryInfo(); changeStory != nil {
		lineupList := g.GetPd().GetLineUpPb(g.GetPd().GetStoryLineById(changeStory.ChangeStoryId))
		rsp.LineupList = append(rsp.LineupList, lineupList)
	}

	g.Send(cmd.GetAllLineupDataScRsp, rsp)
}

// 获取当前队伍
func HandleGetCurLineupDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetCurLineupDataScRsp)
	rsp.Lineup = g.GetPd().GetLineUpPb(g.GetPd().GetCurLineUp())

	g.Send(cmd.GetCurLineupDataScRsp, rsp)
}

// 获取当前队伍角色信息
func GetLineupAvatarDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetLineupAvatarDataScRsp)
	rsp.AvatarDataList = g.GetPd().GetLineupAvatarDataList(g.GetPd().GetCurLineUp())

	g.Send(cmd.GetLineupAvatarDataScRsp, rsp)
}

// 更新队伍中的指定角色
func HandleJoinLineupCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.JoinLineupCsReq)

	g.GetPd().UnDbLineUp(req.Index, req.Slot, req.BaseAvatarId)

	// 队伍更新通知
	g.SyncLineupNotify(g.GetPd().GetLineUpById(req.Index))
	g.SceneGroupRefreshScNotify(g.GetPd().GetLineUpById(req.Index))

	rsp := new(proto.LineupAvatar)
	g.Send(cmd.JoinLineupScRsp, rsp)
}

func HandleSwitchLineupIndexCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SwitchLineupIndexCsReq)

	db := g.GetPd().GetLineUp()
	db.MainLineUp = req.Index
	// 队伍更新通知
	g.SyncLineupNotify(g.GetPd().GetCurLineUp())
	g.SceneGroupRefreshScNotify(g.GetPd().GetCurLineUp())

	rsp := &proto.SwitchLineupIndexScRsp{Index: req.Index}

	g.Send(cmd.SwitchLineupIndexScRsp, rsp)
}

// 2.5.0 遗弃
// func HandleSwapLineupCsReq(g *GamePlayer, payloadMsg pb.Message) {
// 	req := payloadMsg.(*proto.SwapLineupCsReq)
//
// 	// 交换角色
// 	g.GetPd().SwapLineup(req.Index, req.SrcSlot, req.DstSlot)
//
// 	// 队伍更新通知
// 	g.SyncLineupNotify(g.GetPd().GetLineUpById(req.Index))
// 	g.SceneGroupRefreshScNotify(req.Index)
//
// 	rsp := &proto.SwapLineupCsReq{}
//
// 	g.Send(cmd.SwapLineupScRsp, rsp)
// }

func SetLineupNameCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetLineupNameCsReq)
	db := g.GetPd().GetLineUpById(req.Index)
	db.Name = req.Name

	// 队伍更新通知
	g.SyncLineupNotify(db)
	g.SceneGroupRefreshScNotify(db)

	rsp := &proto.SetLineupNameScRsp{
		Index: req.Index,
		Name:  req.Name,
	}

	g.Send(cmd.SetLineupNameScRsp, rsp)
}

func ReplaceLineupCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ReplaceLineupCsReq)

	index := req.Index
	isBattleLine := false
	var db *spb.Line
	if req.GameStoryLineId != 0 {
		db = g.GetPd().GetStoryLineById(req.GameStoryLineId)
	} else {
		switch req.ExtraLineupType {
		case proto.ExtraLineupType_LINEUP_NONE:
			db = g.GetPd().GetLineUpById(index)
		case proto.ExtraLineupType_LINEUP_CHALLENGE:
			index = model.Challenge_1
			db = g.GetPd().GetBattleLineUpById(index)
			isBattleLine = true
		case proto.ExtraLineupType_LINEUP_CHALLENGE_2:
			index = model.Challenge_2
			db = g.GetPd().GetBattleLineUpById(index)
			isBattleLine = true
		case proto.ExtraLineupType_LINEUP_ROGUE:
			index = model.Rogue
			db = g.GetPd().GetBattleLineUpById(index)
			isBattleLine = true
		case proto.ExtraLineupType_LINEUP_TOURN_ROGUE:
			index = model.RogueTourn
			db = g.GetPd().GetBattleLineUpById(index)
			isBattleLine = true
		}
	}

	db.LeaderSlot = req.LeaderSlot
	db.LineType = spb.ExtraLineupType(req.ExtraLineupType)
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	for _, avatarList := range req.LineupSlotList {
		db.AvatarIdList[avatarList.Slot] = &spb.LineAvatarList{
			AvatarId:       avatarList.Id,
			Slot:           avatarList.Slot,
			LineAvatarType: spb.AvatarType(avatarList.AvatarType),
		}
	}

	// 队伍更新通知
	g.SyncLineupNotify(db)
	if isBattleLine {
		// 将角色属性拷贝出来
		// for _, avatar := range req.LineupSlotList {
		// 	avatarBin := g.GetPd().GetAvatarBinById(avatar.Id)
		// 	g.GetPd().CopyBattleAvatar(avatarBin)
		// }
	}

	// 是当前队伍的时候就需要场景通知
	if g.GetPd().GetCurLineUp().Index == req.Index {
		g.SceneGroupRefreshScNotify(db)
	}

	g.Send(cmd.ReplaceLineupScRsp, &proto.ReplaceLineupScRsp{})
}

func ChangeLineupLeaderCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ChangeLineupLeaderCsReq)

	rsp := &proto.ChangeLineupLeaderScRsp{Slot: req.Slot}

	db := g.GetPd().GetCurLineUp()
	db.LeaderSlot = req.Slot

	g.Send(cmd.ChangeLineupLeaderScRsp, rsp)
}

func TriggerVoiceCsReq(g *GamePlayer, payloadMsg pb.Message) {
	g.Send(cmd.TriggerVoiceScRsp, new(proto.TriggerVoiceScRsp))
}

func QuitLineupCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.QuitLineupCsReq)
	db := g.GetPd().GetCurLineUp()

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
	g.SceneGroupRefreshScNotify(db)

	g.Send(cmd.QuitLineupScRsp, nil)
}

func (g *GamePlayer) SetBattleLineUp(index uint32, avatarList []uint32) {
	if avatarList == nil { // 没有传入角色
		avatarList = make([]uint32, 0)
		for _, info := range g.GetPd().GetCurLineUp().AvatarIdList {
			avatarList = append(avatarList, info.AvatarId)
		}
	}
	var lineUpType spb.ExtraLineupType
	switch index {
	case model.Challenge_1:
		lineUpType = spb.ExtraLineupType_LINEUP_CHALLENGE
	case model.Challenge_2:
		lineUpType = spb.ExtraLineupType_LINEUP_CHALLENGE_2
	case model.Rogue:
		lineUpType = spb.ExtraLineupType_LINEUP_ROGUE
	case model.RogueTourn:
		lineUpType = spb.ExtraLineupType_LINEUP_TOURN_ROGUE
	case model.Activity:
		lineUpType = spb.ExtraLineupType_LINEUP_STAGE_TRIAL
	default:
		logger.Warn("未知的队伍类型:%v", index)
		return
	}
	db := g.GetPd().GetBattleLineUpById(index)
	db.LeaderSlot = 0
	db.LineType = lineUpType
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	var id uint32 = 0
	for _, avatarId := range avatarList {
		if ok, avatarType := g.GetPd().SpecialMainAvatar(avatarId); ok {
			db.AvatarIdList[id] = &spb.LineAvatarList{AvatarId: avatarId, Slot: id, LineAvatarType: avatarType}
			id++
		}
	}
	// 拷贝角色
	// for _, avatar := range avatarList {
	// 	avatarBin := g.GetPd().GetAvatarBinById(avatar)
	// 	g.GetPd().CopyBattleAvatar(avatarBin)
	// }
	db.Mp = 5
	g.SyncLineupNotify(db)
}
