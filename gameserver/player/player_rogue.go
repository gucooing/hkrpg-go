package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetRogueScoreRewardInfoCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetRogueScoreRewardInfoScRsp{
		Retcode: 0,
		Info:    g.GetPd().GetRogueScoreRewardInfo(),
	}

	g.Send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueInitialScoreCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetRogueInitialScoreScRsp{
		RogueScoreRewardInfo: g.GetPd().GetRogueScoreRewardInfo(),
		Retcode:              0,
	}

	g.Send(cmd.GetRogueInitialScoreScRsp, rsp)
}

func (g *GamePlayer) GetRogueTalentInfoCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetRogueTalentInfoScRsp{
		TalentInfoList: &proto.RogueTalentInfoList{
			TalentInfo: make([]*proto.RogueTalentInfo, 0),
		},
	}

	for _, talent := range gdconf.GetTalentIDList() {
		rogueTalent := &proto.RogueTalentInfo{
			Status:   proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE,
			TalentId: talent,
		}
		rsp.TalentInfoList.TalentInfo = append(rsp.TalentInfoList.TalentInfo, rogueTalent)
	}

	g.Send(cmd.GetRogueTalentInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueInfoCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetRogueInfoScRsp)
	rsp.RogueGameInfo = g.GetPd().GetRogueInfo()

	g.Send(cmd.GetRogueInfoScRsp, rsp)
}

func (g *GamePlayer) StartRogueCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartRogueCsReq)
	rsp := &proto.StartRogueScRsp{}
	conf := gdconf.GetRogueAreaConfigById(req.AreaId)
	if conf == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	mapId := conf.AreaProgress*100 + conf.Difficulty
	rogueMap := gdconf.GetRogueMapById(mapId) //   取关卡配置
	if rogueMap == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	// 更新队伍
	g.SetBattleLineUp(model.Rogue, req.BaseAvatarIdList)
	// 取房间
	rogueRoomMap := make(map[uint32]*spb.RogueRoom, 0)
	switch conf.AreaProgress {
	case 0:
		rogueRoomMap[1] = &spb.RogueRoom{
			RoomId:         100,
			RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY,
			NextSiteIdList: make([]uint32, 0),
		}
	case 1:
		for id, rogue := range rogueMap.SiteList {
			rogueRoomMap[id] = &spb.RogueRoom{
				RoomId:         gdconf.GetRogueRoomTypeBy100(id),
				RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
				NextSiteIdList: rogue.NextSiteIDList,
			}
		}
	default:
		for id, rogue := range rogueMap.SiteList {
			rogueRoomMap[id] = &spb.RogueRoom{
				RoomId:         gdconf.GetRogueRoomTypeBySiteID(id),
				RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
				NextSiteIdList: rogue.NextSiteIDList,
			}
		}
	}
	rogueRoomMap[rogueMap.StartId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY

	// 初始化祝福列表
	g.GetPd().NewGetRogueBuffByType()
	// 更新db
	g.GetPd().SetMaterialById(model.Cf, 100) // 将宇宙碎片重置成100个
	db := g.GetPd().GetDbRogue()
	db.CurRogue = &spb.CurRogue{
		CurAreaId:             req.AreaId,
		AeonId:                req.AeonId,
		CurSiteId:             rogueMap.StartId,
		RogueRoomMap:          rogueRoomMap,
		RogueMapId:            mapId,
		RogueActivityModuleID: model.QuestRogue,
		Status:                spb.RogueStatus_ROGUE_STATUS_DOING,
	}
	// 设置状态
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_ROGUE)
	// 准备工作就绪,告知客户端
	g.Send(cmd.SyncRogueStatusScNotify, &proto.SyncRogueStatusScNotify{Status: proto.RogueStatus_ROGUE_STATUS_DOING})

	rsp.Lineup = g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Rogue))
	rsp.Scene = g.GetPd().GetRogueScene(rogueRoomMap[rogueMap.StartId].RoomId)
	rsp.RogueGameInfo = g.GetPd().GetRogueInfo()

	g.Send(cmd.StartRogueScRsp, rsp)
}

