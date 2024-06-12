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
	msg := g.DecodePayloadToProto(cmd.StartRogueCsReq, payloadMsg)
	req := msg.(*proto.StartRogueCsReq)
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
	lineUpDb := g.GetBattleLineUpById(Rogue)
	lineUpDb.LeaderSlot = 0
	if req.BaseAvatarIdList != nil {
		lineUpDb.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
		for id, avatarId := range req.BaseAvatarIdList {
			lineUpDb.AvatarIdList[uint32(id)] = &spb.LineAvatarList{AvatarId: avatarId, Slot: uint32(id)}
		}
	} else {
		curAvatarList := g.GetCurLineUp()
		if curAvatarList == nil {
			rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
			g.Send(cmd.StartRogueScRsp, rsp)
			return
		}
		for id, avatar := range curAvatarList.AvatarIdList {
			lineUpDb.AvatarIdList[id] = &spb.LineAvatarList{AvatarId: avatar.AvatarId, Slot: id}
		}
	}
	// 将角色属性拷贝出来
	for _, avatar := range lineUpDb.AvatarIdList {
		avatarBin := g.GetAvatarBinById(avatar.AvatarId)
		g.CopyBattleAvatar(avatarBin)
	}
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
	g.NewGetRogueBuffByType()
	// 更新db
	g.SetMaterialById(Cf, 100) // 将宇宙碎片重置成100个
	db := g.GetDbRogue()
	db.CurRogue = &spb.CurRogue{
		CurAreaId:    req.AreaId,
		AeonId:       req.AeonId,
		CurSiteId:    rogueMap.StartId,
		RogueRoomMap: rogueRoomMap,
		RogueMapId:   mapId,
	}
	// 设置状态
	g.SetBattleStatus(spb.BattleType_Battle_ROGUE)
	// 准备工作就绪,告知客户端
	g.Send(cmd.SyncRogueStatusScNotify, &proto.SyncRogueStatusScNotify{Status: proto.RogueStatus_ROGUE_STATUS_DOING})

	rsp.Lineup = g.GetBattleLineUpPb(Rogue)
	rsp.Scene = g.GetRogueScene(rogueRoomMap[rogueMap.StartId].RoomId)
	rsp.RogueInfo = g.GetRogueInfo()
	// rsp.RotateInfo

	g.Send(cmd.StartRogueScRsp, rsp)
}

// 模拟宇宙攻击事件结算
func (g *GamePlayer) RoguePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq) {
	// // buff同步
	battleDb := g.GetBattleBackupById(req.BattleId)
	g.SyncEntityBuffChangeListScNotify(battleDb.AttackedByEntityId)
	// // 积分同步
	// g.SyncRogueVirtualItemInfoScNotify()
	// // 物品增加通知
	// var pileItem []*Material
	// pileItem = append(pileItem, &Material{
	// 	Tid: 31,
	// 	Num: 21,
	// })
	// g.AddMaterial(pileItem)

	// 祝福选择页通知 SyncRogueCommonPendingActionScNotify
	// buffIdList := gdconf.GetBuffListByNum(3)
	buffIdList := make([]uint32, 0)
	for i := 0; i < 3; i++ {
		buffIdList = append(buffIdList, g.GetRogueBuff())
	}
	g.SyncRogueCommonPendingActionScNotify(buffIdList)
	// 刷新门
}

// 区域通知
func (g *GamePlayer) SyncRogueMapRoomScNotify() {
	rogue := g.GetDbRogue()

	notify := &proto.SyncRogueMapRoomScNotify{
		// CurRoom: &proto.RogueRoom{
		// 	CurStatus: proto.RogueRoomStatus(rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomStatus),
		// 	SiteId:    rogue.CurRogue.CurSiteId,
		// 	RoomId:    rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomId,
		// },
		MapId: rogue.CurRogue.RogueMapId,
	}
	g.Send(cmd.SyncRogueMapRoomScNotify, notify)
}

