package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 角色状态改变时需要发送通知
func (g *GamePlayer) PlayerPlayerSyncScNotify() {
	notify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   g.BasicBin.Nickname,
			Level:      g.BasicBin.Level,
			Exp:        g.BasicBin.Exp,
			Hcoin:      g.GetItem().MaterialMap[1],
			Scoin:      g.GetItem().MaterialMap[2],
			Mcoin:      g.GetItem().MaterialMap[3],
			Stamina:    g.GetItem().MaterialMap[11],
			WorldLevel: g.BasicBin.WorldLevel,
		},
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) AddTrailblazerExp(num uint32) {
	g.BasicBin.Exp += num
	level, exp, worldLevel := gdconf.GetPlayerLevelConfigByLevel(g.BasicBin.Exp, g.BasicBin.Level, g.BasicBin.WorldLevel)
	if level == 0 && exp == 0 {
		return
	} else {
		g.BasicBin.Exp = exp
		g.BasicBin.Level = level
		g.BasicBin.WorldLevel = worldLevel
		g.PlayerPlayerSyncScNotify()
	}

}
