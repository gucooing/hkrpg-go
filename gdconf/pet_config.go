package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type Pet struct {
	PetMapById     map[uint32]*PetConfig
	PetMapByItemID map[uint32]*PetConfig
}

type PetConfig struct {
	PetID        uint32 `json:"PetID"`
	SummonUnitID uint32 `json:"SummonUnitID"`
	PetItemID    uint32 `json:"PetItemID"`
}

func (g *GameDataConfig) loadPetConfig() {
	g.Pet = &Pet{
		PetMapById:     make(map[uint32]*PetConfig),
		PetMapByItemID: make(map[uint32]*PetConfig),
	}
	petConfigList := make([]*PetConfig, 0)
	name := "PetConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &petConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range petConfigList {
		if g.Pet.PetMapById == nil {
			g.Pet.PetMapById = make(map[uint32]*PetConfig)
		}
		if g.Pet.PetMapByItemID == nil {
			g.Pet.PetMapByItemID = make(map[uint32]*PetConfig)
		}
		g.Pet.PetMapById[v.PetID] = v
		g.Pet.PetMapByItemID[v.PetItemID] = v
	}

	logger.Info(text.GetText(17), len(g.Pet.PetMapById), name)
}

func GetPetConfigByItemId(itemId uint32) *PetConfig {
	return getConf().Pet.PetMapByItemID[itemId]
}

func GetPetConfigById(Id uint32) *PetConfig {
	return getConf().Pet.PetMapById[Id]
}
