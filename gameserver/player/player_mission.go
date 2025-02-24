package player

import (
	"fmt"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/push/client"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func init() {
	MissionFinishType = map[constant.QuestFinishType]func(this *QuestFinishType){
		constant.Talk:                          Break,
		constant.Unknown:                       Unknown,                       // 直接完成
		constant.PropState:                     MissionPropState,              // 物品状态
		constant.GetTrialAvatar:                GetTrialAvatar,                // 加载试用角色
		constant.DelTrialAvatar:                DelTrialAvatar,                // 卸载试用角色
		constant.EnterFloor:                    EnterFloor,                    // 传送
		constant.EnterRaidScene:                EnterRaidScene,                // raid传送
		constant.SubMissionFinishCnt:           SubMissionFinishCnt,           // 完成列表中的子任务即可
		constant.FinishMission:                 FinishMission,                 // 完成列表中的主任务即可
		constant.RaidFinishCnt:                 RaidFinishCnt,                 // 完成raid
		constant.MessagePerformSectionFinish:   MessagePerformSectionFinish,   // 发送对话框
		constant.MessageSectionFinish:          MessageSectionFinish,          // 发送消息
		constant.UseSelectedItem:               UseSelectedItem,               // 使用消耗品
		constant.AetherDivideCollectSpiritType: AetherDivideCollectSpiritType, // 以太战线获得新角色
	}

	FinishActionMap = map[constant.FinishActionType]func(this *FinishActionType) bool{
		constant.ChangeLineup:          FinishActionChangeLineup,          // 强制更新队伍
		constant.Recover:               FinishActionRecover,               // 恢复队伍
		constant.AddMissionItem:        FinishActionAddMissionItem,        // 添加任务道具
		constant.AddRecoverMissionItem: FinishActionAddRecoverMissionItem, // 添加任务恢复道具
		constant.DelMissionItem:        FinishActionDelMissionItem,        // 删除任务道具
		constant.DelMission:            FinishActionDelMission,            // 结束任务
		constant.DisableMission:        FinishActionDisableMission,        // 删除主线任务
		constant.DelSubMission:         FinishActionDelSubMission,         // 删除子任务
		constant.EnterEntryIfNotThere:  FinishActionEnterEntryIfNotThere,  // 传送到目标场景
		constant.SetFloorSavedValue:    FinishActionSetFloorSavedValue,    // 设置物品状态
		constant.MoveToAnchor:          FinishActionMoveToAnchor,          // 移动到锚点
		constant.SetGroupState:         FinishActionSetGroupState,         // 设置组状态
		constant.FATChangeStoryLine:    FinishActionFATChangeStoryLine,    // 强制添加并开启故事线
	}
}

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
	finishSubMission := g.GetPd().TalkStrSubMission(req) // 获取子任务
	if len(finishSubMission) != 0 {
		if len(finishSubMission) > 0 {
			g.InspectMission(finishSubMission...)
		}
	}
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
func (g *GamePlayer) InspectMission(finishSubMission ...uint32) {
	if g.GetPd().GetBasicBin().IsJumpMission {
		return
	}
	addItem := model.NewAddItem(nil)

	g.GetPd().AddFinishSubMission(finishSubMission, addItem)
	finishMainList := make([]uint32, 0)
	newFinishSubList := make([]uint32, 0)
	newProgressSubList := make([]uint32, 0)
	newFinishSubList = append(newFinishSubList, finishSubMission...)
	for {
		// 接取检查
		mainList, subList := g.AcceptInspectMission()
		newProgressSubList = append(newProgressSubList, subList...) // 将接取的任务添加到同步列表
		// 完成检查
		finishList, finishSubList, progressSubList := g.FinishInspectMission(addItem)
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
			g.AutoServerMissionFinishAction(finishId, addItem)
			g.Send(cmd.StartFinishSubMissionScNotify,
				&proto.StartFinishSubMissionScNotify{SubMissionId: finishId})
			if newProgressSubList[i] == finishId {
				newProgressSubList = append(newProgressSubList[:i], newProgressSubList[i+1:]...)
				break
			}
		}
	}

	g.GetPd().AddItem(addItem)
	addItem.AllSync.MissionFinishMainList = finishMainList
	addItem.AllSync.MissionFinishSubList = newFinishSubList
	addItem.AllSync.MissionProgressSubList = newProgressSubList
	// 检查场景卸加载
	uninstallGroup, loadedGroupList := g.GetPd().AutoEntryGroup()
	g.UpSceneGroupRefreshScNotify(uninstallGroup, loadedGroupList)
	// 命途解锁检查
	g.GetPd().CheckUnlockMultiPath(addItem.AllSync)
	if raidId, ok := g.GetPd().CheckRaid(); ok {
		g.RaidInfoNotify(raidId) // raid完成检查
	}
	if len(addItem.AllSync.MaterialList) != 0 ||
		len(addItem.AllSync.MissionFinishMainList) != 0 ||
		len(addItem.AllSync.MissionFinishSubList) != 0 ||
		len(addItem.AllSync.MissionProgressSubList) != 0 {
		g.AllPlayerSyncScNotify(addItem.AllSync)
		g.AllScenePlaneEventScNotify(addItem.PileItem)
		g.InspectMission()
	}
}

