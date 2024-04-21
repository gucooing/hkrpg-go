package gs

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

const (
	Ticker = 5 // 心跳包间隔时间 / s
)

var GAMESERVER *GameServer

type GameServer struct {
	Config     *config.Config
	Store      *db.Store
	Port       string
	AppId      uint32
	GSListener net.Listener
	nodeConn   net.Conn
	PlayerMap  map[int64]*player.GamePlayer

	gateList     map[uint32]*gateServer // gate列表
	gateListLock sync.Mutex             // gate列表同步锁

	RecvCh chan *TcpNodeMsg
	Ticker *time.Ticker
	Stop   chan struct{}
}

type TcpNodeMsg struct {
	cmdId      uint16
	serviceMsg pb.Message
}

func NewGameServer(cfg *config.Config, appid string) *GameServer {
	s := new(GameServer)

	GAMESERVER = s

	s.Config = cfg
	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	s.AppId = alg.GetAppIdUint32(appid)
	player.SNOWFLAKE = alg.NewSnowflakeWorker(1)
	logger.Info("GameServer AppId:%s", appid)
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

	s.RecvCh = make(chan *TcpNodeMsg)
	s.Ticker = time.NewTicker(Ticker * time.Second)
	s.Stop = make(chan struct{})
	s.ServiceStart()

	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		log.Println("nodeserver error")
		os.Exit(0)
	}
	s.nodeConn = tcpConn
	s.PlayerMap = make(map[int64]*player.GamePlayer)
	s.gateList = make(map[uint32]*gateServer)

	go s.recvNode()
	go s.AutoUpDataPlayer()
	// 向node注册
	s.ServiceConnectionReq()

	return s
}

func (s *GameServer) GetPlayerNum() int64 {
	return 0
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

func KickPlayer(g *player.GamePlayer) {
	if err := UpDataPlayer(g); err != nil {
		logger.Error("[UID:%v]保存数据失败", g.Uid)
	}
	GAMESERVER.DelPlayerMap(g.Uuid)
	if g.GateConn != nil {
		g.GateConn.Close()
	}
	logger.Info("[UID:%v]玩家离线game", g.Uid)
}

func UpDataPlayer(g *player.GamePlayer) error {
	var err error
	if g.PlayerPb == nil {
		return nil
	}
	if g.Uid == 0 {
		return nil
	}
	if bin, ok := db.DBASE.GetPlayerStatus(strconv.Itoa(int(g.AccountId))); !ok {
		return nil
	} else {
		statu := new(spb.PlayerStatusRedisData)
		err := pb.Unmarshal(bin, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			return err
		}
		if statu.GameserverId != GAMESERVER.AppId || statu.Uuid != g.Uuid {
			// 脏数据
			return nil
		}
	}
	dbDate := new(db.PlayerData)
	dbDate.Uid = g.Uid
	dbDate.Level = g.PlayerPb.Level
	dbDate.Exp = g.PlayerPb.Exp
	dbDate.Nickname = g.PlayerPb.Nickname
	dbDate.BinData, err = pb.Marshal(g.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return err
	}

	if err = db.DBASE.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return err
	}
	return nil
}
