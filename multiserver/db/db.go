package db

import (
	"github.com/gucooing/hkrpg-go/multiserver/config"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Store struct {
	ConfMysql   *gorm.DB      // 服务器配置
	MailRedis   *redis.Client // 玩家邮件库
	ChatRedis   *redis.Client // 聊天临时数据
	StatusRedis *redis.Client // 在线状态数据库
}

func NewStore(conf *config.Config) *Store {
	s := new(Store)
	confMysql := conf.MysqlConf["conf"]
	s.ConfMysql = database.NewMysql(confMysql.Dsn)
	s.ConfMysql.AutoMigrate(&database.Mail{})

	redisMailConf := conf.RedisConf["player_mail"]
	s.MailRedis = database.NewRedis(redisMailConf.Addr, redisMailConf.Password, redisMailConf.DB)
	redisChatConf := conf.RedisConf["player_chat"]
	s.ChatRedis = database.NewRedis(redisChatConf.Addr, redisChatConf.Password, redisChatConf.DB)
	redisStatusConf := conf.RedisConf["player_status"]
	s.StatusRedis = database.NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)
	logger.Info("数据库连接成功")
	return s
}
