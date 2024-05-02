package db

import (
	"github.com/gucooing/hkrpg-go/nodeserver/config"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	StatusRedis *redis.Client // 在线状态数据库
}

func NewStore(conf *config.Config) *Store {
	s := new(Store)
	redisStatusConf := conf.RedisConf["player_status"]
	s.StatusRedis = database.NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)
	logger.Info("数据库连接成功")
	return s
}
