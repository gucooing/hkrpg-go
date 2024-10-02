package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gucooing/hkrpg-go/dispatch"
	"github.com/gucooing/hkrpg-go/dispatch/app"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "dispatch.json"
	err := dispatch.LoadConfig(confName)
	if err != nil {
		if err == dispatch.FileNotExist {
			p, _ := json.MarshalIndent(dispatch.DefaultConfig, "", "  ")
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
	logger.InitLogger("discord"+"["+appid+"]", strings.ToUpper(dispatch.GetConfig().LogLevel))
	logger.Info("hkrpg-go")
	logger.Info("AppVersion:%s", pkg.GetAppVersion())
	logger.Info("GameVersion:%s", pkg.GetGameVersion())
	cfg := dispatch.GetConfig()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if err = app.Run(done, cfg, appid); err != nil {
		logger.Error(err.Error())
		logger.CloseLogger()
	}
}
