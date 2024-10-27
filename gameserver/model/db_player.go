package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

type PlayerData struct {
	OnlineData *OnlineData             // 玩家在线数据
	BasicBin   *spb.PlayerBasicCompBin // 玩家pb数据
}

type OnlineData struct {
	LoginToday            bool                     // 是否是今天第一次登录
	BattleId              uint32                   // 战斗id
	IsPaused              bool                     // 是否暂停
	GameObjectGuidCounter uint32                   // 游戏对象guid计数器 貌似一个玩家用一个就行了
	IsNickName            bool                     // 是否修改昵称
	SceneMap              *SceneMap                // 在线场景管理
	BlockMap              map[uint32]*spb.BlockBin // 缓存场景
	CurBattle             *CurBattle               // 正在进行的战斗
	Mission               *MissionInfo             // 任务预处理
}

func NewPlayerData() *PlayerData {
	g := new(PlayerData)
	g.BasicBin = newBasicBin()
	// 添加默认数据
	g.AddAvatar(1001)
	g.AddAvatar(8001)
	g.NewTrialLine([]uint32{1001005, 0, 0, 0, 1001005})
	return g
}

func newBasicBin() *spb.PlayerBasicCompBin {
	basicBin := &spb.PlayerBasicCompBin{
		Level:                1,
		Nickname:             "hkrpg-go",
		WorldLevel:           0,
		Activity:             NewActivity(),
		PojokNostalgia:       newPojokNostalgia(),
		Signature:            "",
		HeadImageAvatarId:    0,
		Birthday:             0,
		Scene:                NewScene(),
		ChangeStory:          NewChangeStory(),
		Avatar:               NewAvatar(),
		LineUp:               NewLineUp(),
		Item:                 NewItem(),
		Gacha:                NewGacha(),
		Battle:               NewBattle(),
		RewardTakenLevelList: make([]uint32, 0),
		RegisterTime:         0,
		TotalLoginDays:       0,
		TotalGameTime:        0,
		LastLoginTime:        0,
		LastStaminaTime:      0,
		LastLogoutTime:       0,
		Mail:                 NewMail(),
		Friend:               NewFriend(),
		Mission:              newMission(),
		Day:                  NewDays(),
		DataVersion:          0,
		LastDailyRefreshTime: 0,
		Tutorial:             NewTutorialDb(),
		IsProficientPlayer:   false,
		LanguageType:         0,
		ClientAppVersion:     "",
		ClientDeviceInfo:     "",
		ClientSystemVersion:  "",
		SetLanguageTag:       0,
		GuidSeqId:            0,
		MessageGroupList:     NewMessageGroup(),
		PlayerStatId:         0,
		NicknameAuditBin:     nil,
		IpCountryCode:        "",
		IpRegionName:         "",
		IsJumpMission:        false,
	}

	return basicBin
}

func (g *PlayerData) GetIsProficientPlayer() bool {
	db := g.GetBasicBin()
	return db.IsProficientPlayer
}

func (g *PlayerData) GetBasicBin() *spb.PlayerBasicCompBin {
	if g.BasicBin == nil {
		g.BasicBin = newBasicBin()
	}
	return g.BasicBin
}

func (g *PlayerData) GetOnlineData() *OnlineData {
	if g.OnlineData == nil {
		g.OnlineData = &OnlineData{
			LoginToday:            false,
			IsPaused:              false,
			GameObjectGuidCounter: 0,
			IsNickName:            false,
			SceneMap:              NewSceneMap(),
			BlockMap:              NewBlockMap(),
			CurBattle:             g.NewCurBattle(),
			BattleId:              10000,
			Mission:               nil,
		}
	}

	return g.OnlineData
}

func (g *PlayerData) GetNextGameObjectGuid() uint32 {
	db := g.GetOnlineData()
	db.GameObjectGuidCounter++
	return 0 + db.GameObjectGuidCounter
}

func (g *PlayerData) GetBattleIdGuid() uint32 {
	db := g.GetOnlineData()
	if db.BattleId <= 0 {
		db.BattleId = 10000
	}
	defer g.AddBattleIdGuid()
	return db.BattleId
}

func (g *PlayerData) AddBattleIdGuid() {
	db := g.GetOnlineData()
	db.BattleId++
}

