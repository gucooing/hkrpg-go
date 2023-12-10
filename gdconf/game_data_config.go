package gdconf

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var CONF *GameDataConfig = nil

type GameDataConfig struct {
	// 配置表路径前缀
	excelPrefix  string
	configPrefix string
	// 配置表数据
	AvatarDataMap          map[string]*AvatarData                 // 角色
	EquipmentConfigMap     map[string]*EquipmentConfig            // 光锥
	EquipmentExpTypeMap    map[string]map[string]*EquipmentExp    // 光锥经验配置
	RelicMap               map[string]*Relic                      // 遗器
	ItemConfigEquipmentMap map[string]*ItemConfigEquipment        // 背包光锥配置
	ItemConfigRelicMap     map[string]*ItemConfigRelic            // 背包遗器配置
	RogueAreaMap           map[string]*RogueArea                  // 地图库
	AvatarSkilltreeMap     map[string]map[string]*AvatarSkilltree // 技能库
	MazePlaneMap           map[string]*MazePlane                  // 场景id
	GroupMap               map[uint32]map[uint32][]*LevelGroup    // 场景实体
	MapEntranceMap         map[string]*MapEntrance                // 地图入口
	BannersMap             map[uint32]*Banners                    // 卡池信息
	ActivityPanelMap       map[string]*ActivityPanel              // 活动
	QuestDataMap           map[string]*QuestData                  // 任务
	MonsterConfigMap       map[string]*MonsterConfig              // 怪物配置
	ChallengeMazeConfigMap map[string]*ChallengeMazeConfig        // 挑战配置
	BackGroundMusicMap     map[string]*BackGroundMusic            // 背景音乐
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
	g.loadAvatarData()          // 角色
	g.loadEquipmentConfig()     // 光锥
	g.loadEquipmentExpType()    // 光锥经验配置
	g.loadRelic()               // 遗器
	g.loadItemConfigEquipment() // 背包光锥配置
	g.loadItemConfigRelic()     // 背包遗器配置
	g.loadRogueArea()           // 副本
	g.loadAvatarSkilltree()     // 技能库
	g.loadMazePlane()           // 场景id
	g.loadGroup()               // 场景实体
	g.loadMapEntrance()         // 地图入口
	g.loadBanners()             // 卡池信息
	g.loadActivityPanel()       // 活动
	g.loadQuestData()           // 任务
	g.loadMonsterConfig()       // 怪物配置
	g.loadChallengeMazeConfig() // 挑战配置
	g.loadBackGroundMusic()     // 背景音乐
}
