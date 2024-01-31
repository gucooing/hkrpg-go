package game

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var GAMESERVER *GameServer

type GameServer struct {
	Config     *config.Config
	Store      *db.Store
	Port       string
	AppId      string
	GSListener net.Listener
	nodeConn   net.Conn
	PlayerMap  map[uint32]*player.GamePlayer
}

func NewGameServer(cfg *config.Config) *GameServer {
	s := new(GameServer)

	GAMESERVER = s

	s.Config = cfg
	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	s.AppId = alg.GetAppId()
	player.SNOWFLAKE = alg.NewSnowflakeWorker(1)
	logger.Info("GameServer AppId:%s", s.AppId)
	port := s.Config.AppList[s.AppId].App["port_gt"].Port
	if port == "" {
		log.Println("GameServer Port error")
		os.Exit(0)
	}
	s.Port = port
	addr := "0.0.0.0:" + port
	gSListener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}
	s.GSListener = gSListener
	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		log.Println("nodeserver error")
		os.Exit(0)
	}
	s.nodeConn = tcpConn
	s.PlayerMap = make(map[uint32]*player.GamePlayer)

	go s.recvNode()
	go s.AutoUpDataPlayer()
	// 向node注册
	s.Connection()

	return s
}

func (s *GameServer) StartGameServer() error {
	for {
		conn, err := s.GSListener.Accept()
		if err != nil {
			logger.Info("GameServer接受连接失败:%s", err.Error())
			continue
		}
		g := NewPlayer(conn)
		go s.recvGate(g)
	}
}

func NewPlayer(conn net.Conn) *player.GamePlayer {
	g := new(player.GamePlayer)
	g.GateConn = conn

	return g
}

func (s *GameServer) AutoUpDataPlayer() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		for _, g := range s.PlayerMap {
			if g.Uid == 0 {
				return
			}
			lastActiveTime := g.LastActiveTime
			timestamp := time.Now().Unix()
			if timestamp-lastActiveTime >= 120 {
				logger.Info("[UID:%v]玩家超时离线", g.Uid)
				KickPlayer(g)
			}
		}
	}
}

func Close() error {
	for _, gamePlayer := range GAMESERVER.PlayerMap {
		KickPlayer(gamePlayer)
	}
	return nil
}
