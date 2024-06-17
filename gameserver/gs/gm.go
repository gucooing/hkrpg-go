package gs

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	play := s.getPlayerByUid(req.PlayerUid)
	if play == nil {
		return
	}
	play.p.GmGive(serviceMsg)
}

func (s *GameServer) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	play := s.getPlayerByUid(req.PlayerUid)
	if play == nil {
		return
	}
	play.p.GmWorldLevel(serviceMsg)
}

func (s *GameServer) DelItem(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.DelItem)
	play := s.getPlayerByUid(req.PlayerUid)
	if play == nil {
		return
	}
	play.p.DelItem(serviceMsg)
}

func (s *GameServer) GmMaxCurAvatar(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.MaxCurAvatar)
	play := s.getPlayerByUid(req.PlayerUid)
	if play == nil {
		return
	}
	play.p.GmMaxCurAvatar(serviceMsg)
}

func (s *GameServer) GmMission(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmMission)
	play := s.getPlayerByUid(req.PlayerUid)
	if play == nil {
		return
	}
	play.p.GmMission(req)
}
