package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type AvatarPlayerIcon struct {
	ID       uint32 `json:"ID"`
	AvatarID uint32 `json:"AvatarID"`
}

func (g *GameDataConfig) loadAvatarPlayerIcon() {
	g.AvatarPlayerIconMap = make(map[uint32]*AvatarPlayerIcon)
	avatarPlayerIconList := make([]*AvatarPlayerIcon, 0)
	name := "AvatarPlayerIcon.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarPlayerIconList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range avatarPlayerIconList {
		g.AvatarPlayerIconMap[v.AvatarID] = v
	}
	logger.Info(text.GetText(17), len(g.AvatarPlayerIconMap), name)
}

func GetAvatarPlayerIcon(avatarID uint32) *AvatarPlayerIcon {
	return getConf().AvatarPlayerIconMap[avatarID]
}
