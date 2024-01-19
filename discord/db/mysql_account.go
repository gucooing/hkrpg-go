package db

// 查询账号
func (s *Store) QueryAccountByFieldUsername(Username string) *Account {
	var account Account
	s.Db.Model(&Account{}).Where("Username = ?", Username).First(&account)
	return &account
}
func (s *Store) QueryAccountByFieldAccountId(AccountId uint) *Account {
	var account Account
	s.Db.Model(&Account{}).Where("account_id = ?", AccountId).First(&account)
	return &account
}

// 添加新账号
func (s *Store) UpdateAccountFieldByFieldName(account *Account) (uint, error) {
	if err := s.Db.Create(account).Error; err == nil {
		return account.AccountId, nil
	} else {
		return 0, err
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
