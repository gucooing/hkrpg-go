package node

import (
	"bufio"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) dispatchRecvHandle() {
	payload := make([]byte, PacketMaxLen)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! DISPATCH SERVICE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.killService()
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := bufio.NewReader(s.Conn).Read(payload)
		if err != nil {
			s.killService()
			break
		}
		bin = payload[:recvLen]
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
	req := serviceMsg.(*spb.GetAllServiceGateReq)
	if req.ServiceType != s.ServerType {
		logger.Debug("Service registration failed")
		s.killService()
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
