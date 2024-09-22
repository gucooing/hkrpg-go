package database

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"gorm.io/gorm"
)

func NewPush(db *gorm.DB) {
	db.AutoMigrate(
		&constant.ServiceLog{}, // log
	)
	logger.Info("数据库连接成功")
}
