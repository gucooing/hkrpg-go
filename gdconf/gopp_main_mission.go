package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type GoppMission struct {
	GoppMainMission    map[uint32]*GoppMainMission // 预处理主线任务
	GoppSubMainMission map[uint32]*SubMission      // 预处理主线子任务
	GoppMissionJson    map[uint32]*MissionJson     // json中任务配置
}

type GoppMainMission struct {
	MainMissionID        uint32        `json:"MainMissionID"`
	StartSubMissionList  []uint32      `json:"StartSubMissionList"`
	FinishSubMissionList []uint32      `json:"FinishSubMissionList"`
	SubMissionList       []*SubMission `json:"SubMissionList"`
}

type SubMission struct {
	ID                uint32                    `json:"ID"`
	MainMissionID     uint32                    `json:"MainMissionID"`
	MissionJsonPath   string                    `json:"MissionJsonPath"`
	LevelPlaneID      uint32                    `json:"LevelPlaneID"`
	LevelFloorID      uint32                    `json:"LevelFloorID"`
	AudioEmotionState string                    `json:"AudioEmotionState"`
	TakeType          constant.MissionBeginType `json:"TakeType"`
	TakeParamIntList  []uint32                  `json:"TakeParamIntList"`
	MazePlaneID       uint32                    `json:"MazePlaneID"`
	MazeFloorID       uint32                    `json:"MazeFloorID"`
	FinishType        constant.QuestFinishType  `json:"FinishType"`
	ParamType         string                    `json:"ParamType"`
	ParamInt1         uint32                    `json:"ParamInt1"`
	ParamInt2         uint32                    `json:"ParamInt2"`
	ParamInt3         uint32                    `json:"ParamInt3"`
	ParamStr1         string                    `json:"ParamStr1"`
	ParamIntList      []uint32                  `json:"ParamIntList"`
	ParamItemList     []*ParamItem              `json:"ParamItemList"`
	SubRewardID       uint32                    `json:"SubRewardID"`
	FinishActionList  []*FinishAction           `json:"FinishActionList"`
	Progress          uint32                    `json:"Progress"`
	IsShow            bool                      `json:"IsShow"`
	WayPointFloorID   uint32                    `json:"WayPointFloorID"`
	WayPointGroupID   uint32                    `json:"WayPointGroupID"`
	WayPointEntityID  uint32                    `json:"WayPointEntityID"`
	MapNPCList        []*MapNPC                 `json:"MapNPCList"`
	MapPropList       []*MapProp                `json:"MapPropList"`
}

type ParamItem struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}

type FinishAction struct {
	FinishActionType       constant.FinishActionType `json:"FinishActionType"`
	FinishActionPara       []uint32                  `json:"FinishActionPara"`
	FinishActionParaString []string                  `json:"FinishActionParaString"`
}

type MapNPC struct {
	GroupID uint32 `json:"GroupID"`
	NPCID   uint32 `json:"NPCID"`
}

type MapProp struct {
	GroupID uint32 `json:"GroupID"`
	PropID  uint32 `json:"PropID"`
}

type MissionJson struct {
	OnStartSequece []*OnStartSequece `json:"OnStartSequece"`
	OnInitSequece  []*OnStartSequece `json:"OnInitSequece"`
	Type           string            `json:"Type"`
}

type OnStartSequece struct {
	TaskList []*Task `json:"TaskList"`
}

type Task struct {
	Type string `json:"$type"`
	// 传送相关
	EntranceID  uint32 `json:"EntranceID"`
	TaskEnabled bool   `json:"TaskEnabled"`
	// 战斗相关
	EventID      *EventID    `json:"EventID"`
	GroupID      interface{} `json:"GroupID"`
	AnchorID     interface{} `json:"AnchorID"`
	BattleAreaID interface{} `json:"BattleAreaID"`
}

type EventID struct {
	FixedValue *FixedValue `json:"FixedValue"`
}

type GroupID struct {
	IsDynamic  bool        `json:"IsDynamic"`
	FixedValue *FixedValue `json:"FixedValue"`
}

type FixedValue struct {
	Value float64 `json:"Value"`
}

