package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/hkrpg-go/muipserver/config"
	"github.com/gucooing/hkrpg-go/muipserver/muip"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "muipserver.json"
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
	logger.InitLogger("muipserver"+"["+alg.GetAppId()+"]", strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")

	cfg := config.GetConfig()
	muips := muip.NewMuip(cfg)
	// 初始化

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动SDK服务
	go func() {
		if err = muips.Start(); err != nil {
			logger.Error("无法启动muipserver服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}
}
