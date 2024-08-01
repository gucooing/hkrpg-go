package gdconf

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var CONF *GameDataConfig = nil

type loadFunc func()

type GameDataConfig struct {
	loadFunc []loadFunc // 多线程读取方法
	wg       sync.WaitGroup
	goppFunc []loadFunc // 多线程读取预处理方法
	// 配置表路径前缀
	pathPrefix   string
	resPrefix    string
	excelPrefix  string
	configPrefix string
	dataPrefix   string
	// 配置表数据
	AvatarDataMap                map[uint32]*AvatarData                          // 角色
	AvatarExpItemConfigMap       map[uint32]*AvatarExpItemConfig                 // 角色升级经验材料配置
	AvatarPromotionConfigMap     map[uint32]map[uint32]*AvatarPromotionConfig    // 角色突破配置
	MultiplePathAvatarConfigMap  map[uint32]*MultiplePathAvatarConfig            // 多命途角色配置
	ExpTypeMap                   map[uint32]map[uint32]*ExpType                  // 经验配置
	EquipmentConfigMap           map[uint32]*EquipmentConfig                     // 光锥
	EquipmentExpTypeMap          map[uint32]map[uint32]*EquipmentExp             // 光锥经验配置
	EquipmentPromotionConfigMap  map[uint32]map[uint32]*EquipmentPromotionConfig // 光锥突破配置
	EquipmentSkillConfigMap      map[uint32]map[uint32]*EquipmentSkillConfig     // 光锥效果配置
	RelicMap                     map[uint32]*Relic                               // 遗器
	RelicMainAffixConfigMap      map[uint32]map[uint32]*RelicMainAffixConfig     // 圣遗物主属性配置
	RelicSubAffixConfigMap       map[uint32]map[uint32]*RelicSubAffixConfig      // 圣遗物副属性配置
	RelicExpTypeMap              map[uint32]map[uint32]*RelicExpType             // 圣遗物经验配置
	ItemConfigMap                *ItemList                                       // 材料
	ItemConfigEquipmentMap       map[uint32]*ItemConfigEquipment                 // 背包光锥配置
	ItemConfigRelicMap           map[uint32]*ItemConfigRelic                     // 背包遗器配置
	ItemUseBuffDataMap           map[uint32]*ItemUseBuffData                     // 物品增益配置表
	ItemComposeConfigMap         map[uint32]*ItemComposeConfig                   // 合成配置表
	RogueTalentMap               map[uint32]*RogueTalent                         // 模拟宇宙天赋
	RogueMapGenMap               map[uint32][]uint32                             // 模拟宇宙id场景映射表
	RogueManagerMap              map[uint32]*RogueManagerList                    // 模拟宇宙排期表
	RogueMonsterMap              map[uint32]*RogueMonster                        // 模拟宇宙怪物配置
	RogueMonsterGroupMap         map[uint32]*RogueMonsterGroups                  // 模拟宇宙怪物生成配置
	RogueBuffMap                 *RogueBuffList                                  // 模拟宇宙buff列表
	RogueAreaConfigMap           map[uint32]*RogueAreaConfig                     // 模拟宇宙关卡配置
	RogueMap                     map[uint32]*RogueMap                            // 模拟宇宙关卡地图表
	CocoonConfigMap              map[uint32]map[uint32]*CocoonConfig             // 挑战/周本
	MappingInfoMap               map[uint32]map[uint32]*MappingInfo              // 挑战/周本奖励
	AvatarSkilltreeMap           map[uint32]map[uint32]*AvatarSkilltree          // 技能库
	MazeBuffMap                  map[uint32]map[uint32]*MazeBuff                 // 技能buff库
	MazePlaneMap                 map[uint32]*MazePlane                           // 场景id
	NPCMonsterDataMap            map[uint32]*NPCMonsterData                      // NPC怪物表？
	MazePropMap                  map[uint32]*MazeProp                            // 实体列表？
	NPCDataMap                   map[uint32]*NPCData                             // NPC列表？
	GroupMap                     map[uint32]map[uint32]map[uint32]*LevelGroup    // 场景实体
	FloorMap                     map[uint32]map[uint32]*LevelFloor               // ?
	MapEntranceMap               map[uint32]*MapEntrance                         // 地图入口
	SpecialPropMap               map[uint32]*SpecialProp                         // 物品特殊状态
	BannersMap                   map[uint32]*Banners                             // 卡池信息
	ActivityPanelMap             map[uint32]*ActivityPanel                       // 活动
	ActivityLoginConfigMap       map[uint32]*ActivityLoginConfig                 // 登录活动表
	AvatarDemoConfigMap          map[uint32]*AvatarDemoConfig                    // 角色试用信息
	SpecialAvatarMap             map[uint32]map[uint32]*SpecialAvatar            // 预设角色映射表
	ActivitySchedulingMap        []*ActivityScheduling                           // 活动排期
	QuestDataMap                 map[uint32]*QuestData                           // 任务
	MonsterConfigMap             map[uint32]*MonsterConfig                       // 怪物配置
	ChallengeMazeConfigMap       map[uint32]*ChallengeMazeConfig                 // 忘却之庭配置
	ChallengeTargetConfigMap     map[uint32]*ChallengeTargetConfig               // 忘却之庭结算配置
	ChallengeStoryMazeExtraMap   map[uint32]*ChallengeStoryMazeExtra             // 忘却之庭活动积分规则
	BackGroundMusicMap           map[uint32]*BackGroundMusic                     // 背景音乐
	PlayerLevelConfigMap         map[uint32]*PlayerLevelConfig                   // 账号等级经验配置
	TextJoinConfigMap            map[uint32]*TextJoinConfig                      // 文本？
	PlaneEventMap                map[uint32]map[uint32]*PlaneEvent               // 大世界怪物信息
	StageConfigMap               map[uint32]*StageConfig                         // 具体怪物群信息
	LoadingDescMap               map[uint32]*LoadingDesc                         // 战斗随机种子
	ShopConfigMap                map[uint32][]uint32                             // 商店配置
	ShopGoodsConfigMap           map[uint32]map[uint32]*ShopGoodsConfig          // 商品配置
	RewardDataMap                map[uint32]*RewardData                          // 奖励配置
	MainMissionMap               map[uint32]*MainMission                         // 主线任务
	EventMissionMap              map[uint32]*EventMission                        // 事件任务？
	VideoVersionKey              *VideoVersionKey                                // 视频key
	InteractConfigMap            map[uint32]*InteractConfig                      // 互动配置
	MessageGroupConfig           *MessageGroupConfig                             // 消息配置
	MessageSectionConfigMap      map[uint32]*MessageSectionConfig                // 消息配置2
	TutorialDataMap              map[uint32]*TutorialData                        // 教程
	TutorialGuideGroupMap        map[uint32]*TutorialGuideGroup                  // 图鉴教程
	AdventurePlayerMap           map[uint32]*AdventurePlayer                     // 角色场景技能列表
	AvatarMazeBuffMap            map[uint32]map[uint32]*AvatarMazeBuff           // 角色场景技能效果
	RogueTournPermanentTalentMap map[uint32]*RogueTournPermanentTalent           // 灵感回路信息
	RogueTournDifficultyCompMap  map[uint32]*RogueTournDifficultyComp            // 差分宇宙难度
	RogueTournFormulaMap         map[uint32]*RogueTournFormula                   // 差分宇宙方程
	RogueTournAreaMap            map[uint32]*RogueTournArea                      // 差分宇宙关卡配置
	RogueTournExpScoreMap        map[uint32]*RogueTournExpScore                  // 差分宇宙经验获取配置
	RogueTournExpRewardMap       map[uint32]map[uint32]*RogueTournExpReward      // 差分宇宙等级奖励配置
	RogueTournRoom               *RogueTournRoom                                 // 差分宇宙地图配置
	RaidConfigMap                map[uint32]map[uint32]*RaidConfig               // Raid配置
	// 下面是预处理
	ServerGroupMap map[uint32]map[uint32]map[uint32]*GoppLevelGroup // 预处理服务器场景
	Teleports      map[uint32]map[uint32]*Teleports                 // 预处理传送锚点
	GoppMission    *GoppMission                                     // 预处理任务
	RogueRoomMap   *RogueRoomMap                                    // 模拟宇宙地图配置表
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
	g.pathPrefix = pathPrefix

	g.resPrefix = pathPrefix
	g.resPrefix += "/"

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
	g.wg.Add(len(g.loadFunc))
	for _, fn := range g.loadFunc {
		go func() {
			fn()
			g.wg.Done()
		}()
	}
	g.wg.Wait()

	g.gopp()
	g.wg.Add(len(g.goppFunc))
	for _, fn := range g.goppFunc {
		go func() {
			fn()
			g.wg.Done()
		}()
	}
	g.wg.Wait()
}

