package discord

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/discord/config"
	"github.com/gucooing/hkrpg-go/discord/db"
	"github.com/gucooing/hkrpg-go/discord/sdk"
	"github.com/gucooing/hkrpg-go/pkg/alg"
)

// 初始化所有服务
func NewServer(cfg *config.Config) *sdk.Server {
	s := &sdk.Server{}
	s.Config = cfg
	s.Store = db.NewStore(s.Config) // 初始化数据库连接
	gin.SetMode(gin.ReleaseMode)    // 初始化gin
	s.Router = gin.New()            // gin.Default()
	s.Router.Use(gin.Recovery())
	cfg.Ec2b = alg.GetEc2b() // 读取ec2b密钥

	return s
}
