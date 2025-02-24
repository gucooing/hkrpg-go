package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type FarmElementConfigMap struct {
	FarmElementConfigById      map[uint32]map[uint32]*FarmElementConfig // WorldLevel ID
	FarmElementConfigByStageID map[uint32]*FarmElementConfig
}

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
	g.FarmElementConfigMap = &FarmElementConfigMap{
		FarmElementConfigById:      make(map[uint32]map[uint32]*FarmElementConfig),
		FarmElementConfigByStageID: make(map[uint32]*FarmElementConfig),
	}
	farmElementConfiglist := make([]*FarmElementConfig, 0)
	name := "FarmElementConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &farmElementConfiglist)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range farmElementConfiglist {
		if g.FarmElementConfigMap.FarmElementConfigById[v.ID] == nil {
			g.FarmElementConfigMap.FarmElementConfigById[v.ID] = make(map[uint32]*FarmElementConfig)
		}
		g.FarmElementConfigMap.FarmElementConfigById[v.ID][v.WorldLevel] = v
		g.FarmElementConfigMap.FarmElementConfigByStageID[v.StageID] = v
	}

	logger.Info(text.GetText(17), len(g.FarmElementConfigMap.FarmElementConfigById), name)
}

func GetFarmElementConfigByStageID(stageID uint32) *FarmElementConfig {
	return getConf().FarmElementConfigMap.FarmElementConfigByStageID[stageID]
}

func GetFarmElementConfigByID(id, worldLevel uint32) *FarmElementConfig {
	if getConf().FarmElementConfigMap.FarmElementConfigById[id] == nil {
		return nil
	}
	return getConf().FarmElementConfigMap.FarmElementConfigById[id][worldLevel]
}
