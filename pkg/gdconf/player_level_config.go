package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type PlayerLevelConfig struct {
	Level         uint32 `json:"Level"`
	PlayerExp     uint32 `json:"PlayerExp"`
	StaminaLimit  uint32 `json:"StaminaLimit"`
	LevelRewardID uint32 `json:"LevelRewardID"`
}

func (g *GameDataConfig) loadPlayerLevelConfig() {
	g.PlayerLevelConfigMap = make(map[uint32]*PlayerLevelConfig)
	playerElementsFilePath := g.excelPrefix + "PlayerLevelConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.PlayerLevelConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v PlayerLevelConfig", len(g.PlayerLevelConfigMap))
}

func GetPlayerLevelConfig(level uint32) *PlayerLevelConfig {
	return CONF.PlayerLevelConfigMap[level]
}

func GetPlayerLevelConfigByLevel(exp, level, worldLevel uint32) (uint32, uint32, uint32) {
	for ; level < 71; level++ {
		var newExp uint32
		olConf := CONF.PlayerLevelConfigMap[level]
		newConf := CONF.PlayerLevelConfigMap[level+1]
		if newConf == nil {
			newExp = olConf.PlayerExp
		} else {
			newExp = newConf.PlayerExp - olConf.PlayerExp
		}
		if exp < newExp {
			switch worldLevel {
			case 0:
				if level >= 20 {
					return 20, exp, 0
				} else {
					return level, exp, 0
				}
			case 1:
				if level >= 30 {
					return 30, exp, 1
				} else {
					return level, exp, 1
				}
			case 2:
				if level >= 40 {
					return 40, exp, 2
				} else {
					return level, exp, 2
				}
			case 3:
				if level >= 50 {
					return 50, exp, 3
				} else {
					return level, exp, 3
				}
			case 4:
				if level >= 60 {
					return 60, exp, 4
				} else {
					return level, exp, 4
				}
			case 5:
				if level >= 65 {
					return 65, exp, 5
				} else {
					return level, exp, 5
				}
			case 6:
				if level >= 70 {
					return 70, exp, 6
				} else {
					return level, exp, 6
				}
			}
			if level >= 70 {
				return 70, exp, 6
			} else {
				return level, exp, 6
			}
		} else {
			if level == 70 {
				if exp < CONF.PlayerLevelConfigMap[70].PlayerExp {
					return 70, exp, 6
				} else {
					return 70, CONF.PlayerLevelConfigMap[70].PlayerExp, 6
				}
			}
			exp -= newExp
			level++
		}
	}
	return 70, CONF.PlayerLevelConfigMap[70].PlayerExp, 6
}
