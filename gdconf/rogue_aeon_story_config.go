package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueAeonStoryConfig struct {
	RogueAeonID uint32 `json:"RogueAeonID"`
	AeonStoryID uint32 `json:"AeonStoryID"`
	UnlockID    uint32 `json:"UnlockID"`
}

func (g *GameDataConfig) loadRogueAeonStoryConfig() {
	g.RogueAeonStoryConfigMap = make(map[uint32]map[uint32]*RogueAeonStoryConfig)
	rogueAeonStoryConfigList := make([]*RogueAeonStoryConfig, 0)
	name := "RogueAeonStoryConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueAeonStoryConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueAeonStoryConfigList {
		if g.RogueAeonStoryConfigMap[v.RogueAeonID] == nil {
			g.RogueAeonStoryConfigMap[v.RogueAeonID] = make(map[uint32]*RogueAeonStoryConfig)
		}
		g.RogueAeonStoryConfigMap[v.RogueAeonID][v.AeonStoryID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueAeonStoryConfigMap), name)
}

func GetRogueAeonStoryConfigMap() map[uint32]map[uint32]*RogueAeonStoryConfig {
	return getConf().RogueAeonStoryConfigMap
}