func (g *GamePlayer) AcceptInspectMission() ([]uint32, []uint32) {
	mainList := g.GetPd().AcceptMainMission() // 接取主任务
	g.GetPd().AddMainMission(mainList)
	subList := g.GetPd().AcceptSubMission() // 接取子任务
	g.GetPd().AddSubMission(subList)
	g.MissionAcceptScNotify(subList)

	return mainList, subList
}

func (g *GamePlayer) FinishInspectMission(addItem *model.AddItem) ([]uint32, []uint32, []uint32) {
	finishSubList, progressSubList := g.FinishServerSubMission()
	g.GetPd().AddFinishSubMission(finishSubList, addItem)
	// 主任务完成检查
	finishMainList := g.GetPd().FinishServerMainMission()
	g.GetPd().AddFinishMainMission(finishMainList, addItem)
	// 完成主任务中的未完成子任务
	finishSubLists := g.GetPd().CheckMainMission(finishMainList)
	g.GetPd().AddFinishSubMission(finishSubLists, addItem)
	finishSubList = append(finishSubList, finishSubLists...)

	return finishMainList, finishSubList, progressSubList
}

// 子任务完成检查
func (g *GamePlayer) FinishServerSubMission() ([]uint32, []uint32) {
	this := &QuestFinishType{
		GamePlayer: g,
		conf:       nil,
		finishMap:  make(map[uint32]int),
		notifyMap:  make(map[uint32]int),
	}
	subMissionList := g.GetPd().GetSubMainMissionList() // 已接取的子任务
	finishServerSubMissionList := make([]uint32, 0)
	progressSubMissionList := make([]uint32, 0)
	for id, _ := range subMissionList {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			this.finishMap[id] = 114514 // 直接完成掉
			logger.Debug("Mission:%v Conf Error", id)
			continue
		}
		handle, ok := MissionFinishType[conf.FinishType]
		if !ok {
			// client.PushServer(&constant.LogPush{
			// 	PushMessage: constant.PushMessage{
			// 		Tag: "Mission",
			// 	},
			// 	LogMsg: fmt.Sprintf("未知的任务完成条件,MissionId:%v,完成条件:%s",
			// 		conf.ID, conf.FinishType),
			// 	LogLevel: constant.ERROR,
			// })
			continue
		}
		this.conf = conf
		handle(this)
	}
	for id, info := range this.finishMap {
		if info == 114514 {
			finishServerSubMissionList = append(finishServerSubMissionList, id)
		}
	}
	for id, info := range this.notifyMap {
		if info == 114514 {
			progressSubMissionList = append(progressSubMissionList, id)
		}
	}

	return finishServerSubMissionList, progressSubMissionList
}

// 完成任务后完成服务端动作（不结束任务
func (g *GamePlayer) AutoServerMissionFinishAction(id uint32, addItem *model.AddItem) {
	conf := gdconf.GetSubMainMissionById(id)
	addItem = model.NewAddItem(addItem)
	if conf == nil || conf.FinishActionList == nil {
		return
	}
	this := &FinishActionType{
		GamePlayer:   g,
		conf:         conf,
		addItem:      addItem,
		finishAction: nil,
	}
	for _, finishAction := range conf.FinishActionList {
		handle, ok := FinishActionMap[finishAction.FinishActionType]
		if !ok {
			client.PushServer(&constant.LogPush{
				PushMessage: constant.PushMessage{
					Tag: "Mission",
				},
				LogMsg: fmt.Sprintf("未知的任务自动执行,MissionId:%v,自动执行:%s",
					conf.ID, finishAction.FinishActionType),
				LogLevel: constant.ERROR,
			})
			continue
		}
		this.finishAction = finishAction
		handle(this)
	}
}

/**********************************任务方法********************************/

// 子任务完成检查
var MissionFinishType map[constant.QuestFinishType]func(this *QuestFinishType)

