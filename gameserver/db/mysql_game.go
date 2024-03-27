package db

// 使用账号id拉取数据
func (s *Store) QueryAccountUidByFieldPlayer(AccountId uint32) *Player {
	var player Player
	s.Db.Model(&Player{}).Where("account_id = ?", AccountId).First(&player)
	return &player
}

// 添加新账号数据
func (s *Store) AddDatePlayerFieldByFieldName(player *Player) error {
	if err := s.Db.Create(player).Error; err != nil {
		return err
	}
	return nil
}

// 更新账号
func (s *Store) UpdatePlayer(player *Player) error {
	if player.Uid == 0 {
		return nil
	}
	if err := s.Db.Model(&Player{}).Where("account_id = ?", player.AccountId).Updates(player).Error; err == nil {
		return nil
	} else {
		return err
	}
}