// 模拟宇宙攻击事件结算
func (g *GamePlayer) RoguePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq, sce *model.SceneCastEntity) {
	// // buff同步
	// battleDb := g.GetBattleBackupById(req.BattleId)
	// g.SyncEntityBuffChangeListScNotify(battleDb.AttackedByEntityId)

	curRoom := g.GetPd().GetCurRogueRoom()
	if len(curRoom.NextSiteIdList) == 0 {
		// 没有关卡了,结算!
		g.GetPd().FinishRogueRoom(g.GetPd().GetCurRogue().CurSiteId)
		g.SyncRogueMapRoomScNotify(g.GetPd().GetCurRogue().CurSiteId)
		g.Send(cmd.SyncRogueExploreWinScNotify, &proto.SyncRogueExploreWinScNotify{IsExploreWin: true})
	} else {
		// 祝福选择页通知 SyncRogueCommonPendingActionScNotify
		// for x := 0; x < monsterNum; x++ {
		buffIdList := make([]uint32, 0)
		for i := 0; i < 3; i++ {
			// buffIdList = append(buffIdList, g.GetRogueBuff())
			buffIdList = append(buffIdList, gdconf.GetRogueBuff())
		}
		g.SyncRogueCommonPendingActionScNotify(buffIdList)
		// }

		// 刷新门
	}
}

func (g *GamePlayer) HandleRogueCommonPendingActionCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.HandleRogueCommonPendingActionCsReq)
	action := req.Action
	if action != nil {
		switch action.(type) {
		case *proto.HandleRogueCommonPendingActionCsReq_BuffSelectResult: // 添加buff
			buffId := action.(*proto.HandleRogueCommonPendingActionCsReq_BuffSelectResult).BuffSelectResult.BuffSelectId
			g.GetPd().AddRogueBuff(buffId)
			g.SyncRogueCommonActionResultScNotify(buffId)
		case *proto.HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult: // 添加方程
			formulaId := action.(*proto.HandleRogueCommonPendingActionCsReq_RogueTournFormulaResult).RogueTournFormulaResult.TournFormulaId
			g.GetPd().AddCurRogueTournFormula(formulaId)
			g.FormulaSyncRogueCommonActionResultScNotify(formulaId)
		}
	}

	rsp := &proto.HandleRogueCommonPendingActionScRsp{
		QueuePosition: g.GetPd().GetRogueBuffNum(),
		Retcode:       0,
		QueueLocation: req.QueueLocation,
		Action:        nil,
	}

	g.Send(cmd.HandleRogueCommonPendingActionScRsp, rsp)
}

// 区域通知
func (g *GamePlayer) SyncRogueMapRoomScNotify(siteId uint32) {
	curRogue := g.GetPd().GetCurRogue()
	db := g.GetPd().GetRogueRoom()[siteId]
	if db == nil {
		return
	}

	notify := &proto.SyncRogueMapRoomScNotify{
		CurRoom: &proto.RogueRoom{
			CurStatus: proto.RogueRoomStatus(db.RoomStatus),
			SiteId:    siteId,
			RoomId:    db.RoomId,
		},
		MapId: curRogue.RogueMapId,
	}
	g.Send(cmd.SyncRogueMapRoomScNotify, notify)
}

