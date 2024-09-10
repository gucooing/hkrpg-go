package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "ContentPackageConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &contentPackageConfigList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, v := range contentPackageConfigList {
		g.ContentPackageConfigMap[v.ContentID] = v
	}

	logger.Info("load %v ContentPackageConfig", len(g.ContentPackageConfigMap))
}

func GetContentPackageConfigMap() map[uint32]*ContentPackageConfig {
	return CONF.ContentPackageConfigMap
}
