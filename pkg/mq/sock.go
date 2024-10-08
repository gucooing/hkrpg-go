package mq

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/gucooing/gunet"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg/alg"
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
	nodeTcp                *NodeTcp                 // 和node的tcp连接
	discoveryClient        *rpc.NodeDiscoveryClient // 和node连接的rpc
	gateTcpMqInstMap       map[spb.ServerType]map[uint32]*GateTcpMqInst
	gateTcpMqSync          *sync.RWMutex // 读写锁
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

type NodeTcp struct {
	nodeMqAddr string
	conn       *gunet.TcpConn
	state      nodeConn
}
type nodeConn int

const (
	nodeConnLost  nodeConn = 0
	nodeConnEct   nodeConn = 1
	nodeConnClose nodeConn = 2
)

func NewMessageQueue(serverType spb.ServerType, appId uint32,
	discoveryClient *rpc.NodeDiscoveryClient, gateAddr, nodeMqAddr, regionName string) (r *MessageQueue) {
	if discoveryClient == nil && serverType != spb.ServerType_SERVICE_NODE {
		logger.Error("discoveryClient error")
		return nil
	}
	r = new(MessageQueue)
	r.appId = appId
	r.serverType = serverType
	r.regionName = regionName
	r.discoveryClient = discoveryClient
	r.netMsgInput = make(chan *NetMsg, 1000)
	r.netMsgOutput = make(chan *NetMsg, 1000)
	r.gateTcpMqEventChan = make(chan *GateTcpMqEvent, 1000)
	r.gateTcpMqDeadEventChan = make(chan string, 1000)
	r.gateTcpMqSync = new(sync.RWMutex)
	r.nodeTcp = &NodeTcp{
		nodeMqAddr: nodeMqAddr,
	}

	if serverType == spb.ServerType_SERVICE_GATE ||
		serverType == spb.ServerType_SERVICE_NODE {
		go r.runGateTcpMqServer(gateAddr)
	} else if serverType == spb.ServerType_SERVICE_DISPATCH ||
		serverType == spb.ServerType_SERVICE_GAME ||
		serverType == spb.ServerType_SERVICE_MUIP ||
		serverType == spb.ServerType_SERVICE_MULTI {
		go r.runGateTcpMqClient()
	}
	if serverType != spb.ServerType_SERVICE_NODE {
		go r.nodeTcpMq()
	}
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
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	m.nodeTcp.Close()
}

func (m *MessageQueue) runGateTcpMqServer(gateAddr string) {
	listener, err := gunet.NewTcpS(gateAddr)
	if err != nil {
		logger.Error("tcp mq listen error: %v", err)
		return
	}
	logger.Info("new mq server")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("tcp mq accept error: %v", err)
			return
		}
		logger.Info("accept tcp mq, server addr: %v", conn.RemoteAddr().String())
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
	protoObj := smd.DecodePayloadToProto(&alg.PackMsg{
		CmdId:     netMsg.CmdId,
		ProtoData: netMsg.ServiceMsgByte,
	})
	if protoObj == nil {
		return
	}
	req := protoObj.(*spb.GateTcpMqHandshakeReq)
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
	m.gateTcpMqConn(gateServerConnAddrMap)
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

func (m *MessageQueue) GetGateTcpMqInst(serverType spb.ServerType, appId uint32) *GateTcpMqInst {
	m.gateTcpMqSync.RLock()
	defer m.gateTcpMqSync.RUnlock()
	instMap, exist := m.gateTcpMqInstMap[serverType]
	if !exist {
		return nil
	}
	inst, exist := instMap[appId]
	if !exist {
		return nil
	}
	return inst
}

