package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver"
	"github.com/gucooing/hkrpg-go/gateserver/service"
	"github.com/gucooing/hkrpg-go/gateserver/session"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func Run(done chan os.Signal, cfg *gateserver.Config, appid string) error {
	appInfo, ok := cfg.AppList[appid]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	netInfo, ok := appInfo.App["port_player"]
	if !ok {
		return fmt.Errorf("app not exist")
	}
	// new grpc
	nodeAddr, ok := cfg.NetConf["Node"]
	if !ok {
		return fmt.Errorf("node not exist")
	}
	discoveryClient, err := rpc.NewNodeRpcClient(nodeAddr)
	if err != nil {
		return err
	}
	// 注册到node
	_, err = discoveryClient.RegisterServer(context.TODO(), &nodeapi.RegisterServerReq{
		Type:       nodeapi.ServerType_SERVICE_GATE,
		AppVersion: pkg.GetAppVersion(),
		RegionName: appInfo.RegionName,
		MqAddr:     appInfo.MqAddr,
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
			Type:       nodeapi.ServerType_SERVICE_GATE,
			AppId:      alg.GetAppIdUint32(appid),
			RegionName: appInfo.RegionName,
		})
	}()
	// new mq
	messageQueue := mq.NewMessageQueue(spb.ServerType_SERVICE_GATE,
		alg.GetAppIdUint32(appid), discoveryClient, appInfo.MqAddr, appInfo.RegionName)
	if messageQueue == nil {
		return fmt.Errorf("message queue nil")
	}
	defer messageQueue.Close()
	// new db
	database.NewGateStore(cfg.MysqlConf, cfg.RedisConf)
	// new gate
	g := service.NewGateServer(discoveryClient, messageQueue, netInfo, appInfo, alg.GetAppIdUint32(appid))
	if g == nil {
		return fmt.Errorf("gateserver nil")
	}
	session.MAX_CLIENT__CONN_NUM = cfg.MaxPlayer
	// run kcp
	go func() {
		err = g.KcpConn.RunKcp()
		if err != nil {
			logger.Error("kcp conn err: %v", err)
			return
		}
	}()

	logger.Info("gate service start")

	// close
	select {
	case <-done:
		_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		logger.Info("gate服务正在关闭")
		g.KcpConn.Close()
		logger.Info("gate服务已停止")
		logger.CloseLogger()
		return nil
	}
}
