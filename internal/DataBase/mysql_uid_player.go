package DataBase

// 使用账号uid拉取数据
func (s *Store) QueryUidPlayerUidByFieldPlayer(AccountId uint32) *UidPlayer {
	var uidplayer UidPlayer
	s.Db.Model(&UidPlayer{}).Where("account_id = ?", AccountId).First(&uidplayer)
	return &uidplayer
}

// 更新账号
func (s *Store) UpdateUidPlayer(accountId uint, uidPlayer *UidPlayer) error {
	if err := s.Db.Model(&UidPlayer{}).Where("account_id = ?", accountId).Updates(uidPlayer).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 添加新账号
func (s *Store) AddUidPlayer(uidPlayer *UidPlayer) error {
	// 检查是否已存在具有相同PlayerID的记录
	var count int64
	s.Db.Model(&UidPlayer{}).Where("account_id = ?", uidPlayer.AccountId).Count(&count)
	if count > 0 {
		if err := s.Db.Model(&UidPlayer{}).Where("account_id = ?", uidPlayer.AccountId).Updates(uidPlayer).Error; err == nil {
			return nil
		} else {
			return err
		}
	}
	// 添加新记录
	if err := s.Db.Create(uidPlayer).Error; err != nil {
		return err
	}

	return nil
}
