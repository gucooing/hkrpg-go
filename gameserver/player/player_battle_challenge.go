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
	// if challengeState.ChallengeCount == 1 {
	// 	rsp.ChallengeInfo.StoryInfo.StoryBuffs = &proto.ChallengeStoryInfo_CurStoryBuffs{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: []uint32{challengeState.StoryBuffOne}}}
	// 	// rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffOne)
	// } else {
	// 	rsp.ChallengeInfo.StoryInfo.StoryBuffs = &proto.ChallengeStoryInfo_CurStoryBuffs{CurStoryBuffs: &proto.ChallengeStoryBuffInfo{BuffList: []uint32{challengeState.StoryBuffTwo}}}
	// 	// rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList = append(rsp.ChallengeInfo.StoryInfo.CurStoryBuffs.BuffList, challengeState.StoryBuffTwo)
	// }
	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

// 进入忘却之庭

func (g *GamePlayer) StartChallengeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartChallengeCsReq, payloadMsg)
	req := msg.(*proto.StartChallengeCsReq)
	var lineUpId uint32
	// 设置当前战斗的忘却之庭
	curChallenge := g.SetCurChallenge(req.ChallengeId)
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
	// 设置战斗状态
	g.SetBattleStatus(spb.BattleType_Battle_CHALLENGE)

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
	// battleStatus := g.GetBattleStatus()
	// curChallenge := g.GetCurChallenge()
	// 战斗失败
	if req.EndStatus == proto.BattleEndStatus_BATTLE_END_LOSE {
		return
	}
	// 更新状态
	g.SetCurChallengeRoundCount(req.Stt.GetRoundCnt() + 1) // 更新已使用回合数

	// // 是否还有一关
	// if g.IsNextChallenge() {
	// 	// 战斗正常结束进入结算
	// 	// 计算分数
	// 	var stage uint32 = 0
	// 	for _, challengeTargetID := range challengeState.ChallengeTargetID {
	// 		challengeTargetConfig := gdconf.GetChallengeTargetConfigById(challengeTargetID)
	// 		if challengeTargetConfig.ChallengeTargetType == "DEAD_AVATAR" {
	// 			// 是否有角色死亡
	// 			stage += 3
	// 		} else {
	// 			if (challengeState.ChallengeCountDown - challengeState.RoundCount) >= challengeTargetConfig.ChallengeTargetParam1 {
	// 				stage += 2
	// 			}
	// 		}
	// 	}
	//
	// 	// 将战斗结果储存到数据库
	// 	challengeDb := g.GetChallengeById(challengeState.ChallengeId)
	// 	if challengeDb.Stars < stage {
	// 		challengeDb.Stars = stage
	// 	}
	// 	// 发送战斗胜利通知
	// 	challengeSettleNotify := &proto.ChallengeSettleNotify{
	// 		Stars:       stage,
	// 		Reward:      nil, // TODO 记得发奖励
	// 		ChallengeId: challengeState.ChallengeId,
	// 		IsWin:       true,
	// 	}
	// 	g.Send(cmd.ChallengeSettleNotify, challengeSettleNotify)
	// 	// 战斗正式结束，还原战斗信息
	// 	battleState.BattleType = spb.BattleType_Battle_NONE
	// 	challengeState.Status = proto.ChallengeStatus_CHALLENGE_FINISH
	// }

	if g.IsNextChallenge() {
		// 还没结束
		g.AddChallengeCurStage(1)
		// 添加怪物
		g.ChallengeAddSceneGroupRefreshScNotify()
		// 添加角色
		g.ChallengeAddAvatarSceneGroupRefreshScNotify()
		// 更新新的队伍
		g.SyncLineupNotify(Challenge_2, true)
		// // 通知当前战斗的队伍
		// challengeLineupNotify := &proto.ChallengeLineupNotify{
		// 	ExtraLineupType: proto.ExtraLineupType_LINEUP_CHALLENGE_2,
		// }
		// g.Send(cmd.ChallengeLineupNotify, challengeLineupNotify)
		// 通知坐标
		// g.SceneEntityMoveScNotify(pos, rot, challengeState.EntranceID)
	}
}

