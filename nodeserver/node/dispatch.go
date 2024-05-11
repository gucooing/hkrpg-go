package node

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) dispatchRecvHandle() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! DISPATCH SERVICE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.n.killService(s)
			return
		}
	}()

	for {
		bin, err := s.Conn.Read()
		if err != nil {
			s.n.killService(s)
			break
		}
		msgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &msgList, nil)
		for _, msg := range msgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.dispatchRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) dispatchRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.GetAllServiceGateReq: // 心跳
		s.dispatchGetAllServiceGateReq(serviceMsg)
	default:
		logger.Info("dispatch -> node error cmdid:%v", cmdId)
	}
}

func (s *Service) dispatchGetAllServiceGateReq(serviceMsg pb.Message) {
	s.lastAliveTime = time.Now().Unix()
	req := serviceMsg.(*spb.GetAllServiceGateReq)
	if req.ServiceType != s.ServerType {
		logger.Debug("Service registration failed")
		s.n.killService(s)
		return
	}
	rsp := &spb.GetAllServiceGateRsp{
		GateServiceList: make([]*spb.ServiceAll, 0),
		DispatchTime:    req.DispatchTime,
		NodeTime:        time.Now().UnixNano() / 1e6,
	}
	for _, service := range NODE.MapService[spb.ServerType_SERVICE_GATE] {
		serviceAll := &spb.ServiceAll{
			ServiceType: service.ServerType,
			Addr:        service.Addr,
			Port:        service.Port,
			PlayerNum:   service.PlayerNum,
			AppId:       service.AppId,
		}
		rsp.GateServiceList = append(rsp.GateServiceList, serviceAll)
	}
	s.sendHandle(cmd.GetAllServiceGateRsp, rsp)
}
