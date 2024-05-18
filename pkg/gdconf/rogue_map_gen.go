package gdconf

import (
	"math/rand"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

func (g *GameDataConfig) loadRogueMapGen() {
	g.RogueMapGenMap = make(map[string][]uint32)
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
}

func GetRogueRoomIDBySiteID(siteID uint32) uint32 {
	rogue := CONF.RogueMapGenMap[strconv.Itoa(int(siteID))]
	idIndex := rand.Intn(len(rogue))
	rogueId := rogue[idIndex]

	return rogueId
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
