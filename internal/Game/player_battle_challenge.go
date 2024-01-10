package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/***********************************忘却之庭***********************************/

// 获取状态

func (g *Game) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurChallengeScRsp)

	challengeState := g.GetChallengeState()

	rsp.ChallengeInfo = &proto.ChallengeInfo{
		ChallengeId:     challengeState.ChallengeId,
		Status:          challengeState.Status,
		RoundCount:      challengeState.RoundCount,
		ExtraLineupType: challengeState.ExtraLineupType,
		Score:           challengeState.ChallengeScore,
		StoryInfo:       &proto.ChallengeStoryInfo{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: make([]uint32, 0)}},
	}
	if challengeState.ChallengeCount == 1 {
		rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffOne)
	} else {
		rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffTwo)
	}

	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

// 进入忘却之庭

func (g *Game) StartChallengeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartChallengeCsReq, payloadMsg)
	req := msg.(*proto.StartChallengeCsReq)
	battleState := g.GetBattleState()
	challengeState := g.GetChallengeState()

	if req.StoryInfo.StoryBuffInfo.StoryBuffOne != 0 {
		battleState.BattleType = spb.BattleType_Battle_CHALLENGE_Story
		challengeState.StoryBuffOne = req.StoryInfo.StoryBuffInfo.StoryBuffOne
		challengeState.StoryBuffTwo = req.StoryInfo.StoryBuffInfo.StoryBuffTwo
	} else {
		battleState.BattleType = spb.BattleType_Battle_CHALLENGE
	}

	// 如果是新战斗就添加
	challengeState.ChallengeId = req.ChallengeId
	challengeState.Status = proto.ChallengeStatus_CHALLENGE_DOING
	challengeState.RoundCount = 0
	challengeState.ExtraLineupType = proto.ExtraLineupType_LINEUP_CHALLENGE

	challengeInfo := &proto.ChallengeInfo{
		ChallengeId:     challengeState.ChallengeId,
		Status:          challengeState.Status,
		RoundCount:      challengeState.RoundCount,
		ExtraLineupType: challengeState.ExtraLineupType,
		Score:           challengeState.ChallengeScore,
		StoryInfo:       &proto.ChallengeStoryInfo{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: make([]uint32, 0)}},
	}
	challengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(challengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffOne)

	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(strconv.Itoa(int(req.ChallengeId)))
	if challengeInfo == nil || challengeMazeConfig == nil {
		rsp := &proto.StartChallengeScRsp{
			Retcode: 2,
		}
		g.Send(cmd.StartChallengeScRsp, rsp)
		return
	}
	// 获取映射信息
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(challengeMazeConfig.MapEntranceID)))
	sceneGroup := gdconf.GetNGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, challengeMazeConfig.MazeGroupID1)
	if sceneGroup.AnchorList == nil {
		rsp := &proto.StartChallengeScRsp{
			Retcode: 2,
		}
		g.Send(cmd.StartChallengeScRsp, rsp)
		return
	}
	// 获取坐标信息
	challengeState.Pos = &spb.VectorBin{
		X: int32(sceneGroup.AnchorList[0].PosX * 1000),
		Y: int32(sceneGroup.AnchorList[0].PosY * 1000),
		Z: int32(sceneGroup.AnchorList[0].PosZ * 1000),
	}
	challengeState.Rot = &spb.VectorBin{
		X: int32(sceneGroup.AnchorList[0].RotX * 1000),
		Y: int32(sceneGroup.AnchorList[0].RotY * 1000),
		Z: int32(sceneGroup.AnchorList[0].RotZ * 1000),
	}
	challengeState.NPCMonsterPos = &spb.VectorBin{
		X: int32(sceneGroup.MonsterList[0].PosX * 1000),
		Y: int32(sceneGroup.MonsterList[0].PosY * 1000),
		Z: int32(sceneGroup.MonsterList[0].PosZ * 1000),
	}
	challengeState.NPCMonsterRot = &spb.VectorBin{
		X: int32(sceneGroup.MonsterList[0].RotX * 1000),
		Y: int32(sceneGroup.MonsterList[0].RotY * 1000),
		Z: int32(sceneGroup.MonsterList[0].RotZ * 1000),
	}
	challengeState.PlaneID = mapEntrance.PlaneID
	challengeState.FloorID = mapEntrance.FloorID
	challengeState.EntranceID = challengeMazeConfig.MapEntranceID

	challengeState.ChallengeCount = challengeMazeConfig.StageNum
	challengeState.CurChallengeCount = 1
	challengeState.ChallengeTargetID = challengeMazeConfig.ChallengeTargetID
	challengeState.ChallengeCountDown = challengeMazeConfig.ChallengeCountDown
	if challengeState.ChallengeCountDown == 0 {
		challengeState.ChallengeCountDown = 5
	}
	battleState.BuffList = append(battleState.BuffList, challengeMazeConfig.MazeBuffID)
	// 添加波次
	challengeState.CurChallengeBattle = make(map[uint32]*CurChallengeBattle)
	for id, challengeRoom := range challengeMazeConfig.ChallengeState {
		curChallengeBattle := &CurChallengeBattle{
			NPCMonsterID: challengeRoom.NPCMonsterID,
			EventID:      challengeRoom.EventID,
			GroupID:      challengeRoom.GroupID,
			ConfigID:     challengeRoom.ConfigID,
		}
		challengeState.CurChallengeBattle[id] = curChallengeBattle
	}

	scene := g.GetChallengeScene()

	lineup := g.GetLineUpPb(6)

	rsp := &proto.StartChallengeScRsp{
		ChallengeInfo: challengeInfo,
		Scene:         scene,
		Lineup:        lineup,
	}

	battleState.BattleType = spb.BattleType_Battle_CHALLENGE
	g.Send(cmd.StartChallengeScRsp, rsp)
}

