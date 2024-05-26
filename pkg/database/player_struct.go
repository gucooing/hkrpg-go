package database

import (
	"time"
)

type PlayerData struct {
	Uid         uint32 `gorm:"primarykey"`
	Nickname    string
	Level       uint32
	Exp         uint32
	DataVersion uint32
	BinData     []byte
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
