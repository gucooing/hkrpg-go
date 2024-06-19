package database

import (
	"context"
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
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

// 获取玩家好友申请信息
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

// 设置玩家好友申请信息
func SetPlayerFriend(rc *redis.Client, uid uint32, value []byte) bool {
	key := "player_friend:" + strconv.Itoa(int(uid))
	err := rc.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}
	return true
}

// 获取玩家待加入数据库好友信息
func GetAcceptApplyFriend(rc *redis.Client, uid uint32) ([]byte, bool) {
	key := "accept_apply_friend:" + strconv.Itoa(int(uid))
	bin, err := rc.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 设置玩家待加入数据库好友信息
func SetAcceptApplyFriend(rc *redis.Client, uid uint32, value []byte) bool {
	key := "accept_apply_friend:" + strconv.Itoa(int(uid))
	err := rc.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}
	return true
}

// 删除玩家待加入数据库好友信息
func DelAcceptApplyFriend(rc *redis.Client, uid uint32) {
	key := "accept_apply_friend:" + strconv.Itoa(int(uid))
	rc.Del(ctx, key)
}

// 获取ComboToken
func GetComboTokenByAccountId(rc *redis.Client, accountId string) string {
	key := "player_comboToken:" + accountId
	comboToken, err := rc.Get(ctx, key).Result()
	if err != nil {
		comboToken = random.GetRandomByteHexStr(20)
		SetComboTokenByAccountId(rc, accountId, comboToken)
	}
	return comboToken
}

// 设置ComboToken
func SetComboTokenByAccountId(rc *redis.Client, accountId, comboToken string) string {
	key := "player_comboToken:" + accountId
	err := rc.Set(ctx, key, comboToken, 168*time.Hour).Err()
	if err != nil {
		return ""
	}
	logger.Debug("[accountId: %s] 生产新的comboToken: %s", accountId, comboToken)
	return comboToken
}
