package db

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
)

// 使用账号id拉取数据
func (s *Store) QueryAccountUidByFieldPlayer(uid uint32) *constant.PlayerData {
	var playerData constant.PlayerData
	s.PlayerDataMysql.Model(&constant.PlayerData{}).Where("uid = ?", uid).First(&playerData)
	return &playerData
}
