package gdconf

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	"github.com/hjson/hjson-go/v4"
)

type LevelFloor struct {
	FloorID             uint32                      `json:"FloorID"`             // 地图区域ID
	StartGroupIndex     uint32                      `json:"StartGroupIndex"`     // 开始组索引
	StartAnchorID       uint32                      `json:"StartAnchorID"`       // 开始锚点ID
	StageData           string                      `json:"StageData"`           // 场景配置路径
	CameraType          constant.MapCameraType      `json:"CameraType"`          // 关卡相机类型
	LayerToAreaMask     []uint32                    `json:"LayerToAreaMask"`     // Layer对应的NavMesh Area Mask
	LevelFeatureModules []constant.LevelFeatureType `json:"LevelFeatureModules"` // 关卡支持的插件模块
	NavmapConfigPath    string                      `json:"NavmapConfigPath"`    // 导航图路径
	RegionConfigPath    string                      `json:"RegionConfigPath"`    // Region相关配置路径
	CustomValues        []*FloorCustomValueConfig   `json:"CustomValues"`        // 关卡变量列表(客户端)
	GroupInstanceList   []*RtLevelGroupInstanceInfo `json:"GroupInstanceList"`   // 组实例列表
	DimensionList       []*RtLevelDimensionInfo     `json:"DimensionList"`       // 位面列表
	// GroupInstanceCommonMap     map[uint32]*RtLevelGroupInstanceCommonInfo `json:"GroupInstanceCommonMap"`   // 在加载Floor时就需要了解的Group位面无关公共信息
	IsRestartLevelGraph        bool    `json:"IsRestartLevelGraph"`      // 跨场景后是否重新执行levelGraph
	EnableGroupStreaming       bool    `json:"EnableGroupStreaming"`     // 开启GroupStreaming功能
	EnableGroupSpaceConflict   bool    `json:"EnableGroupSpaceConflict"` // 开启Group空间冲突检测
	EnableGroupRegionStreaming bool    `json:"EnableGroupRegionStreaming"`
	TempGroupUnloadByY         float32 `json:"TempGroupUnloadByY"` // 在特定地图Entity的Y距离超限制直接卸载，为了解决头顶Entity穿帮问题
}
type RtLevelGroupInstanceInfo struct {
	ID        uint32 `json:"ID"`
	Name      string `json:"Name"`
	GroupPath string `json:"GroupPath"`
	IsDelete  bool   `json:"IsDelete"`
}
type FloorCustomValueConfig struct {
	ID         uint32 `json:"ID"`   // 关卡变量ID
	Name       string `json:"Name"` // 关卡变量名
	GroupID    uint32 `json:"GroupID"`
	InstanceID uint32 `json:"InstanceID"`
	AgentType  string `json:"AgentType"`
}

type RtLevelDimensionInfo struct {
	ID                   uint32                          `json:"ID"` // 位面ID, 0是基础位面
	Category             constant.LevelDimensionCategory `json:"Category"`
	DefaultEnviroProfile string                          `json:"DefaultEnviroProfile"` // 环境配置路径
	CameraType           constant.MapCameraType          `json:"CameraType"`           // 关卡相机类型
	// BattleAreaList               []*LevelBattleAreaInfo                `json:"BattleAreaList"`       // 战斗区域列表
	CameraPrefabPath             string                                `json:"CameraPrefabPath"`
	ConstValues                  []*FloorCustomValueConfig             `json:"ConstValues"`                  // 关卡常量列表
	SavedValues                  []*FloorSavedValueConfig              `json:"SavedValues"`                  // 关卡存档变量列表
	SavedValueConfigDict         map[string]*SavedValueDimensionConfig `json:"SavedValueConfigDict"`         // 关卡存档变量位面配置
	IsExclusiveSaveMapSection    bool                                  `json:"IsExclusiveSaveMapSection"`    // 是否在该位面单独存档迷雾状态
	UnlockAllMapSectionOnInitial bool                                  `json:"UnlockAllMapSectionOnInitial"` // 是否在该位面初始解锁迷雾
	GroupIndexList               []uint32                              `json:"GroupIndexList"`               // 组索引列表
	ExclusiveSavedGroupIDList    []uint32                              `json:"ExclusiveSavedGroupIDList"`    // 单独存档的组的ID列表
	StartAnchorID                uint32                                `json:"StartAnchorID"`                // 开始锚点ID
	StartGroupIndex              uint32                                `json:"StartGroupIndex"`              // 开始组索引
}

type FloorSavedValueConfig struct {
	ID            uint32  `json:"ID"`            // 关卡存档变量ID
	IsDelete      bool    `json:"IsDelete"`      // 是否被标记删除
	Name          string  `json:"Name"`          // 关卡存档变量名
	DefaultValue  int32   `json:"DefaultValue"`  // 默认值
	AllowedValues []int32 `json:"AllowedValues"` // 合法值列表
	MaxValue      int32   `json:"MaxValue"`      // 最大值
	MinValue      int32   `json:"MinValue"`      // 最小值
}

type SavedValueDimensionConfig struct {
	IsExclusiveSaved          bool   `json:"IsExclusiveSaved"`          // 是否单独存档
	IsNeedMergeBack           bool   `json:"IsNeedMergeBack"`           // 是否需要合回主位面
	RelatedContentID          uint32 `json:"RelatedContentID"`          // 关联的ContentID，不为0时才有效，需求来源：https://www.tapd.cn/41546042/prong/stories/view/1141546042002531844
	RelatedContentDimensionID uint32 `json:"RelatedContentDimensionID"` // 关联的Content该与哪个位面共享存档
}

