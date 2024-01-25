package game

import (
	"github.com/gucooing/hkrpg-go/gameserver/logger"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 向node注册
func (s *GameServer) Connection() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_GAME,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp,
		Port:       s.Port,
	}

	s.sendNode(cmd.ServiceConnectionReq, req)
}

// 从node接收消息
func (s *GameServer) recvNode() {
	nodeMsg := make([]byte, player.PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := s.nodeConn.Read(nodeMsg)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.NodeRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

// 发送到node
func (s *GameServer) sendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdId error")
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := s.nodeConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

func (s *GameServer) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_GAME && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
	// TODO 发送心跳包
}

/************************************gate********************************/

// 从gate接收消息
func (s *GameServer) recvGate(g *player.GamePlayer) {
	nodeMsg := make([]byte, player.PacketMaxLen)
	for {
		var bin []byte = nil
		recvLen, err := g.GateConn.Read(nodeMsg)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.GateRegisterMessage(g, msg.CmdId, serviceMsg)
		}
	}
}

func (s *GameServer) GateRegisterMessage(g *player.GamePlayer, cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLoginReq:
		s.PlayerLoginReq(g, payloadMsg) // gate玩家登录通知
	case cmd.PlayerToGameByGateReq:
		s.PlayerToGameByGateReq(g, payloadMsg)
	}
}

func (s *GameServer) PlayerLoginReq(g *player.GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*spb.PlayerLoginReq)
	if req.PlayerUid == 0 {
		return
	}
	s.PlayerMap[req.PlayerUid] = g
	g.Uid = req.PlayerUid
	g.GetPlayerDate()

	rsp := &spb.PlayerLoginRsp{PlayerUid: req.PlayerUid}
	g.SendGate(cmd.PlayerLoginRsp, rsp)
}

// 从gate收到的玩家数据包
func (s *GameServer) PlayerToGameByGateReq(g *player.GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*spb.PlayerToGameByGateReq)
	playerMsgList := make([]*alg.PackMsg, 0)
	alg.DecodeBinToPayload(req.PlayerBin, &playerMsgList, nil)
	for _, msg := range playerMsgList {
		g.RegisterMessage(msg.CmdId, msg.ProtoData)
	}
}
