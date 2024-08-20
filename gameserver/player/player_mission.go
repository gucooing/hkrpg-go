package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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

func (g *GamePlayer) GetDailyActiveInfoCsReq(payloadMsg pb.Message) {
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

func (g *GamePlayer) GetMainMissionCustomValueCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetMainMissionCustomValueCsReq)
	rsp := &proto.GetMainMissionCustomValueScRsp{MainMissionList: make([]*proto.MainMission, 0)}
	mainMissionList := g.GetPd().GetMainMissionList()             // 已接取的主任务
	finishMainMissionList := g.GetPd().GetFinishMainMissionList() // 已完成的主任务
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

func (g *GamePlayer) UpdateTrackMainMissionIdCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UpdateTrackMainMissionIdCsReq)
	g.Send(cmd.UpdateTrackMainMissionIdScRsp, &proto.UpdateTrackMainMissionIdScRsp{TrackMissionId: req.TrackMissionId})
}

func (g *GamePlayer) GetMissionEventDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetMissionEventDataScRsp{
		Retcode:          0,
		ChallengeEventId: 0,
		MissionEventList: make([]*proto.Mission, 0),
	}
	for _, mission := range g.GetPd().GetMainMissionList() {
		rsp.MissionEventList = append(rsp.MissionEventList, &proto.Mission{
			Id:       mission.MissionId,
			Progress: mission.MissionId,
			Status:   proto.MissionStatus(mission.Status),
		})
	}
	g.Send(cmd.GetMissionEventDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetMissionStatusCsReq(payloadMsg pb.Message) {
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

func (g *GamePlayer) GetMissionDataCsReq(payloadMsg pb.Message) {
	mainMissionList := g.GetPd().GetMainMissionList()
	subMainMissionList := g.GetPd().GetSubMainMissionList()

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

func (g *GamePlayer) FinishTalkMissionCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishTalkMissionCsReq)
	finishSubMission := g.GetPd().TalkStrSubMission(req) // 获取子任务
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
	g.Send(cmd.FinishTalkMissionScRsp, &proto.FinishTalkMissionScRsp{TalkStr: req.TalkStr, CustomValueList: req.CustomValueList})
}

func (g *GamePlayer) FinishCosumeItemMissionCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishCosumeItemMissionCsReq)
	allSync := new(model.AllPlayerSync)
	if g.GetPd().FinishCosumeItemMission(req.SubMissionId, allSync) {
		g.InspectMission([]uint32{req.SubMissionId})
	}
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.FinishCosumeItemMissionScRsp, &proto.FinishCosumeItemMissionScRsp{SubMissionId: req.SubMissionId})
}

func (g *GamePlayer) MissionRewardScNotify() {

}

/*********************************检查操作*********************************/

// 登录任务检查
func (g *GamePlayer) LoginReadyMission() {
	if g.IsJumpMission {
		return
	}
	g.InspectMission(nil)
	g.AllCheckMainMission()
}

