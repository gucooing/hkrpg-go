package dbconf

import (
	"gorm.io/gorm"
)

func GameServer(dsn *gorm.DB) {
	NewAllMail(dsn) // 全服邮件
	NewRogue(dsn)   // 模拟宇宙排期
}
