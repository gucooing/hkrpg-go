package gdconf

import (
	"fmt"
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
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueMapGenMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RogueMapGen", len(g.RogueMapGenMap))
}

func GetRogueRoomIDBySiteID() map[uint32]uint32 {
	var rogueMap map[uint32]uint32
	rogueMap = make(map[uint32]uint32)
	for id, rogue := range CONF.RogueMapGenMap {
		idIndex := rand.Intn(len(rogue))
		rogueMap[stou32(id)] = rogue[idIndex]
	}

	return rogueMap
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
