package player

import (
	"strings"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

// 通知客户端进入场景
func (g *GamePlayer) EnterSceneByServerScNotify(entryId, teleportId, groupID, anchorID uint32) {
	if entryId == 0 {
		entryId = g.GetPd().GetCurEntryId()
	}
	rsp := new(proto.EnterSceneByServerScNotify)
	mapEntrance := gdconf.GetMapEntranceById(entryId)
	if mapEntrance == nil {
		return
	}
	teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if teleportsMap == nil {
		return
	}

	if anchorID == 0 || groupID == 0 {
		anchorID = mapEntrance.StartAnchorID
		groupID = mapEntrance.StartGroupID
	}

	var pos *proto.Vector
	var rot *proto.Vector

	// 获取队伍
	curLine := g.GetPd().GetCurLineUp()
	rsp.Lineup = g.GetPd().GetLineUpPb(curLine)
	// 获取坐标
	var anchor *gdconf.AnchorList
	if teleportsMap.Teleports[teleportId] != nil {
		anchorID = teleportsMap.Teleports[teleportId].AnchorID
		groupID = teleportsMap.Teleports[teleportId].AnchorGroupID
		anchor = teleportsMap.TeleportsByGroupId[groupID].AnchorList[anchorID]
	} else {
		if teleportId == 0 {
			if anchorID == 0 || groupID == 0 {
				anchor = gdconf.GetAnchorByIndex(mapEntrance.PlaneID, mapEntrance.FloorID)
			} else {
				anchor = gdconf.GetAnchor(mapEntrance.PlaneID, mapEntrance.FloorID, groupID, anchorID)
			}
		} else {
			confMi := gdconf.GetMappingInfoById(teleportId, 0)
			if confMi != nil {
				group := gdconf.GetNGroupById(confMi.PlaneID, confMi.FloorID, confMi.GroupID)
				anchor = group.AnchorList[0] // 相当于随机取了
			}
		}
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
		logger.Error("entryId:%v,teleportId:%v error", entryId, teleportId)
		g.Send(cmd.EnterSceneByServerScNotify, rsp)
		return
	}
	if db := g.GetPd().GetCurChangeStoryInfo(); db != nil {
		g.StoryLineInfoScNotify()
		g.SyncLineupNotify(g.GetPd().GetCurLineUp())
		g.Send(cmd.SyncServerSceneChangeNotify, &proto.SyncServerSceneChangeNotify{})
	}
	rsp.Scene = g.GetPd().GetSceneInfo(entryId, pos, rot, curLine)
	finishSubMission := g.GetPd().EnterMapByEntrance(entryId) // 任务检查
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
	g.ChangeStoryLineFinishScNotify()
}

// 传送到指定位置
func (g *GamePlayer) SceneByServerScNotify(entryId uint32, pos, rot *proto.Vector) {
	rsp := new(proto.EnterSceneByServerScNotify)
	// 获取队伍
	curLine := g.GetPd().GetCurLineUp()
	rsp.Lineup = g.GetPd().GetLineUpPb(curLine)
	rsp.Scene = g.GetPd().GetSceneInfo(entryId, pos, rot, curLine)

	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *GamePlayer) HandleGetEnteredSceneCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetEnteredSceneScRsp)
	db := g.GetPd().GetScene()
	mapEntrance := gdconf.GetMapEntranceById(db.EntryId)
	if mapEntrance == nil {
		return
	}
	enteredSceneInfo := &proto.EnteredSceneInfo{
		FloorId: mapEntrance.FloorID,
		PlaneId: mapEntrance.PlaneID,
	}
	rsp.EnteredSceneInfoList = []*proto.EnteredSceneInfo{enteredSceneInfo}

	g.Send(cmd.GetEnteredSceneScRsp, rsp)
}