// 任务检查
func (g *GamePlayer) InspectMission(finishSubMission []uint32) {
	allSync := &model.AllPlayerSync{
		IsBasic:                true,
		MaterialList:           make([]uint32, 0),
		MissionFinishMainList:  make([]uint32, 0),
		MissionFinishSubList:   make([]uint32, 0),
		MissionProgressSubList: make([]uint32, 0),
	}
	var pileItem []*model.Material
	g.GetPd().AddFinishSubMission(finishSubMission, allSync, pileItem)
	finishMainList := make([]uint32, 0)
	newFinishSubList := make([]uint32, 0)
	newProgressSubList := make([]uint32, 0)
	newFinishSubList = append(newFinishSubList, finishSubMission...)
	for {
		// 接取检查
		mainList, subList := g.AcceptInspectMission()
		newProgressSubList = append(newProgressSubList, subList...) // 将接取的任务添加到同步列表
		// 完成检查
		finishList, finishSubList, progressSubList := g.FinishInspectMission(allSync, pileItem)
		finishMainList = append(finishMainList, finishList...)              // 将完成的任务添加到同步列表
		newFinishSubList = append(newFinishSubList, finishSubList...)       // 将完成的任务添加到同步列表
		newProgressSubList = append(newProgressSubList, progressSubList...) // 将接取的任务添加到同步列表

		if len(mainList) == 0 &&
			len(subList) == 0 &&
			len(finishList) == 0 &&
			len(finishSubList) == 0 &&
			len(progressSubList) == 0 {
			break
		}
	}

	for _, finishId := range finishMainList {
		g.Send(cmd.StartFinishMainMissionScNotify,
			&proto.StartFinishMainMissionScNotify{MainMissionId: finishId})
	}

	for i := len(newProgressSubList) - 1; i >= 0; i-- {
		for _, finishId := range newFinishSubList {
			g.AutoServerMissionFinishAction(finishId, allSync, pileItem)
			g.Send(cmd.StartFinishSubMissionScNotify,
				&proto.StartFinishSubMissionScNotify{SubMissionId: finishId})
			if newProgressSubList[i] == finishId {
				newProgressSubList = append(newProgressSubList[:i], newProgressSubList[i+1:]...)
				break
			}
		}
	}

	g.GetPd().AddItem(pileItem, allSync)
	allSync.MissionFinishMainList = finishMainList
	allSync.MissionFinishSubList = newFinishSubList
	allSync.MissionProgressSubList = newProgressSubList
	// 检查场景卸加载
	uninstallGroup, loadedGroupList := g.GetPd().AutoEntryGroup()
	g.UpSceneGroupRefreshScNotify(uninstallGroup, loadedGroupList)
	// 命途解锁检查
	g.GetPd().CheckUnlockMultiPath(allSync)
	if raidId, ok := g.GetPd().CheckRaid(); ok {
		g.RaidInfoNotify(raidId) // raid完成检查
	}
	if len(allSync.MaterialList) != 0 ||
		len(allSync.MissionFinishMainList) != 0 ||
		len(allSync.MissionFinishSubList) != 0 ||
		len(allSync.MissionProgressSubList) != 0 {
		g.AllPlayerSyncScNotify(allSync)
		g.AllScenePlaneEventScNotify(pileItem)
		g.InspectMission(nil)
	}
}

func (g *GamePlayer) AcceptInspectMission() ([]uint32, []uint32) {
	mainList := g.GetPd().AcceptMainMission() // 接取主任务
	g.GetPd().AddMainMission(mainList)
	subList := g.GetPd().AcceptSubMission() // 接取子任务
	g.GetPd().AddSubMission(subList)
	g.Send(cmd.MissionAcceptScNotify, &proto.MissionAcceptScNotify{SubMissionIdList: subList})

	return mainList, subList
}

func (g *GamePlayer) FinishInspectMission(allSync *model.AllPlayerSync, pileItem []*model.Material) ([]uint32, []uint32, []uint32) {
	finishSubList, progressSubList := g.FinishServerSubMission()
	g.GetPd().AddFinishSubMission(finishSubList, allSync, pileItem)
	// 主任务完成检查
	finishMainList := g.GetPd().FinishServerMainMission()
	g.GetPd().AddFinishMainMission(finishMainList, allSync, pileItem)
	// 完成主任务中的未完成子任务
	finishSubLists := g.CheckMainMission(finishMainList)
	g.GetPd().AddFinishSubMission(finishSubLists, allSync, pileItem)
	finishSubList = append(finishSubList, finishSubLists...)

	return finishMainList, finishSubList, progressSubList
}

