package db

import (
	"github.com/gucooing/hkrpg-go/gameserver/config"
	"gorm.io/gorm"
)

type Store struct {
	config *config.Config
	Db     *gorm.DB
}

type Player struct {
	Uid          uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	AccountId    uint32
	PlayerDataPb []byte
}
