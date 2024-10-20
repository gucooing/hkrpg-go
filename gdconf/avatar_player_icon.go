package gdconf

import (
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AvatarPlayerIcon struct {
	ID       uint32 `json:"ID"`
	AvatarID uint32 `json:"AvatarID"`
}

func (g *GameDataConfig) loadAvatarPlayerIcon() {
	g.AvatarPlayerIconMap = make(map[uint32]*AvatarPlayerIcon)
	avatarPlayerIconList := make([]*AvatarPlayerIcon, 0)
	playerElementsFilePath := g.excelPrefix + "AvatarPlayerIcon.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		logger.Error("open file error: %v", err)
		return
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarPlayerIconList)
	if err != nil {
		logger.Error("parse file error: %v", err)
		return
	}
	for _, v := range avatarPlayerIconList {
		g.AvatarPlayerIconMap[v.AvatarID] = v
	}
	logger.Info("load %v AvatarPlayerIcon", len(g.AvatarPlayerIconMap))
}

func GetAvatarPlayerIcon(avatarID uint32) *AvatarPlayerIcon {
	return CONF.AvatarPlayerIconMap[avatarID]
}
