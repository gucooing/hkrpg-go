package rpc

import (
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"google.golang.org/grpc"
)

// node
type NodeDiscoveryClient struct {
	nodeapi.NodeDiscoveryClient
}

func NewNodeRpcClient(grpcAddr string) (*NodeDiscoveryClient, error) {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	discoveryClient, err := newNodeDiscoveryClient(conn)
	if err != nil {
		return nil, err
	}
	return discoveryClient, nil
}

func newNodeDiscoveryClient(conn *grpc.ClientConn) (*NodeDiscoveryClient, error) {
	cl := nodeapi.NewNodeDiscoveryClient(conn)
	return &NodeDiscoveryClient{cl}, nil
}

// gm
