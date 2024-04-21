package gate

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

func (s *GateServer) ServiceStart() {
	go func() {
		for {
			select {
			case msg := <-s.RecvCh:
				s.nodeRegisterMessage(msg.cmdId, msg.serviceMsg)
			case <-s.Ticker.C:
				s.gateGetAllServiceGameReq()
			case <-s.Stop:
				s.Ticker.Stop()
				fmt.Println("Player goroutine stopped")
				return
			}
		}
	}()
}

func (s *GateServer) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg) // 注册包
	case cmd.GetAllServiceGameRsp:
		s.GetAllServiceGameRsp(serviceMsg) // 心跳包
	default:
		logger.Info("nodeRegister error cmdid:%v", cmdId)
	}
}

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

func (s *GateServer) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_GATE && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (s *GateServer) gateGetAllServiceGameReq() {
	// 心跳包
	req := &spb.GetAllServiceGameReq{
		ServiceType: spb.ServerType_SERVICE_GATE,
		GateTime:    time.Now().UnixNano() / 1e6,
		PlayerNum:   uint64(CLIENT_CONN_NUM),
	}
	s.sendNode(cmd.GetAllServiceGameReq, req)
}

func (s *GateServer) GetAllServiceGameRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetAllServiceGameRsp)
	for _, service := range rsp.GameServiceList {
		if service.Addr == "" || service.AppId == 0 || service.ServiceType != spb.ServerType_SERVICE_GAME {
			continue
		}
		if s.getGsByAppid(service.AppId) == nil {
			addr := service.Addr + ":" + service.Port
			s.newGs(addr, service.AppId)
		}
	}

	logger.Debug("gate <--> node ping:%v", (rsp.NodeTime-rsp.GateTime)/2)
}

/******************************************NewLogin***************************************/

func (s *GateServer) PlayerLogoutNotify(uid uint32) {
	notify := &spb.PlayerLogoutNotify{
		Uid: uid,
	}
	s.sendNode(cmd.PlayerLogoutNotify, notify)
}
