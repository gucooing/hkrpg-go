package node

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	if req.PlayerUid == 0 || NODE.PlayerMap[req.PlayerUid] == nil {
		return
	}
	GetPlayerGame(req.PlayerUid).sendHandle(cmd.GmGive, serviceMsg)
}

func (s *Service) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	if req.PlayerUid == 0 || NODE.PlayerMap[req.PlayerUid] == nil {
		return
	}
	GetPlayerGame(req.PlayerUid).sendHandle(cmd.GmWorldLevel, serviceMsg)
}
