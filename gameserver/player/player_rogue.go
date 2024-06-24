package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetRogueScoreRewardInfoCsReq(payloadMsg []byte) {
	rsp := &proto.GetRogueScoreRewardInfoScRsp{
		Retcode: 0,
		Info:    g.GetRogueScoreRewardInfo(),
	}

	g.Send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueInitialScoreCsReq(payloadMsg []byte) {
	rsp := &proto.GetRogueInitialScoreScRsp{
		RogueScoreRewardInfo: g.GetRogueScoreRewardInfo(),
		Retcode:              0,
	}

	g.Send(cmd.GetRogueInitialScoreScRsp, rsp)
}

func (g *GamePlayer) GetRogueTalentInfoCsReq(payloadMsg []byte) {
	rsp := &proto.GetRogueTalentInfoScRsp{
		RogueTalentInfo: &proto.RogueTalentInfo{
			RogueTalentList: make([]*proto.RogueTalent, 0),
		},
	}

	for _, talent := range gdconf.GetTalentIDList() {
		rogueTalent := &proto.RogueTalent{
			Status:   proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE,
			TalentId: talent,
		}
		rsp.RogueTalentInfo.RogueTalentList = append(rsp.RogueTalentInfo.RogueTalentList, rogueTalent)
	}

	g.Send(cmd.GetRogueTalentInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetRogueInfoScRsp)
	rsp.RogueInfo = g.GetRogueInfo()

	g.Send(cmd.GetRogueInfoScRsp, rsp)
}

func (g *GamePlayer) StartRogueCsReq(payloadMsg []byte) {
	// msg := g.DecodePayloadToProto(cmd.StartRogueCsReq, payloadMsg)
	// req := msg.(*proto.StartRogueCsReq)
	// rsp := &proto.StartRogueScRsp{}
	// conf := gdconf.GetRogueAreaConfigById(req.AreaId)
	// if conf == nil {
	// 	rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
	// 	g.Send(cmd.StartRogueScRsp, rsp)
	// 	return
	// }
	// mapId := conf.AreaProgress*100 + conf.Difficulty
	// rogueMap := gdconf.GetRogueMapById(mapId) //   取关卡配置
	// if rogueMap == nil {
	// 	rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
	// 	g.Send(cmd.StartRogueScRsp, rsp)
	// 	return
	// }
	// // 更新队伍
	// lineUpDb := g.GetBattleLineUpById(Rogue)
	// lineUpDb.LeaderSlot = 0
	// if req.BaseAvatarIdList != nil {
	// 	lineUpDb.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	// 	for id, avatarId := range req.BaseAvatarIdList {
	// 		lineUpDb.AvatarIdList[uint32(id)] = &spb.LineAvatarList{AvatarId: avatarId, Slot: uint32(id)}
	// 	}
	// } else {
	// 	curAvatarList := g.GetCurLineUp()
	// 	if curAvatarList == nil {
	// 		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
	// 		g.Send(cmd.StartRogueScRsp, rsp)
	// 		return
	// 	}
	// 	for id, avatar := range curAvatarList.AvatarIdList {
	// 		lineUpDb.AvatarIdList[id] = &spb.LineAvatarList{AvatarId: avatar.AvatarId, Slot: id}
	// 	}
	// }
	// // 将角色属性拷贝出来
	// for _, avatar := range lineUpDb.AvatarIdList {
	// 	avatarBin := g.GetAvatarBinById(avatar.AvatarId)
	// 	g.CopyBattleAvatar(avatarBin)
	// }
	// // 取房间
	// rogueRoomMap := make(map[uint32]*spb.RogueRoom, 0)
	// switch conf.AreaProgress {
	// case 0:
	// 	rogueRoomMap[1] = &spb.RogueRoom{
	// 		RoomId:         100,
	// 		RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY,
	// 		NextSiteIdList: make([]uint32, 0),
	// 	}
	// case 1:
	// 	for id, rogue := range rogueMap.SiteList {
	// 		rogueRoomMap[id] = &spb.RogueRoom{
	// 			RoomId:         gdconf.GetRogueRoomTypeBy100(id),
	// 			RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
	// 			NextSiteIdList: rogue.NextSiteIDList,
	// 		}
	// 	}
	// default:
	// 	for id, rogue := range rogueMap.SiteList {
	// 		rogueRoomMap[id] = &spb.RogueRoom{
	// 			RoomId:         gdconf.GetRogueRoomTypeBySiteID(id),
	// 			RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
	// 			NextSiteIdList: rogue.NextSiteIDList,
	// 		}
	// 	}
	// }
	// rogueRoomMap[rogueMap.StartId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY
	//
	// // 初始化祝福列表
	// g.NewGetRogueBuffByType()
	// // 更新db
	// g.SetMaterialById(Cf, 100) // 将宇宙碎片重置成100个
	// db := g.GetDbRogue()
	// db.CurRogue = &spb.CurRogue{
	// 	CurAreaId:             req.AreaId,
	// 	AeonId:                req.AeonId,
	// 	CurSiteId:             rogueMap.StartId,
	// 	RogueRoomMap:          rogueRoomMap,
	// 	RogueMapId:            mapId,
	// 	RogueActivityModuleID: QuestRogue,
	// }
	// // 设置状态
	// g.SetBattleStatus(spb.BattleType_Battle_ROGUE)
	// // 准备工作就绪,告知客户端
	// g.Send(cmd.SyncRogueStatusScNotify, &proto.SyncRogueStatusScNotify{Status: proto.RogueStatus_ROGUE_STATUS_DOING})
	//
	// rsp.Lineup = g.GetBattleLineUpPb(Rogue)
	// rsp.Scene = g.GetRogueScene(rogueRoomMap[rogueMap.StartId].RoomId)
	// rsp.RogueInfo = g.GetRogueInfo()
	// // rsp.RotateInfo
	//
	// g.Send(cmd.StartRogueScRsp, rsp)
}

// 模拟宇宙攻击事件结算
func (g *GamePlayer) RoguePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq) {
	// // buff同步
	battleDb := g.GetBattleBackupById(req.BattleId)
	g.SyncEntityBuffChangeListScNotify(battleDb.AttackedByEntityId)

	// 祝福选择页通知 SyncRogueCommonPendingActionScNotify
	buffIdList := make([]uint32, 0)
	for i := 0; i < 3; i++ {
		buffIdList = append(buffIdList, g.GetRogueBuff())
	}
	g.SyncRogueCommonPendingActionScNotify(buffIdList)
	// 刷新门
}