func (g *GamePlayer) SyncRogueVirtualItemInfoScNotify() {
	notify := &proto.SyncRogueVirtualItemInfoScNotify{
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
			// RogueCoin: g.GetDbRogue().CurRogue.CosmicFragment,
			// X:     8,
		},
	}

	g.Send(cmd.SyncRogueVirtualItemInfoScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonPendingActionScNotify(buffIdList []uint32) {
	// db := g.GetCurRogue()
	rogueCommonBuffList := make([]*proto.RogueCommonBuff, 0)
	firstBuffTypeList := make([]uint32, 0)
	for _, buffId := range buffIdList {
		rogueCommonBuffList = append(rogueCommonBuffList, &proto.RogueCommonBuff{
			BuffLevel: 1,
			BuffId:    buffId,
		})
		conf := gdconf.GetBuffById(buffId, 1)
		if conf != nil {
			firstBuffTypeList = append(firstBuffTypeList, conf.RogueBuffType)
		}
		g.AddRogueBuffNum()
	}
	notify := &proto.SyncRogueCommonPendingActionScNotify{
		Action: &proto.RogueCommonPendingAction{
			RogueAction: &proto.RogueAction{
				Action: &proto.RogueAction_BuffSelectInfo{
					BuffSelectInfo: &proto.RogueCommonBuffSelectInfo{
						CanRoll:           true,
						FirstBuffTypeList: firstBuffTypeList,
						SelectBuffList:    rogueCommonBuffList, // buff信息列表
						SourceCurCount:    1,
						SourceHintId:      1,
						RollBuffCount:     1,
						SourceTotalCount:  1,
						RollBuffMaxCount:  1,

						RollBuffCostData: &proto.ItemCostData{ItemList: []*proto.ItemCost{
							{
								ItemOneofCase: &proto.ItemCost_PileItem{
									PileItem: &proto.PileItem{
										ItemId:  Cf,
										ItemNum: 30,
									},
								},
							},
						}},
						RollBuffFreeCount:        0,
						SourceType:               0,
						HandbookUnlockBuffIdList: buffIdList,
					},
				},
			},
			QueuePosition: g.GetRogueBuffNum(),
		},
		RogueVersionId: 101,
	}

	g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
}

func (g *GamePlayer) SyncEntityBuffChangeListScNotify(entityId uint32) {
	notify := &proto.SyncEntityBuffChangeListScNotify{
		EntityBuffChangeList: make([]*proto.EntityBuffChange, 0),
	}
	notify.EntityBuffChangeList = append(notify.EntityBuffChangeList, &proto.EntityBuffChange{
		EntityId: entityId,
	})
	g.Send(cmd.SyncEntityBuffChangeListScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonActionResultScNotify(buffId uint32) {
	// rogue := g.GetRogue()
	// if rogue.BuffList[buffId] == nil {
	// 	return
	// }
	// notify := &proto.SyncRogueCommonActionResultScNotify{
	// 	/*
	// 		Action: &proto.RogueActionResult{
	// 			ActionData: &proto.RogueActionResultData{
	// 				AddBuffList: &proto.RogueBuffData{
	// 					Level:  rogue.BuffList[buffId].Level,
	// 					BuffId: buffId,
	// 				},
	// 			},
	// 			Source: proto.RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_SELECT,
	// 		},
	// 	*/
	// 	RogueVersionId: g.GetDbRogue().CurRogue.RogueMapID,
	// }
	//
	// g.Send(cmd.SyncRogueCommonActionResultScNotify, notify)
}

func (g *GamePlayer) GetRogueHandbookDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetRogueHandbookDataScRsp{
		HandbookInfo: &proto.RogueHandbook{
			// MiracleList: make([]*proto.RogueHandbookMiracle, 0),
			// RogueEvent:  make([]*proto.RogueHandbookEvent, 0),
			// BuffList:    make([]*proto.RogueHandbookBuff, 0),
		},
	}

	// 解锁全部祝福
	// allBuffList := gdconf.GetAllBuff()
	// for _, id := range allBuffList {
	// 	buff := &proto.RogueHandbookBuff{
	// 		BuffId: id,
	// 	}
	// 	rsp.HandbookInfo.BuffList = append(rsp.HandbookInfo.BuffList, buff)
	// }

	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

// 模拟宇宙攻击事件
func (g *GamePlayer) RogueSceneCastSkillCsReq(rsp *proto.SceneCastSkillScRsp) {
	rsp.BattleInfo.BattleTargetInfo = make(map[uint32]*proto.BattleTargetList)
	rsp.BattleInfo.BattleTargetInfo[2] = &proto.BattleTargetList{
		BattleTargetList: make([]*proto.BattleTarget, 0),
	}
	battleTargetList := make([]*proto.BattleTarget, 0)
	bTL := []uint32{50008, 50007, 50006, 50005, 50004, 50003, 50001, 30072, 30071, 30070, 30069, 30068, 30067, 30066, 30065, 30064, 30063, 30062, 30061, 30060, 30059, 30058, 30057, 30056, 30055, 30054, 30053, 30046, 30045, 30044, 30043, 30042, 30041, 30040, 30039, 30037, 30034, 30024, 30021, 30016, 30012, 30008, 30006, 30004, 30002, 20074, 20069, 20068, 20067, 20066, 20064, 20060, 20058, 20052, 20050, 20049, 20047, 20046, 20045, 20044, 20043, 20040, 20039, 20034, 20033, 20032, 20031, 20024, 20021, 20019, 20018, 20017, 20016, 20015, 20013, 20010, 20009, 20008, 20002}
	for _, id := range bTL {
		battleTarget := &proto.BattleTarget{
			Id: id,
		}
		battleTargetList = append(battleTargetList, battleTarget)
	}
	rsp.BattleInfo.BattleTargetInfo[2].BattleTargetList = battleTargetList
	// 添加角色
	// rsp.BattleInfo.BattleAvatarList = g.GetBattleAvatarList(uint32(proto.ExtraLineupType_LINEUP_ROGUE))
	// 添加buff
	rsp.BattleInfo.BuffList = make([]*proto.BattleBuff, 0)
	// for id, buff := range g.GetRogueBuff() {
	// 	battleBuff := &proto.BattleBuff{
	// 		Id:         id,
	// 		Level:      buff.Level,
	// 		OwnerIndex: 4294967295,
	// 		WaveFlag:   4294967295,
	// 	}
	// 	rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, battleBuff)
	// }
	// 路途buff！！！
	rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, &proto.BattleBuff{
		Id:       1000111,
		Level:    1,
		WaveFlag: 4294967295,
		DynamicValues: map[string]float32{
			"SkillIndex": 1,
		},
		TargetIndexList: []uint32{1},
	})
	g.Send(cmd.SceneCastSkillScRsp, rsp)
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

func (g *GamePlayer) HandleRogueCommonPendingActionCsReq(payloadMsg []byte) {
	// msg := g.DecodePayloadToProto(cmd.HandleRogueCommonPendingActionCsReq, payloadMsg)
	// req := msg.(*proto.HandleRogueCommonPendingActionCsReq)

	// 添加后通知启动
	// g.SyncRogueCommonActionResultScNotify(buffSelectResult.BuffId)
	// 模拟宇宙图鉴更新通知？ SyncRogueHandbookDataUpdateScNotify
	// 模拟宇宙常见操作结果通知 SyncRogueCommonActionResultScNotify // add buff, buff状态
	rsp := &proto.HandleRogueCommonPendingActionScRsp{
		QueuePosition: g.GetRogueBuffNum(),
		Retcode:       0,
		QueueLocation: g.GetRogueBuffNum(),
		Action:        nil,
	}

	g.Send(cmd.HandleRogueCommonPendingActionScRsp, rsp)
}

func (g *GamePlayer) EnterRogueMapRoomCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.EnterRogueMapRoomCsReq, payloadMsg)
	req := msg.(*proto.EnterRogueMapRoomCsReq)
	// curRogue := g.GetCurDbRogue()
	// if curRogue.RogueSceneMap[req.SiteId] == nil && curRogue.RogueSceneMap[req.SiteId].RoomId != req.RoomId {
	// 	return
	// }
	// curRogue.RogueSceneMap[curRogue.CurSiteId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_FINISH
	// curRogue.RogueSceneMap[req.SiteId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY
	// curRogue.CurSiteId = req.SiteId
	// 更新通知
	g.SyncRogueMapRoomScNotify()
	scene := g.GetRogueScene(req.RoomId)
	if scene == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}

	rsp := &proto.EnterRogueMapRoomScRsp{
		Lineup:    g.GetBattleLineUpPb(Rogue),
		CurSiteId: req.SiteId,
		Retcode:   0,
		Scene:     scene,
	}

	g.Send(cmd.EnterRogueMapRoomScRsp, rsp)
}
