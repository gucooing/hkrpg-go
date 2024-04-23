package gate

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type PlayerGame struct {
	gs             *gameServer
	Status         spb.PlayerStatus
	Uid            uint32 // uid
	AccountId      uint32
	Uuid           int64 // 唯一临时uuid
	Seed           uint64
	XorKey         []byte // 密钥
	KcpConn        *kcp.UDPSession
	LastActiveTime int64 // 最近一次的活跃时间
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
		cmd.PlayerHeartBeatCsReq: p.HandlePlayerHeartBeatCsReq,
		cmd.PlayerLogoutCsReq:    p.PlayerLogoutCsReq,
		cmd.GetAuthkeyCsReq:      p.nilProto,
	}
}

func (p *PlayerGame) PlayerRegisterMessage(cmdId uint16, tcpMsg *alg.PackMsg) {
	handlerFunc, ok := p.RouteManager.handlerFuncRouteMap[cmdId]
	if !ok {
		p.GateToGame(tcpMsg)
		return
	}
	handlerFunc(tcpMsg)
	return
}

func (p *PlayerGame) nilProto(tcpMsg *alg.PackMsg) {}

// 将玩家消息转发到game
func (p *PlayerGame) GateToGame(tcpMsg *alg.PackMsg) {
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	msg := &spb.GateToGameMsgNotify{
		Uid:  p.Uid,
		Uuid: p.Uuid,
		Msg:  binMsg,
	}
	p.gs.sendGame(cmd.GateToGameMsgNotify, msg)
}

// 将消息发送给客户端
func (p *PlayerGame) GateToPlayer(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	SendHandle(p, tcpMsg)
}

func (s *GateServer) AddPlayerMap(uuid int64, player *PlayerGame) {
	s.playerMapLock.Lock()
	s.playerMap[uuid] = player
	s.playerMapLock.Unlock()
	go player.gs.AddPlayerMap(uuid, player)
}

func (gs *gameServer) AddPlayerMap(uuid int64, player *PlayerGame) {
	gs.playerMapLock.Lock()
	gs.playerMap[uuid] = player
	gs.playerMapLock.Unlock()
}

func (s *GateServer) DelPlayerMap(uuid int64) {
	s.playerMap[uuid].gs.DelPlayerMap(uuid)
	s.playerMapLock.Lock()
	if s.playerMap[uuid] != nil {
		delete(s.playerMap, uuid)
	}
	s.playerMapLock.Unlock()
}

func (gs *gameServer) DelPlayerMap(uuid int64) {
	gs.playerMapLock.Lock()
	if gs.playerMap[uuid] != nil {
		delete(gs.playerMap, uuid)
	}
	gs.playerMapLock.Unlock()
}

func (s *GateServer) GetPlayerByUuid(uuid int64) (*PlayerGame, bool) {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	player, ok := s.playerMap[uuid]
	return player, ok
}

func (s *GateServer) GetAllPlayer() map[int64]*PlayerGame {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	players := make(map[int64]*PlayerGame)
	for k, v := range s.playerMap {
		players[k] = v
	}
	return players
}

func (gs *gameServer) GetAllPlayer() map[int64]*PlayerGame {
	gs.playerMapLock.Lock()
	defer gs.playerMapLock.Unlock()
	players := make(map[int64]*PlayerGame)
	for k, v := range gs.playerMap {
		players[k] = v
	}
	return players
}

// 玩家主动离线处理
func (p *PlayerGame) PlayerLogoutCsReq(tcpMsg *alg.PackMsg) {
	p.KcpConn.Close()
	notify := &spb.GateToGamePlayerLogoutNotify{
		Uid:  p.Uid,
		Uuid: p.Uuid,
	}
	p.gs.sendGame(cmd.GateToGamePlayerLogoutNotify, notify)
	p.gs.gate.DelPlayerMap(p.Uuid)
	logger.Info("[UID:%v][UUID:%v]玩家离线成功", p.Uid, p.Uuid)
}

// 玩家超时离线
func (s *GateServer) AutoDelPlayer() {
	ticker := time.NewTicker(time.Second * 120)
	for {
		<-ticker.C
		plays := s.GetAllPlayer()
		for _, play := range plays {
			if time.Now().Unix()-play.LastActiveTime > 30 {
				play.GateToPlayer(cmd.PlayerKickOutScNotify, nil)
				play.PlayerLogoutCsReq(nil)
			}
		}
	}
}

/*
1.通知客户端下线
2.删除玩家内存
*/
func (s *GateServer) killPlayer(p *PlayerGame) {
	p.GateToPlayer(cmd.PlayerKickOutScNotify, nil)
	s.DelPlayerMap(p.Uuid)
	logger.Info("[UID:%v][UUID:%v]玩家下线gate", p.Uid, p.Uuid)
}

// 玩家登录超时离线
func (p *PlayerGame) loginTicker() {
	select {
	case <-p.ticker.C:
		logger.Info("[UID:%v][UUID:%v]玩家登录超时", p.Uid, p.Uuid)
		p.GateToPlayer(cmd.PlayerKickOutScNotify, nil)
		p.KcpConn.Close()
		p.gs.gate.DelPlayerMap(p.Uuid)
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
