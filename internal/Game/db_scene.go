package Game

type DbScene struct {
	WorldId uint32 // 世界ID
	EntryId uint32 //
}

func NewScene(data *PlayerData) *PlayerData {
	dbScene := new(DbScene)
	dbScene.WorldId = 201
	dbScene.EntryId = 1010101

	data.DbScene = dbScene
	return data
}
