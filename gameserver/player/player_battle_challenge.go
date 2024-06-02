package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/***********************************忘却之庭***********************************/

func (g *GamePlayer) HandleGetChallengeCsReq(payloadMsg []byte) {
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

func (g *GamePlayer) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := &proto.GetCurChallengeScRsp{
		CurChallenge: g.GetChallengeInfo(),
	}
	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

// 进入忘却之庭

func (g *GamePlayer) StartChallengeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartChallengeCsReq, payloadMsg)
	req := msg.(*proto.StartChallengeCsReq)
	// 设置战斗状态
	storyInfo := req.GetPlayerInfo()
	if storyInfo == nil {
		g.SetBattleStatus(spb.BattleType_Battle_CHALLENGE)
	} else {
		g.SetBattleStatus(spb.BattleType_Battle_CHALLENGE_Story)
	}

	var lineUpId uint32
	// 设置当前战斗的忘却之庭
	curChallenge := g.SetCurChallenge(req.ChallengeId, storyInfo)
	switch curChallenge.CurStage {
	case 1:
		lineUpId = Challenge_1
	case 2:
		lineUpId = Challenge_2
	}
	rsp := &proto.StartChallengeScRsp{
		CurChallenge: g.GetChallengeInfo(),
		Scene:        g.GetChallengeScene(),
		Lineup:       g.GetBattleLineUpPb(lineUpId),
	}

	g.Send(cmd.StartChallengeScRsp, rsp)
}

// 忘却之庭战斗退出/结束

func (g *GamePlayer) LeaveChallengeCsReq(payloadMsg []byte) {
	curChallenge := g.GetCurChallenge()
	if curChallenge == nil {
		return
	}
	if curChallenge.Status == spb.ChallengeStatus_CHALLENGE_DOING {
		g.Send(cmd.QuitBattleScNotify, nil) // 战斗没结束就退出是主动退出
	}
	g.Send(cmd.LeaveChallengeScRsp, nil)

	g.EnterSceneByServerScNotify(g.GetCurEntryId(), 0)
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
	// 回合处理
	if g.IsNextChallenge() {
		// 还没结束
		g.AddChallengeCurStage(1)
		// 添加怪物
		g.ChallengeAddSceneGroupRefreshScNotify()
		// 添加角色
		g.ChallengeAddAvatarSceneGroupRefreshScNotify()
		// 更新新的队伍
		g.SyncLineupNotify(Challenge_2, true)
	} else {
		// 结算
		g.ChallengeSettle()
		// 发送战斗胜利通知
		g.ChallengeSettleNotify()
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
				AddEntity: &proto.SceneEntityInfo{
					GroupId:  mazeGroupID,
					InstId:   monster.ID,
					EntityId: entityId,
					Motion: &proto.MotionInfo{
						Pos: monsterPos,
						Rot: monsterRot,
					},
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
				},
				EventID: eventIDList[id],
			})
			sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
		}
	}
	return sceneEntityRefreshInfo
}
