package db

import (
	"context"
	"time"

	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gromlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var ctx = context.Background()

func (s *Store) init() {
	var err error
	dsn := s.config.MysqlDsn
	s.MysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gromlogger.Default.LogMode(gromlogger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error("MySQL数据库连接失败,错误原因:%s", err)
		return
	}
	logger.Info("MySQL数据库连接成功")
	sqlDB, err := s.MysqlDb.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(5)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10 秒钟
	// 初始化表
	err = s.MysqlDb.AutoMigrate(&Account{})
	if err != nil {
		logger.Error("MySQL数据库初始化失败")
		return
	}
	logger.Info("MySQL数据库初始化成功")
}

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	s.init()
	s.RedisDb = NewRedis(config) // 初始化redis
	return s
}

func NewRedis(config *config.Config) *redis.Client {
	reconf := config.RedisConf["player_login"]
	rdb := redis.NewClient(&redis.Options{
		Network:               "",
		Addr:                  reconf.Addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              reconf.Password,
		CredentialsProvider:   nil,
		DB:                    reconf.DB,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		MaxActiveConns:        0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
		DisableIndentity:      false,
		IdentitySuffix:        "",
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err.Error())
	}
	logger.Info("Redis数据库连接成功")
	return rdb
}
