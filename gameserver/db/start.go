package db

import (
	"context"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var DBASE *Store
var ctx = context.Background()

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	DBASE = s
	mysqlPlayerDataConf := config.MysqlConf["player_data"]
	s.PlayerDataMysql = alg.NewMysql(mysqlPlayerDataConf.Dsn)
	s.PlayerDataMysql.AutoMigrate(&PlayerData{})

	redisLoginConf := config.RedisConf["player_login"]
	s.LoginRedis = alg.NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)
	redisStatusConf := config.RedisConf["player_status"]
	s.StatusRedis = alg.NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)

	logger.Info("数据库连接成功")
	return s
}
