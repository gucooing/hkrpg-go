package service

import (
	"context"
	"encoding/base64"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/dispatch/sdk"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

type Dispatch struct {
	DiscoveryClient *rpc.NodeDiscoveryClient
	MessageQueue    *mq.MessageQueue
	Server          *sdk.Server
	AppId           uint32
	RegionName      string
	OuterPort       string
	OuterAddr       string
}

func NewDispatch(discoveryClient *rpc.NodeDiscoveryClient, messageQueue *mq.MessageQueue,
	netInfo constant.AppNet, appInfo constant.AppList, appId uint32) *Dispatch {
	d := &Dispatch{
		DiscoveryClient: discoveryClient,
		MessageQueue:    messageQueue,
		RegionName:      appInfo.RegionName,
		AppId:           appId,
		OuterAddr:       netInfo.OuterAddr,
		OuterPort:       netInfo.OuterPort,
		Server: &sdk.Server{
			RegionInfo:         make(map[string]*sdk.RegionInfo),
			UpstreamServer:     make(map[string]*sdk.UrlList),
			UpstreamServerLock: new(sync.RWMutex),
		},
	}
	d.getRegionInfo()
	go d.keepaliveServer()
	go d.messageQueue()

	return d
}

// mq收包
func (d *Dispatch) messageQueue() {
	for {
		netMsg := <-d.MessageQueue.GetNetMsg()
		switch netMsg.OriginServerType {
		case spb.ServerType_SERVICE_NODE:
			logger.Info("node mq msg:%s", base64.StdEncoding.EncodeToString(netMsg.ServiceMsgByte))
		default:
			logger.Info("error ServerType:%s", netMsg.OriginServerType.String())
		}
	}
}

// 心跳
func (d *Dispatch) keepaliveServer() {
	ticker := time.NewTicker(time.Second * 15)
	gateTicker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-ticker.C:
			rsp, err := d.DiscoveryClient.KeepaliveServer(context.TODO(), &nodeapi.KeepaliveServerReq{
				Type:       nodeapi.ServerType_SERVICE_DISPATCH,
				AppVersion: pkg.GetAppVersion(),
				RegionName: d.RegionName,
				AppId:      d.AppId,
				OuterPort:  d.OuterPort,
				OuterAddr:  d.OuterAddr,
				LoadCount:  0,
			})
			if err != nil {
				logger.Error("keepalive error: %v", err)
				continue
			}
			if rsp.RetCode == nodeapi.Retcode_RET_Reconnect {
				// TODO 代表是重连
			}
		case <-gateTicker.C:
			d.getRegionInfo()
		}
	}
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
		if d.MessageQueue.GetGateTcpMqInst(spb.ServerType_SERVICE_GATE, info.MinGateAppId) != nil {
			regionInfo.MinGateAddr = info.MinGateAddr
			regionInfo.MinGatePort = alg.S2U32(info.MinGatePort)
			regionInfo.MinGateTcp = info.MinGateTcp
		}
		ec2b, _ := random.LoadEc2bKey(info.ClientSecretKey)
		regionInfo.Ec2b = ec2b
		d.Server.RegionInfo[info.Name] = regionInfo
	}
}
