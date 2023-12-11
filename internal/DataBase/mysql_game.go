package DataBase

// 使用账号uid拉取数据
func (s *Store) QueryAccountUidByFieldPlayer(AccountId uint32) *Player {
	var player Player
	s.Db.Model(&Player{}).Where("account_uid = ?", AccountId).First(&player)
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
	if player.AccountUid == 0 {
		return nil
	}
	if err := s.Db.Model(&Player{}).Where("account_uid = ?", player.AccountUid).Updates(player).Error; err == nil {
		return nil
	} else {
		return err
	}
}
