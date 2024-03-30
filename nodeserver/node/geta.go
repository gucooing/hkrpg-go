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
	case cmd.PlayerLoginRsp: // 玩家退出回复
		s.gatePlayerLoginRsp(serviceMsg)
	case cmd.GetAllServiceGameReq: // 心跳包
		s.gateGetAllServiceGameReq(serviceMsg)

	case cmd.PlayerLogoutNotify:
		s.gatePlayerLogoutNotify(serviceMsg)
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
		logger.Info("[UID:%v]玩家已登录gate:%v", req.PlayerUid, s.AppId)
		s.sendHandle(cmd.PlayerLoginRsp, rsp)
		s.PlayerNum++
	} else {
		notify := &spb.PlayerLogoutReq{
			PlayerUid:     req.PlayerUid,
			OfflineReason: spb.PlayerOfflineReason_OFFLINE_REPEAT_LOGIN,
		}
		if player.GateAppId == s.AppId {
			// 同网关登录情况
			// TODO 此处状态问题，没有保存新game，如果是同gate，同game登录无事
			logger.Info("[UID:%v]玩家同网关登录，通知game离线该玩家", req.PlayerUid)
			player.PlayerStatus = &PlayerStatus{
				Status:     spb.PlayerStatus_PlayerStatus_RepeatLogin, // 登录中的状态
				GateStatus: spb.PlayerGateStatus_PlayerGateStatus_GateLogout,
				GameStatus: spb.PlayerGameStatus_PlayerGameStatus_GameWaitLogout,
			}
		} else {
			logger.Info("[UID:%v]玩家重复登录，通知gate,game离线该玩家", req.PlayerUid)
			player.PlayerStatus = &PlayerStatus{
				Status:     spb.PlayerStatus_PlayerStatus_RepeatLogin, // 登录中的状态
				GateStatus: spb.PlayerGateStatus_PlayerGateStatus_GateWaitLogout,
				GameStatus: spb.PlayerGameStatus_PlayerGameStatus_GameWaitLogout,
			}
			s.PlayerNum++
		}
		if gate := GetPlayerGate(req.PlayerUid); gate != nil {
			if gate.AppId != s.AppId {
				// 通知旧gate玩家下线
				gate.PlayerNum--
				gate.sendHandle(cmd.PlayerLogoutReq, notify)
			}
		}
		if game := GetPlayerGame(req.PlayerUid); game != nil {
			game.PlayerNum--
			// 通知旧game玩家下线
			game.sendHandle(cmd.PlayerLogoutReq, notify)
		}

		player.GateAppId = s.AppId
		player.GameAppId = req.AppId
	}
}

// node -> gate 玩家离线通知
func (s *Service) gatePlayerLoginRsp(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLoginRsp)
	if player := NODE.PlayerMap[req.PlayerUid]; player != nil {
		logger.Info("[UID:%v]gate退出登录成功", req.PlayerUid)
		player.PlayerStatus.GateStatus = spb.PlayerGateStatus_PlayerGateStatus_GateLogout
	}
	repeatLogin(req.PlayerUid)
}

// 重复登录后处理结果处理
func repeatLogin(uid uint32) {
	if player := NODE.PlayerMap[uid]; player != nil { // 是否有这个玩家
		if status := player.PlayerStatus; status != nil { // 是否有状态 （意义不明
			// 是否符合登录条件
			if status.GateStatus == spb.PlayerGateStatus_PlayerGateStatus_GateLogout && status.GameStatus == spb.PlayerGameStatus_PlayerGameStatus_GameLogout {
				if gate := GetPlayerGate(uid); gate != nil {
					status.Status = spb.PlayerStatus_PlayerStatus_LoggingIn
					gate.PlayerNum++
					gate.sendHandle(cmd.PlayerLoginRsp, &spb.PlayerLoginRsp{PlayerUid: uid})
					logger.Info("[UID:%v]玩家已登录gate:%v", uid, gate.AppId)
				}
			}
		}
	}
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

/******************************************NewLogin***************************************/

func (s *Service) gatePlayerLogoutNotify(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutNotify)
	logger.Info("[UID:%v]收到玩家被动下线通知", req.Uid)
}
