package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/dispatch"
	"github.com/gucooing/hkrpg-go/dispatch/service"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func Run(done chan os.Signal, cfg *dispatch.Config, appid string) error {
	// new db
	database.NewDisaptchStore(cfg.MysqlConf, cfg.RedisConf)
	appInfo, ok := cfg.AppList[appid]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	netInfo, ok := appInfo.App["port_http"]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	// new grpc
	nodeGrpc, ok := cfg.NetConf["NodeGrpc"]
	if !ok {
		return fmt.Errorf("NodeGrpc not exist")
	}
	discoveryClient, err := rpc.NewNodeRpcClient(nodeGrpc)
	if err != nil {
		return err
	}
	// 注册到node
	_, err = discoveryClient.RegisterServer(context.TODO(), &nodeapi.RegisterServerReq{
		Type:       nodeapi.ServerType_SERVICE_DISPATCH,
		AppVersion: pkg.GetAppVersion(),
		RegionName: appInfo.RegionName,
		AppId:      alg.GetAppIdUint32(appid),
		OuterPort:  netInfo.OuterPort,
		OuterAddr:  netInfo.OuterAddr,
	})
	if err != nil {
		return fmt.Errorf("register server error: %v ", err)
	}
	logger.Info("register server success")
	defer func() {
		_, _ = discoveryClient.CloseServer(context.TODO(), &nodeapi.CloseServerReq{
			Type:       nodeapi.ServerType_SERVICE_DISPATCH,
			AppId:      alg.GetAppIdUint32(appid),
			RegionName: appInfo.RegionName,
		})
	}()
	// new mq
	nodeMq, ok := cfg.NetConf["NodeMq"]
	if !ok {
		return fmt.Errorf("NodeMq not exist")
	}
	messageQueue := mq.NewMessageQueue(spb.ServerType_SERVICE_DISPATCH,
		alg.GetAppIdUint32(appid), discoveryClient, "", nodeMq, appInfo.RegionName)
	if messageQueue == nil {
		return fmt.Errorf("message queue nil")
	}
	defer messageQueue.Close()
	// new sdk
	d := service.NewDispatch(discoveryClient, messageQueue, netInfo, appInfo, alg.GetAppIdUint32(appid))
	d.Server.IsAutoCreate = cfg.AutoCreate
	d.Server.OuterAddr = fmt.Sprintf("http://%s:%s", netInfo.OuterAddr, netInfo.OuterPort)
	d.Server.UpstreamServerList = cfg.UpstreamServerList
	gin.SetMode(gin.ReleaseMode) // 初始化gin
	sdkRouter := gin.New()
	sdkRouter.Use(gin.Recovery())
	d.Server.GetSdkRouter(sdkRouter) // 初始化路由
	go d.Server.UpUpstreamServer()
	go func() {
		err = alg.NewHttp(netInfo, sdkRouter)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	}()
	logger.Info("dispatch service start")

	// close
	select {
	case <-done:
		_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		logger.Info("dispatch服务正在关闭")
		logger.Info("dispatch服务已停止")
		logger.CloseLogger()
		return nil
	}
}
