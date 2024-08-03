package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type AllPlayerSync struct {
	IsBasic                bool     // 基本信息
	AvatarList             []uint32 // 角色列表
	MaterialList           []uint32 // 物品id列表
	EquipmentList          []uint32 // 光锥列表
	DelEquipmentList       []uint32 // 删除列表
	RelicList              []uint32 // 圣遗物列表
	DelRelicList           []uint32 // 删除列表
	MissionFinishMainList  []uint32 // 已完成的主任务
	MissionFinishSubList   []uint32 // 已完成的子任务
	MissionProgressSubList []uint32 // 需要通知的子任务
}

// 玩家ping包处理
func (g *GamePlayer) HandlePlayerHeartBeatCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PlayerHeartBeatCsReq)
	sTime := getCurTime()

	rsp := new(proto.PlayerHeartBeatScRsp)
	rsp.ServerTimeMs = sTime
	rsp.ClientTimeMs = req.ClientTimeMs

	logger.Info("[UID:%v]PlayerHeartBeatCsReq", g.Uid)

	g.Send(cmd.PlayerHeartBeatScRsp, rsp)
}

func (g *GamePlayer) GetSpringRecoverDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetSpringRecoverDataScRsp)
	rsp.SpringRecoverConfig = g.GetSpringRecoverConfig()
	rsp.HealPoolInfo = g.GetHealPoolInfo()
	g.Send(cmd.GetSpringRecoverDataScRsp, rsp)
}

// 角色状态改变时需要发送通知
func (g *GamePlayer) PlayerPlayerSyncScNotify() {
	db := g.GetMaterialMap()
	notify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   g.GetNickname(),
			Level:      g.GetLevel(),
			Exp:        db[Exp],
			Hcoin:      db[Hcoin],
			Scoin:      db[Scoin],
			Mcoin:      db[Mcoin],
			Stamina:    db[Stamina],
			WorldLevel: g.GetWorldLevel(),
		},
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) AddTrailblazerExp(num uint32) {
	material := g.GetMaterialMap()
	db := g.GetBasicBin()
	material[Exp] += num
	level, exp, worldLevel := gdconf.GetPlayerLevelConfigByLevel(material[Exp], g.GetLevel(), g.GetWorldLevel())
	material[Exp] = exp
	db.Level = level
	db.WorldLevel = worldLevel
	g.PlayerPlayerSyncScNotify()
}

func (g *GamePlayer) SetPlayerInfoCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetPlayerInfoCsReq)

	g.SetNickname(req.Nickname)
	main := g.GetAvatarBinById(8001)

	if req.IsModify {
		if req.Gender == proto.Gender_GenderWoman {
			g.Send(cmd.AvatarPathChangedNotify, &proto.AvatarPathChangedNotify{
				CurMultiPathAvatarType: proto.MultiPathAvatarType_GirlWarriorType,
				BaseAvatarId:           8001,
			})
			db := g.GetAvatar()
			db.Gender = spb.Gender_GenderWoman
			main.CurPath = uint32(spb.HeroBasicType_GirlWarrior)
			g.AddMultiPathAvatar(uint32(spb.HeroBasicType_GirlWarrior))
			g.AvatarPlayerSyncScNotify(8001)
		}
		g.CreateCharacterSubMission()
	}
	rsp := &proto.SetPlayerInfoScRsp{
		Retcode:       0,
		CurAvatarPath: proto.MultiPathAvatarType(main.CurPath),
		IsModify:      req.IsModify,
	}
	g.PlayerPlayerSyncScNotify() // 角色信息通知
	g.Send(cmd.SetPlayerInfoScRsp, rsp)
}

