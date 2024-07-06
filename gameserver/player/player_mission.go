package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) MissionAcceptScNotify() {
	notify := &proto.MissionAcceptScNotify{
		SubMissionIdList: make([]uint32, 0),
	}
	g.Send(cmd.MissionAcceptScNotify, notify)
}

func (g *GamePlayer) GetMissionEventDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetMissionEventDataScRsp{
		Retcode:          0,
		ChallengeEventId: 0,
		MissionEventList: make([]*proto.Mission, 0),
	}
	conf := gdconf.GetEventMission()
	for id, mission := range conf {
		if mission.TakeType == "Auto" {
			rsp.MissionEventList = append(rsp.MissionEventList, &proto.Mission{
				Id:       id,
				Progress: 0,
				Status:   proto.MissionStatus_MISSION_DOING,
			})
		}
	}
	g.Send(cmd.GetMissionEventDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetMissionStatusCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetMissionStatusCsReq, payloadMsg)
	req := msg.(*proto.GetMissionStatusCsReq)

	rsp := &proto.GetMissionStatusScRsp{
		FinishedMainMissionIdList:   make([]uint32, 0),
		MissionEventStatusList:      make([]*proto.Mission, 0),
		SubMissionStatusList:        make([]*proto.Mission, 0),
		Retcode:                     0,
		UnfinishedMainMissionIdList: make([]uint32, 0),
		DisabledMainMissionIdList:   make([]uint32, 0),
	}
	rsp.FinishedMainMissionIdList = []uint32{}
	rsp.SubMissionStatusList = make([]*proto.Mission, 0)
	if g.IsJumpMission {
		for _, id := range req.MainMissionIdList {
			rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, id)
		}
		for _, id := range req.SubMissionIdList {
			rsp.SubMissionStatusList = append(rsp.SubMissionStatusList, &proto.Mission{
				Id:     id,
				Status: proto.MissionStatus_MISSION_FINISH,
			})
		}
		g.Send(cmd.GetMissionStatusScRsp, rsp)
		return
	}
	finishMainDb := g.GetFinishMainMissionList()
	finishMainSubDb := g.GetFinishSubMainMissionList()
	curMainSubDb := g.GetSubMainMissionList()
	// 处理主线任务
	for _, id := range req.MainMissionIdList {
		if finishMainDb[id] != nil {
			rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, id)
		} else {
			rsp.UnfinishedMainMissionIdList = append(rsp.UnfinishedMainMissionIdList, id)
		}
	}
	// 处理子任务
	for _, id := range req.SubMissionIdList {
		status := proto.MissionStatus_MISSION_NONE
		if curMainSubDb[id] != nil {
			status = proto.MissionStatus_MISSION_DOING
		}
		if finishMainSubDb[id] != nil {
			status = proto.MissionStatus_MISSION_FINISH
		}
		rsp.SubMissionStatusList = append(rsp.SubMissionStatusList, &proto.Mission{
			Id:     id,
			Status: status,
		})
	}

	g.Send(cmd.GetMissionStatusScRsp, rsp)
}

func (g *GamePlayer) GetMissionDataCsReq(payloadMsg []byte) {
	mainMissionList := g.GetMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()

	rsp := &proto.GetMissionDataScRsp{
		MissionDataList: make([]*proto.MissionData, 0),
		MissionList:     make([]*proto.Mission, 0),
		Retcode:         0,
	}
	// add main
	for _, main := range mainMissionList {
		rsp.MissionDataList = append(rsp.MissionDataList, &proto.MissionData{
			Status: proto.MissionStatus(main.Status),
			Id:     main.MissionId,
		})
	}
	// add sub mission
	for _, subMission := range subMainMissionList {
		rsp.MissionList = append(rsp.MissionList, &proto.Mission{
			Id:       subMission.MissionId,
			Progress: subMission.Progress,
			Status:   proto.MissionStatus(subMission.Status),
		})
	}
	g.Send(cmd.GetMissionDataScRsp, rsp)
}

func (g *GamePlayer) FinishTalkMissionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.FinishTalkMissionCsReq, payloadMsg)
	req := msg.(*proto.FinishTalkMissionCsReq)
	g.TalkStrSubMission(req.TalkStr) // 获取子任务
	g.Send(cmd.FinishTalkMissionScRsp, &proto.FinishTalkMissionScRsp{TalkStr: req.TalkStr})
}

func (g *GamePlayer) MissionPlayerSyncScNotify(nextSub, finSub, finishMain []uint32) {
	if len(nextSub) == 0 && len(finSub) == 0 && len(finishMain) == 0 {
		return
	}
	notify := &proto.PlayerSyncScNotify{
		MissionSync: &proto.MissionSync{
			MissionList:       make([]*proto.Mission, 0),
			MainMissionIdList: finishMain,
		},
	}
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	for _, sub := range finSub {
		db := finishSubMainMissionList[sub]
		if db == nil {
			continue
		}
		notify.MissionSync.MissionList = append(notify.MissionSync.MissionList, &proto.Mission{
			Id:       db.MissionId,
			Progress: db.Progress,
			Status:   proto.MissionStatus(db.Status),
		})
	}
	for _, sub := range nextSub {
		db := subMainMissionList[sub]
		if db == nil {
			continue
		}
		notify.MissionSync.MissionList = append(notify.MissionSync.MissionList, &proto.Mission{
			Id:       db.MissionId,
			Progress: db.Progress,
			Status:   proto.MissionStatus(db.Status),
		})
	}
	g.Send(cmd.PlayerSyncScNotify, notify)
}
