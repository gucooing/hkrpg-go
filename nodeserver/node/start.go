package node

import (
	"bufio"
	"log"
	"net"
	"os"
	"sync"

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
var syncPlayerUuidMap sync.Mutex
var syncPlayerMap sync.Mutex

type Node struct {
	AppId         uint32
	Port          string
	Config        *config.Config
	MapService    map[spb.ServerType]map[uint32]*Service // [ServerType][appid][Service]
	PlayerUuidMap map[uint32]int64                       // [uid][uuid]
	PlayerMap     map[int64]*PlayerService               // [uid][gateAppId][GameAppId]
}

type Service struct {
	Conn       net.Conn
	AppId      uint32
	ServerType spb.ServerType
	Addr       string
	Port       string
	PlayerNum  uint64
}

type PlayerService struct {
	Uuid      int64
	Uid       uint32
	GateAppId uint32
	GameAppId uint32
}

func NewNode(cfg *config.Config, appid string) *Node {
	NODE = new(Node)
	NODE.Config = cfg
	NODE.AppId = alg.GetAppIdUint32(appid)
	logger.Info("NodeServer AppId:%s", appid)
	port := NODE.Config.AppList[appid].App["port_service"].Port
	if port == "" {
		log.Println("Node port error")
		os.Exit(0)
	}
	NODE.Port = port
	NODE.MapService = GetMapService()
	NODE.PlayerUuidMap = make(map[uint32]int64)
	NODE.PlayerMap = make(map[int64]*PlayerService)

	return NODE
}

func GetMapService() map[spb.ServerType]map[uint32]*Service {
	if NODE.MapService == nil {
		NODE.MapService = make(map[spb.ServerType]map[uint32]*Service)
		NODE.MapService[spb.ServerType_SERVICE_NODE] = make(map[uint32]*Service)
		NODE.MapService[spb.ServerType_SERVICE_GAME] = make(map[uint32]*Service)
		NODE.MapService[spb.ServerType_SERVICE_GATE] = make(map[uint32]*Service)
		NODE.MapService[spb.ServerType_SERVICE_DISPATCH] = make(map[uint32]*Service)
		NODE.MapService[spb.ServerType_SERVICE_MULTI] = make(map[uint32]*Service)
		NODE.MapService[spb.ServerType_SERVICE_MUIP] = make(map[uint32]*Service)
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

func AddPlayerUuidMap(uuid int64, uid uint32) {
	syncPlayerUuidMap.Lock()
	NODE.PlayerUuidMap[uid] = uuid
	syncPlayerUuidMap.Unlock()
}

func DelPlayerUuidMap(uid uint32) {
	if NODE.PlayerUuidMap[uid] != 0 {
		syncPlayerUuidMap.Lock()
		delete(NODE.PlayerUuidMap, uid)
		syncPlayerUuidMap.Unlock()
	}
}

func AddPlayerMap(uuid int64, p *PlayerService) {
	syncPlayerMap.Lock()
	NODE.PlayerMap[uuid] = p
	syncPlayerMap.Unlock()
}

func DelPlayerMap(uuid int64) {
	if NODE.PlayerMap[uuid] != nil {
		syncPlayerMap.Lock()
		delete(NODE.PlayerMap, uuid)
		syncPlayerMap.Unlock()
	}
}

func getPlayerServiceByUuid(uid uint32) *PlayerService {
	return NODE.PlayerMap[NODE.PlayerUuidMap[uid]]
}

func getGsByAppId(appid uint32) *Service {
	return NODE.MapService[spb.ServerType_SERVICE_GAME][appid]
}
