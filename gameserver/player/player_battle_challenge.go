package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

/***********************************忘却之庭***********************************/

// 获取状态

func (g *GamePlayer) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := &proto.GetCurChallengeScRsp{
		ChallengeInfo: g.GetChallengeInfo(),
	}
	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

// 进入忘却之庭

func (g *GamePlayer) StartChallengeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartChallengeCsReq, payloadMsg)
	req := msg.(*proto.StartChallengeCsReq)
	// 设置战斗状态
	storyInfo := req.GetStoryInfo()
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
		ChallengeInfo: g.GetChallengeInfo(),
		Scene:         g.GetChallengeScene(),
		Lineup:        g.GetBattleLineUpPb(lineUpId),
	}

	g.Send(cmd.StartChallengeScRsp, rsp)
}

// 忘却之庭战斗退出/结束

func (g *GamePlayer) LeaveChallengeCsReq(payloadMsg []byte) {
	curChallenge := g.GetCurChallenge()
	if proto.ChallengeStatus(curChallenge.Status) == proto.ChallengeStatus_CHALLENGE_DOING {
		g.Send(cmd.QuitBattleScNotify, nil)
	}
	g.Send(cmd.LeaveChallengeScRsp, nil)

	g.EnterSceneByServerScNotify(g.GetScene().EntryId, 0)
	g.GetBattleState().BattleType = spb.BattleType_Battle_NONE
	g.GetBattleState().BuffList = make([]uint32, 0)
}

// 忘却之庭世界战斗结算事件

func (g *GamePlayer) ChallengePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_NONE:
		return
	case proto.BattleEndStatus_BATTLE_END_LOSE: // 战斗失败
		g.SetCurChallengeStatus(spb.ChallengeStatus_CHALLENGE_UNKNOWN)
		g.ChallengeSettleNotify()
		// 设置战斗状态为空
		g.SetBattleStatus(spb.BattleType_Battle_NONE)
		// 清空忘却之庭
		g.NewCurChallenge()
		return
	case proto.BattleEndStatus_BATTLE_END_QUIT: // 退出战斗
		// 设置战斗状态为空
		g.SetBattleStatus(spb.BattleType_Battle_NONE)
		// 清空忘却之庭
		g.NewCurChallenge()
		return
	}
	// 更新状态
	g.SetCurChallengeRoundCount(req.Stt.GetRoundCnt()) // 更新已使用回合数
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
		// 设置战斗状态为空
		g.SetBattleStatus(spb.BattleType_Battle_NONE)
		// 清空忘却之庭
		g.NewCurChallenge()
	}
}

func (g *GamePlayer) ChallengeSettleNotify() {
	db := g.GetCurChallenge()
	notify := &proto.ChallengeSettleNotify{
		Stars:          db.Stars,                  // 得分
		Reward:         g.GetChallengeReward(),    // 奖励
		ChallengeId:    db.ChallengeId,            // 关卡id
		IsWin:          db.IsWin,                  // 是否赢
		ScoreTwo:       db.ScoreTwo,               // 二层挑战得分
		ChallengeScore: db.ScoreOne + db.ScoreTwo, // 总得分
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
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(challengeMazeConfig.MapEntranceID)))
	if mapEntrance == nil {
		return
	}
	foorMap := gdconf.GetMazeByGroupId(mapEntrance.PlaneID, mapEntrance.FloorID, mazeGroupID)
	if foorMap == nil {
		return
	}
	pos, rot := g.GetChallengesAnchor(foorMap.AnchorList)
	if pos == nil || rot == nil {
		return
	}

	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshInfo: make([]*proto.SceneGroupRefreshInfo, 0),
	}
	sceneGroupRefreshInfo := &proto.SceneGroupRefreshInfo{
		RefreshEntity: g.GetAddAvatarSceneEntityRefreshInfo(lineUp, pos, rot),
	}
	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)
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
	mapEntrance := gdconf.GetMapEntranceById(strconv.Itoa(int(challengeMazeConfig.MapEntranceID)))
	if mapEntrance == nil {
		return
	}
	foorMap := gdconf.GetMazeByGroupId(mapEntrance.PlaneID, mapEntrance.FloorID, mazeGroupID)
	if foorMap == nil || len(npcMonsterIDList) != len(eventIDList) || len(eventIDList) != len(configList) {
		return
	}

	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshInfo: make([]*proto.SceneGroupRefreshInfo, 0),
	}
	sceneGroupRefreshInfo := &proto.SceneGroupRefreshInfo{
		GroupId:       mazeGroupID,
		RefreshEntity: g.GetAddMonsterSceneEntityRefreshInfo(mazeGroupID, configList, eventIDList, npcMonsterIDList, foorMap.MonsterList),
	}
	notify.GroupRefreshInfo = append(notify.GroupRefreshInfo, sceneGroupRefreshInfo)

	g.Send(cmd.SceneGroupRefreshScNotify, notify)
}
