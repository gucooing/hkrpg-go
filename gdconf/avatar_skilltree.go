package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	"github.com/hjson/hjson-go/v4"
)

type AvatarSkilltree struct {
	PointID        uint32   `json:"PointID"`
	Level          uint32   `json:"Level"`
	MaxLevel       uint32   `json:"MaxLevel"`
	AvatarID       uint32   `json:"AvatarID"`
	DefaultUnlock  bool     `json:"DefaultUnlock"` // 是否默认解锁?
	LevelUpSkillID []uint32 `json:"LevelUpSkillID"`
	PrePoint       []uint32 `json:"PrePoint"` // 前置解锁技能
}

func (g *GameDataConfig) loadAvatarSkilltree() {
	g.AvatarSkilltreeMap = make(map[uint32]map[uint32]*AvatarSkilltree)
	avatarSkilltreeMap := make([]*AvatarSkilltree, 0)
	name := "AvatarSkillTreeConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarSkilltreeMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range avatarSkilltreeMap {
		if g.AvatarSkilltreeMap[v.PointID] == nil {
			g.AvatarSkilltreeMap[v.PointID] = make(map[uint32]*AvatarSkilltree)
		}
		g.AvatarSkilltreeMap[v.PointID][v.Level] = v
	}
	logger.Info(text.GetText(17), len(g.AvatarSkilltreeMap), name)
}

func GetAvatarSkilltreeListById(avatarId uint32) map[uint32]uint32 {
	skilltreeList := make(map[uint32]uint32)
	for _, gdconf := range getConf().AvatarSkilltreeMap {
		if gdconf[1].AvatarID == avatarId {
			skilltree := &proto.AvatarSkillTree{
				PointId: gdconf[1].PointID,
			}
			if gdconf[1].DefaultUnlock {
				skilltree.Level = 1
			} else {
				skilltree.Level = 0
			}
			skilltreeList[skilltree.PointId] = skilltree.Level
		}
	}
	return skilltreeList
}

func GetAvatarSkilltreeBySkillId(skillId, level uint32) *AvatarSkilltree {
	return getConf().AvatarSkilltreeMap[skillId][level]
}

func GetAvatarSkilltreeMap() map[uint32]map[uint32]*AvatarSkilltree {
	return getConf().AvatarSkilltreeMap
}
