package gameserver

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/dbconf"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	Ticker                 = 5   // 定时器间隔时间 / s
	AutoUpDataPlayerTicker = 120 // 定时执行玩家数据保存间隔时间 / s
)

var PLAYERNUM int64 // 玩家人数

type GameServer struct {
	Config           *Config
	Store            *database.GameStore
	Port             string
	InnerAddr        string
	OuterAddr        string
	AppId            uint32
	GSListener       *gunet.TcpListener
	node             *NodeService
	gateList         map[uint32]*gateServer // gate列表
	gateListLock     sync.Mutex             // gate列表同步锁
	playerMap        map[uint32]*GamePlayer // 玩家列表
	playerMapLock    sync.Mutex             // 玩家列表互斥锁
	Ticker           *time.Ticker
	everyDay4        *time.Ticker
	autoUpDataPlayer *time.Ticker
	Stop             chan struct{}
}

func NewGameServer(cfg *Config, appid string) *GameServer {
	s := new(GameServer)
	s.Config = cfg
	s.Store = database.NewGameStore(s.Config.MysqlConf, s.Config.RedisConf) // 初始化数据库连接
	dbconf.GameServer(s.Store.ServerConf)
	s.AppId = alg.GetAppIdUint32(appid)
	s.gateList = make(map[uint32]*gateServer)
	s.playerMap = make(map[uint32]*GamePlayer)
	logger.Info("GameServer AppId:%s", appid)
	// 开启tcp服务
	appConf := s.Config.AppList[appid].App["port_gt"]
	if appConf.Port == "" {
		log.Println("GameServer Port error")
		os.Exit(0)
	}
	s.Port = appConf.Port
	s.InnerAddr = appConf.InnerAddr
	s.OuterAddr = appConf.OuterAddr
	addr := s.InnerAddr + ":" + s.Port
	gSListener, err := gunet.NewTcpS(addr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}
	s.GSListener = gSListener
	// 开启game定时器
	s.Ticker = time.NewTicker(Ticker * time.Second)
	s.autoUpDataPlayer = time.NewTicker(AutoUpDataPlayerTicker * time.Second)
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	s.everyDay4 = time.NewTicker(everyDay4)
	s.Stop = make(chan struct{})
	go s.gameTicker()
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
			serviceMsg := cmd.DecodePayloadToProto(msg)
			switch msg.CmdId {
			case cmd.GateLoginGameReq:
				rsp := serviceMsg.(*spb.GateLoginGameReq)
				go s.newGate(conn, rsp.AppId)
			}
			return
		}
	}
}

func (s *GameServer) AutoUpDataPlayer() {
	logger.Info("开始自动保存玩家数据")
	timestamp := time.Now().Unix()
	playerList := s.getAllPlayer()
	for _, g := range playerList {
		if g.p.Uid == 0 {
			continue
		}
		if g.lastActiveTime+50 < timestamp {
			logger.Info("[UID:%v]超时离线", g.p.Uid)
			s.killPlayer(g)
			continue
		}
		if timestamp-g.p.LastUpDataTime >= 180 {
			logger.Debug("[UID:%v]玩家数据自动保存", g.p.Uid)
			g.p.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_ONLINE)
			s.AddPlayerStatus(g) // 刷新状态
			g.p.LastUpDataTime = timestamp + rand.Int63n(120)
		}
	}
	logger.Info("自动保存玩家数据结束")
}

func (s *GameServer) Close() error {
	for _, g := range s.getAllPlayer() {
		if g.p.Uid == 0 {
			continue
		}
		g.p.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_OFFLINE)
	}
	return nil
}

func (s *GameServer) gameTicker() {
	for {
		select {
		case <-s.Ticker.C:
			s.GlobalRotationEvent5s()
		case <-s.autoUpDataPlayer.C:
			s.AutoUpDataPlayer()
		case <-s.everyDay4.C: // 4点事件
			s.GlobalRotationEvent4h()
		case <-s.Stop:
			s.Ticker.Stop()
			return
		}
	}
}

func (s *GameServer) GlobalRotationEvent5s() {
	// 检查node是否存在
	if s.node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}

func (s *GameServer) GlobalRotationEvent4h() {
	playes := s.getAllPlayer()
	for _, g := range playes {
		if g.p.RecvChan != nil {
			g.p.RecvChan <- player.Msg{
				MsgType: player.DailyTask,
			}
		}
	}
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	s.everyDay4 = time.NewTicker(everyDay4)
}
