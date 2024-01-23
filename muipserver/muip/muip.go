package muip

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/muipserver/config"
	"github.com/gucooing/hkrpg-go/muipserver/logger"
	"github.com/gucooing/hkrpg-go/pkg/alg"
)

var MUIP *Muip

type Muip struct {
	Config   *config.Config
	AppId    string
	Port     string
	NodeConn net.Conn
	Router   *gin.Engine
}

func NewMuip(cfg *config.Config) *Muip {
	s := new(Muip)
	MUIP = s

	s.Config = cfg
	s.AppId = alg.GetAppId()
	logger.Info("MuipServer AppId:%s", s.AppId)
	port := s.Config.AppList[s.AppId].App["port_http"].Port
	if port == "" {
		panic("MuipServer Port error")
	}
	s.Port = port
	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		panic(err.Error())
		return nil
	}
	s.NodeConn = tcpConn
	go s.RecvNode()
	// 向node注册
	s.Connection()

	gin.SetMode(gin.ReleaseMode) // 初始化gin
	s.Router = gin.Default()     // gin.New()
	s.Router.Use(gin.Recovery())

	return s
}
