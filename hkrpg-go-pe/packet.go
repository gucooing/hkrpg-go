package hkrpg_go_pe

import (
	"strings"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type packetFunc func(h *HkRpgGoServer, p *PlayerGame, protoMsg pb.Message)

var packetCaptureMap = map[uint16]packetFunc{ // 抽包
	cmd.PlayerLogoutCsReq:               PlayerLogoutCsReq,               // 下线请求
	cmd.SendMsgCsReq:                    SendMsgCsReq,                    // 发送聊天消息
	cmd.ApplyFriendCsReq:                ApplyFriendCsReq,                // 发起好友申请
	cmd.GetFriendRecommendListInfoCsReq: GetFriendRecommendListInfoCsReq, // 获取附近的人
}

func (h *HkRpgGoServer) packetCapture(p *PlayerGame, cmdId uint16, protoMsg pb.Message) {
	handelFunc, ok := packetCaptureMap[cmdId]
	if !ok {
		p.sendGameMsg(player.Client, cmdId, protoMsg)
		return
	}
	handelFunc(h, p, protoMsg)
}

func PlayerLogoutCsReq(h *HkRpgGoServer, p *PlayerGame, payloadMsg pb.Message) {
	h.DelPlayer(p.Conn.GetSession().Uid)
}

func SendMsgCsReq(h *HkRpgGoServer, p *PlayerGame, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SendMsgCsReq)

	targetList := req.TargetList
	notify := &proto.RevcMsgScNotify{
		SourceUid:   p.Conn.GetSession().Uid,
		MessageText: req.MessageText,
		ExtraId:     req.ExtraId,
		MessageType: req.MessageType,
		BNABNCCMILM: req.BNABNCCMILM,
		ChatType:    req.ChatType,
	}
	for _, targetUid := range targetList {
		notify.TargetUid = targetUid
		p.toSession(player.Msg{
			CmdId:     cmd.RevcMsgScNotify,
			PlayerMsg: notify,
		})
		if targetUid == 0 {
			bot := &proto.RevcMsgScNotify{
				SourceUid:   0,
				TargetUid:   p.Conn.GetSession().Uid,
				MessageText: p.GamePlayer.EnterCommand(player.Msg{CommandList: strings.Split(req.MessageText, " ")}),
				ExtraId:     req.ExtraId,
				MessageType: req.MessageType,
				BNABNCCMILM: req.BNABNCCMILM,
				ChatType:    req.ChatType,
			}
			p.toSession(player.Msg{
				CmdId:     cmd.RevcMsgScNotify,
				PlayerMsg: bot,
			})
			continue
		}
		target := h.GetPlayer(targetUid)
		if target == nil {
			continue
		}
		target.toSession(player.Msg{
			CmdId:     cmd.RevcMsgScNotify,
			PlayerMsg: notify,
		})
	}

	// if req.MessageText == "dump" {
	// 	p.GamePlayer.Dump()
	// }

	rsp := &proto.SendMsgScRsp{
		EndTime: uint64(time.Now().Unix()),
		Retcode: 0,
	}
	p.toSession(player.Msg{
		CmdId:     cmd.SendMsgScRsp,
		PlayerMsg: rsp,
	})
}

func ApplyFriendCsReq(h *HkRpgGoServer, p *PlayerGame, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ApplyFriendCsReq)
	rsp := &proto.ApplyFriendScRsp{
		Uid:     req.Uid,
		Retcode: 0,
	}
	p.toSession(player.Msg{
		CmdId:     cmd.ApplyFriendScRsp,
		PlayerMsg: rsp,
	})
}

func GetFriendRecommendListInfoCsReq(h *HkRpgGoServer, p *PlayerGame, payloadMsg pb.Message) {
	rsp := &proto.GetFriendRecommendListInfoScRsp{
		PlayerInfoList: make([]*proto.FriendRecommendInfo, 0),
		Retcode:        0,
	}
	for _, s := range h.GetAllPlayer() {
		if s.GamePlayer.Uid == p.Conn.GetSession().Uid {
			continue
		}
		rsp.PlayerInfoList = append(rsp.PlayerInfoList, &proto.FriendRecommendInfo{
			PlayerInfo: model.GetPlayerSimpleInfo(s.GamePlayer.Uid),
		})
	}
	p.toSession(player.Msg{
		CmdId:     cmd.GetFriendRecommendListInfoScRsp,
		PlayerMsg: rsp,
	})
}