func (g *PlayerData) GetNickname() string {
	db := g.GetBasicBin()
	if db.Nickname == "" {
		db.Nickname = "hkrpg-go"
	}
	return db.Nickname
}

func (g *PlayerData) SetNickname(name string) {
	db := g.GetBasicBin()
	db.Nickname = name
}

func (g *PlayerData) GetLevel() uint32 {
	db := g.GetBasicBin()
	if db.Level <= 0 {
		db.Level = 1
	}
	return db.Level
}

func (g *PlayerData) GetWorldLevel() uint32 {
	db := g.GetBasicBin()
	if db.WorldLevel < 0 {
		db.WorldLevel = 0
	}
	return db.WorldLevel
}

func (g *PlayerData) GetIsJumpMission() bool {
	return g.GetBasicBin().IsJumpMission
}

func (g *PlayerData) AddWorldLevel(num uint32) {
	g.SetWorldLevel(g.GetWorldLevel() + num)
}

func (g *PlayerData) SetWorldLevel(worldLevel uint32) {
	if worldLevel < 0 || worldLevel > 6 {
		return
	}
	db := g.GetBasicBin()
	db.WorldLevel = worldLevel
}

func (g *PlayerData) GetHeadIcon() uint32 {
	db := g.GetBasicBin()
	if db.HeadImageAvatarId == 0 {
		db.HeadImageAvatarId = 0
	}
	return db.HeadImageAvatarId
}

func (g *PlayerData) GetDataVersion() uint32 {
	db := g.GetBasicBin()
	return db.DataVersion
}

func (g *PlayerData) GetSignature() string {
	db := g.GetBasicBin()
	return db.Signature
}

func NewTutorialDb() *spb.TutorialDb {
	return new(spb.TutorialDb)
}

func (g *PlayerData) GetTutorialDb() *spb.TutorialDb {
	db := g.GetBasicBin()
	if db.Tutorial == nil {
		db.Tutorial = NewTutorialDb()
	}
	return db.Tutorial
}

func (g *PlayerData) GetTutorial() map[uint32]*spb.TutorialInfo {
	db := g.GetTutorialDb()
	if db.Tutorial == nil {
		db.Tutorial = make(map[uint32]*spb.TutorialInfo)
	}
	return db.Tutorial
}

func (g *PlayerData) GetTutorialGuide() map[uint32]*spb.TutorialInfo {
	db := g.GetTutorialDb()
	if db.TutorialGuide == nil {
		db.TutorialGuide = make(map[uint32]*spb.TutorialInfo)
	}
	return db.TutorialGuide
}

func (g *PlayerData) UnlockTutorial(id uint32) {
	db := g.GetTutorial()
	db[id] = &spb.TutorialInfo{
		Id:     id,
		Status: spb.TutorialStatus_TUTORIAL_UNLOCK,
	}
}

func (g *PlayerData) FinishTutorial(id uint32) {
	db := g.GetTutorial()
	if db[id] == nil {
		return
	}
	db[id].Status = spb.TutorialStatus_TUTORIAL_FINISH
}

func (g *PlayerData) UnlockTutorialGuide(id uint32) {
	db := g.GetTutorialGuide()
	db[id] = &spb.TutorialInfo{
		Id:     id,
		Status: spb.TutorialStatus_TUTORIAL_UNLOCK,
	}
}

func (g *PlayerData) FinishTutorialGuide(id uint32, addItem *AddItem) {
	db := g.GetTutorialGuide()
	if db[id] != nil {
		addItem = NewAddItem(addItem)
		db[id].Status = spb.TutorialStatus_TUTORIAL_FINISH
		conf := gdconf.GetTutorialGuideGroup(id)
		pile := GetRewardData(conf.RewardID)
		addItem.PileItem = append(addItem.PileItem, pile...)
		g.AddItem(addItem)
	}
}

func (g *PlayerData) GetRewardTakenLevelList() []uint32 {
	db := g.GetBasicBin()
	if db.RewardTakenLevelList == nil {
		db.RewardTakenLevelList = make([]uint32, 0)
	}
	return db.RewardTakenLevelList
}

func (g *PlayerData) AddRewardTakenLevelList(id uint32) {
	db := g.GetRewardTakenLevelList()
	isAdd := true
	for _, level := range db {
		if level == id {
			isAdd = false
			break
		}
	}
	if isAdd {
		db = append(db, id)
	}
}
