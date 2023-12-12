package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	"github.com/hjson/hjson-go/v4"
)

type AvatarSkilltree struct {
	PointID       uint32 `json:"PointID"`
	Level         uint32 `json:"Level"`
	MaxLevel      uint32 `json:"MaxLevel"`
	AvatarID      uint32 `json:"AvatarID"`
	DefaultUnlock bool   `json:"DefaultUnlock"` // 是否默认解锁?
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

func GetAvatarSkilltreeById(avatarId, level uint32) []*proto.AvatarSkillTree {
	skilltreeList := make([]*proto.AvatarSkillTree, 0)
	for _, avatarSkilltreeList := range CONF.AvatarSkilltreeMap {
		for _, levelList := range avatarSkilltreeList {
			if levelList.AvatarID == avatarId && levelList.Level == level && levelList.DefaultUnlock {
				skilltree := &proto.AvatarSkillTree{
					PointId: levelList.PointID,
					Level:   level,
				}
				skilltreeList = append(skilltreeList, skilltree)
			}
		}
	}
	return skilltreeList
}

func GetAvatarSkilltreeMap() map[string]map[string]*AvatarSkilltree {
	return CONF.AvatarSkilltreeMap
}
