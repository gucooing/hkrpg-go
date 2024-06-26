package gdconf

import (
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

func (g *GameDataConfig) loadRogueMapGen() {
	g.RogueMapGenMap = make(map[uint32][]uint32)
	playerElementsFilePath := g.dataPrefix + "RogueMapGen.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		logger.Error("open file error: %v", err)
		return
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueMapGenMap)
	if err != nil {
		logger.Error("parse file error: %v", err)
		return
	}
	logger.Info("load %v RogueMapGen", len(g.RogueMapGenMap))
	g.wg.Done()
}

func GetRogueRoomTypeBySiteID(siteID uint32) uint32 {
	rogue := CONF.RogueMapGenMap[siteID]
	idIndex := rand.Intn(len(rogue))
	typeId := rogue[idIndex]

	return GetRogueRoomByType(typeId)
}

func GetRogueRoomTypeBy100(siteID uint32) uint32 {
	roomMap := map[uint32]uint32{
		1: 2,
		2: 6,
		3: 5,
		4: 2,
		5: 3,
		6: 5,
		7: 7,
	}
	return GetRogueRoomByType(roomMap[siteID])
}
