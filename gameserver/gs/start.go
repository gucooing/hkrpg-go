package gs

import (
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	Ticker = 5 // 心跳包间隔时间 / s
)

var GAMESERVER *GameServer
var PLAYERNUM int64 // 玩家人数

type GameServer struct {
	Config       *config.Config
	Store        *db.Store
	Port         string
	AppId        uint32
	GSListener   net.Listener
	node         *NodeService
	PlayerMap    map[int64]*GamePlayer
	gateList     map[uint32]*gateServer // gate列表
	gateListLock sync.Mutex             // gate列表同步锁
	Ticker       *time.Ticker
	Stop         chan struct{}
}

func NewGameServer(cfg *config.Config, appid string) *GameServer {
	s := new(GameServer)

	GAMESERVER = s

	s.Config = cfg
	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	s.AppId = alg.GetAppIdUint32(appid)
	s.PlayerMap = make(map[int64]*GamePlayer)
	s.gateList = make(map[uint32]*gateServer)
	player.SNOWFLAKE = alg.NewSnowflakeWorker(1)
	logger.Info("GameServer AppId:%s", appid)
	// 开启tcp服务
	port := s.Config.AppList[appid].App["port_gt"].Port
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
	// 开启game定时器
	s.Ticker = time.NewTicker(Ticker * time.Second)
	s.Stop = make(chan struct{})
	go s.gameTicker()
	go s.AutoUpDataPlayer()
	return s
}

func (s *GameServer) GetPlayerNum() int64 {
	return PLAYERNUM
}

func (s *GameServer) StartGameServer() error {
	for {
		conn, err := s.GSListener.Accept()
		if err != nil {
			logger.Info("GameServer接受连接失败:%s", err.Error())
			continue
		}
		go s.recvNil(conn)
	}
}

func (s *GameServer) recvNil(conn net.Conn) {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATESERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			return
		}
	}()
	nodeMsg := make([]byte, player.PacketMaxLen)
	var bin []byte = nil
	recvLen, err := conn.Read(nodeMsg)
	if err != nil {
		logger.Debug("exit recv loop, conn read err: %v", err)
		return
	}
	bin = nodeMsg[:recvLen]
	nodeMsgList := make([]*alg.PackMsg, 0)
	alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
	for _, msg := range nodeMsgList {
		serviceMsg := alg.DecodePayloadToProto(msg)
		if msg.CmdId == cmd.GateLoginGameReq {
			rsp := serviceMsg.(*spb.GateLoginGameReq)
			switch rsp.ServerType {
			case spb.ServerType_SERVICE_GATE:
				go s.recvGate(conn, rsp.AppId)
				return
			}
		}
	}
	conn.Close()
	return
}

func (s *GameServer) AutoUpDataPlayer() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		for _, g := range s.PlayerMap {
			if g.p.Uid == 0 {
				return
			}
			lastActiveTime := g.p.LastActiveTime
			timestamp := time.Now().Unix()
			if timestamp-lastActiveTime >= 120 {
				logger.Info("[UID:%v]玩家超时离线", g.p.Uid)
				KickPlayer(g.p)
			}
		}
	}
}

func Close() error {
	for _, gamePlayer := range GAMESERVER.PlayerMap {
		KickPlayer(gamePlayer.p)
	}
	return nil
}

func (s *GameServer) gameTicker() {
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

func (s *GameServer) GlobalRotationEvent() {
	// 检查node是否存在
	if s.node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}
