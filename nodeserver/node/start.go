package node

import (
	"bufio"
	"log"
	"net"
	"os"

	"github.com/gucooing/hkrpg-go/nodeserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen = 343 * 1024 // 最大应用层包长度
)

var NODE *Node = nil

type Node struct {
	AppId      string
	Port       string
	Config     *config.Config
	MapService map[spb.ServerType]map[string]*Service // [ServerType][appid][Service]
	PlayerMap  map[uint32]*PlayerService              // [uid][gateAppId][GameAppId]
}

type Service struct {
	Conn       net.Conn
	AppId      string
	ServerType spb.ServerType
	Addr       string
	Port       string
	PlayerNum  uint64
}

type PlayerService struct {
	GateAppId        string
	GameAppId        string
	PlayerOnlineData []byte
	PlayerStatus     *PlayerStatus
}

type PlayerStatus struct {
	Status     spb.PlayerStatus
	GateStatus spb.PlayerGateStatus
	GameStatus spb.PlayerGameStatus
}

func GetPlayerGame(uid uint32) *Service {
	if NODE.PlayerMap[uid] == nil {
		return nil
	}
	return NODE.MapService[spb.ServerType_SERVICE_GAME][NODE.PlayerMap[uid].GameAppId]
}

func GetPlayerGate(uid uint32) *Service {
	if NODE.PlayerMap[uid] == nil {
		return nil
	}
	return NODE.MapService[spb.ServerType_SERVICE_GATE][NODE.PlayerMap[uid].GateAppId]
}

func NewNode(cfg *config.Config) *Node {
	NODE = new(Node)
	NODE.Config = cfg
	NODE.AppId = alg.GetAppId()
	logger.Info("NodeServer AppId:%s", NODE.AppId)
	port := NODE.Config.AppList[NODE.AppId].App["port_service"].Port
	if port == "" {
		log.Println("Node port error")
		os.Exit(0)
	}
	NODE.Port = port
	NODE.MapService = GetMapService()
	NODE.PlayerMap = make(map[uint32]*PlayerService)

	return NODE
}

func GetMapService() map[spb.ServerType]map[string]*Service {
	if NODE.MapService == nil {
		NODE.MapService = make(map[spb.ServerType]map[string]*Service)
		NODE.MapService[spb.ServerType_SERVICE_NODE] = make(map[string]*Service)
		NODE.MapService[spb.ServerType_SERVICE_GAME] = make(map[string]*Service)
		NODE.MapService[spb.ServerType_SERVICE_GATE] = make(map[string]*Service)
		NODE.MapService[spb.ServerType_SERVICE_DISPATCH] = make(map[string]*Service)
		NODE.MapService[spb.ServerType_SERVICE_MULTI] = make(map[string]*Service)
		NODE.MapService[spb.ServerType_SERVICE_MUIP] = make(map[string]*Service)
	}
	return NODE.MapService
}

func (n *Node) NewNode() {
	logger.Info("此NodeServer端口为:%v", n.Port)
	// 监听地址和端口
	listen, err := net.Listen("tcp", "0.0.0.0:"+n.Port)
	if err != nil {
		logger.Error("NodeServer监听失败:%s", err.Error())
		os.Exit(0)
	}
	defer listen.Close()

	logger.Info("NodeServer已启动")

	for {
		conn, err := listen.Accept()
		if err != nil {
			logger.Error("NodeServer接受连接失败:%s", err.Error())
			continue
		}
		// logger.Info("未知服务尝试连接addr:%s", conn.RemoteAddr().String())
		s := NewService(conn)
		go recvHandle(s)
	}
}

func NewService(conn net.Conn) *Service {
	s := new(Service)
	s.Conn = conn
	return s
}

func recvHandle(s *Service) {
	payload := make([]byte, PacketMaxLen)
	var bin []byte = nil
	recvLen, err := bufio.NewReader(s.Conn).Read(payload)
	if err != nil {
		// s.killService()
		return
	}
	bin = payload[:recvLen]
	msgList := make([]*alg.PackMsg, 0)
	alg.DecodeBinToPayload(bin, &msgList, nil)
	for _, msg := range msgList {
		serviceMsg := alg.DecodePayloadToProto(msg)
		// s.RegisterMessage(msg.CmdId, serviceMsg)
		if msg.CmdId == cmd.ServiceConnectionReq {
			s.ServiceConnectionReq(serviceMsg)
		} else {
			continue
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

func (s *Service) killService() {
	logger.Info("[%s]服务离线:%s", s.ServerType, s.Conn.RemoteAddr().String())
	s.Conn.Close()
	delete(NODE.MapService[s.ServerType], s.AppId)
}
