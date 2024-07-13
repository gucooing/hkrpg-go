package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetRaidInfoCsReq(payloadMsg []byte) {
	rsp := &proto.GetRaidInfoScRsp{
		ChallengeTakenRewardIdList: make([]uint32, 0),
		ChallengeRaidList:          make([]*proto.ChallengeRaid, 0),
		FinishedRaidInfoList:       make([]*proto.FinishedRaidInfo, 0),
		Retcode:                    0,
	}
	for _, db := range g.GetFinishRaidMap() {
		rsp.FinishedRaidInfoList = append(rsp.FinishedRaidInfoList, &proto.FinishedRaidInfo{
			WorldLevel: db.HardLevel,
			RaidId:     db.RaidId,
		})
	}

	g.Send(cmd.GetRaidInfoScRsp, rsp)
}

func (g *GamePlayer) StartRaidCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartRaidCsReq, payloadMsg)
	req := msg.(*proto.StartRaidCsReq)
	rsp := &proto.StartRaidScRsp{}
	g.NewRaidInfo(req.RaidId) // 重置
	rsp.Retcode = uint32(g.newRaidInfo(req))
	// 设置状态
	db := g.GetRaidInfo(req.RaidId)
	g.SetBattleStatus(spb.BattleType_Battle_RAID)
	g.RaidInfoNotify(req.RaidId)
	g.RaidEnterSceneByServerScNotify(db.EntryId)

	g.Send(cmd.StartRaidScRsp, rsp)
}

func (g *GamePlayer) LeaveRaidCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.LeaveRaidCsReq, payloadMsg)
	req := msg.(*proto.LeaveRaidCsReq)
	rsp := &proto.LeaveRaidScRsp{}
	db := g.GetFinishRaidInfo(req.RaidId)
	if db == nil {
		g.Send(cmd.LeaveRaidScRsp, rsp)
		return
	}
	var teleportToAnchor = true
	// 设置状态
	g.SetBattleStatus(spb.BattleType_Battle_NONE)
	conf := gdconf.GetRaidConfig(db.RaidId, db.HardLevel)
	g.NewRaidInfo(req.RaidId) // 重置
	if conf == nil {
		g.Send(cmd.LeaveRaidScRsp, rsp)
		return
	}
	if (conf.MainMissionIDBefore != conf.MainMissionIDAfter) || conf.MainMissionIDBefore == 0 {
		teleportToAnchor = false
	}
	if teleportToAnchor {
		// 回到之前的位置
		g.SceneByServerScNotify(g.GetCurEntryId(), g.GetPosPb(), g.GetRotPb())
	}

	g.Send(cmd.LeaveRaidScRsp, rsp)
}
