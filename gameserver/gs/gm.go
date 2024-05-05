package gs

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	play := s.GetPlayerByUuid(req.Uuid)
	if play == nil {
		return
	}
	play.p.GmGive(serviceMsg)
}

func (s *GameServer) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	play := s.GetPlayerByUuid(req.Uuid)
	if play == nil {
		return
	}
	play.p.GmWorldLevel(serviceMsg)
}

func (s *GameServer) DelItem(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.DelItem)
	play := s.GetPlayerByUuid(req.Uuid)
	if play == nil {
		return
	}
	play.p.DelItem(serviceMsg)
}
