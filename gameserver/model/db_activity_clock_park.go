package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

/*****************************接口***************************/

func (g *PlayerData) GetClockParkInfoList() []*proto.ClockParkInfo {
	list := make([]*proto.ClockParkInfo, 0)
	for _, conf := range gdconf.GetClockParkScriptConfigMap() {
		if g.GetUnlockStatus(conf.ScriptUnlockCondition) {
			list = append(list, &proto.ClockParkInfo{
				ScriptId:   conf.ActivityStudioScriptID,
				ChapterIds: []uint32{conf.StartChapterID},
			})
		}
	}

	return list
}
