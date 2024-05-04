package node

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) gateRecvHandle() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE SERVICE MAIN LOOP PANIC !!!")
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
			s.gateRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) gateRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.GetAllServiceGameReq: // 心跳包
		s.gateGetAllServiceGameReq(serviceMsg)
	default:
		logger.Info("gateRegister error cmdid:%v", cmdId)
	}
}

func (s *Service) gateGetAllServiceGameReq(serviceMsg pb.Message) {
	s.lastAliveTime = time.Now().Unix()
	req := serviceMsg.(*spb.GetAllServiceGameReq)
	if req.ServiceType != s.ServerType {
		logger.Debug("Service registration failed")
		s.n.killService(s)
		return
	}
	s.PlayerNum = req.PlayerNum
	rsp := &spb.GetAllServiceGameRsp{
		GameServiceList: make([]*spb.ServiceAll, 0),
		GateTime:        req.GateTime,
		NodeTime:        time.Now().UnixNano() / 1e6,
	}
	for _, service := range s.n.GetAllServiceByType(spb.ServerType_SERVICE_GAME) {
		serviceAll := &spb.ServiceAll{
			ServiceType: service.ServerType,
			Addr:        service.Addr,
			Port:        service.Port,
			PlayerNum:   service.PlayerNum,
			AppId:       service.AppId,
		}
		rsp.GameServiceList = append(rsp.GameServiceList, serviceAll)
	}
	s.sendHandle(cmd.GetAllServiceGameRsp, rsp)
}
