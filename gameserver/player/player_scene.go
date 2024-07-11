package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

// 通知客户端进入场景
func (g *GamePlayer) EnterSceneByServerScNotify(entryId, teleportId, groupID, anchorID uint32) {
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
	curLine := g.GetCurLineUp()
	rsp.Lineup = g.GetLineUpPb(curLine)
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
	rsp.Scene = g.GetSceneInfo(entryId, pos, rot, curLine)
	g.SetCurEntryId(entryId)
	g.EnterMapByEntrance(entryId) // 任务检查
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
	// "2010101"
}

// 传送到指定位置
func (g *GamePlayer) SceneByServerScNotify(entryId uint32, pos, rot *proto.Vector) {
	rsp := new(proto.EnterSceneByServerScNotify)
	// 获取队伍
	curLine := g.GetCurLineUp()
	rsp.Lineup = g.GetLineUpPb(curLine)
	rsp.Scene = g.GetSceneInfo(entryId, pos, rot, curLine)

	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *GamePlayer) HandleGetEnteredSceneCsReq(payloadMsg []byte) {
	rsp := new(proto.GetEnteredSceneScRsp)
	db := g.GetScene()
	mapEntrance := gdconf.GetMapEntranceById(db.EntryId)
	if mapEntrance == nil {
		return
	}
	enteredSceneInfo := &proto.EnteredScene{
		FloorId: mapEntrance.FloorID,
		PlaneId: mapEntrance.PlaneID,
	}
	rsp.EnteredSceneList = []*proto.EnteredScene{enteredSceneInfo}

	g.Send(cmd.GetEnteredSceneScRsp, rsp)
}

