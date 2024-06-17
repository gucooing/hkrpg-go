package sdk

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/dispatch/db"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

type Server struct {
	AppId      uint32
	Port       string
	InnerAddr  string
	OuterAddr  string
	node       *NodeService
	Config     *config.Config
	Store      *db.Store
	Router     *gin.Engine
	server     *http.Server
	AutoCreate sync.Mutex
	Ec2b       *random.Ec2b
	Ticker     *time.Ticker
	Stop       chan struct{}
}

// 初始化所有服务
func NewServer(cfg *config.Config, appid string) *Server {
	s := &Server{}
	s.Config = cfg
	s.AppId = alg.GetAppIdUint32(appid)
	logger.Info("Dispatch AppId:%s", appid)
	appConf := s.Config.AppList[appid].App["port_http"]
	if appConf.Port == "" {
		log.Println("Dispatch Port error")
		os.Exit(0)
	}
	s.Port = appConf.Port
	s.InnerAddr = appConf.InnerAddr
	s.OuterAddr = appConf.OuterAddr
	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	gin.SetMode(gin.ReleaseMode)    // 初始化gin
	s.Router = gin.New()            // gin.Default()
	s.Router.Use(gin.Recovery())
	s.Ec2b = alg.GetEc2b() // 读取ec2b密钥
	// 开启dispatch定时器
	s.Ticker = time.NewTicker(5 * time.Second)
	s.Stop = make(chan struct{})
	go s.dispatchTicker()

	return s
}

func (s *Server) Start() error {
	// 初始化路由
	s.InitRouter()
	httpsAddr := s.InnerAddr + ":" + s.Port
	err := s.startServer(httpsAddr)
	return err
}

func (s *Server) startServer(addr string) error {
	var err error
	server := &http.Server{Addr: addr, Handler: s.Router}
	logger.Info("dispatch监听地址:%s", addr)
	logger.Info("dispatch对外地址:%s", s.OuterAddr)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("dispatch 服务器启动失败, 原因: %s", err)
		return err
	}
	return nil
}

func (s *Server) Shutdown(context.Context) error {
	if s.server == nil {
		return nil
	}
	close(s.Stop)
	return s.server.Close()
}

func clientIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// logger.Debug("http req:%s", c.Request.RequestURI)
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.Next()
			return
		}

		// 将 IP 信息存储在 gin.Context 中
		c.Set("IP", ip)
		c.Next()
	}
}

func (s *Server) dispatchTicker() {
	go func() {
		for {
			select {
			case <-s.Ticker.C:
				s.GlobalRotationEvent()
			case <-s.Stop:
				s.Ticker.Stop()
				return
			}
		}
	}()
}

func (s *Server) GlobalRotationEvent() {
	// 检查node是否存在
	if s.node == nil {
		logger.Info("尝试连接node")
		s.newNode()
	}
}