// 客户端登录需要的包，不是传送的通知包
func (g *GamePlayer) HandleGetCurSceneInfoCsReq(payloadMsg pb.Message) {
	pos := g.GetPd().GetPosPb()
	rot := g.GetPd().GetRotPb()
	dbScene := g.GetPd().GetScene()
	curLine := g.GetPd().GetCurLineUp()

	rsp := new(proto.GetCurSceneInfoScRsp)
	rsp.Scene = g.GetPd().GetSceneInfo(dbScene.EntryId, pos, rot, curLine)
	g.Send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *GamePlayer) HanldeGetSceneMapInfoCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetSceneMapInfoCsReq)

	rsp := new(proto.GetSceneMapInfoScRsp)
	for _, entryId := range req.EntryIdList {
		mapEntrance := gdconf.GetMapEntranceById(entryId)
		if mapEntrance != nil {
			groupList := gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)
			block := g.GetPd().GetBlock(entryId)
			if groupList != nil {
				mapList := &proto.SceneMapInfo{
					LightenSectionList: make([]uint32, 0),
					ChestList: []*proto.ChestInfo{
						{ChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_NORMAL, OpenedNum: 1},
						{ChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_CHALLENGE, OpenedNum: 1},
						{ChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_PUZZLE, OpenedNum: 1},
					},
					UnlockTeleportList: make([]uint32, 0),
					DimensionId:        g.GetPd().GetDimensionId(),
					EntryStoryLineId:   req.EntryStoryLineId,
					FloorSavedData:     g.GetPd().GetFloorSavedData(entryId),
					EntryId:            entryId,
				}

				for i := uint32(0); i < 100; i++ {
					mapList.LightenSectionList = append(mapList.LightenSectionList, i)
				}

				for _, group := range groupList {
					mapList.MazeGroupList = append(mapList.MazeGroupList, &proto.MazeGroup{GroupId: group.GroupId})
					for _, prop := range group.PropList {
						if prop.MappingInfoID != 0 {
							mazeProp := &proto.MazePropState{
								State:    g.GetPd().GetPropState(block, group.GroupId, prop.ID, prop.State),
								GroupId:  group.GroupId,
								ConfigId: prop.ID,
							}
							mapList.MazePropList = append(mapList.MazePropList, mazeProp)
							mapList.UnlockTeleportList = append(mapList.UnlockTeleportList, prop.MappingInfoID)
						}
					}
				}
				rsp.SceneMapInfo = append(rsp.SceneMapInfo, mapList)
			}
		}
	}

	g.Send(cmd.GetSceneMapInfoScRsp, rsp)
}

func (g *GamePlayer) EnterSceneCsReq(payloadMsg pb.Message) {
	g.ContentPackageSyncDataScNotify()
	req := payloadMsg.(*proto.EnterSceneCsReq)
	rsp := &proto.EnterSceneScRsp{
		ContentId:       req.ContentId,
		GameStoryLineId: req.GameStoryLineId,
		IsCloseMap:      req.IsCloseMap,
		Retcode:         0,
		IsOverMap:       true,
	}
	entryId := req.EntryId
	teleportId := req.TeleportId
	var groupId uint32 = 0
	var anchorId uint32 = 0

	if conf := gdconf.GetStoryLine(req.GameStoryLineId); conf != nil {
		changeStory := g.GetPd().GetChangeStoryInfo(req.GameStoryLineId)
		if changeStory == nil {
			entryIds, anchorGroups, anchorIds, ok := g.GetPd().MissionAddChangeStoryLine(
				[]uint32{req.GameStoryLineId, 0, 0, 0})
			if ok {
				g.EnterSceneByServerScNotify(entryIds, 0, anchorGroups, anchorIds)
			}
			return
		}
		db := g.GetPd().GetChangeStory()
		db.IsChangeStory = true
		db.CurChangeStory = req.GameStoryLineId
		entryId = changeStory.Scene.EntryId
		groupId = changeStory.Scene.GroupId
		anchorId = changeStory.Scene.AnchorId
		g.GetPd().NewStoryLine(req.GameStoryLineId)
	} else if req.GameStoryLineId == 0 {
		db := g.GetPd().GetChangeStory()
		db.IsChangeStory = false
		if entryId == 0 {
			entryId = g.GetPd().GetCurEntryId()
		}
	}

	if entryId != 0 &&
		req.ContentId == 0 &&
		req.GameStoryLineId == 0 {
		g.GetPd().SetCurEntryId(entryId)
	}

	if req.ContentId == 0 {
		g.EnterSceneByServerScNotify(entryId, teleportId, groupId, anchorId)
	} else {

	}

	g.Send(cmd.EnterSceneScRsp, rsp)
	// g.Send(cmd.SceneUpdatePositionVersionNotify, rsp)
}

