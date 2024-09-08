package mq

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func (m *MessageQueue) SendToGate(appid uint32, msg *NetMsg) {
	msg.ServerType = spb.ServerType_SERVICE_GATE
	msg.AppId = appid
	msg.OriginServerAppId = m.appId
	msg.OriginServerType = m.serverType
	m.netMsgInput <- msg
}

func (m *MessageQueue) SendToGame(appid uint32, msg *NetMsg) {
	msg.ServerType = spb.ServerType_SERVICE_GAME
	msg.AppId = appid
	msg.OriginServerAppId = m.appId
	msg.OriginServerType = m.serverType
	m.netMsgInput <- msg
}
