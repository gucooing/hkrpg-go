package muip

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type NodeService struct {
	muip         *MuipServer
	nodeConn     *gunet.TcpConn
	tickerCancel context.CancelFunc
	ticker       *time.Ticker // 定时器
}

type Service struct {
	AppId     string
	PlayerNum int64
}

func (s *MuipServer) newNode() {
	n := new(NodeService)
	tcpConn, err := gunet.NewTcpC(s.Config.NetConf["Node"])
	if err != nil {
		logger.Error("nodeserver error:%s", err.Error())
		return
	}
	n.nodeConn = tcpConn
	n.muip = s
	n.ticker = time.NewTicker(5 * time.Second)
	tickerCtx, tickerCancel := context.WithCancel(context.Background())
	n.tickerCancel = tickerCancel
	s.node = n
	go n.recvNode()
	// 向node注册
	n.ServiceConnectionReq()
	// 开启node定时器
	n.nodeTicler(tickerCtx)
}

func (n *NodeService) nodeKill() {
	n.nodeConn.Close()
	n.tickerCancel()
	logger.Info("node server离线")
	n.muip.node = nil
}

func (n *NodeService) nodeTicler(tickerCtx context.Context) {
	for {
		select {
		case <-n.ticker.C:
			n.MuipToNodePingReq()
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
		AppId:      n.muip.AppId,
		Addr:       n.muip.ApiAddr,
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
	case cmd.ServiceConnectionRsp:
		n.ServiceConnectionRsp(serviceMsg)
	case cmd.MuipToNodePingRsp:
		n.MuipToNodePingRsp(serviceMsg)
	default:
		logger.Info("node -> muip error cmdid:%v", cmdId)
	}
}

// 来自api的消息转发到nodeserver
func (a *Api) ToNode(c *gin.Context, cmdId uint16, message pb.Message) {
	a.muip.node.sendNode(cmdId, message)
	c.JSON(200, gin.H{
		"code": 0,
	})
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
	if rsp.ServerType == spb.ServerType_SERVICE_MUIP && rsp.AppId == n.muip.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (n *NodeService) MuipToNodePingReq() {
	req := &spb.MuipToNodePingReq{
		MuipServerTime: time.Now().Unix(),
	}
	n.sendNode(cmd.MuipToNodePingReq, req)
}

func (n *NodeService) MuipToNodePingRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.MuipToNodePingRsp)
	if rsp.NodeServerTime-rsp.MuipServerTime > 5 {
		logger.Error("调用时间过长")
	}
	allService := make(map[spb.ServerType][]*Service, 0)
	for id, serviceList := range rsp.ServiceList {
		if len(serviceList.ServiceList) == 0 {
			continue
		}
		if allService[spb.ServerType(id)] == nil {
			allService[spb.ServerType(id)] = make([]*Service, 0)
		}
		for _, service := range serviceList.ServiceList {
			allService[spb.ServerType(id)] = append(allService[spb.ServerType(id)], &Service{
				AppId:     alg.GetAppIdStr(service.AppId),
				PlayerNum: service.PlayerNum,
			})
		}
	}

	n.muip.setAllService(allService)
}
