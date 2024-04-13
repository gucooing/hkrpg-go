package gs

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.GameToNodePingRsp:
		s.GameToNodePingRsp(serviceMsg)
	case cmd.NodeToGsPlayerLogoutNotify:
		s.NodeToGsPlayerLogoutNotify(serviceMsg)
	// 下面是gm
	case cmd.GmGive:
		s.GmGive(serviceMsg) // 获取物品
	case cmd.GmWorldLevel:
		s.GmWorldLevel(serviceMsg) // 设置世界等级
	case cmd.DelItem:
		s.DelItem(serviceMsg) // 清空背包
	default:
		logger.Info("node -> game error cmdid:%v", cmdId)
	}
}