func (g *GamePlayer) ChallengeSyncLineupNotify(index uint32) {
	rsq := new(proto.SyncLineupNotify)
	lineUp := g.GetBattleLineUpById(index)
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           0,
		ExtraLineupType: proto.ExtraLineupType(index),
		MaxMp:           5,
		Mp:              5,
		PlaneId:         0,
	}
	for slot, lineAvatar := range lineUp.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: proto.AvatarType(avatarBin.AvatarType),
			Slot:       slot,
			Satiety:    0,
			Hp:         10000,
			Id:         avatarBin.AvatarId,
			SpBar: &proto.SpBarInfo{
				CurSp: 10000,
				MaxSp: 10000,
			},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}
	rsq.Lineup = lineupList

	g.Send(cmd.SyncLineupNotify, rsq)
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

// 下面是活动

func (g *GamePlayer) ChallengeStorySceneCastSkillCsReq(rsp *proto.SceneCastSkillScRsp) {
	challengeState := g.GetChallengeState()
	// var lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE)
	// // var targetIndex uint32 = 0
	storyMazeExtra := gdconf.GetChallengeStoryMazeExtraById(challengeState.ChallengeId)
	//
	// // 通过波次获取队伍
	// if challengeState.ExtraLineupType == proto.ExtraLineupType_LINEUP_CHALLENGE {
	// 	lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE)
	// } else {
	// 	lineUpId = uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2)
	// }
	//
	// // 添加角色
	// rsp.BattleInfo.BattleAvatarList = g.GetBattleAvatarList(lineUpId)
	// 添加回合限制
	rsp.BattleInfo.RoundsLimit = challengeState.ChallengeCountDown

	// 添加关卡buff
	if challengeState.CurChallengeCount == 1 {
		buffListStory := &proto.BattleBuff{
			Id:              challengeState.StoryBuffOne,
			Level:           1,
			OwnerId:         0,
			TargetIndexList: []uint32{0},
			WaveFlag:        4294967295, // 失效时间
			DynamicValues:   make(map[string]float32),
		}
		buffListStory.DynamicValues["SkillIndex"] = 1
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffListStory)
	} else {
		buffListStory := &proto.BattleBuff{
			Id:              challengeState.StoryBuffTwo,
			Level:           1,
			OwnerId:         0,
			TargetIndexList: []uint32{0},
			WaveFlag:        4294967295, // 失效时间
			DynamicValues:   make(map[string]float32),
		}
		buffListStory.DynamicValues["SkillIndex"] = 1
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffListStory)
	}

	// 添加场景buff
	for _, buffId := range challengeState.SceneBuffList {
		buffList := &proto.BattleBuff{
			Id:       buffId,
			Level:    1,
			OwnerId:  4294967295,
			WaveFlag: 4294967295, // 失效时间
		}
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
	}
	// 添加角色buff
	/*
		for _, buffId := range challengeState.AvatarBuffList {
			buffList := &proto.BattleBuff{
				Id:              buffId,
				Level:           1,
				OwnerId:         targetIndex,
				TargetIndexList: []uint32{targetIndex},
				WaveFlag:        4294967295, // 失效时间
			}
			rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, buffList)
			targetIndex++
		}
	*/
	rsp.BattleInfo.BattleTargetInfo = make(map[uint32]*proto.BattleTargetList)
	rsp.BattleInfo.BattleTargetInfo[1] = &proto.BattleTargetList{
		BattleTargetList: []*proto.BattleTarget{{
			Id: 10001,
		}},
	}
	battleTargetList := make([]*proto.BattleTarget, 0)
	for _, id := range storyMazeExtra.BattleTargetID {
		battleTarget := &proto.BattleTarget{
			Id:       id,
			Progress: 0,
		}
		battleTargetList = append(battleTargetList, battleTarget)
	}
	rsp.BattleInfo.BattleTargetInfo[5] = &proto.BattleTargetList{
		BattleTargetList: battleTargetList,
	}

	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

