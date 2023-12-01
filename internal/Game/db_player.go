package Game

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
}

type Vector struct {
	X int
	Y int
	Z int
}

// 初始化账号数据
func (g *Game) AddPalyerData(uid uint32) *PlayerData {
	data := new(PlayerData)
	data.PlayerId = uid
	data.MainAvatar = 8001
	data.NickName = "hkrpg-go"
	data.Level = 1
	data.Exp = 0
	data.Stamina = 240
	data.ReserveStamina = 4
	data.WorldLevel = 0
	data.Signature = "hkrpg-go"
	data.HeadImage = 208001
	data.Pos = &Vector{
		X: 99,
		Y: 62,
		Z: -4800,
	}
	data.DbAvatar = new(DbAvatar)
	data.DbAvatar.Avatar = make(map[uint32]*Avatar)
	avatarIdList := []uint32{8001} // 设置初始化时给予多少角色
	for _, a := range avatarIdList {
		data.DbAvatar.Avatar[a] = AddAvatar(a)
	}
	// 将主角写入队伍
	data = g.GetDbLineUp(data)
	data.DbLineUp.LineUpList[0].AvatarIdList[0] = 8001
	return data
}
