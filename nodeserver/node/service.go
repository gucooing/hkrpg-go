package node

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
		if minService == "" {
			minService = service.AppId
		}
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
	case spb.ServerType_SERVICE_DISPATCH:
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
		rsp.AppId = appId
	}

	s.sendHandle(cmd.GetServerOuterAddrRsp, rsp)
}

func (s *Service) GetAllServiceReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetAllServiceReq)
	if req.ServiceType == spb.ServerType_SERVICE_NONE {
		return
	}
	switch req.ServiceType {
	case spb.ServerType_SERVICE_MUIP:
		s.MuipGetAllServiceReq()
		return
	}
	rsp := &spb.GetAllServiceRsp{
		ServiceType: req.ServiceType,
		ServiceList: make([]*spb.ServiceAll, 0),
	}
	for _, service := range NODE.MapService[req.GetServiceType_] {
		serviceList := &spb.ServiceAll{
			ServiceType: service.ServerType,
			Addr:        service.Addr + ":" + service.Port,
			PlayerNum:   service.PlayerNum,
			AppId:       service.AppId,
		}
		rsp.ServiceList = append(rsp.ServiceList, serviceList)
	}

	s.sendHandle(cmd.GetAllServiceRsp, rsp)
}

func (s *Service) MuipGetAllServiceReq() {
	rsp := &spb.GetAllServiceRsp{
		ServiceType: spb.ServerType_SERVICE_MUIP,
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

// 注意！只有gate能发登录通知包
func (s *Service) PlayerLoginReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLoginReq)
	rsp := new(spb.PlayerLoginRsp)
	if req.PlayerUid == 0 {
		s.sendHandle(cmd.PlayerLoginRsp, rsp)
		return
	}
	if NODE.PlayerMap[req.PlayerUid] == nil {
		rsp.PlayerUid = req.PlayerUid
	} else {
		// 发送重复登录下线通知
		if s.ServerType == spb.ServerType_SERVICE_GATE {
			// 旧game服务玩家数减少
			GetPlayerGame(req.PlayerUid).PlayerNum--
			// 通知旧gate玩家下线
			notify := &spb.PlayerLogoutNotify{
				PlayerUid: req.PlayerUid,
			}
			GetPlayerGate(req.PlayerUid).sendHandle(cmd.PlayerLogoutNotify, notify)
		}
	}
	// 添加在线玩家列表到map/game
	NODE.PlayerMap[req.PlayerUid] = &PlayerService{
		GateAppId: s.AppId,
		GameAppId: req.AppId,
	}
	// 目标game添加玩家数
	GetPlayerGame(req.PlayerUid).PlayerNum++
	// 目标gate添加玩家数
	s.PlayerNum++
	logger.Info("[UID:%v]登录目标GameServer:%v", req.PlayerUid, req.AppId)

	s.sendHandle(cmd.PlayerLoginRsp, rsp)
}

// 注意！只有gate能发离线通知包
func (s *Service) PlayerLogoutReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutReq)
	if req.PlayerUid == 0 || NODE.PlayerMap[req.PlayerUid] == nil {
		return
	}
	logger.Info("[UID:%v]离线目标GameServer:%v", req.PlayerUid, NODE.PlayerMap[req.PlayerUid].GameAppId)
	// 通知game玩家离线
	GetPlayerGame(req.PlayerUid).sendHandle(cmd.PlayerLogoutReq, serviceMsg)
	// 减少gate人数
	s.PlayerNum--
	// 减少game人数
	GetPlayerGame(req.PlayerUid).PlayerNum--
	// 删除玩家
	delete(NODE.PlayerMap, req.PlayerUid)
}

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
