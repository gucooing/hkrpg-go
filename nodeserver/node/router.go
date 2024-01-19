package node

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) RegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionReq:
		s.ServiceConnectionReq(serviceMsg)
	case cmd.GetServerOuterAddrReq:
		s.GetServerOuterAddrReq(serviceMsg)
	}
}
