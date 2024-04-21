package gs

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
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
func (ge *gateServer) killPlayer(p *player.GamePlayer) {
	ge.seedGate(cmd.GameToGatePlayerLogoutNotify, &spb.GameToGatePlayerLogoutNotify{
		Uid:  p.Uid,
		Uuid: p.Uuid,
	})
	ge.upDataPlayer(p)
	db.DBASE.DelPlayerStatus(strconv.Itoa(int(p.AccountId)))
	ge.game.DelPlayerMap(p.Uuid)
	logger.Debug("[UID:%v][UUID:%v]玩家重复登录下线成功", p.Uid, p.Uuid)
}

func (ge *gateServer) upDataPlayer(p *player.GamePlayer) {
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
		logger.Debug("[UID:%v][UUID:%v]数据过期，已丢弃", p.Uid, p.Uuid)
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
