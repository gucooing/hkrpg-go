package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) SyncRogueStatusScNotify() {
	notify := &proto.SyncRogueStatusScNotify{}
	defer g.Send(cmd.SyncRogueStatusScNotify, notify)
	db := g.GetPd().GetCurRogue()
	if db == nil {
		notify.Status = proto.RogueStatus_ROGUE_STATUS_NONE
		return
	}
	notify.Status = proto.RogueStatus(db.Status)
}

func (g *GamePlayer) SyncRogueHandbookDataUpdateScNotify() {
	notify := &proto.SyncRogueHandbookDataUpdateScNotify{
		MagicUnitList:        make([]*proto.RogueMagicUnitInfo, 0),
		HandbookMiracleList:  g.GetPd().GetRogueHandbookMiracleInfoList(),
		HandbookEventList:    g.GetPd().GetRogueHandbookEventInfoList(),
		HandbookMazeBuffList: g.GetPd().GetRogueHandbookMazeBuffList(),
		MagicScepterList:     make([]*proto.RogueMagicScepterInfo, 0),
	}

	g.Send(cmd.SyncRogueHandbookDataUpdateScNotify, notify)
}

func GetRogueHandbookDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetRogueHandbookDataScRsp{
		HandbookInfo: &proto.RogueHandbook{
			MiracleList: g.GetPd().GetRogueHandbookMiracleInfoList(),
			BuffList:    g.GetPd().GetRogueHandbookMazeBuffList(),
			EventList:   g.GetPd().GetRogueHandbookEventInfoList(),
			AeonList:    g.GetPd().GetRogueHandbookAeonInfoList(),
		},
	}

	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func TakeRogueEventHandbookRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeRogueEventHandbookRewardCsReq)
	addItem := model.NewAddItem(nil)
	rsp := &proto.TakeRogueEventHandbookRewardScRsp{
		Reward:                  &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		Retcode:                 0,
		RewardHandbookEventList: make([]uint32, 0),
	}

	for _, eventId := range req.HandbookEventList {
		conf := gdconf.GetRogueHandBookEvent(eventId)
		if conf == nil {
			continue
		}
		if db := g.GetPd().GetRogueHandbookEvent(eventId); db != nil && !db.IsTakenReward {
			addItem.PileItem = append(addItem.PileItem, model.GetRewardData(conf.EventReward)...)
			db.IsTakenReward = true
			rsp.RewardHandbookEventList = append(rsp.RewardHandbookEventList, eventId)
		}
	}

	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	rsp.Reward.ItemList = addItem.ItemList
	g.Send(cmd.TakeRogueEventHandbookRewardScRsp, rsp)
}

func TakeRogueMiracleHandbookRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeRogueMiracleHandbookRewardCsReq)
	addItem := model.NewAddItem(nil)
	rsp := &proto.TakeRogueMiracleHandbookRewardScRsp{
		Reward:                    &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		Retcode:                   0,
		RewardHandbookMiracleList: make([]uint32, 0),
	}

	for _, miracleId := range req.HandbookMiracleList {
		conf := gdconf.GetRogueHandbookMiracle(miracleId)
		if conf == nil {
			continue
		}
		if db := g.GetPd().GetRogueHandbookMiracle(miracleId); db != nil && !db.IsTakenReward {
			addItem.PileItem = append(addItem.PileItem, model.GetRewardData(conf.MiracleReward)...)
			db.IsTakenReward = true
			rsp.RewardHandbookMiracleList = append(rsp.RewardHandbookMiracleList, miracleId)
		}
	}

	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	rsp.Reward.ItemList = addItem.ItemList
	g.Send(cmd.TakeRogueMiracleHandbookRewardScRsp, rsp)
}

func CommonRogueQueryCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.CommonRogueQueryScRsp{
		RogueGetInfo: &proto.OINKPDBJJOE{
			ONGHNJCGJLA: &proto.ADILAOJIMPO{
				HIPAMHDLBDB: []proto.RogueUnlockFunctionType{
					proto.RogueUnlockFunctionType_ROGUE_UNLOCK_FUNCTION_TYPE_MIRACLE,
					proto.RogueUnlockFunctionType_ROGUE_UNLOCK_FUNCTION_TYPE_SHOW_HINT,
					proto.RogueUnlockFunctionType_ROGUE_UNLOCK_FUNTION_TYPE_EXHIBITION,
					proto.RogueUnlockFunctionType_ROGUE_UNLOCK_FUNTION_TYPE_COLLECTION,
					proto.RogueUnlockFunctionType_ROGUE_UNLOCK_FUNCTION_TYPE_COSMOS_BAN_AEON,
				},
			},
			RogueAreaInfo: &proto.DEIPJCNOIBO{
				MPOOEPKBGCK: 501,
				PKDKPGCKKKI: 501,
				PDDMCKELMIJ: 163,
			},
		},
		GGCGPNABJGA: 0,
		Retcode:     0,
	}

	g.Send(cmd.CommonRogueQueryScRsp, rsp)
}

// 模拟宇宙自动化
func (g *GamePlayer) RogueAction() {
	curRogue := g.GetPd().GetCurRogue()
	if curRogue == nil {
		return
	}
	var rogueSubMode uint32 = 0
	switch x := curRogue.RogueInfo.(type) {
	case *spb.CurRogue_QuestRogue:
		rogueSubMode = constant.RogueTypeQuest
	default:
		logger.Error(text.GetText(100), x)
		return
	}
	// 添加 RogueAction
	g.SyncRogueCommonPendingActionScNotify(rogueSubMode, curRogue.Action)
}

func (g *GamePlayer) SyncRogueCommonPendingActionScNotify(rogueSubMode uint32, actionMap map[uint32]*spb.RogueAction) {
	for queuePosition, action := range actionMap {
		notify := &proto.SyncRogueCommonPendingActionScNotify{
			Action:       &proto.RogueCommonPendingAction{},
			RogueSubMode: rogueSubMode,
		}
		if a := g.GetPd().GetRogueCommonPendingAction(queuePosition, action); a != nil {
			notify.Action = a
			g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
		}
	}
}

func HandleRogueCommonPendingActionCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.HandleRogueCommonPendingActionCsReq)
	rsp := &proto.HandleRogueCommonPendingActionScRsp{
		QueuePosition: req.QueueLocation,
		Retcode:       0,
		QueueLocation: 0,
		Action:        nil,
	}

	switch x := req.Action.(type) {
	case *proto.HandleRogueCommonPendingActionCsReq_BonusSelectResult: // 选择开拓祝福
		g.HandleRogueActionBonusSelectResult(rsp, x, req.QueueLocation)
	case *proto.HandleRogueCommonPendingActionCsReq_BuffSelectResult: // 选择buff
		g.HandleRogueActionBuffSelectResult(rsp, x, req.QueueLocation)
	case *proto.HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult: // 刷新buff

	default:
		logger.Error(text.GetText(101), x)
	}

	g.RogueAction() // 模拟宇宙自动化

	g.Send(cmd.HandleRogueCommonPendingActionScRsp, rsp)
}

func (g *GamePlayer) HandleRogueActionBonusSelectResult(rsp *proto.HandleRogueCommonPendingActionScRsp,
	callback *proto.HandleRogueCommonPendingActionCsReq_BonusSelectResult, queueLocation uint32) {
	action := g.GetPd().GetRogueActionByQueuePosition(queueLocation)
	if action == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_ROGUE_SELECT_BUFF_NOT_EXIST)
		return
	}
	if a, ok := action.Action.(*spb.RogueAction_BonusSelect); ok {
		bonusId := callback.BonusSelectResult.BonusId
		if a.BonusSelect.BonusIdMap[bonusId] {
			// TODO 执行Bonus
			// c := gdconf.GetRogueBonus(bonusId)
			rsp.QueueLocation = queueLocation
			g.GetPd().DelRogueActionByQueuePosition(queueLocation)
		} else {
			rsp.Retcode = uint32(proto.Retcode_RET_ROGUE_SELECT_BUFF_NOT_EXIST)
			return
		}
	}
}

