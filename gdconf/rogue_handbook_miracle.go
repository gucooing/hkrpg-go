package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueHandbookMiracle struct {
	MiracleHandbookID         uint32   `json:"MiracleHandbookID"`
	MiracleReward             uint32   `json:"MiracleReward"`
	MiracleTypeList           []uint32 `json:"MiracleTypeList"`
	MiracleDisplayID          uint32   `json:"MiracleDisplayID"`
	Order                     uint32   `json:"Order"`
	MiracleIDForEffectDisplay uint32   `json:"MiracleIDForEffectDisplay"`
}

func (g *GameDataConfig) loadRogueHandbookMiracle() {
	g.RogueHandbookMiracleMap = make(map[uint32]*RogueHandbookMiracle)
	rogueHandbookMiracleList := make([]*RogueHandbookMiracle, 0)
	name := "RogueHandbookMiracle.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueHandbookMiracleList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueHandbookMiracleList {
		g.RogueHandbookMiracleMap[v.MiracleHandbookID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueHandbookMiracleMap), name)
}

func GetRogueHandbookMiracleMap() map[uint32]*RogueHandbookMiracle {
	return getConf().RogueHandbookMiracleMap
}

func GetRogueHandbookMiracle(miracleID uint32) *RogueHandbookMiracle {
	return getConf().RogueHandbookMiracleMap[miracleID]
}
