package gdconf

import (
	"fmt"
	"os"
	"strconv"

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
	g.PlayerLevelConfigMap = make(map[string]*PlayerLevelConfig)
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

func GetPlayerLevelConfigByLevel(exp, level, worldLevel uint32) (uint32, uint32) {
	level++
	for ; level < 71; level++ {
		if exp < CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].PlayerExp {
			switch worldLevel {
			case 0:
				if level >= 20 {
					return 20, exp
				} else {
					return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
				}
			case 1:
				if level >= 30 {
					return 30, exp
				} else {
					return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
				}
			case 2:
				if level >= 40 {
					return 40, exp
				} else {
					return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
				}
			case 3:
				if level >= 50 {
					return 50, exp
				} else {
					return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
				}
			case 4:
				if level >= 60 {
					return 60, exp
				} else {
					return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
				}
			case 5:
				if level >= 65 {
					return 65, exp
				} else {
					return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
				}
			case 6:
				if level >= 70 {
					return 70, exp
				} else {
					return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
				}
			}
			if level >= 70 {
				return 70, exp
			} else {
				return CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].Level - 1, exp
			}
		} else {
			if level == 70 {
				if exp < CONF.PlayerLevelConfigMap[strconv.Itoa(70)].PlayerExp {
					return 70, exp
				} else {
					return 70, CONF.PlayerLevelConfigMap[strconv.Itoa(70)].PlayerExp
				}
			}
			exp -= CONF.PlayerLevelConfigMap[strconv.Itoa(int(level))].PlayerExp
		}
	}
	return 0, 0
}
