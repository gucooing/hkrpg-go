package multi

import (
	"context"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type NodeService struct {
	s        *Multi
	Addr     string
	nodeConn *gunet.TcpConn

	tickerCancel context.CancelFunc
	ticker       *time.Ticker // 定时器
}

func (s *Multi) newNode() {
	n := new(NodeService)
	n.s = s
	n.Addr = s.Config.NetConf["Node"]

	tcpConn, err := gunet.NewTcpC(n.Addr)
	if err != nil {
		logger.Error("nodeserver error:%s", err.Error())
		return
	}
	n.nodeConn = tcpConn
	n.ticker = time.NewTicker(5 * time.Second)
	tickerCtx, tickerCancel := context.WithCancel(context.Background())
	n.tickerCancel = tickerCancel

	go n.recvNode()
	// 向node注册
	n.ServiceConnectionReq()
	// 开启node定时器
	go n.nodeTicler(tickerCtx)
}

func (n *NodeService) nodeKill() {
	n.nodeConn.Close()
	if n.tickerCancel != nil {
		n.tickerCancel()
	}
	logger.Info("node server离线")
	n.s.Node = nil
}

func (n *NodeService) nodeTicler(tickerCtx context.Context) {
	for {
		select {
		case <-n.ticker.C:
			n.MultiToNodePingReq() // ping包
		case <-tickerCtx.Done():
			n.ticker.Stop()
			return
		}
	}
}

// 向node注册
func (n *NodeService) ServiceConnectionReq() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_MULTI,
		AppId:      n.s.AppId,
		Addr:       n.s.addr,
		Port:       n.s.port,
	}

	n.sendNode(cmd.ServiceConnectionReq, req)
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
			serviceMsg := alg.DecodePayloadToProto(msg)
			n.nodeRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (n *NodeService) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp: // 注册回包
		n.ServiceConnectionRsp(serviceMsg)
	case cmd.MultiToNodePingRsp: // 心跳包
		n.MultiToNodePingRsp(serviceMsg)
	default:
		logger.Info("node -> multi error cmdid:%v", cmdId)
	}
}

// 发送到node
func (n *NodeService) sendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
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
	if rsp.ServerType == spb.ServerType_SERVICE_MULTI && rsp.AppId == n.s.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (n *NodeService) MultiToNodePingReq() {
	req := &spb.MultiToNodePingReq{
		MultiServerTime: time.Now().UnixNano() / 1e6,
	}
	n.sendNode(cmd.MultiToNodePingReq, req)
}

func (n *NodeService) MultiToNodePingRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.MultiToNodePingRsp)
	if rsp.NodeServerTime-rsp.MultiServerTime > 5 {
		logger.Warn("multi <-> node 调用时间过长")
	}
}
