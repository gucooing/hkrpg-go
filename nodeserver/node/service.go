package node

import (
	"github.com/gucooing/hkrpg-go/nodeserver/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func getMinService(ServerType spb.ServerType) string {
	var minService string
	var minNum uint64

	if len(NODE.MapService[ServerType]) == 0 {
		return ""
	}

	for _, service := range NODE.MapService[ServerType] {
		if service.PlayerNum == 0 || service.PlayerNum < minNum {
			minService = service.Addr
		}
	}

	return minService
}

func (s *Service) ServiceConnectionReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.ServiceConnectionReq)
	if req.AppId == "" || req.ServerType == 0 {
		logger.Debug("Service registration failed")
		s.killService()
		return
	}
	s.AppId = req.AppId
	s.ServerType = req.ServerType
	s.Addr = req.Addr
	NODE.MapService[s.ServerType][s.AppId] = s

	logger.Info("AppId:%s Service:%s Service registration successful", s.AppId, s.ServerType)

	rsp := &spb.ServiceConnectionRsp{
		ServerType: req.ServerType,
		AppId:      req.AppId,
	}

	s.sendHandle(cmd.ServiceConnectionRsp, rsp)
}
