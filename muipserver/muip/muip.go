package muip

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	api2 "github.com/gucooing/hkrpg-go/muipserver/api"
	"github.com/gucooing/hkrpg-go/muipserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

const (
	Ticker = 5 // 定时器间隔时间 / s
)

type Muip struct {
	Config *config.Config
	AppId  uint32
	Api    *api2.Api
	Node   *NodeService

	Ticker *time.Ticker
	Stop   chan struct{}
}

type AllService struct {
	AppId     uint32
	PlayerNum int64
}

func NewMuip(cfg *config.Config, appid string) *Muip {
	s := new(Muip)
	s.Config = cfg
	s.AppId = alg.GetAppIdUint32(appid)
	logger.Info("MuipServer AppId:%s", appid)

	// newApi
	port := s.Config.AppList[appid].App["port_http"].Port
	if port == "" {
		log.Println("Api Port error")
		os.Exit(0)
	}
	s.Api = &api2.Api{
		Addr:   "0.0.0.0:" + port,
		Router: gin.Default(), // gin.New(),
	}
	gin.SetMode(gin.ReleaseMode) // 初始化gin
	s.Api.Router.Use(gin.Recovery())
	s.Api.InitRouter()

	// 启动muip定时器
	s.Ticker = time.NewTicker(Ticker * time.Second)
	s.Stop = make(chan struct{})
	go s.gameTicker()

	return s
}

func (s *Muip) gameTicker() {
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

func (s *Muip) GlobalRotationEvent() {
	// 检查node是否存在
	if s.Node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}
