package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func (g *PlayerData) GetChallenge() *spb.Challenge {
	db := g.GetBattle()
	if db.Challenge == nil {
		db.Challenge = &spb.Challenge{
			ChallengeGroupList: make(map[uint32]*spb.ChallengeGroupInfo),
			CurChallenge:       &spb.CurChallenge{},
		}
	}
	return db.Challenge
}

func (g *PlayerData) GetChallengeGroupList() map[uint32]*spb.ChallengeGroupInfo {
	db := g.GetChallenge()
	if db.ChallengeGroupList == nil {
		db.ChallengeGroupList = make(map[uint32]*spb.ChallengeGroupInfo)
	}
	return db.ChallengeGroupList
}

func (g *PlayerData) GetChallengeGroupInfoById(groupId uint32) *spb.ChallengeGroupInfo {
	db := g.GetChallengeGroupList()
	if db[groupId] == nil {
		db[groupId] = &spb.ChallengeGroupInfo{}
	}
	return db[groupId]
}

func (g *PlayerData) GetChallengeInfoById(groupId, challengeId uint32) *spb.ChallengeInfo {
	db := g.GetChallengeGroupInfoById(groupId)
	if db.ChallengeInfoList == nil {
		db.ChallengeInfoList = make(map[uint32]*spb.ChallengeInfo)
	}
	if db.ChallengeInfoList[challengeId] == nil {
		db.ChallengeInfoList[challengeId] = &spb.ChallengeInfo{
			ChallengeId: challengeId,
		}
	}
	return db.ChallengeInfoList[challengeId]
}

func (g *PlayerData) UpdateChallengeList(groupId uint32, curChallenge *spb.CurChallenge) {
	group := g.GetChallengeGroupInfoById(groupId)
	group.RecordId++
	group.MaxChallengeId = alg.MaxUin32(group.MaxChallengeId, curChallenge.ChallengeId)
	db := g.GetChallengeInfoById(groupId, curChallenge.ChallengeId)
	db.RecordId++
	if db.Stars < curChallenge.Stars ||
		db.ScoreOne+db.ScoreTwo < curChallenge.ScoreOne+curChallenge.ScoreTwo {
		newDb := &spb.ChallengeInfo{
			Stars:       curChallenge.Stars,
			ScoreOne:    curChallenge.ScoreOne,
			ScoreTwo:    curChallenge.ScoreTwo,
			ChallengeId: curChallenge.ChallengeId,
			IsReward:    false,
			RecordId:    db.RecordId,
			BuffOne:     curChallenge.BuffOne,
			BuffTwo:     curChallenge.BuffTwo,
			LineupList:  curChallenge.LineupList,
			Floor:       curChallenge.Floor,
		}

		group.ChallengeInfoList[curChallenge.ChallengeId] = newDb
	}
}

// 这玩意是用来清空当前忘却之庭战斗的
func (g *PlayerData) NewCurChallenge() {
	db := g.GetChallenge()
	db.CurChallenge = nil
}

func (g *PlayerData) SetCurChallenge(req *proto.StartChallengeCsReq) *spb.CurChallenge {
	db := g.GetChallenge()
	storyInfo := req.GetStageInfo()
	var buffOne uint32 = 0
	var buffTwe uint32 = 0
	var isBoos = false
	if storyInfo != nil {
		if storyInfo.StoryInfo == nil {
			isBoos = true
			buffOne = storyInfo.GetBossInfo().GetBuffOne()
			buffTwe = storyInfo.GetBossInfo().GetBuffTwo()
		} else {
			buffOne = storyInfo.GetStoryInfo().GetBuffOne()
			buffTwe = storyInfo.GetStoryInfo().GetBuffTwo()
		}
	}
	conf := gdconf.GetChallengeMazeConfigById(req.ChallengeId)
	db.CurChallenge = &spb.CurChallenge{
		ChallengeId: req.ChallengeId,
		StageNum:    conf.StageNum,
		CurStage:    1,
		Status:      spb.ChallengeStatus_CHALLENGE_DOING,
		RoundCount:  0,
		BuffOne:     buffOne,
		BuffTwo:     buffTwe,
		MazeBuffId:  conf.MazeBuffID,
		IsBoos:      isBoos,
		GroupId:     conf.GroupID,
		Floor:       conf.Floor,
		LineupList:  make([]*spb.ChallengeLineup, 0),
	}
	db.CurChallenge.LineupList = append(db.CurChallenge.LineupList,
		g.GetSpbChallengeLineup(g.GetBattleLineUpById(Challenge_1)))
	if req.SecondLineup != nil {
		db.CurChallenge.LineupList = append(db.CurChallenge.LineupList,
			g.GetSpbChallengeLineup(g.GetBattleLineUpById(Challenge_2)))
	}
	return db.CurChallenge
}

