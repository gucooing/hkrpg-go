package gs

import (
	"encoding/json"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.PlayerLogoutReq:
		s.PlayerLogoutReq(serviceMsg) // 玩家离线通知
	case cmd.SyncPlayerOnlineDataNotify:
		s.SyncPlayerOnlineDataNotify(serviceMsg) // 在线数据同步
	case cmd.GetAllServiceRsp:
		s.GetAllServiceRsp(serviceMsg)
	// 下面是gm
	case cmd.GmGive:
		s.GmGive(serviceMsg) // 获取物品
	case cmd.GmWorldLevel:
		s.GmWorldLevel(serviceMsg) // 设置世界等级
	default:
		logger.Info("node -> game error cmdid:%v", cmdId)
	}
}

func (s *GameServer) SyncPlayerOnlineDataNotify(serviceMsg pb.Message) {
	noti := serviceMsg.(*spb.SyncPlayerOnlineDataNotify)
	if noti.PlayerUid == 0 || noti.PlayerOnlineData == nil {
		return
	}
	var data *player.PlayerData
	err := json.Unmarshal(noti.PlayerOnlineData, &data)
	if err != nil {
		return
	}
	logger.Info("[UID:%v]在线数据同步成功", noti.PlayerUid)
	s.PlayerMap[noti.PlayerUid].Player = data
}
