package app

import (
	"context"
	"fmt"
	"net"

	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/nodeserver/config"
	"github.com/gucooing/hkrpg-go/nodeserver/service"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, cfg *config.Config, appid string) error {
	appInfi, ok := cfg.AppList[appid]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	netInfo, ok := appInfi.App["port_service"]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	// new grpc
	rpcAddr := fmt.Sprintf("%s:%s", netInfo.OuterAddr, netInfo.Port)
	lis, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		return err
	}
	var opts grpc.ServerOption
	grpcServer := grpc.NewServer(opts)
	nodeapi.RegisterNodeDiscoveryServer(grpcServer, new(service.NodeDiscoveryService))
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	logger.Info("grpc server start!:%s", rpcAddr)
	defer grpcServer.GracefulStop()
	// new db
	database.NewNodeStore(cfg.MysqlConf, cfg.RedisConf)

	// stop
	select {
	case <-ctx.Done():
		return nil
	}
}
