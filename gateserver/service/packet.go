package service

import (
	"github.com/gucooing/hkrpg-go/gateserver/session"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
)

type packetFunc func(g *GateServer, s *session.Session, packMsg *alg.PackMsg)

var packetCaptureMap = map[uint16]packetFunc{ // 抽包
	cmd.PlayerLogoutCsReq: PlayerLogoutCsReq, // 下线请求
}

func (g *GateServer) packetCapture(s *session.Session, packMsg *alg.PackMsg) {
	handelFunc, ok := packetCaptureMap[packMsg.CmdId]
	if !ok {
		netMsg := &mq.NetMsg{
			MsgType:        mq.GameServer,
			Uid:            s.Uid,
			CmdId:          packMsg.CmdId,
			ServiceMsgByte: packMsg.ProtoData,
		}
		g.MessageQueue.SendToGame(s.GameAppId, netMsg)
		return
	}
	handelFunc(g, s, packMsg)
}

func PlayerLogoutCsReq(g *GateServer, s *session.Session, packMsg *alg.PackMsg) {
	g.MessageQueue.SendToGame(s.GameAppId, &mq.NetMsg{ // 通知GS下线
		MsgType: mq.PlayerLogout,
		Uid:     s.Uid,
	})
	g.DelSession(s)
}
