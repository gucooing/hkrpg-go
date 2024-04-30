package gs

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
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
	gate         *gateServer
	game         *GameServer
	p            *player.GamePlayer
	RouteManager *RouteManager
}

func (s *GameServer) killPlayer(p *player.GamePlayer) {
	s.upDataPlayer(p)
	db.DBASE.DistUnlockPlayerStatus(strconv.Itoa(int(p.AccountId)))
	s.DelPlayerMap(p.Uuid)
	logger.Info("[UID:%v][UUID:%v]玩家下线成功", p.Uid, p.Uuid)
}

func (s *GameServer) upDataPlayer(p *player.GamePlayer) {
	redisDb, ok := db.DBASE.GetPlayerStatus(strconv.Itoa(int(p.AccountId)))
	if !ok {
		return
	}
	statu := new(spb.PlayerStatusRedisData)
	err := pb.Unmarshal(redisDb, statu)
	if err != nil {
		logger.Error("PlayerStatusRedisData Unmarshal error")
		db.DBASE.DistUnlockPlayerStatus(strconv.Itoa(int(p.AccountId)))
		return
	}
	if statu.GameserverId != GAMESERVER.AppId || statu.Uuid != p.Uuid {
		// 脏数据
		logger.Info("[UID:%v][UUID:%v]数据过期，已丢弃", p.Uid, p.Uuid)
		return
	}
	dbDate := new(db.PlayerData)
	dbDate.Uid = p.Uid
	dbDate.Level = p.PlayerPb.Level
	dbDate.Exp = p.PlayerPb.Exp
	dbDate.Nickname = p.PlayerPb.Nickname
	dbDate.BinData, err = pb.Marshal(p.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return
	}

	if err = db.DBASE.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return
	}
	if !s.SetPlayerPlayerBasicBriefData(p) {
		logger.Error("[UID:%v][UUID:%v]玩家简要信息保存失败", p.Uid, p.Uuid)
	}
	return
}

/************************************接口*********************************/

func (s *GameServer) AddPlayerMap(uuid int64, g *player.GamePlayer, ge *gateServer) {
	syncGD.Lock()
	gamePlayer := &GamePlayer{
		gate:         ge,
		game:         s,
		p:            g,
		RouteManager: NewRouteManager(g),
	}
	s.PlayerMap[uuid] = gamePlayer
	// 初始化在线数据
	if s.PlayerMap[g.Uuid].p.Player == nil {
		s.PlayerMap[g.Uuid].p.Player = &player.PlayerData{
			Battle: make(map[uint32]*player.Battle),
			BattleState: &player.BattleState{
				ChallengeState: &player.ChallengeState{},
			},
		}
	}
	PLAYERNUM = int64(len(s.PlayerMap))
	syncGD.Unlock()
	go func() {
		ge.AddPlayerMap(uuid, gamePlayer)
	}()
}

func (ge *gateServer) AddPlayerMap(uuid int64, playerGame *GamePlayer) {
	ge.playerMapLock.Lock()
	ge.playerMap[uuid] = playerGame
	ge.playerMapLock.Unlock()
}

func (s *GameServer) DelPlayerMap(uuid int64) {
	syncGD.Lock()
	if s.PlayerMap[uuid] != nil {
		func() {
			if ge := s.getGeByAppid(s.PlayerMap[uuid].p.GateAppId); ge != nil {
				ge.DelPlayerMap(uuid)
			}
		}()
		delete(s.PlayerMap, uuid)
	}
	PLAYERNUM = int64(len(s.PlayerMap))
	syncGD.Unlock()
}

func (ge *gateServer) DelPlayerMap(uuid int64) {
	ge.playerMapLock.Lock()
	if ge.playerMap[uuid] != nil {
		delete(ge.playerMap, uuid)
	}
	ge.playerMapLock.Unlock()
}

func (s *GameServer) GetAllPlayer() map[int64]*GamePlayer {
	players := make(map[int64]*GamePlayer)
	syncGD.Lock()
	defer syncGD.Unlock()
	for uuid, play := range s.PlayerMap {
		players[uuid] = play
	}
	return players
}

func (ge *gateServer) GetAllPlayer() map[int64]*GamePlayer {
	players := make(map[int64]*GamePlayer)
	syncGD.Lock()
	defer syncGD.Unlock()
	for uuid, play := range ge.playerMap {
		players[uuid] = play
	}
	return players
}

func (s *GameServer) GetPlayerByUuid(uuid int64) *GamePlayer {
	syncGD.Lock()
	defer syncGD.Unlock()
	return s.PlayerMap[uuid]
}

func (s *GameServer) AddPlayerStatus(p *player.GamePlayer) error {
	bin := &spb.PlayerStatusRedisData{
		Status:       spb.PlayerStatusType_PLAYER_STATUS_ONLINE,
		GameserverId: s.AppId,
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
	if ok := s.Store.DistLockPlayerStatus(strconv.Itoa(int(p.AccountId)), value); !ok {
		logger.Info("玩家状态锁加锁失败")
	}
	return err
}

func (s *GameServer) KickPlayer(g *player.GamePlayer) {
	if err := s.UpDataPlayer(g); err != nil {
		logger.Error("[UID:%v]保存数据失败", g.Uid)
	}
	GAMESERVER.DelPlayerMap(g.Uuid)
	if g.GateConn != nil {
		g.GateConn.Close()
	}
	logger.Info("[UID:%v]玩家离线game", g.Uid)
}

func (s *GameServer) UpDataPlayer(g *player.GamePlayer) error {
	var err error
	if g.PlayerPb == nil {
		return nil
	}
	if g.Uid == 0 {
		return nil
	}
	if bin, ok := db.DBASE.GetPlayerStatus(strconv.Itoa(int(g.AccountId))); !ok {
		return nil
	} else {
		statu := new(spb.PlayerStatusRedisData)
		err := pb.Unmarshal(bin, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			return err
		}
		if statu.GameserverId != GAMESERVER.AppId || statu.Uuid != g.Uuid {
			// 脏数据
			return nil
		}
	}
	dbDate := new(db.PlayerData)
	dbDate.Uid = g.Uid
	dbDate.Level = g.PlayerPb.Level
	dbDate.Exp = g.PlayerPb.Exp
	dbDate.Nickname = g.PlayerPb.Nickname
	dbDate.BinData, err = pb.Marshal(g.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return err
	}

	if err = db.DBASE.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return err
	}
	if !s.SetPlayerPlayerBasicBriefData(g) {
		logger.Error("[UID:%v][UUID:%v]玩家简要信息保存失败", g.Uid, g.Uuid)
	}
	return nil
}

func (s *GameServer) SetPlayerPlayerBasicBriefData(g *player.GamePlayer) bool {
	playerBasicBrief := &spb.PlayerBasicBriefData{
		Nickname:          g.GetNickname(),
		Level:             g.GetLevel(),
		WorldLevel:        g.GetWorldLevel(),
		LastLoginTime:     time.Now().Unix(),
		HeadImageAvatarId: g.GetHeadIcon(),
		Exp:               g.GetPlayerPb().Exp,
		PlatformType:      0,
		Uid:               g.Uid,
	}

	bin, err := pb.Marshal(playerBasicBrief)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return false
	}

	return s.Store.SetPlayerPlayerBasicBriefData(g.Uid, bin)
}