func (g *GamePlayer) InteractPropCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.InteractPropCsReq)
	rsp := &proto.InteractPropScRsp{
		Retcode:      0,
		PropState:    0,
		PropEntityId: req.PropEntityId,
	}
	var pileItem []*model.Material
	var propEntityIdList []uint32
	allSync := &model.AllPlayerSync{
		IsBasic:      true,
		MaterialList: make([]uint32, 0),
	}

	pe := g.GetPd().GetPropEntityById(req.PropEntityId)
	if pe == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	blockBin := g.GetPd().GetBlock(pe.EntryId)
	mapEntrance := gdconf.GetMapEntranceById(blockBin.EntryId)
	if mapEntrance == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	confProp := gdconf.GetServerPropById(mapEntrance.PlaneID, mapEntrance.FloorID, pe.GroupId, pe.InstId)
	if confProp == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	confInteract := gdconf.GetInteractConfigById(req.InteractId)
	if confInteract == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	oldState := g.GetPd().GetPropState(blockBin, pe.GroupId, pe.InstId, confProp.State)
	newState := gdconf.GetStateValue(confInteract.TargetState)
	setState := newState
	mazeProp := gdconf.GetMazePropId(pe.PropId)
	if mazeProp == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	switch mazeProp.PropType {
	case gdconf.PROP_TREASURE_CHEST: // 宝箱
		if oldState == 12 && newState == 13 {
			pileItem = append(pileItem, &model.Material{Tid: model.Mcoin, Num: 1000})
		}
	case gdconf.PROP_DESTRUCT: // 破坏物
		if newState == 0 {
			setState = 1
		}
	case gdconf.PROP_MAZE_PUZZLE: // 拼图
	}

	if strings.Contains(confProp.Name, "Console") {
		setState = oldState
	}

	if confProp.GoppValue != nil {
		for _, goppValue := range confProp.GoppValue {
			if enep := g.GetPd().GetPropEntity(goppValue.GroupId, goppValue.InstId); enep != nil {
				g.GetPd().UpPropState(blockBin, goppValue.GroupId, goppValue.InstId, 1) // 更新状态
				propEntityIdList = append(propEntityIdList, enep.EntityId)
			}
		}
	}

	propEntityIdList = append(propEntityIdList, req.PropEntityId)
	g.GetPd().UpPropState(blockBin, pe.GroupId, pe.InstId, setState) // 更新地图
	rsp.PropState = setState
	// 统一通知
	g.GetPd().AddItem(pileItem, allSync)
	g.AllPlayerSyncScNotify(allSync)
	g.AllScenePlaneEventScNotify(pileItem)
	g.PropSceneGroupRefreshScNotify(propEntityIdList, blockBin)  // 通知状态更改
	finishSubMission := g.GetPd().UpInteractSubMission(blockBin) // 检查交互任务
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
	g.Send(cmd.InteractPropScRsp, rsp)
}

