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
	AppId      uint32
	Port       string
	NodeConn   net.Conn
	Router     *gin.Engine
	AllService map[string][]*AllService
}

type AllService struct {
	AppId     uint32
	PlayerNum uint64
}

func NewMuip(cfg *config.Config, appid string) *Muip {
	s := new(Muip)
	MUIP = s

	s.Config = cfg
	s.AppId = alg.GetAppIdUint32(appid)
	logger.Info("MuipServer AppId:%s", appid)
	port := s.Config.AppList[appid].App["port_http"].Port
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
