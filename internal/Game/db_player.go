package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

type PlayerData struct {
	PlayerId   uint32             // 玩家uid
	NickName   string             // 昵称
	Level      uint32             // 玩家等级
	Exp        uint32             // 玩家经验
	WorldLevel uint32             // 世界等级
	Mcoin      uint32             // 氪金的
	Signature  string             // 签名
	HeadImage  uint32             // 头像
	Birthday   uint32             // 生日
	DbScene    *DbScene           // 场景
	Pos        *Vector            // 存档坐标
	Rot        *Vector            // 存档朝向
	DbAvatar   *DbAvatar          // 角色数据
	DbLineUp   *DbLineUp          // 队伍
	DbItem     *DbItem            // 背包
	DbGacha    *Dbgacha           // 卡池抽取情况
	Battle     map[uint32]*Battle // 正在进行的战斗
	Challenge  *Challenge         // 忘却之庭
	// 下面是在线数据
	BattleId              uint32                 `json:"-"` // 战斗id
	EntityBattleId        uint32                 `json:"-"` // 攻击实体id
	IsPaused              bool                   `json:"-"` // 是否暂停
	GameObjectGuidCounter uint64                 `json:"-"` // 游戏对象guid计数器
	IsNickName            bool                   `json:"-"` // 是否修改昵称
	EntityList            map[uint32]*EntityList `json:"-"` // 实体ID映射表
	IsBattle              bool                   `json:"-"` // 是否在战斗场景中
}

type Vector struct {
	X int
	Y int
	Z int
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
	mainAvatar := proto.HeroBasicType_BoyWarrior
	data.PlayerId = uid
	data.NickName = "hkrpg-go"
	data.Level = 20
	data.Exp = 0
	data.WorldLevel = 0
	data.Mcoin = 9999
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
	data = NewChallenge(data)

	return data
}
