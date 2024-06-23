package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) RogueTournQueryCsReq(payloadMsg []byte) {
	rsp := &proto.RogueTournQueryScRsp{
		Retcode: 0,
		// RogueTournCurInfo: g.GetRogueTournCurInfo(),
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

func (g *GamePlayer) RogueTournStartCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.RogueTournStartCsReq, payloadMsg)
	req := msg.(*proto.RogueTournStartCsReq)
	rsp := new(proto.RogueTournStartScRsp)
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
	g.SetBattleStatus(spb.BattleType_Battle_ROGUE_TOURN)

	rsp.RogueTournCurSceneInfo = &proto.RogueTournCurSceneInfo{
		Lineup: g.GetBattleLineUpPb(RogueTourn),
		Scene:  g.GetRogueTournScene(8040701),
		// FFKCPBBDCGL: nil,
	}
	rsp.RogueTournCurInfo = g.GetRogueTournCurInfo()
	g.Send(cmd.RogueTournStartScRsp, rsp)
}
