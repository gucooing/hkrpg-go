package gate

import (
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GateServer) NodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.GetAllServiceRsp:
		s.GetAllServiceRsp(serviceMsg)
	default:

	}
}

func (p *PlayerGame) GameRegisterMessage(cmdId uint16, playerMsg pb.Message) {
	switch cmdId {
	default:
		p.GameToGate(cmdId, playerMsg)
	}
}

func (p *PlayerGame) PlayerRegisterMessage(cmdId uint16, tcpMsg *alg.PackMsg) {
	switch cmdId {
	case cmd.PlayerLogoutCsReq:
		req := &spb.PlayerLogoutReq{
			PlayerUid: p.Uid,
		}
		GAMESERVER.sendNode(cmd.PlayerLogoutReq, req)
		p.GateToGame(tcpMsg)
		p.KcpConn.Close()
	default:
		p.GateToGame(tcpMsg)
	}
}