func (g *GamePlayer) HandleRogueCommonPendingActionCsReq(payloadMsg []byte) {
	// msg := g.DecodePayloadToProto(cmd.HandleRogueCommonPendingActionCsReq, payloadMsg)
	// req := msg.(*proto.HandleRogueCommonPendingActionCsReq)
	// action := req.Action
	// if action != nil {
	// 	switch action.(type) {
	// 	case *proto.HandleRogueCommonPendingActionCsReq_BuffSelectResult: // 添加buff
	// 		buffId := action.(*proto.HandleRogueCommonPendingActionCsReq_BuffSelectResult).BuffSelectResult.BuffId
	// 		g.AddRogueBuff(buffId)
	// 		g.SyncRogueCommonActionResultScNotify(buffId)
	// 	}
	// }
	//
	// // 模拟宇宙图鉴更新通知？ SyncRogueHandbookDataUpdateScNotify
	// // 模拟宇宙常见操作结果通知 SyncRogueCommonActionResultScNotify // add buff, buff状态
	// rsp := &proto.HandleRogueCommonPendingActionScRsp{
	// 	QueuePosition: g.GetRogueBuffNum(),
	// 	Retcode:       0,
	// 	QueueLocation: g.GetRogueBuffNum(),
	// 	Action:        nil,
	// }
	//
	// g.Send(cmd.HandleRogueCommonPendingActionScRsp, rsp)
}

// 区域通知
func (g *GamePlayer) SyncRogueMapRoomScNotify() {
	// curRogue := g.GetCurRogue()
	// curRoom := g.GetCurRogueRoom()
	//
	// notify := &proto.SyncRogueMapRoomScNotify{
	// 	CurRoom: &proto.RogueRoom{
	// 		CurStatus: proto.RogueRoomStatus(curRoom.RoomStatus),
	// 		SiteId:    curRogue.CurSiteId,
	// 		RoomId:    curRoom.RoomId,
	// 	},
	// 	MapId: curRogue.RogueMapId,
	// }
	// g.Send(cmd.SyncRogueMapRoomScNotify, notify)
}

