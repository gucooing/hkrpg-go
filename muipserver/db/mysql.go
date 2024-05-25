package db

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
)

// 使用账号id拉取数据
func (s *Store) QueryAccountUidByFieldPlayer(uid uint32) *database.PlayerData {
	var playerData database.PlayerData
	s.PlayerDataMysql.Model(&database.PlayerData{}).Where("uid = ?", uid).First(&playerData)
	return &playerData
}
