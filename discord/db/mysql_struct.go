package db

import (
	"github.com/gucooing/hkrpg-go/discord/config"
	"gorm.io/gorm"
)

type Store struct {
	config *config.Config
	Db     *gorm.DB
}

type Account struct {
	AccountId  uint `gorm:"primarykey;AUTO_INCREMENT"`
	Username   string
	Token      string
	CreateTime int64
}

type UidPlayer struct {
	AccountUid uint `gorm:"primarykey;AUTO_INCREMENT"`
	AccountId  uint
	IsBan      bool
	ComboToken string
}
