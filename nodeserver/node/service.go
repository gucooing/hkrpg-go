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

// 重复登录后处理结果处理
func repeatLogin(uid uint32) {
	if player := NODE.PlayerMap[uid]; player != nil {
		if status := player.PlayerStatus; status != nil {
			if status.GateStatus == spb.PlayerGateStatus_PlayerGateStatus_GateLogout && status.GameStatus == spb.PlayerGameStatus_PlayerGameStatus_GameLogout {
				if gate := GetPlayerGate(uid); gate != nil {
					status.Status = spb.PlayerStatus_PlayerStatus_LoggingIn
					gate.PlayerNum++
					gate.sendHandle(cmd.PlayerLoginRsp, &spb.PlayerLoginRsp{PlayerUid: uid})
				}
			}
		}
	}
}

func (s *Service) SyncPlayerOnlineDataNotify(serviceMsg pb.Message) {
	reqn := serviceMsg.(*spb.SyncPlayerOnlineDataNotify)
	rspn := new(spb.SyncPlayerOnlineDataNotify)
	if reqn.PlayerUid == 0 || NODE.PlayerMap[reqn.PlayerUid] == nil {
		return
	}
	rspn.PlayerUid = reqn.PlayerUid
	if reqn.PlayerOnlineData == nil {
		rspn.PlayerOnlineData = NODE.PlayerMap[reqn.PlayerUid].PlayerOnlineData
	} else {
		NODE.PlayerMap[reqn.PlayerUid].PlayerOnlineData = reqn.PlayerOnlineData
	}

	s.sendHandle(cmd.SyncPlayerOnlineDataNotify, rspn)
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
