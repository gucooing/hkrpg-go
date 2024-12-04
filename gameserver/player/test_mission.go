package player

import (
	"fmt"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/suppl/pushc"
)

func (g *GamePlayer) InspectMission(finishSubMission ...uint32) {
	if g.GetPd().GetBasicBin().IsJumpMission {
		return
	}
	addItem := model.NewAddItem(nil)
	var finishSubList []uint32
	for _, v := range finishSubMission {
		finishSubList = append(finishSubList, v)
	}
	g.FinishSubMission(finishSubList, addItem)
	// 进入检查循环
	for {
		// 已接取的子任务检查
		finish, progress := g.FinishServerSubMission()
		g.FinishSubMission(finish, addItem)
		// 已接取的主任务检查
		finishMainList, progress2 := g.FinishServerMainMission()
		g.FinishMainMission(finishMainList, addItem)
		// 接取主任务
		mainList := g.GetPd().AcceptMainMission()
		g.GetPd().AddMainMission(mainList)

		addItem.AllSync.MissionFinishMainList = append(
			addItem.AllSync.MissionFinishMainList, finishMainList...)
		addItem.AllSync.MissionFinishSubList = append(
			addItem.AllSync.MissionFinishSubList, finish...)
		addItem.AllSync.MissionProgressSubList = append(
			addItem.AllSync.MissionProgressSubList, progress...)
		addItem.AllSync.MissionProgressSubList = append(
			addItem.AllSync.MissionProgressSubList, progress2...)
		if len(finish) == 0 &&
			len(progress) == 0 &&
			len(finishMainList) == 0 &&
			len(progress2) == 0 &&
			len(mainList) == 0 {
			break
		}
	}
	g.GetPd().AddItem(addItem)
	// 检查场景卸加载
	uninstallGroup, loadedGroupList := g.GetPd().AutoEntryGroup()
	g.UpSceneGroupRefreshScNotify(uninstallGroup, loadedGroupList)
	// 进入由任务解锁检查

	// 命途解锁检查
	g.GetPd().CheckUnlockMultiPath(addItem.AllSync)
	if raidId, ok := g.GetPd().CheckRaid(); ok {
		g.RaidInfoNotify(raidId) // raid完成检查
	}

	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.AllScenePlaneEventScNotify(addItem.PileItem)
}

// FinishSubMission 完成子任务
func (g *GamePlayer) FinishSubMission(finishSubList []uint32, addItem *model.AddItem) {
	for _, v := range finishSubList {
		// 执行完成动作
		g.AutoServerMissionFinishAction(v, addItem)
		// 通知
		g.StartFinishSubMissionScNotify(v)
	}
	// 写入数据库
	g.GetPd().AddFinishSubMission(finishSubList, addItem)
}

// FinishMainMission 完成主任务
func (g *GamePlayer) FinishMainMission(finishMainList []uint32, addItem *model.AddItem) {
	// 通知
	for _, finishId := range finishMainList {
		g.StartFinishMainMissionScNotify(finishId)
	}
	g.GetPd().AddFinishMainMission(finishMainList, addItem)
}

/*********************************完成检查**********************************/

var JumpMainMission = map[uint32]bool{
	4030001: true,
	4030002: true,
}

// FinishServerMainMission 已接取的主任务检查
func (g *GamePlayer) FinishServerMainMission() ([]uint32, []uint32) {
	mainMissionList := g.GetPd().GetMainMissionList()               // 已接取的主任务
	finishSubMissionList := g.GetPd().GetFinishSubMainMissionList() // 已完成的子任务
	finishServerMainMissionList := make([]uint32, 0)
	progressSubMissionList := make([]uint32, 0)
	for _, mainMission := range mainMissionList {
		if JumpMainMission[mainMission.MissionId] {
			finishServerMainMissionList = append(finishServerMainMissionList, mainMission.MissionId)
			continue
		}
		iSFinishMain := true
		conf := gdconf.GetGoppMainMissionById(mainMission.MissionId)
		if conf != nil {
			for _, subMissionId := range conf.FinishSubMissionList { // 检查主线任务是否满足完成条件
				if finishSubMissionList[subMissionId] == nil {
					iSFinishMain = false
					break
				}
			}
			if iSFinishMain {
				finishServerMainMissionList = append(finishServerMainMissionList, mainMission.MissionId)
			} else { // 不需要完成此主任务，去检查主任务下是否有子任务需要接取
				for _, subConf := range conf.SubMissionList {
					if g.GetPd().GetFinishSubMainMissionById(subConf.ID) == nil &&
						g.GetPd().GetSubMainMissionById(subConf.ID) == nil {
						handle, ok := SubMissionTakeTypeList[subConf.TakeType]
						if !ok {
							pushc.PushServer(&constant.LogPush{
								PushMessage: constant.PushMessage{
									Tag: "Mission",
								},
								LogMsg: fmt.Sprintf("未知的子任务接取条件,MissionId:%v,接取条件:%s",
									subConf.ID, subConf.TakeType),
								LogLevel: constant.ERROR,
							})
							continue
						}
						switch handle(g, subConf) {
						case unknownMission:
						case progressMission:
							progressSubMissionList = append(progressSubMissionList, mainMission.MissionId)
						}
					}
				}
			}
		}
	}

	return finishServerMainMissionList, progressSubMissionList
}