func (g *GamePlayer) GroupStateChangeCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GroupStateChangeCsReq)
	rsp := &proto.GroupStateChangeScRsp{
		GroupStateInfo: req.GroupStateInfo,
		Retcode:        0,
	}
	if req.GroupStateInfo != nil {
		blockBin := g.GetPd().GetBlock(req.GroupStateInfo.EntryId)
		mapEntrance := gdconf.GetMapEntranceById(blockBin.EntryId)
		if mapEntrance == nil {
			g.Send(cmd.GroupStateChangeScRsp, rsp)
			return
		}
		confGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, req.GroupStateInfo.GroupId)
		if confGroup == nil || confGroup.PropList == nil {
			g.Send(cmd.GroupStateChangeScRsp, rsp)
			return
		}
		var propEntityIdList []uint32
		for _, prop := range confGroup.PropList {
			if strings.Contains(prop.Name, "Console") {
				if enep := g.GetPd().GetPropEntity(confGroup.GroupId, prop.ID); enep != nil {
					g.GetPd().UpPropState(blockBin, confGroup.GroupId, prop.ID, 1) // 更新状态
					propEntityIdList = append(propEntityIdList, enep.EntityId)
				}
			}
		}
		g.PropSceneGroupRefreshScNotify(propEntityIdList, blockBin) // 通知状态更改
	}
	g.Send(cmd.GroupStateChangeScNotify, &proto.GroupStateChangeScNotify{
		GroupStateInfo: req.GroupStateInfo,
	})

	g.Send(cmd.GroupStateChangeScRsp, rsp)
}

func (g *GamePlayer) DeployRotaterCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.DeployRotaterCsReq)
	// 设置旋转
	rsp := &proto.DeployRotaterScRsp{
		Retcode:     0,
		EnergyInfo:  nil,
		RotaterData: req.RotaterData,
	}
	g.Send(cmd.DeployRotaterScRsp, rsp)
}

func (g *GamePlayer) StartWolfBroGameCsReq(payloadMsg pb.Message) {
	g.Send(cmd.WolfBroGameDataChangeScNotify, &proto.WolfBroGameDataChangeScNotify{
		WolfBroGameData: &proto.WolfBroGameData{
			// KHOGNFEGNLC: &proto.WolfBroGameInfo{
			// 	Motion: &proto.MotionInfo{
			// 		Pos: g.GetPosPb(),
			// 		Rot: g.GetRotPb(),
			// 	},
			// 	BOLDFGOJGII: 0,
			// 	ADLJJIGGBHE: make([]*proto.Vector, 0),
			// 	OAOLMHLHNAI: false,
			// },
			// LINLMHBEAPC: "114514",
			Id: 3,
		},
	})
	g.Send(cmd.StartWolfBroGameScRsp, &proto.StartWolfBroGameScRsp{})
}

func (g *GamePlayer) SetGroupCustomSaveDataCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetGroupCustomSaveDataCsReq)
	g.Send(cmd.GroupStateChangeScNotify, &proto.GroupStateChangeScNotify{
		GroupStateInfo: &proto.GroupStateInfo{
			EntryId:    req.EntryId,
			GroupState: 1,
			GroupId:    req.GroupId,
		},
	})
	g.Send(cmd.SetGroupCustomSaveDataScRsp, &proto.SetGroupCustomSaveDataScRsp{
		EntryId: req.EntryId,
		GroupId: req.GroupId,
		Retcode: 0,
	})
}

func (g *GamePlayer) SpringRecoverSingleAvatarCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SpringRecoverSingleAvatarCsReq)
	g.GetPd().AvatarRecover(req.Id)

	rsp := &proto.SpringRecoverSingleAvatarScRsp{
		HpFieldNumber: 10000,
		Retcode:       0,
		AvatarType:    req.AvatarType,
		Id:            req.Id,
	}
	g.SyncLineupNotify(g.GetPd().GetBattleLineUp())
	g.Send(cmd.SpringRecoverSingleAvatarScRsp, rsp)
}

