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
	"github.com/gucooing/hkrpg-go/protocol/cmd"
)

const (
	PacketMaxLen            = 343 * 1024 // 最大应用层包长度
	KcpConnAddrChangeNotify = "KcpConnAddrChangeNotify"
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数
var GATESERVER *GateServer

type GateServer struct {
	AppId            uint32
	WorkerId         int64
	Port             string
	Config           *config.Config
	Store            *Store
	snowflake        *alg.SnowflakeWorker // 雪花唯一id生成器
	kcpListener      *kcp.Listener
	node             *NodeService
	kcpFin           bool
	sessionIdCounter uint32
	playerMap        map[int64]*PlayerGame // 玩家内存
	playerMapLock    sync.Mutex            // 玩家列表互斥锁
	kcpEventChan     chan *KcpEvent
	Ec2b             *random.Ec2b
	gsList           map[uint32]*gameServer // gs列表
	gsListLock       sync.Mutex             // gs列表互斥锁
	Ticker           *time.Ticker
	Stop             chan struct{}
}

type KcpEvent struct {
	SessionId    uint32
	EventId      string
	EventMessage any
}

func NewGate(cfg *config.Config, appid string) *GateServer {
	s := new(GateServer)
	GATESERVER = s

	s.Ec2b = alg.GetEc2b()
	s.Config = cfg
	s.Store = NewStore(s.Config) // 初始化数据库连接
	s.playerMap = make(map[int64]*PlayerGame)
	s.gsList = make(map[uint32]*gameServer)
	s.AppId = alg.GetAppIdUint32(appid)
	s.WorkerId = 1
	s.snowflake = alg.NewSnowflakeWorker(s.WorkerId)
	logger.Info("GateServer AppId:%s", appid)
	// 开启kcp服务
	port := s.Config.AppList[appid].App["port_player"].Port
	if port == "" {
		log.Println("GateServer Port error")
		os.Exit(0)
	}
	s.Port = port
	addr := "0.0.0.0:" + s.Port
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
			Close()
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
	g.LastActiveTime = time.Now().Unix()
	// 初始化路由
	g.RouteManager = NewRouteManager(g)

	return g
}

// gate定时器
func (s *GateServer) gateTicker() {
	for {
		select {
		case <-s.Ticker.C:
			s.GlobalRotationEvent()
		case <-s.Stop:
			s.Ticker.Stop()
			return
		}
	}
}

func (s *GateServer) GlobalRotationEvent() {
	// 检查node是否存在
	if s.node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}

func Close() error {
	ges := GATESERVER
	ges.kcpFin = true
	plays := ges.GetAllPlayer()
	for _, play := range plays {
		play.GateToPlayer(cmd.PlayerKickOutScNotify, nil)
		play.PlayerLogoutCsReq(nil)
	}
	close(ges.Stop)
	return nil
}