func (g *GameDataConfig) load() {
	g.loadFunc = []loadFunc{
		g.loadAvatarData,               // 角色
		g.loadAvatarExpItemConfig,      // 角色升级经验材料配置
		g.loadAvatarPromotionConfig,    // 角色突破配置
		g.loadMultiplePathAvatarConfig, // 多命途角色配置
		g.loadExpType,                  // 经验配置
		g.loadEquipmentConfig,          // 光锥
		g.loadEquipmentExpType,         // 光锥经验配置
		g.loadEquipmentPromotionConfig, // 光锥突破配置
		g.loadEquipmentSkillConfig,     // 光锥效果配置
		g.loadRelic,                    // 遗器
		g.loadRelicMainAffixConfig,     // 圣遗物主属性配置
		g.loadRelicSubAffixConfig,      // 圣遗物副属性配置
		g.loadRelicExpType,             // 圣遗物经验配置
		g.loadItemConfig,               // 材料
		g.loadItemConfigEquipment,      // 背包光锥配置
		g.loadItemConfigRelic,          // 背包遗器配置
		g.loadItemUseBuffData,          // 物品增益配置表
		g.loadItemComposeConfig,        // 合成配置表
		g.loadRogueTalent,              // 模拟宇宙天赋
		g.loadRogueMapGen,              // 模拟宇宙id场景映射表
		g.loadRogueManager,             // 模拟宇宙排期表
		g.loadRogueMonster,             // 模拟宇宙怪物配置
		g.loadRogueMonsterGroup,        // 模拟宇宙怪物生成配置
		g.loadRogueBuff,                // 模拟宇宙buff列表
		g.loadRogueAreaConfig,          // 模拟宇宙关卡配置
		g.loadRogueMap,                 // 模拟宇宙关卡地图表
		g.loadCocoonConfig,             // 挑战/周本
		g.loadMappingInfo,              // 挑战/周本奖励
		g.loadAvatarSkilltree,          // 技能库
		g.loadMazeBuff,                 // 技能buff库
		g.loadMazePlane,                // 场景id
		g.loadNPCMonsterData,           // 怪物表
		g.loadMazeProp,                 // 实体列表？
		g.loadNPCData,                  // NPC列表？
		// g.loadGroup,                     // 场景实体
		g.loadFloor,                     // 场景
		g.loadMapEntrance,               // 地图入口
		g.loadSpecialProp,               // 物品特殊状态
		g.loadBanners,                   // 卡池信息
		g.loadActivityPanel,             // 活动
		g.loadAvatarDemoConfig,          // 角色试用信息
		g.loadSpecialAvatar,             // 预设角色映射表
		g.loadActivityScheduling,        // 活动排期
		g.loadActivityLoginConfig,       // 登录活动表
		g.loadQuestData,                 // 任务
		g.loadMonsterConfig,             // 怪物配置
		g.loadChallengeMazeConfig,       // 忘却之庭配置
		g.loadChallengeTargetConfig,     // 忘却之庭结算配置
		g.loadChallengeStoryMazeExtra,   // 忘却之庭活动积分规则
		g.loadBackGroundMusic,           // 背景音乐
		g.loadPlayerLevelConfig,         // 账号等级经验配置
		g.loadTextJoinConfig,            // 文本？
		g.loadPlaneEvent,                // 大世界怪物信息
		g.loadStageConfig,               // 具体怪物群信息
		g.loadLoadingDesc,               // 战斗随机种子
		g.loadShopConfig,                // 商店配置
		g.loadShopGoodsConfig,           // 商品配置
		g.loadRewardData,                // 奖励配置
		g.loadMainMission,               // 主线任务
		g.loadEventMission,              // 事件任务？
		g.loadVideoVersionKey,           // 视频key
		g.loadInteractConfig,            // 互动配置
		g.loadMessageGroupConfig,        // 消息配置
		g.loadMessageSectionConfig,      // 消息配置2
		g.loadTutorialData,              // 教程
		g.loadTutorialGuideGroup,        // 图鉴教程
		g.loadAdventurePlayer,           // 角色场景技能列表
		g.loadAvatarMazeBuff,            // 角色场景技能效果
		g.loadRogueTournPermanentTalent, // 灵感回路信息
		g.loadRogueTournDifficultyComp,  // 差分宇宙难度
		g.loadRogueTournFormula,         // 差分宇宙方程
		g.loadRogueTournArea,            // 差分宇宙关卡配置
		g.loadRogueTournExpScore,        // 差分宇宙经验获取配置
		g.loadRogueTournExpReward,       // 差分宇宙等级奖励配置
		g.loadRogueTournRoomGen,         // 差分宇宙地图配置
		g.loadRaidConfig,                // Raid配置
	}
}

func (g *GameDataConfig) gopp() {
	g.goppFunc = []loadFunc{
		g.goppServerGroup, // 预处理服务器场景
		g.goppTeleports,   // 预处理传送锚点
		g.goppMainMission, // 预处理主线任务
		g.goppRogueRoom,   // 预处理模拟宇宙地图配置表
	}
}