func (g *GamePlayer) SyncRogueVirtualItemInfoScNotify() {
	notify := &proto.SyncRogueVirtualItemInfoScNotify{
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{},
	}

	g.Send(cmd.SyncRogueVirtualItemInfoScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonPendingActionScNotify(buffIdList []uint32) {
	rogueCommonBuffList := make([]*proto.RogueCommonBuff, 0)
	firstBuffTypeList := make([]uint32, 0)
	for _, buffId := range buffIdList {
		rogueCommonBuffList = append(rogueCommonBuffList, &proto.RogueCommonBuff{
			BuffLevel: 1,
			BuffId:    buffId,
		})
		conf := gdconf.GetBuffByIdAndLevel(buffId, 1)
		if conf != nil {
			firstBuffTypeList = append(firstBuffTypeList, conf.RogueBuffType)
		}
		g.GetPd().AddRogueBuffNum()
	}
	notify := &proto.SyncRogueCommonPendingActionScNotify{
		Action: &proto.RogueCommonPendingAction{
			RogueAction: &proto.RogueAction{
				Action: &proto.RogueAction_BuffSelectInfo{
					BuffSelectInfo: &proto.RogueCommonBuffSelectInfo{
						CanRoll:           true,                // 是否可以刷新
						RollBuffCount:     0,                   // 已刷新次数
						RollBuffMaxCount:  1,                   // 刷新最大次数
						RollBuffFreeCount: 1,                   // 免费刷新次数
						FirstBuffTypeList: firstBuffTypeList,   // buff属性列表
						SelectBuffList:    rogueCommonBuffList, // buff信息列表
						SourceCurCount:    1,                   // 提示
						SourceHintId:      1,                   // 提示文本
						SourceTotalCount:  1,                   // Source To Count
						RollBuffCostData: &proto.ItemCostData{ItemList: []*proto.ItemCost{ // 刷新需要的东西
							{
								PileItem: &proto.PileItem{
									ItemId:  model.Cf,
									ItemNum: g.GetPd().GetMaterialById(model.Cf),
								},
							},
						}},
						SourceType:               0,
						HandbookUnlockBuffIdList: buffIdList,
					},
				},
			},
			QueuePosition: g.GetPd().GetRogueBuffNum(),
		},
		RogueSubMode: 101,
	}

	g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
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

func (g *GamePlayer) GetRogueHandbookDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetRogueHandbookDataScRsp{
		// HandbookInfo: &proto.RogueHandbook{},
	}
	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func (g *GamePlayer) QuitRogueCsReq(payloadMsg pb.Message) {
	db := g.GetPd().GetCurRogue()
	db.Status = spb.RogueStatus_ROGUE_STATUS_FINISH
	db.IsWin = true
	g.Send(cmd.SyncRogueStatusScNotify, &proto.SyncRogueStatusScNotify{Status: proto.RogueStatus(db.Status)})
	rsp := &proto.QuitRogueScRsp{
		RogueGameInfo: g.GetPd().GetRogueInfo(),
	}
	g.Send(cmd.QuitRogueScRsp, rsp)
	g.GetPd().NewCurRogue()
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
}

func (g *GamePlayer) LeaveRogueCsReq(payloadMsg pb.Message) {
	curLine := g.GetPd().GetCurLineUp()
	// SyncRogueFinishScNotify
	rsp := &proto.LeaveRogueScRsp{
		RogueGameInfo: g.GetPd().GetRogueInfo(),
		Lineup:        g.GetPd().GetLineUpPb(curLine),
		Scene: g.GetPd().GetSceneInfo(
			g.GetPd().GetScene().EntryId, g.GetPd().GetPosPb(), g.GetPd().GetRotPb(), curLine),
	}

	g.Send(cmd.LeaveRogueScRsp, rsp)
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
}

func (g *GamePlayer) EnterRogueMapRoomCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.EnterRogueMapRoomCsReq)
	g.GetPd().FinishRogueRoom(g.GetPd().GetCurRogue().CurSiteId)
	// 通知old
	g.SyncRogueMapRoomScNotify(g.GetPd().GetCurRogue().CurSiteId)
	g.GetPd().UpCurRogueRoom(req.SiteId)
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

func (g *GamePlayer) GetRogueBuffEnhanceInfoCsReq(payloadMsg pb.Message) {
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

func (g *GamePlayer) GetRogueAdventureRoomInfoCsReq(payloadMsg pb.Message) {
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
