package database

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedis(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:               "",
		Addr:                  addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              password,
		CredentialsProvider:   nil,
		DB:                    db,
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
		panic("redis connect fail" + err.Error())
	}
	return rdb
}

// 获取玩家好友信息
func GetPlayerFriend(rc *redis.Client, uid uint32) ([]byte, bool) {
	key := "player_friend:" + strconv.Itoa(int(uid))
	bin, err := rc.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 设置玩家好友信息
func SetPlayerFriend(rc *redis.Client, uid uint32, value []byte) bool {
	key := "player_friend:" + strconv.Itoa(int(uid))
	err := rc.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}
	return true
}
