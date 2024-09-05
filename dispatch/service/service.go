package service

import (
	"context"

	"github.com/gucooing/hkrpg-go/dispatch/sdk"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
)

type Dispatch struct {
	DiscoveryClient *rpc.NodeDiscoveryClient
	MessageQueue    *mq.MessageQueue
	Server          *sdk.Server
}

func NewDispatch(discoveryClient *rpc.NodeDiscoveryClient, messageQueue *mq.MessageQueue) *Dispatch {
	d := &Dispatch{
		DiscoveryClient: discoveryClient,
		MessageQueue:    messageQueue,
		Server:          new(sdk.Server),
	}
	d.getRegionInfo()

	return d
}

func (d *Dispatch) getRegionInfo() { // 拉取区服信息
	rsp, err := d.DiscoveryClient.GetAllRegionInfo(context.TODO(), &nodeapi.GetAllRegionInfoReq{})
	if err != nil {
		logger.Error("get all region info error:%s", err.Error())
		return
	}
	d.Server.RegionInfo = make(map[string]*sdk.RegionInfo)
	for _, info := range rsp.RegionInfoList {
		regionInfo := &sdk.RegionInfo{
			Name:  info.Name,
			Title: info.Title,
			Type:  info.Type,
		}
		ec2b, _ := random.LoadEc2bKey(info.ClientSecretKey)
		regionInfo.Ec2b = ec2b
		d.Server.RegionInfo[info.Name] = regionInfo
	}
}
