package Game

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type PlayerData struct {
	Battle                map[uint32]*Battle     // 正在进行的战斗
	BattleState           *BattleState           // 战斗情况
	BattleId              uint32                 // 战斗id
	EntityBattleId        uint32                 // 攻击实体id
	IsPaused              bool                   // 是否暂停
	GameObjectGuidCounter uint64                 // 游戏对象guid计数器
	IsNickName            bool                   // 是否修改昵称
	EntityList            map[uint32]*EntityList // 实体ID映射表
	NpcList               map[uint32]uint32
}

type EntityList struct {
	Entity  uint32 // 实体Id
	GroupId uint32 // 地图块
	Pos     *Vector
	Rot     *Vector
}

type Vector struct {
	X int32
	Y int32
	Z int32
}

func (g *Game) GetSceneNpcList() map[uint32]uint32 {
	if g.Player.NpcList == nil {
		g.Player.NpcList = make(map[uint32]uint32)
	}
	return g.Player.NpcList
}

func (g *Game) GetNextGameObjectGuid() uint64 {
	g.Player.GameObjectGuidCounter++
	return 0 + g.Player.GameObjectGuidCounter
}

func (g *Game) GetBattleIdGuid() uint32 {
	g.Player.BattleId++
	return 1 + g.Player.BattleId
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
		Battle:                  g.GetBattle(),
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