func (g *GamePlayer) ChallengeStoryPVEBattleResultCsReq(req *proto.PVEBattleResultCsReq) {
	battleState := g.GetBattleState()
	challengeState := g.GetChallengeState()
	pos := challengeState.Pos
	rot := challengeState.Rot

	if challengeState.ExtraLineupType == proto.ExtraLineupType_LINEUP_CHALLENGE {
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE))
	} else {
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2))
	}

	// 删除实体
	nitify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshInfo: make([]*proto.SceneGroupRefreshInfo, 0),
	}
	// for _, eventId := range challengeState.MonsterEntityMap {
	// 	entity := g.GetSceneEntity().MonsterEntity[eventId]
	// 	if entity != nil {
	// 		groupRefreshInfo := &proto.SceneGroupRefreshInfo{
	// 			GroupId: entity.GroupId,
	// 			RefreshEntity: []*proto.SceneEntityRefreshInfo{
	// 				{
	// 					DelEntity: eventId,
	// 				},
	// 			},
	// 		}
	// 		nitify.GroupRefreshInfo = append(nitify.GroupRefreshInfo, groupRefreshInfo)
	// 		delete(g.GetSceneEntity().MonsterEntity, eventId)
	// 	}
	// }
	g.Send(cmd.SceneGroupRefreshScNotify, nitify)

	// 获取分数
	challengeState.ChallengeScore += req.Stt.ChallengeScore
	if challengeState.CurChallengeCount == 1 {
		challengeState.ScoreOne = req.Stt.ChallengeScore
	} else {
		challengeState.ScoreTwo = req.Stt.ChallengeScore
	}
	// 通过波次数判断是否还有一关
	if challengeState.CurChallengeCount == challengeState.ChallengeCount {
		// 战斗正常结束进入结算

		// 计算分数
		var stage uint32 = 0
		if challengeState.ChallengeScore >= 30000 {
			stage++
		}
		for _, challengeTargetID := range challengeState.ChallengeTargetID {
			challengeTargetConfig := gdconf.GetChallengeTargetConfigById(challengeTargetID)
			if challengeState.ChallengeScore >= challengeTargetConfig.ChallengeTargetParam1 {
				stage += 2
			}
		}

		// 将战斗结果储存到数据库
		challengeDb := g.GetChallengeById(challengeState.ChallengeId)
		if challengeDb.Stars < stage {
			challengeDb.Stars = stage
		}
		if challengeDb.ScoreOne < challengeState.ScoreOne {
			challengeDb.ScoreOne = challengeState.ScoreOne
		}
		if challengeDb.ScoreTwo < challengeState.ScoreTwo {
			challengeDb.ScoreTwo = challengeState.ScoreTwo
		}
		// 发送战斗胜利通知
		challengeSettleNotify := &proto.ChallengeSettleNotify{
			Stars:          stage,
			Reward:         nil, // TODO 记得发奖励
			ChallengeId:    challengeState.ChallengeId,
			IsWin:          true,
			ChallengeScore: challengeState.ScoreOne,
			ScoreTwo:       challengeState.ScoreTwo,
		}
		g.Send(cmd.ChallengeSettleNotify, challengeSettleNotify)
		// 战斗正式结束，还原战斗信息
		battleState.BattleType = spb.BattleType_Battle_NONE
		challengeState.Status = proto.ChallengeStatus_CHALLENGE_FINISH
	} else {
		// 还差一波
		challengeState.CurChallengeCount++
		challengeState.ExtraLineupType = proto.ExtraLineupType_LINEUP_CHALLENGE_2
		// challengeState.AvatarBuffList = make([]uint32, 0)
		// 添加怪物
		g.ChallengeAddSceneGroupRefreshScNotify()
		// 添加角色
		g.ChallengeAddAvatarSceneGroupRefreshScNotify()
		// 更新新的队伍
		g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_CHALLENGE_2))
		// 通知当前战斗的队伍
		challengeLineupNotify := &proto.ChallengeLineupNotify{
			ExtraLineupType: challengeState.ExtraLineupType,
		}
		g.Send(cmd.ChallengeLineupNotify, challengeLineupNotify)
		// 通知坐标
		g.SceneEntityMoveScNotify(pos, rot, challengeState.EntranceID)
	}
}
