package muip

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/muipserver/config"
	"github.com/gucooing/hkrpg-go/muipserver/db"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	Ticker = 5 // 定时器间隔时间 / s
)

type MuipServer struct {
	AppId          uint32
	ApiAddr        string
	Store          *db.Store
	Config         *config.Config
	Api            *Api
	node           *NodeService
	allService     map[spb.ServerType][]*Service
	allServiceSync sync.Mutex
	Ticker         *time.Ticker
	Stop           chan struct{}
}

type Api struct {
	muip   *MuipServer
	Router *gin.Engine
}

func NewMuip(config *config.Config, appid string) *MuipServer {
	s := new(MuipServer)
	s.Config = config
	s.AppId = alg.GetAppIdUint32(appid)
	s.Store = db.NewStore(config)
	logger.Info("MuipServer AppId:%s", appid)
	appConf := s.Config.AppList[appid].App["port_http"]
	if appConf.Port == "" {
		log.Println("MuipServer Port error")
		os.Exit(0)
	}
	s.ApiAddr = appConf.OuterAddr + ":" + appConf.Port
	s.Api = s.newApi()
	s.allService = make(map[spb.ServerType][]*Service)
	// 开启game定时器
	s.Ticker = time.NewTicker(Ticker * time.Second)
	s.Stop = make(chan struct{})
	go s.muipTicker()
	return s
}

func (s *MuipServer) newApi() *Api {
	a := new(Api)
	a.muip = s
	gin.SetMode(gin.ReleaseMode) // 初始化gin
	a.Router = gin.New()         // gin.Default()
	a.Router.Use(gin.Recovery())
	// 初始化路由
	a.InitRouter()

	return a
}

func (a *Api) StartApi() error {
	server := &http.Server{Addr: a.muip.ApiAddr, Handler: a.Router}
	logger.Info("Api 在 %s 启动", a.muip.ApiAddr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {

		return err
	}
	return nil
}

func (s *MuipServer) muipTicker() {
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

func (s *MuipServer) GlobalRotationEvent() {
	// 检查node是否存在
	if s.node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}

func (s *MuipServer) setAllService(allService map[spb.ServerType][]*Service) {
	s.allServiceSync.Lock()
	s.allService = allService
	s.allServiceSync.Unlock()
}

func (s *MuipServer) getAllService() map[spb.ServerType][]*Service {
	s.allServiceSync.Lock()
	allService := make(map[spb.ServerType][]*Service, 0)
	for id, serviceList := range s.allService {
		if allService[id] == nil {
			allService[id] = make([]*Service, 0)
		}
		for _, service := range serviceList {
			allService[id] = append(allService[id], service)
		}
	}
	s.allServiceSync.Unlock()
	return allService
}

func (s *MuipServer) Close() {
	// 通知node服务下线
}
