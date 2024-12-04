package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

/***********************************忘却之庭***********************************/

func HandleGetChallengeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetChallengeScRsp)
	rsp.ChallengeList = make([]*proto.Challenge, 0)
	rsp.ChallengeGroupList = make([]*proto.ChallengeGroup, 0)
	for groupId, infoList := range g.GetPd().GetChallengeGroupList() {
		challengeGroup := &proto.ChallengeGroup{
			TakenStarsCountReward: infoList.ChallengeReward,
			GroupId:               groupId,
		}
		rsp.ChallengeGroupList = append(rsp.ChallengeGroupList, challengeGroup)
		for _, info := range infoList.ChallengeInfoList {
			challenge := &proto.Challenge{
				ChallengeId: info.ChallengeId,
				Star:        info.Stars,
				ScoreId:     info.ScoreOne,
				ScoreTwo:    info.ScoreTwo,
				TakenReward: 0,
			}
			rsp.ChallengeList = append(rsp.ChallengeList, challenge)
		}
	}
	g.Send(cmd.GetChallengeScRsp, rsp)
}

func TakeChallengeRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeChallengeRewardCsReq)
	rsp := &proto.TakeChallengeRewardScRsp{
		TakenRewardList: make([]*proto.TakenChallengeRewardInfo, 0),
		GroupId:         req.GroupId,
		Retcode:         0,
	}
	addItem := model.NewAddItem(nil)

	db := g.GetPd().GetChallengeGroupInfoById(req.GroupId)
	conf := gdconf.GetChallengeGroupConfig(req.GroupId)
	if conf == nil {
		g.Send(cmd.TakeChallengeRewardScRsp, rsp)
		return
	}
	var start uint32
	for _, v := range db.ChallengeInfoList {
		start += model.GetChallengeStars(v.Stars)
	}
	curStart := model.GetTakenRewards(db.ChallengeReward)
	for curStart < start {
		curStart++
		rewardID := gdconf.GetChallengeRewardLineRewardID(conf.RewardLineGroupID, curStart)
		if rewardID == 0 {
			continue
		}
		pile := model.GetRewardData(rewardID)
		addItem.PileItem = append(addItem.PileItem, pile...)
		g.GetPd().AddItem(addItem)
		db.ChallengeReward = model.SetTakenReward(db.ChallengeReward, curStart)
	}
	rsp.TakenRewardList = append(rsp.TakenRewardList, &proto.TakenChallengeRewardInfo{
		StarCount: curStart,
		Reward: &proto.ItemList{
			ItemList: addItem.ItemList,
		},
	})
	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.Send(cmd.TakeChallengeRewardScRsp, rsp)
}

func GetChallengeGroupStatisticsCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetChallengeGroupStatisticsCsReq)
	rsp := &proto.GetChallengeGroupStatisticsScRsp{
		GroupId: req.GroupId,
		Retcode: 0,
	}
	switch req.GroupId / 1000 {
	case 1:
		rsp.Challenge = g.GetPd().GetChallengeGroupStatisticsChallengeDefault(req.GroupId)
	case 2:
		rsp.Challenge = g.GetPd().GetChallengeGroupStatisticsChallengeStory(req.GroupId)
	case 3:
		rsp.Challenge = g.GetPd().GetChallengeGroupStatisticsChallengeBoss(req.GroupId)
	default:
		logger.Warn("challenge error")
	}
	g.Send(cmd.GetChallengeGroupStatisticsScRsp, rsp)
}

// 获取状态

func GetCurChallengeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetCurChallengeScRsp{
		CurChallenge: g.GetPd().GetChallengeInfo(),
		LineupList: []*proto.LineupInfo{
			g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Challenge_1)),
			g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Challenge_2)),
		},
	}
	g.Send(cmd.GetCurChallengeScRsp, rsp)
}

// 进入忘却之庭

func StartChallengeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartChallengeCsReq)
	// 设置战斗状态
	storyInfo := req.GetStageInfo()
	if storyInfo == nil {
		g.GetPd().SetBattleStatus(spb.BattleType_Battle_CHALLENGE)
	} else {
		g.GetPd().SetBattleStatus(spb.BattleType_Battle_CHALLENGE_Story)
	}

	// 设置队伍
	if req.FirstLineup == nil {
		g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
		g.Send(cmd.StartChallengeScRsp, &proto.StartChallengeScRsp{})
		return
	}
	g.Send(cmd.SyncServerSceneChangeNotify, &proto.SyncServerSceneChangeNotify{})
	g.SetBattleLineUp(model.Challenge_1, req.FirstLineup)
	if req.SecondLineup != nil {
		g.SetBattleLineUp(model.Challenge_2, req.SecondLineup)
	}
	// 设置当前战斗的忘却之庭
	g.GetPd().SetCurChallenge(req)
	rsp := &proto.StartChallengeScRsp{
		CurChallenge: g.GetPd().GetChallengeInfo(),
		Scene:        g.GetChallengeScene(),
		LineupList: []*proto.LineupInfo{
			g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Challenge_1)),
			g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Challenge_2)),
		},
	}

	g.Send(cmd.StartChallengeScRsp, rsp)
}

// 忘却之庭重新挑战
func RestartChallengePhaseCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.RestartChallengePhaseScRsp{
		Retcode: 0,
		Scene:   nil,
	}
	db := g.GetPd().GetCurChallenge()
	if db == nil || db.ChallengeId == 0 { // 判断是否合法
		rsp.Retcode = uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN)
	} else {
		// 关卡回调
		if db.CurStage == 1 {
			db.ScoreOne = 0
		} else {
			db.ScoreTwo = 0
		}
		db.IsWin = false
		db.CurStage--
		db.KillMonster = 0
		db.DeadAvatar = 0 // 如果第一关死了在第二关 重复挑战怎么办
		g.GetPd().SetCurChallengeRoundCount(0)
		g.GetPd().SetCurChallengeScore(0)
		rsp.Scene = g.GetChallengeScene()
	}

	g.Send(cmd.RestartChallengePhaseScRsp, rsp)
}

// 忘却之庭战斗退出/结束

func LeaveChallengeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	curChallenge := g.GetPd().GetCurChallenge()
	if curChallenge == nil {
		return
	}
	if curChallenge.Status == spb.ChallengeStatus_CHALLENGE_DOING {
		g.Send(cmd.QuitBattleScNotify, nil) // 战斗没结束就退出是主动退出
	}
	g.ChallengeSettleNotify()
	g.Send(cmd.LeaveChallengeScRsp, &proto.LeaveChallengeScRsp{})

	g.EnterSceneByServerScNotify(g.GetPd().GetCurEntryId(), 0, 0, 0)
	// 设置战斗状态为空
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
	// 清空忘却之庭
	g.GetPd().NewCurChallenge()
}

// 忘却之庭世界战斗结算事件

func (g *GamePlayer) ChallengePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq, battleBin *model.BattleBackup) {
	db := g.GetPd().GetCurChallenge()
	if db == nil {
		return
	}
	conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
	if conf == nil {
		return
	}
	stt := req.GetStt()
	if stt != nil {
		g.GetPd().SetCurChallengeRoundCount(req.Stt.GetRoundCnt())  // 更新已使用回合数
		g.GetPd().SetCurChallengeScore(req.Stt.GetChallengeScore()) // 更新分数
	}
	switch req.EndStatus {
	case proto.BattleEndStatus_BATTLE_END_NONE:
		return
	case proto.BattleEndStatus_BATTLE_END_LOSE: // 战斗失败
		if !g.GetPd().ChallengeBattleEndLose() {
			if !db.IsBoos {
				g.ChallengeSettleNotify()
			}
			return
		}
	case proto.BattleEndStatus_BATTLE_END_QUIT: // 退出战斗
		return
	}
	if battleBin != nil && battleBin.Sce != nil {
		g.GetPd().AddCurChallengeKillMonster(uint32(len(battleBin.Sce.MonsterEntityIdList)))
	}
	// 场景上是否还有未处理敌人
	if g.GetPd().GetCurChallengeMonsterNum() > g.GetPd().GetCurChallengeKillMonster() {
		return // 还有就不更新状态，继续进行
	}
	// 更新状态
	g.GetPd().SetCurChallengeKillMonster(0) // 切换关卡，标记为0
	if db.IsBoos {
		// 结算
		g.GetPd().ChallengeSettle()
		g.ChallengeBossPhaseSettleNotify(req.Stt.GetBattleTargetInfo())
	}
	// 回合处理
	if g.GetPd().IsNextChallenge() {
		// 还没结束
		g.GetPd().AddChallengeCurStage(1)
		if !db.IsBoos {
			// 添加怪物
			g.ChallengeAddSceneGroupRefreshScNotify()
			// 添加角色
			g.ChallengeAddAvatarSceneGroupRefreshScNotify()
			// 更新新的队伍
			g.SyncLineupNotify(g.GetPd().GetBattleLineUpById(model.Challenge_2))
		}
	} else {
		// 发送战斗胜利通知
		if !db.IsBoos {
			// 结算
			g.GetPd().ChallengeSettle()
			g.ChallengeSettleNotify()
		}
		// 将战斗结果储存到数据库
		g.GetPd().UpdateChallengeList(conf.GroupID, db)
		// 更改状态
		g.GetPd().SetCurChallengeStatus(spb.ChallengeStatus_CHALLENGE_FINISH)
	}
}

