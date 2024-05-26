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

// 添加新账号数据
func (s *Store) AddDatePlayerFieldByFieldName(player *database.PlayerData) error {
	if err := s.PlayerDataMysql.Create(player).Error; err != nil {
		return err
	}
	return nil
}

// 更新账号
func (s *Store) UpdatePlayer(player *database.PlayerData) error {
	if player.Uid == 0 {
		return nil
	}
	if err := s.PlayerDataMysql.Model(&database.PlayerData{}).Where("uid = ?", player.Uid).Updates(player).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 拉取全部邮件
func (s *Store) GetAllMail() []*database.Mail {
	var mailMap []*database.Mail
	s.ServerConf.Find(&mailMap)
	return mailMap
}
