package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

// 角色状态改变时需要发送通知
func (g *GamePlayer) PlayerPlayerSyncScNotify() {
	db := g.GetMaterialMap()
	notify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   g.GetNickname(),
			Level:      g.GetLevel(),
			Exp:        db[Exp],
			Hcoin:      db[Hcoin],
			Scoin:      db[Scoin],
			Mcoin:      db[Mcoin],
			Stamina:    db[Stamina],
			WorldLevel: g.GetWorldLevel(),
		},
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) AddTrailblazerExp(num uint32) {
	material := g.GetMaterialMap()
	db := g.GetBasicBin()
	material[Exp] += num
	level, exp, worldLevel := gdconf.GetPlayerLevelConfigByLevel(material[22], g.GetLevel(), g.GetWorldLevel())
	if level == 0 && exp == 0 {
		return
	} else {
		material[Exp] = exp
		db.Level = level
		db.WorldLevel = worldLevel
		g.PlayerPlayerSyncScNotify()
	}
}
