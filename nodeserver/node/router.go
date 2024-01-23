package node

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) RegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionReq: // 服务注册
		s.ServiceConnectionReq(serviceMsg)
	case cmd.GetServerOuterAddrReq: // 心跳
		s.GetServerOuterAddrReq(serviceMsg)
	case cmd.PlayerLoginReq: // 玩家登录通知
		s.PlayerLoginReq(serviceMsg)
	case cmd.PlayerLogoutReq: // 玩家退出登录通知
		s.PlayerLogoutReq(serviceMsg)
		// 下面是GM包
	case cmd.GmGive:
		s.GmGive(serviceMsg)
	case cmd.GmWorldLevel:
		s.GmWorldLevel(serviceMsg)
	}
}
