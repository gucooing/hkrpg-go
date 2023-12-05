package internal

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/internal/SDK"
	"github.com/gucooing/hkrpg-go/pkg/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

// 初始化所有服务
func NewServer(cfg *config.Config) *SDK.Server {
	s := &SDK.Server{}
	s.Config = cfg
	s.Store = DataBase.NewStore(s.Config) // 初始化数据库连接
	gin.SetMode(gin.ReleaseMode)          // 初始化gin
	s.Router = gin.Default()              // gin.New()
	s.Router.Use(gin.Recovery())
	cfg.Ec2b = getEc2b() // 读取ec2b密钥

	return s
}

func getEc2b() *random.Ec2b {
	open, err := os.Open("data/Ec2b.bin")
	defer open.Close()
	if err != nil {
		ec2p := random.NewEc2b().Bytes()
		ioutil.WriteFile("data/Ec2b.bin", ec2p, 0644)
		logger.Info("ec2b不存在,生成ec2b文件中")
		ec2b, err := random.LoadEc2bKey(ec2p)
		if err != nil {
			logger.Error("parse region ec2b error: %v", err)
			return nil
		}
		return ec2b
	} else {
		ec2p, err := io.ReadAll(open)
		if err != nil {
			logger.Error("read Ec2b error")
			return nil
		}
		defer open.Close()
		ec2b, err := random.LoadEc2bKey(ec2p)
		if err != nil {
			logger.Error("parse region ec2b error: %v", err)
			return nil
		}
		return ec2b
	}
}
