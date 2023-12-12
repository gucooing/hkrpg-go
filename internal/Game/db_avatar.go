package Game

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

type DbAvatar struct {
	Avatar     map[uint32]*Avatar
	MainAvatar proto.HeroBasicType // 默认主角
}

type Avatar struct {
	AvatarId          uint32            // 角色id
	Exp               uint32            // 经验
	Level             uint32            // 等级
	FirstMetTimestamp uint64            // 获得时间戳
	Promotion         uint32            // 突破等阶
	Rank              uint32            // 命座
	Hp                uint32            // 血量
	SkilltreeList     map[uint32]uint32 `json:"-"` // 技能等级数据
	EquipmentUniqueId uint32            // 装备光锥
}

func NewAvatar(data *PlayerData, mainAvatar proto.HeroBasicType) *PlayerData {
	data.DbAvatar = new(DbAvatar)
	data.DbAvatar.MainAvatar = mainAvatar
	data.DbAvatar.Avatar = make(map[uint32]*Avatar)

	return data
}

func (g *Game) AddAvatar(avatarId uint32) {
	avatar := new(Avatar)
	// TODO
	avatar.AvatarId = avatarId
	avatar.Exp = 0
	avatar.Level = 1
	avatar.FirstMetTimestamp = uint64(time.Now().Unix())
	avatar.Promotion = 0
	avatar.Rank = 0
	avatar.Hp = 10000
	avatar.EquipmentUniqueId = 0

	g.Player.DbAvatar.Avatar[avatarId] = avatar
	g.AvatarPlayerSyncScNotify(avatarId)
	g.ScenePlaneEventScNotify(avatarId, 1)
}
func GetKilltreeList(avatarId, level uint32) []*proto.AvatarSkillTree {
	skilltreeList := gdconf.GetAvatarSkilltreeById(avatarId, level)
	return skilltreeList
}

func (g *Game) AvatarPlayerSyncScNotify(avatarId uint32) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
	}
	avatardb := g.Player.DbAvatar.Avatar[avatarId]
	avatar := &proto.Avatar{
		SkilltreeList:     GetKilltreeList(avatarId, 1),
		Exp:               avatardb.Exp,
		BaseAvatarId:      avatarId,
		Rank:              avatardb.Rank,
		EquipmentUniqueId: avatardb.EquipmentUniqueId,
		EquipRelicList:    make([]*proto.EquipRelic, 0),
		TakenRewards:      make([]uint32, 0),
		FirstMetTimestamp: avatardb.FirstMetTimestamp,
		Promotion:         avatardb.Promotion,
		Level:             avatardb.Level,
	}
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	g.send(cmd.PlayerSyncScNotify, notify)

	g.UpDataPlayer()
}
