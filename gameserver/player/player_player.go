package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/lua"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

// 玩家ping包处理
func HandlePlayerHeartBeatCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PlayerHeartBeatCsReq)
	sTime := getCurTime()

	rsp := new(proto.PlayerHeartBeatScRsp)
	rsp.ServerTimeMs = sTime
	rsp.ClientTimeMs = req.ClientTimeMs

	for _, v := range lua.GetPingLua() {
		g.ClientDownloadDataScNotify(v)
	}

	g.Send(cmd.PlayerHeartBeatScRsp, rsp)
}

func SetPlayerInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetPlayerInfoCsReq)

	g.GetPd().SetNickname(req.Nickname)
	main := g.GetPd().GetAvatarBinById(8001)
	allSync := &model.AllPlayerSync{IsBasic: true, AvatarList: make([]uint32, 0)}

	if req.IsModify {
		if req.Gender == proto.Gender_GenderWoman {
			g.Send(cmd.AvatarPathChangedNotify, &proto.AvatarPathChangedNotify{
				CurMultiPathAvatarType: proto.MultiPathAvatarType_GirlWarriorType,
				BaseAvatarId:           8001,
			})
			db := g.GetPd().GetAvatar()
			db.Gender = spb.Gender_GenderWoman
			main.CurPath = uint32(spb.HeroBasicType_GirlWarrior)
			g.GetPd().AddMultiPathAvatar(uint32(spb.HeroBasicType_GirlWarrior))
			allSync.AvatarList = append(allSync.AvatarList, 8001)
		}
		finishSubMission := g.GetPd().CreateCharacterSubMission()
		if len(finishSubMission) > 0 {
			g.InspectMission(finishSubMission...)
		}
	}
	rsp := &proto.SetPlayerInfoScRsp{
		Retcode:       0,
		CurAvatarPath: proto.MultiPathAvatarType(main.CurPath),
		IsModify:      req.IsModify,
	}
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.SetPlayerInfoScRsp, rsp)
}

