package gameserver

import (
	"context"
	"strconv"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

type gateServer struct {
	game  *GameServer
	appid uint32
	conn  *gunet.TcpConn // gate tcp通道
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

func (s *GameServer) newGate(conn *gunet.TcpConn, appid uint32) {
	ge := &gateServer{
		game:  s,
		appid: appid,
		conn:  conn,
	}
	s.addGeList(ge)
	ge.seedGate(cmd.GateLoginGameRsp, &spb.GateLoginGameRsp{
		Retcode: 0,
	})
	logger.Info("gate:[%v]在game注册成功", appid)
	ge.recvGate()
}

// 从gate接收消息
func (ge *gateServer) recvGate() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker gate: %v", ge.appid)
			ge.killGate()
		}
	}()

	for {
		bin, err := ge.conn.Read()
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			ge.killGate()
			return
		}
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := cmd.DecodePayloadToProto(msg)
			go ge.gateRegisterMessage(msg.CmdId, serviceMsg)
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
		ge.GateToGamePlayerLogoutNotify(payloadMsg)
	case cmd.GateToGameMsgNotify:
		ge.GateToGameMsgNotify(payloadMsg) // gate转发客户端消息到gs
	default:
		logger.Error("gate -> game cmdid error: %v", cmdId)
	}
}

