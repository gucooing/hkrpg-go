package node

import (
	"github.com/gucooing/hkrpg-go/nodeserver/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) GetServerOuterAddrReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetServerOuterAddrReq)
	var serverType spb.ServerType
	if req.AppId != s.AppId {
		logger.Debug("Service registration failed")
		s.killService()
		return
	}
	s.PlayerNum = req.PlayerNum
	switch req.ServerType {
	case spb.ServerType_SERVICE_DISCORD:
		serverType = spb.ServerType_SERVICE_GETA
	case spb.ServerType_SERVICE_GETA:
		serverType = spb.ServerType_SERVICE_GAME
	}
	rsp := &spb.GetServerOuterAddrRsp{
		ServerType: req.ServerType,
		Addr:       getMinService(serverType),
	}
	s.sendHandle(cmd.GetServerOuterAddrRsp, rsp)
}