func (g *GamePlayer) SyncRogueVirtualItemInfoScNotify() {
	notify := &proto.SyncRogueVirtualItemInfoScNotify{
		// RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
		// },
	}

	g.Send(cmd.SyncRogueVirtualItemInfoScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonPendingActionScNotify(buffIdList []uint32) {
	// rogueCommonBuffList := make([]*proto.RogueCommonBuff, 0)
	// firstBuffTypeList := make([]uint32, 0)
	// for _, buffId := range buffIdList {
	// 	rogueCommonBuffList = append(rogueCommonBuffList, &proto.RogueCommonBuff{
	// 		BuffLevel: 1,
	// 		BuffId:    buffId,
	// 	})
	// 	conf := gdconf.GetBuffByIdAndLevel(buffId, 1)
	// 	if conf != nil {
	// 		firstBuffTypeList = append(firstBuffTypeList, conf.RogueBuffType)
	// 	}
	// 	g.AddRogueBuffNum()
	// }
	// notify := &proto.SyncRogueCommonPendingActionScNotify{
	// 	Action: &proto.RogueCommonPendingAction{
	// 		RogueAction: &proto.RogueAction{
	// 			// Action: &proto.RogueAction_BuffSelectInfo{
	// 			BuffSelectInfo: &proto.RogueCommonBuffSelectInfo{
	// 				CanRoll:           true,                // 是否可以刷新
	// 				RollBuffCount:     0,                   // 已刷新次数
	// 				RollBuffMaxCount:  1,                   // 刷新最大次数
	// 				RollBuffFreeCount: 1,                   // 免费刷新次数
	// 				FirstBuffTypeList: firstBuffTypeList,   // buff属性列表
	// 				SelectBuffList:    rogueCommonBuffList, // buff信息列表
	// 				SourceCurCount:    1,                   // 提示
	// 				SourceHintId:      1,                   // 提示文本
	// 				SourceTotalCount:  1,                   // Source To Count
	// 				RollBuffCostData: &proto.ItemCostData{ItemList: []*proto.ItemCost{ // 刷新需要的东西
	// 					{
	// 						ItemOneofCase: &proto.ItemCost_PileItem{
	// 							PileItem: &proto.PileItem{
	// 								ItemId:  Cf,
	// 								ItemNum: 30,
	// 							},
	// 						},
	// 					},
	// 				}},
	// 				SourceType:               0,
	// 				HandbookUnlockBuffIdList: buffIdList,
	// 				// },
	// 			},
	// 		},
	// 		QueuePosition: g.GetRogueBuffNum(),
	// 	},
	// 	RogueVersionId: 101,
	// }
	//
	// g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonActionResultScNotify(buffId uint32) {
	// db := g.GetRogueBuffById(buffId)
	// if db == nil {
	// 	return
	// }
	// notify := &proto.SyncRogueCommonActionResultScNotify{
	// 	ActionResult:   make([]*proto.RogueCommonActionResult, 0),
	// 	RogueVersionId: 101,
	// }
	// notify.ActionResult = append(notify.ActionResult, &proto.RogueCommonActionResult{
	// 	Source: 0,
	// 	RogueAction: &proto.RogueCommonActionResultData{
	// 		ResultData: &proto.RogueCommonActionResultData_GetBuffList{
	// 			GetBuffList: &proto.RogueCommonBuff{
	// 				BuffId:    buffId,
	// 				BuffLevel: db.BuffLevel,
	// 			},
	// 		},
	// 	},
	// })
	//
	// g.Send(cmd.SyncRogueCommonActionResultScNotify, notify)
}

func (g *GamePlayer) GetRogueHandbookDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetRogueHandbookDataScRsp{
		// HandbookInfo: &proto.RogueHandbook{},
	}
	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func (g *GamePlayer) QuitRogueCsReq(payloadMsg []byte) {

	g.Send(cmd.QuitRogueScRsp, nil)
	g.SetBattleStatus(spb.BattleType_Battle_NONE)
}

func (g *GamePlayer) LeaveRogueCsReq(payloadMsg []byte) {
	curLine := g.GetCurLineUp()
	rsp := &proto.LeaveRogueScRsp{
		RogueInfo: g.GetRogueInfo(),
		Lineup:    g.GetLineUpPb(curLine),
		Scene:     g.GetSceneInfo(g.GetScene().EntryId, g.GetPosPb(), g.GetRotPb(), curLine),
	}

	g.Send(cmd.LeaveRogueScRsp, rsp)
	g.SetBattleStatus(spb.BattleType_Battle_NONE)
}

func (g *GamePlayer) EnterRogueMapRoomCsReq(payloadMsg []byte) {
	// msg := g.DecodePayloadToProto(cmd.EnterRogueMapRoomCsReq, payloadMsg)
	// req := msg.(*proto.EnterRogueMapRoomCsReq)
	// g.UpCurRogueRoom(req.SiteId)
	// // 更新通知
	// g.SyncRogueMapRoomScNotify()
	// rsp := &proto.EnterRogueMapRoomScRsp{
	// 	RotateInfo: g.GetRogueMapRotateInfo(),
	// 	Lineup:     g.GetBattleLineUpPb(Rogue),
	// 	CurSiteId:  req.SiteId,
	// 	Retcode:    0,
	// 	Scene:      g.GetRogueScene(req.RoomId),
	// }
	//
	// g.Send(cmd.EnterRogueMapRoomScRsp, rsp)
}

func (g *GamePlayer) GetRogueBuffEnhanceInfoCsReq(payloadMsg []byte) {
	// rsp := &proto.GetRogueBuffEnhanceInfoScRsp{
	// 	BuffEnhanceInfo: &proto.RogueBuffEnhanceInfo{
	// 		EnhanceInfo: make([]*proto.RogueBuffEnhance, 0),
	// 	},
	// }
	// db := g.GetRogueBuffList()
	// for _, info := range db {
	// 	rsp.BuffEnhanceInfo.EnhanceInfo = append(rsp.BuffEnhanceInfo.EnhanceInfo, &proto.RogueBuffEnhance{
	// 		CostData: &proto.ItemCostData{
	// 			ItemList: make([]*proto.ItemCost, 0),
	// 		},
	// 		BuffId: info.BuffId,
	// 	})
	// }
	// g.Send(cmd.GetRogueBuffEnhanceInfoScRsp, rsp)
}
