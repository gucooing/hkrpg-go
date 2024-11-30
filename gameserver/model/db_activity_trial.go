package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func (g *PlayerData) GetTrialActivity() *spb.TrialActivity {
	db := g.GetActivity()
	if db.TrialActivity == nil {
		db.TrialActivity = &spb.TrialActivity{}
	}
	return db.TrialActivity
}

func (g *PlayerData) GetTrialActivityInfo() map[uint32]*spb.TrialActivityInfo {
	db := g.GetTrialActivity()
	if db.TrialInfoMap == nil {
		db.TrialInfoMap = make(map[uint32]*spb.TrialActivityInfo)
	}
	return db.TrialInfoMap
}

func (g *PlayerData) GetTrialActivityInfoById(stageId uint32) *spb.TrialActivityInfo {
	db := g.GetTrialActivityInfo()
	if db[stageId] == nil {
		db[stageId] = &spb.TrialActivityInfo{
			StageId:     stageId,
			TakenReward: false,
		}
	}
	return db[stageId]
}

func (g *PlayerData) GetCurTrialActivityInfo() *spb.CurTrialActivity {
	db := g.GetTrialActivity()
	if db.CurTrial == nil {
		db.CurTrial = &spb.CurTrialActivity{}
	}
	return db.CurTrial
}

func (g *PlayerData) NewCurTrialActivityInfo(info *spb.CurTrialActivity) {
	db := g.GetTrialActivity()
	db.CurTrial = info
}

func (g *PlayerData) GetTrialActivityScene() *proto.SceneInfo {
	db := g.GetCurTrialActivityInfo()
	avatarDemo := gdconf.GetAvatarDemoConfigById(db.StageId)
	if avatarDemo == nil {
		return nil
	}
	mapEntrance := gdconf.GetMapEntranceById(avatarDemo.MapEntranceID)
	if mapEntrance == nil {
		return nil
	}
	pos, rot := gdconf.GetAnchorByIndexPosRot(mapEntrance.PlaneID, mapEntrance.FloorID)
	scene := g.GetBattleLoadScene(avatarDemo.MapEntranceID, pos, rot,
		g.GetBattleLineUpById(Activity))
	// 添加实体
	for _, levelGroup := range gdconf.GetServerGroup(mapEntrance.PlaneID, mapEntrance.FloorID) {
		// scene.GroupIdList = append(scene.GroupIdList, levelGroup.GroupId)
		sceneGroupState := &proto.SceneGroupState{
			GroupId:   levelGroup.GroupId,
			IsDefault: true,
		}
		scene.GroupStateList = append(scene.GroupStateList, sceneGroupState)

		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
		}
		// 添加物品实体
		g.GetPropByID(entityGroupLists, levelGroup, nil, avatarDemo.MapEntranceID)
		// 添加怪物实体
		if levelGroup.GroupId == avatarDemo.MazeGroupID1 {
			for _, monsterList := range levelGroup.MonsterList {
				entityId := g.GetNextGameObjectGuid()
				monsterPos := &proto.Vector{
					X: int32(monsterList.PosX * 1000),
					Y: int32(monsterList.PosY * 1000),
					Z: int32(monsterList.PosZ * 1000),
				}
				monsterRot := &proto.Vector{
					X: int32(monsterList.RotX * 1000),
					Y: int32(monsterList.RotY * 1000),
					Z: int32(monsterList.RotZ * 1000),
				}
				entityList := &proto.SceneEntityInfo{
					GroupId:  levelGroup.GroupId,
					InstId:   monsterList.ID,
					EntityId: entityId,
					Motion: &proto.MotionInfo{
						Pos: monsterPos,
						Rot: monsterRot,
					},
					EntityOneofCase: &proto.SceneEntityInfo_NpcMonster{
						NpcMonster: &proto.SceneNpcMonsterInfo{
							MonsterId: avatarDemo.NpcMonsterIDList1[0],
							EventId:   avatarDemo.EventIDList1[0],
						},
					},
				}
				// 添加怪物实体
				g.AddEntity(levelGroup.GroupId, &MonsterEntity{
					Entity: Entity{
						InstId:   monsterList.ID,
						EntityId: entityId,
						GroupId:  levelGroup.GroupId,
						Pos:      monsterPos,
						Rot:      monsterRot,
					},
					EventID: avatarDemo.EventIDList1[0],
				})
				entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
			}
		} else {
			g.GetNPCMonsterByID(entityGroupLists, levelGroup)
		}

		// 添加NPC实体
		g.GetSceneNPCByConf(entityGroupLists, levelGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}

	return scene
}
