package multi

import (
	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type gateServer struct {
	multi *Multi
	appid uint32
	conn  *gunet.TcpConn // gate tcp通道
}

func (s *Multi) addGeList(ge *gateServer) {
	s.gateListLock.Lock()
	s.gateList[ge.appid] = ge
	s.gateListLock.Unlock()
}

func (s *Multi) delGeList(appid uint32) {
	s.gateListLock.Lock()
	delete(s.gateList, appid)
	s.gateListLock.Unlock()
}

func (s *Multi) getGeByAppid(appid uint32) *gateServer {
	s.gateListLock.Lock()
	defer s.gateListLock.Unlock()
	return s.gateList[appid]
}

// 从gate接收消息
func (s *Multi) recvGate(conn *gunet.TcpConn, appid uint32) {
	ge := &gateServer{
		multi: s,
		appid: appid,
		conn:  conn,
	}
	s.addGeList(ge)
	rsp := &spb.GateLoginMultiRsp{
		Retcode: 0,
	}
	ge.seedGate(cmd.GateLoginMultiRsp, rsp)
	logger.Info("gate:[%v]在multi注册成功", appid)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker gate: %v", appid)
			ge.killGate()
		}
	}()

	for {
		bin, err := conn.Read()
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			ge.killGate()
			return
		}
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			go ge.gateRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (ge *gateServer) gateRegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {
	default:
		logger.Error("gate -> multi cmdid error: %v", cmdId)
	}
}

func (ge *gateServer) seedGate(cmdId uint16, payloadMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = payloadMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdid error")
		return
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err = ge.conn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

// gate离线
func (ge *gateServer) killGate() {
	ge.multi.delGeList(ge.appid)
	logger.Info("[APPID:%v]gate server离线", ge.appid)
}
