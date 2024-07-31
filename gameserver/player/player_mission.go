package player

import (
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
	rsp := &proto.GetMainMissionCustomValueScRsp{MainMissionList: make([]*proto.MainMission, 0)}
	mainMissionList := g.GetMainMissionList()             // 已接取的主任务
	finishMainMissionList := g.GetFinishMainMissionList() // 已完成的主任务
	for _, id := range req.MainMissionIdList {
		if mainMissionList[id] != nil {
			mission := &proto.MainMission{
				Id:              id,
				CustomValueList: make([]*proto.MissionCustomValue, 0),
				Status:          proto.MissionStatus(mainMissionList[id].Status),
			}
			if mainMissionList[id].MissionCustomValue != nil {
				for _, v := range mainMissionList[id].MissionCustomValue {
					mission.CustomValueList = append(mission.CustomValueList, &proto.MissionCustomValue{
						CustomValue: v.CustomValue,
						Index:       v.Index,
					})
				}
			}
			rsp.MainMissionList = append(rsp.MainMissionList, mission)
		}
		if finishMainMissionList[id] != nil {
			mission := &proto.MainMission{
				Id:              id,
				CustomValueList: make([]*proto.MissionCustomValue, 0),
				Status:          proto.MissionStatus(finishMainMissionList[id].Status),
			}
			if finishMainMissionList[id].MissionCustomValue != nil {
				for _, v := range finishMainMissionList[id].MissionCustomValue {
					mission.CustomValueList = append(mission.CustomValueList, &proto.MissionCustomValue{
						CustomValue: v.CustomValue,
						Index:       v.Index,
					})
				}
			}
			rsp.MainMissionList = append(rsp.MainMissionList, mission)
		}
	}
	g.Send(cmd.GetMainMissionCustomValueScRsp, rsp)
}

func (g *GamePlayer) UpdateTrackMainMissionIdCsReq(payloadMsg []byte) {
	g.Send(cmd.UpdateTrackMainMissionIdScRsp, &proto.UpdateTrackMainMissionIdScRsp{})
}

func (g *GamePlayer) GetMissionEventDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetMissionEventDataScRsp{
		Retcode:          0,
		ChallengeEventId: 0,
		MissionEventList: make([]*proto.Mission, 0),
	}
	for _, mission := range g.GetMainMissionList() {
		rsp.MissionEventList = append(rsp.MissionEventList, &proto.Mission{
			Id:       mission.MissionId,
			Progress: mission.MissionId,
			Status:   proto.MissionStatus(mission.Status),
		})
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
		MainMissionMcvList:          make([]*proto.MainMissionCustomValue, 0),
	}
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
	finishMainDb := g.GetFinishMainMissionList() // 完成的主线任务
	finishSubMissionList := g.GetFinishSubMainMissionList()
	subMissionList := g.GetSubMainMissionList()
	// 处理主线任务
	for _, id := range req.MainMissionIdList {
		rsp.MainMissionMcvList = append(rsp.MainMissionMcvList, &proto.MainMissionCustomValue{MainMissionId: id})
		if finishMainDb[id] != nil {
			rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, id)
		} else {
			rsp.UnfinishedMainMissionIdList = append(rsp.UnfinishedMainMissionIdList, id)
		}
	}
	// 处理子任务
	for _, id := range req.SubMissionIdList {
		status := proto.MissionStatus_MISSION_NONE
		if subMissionList[id] != nil {
			status = proto.MissionStatus_MISSION_DOING
		}
		if finishSubMissionList[id] != nil {
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
		MainMissionList: make([]*proto.MainMission, 0), // doing mainMissionList
		MissionList:     make([]*proto.Mission, 0),     // doing subMissionList
		Retcode:         0,
		// GOBNFADAILM:     1021201, // 102120113 // cur mainMission
	}
	// add main
	for _, main := range mainMissionList {
		rsp.MainMissionList = append(rsp.MainMissionList, &proto.MainMission{
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
	g.Send(cmd.FinishTalkMissionScRsp, &proto.FinishTalkMissionScRsp{TalkStr: req.TalkStr, CustomValueList: req.CustomValueList})
}

func (g *GamePlayer) FinishCosumeItemMissionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.FinishCosumeItemMissionCsReq, payloadMsg)
	req := msg.(*proto.FinishCosumeItemMissionCsReq)
	g.FinishCosumeItemMission(req.SubMissionId)
	g.Send(cmd.FinishCosumeItemMissionScRsp, &proto.FinishCosumeItemMissionScRsp{SubMissionId: req.SubMissionId})
}

func (g *GamePlayer) MissionRewardScNotify() {

}
