package sdk

import (
	"context"
	"sync"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type NodeService struct {
	s            *Server
	gateList     map[uint32]*Gate
	gateListLock sync.Mutex // gate列表互斥锁
	nodeConn     *gunet.TcpConn
	tickerCancel context.CancelFunc
	ticker       *time.Ticker // 定时器
}

// TODO 结构上要考虑熔断设计（综合参考成功率，流量，人数，系统状态
type Gate struct {
	Ip   string
	Port uint32
	Num  int64
}

func (s *Server) newNode() {
	n := new(NodeService)
	tcpConn, err := gunet.NewTcpC(s.Config.NetConf["Node"])
	if err != nil {
		logger.Error("nodeserver error:%s", err.Error())
		return
	}
	n.nodeConn = tcpConn
	n.s = s
	n.gateList = make(map[uint32]*Gate)
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
	n.s.node = nil
}

func (n *NodeService) nodeTicler(tickerCtx context.Context) {
	for {
		select {
		case <-n.ticker.C:
			n.getAllServiceGateReq() // ping包
		case <-tickerCtx.Done():
			n.ticker.Stop()
			return
		}
	}
}

// 从node接收消息
func (n *NodeService) recvNode() {
	for {
		var bin []byte = nil
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
	case cmd.GetAllServiceGateRsp: // 心跳包
		n.GetAllServiceGateRsp(serviceMsg)
	default:
		logger.Info("node -> dispatch error cmdid:%v", cmdId)
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

// 向node注册
func (n *NodeService) ServiceConnectionReq() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_DISPATCH,
		AppId:      n.s.AppId,
		Addr:       n.s.OuterAddr,
		Port:       n.s.Port,
	}
	n.sendNode(cmd.ServiceConnectionReq, req)
}

// 向node注册回包
func (n *NodeService) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_DISPATCH && rsp.AppId == n.s.AppId {
		logger.Info("已向node注册成功！")
	}
}

// ping包请求
func (n *NodeService) getAllServiceGateReq() {
	req := &spb.GetAllServiceGateReq{
		ServiceType:  spb.ServerType_SERVICE_DISPATCH,
		DispatchTime: time.Now().UnixNano() / 1e6,
	}
	n.sendNode(cmd.GetAllServiceGateReq, req)
}

// ping包回应
func (n *NodeService) GetAllServiceGateRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetAllServiceGateRsp)

	gateList := make(map[uint32]*Gate)
	for _, service := range rsp.GateServiceList {
		gate := &Gate{
			Ip:   service.Addr,
			Port: alg.S2U32(service.Port),
			Num:  service.PlayerNum,
		}
		gateList[service.AppId] = gate
	}
	n.setGateList(gateList)
	logger.Debug("dispatch <--> node ping:%v", (rsp.NodeTime-rsp.DispatchTime)/2)
}

func (n *NodeService) setGateList(gateList map[uint32]*Gate) {
	n.gateListLock.Lock()
	n.gateList = gateList
	n.gateListLock.Unlock()
}

func (n *NodeService) getGate() *Gate {
	var minAppId uint32
	var minNum int64
	n.gateListLock.Lock()
	for id, gate := range n.gateList {
		if minAppId == 0 || minNum > gate.Num {
			minAppId = id
			minNum = gate.Num
		}
	}
	gate := n.gateList[minAppId]
	n.gateListLock.Unlock()
	return gate
}
