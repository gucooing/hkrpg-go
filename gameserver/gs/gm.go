package gs

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	if req.PlayerUid == 0 || s.PlayerMap[req.Uuid] == nil {
		return
	}
	s.PlayerMap[req.Uuid].GmGive(serviceMsg)
}

func (s *GameServer) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	if req.PlayerUid == 0 || s.PlayerMap[req.Uuid] == nil {
		return
	}
	s.PlayerMap[req.Uuid].GmWorldLevel(serviceMsg)
}

func (s *GameServer) DelItem(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.DelItem)
	if req.PlayerUid == 0 || s.PlayerMap[req.Uuid] == nil {
		return
	}
	s.PlayerMap[req.Uuid].DelItem(serviceMsg)
}
