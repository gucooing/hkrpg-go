package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

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
