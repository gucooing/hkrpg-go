package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/nodeserver/config"
	"github.com/gucooing/hkrpg-go/nodeserver/service"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"google.golang.org/grpc"
)

func Run(done chan os.Signal, cfg *config.Config, appid string) error {
	appInfo, ok := cfg.AppList[appid]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	netInfo, ok := appInfo.App["port_service"]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	// new grpc
	rpcAddr := fmt.Sprintf("%s:%s", netInfo.InnerAddr, netInfo.InnerPort)
	lis, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		return err
	}
	optss := make([]grpc.ServerOption, 0)
	grpcServer := grpc.NewServer(optss...)
	node := new(service.NodeDiscoveryService)
	nodeapi.RegisterNodeDiscoveryServer(grpcServer, node)
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			logger.Error("grpc server start error: %v", err)
			return
		}
	}()

	logger.Info("grpc server start addr:%s", rpcAddr)
	defer grpcServer.GracefulStop()
	// new db
	database.NewNodeStore(cfg.MysqlConf, cfg.RedisConf)
	// new node
	service.NewNodeService(node)
	logger.Info("node service start")

	// stop
	select {
	case <-done:
		_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		logger.Info("node server服务正在关闭")
		logger.Info("node server服务已停止")
		logger.CloseLogger()
		return nil
	}
}
