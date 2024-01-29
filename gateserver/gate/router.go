package gate

import (
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	case cmd.PlayerLogoutNotify:
		s.PlayerLogoutNotify(serviceMsg)
	default:

	}
}

func (p *PlayerGame) GameRegisterMessage(cmdId uint16, playerMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLoginRsp:
		p.IsConnect = true
	default:
		p.GameToGate(cmdId, playerMsg)
	}
}

func (p *PlayerGame) PlayerRegisterMessage(cmdId uint16, tcpMsg *alg.PackMsg) {
	switch cmdId {
	case cmd.PlayerHeartBeatCsReq:
		p.HandlePlayerHeartBeatCsReq(tcpMsg.ProtoData) // 心跳包
	case cmd.PlayerLogoutCsReq: // 退出游戏
		logger.Info("[UID:%v]离线目标GameServer:%v", p.Uid, p.GameAppId)
		req := &spb.PlayerLogoutReq{
			PlayerUid: p.Uid,
		}
		p.PlayerOfflineReason = spb.PlayerOfflineReason_OFFLINE_DRIVING
		GAMESERVER.sendNode(cmd.PlayerLogoutReq, req)
		p.KcpConn.Close()
		p.GameConn.Close()
		delete(GAMESERVER.sessionMap, p.Uid)
	case cmd.PlayerLoginCsReq:
		p.PlayerOfflineReason = spb.PlayerOfflineReason_OFFLINE_GAME_ERROR
		p.GateToGame(tcpMsg)
	default:
		p.GateToGame(tcpMsg)
	}
}
