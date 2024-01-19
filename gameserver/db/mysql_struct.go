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
	AccountUid   uint32
	PlayerDataPb []byte
}
