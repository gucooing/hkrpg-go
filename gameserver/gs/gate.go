package gs

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

var syncGD sync.Mutex

type gateServer struct {
	game          *GameServer
	appid         uint32
	playerMap     map[int64]*player.GamePlayer // 玩家列表
	playerMapLock sync.Mutex                   // 玩家列表互斥锁
	conn          net.Conn                     // gate tcp通道

	msgChan chan player.Msg // 消息通道
}

func (s *GameServer) addGsList(ge *gateServer) {
	s.gateListLock.Lock()
	s.gateList[ge.appid] = ge
	s.gateListLock.Unlock()
}

func (s *GameServer) delGsList(appid uint32) {
	s.gateListLock.Lock()
	delete(s.gateList, appid)
	s.gateListLock.Unlock()
}

// 从gate接收消息
func (s *GameServer) recvGate(conn net.Conn, appid uint32) {
	ge := &gateServer{
		game:      s,
		appid:     appid,
		playerMap: make(map[int64]*player.GamePlayer),
		conn:      conn,
	}
	s.addGsList(ge)
	rsp := &spb.GateLoginGameRsp{
		Retcode: 0,
	}
	ge.seedGate(cmd.GateLoginGameRsp, rsp)
	ge.msgChan = make(chan player.Msg, 10)
	go ge.recvPlayer()
	logger.Info("gate:[%v]在game注册成功", appid)

	nodeMsg := make([]byte, player.PacketMaxLen)

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker gate: %v", appid)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := conn.Read(nodeMsg)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			// KickPlayer(g)
			return
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			ge.gateRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

// 接收player传来的消息
func (ge *gateServer) recvPlayer() {
	for {
		bin := <-ge.msgChan
		ge.playerToGame(bin)
	}
}

func (ge *gateServer) gateRegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {
	case cmd.GateGamePingReq:
		ge.GateGamePingReq(payloadMsg) // 来自gate的ping包
	case cmd.GateGamePlayerLoginReq:
		ge.GateGamePlayerLoginReq(payloadMsg) // 来自gate的玩家登录请求
	case cmd.GetToGamePlayerLogoutReq:
		ge.GetToGamePlayerLogoutReq(payloadMsg) // gate直接向目标game申请下线玩家请求
	case cmd.GateToGameMsgNotify:
		ge.GateToGameMsgNotify(payloadMsg) // gate转发客户端消息到gs
	}
}

func (ge *gateServer) playerToGame(msg player.Msg) {
	switch msg.CmdId {
	case cmd.GameToGateMsgNotify:
		ge.GameToGateMsgNotify(msg.PlayerMsg)
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
	_, err := ge.conn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

func (ge *gateServer) GateGamePingReq(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GateGamePingReq)
	rsp := &spb.GateGamePingRsp{
		GateServerTime: req.GateServerTime,
		GameServerTime: time.Now().Unix(),
		PlayerNum:      ge.game.GetPlayerNum(),
	}
	ge.seedGate(cmd.GateGamePingRsp, rsp)
}

func (ge *gateServer) GateGamePlayerLoginReq(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GateGamePlayerLoginReq)
	if req.Uid == 0 || req.Uuid == 0 || req.AccountId == 0 {
		logger.Error("player login uid or uuid error")
		return
	}
	g := NewPlayer(req.Uid, req.AccountId, req.Uuid, ge.msgChan)
	// 拉取账户数据
	g.GetPlayerDate(req.Uid)
	ge.game.AddPlayerMap(req.Uuid, g)
	logger.Info("[UID:%v]|[UUID:%v]登录game", g.Uid, req.Uuid)
	ge.AddPlayerStatus(g)
	rsp := &spb.GateGamePlayerLoginRsp{
		Retcode: 0,
		Uid:     req.Uid,
		Uuid:    req.Uuid,
	}
	ge.seedGate(cmd.GateGamePlayerLoginRsp, rsp)
}

func (ge *gateServer) GetToGamePlayerLogoutReq(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GetToGamePlayerLogoutReq)
	play := ge.game.GetPlayerByUuid(req.GetOldUuid())
	if play == nil {
		return
	}
	// 下线玩家
	ge.killPlayer(play)

	rsp := &spb.GetToGamePlayerLogoutRsp{
		Retcode:         spb.Retcode_RET_SUCC,
		Uid:             req.Uid,
		NewUuid:         req.NewUuid,
		NewGameServerId: req.NewGameServerId,
	}
	ge.seedGate(cmd.GetToGamePlayerLogoutRsp, rsp)
}

func NewPlayer(uid, accountId uint32, uuid int64, msg chan player.Msg) *player.GamePlayer {
	g := new(player.GamePlayer)
	g.LastActiveTime = time.Now().Unix()
	g.Uid = uid
	g.AccountId = accountId
	g.Uuid = uuid
	g.MsgChan = msg

	return g
}

func (ge *gateServer) GateToGameMsgNotify(payloadMsg pb.Message) {
	rsp := payloadMsg.(*spb.GateToGameMsgNotify)
	paler := ge.game.GetPlayerByUuid(rsp.Uuid)
	if paler != nil {
		msgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(rsp.Msg, &msgList, nil)
		for _, msg := range msgList {
			paler.RegisterMessage(msg.CmdId, msg.ProtoData)
		}
	}
}

func (ge *gateServer) GameToGateMsgNotify(payloadMsg pb.Message) {
	ge.seedGate(cmd.GameToGateMsgNotify, payloadMsg)
}

func (s *GameServer) AddPlayerMap(uuid int64, g *player.GamePlayer) {
	syncGD.Lock()
	s.PlayerMap[uuid] = g
	// 初始化在线数据
	if s.PlayerMap[g.Uuid].Player == nil {
		s.PlayerMap[g.Uuid].Player = &player.PlayerData{
			Battle: make(map[uint32]*player.Battle),
			BattleState: &player.BattleState{
				ChallengeState: &player.ChallengeState{},
			},
		}
	}
	syncGD.Unlock()
}

func (s *GameServer) DelPlayerMap(uuid int64) {
	syncGD.Lock()
	if s.PlayerMap[uuid] != nil {
		delete(s.PlayerMap, uuid)
	}
	syncGD.Unlock()
}

func (s *GameServer) GetPlayerByUuid(uuid int64) *player.GamePlayer {
	syncGD.Lock()
	defer syncGD.Unlock()
	return s.PlayerMap[uuid]
}

func (ge *gateServer) AddPlayerStatus(p *player.GamePlayer) error {
	bin := &spb.PlayerStatusRedisData{
		Status:       spb.PlayerStatusType_PLAYER_STATUS_ONLINE,
		GameserverId: ge.game.AppId,
		LoginRand:    0,
		LoginTime:    time.Now().Unix(),
		Uid:          p.Uid,
		Uuid:         p.Uuid,
	}
	value, err := pb.Marshal(bin)
	if err != nil {
		logger.Error("pb marshal error: %v\n", err)
		return err
	}
	err = ge.game.Store.SetPlayerStatus(strconv.Itoa(int(p.AccountId)), value)
	return err
}
