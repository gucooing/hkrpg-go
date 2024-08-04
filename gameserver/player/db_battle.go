// 记录战斗关键数据并储存，用于战斗结算

package player

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

var BattleBackupLock sync.Mutex // 战斗列表互斥锁

type CurBattle struct {
	BattleBackup       map[uint32]*BattleBackup     // 正在进行的战斗[战斗id]战斗细节
	RogueInfoOnline    *RogueInfoOnline             // 模拟宇宙临时数据
	AvatarBuff         map[uint32]*OnLineAvatarBuff // 角色在线buff
	ActivityInfoOnline *ActivityInfoOnline          // 角色试用在线数据
}

type BattleBackup struct {
	BattleId           uint32                   // 战斗id
	BattleAvatarList   map[uint32]*BattleAvatar // 参加战斗的角色
	monsterEntity      []uint32                 // 参战怪物实体id
	CocoonId           uint32                   // 关卡id
	WorldLevel         uint32                   // 关卡等级
	EventId            uint32                   // 任务用的
	AttackedByEntityId uint32                   // 发起攻击的实体id
	// 奖励
}

func (g *GamePlayer) NewCurBattle() *CurBattle {
	return &CurBattle{
		BattleBackup: make(map[uint32]*BattleBackup),
	}
}

func (g *GamePlayer) GetCurBattle() *CurBattle {
	db := g.GetOnlineData()
	if db.CurBattle == nil {
		db.CurBattle = g.NewCurBattle()
	}
	return db.CurBattle
}

func (g *GamePlayer) GetBattleBackup() map[uint32]*BattleBackup {
	db := g.GetCurBattle()
	if db.BattleBackup == nil {
		db.BattleBackup = make(map[uint32]*BattleBackup)
	}
	return db.BattleBackup
}

func (g *GamePlayer) GetBattleBackupById(battleId uint32) *BattleBackup {
	BattleBackupLock.Lock()
	defer BattleBackupLock.Unlock()
	return g.GetBattleBackup()[battleId]
}

func (g *GamePlayer) AddBattleBackup(bb *BattleBackup) {
	BattleBackupLock.Lock()
	defer BattleBackupLock.Unlock()
	g.GetBattleBackup()[bb.BattleId] = bb
}

func (g *GamePlayer) DelBattleBackupById(battleId uint32) {
	BattleBackupLock.Lock()
	defer BattleBackupLock.Unlock()
	battle := g.GetBattleBackup()[battleId]
	for _, entityId := range battle.monsterEntity {
		me := g.GetMonsterEntityById(entityId)
		if me != nil {
			g.UpKillMonsterSubMission(me)
		}
	}
	delete(g.GetBattleBackup(), battleId)
}

type OnLineAvatarBuff struct {
	AvatarId uint32 // 角色
	BuffId   uint32 // buffid
	Count    uint32 // 使用次数
	AddTime  uint64 // 添加时间
	LifeTime uint32 // 有效时间
}

func (g *GamePlayer) NewOnLineAvatarBuff() {
	db := g.GetCurBattle()
	db.AvatarBuff = make(map[uint32]*OnLineAvatarBuff)
}

func (g *GamePlayer) GetOnLineAvatarBuff() map[uint32]*OnLineAvatarBuff {
	db := g.GetCurBattle()
	if db.AvatarBuff == nil {
		db.AvatarBuff = make(map[uint32]*OnLineAvatarBuff)
	}
	return db.AvatarBuff
}

func (g *GamePlayer) GetOnLineAvatarBuffById(id uint32) *OnLineAvatarBuff {
	db := g.GetOnLineAvatarBuff()
	return db[id]
}

func (g *GamePlayer) HandleAvatarSkill(entityId uint32) (bool, bool) {
	avatar := g.GetAvatarEntity(entityId)
	isBattle := true
	isDelMp := false
	if avatar == nil {
		return true, false
	}
	confAvatar := gdconf.GetAdventurePlayerByAvatarId(avatar.AvatarId)
	if confAvatar == nil {
		return true, false
	}
	for _, mazeSkillId := range confAvatar.MazeSkillIdList {
		confBuff := gdconf.GetAvatarMazeBuffById(mazeSkillId, 1)
		if confBuff == nil {
			continue
		}
		switch confBuff.UseType {
		case "AddBattleBuff":
			g.AddOnLineAvatarBuff(avatar.AvatarId, confBuff.ID)
		}
		if strings.Contains(confBuff.BuffEffect, "MazeBuffEffect") {
			// isBattle = false
			summonId := strconv.FormatUint(uint64(mazeSkillId), 10)[:4] + strconv.FormatUint(uint64(mazeSkillId), 10)[5:]
			g.Send(cmd.SceneGroupRefreshScNotify, &proto.SceneGroupRefreshScNotify{
				GroupRefreshList: g.GetAddBuffSceneEntityRefreshInfo(entityId, alg.S2U32(summonId), g.GetRotPb(), g.GetPosPb()),
			})
		}
		if confBuff.InBattleBindingKey != "" {
			isDelMp = true
		}
	}

	return isBattle, isDelMp
}

