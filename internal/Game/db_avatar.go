package Game

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

type DbAvatar struct {
	Avatar         map[uint32]*Avatar
	MainAvatar     proto.HeroBasicType // 默认主角
	MainAvatarList []uint32            // 主角们
}

type Avatar struct {
	AvatarId          uint32           // 角色id
	Exp               uint32           // 经验
	Level             uint32           // 等级
	Type              proto.AvatarType // 角色状态
	FirstMetTimestamp uint64           // 获得时间戳
	Promotion         uint32           // 突破等阶
	Rank              uint32           // 命座
	Hp                uint32           // 血量
	SpBar             *SpBarInfo       // 能量
	SkilltreeList     []*Skilltree     // 技能等级数据 [id]level
	EquipmentUniqueId uint32           // 装备光锥
	TakenRewards      []uint32         // 已领取的突破奖励
	BuffList          uint32           // 开启战斗的buff
}

type Skilltree struct {
	PointId uint32
	Level   uint32
}

type SpBarInfo struct {
	CurSp uint32
	MaxSp uint32
}

func (g *Game) GetAvatarHp(avatarId uint32) uint32 {
	level := g.Player.DbAvatar.Avatar[avatarId].Level
	promotion := g.Player.DbAvatar.Avatar[avatarId].Promotion

	avatarPromotion := gdconf.GetAvatarPromotionConfigMap()[strconv.Itoa(int(avatarId))][strconv.Itoa(int(promotion))]

	hp := avatarPromotion.HPBase.Value + (avatarPromotion.HPAdd.Value * float64(level))
	return uint32(hp)
}

func NewAvatar(data *PlayerData, mainAvatar proto.HeroBasicType) *PlayerData {
	data.DbAvatar = new(DbAvatar)
	data.DbAvatar.MainAvatar = mainAvatar
	data.DbAvatar.MainAvatarList = []uint32{8001, 8002, 8003, 8004}
	data.DbAvatar.Avatar = make(map[uint32]*Avatar)

	return data
}

func (g *Game) AddAvatar(avatarId uint32) {
	for _, avatar := range g.Player.DbAvatar.Avatar {
		if avatar.AvatarId == avatarId {
			g.AddMaterial(avatarId+10000, 1)
			return
		}
	}
	avatar := new(Avatar)
	// TODO
	avatar.AvatarId = avatarId
	avatar.Exp = 0
	avatar.Level = 1
	avatar.Type = proto.AvatarType_AVATAR_FORMAL_TYPE
	avatar.FirstMetTimestamp = uint64(time.Now().Unix())
	avatar.Promotion = 0
	avatar.Rank = 0
	avatar.Hp = 10000
	avatar.EquipmentUniqueId = 0
	avatar.SpBar = &SpBarInfo{
		CurSp: 10000,
		MaxSp: 10000,
	}
	avatar.TakenRewards = make([]uint32, 0)
	avatar.BuffList = 0
	for id, level := range gdconf.GetAvatarSkilltreeListById(avatarId) {
		skilltreeList := &Skilltree{
			PointId: id,
			Level:   level,
		}
		avatar.SkilltreeList = append(avatar.SkilltreeList, skilltreeList)
	}

	g.Player.DbAvatar.Avatar[avatarId] = avatar
	g.AvatarPlayerSyncScNotify(avatarId)
	// g.ScenePlaneEventScNotify(avatarId, 1)
}

func (g *Game) GetSkilltree(avatarId uint32) []*proto.AvatarSkillTree {
	skilltreeList := make([]*proto.AvatarSkillTree, 0)
	for _, dbSkilltreeList := range g.Player.DbAvatar.Avatar[avatarId].SkilltreeList {
		if dbSkilltreeList.Level == 0 {
			continue
		}
		skilltree := &proto.AvatarSkillTree{
			PointId: dbSkilltreeList.PointId,
			Level:   dbSkilltreeList.Level,
		}
		skilltreeList = append(skilltreeList, skilltree)
	}
	return skilltreeList
}

func (g *Game) AvatarPlayerSyncScNotify(avatarId uint32) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
	}
	avatardb := g.Player.DbAvatar.Avatar[avatarId]
	avatar := &proto.Avatar{
		SkilltreeList:     g.GetSkilltree(avatarId),
		Exp:               avatardb.Exp,
		BaseAvatarId:      avatarId,
		Rank:              avatardb.Rank,
		EquipmentUniqueId: avatardb.EquipmentUniqueId,
		EquipRelicList:    make([]*proto.EquipRelic, 0),
		TakenRewards:      avatardb.TakenRewards,
		FirstMetTimestamp: avatardb.FirstMetTimestamp,
		Promotion:         avatardb.Promotion,
		Level:             avatardb.Level,
	}
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	g.Send(cmd.PlayerSyncScNotify, notify)
}
