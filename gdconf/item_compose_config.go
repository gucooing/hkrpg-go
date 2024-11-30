package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ItemComposeConfig struct {
	ID                        uint32               `json:"ID"`
	FormulaType               constant.FormulaType `json:"FormulaType"`
	RelicList                 []uint32             `json:"RelicList"`
	SpecialMaterialCost       []uint32             `json:"SpecialMaterialCost"`
	SpecialMaterialCostNumber uint32               `json:"SpecialMaterialCostNumber"`
	ItemID                    uint32               `json:"ItemID"`
	MaterialCost              []*ItemC             `json:"MaterialCost"` // 消耗材料
	CoinCost                  uint32               `json:"CoinCost"`     // Scoin
	Type                      uint32               `json:"Type"`
	Order                     uint32               `json:"Order"`
	WorldLevelRequire         uint32               `json:"WorldLevelRequire"`
	MaxCount                  uint32               `json:"MaxCount"`
	IsShowHoldNumber          bool                 `json:"IsShowHoldNumber"`
	ItemComposeTag            []uint32             `json:"ItemComposeTag"`
	LimitType                 constant.LimitType   `json:"LimitType"`
	FuncType                  constant.FuncType    `json:"FuncType"`
}

type ItemC struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}

func (g *GameDataConfig) loadItemComposeConfig() {
	g.ItemComposeConfigMap = make(map[uint32]*ItemComposeConfig)
	itemComposeConfig := make([]*ItemComposeConfig, 0)
	name := "ItemComposeConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &itemComposeConfig)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range itemComposeConfig {
		g.ItemComposeConfigMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.ItemComposeConfigMap), name)
}

func GetItemComposeConfig(id uint32) *ItemComposeConfig {
	return getConf().ItemComposeConfigMap[id]
}
