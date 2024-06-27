package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	FinishType        constant.QuestFinishType  `json:"FinishType"`
	ParamType         string                    `json:"ParamType"`
	ParamInt1         uint32                    `json:"ParamInt1"`
	ParamInt2         uint32                    `json:"ParamInt2"`
	ParamInt3         uint32                    `json:"ParamInt3"`
	ParamStr1         string                    `json:"ParamStr1"`
	ParamIntList      []uint32                  `json:"ParamIntList"`
	ParamItemList     []*ParamItem              `json:"ParamItemList"`
	FinishActionList  []*FinishAction           `json:"FinishActionList"`
	Progress          uint32                    `json:"Progress"`
	IsShow            bool                      `json:"IsShow"`
	WayPointFloorID   uint32                    `json:"WayPointFloorID"`
	MapNPCList        []*MapNPC                 `json:"MapNPCList"`
	MapPropList       []*MapProp                `json:"MapPropList"`
}

type ParamItem struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}

type FinishAction struct {
	FinishActionType constant.FinishActionType `json:"FinishActionType"`
	FinishActionPara []uint32                  `json:"FinishActionPara"`
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
	Type           string            `json:"Type"`
}

type OnStartSequece struct {
	TaskList []*Task `json:"TaskList"`
}

type Task struct {
	Type string `json:"$type"`
	// 传送相关
	EntranceID uint32 `json:"EntranceID"`
	// 战斗相关
	EventID      *EventID    `json:"EventID"`
	GroupID      interface{} `json:"GroupID"`
	BattleAreaID interface{} `json:"BattleAreaID"`
	TaskEnabled  bool        `json:"TaskEnabled"`
}

type EventID struct {
	FixedValue *FixedValue `json:"FixedValue"`
}

type GroupID struct {
	IsDynamic  bool        `json:"IsDynamic"`
	FixedValue *FixedValue `json:"FixedValue"`
}

type FixedValue struct {
	Value uint32 `json:"Value"`
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
			logger.Debug("open MainMission error: %v", err)

			return
		}
		err = hjson.Unmarshal(playerElementsFile, &goppMainMission)
		if err != nil {
			info := fmt.Sprintf("parse MainMission error: %v", err)
			panic(info)
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
					logger.Debug("open MissionJsonPath error:%s", err)
				} else {
					g.GoppMission.GoppMissionJson[subMission.ID] = mj
				}
			}
			g.GoppMission.GoppSubMainMission[subMission.ID] = subMission
		}
	}

	logger.Info("gopp %v MainMission", len(g.GoppMission.GoppMainMission))
	logger.Info("gopp %v SubMainMission", len(g.GoppMission.GoppSubMainMission))
	logger.Info("gopp %v MissionJson", len(g.GoppMission.GoppMissionJson))
}

func GetGoppMainMission() map[uint32]*GoppMainMission {
	return CONF.GoppMission.GoppMainMission
}

func GetGoppMainMissionById(id uint32) *GoppMainMission {
	return CONF.GoppMission.GoppMainMission[id]
}

func GetSubMainMission() map[uint32]*SubMission {
	return CONF.GoppMission.GoppSubMainMission
}

func GetSubMainMissionById(id uint32) *SubMission {
	return CONF.GoppMission.GoppSubMainMission[id]
}

func GetEntryId(id uint32) (uint32, bool) {
	var entryId uint32
	var isFloor = false
	conf := CONF.GoppMission.GoppMissionJson[id]
	if conf == nil {
		return 0, false
	}
	for _, info := range conf.OnStartSequece {
		if info.TaskList == nil {
			continue
		}
		for _, task := range info.TaskList {
			entryId = task.EntranceID
			if CONF.MapEntranceMap[entryId] != nil {
				isFloor = true
				break
			}
		}
	}
	return entryId, isFloor
}

func IsBattleMission(id, eventId uint32) bool {
	isFinish := false
	conf := CONF.GoppMission.GoppMissionJson[id]
	if conf == nil {
		return false
	}
	for _, info := range conf.OnStartSequece {
		if info.TaskList == nil {
			continue
		}
		for _, task := range info.TaskList {
			if task.EventID != nil && task.EventID.FixedValue.Value == eventId {
				isFinish = true
				break
			}
		}
	}

	return isFinish
}
