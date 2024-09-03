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

	hkrpggo "github.com/gucooing/hkrpg-go/hkrpg-go"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	confName := "hkrpg-go-pe.json"
	err := hkrpggo.LoadConfig(confName)
	if err != nil {
		if err == hkrpggo.FileNotExist {
			p, _ := json.MarshalIndent(hkrpggo.DefaultConfig, "", "  ")
			cf, _ := os.Create("./conf/" + confName)
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 %s \n", confName)
			main()
		} else {
			panic(err)
		}
	}
	cfg := hkrpggo.GetConfig()
	// 初始化日志
	logger.InitLogger("hkrpg_go_pe", strings.ToUpper(cfg.LogLevel))
	logger.Info("hkrpg_go_pe")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// 初始化服务器
	s := hkrpggo.NewServer(cfg)

	// 启动SDK服务
	go func() {
		if err = s.Dispatch.Start(); err != nil {
			logger.Error("无法启动dispatch服务器")
		}
	}()
	// 启动GameServer服务
	go func() {
		if err = s.RunGameServer(); err != nil {
			logger.Error("无法启动gameserver服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			logger.Info("hkrpg_go_pe服务正在关闭")
			s.Close()
			logger.Info("hkrpg_go_pe服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()

	select {}
}