func (g *GamePlayer) AllPlayerSyncScNotify(allSync *model.AllPlayerSync) {
	if allSync == nil {
		return
	}
	if !allSync.IsBasic &&
		len(allSync.UnlockPamSkinList) == 0 &&
		len(allSync.AvatarList) == 0 &&
		len(allSync.UnlockedHeadIconList) == 0 &&
		len(allSync.MaterialList) == 0 &&
		len(allSync.EquipmentList) == 0 &&
		len(allSync.DelEquipmentList) == 0 &&
		len(allSync.RelicList) == 0 &&
		len(allSync.DelRelicList) == 0 &&
		len(allSync.MissionFinishMainList) == 0 &&
		len(allSync.MissionFinishSubList) == 0 &&
		len(allSync.MissionProgressSubList) == 0 &&
		len(allSync.MessageGroupList) == 0 &&
		len(allSync.MessageSectionList) == 0 {
		return
	}

	if len(allSync.UnlockPamSkinList) != 0 {
		g.UnlockPamSkinScNotify(allSync.UnlockPamSkinList)
	}
	syncStatus := &proto.SyncStatus{
		SectionStatus:      make([]*proto.SectionStatus, 0),
		MessageGroupStatus: make([]*proto.GroupStatus, 0),
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
		PlayerboardModuleSync: &proto.PlayerBoardModuleSync{
			UnlockedHeadIconList: make([]*proto.HeadIconData, 0),
		},
		SyncStatus: syncStatus,
	}
	db := g.GetPd().GetMaterialMap()
	// 添加账户基本信息
	if allSync.IsBasic {
		notify.BasicInfo = &proto.PlayerBasicInfo{
			Nickname:   g.GetPd().GetNickname(),
			Level:      g.GetPd().GetLevel(),
			Exp:        db[model.Exp],
			Hcoin:      db[model.Hcoin],
			Scoin:      db[model.Scoin],
			Mcoin:      db[model.Mcoin],
			Stamina:    db[model.Stamina],
			WorldLevel: g.GetPd().GetWorldLevel(),
		}
	}
	// 添加角色
	if allSync.AvatarList != nil {
		for _, avatarId := range allSync.AvatarList {
			avatarDb := g.GetPd().GetAvatarBinById(avatarId)
			if avatarDb == nil {
				continue
			}
			if avatarDb.IsMultiPath {
				notify.MultiPathAvatarInfoList = append(notify.MultiPathAvatarInfoList,
					g.GetPd().GetMultiPathAvatarInfo(avatarId)...)
			}
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList,
				g.GetPd().GetProtoAvatarById(avatarId))

		}
	}
	// 添加头像
	if len(allSync.UnlockedHeadIconList) > 0 {
		for _, headIconId := range allSync.UnlockedHeadIconList {
			notify.PlayerboardModuleSync.UnlockedHeadIconList = append(notify.PlayerboardModuleSync.UnlockedHeadIconList,
				&proto.HeadIconData{Id: headIconId})
		}
	}
	// 添加物品
	if allSync.MaterialList != nil {
		for _, materialId := range allSync.MaterialList {
			if materialId == model.Exp {
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
			notify.EquipmentList = append(notify.EquipmentList, g.GetPd().GetProtoEquipment(uniqueId))
		}
	}
	// 删除光锥
	notify.DelEquipmentList = allSync.DelEquipmentList
	// 添加圣遗物
	if allSync.RelicList != nil {
		for _, uniqueId := range allSync.RelicList {
			notify.RelicList = append(notify.RelicList, g.GetPd().GetProtoRelicById(uniqueId))
		}
	}
	// 删除圣遗物
	notify.DelRelicList = allSync.DelRelicList
	// 添加完成的主任务
	notify.MissionSync.FinishedMainMissionIdList = allSync.MissionFinishMainList
	// 添加需要通知的子任务
	if allSync.MissionProgressSubList != nil {
		subMissionList := g.GetPd().GetSubMainMissionList()
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
		subMissionList := g.GetPd().GetFinishSubMainMissionList()
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
	// 添加对话组
	if allSync.MessageGroupList != nil {
		for _, messageGroupId := range allSync.MessageGroupList {
			group := g.GetPd().GetMessageGroupByContactId(messageGroupId)
			if group != nil {
				syncStatus.MessageGroupStatus = append(syncStatus.MessageGroupStatus, &proto.GroupStatus{
					GroupStatus: proto.MessageGroupStatus(group.Status),
					RefreshTime: group.RefreshTime,
					GroupId:     group.Id,
				})
			}
		}
	}

	// 添加对话
	if allSync.MessageSectionList != nil {
		for _, sectionId := range allSync.MessageSectionList {
			info := g.GetPd().GetMessageSection(sectionId)
			if info != nil {
				syncStatus.SectionStatus = append(syncStatus.SectionStatus, &proto.SectionStatus{
					SectionId:     info.Id,
					SectionStatus: proto.MessageSectionStatus(info.Status),
				})
			}
		}
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) AllScenePlaneEventScNotify(addPileItem []*model.Material) {
	if addPileItem == nil || len(addPileItem) == 0 {
		return
	}
	// 通知客户端增加了物品
	notify := &proto.ScenePlaneEventScNotify{
		GetItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	itemList := make([]*proto.Item, 0)
	// 添加物品
	for _, item := range addPileItem {
		itemList = append(itemList, &proto.Item{
			ItemId: item.Tid,
			Num:    item.Num,
		})
	}
	notify.GetItemList.ItemList = itemList

	g.Send(cmd.ScenePlaneEventScNotify, notify)
}

func (g *GamePlayer) UnlockPamSkinScNotify(pamSkinList []uint32) {
	for _, v := range pamSkinList {
		notify := &proto.UnlockPamSkinScNotify{
			PamSkinId: v,
		}
		g.Send(cmd.UnlockPamSkinScNotify, notify)
	}

}
