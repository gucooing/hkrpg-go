package game

import (
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
	s.PlayerMap[req.PlayerUid].KickPlayer()
	delete(s.PlayerMap, req.PlayerUid)
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
