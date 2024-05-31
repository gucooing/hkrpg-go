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

type BlockData struct { // 地图db
	Uid         uint32 `gorm:"primaryKey"`
	EntryId     uint32 `gorm:"primaryKey"`
	DataVersion uint32
	BinData     []byte
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
