package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type EquipmentConfig struct {
	EquipmentID          uint32    `json:"EquipmentID"`
	ItemID               uint32    `json:"ItemID"`
	Release              bool      `json:"Release"`
	Rarity               string    `json:"Rarity"`
	AvatarBaseType       string    `json:"AvatarBaseType"`
	SkillID              uint32    `json:"SkillID"`
	ExpType              uint32    `json:"ExpType"`
	ExpProvide           uint32    `json:"ExpProvide"`
	CoinCost             uint32    `json:"CoinCost"`
	ItemRightPanelOffset []float64 `json:"ItemRightPanelOffset"`
	AvatarDetailOffset   []float64 `json:"AvatarDetailOffset"`
	BattleDialogOffset   []float64 `json:"BattleDialogOffset"`
	GachaResultOffset    []float64 `json:"GachaResultOffset"`
}

func (g *GameDataConfig) loadEquipmentConfig() {
	g.EquipmentConfigMap = make(map[uint32]*EquipmentConfig)

	fileList := []string{"EquipmentConfig.json", "EquipmentExpItemConfig.json"}
	for _, file := range fileList {
		equipmentConfigList := make([]*EquipmentConfig, 0)
		files := g.excelPrefix + file
		bin, err := os.ReadFile(files)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(18), file, err))
		}
		err = hjson.Unmarshal(bin, &equipmentConfigList)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(19), file, err))
		}
		for _, v := range equipmentConfigList {
			if v.EquipmentID == 0 {
				g.EquipmentConfigMap[v.ItemID] = v
			} else {
				g.EquipmentConfigMap[v.EquipmentID] = v
			}
		}
	}

	logger.Info(text.GetText(17), len(g.EquipmentConfigMap), "EquipmentConfig")
}

func GetEquipmentConfigById(ID uint32) *EquipmentConfig {
	return getConf().EquipmentConfigMap[ID]
}

func GetEquipmentConfigMap() map[uint32]*EquipmentConfig {
	return getConf().EquipmentConfigMap
}

func GetEquipmentList() []uint32 {
	var equipmentList []uint32
	for _, equipment := range getConf().EquipmentConfigMap {
		equipmentList = append(equipmentList, equipment.EquipmentID)
	}
	return equipmentList
}
