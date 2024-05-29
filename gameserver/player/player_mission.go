package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) MissionAcceptScNotify() {
	notify := &proto.MissionAcceptScNotify{
		SubMissionIdList: []uint32{100010102, 100010191, 401013101},
	}
	g.Send(cmd.MissionAcceptScNotify, notify)
}
