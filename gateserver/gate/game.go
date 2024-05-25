package gate

import (
	"context"
	"strconv"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type gameServer struct {
	gate         *GateServer
	appid        uint32         // appid
	conn         *gunet.TcpConn // gs tcp通道
	playerNum    int64          // 所连接的gs玩家数
	tickerCancel context.CancelFunc
	ticker       *time.Ticker // 定时器
}

func (s *GateServer) newGs(addr string, appid uint32) {
	gameConn, err := gunet.NewTcpC(addr)
	if err != nil {
		logger.Error("无法连接到GAME:", err)
		return
	}
	gs := &gameServer{
		gate:  s,
		appid: appid,
		conn:  gameConn,
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
	for {
		bin, err := gs.conn.Read()
		if err != nil {
			gs.gameKill()
			return
		}
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			go gs.gameRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

// gameserver离线时
func (gs *gameServer) gameKill() {
	plays := gs.gate.getAllPlayer()
	for _, play := range plays {
		if play.gs.appid == gs.appid {
			gs.gate.passPlayerKill(play, spb.Retcode_RET_PLAYER_GAME_LOGIN)
		}
	}
	if gs.tickerCancel != nil {
		gs.tickerCancel()
	}
	gs.gate.delGsList(gs.appid)
	logger.Info("[APPID:%v]game server离线", gs.appid)
}

func (gs *gameServer) GateLoginGameRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.GateLoginGameRsp)
	if rsp.Retcode != 0 {
		gs.conn.Close()
		return
	}
	// 注册成功，将gs放入可连接列表
	gs.gate.addGsList(gs)
	gs.ticker = time.NewTicker(5 * time.Second)
	tickerCtx, tickerCancel := context.WithCancel(context.Background())
	gs.tickerCancel = tickerCancel
	logger.Info("gate在game:[%v]注册成功", gs.appid)
	go gs.gsTicker(tickerCtx)
}

// gs定时器
func (gs *gameServer) gsTicker(tickerCtx context.Context) {
	for {
		select {
		case <-gs.ticker.C:
			gs.GateGamePingReq() // ping包
		case <-tickerCtx.Done():
			gs.ticker.Stop()
			return
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
	case cmd.GetToGamePlayerLogoutRsp:
		gs.GetToGamePlayerLogoutRsp(playerMsg) // gate直接向目标game申请下线玩家回复
	case cmd.GameToGatePlayerLogoutNotify:
		gs.GameToGatePlayerLogoutNotify(playerMsg) // game告知gate玩家要下线了
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

func (s *GateServer) getGsByAppid(appid uint32) *gameServer {
	gs := new(gameServer)
	s.gsListLock.Lock()
	gs = s.gsList[appid]
	s.gsListLock.Unlock()
	return gs
}

func (s *GateServer) getMinGsAppId() *gameServer {
	if s.node == nil {
		return nil
	}
	var minAppId uint32
	var minNum int64
	s.gsListLock.Lock()
	for id, game := range s.gsList {
		if minAppId == 0 || minNum > game.playerNum {
			minAppId = id
			minNum = game.playerNum
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
	if s.gsList[appid] != nil {
		delete(s.gsList, appid)
	}
	s.gsListLock.Unlock()
}

// gs ping 请求
func (gs *gameServer) GateGamePingReq() {
	req := &spb.GateGamePingReq{
		GateServerTime: time.Now().Unix(),
	}
	gs.sendGame(cmd.GateGamePingReq, req)
}

// gs ping 回复
func (gs *gameServer) GateGamePingRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.GateGamePingRsp)
	gs.playerNum = rsp.PlayerNum
}

// 玩家在gs注册请求
func (gs *gameServer) GateGamePlayerLoginReq(uid, accountId uint32) {
	logger.Debug("[UID:%v][AccountId:%v]发送登录通知", uid, accountId)
	req := &spb.GateGamePlayerLoginReq{
		Uid:       uid,
		AccountId: accountId,
	}
	gs.sendGame(cmd.GateGamePlayerLoginReq, req)
}

// 玩家在gs注册回复
func (gs *gameServer) GateGamePlayerLoginRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.GateGamePlayerLoginRsp)
	switch rsp.Retcode {
	case spb.Retcode_RET_PLAYER_ID_ERR:
		logger.Debug("由于玩家id丢失，玩家无法登录")
		return
	case spb.Retcode_RET_NODE_ERR:
		logger.Debug("由于node意外离线，玩家无法登录")
		return
	}
	player := gs.gate.getLoginPlayerByUid(rsp.Uid)
	if player == nil {
		logger.Warn("[UID:%v]不存在此玩家", rsp.Uid)
		return
	}
	if player.gs.appid != gs.appid {
		logger.Warn("不存在此gameserver")
		return
	}
	// 删除登录玩家
	gs.gate.delLoginPlayerByUid(rsp.Uid)
	// 将玩家添加到已登录玩家列表
	if !gs.gate.addPlayer(rsp.Uid, player) {
		logger.Warn("[UID:%v]超出预期的玩家重复登录", rsp.Uid)
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
	player.GateToPlayer(cmd.PlayerGetTokenScRsp, prsp)
	// 结束定时器
	player.closeStop()
	// 删除登录锁
	gs.gate.Store.DistUnlock(strconv.Itoa(int(player.AccountId)))
	logger.Info("[AccountId:%v][UID:%v]登录gate", player.AccountId, player.Uid)
}

func (p *PlayerGame) closeStop() {
	if !p.isChannelClosed() {
		close(p.stop)
	}
}

// gs玩家下线回复
func (gs *gameServer) GetToGamePlayerLogoutRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.GetToGamePlayerLogoutRsp)
	if rsp.Retcode != 0 {
		return
	}
	play := gs.gate.getPlayerByUid(rsp.Uid)
	loginPlay := gs.gate.getLoginPlayerByUid(rsp.Uid)
	if play != nil {
		switch play.Status {
		case spb.PlayerStatus_PlayerStatus_LoggingIn: // 登录中收到下线，肯定是重复登录下线回复
			logger.Warn("[UID:%v]🖥️🦐不是，哥们！你登录流程都没跑完怎么收到的下线通知?", rsp.Uid)
		case spb.PlayerStatus_PlayerStatus_PostLogin: // 已登录状态收到下线，滚
			gs.gate.passPlayerKill(play, spb.Retcode_RET_PLAYER_GATE_REPEAT_LOGIN)
		case spb.PlayerStatus_PlayerStatus_Logout_Wait: // 离线等待中收到下线
			play.Status = spb.PlayerStatus_PlayerStatus_Logout
			gs.gate.delPlayerByUid(play.Uid)
		}
	}
	// 登录中收到下线，肯定是重复登录下线回复
	if loginPlay != nil {
		newGs := gs.gate.getGsByAppid(rsp.NewGameServerId)
		if newGs == nil {
			return
		}
		newGs.playerLogin(loginPlay)
	}

	logger.Debug("[UID:%v]下线玩家成功", rsp.Uid)
}

// game通知gate玩家消息
func (gs *gameServer) GameToGateMsgNotify(playerMsg pb.Message) {
	notify := playerMsg.(*spb.GameToGateMsgNotify)
	player := gs.gate.getPlayerByUid(notify.Uid)
	if player == nil {
		return
	}
	msgList := make([]*alg.PackMsg, 0)
	alg.DecodeBinToPayload(notify.Msg, &msgList, nil)
	for _, msg := range msgList {
		SendHandle(player, msg)
	}
}

// game通知gate玩家下线
func (gs *gameServer) GameToGatePlayerLogoutNotify(playerMsg pb.Message) {
	notify := playerMsg.(*spb.GameToGatePlayerLogoutNotify)
	if play := gs.gate.getPlayerByUid(notify.Uid); play != nil {
		gs.gate.passPlayerKill(play, spb.Retcode_RET_PLAYER_GAME_LOGIN)
	}
}
