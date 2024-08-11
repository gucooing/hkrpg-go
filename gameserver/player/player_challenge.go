package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

/***********************************忘却之庭***********************************/

func (g *GamePlayer) HandleGetChallengeCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetChallengeScRsp)
	rsp.ChallengeList = make([]*proto.Challenge, 0)
	rsp.ChallengeGroupList = make([]*proto.ChallengeGroup, 0)
	for id, stars := range g.GetChallengeList() {
		challenge := &proto.Challenge{
			ChallengeId: id,
			Star:        stars.Stars,
			ScoreId:     stars.ScoreOne,
			ScoreTwo:    stars.ScoreTwo,
			TakenReward: 0,
		}
		rsp.ChallengeList = append(rsp.ChallengeList, challenge)
	}
	for taken, id := range g.GetChallengeRewardList() {
		challengeGroup := &proto.ChallengeGroup{
			TakenStarsCountReward: taken,
			GroupId:               id,
		}
		rsp.ChallengeGroupList = append(rsp.ChallengeGroupList, challengeGroup)
	}
	g.Send(cmd.GetChallengeScRsp, rsp)
}

// 获取状态

func (g *GamePlayer) GetCurChallengeCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetCurChallengeScRsp{
		CurChallenge: g.GetChallengeInfo(),
		LineupList: []*proto.LineupInfo{
			g.GetLineUpPb(g.GetBattleLineUpById(Challenge_1)),
			g.GetLineUpPb(g.GetBattleLineUpById(Challenge_2)),
		},
	}
	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

// 进入忘却之庭

func (g *GamePlayer) StartChallengeCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartChallengeCsReq)
	// 设置战斗状态
	storyInfo := req.GetPlayerInfo()
	if storyInfo == nil {
		g.SetBattleStatus(spb.BattleType_Battle_CHALLENGE)
	} else {
		g.SetBattleStatus(spb.BattleType_Battle_CHALLENGE_Story)
	}

	// 设置队伍
	if req.FirstLineup == nil {
		g.SetBattleStatus(spb.BattleType_Battle_NONE)
		g.Send(cmd.StartChallengeScRsp, &proto.StartChallengeScRsp{})
		return
	}
	g.Send(cmd.SyncServerSceneChangeNotify, &proto.SyncServerSceneChangeNotify{})
	g.SetBattleLineUp(Challenge_1, req.FirstLineup)
	if req.SecondLineup != nil {
		g.SetBattleLineUp(Challenge_2, req.SecondLineup)
	}
	// 设置当前战斗的忘却之庭
	g.SetCurChallenge(req.ChallengeId, storyInfo)
	rsp := &proto.StartChallengeScRsp{
		CurChallenge: g.GetChallengeInfo(),
		Scene:        g.GetChallengeScene(),
		LineupList: []*proto.LineupInfo{
			g.GetLineUpPb(g.GetBattleLineUpById(Challenge_1)),
			g.GetLineUpPb(g.GetBattleLineUpById(Challenge_2)),
		},
	}

	g.Send(cmd.StartChallengeScRsp, rsp)
}

// 忘却之庭战斗退出/结束

func (g *GamePlayer) LeaveChallengeCsReq(payloadMsg pb.Message) {
	curChallenge := g.GetCurChallenge()
	if curChallenge == nil {
		return
	}
	if curChallenge.Status == spb.ChallengeStatus_CHALLENGE_DOING {
		g.Send(cmd.QuitBattleScNotify, nil) // 战斗没结束就退出是主动退出
	}
	g.ChallengeSettleNotify()
	g.Send(cmd.LeaveChallengeScRsp, nil)

	g.EnterSceneByServerScNotify(g.GetCurEntryId(), 0, 0, 0)
	// 设置战斗状态为空
	g.SetBattleStatus(spb.BattleType_Battle_NONE)
	// 清空忘却之庭
	g.NewCurChallenge()
}

// 忘却之庭世界战斗结算事件

