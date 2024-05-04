package node

import (
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type Service struct {
	n             *Node
	Conn          *gunet.TcpConn
	AppId         uint32
	ServerType    spb.ServerType
	Addr          string
	Port          string
	PlayerNum     int64
	lastAliveTime int64
}

func (n *Node) recvHandle(s *Service) {
	bin, err := s.Conn.Read()
	if err != nil {
		logger.Warn("已切断异常连接Addr:%s", s.Conn.RemoteAddr().String())
		s.Conn.Close()
		return
	}
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
	if n.DelMapService(s) {
		logger.Info("[%s]服务离线:%s", s.ServerType, s.Conn.RemoteAddr().String())
	}
}

func (n *Node) AddMapService(s *Service) {
	n.serviceMapLock.Lock()
	n.MapService[s.ServerType][s.AppId] = s
	n.serviceMapLock.Unlock()
	logger.Info("AppId:%s Service:%s Service registration successful", alg.GetAppIdStr(s.AppId), s.ServerType)
}

func (n *Node) DelMapService(s *Service) bool {
	n.serviceMapLock.Lock()
	defer n.serviceMapLock.Unlock()
	if n.MapService[s.ServerType] == nil || n.MapService[s.ServerType][s.AppId] == nil {
		return false
	} else {
		delete(n.MapService[s.ServerType], s.AppId)
		return true
	}
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

func (n *Node) GetAllServiceByType(serverType spb.ServerType) []*Service {
	servicesList := make([]*Service, 0)
	n.serviceMapLock.Lock()
	all := n.MapService[serverType]
	n.serviceMapLock.Unlock()
	for _, service := range all {
		servicesList = append(servicesList, service)
	}
	return servicesList
}

func (n *Node) GetAllServiceByTypeId(serverType spb.ServerType, appid uint32) *Service {
	n.serviceMapLock.Lock()
	defer n.serviceMapLock.Unlock()
	serviceList := n.MapService[serverType]
	if serviceList == nil {
		return nil
	}
	return serviceList[appid]
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
	case spb.ServerType_SERVICE_MULTI:
		go s.multiRecvHandle()
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
	s.lastAliveTime = time.Now().Unix()
	n.AddMapService(s)

	rsp := &spb.ServiceConnectionRsp{
		ServerType: req.ServerType,
		AppId:      req.AppId,
	}

	s.sendHandle(cmd.ServiceConnectionRsp, rsp)
}

func (n *Node) removeDeadServer() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		<-ticker.C
		nowTime := time.Now().Unix()
		for _, serviceList := range n.GetAllService() {
			for _, service := range serviceList {
				if nowTime-service.lastAliveTime > 30 {
					n.killService(service)
					logger.Info("[APPID:%v]删除死服务器", service.AppId)
				}
			}
		}
	}
}
