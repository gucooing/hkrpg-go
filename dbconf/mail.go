package dbconf

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
	"gorm.io/gorm"
)

var MAIL *Mail

type Mail struct {
	AllMailMap map[uint32]*constant.Mail
}

func GetMail() *Mail {
	if MAIL == nil {
		MAIL = &Mail{
			AllMailMap: make(map[uint32]*constant.Mail),
		}
	}
	return MAIL
}

// 拉取全服邮件
func NewAllMail(dsn *gorm.DB) {
	m := GetMail()
	db := database.GetDbAllMail(dsn)
	m.AllMailMap = make(map[uint32]*constant.Mail)
	for _, v := range db {
		mail := v
		mail.ItemList = make([]*constant.Item, 0)
		if v.Item != "" {
			err := hjson.Unmarshal([]byte(v.Item), &mail.ItemList)
			if err != nil {
				logger.Error("mail item error: %v", err)
			}
		}
		m.AllMailMap[v.Id] = mail
	}
}

func GetAllMail() map[uint32]*constant.Mail {
	return GetMail().AllMailMap
}

func GetAllMailById(id uint32) *constant.Mail {
	if db := GetMail().AllMailMap; db == nil {
		return nil
	} else {
		return db[id]
	}
}
