package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "CocoonConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &cocoonConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range cocoonConfigMap {
		if g.CocoonConfigMap[v.ID] == nil {
			g.CocoonConfigMap[v.ID] = make(map[uint32]*CocoonConfig)
		}
		g.CocoonConfigMap[v.ID][v.WorldLevel] = v
	}

	logger.Info(text.GetText(17), len(g.CocoonConfigMap), name)
}

func GetCocoonConfigById(id, worldLevel uint32) *CocoonConfig {
	return getConf().CocoonConfigMap[id][worldLevel]
}
