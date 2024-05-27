package gate

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

const (
	PacketMaxLen            = 343 * 1024 // 最大应用层包长度
	KcpConnAddrChangeNotify = "KcpConnAddrChangeNotify"
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数

type GateServer struct {
	AppId              uint32
	WorkerId           int64
	Port               string
	Config             *config.Config
	Store              *Store
	snowflake          *alg.SnowflakeWorker // 雪花唯一id生成器
	kcpListener        *kcp.Listener
	node               *NodeService
	sessionIdCounter   uint32
	kcpEventChan       chan *KcpEvent
	Ec2b               *random.Ec2b
	gsList             map[uint32]*gameServer // gs列表
	gsListLock         sync.Mutex             // gs列表互斥锁
	Ticker             *time.Ticker
	Stop               chan struct{}
	loginPlayerMap     map[uint32]*PlayerGame // 正在登录的玩家列表
	loginPlayerMapLock sync.Mutex             // 正在登录的玩家列表互斥锁
	playerMap          map[uint32]*PlayerGame // 玩家列表
	playerMapLock      sync.Mutex             // 玩家列表互斥锁
}

type KcpEvent struct {
	SessionId    uint32
	EventId      string
	EventMessage any
}

func NewGate(cfg *config.Config, appid string) *GateServer {
	s := new(GateServer)

	s.Ec2b = alg.GetEc2b()
	s.Config = cfg
	s.Store = NewStore(s.Config) // 初始化数据库连接
	s.gsList = make(map[uint32]*gameServer)
	s.AppId = alg.GetAppIdUint32(appid)
	s.WorkerId = 1
	s.snowflake = alg.NewSnowflakeWorker(s.WorkerId)
	s.loginPlayerMap = make(map[uint32]*PlayerGame)
	s.playerMap = make(map[uint32]*PlayerGame)
	logger.Info("GateServer AppId:%s", appid)
	// 开启kcp服务
	port := s.Config.AppList[appid].App["port_player"].Port
	if port == "" {
		log.Println("GateServer Port error")
		os.Exit(0)
	}
	s.Port = port

	addr := s.Config.InnerIp + ":" + s.Port
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		log.Printf("listen kcp err: %v\n", err)
		os.Exit(0)
	}
	s.kcpListener = kcpListener
	// 开启gate定时器
	s.Ticker = time.NewTicker(5 * time.Second)
	s.Stop = make(chan struct{})
	go s.gateTicker()
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATESERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.Close()
			os.Exit(0)
		}
	}()
	go s.kcpNetInfo()
	go s.kcpEnetHandle(kcpListener)
	go s.AutoDelPlayer()

	return s
}

func (s *GateServer) NewGame(kcpConn *kcp.UDPSession) *PlayerGame {
	g := new(PlayerGame)
	g.KcpConn = kcpConn
	g.XorKey = s.Ec2b.XorKey()
	g.LastActiveTime = getCurTime()
	// 初始化路由
	g.RouteManager = NewRouteManager(g)

	return g
}

// gate定时器
func (s *GateServer) gateTicker() {
	for {
		select {
		case <-s.Ticker.C:
			s.GlobalRotationEvent5s()
		case <-s.Stop:
			s.Ticker.Stop()
			return
		}
	}
}

func (s *GateServer) GlobalRotationEvent5s() {
	// 检查node是否存在
	if s.node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}

func (s *GateServer) Close() error {
	return nil
}
