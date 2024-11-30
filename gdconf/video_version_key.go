package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type VideoVersionKey struct {
	ActivityVideoKeyInfoList []*VideoKeyInfo `json:"activityVideoKeyInfoList"`
	VideoKeyInfoList         []*VideoKeyInfo `json:"videoKeyInfoList"`
}

type VideoKeyInfo struct {
	Id       uint32 `json:"id"`
	VideoKey uint64 `json:"videoKey"`
}

func (g *GameDataConfig) loadVideoVersionKey() {
	g.VideoVersionKey = new(VideoVersionKey)
	name := "VideoVersionKey.json"
	playerElementsFile, err := os.ReadFile(g.dataPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &g.VideoVersionKey)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	logger.Info(text.GetText(17), len(g.VideoVersionKey.VideoKeyInfoList)+len(g.VideoVersionKey.ActivityVideoKeyInfoList), name)
}

func GetVideoVersionKey() *VideoVersionKey {
	return getConf().VideoVersionKey
}
