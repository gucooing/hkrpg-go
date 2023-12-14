package SDK

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/pkg/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

type Server struct {
	Config     *config.Config
	Store      *DataBase.Store
	Router     *gin.Engine
	server     *http.Server
	AutoCreate sync.Mutex
}

func (s *Server) Start() error {
	// 初始化路由
	s.InitRouter()
	// 获取地址
	addr := s.Config.Http.Addr
	fullAddr := addr + ":" + strconv.FormatInt(s.Config.Http.Port, 10)
	go s.startServer(fullAddr, "HTTP")

	return nil
}

// startServer 启动一个 HTTP 服务器
func (s *Server) startServer(addr string, serverType string) {
	s.server = &http.Server{Addr: addr, Handler: s.Router}
	logger.Info("hkrpg-go SDK 服务器在%s启动", addr)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("hkrpg-go SDK 服务器启动失败,原因:%s", err)

	}
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
