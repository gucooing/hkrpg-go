package gate

import (
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (s *GateServer) NodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.GetGameOuterAddrRsp:
		s.GetGameOuterAddrRsp(serviceMsg)
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
	default:
		p.GateToGame(tcpMsg)
	}
}
