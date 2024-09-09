package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gucooing/hkrpg-go/gameserver"
	"github.com/gucooing/hkrpg-go/gameserver/app"
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
	if err = app.Run(done, cfg, appid); err != nil {
		logger.Error(err.Error())
		logger.CloseLogger()
	}
}
