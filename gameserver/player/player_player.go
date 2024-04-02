package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
)

func (g *GamePlayer) AddTrailblazerExp(num uint32) {
	g.PlayerPb.Exp += num
	level, exp := gdconf.GetPlayerLevelConfigByLevel(g.PlayerPb.Exp, g.PlayerPb.Level, g.PlayerPb.WorldLevel)
	if level == 0 && exp == 0 {
		return
	} else {
		g.PlayerPb.Exp = exp
		g.PlayerPb.Level = level
		g.PlayerPlayerSyncScNotify()
	}

}
