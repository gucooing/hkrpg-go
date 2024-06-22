package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) GetRogueTournSeasonInfo() *proto.RogueTournSeasonInfo {
	info := &proto.RogueTournSeasonInfo{
		SubTournId:  1,
		MainTournId: 1,
	}
	return info
}

func (g *GamePlayer) GetInspirationCircuitInfo() *proto.InspirationCircuitInfo {
	info := &proto.InspirationCircuitInfo{
		RogueTalentInfo: &proto.RogueTalentInfo{RogueTalentList: make([]*proto.RogueTalent, 0)},
		InspirationNum:  g.GetMaterialById(Inspiration),
	}
	for v, k := range gdconf.GetRogueTournPermanentTalentMap() {
		status := proto.RogueTalentStatus_ROGUE_TALENT_STATUS_LOCK
		if k.IsImportant {
			status = proto.RogueTalentStatus_ROGUE_TALENT_STATUS_UNLOCK
		}
		status = proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE
		info.RogueTalentInfo.RogueTalentList = append(info.RogueTalentInfo.RogueTalentList, &proto.RogueTalent{
			Status:   status,
			TalentId: v,
		})
	}
	return info
}

func (g *GamePlayer) GetExtraScoreInfo() *proto.ExtraScoreInfo {
	info := &proto.ExtraScoreInfo{
		LFADDDLCGBM: 1,
		IsFinish:    true,
		EndTime:     1719172800,
		JLDGFGMEMJH: 1000,
	}
	return info
}

func (g *GamePlayer) GetSynchronicityLevelInfo() *proto.SynchronicityLevelInfo {
	info := &proto.SynchronicityLevelInfo{
		ReceivedLevelRewards: make([]uint32, 0),
		Exp:                  0,
	}

	return info
}

func (g *GamePlayer) GetRogueTournDifficultyInfo() []*proto.RogueTournDifficultyInfo {
	info := make([]*proto.RogueTournDifficultyInfo, 0)
	for _, id := range []uint32{10101, 10102, 10103, 10104, 10105, 10106} {
		info = append(info, &proto.RogueTournDifficultyInfo{
			DifficultyId: id,
			IsUnlock:     true,
		})
	}

	return info
}

func (g *GamePlayer) GetRogueTournAreaInfo() []*proto.RogueTournAreaInfo {
	info := make([]*proto.RogueTournAreaInfo, 0)
	for _, id := range []uint32{101, 201, 202, 203, 204, 205, 1011, 1012, 1013, 1014, 1015} {
		info = append(info, &proto.RogueTournAreaInfo{
			AreaId:                      id,
			IsFinish:                    true,
			IsTakenReward:               false,
			IsUnlock:                    true,
			UnlockedTournDifficultyList: make([]uint32, 0),
		})
	}
	return info
}

func (g *GamePlayer) GetRogueTournCurInfo() *proto.RogueTournCurInfo {
	info := &proto.RogueTournCurInfo{
		IHELIGMBGIL: nil,
		LDKKBIIEKGK: nil,
	}
	return info
}

func (g *GamePlayer) GetRogueTournScene(entryId uint32) *proto.SceneInfo {
	mapEntrance := gdconf.GetMapEntranceById(entryId)
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
		EntryId:            entryId,
		GameModeType:       17, // gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
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
	startGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, 64)
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
	lineUp := g.GetBattleLineUpById(Rogue)

	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroupList, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)

	// for groupID, ida := range rogueRoom.GroupWithContent {
	// 	sceneGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, groupID)
	// 	if sceneGroup == nil {
	// 		continue
	// 	}
	// 	scene.GroupIdList = append(scene.GroupIdList, groupID)
	// 	sceneGroupState := &proto.SceneGroupState{
	// 		GroupId:   groupID,
	// 		IsDefault: true,
	// 	}
	// 	scene.GroupStateList = append(scene.GroupStateList, sceneGroupState)
	//
	// 	entityGroupLists := &proto.SceneEntityGroupInfo{
	// 		GroupId:    groupID,
	// 		EntityList: make([]*proto.SceneEntityInfo, 0),
	// 	}
	// 	// 添加物品实体
	// 	g.GetRoguePropByID(entityGroupLists, sceneGroup)
	// 	// 添加怪物实体
	// 	g.GetRogueNPCMonsterByID(entityGroupLists, sceneGroup, ida)
	// 	// 添加NPC实体
	// 	g.GetNPCByID(entityGroupLists, sceneGroup)
	// 	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	// }

	return scene
}