// 子任务完成检查
func (g *GamePlayer) FinishServerSubMission() ([]uint32, []uint32) {
	subMissionList := g.GetPd().GetSubMainMissionList() // 已接取的子任务
	// finishSubMissionList := g.GetFinishSubMainMissionList() // 已完成的子任务
	finishServerSubMissionList := make([]uint32, 0)
	progressSubMissionList := make([]uint32, 0)
	for id, _ := range subMissionList {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		ifFinish := false
		switch conf.FinishType {
		case constant.GetTrialAvatar: // 加载试用角色
			lineAvatar := g.GetPd().GetTrialAvatar(conf.ParamInt1)
			g.AddAvatarSceneGroupRefreshScNotify(lineAvatar, g.GetPd().GetPosPb(), g.GetPd().GetRotPb())
			g.SyncLineupNotify(g.GetPd().GetBattleLineUp())
			ifFinish = true
			break
		case constant.DelTrialAvatar: // 卸载试用角色
			g.GetPd().DelTrialAvatar(conf.ParamInt1)
			g.GetPd().GetAddAvatarSceneEntityRefreshInfo(
				g.GetPd().GetBattleLineUp(), g.GetPd().GetPosPb(), g.GetPd().GetRotPb())
			g.SyncLineupNotify(g.GetPd().GetBattleLineUp())
			ifFinish = true
			break
		case constant.EnterFloor: // 传送
			if entryId, groupID, anchorID, ok := gdconf.GetEntryId(id); ok {
				g.GetPd().SetCurEntryId(entryId)
				g.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
				ifFinish = true
			} else {
				logger.Error("EnterFloor MissionId:%v error", id)
			}
			break
		case constant.EnterRaidScene: // raid传送
			g.RaidEnterSceneByServerScNotify(conf.ParamInt2)
			ifFinish = true
			break
		case constant.SubMissionFinishCnt: // 完成列表中的子任务即可
			finish, progress := g.GetPd().SubMissionFinishCnt(id)
			if finish != 0 {
				finishServerSubMissionList = append(finishServerSubMissionList, finish)
			}
			if progress != 0 {
				progressSubMissionList = append(progressSubMissionList, progress)
			}
			break
		case constant.FinishMission: // 完成列表中的主任务即可
			finish, progress := g.GetPd().FinishMainMission(id)
			if finish != 0 {
				finishServerSubMissionList = append(finishServerSubMissionList, finish)
			}
			if progress != 0 {
				progressSubMissionList = append(progressSubMissionList, progress)
			}
			break
		case constant.RaidFinishCnt: // 完成raid
			ifFinish = true
			for _, raid := range conf.ParamIntList {
				if g.GetPd().GetFinishRaidInfo(raid) == nil {
					ifFinish = false
					break
				}
			}
		case constant.MessagePerformSectionFinish: // 发送对话框
			contactId := g.GetPd().AddMessageGroup(conf.ParamInt1)
			g.MessageGroupPlayerSyncScNotify(contactId)
			break
		case constant.MessageSectionFinish: // 发送消息
			contactId := g.GetPd().AddMessageGroup(conf.ParamInt1)
			g.MessageGroupPlayerSyncScNotify(contactId)
			ifFinish = true
			break
		case constant.Unknown: // 直接完成
			ifFinish = true
			break
		case constant.PropState:
			ifFinish = g.GetPd().MissionPropState(id)
			break
		case constant.UseSelectedItem: // 使用消耗品
			ifFinish = true
			break
		}
		if ifFinish {
			finishServerSubMissionList = append(finishServerSubMissionList, id)
		}
	}

	return finishServerSubMissionList, progressSubMissionList
}