func (g *GamePlayer) HandleRogueActionBuffSelectResult(rsp *proto.HandleRogueCommonPendingActionScRsp,
	callback *proto.HandleRogueCommonPendingActionCsReq_BuffSelectResult, queueLocation uint32) {
	action := g.GetPd().GetRogueActionByQueuePosition(queueLocation)
	if action == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_ROGUE_SELECT_BUFF_NOT_EXIST)
		return
	}
	buffMap := make(map[uint32]*spb.RogueBuff)
	if a, ok := action.Action.(*spb.RogueAction_BuffSelect); ok {
		bufffId := callback.BuffSelectResult.BuffSelectId
		if info := a.BuffSelect.BuffMap[bufffId]; info != nil {
			buffMap[bufffId] = info
			rsp.QueueLocation = queueLocation
			rsp.Action = &proto.HandleRogueCommonPendingActionScRsp_BuffRerollCallback{
				BuffRerollCallback: &proto.RogueBuffRerollCallback{},
			}
			g.AddAndSyncEntityBuffChangeListScNotify(buffMap)
			g.GetPd().DelRogueActionByQueuePosition(queueLocation)
		} else {
			rsp.Retcode = uint32(proto.Retcode_RET_ROGUE_SELECT_BUFF_NOT_EXIST)
			return
		}
	}
}

func (g *GamePlayer) AddAndSyncEntityBuffChangeListScNotify(buffMap map[uint32]*spb.RogueBuff) {
	notify := &proto.SyncEntityBuffChangeListScNotify{
		EntityBuffChangeList: make([]*proto.EntityBuffChangeInfo, 0),
	}
	g.GetPd().AddRogueBuff(buffMap)
	for buffId, _ := range buffMap {
		notify.EntityBuffChangeList = append(notify.EntityBuffChangeList, &proto.EntityBuffChangeInfo{
			EntityId:       g.GetPd().GetCurAvatarEntity(),
			Reason:         0,
			BuffChangeInfo: g.GetPd().GetRogueBuffInfoById(buffId),
			RemoveBuffId:   0,
		})
	}
	g.Send(cmd.SyncEntityBuffChangeListScNotify, notify)
}

// 模拟宇宙攻击事件结算
func (g *GamePlayer) RoguePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq, sce *model.SceneCastEntity) {
	// // buff同步
	// battleDb := g.GetBattleBackupById(req.BattleId)
	// g.SyncEntityBuffChangeListScNotify(battleDb.AttackedByEntityId)

	win := func() { // 胜利
		curRoom := g.GetPd().GetCurQuestRogueRoom()
		if len(curRoom.NextSiteIdList) == 0 {
			// 没有关卡了,结算!
			// g.GetPd().SetQuestRogueRoomStatus(g.GetPd().GetCurRogue().)
			// g.SyncRogueMapRoomScNotify(g.GetPd().GetCurRogue().CurSiteId)
			g.Send(cmd.SyncRogueExploreWinScNotify, &proto.SyncRogueExploreWinScNotify{IsExploreWin: true})
		} else {
			// 祝福选择页通知
			for _, _ = range sce.MonsterEntityIdList {
				g.GetPd().AddRogueActionBuffSelect([]uint32{g.GetPd().GetRogueBuff(),
					g.GetPd().GetRogueBuff(), g.GetPd().GetRogueBuff()})
			}

			// 刷新门
		}
	}

	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_WIN:
		win()
	}

	g.RogueAction() // 模拟宇宙自动化
}

// 区域通知
func (g *GamePlayer) SyncRogueMapRoomScNotify(siteId uint32) {
	// curRogue := g.GetPd().GetCurRogue()
	db := g.GetPd().GetQuestRogueRoom()[siteId]
	if db == nil {
		return
	}

	notify := &proto.SyncRogueMapRoomScNotify{
		CurRoom: &proto.RogueRoom{
			CurStatus: proto.RogueRoomStatus(db.RoomStatus),
			SiteId:    siteId,
			RoomId:    db.RoomId,
		},
		// MapId: curRogue.RogueMapId,
	}
	g.Send(cmd.SyncRogueMapRoomScNotify, notify)
}

