package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

type PlayerData struct {
	PlayerId       uint32    // 玩家uid
	NickName       string    // 昵称
	Level          uint32    // 玩家等级
	Exp            uint32    // 玩家经验
	Stamina        uint32    // 体力
	ReserveStamina uint32    // 储备体力
	WorldLevel     uint32    // 世界等级
	Signature      string    // 签名
	HeadImage      uint32    // 头像
	Birthday       uint32    // 生日
	DbScene        *DbScene  // 场景
	Pos            *Vector   // 存档坐标
	Rot            *Vector   // 存档朝向
	DbAvatar       *DbAvatar // 角色数据
	DbLineUp       *DbLineUp // 队伍
	DbItem         *DbItem   // 背包
	DbGacha        *Dbgacha  // 卡池抽取情况
	// 下面是在线数据
	IsPaused              bool   `json:"-"` // 是否暂停
	GameObjectGuidCounter uint64 `json:"-"` // 游戏对象guid计数器
	IsNickName            bool   `json:"-"` // 是否修改昵称
}

type Vector struct {
	X int
	Y int
	Z int
}

func (g *Game) GetNextGameObjectGuid() uint64 {
	g.Player.GameObjectGuidCounter++
	return 0 + g.Player.GameObjectGuidCounter
}

// 初始化账号数据
func (g *Game) AddPalyerData(uid uint32) *PlayerData {
	data := new(PlayerData)
	mainAvatar := proto.HeroBasicType_BoyWarrior
	data.PlayerId = uid
	data.NickName = "hkrpg-go"
	data.Level = 1
	data.Exp = 0
	data.Stamina = 240
	data.ReserveStamina = 2400
	data.WorldLevel = 0
	data.Signature = "hkrpg-go"
	data.HeadImage = 208001
	data = NewScene(data)
	data.Pos = &Vector{
		X: -43300,
		Y: 6,
		Z: -37960,
	}
	data.Rot = &Vector{
		X: 0,
		Y: 90000,
		Z: 0,
	}
	data = NewAvatar(data, mainAvatar)
	// 将主角写入队伍
	data = NewDbLineUp(data)
	data = NewItem(data)
	data = NewGaCha(data)

	return data
}
