package gdconf

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var CONF *GameDataConfig = nil

type GameDataConfig struct {
	// 配置表路径前缀
	excelPrefix  string
	configPrefix string
	dataPrefix   string
	// 配置表数据
	AvatarDataMap               map[string]*AvatarData                          // 角色
	AvatarExpItemConfigMap      map[string]*AvatarExpItemConfig                 // 角色升级经验材料配置
	AvatarPromotionConfigMap    map[string]map[string]*AvatarPromotionConfig    // 角色突破配置
	ExpTypeMap                  map[string]map[string]*ExpType                  // 经验配置
	EquipmentConfigMap          map[string]*EquipmentConfig                     // 光锥
	EquipmentExpTypeMap         map[string]map[string]*EquipmentExp             // 光锥经验配置
	EquipmentPromotionConfigMap map[string]map[string]*EquipmentPromotionConfig // 光锥突破配置
	RelicMap                    map[string]*Relic                               // 遗器
	RelicMainAffixConfigMap     map[uint32]map[uint32]*RelicMainAffixConfig     // 圣遗物主属性配置
	RelicSubAffixConfigMap      map[uint32]map[uint32]*RelicSubAffixConfig      // 圣遗物副属性配置
	RelicExpTypeMap             map[string]map[string]*RelicExpType             // 圣遗物经验配置
	ItemConfigMap               *ItemList                                       // 材料
	ItemConfigEquipmentMap      map[string]*ItemConfigEquipment                 // 背包光锥配置
	ItemConfigRelicMap          map[string]*ItemConfigRelic                     // 背包遗器配置
	RogueTalentMap              map[string]*RogueTalent                         // 模拟宇宙天赋
	RogueMapGenMap              map[string][]uint32                             // 模拟宇宙id场景映射表
	RogueManagerMap             *RogueManager                                   // 模拟宇宙排期表
	RogueMonsterMap             map[uint32]*RogueMonster                        // 模拟宇宙怪物配置
	RogueMonsterGroupMap        map[uint32]*RogueMonsterGroup                   // 模拟宇宙怪物生成配置
	RogueBuffMap                *RogueBuffList                                  // 模拟宇宙buff列表
	RogueAreaConfigMap          map[string]*RogueAreaConfig                     // 模拟宇宙关卡配置
	RogueMap                    map[uint32]*RogueMap                            // 模拟宇宙关卡地图表
	RogueRoomMap                map[uint32]*RogueRoom                           // 模拟宇宙地图配置表
	CocoonConfigMap             map[string]map[string]*CocoonConfig             // 挑战/周本
	MappingInfoMap              map[string]map[string]*MappingInfo              // 挑战/周本奖励
	AvatarSkilltreeMap          map[string]map[string]*AvatarSkilltree          // 技能库
	MazeBuffMap                 map[string]map[string]*MazeBuff                 // 技能buff库
	MazePlaneMap                map[string]*MazePlane                           // 场景id
	NPCMonsterDataMap           map[string]*NPCMonsterData                      // NPC怪物表？
	MazePropMap                 map[uint32]*MazeProp                            // 实体列表？
	NPCDataMap                  map[string]*NPCData                             // NPC列表？
	GroupMap                    map[uint32]map[uint32]map[uint32]*LevelGroup    // 场景实体
	FloorMap                    map[uint32]map[uint32]*LevelFloor               // ?
	MapEntranceMap              map[string]*MapEntrance                         // 地图入口
	BannersMap                  map[uint32]*Banners                             // 卡池信息
	ActivityPanelMap            map[string]*ActivityPanel                       // 活动
	ActivityLoginConfigMap      map[string]*ActivityLoginConfig                 // 登录活动表
	AvatarDemoConfigMap         map[string]*AvatarDemoConfig                    // 角色试用信息
	SpecialAvatarMap            map[string]map[string]*SpecialAvatar            // 预设角色映射表
	ActivitySchedulingMap       []*ActivityScheduling                           // 活动排期
	QuestDataMap                map[string]*QuestData                           // 任务
	MonsterConfigMap            map[string]*MonsterConfig                       // 怪物配置
	ChallengeMazeConfigMap      map[uint32]*ChallengeMazeConfig                 // 忘却之庭配置
	ChallengeTargetConfigMap    map[string]*ChallengeTargetConfig               // 忘却之庭结算配置
	ChallengeStoryMazeExtraMap  map[string]*ChallengeStoryMazeExtra             // 忘却之庭活动积分规则
	BackGroundMusicMap          map[string]*BackGroundMusic                     // 背景音乐
	PlayerLevelConfigMap        map[string]*PlayerLevelConfig                   // 账号等级经验配置
	TextJoinConfigMap           map[string]*TextJoinConfig                      // 文本？
	PlaneEventMap               map[string]map[string]*PlaneEvent               // 大世界怪物信息
	StageConfigMap              map[string]*StageConfig                         // 具体怪物群信息
	LoadingDescMap              map[string]*LoadingDesc                         // 战斗随机种子
	ShopConfigMap               map[uint32][]uint32                             // 商店配置
	ShopGoodsConfigMap          map[uint32]map[uint32]*ShopGoodsConfig          // 商品配置
	RewardDataMap               map[string]*RewardData                          // 奖励配置
}

