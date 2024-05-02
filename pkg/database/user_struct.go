package database

type PlayerUid struct {
	Uid          uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	AccountType  uint32
	AccountId    uint32
	CreateTime   int64
	IsBan        bool
	BanBeginTime int64
	BanEndTime   int64
	BanMsg       string
}
