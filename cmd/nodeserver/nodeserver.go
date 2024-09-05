package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/goccy/go-json"
	"github.com/gucooing/hkrpg-go/nodeserver/app"
	"github.com/gucooing/hkrpg-go/nodeserver/config"
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

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	if err = app.Run(done, cfg, appid); err != nil {
		logger.Error("node server error:%s", err.Error())
		logger.CloseLogger()
	}
}
