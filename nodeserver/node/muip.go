package node

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) muipRecvHandle() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! DISPATCH SERVICE MAIN LOOP PANIC !!!")
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
			s.muipRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) muipRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.MuipToNodePingReq:
		s.MuipToNodePingReq(serviceMsg)
	case cmd.GmGive:
		s.GmGive(serviceMsg)
	case cmd.GmWorldLevel:
		s.GmWorldLevel(serviceMsg)
	case cmd.DelItem:
		s.DelItem(serviceMsg)
	case cmd.MaxCurAvatar:
		s.MaxCurAvatar(serviceMsg)
	default:
		logger.Info("muip -> node error cmdid:%v", cmdId)
	}
}

func (s *Service) MuipToNodePingReq(serviceMsg pb.Message) {
	s.lastAliveTime = time.Now().Unix()
	req := serviceMsg.(*spb.MuipToNodePingReq)
	rsp := &spb.MuipToNodePingRsp{
		MuipServerTime: req.MuipServerTime,
		NodeServerTime: time.Now().Unix(),
		ServiceList:    make(map[uint32]*spb.MuipServiceAll),
	}
	for serverType, serviceList := range s.n.GetAllService() {
		muipServiceAll := &spb.MuipServiceAll{
			ServiceList: make([]*spb.ServiceAll, 0),
		}
		for _, service := range serviceList {
			muipServiceAll.ServiceList = append(muipServiceAll.ServiceList, &spb.ServiceAll{
				ServiceType: service.ServerType,
				Addr:        service.Addr,
				PlayerNum:   service.PlayerNum,
				AppId:       service.AppId,
				Port:        service.Port,
			})
		}
		rsp.ServiceList[serverType] = muipServiceAll
	}

	s.sendHandle(cmd.MuipToNodePingRsp, rsp)
}

func (s *Service) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	if gs, _, ok := s.getPlayerStatusRedis(req.PlayerUid); ok {
		notify := &spb.GmGive{
			PlayerUid: req.PlayerUid,
			ItemId:    req.ItemId,
			ItemCount: req.ItemCount,
			GiveAll:   req.GiveAll,
		}
		gs.sendHandle(cmd.GmGive, notify)
	}
}

func (s *Service) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	if gs, _, ok := s.getPlayerStatusRedis(req.PlayerUid); ok {
		notify := &spb.GmWorldLevel{
			PlayerUid:  req.PlayerUid,
			WorldLevel: req.WorldLevel,
		}
		gs.sendHandle(cmd.GmWorldLevel, notify)
	}
}

func (s *Service) DelItem(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.DelItem)
	if gs, _, ok := s.getPlayerStatusRedis(req.PlayerUid); ok {
		notify := &spb.DelItem{
			PlayerUid: req.PlayerUid,
		}
		gs.sendHandle(cmd.DelItem, notify)
	}
}

func (s *Service) MaxCurAvatar(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.MaxCurAvatar)
	if gs, _, ok := s.getPlayerStatusRedis(req.PlayerUid); ok {
		notify := &spb.MaxCurAvatar{
			PlayerUid: req.PlayerUid,
			AvatarId:  req.AvatarId,
			All:       req.All,
		}
		gs.sendHandle(cmd.MaxCurAvatar, notify)
	}
}

func (s *Service) getPlayerStatusRedis(accountId uint32) (*Service, *spb.PlayerStatusRedisData, bool) {
	if bin, ok := s.n.Store.GetPlayerStatus(strconv.Itoa(int(accountId))); ok {
		statu := new(spb.PlayerStatusRedisData)
		err := pb.Unmarshal(bin, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			return nil, nil, false
		}
		gs := s.n.GetAllServiceByTypeId(spb.ServerType_SERVICE_GAME, statu.GameserverId)
		if gs == nil {
			return nil, nil, false
		}
		return gs, statu, true
	}
	return nil, nil, false
}
