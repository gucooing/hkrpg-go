package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
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
			WorldLevel:       g.GetPd().GetWorldLevel(),
			Level:            uint32(i),
			DailyActivePoint: uint32(i * 100),
			IsHasTaken:       true,
		}
		notify.DailyActiveLevelList = append(notify.DailyActiveLevelList, dailyActivityInfo)
	}
	g.Send(cmd.DailyActiveInfoNotify, notify)
}

func GetDailyActiveInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetDailyActiveInfoScRsp{
		DailyActiveLevelList:   make([]*proto.DailyActivityInfo, 0),
		DailyActiveQuestIdList: dailyActiveQuestIdList,
		DailyActivePoint:       500,
	}

	for i := 1; i < 6; i++ {
		dailyActivityInfo := &proto.DailyActivityInfo{
			WorldLevel:       g.GetPd().GetWorldLevel(),
			Level:            uint32(i),
			DailyActivePoint: uint32(i * 100),
			IsHasTaken:       true,
		}
		rsp.DailyActiveLevelList = append(rsp.DailyActiveLevelList, dailyActivityInfo)
	}

	g.Send(cmd.GetDailyActiveInfoScRsp, rsp)
}

/*******************任务****************/

func (g *GamePlayer) StartFinishSubMissionScNotify(subMissionId uint32) {
	g.Send(cmd.StartFinishSubMissionScNotify, &proto.StartFinishSubMissionScNotify{
		SubMissionId: subMissionId,
	})
}

func (g *GamePlayer) StartFinishMainMissionScNotify(mainMissionId uint32) {
	g.Send(cmd.StartFinishMainMissionScNotify, &proto.StartFinishMainMissionScNotify{
		MainMissionId: mainMissionId,
	})
}

func GetMainMissionCustomValueCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetMainMissionCustomValueCsReq)
	rsp := &proto.GetMainMissionCustomValueScRsp{MainMissionList: make([]*proto.MainMission, 0)}
	mainMissionList := g.GetPd().GetMainMissionList()             // 已接取的主任务
	finishMainMissionList := g.GetPd().GetFinishMainMissionList() // 已完成的主任务
	if g.GetPd().GetBasicBin().IsJumpMission {
		for _, id := range req.MainMissionIdList {
			rsp.MainMissionList = append(rsp.MainMissionList, &proto.MainMission{
				Id:              id,
				CustomValueList: make([]*proto.MissionCustomValue, 0),
				Status:          proto.MissionStatus_MISSION_FINISH,
			})
		}
	} else {
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
	}

	g.Send(cmd.GetMainMissionCustomValueScRsp, rsp)
}

func UpdateTrackMainMissionIdCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UpdateTrackMainMissionIdCsReq)
	g.Send(cmd.UpdateTrackMainMissionIdScRsp, &proto.UpdateTrackMainMissionIdScRsp{TrackMissionId: req.TrackMissionId})
}

func GetMissionEventDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetMissionEventDataScRsp{
		Retcode:          0,
		ChallengeEventId: 0,
		MissionEventList: make([]*proto.Mission, 0),
	}
	if !g.GetPd().GetBasicBin().IsJumpMission {
		for _, mission := range g.GetPd().GetMainMissionList() {
			rsp.MissionEventList = append(rsp.MissionEventList, &proto.Mission{
				Id:       mission.MissionId,
				Progress: 1,
				Status:   proto.MissionStatus(mission.Status),
			})
		}
	}

	g.Send(cmd.GetMissionEventDataScRsp, rsp)
}

func HandleGetMissionStatusCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetMissionStatusCsReq)

	rsp := &proto.GetMissionStatusScRsp{
		FinishedMainMissionIdList:   make([]uint32, 0),
		MissionEventStatusList:      make([]*proto.Mission, 0),
		SubMissionStatusList:        make([]*proto.Mission, 0),
		Retcode:                     0,
		UnfinishedMainMissionIdList: make([]uint32, 0),
		DisabledMainMissionIdList:   make([]uint32, 0),
		MainMissionMcvList:          make([]*proto.MainMissionCustomValue, 0),
	}
	if g.GetPd().GetBasicBin().IsJumpMission {
		rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, req.MainMissionIdList...)
		for _, id := range req.SubMissionIdList {
			rsp.SubMissionStatusList = append(rsp.SubMissionStatusList, &proto.Mission{
				Id:     id,
				Status: proto.MissionStatus_MISSION_FINISH,
			})
		}
		for _, id := range req.MissionEventIdList {
			rsp.MissionEventStatusList = append(rsp.MissionEventStatusList, &proto.Mission{
				Id:     id,
				Status: proto.MissionStatus_MISSION_FINISH,
			})
		}
		g.Send(cmd.GetMissionStatusScRsp, rsp)
		return
	}
	finishMainDb := g.GetPd().GetFinishMainMissionList() // 完成的主线任务
	finishSubMissionList := g.GetPd().GetFinishSubMainMissionList()
	subMissionList := g.GetPd().GetSubMainMissionList()
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

func GetMissionDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	mainMissionList := g.GetPd().GetMainMissionList()
	subMainMissionList := g.GetPd().GetSubMainMissionList()

	rsp := &proto.GetMissionDataScRsp{
		MainMissionList: make([]*proto.MainMission, 0), // doing mainMissionList
		MissionList:     make([]*proto.Mission, 0),     // doing subMissionList
		Retcode:         0,
		// GOBNFADAILM:     1021201, // 102120113 // cur mainMission
	}

	if !g.GetPd().GetBasicBin().IsJumpMission {
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
	}
	g.Send(cmd.GetMissionDataScRsp, rsp)
}

func FinishTalkMissionCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishTalkMissionCsReq)
	g.InspectMission(g.GetPd().TalkStrSubMission(req)...)
	g.Send(cmd.FinishTalkMissionScRsp, &proto.FinishTalkMissionScRsp{TalkStr: req.TalkStr, CustomValueList: req.CustomValueList})
}

func FinishCosumeItemMissionCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishCosumeItemMissionCsReq)
	allSync := new(model.AllPlayerSync)
	if g.GetPd().FinishCosumeItemMission(req.SubMissionId, allSync) {
		g.InspectMission(req.SubMissionId)
	}
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.FinishCosumeItemMissionScRsp, &proto.FinishCosumeItemMissionScRsp{SubMissionId: req.SubMissionId})
}

func (g *GamePlayer) MissionRewardScNotify() {

}

func (g *GamePlayer) MissionAcceptScNotify(subList []uint32) {
	if len(subList) == 0 {
		return
	}
	g.Send(cmd.MissionAcceptScNotify, &proto.MissionAcceptScNotify{SubMissionIdList: subList})
}

/*********************************检查操作*********************************/

// 登录任务检查
func (g *GamePlayer) LoginReadyMission() {
	if g.GetPd().GetBasicBin().IsJumpMission {
		return
	}
	g.InspectMission()
	g.GetPd().AllCheckMainMission()
}

// 过期任务检查
func (g *GamePlayer) ExpiredMission() {
	mainMissionList := g.GetPd().GetMainMissionList() // 已接取的主任务
	delMainMissionList := make([]uint32, 0)
	addItem := model.NewAddItem(nil)
	for _, info := range mainMissionList {
		if conf := gdconf.GetMainMissionById(info.MissionId); conf != nil {
			if conf.Type == "Branch" { // 活动任务处理

			}
		} else {
			delMainMissionList = append(delMainMissionList, info.MissionId)
		}
	}
	g.GetPd().AddFinishMainMission(delMainMissionList, addItem)
	// 完成主任务中的未完成子任务
	finishSubLists := g.GetPd().CheckMainMission(delMainMissionList)
	g.GetPd().AddFinishSubMission(finishSubLists, addItem)
	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
}

// 任务检查
func (g *GamePlayer) _InspectMission(finishSubMission []uint32) {
	if g.GetPd().GetBasicBin().IsJumpMission {
		return
	}
	// addItem := model.NewAddItem(nil)
	//
	// g.GetPd().AddFinishSubMission(finishSubMission, addItem)
	// finishMainList := make([]uint32, 0)
	// newFinishSubList := make([]uint32, 0)
	// newProgressSubList := make([]uint32, 0)
	// newFinishSubList = append(newFinishSubList, finishSubMission...)
	// for {
	// 	// 接取检查
	// 	mainList, subList := g.AcceptInspectMission()
	// 	newProgressSubList = append(newProgressSubList, subList...) // 将接取的任务添加到同步列表
	// 	// 完成检查
	// 	finishList, finishSubList, progressSubList := g.FinishInspectMission(addItem)
	// 	finishMainList = append(finishMainList, finishList...)              // 将完成的任务添加到同步列表
	// 	newFinishSubList = append(newFinishSubList, finishSubList...)       // 将完成的任务添加到同步列表
	// 	newProgressSubList = append(newProgressSubList, progressSubList...) // 将接取的任务添加到同步列表
	//
	// 	if len(mainList) == 0 &&
	// 		len(subList) == 0 &&
	// 		len(finishList) == 0 &&
	// 		len(finishSubList) == 0 &&
	// 		len(progressSubList) == 0 {
	// 		break
	// 	}
	// }
	//
	// for _, finishId := range finishMainList {
	// 	g.Send(cmd.StartFinishMainMissionScNotify,
	// 		&proto.StartFinishMainMissionScNotify{MainMissionId: finishId})
	// }
	//
	// for i := len(newProgressSubList) - 1; i >= 0; i-- {
	// 	for _, finishId := range newFinishSubList {
	// 		g.AutoServerMissionFinishAction(finishId, addItem)
	// 		g.Send(cmd.StartFinishSubMissionScNotify,
	// 			&proto.StartFinishSubMissionScNotify{SubMissionId: finishId})
	// 		if newProgressSubList[i] == finishId {
	// 			newProgressSubList = append(newProgressSubList[:i], newProgressSubList[i+1:]...)
	// 			break
	// 		}
	// 	}
	// }
	//
	// g.GetPd().AddItem(addItem)
	// addItem.AllSync.MissionFinishMainList = finishMainList
	// addItem.AllSync.MissionFinishSubList = newFinishSubList
	// addItem.AllSync.MissionProgressSubList = newProgressSubList
	// // 检查场景卸加载
	// uninstallGroup, loadedGroupList := g.GetPd().AutoEntryGroup()
	// g.UpSceneGroupRefreshScNotify(uninstallGroup, loadedGroupList)
	// // 命途解锁检查
	// g.GetPd().CheckUnlockMultiPath(addItem.AllSync)
	// if raidId, ok := g.GetPd().CheckRaid(); ok {
	// 	g.RaidInfoNotify(raidId) // raid完成检查
	// }
	// if len(addItem.AllSync.MaterialList) != 0 ||
	// 	len(addItem.AllSync.MissionFinishMainList) != 0 ||
	// 	len(addItem.AllSync.MissionFinishSubList) != 0 ||
	// 	len(addItem.AllSync.MissionProgressSubList) != 0 {
	// 	g.AllPlayerSyncScNotify(addItem.AllSync)
	// 	g.AllScenePlaneEventScNotify(addItem.PileItem)
	// 	g.InspectMission(nil)
	// }
}

/**********************************任务方法********************************/

// 任务分类TYPE
