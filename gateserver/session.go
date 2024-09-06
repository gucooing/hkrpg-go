package gateserver

import (
	"encoding/base64"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

type PlayerGame struct {
	gs             *gameServer
	ga             *GateServer
	Status         spb.PlayerStatus
	Uid            uint32 // uid
	AccountId      uint32
	Seed           uint64
	XorKey         []byte // 密钥
	KcpConn        *kcp.UDPSession
	LastActiveTime uint64 // 最近一次的活跃时间
	RouteManager   *RouteManager
	ticker         *time.Timer
	stop           chan struct{}
}

type HandlerFunc func(tcpMsg *alg.PackMsg)

type RouteManager struct {
	handlerFuncRouteMap map[uint16]HandlerFunc
}

func NewRouteManager(p *PlayerGame) (r *RouteManager) {
	r = new(RouteManager)
	r.initRoute(p)
	return r
}

func (r *RouteManager) initRoute(p *PlayerGame) {
	r.handlerFuncRouteMap = map[uint16]HandlerFunc{
		// cmd.PlayerHeartBeatCsReq: p.HandlePlayerHeartBeatCsReq,
		cmd.PlayerLogoutCsReq: p.PlayerLogoutCsReq,
		cmd.GetAuthkeyCsReq:   p.nilProto,
		// 好友
		cmd.ApplyFriendCsReq:  p.ApplyFriendCsReq,  // 发送好友申请
		cmd.HandleFriendCsReq: p.HandleFriendCsReq, // 处理好友申请
		cmd.SendMsgCsReq:      p.SendMsgCsReq,      // 发送聊天信息
	}
}

func (p *PlayerGame) PlayerRegisterMessage(cmdId uint16, tcpMsg *alg.PackMsg) {
	p.LastActiveTime = getCurTime()
	handlerFunc, ok := p.RouteManager.handlerFuncRouteMap[cmdId]
	if !ok {
		p.GateToGame(tcpMsg)
		return
	}
	handlerFunc(tcpMsg)
}

func (p *PlayerGame) nilProto(tcpMsg *alg.PackMsg) {
	logger.Info("", tcpMsg.ProtoData)
}

// 将玩家消息转发到game
func (p *PlayerGame) GateToGame(tcpMsg *alg.PackMsg) {
	// logger.Debug("[UID:%v]gate->game:%s", p.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(tcpMsg.CmdId))
	// msg := &spb.GateToGameMsgNotify{
	// 	Uid:   p.Uid,
	// 	CmdId: int32(tcpMsg.CmdId),
	// 	Msg:   tcpMsg.ProtoData,
	// }
	// p.gs.sendGame(cmd.GateToGameMsgNotify, msg)
}

func testMsg(cmdId uint16, B64Msg string) {
	payloadMsg, _ := base64.StdEncoding.DecodeString(B64Msg)
	protoObj := cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Warn("get new proto object is nil")
		return
	}
	err := pb.Unmarshal(payloadMsg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data NAME: %s  err: %v || b64:%s", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), err, base64.StdEncoding.EncodeToString(payloadMsg))
		return
	}
}

func testsMsg(cmdId uint16, payloadMsg []byte) {
	protoObj := cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Warn("get new proto object is nil")
		return
	}
	err := pb.Unmarshal(payloadMsg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data NAME: %s  err: %v || b64:%s", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), err, base64.StdEncoding.EncodeToString(payloadMsg))
		return
	}
	data := protojson.Format(protoObj)
	logger.Debug("S --> C : NAME: %s KcpMsg: \n%s\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
}

