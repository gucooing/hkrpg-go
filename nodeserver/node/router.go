package node

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

// 公共接口
func (s *Service) RegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionReq: // 服务注册
		s.ServiceConnectionReq(serviceMsg)
	case cmd.SyncPlayerOnlineDataNotify:
		s.SyncPlayerOnlineDataNotify(serviceMsg) // 同步在线数据 TODO 应改成堵塞
	default:
		logger.Info("error -> node error cmdid:%v", cmdId)
	}
}

/*
func (s *Service) GetAllServiceReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetAllServiceReq)
	rsp := &spb.GetAllServiceRsp{
		ServiceType: req.ServiceType,
		ServiceList: make([]*spb.ServiceAll, 0),
	}
	if req.ServiceType == spb.ServerType_SERVICE_MUIP {
		for typ, app := range NODE.MapService {
			for appid, service := range app {
				serviceList := &spb.ServiceAll{
					ServiceType: typ,
					Addr:        appid,
					PlayerNum:   service.PlayerNum,
					AppId:       service.AppId,
					Port:        service.Port,
				}
				rsp.ServiceList = append(rsp.ServiceList, serviceList)
			}
		}
	}

	s.sendHandle(cmd.GetAllServiceRsp, rsp)
}
*/
