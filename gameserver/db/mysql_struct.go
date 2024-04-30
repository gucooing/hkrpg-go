package db

import (
	"github.com/gucooing/hkrpg-go/gameserver/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Store struct {
	config               *config.Config
	PlayerDataMysql      *gorm.DB
	LoginRedis           *redis.Client
	StatusRedis          *redis.Client
	PlayerBriefDataRedis *redis.Client // 玩家简要信息
}

type PlayerData struct {
	Uid      uint32 `gorm:"primarykey"`
	Nickname string
	Level    uint32
	Exp      uint32
	BinData  []byte
}