// SubMissionTakeTypeFunc 子任务接取检查
type SubMissionTakeTypeFunc func(g *GamePlayer, conf *gdconf.SubMission) int

var SubMissionTakeTypeList = map[constant.MissionBeginType]SubMissionTakeTypeFunc{
	constant.MissionBeginTypeNil:           MissionBeginTypeNil,
	constant.MissionBeginTypeAuto:          MissionBeginTypeAuto,
	constant.MissionBeginTypeUnknown:       MissionBeginTypeUnknown,
	constant.MissionBeginTypeAnySequence:   MissionBeginTypeAnySequence,
	constant.MissionBeginTypeMultiSequence: MissionBeginTypeMultiSequence,
	constant.MissionBeginTypeCustomValue:   MissionBeginTypeCustomValue,
}

func MissionBeginTypeNil(g *GamePlayer, conf *gdconf.SubMission) int {
	return progressMission
}

func MissionBeginTypeAuto(g *GamePlayer, conf *gdconf.SubMission) int {
	return progressMission
}
func MissionBeginTypeUnknown(g *GamePlayer, conf *gdconf.SubMission) int {
	return progressMission
}
func MissionBeginTypeAnySequence(g *GamePlayer, conf *gdconf.SubMission) int {
	for _, takeParamId := range conf.TakeParamIntList {
		if g.GetPd().GetFinishSubMainMissionById(takeParamId) != nil {
			return progressMission
		}
	}
	return unknownMission
}
func MissionBeginTypeMultiSequence(g *GamePlayer, conf *gdconf.SubMission) int {
	for _, takeParamId := range conf.TakeParamIntList {
		if g.GetPd().GetFinishSubMainMissionById(takeParamId) == nil {
			return unknownMission
		}
	}
	return progressMission
}
func MissionBeginTypeCustomValue(g *GamePlayer, conf *gdconf.SubMission) int {
	if g.GetPd().MissionCustomValue(conf.ID) {
		return progressMission
	}
	return unknownMission
}

// FinishServerSubMission 已接取的子任务检查
func (g *GamePlayer) FinishServerSubMission() ([]uint32, []uint32) {
	subMissionList := g.GetPd().GetSubMainMissionList() // 已接取的子任务
	finishServerSubMissionList := make([]uint32, 0)
	progressSubMissionList := make([]uint32, 0)
	for id, _ := range subMissionList {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil ||
			JumpSubMission[id] {
			finishServerSubMissionList = append(finishServerSubMissionList, id)
			logger.Debug("Mission:%v Conf Error", id)
			continue
		}
		handle, ok := SubMissionFinishType[conf.FinishType]
		if !ok {
			continue
		}
		switch handle(g, conf) {
		case unknownMission:
		case progressMission:
			progressSubMissionList = append(progressSubMissionList, id)
		case finishMission:
			finishServerSubMissionList = append(finishServerSubMissionList, id)
		}
	}

	return finishServerSubMissionList, progressSubMissionList
}

// 服务端完成子任务

var JumpSubMission = map[uint32]bool{}

type SubMissionFinishFunc func(g *GamePlayer, conf *gdconf.SubMission) int

const (
	unknownMission = iota
	progressMission
	finishMission
)

