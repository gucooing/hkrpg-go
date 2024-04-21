package gate

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type gameServer struct {
	gate          *GateServer
	appid         uint32            // appid
	playerMap     map[uint32]uint64 // 玩家列表
	playerMapLock sync.Mutex        // 玩家列表互斥锁
	conn          net.Conn          // gs tcp通道
	playerNum     int64             // 所连接的gs玩家数

	gsChan chan struct{} // gs通道
	ticker *time.Ticker  // 定时器
	msg    chan Msg      // 玩家上发消息
}

type Msg struct {
	appId     uint32 // gs appid
	cmdId     uint16
	playerMsg pb.Message
}

func (s *GateServer) sendGs(msg *Msg) {
	gs := s.getGsByAppid(msg.appId)
	if gs == nil {
		return
	}

	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = msg.cmdId
	rspMsg.PayloadMessage = msg.playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdid error")
		return
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := gs.conn.Write(binMsg)
	if err != nil {
		logger.Debug("[GS:%v]gate->game error: %s", msg.appId, err.Error())
		return
	}
}

func (s *GateServer) getGsByAppid(appid uint32) *gameServer {
	gs := new(gameServer)
	s.gsListLock.Lock()
	gs = s.gsList[appid]
	s.gsListLock.Unlock()
	return gs
}

func (s *GateServer) getMinGsAppId() *gameServer {
	var minAppId uint32
	var minNum int
	s.gsListLock.Lock()
	for id, game := range s.gsList {
		if minAppId == 0 || minNum > len(game.playerMap) {
			minAppId = id
			minNum = len(game.playerMap)
		}
	}
	gs := s.gsList[minAppId]
	s.gsListLock.Unlock()
	return gs

}

func (s *GateServer) addGsList(gs *gameServer) {
	s.gsListLock.Lock()
	s.gsList[gs.appid] = gs
	s.gsListLock.Unlock()
}

func (s *GateServer) delGsList(appid uint32) {
	s.gsListLock.Lock()
	delete(s.gsList, appid)
	s.gsListLock.Unlock()
}

func (s *GateServer) newGs(addr string, appid uint32) {
	gameConn, err := net.Dial("tcp", addr)
	if err != nil {
		logger.Error("无法连接到GAME:", err)
		return
	}
	gs := &gameServer{
		gate:      s,
		appid:     appid,
		playerMap: make(map[uint32]uint64),
		conn:      gameConn,
		gsChan:    make(chan struct{}),
	}
	s.addGsList(gs)
	go gs.recvGame()
	gs.sendGame(cmd.GateLoginGameReq, &spb.GateLoginGameReq{
		ServerType: spb.ServerType_SERVICE_GATE,
		AppId:      s.AppId,
	})
}

// 从game接收消息
func (gs *gameServer) recvGame() {
	nodeMsg := make([]byte, PacketMaxLen)
	for {
		var bin []byte = nil
		recvLen, err := gs.conn.Read(nodeMsg)
		if err != nil {
			logger.Debug("[GS:%v]game->gate error: %s", gs.appid, err.Error())
			return
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			playerMsg := alg.DecodePayloadToProto(msg)
			gs.gameRegisterMessage(msg.CmdId, playerMsg)
		}
	}
}

func (gs *gameServer) gameRegisterMessage(cmdId uint16, playerMsg pb.Message) {
	switch cmdId {
	case cmd.GateLoginGameRsp:
		gs.GateLoginGameRsp(playerMsg) // gate在game注册回复包
	case cmd.GateGamePingRsp:
		gs.GateGamePingRsp(playerMsg) // gate发送给gs的ping回复包
	case cmd.GateGamePlayerLoginRsp:
		gs.GateGamePlayerLoginRsp(playerMsg) // game玩家登录成功通知
	case cmd.GameToGateMsgNotify:
		gs.GameToGateMsgNotify(playerMsg)

	default:
		logger.Error("game -> gate register error, cmdId:%v", cmdId)
	}
}

// 发送到game
func (gs *gameServer) sendGame(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdid error")
		return
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := gs.conn.Write(binMsg)
	if err != nil {
		logger.Debug("[GS:%v]gate->game error: %s", gs.appid, err.Error())
		return
	}
}

func (gs *gameServer) GateLoginGameRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.GateLoginGameRsp)
	if rsp.Retcode != 0 {
		// TODO 销毁这个gs连接
		return
	}
	// 注册成功，将gs放入可连接列表
	gs.gate.addGsList(gs)
	gs.ticker = time.NewTicker(5 * time.Second)
	logger.Info("gate在game:[%v]注册成功", gs.appid)
	go gs.gsMsgTo()
}

// 同一个gs的玩家共用一个协程
func (gs *gameServer) gsMsgTo() {
	for {
		select {
		case <-gs.ticker.C:
			gs.GateGamePingReq() // ping包
		}
	}
}

