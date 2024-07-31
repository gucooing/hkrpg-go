package db

import (
	"context"

	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var Db *Store

type Store struct {
	config               *config.Config
	PlayerDataMysql      *gorm.DB
	ServerConf           *gorm.DB
	LoginRedis           *redis.Client
	StatusRedis          *redis.Client
	PlayerBriefDataRedis *redis.Client // 玩家简要信息
}

var ctx = context.Background()

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	Db = s
	mysqlPlayerDataConf := config.MysqlConf["player"]
	s.PlayerDataMysql = database.NewMysql(mysqlPlayerDataConf.Dsn)
	s.PlayerDataMysql.AutoMigrate(&database.PlayerData{}, &database.BlockData{})
	mysqlServerConf := config.MysqlConf["conf"]
	s.ServerConf = database.NewMysql(mysqlServerConf.Dsn)
	s.ServerConf.AutoMigrate(&database.Mail{}, &database.RogueConf{}, &database.ScheduleConf{})

	redisLoginConf := config.RedisConf["player_login"]
	s.LoginRedis = database.NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)
	redisStatusConf := config.RedisConf["player_status"]
	s.StatusRedis = database.NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)
	playerBriefDataRedis := config.RedisConf["player_brief_data"]
	s.PlayerBriefDataRedis = database.NewRedis(playerBriefDataRedis.Addr, playerBriefDataRedis.Password, playerBriefDataRedis.DB)

	logger.Info("数据库连接成功")
	database.GetDbConf(s.ServerConf) // 初始化数据库配置表
	return s
}

func GetDb() *Store {
	return Db
}
