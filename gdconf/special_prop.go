package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type SpecialProp struct {
	EntryId   uint32                `json:"EntryId"`
	GroupList map[uint32]*GroupList `json:"GroupList"`
}
type GroupList struct {
	GroupId   uint32            `json:"GroupId"`
	PropState map[uint32]string `json:"PropState"`
}

func (g *GameDataConfig) loadSpecialProp() {
	g.SpecialPropMap = make(map[uint32]*SpecialProp)
	playerElementsFilePath := g.dataPrefix + "SpecialProp.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.SpecialPropMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v SpecialProp", len(g.SpecialPropMap))
}

func GetSpecialProp(entryId uint32) *SpecialProp {
	return CONF.SpecialPropMap[entryId]
}
