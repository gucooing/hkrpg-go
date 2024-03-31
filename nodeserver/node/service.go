package node

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 公共服务注册方法
func (s *Service) ServiceConnectionReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.ServiceConnectionReq)
	if req.AppId == "" || req.ServerType == 0 {
		logger.Debug("Service registration failed")
		s.killService()
		return
	}
	switch req.ServerType {
	case spb.ServerType_SERVICE_GATE:
		go s.gateRecvHandle()
	case spb.ServerType_SERVICE_GAME:
		go s.gameRecvHandle()
	case spb.ServerType_SERVICE_DISPATCH:
		go s.dispatchRecvHandle()
	case spb.ServerType_SERVICE_MUIP:
		go s.muipRecvHandle()
	default:
		logger.Info("Service registration failed")
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

func (s *Service) GetAllServiceReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetAllServiceReq)
	if req.ServiceType == spb.ServerType_SERVICE_NONE {
		return
	}
	rsp := &spb.GetAllServiceRsp{
		ServiceType: req.ServiceType,
		ServiceList: make([]*spb.ServiceAll, 0),
	}
	for _, serviceList := range NODE.MapService {
		for _, service := range serviceList {
			serviceLists := &spb.ServiceAll{
				ServiceType: service.ServerType,
				Addr:        service.Addr + ":" + service.Port,
				PlayerNum:   service.PlayerNum,
				AppId:       service.AppId,
			}
			rsp.ServiceList = append(rsp.ServiceList, serviceLists)
		}
	}

	s.sendHandle(cmd.GetAllServiceRsp, rsp)
}
