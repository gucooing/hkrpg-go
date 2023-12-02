package Game

import (
	"github.com/gucooing/hkrpg-go/gdconf"
)

type PlayerData struct {
	PlayerId       uint32    // 玩家uid
	MainAvatar     uint32    // 默认主角
	NickName       string    // 昵称
	Level          uint32    // 玩家等级
	Exp            uint32    // 玩家经验
	Stamina        uint32    // 体力
	ReserveStamina uint32    // 储备体力
	WorldLevel     uint32    // 世界等级
	Signature      string    // 签名
	HeadImage      uint32    // 头像
	Birthday       []uint8   // 生日
	Pos            *Vector   // 存档坐标
	Rot            *Vector   // 存档朝向
	DbAvatar       *DbAvatar // 角色数据
	DbLineUp       *DbLineUp // 队伍
	// 下面是在线数据
	IsPaused              bool   `json:"-"` // 是否暂停
	GameObjectGuidCounter uint64 `json:"-"` // 游戏对象guid计数器
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
	data.PlayerId = uid
	data.MainAvatar = 1215
	data.NickName = "hkrpg-go"
	data.Level = 1
	data.Exp = 0
	data.Stamina = 240
	data.ReserveStamina = 2400
	data.WorldLevel = 0
	data.Signature = "hkrpg-go"
	data.HeadImage = 201217
	data.Pos = &Vector{
		X: 99,
		Y: 62,
		Z: -4800,
	}
	data.DbAvatar = new(DbAvatar)
	data.DbAvatar.Avatar = make(map[uint32]*Avatar)
	for _, a := range gdconf.GetAvatarDataMap() {
		avatarId := a.AvatarId
		if avatarId == 8002 || avatarId == 8003 || avatarId == 8004 {
			continue
		}
		data.DbAvatar.Avatar[avatarId] = AddAvatar(avatarId)
	}
	// 将主角写入队伍
	data = g.GetDbLineUp(data)
	data.DbLineUp.LineUpList[0].AvatarIdList[0] = 1215
	return data
}
