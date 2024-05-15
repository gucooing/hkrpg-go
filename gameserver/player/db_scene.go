package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type EntityAll interface {
	AddEntity(g *GamePlayer)
}

type Entity struct {
	EntityId uint32 // 实体id
	GroupId  uint32 // 区域
	Pos      *proto.Vector
	Rot      *proto.Vector
}

type AvatarEntity struct {
	Entity
	AvatarId uint32 // 角色id
}

type MonsterEntity struct {
	Entity
	EventID uint32 // 怪物id
}

type NpcEntity struct {
	Entity
	NpcId uint32 // ncp id
}

type PropEntity struct {
	Entity
	PropId uint32 // 物品id
}

func (g *GamePlayer) GetEntity() map[uint32]EntityAll {
	db := g.GetOnlineData()
	if db.EntityMap == nil {
		db.EntityMap = make(map[uint32]EntityAll)
	}
	return db.EntityMap
}

func (g *GamePlayer) NewEntity() map[uint32]EntityAll { // 清空实体列表用的
	db := g.GetOnlineData()
	db.EntityMap = make(map[uint32]EntityAll)
	return db.EntityMap
}

func (ae *AvatarEntity) AddEntity(g *GamePlayer) {
	db := g.GetEntity()
	db[ae.EntityId] = ae
}

func (me *MonsterEntity) AddEntity(g *GamePlayer) {
	db := g.GetEntity()
	db[me.EntityId] = me
}

func (ne *NpcEntity) AddEntity(g *GamePlayer) {
	db := g.GetEntity()
	db[ne.EntityId] = ne
}

func (pe *PropEntity) AddEntity(g *GamePlayer) {
	db := g.GetEntity()
	db[pe.EntityId] = pe
}

func (g *GamePlayer) AddEntity(t EntityAll) {
	t.AddEntity(g)
}

func (g *GamePlayer) GetEntityById(id uint32) EntityAll { // 根据实体id拉取实体
	db := g.GetEntity()
	return db[id]
}

func (g *GamePlayer) GetMonsterEntityById(id uint32) *MonsterEntity {
	db := g.GetEntityById(id)
	switch db.(type) {
	case *MonsterEntity:
		return db.(*MonsterEntity)
	}
	return nil
}

func (g *GamePlayer) NewScene() *spb.Scene {
	return &spb.Scene{
		EntryId: 1010101,
	}
}

func (g *GamePlayer) GetScene() *spb.Scene {
	db := g.BasicBin
	if db.Scene == nil {
		db.Scene = g.NewScene()
	}
	return db.Scene
}

func (g *GamePlayer) GetCurEntryId() uint32 {
	db := g.GetScene()
	return db.EntryId
}

func (g *GamePlayer) SetCurEntryId(id uint32) {
	db := g.GetScene()
	db.EntryId = id
}

func (g *GamePlayer) NewPos() *spb.VectorBin {
	return &spb.VectorBin{
		X: -43300,
		Y: 6,
		Z: -37960,
	}
}

func (g *GamePlayer) NewRot() *spb.VectorBin {
	return &spb.VectorBin{
		X: 0,
		Y: 90000,
		Z: 0,
	}
}

func (g *GamePlayer) GetPos() *spb.VectorBin {
	db := g.GetBasicBin()
	if db.Pos == nil {
		db.Pos = g.NewPos()
	}
	return db.Pos
}

func (g *GamePlayer) GetRot() *spb.VectorBin {
	db := g.GetBasicBin()
	if db.Rot == nil {
		db.Rot = g.NewRot()
	}
	return db.Rot
}

func (g *GamePlayer) SetPos(x, y, z int32) {
	db := g.GetPos()
	db.X = x
	db.Y = y
	db.Z = z
}

