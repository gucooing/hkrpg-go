package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type FarmElementConfig struct {
	ID            uint32   `json:"ID"`
	WorldLevel    uint32   `json:"WorldLevel"`
	MappingInfoID uint32   `json:"MappingInfoID"`
	DropList      []uint32 `json:"DropList"`
	StaminaCost   uint32   `json:"StaminaCost"`
	DamageType    []string `json:"DamageType"`
	StageID       uint32   `json:"StageID"`
}

func (g *GameDataConfig) loadFarmElementConfig() {
	g.FarmElementConfigMap = make(map[uint32]*FarmElementConfig)
	farmElementConfiglist := make([]*FarmElementConfig, 0)
	playerElementsFilePath := g.excelPrefix + "FarmElementConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &farmElementConfiglist)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range farmElementConfiglist {
		g.FarmElementConfigMap[v.StageID] = v
	}
	logger.Info("load %v FarmElementConfig", len(g.FarmElementConfigMap))
}

func GetFarmElementConfig(id uint32) *FarmElementConfig {
	return CONF.FarmElementConfigMap[id]
}
