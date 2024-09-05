package mq

import (
	"context"
	"io"
	"time"

	"github.com/gucooing/gunet"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	smd "github.com/gucooing/hkrpg-go/protocol/server"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	EventConnect = iota
	EventDisconnect
)

var ctx = context.TODO()

type MessageQueue struct {
	appId                  uint32
	serverType             spb.ServerType
	regionName             string                   // 区服配置
	netMsgInput            chan *NetMsg             // 发包
	netMsgOutput           chan *NetMsg             // 收包
	gateTcpMqEventChan     chan *GateTcpMqEvent     // 连接管理
	gateTcpMqDeadEventChan chan string              // gate 管理
	discoveryClient        *rpc.NodeDiscoveryClient // 和node连接的rpc
}

type GateTcpMqEvent struct {
	event int
	inst  *GateTcpMqInst
}

type GateTcpMqInst struct {
	conn       *gunet.TcpConn
	serverType spb.ServerType
	appId      uint32
}

func NewMessageQueue(serverType spb.ServerType, appId uint32,
	discoveryClient *rpc.NodeDiscoveryClient, gateAddr, regionName string) (r *MessageQueue) {
	if discoveryClient == nil {
		logger.Error("discoveryClient error")
		return nil
	}
	r = new(MessageQueue)
	r.appId = appId
	r.serverType = serverType
	r.regionName = regionName
	r.discoveryClient = discoveryClient
	r.netMsgInput = make(chan *NetMsg, 100)
	r.netMsgOutput = make(chan *NetMsg, 100)
	r.gateTcpMqEventChan = make(chan *GateTcpMqEvent, 100)
	r.gateTcpMqDeadEventChan = make(chan string, 100)

	if serverType == spb.ServerType_SERVICE_GATE {
		go r.runGateTcpMqServer(gateAddr)
	} else if serverType == spb.ServerType_SERVICE_DISPATCH ||
		serverType == spb.ServerType_SERVICE_GAME ||
		serverType == spb.ServerType_SERVICE_MUIP ||
		serverType == spb.ServerType_SERVICE_MULTI {
		go r.runGateTcpMqClient()
	}
	go r.nodeGrpcRecvHandle()
	go r.sendHandler()
	return
}

func (m *MessageQueue) GetNetMsg() chan *NetMsg {
	return m.netMsgOutput
}

