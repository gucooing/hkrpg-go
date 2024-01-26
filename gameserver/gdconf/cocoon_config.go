package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type CocoonConfig struct {
	ID            uint32   `json:"ID"`
	WorldLevel    uint32   `json:"WorldLevel"`
	PropID        uint32   `json:"PropID"`
	CocoonType    string   `json:"CocoonType"`
	MappingInfoID uint32   `json:"MappingInfoID"`
	StageID       uint32   `json:"StageID"`
	StageIDList   []uint32 `json:"StageIDList"`
	ParamList     []*Value `json:"ParamList"`
	DropList      []uint32 `json:"DropList"`
	StaminaCost   uint32   `json:"StaminaCost"` // 扣除体力
	MaxWave       uint32   `json:"MaxWave"`
	OpenDate      []uint32 `json:"OpenDate"`
	DamageType    []string `json:"DamageType"`
	FarmType      string   `json:"FarmType"`
}

func (g *GameDataConfig) loadCocoonConfig() {
	g.CocoonConfigMap = make(map[string]map[string]*CocoonConfig)
	playerElementsFilePath := g.excelPrefix + "CocoonConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.CocoonConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v CocoonConfig", len(g.CocoonConfigMap))
}

func GetCocoonConfigById(stageID, worldLevel uint32) *CocoonConfig {
	return CONF.CocoonConfigMap[strconv.Itoa(int(stageID))][strconv.Itoa(int(worldLevel))]
}
