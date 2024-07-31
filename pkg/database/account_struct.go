package database

import (
	"time"
)

type Account struct {
	AccountId  uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	Username   string
	Token      string
	ComboToken string
	CreateTime int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
