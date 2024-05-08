package database

type PlayerData struct {
	Uid         uint32 `gorm:"primarykey"`
	Nickname    string
	Level       uint32
	Exp         uint32
	DataVersion uint32
	BinData     []byte
}
