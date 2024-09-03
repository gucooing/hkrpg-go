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

	"github.com/gucooing/hkrpg-go/gameserver"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "gameserver.json"
	err := gameserver.LoadConfig(confName)
	if err != nil {
		if err == gameserver.FileNotExist {
			p, _ := json.MarshalIndent(gameserver.DefaultConfig, "", "  ")
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
	cfg := gameserver.GetConfig()
	// 初始化日志
	logger.InitLogger("gameserver"+"["+appid+"]", strings.ToUpper(cfg.LogLevel))
	logger.Info("hkrpg-go")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// 初始化game
	gs := gameserver.NewGameServer(cfg, appid)
	if gs == nil {
		logger.Error("game初始化失败")
		return
	}

	go func() {
		http.ListenAndServe("0.0.0.0:9991", nil)
	}()

	// 加载res
	gdconf.InitGameDataConfig(gs.Config.GameDataConfigPath)

	// 启动game
	go func() {
		if err = gs.StartGameServer(); err != nil {
			logger.Error("无法启动game服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			logger.Info("game服务正在关闭")
			if err = gs.Close(); err != nil {
				logger.Error("无法正常关闭game服务")
			}
			logger.Info("game服务已停止")

			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}

}
