package node

import (
	"bufio"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) muipRecvHandle() {
	payload := make([]byte, PacketMaxLen)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! DISPATCH SERVICE MAIN LOOP PANIC !!!")
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
			s.muipRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) muipRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.GmGive:
		s.GmGive(serviceMsg)
	case cmd.GmWorldLevel:
		s.GmWorldLevel(serviceMsg)
	case cmd.GetAllServiceReq:
		s.GetAllServiceReq(serviceMsg)
	default:
		logger.Info("muip -> node error cmdid:%v", cmdId)
	}
}

func (s *Service) GmGive(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmGive)
	if req.PlayerUid == 0 || NODE.PlayerUuidMap[req.PlayerUid] == 0 {
		return
	}
	ps := getPlayerServiceByUuid(req.PlayerUid)
	notify := &spb.GmGive{
		PlayerUid: req.PlayerUid,
		ItemId:    req.ItemId,
		ItemCount: req.ItemCount,
		GiveAll:   req.GiveAll,
		Uuid:      ps.Uuid,
	}
	getGsByAppId(ps.GameAppId).sendHandle(cmd.GmGive, notify)
}

func (s *Service) GmWorldLevel(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GmWorldLevel)
	if req.PlayerUid == 0 || NODE.PlayerUuidMap[req.PlayerUid] == 0 {
		return
	}
	ps := getPlayerServiceByUuid(req.PlayerUid)
	notify := &spb.GmWorldLevel{
		PlayerUid:  req.PlayerUid,
		WorldLevel: req.WorldLevel,
		Uuid:       ps.Uuid,
	}
	getGsByAppId(ps.GameAppId).sendHandle(cmd.GmWorldLevel, notify)
}
