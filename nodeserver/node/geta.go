package node

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) gateRecvHandle() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE SERVICE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.n.killService(s)
			return
		}
	}()

	for {
		bin, err := s.Conn.Read()
		if err != nil {
			s.n.killService(s)
			break
		}
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
	case cmd.GateToNodePingReq: // 心跳包
		s.GateToNodePingReq(serviceMsg)
	case cmd.PlayerMsgGateToNodeNotify:
		s.PlayerMsgGateToNodeNotify(serviceMsg)
	default:
		logger.Info("gateRegister error cmdid:%v", cmdId)
	}
}

func (s *Service) GateToNodePingReq(serviceMsg pb.Message) {
	s.lastAliveTime = time.Now().Unix()
	req := serviceMsg.(*spb.GateToNodePingReq)
	if req.ServiceType != s.ServerType {
		logger.Debug("Service registration failed")
		s.n.killService(s)
		return
	}
	s.PlayerNum = req.PlayerNum
	rsp := &spb.GateToNodePingRsp{
		GameServiceList: make([]*spb.ServiceAll, 0),
		GateTime:        req.GateTime,
		NodeTime:        time.Now().UnixNano() / 1e6,
	}
	for _, service := range s.n.GetAllServiceByType(spb.ServerType_SERVICE_GAME) {
		serviceAll := &spb.ServiceAll{
			ServiceType: service.ServerType,
			Addr:        service.Addr,
			Port:        service.Port,
			PlayerNum:   service.PlayerNum,
			AppId:       service.AppId,
		}
		rsp.GameServiceList = append(rsp.GameServiceList, serviceAll)
	}
	for _, service := range s.n.GetAllServiceByType(spb.ServerType_SERVICE_MULTI) {
		rsp.MultiService = &spb.ServiceAll{
			ServiceType: service.ServerType,
			Addr:        service.Addr,
			Port:        service.Port,
			PlayerNum:   service.PlayerNum,
			AppId:       service.AppId,
		}
	}
	s.sendHandle(cmd.GateToNodePingRsp, rsp)
}

func (s *Service) PlayerMsgGateToNodeNotify(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerMsgGateToNodeNotify)
	switch req.MsgType {
	case spb.PlayerMsgType_PMT_APPLYFRIEND:
		s.ApplyFriend(req)
	}
}

// 添加玩家操作
func (s *Service) ApplyFriend(req *spb.PlayerMsgGateToNodeNotify) {
	if gs, _, ok := s.getPlayerStatusRedis(req.SendUid); ok {
		logger.Debug("玩家:%v,gs:%v", req.SendUid, gs.AppId)
	}
	logger.Info("玩家:%v,向玩家:%v,发起好友申请", req.ApplyUid, req.SendUid)
	bin, _ := database.GetPlayerFriend(s.n.Store.PlayerBriefDataRedis, req.SendUid)
	friend := new(spb.PlayerFriend)
	pb.Unmarshal(bin, friend)
	if friend.RecvApplyFriend == nil {
		friend.RecvApplyFriend = make(map[uint32]*spb.ReceiveApply)
	}
	friend.RecvApplyFriend[req.ApplyUid] = &spb.ReceiveApply{
		ApplyUid:  req.ApplyUid,
		ApplyTime: time.Now().Unix(),
	}
	ubin, err := pb.Marshal(friend)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return
	}
	database.SetPlayerFriend(s.n.Store.PlayerBriefDataRedis, req.SendUid, ubin)

}
