package Game

type DbItem struct {
	RelicMap map[uint32]*Item
}

type Item struct {
	Tid uint32
	Num uint32 // 个数
}

func NewItem(data *PlayerData) *PlayerData {
	dbItem := new(DbItem)
	dbItem.RelicMap = make(map[uint32]*Item)

	dbItem.RelicMap[101] = &Item{Tid: 101, Num: 2000}
	dbItem.RelicMap[102] = &Item{Tid: 102, Num: 2000}

	data.DbItem = dbItem

	return data
}
