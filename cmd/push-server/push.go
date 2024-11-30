package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/push"
	"github.com/gucooing/hkrpg-go/pkg/push/server"
)

func main() {
	confName := "push.json"
	err := server.LoadConfig(confName)
	if err != nil {
		if err == server.FileNotExist {
			p, _ := json.MarshalIndent(server.DefaultConfig, "", "  ")
			cf, _ := os.Create("./conf/" + confName)
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 %s \n", confName)
			return
		} else {
			panic(err)
		}
	}
	cfg := server.GetConfig()
	logger.InitLogger("push", strings.ToUpper(cfg.LogLevel))
	logger.Info("push server AppVersion:%s", push.AppVersion)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if err = server.NewServer(done, cfg); err != nil {
		logger.Error(err.Error())
		logger.CloseLogger()
	}
}
