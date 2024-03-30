package gs

import (
	"sync"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

var syncGD sync.Mutex

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
	case cmd.PlayerToGameByGateReq:
		s.PlayerToGameByGateReq(g, payloadMsg)

	case cmd.PlayerLoginNotify:
		s.PlayerLoginNotify(g, payloadMsg) // gate请求登录
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

/******************************************NewLogin***************************************/

func (s *GameServer) PlayerLoginNotify(g *player.GamePlayer, payloadMsg pb.Message) {
	notify := payloadMsg.(*spb.PlayerLoginNotify)

	if notify.Uuid == 0 || notify.AccountId == 0 || notify.GameServerAppId != s.AppId || notify.Uid == 0 {
		return
	}
	g.Uid = notify.Uid
	g.AccountId = notify.AccountId
	g.Uuid = notify.Uuid
	g.GateAppId = notify.GateServerAppId

	// 异步拉取账户数据
	go func() {
		g.GetPlayerDate(notify.Uid)
		syncGD.Lock()
		s.AddPlayer(notify.Uuid, g)
		// 初始化在线数据
		if s.PlayerMapS[notify.Uuid].Player == nil {
			s.PlayerMapS[notify.Uuid].Player = &player.PlayerData{
				Battle: make(map[uint32]*player.Battle),
				BattleState: &player.BattleState{
					ChallengeState: &player.ChallengeState{},
				},
			}
		}
		syncGD.Unlock()
	}()
	logger.Info("[UID:%v]|[UUID:%v]登录game", g.Uid, notify.Uuid)

	// 通知node玩家登录
	s.sendNode(cmd.PlayerLoginNotify, &spb.PlayerLoginNotify{
		Uuid:            g.Uuid,
		AccountId:       g.AccountId,
		Uid:             g.Uid,
		GateServerAppId: g.GateAppId,
		GameServerAppId: s.AppId,
	})
}

func (s *GameServer) AddPlayer(uuid int64, g *player.GamePlayer) {
	s.PlayerMapS[uuid] = g
}