type QuestFinishType struct {
	*GamePlayer
	conf      *gdconf.SubMission // 任务配置
	finishMap map[uint32]int     // 完成列表
	notifyMap map[uint32]int     // 更改列表
}

func (this *QuestFinishType) addFinishMap() {
	if this.finishMap == nil {
		this.finishMap = make(map[uint32]int)
	}
	if this.notifyMap[this.conf.ID] != 0 {
		delete(this.notifyMap, this.conf.ID)
	}
	this.finishMap[this.conf.ID] = 114514
}

func (this *QuestFinishType) addNotifyMap() {
	if this.notifyMap == nil {
		this.notifyMap = make(map[uint32]int)
	}
	if this.finishMap[this.conf.ID] == 0 {
		this.notifyMap[this.conf.ID] = 114514
	}
}

func Break(this *QuestFinishType) {}

func Unknown(this *QuestFinishType) {
	this.addNotifyMap()
}

func GetTrialAvatar(this *QuestFinishType) {
	lineAvatar := this.GetPd().AddCurLineUpAvatar(this.conf.ParamInt1)
	this.AddAvatarSceneGroupRefreshScNotify(lineAvatar, this.GetPd().GetPosPb(), this.GetPd().GetRotPb())
	this.SyncLineupNotify(this.GetPd().GetCurLineUp())

	this.addFinishMap()
}

func DelTrialAvatar(this *QuestFinishType) {
	this.GetPd().DelCurLineUpAvatar(this.conf.ParamInt1)
	this.GetPd().GetAddAvatarSceneEntityRefreshInfo(
		this.GetPd().GetCurLineUp(), this.GetPd().GetPosPb(), this.GetPd().GetRotPb())
	this.SyncLineupNotify(this.GetPd().GetCurLineUp())

	this.addFinishMap()
}

func EnterFloor(this *QuestFinishType) {
	if entryId, groupID, anchorID, ok := gdconf.GetEntryId(this.conf.ID); ok {
		this.GetPd().SetCurEntryId(entryId)
		this.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)

		this.addFinishMap()
	} else {
		logger.Error("EnterFloor MissionId:%v error", this.conf.ID)
	}
}

func EnterRaidScene(this *QuestFinishType) {
	this.RaidEnterSceneByServerScNotify(this.conf.ParamInt2)

	this.addFinishMap()
}

func SubMissionFinishCnt(this *QuestFinishType) {
	db := this.GetPd().GetSubMainMissionList()[this.conf.ID]
	finishSubMissionList := this.GetPd().GetFinishSubMainMissionList()
	conf := gdconf.GetSubMainMissionById(this.conf.ID)
	if conf == nil || db == nil {
		return
	}
	OldProgress := db.Progress
	db.Progress = 0
	isFinish := true
	for _, paramInt := range conf.ParamIntList {
		if finishSubMissionList[paramInt] != nil {
			db.Progress++
		} else {
			isFinish = false
		}
	}
	if db.Progress == conf.Progress {
		isFinish = true
	}
	if isFinish { // 完成任务
		db.Progress = conf.Progress

		this.addFinishMap()
	} else {
		if OldProgress != db.Progress {
			this.addFinishMap()
		}
	}
}

func FinishMission(this *QuestFinishType) {
	db := this.GetPd().GetSubMainMissionList()[this.conf.ID]
	finishMainMissionList := this.GetPd().GetFinishMainMissionList()
	conf := gdconf.GetSubMainMissionById(this.conf.ID)
	if conf == nil || db == nil {
		return
	}
	OldProgress := db.Progress
	db.Progress = 0
	isFinish := true
	for _, paramInt := range conf.ParamIntList {
		if finishMainMissionList[paramInt] != nil {
			db.Progress++
		} else {
			isFinish = false
		}
	}
	if isFinish { // 完成任务
		db.Progress = conf.Progress

		this.addFinishMap()
	} else {
		if OldProgress != db.Progress {
			this.addFinishMap()
		}
	}
}

func RaidFinishCnt(this *QuestFinishType) {
	var ifFinish = true
	for _, raid := range this.conf.ParamIntList {
		if this.GetPd().GetFinishRaidInfo(raid) == nil {
			ifFinish = false
			break
		}
	}
	if ifFinish {
		this.addFinishMap()
	}
}

func MessagePerformSectionFinish(this *QuestFinishType) {
	contactId := this.GetPd().AddMessageGroup(this.conf.ParamInt1)
	this.MessageGroupPlayerSyncScNotify(contactId)
}

