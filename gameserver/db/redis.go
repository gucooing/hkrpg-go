package db

import (
	"context"
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
)

const (
	MaxLockAliveTime  = 0  // 单个锁的最大存活时间 毫秒
	LockRetryWaitTime = 50 // 同步加锁重试间隔时间 毫秒
	MaxLockRetryTimes = 2  // 同步加锁最大重试次数
)

// 标记玩家状态
func (s *Store) DistLockPlayerStatus(accountId string, value []byte) bool {
	var result = false
	for i := 0; i < MaxLockRetryTimes; i++ {
		var err error = nil
		key := "player_status_lock:" + accountId
		result, err = s.StatusRedis.SetNX(context.TODO(),
			key,
			value,
			time.Millisecond*time.Duration(MaxLockAliveTime)).Result()
		if err != nil {
			logger.Error("redis lock setnx error: %v", err)
			return false
		}
		if result == true {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(LockRetryWaitTime))
	}
	return result
}

// 获取玩家状态
func (s *Store) GetPlayerStatus(accountId string) ([]byte, bool) {
	key := "player_status_lock:" + accountId
	bin, err := s.StatusRedis.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 删除玩家状态
func (s *Store) DistUnlockPlayerStatus(accountId string) {
	var result int64 = 0
	var err error = nil
	key := "player_status_lock:" + accountId
	result, err = s.StatusRedis.Del(context.TODO(), key).Result()
	if err != nil {
		logger.Error("redis lock del error: %v", err)
		return
	}
	if result == 0 {
		logger.Error("redis lock del result is fail")
		return
	}
}

func (s *Store) SetPlayerPlayerBasicBriefData(uid uint32, value []byte) bool {
	key := "player_brief_data:" + strconv.Itoa(int(uid))
	err := s.PlayerBriefDataRedis.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}
	return true
}

func (s *Store) GetPlayerPlayerBasicBriefData(uid uint32) ([]byte, bool) {
	key := "player_brief_data:" + strconv.Itoa(int(uid))
	bin, err := s.PlayerBriefDataRedis.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 获取玩家好友信息
func (s *Store) GetPlayerFriend(uid uint32) ([]byte, bool) {
	key := "player_friend:" + strconv.Itoa(int(uid))
	bin, err := s.PlayerBriefDataRedis.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 设置玩家好友信息
func (s *Store) SetPlayerFriend(uid uint32, value []byte) bool {
	key := "player_friend:" + strconv.Itoa(int(uid))
	err := s.PlayerBriefDataRedis.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false
	}
	return true
}