func (m *MessageQueue) Close() {
	// 等待所有待发送的消息发送完毕
	for {
		if len(m.netMsgInput) == 0 {
			time.Sleep(time.Millisecond * 100)
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (m *MessageQueue) runGateTcpMqServer(gateAddr string) {
	listener, err := gunet.NewTcpS(gateAddr)
	if err != nil {
		logger.Error("tcp mq listen error: %v", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("tcp mq accept error: %v", err)
			return
		}
		logger.Info("accept gate tcp mq, server addr: %v", conn.RemoteAddr().String())
		go m.gateTcpMqHandshake(conn)
	}
}

func (m *MessageQueue) gateTcpMqHandshake(conn *gunet.TcpConn) {
	bin, err := conn.Read()
	if err != nil {
		logger.Error("tcp mq handshake error: %v", err)
		return
	}
	inst := &GateTcpMqInst{
		conn:       conn,
		serverType: 0,
		appId:      0,
	}
	netMsg := DecodeBinToPayload(bin)
	if netMsg.CmdId != smd.GateTcpMqHandshakeReq {
		logger.Error("tcp mq handshake error: %v", conn.RemoteAddr().String())
		return
	}
	if !DecodePayloadToProto(netMsg) {
		return
	}
	req := netMsg.ServiceMsg.(*spb.GateTcpMqHandshakeReq)
	switch req.Type {
	case spb.ServerType_SERVICE_DISPATCH:
		inst.serverType = spb.ServerType_SERVICE_DISPATCH
	case spb.ServerType_SERVICE_GATE:
		inst.serverType = spb.ServerType_SERVICE_GATE
	case spb.ServerType_SERVICE_MUIP:
		inst.serverType = spb.ServerType_SERVICE_MUIP
	case spb.ServerType_SERVICE_MULTI:
		inst.serverType = spb.ServerType_SERVICE_MULTI
	case spb.ServerType_SERVICE_GAME:
		inst.serverType = spb.ServerType_SERVICE_GAME
	default:
		logger.Error("unknown gate tcp mq handshake req: %v", req.Type)
		return
	}
	inst.appId = req.AppId
	if inst.appId == 0 {
		logger.Error("unknown gate tcp mq handshake req: %v", inst.appId)
		return
	}
	m.gateTcpMqEventChan <- &GateTcpMqEvent{
		event: EventConnect,
		inst:  inst,
	}
	go m.gateTcpMqRecvHandle(inst)
}

func (m *MessageQueue) runGateTcpMqClient() {
	gateServerConnAddrMap := make(map[string]bool)
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case addr := <-m.gateTcpMqDeadEventChan:
			// gate 断开连接
			delete(gateServerConnAddrMap, addr)
		case <-ticker.C: // 定时获取全部gate地址
			m.gateTcpMqConn(gateServerConnAddrMap)
		}
	}
}

// 发信
func (m *MessageQueue) sendHandler() {
	gateTcpMqInstMap := map[spb.ServerType]map[uint32]*GateTcpMqInst{
		spb.ServerType_SERVICE_GATE:     make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_DISPATCH: make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_GAME:     make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_MUIP:     make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_MULTI:    make(map[uint32]*GateTcpMqInst),
	}
	for {
		select {
		case netMsg := <-m.netMsgInput:
			logger.Info("", netMsg)
			instMap, exist := gateTcpMqInstMap[netMsg.ServerType]
			if !exist {
				logger.Error("unknown server type: %v", netMsg.ServerType)
				continue
			}
			inst, exist := instMap[netMsg.AppId]
			if !exist {
				continue
			}
			if !EncodeProtoToPayload(netMsg) {
				continue
			}
			bin := EncodePayloadToBin(netMsg)
			_, err := inst.conn.Write(bin)
			if err != nil {
				logger.Error("write error: %v", err)
				continue
			}
		case event := <-m.gateTcpMqEventChan:
			inst := event.inst
			switch event.event {
			case EventConnect: // 创建连接
				logger.Info("gate tcp mq connect, addr: %v, server type: %v, appid: %v", inst.conn.RemoteAddr().String(), inst.serverType, inst.appId)
				gateTcpMqInstMap[inst.serverType][inst.appId] = inst
			case EventDisconnect: // 删除连接
				logger.Info("gate tcp mq disconnect, addr: %v, server type: %v, appid: %v", inst.conn.RemoteAddr().String(), inst.serverType, inst.appId)
				delete(gateTcpMqInstMap[inst.serverType], inst.appId)
				m.gateTcpMqDeadEventChan <- inst.conn.RemoteAddr().String()
			}
		}
	}
}

// 收信
func (m *MessageQueue) gateTcpMqRecvHandle(inst *GateTcpMqInst) {
	for {
		bin, err := inst.conn.Read()
		if err != nil {
			logger.Error("tcp mq receive error: %v", err)
			m.gateTcpMqEventChan <- &GateTcpMqEvent{
				event: EventDisconnect,
				inst:  inst,
			}
			inst.conn.Close()
			return
		}
		netMsg := DecodeBinToPayload(bin)
		if DecodePayloadToProto(netMsg) {
			m.netMsgOutput <- netMsg
		}
	}
}

func (m *MessageQueue) nodeGrpcRecvHandle() {
	stream, err := m.discoveryClient.NodeStreamMessages(ctx, &nodeapi.NodeStreamMessagesReq{})
	if err != nil {
		logger.Error("node grpc receive error: %v", err)
	}
	for {
		rsp, err := stream.Recv()
		if err == io.EOF {
			logger.Info("node grpc server closed")
			return
		}
		if err != nil {
			logger.Error("node grpc receive error: %v", err)
			return
		}
		logger.Info("", rsp) // TODO
		netMsg := &NetMsg{
			ServerType:        0,
			AppId:             0,
			OriginServerType:  0,
			OriginServerAppId: 0,
			MsgType:           NodeMsg,
			Uid:               0,
			CmdId:             0,
			ServiceMsgByte:    nil,
			ServiceMsg:        nil,
		}
		m.netMsgOutput <- netMsg
	}
}

func (m *MessageQueue) gateTcpMqConn(gateServerConnAddrMap map[string]bool) {
	rsp, err := m.discoveryClient.GetAllGateServerMq(ctx, &nodeapi.GetAllGateServerMqReq{RegionName: m.regionName})
	if err != nil {
		logger.Error("gate tcp mq get gate list error: %v", err)
		return
	}
	for _, server := range rsp.ServerList {
		_, exist := gateServerConnAddrMap[server.MqAddr]
		// GATE连接已存在
		if exist {
			logger.Debug("gate tcp mq conn already exist addr: %v", server.MqAddr)
			continue
		}

		conn, err := gunet.NewTcpC(server.MqAddr)
		if err != nil {
			continue
		}
		netMsg := &NetMsg{
			CmdId: smd.GateTcpMqHandshakeReq,
			ServiceMsg: &spb.GateTcpMqHandshakeReq{
				Type:  spb.ServerType_SERVICE_DISPATCH,
				AppId: 2,
			},
		}
		if !EncodeProtoToPayload(netMsg) {
			continue
		}
		bin := EncodePayloadToBin(netMsg)
		_, err = conn.Write(bin)
		if err != nil {
			continue
		}
		inst := &GateTcpMqInst{
			conn:       conn,
			serverType: spb.ServerType_SERVICE_GATE,
			appId:      server.AppId,
		}
		m.gateTcpMqEventChan <- &GateTcpMqEvent{
			event: EventConnect,
			inst:  inst,
		}
		gateServerConnAddrMap[server.MqAddr] = true
		logger.Info("connect gate tcp mq, gate addr: %v", conn.RemoteAddr().String())
		go m.gateTcpMqRecvHandle(inst)
	}
}
