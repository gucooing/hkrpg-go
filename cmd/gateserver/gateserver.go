package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gucooing/hkrpg-go/gateserver"
	"github.com/gucooing/hkrpg-go/gateserver/app"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "gateserver.json"
	err := gateserver.LoadConfig(confName)
	if err != nil {
		if err == gateserver.FileNotExist {
			p, _ := json.MarshalIndent(gateserver.DefaultConfig, "", "  ")
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
	logger.InitLogger("gateserver"+"["+appid+"]", strings.ToUpper(gateserver.GetConfig().LogLevel))
	logger.Info("hkrpg-go")
	logger.Info("AppVersion:%s", pkg.GetAppVersion())
	logger.Info("GameVersion:%s", pkg.GetGameVersion())
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	cfg := gateserver.GetConfig()

	if err = app.Run(done, cfg, appid); err != nil {
		logger.Error(err.Error())
		logger.CloseLogger()
	}

}
