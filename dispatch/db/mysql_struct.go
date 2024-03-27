package db

import (
	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Store struct {
	config  *config.Config
	MysqlDb *gorm.DB
	RedisDb *redis.Client
}

type Account struct {
	AccountId  uint `gorm:"primarykey;AUTO_INCREMENT"`
	Username   string
	Token      string
	CreateTime int64
}
