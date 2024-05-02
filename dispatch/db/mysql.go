package db

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
)

// 查询账号
func (s *Store) QueryAccountByFieldUsername(Username string) *database.Account {
	var account database.Account
	s.AccountMysql.Model(&database.Account{}).Where("Username = ?", Username).First(&account)
	return &account
}
func (s *Store) QueryAccountByFieldAccountId(AccountId uint) *database.Account {
	var account database.Account
	s.AccountMysql.Model(&database.Account{}).Where("account_id = ?", AccountId).First(&account)
	return &account
}

// 添加新账号
func (s *Store) UpdateAccountFieldByFieldName(account *database.Account) (uint, error) {
	if err := s.AccountMysql.Create(account).Error; err == nil {
		return account.AccountId, nil
	} else {
		return 0, err
	}
}
