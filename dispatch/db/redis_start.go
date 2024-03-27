package db

import (
	"context"

	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
)

func NewRedis(config *config.Config) *redis.Client {
	reconf := config.RedisConf["player_token"]
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
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err.Error())
	}
	logger.Info("Redis数据库连接成功")
	return rdb
}
