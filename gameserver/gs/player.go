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
	s.delPlayerMap(p.p.Uid)
	logger.Info("[UID:%v]玩家下线成功", p.p.Uid)
}

// 玩家数据唯一修改方法
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
	if statu.GameserverId != s.AppId && statu.DataVersion != p.GetDataVersion() {
		// 脏数据
		logger.Info("[UID:%v]数据过期，已丢弃", p.Uid)
		return
	}
	//  确认写入，更新数据版本
	p.AddDataVersion()
	dbDate := new(database.PlayerData)
	dbDate.Uid = p.Uid
	dbDate.Level = p.GetLevel()
	dbDate.Exp = p.PlayerPb.Exp
	dbDate.Nickname = p.GetNickname()
	dbDate.BinData, err = pb.Marshal(p.PlayerPb)
	dbDate.DataVersion = p.GetDataVersion()
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return
	}

	if err = s.Store.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return
	}
	if !s.SetPlayerPlayerBasicBriefData(p) {
		logger.Error("[UID:%v]玩家简要信息保存失败", p.Uid)
	}
	return
}

/************************************接口*********************************/

func (s *GameServer) addPlayerMap(uid uint32, g *player.GamePlayer, ge *gateServer) (*GamePlayer, bool) {
	gamePlayer := &GamePlayer{
		gate:         ge,
		game:         s,
		p:            g,
		RouteManager: NewRouteManager(g),
	}
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()

	if gamePlayer.p.Player == nil {
		gamePlayer.p.Player = &player.PlayerData{
			Battle: make(map[uint32]*player.Battle),
			BattleState: &player.BattleState{
				ChallengeState: &player.ChallengeState{},
			},
		}
	}
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
	for uuid, play := range s.playerMap {
		players[uuid] = play
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
		DataVersion:  g.p.GetDataVersion(),
	}
	value, err := pb.Marshal(bin)
	if err != nil {
		logger.Error("pb marshal error: %v\n", err)
		return err
	}
	if ok := s.Store.DistLockPlayerStatus(strconv.Itoa(int(g.p.AccountId)), value); !ok {
		logger.Info("玩家状态锁加锁失败")
	}
	return err
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
