package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

// 通知客户端进入场景
func (g *GamePlayer) EnterSceneByServerScNotify(entryId, teleportId uint32) {
	rsp := new(proto.EnterSceneByServerScNotify)
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))
	if mapEntrance == nil {
		return
	}
	foorMap := gdconf.GetFloorById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return
	}

	var groupID = mapEntrance.StartGroupID
	var anchorID = mapEntrance.StartAnchorID

	if teleportId != 0 {
		groupID = foorMap.Teleports[teleportId].AnchorGroupID
		anchorID = foorMap.Teleports[teleportId].AnchorID
	} else if anchorID == 0 {
		groupID = foorMap.StartGroupID
		anchorID = foorMap.StartAnchorID
	}

	// 获取队伍
	rsp.Lineup = g.GetLineUpPb(g.GetLineUp().MainLineUp)

	rsp.Scene = &proto.SceneInfo{
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).WorldID,
		LeaderEntityId:     leaderEntityId,
		FloorId:            mapEntrance.FloorID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).PlaneType),
		PlaneId:            mapEntrance.PlaneID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
	}

	for i := uint32(0); i < 100; i++ {
		rsp.Scene.LightenSectionList = append(rsp.Scene.LightenSectionList, i)
	}

	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}

	// 获取场景实体
	monsterEntity := make(map[uint32]*MonsterEntity, 0)
	avatarEntity := make(map[uint32]*AvatarEntity, 0)
	npcEntity := make(map[uint32]*NpcEntity, 0)

	// 添加队伍角色进实体列表，并设置坐标
	if foorMap.Groups[groupID] == nil {
		return
	}
	for _, anchor := range foorMap.Groups[groupID].AnchorList {
		if anchor.ID == anchorID {
			for id, avatarid := range g.GetLineUpById(g.GetLineUp().MainLineUp).AvatarIdList {
				if avatarid == 0 {
					continue
				}
				entityId := uint32(g.GetNextGameObjectGuid())
				entityList := &proto.SceneEntityInfo{
					Actor: &proto.SceneActorInfo{
						AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
						BaseAvatarId: avatarid,
					},
					Motion: &proto.MotionInfo{
						Pos: &proto.Vector{
							X: int32(anchor.PosX * 1000),
							Y: int32(anchor.PosY * 1000),
							Z: int32(anchor.PosZ * 1000),
						},
						Rot: &proto.Vector{
							X: int32(anchor.RotX * 1000),
							Y: int32(anchor.RotY * 1000),
							Z: int32(anchor.RotZ * 1000),
						},
					},
				}
				// 为进入场景的角色设置与上面相同的实体id
				if id == 0 {
					entityList.EntityId = leaderEntityId
					avatarEntity[leaderEntityId] = &AvatarEntity{
						AvatarId: avatarid,
						GroupId:  groupID,
					}
				} else {
					entityList.EntityId = entityId
					avatarEntity[entityId] = &AvatarEntity{
						AvatarId: avatarid,
						GroupId:  groupID,
					}
				}
				entityGroup.EntityList = append(entityGroup.EntityList, entityList)
			}

			// TODO 数据保存问题？
			g.GetPos().X = int32(anchor.PosX * 1000)
			g.GetPos().Y = int32(anchor.PosY * 1000)
			g.GetPos().Z = int32(anchor.PosZ * 1000)
			g.GetRot().X = int32(anchor.RotX * 1000)
			g.GetRot().Y = int32(anchor.RotY * 1000)
			g.GetRot().Z = int32(anchor.RotZ * 1000)
			g.GetScene().EntryId = entryId
			g.GetScene().PlaneId = mapEntrance.PlaneID
			g.GetScene().FloorId = mapEntrance.FloorID
			g.GetLineUp().MainAvatarId = 0
			break
		}
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	for _, levelGroup := range foorMap.Groups {
		if levelGroup.GroupId == 0 {
			continue
		}
		if len(levelGroup.PropList) == 0 && len(levelGroup.NPCList) == 0 && len(levelGroup.MonsterList) == 0 {
			continue
		}
		rsp.Scene.GroupIdList = append(rsp.Scene.GroupIdList, levelGroup.GroupId)

		// 添加物品实体
		entityGroupList := g.GetPropByID(levelGroup, levelGroup.GroupId)
		// 添加怪物实体
		entityGroupList, x := g.GetNPCMonsterByID(entityGroupList, levelGroup, levelGroup.GroupId, monsterEntity)
		monsterEntity = x
		// 添加NPC实体
		entityGroupList = g.GetNPCByID(entityGroupList, levelGroup, levelGroup.GroupId)
		if len(entityGroupList.EntityList) != 0 {
			rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroupList)
		}
	}

	g.GetSceneEntity().MonsterEntity = monsterEntity
	g.GetSceneEntity().AvatarEntity = avatarEntity
	g.GetSceneEntity().NpcEntity = npcEntity
	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

