package muip

import (
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

func (s *Muip) NodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp: // 注册回包
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.GetAllServiceRsp: // 心跳包
		s.GetAllServiceRsp(serviceMsg)
	default:

	}
}

// 从node接收消息
func (s *Muip) RecvNode() {
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
			s.NodeRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

// 发送到node
func (s *Muip) SendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdId error")
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err = s.NodeConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

// 向node注册
func (s *Muip) Connection() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_MUIP,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp,
		Port:       s.Port,
	}

	s.SendNode(cmd.ServiceConnectionReq, req)
}

func (s *Muip) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_MUIP && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
	// 获取game地址/心跳包
	go s.GetAllServiceReq()
}

func (s *Muip) GetAllServiceReq() {
	// 心跳包
	for {
		req := &spb.GetAllServiceReq{
			ServiceType: spb.ServerType_SERVICE_MUIP,
		}
		s.SendNode(cmd.GetAllServiceReq, req)
		time.Sleep(time.Second * 5)
	}
}

func (s *Muip) GetAllServiceRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetAllServiceRsp)
	if rsp.ServiceType != spb.ServerType_SERVICE_MUIP {
		return
	}
	allService := make(map[string][]*AllService)

	for _, service := range rsp.ServiceList {
		allService[service.ServiceType.String()] = append(allService[service.ServiceType.String()], &AllService{
			AppId:     service.AppId,
			PlayerNum: service.PlayerNum,
		})
	}

	s.AllService = allService
}
