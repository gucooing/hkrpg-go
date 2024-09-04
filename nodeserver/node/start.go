package node

import (
	"log"
	"os"
	"sync"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/nodeserver/config"
	"github.com/gucooing/hkrpg-go/nodeserver/db"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

var NODE *Node = nil

type Node struct {
	AppId          uint32
	Port           string
	Config         *config.Config
	Store          *db.Store
	MapService     map[spb.ServerType]map[uint32]*Service // [ServerType][appid][Service]
	serviceMapLock sync.Mutex                             // 服务列表互斥锁
}

func NewNode(cfg *config.Config, appid string, store *db.Store) *Node {
	n := new(Node)
	NODE = n
	n.Config = cfg
	n.Store = store
	n.AppId = alg.GetAppIdUint32(appid)
	logger.Info("NodeServer AppId:%s", appid)
	port := n.Config.AppList[appid].App["port_service"].Port
	if port == "" {
		log.Println("Node port error")
		os.Exit(0)
	}
	n.Port = port
	n.MapService = n.GetMapService()

	return n
}

func (n *Node) GetMapService() map[spb.ServerType]map[uint32]*Service {
	if n.MapService == nil {
		n.MapService = make(map[spb.ServerType]map[uint32]*Service)
		n.MapService[spb.ServerType_SERVICE_GAME] = make(map[uint32]*Service)
		n.MapService[spb.ServerType_SERVICE_GATE] = make(map[uint32]*Service)
		n.MapService[spb.ServerType_SERVICE_DISPATCH] = make(map[uint32]*Service)
		n.MapService[spb.ServerType_SERVICE_MULTI] = make(map[uint32]*Service)
		n.MapService[spb.ServerType_SERVICE_MUIP] = make(map[uint32]*Service)
	}
	return n.MapService
}

func (n *Node) NewNode() {
	logger.Info("此NodeServer端口为:%v", n.Port)
	// 监听地址和端口
	listen, err := gunet.NewTcpS("0.0.0.0:" + n.Port)
	if err != nil {
		logger.Error("NodeServer监听失败:%s", err.Error())
		os.Exit(0)
	}
	defer listen.Close()
	logger.Info("NodeServer已启动")
	go n.removeDeadServer()
	for {
		conn, err := listen.Accept()
		if err != nil {
			logger.Error("NodeServer接受连接失败:%s", err.Error())
			continue
		}
		s := n.newService(conn)
		go n.recvHandle(s)
	}
}

func (n *Node) newService(conn *gunet.TcpConn) *Service {
	s := new(Service)
	s.Conn = conn
	s.n = n
	return s
}

func (n *Node) Close() {

}
