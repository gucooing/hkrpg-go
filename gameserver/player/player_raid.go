package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetRaidInfoCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetRaidInfoScRsp{
		ChallengeTakenRewardIdList: make([]uint32, 0),
		ChallengeRaidList:          make([]*proto.ChallengeRaid, 0),
		FinishedRaidInfoList:       make([]*proto.FinishedRaidInfo, 0),
		Retcode:                    0,
	}
	for _, db := range g.GetPd().GetFinishRaidMap() {
		rsp.FinishedRaidInfoList = append(rsp.FinishedRaidInfoList, &proto.FinishedRaidInfo{
			WorldLevel: db.HardLevel,
			RaidId:     db.RaidId,
			// CKJLBFCBDDB: make([]uint32, 0),
		})
	}

	g.Send(cmd.GetRaidInfoScRsp, rsp)
}

func (g *GamePlayer) StartRaidCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartRaidCsReq)
	rsp := &proto.StartRaidScRsp{}
	g.GetPd().NewRaidInfoDb(req.RaidId) // 重置
	conf := gdconf.GetRaidConfig(req.RaidId, req.WorldLevel)
	if conf == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_RECENT_ELEMENT_STAGE_NOT_MATCH)
		g.Send(cmd.StartRaidScRsp, rsp)
		return
	}
	// 调整队伍
	avatarList := g.GetPd().SetRaidLineUp(req, conf)
	g.SetBattleLineUp(model.Raid, avatarList)
	rsp.Retcode = uint32(g.GetPd().NewRaidInfo(req))
	// 检查任务
	g.InspectMission(nil)
	// 设置状态
	db := g.GetPd().GetRaidInfo(req.RaidId)
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_RAID)
	g.RaidInfoNotify(req.RaidId)
	if g.RaidEnterSceneByServerScNotify(db.EntryId) {
		finishSubMission := g.GetPd().EnterMapByEntrance(db.EntryId) // 任务检查
		if len(finishSubMission) != 0 {
			g.InspectMission(finishSubMission)
		}
	}

	g.Send(cmd.StartRaidScRsp, rsp)
}

func (g *GamePlayer) LeaveRaidCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.LeaveRaidCsReq)
	rsp := &proto.LeaveRaidScRsp{}
	db := g.GetPd().GetFinishRaidInfo(req.RaidId)
	if db == nil {
		g.SceneByServerScNotify(g.GetPd().GetCurEntryId(), g.GetPd().GetPosPb(), g.GetPd().GetRotPb())
		g.Send(cmd.LeaveRaidScRsp, rsp)
		return
	}
	var teleportToAnchor = true
	// 设置状态
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
	conf := gdconf.GetRaidConfig(db.RaidId, db.HardLevel)
	g.GetPd().NewRaidInfoDb(req.RaidId) // 重置
	if conf == nil {
		g.Send(cmd.LeaveRaidScRsp, rsp)
		return
	}
	switch conf.Type {
	case constant.RaidConfigTypeMission:
		if (conf.MainMissionIDBefore != conf.MainMissionIDAfter) || conf.MainMissionIDBefore == 0 {
			teleportToAnchor = false
		}
	}

	if teleportToAnchor {
		// 回到之前的位置
		g.SceneByServerScNotify(g.GetPd().GetCurEntryId(), g.GetPd().GetPosPb(), g.GetPd().GetRotPb())
	}

	// 任务检查
	g.InspectMission(nil)
	g.Send(cmd.LeaveRaidScRsp, rsp)
}

func (g *GamePlayer) RaidInfoNotify(raidID uint32) {
	db := g.GetPd().GetRaidInfo(raidID)
	if db == nil {
		return
	}
	allSync := &model.AllPlayerSync{IsBasic: true, MaterialList: make([]uint32, 0)}
	notify := &proto.RaidInfoNotify{
		ItemList:       &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		RaidId:         db.RaidId,
		Status:         proto.RaidStatus(db.Status),
		WorldLevel:     db.HardLevel,
		RaidFinishTime: db.FinishTime,
	}
	// TODO 有重复领取的问题，db加个字段就行了
	if db.Status == spb.RaidStatus_RAID_STATUS_FINISH {
		notify.ItemList.ItemList = g.GetPd().RaidReward(db.RaidId, db.HardLevel, allSync)
	}

	g.AllPlayerSyncScNotify(allSync)

	g.Send(cmd.RaidInfoNotify, notify)
}

func (g *GamePlayer) RaidEnterSceneByServerScNotify(entryId uint32) bool {
	rsp := new(proto.EnterSceneByServerScNotify)
	mapEntrance := gdconf.GetMapEntranceById(entryId)
	if mapEntrance == nil {
		return false
	}
	teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if teleportsMap == nil {
		return false
	}
	anchorID := mapEntrance.StartAnchorID
	groupID := mapEntrance.StartGroupID

	var pos *proto.Vector
	var rot *proto.Vector

	// 获取队伍
	lineDb := g.GetPd().GetBattleLineUpById(model.Raid)
	rsp.Lineup = g.GetPd().GetLineUpPb(lineDb)
	// 获取坐标
	var anchor *gdconf.AnchorList
	if anchorID == 0 || groupID == 0 {
		anchor = gdconf.GetAnchorByIndex(mapEntrance.PlaneID, mapEntrance.FloorID)
	} else {
		anchor = gdconf.GetAnchor(mapEntrance.PlaneID, mapEntrance.FloorID, groupID, anchorID)
	}

	if anchor != nil {
		pos = &proto.Vector{
			X: int32(anchor.PosX * 1000),
			Y: int32(anchor.PosY * 1000),
			Z: int32(anchor.PosZ * 1000),
		}
		rot = &proto.Vector{
			X: int32(anchor.RotX * 1000),
			Y: int32(anchor.RotY * 1000),
			Z: int32(anchor.RotZ * 1000),
		}
	}

	if pos == nil {
		// 这都没有那就不要传送了
		logger.Error("raid entryId:%v error", entryId)
		g.Send(cmd.EnterSceneByServerScNotify, rsp)
		return false
	}
	rsp.Scene = g.GetPd().GetRaidSceneInfo(entryId, pos, rot, lineDb)
	// g.SetCurEntryId(entryId)
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
	return true
}
