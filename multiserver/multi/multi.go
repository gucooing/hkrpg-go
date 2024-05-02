package multi

import (
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/multiserver/db"

	"github.com/gucooing/hkrpg-go/multiserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

const (
	Ticker = 5 // 定时器间隔时间 / s
)

type Multi struct {
	Config *config.Config
	AppId  uint32
	Node   *NodeService
	store  *db.Store

	gsList     map[uint32]*gameServer // gs列表
	gsListLock sync.Mutex             // gs列表互斥锁

	Ticker *time.Ticker
	Stop   chan struct{}
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

	// 启动muip定时器
	s.Ticker = time.NewTicker(Ticker * time.Second)
	s.Stop = make(chan struct{})
	go s.gameTicker()

	return s
}

func (s *Multi) gameTicker() {
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

func (s *Multi) GlobalRotationEvent() {
	// 检查node是否存在
	if s.Node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}

func (s *Multi) Close() {

}
