package node

import (
	"github.com/gucooing/hkrpg-go/nodeserver/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) GetGateOuterAddrReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetGateOuterAddrReq)
	if req.AppId != s.AppId {
		logger.Debug("Service registration failed")
		s.killService()
		return
	}
	rsp := &spb.GetGateOuterAddrRsp{
		ServerType: req.ServerType,
		GateAddr:   getMinService(spb.ServerType_SERVICE_GETA),
	}
	s.sendHandle(cmd.GetGateOuterAddrRsp, rsp)
}
