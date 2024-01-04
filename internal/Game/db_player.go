package Game

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type PlayerData struct {
	Battle    map[uint32]*Battle // 正在进行的战斗
	Challenge *Challenge         // 忘却之庭
	// 下面是在线数据
	BattleId              uint32                 `json:"-"` // 战斗id
	EntityBattleId        uint32                 `json:"-"` // 攻击实体id
	IsPaused              bool                   `json:"-"` // 是否暂停
	GameObjectGuidCounter uint64                 `json:"-"` // 游戏对象guid计数器
	IsNickName            bool                   `json:"-"` // 是否修改昵称
	EntityList            map[uint32]*EntityList `json:"-"` // 实体ID映射表
	IsBattle              bool                   `json:"-"` // 是否在战斗场景中
}

type EntityList struct {
	Entity  uint32 // 实体Id
	GroupId uint32 // 地图块
}

func (g *Game) GetNextGameObjectGuid() uint64 {
	g.Player.GameObjectGuidCounter++
	return 0 + g.Player.GameObjectGuidCounter
}

func (g *Game) GetBattleIdGuid() uint32 {
	g.Player.BattleId++
	return 1 + g.Player.BattleId
}

// 初始化账号数据
func (g *Game) AddPalyerData(uid uint32) *PlayerData {
	data := new(PlayerData)
	data = NewChallenge(data)

	return data
}

func (g *Game) NewPlayer(uid uint32) *spb.PlayerBasicCompBin {
	g.PlayerPb = new(spb.PlayerBasicCompBin)
	g.PlayerPb = &spb.PlayerBasicCompBin{
		Uid:                     uid,
		Level:                   1,
		Exp:                     0,
		Nickname:                "hkrpg-go",
		WorldLevel:              0,
		Mcoin:                   0,
		Signature:               "签名",
		HeadImageAvatarId:       208001,
		Birthday:                0,
		Scene:                   g.GetScene(),
		Pos:                     g.GetPos(),
		Rot:                     g.GetRot(),
		Avatar:                  g.GetAvatar(),
		LineUp:                  g.GetLineUp(),
		Item:                    g.GetItem(),
		Gacha:                   g.GetGacha(),
		Battle:                  nil,
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
