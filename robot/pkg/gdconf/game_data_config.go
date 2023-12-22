package gdconf

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/robot/pkg/config"
)

var CONF *GameDataConfig = nil

type GameDataConfig struct {
	// 配置表路径前缀
	excelPrefix  string
	configPrefix string
	// 配置表数据
	MapEntranceMap map[string]*MapEntrance // 地图入口
}

func InitGameDataConfig() {
	CONF = new(GameDataConfig)
	startTime := time.Now().Unix()
	CONF.loadAll()
	runtime.GC()
	endTime := time.Now().Unix()
	logger.Info("load all game data config finish, cost: %v(s)", endTime-startTime)
}

func (g *GameDataConfig) loadAll() {
	pathPrefix := config.GetConfig().GameDataConfigPath

	dirInfo, err := os.Stat(pathPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("open game data config dir error: %v", err)
		panic(info)
	}

	g.excelPrefix = pathPrefix + "/ExcelOutput"
	dirInfo, err = os.Stat(g.excelPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("open game data config ExcelOutput dir error: %v", err)
		panic(info)
	}
	g.excelPrefix += "/"

	g.configPrefix = pathPrefix + "/Config"
	dirInfo, err = os.Stat(g.configPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("open game data config Config dir error: %v", err)
		panic(info)
	}
	g.configPrefix += "/"

	g.load()
}

func (g *GameDataConfig) load() {
	g.loadMapEntrance() // 地图入口
}
