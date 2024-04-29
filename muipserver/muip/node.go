package muip

import (
	"context"
	"net"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

type NodeService struct {
	s        *Muip
	Addr     string
	nodeConn net.Conn

	tickerCancel context.CancelFunc
	ticker       *time.Ticker // 定时器
}

func (s *Muip) newNode() {
	n := new(NodeService)
	n.s = s
	n.Addr = s.Config.NetConf["Node"]

	tcpConn, err := net.Dial("tcp", n.Addr)
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
	n.nodeTicler(tickerCtx)
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
			n.MuipToNodePingReq() // ping包
		case <-tickerCtx.Done():
			n.ticker.Stop()
			return
		}
	}
}

// 向node注册
func (n *NodeService) ServiceConnectionReq() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_MUIP,
		AppId:      n.s.AppId,
	}

	n.sendNode(cmd.ServiceConnectionReq, req)
}

// 从node接收消息
func (n *NodeService) recvNode() {
	nodeMsg := make([]byte, player.PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := n.nodeConn.Read(nodeMsg)
		if err != nil {
			logger.Error("node error")
			n.nodeKill()
			return
		}
		bin = nodeMsg[:recvLen]
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
	case cmd.MuipToNodePingRsp: // 心跳包
		n.MuipToNodePingRsp(serviceMsg)
	default:
		logger.Info("node -> muip error cmdid:%v", cmdId)
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
	if rsp.ServerType == spb.ServerType_SERVICE_MUIP && rsp.AppId == n.s.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (n *NodeService) MuipToNodePingReq() {
	req := &spb.MuipToNodePingReq{
		MuipServerTime: time.Now().UnixNano() / 1e6,
	}
	n.sendNode(cmd.MuipToNodePingReq, req)
}

func (n *NodeService) MuipToNodePingRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.MuipToNodePingRsp)
	if rsp.NodeServerTime-rsp.MuipServerTime > 5 {
		logger.Warn("muip <-> node 调用时间过长")
	}
	logger.Info("ping rsp msg :%s", protojson.Format(rsp))
}
