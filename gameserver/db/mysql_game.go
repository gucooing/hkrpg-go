package db

// 使用账号id拉取数据
func (s *Store) QueryAccountUidByFieldPlayer(uid uint32) *PlayerData {
	var playerData PlayerData
	s.Mysql.Model(&PlayerData{}).Where("uid = ?", uid).First(&playerData)
	return &playerData
}

// 添加新账号数据
func (s *Store) AddDatePlayerFieldByFieldName(player *PlayerData) error {
	if err := s.Mysql.Create(player).Error; err != nil {
		return err
	}
	return nil
}

// 更新账号
func (s *Store) UpdatePlayer(player *PlayerData) error {
	if player.Uid == 0 {
		return nil
	}
	if err := s.Mysql.Model(&PlayerData{}).Where("uid = ?", player.Uid).Updates(player).Error; err == nil {
		return nil
	} else {
		return err
	}
}
