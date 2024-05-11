package database

// 全服邮件
type Mail struct {
	Id        uint64 `gorm:"primarykey;AUTO_INCREMENT"` // 邮件id
	Title     string // 邮件标题
	Sender    string // 发件人
	BeginTime int64  // 开始时间
	EndTime   int64  // 结束时间
	Content   string // 内容
}
