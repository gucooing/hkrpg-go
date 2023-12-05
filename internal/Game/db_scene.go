package Game

type DbScene struct {
	WorldId uint32 // 世界ID
	EntryId uint32 //
}

func NewScene(data *PlayerData) *PlayerData {
	dbScene := new(DbScene)
	dbScene.WorldId = 101
	dbScene.EntryId = 1000101

	data.DbScene = dbScene
	return data
}
