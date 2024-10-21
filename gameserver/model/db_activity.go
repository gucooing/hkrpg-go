package model

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewActivity() *spb.Activity {
	return &spb.Activity{
		TrialActivity: make(map[uint32]*spb.TrialActivityInfo),
		ActivityLogin: make(map[uint32]uint32),
	}
}

func (g *PlayerData) GetActivity() *spb.Activity {
	db := g.GetBasicBin()
	if db.Activity == nil {
		db.Activity = NewActivity()
	}
	return db.Activity
}

func (g *PlayerData) GetTrialActivity() map[uint32]*spb.TrialActivityInfo {
	if g.GetActivity().TrialActivity == nil {
		g.GetActivity().TrialActivity = make(map[uint32]*spb.TrialActivityInfo)
	}
	return g.GetActivity().TrialActivity
}

func (g *PlayerData) GetTrialActivityById(stageId uint32) *spb.TrialActivityInfo {
	db := g.GetTrialActivity()
	if db[stageId] == nil {
		db[stageId] = &spb.TrialActivityInfo{
			StageId:     stageId,
			TakenReward: false,
		}
	}
	return db[stageId]
}

func (g *PlayerData) GetLoginActivity() map[uint32]uint32 {
	if g.GetActivity().ActivityLogin == nil {
		g.GetActivity().ActivityLogin = make(map[uint32]uint32)
	}
	return g.GetActivity().ActivityLogin
}
