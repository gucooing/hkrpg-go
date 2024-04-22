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
func (s *GameServer) killPlayer(p *player.GamePlayer) {
	s.upDataPlayer(p)
	db.DBASE.DelPlayerStatus(strconv.Itoa(int(p.AccountId)))
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
		db.DBASE.DelPlayerStatus(strconv.Itoa(int(p.AccountId)))
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
	return
}

/************************************接口*********************************/

func (s *GameServer) AddPlayerMap(uuid int64, g *player.GamePlayer) {
	syncGD.Lock()
	s.PlayerMap[uuid] = g
	// 初始化在线数据
	if s.PlayerMap[g.Uuid].Player == nil {
		s.PlayerMap[g.Uuid].Player = &player.PlayerData{
			Battle: make(map[uint32]*player.Battle),
			BattleState: &player.BattleState{
				ChallengeState: &player.ChallengeState{},
			},
		}
	}
	syncGD.Unlock()
	go func() {
		if ge := s.getGeByAppid(g.GateAppId); ge != nil {
			ge.AddPlayerMap(uuid, g)
		}
	}()
}

func (ge *gateServer) AddPlayerMap(uuid int64, g *player.GamePlayer) {
	ge.playerMapLock.Lock()
	ge.playerMap[uuid] = g
	ge.playerMapLock.Unlock()
}

func (s *GameServer) DelPlayerMap(uuid int64) {
	syncGD.Lock()
	if s.PlayerMap[uuid] != nil {
		func() {
			if ge := s.getGeByAppid(s.PlayerMap[uuid].GateAppId); ge != nil {
				ge.DelPlayerMap(uuid)
			}
		}()
		delete(s.PlayerMap, uuid)
	}
	syncGD.Unlock()
}

func (ge *gateServer) DelPlayerMap(uuid int64) {
	ge.playerMapLock.Lock()
	if ge.playerMap[uuid] != nil {
		delete(ge.playerMap, uuid)
	}
	ge.playerMapLock.Unlock()
}

func (s *GameServer) GetAllPlayer() map[int64]*player.GamePlayer {
	players := make(map[int64]*player.GamePlayer)
	syncGD.Lock()
	defer syncGD.Unlock()
	for _, play := range s.PlayerMap {
		players[play.Uuid] = play
	}
	return players
}

func (ge *gateServer) GetAllPlayer() map[int64]*player.GamePlayer {
	players := make(map[int64]*player.GamePlayer)
	syncGD.Lock()
	defer syncGD.Unlock()
	for _, play := range ge.playerMap {
		players[play.Uuid] = play
	}
	return players
}

func (s *GameServer) GetPlayerByUuid(uuid int64) *player.GamePlayer {
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
	err = s.Store.SetPlayerStatus(strconv.Itoa(int(p.AccountId)), value)
	return err
}
