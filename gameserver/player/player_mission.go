package player

import (
	"regexp"
	"strconv"

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
	finishMainDb := g.GetFinishMainMissionList()
	for _, id := range req.MainMissionIdList {
		if finishMainDb[id] != nil {
			rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, id)
		} else {
			rsp.UnfinishedMainMissionIdList = append(rsp.UnfinishedMainMissionIdList, id)
		}
	}
	for _, id := range req.SubMissionIdList {
		rsp.SubMissionStatusList = append(rsp.SubMissionStatusList, &proto.Mission{
			Id:     id,
			Status: proto.MissionStatus_MISSION_FINISH,
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
			Status:          proto.MissionStatus(main.Status),
			CustomValueList: make([]*proto.AHJMIKLBOEK, 0),
			Id:              main.MissionId,
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
	subMissionId := getMissionUin32(req.TalkStr) // 获取子任务
	g.UpSubMainMission(subMissionId)             // 完成子任务
	nextSub := g.GetNextSubMission(subMissionId) // 获取子任务接下来的任务
	g.Send(cmd.StartFinishSubMissionScNotify, &proto.StartFinishSubMissionScNotify{SubMissionId: subMissionId})
	g.MissionPlayerSyncScNotify(nextSub, []uint32{subMissionId}) // 发送通知
	g.Send(cmd.FinishTalkMissionScRsp, &proto.FinishTalkMissionScRsp{TalkStr: req.TalkStr})
}

func getMissionUin32(talkStr string) uint32 {
	pattern := regexp.MustCompile(`\d+`)
	matches := pattern.FindAllString(talkStr, -1)
	var numbers uint32
	for _, match := range matches {
		num, err := strconv.ParseUint(match, 10, 32)
		if err == nil {
			numbers = uint32(num)
			break
		}
	}
	return numbers
}

func (g *GamePlayer) MissionPlayerSyncScNotify(nextSub, finish []uint32) {
	notify := &proto.PlayerSyncScNotify{
		MissionSync: &proto.MissionSync{
			MissionList: make([]*proto.Mission, 0),
		},
	}
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	for _, sub := range finish {
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
