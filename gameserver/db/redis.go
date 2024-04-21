package db

import (
	"github.com/redis/go-redis/v9"
)

// 标记玩家状态
func (s *Store) SetPlayerStatus(accountId string, value []byte) error {
	key := "player_status:" + accountId
	err := s.RedisDb.Set(ctx, key, value, 0).Err()
	return err
}

// 获取玩家状态
func (s *Store) GetPlayerStatus(accountId string) ([]byte, bool) {
	key := "player_status:" + accountId
	bin, err := s.RedisDb.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 删除玩家状态
func (s *Store) DelPlayerStatus(accountId string) error {
	key := "player_status:" + accountId
	err := s.RedisDb.Del(ctx, key).Err()
	return err
}
