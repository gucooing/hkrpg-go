package gdconf

import (
	"fmt"
	"os"

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
	g.CocoonConfigMap = make(map[uint32]map[uint32]*CocoonConfig)
	cocoonConfigMap := make([]*CocoonConfig, 0)
	playerElementsFilePath := g.excelPrefix + "CocoonConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &cocoonConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, v := range cocoonConfigMap {
		if g.CocoonConfigMap[v.StageID] == nil {
			g.CocoonConfigMap[v.StageID] = make(map[uint32]*CocoonConfig)
		}
		g.CocoonConfigMap[v.StageID][v.WorldLevel] = v
	}

	logger.Info("load %v CocoonConfig", len(g.CocoonConfigMap))
}

func GetCocoonConfigById(stageID, worldLevel uint32) *CocoonConfig {
	return CONF.CocoonConfigMap[stageID][worldLevel]
}