func (g *GamePlayer) ChallengePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	stt := req.GetStt()
	if stt != nil {
		g.SetCurChallengeRoundCount(req.Stt.GetRoundCnt())  // 更新已使用回合数
		g.SetCurChallengeScore(req.Stt.GetChallengeScore()) // 更新分数
	}
	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_NONE:
		return
	case proto.BattleEndStatus_BATTLE_END_LOSE: // 战斗失败
		if !g.ChallengeBattleEndLose() {
			return
		}
	case proto.BattleEndStatus_BATTLE_END_QUIT: // 退出战斗
		return
	}
	g.AddCurChallengeKillMonster(1) // TODO 默认一次战斗一个怪物
	// 场景上是否还有未处理敌人
	if g.GetCurChallengeMonsterNum() > g.GetCurChallengeKillMonster() {
		return // 还有就不更新状态，继续进行
	}
	// 更新状态
	g.SetCurChallengeKillMonster(0) // 切换关卡，标记为0
	if db.IsBoos {
		g.ChallengeBossPhaseSettleNotify(req.Stt.GetBattleTargetInfo())
	}
	// 回合处理
	if g.IsNextChallenge() {
		// 还没结束
		g.AddChallengeCurStage(1)
		// 添加怪物
		g.ChallengeAddSceneGroupRefreshScNotify()
		if !db.IsBoos {
			// 添加角色
			g.ChallengeAddAvatarSceneGroupRefreshScNotify()
			// 更新新的队伍
			g.SyncLineupNotify(g.GetBattleLineUpById(Challenge_2))
		}
	} else {
		// 结算
		g.ChallengeSettle()
		// 发送战斗胜利通知
		if !db.IsBoos {
			g.ChallengeSettleNotify()
		}
		// 将战斗结果储存到数据库
		g.UpdateChallengeList(db)
		// 更改状态
		g.SetCurChallengeStatus(spb.ChallengeStatus_CHALLENGE_FINISH)
	}
}

func (g *GamePlayer) ChallengeSettleNotify() {
	db := g.GetCurChallenge()
	notify := &proto.ChallengeSettleNotify{
		Star:           db.Stars,               // 得分
		Reward:         g.GetChallengeReward(), // 奖励
		ChallengeId:    db.ChallengeId,         // 关卡id
		IsWin:          db.IsWin,               // 是否赢
		ScoreTwo:       db.ScoreTwo,            // 二层挑战得分
		ChallengeScore: db.ScoreOne,            // 一层挑战得分
	}
	g.Send(cmd.ChallengeSettleNotify, notify)
}

func (g *GamePlayer) ChallengeBossPhaseSettleNotify(targeList map[uint32]*proto.BattleTargetList) {
	db := g.GetCurChallenge()
	notify := &proto.ChallengeBossPhaseSettleNotify{
		IsRemainingAction: true,
		Star:              7,
		Phase:             db.CurStage,
		ChallengeScore:    db.ScoreOne,
		ScoreTwo:          db.ScoreTwo,
		IsWin:             true,
		BattleTargetList:  targeList[1].BattleTargetList,
		ChallengeId:       db.ChallengeId,
	}
	g.Send(cmd.ChallengeBossPhaseSettleNotify, notify)
}

func (g *GamePlayer) ChallengeAddAvatarSceneGroupRefreshScNotify() {
	curChallenge := g.GetCurChallenge()
	mazeGroupID := g.GetChallengesMazeGroupID()
	lineUp := g.GetChallengesLineUp()
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return
	}
	mapEntrance := gdconf.GetMapEntranceById(challengeMazeConfig.MapEntranceID)
	if mapEntrance == nil {
		return
	}
	foorMap := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, mazeGroupID)
	if foorMap == nil {
		return
	}
	pos, rot := g.GetChallengesAnchor(foorMap.AnchorList)
	if pos == nil || rot == nil {
		return
	}

	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		RefreshEntity: g.GetAddAvatarSceneEntityRefreshInfo(lineUp, pos, rot),
	}
	notify.GroupRefreshList = append(notify.GroupRefreshList, sceneGroupRefreshInfo)
	g.Send(cmd.SceneGroupRefreshScNotify, notify)

	// 通知当前位置
	sceneEntityMoveScNotify := &proto.SceneEntityMoveScNotify{
		EntryId:          challengeMazeConfig.MapEntranceID,
		ClientPosVersion: 0,
		Motion: &proto.MotionInfo{
			Pos: pos,
			Rot: rot,
		},
	}
	g.Send(cmd.SceneEntityMoveScNotify, sceneEntityMoveScNotify)
}

