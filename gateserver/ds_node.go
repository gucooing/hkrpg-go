package gateserver

import (
	"context"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

type NodeService struct {
	gate     *GateServer
	nodeConn *gunet.TcpConn

	tickerCancel context.CancelFunc
	ticker       *time.Ticker // 定时器
}

func (s *GateServer) sendNode(cmdId uint16, playerMsg pb.Message) {
	node := s.node
	if node == nil {
		return
	}
	node.sendNode(cmdId, playerMsg)
}

func (s *GateServer) newNode() {
	n := new(NodeService)
	tcpConn, err := gunet.NewTcpC(s.Config.NetConf["Node"])
	if err != nil {
		logger.Error("nodeserver error:%s", err.Error())
		return
	}
	n.nodeConn = tcpConn
	n.gate = s
	n.ticker = time.NewTicker(5 * time.Second)
	tickerCtx, tickerCancel := context.WithCancel(context.Background())
	n.tickerCancel = tickerCancel
	s.node = n
	go n.recvNode()
	// 向node注册
	n.ServiceConnectionReq()
	// 开启node定时器
	go n.nodeTicler(tickerCtx)
}

func (n *NodeService) nodeKill() {
	n.nodeConn.Close()
	n.tickerCancel()
	logger.Info("node server离线")
	n.gate.node = nil
}

func (n *NodeService) nodeTicler(tickerCtx context.Context) {
	for {
		select {
		case <-n.ticker.C:
			n.GateToNodePingReq() // ping包
		case <-tickerCtx.Done():
			n.ticker.Stop()
			return
		}
	}
}

// 向node注册
func (n *NodeService) ServiceConnectionReq() {
	// req := &spb.ServiceConnectionReq{
	// 	ServerType: spb.ServerType_SERVICE_GATE,
	// 	AppId:      n.gate.AppId,
	// 	Addr:       n.gate.OuterAddr,
	// 	Port:       n.gate.Port,
	// }
	//
	// n.sendNode(cmd.ServiceConnectionReq, req)
}

// 从node接收消息
func (n *NodeService) recvNode() {
	for {
		bin, err := n.nodeConn.Read()
		if err != nil {
			logger.Error("node error")
			n.nodeKill()
			return
		}
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := cmd.DecodePayloadToProto(msg)
			n.nodeRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (n *NodeService) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	// case cmd.ServiceConnectionRsp:
	// 	n.ServiceConnectionRsp(serviceMsg) // 注册包
	// case cmd.GateToNodePingRsp:
	// 	n.GateToNodePingRsp(serviceMsg) // 心跳包
	default:
		logger.Info("nodeRegister error cmdid:%v", cmdId)
	}
}

// 发送到node
func (n *NodeService) sendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(cmd.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := cmd.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdId error")
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := n.nodeConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

func (n *NodeService) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_GATE && rsp.AppId == n.gate.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (n *NodeService) GateToNodePingReq() {
	// // 心跳包
	// req := &spb.GateToNodePingReq{
	// 	ServiceType: spb.ServerType_SERVICE_GATE,
	// 	GateTime:    time.Now().UnixNano() / 1e6,
	// 	PlayerNum:   int64(CLIENT_CONN_NUM),
	// }
	// n.sendNode(cmd.GateToNodePingReq, req)
}

func (n *NodeService) GateToNodePingRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GateToNodePingRsp)
	for _, service := range rsp.GameServiceList {
		if service.Addr == "" || service.AppId == 0 || service.ServiceType != spb.ServerType_SERVICE_GAME {
			continue
		}
		if n.gate.getGsByAppid(service.AppId) == nil {
			logger.Info("[AppId:%v]发现新的gameserver接入,申请连接中", service.AppId)
			addr := service.Addr + ":" + service.Port
			n.gate.newGs(addr, service.AppId)
		}
	}

	logger.Debug("gate <--> node ping:%v", (rsp.NodeTime-rsp.GateTime)/2)
}
