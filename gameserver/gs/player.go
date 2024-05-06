package gs

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/database"
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
	gate           *gateServer
	game           *GameServer
	p              *player.GamePlayer
	RouteManager   *RouteManager
	LastActiveTime int64 // 最近一次的保存时间
}

// 这个kill玩家不会通知给gate
func (s *GameServer) killPlayer(p *GamePlayer) {
	s.upDataPlayer(p.p)
	s.Store.DistUnlockPlayerStatus(strconv.Itoa(int(p.p.AccountId)))
	p.gate.DelPlayerMap(p.p.Uuid)
	logger.Info("[UID:%v][UUID:%v]玩家下线成功", p.p.Uid, p.p.Uuid)
}

func (s *GameServer) upDataPlayer(p *player.GamePlayer) {
	redisDb, ok := s.Store.GetPlayerStatus(strconv.Itoa(int(p.AccountId)))
	if !ok {
		return
	}
	statu := new(spb.PlayerStatusRedisData)
	err := pb.Unmarshal(redisDb, statu)
	if err != nil {
		logger.Error("PlayerStatusRedisData Unmarshal error")
		s.Store.DistUnlockPlayerStatus(strconv.Itoa(int(p.AccountId)))
		return
	}
	if statu.GameserverId != GAMESERVER.AppId || statu.Uuid != p.Uuid {
		// 脏数据
		logger.Info("[UID:%v][UUID:%v]数据过期，已丢弃", p.Uid, p.Uuid)
		return
	}
	dbDate := new(database.PlayerData)
	dbDate.Uid = p.Uid
	dbDate.Level = p.PlayerPb.Level
	dbDate.Exp = p.PlayerPb.Exp
	dbDate.Nickname = p.PlayerPb.Nickname
	dbDate.BinData, err = pb.Marshal(p.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return
	}

	if err = s.Store.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return
	}
	if !s.SetPlayerPlayerBasicBriefData(p) {
		logger.Error("[UID:%v][UUID:%v]玩家简要信息保存失败", p.Uid, p.Uuid)
	}
	return
}

/************************************接口*********************************/

func (ge *gateServer) AddPlayerMap(uuid int64, g *player.GamePlayer) {
	gamePlayer := &GamePlayer{
		gate:         ge,
		game:         ge.game,
		p:            g,
		RouteManager: NewRouteManager(g),
	}
	ge.playerMapLock.Lock()
	ge.playerMap[uuid] = gamePlayer
	ge.playerMapLock.Unlock()

	if gamePlayer.p.Player == nil {
		gamePlayer.p.Player = &player.PlayerData{
			Battle: make(map[uint32]*player.Battle),
			BattleState: &player.BattleState{
				ChallengeState: &player.ChallengeState{},
			},
		}
	}
	PLAYERNUM++
}

func (ge *gateServer) DelPlayerMap(uuid int64) {
	ge.playerMapLock.Lock()
	if ge.playerMap[uuid] != nil {
		delete(ge.playerMap, uuid)
	}
	ge.playerMapLock.Unlock()
}

func (ge *gateServer) GetAllPlayer() map[int64]*GamePlayer {
	players := make(map[int64]*GamePlayer)
	ge.playerMapLock.Lock()
	defer ge.playerMapLock.Unlock()
	for uuid, play := range ge.playerMap {
		players[uuid] = play
	}
	return players
}

func (ge *gateServer) GetPlayerByUuid(uuid int64) *GamePlayer {
	ge.playerMapLock.Lock()
	defer ge.playerMapLock.Unlock()
	return ge.playerMap[uuid]
}

// 这个玩意不要过多执行，会卡死的
func (s *GameServer) GetPlayerByUuid(uuid int64) *GamePlayer {
	for _, ge := range s.gateList {
		playerList := ge.GetAllPlayer()
		if playerList[uuid] != nil {
			return playerList[uuid]
		}
	}
	return nil
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

func (s *GameServer) UpDataPlayer(g *player.GamePlayer) error {
	var err error
	if g.PlayerPb == nil {
		return nil
	}
	if g.Uid == 0 {
		return nil
	}
	if bin, ok := s.Store.GetPlayerStatus(strconv.Itoa(int(g.AccountId))); !ok {
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
	dbDate := new(database.PlayerData)
	dbDate.Uid = g.Uid
	dbDate.Level = g.PlayerPb.Level
	dbDate.Exp = g.PlayerPb.Exp
	dbDate.Nickname = g.PlayerPb.Nickname
	dbDate.BinData, err = pb.Marshal(g.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return err
	}

	if err = s.Store.UpdatePlayer(dbDate); err != nil {
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