// 发信
func (m *MessageQueue) sendHandler() {
	m.gateTcpMqInstMap = map[spb.ServerType]map[uint32]*GateTcpMqInst{
		spb.ServerType_SERVICE_GATE:     make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_DISPATCH: make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_GAME:     make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_MUIP:     make(map[uint32]*GateTcpMqInst),
		spb.ServerType_SERVICE_MULTI:    make(map[uint32]*GateTcpMqInst),
	}
	for {
		select {
		case netMsg := <-m.netMsgInput:
			var conn *gunet.TcpConn
			if netMsg.ServerType == spb.ServerType_SERVICE_NODE {
				conn = m.nodeTcp.conn
			} else {
				inst := m.GetGateTcpMqInst(netMsg.ServerType, netMsg.AppId)
				if inst == nil {
					logger.Error("unknown server type: %v", netMsg.ServerType)
					continue
				}
				conn = inst.conn
			}
			if conn == nil {
				logger.Error("unknown server type: %v", netMsg.ServerType)
				continue
			}
			if netMsg.MsgType == ServerMsg {
				netMsg.ServiceMsgByte = smd.EncodeProtoToPayload(
					&smd.ProtoMsg{
						CmdId:          netMsg.CmdId,
						PayloadMessage: netMsg.ServiceMsgPb,
					})
			}

			bin := EncodePayloadToBin(netMsg)
			_, err := conn.Write(bin)
			if err != nil {
				logger.Error("write error: %v", err)
				continue
			}
		case event := <-m.gateTcpMqEventChan:
			inst := event.inst
			switch event.event {
			case EventConnect: // 创建连接
				logger.Info("tcp mq connect, addr: %v, server type: %v, appid: %v", inst.conn.RemoteAddr().String(), inst.serverType, inst.appId)
				m.gateTcpMqSync.Lock()
				m.gateTcpMqInstMap[inst.serverType][inst.appId] = inst
				m.gateTcpMqSync.Unlock()
			case EventDisconnect: // 删除连接
				logger.Info("tcp mq disconnect, addr: %v, server type: %v, appid: %v", inst.conn.RemoteAddr().String(), inst.serverType, inst.appId)
				m.gateTcpMqSync.Lock()
				delete(m.gateTcpMqInstMap[inst.serverType], inst.appId)
				m.netMsgOutput <- &NetMsg{
					OriginServerType:  inst.serverType,
					OriginServerAppId: inst.appId,
					MsgType:           ServiceLogout,
				}
				m.gateTcpMqSync.Unlock()
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
			ServiceMsgByte: smd.EncodeProtoToPayload(
				&smd.ProtoMsg{
					CmdId: smd.GateTcpMqHandshakeReq,
					PayloadMessage: &spb.GateTcpMqHandshakeReq{
						Type:  m.serverType,
						AppId: m.appId,
					},
				}),
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

func (m *MessageQueue) nodeTcpMq() {
	nodeTcp := m.nodeTcp
	for {
		if nodeTcp.state == nodeConnClose { // 连接验证
			return
		}
		conn, err := gunet.NewTcpC(nodeTcp.nodeMqAddr)
		if err != nil {
			logger.Error("node tcp mq connect error: %v", err)
			return
		}
		netMsg := &NetMsg{
			CmdId: smd.GateTcpMqHandshakeReq,
			ServiceMsgByte: smd.EncodeProtoToPayload(
				&smd.ProtoMsg{
					CmdId: smd.GateTcpMqHandshakeReq,
					PayloadMessage: &spb.GateTcpMqHandshakeReq{
						Type:  m.serverType,
						AppId: m.appId,
					},
				}),
		}
		bin := EncodePayloadToBin(netMsg)
		_, err = conn.Write(bin)
		if err != nil {
			logger.Error("node tcp mq handshake error: %v", err)
			return
		}
		nodeTcp.conn = conn

		nodeTcp.state = nodeConnEct
		logger.Info("node tcp mq connect, addr: %v", nodeTcp.nodeMqAddr)
		err = m.nodeTcpMqRecvHandle()
		if err != nil {
			logger.Error("node tcp mq receive error: %v", err)
		}
		time.Sleep(time.Second * 5)
	}
}

// node收信
func (m *MessageQueue) nodeTcpMqRecvHandle() error {
	nodeTcp := m.nodeTcp
	if nodeTcp.conn == nil {
		return errors.New("node tcp mq connect fail")
	}
	for {
		bin, err := nodeTcp.conn.Read()
		if err != nil {
			logger.Error("node tcp mq receive error: %v", err)
			nodeTcp.state = nodeConnLost
			return err
		}
		netMsg := DecodeBinToPayload(bin)
		m.netMsgOutput <- netMsg
	}
}

func (n *NodeTcp) Close() {
	if n.state == nodeConnClose {
		return
	}
	n.state = nodeConnClose
	if n.conn != nil {
		n.conn.Close()
	}
}