func (g *GameDataConfig) loadFloor() {
	g.FloorMap = make(map[uint32]map[uint32]*LevelFloor)
	playerElementsFilePath := g.configPrefix + "LevelOutput/RuntimeFloor"
	files, err := scanFiles(playerElementsFilePath)
	if err != nil {
		logger.Error(text.GetText(16), playerElementsFilePath)
		return
	}

	for _, file := range files {
		levelFloor := new(LevelFloor)
		planeId, floorId := extractNumbersFloor(filepath.Base(file))

		playerElementsFile, err := os.ReadFile(file)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(18), file, err))
		}

		err = hjson.Unmarshal(playerElementsFile, levelFloor)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(19), file, err))
		}

		if g.FloorMap[planeId] == nil {
			g.FloorMap[planeId] = make(map[uint32]*LevelFloor)
		}
		if g.FloorMap[planeId][floorId] == nil {
			g.FloorMap[planeId][floorId] = new(LevelFloor)
		}
		g.FloorMap[planeId][floorId] = levelFloor
	}

	logger.Info(text.GetText(17), len(g.FloorMap), "Floor")
	g.loadGroup() // 场景实体
}

func GetFloor() map[uint32]map[uint32]*LevelFloor {
	return getConf().FloorMap
}

func GetFloorById(planeId, floorId uint32) *LevelFloor {
	return getConf().FloorMap[planeId][floorId]
}

func GetFloorMap() map[uint32]map[uint32]*LevelFloor {
	return getConf().FloorMap
}

func extractNumbersFloor(filename string) (uint32, uint32) {
	filename = strings.TrimSuffix(filename, ".json")

	parts := strings.Split(filename, "_")
	if len(parts) != 2 {
		return 0, 0
	}

	pValueStr := strings.TrimLeft(parts[0], "P")
	fValueStr := strings.TrimLeft(parts[1], "F")

	pValue, _ := strconv.ParseUint(pValueStr, 10, 32)
	fValue, _ := strconv.ParseUint(fValueStr, 10, 32)

	return uint32(pValue), uint32(fValue)
}

func GetAnchorByIndex(planeId, floorId uint32) *AnchorList {
	floor := GetFloorById(planeId, floorId)
	if floor == nil {
		return nil
	}
	if uint32(len(floor.GroupInstanceList)) < floor.StartGroupIndex {
		return nil
	}
	groupInstance := floor.GroupInstanceList[floor.StartGroupIndex]
	if groupInstance == nil {
		return nil
	}
	group := GetNGroupById(planeId, floorId, groupInstance.ID)
	if group == nil {
		return nil
	}
	// if uint32(len(group.AnchorList)) < floor.StartAnchorID {
	// 	return nil
	// }
	// return group.AnchorList[floor.StartAnchorID]
	for _, anchorInfo := range group.AnchorList {
		if anchorInfo.ID == floor.StartAnchorID {
			return anchorInfo
		}
	}
	return nil
}

func GetAnchor(planeId, floorId, startGroupID, startAnchorID uint32) *AnchorList {
	if startGroupID == 0 || startAnchorID == 0 {
		return GetAnchorByIndex(planeId, floorId)
	}
	group := GetNGroupById(planeId, floorId, startGroupID)
	if group == nil {
		return nil
	}
	for _, anchorInfo := range group.AnchorList {
		if anchorInfo.ID == startAnchorID {
			return anchorInfo
		}
	}
	return nil
}

func GetAnchorByIndexPosRot(planeId, floorId uint32) (*proto.Vector, *proto.Vector) {
	conf := GetAnchorByIndex(planeId, floorId)
	if conf == nil {
		return nil, nil
	}
	return &proto.Vector{
			X: int32(conf.PosX * 1000),
			Y: int32(conf.PosY * 1000),
			Z: int32(conf.PosZ * 1000),
		}, &proto.Vector{
			X: int32(conf.RotX * 1000),
			Y: int32(conf.RotY * 1000),
			Z: int32(conf.RotZ * 1000),
		}
}

func GetSavedValue(planeId, floorId uint32, name string) (uint32, uint32) {
	// floor := GetFloorById(planeId, floorId)
	// var group *LevelGroup
	// var porp *PropList
	// if floor != nil && floor.SavedValues != nil {
	// 	for _, v := range floor.SavedValues {
	// 		if v.Name == name {
	// 			if len(v.AllowedValues) < 2 {
	// 				return 0, 0
	// 			}
	// 			groupInstance := floor.GroupInstanceList[v.AllowedValues[0]]
	// 			if groupInstance == nil {
	// 				return 0, 0
	// 			}
	// 			group = GetNGroupById(planeId, floorId, groupInstance.ID)
	// 			if group == nil {
	// 				return 0, 0
	// 			}
	// 			if uint32(len(group.PropList)) < v.AllowedValues[1] {
	// 				return 0, 0
	// 			}
	// 			porp = group.PropList[v.AllowedValues[1]]
	// 		}
	// 	}
	// }
	// if group != nil && porp != nil {
	// 	return group.GroupId, porp.ID
	// }
	return 0, 0
}