func (g *Game) GetChallengeScene() *proto.SceneInfo {
	challengeState := g.GetChallengeState()

	entityMap := make(map[uint32]*EntityList) // [实体id]怪物群id

	leaderEntityId := uint32(g.GetNextGameObjectGuid())
	// 获取映射信息

	anchorPos := challengeState.Pos
	anchorRot := challengeState.Rot
	curChallengeBattle := challengeState.CurChallengeBattle[1]
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            challengeState.PlaneID,
		FloorId:            challengeState.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(strconv.Itoa(int(challengeState.PlaneID))).WorldID,
		EntryId:            challengeState.EntranceID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(strconv.Itoa(int(challengeState.PlaneID))).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}

	// 将进入场景的角色添加到实体列表里
	entityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    0,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for id, slots := range g.GetLineUpPb(6).AvatarList {
		if slots == nil {
			continue
		}
		entityId := uint32(g.GetNextGameObjectGuid())
		entityList := &proto.SceneEntityInfo{
			Actor: &proto.SceneActorInfo{
				AvatarType:   slots.AvatarType, // TODO
				BaseAvatarId: slots.Id,
			},
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: anchorPos.X,
					Y: anchorPos.Y,
					Z: anchorPos.Z,
				},
				Rot: &proto.Vector{
					X: anchorRot.X,
					Y: anchorRot.Y,
					Z: anchorRot.Z,
				},
			},
		}
		// 为进入场景的角色设置与上面相同的实体id
		if id == 0 {
			entityList.EntityId = leaderEntityId
			entityMap[leaderEntityId] = &EntityList{
				Entity:  slots.Slot,
				GroupId: 0,
			}
		} else {
			entityMap[entityId] = &EntityList{
				Entity:  slots.Slot,
				GroupId: 0,
			}
			entityList.EntityId = entityId
		}
		entityGroup.EntityList = append(entityGroup.EntityList, entityList)
	}
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)

	// 添加怪物实体
	entityGroupNPCMonster := &proto.SceneEntityGroupInfo{
		GroupId:    curChallengeBattle.GroupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	entityId := uint32(g.GetNextGameObjectGuid())
	entityList := &proto.SceneEntityInfo{
		GroupId:  curChallengeBattle.GroupID,
		InstId:   curChallengeBattle.ConfigID,
		EntityId: entityId,
		Motion: &proto.MotionInfo{
			Pos: &proto.Vector{
				X: challengeState.NPCMonsterPos.X,
				Y: challengeState.NPCMonsterPos.Y,
				Z: challengeState.NPCMonsterPos.Z,
			},
			Rot: &proto.Vector{
				X: 0,
				Y: challengeState.NPCMonsterRot.Y,
				Z: 0,
			},
		},
		NpcMonster: &proto.SceneNpcMonsterInfo{
			WorldLevel: g.PlayerPb.WorldLevel,
			MonsterId:  curChallengeBattle.NPCMonsterID,
			EventId:    curChallengeBattle.EventID,
		},
	}
	entityMap[entityId] = &EntityList{
		Entity:  curChallengeBattle.EventID,
		GroupId: curChallengeBattle.GroupID,
	}
	entityGroupNPCMonster.EntityList = append(entityGroupNPCMonster.EntityList, entityList)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupNPCMonster)

	g.Player.EntityList = entityMap
	return scene
}

// 忘却之庭战斗退出/结束

func (g *Game) LeaveChallengeCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	if g.GetBattleState().ChallengeState.Status == proto.ChallengeStatus_CHALLENGE_DOING {
		g.Send(cmd.QuitBattleScNotify, rsp)
	}
	g.Send(cmd.LeaveChallengeScRsp, rsp)

	g.EnterSceneByServerScNotify(g.GetScene().EntryId, 0)
	g.GetBattleState().BattleType = spb.BattleType_Battle_NONE
	g.GetBattleState().BuffList = make([]uint32, 0)
}

// 忘却之庭世界发生攻击事件

func (g *Game) ChallengeSceneCastSkillCsReq(payloadMsg []byte) {
}

// 忘却之庭世界战斗结算事件

func (g *Game) ChallengePVEBattleResultCsReq(payloadMsg []byte) {
}