func (ge *gateServer) seedGate(cmdId uint16, payloadMsg pb.Message) {
	rspMsg := new(cmd.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = payloadMsg
	tcpMsg := cmd.EncodeProtoToPayload(rspMsg)
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
	}
	if req.Uid == 0 || req.AccountId == 0 {
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
	p := ge.NewPlayer(req.Uid, req.AccountId)
	// 拉取账户数据
	go p.RecvMsg()
	go ge.recvPlayer(p)
	p.GetPlayerDateByDb()
	g, ok := ge.game.addPlayerMap(req.Uid, p, ge)
	if !ok {
		logger.Warn("[UID:%v]超出预期的玩家重复登录", p.Uid)
		return
	}
	logger.Info("[UID:%v]登录game", p.Uid)
	ge.game.AddPlayerStatus(g)
	ge.GateGamePlayerLoginRsp(rsp)
}

func (ge *gateServer) recvPlayer(p *player.GamePlayer) {
	for {
		select {
		case bin := <-p.SendChan:
			switch bin.MsgType {
			case player.Server:
				ge.sendPlayer(p, bin.CmdId, bin.PlayerMsg)
			}
		case <-p.SendCtx.Done():
			p.IsClosed = true
			return
		}
	}
}

func (ge *gateServer) sendPlayer(p *player.GamePlayer, cmdId uint16, playerMsg pb.Message) {
	if p.IsClosed {
		return
	}
	rspMsg := new(cmd.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	kcpMsg := cmd.EncodeProtoToPayload(rspMsg)
	// logger.Debug("[UID:%v]game->gate:%s", p.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId))
	ge.GameToGateMsgNotify(&spb.GameToGateMsgNotify{
		Uid:   p.Uid,
		CmdId: int32(cmdId),
		Msg:   kcpMsg.ProtoData,
	})
}

func (ge *gateServer) GateGamePlayerLoginRsp(rsp *spb.GateGamePlayerLoginRsp) {
	ge.seedGate(cmd.GateGamePlayerLoginRsp, rsp)
}

func (ge *gateServer) GetToGamePlayerLogoutReq(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GetToGamePlayerLogoutReq)
	play := ge.game.getPlayerByUid(req.Uid)
	if play == nil {
		logger.Info("[UID:%v]没有找到此玩家", req.Uid)
		database.DelPlayerStatus(ge.game.Store.StatusRedis, strconv.Itoa(int(req.Uid)))
		ge.playerRepeatLogin(req)
		return
	}
	ge.game.killPlayer(play)
	logger.Info("[UID:%v][AccountId:%v]下线玩家,原因:%s", req.Uid, req.AccountId, req.Retcode.String())
	switch req.Retcode {
	case spb.Retcode_RET_PLAYER_REPEAT_LOGIN: // 异网关重复登录
		play.gate.seedGate(cmd.GameToGatePlayerLogoutNotify, &spb.GameToGatePlayerLogoutNotify{
			Uid: play.p.Uid,
		})
		ge.playerRepeatLogin(req)
	case spb.Retcode_RET_PLAYER_GATE_REPEAT_LOGIN: // 同网关重复登录
		ge.playerRepeatLogin(req)
	case spb.Retcode_RET_PLAYER_TIMEOUT: // 超时
		ge.playerRepeatLogin(req)
	case spb.Retcode_RET_PLAYER_LOGOUT: // 正常离线
		ge.playerRepeatLogin(req)
	default:
		return
	}
}

// 玩家重复登录处理
func (ge *gateServer) playerRepeatLogin(req *spb.GetToGamePlayerLogoutReq) {
	rsp := &spb.GetToGamePlayerLogoutRsp{
		Retcode:         spb.Retcode_RET_SUCC,
		Uid:             req.Uid,
		NewGameServerId: req.NewGameServerId,
	}
	ge.seedGate(cmd.GetToGamePlayerLogoutRsp, rsp)
}

func (ge *gateServer) NewPlayer(uid, accountId uint32) *player.GamePlayer {
	g := new(player.GamePlayer)
	g.Uid = uid
	g.AccountId = accountId
	g.SendChan = make(chan player.Msg, 10)
	g.RecvChan = make(chan player.Msg, 10)
	g.SendCtx, g.SendCal = context.WithCancel(context.Background())
	g.RecvCtx, g.RecvCal = context.WithCancel(context.Background())
	g.IsJumpMission = ge.game.Config.IsJumpMission
	g.Store = ge.game.Store

	return g
}

func (ge *gateServer) GateToGameMsgNotify(payloadMsg pb.Message) {
	notify := payloadMsg.(*spb.GateToGameMsgNotify)
	paler := ge.game.getPlayerByUid(notify.Uid)
	if paler != nil {
		// logger.Debug("[UID:%v]gate->game:%s", paler.p.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(uint16(notify.CmdId)))
		paler.lastActiveTime = time.Now().Unix()
		playerMsg := cmd.DecodePayloadToProto(&alg.PackMsg{
			CmdId:     uint16(notify.CmdId),
			HeadData:  make([]byte, 0),
			ProtoData: notify.Msg,
		})
		if playerMsg == nil {
			logger.Warn("[UID:%v]DecodePayloadToProto error", paler.p.Uid)
		}
		if paler.p.RecvChan != nil {
			paler.p.RecvChan <- player.Msg{
				CmdId:     uint16(notify.CmdId),
				MsgType:   player.Client,
				PlayerMsg: playerMsg,
			}
			if paler.p.IsClosed {
				close(paler.p.RecvChan)
			}
		}
	}
}

func (ge *gateServer) GameToGateMsgNotify(payloadMsg pb.Message) {
	ge.seedGate(cmd.GameToGateMsgNotify, payloadMsg)
}

func (ge *gateServer) GateToGamePlayerLogoutNotify(payloadMsg pb.Message) {
	notify := payloadMsg.(*spb.GateToGamePlayerLogoutNotify)
	play := ge.game.getPlayerByUid(notify.Uid)
	if play == nil {
		return
	} else {
		// 下线玩家
		ge.game.killPlayer(play)
	}
}

// gate离线
func (ge *gateServer) killGate() {
	plays := ge.game.getAllPlayer()
	for _, play := range plays {
		if play.gate.appid == ge.appid {
			ge.game.killPlayer(play)
		}
	}
	ge.game.delGeList(ge.appid)
	logger.Info("[APPID:%v]gate server离线", ge.appid)
}