func (g *GamePlayer) AddOnLineAvatarBuff(avatarID, buffId uint32) {
	db := g.GetOnLineAvatarBuff()
	db[avatarID] = &OnLineAvatarBuff{
		AvatarId: avatarID,
		BuffId:   buffId,
		Count:    1,
		AddTime:  uint64(time.Now().Unix()),
		LifeTime: 15,
	}
}

func (g *GamePlayer) DelOnLineAvatarBuff(avatarID, buffId uint32) {
	db := g.GetOnLineAvatarBuff()
	delete(db, avatarID)
}

func NewBattle() *spb.Battle {
	return &spb.Battle{
		BattleType: 0,
		Rogue:      nil,
		Challenge:  nil,
		Rain:       nil,
	}
}

func (g *GamePlayer) GetBattle() *spb.Battle {
	db := g.GetBasicBin()
	if db.Battle == nil {
		db.Battle = NewBattle()
	}
	return db.Battle
}

func (g *GamePlayer) GetBattleStatus() spb.BattleType {
	db := g.GetBattle()
	return db.BattleType
}

func (g *GamePlayer) SetBattleStatus(status spb.BattleType) {
	db := g.GetBattle()
	db.BattleType = status
}

/*************忘却之庭*************/

func (g *GamePlayer) GetChallenge() *spb.Challenge {
	db := g.GetBattle()
	if db.Challenge == nil {
		db.Challenge = &spb.Challenge{
			ChallengeList:       make(map[uint32]*spb.ChallengeList),
			ChallengeRewardList: make(map[uint64]uint32),
			CurChallenge:        &spb.CurChallenge{},
		}
	}
	return db.Challenge
}

func (g *GamePlayer) GetChallengeList() map[uint32]*spb.ChallengeList {
	db := g.GetChallenge()
	if db.ChallengeList == nil {
		db.ChallengeList = make(map[uint32]*spb.ChallengeList)
	}
	return db.ChallengeList
}

func (g *GamePlayer) GetChallengeRewardList() map[uint64]uint32 {
	db := g.GetChallenge()
	if db.ChallengeRewardList == nil {
		db.ChallengeRewardList = make(map[uint64]uint32)
	}
	return db.ChallengeRewardList
}

func (g *GamePlayer) GetChallengeById(id uint32) *spb.ChallengeList {
	db := g.GetChallengeList()
	if db[id] == nil {
		db[id] = &spb.ChallengeList{}
	}
	return db[id]
}

func (g *GamePlayer) UpdateChallengeList(curChallenge *spb.CurChallenge) {
	db := g.GetChallengeById(curChallenge.ChallengeId)
	db.Stars = alg.MaxUin32(db.Stars, curChallenge.Stars)
	db.ScoreOne = alg.MaxUin32(db.ScoreOne, curChallenge.ScoreOne)
	db.ScoreTwo = alg.MaxUin32(db.ScoreTwo, curChallenge.ScoreTwo)
}

// 这玩意是用来清空当前忘却之庭战斗的
func (g *GamePlayer) NewCurChallenge() {
	db := g.GetChallenge()
	db.CurChallenge = nil
}

func (g *GamePlayer) SetCurChallenge(challengeId uint32, storyInfo *proto.ChallengeBuffInfo) *spb.CurChallenge {
	db := g.GetChallenge()
	var buffOne uint32 = 0
	var buffTwe uint32 = 0
	var isBoos = false
	if storyInfo != nil {
		if storyInfo.StoryBuffInfo == nil {
			isBoos = true
			buffOne = storyInfo.GetBossBuffInfo().GetBuffOne()
			buffTwe = storyInfo.GetBossBuffInfo().GetBuffTwo()
		} else {
			buffOne = storyInfo.GetStoryBuffInfo().GetBuffOne()
			buffTwe = storyInfo.GetStoryBuffInfo().GetBuffTwo()
		}
	}
	conf := gdconf.GetChallengeMazeConfigById(challengeId)
	db.CurChallenge = &spb.CurChallenge{
		ChallengeId: challengeId,
		StageNum:    conf.StageNum,
		CurStage:    1,
		Status:      spb.ChallengeStatus_CHALLENGE_DOING,
		RoundCount:  0,
		BuffOne:     buffOne,
		BuffTwo:     buffTwe,
		MazeBuffId:  conf.MazeBuffID,
		IsBoos:      isBoos,
	}
	return db.CurChallenge
}