func (g *GamePlayer) ChallengeSettleNotify() {
	cur := g.GetPd().GetCurChallenge()
	db := g.GetPd().GetChallengeInfoById(cur.GroupId, cur.ChallengeId)
	addItem := model.NewAddItem(nil)
	if !db.IsReward {
		g.GetPd().GetChallengeReward(addItem)
		db.IsReward = true
		g.AllPlayerSyncScNotify(addItem.AllSync)
	}

	notify := &proto.ChallengeSettleNotify{
		Star:           cur.Stars,                                   // 得分
		Reward:         &proto.ItemList{ItemList: addItem.ItemList}, // 奖励
		ChallengeId:    cur.ChallengeId,                             // 关卡id
		IsWin:          cur.IsWin,                                   // 是否赢
		ScoreTwo:       cur.ScoreTwo,                                // 二层挑战得分
		ChallengeScore: cur.ScoreOne,                                // 一层挑战得分
	}
	if cur.IsWin {
		finishSubMission := g.GetPd().ChallengeFinishCnt()
		g.InspectMission(finishSubMission...)
	}
	g.Send(cmd.ChallengeSettleNotify, notify)
}

func (g *GamePlayer) ChallengeBossPhaseSettleNotify(targeList map[uint32]*proto.BattleTargetList) {
	cur := g.GetPd().GetCurChallenge()
	db := g.GetPd().GetChallengeInfoById(cur.GroupId, cur.ChallengeId)
	notify := &proto.ChallengeBossPhaseSettleNotify{
		IsRemainingAction: true,
		Star:              cur.Stars,
		Phase:             cur.CurStage,
		ChallengeScore:    cur.ScoreOne,
		ScoreTwo:          cur.ScoreTwo,
		IsWin:             cur.IsWin,
		BattleTargetList:  targeList[1].BattleTargetList,
		ChallengeId:       cur.ChallengeId,
		IsReward:          !db.IsReward,
	}
	g.Send(cmd.ChallengeBossPhaseSettleNotify, notify)
}

