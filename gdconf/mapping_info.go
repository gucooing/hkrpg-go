package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MappingInfo struct {
	ID              uint32        `json:"ID"`
	WorldLevel      uint32        `json:"WorldLevel"`
	Type            string        `json:"Type"`
	FarmType        string        `json:"FarmType"`
	IsTeleport      bool          `json:"IsTeleport"`
	IsShowInFog     bool          `json:"IsShowInFog"`
	PlaneID         uint32        `json:"PlaneID"`
	FloorID         uint32        `json:"FloorID"`
	GroupID         uint32        `json:"GroupID"`
	ConfigID        uint32        `json:"ConfigID"`
	ShowMonsterList []uint32      `json:"ShowMonsterList"`
	DisplayItemList []*RewardList `json:"DisplayItemList"`
}

func (g *GameDataConfig) loadMappingInfo() {
	g.MappingInfoMap = make(map[uint32]map[uint32]*MappingInfo)
	mappingInfoMap := make([]*MappingInfo, 0)
	name := "MappingInfo.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &mappingInfoMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range mappingInfoMap {
		if g.MappingInfoMap[v.ID] == nil {
			g.MappingInfoMap[v.ID] = make(map[uint32]*MappingInfo)
		}
		g.MappingInfoMap[v.ID][v.WorldLevel] = v
	}

	logger.Info(text.GetText(17), len(g.MappingInfoMap), name)
}

func GetMappingInfoById(id, worldLevel uint32) *MappingInfo {
	return getConf().MappingInfoMap[id][worldLevel]
}
