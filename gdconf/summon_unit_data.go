package gdconf

import (
	"fmt"
	"os"
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type SummonUnitDataInfo struct {
	SummonUnitDataMap     map[uint32]*SummonUnitData
	SummonUnitDataJsonMap map[uint32]*SummonUnitDataJson
}

type SummonUnitData struct {
	ID                      uint32 `json:"ID"`
	JsonPath                string `json:"JSONPath"`
	IsClient                bool   `json:"IsClient"`
	IsTeamSummon            bool   `json:"IsTeamSummon"`
	DestroyOnEnterBattle    bool   `json:"DestroyOnEnterBattle"`
	RemoveMazeBuffOnDestroy bool   `json:"RemoveMazeBuffOnDestroy"`
	MaxSummonCount          uint32 `json:"MaxSummonCount"`
	UniqueGroup             string `json:"UniqueGroup"`
}

type SummonUnitDataJson struct {
	AttachPoint   string                        `json:"AttachPoint"`
	TriggerConfig *SummonUnitTriggers           `json:"TriggerConfig"`
	Actions       map[string][]*MazeSkillAction `json:"-"`
}

type SummonUnitTriggers struct {
	CustomTriggers []*SummonUnitCustomTrigger `json:"CustomTriggers"`
}

type SummonUnitCustomTrigger struct {
	TriggerName    string      `json:"TriggerName"`
	OnTriggerEnter []*TaskInfo `json:"OnTriggerEnter"`
}

func (g *GameDataConfig) loadSummonUnitData() {
	g.SummonUnitDataInfo = &SummonUnitDataInfo{
		SummonUnitDataMap:     make(map[uint32]*SummonUnitData),
		SummonUnitDataJsonMap: make(map[uint32]*SummonUnitDataJson),
	}
	summonUnitDataList := make([]*SummonUnitData, 0)
	playerElementsFilePath := g.excelPrefix + "SummonUnitData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &summonUnitDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range summonUnitDataList {
		jsonData := new(SummonUnitDataJson)
		confElementsFile, err := os.ReadFile(g.pathPrefix + "/" + v.JsonPath)
		if err != nil {
			logger.Error("open file error: %v", err)
			continue
		}
		err = hjson.Unmarshal(confElementsFile, &jsonData)
		if err != nil {
			logger.Error("parse file error: %v", err)
			continue
		}
		jsonData.Actions = make(map[string][]*MazeSkillAction)
		if jsonData.TriggerConfig != nil && jsonData.TriggerConfig.CustomTriggers != nil {
			for _, customTrigger := range jsonData.TriggerConfig.CustomTriggers {
				jsonData.Actions[customTrigger.TriggerName] = BuildSummonUnitMazeSkillActions(customTrigger)
			}
		}
		g.SummonUnitDataInfo.SummonUnitDataJsonMap[v.ID] = jsonData
		g.SummonUnitDataInfo.SummonUnitDataMap[v.ID] = v
	}

	logger.Info("load %v SummonUnitData", len(g.SummonUnitDataInfo.SummonUnitDataMap))
	logger.Info("load %v SummonUnitDataJson", len(g.SummonUnitDataInfo.SummonUnitDataJsonMap))
}

func GetSummonUnitData(summonId uint32) *SummonUnitData {
	return CONF.SummonUnitDataInfo.SummonUnitDataMap[summonId]
}

func BuildSummonUnitMazeSkillActions(customTriggers *SummonUnitCustomTrigger) []*MazeSkillAction {
	actionList := make([]*MazeSkillAction, 0)
	for _, task := range customTriggers.OnTriggerEnter {
		if strings.Contains(task.Type, "AddMazeBuff") {
			actionList = append(actionList, &MazeSkillAction{
				Type: constant.AddMazeBuff,
				Id:   task.ID,
			})
		} else if strings.Contains(task.Type, "TriggerHitProp") {

		}
	}
	return actionList
}

func GetSummonUnitMazeSkillAction(summonId uint32, triggerName string) []*MazeSkillAction {
	if CONF.SummonUnitDataInfo.SummonUnitDataJsonMap[summonId] == nil ||
		CONF.SummonUnitDataInfo.SummonUnitDataJsonMap[summonId].Actions == nil {
		return nil
	}
	return CONF.SummonUnitDataInfo.SummonUnitDataJsonMap[summonId].Actions[triggerName]
}
