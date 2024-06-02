package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "MappingInfo.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MappingInfoMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v MappingInfo", len(g.MappingInfoMap))
}

func GetMappingInfoById(stageID, worldLevel uint32) *MappingInfo {
	return CONF.MappingInfoMap[stageID][worldLevel]
}

func GetEntryId(planeID, floorID uint32) uint32 {
	var entryId uint32 = 0
tv:
	for _, mappingInfo := range CONF.MappingInfoMap {
		for _, mapEntrance := range mappingInfo {
			if mapEntrance.PlaneID == planeID && mapEntrance.FloorID == floorID {
				entryId = mapEntrance.ID
				break tv
			}
		}
	}
	return entryId
}
