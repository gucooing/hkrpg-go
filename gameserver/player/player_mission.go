package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

/*******************每日任务****************/

// 每日实训
var dailyActiveQuestIdList = []uint32{2100132, 2100133, 2100134, 2100139, 2100150, 2100152, 2100153, 2100154}

func (g *GamePlayer) DailyActiveInfoNotify() {
	notify := &proto.DailyActiveInfoNotify{
		DailyActiveLevelList:   make([]*proto.DailyActivityInfo, 0),
		DailyActiveQuestIdList: dailyActiveQuestIdList,
		DailyActivePoint:       500,
	}
	for i := 1; i < 6; i++ {
		dailyActivityInfo := &proto.DailyActivityInfo{
			WorldLevel:       g.GetWorldLevel(),
			Level:            uint32(i),
			DailyActivePoint: uint32(i * 100),
			IsHasTaken:       true,
		}
		notify.DailyActiveLevelList = append(notify.DailyActiveLevelList, dailyActivityInfo)
	}
	g.Send(cmd.DailyActiveInfoNotify, notify)
}

func (g *GamePlayer) GetDailyActiveInfoCsReq(payloadMsg []byte) {
	rsp := &proto.GetDailyActiveInfoScRsp{
		DailyActiveLevelList:   make([]*proto.DailyActivityInfo, 0),
		DailyActiveQuestIdList: dailyActiveQuestIdList,
		DailyActivePoint:       500,
	}

	for i := 1; i < 6; i++ {
		dailyActivityInfo := &proto.DailyActivityInfo{
			WorldLevel:       g.GetWorldLevel(),
			Level:            uint32(i),
			DailyActivePoint: uint32(i * 100),
			IsHasTaken:       true,
		}
		rsp.DailyActiveLevelList = append(rsp.DailyActiveLevelList, dailyActivityInfo)
	}

	g.Send(cmd.GetDailyActiveInfoScRsp, rsp)
}

/*******************任务****************/

func (g *GamePlayer) GetMainMissionCustomValueCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetMainMissionCustomValueCsReq, payloadMsg)
	req := msg.(*proto.GetMainMissionCustomValueCsReq)
	rsp := &proto.GetMainMissionCustomValueScRsp{MissionDataList: make([]*proto.MissionData, 0)}
	mainMissionList := g.GetMainMissionList()             // 已接取的主任务
	finishMainMissionList := g.GetFinishMainMissionList() // 已完成的主任务
	for _, id := range req.MainMissionIdList {
		if mainMissionList[id] != nil {
			rsp.MissionDataList = append(rsp.MissionDataList, &proto.MissionData{
				Id:              id,
				CustomValueList: nil,
				Status:          proto.MissionStatus(mainMissionList[id].Status),
			})
		}
		if finishMainMissionList[id] != nil {
			rsp.MissionDataList = append(rsp.MissionDataList, &proto.MissionData{
				Id:              id,
				CustomValueList: nil,
				Status:          proto.MissionStatus(finishMainMissionList[id].Status),
			})
		}
	}
	g.Send(cmd.GetMainMissionCustomValueScRsp, rsp)
}

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
	g.TalkStrSubMission(req) // 获取子任务
	g.Send(cmd.FinishTalkMissionScRsp, &proto.FinishTalkMissionScRsp{TalkStr: req.TalkStr})
}

func (g *GamePlayer) FinishCosumeItemMissionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.FinishCosumeItemMissionCsReq, payloadMsg)
	req := msg.(*proto.FinishCosumeItemMissionCsReq)
	g.FinishCosumeItemMission(req.SubMissionId)
	g.Send(cmd.FinishCosumeItemMissionScRsp, &proto.FinishCosumeItemMissionScRsp{SubMissionId: req.SubMissionId})
}

func (g *GamePlayer) MissionRewardScNotify() {

}
