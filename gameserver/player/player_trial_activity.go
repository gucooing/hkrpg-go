package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type ActivityInfoOnline struct {
	StageId uint32 // 关卡id
}

func (g *GamePlayer) GetActivityInfoOnline() *ActivityInfoOnline {
	db := g.GetCurBattle()
	if db.ActivityInfoOnline == nil {
		db.ActivityInfoOnline = &ActivityInfoOnline{}
	}
	return db.ActivityInfoOnline
}

func (g *GamePlayer) StartTrialActivityCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartTrialActivityCsReq, payloadMsg)
	req := msg.(*proto.StartTrialActivityCsReq)

	avatarDemo := gdconf.GetAvatarDemoConfigById(req.StageId)
	if avatarDemo == nil {
		return
	}
	// 记录关卡
	db := g.GetActivityInfoOnline()
	db.StageId = req.StageId
	// 更新角色
	lineup := g.GetBattleLineUpById(Activity)
	lineup.LeaderSlot = 0
	lineup.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	for id, avatarId := range avatarDemo.TrialAvatarList {
		lineup.AvatarIdList[uint32(id-1)] = &spb.LineAvatarList{
			Slot:           uint32(id - 1),
			AvatarId:       avatarId,
			LineAvatarType: spb.LineAvatarType_LineAvatarType_TRIAL,
		}
	}

	g.StartTrialEnterSceneByServerScNotify()

	rsp := &proto.StartTrialActivityScRsp{StageId: req.StageId}
	g.Send(cmd.StartTrialActivityScRsp, rsp)
}

func (g *GamePlayer) StartTrialEnterSceneByServerScNotify() {
	notify := &proto.EnterSceneByServerScNotify{
		Scene:  g.GetTrialActivityScene(),
		Lineup: g.GetBattleLineUpPb(Activity),
	}
	g.Send(cmd.EnterSceneByServerScNotify, notify)
}

func (g *GamePlayer) GetTrialActivityScene() *proto.SceneInfo {
	db := g.GetActivityInfoOnline()
	avatarDemo := gdconf.GetAvatarDemoConfigById(db.StageId)
	if avatarDemo == nil {
		return nil
	}
	mapEntrance := gdconf.GetMapEntranceById(avatarDemo.MapEntranceID)
	if mapEntrance == nil {
		return nil
	}
	leaderEntityId := g.GetNextGameObjectGuid()
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            mapEntrance.PlaneID,
		FloorId:            mapEntrance.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID,
		EntryId:            avatarDemo.MapEntranceID,
		GameModeType:       14, // gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}
	// 获取场景实体
	entityGroupList := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	startGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, avatarDemo.MazeGroupID1)
	var pos *proto.Vector
	var rot *proto.Vector
	for _, anchor := range startGroup.AnchorList {
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
		break
	}
	lineUp := g.GetBattleLineUpById(Activity)

	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroupList, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)

	// 添加实体
	levelGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, avatarDemo.MazeGroupID1)
	scene.GroupIdList = append(scene.GroupIdList, levelGroup.GroupId)
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
			NpcMonster: &proto.SceneNpcMonsterInfo{
				MonsterId: avatarDemo.NpcMonsterIDList1[0],
				EventId:   avatarDemo.EventIDList1[0],
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
	// 添加NPC实体
	g.GetNPCByID(entityGroupLists, levelGroup)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)

	return scene
}

func (g *GamePlayer) TrialActivityPVEBattleResultScRsp(rsp *proto.PVEBattleResultScRsp) {
	// rsp.BattleAvatarList = g.TrialActivityGetBattleAvatarList()
	// if rsp.EndStatus == proto.BattleEndStatus_BATTLE_END_WIN {
	// 	// 传送回原来的场景
	// 	g.SceneByServerScNotify(g.GetScene().EntryId, g.GetPosPb(), g.GetRotPb())
	// 	// 储存通关状态
	// 	g.GetActivity().TrialActivity = append(g.GetActivity().TrialActivity, g.GetTrialActivityState().AvatarDemoId)
	// 	// 发送通关通知
	// 	scNotify := &proto.TrialActivityDataChangeScNotify{
	// 		TrialActivityInfo: &proto.TrialActivityInfo{
	// 			StageId:     g.GetTrialActivityState().AvatarDemoId,
	// 			TakenReward: false,
	// 		},
	// 	}
	// 	g.Send(cmd.TrialActivityDataChangeScNotify, scNotify)
	// 	notify := &proto.CurTrialActivityScNotify{
	// 		// TrialActivityId: g.GetTrialActivityState().AvatarDemoId,
	// 		Status: proto.TrialActivityStatus_TRIAL_ACTIVITY_STATUS_FINISH,
	// 	}
	// 	g.Send(cmd.CurTrialActivityScNotify, notify)
	// 	// 恢复战斗状态为空
	// 	g.GetBattleState().BattleType = spb.BattleType_Battle_NONE
	// }
	g.Send(cmd.PVEBattleResultScRsp, rsp)
}
