package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	specialAvatarMap := make([]*SpecialAvatar, 0)
	name := "SpecialAvatar.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &specialAvatarMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range specialAvatarMap {
		if g.SpecialAvatarMap[v.SpecialAvatarID] == nil {
			g.SpecialAvatarMap[v.SpecialAvatarID] = make(map[uint32]*SpecialAvatar)
		}
		g.SpecialAvatarMap[v.SpecialAvatarID][v.WorldLevel] = v
	}

	logger.Info(text.GetText(17), len(g.SpecialAvatarMap), name)
}

func GetSpecialAvatarById(stageID uint32) *SpecialAvatar {
	return getConf().SpecialAvatarMap[stageID][0]
}

func SpecialAvatarGetBaseAvatarID(specialAvatarID uint32) uint32 {
	sac := getConf().SpecialAvatarMap[specialAvatarID][0]
	if sac == nil {
		return specialAvatarID
	}
	if mpac := GetMultiplePathAvatarConfig(sac.AvatarID); mpac != nil {
		return mpac.BaseAvatarID
	}
	return sac.AvatarID
}

func SpecialAvatarGetPlayerID(specialAvatarID uint32) uint32 {
	sac := getConf().SpecialAvatarMap[specialAvatarID][0]
	if sac == nil {
		return specialAvatarID
	}
	if mpac := GetMultiplePathAvatarConfig(sac.AvatarID); mpac != nil {
		return mpac.BaseAvatarID
	}
	return sac.PlayerID
}
