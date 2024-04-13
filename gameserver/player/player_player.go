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
			Nickname:   g.PlayerPb.Nickname,
			Level:      g.PlayerPb.Level,
			Exp:        g.PlayerPb.Exp,
			Hcoin:      g.GetItem().MaterialMap[1],
			Scoin:      g.GetItem().MaterialMap[2],
			Mcoin:      g.GetItem().MaterialMap[3],
			Stamina:    g.GetItem().MaterialMap[11],
			WorldLevel: g.PlayerPb.WorldLevel,
		},
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) AddTrailblazerExp(num uint32) {
	g.PlayerPb.Exp += num
	level, exp, worldLevel := gdconf.GetPlayerLevelConfigByLevel(g.PlayerPb.Exp, g.PlayerPb.Level, g.PlayerPb.WorldLevel)
	if level == 0 && exp == 0 {
		return
	} else {
		g.PlayerPb.Exp = exp
		g.PlayerPb.Level = level
		g.PlayerPb.WorldLevel = worldLevel
		g.PlayerPlayerSyncScNotify()
	}

}
