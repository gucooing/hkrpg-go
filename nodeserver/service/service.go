package service

import (
	"context"
	"errors"
	"sync"
	"time"

	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"google.golang.org/grpc"
)

type NodeDiscoveryService struct {
	nodeapi.UnimplementedNodeDiscoveryServer
	regionMap     map[string]*RegionInfo
	regionMapLock sync.RWMutex
}

type RegionInfo struct {
	Name            string
	Title           string
	Type            uint32
	ClientSecretKey *random.Ec2b
	AutoCreate      bool
	serviceMap      map[nodeapi.ServerType]map[uint32]*Service // [ServerType][appid][Service]
	serviceMapLock  sync.RWMutex                               // 服务列表互斥锁
}

type Service struct {
	appId         uint32               // appid
	serverType    nodeapi.ServerType   // type
	regionName    string               // 区服
	mqAddr        string               // mq地址
	playerNum     int64                // 玩家数量
	lastAliveTime int64                // 最后一次ping时间
	outerPort     string               // 外网端口
	outerAddr     string               // 外网地址
	status        nodeapi.ServerStatus // 服务状态
}

func newMapService() map[nodeapi.ServerType]map[uint32]*Service {
	return map[nodeapi.ServerType]map[uint32]*Service{
		nodeapi.ServerType_SERVICE_DISPATCH: make(map[uint32]*Service),
		nodeapi.ServerType_SERVICE_GATE:     make(map[uint32]*Service),
		nodeapi.ServerType_SERVICE_GAME:     make(map[uint32]*Service),
		nodeapi.ServerType_SERVICE_MUIP:     make(map[uint32]*Service),
		nodeapi.ServerType_SERVICE_MULTI:    make(map[uint32]*Service),
	}
}

func NewNodeService(s *NodeDiscoveryService) {
	// 拉取全部区服配置
	regionConfList := database.GetAllRegionConf(database.NODE.ServerConf)
	s.regionMap = make(map[string]*RegionInfo)
	for _, regionConf := range regionConfList {
		ec2b := random.NewEc2b()
		regionConf.ClientSecretKey = ec2b.Bytes()
		err := database.SetRegionConf(database.NODE.ServerConf, regionConf)
		if err != nil {
			logger.Warn("更新区服配置失败:%s", err.Error())
		}
		info := &RegionInfo{
			Name:            regionConf.Name,
			Title:           regionConf.Title,
			Type:            regionConf.Type,
			ClientSecretKey: ec2b,
			AutoCreate:      regionConf.AutoCreate,
			serviceMap:      newMapService(),
		}
		s.regionMap[regionConf.Name] = info
	}
}

func (s *NodeDiscoveryService) GetRegionMap() map[string]*RegionInfo {
	if s.regionMap == nil {
		s.regionMap = make(map[string]*RegionInfo)
	}
	return s.regionMap
}

func (s *NodeDiscoveryService) GetRegion(regionName string) *RegionInfo {
	if s.regionMap == nil {
		s.regionMap = make(map[string]*RegionInfo)
	}
	if s.regionMap[regionName] == nil {
		s.regionMapLock.Lock()
		// 没有配置,那就去数据库拉取
		regionConf, err := database.GetRegionConf(database.NODE.ServerConf, regionName)
		if err != nil {
			logger.Error("拉取区服配置失败:%s", err.Error())
			s.regionMapLock.Unlock()
			return nil
		}
		ec2b := random.NewEc2b()
		regionConf.ClientSecretKey = ec2b.Bytes()
		err = database.SetRegionConf(database.NODE.ServerConf, regionConf)
		if err != nil {
			logger.Warn("更新区服配置失败:%s", err.Error())
		}
		info := &RegionInfo{
			Name:            regionConf.Name,
			Title:           regionConf.Title,
			Type:            regionConf.Type,
			ClientSecretKey: ec2b,
			AutoCreate:      regionConf.AutoCreate,
			serviceMap:      newMapService(),
		}
		s.regionMap[regionName] = info
		s.regionMapLock.Unlock()
	}
	s.regionMapLock.RLock()
	defer s.regionMapLock.RUnlock()
	return s.regionMap[regionName]
}