func (g *GamePlayer) SyncRogueVirtualItemInfoScNotify() {
	notify := &proto.SyncRogueVirtualItemInfoScNotify{
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{},
	}

	g.Send(cmd.SyncRogueVirtualItemInfoScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonActionResultScNotify(buffId uint32) {
	db := g.GetPd().GetRogueBuffById(buffId)
	if db == nil {
		return
	}
	notify := &proto.SyncRogueCommonActionResultScNotify{
		ActionResultList: make([]*proto.RogueCommonActionResult, 0),
		RogueSubMode:     101,
	}
	notify.ActionResultList = append(notify.ActionResultList, &proto.RogueCommonActionResult{
		Source: 0,
		RogueAction: &proto.RogueCommonActionResultData{
			ResultData: &proto.RogueCommonActionResultData_GetBuffList{
				GetBuffList: &proto.RogueCommonBuff{
					BuffId:    buffId,
					BuffLevel: db.BuffLevel,
				},
			},
		},
	})

	g.Send(cmd.SyncRogueCommonActionResultScNotify, notify)
}

func QuitRogueCsReq(g *GamePlayer, payloadMsg pb.Message) {
	db := g.GetPd().GetCurRogue()
	db.Status = spb.RogueStatus_ROGUE_STATUS_FINISH
	// db.IsWin = true
	g.Send(cmd.SyncRogueStatusScNotify, &proto.SyncRogueStatusScNotify{Status: proto.RogueStatus(db.Status)})
	rsp := &proto.QuitRogueScRsp{
		RogueGameInfo: g.GetPd().GetQuestRogueInfo(),
	}
	g.Send(cmd.QuitRogueScRsp, rsp)
	// g.GetPd().NewCurRogue()
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
}

func LeaveRogueCsReq(g *GamePlayer, payloadMsg pb.Message) {
	curLine := g.GetPd().GetCurLineUp()
	// SyncRogueFinishScNotify
	rsp := &proto.LeaveRogueScRsp{
		RogueGameInfo: g.GetPd().GetQuestRogueInfo(),
		Lineup:        g.GetPd().GetLineUpPb(curLine),
		Scene: g.GetPd().GetSceneInfo(
			g.GetPd().GetScene().EntryId, g.GetPd().GetPosPb(), g.GetPd().GetRotPb(), curLine),
	}

	g.Send(cmd.LeaveRogueScRsp, rsp)
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
}

func EnterRogueMapRoomCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.EnterRogueMapRoomCsReq)
	// g.GetPd().FinishRogueRoom(g.GetPd().GetCurRogue().CurSiteId)
	// 通知old
	// g.SyncRogueMapRoomScNotify(g.GetPd().GetCurRogue().CurSiteId)
	// g.GetPd().UpCurRogueRoom(req.SiteId)
	// 更新通知
	g.SyncRogueMapRoomScNotify(req.SiteId)
	rsp := &proto.EnterRogueMapRoomScRsp{
		RotateInfo: g.GetPd().GetRogueMapRotateInfo(req.RoomId),
		Lineup:     g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Rogue)),
		CurSiteId:  req.SiteId,
		Retcode:    0,
		Scene:      g.GetPd().GetRogueScene(req.RoomId),
	}

	g.Send(cmd.EnterRogueMapRoomScRsp, rsp)
}

func GetRogueBuffEnhanceInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetRogueBuffEnhanceInfoScRsp{
		BuffEnhanceInfo: &proto.RogueBuffEnhanceInfoList{
			EnhanceInfoList: make([]*proto.RogueBuffEnhanceInfo, 0),
		},
	}
	db := g.GetPd().GetRogueBuffList()
	for _, info := range db {
		rsp.BuffEnhanceInfo.EnhanceInfoList = append(rsp.BuffEnhanceInfo.EnhanceInfoList, &proto.RogueBuffEnhanceInfo{
			CostData: &proto.ItemCostData{
				ItemList: make([]*proto.ItemCost, 0),
			},
			BuffId: info.BuffId,
		})
	}
	g.Send(cmd.GetRogueBuffEnhanceInfoScRsp, rsp)
}

func GetRogueAdventureRoomInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetRogueAdventureRoomInfoScRsp{
		// NKIEHEJPKPK: &proto.BDJFNCAHDCP{
		// 	OLHEOHGEGEP: 16,
		// 	IMPECOKHIHL: 1,
		// 	Status:      0,
		// 	MIBJHFNEHJC: 0,
		// 	ScoreId:     0,
		// },
	}
	g.Send(cmd.GetRogueAdventureRoomInfoScRsp, rsp)
}
