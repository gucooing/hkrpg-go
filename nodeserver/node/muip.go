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

func (s *Service) muipRecvHandle() {
	payload := make([]byte, PacketMaxLen)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! DISPATCH SERVICE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.n.killService(s)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := bufio.NewReader(s.Conn).Read(payload)
		if err != nil {
			s.n.killService(s)
			break
		}
		bin = payload[:recvLen]
		msgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &msgList, nil)
		for _, msg := range msgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.muipRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) muipRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.MuipToNodePingReq:
		s.MuipToNodePingReq(serviceMsg)
	default:
		logger.Info("muip -> node error cmdid:%v", cmdId)
	}
}

func (s *Service) MuipToNodePingReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.MuipToNodePingReq)
	rsp := &spb.MuipToNodePingRsp{
		MuipServerTime: req.MuipServerTime,
		NodeServerTime: time.Now().UnixNano() / 1e6,
		ServiceList:    make(map[uint32]*spb.MuipServiceAll),
	}

	for serverType, serviceList := range s.n.GetAllService() {
		muipServiceAll := &spb.MuipServiceAll{
			ServiceList: make([]*spb.ServiceAll, 0),
		}
		for _, service := range serviceList {
			muipServiceAll.ServiceList = append(muipServiceAll.ServiceList, &spb.ServiceAll{
				ServiceType: service.ServerType,
				Addr:        service.Addr,
				PlayerNum:   service.PlayerNum,
				AppId:       service.AppId,
				Port:        service.Port,
			})
		}
		rsp.ServiceList[serverType] = muipServiceAll
	}

	s.sendHandle(cmd.MuipToNodePingRsp, rsp)
}
