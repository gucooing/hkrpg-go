package dispatch

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/dispatch/db"
	"github.com/gucooing/hkrpg-go/dispatch/logger"
	"github.com/gucooing/hkrpg-go/dispatch/sdk"
	"github.com/gucooing/hkrpg-go/pkg/alg"
)

// 初始化所有服务
func NewServer(cfg *config.Config) *sdk.Server {
	s := &sdk.Server{}
	s.Config = cfg
	s.AppId = alg.GetAppId()
	logger.Info("Dispatch AppId:%s", s.AppId)
	port := s.Config.AppList[s.AppId].App["port_http"].Port
	if port == "" {
		panic("Dispatch Port error")
	}
	s.Port = port
	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		logger.Error("node error:", err)
		return nil
	}
	s.NodeConn = tcpConn
	go s.RecvNode()
	// 向node注册
	s.Connection()

	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	gin.SetMode(gin.ReleaseMode)    // 初始化gin
	s.Router = gin.Default()        // gin.New()
	s.Router.Use(gin.Recovery())
	cfg.Ec2b = alg.GetEc2b() // 读取ec2b密钥

	return s
}