func (g *GamePlayer) ChallengeAddSceneGroupRefreshScNotify() {
	curChallenge := g.GetCurChallenge()
	mazeGroupID := g.GetChallengesMazeGroupID()
	configList := g.GetChallengesConfigList()
	eventIDList := g.GetChallengesEventIDList()
	npcMonsterIDList := g.GetChallengesNpcMonsterIDList()
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return
	}
	mapEntrance := gdconf.GetMapEntranceById(challengeMazeConfig.MapEntranceID)
	if mapEntrance == nil {
		return
	}
	foorMap := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, mazeGroupID)
	if foorMap == nil || len(npcMonsterIDList) != len(eventIDList) || len(eventIDList) != len(configList) {
		return
	}

	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		GroupId:       mazeGroupID,
		RefreshEntity: g.ChallengesAddMonsterSceneEntityRefreshInfo(mazeGroupID, configList, eventIDList, npcMonsterIDList, foorMap.MonsterList),
	}
	notify.GroupRefreshList = append(notify.GroupRefreshList, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}

// 添加怪物
func (g *GamePlayer) ChallengesAddMonsterSceneEntityRefreshInfo(mazeGroupID uint32, configList, eventIDList, npcMonsterIDList []uint32, monsterList map[uint32]*gdconf.MonsterList) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for id, config := range configList {
		for _, monster := range monsterList {
			if monster.ID != config {
				continue
			}
			entityId := g.GetNextGameObjectGuid()
			monsterPos := &proto.Vector{
				X: int32(monster.PosX * 1000),
				Y: int32(monster.PosY * 1000),
				Z: int32(monster.PosZ * 1000),
			}
			monsterRot := &proto.Vector{
				X: int32(monster.RotX * 1000),
				Y: int32(monster.RotY * 1000),
				Z: int32(monster.RotZ * 1000),
			}
			seri := &proto.SceneEntityRefreshInfo{
				Refresh: &proto.SceneEntityRefreshInfo_AddEntity{
					AddEntity: &proto.SceneEntityInfo{
						GroupId:  mazeGroupID,
						InstId:   monster.ID,
						EntityId: entityId,
						Motion: &proto.MotionInfo{
							Pos: monsterPos,
							Rot: monsterRot,
						},
						EntityOneofCase: &proto.SceneEntityInfo_NpcMonster{
							NpcMonster: &proto.SceneNpcMonsterInfo{
								MonsterId: npcMonsterIDList[id],
								EventId:   eventIDList[id],
							},
						},
					},
				},
			}
			// 添加怪物实体
			g.AddEntity(mazeGroupID, &MonsterEntity{
				Entity: Entity{
					EntityId: entityId,
					GroupId:  mazeGroupID,
					Pos:      monsterPos,
					Rot:      monsterRot,
					InstId:   monster.ID,
				},
				EventID: eventIDList[id],
			})
			sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
		}
	}
	return sceneEntityRefreshInfo
}

func (g *GamePlayer) StartPartialChallengeCsReq(payloadMsg pb.Message) {
	// msg := g.DecodePayloadToProto(cmd.StartPartialChallengeCsReq, payloadMsg)
	// req := msg.(*proto.StartPartialChallengeCsReq)
	// g.SetBattleStatus(spb.BattleType_Battle_CHALLENGE_Story_2)
	// // 设置当前战斗的忘却之庭
	// g.SetCurChallenge(req.ChallengeId, storyInfo)
}

