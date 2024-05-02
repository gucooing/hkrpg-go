package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/gateserver/gate"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "gateserver.json"
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
	appid := alg.GetAppId()
	// 初始化日志
	logger.InitLogger("gateserver"+"["+appid+"]", strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	cfg := config.GetConfig()
	// 初始化geta
	gaten := gate.NewGate(cfg, appid)

	// 启动gate服务
	go func() {
		if err = gaten.RunKcp(); err != nil {
			logger.Error("无法启动kcp服务")
		}
	}()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			logger.Info("geta服务正在关闭")
			if err = gate.Close(); err != nil {
				logger.Error("无法正常关闭geta服务")
			}
			logger.Info("geta服务已停止")

			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}

}