func (s *NodeDiscoveryService) AddService(info *Service) bool {
	list := s.GetRegion(info.regionName)
	if list == nil {
		return false
	}
	list.serviceMapLock.Lock()
	defer list.serviceMapLock.Unlock()
	if list.serviceMap == nil {
		list.serviceMap = make(map[nodeapi.ServerType]map[uint32]*Service)
	}
	if list.serviceMap[info.serverType] == nil {
		list.serviceMap[info.serverType] = make(map[uint32]*Service)
	}
	if _, ok := list.serviceMap[info.serverType][info.appId]; ok {
		return false
	}
	list.serviceMap[info.serverType][info.appId] = info
	return true
}

func (s *NodeDiscoveryService) DelService(info *Service) {
	list := s.GetRegion(info.regionName)
	if list == nil {
		return
	}
	list.serviceMapLock.Lock()
	defer list.serviceMapLock.Unlock()
	if list.serviceMap == nil {
		return
	}
	if list.serviceMap[info.serverType] == nil {
		return
	}
	if _, ok := list.serviceMap[info.serverType][info.appId]; ok {
		delete(list.serviceMap[info.serverType], info.appId)
	}
}

func (s *NodeDiscoveryService) GetService(regionName string, st nodeapi.ServerType, appId uint32) *Service {
	region := s.GetRegion(regionName)
	if region == nil {
		return nil
	}
	region.serviceMapLock.RLock()
	defer region.serviceMapLock.RUnlock()
	if region.serviceMap == nil {
		region.serviceMap = make(map[nodeapi.ServerType]map[uint32]*Service)
	}
	list := region.serviceMap[st]
	if list == nil {
		return nil
	}
	return list[appId]
}

func (s *NodeDiscoveryService) GetMinGate(regionName string) *Service {
	region := s.GetRegion(regionName)
	if region == nil {
		return nil
	}
	region.serviceMapLock.RLock()
	defer region.serviceMapLock.RUnlock()
	if region.serviceMap == nil {
		region.serviceMap = make(map[nodeapi.ServerType]map[uint32]*Service)
	}
	list := region.serviceMap[nodeapi.ServerType_SERVICE_GATE]
	if list == nil {
		return nil
	}
	var minNum int64 = 0
	var minAppid uint32 = 0
	for _, service := range list {
		if minAppid == 0 {
			minAppid = service.appId
			minNum = service.playerNum
			continue
		}
		if service.playerNum < minNum {
			minNum = service.playerNum
			minAppid = service.appId
		}
	}
	return list[minAppid]
}

func (s *NodeDiscoveryService) Test(ctx context.Context, req *nodeapi.TestReq) (*nodeapi.TestRsp, error) {
	return &nodeapi.TestRsp{
		ReqMsg: req.Msg,
	}, nil
}

// 向node注册服务
func (s *NodeDiscoveryService) RegisterServer(ctx context.Context, req *nodeapi.RegisterServerReq) (*nodeapi.RegisterServerRsp, error) {
	info := &Service{
		appId:         req.AppId,
		serverType:    req.Type,
		regionName:    req.RegionName,
		mqAddr:        req.MqAddr,
		playerNum:     0,
		lastAliveTime: time.Now().Unix(),
		outerPort:     req.OuterPort,
		outerAddr:     req.OuterAddr,
	}
	rsp := &nodeapi.RegisterServerRsp{}
	if s.AddService(info) {
		logger.Info("add service:%s regionName:%s appId:%v", req.Type, req.RegionName, req.AppId)
		return rsp, nil
	} else {
		logger.Error("add repeatedly service:%s regionName:%s appId:%v", req.Type, req.RegionName, req.AppId)
		return rsp, errors.New("add service failed")
	}
}

