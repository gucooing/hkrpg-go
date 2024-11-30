package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

func (g *GameDataConfig) loadRogueMapGen() {
	g.RogueMapGenMap = make(map[uint32][]uint32)
	name := "RogueMapGen.json"
	playerElementsFile, err := os.ReadFile(g.dataPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueMapGenMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	logger.Info(text.GetText(17), len(g.RogueMapGenMap), name)
}

func GetRogueRoomTypeBySiteID(siteID uint32) uint32 {
	rogue := getConf().RogueMapGenMap[siteID]
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
