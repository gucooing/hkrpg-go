package dispatch

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

type Store struct {
	config       *config.Config
	AccountMysql *gorm.DB
	LoginRedis   *redis.Client
	HkrpgGoPe    *gorm.DB
}

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	accountMysqlConf := config.MysqlConf["account"]
	s.AccountMysql = database.NewMysql(accountMysqlConf.Dsn)
	s.AccountMysql.AutoMigrate(&database.Account{})

	redisLoginConf := config.RedisConf["player_login"]
	s.LoginRedis = database.NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)

	logger.Info("数据库连接成功")
	return s
}
