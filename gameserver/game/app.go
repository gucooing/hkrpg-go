package game

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GAMESERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			Close()
			os.Exit(0)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := s.nodeConn.Read(nodeMsg)
		if err != nil {
			log.Println("node error")
			os.Exit(0)
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
	// 发送心跳包
	go s.GetAllServiceReq()
}

func (s *GameServer) GetAllServiceReq() {
	// 心跳包
	for {
		req := &spb.GetAllServiceReq{
			ServiceType: spb.ServerType_SERVICE_GAME,
		}
		s.sendNode(cmd.GetAllServiceReq, req)
		time.Sleep(time.Second * 5)
	}
}

/************************************gate********************************/

// 从gate接收消息
func (s *GameServer) recvGate(g *player.GamePlayer) {
	nodeMsg := make([]byte, player.PacketMaxLen)

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker player uid: %v", g.Uid)
			KickPlayer(g)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := g.GateConn.Read(nodeMsg)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			KickPlayer(g)
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
	case cmd.PlayerLogoutNotify:
		s.PlayerLogoutNotify(g, payloadMsg)
	}
}

func (s *GameServer) PlayerLoginReq(g *player.GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*spb.PlayerLoginReq)
	if req.PlayerUid == 0 {
		return
	}
	if pla := s.PlayerMap[req.PlayerUid]; pla != nil {
		KickPlayer(s.PlayerMap[req.PlayerUid])
	}
	logger.Info("[UID:%v]玩家登录gs", req.PlayerUid)
	g.Uid = req.PlayerUid
	s.PlayerMap[req.PlayerUid] = g
	// 拉取账户数据
	g.GetPlayerDate()
	if s.PlayerMap[req.PlayerUid].Player == nil {
		s.PlayerMap[req.PlayerUid].Player = &player.PlayerData{
			Battle: make(map[uint32]*player.Battle),
			BattleState: &player.BattleState{
				ChallengeState: &player.ChallengeState{},
			},
		}
	}
	// 异步向node同步在线数据
	go func() {
		s.sendNode(cmd.PlayerLoginReq, &spb.PlayerLoginReq{PlayerUid: req.PlayerUid})
		pdsm := &spb.SyncPlayerOnlineDataNotify{
			PlayerUid: g.Uid,
		}
		s.sendNode(cmd.SyncPlayerOnlineDataNotify, pdsm)
	}()

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

func (s *GameServer) PlayerLogoutNotify(g *player.GamePlayer, payloadMsg pb.Message) {
	noti := payloadMsg.(*spb.PlayerLogoutNotify)
	if players := s.PlayerMap[noti.PlayerUid]; players != nil {
		s.sendNode(cmd.PlayerLogoutNotify, payloadMsg)
		KickPlayer(g)
		logger.Info("[UID:%v]game玩家离线成功", noti.PlayerUid)
	}
}

func KickPlayer(g *player.GamePlayer) {
	/*
		TODO
		1.保存数据到数据库
		2.断开gate-game连接
	*/
	logger.Debug("[UID:%v]玩家离线", g.Uid)
	GAMESERVER.SyncPlayerDate(g)
	UpDataPlayer(g)
	g.GateConn.Close()
	delete(GAMESERVER.PlayerMap, g.Uid)
}

func UpDataPlayer(g *player.GamePlayer) error {
	var err error
	if g.PlayerPb == nil {
		return nil
	}
	if g.Uid == 0 {
		return nil
	}
	dbDate := new(db.Player)
	dbDate.AccountUid = g.Uid

	dbDate.PlayerDataPb, err = pb.Marshal(g.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
	}

	if err = db.DBASE.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return err
	}

	logger.Debug("数据库账号:%v 数据更新", g.Uid)
	return nil
}

func (s *GameServer) SyncPlayerDate(g *player.GamePlayer) {
	playerBin, _ := json.Marshal(g.Player)
	pdsm := &spb.SyncPlayerOnlineDataNotify{
		PlayerUid:        g.Uid,
		PlayerOnlineData: playerBin,
	}
	s.sendNode(cmd.SyncPlayerOnlineDataNotify, pdsm)
	logger.Debug("[UID:%v]在线数据已同步到node", g.Uid)
}
