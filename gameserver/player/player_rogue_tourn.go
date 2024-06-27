package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) RogueTournQueryCsReq(payloadMsg []byte) {
	rsp := &proto.RogueTournQueryScRsp{
		Retcode:           0,
		RogueTournCurInfo: g.GetRogueTournCurInfo(),
		RogueTournInfo: &proto.RogueTournInfo{
			RogueTournSaveList:       make([]*proto.RogueTournSaveList, 0),
			RogueTournAreaInfo:       g.GetRogueTournAreaInfo(),
			InspirationCircuit:       g.GetInspirationCircuitInfo(),
			RogueTournSeasonInfo:     g.GetRogueTournSeasonInfo(),
			ExtraScoreInfo:           g.GetExtraScoreInfo(),
			RogueTournExpInfo:        g.GetRogueTournExpInfo(),
			RogueTournCollectionInfo: g.GetRogueTournCollectionInfo(),
			RogueTournDifficultyInfo: g.GetRogueTournDifficultyInfo(),
		},
	}
	g.Send(cmd.RogueTournQueryScRsp, rsp)
}

func (g *GamePlayer) RogueTournGetPermanentTalentInfoCsReq(payloadMsg []byte) {
	rsp := &proto.RogueTournGetPermanentTalentInfoScRsp{
		InspirationCircuit: g.GetInspirationCircuitInfo(),
		Retcode:            0,
	}
	g.Send(cmd.RogueTournGetPermanentTalentInfoScRsp, rsp)
}

func (g *GamePlayer) RogueTournGetMiscRealTimeDataCsReq(payloadMsg []byte) {
	rsp := &proto.RogueTournGetMiscRealTimeDataScRsp{}
	g.Send(cmd.RogueTournGetMiscRealTimeDataScRsp, rsp)
}

func (g *GamePlayer) RogueTournStartCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.RogueTournStartCsReq, payloadMsg)
	req := msg.(*proto.RogueTournStartCsReq)
	rsp := new(proto.RogueTournStartScRsp)
	conf := gdconf.GetRogueTournAreaById(req.AreaId)
	if conf == nil {
		g.Send(cmd.RogueTournStartScRsp, rsp)
		return
	}
	// 更新队伍
	lineUpDb := g.GetBattleLineUpById(RogueTourn)
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
			g.Send(cmd.RogueTournStartScRsp, rsp)
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

	// 更新db
	g.SetBattleStatus(spb.BattleType_Battle_ROGUE_TOURN)
	g.SetMaterialById(Cf, 100) // 将宇宙碎片重置成100个
	g.NewCurRogueTourn(req.AreaId)
	curRoom := g.GetCurRogueTournRoom()

	rsp.RogueTournCurSceneInfo = &proto.RogueTournCurSceneInfo{
		Lineup: g.GetBattleLineUpPb(RogueTourn),
		Scene:  g.GetRogueTournScene(curRoom.RoomId),
	}
	rsp.RogueTournCurInfo = g.GetRogueTournCurInfo()

	// 选择三个初始方程
	g.AddRogueBuffNum()
	g.FormulaSyncRogueCommonPendingActionScNotify([]uint32{130204, 130408, 130307})
	rsp.RogueTournCurInfo.RogueTournCurAreaInfo.PendingAction = &proto.RogueCommonPendingAction{
		QueuePosition: g.GetRogueBuffNum(),
		RogueAction: &proto.RogueAction{
			Action: &proto.RogueAction_RogueFormulaSelectInfo{
				RogueFormulaSelectInfo: &proto.RogueFormulaSelectInfo{
					SelectFormulaIdListFieldNumber: []uint32{130204, 130408, 130307},
				},
			},
		},
	}

	g.Send(cmd.RogueTournStartScRsp, rsp)
}

func (g *GamePlayer) RogueTournEnterCsReq(payloadMsg []byte) {
	curRoom := g.GetCurRogueTournRoom()
	rsp := &proto.RogueTournEnterScRsp{
		RogueTournCurInfo: g.GetRogueTournCurInfo(),
		RogueTournCurSceneInfo: &proto.RogueTournCurSceneInfo{
			Lineup: g.GetBattleLineUpPb(RogueTourn),
			Scene:  g.GetRogueTournScene(curRoom.RoomId),
		},
	}

	g.Send(cmd.RogueTournEnterScRsp, rsp)
}

func (g *GamePlayer) RogueTournSettleCsReq(payloadMsg []byte) {
	curRoom := g.GetCurRogueTournRoom()
	rsp := &proto.RogueTournSettleScRsp{
		Retcode: 0,
		RogueTournCurSceneInfo: &proto.RogueTournCurSceneInfo{
			Lineup: g.GetBattleLineUpPb(RogueTourn),
			Scene:  g.GetRogueTournScene(curRoom.RoomId),
		},
		// IOLFDOIPNKA:            nil,
	}
	db := g.GetRogueTourn()
	db.CurRogueTourn = nil
	g.Send(cmd.RogueTournSettleScRsp, rsp)
}

// 方程选择通知
func (g *GamePlayer) FormulaSyncRogueCommonPendingActionScNotify(formulaList []uint32) {
	notify := &proto.SyncRogueCommonPendingActionScNotify{
		RogueSubMode: 301,
		Action: &proto.RogueCommonPendingAction{
			QueuePosition: g.GetRogueBuffNum(),
			RogueAction: &proto.RogueAction{
				Action: &proto.RogueAction_RogueFormulaSelectInfo{
					RogueFormulaSelectInfo: &proto.RogueFormulaSelectInfo{
						SelectFormulaIdListFieldNumber: formulaList,
					},
				},
			},
		},
	}

	g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
}

func (g *GamePlayer) FormulaSyncRogueCommonActionResultScNotify(formulaId uint32) {
	conf := gdconf.GetRogueTournFormulaById(formulaId)
	if conf == nil {
		return
	}
	notify := &proto.SyncRogueCommonActionResultScNotify{
		ActionResult: make([]*proto.RogueCommonActionResult, 0),
		RogueSubMode: 301,
	}
	notify.ActionResult = append(notify.ActionResult, &proto.RogueCommonActionResult{
		Source: 0,
		RogueAction: &proto.RogueCommonActionResultData{
			ResultData: &proto.RogueCommonActionResultData_GetFormulaList{
				GetFormulaList: &proto.RogueCommonFormula{
					FormulaInfo: &proto.FormulaInfo{
						IsExpand:  false,
						FormulaId: formulaId,
						FormulaBuffTypeList: []*proto.FormulaBuffTypeInfo{
							{
								Num:        conf.MainBuffNum,
								BuffTypeId: conf.MainBuffTypeID,
							},
							{
								Num:        conf.SubBuffNum,
								BuffTypeId: conf.SubBuffTypeID,
							},
						},
					},
				},
			},
		},
	})

	g.Send(cmd.SyncRogueCommonActionResultScNotify, notify)
}
