package service

import (
	"context"

	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
)

type NodeDiscoveryService struct {
	nodeapi.UnimplementedNodeDiscoveryServer
}

func (s *NodeDiscoveryService) Test(ctx context.Context, req *nodeapi.TestReq) (*nodeapi.TestRsp, error) {
	return &nodeapi.TestRsp{
		ReqMsg: req.Msg,
	}, nil
}

// 获取全部gate的消息队列
func (s *NodeDiscoveryService) GetAllGateServerMq(ctx context.Context, req *nodeapi.GetAllGateServerMqReq) (*nodeapi.GetAllGateServerMqRsp, error) {
	return &nodeapi.GetAllGateServerMqRsp{
		ServerList: make([]*nodeapi.GateServerMq, 0),
	}, nil
}