func (g *GamePlayer) GetCurChallenge() *spb.CurChallenge {
	db := g.GetChallenge()
	return db.CurChallenge
}

func (g *GamePlayer) SetCurChallengeRoundCount(rc uint32) {
	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		db := g.GetCurChallenge()
		if db != nil {
			db.RoundCount += rc
		}
	}
}

func (g *GamePlayer) SetCurChallengeScore(score uint32) {
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

func (g *GamePlayer) IsNextChallenge() bool {
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

func (g *GamePlayer) AddChallengeCurStage(num uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.CurStage += num
}

func (g *GamePlayer) SetCurChallengeStatus(status spb.ChallengeStatus) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.Status = status
}

func (g *GamePlayer) AddCurChallengeKillMonster(num uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.KillMonster += num
}

func (g *GamePlayer) SetCurChallengeKillMonster(num uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.KillMonster = num
}

func (g *GamePlayer) GetCurChallengeKillMonster() uint32 {
	db := g.GetCurChallenge()
	if db == nil {
		return 0
	}
	return db.KillMonster
}

func (g *GamePlayer) GetChallengesMazeGroupID() uint32 {
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

func (g *GamePlayer) GetChallengesLineUp() *spb.Line {
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

func (g *GamePlayer) GetChallengesConfigList() []uint32 {
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

func (g *GamePlayer) GetCurChallengeMonsterNum() uint32 {
	conf := g.GetChallengesNpcMonsterIDList()
	if conf == nil {
		return 0
	}
	return uint32(len(conf))
}

func (g *GamePlayer) GetChallengesNpcMonsterIDList() []uint32 {
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

func (g *GamePlayer) GetChallengesEventIDList() []uint32 {
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

func (g *GamePlayer) GetCurChallengeBuffId() uint32 {
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

func (g *GamePlayer) GetChallengesAnchor(anchorList map[uint32]*gdconf.AnchorList) (pos, rot *proto.Vector) {
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
func (g *GamePlayer) AddChallengeDeadAvatar(deadNum uint32) {
	db := g.GetCurChallenge()
	if db == nil {
		return
	}
	db.DeadAvatar += deadNum
}

type MPEM struct {
	IsAvatar        bool     // 是否有玩家参与
	MonsterEntityId []uint32 // 怪物实体id
	MonsterId       []uint32 // 怪物id
	PropEntityId    []uint32 // 物品实体id
	PropId          []uint32 // 怪物id
	AvatarId        uint32   // 角色id
	AvatarEntityId  uint32   // 角色实体id
}

func (g *GamePlayer) GetMem(isMem []uint32, mpem *MPEM) {
	for _, id := range isMem {
		entity := g.GetEntityById(id)
		if entity == nil {
			continue
		}
		switch entity.(type) {
		case *AvatarEntity:
			mpem.IsAvatar = true
			mpem.AvatarId = entity.(*AvatarEntity).AvatarId
			mpem.AvatarEntityId = id
		case *MonsterEntity:
			if mpem.MonsterEntityId == nil {
				mpem.MonsterEntityId = make([]uint32, 0)
			}
			if mpem.MonsterId == nil {
				mpem.MonsterId = make([]uint32, 0)
			}
			mpem.MonsterEntityId = append(mpem.MonsterEntityId, id)
			mpem.MonsterId = append(mpem.MonsterId, entity.(*MonsterEntity).EventID)
		case *PropEntity:
			if mpem.PropEntityId == nil {
				mpem.PropEntityId = make([]uint32, 0)
			}
			if mpem.PropId == nil {
				mpem.PropId = make([]uint32, 0)
			}
			mpem.PropEntityId = append(mpem.PropEntityId, id)
			mpem.PropId = append(mpem.PropId, entity.(*PropEntity).PropId)
		case *NpcEntity:
		default:
			logger.Debug("[EntityId:%v]没有找到相关实体信息", id)
		}
	}
}

// 关卡结束，开始结算
func (g *GamePlayer) ChallengeSettle() {
	db := g.GetCurChallenge()
	conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
	if conf == nil {
		return
	}
	db.IsWin = true
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
	// 额外得分计算
	if db.ScoreOne+db.ScoreTwo >= 30000 {
		db.Stars++
	}
	if db.IsBoos {
		db.Stars++
	}
}

// 忘却之庭战斗失败处理
func (g *GamePlayer) ChallengeBattleEndLose() bool {
	db := g.GetCurChallenge()
	if db == nil {
		return false
	}
	switch g.GetBattleStatus() {
	case spb.BattleType_Battle_CHALLENGE:
		g.SetCurChallengeStatus(spb.ChallengeStatus_CHALLENGE_UNKNOWN)
		g.ChallengeSettleNotify()
		return false
	case spb.BattleType_Battle_CHALLENGE_Story:
		if db.ScoreOne+db.ScoreTwo >= 30000 {
			return true
		}
	}
	return false
}

/*************奖励*************/

func (g *GamePlayer) getBattleDropData(mappingInfoID uint32, allSync *AllPlayerSync, addPileItem []*Material, worldLevel uint32) []*proto.Item {
	conf := gdconf.GetMappingInfoById(mappingInfoID, worldLevel)
	if conf == nil {
		return nil
	}
	itemList := make([]*proto.Item, 0)
	for _, displayItem := range conf.DisplayItemList {
		itemNum := displayItem.ItemNum
		if displayItem.ItemID == Scoin {
			itemNum = 1500 + worldLevel*300
		}
		if itemNum == 0 {
			itemNum = 1 + worldLevel
		}
		addPileItem = append(addPileItem, &Material{
			Tid: displayItem.ItemID,
			Num: itemNum,
		})
		itemList = append(itemList, &proto.Item{
			ItemId: displayItem.ItemID,
			Num:    itemNum,
		})
		allSync.MaterialList = append(allSync.MaterialList, displayItem.ItemID)
	}
	return itemList
}

/****************************************************功能***************************************************/

// 场景战斗
func (g *GamePlayer) GetSceneBattleInfo(mem []uint32, lineUp *spb.Line) (*proto.SceneBattleInfo, *BattleBackup) {
	if mem == nil || lineUp == nil {
		logger.Debug("[UID:%v]战斗获取失败", g.Uid)
		return nil, nil
	}
	battleId := g.GetBattleIdGuid()
	bAList := make(map[uint32]*BattleAvatar, 0)
	for id, lp := range lineUp.AvatarIdList {
		if lp.AvatarId == 0 {
			continue
		}
		bA := &BattleAvatar{
			AssistUid:  0,
			AvatarId:   lp.AvatarId,
			AvatarType: lp.LineAvatarType,
		}
		bAList[id] = bA
	}
	monsterWaveList, stageId, trialAvatarList := g.GetSceneMonsterWave(mem)
	if trialAvatarList != nil && len(trialAvatarList) > 0 {
		bAList = make(map[uint32]*BattleAvatar)
		for id, avatarId := range trialAvatarList {
			if avatarId == 0 {
				continue
			}
			bA := &BattleAvatar{
				AssistUid:  0,
				AvatarId:   avatarId,
				AvatarType: spb.LineAvatarType_LineAvatarType_TRIAL,
			}
			bAList[uint32(id)] = bA
		}
	}
	battleAvatarList, buffList := g.GetProtoBattleAvatar(bAList)
	battleInfo := &proto.SceneBattleInfo{
		LogicRandomSeed:  gdconf.GetLoadingDesc(),                 // 逻辑随机种子
		WorldLevel:       g.GetWorldLevel(),                       // 世界等级
		BattleId:         battleId,                                // 战斗Id
		BattleAvatarList: battleAvatarList,                        // 战斗角色列表
		MonsterWaveList:  monsterWaveList,                         // 怪物列表
		StageId:          stageId,                                 // 起始战斗
		BattleTargetInfo: g.GetBattleTargetInfo(),                 // 战斗目标
		BattleEvent:      make([]*proto.BattleEventBattleInfo, 0), // 战斗信息？？？
		RoundsLimit:      g.GetRoundsLimit(),                      // 回合限制
		BuffList:         g.GetBattleBuff(buffList),               // Buff列表
	}
	// 记录此次战斗
	battleBackup := &BattleBackup{
		BattleId:         battleId,
		BattleAvatarList: bAList,
		monsterEntity:    make([]uint32, 0),
	}

	return battleInfo, battleBackup
}

// 副本战斗
func (g *GamePlayer) GetCocoonBattleInfo(lineUp *spb.Line, req *proto.StartCocoonStageCsReq) (*proto.SceneBattleInfo, *BattleBackup) {
	if lineUp == nil {
		logger.Debug("[UID:%v]战斗获取失败", g.Uid)
		return nil, nil
	}
	var monsterWaveList []*proto.SceneMonsterWave
	bAList := make(map[uint32]*BattleAvatar, 0)
	for _, lp := range lineUp.AvatarIdList {
		bA := &BattleAvatar{
			AssistUid:  0,
			AvatarId:   lp.AvatarId,
			AvatarType: lp.LineAvatarType,
		}
		bAList[lp.AvatarId] = bA
	}
	cocoonConfig := gdconf.GetCocoonConfigById(req.CocoonId, req.WorldLevel)
	if cocoonConfig == nil {
		logger.Warn("No Cocoon like this can be found,cocoonId:%v,worldLevel:%v", req.CocoonId, req.WorldLevel)
		return nil, nil
	}
	// 添加怪物波列表
	stageID := cocoonConfig.StageID
	for _, stage := range cocoonConfig.StageIDList {
		_, bin := g.GetSceneMonsterWaveByStageID(stage)
		monsterWaveList = append(monsterWaveList, bin...)
		if stageID == 0 {
			stageID = stage
		}
	}
	battleId := g.GetBattleIdGuid()
	battleAvatarList, buffList := g.GetProtoBattleAvatar(bAList)
	battleInfo := &proto.SceneBattleInfo{
		LogicRandomSeed:  gdconf.GetLoadingDesc(),                 // 逻辑随机种子
		WorldLevel:       req.GetWorldLevel(),                     // 关卡等级
		BattleId:         battleId,                                // 战斗Id
		StageId:          stageID,                                 // 起始战斗
		BattleAvatarList: battleAvatarList,                        // 战斗角色列表
		MonsterWaveList:  monsterWaveList,                         // 怪物列表
		BattleTargetInfo: g.GetBattleTargetInfo(),                 // 战斗目标
		BattleEvent:      make([]*proto.BattleEventBattleInfo, 0), // 战斗信息？？？
		RoundsLimit:      g.GetRoundsLimit(),                      // 回合限制
		BuffList:         g.GetBattleBuff(buffList),               // Buff列表
	}
	// 记录此次战斗
	battleBackup := &BattleBackup{
		BattleId:         battleId,
		BattleAvatarList: bAList,
		monsterEntity:    make([]uint32, 0),
		CocoonId:         req.CocoonId,
		WorldLevel:       req.WorldLevel,
	}

	return battleInfo, battleBackup
}

func (g *GamePlayer) GetSceneMonsterWave(mem []uint32) ([]*proto.SceneMonsterWave, uint32, []uint32) {
	mWList := make([]*proto.SceneMonsterWave, 0)
	var stageID uint32 = 0
	var trialAvatarList []uint32
	for id, meid := range mem {
		stage := gdconf.GetPlaneEventById(meid, g.GetWorldLevel())
		if stage == nil {
			continue
		}
		list, bin := g.GetSceneMonsterWaveByStageID(stage.StageID)
		mWList = append(mWList, bin...)
		if id == 0 {
			stageID = stage.StageID // 阶段id
		}
		if list != nil && len(list) > 0 {
			trialAvatarList = list
		}
	}
	return mWList, stageID, trialAvatarList
}

func (g *GamePlayer) GetSceneMonsterWaveByStageID(stageID uint32) ([]uint32, []*proto.SceneMonsterWave) {
	mWList := make([]*proto.SceneMonsterWave, 0)
	stageConfig := gdconf.GetStageConfigById(stageID)
	if stageConfig == nil {
		return nil, nil
	}
	for _, monsterListMap := range stageConfig.MonsterList {
		monsterWaveList := &proto.SceneMonsterWave{
			BattleStageId: stageID,
			BattleWaveId:  1,
			DropList:      make([]*proto.ItemList, 0),
			MonsterList:   make([]*proto.SceneMonster, 0),
			MonsterParam:  &proto.SceneMonsterWaveParam{},
		}
		for _, monsterId := range monsterListMap {
			sceneMonster := &proto.SceneMonster{
				MonsterId: monsterId,
			}
			monsterWaveList.MonsterList = append(monsterWaveList.MonsterList, sceneMonster)
			monsterWaveList.DropList = append(monsterWaveList.DropList, &proto.ItemList{})
		}
		mWList = append(mWList, monsterWaveList)
	}
	return stageConfig.TrialAvatarList, mWList
}

// 根据战斗情况添加buff
func (g *GamePlayer) GetBattleBuff(buffList []*proto.BattleBuff) []*proto.BattleBuff {
	if buffList == nil {
		buffList = make([]*proto.BattleBuff, 0)
	}
	status := g.GetBattleStatus()
	switch status {
	case spb.BattleType_Battle_CHALLENGE:
		buffList = append(buffList, g.GetCurChallengeBuff()...)
	case spb.BattleType_Battle_CHALLENGE_Story:
		buffList = append(buffList, g.GetCurChallengeBuff()...)
	case spb.BattleType_Battle_ROGUE:
		buffList = append(buffList, g.GetCurRogueBuff()...)
	}
	return buffList
}

func (g *GamePlayer) GetCurChallengeBuff() []*proto.BattleBuff {
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
	// TODO 添加角色buff
	return buffList
}

func (g *GamePlayer) GetChallengeInfo() *proto.CurChallenge {
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
		PlayerInfo:      g.GetCurChallengeStoryInfo(),     // 挑战buff
		RoundCount:      db.RoundCount,                    // 已使用回合数
		ScoreId:         db.ScoreOne,                      // 第一层得分
		ScoreTwo:        db.ScoreTwo,                      // 第二层得分
	}
	return challengeInfo
}

// 添加自选的关卡buff
func (g *GamePlayer) GetCurChallengeStoryInfo() *proto.ChallengeStoryInfo {
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

// 获取回合限制
func (g *GamePlayer) GetRoundsLimit() uint32 {
	status := g.GetBattleStatus()
	switch status {
	case spb.BattleType_Battle_CHALLENGE:
		db := g.GetCurChallenge()
		if db == nil {
			return 0
		}
		conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
		if conf == nil {
			return 0
		}
		return conf.ChallengeCountDown
	case spb.BattleType_Battle_CHALLENGE_Story:
		db := g.GetCurChallenge()
		if db == nil {
			return 0
		}
		conf := gdconf.GetChallengeStoryMazeExtraById(db.ChallengeId)
		if conf == nil {
			return 0
		}
		return conf.TurnLimit
	}
	return 0
}

// 添加战斗目标
func (g *GamePlayer) GetBattleTargetInfo() map[uint32]*proto.BattleTargetList {
	battleTargetInfoList := make(map[uint32]*proto.BattleTargetList)
	db := g.GetCurChallenge()
	battleTargetInfoList[1] = new(proto.BattleTargetList)
	battleTargetInfoList[2] = new(proto.BattleTargetList)
	battleTargetInfoList[3] = new(proto.BattleTargetList)
	battleTargetInfoList[4] = new(proto.BattleTargetList)
	battleTargetInfoList[5] = new(proto.BattleTargetList)

	if g.GetBattleStatus() != spb.BattleType_Battle_CHALLENGE_Story {
		return battleTargetInfoList
	}
	battleTargetList1 := make([]*proto.BattleTarget, 0)
	if db.IsBoos {
		for _, id := range []uint32{90004, 90005} {
			battleTargetList1 = append(battleTargetList1, &proto.BattleTarget{
				Id: id,
			})
		}
	} else {
		conf := gdconf.GetChallengeStoryMazeExtraById(db.ChallengeId)
		if conf == nil {
			return battleTargetInfoList
		}
		battleTargetList1 = append(battleTargetList1, &proto.BattleTarget{
			Id: 10001,
		})
		battleTargetList5 := make([]*proto.BattleTarget, 0)
		for _, id := range conf.BattleTargetID {
			battleTarget := &proto.BattleTarget{
				Id:       id,
				Progress: 0,
			}
			battleTargetList5 = append(battleTargetList5, battleTarget)
		}
		battleTargetInfoList[5] = &proto.BattleTargetList{
			BattleTargetList: battleTargetList5,
		}
	}
	battleTargetInfoList[1] = &proto.BattleTargetList{
		BattleTargetList: battleTargetList1,
	}
	return battleTargetInfoList
}

// todo 忘却之庭奖励获取
func (g *GamePlayer) GetChallengeReward() *proto.ItemList {
	itemList := &proto.ItemList{
		ItemList: make([]*proto.Item, 0),
	}

	return itemList
}
