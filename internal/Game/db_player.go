package Game

import (
	"github.com/gucooing/hkrpg-go/gdconf"
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
	Pos            *Vector   // 存档坐标
	Rot            *Vector   // 存档朝向
	DbAvatar       *DbAvatar // 角色数据
	DbLineUp       *DbLineUp // 队伍
	Dbgacha        *Dbgacha  // 卡池抽取情况
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
	data.Pos = &Vector{
		X: -47,
		Y: 146,
		Z: 7269,
	}
	data.Rot = &Vector{
		X: 0,
		Y: 0,
		Z: 0,
	}
	data.DbAvatar = new(DbAvatar)
	data.DbAvatar.MainAvatar = mainAvatar
	data.DbAvatar.Avatar = make(map[uint32]*Avatar)
	// TODO 直接给全部角色(包括多个主角，如果出现了问题，那只给一个当前属性主角） *不知道你是不是下一个把四个主角添加到一个队伍的yz
	for _, a := range gdconf.GetAvatarDataMap() {
		avatarId := a.AvatarId
		data.DbAvatar.Avatar[avatarId] = AddAvatar(avatarId)
	}
	// 将主角写入队伍
	data = g.GetDbLineUp(data)
	data.DbLineUp.LineUpList[0].AvatarIdList[0] = uint32(mainAvatar)
	return data
}