func (g *GamePlayer) AllPlayerSyncScNotify(allSync *AllPlayerSync) {
	if allSync == nil {
		return
	}
	if !allSync.IsBasic &&
		len(allSync.AvatarList) == 0 &&
		len(allSync.MaterialList) == 0 &&
		len(allSync.EquipmentList) == 0 &&
		len(allSync.DelEquipmentList) == 0 &&
		len(allSync.RelicList) == 0 &&
		len(allSync.DelRelicList) == 0 &&
		len(allSync.MissionFinishMainList) == 0 &&
		len(allSync.MissionFinishSubList) == 0 &&
		len(allSync.MissionProgressSubList) == 0 {
		return
	}

	notify := &proto.PlayerSyncScNotify{
		AvatarSync:              &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		MultiPathAvatarInfoList: make([]*proto.MultiPathAvatarInfo, 0),
		MaterialList:            make([]*proto.Material, 0),
		EquipmentList:           make([]*proto.Equipment, 0),
		DelEquipmentList:        make([]uint32, 0),
		DelRelicList:            make([]uint32, 0),
		RelicList:               make([]*proto.Relic, 0),
		MissionSync: &proto.MissionSync{
			MissionList:               make([]*proto.Mission, 0),
			FinishedMainMissionIdList: make([]uint32, 0),
		},
	}
	db := g.GetMaterialMap()
	// 添加账户基本信息
	if allSync.IsBasic {
		notify.BasicInfo = &proto.PlayerBasicInfo{
			Nickname:   g.GetNickname(),
			Level:      g.GetLevel(),
			Exp:        db[Exp],
			Hcoin:      db[Hcoin],
			Scoin:      db[Scoin],
			Mcoin:      db[Mcoin],
			Stamina:    db[Stamina],
			WorldLevel: g.GetWorldLevel(),
		}
	}
	// 添加角色
	if allSync.AvatarList != nil {
		for _, avatarId := range allSync.AvatarList {
			avatarDb := g.GetAvatarById(avatarId)
			if avatarDb == nil {
				continue
			}
			if avatarDb.IsMultiPath {
				notify.MultiPathAvatarInfoList = append(notify.MultiPathAvatarInfoList, g.GetMultiPathAvatarInfo(avatarId)...)
			}
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, g.GetProtoAvatarById(avatarId))

		}
	}
	// 添加物品
	if allSync.MaterialList != nil {
		for _, materialId := range allSync.MaterialList {
			if materialId == Exp {
				continue
			}
			notify.MaterialList = append(notify.MaterialList, &proto.Material{
				Tid: materialId,
				Num: db[materialId],
			})
		}
	}
	// 添加光锥
	if allSync.EquipmentList != nil {
		for _, uniqueId := range allSync.EquipmentList {
			notify.EquipmentList = append(notify.EquipmentList, g.GetEquipment(uniqueId))
		}
	}
	// 删除光锥
	notify.DelEquipmentList = allSync.DelEquipmentList
	// 添加圣遗物
	if allSync.RelicList != nil {
		for _, uniqueId := range allSync.RelicList {
			notify.RelicList = append(notify.RelicList, g.GetRelic(uniqueId))
		}
	}
	// 删除圣遗物
	notify.DelRelicList = allSync.DelRelicList
	// 添加完成的主任务
	notify.MissionSync.FinishedMainMissionIdList = allSync.MissionFinishMainList
	// 添加需要通知的子任务
	if allSync.MissionProgressSubList != nil {
		subMissionList := g.GetSubMainMissionList()
		for _, subId := range allSync.MissionProgressSubList {
			if dbSub := subMissionList[subId]; dbSub != nil {
				notify.MissionSync.MissionList = append(notify.MissionSync.MissionList, &proto.Mission{
					Id:       dbSub.MissionId,
					Progress: dbSub.Progress,
					Status:   proto.MissionStatus(dbSub.Status),
				})
			} else {
				logger.Error("subMission db error id:", subId)
			}
		}
	}
	// 添加完成的子任务
	if allSync.MissionFinishSubList != nil {
		subMissionList := g.GetFinishSubMainMissionList()
		for _, subId := range allSync.MissionFinishSubList {
			if dbSub := subMissionList[subId]; dbSub != nil {
				notify.MissionSync.MissionList = append(notify.MissionSync.MissionList, &proto.Mission{
					Id:       dbSub.MissionId,
					Progress: dbSub.Progress,
					Status:   proto.MissionStatus(dbSub.Status),
				})
			} else {
				logger.Error("finishSubMission db error id:", subId)
			}
		}
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}
