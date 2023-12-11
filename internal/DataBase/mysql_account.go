package DataBase

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
