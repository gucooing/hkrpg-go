package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) SetDisplayAvatarCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetDisplayAvatarCsReq)
	db := g.GetPd().GetFriend()
	if db.DisplayAvatarList == nil {
		db.DisplayAvatarList = make(map[uint32]uint32)
	}
	for _, display := range req.DisplayAvatarList {
		db.DisplayAvatarList[display.Pos] = display.AvatarId
	}
	rsp := &proto.SetDisplayAvatarScRsp{
		DisplayAvatarList: req.DisplayAvatarList,
	}
	g.Send(cmd.SetDisplayAvatarScRsp, rsp)
}

func (g *GamePlayer) SetAssistAvatarCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetAssistAvatarCsReq)
	db := g.GetPd().GetFriend()
	if db.AssistAvatarList == nil {
		db.AssistAvatarList = make(map[uint32]uint32)
	}
	for pos, avatarId := range req.AvatarIdList {
		db.AssistAvatarList[uint32(pos)] = avatarId
	}
	rsp := &proto.SetAssistAvatarScRsp{
		AvatarId:     req.AvatarId,
		AvatarIdList: req.AvatarIdList,
	}
	g.Send(cmd.SetAssistAvatarScRsp, rsp)
}

func (g *GamePlayer) HandleGetFriendLoginInfoCsReq(payloadMsg pb.Message) {
	db := g.GetPd().GetFriendList()
	rsp := &proto.GetFriendLoginInfoScRsp{
		FriendUidList: make([]uint32, 0),
		Retcode:       0,
	}
	for uid := range db {
		rsp.FriendUidList = append(rsp.FriendUidList, uid)
	}
	g.Send(cmd.GetFriendLoginInfoScRsp, rsp)
}

func (g *GamePlayer) GetFriendListInfoCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetFriendListInfoScRsp)
	rsp.FriendList = make([]*proto.FriendSimpleInfo, 0)
	for uid := range g.GetPd().GetFriendList() {
		rsp.FriendList = append(rsp.FriendList, g.GetPd().GetFriendSimpleInfo(uid))
	}
	g.Send(cmd.GetFriendListInfoScRsp, rsp)
}

func (g *GamePlayer) GetPlayerDetailInfoCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetPlayerDetailInfoCsReq)
	g.Send(cmd.GetPlayerDetailInfoScRsp, &proto.GetPlayerDetailInfoScRsp{
		DetailInfo: g.GetPd().GetPlayerDetailInfo(req.Uid),
		Retcode:    0,
	})
}

func (g *GamePlayer) GetFriendApplyListInfoCsReq(payloadMsg pb.Message) {
	receiveApplyList := g.GetPd().GetRecvApplyFriend()
	rsp := &proto.GetFriendApplyListInfoScRsp{
		SendApplyList:    make([]uint32, 0),
		Retcode:          0,
		ReceiveApplyList: make([]*proto.FriendApplyInfo, 0),
	}
	for _, receiveApply := range receiveApplyList {
		bin := g.GetPd().GetFriendApplyInfo(receiveApply)
		if bin == nil {
			continue
		}
		rsp.SendApplyList = append(rsp.SendApplyList, receiveApply.ApplyUid)
		rsp.ReceiveApplyList = append(rsp.ReceiveApplyList, bin)
	}
	g.Send(cmd.GetFriendApplyListInfoScRsp, rsp)
}

func (g *GamePlayer) GetChatFriendHistoryCsReq(payloadMsg pb.Message) {
	g.Send(cmd.GetChatFriendHistoryScRsp, &proto.GetChatFriendHistoryScRsp{
		FriendHistoryInfo: []*proto.FriendHistoryInfo{
			{
				LastSendTime: time.Now().Unix(),
				ContactId:    0,
			},
		},
		Retcode: 0,
	})
}

func (g *GamePlayer) SearchPlayerCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SearchPlayerCsReq)
	rsp := &proto.SearchPlayerScRsp{
		Retcode:        0,
		ResultUidList:  make([]uint32, 0),
		SimpleInfoList: make([]*proto.PlayerSimpleInfo, 0),
	}
	for _, uid := range req.UidList {
		bin := model.GetPlayerSimpleInfo(uid)
		if bin == nil {
			continue
		}
		rsp.SimpleInfoList = append(rsp.SimpleInfoList, bin)
		rsp.ResultUidList = append(rsp.ResultUidList, uid)
	}
	g.Send(cmd.SearchPlayerScRsp, rsp)
}

func (g *GamePlayer) HandleFriendCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.HandleFriendCsReq)
	if req.IsAccept {
		g.GetPd().AddFriend(req.Uid)
	}
	rsp := &proto.HandleFriendScRsp{
		IsAccept:   req.IsAccept,
		Uid:        req.Uid,
		Retcode:    0,
		FriendInfo: g.GetPd().GetFriendSimpleInfo(req.Uid),
	}
	g.Send(cmd.HandleFriendScRsp, rsp)
}
