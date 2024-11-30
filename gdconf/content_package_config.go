package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ContentPackageConfig struct {
	ContentID            uint32   `json:"ContentID"`
	MainMissionIDList    []uint32 `json:"MainMissionIDList"`
	EarlyAccessCondition string   `json:"EarlyAccessCondition"`
	EarlyFinishCondition string   `json:"EarlyFinishCondition"`
	ReleaseCondition     string   `json:"ReleaseCondition"`
	InitEntranceID       uint32   `json:"InitEntranceID"`
	GuideConditions      string   `json:"GuideConditions"`
	AfterGuideEntranceID uint32   `json:"AfterGuideEntranceID"`
}

func (g *GameDataConfig) loadContentPackageConfig() {
	g.ContentPackageConfigMap = make(map[uint32]*ContentPackageConfig)
	contentPackageConfigList := make([]*ContentPackageConfig, 0)
	name := "ContentPackageConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &contentPackageConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range contentPackageConfigList {
		g.ContentPackageConfigMap[v.ContentID] = v
	}

	logger.Info(text.GetText(17), len(g.ContentPackageConfigMap), name)
}

func GetContentPackageConfigMap() map[uint32]*ContentPackageConfig {
	return getConf().ContentPackageConfigMap
}