func (g *PlayerData) GetSpbChallengeLineup(line *spb.Line) *spb.ChallengeLineup {
	info := &spb.ChallengeLineup{
		AvatarList: make([]*spb.ChallengeAvatar, 0),
	}
	if line == nil {
		return info
	}
	for _, avatar := range line.AvatarIdList {
		db := g.GetAvatarBinById(avatar.AvatarId)
		if db == nil {
			continue
		}
		info.AvatarList = append(info.AvatarList, &spb.ChallengeAvatar{
			AvatarId: avatar.AvatarId,
			Level:    db.Level,
			Index:    avatar.Slot,
			Type:     spb.AvatarType(db.AvatarType),
		})
	}
	return info
}

func (g *PlayerData) GetCurChallenge() *spb.CurChallenge {
	db := g.GetChallenge()
	return db.CurChallenge
}

func (g *PlayerData) SetCurChallengeRoundCount(rc uint32) {
	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		db := g.GetCurChallenge()
		if db != nil {
			db.RoundCount += rc
		}
	}
}

func (g *PlayerData) SetCurChallengeScore(score uint32) {
	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE_Story:
		db := g.GetCurChallenge()
		if db != nil {
			switch db.CurStage {
			case 1:
				db.ScoreOne = score
			case 2:
				db.ScoreTwo = score
			}
		}
	}
}

func (g *PlayerData) IsNextChallenge() bool {
	db := g.GetCurChallenge()
	if db == nil {
		return false
	}
	if db.StageNum > db.CurStage {
		return true
	} else {
		return false
	}
}

func (g *PlayerData) AddChallengeCurStage(num uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.CurStage += num
}

func (g *PlayerData) SetCurChallengeStatus(status spb.ChallengeStatus) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.Status = status
}

func (g *PlayerData) AddCurChallengeKillMonster(num uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.KillMonster += num
}

func (g *PlayerData) SetCurChallengeKillMonster(num uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.KillMonster = num
}

func (g *PlayerData) GetCurChallengeKillMonster() uint32 {
	db := g.GetCurChallenge()
	if db == nil {
		return 0
	}
	return db.KillMonster
}

func (g *PlayerData) GetChallengesMazeGroupID() uint32 {
	curChallenge := g.GetCurChallenge()
	if curChallenge == nil {
		return 0
	}
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return 0
	}
	switch curChallenge.CurStage {
	case 1:
		return challengeMazeConfig.MazeGroupID1
	case 2:
		return challengeMazeConfig.MazeGroupID2
	}
	return 0
}

func (g *PlayerData) GetChallengesLineUp() *spb.Line {
	curChallenge := g.GetCurChallenge()
	if curChallenge == nil {
		return nil
	}
	switch curChallenge.CurStage {
	case 1:
		return g.GetBattleLineUpById(Challenge_1)
	case 2:
		return g.GetBattleLineUpById(Challenge_2)
	}
	return nil
}

func (g *PlayerData) GetChallengesConfigList() []uint32 {
	curChallenge := g.GetCurChallenge()
	if curChallenge == nil {
		return nil
	}
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return nil
	}
	switch curChallenge.CurStage {
	case 1:
		return challengeMazeConfig.ConfigList1
	case 2:
		return challengeMazeConfig.ConfigList2
	}
	return nil
}

func (g *PlayerData) GetCurChallengeMonsterNum() uint32 {
	conf := g.GetChallengesNpcMonsterIDList()
	if conf == nil {
		return 0
	}
	return uint32(len(conf))
}

