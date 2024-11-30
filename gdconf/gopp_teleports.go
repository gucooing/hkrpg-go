package gdconf

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

type Teleports struct {
	Teleports          map[uint32]*PropList           // 本场景下全部锚点
	TeleportsByGroupId map[uint32]*TeleportsByGroupId // 本场景下细分区域
}

type TeleportsByGroupId struct {
	GroupId    uint32
	Teleports  map[uint32]*PropList
	AnchorList map[uint32]*AnchorList
}

func (g *GameDataConfig) goppTeleports() {
	g.Teleports = make(map[uint32]map[uint32]*Teleports)
	floor := getConf().FloorMap
	if floor == nil {
		logger.Error(text.GetText(25))
		return
	}
	for planeId, list := range floor {
		if g.Teleports[planeId] == nil {
			g.Teleports[planeId] = make(map[uint32]*Teleports)
		}
		for floorId, _ := range list {
			if g.Teleports[planeId][floorId] == nil {
				g.Teleports[planeId][floorId] = &Teleports{
					Teleports:          make(map[uint32]*PropList),
					TeleportsByGroupId: make(map[uint32]*TeleportsByGroupId),
				}
			}
			groupList := GetGroupById(planeId, floorId)
			if groupList == nil {
				logger.Debug(text.GetText(26), planeId, floorId)
				continue
			}
			teleports := make(map[uint32]*PropList)
			for groupID, group := range groupList {
				groupTeleports := make(map[uint32]*PropList)
				anchorList := make(map[uint32]*AnchorList)
				for _, prop := range group.PropList {
					if prop.MappingInfoID != 0 && prop.AnchorID != 0 {
						groupTeleports[prop.MappingInfoID] = prop
						teleports[prop.MappingInfoID] = prop
					}
				}
				for _, anchor := range group.AnchorList {
					anchorList[anchor.ID] = anchor
				}
				g.Teleports[planeId][floorId].TeleportsByGroupId[groupID] = &TeleportsByGroupId{
					GroupId:    groupID,
					Teleports:  groupTeleports,
					AnchorList: anchorList,
				}
			}
			g.Teleports[planeId][floorId].Teleports = teleports
		}
	}

	logger.Info(text.GetText(17), len(g.Teleports), "Teleports")
}

func GetTeleportsById(planeId, floorId uint32) *Teleports {
	if getConf().Teleports[planeId] == nil {
		return nil
	}
	return getConf().Teleports[planeId][floorId]
}

func GetGroupTeleportsById(planeId, floorId, groupID uint32) *TeleportsByGroupId {
	if getConf().Teleports[planeId] == nil || getConf().Teleports[planeId][floorId] == nil {
		return nil
	}
	return getConf().Teleports[planeId][floorId].TeleportsByGroupId[groupID]
}
