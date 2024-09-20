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
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	"google.golang.org/grpc"
)

func Run(done chan os.Signal, cfg *config.Config, appid string) error {
	// new db
	database.NewNodeStore(cfg.MysqlConf, cfg.RedisConf)
	appInfo, ok := cfg.AppList[appid]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	// new grpc
	lis, err := net.Listen("tcp", appInfo.GrpcAddr)
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

	logger.Info("grpc server start addr:%s", appInfo.GrpcAddr)
	defer grpcServer.GracefulStop()
	// new mq
	messageQueue := mq.NewMessageQueue(spb.ServerType_SERVICE_NODE,
		alg.GetAppIdUint32(appid), nil, appInfo.MqAddr, "", appInfo.RegionName)
	if messageQueue == nil {
		return fmt.Errorf("message queue nil")
	}
	node.MessageQueue = messageQueue
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
