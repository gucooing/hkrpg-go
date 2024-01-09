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

	// 启动 HTTP 服务器
	httpAddr := s.Config.Http.Addr + ":" + strconv.FormatInt(s.Config.Http.Port, 10)
	go s.startServer(httpAddr, "HTTP")

	// 根据配置决定是否启动 HTTPS 服务器
	if s.Config.Https != nil && s.Config.Https.Enable {
		httpsAddr := s.Config.Https.Addr + ":" + strconv.FormatInt(s.Config.Https.Port, 10)
		go s.startServer(httpsAddr, "HTTPS")
	}

	return nil
}

// startServer 启动一个 HTTP/HTTPS 服务器
func (s *Server) startServer(addr string, serverType string) {
	server := &http.Server{Addr: addr, Handler: s.Router}
	logger.Info("hkrpg-go SDK 服务器在 %s 启动 (%s)", addr, serverType)

	var err error
	if serverType == "HTTPS" {
		err = server.ListenAndServeTLS(s.Config.Https.CertFile, s.Config.Https.KeyFile)
	} else {
		err = server.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		logger.Error("hkrpg-go SDK 服务器启动失败, 原因: %s", err)
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
