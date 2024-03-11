package game

import (
	"encoding/json"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) NodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.PlayerLogoutReq:
		s.PlayerLogoutReq(serviceMsg) // 玩家离线通知
	case cmd.SyncPlayerOnlineDataNotify:
		s.SyncPlayerOnlineDataNotify(serviceMsg) // 在线数据同步
	case cmd.PlayerLogoutNotify:
		s.PlayerLogoutNotify(serviceMsg) // 异game登录下线通知
	// 下面是gm
	case cmd.GmGive:
		s.GmGive(serviceMsg) // 获取物品
	case cmd.GmWorldLevel:
		s.GmWorldLevel(serviceMsg) // 设置世界等级
	default:

	}
}

func (s *GameServer) PlayerLogoutReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutReq)
	if req.PlayerUid == 0 {
		return
	}
	if pl := s.PlayerMap[req.PlayerUid]; pl != nil {
		KickPlayer(s.PlayerMap[req.PlayerUid])
	}

	s.sendNode(cmd.PlayerLogoutRsp, &spb.PlayerLogoutRsp{PlayerUid: req.PlayerUid})
}

func (s *GameServer) SyncPlayerOnlineDataNotify(serviceMsg pb.Message) {
	noti := serviceMsg.(*spb.SyncPlayerOnlineDataNotify)
	if noti.PlayerUid == 0 || noti.PlayerOnlineData == nil {
		return
	}
	var data *player.PlayerData
	err := json.Unmarshal(noti.PlayerOnlineData, &data)
	if err != nil {
		return
	}
	logger.Info("[UID%v]在线数据同步成功", noti.PlayerUid)
	s.PlayerMap[noti.PlayerUid].Player = data
}

func (s *GameServer) PlayerLogoutNotify(serviceMsg pb.Message) {
	noti := serviceMsg.(*spb.PlayerLogoutNotify)
	if noti.PlayerUid == 0 {
		return
	}
	if s.PlayerMap[noti.PlayerUid] == nil {
		return
	}
	KickPlayer(s.PlayerMap[noti.PlayerUid])
}

func (s *GameServer) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	if req.PlayerUid == 0 || s.PlayerMap[req.PlayerUid] == nil {
		return
	}
	s.PlayerMap[req.PlayerUid].GmGive(serviceMsg)
}

func (s *GameServer) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	if req.PlayerUid == 0 || s.PlayerMap[req.PlayerUid] == nil {
		return
	}
	s.PlayerMap[req.PlayerUid].GmWorldLevel(serviceMsg)
}
