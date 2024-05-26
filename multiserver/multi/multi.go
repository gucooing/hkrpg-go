package multi

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/multiserver/db"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"

	"github.com/gucooing/hkrpg-go/multiserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

const (
	Ticker = 5 // 定时器间隔时间 / s
)

var err error

type Multi struct {
	Config       *config.Config
	AppId        uint32
	addr         string
	port         string
	listener     *gunet.TcpListener
	Node         *NodeService
	store        *db.Store
	gateList     map[uint32]*gateServer // gate列表
	gateListLock sync.Mutex             // gate列表同步锁
	Ticker       *time.Ticker
	everyDay4    *time.Ticker
	Stop         chan struct{}
}

type AllService struct {
	AppId     uint32
	PlayerNum int64
}

func NewMulti(cfg *config.Config, appid string, store *db.Store) *Multi {
	s := new(Multi)
	s.Config = cfg
	s.AppId = alg.GetAppIdUint32(appid)
	logger.Info("MultiServer AppId:%s", appid)
	s.store = store
	s.gateList = make(map[uint32]*gateServer)
	// 开启tcp服务
	s.port = s.Config.AppList[appid].App["port_service"].Port
	if s.port == "" {
		log.Println("MultiServer Port error")
		os.Exit(0)
	}
	s.addr = s.Config.OuterIp
	addr := s.Config.OuterIp + ":" + s.port
	s.listener, err = gunet.NewTcpS(addr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}
	// 启动muip定时器
	everyDay4 := alg.GetEveryDay4()
	logger.Debug("离下一个刷新时间:%v", everyDay4)
	s.everyDay4 = time.NewTicker(everyDay4)
	s.Ticker = time.NewTicker(Ticker * time.Second)
	s.Stop = make(chan struct{})
	go s.gameTicker()

	return s
}

func (s *Multi) StartMultiServer() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			logger.Info("MultiServer接受连接失败:%s", err.Error())
			continue
		}
		go s.recvTcp(conn)
	}
}

func (s *Multi) recvTcp(conn *gunet.TcpConn) {
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
			if msg.CmdId == cmd.GateLoginMultiReq {
				rsp := serviceMsg.(*spb.GateLoginMultiReq)
				go s.recvGate(conn, rsp.AppId)
				return
			} else {

			}
		}
	}
}

func (s *Multi) gameTicker() {
	for {
		select {
		case <-s.Ticker.C:
			s.GlobalRotationEvent()
		case <-s.everyDay4.C:
			everyDay4 := alg.GetEveryDay4()
			logger.Debug("离下一个刷新时间:%v", everyDay4)
			s.everyDay4 = time.NewTicker(everyDay4)
		case <-s.Stop:
			s.Ticker.Stop()
			return
		}
	}
}

func (s *Multi) GlobalRotationEvent() {
	// 检查node是否存在
	if s.Node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}

func (s *Multi) Close() {

}
