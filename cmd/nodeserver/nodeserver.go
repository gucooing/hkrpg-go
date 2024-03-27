package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gucooing/hkrpg-go/nodeserver/config"
	"github.com/gucooing/hkrpg-go/nodeserver/node"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	confName := "nodeserver.json"
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
	logger.InitLogger("nodeserver" + "[" + alg.GetAppId() + "]")
	logger.SetLogLevel(strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")
	cfg := config.GetConfig()

	// 初始化node
	s := node.NewNode(cfg)

	// 开启监听
	go s.NewNode()

	select {}
}
