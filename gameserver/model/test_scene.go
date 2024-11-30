package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

type battleLoadScene struct {
	entryId       uint32
	pos, rot      *proto.Vector
	lineUp        *spb.Line
	loadGroupList map[uint32]*battleGroupInfo
}

type battleGroupInfo struct {
	groupId         uint32 // 组id
	monsterInfoList map[uint32]*monsterInfo
}

type monsterInfo struct {
	monsterId uint32
	eventId   uint32
}

func (g *PlayerData) testGetBattleScene(b *battleLoadScene) *proto.SceneInfo {
	if b == nil {
		return nil
	}
	planeID, floorID, ok := gdconf.GetPFlaneID(b.entryId)
	if !ok {
		// TODO log
		return nil
	}
	leaderEntityId := g.GetNextGameObjectGuid()
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		WorldId:            gdconf.GetWorldId(planeID),
		LeaderEntityId:     leaderEntityId,
		FloorId:            floorID,
		GameModeType:       gdconf.GetPlaneType(planeID),
		PlaneId:            planeID,
		EntryId:            b.entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		FloorSavedData:     g.GetFloorSavedData(b.entryId),
		EntityBuffInfoList: make([]*proto.EntityBuffInfo, 0),
	}
	// 获取场景实体
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroup, b.lineUp, leaderEntityId, b.pos, b.rot)
	// 添加场景
	for groupID, groupInfo := range b.loadGroupList {
		group := gdconf.GetServerGroupById(planeID, floorID, groupInfo.groupId)
		if group == nil {
			continue
		}
		sceneGroupState := &proto.SceneGroupState{
			GroupId:   groupID,
			IsDefault: true,
		}
		scene.GroupStateList = append(scene.GroupStateList, sceneGroupState)

		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    groupID,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加NPC实体
		g.GetSceneNPCByConf(entityGroupLists, group)
		// 添加物品实体
		g.testGetBattlePropByID(entityGroupLists, group, nil, b.entryId)
		// 添加怪物实体
		g.testGetBattleNPCMonsterByID(entityGroupLists, group, groupInfo.monsterInfoList)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}

	return scene
}

func (g *PlayerData) testGetBattlePropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, db *spb.BlockBin, entryId uint32) *proto.SceneEntityGroupInfo {
	for _, propList := range sceneGroup.PropList {
		entityId := g.GetNextGameObjectGuid()
		g.StageObjectCapture(sceneGroup, propList, sceneGroup.GroupId, db)
		propState := g.GetPropState(db, sceneGroup.GroupId, propList.ID, propList.State)
		pos := &proto.Vector{
			X: int32(propList.PosX * 1000),
			Y: int32(propList.PosY * 1000),
			Z: int32(propList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(propList.RotX * 1000),
			Y: int32(propList.RotY * 1000),
			Z: int32(propList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId, // 文件名后那个G
			InstId:   propList.ID,        // ID
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			EntityOneofCase: &proto.SceneEntityInfo_Prop{},
		}
		prop := &proto.ScenePropInfo{
			PropId:    propList.PropID, // PropID
			PropState: propState,
		}
		switch g.GetBattleStatus() {
		case Rogue:
		case RogueTourn:

		}

		entityList.EntityOneofCase = &proto.SceneEntityInfo_Prop{
			Prop: prop,
		}
		// 添加物品实体
		g.AddEntity(sceneGroup.GroupId, &PropEntity{
			Entity: Entity{
				EntityId: entityId,
				InstId:   propList.ID,
				EntryId:  entryId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			PropId:              prop.PropId,
			TriggerBattleString: propList.TriggerBattleString,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *PlayerData) testGetBattleNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, monsterInfoList map[uint32]*monsterInfo) {
	for _, monsterList := range sceneGroup.MonsterList {
		info, ok := monsterInfoList[monsterList.ID]
		if !ok {
			continue
		}
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(monsterList.PosX * 1000),
			Y: int32(monsterList.PosY * 1000),
			Z: int32(monsterList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(monsterList.RotX * 1000),
			Y: int32(monsterList.RotY * 1000),
			Z: int32(monsterList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			EntityOneofCase: &proto.SceneEntityInfo_NpcMonster{
				NpcMonster: &proto.SceneNpcMonsterInfo{
					WorldLevel: g.GetWorldLevel(),
					MonsterId:  info.monsterId,
					EventId:    info.eventId,
				},
			},
		}
		// 添加实体
		g.AddEntity(sceneGroup.GroupId, &MonsterEntity{
			Entity: Entity{
				InstId:   monsterList.ID,
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			EventID: info.eventId,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}
