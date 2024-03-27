package db

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

func (s *Store) GetComboTokenByAccountId(accountId string) string {
	key := "player_comboToken:" + accountId
	comboToken, err := s.RedisDb.Get(ctx, key).Result()
	if err != nil {
		comboToken = random.GetRandomByteHexStr(20)
		s.SetComboTokenByAccountId(accountId, comboToken)
	}
	return comboToken
}

func (s *Store) SetComboTokenByAccountId(accountId, comboToken string) string {
	key := "player_comboToken:" + accountId
	err := s.RedisDb.Set(ctx, key, comboToken, 168*time.Hour).Err()
	if err != nil {
		return ""
	}
	logger.Debug("[accountId: %s] 生产新的comboToken: %s", accountId, comboToken)
	return comboToken
}