func InitGameDataConfig(gameDataConfigPath string) {
	logger.Info("读取资源文件")
	CONF = new(GameDataConfig)
	startTime := time.Now().Unix()
	CONF.loadAll(gameDataConfigPath)
	runtime.GC()
	endTime := time.Now().Unix()
	logger.Info("load all game data config finish, cost: %v(s)", endTime-startTime)
}

func (g *GameDataConfig) loadAll(gameDataConfigPath string) {
	pathPrefix := gameDataConfigPath
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

	g.dataPrefix = "data"
	dirInfo, err = os.Stat(g.dataPrefix)
	if err != nil || !dirInfo.IsDir() {
		logger.Error("open game data config data dir error: %v", err)
	}
	g.dataPrefix += "/"

	g.load()
}

func (g *GameDataConfig) load() {
	g.loadAvatarData()               // 角色
	g.loadAvatarExpItemConfig()      // 角色升级经验材料配置
	g.loadAvatarPromotionConfig()    // 角色突破配置
	g.loadExpType()                  // 经验配置
	g.loadEquipmentConfig()          // 光锥
	g.loadEquipmentExpType()         // 光锥经验配置
	g.loadEquipmentPromotionConfig() // 光锥突破配置
	g.loadRelic()                    // 遗器
	g.loadRelicMainAffixConfig()     // 圣遗物主属性配置
	g.loadRelicSubAffixConfig()      // 圣遗物副属性配置
	g.loadRelicExpType()             // 圣遗物经验配置
	g.loadItemConfig()               // 材料
	g.loadItemConfigEquipment()      // 背包光锥配置
	g.loadItemConfigRelic()          // 背包遗器配置
	g.loadRogueTalent()              // 模拟宇宙天赋
	g.loadRogueMapGen()              // 模拟宇宙id场景映射表
	g.loadRogueManager()             // 模拟宇宙排期表
	g.loadRogueMonster()             // 模拟宇宙怪物配置
	g.loadRogueMonsterGroup()        // 模拟宇宙怪物生成配置
	g.loadRogueBuff()                // 模拟宇宙buff列表
	g.loadRogueAreaConfig()          // 模拟宇宙关卡配置
	g.loadRogueMap()                 // 模拟宇宙关卡地图表
	g.loadRogueRoom()                // 模拟宇宙地图配置表
	g.loadCocoonConfig()             // 挑战/周本
	g.loadMappingInfo()              // 挑战/周本奖励
	g.loadAvatarSkilltree()          // 技能库
	g.loadMazeBuff()                 // 技能buff库
	g.loadMazePlane()                // 场景id
	g.loadNPCMonsterData()           // NPC怪物表？
	g.loadMazeProp()                 // 实体列表？
	g.loadNPCData()                  // NPC列表？
	g.loadGroup()                    // 场景实体
	g.loadFloor()                    // ?
	g.loadMapEntrance()              // 地图入口
	g.loadBanners()                  // 卡池信息
	g.loadActivityPanel()            // 活动
	g.loadAvatarDemoConfig()         // 角色试用信息
	g.loadSpecialAvatar()            // 预设角色映射表
	g.loadActivityScheduling()       // 活动排期
	g.loadActivityLoginConfig()      // 登录活动表
	g.loadQuestData()                // 任务
	g.loadMonsterConfig()            // 怪物配置
	g.loadChallengeMazeConfig()      // 忘却之庭配置
	g.loadChallengeTargetConfig()    // 忘却之庭结算配置
	g.loadChallengeStoryMazeExtra()  // 忘却之庭活动积分规则
	g.loadBackGroundMusic()          // 背景音乐
	g.loadPlayerLevelConfig()        // 账号等级经验配置
	g.loadTextJoinConfig()           // 文本？
	g.loadPlaneEvent()               // 大世界怪物信息
	g.loadStageConfig()              // 具体怪物群信息
	g.loadLoadingDesc()              // 战斗随机种子
	g.loadShopConfig()               // 商店配置
	g.loadShopGoodsConfig()          // 商品配置
	g.loadRewardData()               // 奖励配置
}