func (g *GamePlayer) ChallengeAddAvatarSceneGroupRefreshScNotify() {
	curChallenge := g.GetPd().GetCurChallenge()
	mazeGroupID := g.GetPd().GetChallengesMazeGroupID()
	lineUp := g.GetPd().GetChallengesLineUp()
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
	pos, rot := g.GetPd().GetChallengesAnchor(foorMap.AnchorList)
	if pos == nil || rot == nil {
		return
	}

	notify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		RefreshEntity: g.GetPd().GetAddAvatarSceneEntityRefreshInfo(lineUp, pos, rot),
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
	curChallenge := g.GetPd().GetCurChallenge()
	mazeGroupID := g.GetPd().GetChallengesMazeGroupID()
	configList := g.GetPd().GetChallengesConfigList()
	eventIDList := g.GetPd().GetChallengesEventIDList()
	npcMonsterIDList := g.GetPd().GetChallengesNpcMonsterIDList()
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
		RefreshType:   proto.SceneGroupRefreshType_SCENE_GROUP_REFRESH_TYPE_LOADED,
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
			entityId := g.GetPd().GetNextGameObjectGuid()
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
			g.GetPd().AddEntity(mazeGroupID, &model.MonsterEntity{
				Entity: model.Entity{
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

func StartPartialChallengeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.StartPartialChallengeCsReq)

	rsp := &proto.StartPartialChallengeScRsp{
		Lineup:       nil,
		Scene:        nil,
		Retcode:      0,
		CurChallenge: nil,
	}
	g.Send(cmd.StartPartialChallengeScRsp, rsp)
}

func EnterChallengeNextPhaseCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.EnterChallengeNextPhaseScRsp{
		Scene: g.GetChallengeScene(),
	}

	g.Send(cmd.EnterChallengeNextPhaseScRsp, rsp)
}

func GetFriendChallengeLineupCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetFriendChallengeLineupScRsp{}

	g.Send(cmd.GetFriendChallengeLineupScRsp, rsp)
}

// 获取忘却之庭世界
func (g *GamePlayer) GetChallengeScene() *proto.SceneInfo {
	curChallenge := g.GetPd().GetCurChallenge()
	leaderEntityId := g.GetPd().GetNextGameObjectGuid()
	lineUp := g.GetPd().GetChallengesLineUp()
	mazeGroupID := g.GetPd().GetChallengesMazeGroupID()
	configList := g.GetPd().GetChallengesConfigList()
	npcMonsterIDList := g.GetPd().GetChallengesNpcMonsterIDList()
	eventIDList := g.GetPd().GetChallengesEventIDList()
	mapEntranceID := g.GetPd().GetChallengesMapEntranceID()
	mapEntrance := gdconf.GetMapEntranceById(mapEntranceID)
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, mazeGroupID)
	if foorMap == nil || lineUp == nil || len(npcMonsterIDList) != len(eventIDList) || len(eventIDList) != len(configList) {
		return nil
	}
	pos, rot := g.GetPd().GetChallengesAnchor(foorMap.AnchorList)
	if pos == nil || rot == nil {
		return nil
	}
	// 获取映射信息
	worldId := gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID
	if worldId == 100 {
		worldId = 401
	}
	scene := &proto.SceneInfo{
		ClientPosVersion:  0,
		PlaneId:           mapEntrance.PlaneID,
		FloorId:           mapEntrance.FloorID,
		LeaderEntityId:    leaderEntityId,
		WorldId:           worldId,
		EntryId:           mapEntranceID,
		GameModeType:      4,
		EntityGroupList:   make([]*proto.SceneEntityGroupInfo, 0),
		SceneBuffInfoList: make([]*proto.BuffInfo, 0),
		// GroupIdList:        make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		LightenSectionList: []uint32{0},
		EntityList:         make([]*proto.SceneEntityInfo, 0),
	}

	// 添加场景buff
	if curChallenge.MazeBuffId != 0 {
		scene.SceneBuffInfoList = append(scene.SceneBuffInfoList, &proto.BuffInfo{
			Count:     4294967295,
			LifeTime:  -1,
			BuffId:    curChallenge.MazeBuffId,
			AddTimeMs: uint64(time.Now().UnixMilli()),
			Level:     1,
		})
	}
	// 添加自选buff
	if g.GetPd().GetCurChallengeBuffId() != 0 {
		scene.SceneBuffInfoList = append(scene.SceneBuffInfoList, &proto.BuffInfo{
			Count:     4294967295,
			LifeTime:  -1,
			BuffId:    g.GetPd().GetCurChallengeBuffId(),
			AddTimeMs: uint64(time.Now().UnixMilli()),
			Level:     1,
		})
	}
	// 将进入场景的角色添加到实体列表里
	entityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    0,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	g.GetPd().GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
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
			entityId := g.GetPd().GetNextGameObjectGuid()
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
			g.GetPd().AddEntity(mazeGroupID, &model.MonsterEntity{
				Entity: model.Entity{
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
