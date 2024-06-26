package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type SpecialAvatar struct {
	SpecialAvatarID    uint32 `json:"SpecialAvatarID"`
	WorldLevel         uint32 `json:"WorldLevel"`
	PlayerID           uint32 `json:"PlayerID"`
	AvatarID           uint32 `json:"AvatarID"`
	Level              uint32 `json:"Level"`
	Promotion          uint32 `json:"Promotion"`
	HaveActionDelay    bool   `json:"HaveActionDelay"`
	EquipmentID        uint32 `json:"EquipmentID"`
	EquipmentLevel     uint32 `json:"EquipmentLevel"`
	EquipmentPromotion uint32 `json:"EquipmentPromotion"`
	EquipmentRank      uint32 `json:"EquipmentRank"`
	RelicPropertyType  uint32 `json:"RelicPropertyType"`
	RelicMainValue     uint32 `json:"RelicMainValue"`
	RelicSubValue      uint32 `json:"RelicSubValue"`
}

func (g *GameDataConfig) loadSpecialAvatar() {
	g.SpecialAvatarMap = make(map[uint32]map[uint32]*SpecialAvatar)
	playerElementsFilePath := g.excelPrefix + "SpecialAvatar.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.SpecialAvatarMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v SpecialAvatar", len(g.SpecialAvatarMap))
	g.wg.Done()
}

func GetSpecialAvatarById(stageID uint32) *SpecialAvatar {
	return CONF.SpecialAvatarMap[stageID][0]
}
