package DataBase

import (
	"github.com/gucooing/hkrpg-go/pkg/config"
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

type Player struct {
	AccountUid uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	PlayerData []byte
}