// 完成任务后完成服务端动作（不结束任务
func (g *GamePlayer) AutoServerMissionFinishAction(id uint32, allSync *model.AllPlayerSync, pileItem []*model.Material) {
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil {
		return
	}
	if conf.FinishActionList == nil {
		return
	}
	for _, finishAction := range conf.FinishActionList {
		switch finishAction.FinishActionType {
		case constant.ChangeLineup: // 强制更新队伍
			g.GetPd().NewTrialLine(finishAction.FinishActionPara) // 设置队伍角色
		case constant.Recover: // 恢复队伍
			g.RecoverLine()
		case constant.AddMissionItem: // 添加任务道具
			for index, item := range finishAction.FinishActionPara {
				if len(finishAction.FinishActionPara) < index+2 && index%2 != 0 {
					continue
				}
				pileItem = append(pileItem, &model.Material{
					Tid: item,
					Num: finishAction.FinishActionPara[index+1],
				})
			}
		case constant.AddRecoverMissionItem: // 添加任务恢复道具
			for index, item := range finishAction.FinishActionPara {
				if len(finishAction.FinishActionPara) < index+2 && index%2 != 0 {
					continue
				}
				pileItem = append(pileItem, &model.Material{
					Tid: item,
					Num: finishAction.FinishActionPara[index+1],
				})
			}
		case constant.DelMissionItem: // 删除任务道具

		case constant.DelMission: // 结束任务
			g.InspectMission(finishAction.FinishActionPara)
		case constant.DisableMission: // 删除主线任务
			g.GetPd().AddFinishMainMission(finishAction.FinishActionPara, allSync, pileItem)
			g.InspectMission(nil)
		case constant.DelSubMission: // 删除子任务
			g.InspectMission(finishAction.FinishActionPara)
		case constant.EnterEntryIfNotThere: // 传送到目标场景
			if len(finishAction.FinishActionPara) < 3 {
				continue
			}
			entryId := finishAction.FinishActionPara[0]
			groupID := finishAction.FinishActionPara[1]
			anchorID := finishAction.FinishActionPara[2]
			g.GetPd().SetCurEntryId(entryId)
			g.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
		case constant.SetFloorSavedValue: // 设置物品状态
			g.SetFloorSavedValue(conf, finishAction)
		case constant.MoveToAnchor: // 移动到锚点
			if len(finishAction.FinishActionPara) < 3 {
				continue
			}
			entryId := finishAction.FinishActionPara[0]
			groupID := finishAction.FinishActionPara[1]
			anchorID := finishAction.FinishActionPara[2]
			g.GetPd().SetCurEntryId(entryId)
			g.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
		case constant.SetGroupState: // 设置组状态
			groupID := finishAction.FinishActionPara[0]
			groupState := finishAction.FinishActionPara[1]
			g.GetPd().SetGroupState(g.GetPd().GetBlock(model.FloorTentry(conf.LevelFloorID)), groupID, groupState)
		case constant.FATChangeStoryLine: // 强制添加并开启故事线
			entryId, anchorGroup, anchorId, ok := g.GetPd().MissionAddChangeStoryLine(finishAction.FinishActionPara)
			if ok {
				g.EnterSceneByServerScNotify(entryId, 0, anchorGroup, anchorId)
			}
		}
	}
}

/*********************************数据库操作**********************************/

// 将已完成的主任务下还没有完成的子任务全部完成
func (g *GamePlayer) CheckMainMission(finishMainList []uint32) []uint32 {
	finishSubList := g.GetPd().GetFinishSubMainMissionList()
	finishSubMission := make([]uint32, 0)
	for _, mainId := range finishMainList {
		conf := gdconf.GetGoppMainMissionById(mainId)
		if conf == nil {
			continue
		}
		for _, subInfo := range conf.SubMissionList {
			if finishSubList[subInfo.ID] == nil {
				finishSubMission = append(finishSubMission, subInfo.ID)
			}
		}
	}

	return finishSubMission
}

// 全局检查将已完成的主任务下还没有完成的子任务全部完成
func (g *GamePlayer) AllCheckMainMission() []uint32 {
	finishSubList := g.GetPd().GetFinishSubMainMissionList()
	finishSubMission := make([]uint32, 0)
	for mainId := range g.GetPd().GetFinishMainMissionList() {
		conf := gdconf.GetGoppMainMissionById(mainId)
		if conf == nil {
			continue
		}
		for _, subInfo := range conf.SubMissionList {
			if finishSubList[subInfo.ID] == nil {
				finishSubMission = append(finishSubMission, subInfo.ID)
			}
		}
	}

	return finishSubMission
}
