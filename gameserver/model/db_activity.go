package model

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewActivity() *spb.Activity {
	return &spb.Activity{
		TrialActivity: &spb.TrialActivity{},
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

func (g *PlayerData) GetLoginActivity() map[uint32]uint32 {
	if g.GetActivity().ActivityLogin == nil {
		g.GetActivity().ActivityLogin = make(map[uint32]uint32)
	}
	return g.GetActivity().ActivityLogin
}
