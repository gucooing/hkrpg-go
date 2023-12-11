package Game

import (
	"github.com/gucooing/hkrpg-go/gdconf"
)

func (g *Game) AddTrailblazerExp(num uint32) {
	g.Player.Exp += num
	level, exp := gdconf.GetPlayerLevelConfigByLevel(g.Player.Exp, g.Player.Level, g.Player.WorldLevel)
	if level == 0 && exp == 0 {
		return
	} else {
		g.Player.Exp = exp
		g.Player.Level = level
		g.PlayerPlayerSyncScNotify()
	}

}
