package gdconf

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type LevelGroup struct {
	Index                 uint32 `json:"-"`
	GroupId               uint32
	GroupName             string                              `json:"GroupName"`
	AreaAnchorName        string                              `json:"AreaAnchorName"`
	SaveType              string                              `json:"SaveType"`
	HoYoGroupType         string                              `json:"HoYoGroupType"`
	AtmosphereCondition   *AtmosphereCondition                `json:"AtmosphereCondition"`
	LoadSide              string                              `json:"LoadSide"` // 负载端
	IsHoyoGroup           bool                                `json:"IsHoyoGroup"`
	IsPendedUnload        bool                                `json:"IsPendedUnload"`
	SystemUnlockCondition *LevelGroupSystemUnlockConditionSet `json:"SystemUnlockCondition"`
	SavedValueCondition   *LevelGroupSavedValueConditionSet   `json:"SavedValueCondition"`
	Category              string                              `json:"Category"`             // 类别
	OwnerMainMissionID    uint32                              `json:"OwnerMainMissionID"`   // 主任务id
	LoadCondition         *LevelGroupMissionConditionSet      `json:"LoadCondition"`        // 加载条件
	UnloadCondition       *LevelGroupMissionConditionSet      `json:"UnloadCondition"`      // 卸载条件
	ForceUnloadCondition  *LevelGroupMissionConditionSet      `json:"ForceUnloadCondition"` // 强制卸载条件
	LoadOnInitial         bool                                `json:"LoadOnInitial"`        // 是否默认加载
	PropList              []*PropList                         `json:"PropList"`             // 实体列表
	MonsterList           []*MonsterList                      `json:"MonsterList"`          // 怪物列表
	NPCList               []*NPCList                          `json:"NPCList"`              // NPC列表
	AnchorList            []*AnchorList                       `json:"AnchorList"`           // 锚点列表
}
type AtmosphereCondition struct {
	Conditions []*LevelGroupMissionCondition `json:"Conditions"`
	Operation  string                        `json:"Operation"`
}
type LevelGroupSystemUnlockConditionSet struct {
	Conditions []uint32                    `json:"Conditions"`
	Operation  constant.LogicOperationType `json:"Operation"`
}
type LevelGroupSavedValueConditionSet struct {
	Conditions []*LevelGroupSavedValueCondition `json:"Conditions"`
	Operation  constant.LogicOperationType      `json:"Operation"`
}
type LevelGroupSavedValueCondition struct {
	SavedValueName string               `json:"SavedValueName"`
	Operation      constant.CompareType `json:"Operation"`
	Value          uint32               `json:"Value"`
}
type LevelGroupMissionConditionSet struct {
	Conditions         []*LevelGroupMissionCondition `json:"Conditions"`
	Operation          constant.LogicOperationType   `json:"Operation"`
	DelayToLevelReload bool                          `json:"DelayToLevelReload"`
}
type LevelGroupMissionCondition struct {
	LevelGroupMissionType  constant.LevelGroupMissionType  `json:"Type"`
	LevelGroupMissionPhase constant.LevelGroupMissionPhase `json:"Phase"`
	ID                     uint32                          `json:"ID"`
	SubMissionID           uint32                          `json:"SubMissionID"`
}
type PropList struct {
	ID                       uint32              `json:"ID"`
	PosX                     float64             `json:"PosX"`
	PosY                     float64             `json:"PosY"`
	PosZ                     float64             `json:"PosZ"`
	RotX                     float64             `json:"RotX"`
	RotY                     float64             `json:"RotY"`
	RotZ                     float64             `json:"RotZ "`
	Name                     string              `json:"Name"`
	PropID                   uint32              `json:"PropID"`
	IsDelete                 bool                `json:"IsDelete"`
	IsClientOnly             bool                `json:"IsClientOnly"`
	LoadOnInitial            bool                `json:"LoadOnInitial"`
	IsOverrideInitLevelGraph bool                `json:"IsOverrideInitLevelGraph"`
	CampID                   uint32              `json:"CampID"`
	EventID                  uint32              `json:"EventID"`
	MapLayerID               uint32              `json:"MapLayerID"`
	AnchorGroupID            uint32              `json:"AnchorGroupID"`
	AnchorID                 uint32              `json:"AnchorID"`
	MappingInfoID            uint32              `json:"MappingInfoID"`
	ChestClosed              string              `json:"ChestClosed"`
	State                    string              `json:"State"`
	StageObjectCapture       *StageObjectCapture `json:"StageObjectCapture"`
	ValueSource              *ValueSource        `json:"ValueSource"`
	GoppValue                []*GoppValue        `json:"-"`
	TriggerBattleString      uint32              `json:"-"`
}
type ValueSource struct {
	Values []*Values `json:"Values"`
}
type Values struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}
type AnchorList struct {
	ID         uint32  `json:"ID"`
	PosX       float64 `json:"PosX"`
	PosY       float64 `json:"PosY"`
	PosZ       float64 `json:"PosZ"`
	Name       string  `json:"Name"`
	RotX       float64 `json:"RotX"`
	RotY       float64 `json:"RotY"`
	RotZ       float64 `json:"RotZ "`
	MapLayerID uint32  `json:"MapLayerID"`
}

