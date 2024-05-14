package player

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type PlayerData struct {
	LoginToday            bool                 // 是否是今天第一次登录
	Battle                map[uint32]*Battle   // 正在进行的战斗
	BattleState           *BattleState         // 战斗情况
	BattleId              uint32               // 战斗id
	EntityBattleId        uint32               // 攻击实体id
	IsPaused              bool                 // 是否暂停
	GameObjectGuidCounter uint32               // 游戏对象guid计数器 貌似一个玩家用一个就行了
	IsNickName            bool                 // 是否修改昵称
	EntityMap             map[uint32]EntityAll // 场景实体
	NpcList               map[uint32]uint32
}

type Vector struct {
	X int32
	Y int32
	Z int32
}

func (g *GamePlayer) NewPlayer() *spb.PlayerBasicCompBin {
	g.PlayerPb = new(spb.PlayerBasicCompBin)
	g.PlayerPb = &spb.PlayerBasicCompBin{
		Level:                   1,
		Exp:                     0,
		Nickname:                "hkrpg-go",
		WorldLevel:              0,
		Activity:                g.GetActivity(),
		Signature:               "",
		HeadImageAvatarId:       208001,
		Birthday:                0,
		Scene:                   g.NewScene(),
		Pos:                     g.NewPos(),
		Rot:                     g.NewRot(),
		Avatar:                  g.GetAvatar(),
		LineUp:                  g.NewLineUp(),
		Item:                    g.GetItem(),
		Gacha:                   g.GetGacha(),
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

	return g.PlayerPb
}

func (g *GamePlayer) GetPlayerPb() *spb.PlayerBasicCompBin {
	if g.PlayerPb == nil {
		g.PlayerPb = g.NewPlayer()
	}
	return g.PlayerPb
}

func (g *GamePlayer) GetPlayer() *PlayerData {
	if g.Player == nil {
		g.Player = &PlayerData{
			LoginToday:            false,
			Battle:                nil,
			BattleState:           nil,
			BattleId:              0,
			EntityBattleId:        0,
			IsPaused:              false,
			GameObjectGuidCounter: 0,
			IsNickName:            false,
			EntityMap:             g.NewEntity(),
			NpcList:               nil,
		}
	}

	return g.Player
}

func (g *GamePlayer) GetSceneNpcList() map[uint32]uint32 {
	if g.Player.NpcList == nil {
		g.Player.NpcList = make(map[uint32]uint32)
	}
	return g.Player.NpcList
}

func (g *GamePlayer) GetNextGameObjectGuid() uint32 {
	g.Player.GameObjectGuidCounter++
	return 0 + g.Player.GameObjectGuidCounter
}

func (g *GamePlayer) GetBattleIdGuid() uint32 {
	g.Player.BattleId++
	return 1 + g.Player.BattleId
}

func (g *GamePlayer) GetNickname() string {
	db := g.GetPlayerPb()
	if db.Nickname == "" {
		db.Nickname = "hkrpg-go"
	}
	return db.Nickname
}

func (g *GamePlayer) GetLevel() uint32 {
	db := g.GetPlayerPb()
	if db.Level <= 0 {
		db.Level = 1
	}
	return db.Level
}

func (g *GamePlayer) GetWorldLevel() uint32 {
	db := g.GetPlayerPb()
	if db.WorldLevel < 0 {
		db.WorldLevel = 0
	}
	return db.WorldLevel
}

func (g *GamePlayer) GetHeadIcon() uint32 {
	db := g.GetPlayerPb()
	if db.HeadImageAvatarId == 0 {
		db.HeadImageAvatarId = 208001
	}
	return db.HeadImageAvatarId
}

func (g *GamePlayer) GetDataVersion() uint32 {
	db := g.GetPlayerPb()
	return db.DataVersion
}

func (g *GamePlayer) AddDataVersion() uint32 {
	db := g.GetPlayerPb()
	db.DataVersion++
	return db.DataVersion
}
