package sdk

import (
	"context"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/dispatch/db"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
