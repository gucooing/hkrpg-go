package muip

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/muipserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var MUIP *Muip

type Muip struct {
	Config     *config.Config
	AppId      string
	Port       string
	NodeConn   net.Conn
	Router     *gin.Engine
	AllService map[string][]*AllService
}

type AllService struct {
	AppId     string
	PlayerNum uint64
}

func NewMuip(cfg *config.Config) *Muip {
	s := new(Muip)
	MUIP = s

	s.Config = cfg
	s.AppId = alg.GetAppId()
	logger.Info("MuipServer AppId:%s", s.AppId)
	port := s.Config.AppList[s.AppId].App["port_http"].Port
	if port == "" {
		log.Println("MuipServer Port error")
		os.Exit(0)
	}
	s.Port = port
	// 连接node
	tcpConn, err := net.Dial("tcp", cfg.NetConf["Node"])
	if err != nil {
		log.Println("nodeserver error")
		os.Exit(0)
	}
	s.NodeConn = tcpConn
	s.AllService = make(map[string][]*AllService)
	go s.RecvNode()
	// 向node注册
	s.Connection()

	gin.SetMode(gin.ReleaseMode) // 初始化gin
	s.Router = gin.Default()     // gin.New()
	s.Router.Use(gin.Recovery())

	return s
}