// 传送到指定位置
func (g *GamePlayer) SceneByServerScNotify(entryId uint32, pos, rot *spb.VectorBin) {
	rsp := new(proto.EnterSceneByServerScNotify)
	// 获取队伍
	rsp.Lineup = g.GetLineUpPb(g.GetLineUp().MainLineUp)
	rsp.Scene = g.GetSceneInfo(entryId, pos, rot)

	g.Send(cmd.EnterSceneByServerScNotify, rsp)
}

func (g *GamePlayer) HandleGetEnteredSceneCsReq(payloadMsg []byte) {
	rsp := new(proto.GetEnteredSceneScRsp)
	scene := g.GetScene()
	enteredSceneInfo := &proto.EnteredSceneInfo{
		FloorId: scene.FloorId,
		PlaneId: scene.PlaneId,
	}
	rsp.EnteredSceneInfo = []*proto.EnteredSceneInfo{enteredSceneInfo}

	g.Send(cmd.GetEnteredSceneScRsp, rsp)
}

// 客户端登录需要的包，不是传送的通知包
func (g *GamePlayer) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurSceneInfoScRsp)
	pos := g.GetPos()
	rot := g.GetRot()
	dbScene := g.GetScene()
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(dbScene.EntryId)))

	rsp.Scene = &proto.SceneInfo{
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).WorldID,                        // 世界id
		LeaderEntityId:     leaderEntityId,                                                                                 // 进入场景的角色实体id
		FloorId:            mapEntrance.FloorID,                                                                            // 上面表查询到的，对应文件名中F开头后面的数字
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(mapEntrance.PlaneID))).PlaneType), // 未知
		PlaneId:            mapEntrance.PlaneID,                                                                            // 上面表查询到的，对应文件名中P开头后面的数字
		EntryId:            dbScene.EntryId,                                                                                // 应该是具体的场景条目
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),                                                         // 实体列表
		GroupIdList:        make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
	}

	for i := uint32(0); i < 100; i++ {
		rsp.Scene.LightenSectionList = append(rsp.Scene.LightenSectionList, i)
	}

	// 获取场景实体
	monsterEntity := make(map[uint32]*MonsterEntity, 0)
	avatarEntity := make(map[uint32]*AvatarEntity, 0)
	npcEntity := make(map[uint32]*NpcEntity, 0)
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 将进入场景的角色添加到实体列表里
	for id, avatarid := range g.GetLineUpById(g.PlayerPb.LineUp.MainLineUp).AvatarIdList {
		if avatarid == 0 {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType(g.PlayerPb.Avatar.Avatar[avatarid].AvatarType),
				BaseAvatarId: avatarid,
			},
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: pos.X,
					Y: pos.Y,
					Z: pos.Z,
				},
				Rot: &proto.Vector{
					X: rot.X,
					Y: rot.Y,
					Z: rot.Z,
				},
			},
		}
		// 为进入场景的角色设置与上面相同的实体id
		if id == 0 {
			entityList.EntityId = leaderEntityId
			avatarEntity[leaderEntityId] = &AvatarEntity{
				AvatarId: avatarid,
				GroupId:  0,
			}
		} else {
			avatarEntity[entityId] = &AvatarEntity{
				AvatarId: avatarid,
				GroupId:  0,
			}
			entityList.EntityId = entityId
		}
		entityGroup.EntityList = append(entityGroup.EntityList, entityList)
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	foorMap := gdconf.GetFloorById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return
	}
	for _, levelGroup := range foorMap.Groups {
		if levelGroup.GroupId == 0 {
			continue
		}
		if len(levelGroup.PropList) == 0 && len(levelGroup.NPCList) == 0 && len(levelGroup.MonsterList) == 0 {
			continue
		}
		rsp.Scene.GroupIdList = append(rsp.Scene.GroupIdList, levelGroup.GroupId)

		// 添加物品实体
		entityGroupList := g.GetPropByID(levelGroup, levelGroup.GroupId)
		// 添加怪物实体
		entityGroupList, x := g.GetNPCMonsterByID(entityGroupList, levelGroup, levelGroup.GroupId, monsterEntity)
		monsterEntity = x
		// 添加NPC实体
		entityGroupList = g.GetNPCByID(entityGroupList, levelGroup, levelGroup.GroupId)
		if len(entityGroupList.EntityList) != 0 {
			rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroupList)
		}

	}
	g.GetSceneEntity().MonsterEntity = monsterEntity
	g.GetSceneEntity().AvatarEntity = avatarEntity
	g.GetSceneEntity().NpcEntity = npcEntity
	g.GetLineUp().MainAvatarId = 0

	g.Send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *GamePlayer) HanldeGetSceneMapInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetSceneMapInfoCsReq, payloadMsg)
	req := msg.(*proto.GetSceneMapInfoCsReq)

	// 1000101 1000001

	rsp := new(proto.GetSceneMapInfoScRsp)
	for _, entryId := range req.EntryIdList {
		mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))
		if mapEntrance != nil {
			groupMap := gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)
			if groupMap != nil {
				if int(mapEntrance.StartGroupID) > len(groupMap) {
					return
				}

				mapList := &proto.MazeMapData{
					LightenSectionList: make([]uint32, 0),
					UnlockedChestList: []*proto.MazeChest{
						{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_NORMAL},
						{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_PUZZLE},
						{MapInfoChestType: proto.MapInfoChestType_MAP_INFO_CHEST_TYPE_CHALLENGE},
					},
					UnlockedTeleportList: make([]uint32, 0),
				}

				mapList.EntryId = entryId

				for i := uint32(0); i < 100; i++ {
					mapList.LightenSectionList = append(mapList.LightenSectionList, i)
				}

				for _, groupInfo := range groupMap {
					mazeGroup := &proto.MazeGroup{GroupId: groupInfo.GroupId}
					mapList.MazeGroupList = append(mapList.MazeGroupList, mazeGroup)
				}

				for _, groupMapList := range groupMap {
					for _, propList := range groupMapList.PropList {
						if propList.State != "CheckPointDisable" && propList.State != "CheckPointEnable" {
							continue
						}
						mazeProp := &proto.MazeProp{
							State:    gdconf.GetStateValue(propList.State),
							GroupId:  groupMapList.GroupId,
							ConfigId: propList.ID,
						}
						mapList.MazePropList = append(mapList.MazePropList, mazeProp)
						mapList.UnlockedTeleportList = append(mapList.UnlockedTeleportList, propList.MappingInfoID)
					}
				}
				rsp.MapList = append(rsp.MapList, mapList)
			}
		}
	}

	g.Send(cmd.GetSceneMapInfoScRsp, rsp)
}

func (g *GamePlayer) EnterSceneCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.EnterSceneCsReq, payloadMsg)
	req := msg.(*proto.EnterSceneCsReq)
	rsp := &proto.GetEnteredSceneScRsp{}

	g.EnterSceneByServerScNotify(req.EntryId, req.TeleportId)

	g.Send(cmd.EnterSceneScRsp, rsp)
	g.Send(cmd.SceneUpdatePositionVersionNotify, rsp)
}

func (g *GamePlayer) InteractPropCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.InteractPropCsReq, payloadMsg)
	req := msg.(*proto.InteractPropCsReq)

	rsp := new(proto.InteractPropScRsp)
	rsp.PropEntityId = req.PropEntityId

	g.Send(cmd.InteractPropScRsp, rsp)
}
