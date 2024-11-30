package gdconf

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

type GoppLevelFloor struct {
	FloorID       uint32
	DimensionList map[uint32]*GoppRtLevelDimensionInfo // 位面列表
}

type GoppRtLevelDimensionInfo struct {
	ID              uint32 // 位面ID, 0是基础位面
	Category        constant.LevelDimensionCategory
	SavedValues     []*FloorSavedValueConfig // 关卡存档变量列表
	GroupIndexList  map[uint32]bool          // 组索引列表
	StartAnchorID   uint32                   // 开始组索引
	StartGroupIndex uint32                   // 开始锚点ID
}

func (g *GameDataConfig) goppFloor() {
	g.GoppFloorMap = make(map[uint32]map[uint32]*GoppLevelFloor)
	for planeId, floorMap := range GetFloor() {
		if g.GoppFloorMap[planeId] == nil {
			g.GoppFloorMap[planeId] = make(map[uint32]*GoppLevelFloor)
		}
		for floorId, floor := range floorMap {
			goppFloor := &GoppLevelFloor{
				FloorID:       floorId,
				DimensionList: make(map[uint32]*GoppRtLevelDimensionInfo),
			}
			for _, dim := range floor.DimensionList {
				info := &GoppRtLevelDimensionInfo{
					ID:              dim.ID,
					Category:        dim.Category,
					SavedValues:     dim.SavedValues,
					GroupIndexList:  make(map[uint32]bool),
					StartAnchorID:   dim.StartAnchorID,
					StartGroupIndex: dim.StartGroupIndex,
				}
				for _, groupIndex := range dim.GroupIndexList {
					info.GroupIndexList[groupIndex] = true
				}

				goppFloor.DimensionList[dim.ID] = info
			}

			g.GoppFloorMap[planeId][floorId] = goppFloor
		}
	}
	logger.Info(text.GetText(17), len(g.GoppFloorMap), "GoppLevelFloor")
}

func GetGoppRtLevelDimensionInfo(planeID, floorID, dimId uint32) *GoppRtLevelDimensionInfo {
	f, ok := getConf().GoppFloorMap[planeID][floorID]
	if !ok {
		return nil
	}
	return f.DimensionList[dimId]
}
