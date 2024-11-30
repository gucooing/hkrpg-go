package gdconf

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

var CONF *GameDataConfig = nil
var confSync = &sync.RWMutex{}

var MaxWaitGroup int = 5

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
	AvatarPlayerIconMap          map[uint32]*AvatarPlayerIcon                    // 角色头像配置
	AvatarExpItemConfigMap       map[uint32]*AvatarExpItemConfig                 // 角色升级经验材料配置
	AvatarPromotionConfigMap     map[uint32]map[uint32]*AvatarPromotionConfig    // 角色突破配置
	MultiplePathAvatarConfigMap  map[uint32]*MultiplePathAvatarConfig            // 多命途角色配置
	ExpTypeMap                   map[uint32]map[uint32]*ExpType                  // 经验配置
	EquipmentConfigMap           map[uint32]*EquipmentConfig                     // 光锥
	EquipmentExpTypeMap          map[uint32]map[uint32]*EquipmentExp             // 光锥经验配置
	EquipmentPromotionConfigMap  map[uint32]map[uint32]*EquipmentPromotionConfig // 光锥突破配置
	EquipmentSkillConfigMap      map[uint32]map[uint32]*EquipmentSkillConfig     // 光锥效果配置
	RelicConf                    *RelicConf                                      // 遗器
	RelicMainAffixConfigMap      map[uint32]map[uint32]*RelicMainAffixConfig     // 圣遗物主属性配置
	RelicSubAffixConfigMap       map[uint32]map[uint32]*RelicSubAffixConfig      // 圣遗物副属性配置
	RelicExpTypeMap              map[uint32]map[uint32]*RelicExpType             // 圣遗物经验配置
	ItemConfigMap                *ItemList                                       // 材料
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
	RogueScoreRewardMap          map[uint32]map[uint32]*RogueScoreReward         // 模拟宇宙奖励配置
	CocoonConfigMap              map[uint32]map[uint32]*CocoonConfig             // 挑战/周本
	MappingInfoMap               map[uint32]map[uint32]*MappingInfo              // 挑战/周本奖励
	AvatarSkilltreeMap           map[uint32]map[uint32]*AvatarSkilltree          // 技能库
	MazeBuffMap                  map[uint32]map[uint32]*MazeBuff                 // 技能buff库
	MazePlaneMap                 map[uint32]*MazePlane                           // 场景id
	NPCMonsterDataMap            map[uint32]*NPCMonsterData                      // NPC怪物表？
	MazePropMap                  map[uint32]*MazeProp                            // 实体列表？
	NPCDataMap                   map[uint32]*NPCData                             // NPC列表？
	GroupMap                     map[uint32]map[uint32]map[uint32]*LevelGroup    // 场景实体
	FloorMap                     map[uint32]map[uint32]*LevelFloor               // 地图配置
	GoppFloorMap                 map[uint32]map[uint32]*GoppLevelFloor           // 地图配置预处理
	MapEntranceMap               map[uint32]*MapEntrance                         // 地图入口
	SpecialPropMap               map[uint32]*SpecialProp                         // 物品特殊状态
	BannersMap                   *BannersConf                                    // 卡池信息
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
	ShopConfigMap                *ShopInfo                                       // 商店配置
	ShopGoodsConfigMap           map[uint32]map[uint32]*ShopGoodsConfig          // 商品配置
	CityShopRewardListMap        map[uint32]map[uint32]*CityShopRewardList       // 商店等级配置
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
	StroyLineTrialAvatarDataMap  map[uint32]*StroyLineTrialAvatarData            // 故事线剧情角色配置
	StoryLineMap                 map[uint32]*StoryLine                           // 故事线配置
	StoryLineFloorDataMap        map[uint32]*StoryLineFloorData
	FarmElementConfigMap         map[uint32]*FarmElementConfig
	ItemUseDataMap               map[uint32]*ItemUseData
	DailyMissionDataMap          map[uint32]*DailyMissionData
	MonsterDropMap               map[uint32]map[uint32]*MonsterDrop
	MazeSkillMap                 map[uint32]*MazeSkill
	SummonUnitDataInfo           *SummonUnitDataInfo
	ConfigAdventureAbility       *ConfigAdventureAbility // Ability
	ChatBubbleConfigMap          map[uint32]*ChatBubbleConfig
	PhoneThemeConfigMap          map[uint32]*PhoneThemeConfig
	ContentPackageConfigMap      map[uint32]*ContentPackageConfig
	AetherDividePassiveSkillMap  map[uint32]*AetherDividePassiveSkill
	AetherDivideSpiritMap        map[uint32]*AetherDivideSpirit
	AetherDivideSpiritTrialMap   map[uint32]*AetherDivideSpiritTrial
	AetherDivideChallengeListMap map[uint32]*AetherDivideChallengeList
	ChallengeGroupConfigMap      map[uint32]*ChallengeGroupConfig
	ChallengeRewardLineMap       map[uint32]map[uint32]*ChallengeRewardLine
	FuncUnlockDataMap            map[uint32]*FuncUnlockData
	GroupSystemUnlockDataMap     map[uint32]*GroupSystemUnlockData
	TrainVisitorConfigMap        []*TrainVisitorConfig
	TrainPartyPassengerConfigMap map[uint32]*TrainPartyPassengerConfig
	TrainPartyCardConfigMap      map[uint32]*TrainPartyCardConfig
	TrainPartyAreaConfigMap      map[uint32]*TrainPartyAreaConfig
	TrainPartyStepConfigMap      map[uint32]*TrainPartyStepConfig
	TrainPartyAreaGoalConfigMap  map[uint32]*TrainPartyAreaGoalConfig
	RogueHandbookMiracleMap      map[uint32]*RogueHandbookMiracle            // 模拟宇宙奇物handbook配置
	RogueHandBookEventMap        map[uint32]*RogueHandBookEvent              // 模拟宇宙事件handbook配置
	RogueAeonStoryConfigMap      map[uint32]map[uint32]*RogueAeonStoryConfig // 模拟宇宙星神配置
	RogueBonusMap                map[uint32]*RogueBonus                      // 模拟宇宙祝福配置
	Pet                          *Pet                                        //  宠物
	MatchThreeLevelMap           map[uint32]map[uint32]*MatchThreeLevel      // 折纸小鸟关卡信息
	MatchThreeBirdMap            map[uint32]*MatchThreeBird                  // 折纸小鸟小鸟信息
	ClockParkScriptConfigMap     map[uint32]*ClockParkScriptConfig           // 美梦往事剧本信息
	// 下面是预处理
	ServerGroupMap map[uint32]map[uint32]map[uint32]*GoppLevelGroup // 预处理服务器场景
	Teleports      map[uint32]map[uint32]*Teleports                 // 预处理传送锚点
	GoppMission    *GoppMission                                     // 预处理任务
	RogueRoomMap   *RogueRoomMap                                    // 模拟宇宙地图配置表
	GoppAbility    *GoppAbility
}

