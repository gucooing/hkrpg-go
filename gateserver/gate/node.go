package gate

import (
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/logger"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 向node注册
func (s *GateServer) Connection() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_GATE,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp,
		Port:       s.Port,
	}

	s.sendNode(cmd.ServiceConnectionReq, req)
}

// 发送到node
func (s *GateServer) sendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdId error")
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := s.nodeConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

// 从node接收消息
func (s *GateServer) recvNode() {
	nodeMsg := make([]byte, PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := s.nodeConn.Read(nodeMsg)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.NodeRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *GateServer) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_GATE && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
	// 获取game地址/心跳包
	go s.GetAllServiceReq()
}

func (s *GateServer) GetAllServiceReq() {
	// 心跳包
	for {
		req := &spb.GetAllServiceReq{
			ServiceType:     spb.ServerType_SERVICE_GATE,
			GetServiceType_: spb.ServerType_SERVICE_GAME,
		}
		s.sendNode(cmd.GetAllServiceReq, req)
		time.Sleep(time.Second * 5)
	}
}

func (s *GateServer) GetAllServiceRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetAllServiceRsp)
	if rsp.ServiceType != spb.ServerType_SERVICE_GATE {
		return
	}
	gameAll := make(map[string]*serviceGame, 0)
	var minGameAppId string
	var minGameNum uint64 = 0
	for _, service := range rsp.ServiceList {
		if service.Addr == "" || service.AppId == "" || service.ServiceType != spb.ServerType_SERVICE_GAME {
			return
		}
		if minGameAppId == "" {
			minGameAppId = service.AppId
			minGameNum = service.PlayerNum
		} else {
			if minGameNum > service.PlayerNum {
				minGameAppId = service.AppId
				minGameNum = service.PlayerNum
			}
		}
		serviceG := &serviceGame{
			addr:  service.Addr,
			num:   service.PlayerNum,
			appId: service.AppId,
		}
		gameAll[service.AppId] = serviceG
	}
	if len(gameAll) == 0 {
		return
	}
	s.gameAll = gameAll
	s.gameAppId = minGameAppId
	s.errGameAppId = make([]string, 0)
}
