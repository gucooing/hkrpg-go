package gdconf

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

type Teleports struct {
	GroupId   uint32
	Teleports map[uint32]*PropList
}

func (g *GameDataConfig) goppTeleports() {
	g.Teleports = make(map[uint32]map[uint32]map[uint32]*Teleports)
	floor := CONF.FloorMap
	if floor == nil {
		logger.Error("floor error")
		return
	}
	for planeId, list := range floor {
		if g.Teleports[planeId] == nil {
			g.Teleports[planeId] = make(map[uint32]map[uint32]*Teleports)
		}
		for floorId, _ := range list {
			if g.Teleports[planeId][floorId] == nil {
				g.Teleports[planeId][floorId] = make(map[uint32]*Teleports)
			}
			groupList := GetGroupById(planeId, floorId)
			if groupList == nil {
				logger.Warn("goppTeleports planeId:%v,floorId:%v,error", planeId, floorId)
				continue
			}
			for groupID, group := range groupList {
				teleports := make(map[uint32]*PropList)
				for _, prop := range group.PropList {
					if prop.MappingInfoID != 0 && prop.AnchorID != 0 {
						teleports[prop.MappingInfoID] = prop
					}
				}
				g.Teleports[planeId][floorId][groupID] = &Teleports{
					GroupId:   groupID,
					Teleports: teleports,
				}
			}
		}
	}
	logger.Info("gopp %v Teleports", len(g.Teleports))
}

func GetTeleportsById(planeId, floorId uint32) map[uint32]*Teleports {
	if CONF.Teleports[planeId] == nil {
		return nil
	}
	return CONF.Teleports[planeId][floorId]
}

func GetGroupTeleportsById(planeId, floorId, groupID uint32) *Teleports {
	if CONF.Teleports[planeId] == nil || CONF.Teleports[planeId][floorId] == nil {
		return nil
	}
	return CONF.Teleports[planeId][floorId][groupID]
}
