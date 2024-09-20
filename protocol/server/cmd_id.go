package server

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	GateTcpMqHandshakeReq = 10001
	GateTcpMqHandshakeRsp = 10002
	PlayerLogoutReq       = 10003
	PlayerLogoutRsp       = 10004

	PlayerLogoutNotify = 11001
	GmGive             = 12001
	GmWorldLevel       = 12002
	DelItem            = 12003
	MaxCurAvatar       = 12004
	GmMission          = 12005
)

func (c *CmdProtoMap) registerAllMessage() {
	// seever
	c.regMsg(GateTcpMqHandshakeReq, func() any { return new(spb.GateTcpMqHandshakeReq) })
	c.regMsg(GateTcpMqHandshakeRsp, func() any { return new(spb.GateTcpMqHandshakeRsp) })
	c.regMsg(PlayerLogoutReq, func() any { return new(spb.PlayerLogoutReq) })
	c.regMsg(PlayerLogoutRsp, func() any { return new(spb.PlayerLogoutRsp) })

	c.regMsg(PlayerLogoutNotify, func() any { return new(spb.PlayerLogoutNotify) })
	c.regMsg(GmGive, func() any { return new(spb.GmGive) })
	c.regMsg(GmWorldLevel, func() any { return new(spb.GmWorldLevel) })
	c.regMsg(DelItem, func() any { return new(spb.DelItem) })
	c.regMsg(MaxCurAvatar, func() any { return new(spb.MaxCurAvatar) })
	c.regMsg(GmMission, func() any { return new(spb.GmMission) })
}