var SubMissionFinishType = map[constant.QuestFinishType]SubMissionFinishFunc{
	constant.Talk:                          Break,
	constant.Unknown:                       Unknown,                     // 直接完成
	constant.PropState:                     MissionPropState,            // 物品状态
	constant.GetTrialAvatar:                GetTrialAvatar,              // 加载试用角色
	constant.DelTrialAvatar:                DelTrialAvatar,              // 卸载试用角色
	constant.EnterFloor:                    EnterFloor,                  // 传送
	constant.EnterRaidScene:                EnterRaidScene,              // raid传送
	constant.SubMissionFinishCnt:           SubMissionFinishCnt,         // 完成列表中的子任务即可
	constant.FinishMission:                 FinishMission,               // 完成列表中的主任务即可
	constant.RaidFinishCnt:                 RaidFinishCnt,               // 完成raid
	constant.MessagePerformSectionFinish:   MessagePerformSectionFinish, // 发送对话框
	constant.MessageSectionFinish:          MessageSectionFinish,        // 发送消息
	constant.UseSelectedItem:               UseSelectedItem,             // 使用消耗品
	constant.AetherDivideCollectSpiritType: AetherDivideCollectSpiritType,
}

func Break(g *GamePlayer, conf *gdconf.SubMission) int {
	return unknownMission
}

func Unknown(g *GamePlayer, conf *gdconf.SubMission) int {
	return finishMission
}

func MissionPropState(g *GamePlayer, conf *gdconf.SubMission) int {
	db := g.GetPd().GetBlock(g.GetPd().GetCurEntryId())
	if g.GetPd().GetPropState(db, conf.ParamInt1, conf.ParamInt2, "") == conf.ParamInt3 {
		return finishMission
	}
	return unknownMission
}

func GetTrialAvatar(g *GamePlayer, conf *gdconf.SubMission) int {
	lineAvatar := g.GetPd().AddCurLineUpAvatar(conf.ParamInt1)
	g.AddAvatarSceneGroupRefreshScNotify(lineAvatar, g.GetPd().GetPosPb(), g.GetPd().GetRotPb())
	g.SyncLineupNotify(g.GetPd().GetCurLineUp())

	return finishMission
}

func DelTrialAvatar(g *GamePlayer, conf *gdconf.SubMission) int {
	g.GetPd().DelCurLineUpAvatar(conf.ParamInt1)
	g.GetPd().GetAddAvatarSceneEntityRefreshInfo(
		g.GetPd().GetCurLineUp(), g.GetPd().GetPosPb(), g.GetPd().GetRotPb())
	g.SyncLineupNotify(g.GetPd().GetCurLineUp())

	return finishMission
}

func EnterFloor(g *GamePlayer, conf *gdconf.SubMission) int {
	if entryId, groupID, anchorID, ok := gdconf.GetEntryId(conf.ID); ok {
		g.GetPd().SetCurEntryId(entryId)
		g.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)

		return finishMission
	}
	logger.Error("EnterFloor MissionId:%v error", conf.ID)
	return unknownMission
}

func EnterRaidScene(g *GamePlayer, conf *gdconf.SubMission) int {
	g.RaidEnterSceneByServerScNotify(conf.ParamInt2)

	return finishMission
}

func SubMissionFinishCnt(g *GamePlayer, conf *gdconf.SubMission) int {
	db := g.GetPd().GetSubMainMissionList()[conf.ID]
	finishSubMissionList := g.GetPd().GetFinishSubMainMissionList()
	if db == nil {
		return unknownMission
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
		return finishMission
	} else {
		if OldProgress != db.Progress {
			return progressMission
		}
	}
	return unknownMission
}

func FinishMission(g *GamePlayer, conf *gdconf.SubMission) int {
	db := g.GetPd().GetSubMainMissionList()[conf.ID]
	finishMainMissionList := g.GetPd().GetFinishMainMissionList()
	if db == nil {
		return unknownMission
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
		return finishMission
	} else {
		if OldProgress != db.Progress {
			return progressMission
		}
	}
	return unknownMission
}

func RaidFinishCnt(g *GamePlayer, conf *gdconf.SubMission) int {
	var ifFinish = true
	for _, raid := range conf.ParamIntList {
		if g.GetPd().GetFinishRaidInfo(raid) == nil {
			ifFinish = false
			break
		}
	}
	if ifFinish {
		return finishMission
	}
	return unknownMission
}

func MessagePerformSectionFinish(g *GamePlayer, conf *gdconf.SubMission) int {
	contactId := g.GetPd().AddMessageGroup(conf.ParamInt1)
	g.MessageGroupPlayerSyncScNotify(contactId)
	return progressMission
}