// 更新实体状态
func (g *GamePlayer) PropSceneGroupRefreshScNotify(propEntityIdList []uint32, db *spb.BlockBin) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	for _, propId := range propEntityIdList {
		pe := g.GetPd().GetPropEntityById(propId)
		if pe == nil {
			continue
		}
		isAddend := true
		var info *proto.GroupRefreshInfo
		for _, ninfo := range notify.GroupRefreshList {
			if ninfo.GroupId == pe.GroupId {
				info = ninfo
				isAddend = false
				break
			}
		}
		if isAddend {
			info = new(proto.GroupRefreshInfo)
			info.GroupId = pe.GroupId
			notify.GroupRefreshList = append(notify.GroupRefreshList, info)
		}
		info.RefreshEntity = append(info.RefreshEntity, &proto.SceneEntityRefreshInfo{
			Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
				AddEntity: &proto.SceneEntityInfo{
					EntityId: pe.EntityId,
					GroupId:  pe.GroupId,
					InstId:   pe.InstId,
					Motion: &proto.MotionInfo{
						Pos: pe.Pos,
						Rot: pe.Rot,
					},
					EntityOneofCase: &proto.SceneEntityInfo_Prop{
						Prop: &proto.ScenePropInfo{
							PropId:    pe.PropId, // PropID
							PropState: g.GetPd().GetPropState(db, pe.GroupId, pe.InstId, ""),
						},
					},
				},
			},
		})
	}

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

// 添加角色
func (g *GamePlayer) AddAvatarSceneGroupRefreshScNotify(lineAvatar *spb.LineAvatarList, pos, rot *proto.Vector) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	info := new(proto.GroupRefreshInfo)
	actor := &proto.SceneActorInfo{
		AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
		BaseAvatarId: lineAvatar.AvatarId,
		MapLayer:     0,
		Uid:          0,
	}
	entityId := g.GetPd().GetNextGameObjectGuid()
	if lineAvatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
		conf := gdconf.GetSpecialAvatarById(lineAvatar.AvatarId)
		if conf == nil {
			return
		}
		actor = &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_TRIAL_TYPE,
			BaseAvatarId: conf.AvatarID,
			MapLayer:     0,
			Uid:          0,
		}
		info.RefreshType = proto.SceneGroupRefreshType_SCENE_GROUP_REFRESH_TYPE_UNLOAD
	}
	info.RefreshEntity = append(info.RefreshEntity, &proto.SceneEntityRefreshInfo{
		Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
			AddEntity: &proto.SceneEntityInfo{
				EntityId:        entityId,
				EntityOneofCase: &proto.SceneEntityInfo_Actor{Actor: actor},
				Motion: &proto.MotionInfo{
					Rot: pos,
					Pos: rot,
				},
			},
		},
	})
	g.GetPd().AddEntity(0, &model.AvatarEntity{
		Entity: model.Entity{
			EntityId: entityId,
			GroupId:  0,
			Pos:      pos,
			Rot:      rot,
		},
		AvatarId:   actor.BaseAvatarId,
		LineAvatar: lineAvatar,
	})
	notify.GroupRefreshList = append(notify.GroupRefreshList, info)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

// 卸载/加载场景
func (g *GamePlayer) UpSceneGroupRefreshScNotify(uninstallGroup, loadedGroupList []*model.GroupInfo) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}

	// 卸载
	for _, groupInfo := range uninstallGroup {
		if groupInfo.GroupID == 0 { // 不能卸载角色
			continue
		}
		db := g.GetPd().GetBlock(groupInfo.EntryId)
		groupRefreshInfo := &proto.GroupRefreshInfo{
			GroupId:       groupInfo.GroupID,
			RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
			State:         g.GetPd().GetGroupState(db, groupInfo.GroupID),
		}

		for _, entify := range groupInfo.EntityMap {
			groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity,
				&proto.SceneEntityRefreshInfo{
					Refresh: &proto.SceneEntityRefreshInfo_DeleteEntity{
						DeleteEntity: g.GetPd().GetEntryId(entify)},
				})
		}

		notify.GroupRefreshList = append(notify.GroupRefreshList, groupRefreshInfo)
	}
	// 加载
	for _, groupInfo := range loadedGroupList {
		if groupInfo.GroupID == 0 { // 不能卸载角色
			continue
		}
		group := gdconf.GetServerGroupById(groupInfo.PlaneID, groupInfo.FloorID, groupInfo.GroupID)
		if group == nil {
			continue
		}
		db := g.GetPd().GetBlock(groupInfo.EntryId)
		groupRefreshInfo := &proto.GroupRefreshInfo{
			GroupId:       groupInfo.GroupID,
			RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
			State:         g.GetPd().GetGroupState(db, groupInfo.GroupID),
		}
		// 添加怪物
		groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity,
			g.GetPd().AddMonsterSceneEntityRefreshInfo(groupInfo.GroupID, group.MonsterList)...)
		// 添加npc
		groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity,
			g.GetPd().AddNpcSceneEntityRefreshInfo(groupInfo.GroupID, group.NPCList)...)
		// 添加实体
		groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity,
			g.GetPd().AddPropSceneEntityRefreshInfo(group, groupInfo.GroupID, group.PropList, db)...)
		notify.GroupRefreshList = append(notify.GroupRefreshList, groupRefreshInfo)
		// g.UpdateBlock(db)
	}

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

