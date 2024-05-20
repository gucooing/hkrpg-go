package player

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type OnlineData struct {
	LoginToday            bool                 // 是否是今天第一次登录
	Battle                map[uint32]*Battle   // 正在进行的战斗
	BattleState           *BattleState         // 战斗情况
	BattleId              uint32               // 战斗id
	EntityBattleId        uint32               // 攻击实体id
	IsPaused              bool                 // 是否暂停
	GameObjectGuidCounter uint32               // 游戏对象guid计数器 貌似一个玩家用一个就行了
	IsNickName            bool                 // 是否修改昵称
	EntityMap             map[uint32]EntityAll // 场景实体
	CurBattle             *CurBattle           // 正在进行的战斗
}

type Vector struct {
	X int32
	Y int32
	Z int32
}

func (g *GamePlayer) NewBasicBin() *spb.PlayerBasicCompBin {
	g.BasicBin = new(spb.PlayerBasicCompBin)
	g.BasicBin = &spb.PlayerBasicCompBin{
		Level:                   1,
		Nickname:                "hkrpg-go",
		WorldLevel:              0,
		Activity:                g.NewActivity(),
		Signature:               "",
		HeadImageAvatarId:       208001,
		Birthday:                0,
		Scene:                   g.NewScene(),
		Pos:                     g.NewPos(),
		Rot:                     g.NewRot(),
		Avatar:                  g.NewAvatar(),
		LineUp:                  g.NewLineUp(),
		Item:                    g.NewItem(),
		Gacha:                   g.NewGacha(),
		Battle:                  g.NewBattle(),
		RewardTakenLevelList:    nil,
		OpenStateMap:            nil,
		RegisterTime:            0,
		TotalLoginDays:          0,
		TotalGameTime:           0,
		LastLoginTime:           0,
		LastLoginPlatform:       0,
		LastLogoutTime:          0,
		IsRebateMailSent:        false,
		IsRebateMailReceived:    false,
		DataVersion:             0,
		LastDailyRefreshTime:    0,
		ProfilePictureCostumeId: 0,
		PsnId:                   "",
		LanguageType:            0,
		ClientAppVersion:        "",
		ClientDeviceInfo:        "",
		ClientSystemVersion:     "",
		SetLanguageTag:          0,
		GuidSeqId:               0,
		IsGuest:                 false,
		PivotClientTime:         0,
		PivotUnixTime:           0,
		PlayerStatId:            0,
		NicknameAuditBin:        nil,
		IpCountryCode:           "",
		IpRegionName:            "",
	}

	g.AddAvatar(uint32(g.GetAvatar().CurMainAvatar))

	return g.BasicBin
}

func (g *GamePlayer) GetBasicBin() *spb.PlayerBasicCompBin {
	if g.BasicBin == nil {
		g.BasicBin = g.NewBasicBin()
	}
	return g.BasicBin
}

func (g *GamePlayer) GetOnlineData() *OnlineData {
	if g.OnlineData == nil {
		g.OnlineData = &OnlineData{
			LoginToday:            false,
			Battle:                nil,
			BattleState:           nil,
			BattleId:              0,
			EntityBattleId:        0,
			IsPaused:              false,
			GameObjectGuidCounter: 0,
			IsNickName:            false,
			EntityMap:             g.NewEntity(),
			CurBattle:             g.NewCurBattle(),
		}
	}

	return g.OnlineData
}

func (g *GamePlayer) GetNextGameObjectGuid() uint32 {
	g.OnlineData.GameObjectGuidCounter++
	return 0 + g.OnlineData.GameObjectGuidCounter
}

func (g *GamePlayer) GetBattleIdGuid() uint32 {
	g.OnlineData.BattleId++
	return 1 + g.OnlineData.BattleId
}

func (g *GamePlayer) GetNickname() string {
	db := g.GetBasicBin()
	if db.Nickname == "" {
		db.Nickname = "hkrpg-go"
	}
	return db.Nickname
}

func (g *GamePlayer) GetLevel() uint32 {
	db := g.GetBasicBin()
	if db.Level <= 0 {
		db.Level = 1
	}
	return db.Level
}

func (g *GamePlayer) GetWorldLevel() uint32 {
	db := g.GetBasicBin()
	if db.WorldLevel < 0 {
		db.WorldLevel = 0
	}
	return db.WorldLevel
}

func (g *GamePlayer) SetWorldLevel(worldLevel uint32) {
	if worldLevel < 0 || worldLevel > 6 {
		return
	}
	db := g.GetBasicBin()
	db.WorldLevel = worldLevel
}

func (g *GamePlayer) GetHeadIcon() uint32 {
	db := g.GetBasicBin()
	if db.HeadImageAvatarId == 0 {
		db.HeadImageAvatarId = 208001
	}
	return db.HeadImageAvatarId
}

func (g *GamePlayer) GetDataVersion() uint32 {
	db := g.GetBasicBin()
	return db.DataVersion
}

func (g *GamePlayer) AddDataVersion() uint32 {
	db := g.GetBasicBin()
	db.DataVersion++
	return db.DataVersion
}
