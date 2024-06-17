package database

import (
	"database/sql"
)

// 邮件奖励类型
const (
	MailAvatar   = 1 // 角色
	MailMaterial = 2 // 材料
)

// 全服邮件
type Mail struct {
	Id        uint32       `gorm:"primarykey;AUTO_INCREMENT"` // 邮件id
	Title     string       // 邮件标题
	Sender    string       // 发件人
	BeginTime sql.NullTime // 开始时间
	EndTime   sql.NullTime // 结束时间
	Content   string       // 内容
	Item      string       // 邮件附件
	ItemList  []*Item      `gorm:"-"`
}

type Item struct { // 邮件奖励模板
	ItemType uint32 // 类型
	ItemId   uint32 // id
	Num      uint32 // 数量
}

type ServerConf struct {
	AutoCreate bool // 是否自动注册
}

type RogueConf struct {
	SeasonId  uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	BeginTime sql.NullTime
	EndTime   sql.NullTime
}

type ScheduleConf struct {
	ScheduleId uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	BeginTime  sql.NullTime
	EndTime    sql.NullTime
}
