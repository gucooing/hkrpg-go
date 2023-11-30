package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/internal/SDK"
	"github.com/gucooing/hkrpg-go/pkg/config"
)

// 初始化所有服务
func NewServer(cfg *config.Config) *SDK.Server {
	s := &SDK.Server{}
	s.Config = cfg
	s.Store = DataBase.NewStore(s.Config) // 初始化数据库连接
	gin.SetMode(gin.ReleaseMode)          // 初始化gin
	s.Router = gin.Default()              // gin.New()
	s.Router.Use(gin.Recovery())

	return s
}