func (g *PlayerData) GetChallengesNpcMonsterIDList() []uint32 {
	curChallenge := g.GetCurChallenge()
	if curChallenge == nil {
		return nil
	}
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return nil
	}
	switch curChallenge.CurStage {
	case 1:
		return challengeMazeConfig.NpcMonsterIDList1
	case 2:
		return challengeMazeConfig.NpcMonsterIDList2
	}
	return nil
}

func (g *PlayerData) GetChallengesEventIDList() []uint32 {
	curChallenge := g.GetCurChallenge()
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return nil
	}
	switch curChallenge.CurStage {
	case 1:
		return challengeMazeConfig.EventIDList1
	case 2:
		return challengeMazeConfig.EventIDList2
	}
	return nil
}

func (g *PlayerData) GetChallengesMapEntranceID() uint32 {
	curChallenge := g.GetCurChallenge()
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return 0
	}
	switch curChallenge.CurStage {
	case 1:
		return challengeMazeConfig.MapEntranceID
	case 2:
		return challengeMazeConfig.MapEntranceID2
	}
	return 0
}

func (g *PlayerData) GetCurChallengeBuffId() uint32 {
	curChallenge := g.GetCurChallenge()
	if curChallenge == nil {
		return 0
	}
	switch curChallenge.CurStage {
	case 1:
		return curChallenge.BuffOne
	case 2:
		return curChallenge.BuffTwo
	}
	return 0
}

func (g *PlayerData) GetChallengesAnchor(anchorList map[uint32]*gdconf.AnchorList) (pos, rot *proto.Vector) {
	if anchorList == nil {
		return nil, nil
	}
	for _, anchor := range anchorList {
		pos = &proto.Vector{
			Y: int32(anchor.PosY * 1000),
			X: int32(anchor.PosX * 1000),
			Z: int32(anchor.PosZ * 1000),
		}
		rot = &proto.Vector{
			Y: int32(anchor.RotY * 1000),
			X: int32(anchor.RotX * 1000),
			Z: int32(anchor.RotZ * 1000),
		}
		break
	}
	return
}

// 添加死亡角色数
func (g *PlayerData) AddChallengeDeadAvatar(deadNum uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.DeadAvatar += deadNum
}

// 关卡结束，开始结算
func (g *PlayerData) ChallengeSettle() {
	db := g.GetCurChallenge()
	conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
	if conf == nil {
		return
	}
	// 正式得分计算
	for _, tagId := range conf.ChallengeTargetID {
		tagConf := gdconf.GetChallengeTargetConfigById(tagId)
		switch tagConf.ChallengeTargetType {
		case "DEAD_AVATAR": // 角色存活得分
			if db.DeadAvatar <= tagConf.ChallengeTargetParam1 {
				db.Stars += 3
				// 添加奖励
			}
		case "ROUNDS_LEFT": // 剩余回合得分计算
			if conf.ChallengeCountDown-db.RoundCount > tagConf.ChallengeTargetParam1 {
				db.Stars += 2
				// 添加奖励
			}
		case "TOTAL_SCORE": // 活动得分计算得分
			if db.ScoreOne+db.ScoreTwo >= tagConf.ChallengeTargetParam1 {
				db.Stars += 2
			}
		}
	}
	if !g.IsNextChallenge() {
		db.IsWin = true
		// 额外得分计算
		if db.ScoreOne+db.ScoreTwo >= 30000 {
			db.Stars++
		}
		if db.IsBoos {
			db.Stars++
		}
	}
}

func GetChallengeStars(stars uint32) uint32 {
	total := 0
	for i := 0; i < 3; i++ {
		if stars&(1<<i) != 0 {
			total++
		}
	}
	return uint32(total)
}

func GetTakenRewards(takenStars uint64) uint32 {
	bitMask := takenStars
	var index uint32 = 0
	for bitMask > 0 {
		bitMask >>= 1
		index++
	}

	return index
}

func SetTakenReward(takenStars uint64, star uint32) uint64 {
	takenStars |= 1 << star
	return takenStars
}

