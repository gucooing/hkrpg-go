package node

import (
	"bufio"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) gateRecvHandle() {
	payload := make([]byte, PacketMaxLen)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE SERVICE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.killService()
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := bufio.NewReader(s.Conn).Read(payload)
		if err != nil {
			s.killService()
			break
		}
		bin = payload[:recvLen]
		msgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &msgList, nil)
		for _, msg := range msgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.gateRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) gateRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLoginReq: // 玩家登录通知
		s.gatePlayerLoginReq(serviceMsg)
	case cmd.PlayerLogoutReq: // 玩家退出回复
		s.gatePlayerLogoutReq(serviceMsg)
	case cmd.GetAllServiceGameReq: // 心跳包
		s.gateGetAllServiceGameReq(serviceMsg)
	default:
		logger.Info("gateRegister error cmdid:%v", cmdId)
	}
}

func (s *Service) gatePlayerLoginReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLoginReq)
	rsp := new(spb.PlayerLoginRsp)
	if req.PlayerUid == 0 {
		return
	}
	if player := NODE.PlayerMap[req.PlayerUid]; player == nil {
		NODE.PlayerMap[req.PlayerUid] = &PlayerService{
			GateAppId: s.AppId,
			GameAppId: req.AppId,
			PlayerStatus: &PlayerStatus{
				Status:     spb.PlayerStatus_PlayerStatus_LoggingIn, // 登录中的状态
				GateStatus: spb.PlayerGateStatus_PlayerGateStatus_GateLogin,
				GameStatus: spb.PlayerGameStatus_PlayerGameStatus_GameLogin,
			},
		}
		rsp.PlayerUid = req.PlayerUid
		logger.Info("[UID:%v]登录目标GameServer:%v", req.PlayerUid, req.AppId)
		s.sendHandle(cmd.PlayerLoginRsp, rsp)
	} else {
		logger.Info("[UID:%v]玩家重复登录，通知gate,node离线该玩家", req.PlayerUid)
		player.PlayerStatus = &PlayerStatus{
			Status:     spb.PlayerStatus_PlayerStatus_RepeatLogin, // 登录中的状态
			GateStatus: spb.PlayerGateStatus_PlayerGateStatus_GateWaitLogout,
			GameStatus: spb.PlayerGameStatus_PlayerGameStatus_GameWaitLogout,
		}
		notify := &spb.PlayerLogoutReq{
			PlayerUid:     req.PlayerUid,
			OfflineReason: spb.PlayerOfflineReason_OFFLINE_REPEAT_LOGIN,
		}
		if game := GetPlayerGame(req.PlayerUid); game != nil {
			game.PlayerNum--
			if NODE.PlayerMap[req.PlayerUid].GameAppId != req.AppId {
				// 通知旧game玩家下线
				game.sendHandle(cmd.PlayerLogoutReq, notify)
			}
		}
		if gate := GetPlayerGate(req.PlayerUid); gate != nil {
			gate.PlayerNum--
			if NODE.PlayerMap[req.PlayerUid].GateAppId != s.AppId {
				// 通知旧gate玩家下线
				gate.sendHandle(cmd.PlayerLogoutReq, notify)
			}
		}
	}
	s.PlayerNum++
}

func (s *Service) gatePlayerLogoutReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutReq)
	if player := NODE.PlayerMap[req.PlayerUid]; player != nil {
		logger.Info("[UID:%v]gate退出登录成功", req.PlayerUid)
		player.PlayerStatus.GateStatus = spb.PlayerGateStatus_PlayerGateStatus_GateLogout
	}
	repeatLogin(req.PlayerUid)
}

func (s *Service) gateGetAllServiceGameReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetAllServiceGameReq)
	if req.ServiceType != s.ServerType {
		logger.Debug("Service registration failed")
		s.killService()
		return
	}
	rsp := &spb.GetAllServiceGameRsp{
		GameServiceList: make([]*spb.ServiceAll, 0),
		GateTime:        req.GateTime,
		NodeTime:        time.Now().UnixNano() / 1e6,
	}
	for _, service := range NODE.MapService[spb.ServerType_SERVICE_GAME] {
		serviceAll := &spb.ServiceAll{
			ServiceType: service.ServerType,
			Addr:        service.Addr,
			Port:        service.Port,
			PlayerNum:   service.PlayerNum,
			AppId:       service.AppId,
		}
		rsp.GameServiceList = append(rsp.GameServiceList, serviceAll)
	}
	s.sendHandle(cmd.GetAllServiceGameRsp, rsp)
}