func (g *GameDataConfig) goppMainMission() {
	g.GoppMission = &GoppMission{
		GoppMainMission:    make(map[uint32]*GoppMainMission),
		GoppSubMainMission: make(map[uint32]*SubMission),
	}

	for id := range GetMainMission() {
		goppMainMission := new(GoppMainMission)
		playerElementsFilePath := g.configPrefix + "Level/Mission/" + strconv.Itoa(int(id)) + "/MissionInfo_" + strconv.Itoa(int(id)) + ".json"
		playerElementsFile, err := os.ReadFile(playerElementsFilePath)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(18), playerElementsFilePath, err))
		}
		err = hjson.Unmarshal(playerElementsFile, &goppMainMission)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(19), playerElementsFilePath, err))
		}
		if g.GoppMission.GoppMainMission == nil {
			g.GoppMission.GoppMainMission = make(map[uint32]*GoppMainMission)
		}
		g.GoppMission.GoppMainMission[id] = goppMainMission
		for _, subMission := range goppMainMission.SubMissionList {
			if g.GoppMission.GoppSubMainMission == nil {
				g.GoppMission.GoppSubMainMission = make(map[uint32]*SubMission)
			}
			if g.GoppMission.GoppMissionJson == nil {
				g.GoppMission.GoppMissionJson = make(map[uint32]*MissionJson)
			}
			missionJsonPathFilePath := g.resPrefix + subMission.MissionJsonPath
			missionJsonPathFile, err := os.ReadFile(missionJsonPathFilePath)
			if err == nil {
				mj := new(MissionJson)
				err = hjson.Unmarshal(missionJsonPathFile, &mj)
				if err != nil {
					logger.Error(text.GetText(19), missionJsonPathFilePath, err)
				} else {
					g.GoppMission.GoppMissionJson[subMission.ID] = mj
				}
			}
			g.GoppMission.GoppSubMainMission[subMission.ID] = subMission
		}
	}

	logger.Info(text.GetText(17), len(g.GoppMission.GoppMainMission), "MainMission")
	logger.Info(text.GetText(17), len(g.GoppMission.GoppSubMainMission), "SubMainMission")
	logger.Info(text.GetText(17), len(g.GoppMission.GoppMissionJson), "MissionJson")
}

func GetGoppMainMission() map[uint32]*GoppMainMission {
	return getConf().GoppMission.GoppMainMission
}

func GetGoppMainMissionById(id uint32) *GoppMainMission {
	return getConf().GoppMission.GoppMainMission[id]
}

func GetSubMainMission() map[uint32]*SubMission {
	return getConf().GoppMission.GoppSubMainMission
}

func GetSubMainMissionById(id uint32) *SubMission {
	return getConf().GoppMission.GoppSubMainMission[id]
}

func GetEntryId(id uint32) (uint32, uint32, uint32, bool) {
	conf := GetSubMainMissionById(id)
	jsonConf := getConf().GoppMission.GoppMissionJson[id]
	if jsonConf != nil {
		for _, info := range jsonConf.OnStartSequece {
			if info.TaskList == nil {
				continue
			}
			for _, task := range info.TaskList {
				if getConf().MapEntranceMap[task.EntranceID] != nil {
					return task.EntranceID, getGroupIDUint32(task.GroupID), getGroupIDUint32(task.AnchorID), true
				}
			}
		}
		for _, info := range jsonConf.OnInitSequece {
			if info.TaskList == nil {
				continue
			}
			for _, task := range info.TaskList {
				if getConf().MapEntranceMap[task.EntranceID] != nil {
					return task.EntranceID, getGroupIDUint32(task.GroupID), getGroupIDUint32(task.AnchorID), true
				}
			}
		}
	}
	if conf == nil {
		return 0, 0, 0, false
	}
	str := strconv.Itoa(int(conf.ParamInt2))
	part1 := str[:6]
	part2 := str[6:7]
	newNumStr := part1 + part2
	return alg.S2U32(newNumStr), 0, 0, true
}

func IsBattleMission(id, eventId uint32) bool {
	isFinish := false
	conf := getConf().GoppMission.GoppMissionJson[id]
	if conf == nil {
		return false
	}
	for _, info := range conf.OnStartSequece {
		if info.TaskList == nil {
			continue
		}
		for _, task := range info.TaskList {
			if task.EventID != nil && uint32(task.EventID.FixedValue.Value) == eventId {
				isFinish = true
				break
			}
		}
	}

	return isFinish
}

func getGroupIDUint32(x interface{}) uint32 {
	switch x.(type) {
	case uint32:
		return x.(uint32)
	case *GroupID:
		return uint32(x.(GroupID).FixedValue.Value)
	}
	return 0
}
