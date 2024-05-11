package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/gameserver/gs"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "gameserver.json"
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
	logger.InitLogger("gameserver"+"["+appid+"]", strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	cfg := config.GetConfig()
	// 初始化game
	gameserver := gs.NewGameServer(cfg, appid)
	if gameserver == nil {
		logger.Error("game初始化失败")
		return
	}

	go func() {
		http.ListenAndServe("0.0.0.0:9991", nil)
	}()

	// 加载res
	gdconf.InitGameDataConfig()

	// 启动game
	go func() {
		if err = gameserver.StartGameServer(); err != nil {
			logger.Error("无法启动game服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			logger.Info("game服务正在关闭")
			if err = gameserver.Close(); err != nil {
				logger.Error("无法正常关闭game服务")
			}
			logger.Info("game服务已停止")

			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}

}
