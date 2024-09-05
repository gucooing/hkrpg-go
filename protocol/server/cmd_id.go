package server

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	GateTcpMqHandshakeReq = 10006
	GateTcpMqHandshakeRsp = 10106

	ServiceConnectionReq     = 10000
	ServiceConnectionRsp     = 10100
	GateLoginGameRsp         = 10001
	GateLoginGameReq         = 10101
	GateToGameMsgNotify      = 10002
	GameToGateMsgNotify      = 10102
	GetAllServiceGateReq     = 10003
	GetAllServiceGateRsp     = 10103
	MultiToNodePingReq       = 10004
	MultiToNodePingRsp       = 10104
	MuipToNodePingReq        = 10005
	MuipToNodePingRsp        = 10105
	GameToNodePingReq        = 10007
	GameToNodePingRsp        = 10107
	GateGamePingReq          = 10008
	GateGamePingRsp          = 10108
	GateGamePlayerLoginReq   = 10009
	GateGamePlayerLoginRsp   = 10109
	GetToGamePlayerLogoutReq = 10010
	GetToGamePlayerLogoutRsp = 10110
	GateLoginMultiReq        = 10011
	GateLoginMultiRsp        = 10111
	GateToNodePingReq        = 10012
	GateToNodePingRsp        = 10112

	GateToGamePlayerLogoutNotify = 11000
	PlayerMsgGateToNodeNotify    = 11001
	// PlayerLoginNotify            = 11002
	// NodeToGsPlayerLogoutNotify   = 11003
	GameToGatePlayerLogoutNotify = 11004

	GmGive       = 12001
	GmWorldLevel = 12002
	DelItem      = 12003
	MaxCurAvatar = 12004
	GmMission    = 12005
)

func (c *CmdProtoMap) registerAllMessage() {
	// seever
	c.regMsg(GateTcpMqHandshakeReq, func() any { return new(spb.GateTcpMqHandshakeReq) })
	c.regMsg(GateTcpMqHandshakeRsp, func() any { return new(spb.GateTcpMqHandshakeRsp) })

	c.regMsg(GmGive, func() any { return new(spb.GmGive) })
	c.regMsg(GmWorldLevel, func() any { return new(spb.GmWorldLevel) })
	c.regMsg(DelItem, func() any { return new(spb.DelItem) })
	c.regMsg(MaxCurAvatar, func() any { return new(spb.MaxCurAvatar) })
	c.regMsg(ServiceConnectionReq, func() any { return new(spb.ServiceConnectionReq) })
	c.regMsg(ServiceConnectionRsp, func() any { return new(spb.ServiceConnectionRsp) })
	c.regMsg(GateLoginGameRsp, func() any { return new(spb.GateLoginGameRsp) })
	c.regMsg(GateLoginGameReq, func() any { return new(spb.GateLoginGameReq) })
	c.regMsg(GateToGameMsgNotify, func() any { return new(spb.GateToGameMsgNotify) })
	c.regMsg(GameToGateMsgNotify, func() any { return new(spb.GameToGateMsgNotify) })
	c.regMsg(GetAllServiceGateReq, func() any { return new(spb.GetAllServiceGateReq) })
	c.regMsg(GetAllServiceGateRsp, func() any { return new(spb.GetAllServiceGateRsp) })
	c.regMsg(MultiToNodePingReq, func() any { return new(spb.MultiToNodePingReq) })
	c.regMsg(MultiToNodePingRsp, func() any { return new(spb.MultiToNodePingRsp) })
	c.regMsg(MuipToNodePingReq, func() any { return new(spb.MuipToNodePingReq) })
	c.regMsg(MuipToNodePingRsp, func() any { return new(spb.MuipToNodePingRsp) })
	c.regMsg(GateGamePingReq, func() any { return new(spb.GateGamePingReq) })
	c.regMsg(GateGamePingRsp, func() any { return new(spb.GateGamePingRsp) })
	c.regMsg(GateGamePlayerLoginReq, func() any { return new(spb.GateGamePlayerLoginReq) })
	c.regMsg(GateGamePlayerLoginRsp, func() any { return new(spb.GateGamePlayerLoginRsp) })
	c.regMsg(GetToGamePlayerLogoutReq, func() any { return new(spb.GetToGamePlayerLogoutReq) })
	c.regMsg(GetToGamePlayerLogoutRsp, func() any { return new(spb.GetToGamePlayerLogoutRsp) })
	c.regMsg(GateLoginMultiReq, func() any { return new(spb.GateLoginMultiReq) })
	c.regMsg(GateLoginMultiRsp, func() any { return new(spb.GateLoginMultiRsp) })
	c.regMsg(GameToGatePlayerLogoutNotify, func() any { return new(spb.GameToGatePlayerLogoutNotify) })
	c.regMsg(GateToGamePlayerLogoutNotify, func() any { return new(spb.GateToGamePlayerLogoutNotify) })
	c.regMsg(GmMission, func() any { return new(spb.GmMission) })
	c.regMsg(PlayerMsgGateToNodeNotify, func() any { return new(spb.PlayerMsgGateToNodeNotify) })
	c.regMsg(GameToNodePingReq, func() any { return new(spb.GameToNodePingReq) })
	c.regMsg(GameToNodePingRsp, func() any { return new(spb.GameToNodePingRsp) })
	c.regMsg(GateToNodePingReq, func() any { return new(spb.GateToNodePingReq) })
	c.regMsg(GateToNodePingRsp, func() any { return new(spb.GateToNodePingRsp) })
}
