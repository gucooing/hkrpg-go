package gs

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type gateServer struct {
	game          *GameServer
	appid         uint32
	playerMap     map[int64]*GamePlayer // 玩家列表
	playerMapLock sync.Mutex            // 玩家列表互斥锁
	conn          *gunet.TcpConn        // gate tcp通道

	msgChan          chan player.Msg // 消息通道
	recvPlayerCancel context.CancelFunc
}

func (s *GameServer) addGeList(ge *gateServer) {
	s.gateListLock.Lock()
	s.gateList[ge.appid] = ge
	s.gateListLock.Unlock()
}

func (s *GameServer) delGeList(appid uint32) {
	s.gateListLock.Lock()
	delete(s.gateList, appid)
	s.gateListLock.Unlock()
}

func (s *GameServer) getGeByAppid(appid uint32) *gateServer {
	s.gateListLock.Lock()
	defer s.gateListLock.Unlock()
	return s.gateList[appid]
}

// 从gate接收消息
func (s *GameServer) recvGate(conn *gunet.TcpConn, appid uint32, tmp []byte) {
	ge := &gateServer{
		game:      s,
		appid:     appid,
		playerMap: make(map[int64]*GamePlayer),
		conn:      conn,
	}
	s.addGeList(ge)
	rsp := &spb.GateLoginGameRsp{
		Retcode: 0,
	}
	ge.seedGate(cmd.GateLoginGameRsp, rsp)
	ge.msgChan = make(chan player.Msg, 10)
	recvPlayerCtx, recvPlayerCancel := context.WithCancel(context.Background())
	ge.recvPlayerCancel = recvPlayerCancel
	go ge.recvPlayer(recvPlayerCtx)
	logger.Info("gate:[%v]在game注册成功", appid)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker gate: %v", appid)
			ge.conn.Close()
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

// 接收player传来的消息
func (ge *gateServer) recvPlayer(recvPlayerCtx context.Context) {
	for {
		select {
		case bin := <-ge.msgChan:
			ge.playerToGame(bin)
		case <-recvPlayerCtx.Done():
			close(ge.msgChan)
			return
		}
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
	case cmd.GateToGamePlayerLogoutNotify:
		ge.GateToGamePlayerLogoutNotify(payloadMsg) // gate直接向目标game申请下线玩家通知
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
	logger.Info("[UID:%v][AccountId:%v]收到登录通知", req.Uid, req.AccountId)
	rsp := &spb.GateGamePlayerLoginRsp{
		Retcode: spb.Retcode_RET_SUCC,
		Uid:     req.Uid,
		Uuid:    req.Uuid,
	}
	if req.Uid == 0 || req.Uuid == 0 || req.AccountId == 0 {
		logger.Error("player login uid or uuid error")
		rsp.Retcode = spb.Retcode_RET_PLAYER_ID_ERR
		ge.GateGamePlayerLoginRsp(rsp)
		return
	}
	if ge.game.node == nil {
		logger.Error("player login node error")
		rsp.Retcode = spb.Retcode_RET_NODE_ERR
		ge.GateGamePlayerLoginRsp(rsp)
		return
	}
	g := NewPlayer(req.Uid, req.AccountId, req.Uuid, ge.msgChan)
	// 拉取账户数据
	ge.GetPlayerDate(req.Uid, g)
	ge.AddPlayerMap(req.Uuid, g)
	logger.Info("[UID:%v]|[UUID:%v]登录game", g.Uid, req.Uuid)
	ge.game.AddPlayerStatus(g)
	ge.GateGamePlayerLoginRsp(rsp)
}

func (ge *gateServer) GetPlayerDate(accountId uint32, g *player.GamePlayer) {
	var err error
	dbPlayer := ge.game.Store.QueryAccountUidByFieldPlayer(accountId)
	if dbPlayer == nil || dbPlayer.BinData == nil {
		dbPlayer = new(database.PlayerData)
		logger.Info("新账号登录，进入初始化流程")
		g.PlayerPb = g.NewPlayer()
		// 初始化完毕保存账号数据
		dbPlayer.Uid = g.Uid
		dbPlayer.BinData, err = pb.Marshal(g.PlayerPb)
		if err != nil {
			logger.Error("pb marshal error: %v", err)
		}

		err = ge.game.Store.AddDatePlayerFieldByFieldName(dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败")
			return
		}
	} else {
		g.PlayerPb = new(spb.PlayerBasicCompBin)
		err = pb.Unmarshal(dbPlayer.BinData, g.PlayerPb)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			g.PlayerPb = g.NewPlayer()
			return
		}
	}
}

func (ge *gateServer) GateGamePlayerLoginRsp(rsp *spb.GateGamePlayerLoginRsp) {
	ge.seedGate(cmd.GateGamePlayerLoginRsp, rsp)
}

func (ge *gateServer) GetToGamePlayerLogoutReq(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GetToGamePlayerLogoutReq)
	logger.Info("[UID:%v][AccountId:%v]重复登录，下线玩家中", req.Uid, req.AccountId)
	play := ge.game.GetPlayerByUuid(req.OldUuid)
	if play == nil {
		ge.game.Store.DistUnlockPlayerStatus(strconv.Itoa(int(req.AccountId)))
	} else {
		ge.seedGate(cmd.GameToGatePlayerLogoutNotify, &spb.GameToGatePlayerLogoutNotify{
			Uid:  play.p.Uid,
			Uuid: play.p.Uuid,
		})
		// 下线玩家
		ge.game.killPlayer(play)
	}

	rsp := &spb.GetToGamePlayerLogoutRsp{
		Retcode:         spb.Retcode_RET_SUCC,
		Uid:             req.Uid,
		NewUuid:         req.NewUuid,
		NewGameServerId: req.NewGameServerId,
	}
	ge.seedGate(cmd.GetToGamePlayerLogoutRsp, rsp)
}

func (ge *gateServer) GateToGamePlayerLogoutNotify(payloadMsg pb.Message) {
	notify := payloadMsg.(*spb.GateToGamePlayerLogoutNotify)
	play := ge.game.GetPlayerByUuid(notify.Uuid)
	if play == nil {
		ge.game.Store.DistUnlockPlayerStatus(strconv.Itoa(int(notify.AccountId)))
	} else {
		// 下线玩家
		ge.game.killPlayer(play)
	}
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
			RegisterMessage(msg.CmdId, msg.ProtoData, paler)
		}
	}
}

func (ge *gateServer) GameToGateMsgNotify(payloadMsg pb.Message) {
	ge.seedGate(cmd.GameToGateMsgNotify, payloadMsg)
}

// gate离线
func (ge *gateServer) killGate() {
	plays := ge.GetAllPlayer()
	for _, play := range plays {
		ge.game.killPlayer(play)
	}
	ge.game.delGeList(ge.appid)
	ge.recvPlayerCancel()
	logger.Info("[APPID:%v]gate server离线", ge.appid)
}