// 客户端登录需要的包，不是传送的通知包
func (g *GamePlayer) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	pos := g.GetPosPb()
	rot := g.GetRotPb()
	dbScene := g.GetScene()
	curLine := g.GetCurLineUp()

	rsp := new(proto.GetCurSceneInfoScRsp)
	rsp.Scene = g.GetSceneInfo(dbScene.EntryId, pos, rot, curLine)
	g.Send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *GamePlayer) HanldeGetSceneMapInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetSceneMapInfoCsReq, payloadMsg)
	req := msg.(*proto.GetSceneMapInfoCsReq)

	// 1000101 1000001

	rsp := new(proto.GetSceneMapInfoScRsp)
	for _, entryId := range req.EntryIdList {
		mapEntrance := gdconf.GetMapEntranceById(entryId)
		if mapEntrance != nil {
			groupList := gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)
			// teleportsMap := gdconf.GetTeleportsById(mapEntrance.PlaneID, mapEntrance.FloorID)
			// if teleportsMap != nil {
			// 	mapList := &proto.SceneMapInfo{
			// 		LightenSectionList: make([]uint32, 0),
			// 		ChestList: []*proto.ChestInfo{
			// 			{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_NORMAL},
			// 			{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_CHALLENGE},
			// 			{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_PUZZLE},
			// 		},
			// 		// UnlockedTeleportList: make([]uint32, 0),
			// 	}
			//
			// 	mapList.EntryId = entryId
			//
			// 	// for i := uint32(0); i < 100; i++ {
			// 	// 	mapList.LightenSectionList = append(mapList.LightenSectionList, i)
			// 	// }
			//
			// 	for _, teleports := range teleportsMap.TeleportsByGroupId {
			// 		mazeGroup := &proto.MazeGroup{GroupId: teleports.GroupId}
			// 		mapList.MazeGroupList = append(mapList.MazeGroupList, mazeGroup)
			// 	}
			//
			// 	for _, teleports := range teleportsMap.Teleports {
			// 		mazeProp := &proto.MazePropState{
			// 			State:    gdconf.GetStateValue("CheckPointEnable"), // 默认解锁
			// 			GroupId:  teleports.AnchorGroupID,
			// 			ConfigId: teleports.ID,
			// 		}
			// 		mapList.MazePropList = append(mapList.MazePropList, mazeProp)
			// 		mapList.UnlockTeleportList = append(mapList.UnlockTeleportList, teleports.MappingInfoID)
			// 	}
			// 	rsp.SceneMapInfo = append(rsp.SceneMapInfo, mapList)
			// }

			if groupList != nil {
				mapList := &proto.SceneMapInfo{
					LightenSectionList: make([]uint32, 0),
					ChestList: []*proto.ChestInfo{
						{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_NORMAL},
						{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_CHALLENGE},
						{MapInfoChestType: proto.ChestType_MAP_INFO_CHEST_TYPE_PUZZLE},
					},
					UnlockTeleportList: make([]uint32, 0),
				}

				mapList.EntryId = entryId

				for i := uint32(0); i < 100; i++ {
					mapList.LightenSectionList = append(mapList.LightenSectionList, i)
				}

				for _, group := range groupList {
					mapList.MazeGroupList = append(mapList.MazeGroupList, &proto.MazeGroup{GroupId: group.GroupId})
					for _, prop := range group.PropList {
						if prop.MappingInfoID != 0 {
							var state uint32 = 0
							if prop.AnchorID != 0 {
								state = 8 // 默认解锁
							} else {
								state = 1 // gdconf.GetStateValue(prop.State)
							}
							mazeProp := &proto.MazePropState{
								State:    state,
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

func (g *GamePlayer) EnterSceneCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.EnterSceneCsReq, payloadMsg)
	req := msg.(*proto.EnterSceneCsReq)
	rsp := &proto.GetEnteredSceneScRsp{}

	g.EnterSceneByServerScNotify(req.EntryId, req.TeleportId, 0, 0)

	g.Send(cmd.EnterSceneScRsp, rsp)
	g.Send(cmd.SceneUpdatePositionVersionNotify, rsp)
}

func (g *GamePlayer) InteractPropCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.InteractPropCsReq, payloadMsg)
	req := msg.(*proto.InteractPropCsReq)
	rsp := &proto.InteractPropScRsp{
		Retcode:      0,
		PropState:    0,
		PropEntityId: req.PropEntityId,
	}
	var propEntityIdList []uint32

	pe := g.GetPropEntityById(req.PropEntityId)
	if pe == nil {
		g.Send(cmd.InteractPropScRsp, rsp)
		return
	}
	blockBin := g.GetBlock(pe.EntryId)
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
	oldState := g.GetPropState(blockBin, pe.GroupId, pe.InstId, confProp.State)
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
			g.AddMaterial([]*Material{{Tid: Mcoin, Num: 1000}})
			g.AllPlayerSyncScNotify(&AllPlayerSync{
				MaterialList: []uint32{Mcoin},
			})
		}
	case gdconf.PROP_DESTRUCT: // 破坏物
		if newState == 0 {
			setState = 1
		}
	case gdconf.PROP_MAZE_PUZZLE: // 拼图
	}

	if confProp.GoppValue != nil {
		for _, goppValue := range confProp.GoppValue {
			if enep := g.GetPropEntity(goppValue.GroupId, goppValue.InstId); enep != nil {
				g.UpPropState(blockBin, goppValue.GroupId, goppValue.InstId, 1) // 更新状态
				propEntityIdList = append(propEntityIdList, enep.EntityId)
			}
		}
	}

	propEntityIdList = append(propEntityIdList, req.PropEntityId)
	g.UpPropState(blockBin, pe.GroupId, pe.InstId, setState) // 更新地图
	rsp.PropState = setState
	// 统一通知
	g.PropSceneGroupRefreshScNotify(propEntityIdList, blockBin) // 通知状态更改
	g.UpInteractSubMission(blockBin)                            // 检查交互任务
	// g.UpdateBlock(blockBin)                                     // 保存地图
	g.Send(cmd.InteractPropScRsp, rsp)
}