// 向node取消注册服务
func (s *NodeDiscoveryService) CloseServer(ctx context.Context, req *nodeapi.CloseServerReq) (*nodeapi.CloseServerRsp, error) {
	logger.Info("server cancel, service:%s regionName:%s appId:%v", req.Type, req.RegionName, req.AppId)
	info := &Service{
		appId:      req.AppId,
		serverType: req.Type,
		regionName: req.RegionName,
	}
	switch req.Type {
	case nodeapi.ServerType_SERVICE_GATE:
		// TODO 通知gs保存玩家数据
	case nodeapi.ServerType_SERVICE_GAME:
		// TODO 通知gate下线玩家
	}
	s.DelService(info)
	return &nodeapi.CloseServerRsp{}, nil
}

// 心跳
func (s *NodeDiscoveryService) KeepaliveServer(ctx context.Context, req *nodeapi.KeepaliveServerReq) (*nodeapi.KeepaliveServerRsp, error) {
	info := &Service{
		appId:         req.AppId,
		serverType:    req.Type,
		regionName:    req.RegionName,
		mqAddr:        req.MqAddr,
		playerNum:     req.LoadCount,
		lastAliveTime: time.Now().Unix(),
		outerPort:     req.OuterPort,
		outerAddr:     req.OuterAddr,
	}
	rsp := &nodeapi.KeepaliveServerRsp{}
	if s.AddService(info) {
		logger.Info("reconnect service:%s regionName:%s appId:%v", req.Type, req.RegionName, req.AppId)
		rsp.RetCode = nodeapi.Retcode_RET_Reconnect
	}
	return rsp, nil
}

// 获取全部gate的消息队列
func (s *NodeDiscoveryService) GetAllGateServerMq(ctx context.Context, req *nodeapi.GetAllGateServerMqReq) (*nodeapi.GetAllGateServerMqRsp, error) {
	rsp := &nodeapi.GetAllGateServerMqRsp{
		ServerList: make([]*nodeapi.GateServerMq, 0),
	}
	region := s.GetRegion(req.RegionName)
	region.serviceMapLock.RLock()
	if region.serviceMap != nil {
		for _, gate := range region.serviceMap[nodeapi.ServerType_SERVICE_GATE] {
			rsp.ServerList = append(rsp.ServerList, &nodeapi.GateServerMq{
				AppId:  gate.appId,
				MqAddr: gate.mqAddr,
			})
		}
	}
	region.serviceMapLock.RUnlock()
	return rsp, nil
}

// 获取区服配置
func (s *NodeDiscoveryService) GetAllRegionInfo(ctx context.Context, req *nodeapi.GetAllRegionInfoReq) (*nodeapi.GetAllRegionInfoRsp, error) {
	rsp := &nodeapi.GetAllRegionInfoRsp{
		RegionInfoList: make(map[string]*nodeapi.RegionInfo),
	}
	regionMap := s.GetRegionMap()
	s.regionMapLock.RLock()
	for name, region := range regionMap {
		minGate := s.GetMinGate(name)
		info := &nodeapi.RegionInfo{
			Name:            region.Name,
			Title:           region.Title,
			Type:            region.Type,
			ClientSecretKey: region.ClientSecretKey.Bytes(),
			AutoCreate:      region.AutoCreate,
		}
		if minGate != nil {
			info.MinGateAddr = minGate.outerAddr
			info.MinGatePort = minGate.outerPort
			info.MinGateAppId = minGate.appId
		}
		rsp.RegionInfoList[name] = info
	}
	s.regionMapLock.RUnlock()
	return rsp, nil
}

// 持续接收来自node的消息
func (s *NodeDiscoveryService) NodeStreamMessages(req *nodeapi.NodeStreamMessagesReq, stream grpc.ServerStreamingServer[nodeapi.NodeStreamMessagesRsp]) error {
	// for {
	// 	err := stream.Send(&nodeapi.NodeStreamMessagesRsp{})
	// 	if err != nil {
	// 		logger.Warn("Send error:%v", err)
	// 		return err
	// 	}
	// }
	return nil
}