// 忘却之庭战斗失败处理
func (g *PlayerData) ChallengeBattleEndLose() bool {
	db := g.GetCurChallenge()
	if db == nil {
		return false
	}
	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		g.SetCurChallengeStatus(spb.ChallengeStatus_CHALLENGE_UNKNOWN)
		return false
	case spb.BattleType_Battle_CHALLENGE_Story:
		if db.ScoreOne+db.ScoreTwo >= 30000 {
			return true
		}
	}
	return false
}

func (g *PlayerData) GetChallengeReward(addItem *AddItem) {
	db := g.GetCurChallenge()
	addItem = NewAddItem(addItem)
	conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
	if conf != nil && db.IsWin {
		pile := GetRewardData(conf.RewardID)
		addItem.PileItem = append(addItem.PileItem, pile...)
		g.AddItem(addItem)
	}
}

/***********************接口**********************/

func (g *PlayerData) GetCurChallengeBuff() []*proto.BattleBuff {
	db := g.GetCurChallenge()
	buffList := make([]*proto.BattleBuff, 0)
	// 关卡buff
	if db.MazeBuffId != 0 {
		buffList = append(buffList, &proto.BattleBuff{
			Id:         db.MazeBuffId,
			Level:      1,
			OwnerIndex: 4294967295,
			WaveFlag:   4294967295,
		})
	}
	// 自选buff
	buffId := g.GetCurChallengeBuffId()
	if buffId != 0 {
		buffList = append(buffList, &proto.BattleBuff{
			Id:              buffId,
			Level:           1,
			OwnerIndex:      0,
			TargetIndexList: []uint32{0},
			WaveFlag:        4294967295, // 失效时间
			DynamicValues:   make(map[string]float32),
		})
	}
	return buffList
}

func (g *PlayerData) GetChallengeInfo() *proto.CurChallenge {
	db := g.GetCurChallenge()
	if db == nil {
		return nil
	}
	var lineUpType proto.ExtraLineupType
	switch db.CurStage {
	case 1:
		lineUpType = proto.ExtraLineupType_LINEUP_CHALLENGE
	case 2:
		lineUpType = proto.ExtraLineupType_LINEUP_CHALLENGE_2
	}
	challengeInfo := &proto.CurChallenge{
		ChallengeId:     db.ChallengeId,                   // 挑战关卡
		Status:          proto.ChallengeStatus(db.Status), // 关卡状态
		ExtraLineupType: lineUpType,                       // 队伍type
		StageInfo:       g.GetCurChallengeStoryInfo(),     // 挑战buff
		RoundCount:      db.RoundCount,                    // 已使用回合数
		ScoreId:         db.ScoreOne,                      // 第一层得分
		ScoreTwo:        db.ScoreTwo,                      // 第二层得分
		KillMonsterList: make([]*proto.KillMonster, 0),
	}
	return challengeInfo
}

// 添加自选的关卡buff
func (g *PlayerData) GetCurChallengeStoryInfo() *proto.ChallengeStoryInfo {
	db := g.GetCurChallenge()
	if db == nil {
		return nil
	}
	if db.IsBoos {
		return &proto.ChallengeStoryInfo{
			CurBossBuffs: &proto.ChallengeBossBuffList{
				ChallengeBossConst: 1, // 这玩意不是1就不能进下一节点
				BuffList:           []uint32{db.BuffOne, db.BuffTwo},
			},
		}
	} else {
		return &proto.ChallengeStoryInfo{
			CurStoryBuffs: &proto.ChallengeStoryBuffList{
				BuffList: []uint32{db.BuffOne, db.BuffTwo},
			},
		}
	}
}

/****************************忘却之庭获取挑战信息(明明一模一样还分成三个proto*******************************/