func MessageSectionFinish(g *GamePlayer, conf *gdconf.SubMission) int {
	contactId := g.GetPd().AddMessageGroup(conf.ParamInt1)
	g.MessageGroupPlayerSyncScNotify(contactId)
	return progressMission
}

func UseSelectedItem(g *GamePlayer, conf *gdconf.SubMission) int {
	return finishMission
}

func AetherDivideCollectSpiritType(g *GamePlayer, conf *gdconf.SubMission) int {
	return finishMission
}

// FinishActionTypeFunc 子任务完成自动执行
type FinishActionTypeFunc func(this *FinishActionType)

type FinishActionType struct {
	*GamePlayer
	finishAction *gdconf.FinishAction
	conf         *gdconf.SubMission
	addItem      *model.AddItem
}

var FinishActionTypeMap = map[constant.FinishActionType]FinishActionTypeFunc{
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

// AutoServerMissionFinishAction 完成任务后完成服务端动作（不结束任务
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
		handle, ok := FinishActionTypeMap[finishAction.FinishActionType]
		if !ok {
			pushc.PushServer(&constant.LogPush{
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

func FinishActionChangeLineup(this *FinishActionType) {
	this.GetPd().NewLineByAvatarList(this.finishAction.FinishActionPara) // 设置队伍角色
}

func FinishActionRecover(this *FinishActionType) {
	this.RecoverLine()
}

func FinishActionAddMissionItem(this *FinishActionType) {
	for index, item := range this.finishAction.FinishActionPara {
		if len(this.finishAction.FinishActionPara) < index+2 && index%2 != 0 {
			continue
		}
		this.addItem.PileItem = append(this.addItem.PileItem, &model.Material{
			Tid: item,
			Num: this.finishAction.FinishActionPara[index+1],
		})
	}
}

func FinishActionAddRecoverMissionItem(this *FinishActionType) {
	for index, item := range this.finishAction.FinishActionPara {
		if len(this.finishAction.FinishActionPara) < index+2 && index%2 != 0 {
			continue
		}
		this.addItem.PileItem = append(this.addItem.PileItem, &model.Material{
			Tid: item,
			Num: this.finishAction.FinishActionPara[index+1],
		})
	}
}

func FinishActionDelMissionItem(this *FinishActionType) {
}

func FinishActionDelMission(this *FinishActionType) {
	this.InspectMission(this.finishAction.FinishActionPara...)
}

func FinishActionDisableMission(this *FinishActionType) {
	this.GetPd().AddFinishMainMission(this.finishAction.FinishActionPara, this.addItem)
	this.InspectMission()
}

func FinishActionDelSubMission(this *FinishActionType) {
	this.InspectMission(this.finishAction.FinishActionPara...)
}

func FinishActionEnterEntryIfNotThere(this *FinishActionType) {
	if len(this.finishAction.FinishActionPara) < 3 {
		return
	}
	entryId := this.finishAction.FinishActionPara[0]
	groupID := this.finishAction.FinishActionPara[1]
	anchorID := this.finishAction.FinishActionPara[2]
	this.GetPd().SetCurEntryId(entryId)
	this.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
}

func FinishActionSetFloorSavedValue(this *FinishActionType) {
	this.SetFloorSavedValue(this.conf, this.finishAction)
}

func FinishActionMoveToAnchor(this *FinishActionType) {
	if len(this.finishAction.FinishActionPara) < 3 {
		return
	}
	entryId := this.finishAction.FinishActionPara[0]
	groupID := this.finishAction.FinishActionPara[1]
	anchorID := this.finishAction.FinishActionPara[2]
	this.GetPd().SetCurEntryId(entryId)
	this.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
}

func FinishActionSetGroupState(this *FinishActionType) {
	groupID := this.finishAction.FinishActionPara[0]
	groupState := this.finishAction.FinishActionPara[1]
	this.GetPd().SetGroupState(this.GetPd().GetBlock(model.FloorTentry(this.conf.LevelFloorID)), groupID, groupState)
}

func FinishActionFATChangeStoryLine(this *FinishActionType) {
	entryId, anchorGroup, anchorId, ok := this.GetPd().MissionAddChangeStoryLine(this.finishAction.FinishActionPara)
	if ok {
		this.EnterSceneByServerScNotify(entryId, 0, anchorGroup, anchorId)
	}
}
