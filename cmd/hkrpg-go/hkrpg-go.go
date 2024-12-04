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

	"github.com/gucooing/hkrpg-go/hkrpg-go-pe"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/lua"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/gucooing/hkrpg-go/suppl"
)

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe(":6060", nil))
	// }()
	confName := "hkrpg-go-pe.json"
	err := hkrpg_go_pe.LoadConfig(confName)
	if err != nil {
		if err == hkrpg_go_pe.FileNotExist {
			fmt.Printf("找不到配置文件准备生成默认配置文件 %s \n", confName)
			p, _ := json.MarshalIndent(hkrpg_go_pe.DefaultConfig, "", "  ")
			cf, _ := os.Create("./conf/" + confName)
			_, err := cf.Write(p)
			cf.Close()
			if err != nil {
				fmt.Printf("生成默认配置文件失败 %s \n使用默认配置\n", err.Error())
				hkrpg_go_pe.SetDefaultConfig()
			} else {
				fmt.Printf("生成默认配置文件成功 \n")
				main()
			}
		} else {
			panic(err)
		}
	}
	cfg := hkrpg_go_pe.GetConfig()
	// 初始化日志
	logger.InitLogger("hkrpg_go_pe", strings.ToUpper(cfg.LogLevel))
	logger.Info("hkrpg_go_pe")
	// 初始化语言文件
	text.NewTextMap(cfg.Language, cfg.DataPrefix)
	logger.Info(text.GetText(1), pkg.GetAppVersion())
	logger.Info(text.GetText(2), pkg.GetGameVersion())
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// 启动附加组件
	suppl.Start()
	// 初始化lua
	lua.LoadLua(cfg.GameServer.LoadLua)
	// 初始化服务器
	s := hkrpg_go_pe.NewServer(cfg)

	go func() {
		select {
		case <-done:
			_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			logger.Info(text.GetText(3))
			s.Close()
			logger.Info(text.GetText(4))
			logger.CloseLogger()
			os.Exit(0)
		}
	}()

	select {}
}
