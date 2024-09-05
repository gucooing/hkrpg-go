package gameserver

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

// gate申请离线玩家
/*
1.先通知gate离线玩家
2.拉取redis状态
3.验证是否可以保存数据然后保存数据
4.可保存数据情况下删除redis状态
5.删除玩家内存
*/
type GamePlayer struct {
	gate           *gateServer        // 玩家所属gate
	game           *GameServer        // 玩家所属game
	p              *player.GamePlayer // 玩家内存
	lastActiveTime int64              // 最近一次的通信时间
}

// 这个kill玩家不会通知给gate
func (s *GameServer) killPlayer(p *GamePlayer) {
	p.p.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_OFFLINE)
	database.DelPlayerStatus(s.Store.StatusRedis, strconv.Itoa(int(p.p.Uid)))
	s.delPlayerMap(p.p.Uid)
	logger.Info("[UID:%v]玩家下线成功", p.p.Uid)
	if p.p.SendCal != nil {
		p.p.SendCal()
	}
	if p.p.RecvCal != nil {
		p.p.RecvCal()
	}
}

/************************************接口*********************************/

func (s *GameServer) addPlayerMap(uid uint32, g *player.GamePlayer, ge *gateServer) (*GamePlayer, bool) {
	g.RouteManager = player.NewRouteManager(g)
	gamePlayer := &GamePlayer{
		gate:           ge,
		game:           s,
		p:              g,
		lastActiveTime: time.Now().Unix(),
	}
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()

	if s.playerMap[uid] == nil {
		s.playerMap[uid] = gamePlayer
		PLAYERNUM++
		return gamePlayer, true
	}
	return nil, false
}

func (s *GameServer) delPlayerMap(uid uint32) bool {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	if s.playerMap[uid] != nil {
		delete(s.playerMap, uid)
		PLAYERNUM--
		return true
	}
	return false
}

func (s *GameServer) getAllPlayer() map[uint32]*GamePlayer {
	players := make(map[uint32]*GamePlayer)
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	for uid, play := range s.playerMap {
		players[uid] = play
	}
	return players
}

func (s *GameServer) getPlayerByUid(uid uint32) *GamePlayer {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	return s.playerMap[uid]
}

func (s *GameServer) AddPlayerStatus(g *GamePlayer) error {
	bin := &spb.PlayerStatusRedisData{
		Status:       spb.PlayerStatusType_PLAYER_STATUS_ONLINE,
		GameserverId: s.AppId,
		GateserverId: g.gate.appid,
		LoginRand:    0,
		LoginTime:    time.Now().Unix(),
		Uid:          g.p.Uid,
		DataVersion:  g.p.GetPd().GetDataVersion(),
	}
	value, err := pb.Marshal(bin)
	if err != nil {
		logger.Error("pb marshal error: %v\n", err)
		return err
	}

	if ok := database.AddPlayerStatus(s.Store.StatusRedis, strconv.Itoa(int(g.p.Uid)), value); !ok {
		logger.Info("玩家状态锁加锁失败")
	}
	return err
}
