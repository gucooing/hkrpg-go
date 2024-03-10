package sdk

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen = 343 * 1024 // 最大应用层包长度
)

type TcpNodeMsg struct {
	cmdId      uint16
	serviceMsg pb.Message
}

func (s *Server) ServiceStart() {
	go func() {
		for {
			select {
			case msg := <-s.RecvCh:
				s.nodeRegisterMessage(msg.cmdId, msg.serviceMsg)
			case <-s.Ticker.C:
				s.getAllServiceGateReq()
			case <-s.Stop:
				s.Ticker.Stop()
				fmt.Println("Player goroutine stopped")
				return
			}
		}
	}()
}

func (s *Server) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp: // 注册回包
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.GetAllServiceGateRsp: // 心跳包
		s.GetAllServiceGateRsp(serviceMsg)
	default:
		logger.Info("nodeRegister error cmdid:%v", cmdId)
	}
}

// 从node接收消息
func (s *Server) RecvNode() {
	nodeMsg := make([]byte, PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := s.NodeConn.Read(nodeMsg)
		if err != nil {
			log.Println("node error")
			os.Exit(0)
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			newServiceMsg := new(TcpNodeMsg)
			newServiceMsg.cmdId = msg.CmdId
			newServiceMsg.serviceMsg = serviceMsg
			s.RecvCh <- newServiceMsg
		}
	}
}

// 发送到node
func (s *Server) sendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdId error")
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := s.NodeConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

// 向node注册
func (s *Server) Connection() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_DISPATCH,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp,
		Port:       s.Port,
	}

	s.sendNode(cmd.ServiceConnectionReq, req)
}

func (s *Server) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_DISPATCH && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (s *Server) getAllServiceGateReq() {
	// 心跳包
	req := &spb.GetAllServiceGateReq{
		ServiceType:  spb.ServerType_SERVICE_DISPATCH,
		DispatchTime: time.Now().UnixNano() / 1e6,
	}
	s.sendNode(cmd.GetAllServiceGateReq, req)
}

func (s *Server) GetAllServiceGateRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetAllServiceGateRsp)
	var minPlayerNum uint64 = 1<<63 - 1 // 初始化为较大的数
	var minPlayerService *spb.ServiceAll
	for _, service := range rsp.GateServiceList {
		if service.PlayerNum < minPlayerNum {
			minPlayerNum = service.PlayerNum
		}
	}
	for _, service := range rsp.GateServiceList {
		if service.PlayerNum == minPlayerNum {
			minPlayerService = service
			break
		}
	}

	if minPlayerService == nil {
		return
	}

	s.GateAddr = minPlayerService.Addr
	s.GatePort = minPlayerService.Port

	logger.Info("dispatch <--> node ping:%v | min gate:%s:%s", (rsp.NodeTime-rsp.DispatchTime)/2, s.GateAddr, s.GatePort)
}