func (g *PlayerData) GetChallengeGroupStatisticsChallengeStory(groupId uint32) *proto.GetChallengeGroupStatisticsScRsp_ChallengeStory {
	group := g.GetChallengeGroupInfoById(groupId)
	var db *spb.ChallengeInfo
	if group.ChallengeInfoList != nil {
		db = group.ChallengeInfoList[group.MaxChallengeId]
	}
	info := &proto.GetChallengeGroupStatisticsScRsp_ChallengeStory{
		ChallengeStory: &proto.ChallengeStoryStatistics{
			RecordId:       group.RecordId,
			StageTertinggi: g.GetChallengeStoryStageTertinggi(db),
		},
	}
	return info
}

func (g *PlayerData) GetChallengeStoryStageTertinggi(db *spb.ChallengeInfo) *proto.ChallengeStoryStageTertinggi {
	if db == nil {
		return nil
	}
	info := &proto.ChallengeStoryStageTertinggi{
		LineupList: g.GetChallengeLineupList(db.LineupList),
		// DKFHAHHJILF: 0,
		Level:   db.Floor,
		BuffTwo: db.BuffTwo,
		BuffOne: db.BuffOne,
		ScoreId: db.ScoreOne + db.ScoreTwo,
	}
	return info
}

func (g *PlayerData) GetChallengeGroupStatisticsChallengeDefault(groupId uint32) *proto.GetChallengeGroupStatisticsScRsp_ChallengeDefault {
	group := g.GetChallengeGroupInfoById(groupId)
	var db *spb.ChallengeInfo
	if group.ChallengeInfoList != nil {
		db = group.ChallengeInfoList[group.MaxChallengeId]
	}
	info := &proto.GetChallengeGroupStatisticsScRsp_ChallengeDefault{
		ChallengeDefault: &proto.ChallengeStatistics{
			StageTertinggi: g.GetChallengeStageTertinggi(db),
			RecordId:       group.RecordId,
		},
	}
	return info
}

func (g *PlayerData) GetChallengeStageTertinggi(db *spb.ChallengeInfo) *proto.ChallengeStageTertinggi {
	if db == nil {
		return nil
	}
	info := &proto.ChallengeStageTertinggi{
		LineupList: g.GetChallengeLineupList(db.LineupList),
		// DKFHAHHJILF: 0,
		Level:      db.Floor,
		RoundCount: db.RecordId,
	}
	return info
}

func (g *PlayerData) GetChallengeGroupStatisticsChallengeBoss(groupId uint32) *proto.GetChallengeGroupStatisticsScRsp_ChallengeBoss {
	group := g.GetChallengeGroupInfoById(groupId)
	var db *spb.ChallengeInfo
	if group.ChallengeInfoList != nil {
		db = group.ChallengeInfoList[group.MaxChallengeId]
	}
	info := &proto.GetChallengeGroupStatisticsScRsp_ChallengeBoss{
		ChallengeBoss: &proto.ChallengeBossStatistics{
			RecordId:       group.RecordId,
			StageTertinggi: g.GetChallengeBossStageTertinggi(db),
		},
	}
	return info
}

func (g *PlayerData) GetChallengeBossStageTertinggi(db *spb.ChallengeInfo) *proto.ChallengeBossStageTertinggi {
	if db == nil {
		return nil
	}
	info := &proto.ChallengeBossStageTertinggi{
		LineupList: g.GetChallengeLineupList(db.LineupList),
		// DKFHAHHJILF: 0,
		Level:   db.Floor,
		BuffTwo: db.BuffTwo,
		BuffOne: db.BuffOne,
		ScoreId: db.ScoreOne + db.ScoreTwo,
	}
	return info
}

func (g *PlayerData) GetChallengeLineupList(db []*spb.ChallengeLineup) []*proto.ChallengeLineupList {
	list := make([]*proto.ChallengeLineupList, 0)
	if db == nil {
		return list
	}
	for _, avatarList := range db {
		info := &proto.ChallengeLineupList{
			AvatarList: make([]*proto.ChallengeAvatarInfo, 0),
		}
		for _, avatar := range avatarList.AvatarList {
			info.AvatarList = append(info.AvatarList, &proto.ChallengeAvatarInfo{
				Id:         avatar.AvatarId,
				AvatarType: proto.AvatarType(avatar.Type),
				Index:      avatar.Index,
				Level:      avatar.Level,
			})
		}
		list = append(list, info)
	}

	return list
}
