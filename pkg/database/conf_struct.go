package database

import (
	"database/sql"
)

// 全服邮件
type Mail struct {
	Id        uint64       `gorm:"primarykey;AUTO_INCREMENT"` // 邮件id
	Title     string       // 邮件标题
	Sender    string       // 发件人
	BeginTime sql.NullTime // 开始时间
	EndTime   sql.NullTime // 结束时间
	Content   string       // 内容
}

type ServerConf struct {
	AutoCreate bool // 是否自动注册
}
