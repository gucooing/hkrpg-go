package node

import (
	"time"

	"github.com/gucooing/hkrpg-go/nodeserver/sociality"
	"github.com/gucooing/hkrpg-go/pkg/alg"
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
		s.ApplyFriend(req) // 申请添加好友
	case spb.PlayerMsgType_PMT_ACCEPTFRIEND:
		s.AcceptFriend(req) // 同意好友申请
	}
}

// 申请添加好友
func (s *Service) ApplyFriend(req *spb.PlayerMsgGateToNodeNotify) {
	if gs, _, ok := s.getPlayerStatusRedis(req.SendUid); ok {
		logger.Debug("玩家:%v,gs:%v", req.SendUid, gs.AppId)
	}
	logger.Info("玩家:%v,向玩家:%v,发起好友申请", req.ApplyUid, req.SendUid)
	friend := sociality.GetApplyFriendByUid(req.SendUid)
	if friend.RecvApplyFriend == nil {
		friend.RecvApplyFriend = make(map[uint32]*spb.ReceiveApply)
	}
	friend.RecvApplyFriend[req.ApplyUid] = &spb.ReceiveApply{
		ApplyUid:  req.ApplyUid,
		ApplyTime: time.Now().Unix(),
	}
	if !sociality.UpdateApplyFriend(req.SendUid) {
		logger.Warn("UpdateApplyFriend写入失败")
		s.ApplyFriend(req) // 写失败，再来一次
	}
}

// 同意好友申请
func (s *Service) AcceptFriend(req *spb.PlayerMsgGateToNodeNotify) {
	friend := sociality.GetApplyFriendByUid(req.SendUid)
	if friend.RecvApplyFriend == nil {
		friend.RecvApplyFriend = make(map[uint32]*spb.ReceiveApply)
	}
	if req.IsAcceptFriend {
		if gs, _, ok := s.getPlayerStatusRedis(req.ApplyUid); ok {
			logger.Debug("玩家:%v,gs:%v", req.ApplyUid, gs.AppId)
		} else {
			logger.Debug("目标玩家:%v不在线,将好友缓存入redis", req.ApplyUid)
			accep := sociality.GetAcceptApplyFriendByUid(req.ApplyUid)
			if accep.RecvApplyFriend == nil {
				accep.RecvApplyFriend = make(map[uint32]*spb.ReceiveApply)
			}
			accep.RecvApplyFriend[req.SendUid] = &spb.ReceiveApply{
				ApplyUid:  req.SendUid,
				ApplyTime: time.Now().Unix(),
			}
			sociality.UpdateAcceptApplyFriend(req.ApplyUid)
		}
		logger.Debug("玩家:%v处理来自向玩家:%v,发起的好友申请，结果:%s", req.SendUid, req.ApplyUid, req.IsAcceptFriend)

	}
	if friend.RecvApplyFriend[req.ApplyUid] != nil {
		delete(friend.RecvApplyFriend, req.ApplyUid)
	}
	if !sociality.UpdateApplyFriend(req.SendUid) {
		logger.Warn("UpdateApplyFriend写入失败")
		s.AcceptFriend(req) // 写失败，再来一次
	}
}
