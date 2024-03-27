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
	AppId      string
	Port       string
	NodeConn   net.Conn
	GateAddr   string
	GatePort   string
	Config     *config.Config
	Store      *db.Store
	Router     *gin.Engine
	server     *http.Server
	AutoCreate sync.Mutex
	Ec2b       *random.Ec2b

	RecvCh chan *TcpNodeMsg
	Ticker *time.Ticker
	Stop   chan struct{}
}

// 初始化所有服务
func NewServer(cfg *config.Config) *Server {
	s := &Server{}
	s.Config = cfg
	s.AppId = alg.GetAppId()
	logger.Info("Dispatch AppId:%s", s.AppId)
	port := s.Config.AppList[s.AppId].App["port_http"].Port
	if port == "" {
		log.Println("Dispatch Port error")
		os.Exit(0)
	}
	s.Port = port
	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	gin.SetMode(gin.ReleaseMode)    // 初始化gin
	s.Router = gin.Default()        // gin.New()
	s.Router.Use(gin.Recovery())
	s.Ec2b = alg.GetEc2b() // 读取ec2b密钥

	s.RecvCh = make(chan *TcpNodeMsg)
	s.Ticker = time.NewTicker(5 * time.Second)
	s.Stop = make(chan struct{})
	s.ServiceStart()

	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		log.Println("nodeserver error")
		panic(err)
	}
	s.NodeConn = tcpConn
	go s.RecvNode()
	// 向node注册
	s.Connection()

	return s
}

func (s *Server) Start() error {
	// 初始化路由
	s.InitRouter()
	httpsAddr := s.Config.OuterIp + ":" + s.Port
	err := s.startServer(httpsAddr)
	return err
}

func (s *Server) startServer(addr string) error {
	var err error
	server := &http.Server{Addr: addr, Handler: s.Router}
	logger.Info("dispatch 在 %s 启动", addr)
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
