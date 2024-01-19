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
		ServerType: spb.ServerType_SERVICE_GETA,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp + ":" + s.Port,
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
		logger.Error("cmdid error")
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
	if rsp.ServerType == spb.ServerType_SERVICE_GETA && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
	// 获取game地址
	go s.GetGameOuterAddrReq()
}

func (s *GateServer) GetGameOuterAddrReq() {
	for {
		req := &spb.GetGameOuterAddrReq{
			ServerType: spb.ServerType_SERVICE_GETA,
			AppId:      s.AppId,
		}
		s.sendNode(cmd.GetGameOuterAddrReq, req)
		time.Sleep(time.Microsecond * 5)
	}
}

func (s *GateServer) GetGameOuterAddrRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetGameOuterAddrRsp)
	if rsp.ServerType != spb.ServerType_SERVICE_GETA {
		return
	}
	s.gameAddr = rsp.GameAddr
}
