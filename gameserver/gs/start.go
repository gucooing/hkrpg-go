package gs

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/gucooing/gunet"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	Ticker = 5 // 定时器间隔时间 / s
)

var GAMESERVER *GameServer
var PLAYERNUM int64 // 玩家人数

type GameServer struct {
	Config       *config.Config
	Store        *db.Store
	Port         string
	AppId        uint32
	GSListener   *gunet.TcpListener
	node         *NodeService
	gateList     map[uint32]*gateServer // gate列表
	gateListLock sync.Mutex             // gate列表同步锁
	Ticker       *time.Ticker
	everyDay4    *time.Ticker
	Stop         chan struct{}
}

func NewGameServer(cfg *config.Config, appid string) *GameServer {
	s := new(GameServer)
	GAMESERVER = s
	s.Config = cfg
	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	s.AppId = alg.GetAppIdUint32(appid)
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
	addr := s.Config.OuterIp + ":" + port
	gSListener, err := gunet.NewTcpS(addr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}
	s.GSListener = gSListener
	// 开启game定时器
	s.Ticker = time.NewTicker(Ticker * time.Second)
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	s.everyDay4 = time.NewTicker(everyDay4)
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

func (s *GameServer) recvNil(conn *gunet.TcpConn) {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATESERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			return
		}
	}()
	for {
		bin, err := conn.Read()
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			conn.Close()
			return
		}
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			switch msg.CmdId {
			case cmd.GateLoginGameReq:
				rsp := serviceMsg.(*spb.GateLoginGameReq)
				go s.recvGate(conn, rsp.AppId)
			}
			return
		}
	}
}

func (s *GameServer) AutoUpDataPlayer() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		for _, ge := range s.gateList {
			playerList := ge.GetAllPlayer()
			for _, g := range playerList {
				if g.p.Uid == 0 {
					continue
				}
				lastActiveTime := g.LastActiveTime
				timestamp := time.Now().Unix()
				if timestamp-lastActiveTime >= 180 {
					logger.Info("[UID:%v]玩家数据自动保存", g.p.Uid)
					s.UpDataPlayer(g.p)
					g.LastActiveTime = timestamp + rand.Int63n(120)
				}
			}
		}
	}
}

func (s *GameServer) Close() error {
	for _, ge := range s.gateList {
		playerList := ge.GetAllPlayer()
		for _, g := range playerList {
			if g.p.Uid == 0 {
				continue
			}
			s.UpDataPlayer(g.p)
		}
	}
	return nil
}

func (s *GameServer) gameTicker() {
	for {
		select {
		case <-s.Ticker.C:
			s.GlobalRotationEvent()
		case <-s.everyDay4.C: // 4点事件
			everyDay4 := alg.GetEveryDay4()
			logger.Debug("离下一个刷新时间:%v", everyDay4)
			s.everyDay4 = time.NewTicker(everyDay4)
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
