package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) GetScene() *spb.Scene {
	if g.PlayerPb.Scene == nil {
		g.PlayerPb.Scene = &spb.Scene{
			EntryId: 1010101,
			PlaneId: 10101,
			FloorId: 10101001,
		}
	}
	return g.PlayerPb.Scene
}

func (g *GamePlayer) GetPos() *spb.VectorBin {
	if g.PlayerPb.Pos == nil {
		g.PlayerPb.Pos = &spb.VectorBin{
			X: -43300,
			Y: 6,
			Z: -37960,
		}
	}
	return g.PlayerPb.Pos
}

func (g *GamePlayer) GetRot() *spb.VectorBin {
	if g.PlayerPb.Rot == nil {
		g.PlayerPb.Rot = &spb.VectorBin{
			X: 0,
			Y: 90000,
			Z: 0,
		}
	}
	return g.PlayerPb.Rot
}

func (g *GamePlayer) GetPropByID(sceneGroup *gdconf.LevelGroup, groupID uint32) *proto.SceneEntityGroupInfo {
	entityGroupLists := &proto.SceneEntityGroupInfo{
		GroupId:    groupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for _, propList := range sceneGroup.PropList {
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,     // 文件名后那个G
			InstId:   propList.ID, // ID
			EntityId: uint32(g.GetNextGameObjectGuid()),
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(propList.PosX * 1000),
					Y: int32(propList.PosY * 1000),
					Z: int32(propList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(propList.RotY * 1000),
					Z: 0,
				},
			},
			Prop: &proto.ScenePropInfo{
				PropId:    propList.PropID, // PropID
				PropState: 1,               // gdconf.GetPropState(strconv.Itoa(int(propList.PropID))),
			},
		}
		if propList.State != "CheckPointDisable" && propList.State != "CheckPointEnable" {
			entityList.Prop.PropState = 8 // 解锁
		}
		entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
	}
	return entityGroupLists
}

func (g *GamePlayer) GetNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup, groupID uint32, entityMap map[uint32]*MonsterEntity) (*proto.SceneEntityGroupInfo, map[uint32]*MonsterEntity) {
	for _, monsterList := range sceneGroup.MonsterList {
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(monsterList.PosX * 1000),
					Y: int32(monsterList.PosY * 1000),
					Z: int32(monsterList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(monsterList.RotY * 1000),
					Z: 0,
				},
			},
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.PlayerPb.WorldLevel,
				MonsterId:  monsterList.NPCMonsterID,
				EventId:    monsterList.EventID,
			},
		}
		// 添加实体
		entityMap[entityId] = &MonsterEntity{
			MonsterEId: monsterList.EventID,
			GroupId:    groupID,
			Pos: &Vector{
				X: int32(monsterList.PosX * 1000),
				Y: int32(monsterList.PosY * 1000),
				Z: int32(monsterList.PosZ * 1000),
			},
			Rot: &Vector{
				X: 0,
				Y: int32(monsterList.RotY * 1000),
				Z: 0,
			},
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList, entityMap
}

func (g *GamePlayer) GetNPCByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup, groupID uint32) *proto.SceneEntityGroupInfo {
	for _, npcList := range sceneGroup.NPCList {
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,
			InstId:   npcList.ID,
			EntityId: uint32(g.GetNextGameObjectGuid()),
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(npcList.PosX * 1000),
					Y: int32(npcList.PosY * 1000),
					Z: int32(npcList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(npcList.RotY * 1000),
					Z: 0,
				},
			},
			Npc: &proto.SceneNpcInfo{
				ExtraInfo: nil,
				NpcId:     npcList.NPCID,
			},
		}
		if npcList.FirstDialogueGroupID != 0 {
			g.GetSceneNpcList()[npcList.NPCID] = npcList.FirstDialogueGroupID
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *GamePlayer) GetSceneInfo(entryId uint32, pos, rot *spb.VectorBin) *proto.SceneInfo {
	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetFloorById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return nil
	}

	var groupID = mapEntrance.StartGroupID
	groupID = foorMap.StartGroupID
	scene := &proto.SceneInfo{
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
		scene.LightenSectionList = append(scene.LightenSectionList, i)
	}

	// 获取场景实体
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	monsterEntity := make(map[uint32]*MonsterEntity, 0)
	avatarEntity := make(map[uint32]*AvatarEntity, 0)
	npcEntity := make(map[uint32]*NpcEntity, 0)
	// 添加队伍角色进实体列表，并设置坐标
	if foorMap.Groups[groupID] == nil {
		return nil
	}
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
	g.GetPos().X = pos.X
	g.GetPos().Y = pos.Y
	g.GetPos().Z = pos.Z
	g.GetRot().X = rot.X
	g.GetRot().Y = rot.Y
	g.GetRot().Z = rot.Z
	g.GetScene().EntryId = entryId
	g.GetScene().PlaneId = mapEntrance.PlaneID
	g.GetScene().FloorId = mapEntrance.FloorID
	g.GetLineUp().MainAvatarId = 0
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)

	for _, levelGroup := range foorMap.Groups {
		if levelGroup.GroupId == 0 {
			continue
		}
		if len(levelGroup.PropList) == 0 && len(levelGroup.NPCList) == 0 && len(levelGroup.MonsterList) == 0 {
			continue
		}
		scene.GroupIdList = append(scene.GroupIdList, levelGroup.GroupId)

		// 添加物品实体
		entityGroupList := g.GetPropByID(levelGroup, levelGroup.GroupId)
		// 添加怪物实体
		entityGroupList, x := g.GetNPCMonsterByID(entityGroupList, levelGroup, levelGroup.GroupId, monsterEntity)
		monsterEntity = x
		// 添加NPC实体
		entityGroupList = g.GetNPCByID(entityGroupList, levelGroup, levelGroup.GroupId)
		if len(entityGroupList.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)
		}
	}
	g.GetSceneEntity().MonsterEntity = monsterEntity
	g.GetSceneEntity().AvatarEntity = avatarEntity
	g.GetSceneEntity().NpcEntity = npcEntity

	return scene
}
