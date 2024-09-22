package constant

import (
	"database/sql"
	"time"
)

type ServiceLog struct {
	LogLevel     LogLevel     `gorm:"primarykey"`
	LogId        uint64       `gorm:"primarykey"`
	Tag          string       `gorm:"primarykey"` // 自由tag
	LogAddTime   sql.NullTime // 添加时间
	LogValidTime sql.NullTime // 有效时间
	LogMsg       string       // log 内容
}

type LogLevel int

const (
	DEBUG = 0
	INFO  = 1
	WARN  = 2 // 警告!将会触发通知
	ERROR = 3 // 严重错误!将会触发通知
)

type PushMessageAll interface {
	push()
}

type PushMessage struct {
	Tag string // 自定义tag
}

func (l *LogPush) push() {}

type LogPush struct {
	PushMessage
	LogLevel     LogLevel  // 等级
	LogValidTime time.Time // 有效时间
	LogMsg       string    // 内容
}
