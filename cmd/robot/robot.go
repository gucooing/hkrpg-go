package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/robot/config"
	"github.com/gucooing/hkrpg-go/robot/gdconf"
	"github.com/gucooing/hkrpg-go/robot/robot"
	"golang.org/x/net/context"
)

func main() {
	// 启动读取配置
	confName := "robot.json"
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
	logger.InitLogger("robot" + "[" + alg.GetAppId() + "]")
	logger.SetLogLevel(strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-robot-go")

	gdconf.InitGameDataConfig()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		robot.NewBot()
	}()

	go robot.KcpNetInfo()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			logger.Info("服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()

	select {}
}
