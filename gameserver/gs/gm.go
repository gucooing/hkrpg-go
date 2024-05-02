package gs

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	play := s.getPlayerByUuid(req.Uuid)
	if play == nil {
		return
	}
	play.p.GmGive(serviceMsg)
}

func (s *GameServer) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	play := s.getPlayerByUuid(req.Uuid)
	if play == nil {
		return
	}
	play.p.GmWorldLevel(serviceMsg)
}

func (s *GameServer) DelItem(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.DelItem)
	play := s.getPlayerByUuid(req.Uuid)
	if play == nil {
		return
	}
	play.p.DelItem(serviceMsg)
}

func (s *GameServer) getPlayerByUuid(uuid int64) *GamePlayer {
	playerAll := s.GetAllPlayer()
	if playerAll == nil || playerAll[uuid] == nil {
		logger.Debug("[UUID:%v]找不到该在线玩家", uuid)
		return nil
	} else {
		return playerAll[uuid]
	}
}