func MessageSectionFinish(this *QuestFinishType) {
	contactId := this.GetPd().AddMessageGroup(this.conf.ParamInt1)
	this.MessageGroupPlayerSyncScNotify(contactId)
}

func MissionPropState(this *QuestFinishType) {
	db := this.GetPd().GetBlock(this.GetPd().GetCurEntryId())
	conf := gdconf.GetSubMainMissionById(this.conf.ID)
	if conf == nil || db == nil {
		return // 不存在
	}
	if this.GetPd().GetPropState(db, conf.ParamInt1, conf.ParamInt2, "") == conf.ParamInt3 {
		this.addFinishMap()
	}
}

func UseSelectedItem(this *QuestFinishType) {
	this.addFinishMap()
}

func AetherDivideCollectSpiritType(this *QuestFinishType) {
	this.addFinishMap()
}

// 完成任务后完成服务端动作（不结束任务

var FinishActionMap map[constant.FinishActionType]func(this *FinishActionType) bool

type FinishActionType struct {
	*GamePlayer
	finishAction *gdconf.FinishAction
	conf         *gdconf.SubMission
	addItem      *model.AddItem
}

func FinishActionChangeLineup(this *FinishActionType) bool {
	this.GetPd().NewLineByAvatarList(this.finishAction.FinishActionPara) // 设置队伍角色
	return true
}

func FinishActionRecover(this *FinishActionType) bool {
	this.RecoverLine()
	return true
}

func FinishActionAddMissionItem(this *FinishActionType) bool {
	for index, item := range this.finishAction.FinishActionPara {
		if len(this.finishAction.FinishActionPara) < index+2 && index%2 != 0 {
			continue
		}
		this.addItem.PileItem = append(this.addItem.PileItem, &model.Material{
			Tid: item,
			Num: this.finishAction.FinishActionPara[index+1],
		})
	}
	return true
}

func FinishActionAddRecoverMissionItem(this *FinishActionType) bool {
	for index, item := range this.finishAction.FinishActionPara {
		if len(this.finishAction.FinishActionPara) < index+2 && index%2 != 0 {
			continue
		}
		this.addItem.PileItem = append(this.addItem.PileItem, &model.Material{
			Tid: item,
			Num: this.finishAction.FinishActionPara[index+1],
		})
	}
	return true
}

func FinishActionDelMissionItem(this *FinishActionType) bool {
	return true
}

func FinishActionDelMission(this *FinishActionType) bool {
	this.InspectMission(this.finishAction.FinishActionPara...)
	return true
}

func FinishActionDisableMission(this *FinishActionType) bool {
	this.GetPd().AddFinishMainMission(this.finishAction.FinishActionPara, this.addItem)
	this.InspectMission()
	return true
}

func FinishActionDelSubMission(this *FinishActionType) bool {
	this.InspectMission(this.finishAction.FinishActionPara...)
	return true
}

func FinishActionEnterEntryIfNotThere(this *FinishActionType) bool {
	if len(this.finishAction.FinishActionPara) < 3 {
		return false
	}
	entryId := this.finishAction.FinishActionPara[0]
	groupID := this.finishAction.FinishActionPara[1]
	anchorID := this.finishAction.FinishActionPara[2]
	this.GetPd().SetCurEntryId(entryId)
	this.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
	return true
}

func FinishActionSetFloorSavedValue(this *FinishActionType) bool {
	this.SetFloorSavedValue(this.conf, this.finishAction)
	return true
}

func FinishActionMoveToAnchor(this *FinishActionType) bool {
	if len(this.finishAction.FinishActionPara) < 3 {
		return false
	}
	entryId := this.finishAction.FinishActionPara[0]
	groupID := this.finishAction.FinishActionPara[1]
	anchorID := this.finishAction.FinishActionPara[2]
	this.GetPd().SetCurEntryId(entryId)
	this.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
	return true
}

func FinishActionSetGroupState(this *FinishActionType) bool {
	groupID := this.finishAction.FinishActionPara[0]
	groupState := this.finishAction.FinishActionPara[1]
	this.GetPd().SetGroupState(this.GetPd().GetBlock(model.FloorTentry(this.conf.LevelFloorID)), groupID, groupState)
	return true
}

func FinishActionFATChangeStoryLine(this *FinishActionType) bool {
	entryId, anchorGroup, anchorId, ok := this.GetPd().MissionAddChangeStoryLine(this.finishAction.FinishActionPara)
	if ok {
		this.EnterSceneByServerScNotify(entryId, 0, anchorGroup, anchorId)
	}
	return true
}

// 任务分类TYPE