func (g *GamePlayer) SetRot(x, y, z int32) {
	db := g.GetRot()
	db.X = x
	db.Y = y
	db.Z = z
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetPosPb() *proto.Vector {
	db := g.GetPos()
	return &proto.Vector{
		Y: db.Y,
		X: db.X,
		Z: db.Z,
	}
}

func (g *GamePlayer) GetRotPb() *proto.Vector {
	db := g.GetRot()
	return &proto.Vector{
		Y: db.Y,
		X: db.X,
		Z: db.Z,
	}
}

func (g *GamePlayer) GetSceneAvatarByLineUP(entityGroupList *proto.SceneEntityGroupInfo, lineUp *spb.Line, leaderEntityId uint32, pos, rot *proto.Vector) {
	for sole, lineAvatar := range lineUp.AvatarIdList {
		if lineAvatar.AvatarId == 0 {
			continue
		}
		entityList := &proto.SceneEntityInfo{
			Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
				BaseAvatarId: lineAvatar.AvatarId,
			},
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
		}
		if sole == lineUp.LeaderSlot {
			entityList.EntityId = leaderEntityId
		} else {
			entityId := g.GetNextGameObjectGuid()
			entityList.EntityId = entityId
		}
		g.AddEntity(&AvatarEntity{
			Entity: Entity{
				EntityId: entityList.EntityId,
				GroupId:  0,
				Pos:      pos,
				Rot:      rot,
			},
			AvatarId: lineAvatar.AvatarId,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}

func (g *GamePlayer) GetPropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup) *proto.SceneEntityGroupInfo {
	for _, propList := range sceneGroup.PropList {
		propState := gdconf.GetStateValue(propList)
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(propList.PosX * 1000),
			Y: int32(propList.PosY * 1000),
			Z: int32(propList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: 0,
			Y: int32(propList.RotY * 1000),
			Z: 0,
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId, // 文件名后那个G
			InstId:   propList.ID,        // ID
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			Prop: &proto.ScenePropInfo{
				PropId:    propList.PropID, // PropID
				PropState: propState,
			},
		}
		// 添加物品实体
		g.AddEntity(&PropEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			PropId: propList.PropID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *GamePlayer) GetNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup) *proto.SceneEntityGroupInfo {
	for _, monsterList := range sceneGroup.MonsterList {
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
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.BasicBin.WorldLevel,
				MonsterId:  monsterList.NPCMonsterID,
				EventId:    monsterList.EventID,
			},
		}
		// 添加怪物实体
		g.AddEntity(&MonsterEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			EventID: monsterList.EventID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *GamePlayer) GetNPCByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.LevelGroup) *proto.SceneEntityGroupInfo {
	for _, npcList := range sceneGroup.NPCList {
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(npcList.PosX * 1000),
			Y: int32(npcList.PosY * 1000),
			Z: int32(npcList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(npcList.RotX * 1000),
			Y: int32(npcList.RotY * 1000),
			Z: int32(npcList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId,
			InstId:   npcList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			Npc: &proto.SceneNpcInfo{
				ExtraInfo: nil,
				NpcId:     npcList.NPCID,
			},
		}
		// 添加npc
		g.AddEntity(&NpcEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			NpcId: npcList.NPCID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *GamePlayer) GetSceneInfo(entryId uint32, pos, rot *proto.Vector, lineUp *spb.Line) *proto.SceneInfo {
	leaderEntityId := g.GetNextGameObjectGuid()
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(entryId)))
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetFloorById(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return nil
	}
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
	// 清理老实体列表
	g.NewEntity()
	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)
	for _, levelGroup := range foorMap.Groups {
		if levelGroup.GroupId == 0 {
			continue
		}
		if len(levelGroup.PropList) == 0 && len(levelGroup.NPCList) == 0 && len(levelGroup.MonsterList) == 0 {
			continue
		}
		scene.GroupIdList = append(scene.GroupIdList, levelGroup.GroupId)
		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加物品实体
		g.GetPropByID(entityGroupLists, levelGroup)
		// 添加怪物实体
		g.GetNPCMonsterByID(entityGroupLists, levelGroup)
		// 添加NPC实体
		g.GetNPCByID(entityGroupLists, levelGroup)
		if len(entityGroupLists.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
		}
	}
	return scene
}

func (g *GamePlayer) GetDelSceneGroupRefreshInfo(mem []uint32) []*proto.SceneGroupRefreshInfo {
	sceneGroupRefreshInfo := make([]*proto.SceneGroupRefreshInfo, 0)
	for _, id := range mem {
		entity := g.GetMonsterEntityById(id)
		if entity == nil {
			continue
		}
		sgri := &proto.SceneGroupRefreshInfo{
			State:   0,
			GroupId: entity.GroupId,
			RefreshEntity: []*proto.SceneEntityRefreshInfo{
				{
					DelEntity: entity.EventID,
				},
			},
			GroupRefreshType: 0,
		}
		sceneGroupRefreshInfo = append(sceneGroupRefreshInfo, sgri)
	}
	return sceneGroupRefreshInfo
}
