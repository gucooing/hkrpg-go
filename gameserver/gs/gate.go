package gs

import (
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

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

	// 异步拉取账户数据
	// go func() {
	g.GetPlayerDate()
	if s.PlayerMap[req.PlayerUid].Player == nil {
		s.PlayerMap[req.PlayerUid].Player = &player.PlayerData{
			Battle: make(map[uint32]*player.Battle),
			BattleState: &player.BattleState{
				ChallengeState: &player.ChallengeState{},
			},
		}
	}

	s.gamePlayerLoginRsp(g)
	// }()

	/*
		// 异步向node同步在线数据
		go func() {

			pdsm := &spb.SyncPlayerOnlineDataNotify{
				PlayerUid: g.Uid,
			}
			s.sendNode(cmd.SyncPlayerOnlineDataNotify, pdsm)

			s.gamePlayerLoginRsp(g)
		}()
	*/

}

func (s *GameServer) gamePlayerLoginRsp(g *player.GamePlayer) {
	if g.PlayerPb == nil { // || g.Player == nil {
		return
	} else {
		// 通知node
		s.sendNode(cmd.PlayerLoginReq, &spb.PlayerLoginReq{PlayerUid: g.Uid}) // TODO 改成通知，告诉node玩家上线了
		// 回复gate
		rsp := &spb.PlayerLoginRsp{PlayerUid: g.Uid} // 这里是回复gate，让gate知道gs已经准备好了迎接玩家登录gs
		g.SendGate(cmd.PlayerLoginRsp, rsp)
	}
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
