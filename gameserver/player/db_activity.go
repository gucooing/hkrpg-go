package player

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetActivity() *spb.Activity {
	if g.PlayerPb.Activity == nil {
		g.PlayerPb.Activity = &spb.Activity{
			TrialActivity: make([]uint32, 0),
			ActivityLogin: make(map[uint32]uint32),
		}
	}
	return g.PlayerPb.Activity
}

func (g *GamePlayer) GetTrialActivity() []uint32 {
	if g.GetActivity().TrialActivity == nil {
		g.GetActivity().TrialActivity = make([]uint32, 0)
	}
	return g.GetActivity().TrialActivity
}

func (g *GamePlayer) GetLoginActivity() map[uint32]uint32 {
	if g.GetActivity().ActivityLogin == nil {
		g.GetActivity().ActivityLogin = make(map[uint32]uint32)
	}
	return g.GetActivity().ActivityLogin
}
