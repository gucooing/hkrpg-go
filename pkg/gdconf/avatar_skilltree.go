package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	g.AvatarSkilltreeMap = make(map[string]map[string]*AvatarSkilltree)
	playerElementsFilePath := g.excelPrefix + "AvatarSkillTreeConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarSkilltreeMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarSkillTreeConfig", len(g.AvatarSkilltreeMap))
}

func GetAvatarSkilltreeListById(avatarId uint32) map[uint32]uint32 {
	skilltreeList := make(map[uint32]uint32)
	for _, gdconf := range CONF.AvatarSkilltreeMap {
		if gdconf[strconv.Itoa(1)].AvatarID == avatarId {
			skilltree := &proto.AvatarSkillTree{
				PointId: gdconf[strconv.Itoa(1)].PointID,
			}
			if gdconf[strconv.Itoa(1)].DefaultUnlock {
				skilltree.Level = 1
			} else {
				skilltree.Level = 0
			}
			skilltreeList[skilltree.PointId] = skilltree.Level
		}
	}
	return skilltreeList
}

func GetMainAvatarSkilltreeListById(avatarId uint32) []*proto.AvatarSkillTree {
	skilltreeList := make([]*proto.AvatarSkillTree, 0)
	for _, gdconf := range CONF.AvatarSkilltreeMap {
		if gdconf[strconv.Itoa(1)].AvatarID == avatarId {
			skilltree := &proto.AvatarSkillTree{}
			if gdconf[strconv.Itoa(1)].DefaultUnlock {
				skilltree.Level = 1
			} else {
				continue
			}
			skilltree.PointId = gdconf[strconv.Itoa(1)].PointID
			skilltreeList = append(skilltreeList, skilltree)
		}
	}
	return skilltreeList
}

func GetAvatarSkilltreeMap() map[string]map[string]*AvatarSkilltree {
	return CONF.AvatarSkilltreeMap
}
