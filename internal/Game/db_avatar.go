package Game

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

type DbAvatar struct {
	Avatar map[uint32]*Avatar
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
}

func AddAvatar(avatarId uint32) *Avatar {
	avatar := new(Avatar)
	// TODO
	avatar.AvatarId = avatarId
	avatar.Exp = 10
	avatar.Level = 10
	avatar.FirstMetTimestamp = uint64(time.Now().Unix())
	avatar.Promotion = 0
	avatar.Rank = 6
	avatar.Hp = 10000
	return avatar
}
func GetKilltreeList(avatarId, level string) []*proto.AvatarSkillTree {
	skilltreeList := make([]*proto.AvatarSkillTree, 0)
	skillList := gdconf.GetAvatarSkilltreeMap()
	for _, a := range skillList {
		if a[level].AvatarID == avatarId {
			pointId, _ := strconv.ParseUint(a[level].PointID, 10, 32)
			skilltree := &proto.AvatarSkillTree{
				PointId: uint32(pointId),
				Level:   1,
			}
			skilltreeList = append(skilltreeList, skilltree)
		}
	}
	return skilltreeList
}
