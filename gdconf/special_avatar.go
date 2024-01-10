package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type SpecialAvatar struct {
	SpecialAvatarID uint32 `json:"SpecialAvatarID"`
	WorldLevel      uint32 `json:"WorldLevel"`
	PlayerID        uint32 `json:"PlayerID"`
	AvatarID        uint32 `json:"AvatarID"`
}

func (g *GameDataConfig) loadSpecialAvatar() {
	g.SpecialAvatarMap = make(map[string]map[string]*SpecialAvatar)
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
}

func GetSpecialAvatarById(stageID uint32) *SpecialAvatar {
	return CONF.SpecialAvatarMap[strconv.Itoa(int(stageID))]["0"]
}