func (g *GamePlayer) SyncEntityBuffChangeListScNotify(buffList []uint32) {
	notify := &proto.SyncEntityBuffChangeListScNotify{
		EntityBuffChangeList: make([]*proto.EntityBuffChangeInfo, 0),
	}
	mazeBuffMap := g.GetPd().GetMazeBuffList()
	for _, buffId := range buffList {
		if info, ok := mazeBuffMap[buffId]; ok {
			notify.EntityBuffChangeList = append(notify.EntityBuffChangeList, &proto.EntityBuffChangeInfo{
				Reason:   0,
				EntityId: 6293805,
				BuffChangeInfo: &proto.BuffInfo{
					Count:     info.LifeCount,
					AddTimeMs: info.AddTime,
					BuffId:    info.BuffId,
					Level:     info.Level,
					LifeTime:  -1,
				},
			})
		}
	}

	g.Send(cmd.SyncEntityBuffChangeListScNotify, notify)
}

// 任务设置物品状态
func (g *GamePlayer) SetFloorSavedValue(conf *gdconf.SubMission, finishAction *gdconf.FinishAction) {
	if len(finishAction.FinishActionParaString) < 4 {
		return
	}
	planeId := alg.S2U32(finishAction.FinishActionParaString[0])
	floorId := alg.S2U32(finishAction.FinishActionParaString[1])
	name := finishAction.FinishActionParaString[2]
	state := alg.S2I32(finishAction.FinishActionParaString[3])
	g.GetPd().SetFloorSavedData(model.FloorTentry(floorId), name, state)
	notify := &proto.UpdateFloorSavedValueNotify{
		SavedValue: map[string]int32{name: state},
	}
	g.Send(cmd.UpdateFloorSavedValueNotify, notify)

	db := g.GetPd().GetBlock(model.FloorTentry(conf.WayPointFloorID))
	groupID, instId := gdconf.GetSavedValue(planeId, floorId, name)
	if groupID == 0 || instId == 0 {
		return // TODO subMission 103030204
	}
	g.GetPd().UpPropState(db, groupID, instId, uint32(state))
	if enep := g.GetPd().GetPropEntity(groupID, instId); enep != nil {
		g.PropSceneGroupRefreshScNotify([]uint32{enep.EntityId}, db)
	}
}

/****************************************************领域管理***************************************************/

func (g *GamePlayer) AddSummonUnitSceneGroupRefreshScNotify() {
	db := g.GetPd().GetSummonUnitInfo()
	entityId := db.EntityId
	g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: g.GetPd().GetAddBuffSceneEntityRefreshInfo(
			db.AttachEntityId, db.SummonUnitId, entityId, db.Pos),
	})
	go func() {
		time.Sleep(20 * time.Second)
		g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
			GroupRefreshList: []*proto.GroupRefreshInfo{{
				RefreshType: proto.SceneGroupRefreshType_SCENE_GROUP_REFRESH_TYPE_UNLOAD,
				RefreshEntity: []*proto.SceneEntityRefreshInfo{{
					Refresh: &proto.SceneEntityRefreshInfo_DeleteEntity{
						DeleteEntity: entityId,
					},
				}},
			}},
		})
		g.GetPd().DelSummonUnitInfo()
	}()
}