type MonsterList struct {
	ID            uint32      `json:"ID"`
	PosX          float64     `json:"PosX"`
	PosY          float64     `json:"PosY"`
	PosZ          float64     `json:"PosZ"`
	Name          string      `json:"Name"`
	RotX          float64     `json:"RotX"`
	RotY          float64     `json:"RotY"`
	RotZ          float64     `json:"RotZ "`
	IsDelete      bool        `json:"IsDelete"`
	IsClientOnly  bool        `json:"IsClientOnly"`
	NPCMonsterID  uint32      `json:"NPCMonsterID"`
	CampID        uint32      `json:"CampID"`
	EventID       uint32      `json:"EventID"`
	BattleArea    *BattleArea `json:"BattleArea"`
	PurposeType   string      `json:"PurposeType"`
	FarmElementID uint32      `json:"FarmElementID"`
}

type NPCList struct {
	ID                   uint32   `json:"ID"`
	PosX                 float64  `json:"PosX"`
	PosY                 float64  `json:"PosY"`
	PosZ                 float64  `json:"PosZ"`
	Name                 string   `json:"Name"`
	RotX                 float64  `json:"RotX"`
	RotY                 float64  `json:"RotY"`
	RotZ                 float64  `json:"RotZ "`
	NPCID                uint32   `json:"NPCID"`
	IsDelete             bool     `json:"IsDelete"`
	IsClientOnly         bool     `json:"IsClientOnly"`
	DialogueGroups       []uint32 `json:"DialogueGroups"`
	MapLayerID           uint32   `json:"MapLayerID"`
	BoardShowList        []uint32 `json:"BoardShowList"`
	RaidID               uint32   `json:"RaidID"`
	FirstDialogueGroupID uint32   `json:"FirstDialogueGroupID"`
}

type BattleArea struct {
	GroupID uint32 `json:"GroupID"`
	ID      uint32 `json:"ID"`
}

type StageObjectCapture struct {
	BlockAlias  string `json:"BlockAlias"`
	PrefabAlias string `json:"PrefabAlias"`
}

func (g *GameDataConfig) loadGroup() {
	g.GroupMap = make(map[uint32]map[uint32]map[uint32]*LevelGroup)
	syncs := sync.RWMutex{}
	loadGroups := func(planeId, floorId, index uint32, groupInfo *RtLevelGroupInstanceInfo) {
		func() {
			levelGroup := new(LevelGroup)
			playerElementsFile, err := os.ReadFile(pathPrefix + "/" + groupInfo.GroupPath)
			if err != nil {
				panic(fmt.Sprintf(text.GetText(18), playerElementsFile, err))
			}

			err = hjson.Unmarshal(playerElementsFile, levelGroup)
			if err != nil {
				logger.Error(text.GetText(19), playerElementsFile, err)
				return
			}
			levelGroup.GroupId = groupInfo.ID
			levelGroup.Index = index

			syncs.Lock()
			if g.GroupMap[planeId] == nil {
				g.GroupMap[planeId] = make(map[uint32]map[uint32]*LevelGroup)
			}
			if g.GroupMap[planeId][floorId] == nil {
				g.GroupMap[planeId][floorId] = make(map[uint32]*LevelGroup)
			}
			if g.GroupMap[planeId][floorId][groupInfo.ID] == nil {
				g.GroupMap[planeId][floorId][groupInfo.ID] = levelGroup
			}
			syncs.Unlock()

		}()
	}

	wg := sync.WaitGroup{}
	floor := GetFloor()
	sem := make(chan struct{}, MaxWaitGroup)
	for planeId, floorList := range floor {
		for floorId, floorInfo := range floorList {
			for index, groupInfo := range floorInfo.GroupInstanceList {
				if strings.Contains(groupInfo.GroupPath, "_D") ||
					groupInfo.IsDelete {
					continue
				}
				if MaxWaitGroup == 1 {
					loadGroups(planeId, floorId, uint32(index), groupInfo)
				} else {
					go func() {
						defer func() {
							wg.Done()
							func() { <-sem }()
						}()
						sem <- struct{}{}
						wg.Add(1)
						loadGroups(planeId, floorId, uint32(index), groupInfo)
					}()
				}
			}
		}
	}

	wg.Wait()

	logger.Info(text.GetText(17), len(g.GroupMap), "Groups")
}

func GetNGroupById(planeId, floorId, groupId uint32) *LevelGroup {
	return getConf().GroupMap[planeId][floorId][groupId]
}

func GetGroupById(planeId, floorId uint32) map[uint32]*LevelGroup {
	return getConf().GroupMap[planeId][floorId]
}

func GetGroupMap() map[uint32]map[uint32]map[uint32]*LevelGroup {
	return getConf().GroupMap
}

func scanFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetStateValue(state string) uint32 {
	stateMap := map[string]uint32{
		"Closed":            0,
		"Open":              1,
		"Locked":            0, // 锁定
		"Unlocked":          2,
		"BridgeState1":      3,
		"BridgeState2":      4,
		"BridgeState3":      5,
		"BridgeState4":      6,
		"CheckPointDisable": 8, // 锚点
		"CheckPointEnable":  8,
		"TriggerDisable":    9,
		"TriggerEnable":     10,
		"ChestLocked":       11, //  宝箱
		"ChestClosed":       12,
		"ChestUsed":         13,
		"Elevator1":         14,
		"Elevator2":         15,
		"Elevator3":         16,
		"WaitActive":        17,
		"EventClose":        18, // 事件结束
		"EventOpen":         19, // 事件开始
		"Hidden":            20,
		"TeleportGate0":     21,
		"TeleportGate1":     22,
		"TeleportGate2":     23,
		"TeleportGate3":     24,
		"Destructed":        25,
		"CustomState01":     101,
		"CustomState02":     102,
		"CustomState03":     103,
		"CustomState04":     104,
		"CustomState05":     105,
		"CustomState06":     106,
		"CustomState07":     107,
		"CustomState08":     108,
		"CustomState09":     109,
	}

	value, ok := stateMap[state]
	if !ok {
		return 0
	}

	return value
}

func LoadMonster(groupList *LevelGroup) map[uint32]*MonsterList {
	monsterList := make(map[uint32]*MonsterList)
	if groupList == nil || groupList.MonsterList == nil {
		return nil
	}
	for _, monster := range groupList.MonsterList {
		if monster.IsDelete || monster.IsClientOnly {
			continue
		}
		npcMonsterExcel := GetNPCMonsterId(monster.NPCMonsterID)
		if npcMonsterExcel == nil {
			continue
		}

		monsterList[monster.ID] = monster
	}

	return monsterList
}

func LoadProp(groupList *LevelGroup) map[uint32]*PropList {
	propList := make(map[uint32]*PropList)
	if groupList == nil || groupList.PropList == nil {
		return nil
	}
	for _, prop := range groupList.PropList {
		if prop.IsDelete ||
			prop.IsClientOnly ||
			!prop.LoadOnInitial {
			continue
		}
		MazePropExcel := GetMazePropId(prop.PropID)
		if MazePropExcel == nil {
			continue
		}
		if strings.Contains(prop.Name, "FogDoor") {
			continue
		}
		// 对ValueSource进行预处理
		if prop.ValueSource != nil && prop.ValueSource.Values != nil {
			for _, value := range prop.ValueSource.Values {
				switch value.Value.(type) {
				case string:
					valueStr := value.Value.(string)
					if strings.Contains(value.Key, "Door") ||
						strings.Contains(value.Key, "FlipBridge") ||
						value.Key == "Bridge" ||
						strings.Contains(value.Key, "UnlockTarget") ||
						strings.Contains(value.Key, "Rootcontamination") ||
						strings.Contains(value.Key, "Controller") ||
						strings.Contains(value.Key, "Portal") {
						if prop.GoppValue == nil {
							prop.GoppValue = make([]*GoppValue, 0)
						}
						if groupId, instId, ok := getValue(valueStr); ok {
							prop.GoppValue = append(prop.GoppValue, &GoppValue{
								GroupId: groupId,
								InstId:  instId,
							})
						}
					}
					if strings.Contains(value.Key, "TriggerBattleString") {
						prop.TriggerBattleString = alg.S2U32(valueStr)
					}
				}
			}
		}
		propList[prop.ID] = prop
	}
	return propList
}

func getValue(value string) (uint32, uint32, bool) {
	ok := true
	var groupId uint32
	var instId uint32
	parts := strings.Split(value, ",")
	if len(parts) != 2 {
		ok = false
		return groupId, instId, ok
	}
	num1, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		ok = false
		return groupId, instId, ok
	}
	num2, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		ok = false
		return groupId, instId, ok
	}
	groupId = uint32(num1)
	instId = uint32(num2)
	return groupId, instId, ok
}

func LoadNpc(groupList *LevelGroup, nPCList []*NPCList) map[uint32]*NPCList {
	npcList := make(map[uint32]*NPCList)
	if groupList == nil || groupList.NPCList == nil {
		return nil
	}
	for _, npc := range groupList.NPCList {
		if npc.IsDelete || npc.IsClientOnly { // 过滤不需要发送的
			continue
		}
		NPCDataExcel := GetNPCDataId(npc.NPCID)
		if NPCDataExcel == nil { // 过滤没有的
			continue
		}
		// repeatNpc := false
		// for _, npcl := range nPCList {
		// 	if npcl.NPCID == npc.NPCID {
		// 		repeatNpc = true
		// 		break
		// 	}
		// }
		// if repeatNpc { // 过滤重复的
		// 	continue
		// }

		nPCList = append(nPCList, npc)
		npcList[npc.ID] = npc
	}

	return npcList
}

func LoadAnchor(groupList *LevelGroup) map[uint32]*AnchorList {
	anchorList := make(map[uint32]*AnchorList)
	if groupList == nil || groupList.AnchorList == nil {
		return anchorList
	}
	for _, anchor := range groupList.AnchorList {
		anchorList[anchor.ID] = anchor
	}

	return anchorList
}
