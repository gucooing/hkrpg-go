package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type GoppMainMission struct {
	MainMissionID        uint32        `json:"MainMissionID"`
	StartSubMissionList  []uint32      `json:"StartSubMissionList"`
	FinishSubMissionList []uint32      `json:"FinishSubMissionList"`
	SubMissionList       []*SubMission `json:"SubMissionList"`
}

type SubMission struct {
	ID                uint32          `json:"ID"`
	MainMissionID     uint32          `json:"MainMissionID"`
	MissionJsonPath   string          `json:"MissionJsonPath"`
	LevelPlaneID      uint32          `json:"LevelPlaneID"`
	LevelFloorID      uint32          `json:"LevelFloorID"`
	AudioEmotionState string          `json:"AudioEmotionState"`
	TakeType          string          `json:"TakeType"`
	TakeParamIntList  []uint32        `json:"TakeParamIntList"`
	FinishType        string          `json:"FinishType"`
	ParamType         string          `json:"ParamType"`
	ParamStr1         string          `json:"ParamStr1"`
	ParamItemList     []*ParamItem    `json:"ParamItemList"`
	FinishActionList  []*FinishAction `json:"FinishActionList"`
	Progress          uint32          `json:"Progress"`
	IsShow            bool            `json:"IsShow"`
	WayPointFloorID   uint32          `json:"WayPointFloorID"`
	MapNPCList        []*MapNPC       `json:"MapNPCList"`
	MapPropList       []*MapProp      `json:"MapPropList"`
}

type ParamItem struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}

type FinishAction struct {
	FinishActionType string   `json:"FinishActionType"`
	FinishActionPara []uint32 `json:"FinishActionPara"`
}

type MapNPC struct {
	GroupID uint32 `json:"GroupID"`
	NPCID   uint32 `json:"NPCID"`
}

type MapProp struct {
	GroupID uint32 `json:"GroupID"`
	PropID  uint32 `json:"PropID"`
}

func (g *GameDataConfig) goppMainMission() {
	g.GoppMainMission = make(map[uint32]*GoppMainMission)

	for id := range GetMainMission() {
		goppMainMission := new(GoppMainMission)
		playerElementsFilePath := g.configPrefix + "Level/Mission/" + strconv.Itoa(int(id)) + "/MissionInfo_" + strconv.Itoa(int(id)) + ".json"
		playerElementsFile, err := os.ReadFile(playerElementsFilePath)
		if err != nil {
			logger.Debug("open MainMission error: %v", err)
			continue
		}
		err = hjson.Unmarshal(playerElementsFile, &goppMainMission)
		if err != nil {
			info := fmt.Sprintf("parse MainMission error: %v", err)
			panic(info)
		}
		g.GoppMainMission[id] = goppMainMission
	}

	logger.Info("gopp %v MainMission", len(g.GoppMainMission))
}

func GetGoppMainMission() map[uint32]*GoppMainMission {
	return CONF.GoppMainMission
}

func GetGoppMainMissionById(id uint32) *GoppMainMission {
	return CONF.GoppMainMission[id]
}
