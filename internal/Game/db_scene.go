package Game

type DbScene struct {
	EntryId uint32 //
}

func NewScene(data *PlayerData) *PlayerData {
	dbScene := new(DbScene)
	// dbScene.EntryId = 1010101
	dbScene.EntryId = 2000201

	data.DbScene = dbScene
	return data
}
