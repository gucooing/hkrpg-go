package model

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

/*****************************接口***************************/

func (g *PlayerData) GetMusicRhythmLevelList() []*proto.MusicRhythmLevel {
	list := make([]*proto.MusicRhythmLevel, 0)

	return list
}
