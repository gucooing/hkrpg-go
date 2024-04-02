package db

import (
	"github.com/redis/go-redis/v9"
)

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
