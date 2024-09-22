package server

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
	cfg    *Config
}

func NewServer(done chan os.Signal, cfg *Config) error {
	s := new(Server)
	s.cfg = cfg
	// new db
	db := cfg.Db
	switch db.Type {
	case "mysql":
		s.db = database.NewMysql(db.Host)
	case "sqlite":
		s.db = database.NewSqlite(db.Host)
	default:
		return errors.New("error db type:" + db.Type)
	}
	database.NewPush(s.db)
	// new http server
	gin.SetMode(gin.ReleaseMode)
	s.router = gin.Default() // gin.New()
	s.router.Use(gin.Recovery())
	s.initRouter()

	logger.Info("push server start!")

	go func() {
		err := alg.NewHttp(constant.AppNet{
			InnerAddr: cfg.Host,
			InnerPort: cfg.Port,
			OuterAddr: cfg.Host,
			OuterPort: cfg.Port,
		}, s.router)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	}()

	select {
	case <-done:
		_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		logger.Info("push正在关闭")
		logger.Info("push已停止")
		logger.CloseLogger()
		return nil
	}
}
