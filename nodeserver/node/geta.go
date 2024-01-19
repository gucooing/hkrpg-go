package node

import (
	"github.com/gucooing/hkrpg-go/nodeserver/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) GetGameOuterAddrReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetGameOuterAddrReq)
	if req.AppId != s.AppId {
		logger.Debug("Service registration failed")
		s.killService()
		return
	}
	rsp := &spb.GetGameOuterAddrRsp{
		ServerType: req.ServerType,
		GameAddr:   getMinService(spb.ServerType_SERVICE_GAME),
	}
	s.sendHandle(cmd.GetGameOuterAddrRsp, rsp)
}