func (gs *gameServer) GateGamePingReq() {
	req := &spb.GateGamePingReq{
		GateServerTime: time.Now().Unix(),
	}
	gs.sendGame(cmd.GateGamePingReq, req)
}

func (gs *gameServer) GateGamePingRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.GateGamePingRsp)
	gs.playerNum = rsp.PlayerNum
}

func (gs *gameServer) GateGamePlayerLoginReq(uid uint32, uuid int64) {
	req := &spb.GateGamePlayerLoginReq{
		Uid:  uid,
		Uuid: uuid,
	}
	gs.sendGame(cmd.GateGamePlayerLoginReq, req)
}

func (gs *gameServer) GateGamePlayerLoginRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.GateGamePlayerLoginRsp)
	if player, ok := gs.gate.GetPlayerByUuid(rsp.Uuid); !ok {
		return
	} else {
		if player.gs.appid != gs.appid {
			return
		}
		prsp := &proto.PlayerGetTokenScRsp{
			SecretKeySeed: player.Seed,
			BlackInfo:     &proto.BlackInfo{},
			Uid:           player.Uid,
			Msg:           "",
			Retcode:       0,
		}

		player.Status = spb.PlayerStatus_PlayerStatus_PostLogin
		GateToPlayer(player, cmd.PlayerGetTokenScRsp, prsp)
		// 结束定时器
		player.closeStop()
		logger.Info("[AccountId:%v][UUID:%v]|[UID:%v]登录gate", player.AccountId, player.Uuid, player.Uid)
	}
}

func (gs *gameServer) GameToGateMsgNotify(playerMsg pb.Message) {
	notify := playerMsg.(*spb.GameToGateMsgNotify)
	if player, ok := gs.gate.GetPlayerByUuid(notify.Uuid); !ok {
		return
	} else {
		msgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(notify.Msg, &msgList, nil)
		for _, msg := range msgList {
			SendHandle(player, msg)
		}
	}
}

/*
// 将game消息转发到玩家
func (p *PlayerGame) GameToGate(cmdId uint16, playerMsg pb.Message) {
	rsp := playerMsg.(*spb.PlayerToGameByGateRsp)
	playerMsgList := make([]*alg.PackMsg, 0)
	alg.DecodeBinToPayload(rsp.PlayerBin, &playerMsgList, nil)
	for _, msg := range playerMsgList {
		// 发到玩家
		// logger.Debug("[S->C][UID:%v][CMDID:%v]", p.Uid, msg.CmdId)
		if msg.CmdId == cmd.PlayerLoginScRsp {
			p.playerLoginUp()
		}
		SendHandle(p, msg)

*/

func GateToPlayer(p *PlayerGame, cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	SendHandle(p, tcpMsg)
}

/******************************************NewLogin***************************************/

func (p *PlayerGame) AddPlayerStatus() error {
	bin := &spb.PlayerStatusRedisData{
		Status:       spb.PlayerStatusType_PLAYER_STATUS_ONLINE,
		GameserverId: p.gs.appid,
		LoginRand:    p.Seed,
		LoginTime:    0,
		Uid:          p.Uid,
	}
	value, err := pb.Marshal(bin)
	if err != nil {
		logger.Error("pb marshal error: %v\n", err)
		return err
	}
	err = GATESERVER.Store.SetPlayerStatus(strconv.Itoa(int(p.AccountId)), value)
	return err
}

// gate请求gs离线玩家
func (p *PlayerGame) gateToGsPlayerLogoutReq() {
	/*
		req := &spb.PlayerLogoutReq{
			Uuid:          p.Uuid,
			AccountId:     p.AccountId,
			Uid:           p.Uid,
			OfflineReason: spb.PlayerOfflineReason_OFFLINE_NONE,
		}

		p.sendGame(cmd.PlayerLogoutReq, req)
	*/
}

// gs回应玩家离线
func (p *PlayerGame) gsToGamePlayerLogoutRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.PlayerLogoutRsp)
	if rsp.Retcode != spb.Retcode_RET_SUCC || rsp.Uid != p.Uid || rsp.Uuid != p.Uuid || rsp.AccountId != p.AccountId {
		logger.Info("[gs->gate][UID%v]GS离线失败", p.Uid)
	}
	GateToPlayer(p, cmd.PlayerKickOutScNotify, nil)
	KickPlayer(p)
}

func (p *PlayerGame) closeStop() {
	if !p.isChannelClosed() {
		close(p.stop)
	}
}

func (p *PlayerGame) playerLoginUp() {
	// 登录成功设置
	p.closeStop()

	// 解锁
	GATESERVER.Store.DistUnlock(strconv.Itoa(int(p.AccountId)))
	p.AddPlayerStatus()

	// 通知node玩家登录
	GATESERVER.sendNode(cmd.PlayerLoginNotify, &spb.PlayerLoginNotify{
		Uuid:            p.Uuid,
		AccountId:       p.AccountId,
		Uid:             p.Uid,
		GateServerAppId: GATESERVER.AppId,
		GameServerAppId: p.gs.appid,
	})
}