func InitGameDataConfig(gameDataConfigPath string) {
	logger.Info(text.GetText(14))
	MaxWaitGroup = 5
	CONF = new(GameDataConfig)
	startTime := time.Now().Unix()
	CONF.loadAll(gameDataConfigPath)
	runtime.GC()
	endTime := time.Now().Unix()
	logger.Info(text.GetText(15), endTime-startTime)
	go confTicker(gameDataConfigPath)
}

// 安全访问方法
func getConf() *GameDataConfig {
	confSync.RLock()
	defer confSync.RUnlock()
	return CONF
}

func confTicker(gameDataConfigPath string) {
	ticker := time.NewTicker(time.Minute * 15)
	for {
		select {
		case <-ticker.C:
			func() {
				defer func() {
					runtime.GC()
					if err := recover(); err != nil {
						logger.Error("@LogTag(server_panic)@ err:%s\nstack:%s", err, logger.Stack())
						logger.Warn(text.GetText(106))
					}
				}()
				MaxWaitGroup = 1
				logger.Info(text.GetText(103))
				newConf := new(GameDataConfig)
				startTime := time.Now().Unix()
				newConf.loadAll(gameDataConfigPath)
				endTime := time.Now().Unix()
				logger.Info(text.GetText(104), endTime-startTime)

				confSync.Lock()
				CONF = newConf
				confSync.Unlock()
			}()
		}
	}
}

