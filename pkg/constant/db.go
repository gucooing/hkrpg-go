package constant

import (
	"database/sql"
)

type Account struct {
	AccountId  uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	Username   string
	Token      string
	CreateTime int64
}

/**********************conf*************************/

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

// 玩家邮件
type PlayerMail struct {
	Uid       uint32       `gorm:"primarykey;AUTO_INCREMENT"` // uid
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

type RegionConf struct { // 区服配置
	Id              uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	Name            string // 区服名称
	AutoCreate      bool   // 是否自动注册
	Title           string // 地区
	Type            uint32 // sdk类型
	ClientSecretKey []byte // 加密密钥
}

/**********************player data*************************/

type PlayerData struct {
	Uid         uint32 `gorm:"primarykey"`
	Nickname    string
	Level       uint32
	Exp         uint32
	DataVersion uint32
	BinData     []byte
}

type BlockData struct { // 地图db
	Uid         uint32 `gorm:"primaryKey"`
	EntryId     uint32 `gorm:"primaryKey"`
	DataVersion uint32
	BinData     []byte
}

type PlayerBasic struct {
	Uid     uint32 `gorm:"primaryKey"`
	BinData []byte
}

type PlayerUid struct {
	Uid          uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	AccountType  uint32
	AccountId    uint32
	ComboToken   string
	CreateTime   int64
	IsBan        bool
	BanBeginTime int64
	BanEndTime   int64
	BanMsg       string
}

type ApplyFriend struct {
	Uid          uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	ReceiveApply []byte
}

type AcceptApplyFriend struct {
	Uid               uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	AcceptApplyFriend []byte
}
