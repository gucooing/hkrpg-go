package gdconf

import (
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type VideoVersionKey struct {
	Id       uint32 `json:"ID"`
	VideoKey uint64 `json:"VideoKey"`
}

func (g *GameDataConfig) loadVideoVersionKey() {
	g.VideoVersionKey = make([]*VideoVersionKey, 0)
	playerElementsFilePath := g.dataPrefix + "VideoVersionKey.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		logger.Error("open file error: %v", err)
		return
	}

	err = hjson.Unmarshal(playerElementsFile, &g.VideoVersionKey)
	if err != nil {
		logger.Error("parse file error: %v", err)
		return
	}
	logger.Info("load %v VideoVersionKey", len(g.VideoVersionKey))

}

func GetVideoVersionKey() []*VideoVersionKey {
	return CONF.VideoVersionKey
}