// 将消息发送给客户端
func (p *PlayerGame) GateToPlayer(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(cmd.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := cmd.EncodeProtoToPayload(rspMsg)
	SendHandle(p, tcpMsg)
}

func (s *GateServer) getAllPlayer() map[uint32]*PlayerGame {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	players := make(map[uint32]*PlayerGame)
	for k, v := range s.playerMap {
		players[k] = v
	}
	return players
}

func (s *GateServer) getPlayerByUid(uid uint32) *PlayerGame {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	return s.playerMap[uid]
}

func (s *GateServer) addPlayer(uid uint32, player *PlayerGame) bool {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	if s.playerMap[uid] == nil {
		s.playerMap[uid] = player
		return true
	}
	s.playerMap[uid] = player
	return false
}

func (s *GateServer) delPlayerByUid(uid uint32) bool {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	if s.playerMap[uid] != nil {
		delete(s.playerMap, uid)
		return true
	}
	return false
}

func (s *GateServer) getLoginPlayerByUid(uid uint32) *PlayerGame {
	s.loginPlayerMapLock.Lock()
	defer s.loginPlayerMapLock.Unlock()
	return s.loginPlayerMap[uid]
}

func (s *GateServer) addLoginPlayer(uid uint32, player *PlayerGame) bool {
	s.loginPlayerMapLock.Lock()
	defer s.loginPlayerMapLock.Unlock()
	if s.loginPlayerMap[uid] == nil {
		s.loginPlayerMap[uid] = player
		return true
	}
	return false
}

func (s *GateServer) delLoginPlayerByUid(uid uint32) bool {
	s.loginPlayerMapLock.Lock()
	defer s.loginPlayerMapLock.Unlock()
	if s.loginPlayerMap[uid] != nil {
		delete(s.loginPlayerMap, uid)
		return true
	}
	return false
}

// 玩家主动离线处理
func (p *PlayerGame) PlayerLogoutCsReq(tcpMsg *alg.PackMsg) {
	p.gs.gate.ttiPlayerKill(p, spb.Retcode_RET_PLAYER_LOGOUT)
}

// 玩家超时离线
func (s *GateServer) AutoDelPlayer() {
	ticker := time.NewTicker(time.Second * 30)
	for {
		<-ticker.C
		for _, play := range s.getAllPlayer() {
			if getCurTime()-play.LastActiveTime > 10000 {
				switch play.Status {
				case spb.PlayerStatus_PlayerStatus_PostLogin:
					s.ttiPlayerKill(play, spb.Retcode_RET_PLAYER_TIMEOUT)
				case spb.PlayerStatus_PlayerStatus_Logout_Wait:
					s.delPlayerByUid(play.Uid)
					logger.Debug("[UID:%v]长时间未收到gameserver的离线成功包", play.Uid)
				}
			}
		}
	}
}

// 玩家登录超时离线
func (p *PlayerGame) loginTicker() {
	select {
	case <-p.ticker.C:
		logger.Info("[UID:%v]玩家登录超时", p.Uid)
		p.GateToPlayer(cmd.PlayerKickOutScNotify, nil)
		p.KcpConn.Close()
		p.gs.gate.delLoginPlayerByUid(p.Uid)
		p.ticker.Stop()
		return
	case <-p.stop:
		p.ticker.Stop()
		return
	}
}

func (p *PlayerGame) isChannelClosed() bool {
	// 不适用于有缓存通道
	select {
	case <-p.stop:
		return true
	default:
	}

	return false
}

// 主动离线方法
/*
1.标记玩家状态为离线等待
2.关闭kcp通道
3.通知gameserver离线玩家
4.gameserver离线成功后删除玩家
5.gameserver长时间未回复交由定时器处理
*/
func (s *GateServer) ttiPlayerKill(p *PlayerGame, code spb.Retcode) {
	p.Status = spb.PlayerStatus_PlayerStatus_Logout_Wait
	p.KcpConn.Close()
	// p.gs.sendGame(cmd.GetToGamePlayerLogoutReq, &spb.GetToGamePlayerLogoutReq{
	// 	Retcode:         code,
	// 	Uid:             p.Uid,
	// 	AccountId:       p.AccountId,
	// 	OldGameServerId: p.gs.appid,
	// })
	logger.Debug("[UID:%v]玩家主动离线中,原因:%s", p.Uid, code.String())
}

// 被动离线方法
/*
1.标记玩家状态
2.通知给客户端
3.gate原因需要通知给game
4.关闭kcp通道
5.删除玩家
*/
func (s *GateServer) passPlayerKill(p *PlayerGame, code spb.Retcode) {
	if p == nil {
		return
	}
	p.Status = spb.PlayerStatus_PlayerStatus_Logout
	p.GateToPlayer(cmd.PlayerKickOutScNotify, nil)
	p.KcpConn.Close()
	s.delPlayerByUid(p.Uid)
	switch code {
	case spb.Retcode_RET_PLAYER_SYSTEM_ERROR: // 系统 异常下线
		// p.gs.sendGame(cmd.GateToGamePlayerLogoutNotify, &spb.GateToGamePlayerLogoutNotify{
		// 	Uid:       p.Uid,
		// 	AccountId: p.AccountId,
		// })
	case spb.Retcode_RET_PLAYER_GAME_LOGIN: // game通知下线
	case spb.Retcode_RET_PLAYER_GATE_REPEAT_LOGIN: // 同网关下线

	}
	logger.Debug("[UID:%v]玩家被动离线中,原因:%s", p.Uid, code.String())
}
