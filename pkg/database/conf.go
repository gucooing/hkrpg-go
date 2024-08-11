package database

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"gorm.io/gorm"
)

// 拉取全部邮件
func GetDbAllMail(db *gorm.DB) []*constant.Mail {
	var mailMap []*constant.Mail
	db.Find(&mailMap)
	return mailMap
}

// 拉取全部模拟宇宙
func GetAllRogue(db *gorm.DB) []*constant.RogueConf {
	var rogueMap []*constant.RogueConf
	db.Find(&rogueMap)
	return rogueMap
}
