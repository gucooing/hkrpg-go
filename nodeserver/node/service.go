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
			minService = service.AppId
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
	s.Port = req.Port
	NODE.MapService[s.ServerType][s.AppId] = s

	logger.Info("AppId:%s Service:%s Service registration successful", s.AppId, s.ServerType)

	rsp := &spb.ServiceConnectionRsp{
		ServerType: req.ServerType,
		AppId:      req.AppId,
	}

	s.sendHandle(cmd.ServiceConnectionRsp, rsp)
}

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
		serverType = spb.ServerType_SERVICE_GATE
	case spb.ServerType_SERVICE_GATE:
		serverType = spb.ServerType_SERVICE_GAME
	}

	rsp := &spb.GetServerOuterAddrRsp{
		ServerType: req.ServerType,
	}

	appId := getMinService(serverType)

	if appId == "" {
	} else {
		rsp.Addr = NODE.MapService[serverType][appId].Addr
		rsp.Port = NODE.MapService[serverType][appId].Port
	}

	s.sendHandle(cmd.GetServerOuterAddrRsp, rsp)
}
