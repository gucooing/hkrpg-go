package node

import (
	"bufio"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (n *Node) recvHandle(s *Service) {
	payload := make([]byte, PacketMaxLen)
	var bin []byte = nil
	recvLen, err := bufio.NewReader(s.Conn).Read(payload)
	if err != nil {
		logger.Warn("已切断异常连接Addr:%s", s.Conn.RemoteAddr().String())
		s.Conn.Close()
		return
	}
	bin = payload[:recvLen]
	msgList := make([]*alg.PackMsg, 0)
	alg.DecodeBinToPayload(bin, &msgList, nil)
	for _, msg := range msgList {
		serviceMsg := alg.DecodePayloadToProto(msg)
		if msg.CmdId == cmd.ServiceConnectionReq {
			n.ServiceConnectionReq(serviceMsg, s)
		} else {
			logger.Warn("已切断异常连接Addr:%s", s.Conn.RemoteAddr().String())
			s.Conn.Close()
		}
	}
}

func (s *Service) sendHandle(cmdid uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdid
	rspMsg.PayloadMessage = playerMsg
	serviceMsg := alg.EncodeProtoToPayload(rspMsg)
	if serviceMsg.CmdId == 0 {
		logger.Error("cmdid error")
	}
	binMsg := alg.EncodePayloadToBin(serviceMsg, nil)
	_, err := s.Conn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

func (n *Node) killService(s *Service) {
	s.Conn.Close()
	n.DelMapService(s)
	logger.Info("[%s]服务离线:%s", s.ServerType, s.Conn.RemoteAddr().String())
}

func (n *Node) AddMapService(s *Service) {
	n.serviceMapLock.Lock()
	n.MapService[s.ServerType][s.AppId] = s
	n.serviceMapLock.Unlock()
	logger.Info("AppId:%s Service:%s Service registration successful", alg.GetAppIdStr(s.AppId), s.ServerType)
}

func (n *Node) DelMapService(s *Service) {
	n.serviceMapLock.Lock()
	if n.MapService[s.ServerType] == nil || n.MapService[s.ServerType][s.AppId] == nil {

	} else {
		delete(n.MapService[s.ServerType], s.AppId)
	}
	n.serviceMapLock.Unlock()
}

func (n *Node) GetAllService() map[uint32][]*Service {
	allServices := make(map[uint32][]*Service, 0)
	n.serviceMapLock.Lock()
	for serverType, serviceList := range n.MapService {
		if allServices[uint32(serverType)] == nil {
			allServices[uint32(serverType)] = make([]*Service, 0)
		}
		for _, service := range serviceList {
			allServices[uint32(serverType)] = append(allServices[uint32(serverType)], service)
		}
	}
	n.serviceMapLock.Unlock()
	return allServices
}

// 公共服务注册方法
func (n *Node) ServiceConnectionReq(serviceMsg pb.Message, s *Service) {
	req := serviceMsg.(*spb.ServiceConnectionReq)
	if req.AppId == 0 || req.ServerType == 0 {
		logger.Debug("Service registration failed")
		n.killService(s)
		return
	}
	switch req.ServerType {
	case spb.ServerType_SERVICE_GATE:
		go s.gateRecvHandle()
	case spb.ServerType_SERVICE_GAME:
		go s.gameRecvHandle()
	case spb.ServerType_SERVICE_DISPATCH:
		go s.dispatchRecvHandle()
	case spb.ServerType_SERVICE_MUIP:
		go s.muipRecvHandle()
	default:
		logger.Info("Service registration failed")
		return
	}

	s.AppId = req.AppId
	s.ServerType = req.ServerType
	s.Addr = req.Addr
	s.Port = req.Port
	n.AddMapService(s)

	rsp := &spb.ServiceConnectionRsp{
		ServerType: req.ServerType,
		AppId:      req.AppId,
	}

	s.sendHandle(cmd.ServiceConnectionRsp, rsp)
}
