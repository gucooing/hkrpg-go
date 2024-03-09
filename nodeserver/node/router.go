package node

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

// 公共接口
func (s *Service) RegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionReq: // 服务注册
		s.ServiceConnectionReq(serviceMsg)
	case cmd.GetServerOuterAddrReq: // 心跳
		s.GetServerOuterAddrReq(serviceMsg)
	case cmd.GetAllServiceReq: // 获取目标服务所有
		s.GetAllServiceReq(serviceMsg)
	case cmd.SyncPlayerOnlineDataNotify:
		s.SyncPlayerOnlineDataNotify(serviceMsg) // 同步在线数据
		// 下面是GM包
	case cmd.GmGive:
		s.GmGive(serviceMsg)
	case cmd.GmWorldLevel:
		s.GmWorldLevel(serviceMsg)
	}
}
