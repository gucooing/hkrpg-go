package database

import (
	"time"
)

type Account struct {
	AccountId  uint `gorm:"primarykey;AUTO_INCREMENT"`
	Username   string
	Token      string
	CreateTime int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