func (g *GameDataConfig) loadAll(gameDataConfigPath string) {
	pathPrefix := gameDataConfigPath
	dirInfo, err := os.Stat(pathPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf(text.GetText(16), err)
		panic(info)
	}
	g.pathPrefix = pathPrefix

	g.resPrefix = pathPrefix
	g.resPrefix += "/"

	g.excelPrefix = pathPrefix + "/ExcelOutput"
	dirInfo, err = os.Stat(g.excelPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf(text.GetText(16), err)
		panic(info)
	}
	g.excelPrefix += "/"

	g.configPrefix = pathPrefix + "/Config"
	dirInfo, err = os.Stat(g.configPrefix)
	if err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf(text.GetText(16), err)
		panic(info)
	}
	g.configPrefix += "/"

	g.dataPrefix = "data"
	dirInfo, err = os.Stat(g.dataPrefix)
	if err != nil || !dirInfo.IsDir() {
		logger.Error(text.GetText(16), err)
	}
	g.dataPrefix += "/"

	g.load()
	if MaxWaitGroup == 1 {
		for _, fn := range g.loadFunc {
			fn()
		}
		for _, fn := range g.goppFunc {
			fn()
		}
	} else {
		sem := make(chan struct{}, MaxWaitGroup)
		g.wg.Add(len(g.loadFunc))
		for _, fn := range g.loadFunc {
			sem <- struct{}{}
			go func() {
				fn()
				g.wg.Done()
				func() { <-sem }()
			}()
		}
		g.wg.Wait()

		g.gopp()
		g.wg.Add(len(g.goppFunc))
		for _, fn := range g.goppFunc {
			sem <- struct{}{}
			go func() {
				fn()
				g.wg.Done()
				func() { <-sem }()
			}()
		}
		g.wg.Wait()
		close(sem)
	}
}

func (g *GameDataConfig) load() {
	g.loadFunc = []loadFunc{
		g.loadAvatarData,               // 角色
		g.loadAvatarPlayerIcon,         // 角色头像配置
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
		g.loadRogueScoreReward,         // 模拟宇宙奖励配置
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
		g.loadCityShopRewardList,        // 商店等级配置
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
		g.loadStroyLineTrialAvatarData,  // 故事线剧情角色配置
		g.loadStoryLine,                 // 故事线配置
		g.loadStoryLineFloorData,
		g.loadFarmElementConfig,
		g.loadItemUseData,
		g.loadDailyMissionData,
		g.loadMonsterDrop,
		g.loadMazeSkill,
		g.loadSummonUnitData,
		g.loadChatBubbleConfig,
		g.loadPhoneThemeConfig,
		g.loadContentPackageConfig,
		g.loadAetherDividePassiveSkill,
		g.loadAetherDivideSpirit,
		g.loadAetherDivideSpiritTrial,
		g.loadAetherDivideChallengeList,
		g.loadChallengeGroupConfig,
		g.loadChallengeRewardLine,
		g.loadFuncUnlockData,
		g.loadGroupSystemUnlockData,
		g.loadTrainVisitorConfig,
		g.loadTrainPartyPassengerConfig,
		g.loadTrainPartyCardConfig,
		g.loadTrainPartyAreaConfig,
		g.loadTrainPartyStepConfig,
		g.loadTrainPartyAreaGoalConfig,
		g.loadRogueHandbookMiracle,  // 模拟宇宙奇物handbook配置
		g.loadRogueHandBookEvent,    // 模拟宇宙事件handbook配置
		g.loadRogueAeonStoryConfig,  // 模拟宇宙星神配置
		g.loadRogueBonus,            // 模拟宇宙祝福配置
		g.loadPetConfig,             // 宠物
		g.loadMatchThreeLevel,       // 折纸小鸟关卡信息
		g.loadMatchThreeBird,        // 折纸小鸟小鸟信息
		g.loadClockParkScriptConfig, // 美梦往事剧本信息
	}
}

func (g *GameDataConfig) gopp() {
	g.goppFunc = []loadFunc{
		g.goppFloor,       // 预处理服务器地图
		g.goppServerGroup, // 预处理服务器场景
		g.goppTeleports,   // 预处理传送锚点
		g.goppMainMission, // 预处理主线任务
		g.goppRogueRoom,   // 预处理模拟宇宙地图配置表
		g.goppAbility,     // Ability
		g.loadBanners,     // 卡池信息
	}
}
