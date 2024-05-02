package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/goccy/go-json"
	"github.com/gucooing/hkrpg-go/nodeserver/config"
	"github.com/gucooing/hkrpg-go/nodeserver/db"
	"github.com/gucooing/hkrpg-go/nodeserver/node"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "nodeserver.json"
	err := config.LoadConfig(confName)
	if err != nil {
		if err == config.FileNotExist {
			p, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
			cf, _ := os.Create("./conf/" + confName)
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 %s \n", confName)
			main()
		} else {
			panic(err)
		}
	}
	// 初始化日志
	appid := alg.GetAppId()
	logger.InitLogger("nodeserver"+"["+appid+"]", strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")
	cfg := config.GetConfig()

	// 初始化数据库
	dbs := db.NewStore(cfg)
	// 初始化node
	s := node.NewNode(cfg, appid, dbs)

	// 开启监听
	go s.NewNode()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			logger.Info("NodeServer 正在关闭")
			s.Close()
			logger.Info("NodeServer 服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()

	select {}
}
