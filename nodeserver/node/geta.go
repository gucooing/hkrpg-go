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
	case cmd.GetAllServiceGameReq: // 心跳包
		s.gateGetAllServiceGameReq(serviceMsg)
	case cmd.PlayerLoginNotify:
		s.gatePlayerLoginNotify(serviceMsg)
	case cmd.PlayerLogoutNotify:
		s.gatePlayerLogoutNotify(serviceMsg)
	default:
		logger.Info("gateRegister error cmdid:%v", cmdId)
	}
}

func (s *Service) gateGetAllServiceGameReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetAllServiceGameReq)
	if req.ServiceType != s.ServerType {
		logger.Debug("Service registration failed")
		s.killService()
		return
	}
	s.PlayerNum = req.PlayerNum
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

func (s *Service) gatePlayerLoginNotify(serviceMsg pb.Message) {
	notify := serviceMsg.(*spb.PlayerLoginNotify)
	if notify.Uuid == 0 || notify.Uid == 0 || notify.AccountId == 0 || getGsByAppId(notify.GameServerAppId) == nil || notify.GateServerAppId != s.AppId {
		logger.Error("[UID:%v][gate->node]PlayerLoginNotify通知错误", notify.Uid)
		return
	}
	if NODE.PlayerUuidMap[notify.Uid] != 0 {
		logger.Info("[UID:%v]要上线的玩家还没有下线", notify.Uid)
		return
	}
	AddPlayerUuidMap(notify.Uuid, notify.Uid)
	AddPlayerMap(notify.Uuid, &PlayerService{
		GameAppId: notify.GameServerAppId,
		GateAppId: notify.GateServerAppId,
		Uuid:      notify.Uuid,
		Uid:       notify.Uid,
	})
	s.PlayerNum++
	getGsByAppId(notify.GameServerAppId).PlayerNum++
	logger.Info("[UID:%v][UUID%v]玩家上线", notify.Uid, notify.Uuid)
}

func (s *Service) gatePlayerLogoutNotify(serviceMsg pb.Message) {
	notify := serviceMsg.(*spb.PlayerLogoutNotify)
	if NODE.PlayerUuidMap[notify.Uid] == 0 {
		logger.Info("[UID:%v]找不到要下线的玩家", notify.Uid)
		return
	}
	ps := getPlayerServiceByUuid(notify.Uid)
	gs := getGsByAppId(ps.GameAppId)
	if gs != nil {
		gs.sendHandle(cmd.NodeToGsPlayerLogoutNotify, &spb.NodeToGsPlayerLogoutNotify{Uuid: ps.Uuid})
		gs.PlayerNum--
	}
	DelPlayerUuidMap(ps.Uid)
	DelPlayerMap(ps.Uuid)
	s.PlayerNum--
	logger.Info("[UID:%v][UUID%v]收到玩家被动下线通知", notify.Uid, ps.Uuid)
}
