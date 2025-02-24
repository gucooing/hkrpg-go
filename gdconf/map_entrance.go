package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MapEntranceMap struct {
	MapEntranceList         []*MapEntrance
	MapEntranceMapByEntryId map[uint32]*MapEntrance
	MapEntranceMapByFloorID map[uint32]*MapEntrance
}

type MapEntrance struct {
	ID                    uint32   `json:"ID"`
	EntranceType          string   `json:"EntranceType"`
	PlaneID               uint32   `json:"PlaneID"`
	FloorID               uint32   `json:"FloorID"`
	StartGroupID          uint32   `json:"StartGroupID"`
	StartAnchorID         uint32   `json:"StartAnchorID"`
	BeginMainMissionList  []uint32 `json:"BeginMainMissionList"`
	FinishMainMissionList []uint32 `json:"FinishMainMissionList"`
	FinishSubMissionList  []uint32 `json:"FinishSubMissionList"`
}

func (g *GameDataConfig) loadMapEntrance() {
	g.MapEntranceMap = &MapEntranceMap{
		MapEntranceList:         make([]*MapEntrance, 0),
		MapEntranceMapByEntryId: make(map[uint32]*MapEntrance),
		MapEntranceMapByFloorID: make(map[uint32]*MapEntrance),
	}
	name := "MapEntrance.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MapEntranceMap.MapEntranceList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range g.MapEntranceMap.MapEntranceList {
		g.MapEntranceMap.MapEntranceMapByEntryId[v.ID] = v
		g.MapEntranceMap.MapEntranceMapByFloorID[v.FloorID] = v
	}

	logger.Info(text.GetText(17), len(g.MapEntranceMap.MapEntranceList), name)
}

func GetMapEntranceById(entryId uint32) *MapEntrance {
	return getConf().MapEntranceMap.MapEntranceMapByEntryId[entryId]
}

func GetMapEntranceByFloorID(floorID uint32) *MapEntrance {
	return getConf().MapEntranceMap.MapEntranceMapByFloorID[floorID]
}

func GetPFlaneID(entryId uint32) (uint32, uint32, bool) {
	m := GetMapEntranceById(entryId)
	if m == nil {
		return 0, 0, false
	}
	return m.PlaneID, m.FloorID, true
}

func GetMapEntranceMap() map[uint32]*MapEntrance {
	return getConf().MapEntranceMap.MapEntranceMapByEntryId
}

func GetEntryIdList() []uint32 {
	var entryIdList []uint32
	for _, id := range getConf().MapEntranceMap.MapEntranceList {
		entryIdList = append(entryIdList, id.ID)
	}
	return entryIdList
}