func (g *GamePlayer) EnterChallengeNextPhaseCsReq(payloadMsg pb.Message) {
	rsp := &proto.EnterChallengeNextPhaseScRsp{
		Scene: g.GetChallengeScene(),
	}

	g.Send(cmd.EnterChallengeNextPhaseScRsp, rsp)
}

// 获取忘却之庭世界
func (g *GamePlayer) GetChallengeScene() *proto.SceneInfo {
	curChallenge := g.GetCurChallenge()
	leaderEntityId := g.GetNextGameObjectGuid()
	lineUp := g.GetChallengesLineUp()
	mazeGroupID := g.GetChallengesMazeGroupID()
	configList := g.GetChallengesConfigList()
	npcMonsterIDList := g.GetChallengesNpcMonsterIDList()
	eventIDList := g.GetChallengesEventIDList()
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return nil
	}
	mapEntrance := gdconf.GetMapEntranceById(challengeMazeConfig.MapEntranceID)
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, mazeGroupID)
	if foorMap == nil || lineUp == nil || len(npcMonsterIDList) != len(eventIDList) || len(eventIDList) != len(configList) {
		return nil
	}
	pos, rot := g.GetChallengesAnchor(foorMap.AnchorList)
	if pos == nil || rot == nil {
		return nil
	}
	// 获取映射信息
	worldId := gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID
	if worldId == 100 {
		worldId = 401
	}
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            mapEntrance.PlaneID,
		FloorId:            mapEntrance.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            worldId,
		EntryId:            challengeMazeConfig.MapEntranceID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		EnvBuffList:        make([]*proto.BuffInfo, 0),
		LevelGroupIdList:   make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		LightenSectionList: []uint32{0},
		EntityList:         make([]*proto.SceneEntityInfo, 0),
	}

	// 添加场景buff
	if curChallenge.MazeBuffId != 0 {
		scene.EnvBuffList = append(scene.EnvBuffList, &proto.BuffInfo{
			Count:     4294967295,
			LifeTime:  -1,
			BuffId:    curChallenge.MazeBuffId,
			AddTimeMs: uint64(time.Now().UnixMilli()),
			Level:     1,
		})
	}
	// 添加自选buff
	if g.GetCurChallengeBuffId() != 0 {
		scene.EnvBuffList = append(scene.EnvBuffList, &proto.BuffInfo{
			Count:     4294967295,
			LifeTime:  -1,
			BuffId:    g.GetCurChallengeBuffId(),
			AddTimeMs: uint64(time.Now().UnixMilli()),
			Level:     1,
		})
	}
	// 将进入场景的角色添加到实体列表里
	entityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    0,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	g.GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)
	// 添加怪物实体
	monsterEntityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    mazeGroupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for id, config := range configList {
		for _, monsterList := range foorMap.MonsterList {
			if monsterList.ID != config {
				continue
			}
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
				GroupId:  mazeGroupID,
				InstId:   monsterList.ID,
				EntityId: entityId,
				Motion: &proto.MotionInfo{
					Pos: monsterPos,
					Rot: monsterRot,
				},
				EntityOneofCase: &proto.SceneEntityInfo_NpcMonster{
					NpcMonster: &proto.SceneNpcMonsterInfo{
						MonsterId: npcMonsterIDList[id],
						EventId:   eventIDList[id],
					},
				},
			}
			// 添加怪物实体
			g.AddEntity(mazeGroupID, &MonsterEntity{
				Entity: Entity{
					EntityId: entityId,
					GroupId:  mazeGroupID,
					Pos:      monsterPos,
					Rot:      monsterRot,
					InstId:   monsterList.ID,
				},
				EventID: eventIDList[id],
			})
			monsterEntityGroup.EntityList = append(monsterEntityGroup.EntityList, entityList)
		}
	}
	scene.EntityGroupList = append(scene.EntityGroupList, monsterEntityGroup)
	return scene
}