func (g *GamePlayer) SpringRecoverSingleAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SpringRecoverSingleAvatarCsReq, payloadMsg)
	req := msg.(*proto.SpringRecoverSingleAvatarCsReq)
	g.AvatarRecover(req.Id)

	rsp := &proto.SpringRecoverSingleAvatarScRsp{
		Hp:         10000,
		Retcode:    0,
		AvatarType: req.AvatarType,
		Id:         req.Id,
	}
	g.SyncLineupNotify(g.GetBattleLineUp())
	g.Send(cmd.SpringRecoverSingleAvatarScRsp, rsp)
}

// 更新实体状态
func (g *GamePlayer) PropSceneGroupRefreshScNotify(propEntityIdList []uint32, db *spb.BlockBin) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	for _, propId := range propEntityIdList {
		pe := g.GetPropEntityById(propId)
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
			AddEntity: &proto.SceneEntityInfo{
				EntityId: pe.EntityId,
				GroupId:  pe.GroupId,
				InstId:   pe.InstId,
				Motion: &proto.MotionInfo{
					Pos: pe.Pos,
					Rot: pe.Rot,
				},
				Prop: &proto.ScenePropInfo{
					PropId:    pe.PropId, // PropID
					PropState: g.GetPropState(db, pe.GroupId, pe.InstId, ""),
				},
			},
		})
	}

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

// 添加角色
func (g *GamePlayer) AddAvatarSceneGroupRefreshScNotify(avatarId uint32, isTrial bool, pos, rot *proto.Vector) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	info := new(proto.GroupRefreshInfo)
	actor := &proto.SceneActorInfo{
		AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
		BaseAvatarId: avatarId,
		MapLayer:     0,
		Uid:          0,
	}
	entityId := g.GetNextGameObjectGuid()
	if isTrial {
		conf := gdconf.GetSpecialAvatarById(avatarId)
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
		AddEntity: &proto.SceneEntityInfo{
			EntityId: entityId,
			Actor:    actor,
			Motion: &proto.MotionInfo{
				Rot: pos,
				Pos: rot,
			},
		},
	})
	g.AddEntity(0, &AvatarEntity{
		Entity: Entity{
			EntityId: entityId,
			GroupId:  0,
			Pos:      pos,
			Rot:      rot,
		},
		AvatarId: actor.BaseAvatarId,
	})
	notify.GroupRefreshList = append(notify.GroupRefreshList, info)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

// 卸载/加载场景
func (g *GamePlayer) UpSceneGroupRefreshScNotify(uninstallGroup, loadedGroupList []*GroupInfo) {
	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}

	// 卸载
	for _, groupInfo := range uninstallGroup {
		if groupInfo.GroupID == 0 { // 不能卸载角色
			continue
		}
		groupRefreshInfo := &proto.GroupRefreshInfo{
			GroupId:       groupInfo.GroupID,
			RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
		}

		for _, entify := range groupInfo.EntityMap {
			groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity, &proto.SceneEntityRefreshInfo{
				DeleteEntity: g.GetEntryId(entify),
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
		db := g.GetBlock(groupInfo.EntryId)
		groupRefreshInfo := &proto.GroupRefreshInfo{
			GroupId:       groupInfo.GroupID,
			RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
		}
		// 添加怪物
		groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity, g.AddMonsterSceneEntityRefreshInfo(groupInfo.GroupID, group.MonsterList)...)
		// 添加npc
		groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity, g.AddNpcSceneEntityRefreshInfo(groupInfo.GroupID, group.NPCList)...)
		// 添加实体
		groupRefreshInfo.RefreshEntity = append(groupRefreshInfo.RefreshEntity, g.AddPropSceneEntityRefreshInfo(groupInfo.GroupID, group.PropList, db)...)
		notify.GroupRefreshList = append(notify.GroupRefreshList, groupRefreshInfo)
		// g.UpdateBlock(db)
	}

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}
