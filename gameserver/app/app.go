package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver"
	"github.com/gucooing/hkrpg-go/gameserver/service"
	"github.com/gucooing/hkrpg-go/gdconf"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func Run(done chan os.Signal, cfg *gameserver.Config, appid string) error {
	appInfo, ok := cfg.AppList[appid]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	netInfo, ok := appInfo.App["port_gt"]
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
		Type:       nodeapi.ServerType_SERVICE_GAME,
		AppVersion: pkg.GetAppVersion(),
		RegionName: appInfo.RegionName,
		// MqAddr:     appInfo.MqAddr,
		AppId: alg.GetAppIdUint32(appid),
		// OuterPort:  netInfo.OuterPort,
		// OuterAddr:  netInfo.OuterAddr,
	})
	if err != nil {
		return fmt.Errorf("register server error: %v ", err)
	}
	logger.Info("register server success")
	defer func() {
		_, _ = discoveryClient.CloseServer(context.TODO(), &nodeapi.CloseServerReq{
			Type:       nodeapi.ServerType_SERVICE_GAME,
			AppId:      alg.GetAppIdUint32(appid),
			RegionName: appInfo.RegionName,
		})
	}()
	// new mq
	nodeMq, ok := cfg.NetConf["NodeMq"]
	if !ok {
		return fmt.Errorf("NodeMq not exist")
	}
	messageQueue := mq.NewMessageQueue(spb.ServerType_SERVICE_GAME,
		alg.GetAppIdUint32(appid), discoveryClient, "", nodeMq, appInfo.RegionName)
	if messageQueue == nil {
		return fmt.Errorf("message queue nil")
	}
	defer messageQueue.Close()
	// new conf
	gdconf.InitGameDataConfig(cfg.GameDataConfigPath)
	// new db
	database.NewGameStore(cfg.MysqlConf, cfg.RedisConf)
	// new game
	g := service.NewGameServer(discoveryClient, messageQueue, netInfo, appInfo, alg.GetAppIdUint32(appid))
	if g == nil {
		return fmt.Errorf("create game server error")
	}

	logger.Info("game service start")

	// close
	select {
	case <-done:
		_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		logger.Info("game服务正在关闭")
		logger.Info("game服务已停止")
		logger.CloseLogger()
		return nil
	}
}
