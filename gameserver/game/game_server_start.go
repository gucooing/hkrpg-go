package game

import (
	"net"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/logger"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
)

var GAMESERVER *GameServer

type GameServer struct {
	Config     *config.Config
	Store      *db.Store
	Port       string
	AppId      string
	GSListener net.Listener
	nodeConn   net.Conn
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
		panic("GameServer Port error")
	}
	s.Port = port
	addr := "0.0.0.0:" + port
	gSListener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err.Error())
		return nil
	}
	s.GSListener = gSListener
	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		panic(err.Error())
		return nil
	}
	s.nodeConn = tcpConn
	go s.recvNode()
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
		go g.RecvGate()
	}
}

func NewPlayer(conn net.Conn) *player.GamePlayer {
	g := new(player.GamePlayer)
	g.GateConn = conn

	return g
}
