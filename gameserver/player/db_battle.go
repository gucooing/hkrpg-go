// 记录战斗关键数据并储存，用于战斗结算

package player

import (
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

var BattleBackupLock sync.Mutex // 战斗列表互斥锁

type CurBattle struct {
	BattleBackup map[uint32]*BattleBackup // 正在进行的战斗[战斗id]战斗细节
}

type BattleBackup struct {
	BattleId         uint32                   // 战斗id
	BattleAvatarList map[uint32]*BattleAvatar // 参加战斗的角色
	monsterEntity    []uint32                 // 参战怪物实体id
	CocoonId         uint32                   // 关卡id
	WorldLevel       uint32                   // 关卡等级
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
	delete(g.GetBattleBackup(), battleId)
}

type BattleState struct {
	BattleType         spb.BattleType
	TrialActivityState *TrialActivityState
	BuffList           []uint32 // 进入战斗需要添加的buff
	AvatarBuffList     []uint32 // 角色buff
	RogueState         *RogueState
}
type TrialActivityState struct {
	AvatarDemoId  uint32
	NPCMonsterPos *spb.VectorBin
	NPCMonsterRot *spb.VectorBin
	PlaneID       uint32
	FloorID       uint32
	EntranceID    uint32

	NPCMonsterID uint32
	EventID      uint32
	GroupID      uint32
	ConfigID     uint32
}

type Battle struct {
	BattleId         uint32                // 战斗ID
	Wave             uint32                // 次数
	EventIDList      []uint32              // 怪物群实体id
	LogicRandomSeed  uint32                // 逻辑随机种子
	RoundsLimit      uint32                // 回合限制
	StaminaCost      uint32                // 扣除体力
	DisplayItemList  []*Material           // 奖励物品
	BuffList         []*proto.BattleBuff   // Buff列表
	BattleAvatarList []*proto.BattleAvatar // 战斗角色列表
}

type RogueState struct {
	BuffNum       uint32
	BuffList      map[uint32]*RogueBuff
	AvatarEntity  map[uint32]*AvatarEntity
	MonsterEntity map[uint32]*MonsterEntity
	Battle        map[uint32]*RogueBattle
}
type RogueBattle struct {
	monsterEntityMap []uint32 // 实体列表
}
type RogueBuff struct {
	Level     uint32
	AddTimeMs uint64
}

func (g *GamePlayer) GetBattleState() *BattleState {
	if g.OnlineData.BattleState == nil {
		g.OnlineData.BattleState = &BattleState{
			BattleType:         0,
			BuffList:           make([]uint32, 0),
			TrialActivityState: &TrialActivityState{},
			RogueState:         &RogueState{},
		}
	}
	return g.OnlineData.BattleState
}

func (g *GamePlayer) GetTrialActivityState() *TrialActivityState {
	if g.GetBattleState().TrialActivityState == nil {
		g.GetBattleState().TrialActivityState = &TrialActivityState{}
	}
	return g.GetBattleState().TrialActivityState
}

/************************************区分一下***************************************/

func (g *GamePlayer) NewBattle() *spb.Battle {
	return &spb.Battle{
		BattleType: 0,
		Rogue:      nil,
		Challenge:  nil,
	}
}

func (g *GamePlayer) GetBattle() *spb.Battle {
	db := g.GetBasicBin()
	if db.Battle == nil {
		db.Battle = g.NewBattle()
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

func (g *GamePlayer) SetCurChallenge(challengeId uint32, storyInfo *proto.StartChallengeStoryInfo) *spb.CurChallenge {
	db := g.GetChallenge()
	var buffOne uint32 = 0
	var buffTwe uint32 = 0
	if storyInfo != nil {
		buffOne = storyInfo.GetStoryBuffInfo().GetStoryBuffOne()
		buffTwe = storyInfo.GetStoryBuffInfo().GetStoryBuffTwo()
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
	}
	return db.CurChallenge
}

func (g *GamePlayer) GetCurChallenge() *spb.CurChallenge {
	db := g.GetChallenge()
	return db.CurChallenge
}

func (g *GamePlayer) SetCurChallengeRoundCount(rc uint32) {
	db := g.GetCurChallenge()
	if db != nil {
		db.RoundCount += rc
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

func (g *GamePlayer) GetChallengesAnchor(anchorList []*gdconf.AnchorList) (pos, rot *proto.Vector) {
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
	IsBattle bool     // 是否战斗
	EntityId []uint32 // 实体id
	MPid     []uint32 // 怪物/物品对应id
}

func (g *GamePlayer) GetMem(isMem []uint32) *MPEM {
	mpem := &MPEM{
		IsBattle: false,
		EntityId: make([]uint32, 0),
		MPid:     make([]uint32, 0),
	}
	for _, id := range isMem {
		entity := g.GetEntityById(id)
		switch entity.(type) {
		case *AvatarEntity:
		case *MonsterEntity:
			mpem.EntityId = append(mpem.EntityId, id)
			mpem.MPid = append(mpem.MPid, entity.(*MonsterEntity).EventID)
			mpem.IsBattle = true
		case *PropEntity:
			mpem.EntityId = append(mpem.EntityId, id)
			mpem.MPid = append(mpem.MPid, entity.(*PropEntity).PropId)
		case *NpcEntity:
		default:
			logger.Debug("[EntityId:%v]没有找到相关实体信息", id)
		}
	}
	return mpem
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
}

func (g *GamePlayer) CocoonBattle(cocoonId, worldLevel uint32) {
	cocoonConfig := gdconf.GetCocoonConfigById(cocoonId, worldLevel)
	if cocoonConfig == nil {
		return
	}
}

func (g *GamePlayer) GetDbRogue() *spb.Rogue {
	if g.GetBattle().Rogue == nil {
		g.GetBattle().Rogue = &spb.Rogue{
			RogueArea: make(map[uint32]*spb.RogueArea),
		}

		g.GetBattle().Rogue.RogueArea[100] = &spb.RogueArea{
			AreaId:          100,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS, // 教学关卡默认通关
		}
		g.GetBattle().Rogue.RogueArea[110] = &spb.RogueArea{
			AreaId:          110,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_UNLOCK, // 第一关卡默认解锁
		}
	}

	return g.GetBattle().Rogue
}

func (g *GamePlayer) GetCurDbRogue() *spb.CurRogue {
	rogue := g.GetDbRogue()
	if rogue.CurRogue == nil {
		rogue.CurRogue = new(spb.CurRogue)
	}
	return rogue.CurRogue
}

func (g *GamePlayer) GetCurDbRoom() *spb.RogueRoom {
	curRogue := g.GetCurDbRogue()
	return curRogue.RogueSceneMap[curRogue.CurSiteId]
}

func (g *GamePlayer) GetDbRoomBySiteId(siteId uint32) *spb.RogueRoom {
	curRogue := g.GetCurDbRogue()
	return curRogue.RogueSceneMap[siteId]
}

func (g *GamePlayer) GetDbRogueArea(areaId uint32) *spb.RogueArea {
	rogue := g.GetDbRogue()
	if rogue.RogueArea == nil {
		rogue.RogueArea = make(map[uint32]*spb.RogueArea)
	}
	if rogue.RogueArea[areaId] == nil {
		rogue.RogueArea[areaId] = &spb.RogueArea{
			AreaId:          areaId,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_LOCK,
		}
	}

	return rogue.RogueArea[areaId]
}

/************************************区分一下***************************************/

func (g *GamePlayer) NewRogueState() {
	g.GetBattleState().RogueState = &RogueState{
		BuffNum:       0,
		AvatarEntity:  make(map[uint32]*AvatarEntity),
		MonsterEntity: make(map[uint32]*MonsterEntity),
		Battle:        make(map[uint32]*RogueBattle),
		BuffList:      make(map[uint32]*RogueBuff),
	}
}

func (g *GamePlayer) GetRogue() *RogueState {
	if g.GetBattleState().RogueState == nil {
		g.GetBattleState().RogueState = &RogueState{
			AvatarEntity:  make(map[uint32]*AvatarEntity),
			MonsterEntity: make(map[uint32]*MonsterEntity),
			Battle:        make(map[uint32]*RogueBattle),
			BuffList:      make(map[uint32]*RogueBuff),
		}
	}

	return g.GetBattleState().RogueState
}

func (g *GamePlayer) GetRogueBattle() map[uint32]*RogueBattle {
	if g.GetRogue().Battle == nil {
		g.GetRogue().Battle = make(map[uint32]*RogueBattle)
	}
	return g.GetRogue().Battle
}

func (g *GamePlayer) GetRogueBuff() map[uint32]*RogueBuff {
	if g.GetRogue().BuffList == nil {
		g.GetRogue().BuffList = make(map[uint32]*RogueBuff)
	}
	return g.GetRogue().BuffList
}

func (g *GamePlayer) RogueAddBuff(buffId uint32) {
	rogue := g.GetRogue()
	if rogue.BuffList[buffId] == nil {
		rogue.BuffList[buffId] = &RogueBuff{
			Level:     1,
			AddTimeMs: uint64(time.Now().UnixNano() / 1e6),
		}
	}
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetSceneBattleInfo(mem []uint32, lineUp *spb.Line) (*proto.SceneBattleInfo, *BattleBackup) {
	if mem == nil || lineUp == nil {
		logger.Debug("[UID:%v]战斗获取失败", g.Uid)
		return nil, nil
	}
	bAList := make(map[uint32]*BattleAvatar, 0)
	for _, lp := range lineUp.AvatarIdList {
		bA := &BattleAvatar{
			AssistUid: 0,
			AvatarId:  lp.AvatarId,
			IsAssist:  false,
		}
		bAList[lp.AvatarId] = bA
	}
	battleId := g.GetBattleIdGuid()
	monsterWaveList, stageId := g.GetSceneMonsterWave(mem)
	battleInfo := &proto.SceneBattleInfo{
		LogicRandomSeed:     gdconf.GetLoadingDesc(),                 // 逻辑随机种子
		WorldLevel:          g.GetWorldLevel(),                       // 世界等级
		BattleId:            battleId,                                // 战斗Id
		BattleAvatarList:    g.GetProtoBattleAvatar(bAList),          // 战斗角色列表
		MonsterWaveList:     monsterWaveList,                         // 怪物列表
		StageId:             stageId,                                 // 起始战斗
		BattleTargetInfo:    g.GetBattleTargetInfo(),                 // 战斗目标
		EventBattleInfoList: make([]*proto.BattleEventBattleInfo, 0), // 战斗信息？？？
		RoundsLimit:         g.GetRoundsLimit(),                      // 回合限制
		BuffList:            g.GetBattleBuff(),                       // Buff列表
	}
	// 记录此次战斗
	battleBackup := &BattleBackup{
		BattleId:         battleId,
		BattleAvatarList: bAList,
		monsterEntity:    make([]uint32, 0),
	}

	return battleInfo, battleBackup
}

func (g *GamePlayer) GetCocoonBattleInfo(lineUp *spb.Line, req *proto.StartCocoonStageCsReq) (*proto.SceneBattleInfo, *BattleBackup) {
	if lineUp == nil {
		logger.Debug("[UID:%v]战斗获取失败", g.Uid)
		return nil, nil
	}
	var monsterWaveList []*proto.SceneMonsterWave
	bAList := make(map[uint32]*BattleAvatar, 0)
	for _, lp := range lineUp.AvatarIdList {
		bA := &BattleAvatar{
			AssistUid: 0,
			AvatarId:  lp.AvatarId,
			IsAssist:  false,
		}
		bAList[lp.AvatarId] = bA
	}
	cocoonConfig := gdconf.GetCocoonConfigById(req.CocoonId, req.WorldLevel)
	if cocoonConfig.DropList == nil {
		return nil, nil
	}
	// 添加怪物波列表
	stageID := cocoonConfig.StageID
	for _, stage := range cocoonConfig.StageIDList {
		bin := g.GetSceneMonsterWaveByStageID(stage)
		if bin == nil {
			continue
		}
		if stageID == 0 {
			stageID = stage
		}
		monsterWaveList = append(monsterWaveList, bin...)
	}
	battleId := g.GetBattleIdGuid()
	battleInfo := &proto.SceneBattleInfo{
		LogicRandomSeed:     gdconf.GetLoadingDesc(),                 // 逻辑随机种子
		WorldLevel:          req.GetWorldLevel(),                     // 关卡等级
		BattleId:            battleId,                                // 战斗Id
		StageId:             stageID,                                 // 起始战斗
		BattleAvatarList:    g.GetProtoBattleAvatar(bAList),          // 战斗角色列表
		MonsterWaveList:     monsterWaveList,                         // 怪物列表
		BattleTargetInfo:    g.GetBattleTargetInfo(),                 // 战斗目标
		EventBattleInfoList: make([]*proto.BattleEventBattleInfo, 0), // 战斗信息？？？
		RoundsLimit:         g.GetRoundsLimit(),                      // 回合限制
		BuffList:            g.GetBattleBuff(),                       // Buff列表
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

func (g *GamePlayer) GetSceneMonsterWave(mem []uint32) ([]*proto.SceneMonsterWave, uint32) {
	mWList := make([]*proto.SceneMonsterWave, 0)
	var stageID uint32 = 0
	for id, meid := range mem {
		stage := gdconf.GetPlaneEventById(meid, g.GetWorldLevel())
		if stage == nil {
			continue
		}
		bin := g.GetSceneMonsterWaveByStageID(stage.StageID)
		if bin == nil {
			continue
		}
		mWList = append(mWList, bin...)
		if id == 0 {
			stageID = stage.StageID // 阶段id
		}
	}
	return mWList, stageID
}

func (g *GamePlayer) GetSceneMonsterWaveByStageID(stageID uint32) []*proto.SceneMonsterWave {
	mWList := make([]*proto.SceneMonsterWave, 0)
	stageConfig := gdconf.GetStageConfigById(stageID)
	if stageConfig == nil {
		return nil
	}
	for _, monsterListMap := range stageConfig.MonsterList {
		monsterWaveList := &proto.SceneMonsterWave{
			StageId:     stageID,
			WaveId:      1,
			DropList:    make([]*proto.ItemList, 0),
			MonsterList: make([]*proto.SceneMonster, 0),
			WaveParam:   &proto.SceneMonsterWaveParam{},
		}
		for _, monsterList := range monsterListMap {
			sceneMonster := &proto.SceneMonster{
				MonsterId: monsterList,
			}
			monsterWaveList.MonsterList = append(monsterWaveList.MonsterList, sceneMonster)
		}
		mWList = append(mWList, monsterWaveList)
	}

	return mWList
}

// 根据战斗情况添加buff
func (g *GamePlayer) GetBattleBuff() []*proto.BattleBuff {
	buffList := make([]*proto.BattleBuff, 0)
	status := g.GetBattleStatus()
	switch status {
	case spb.BattleType_Battle_CHALLENGE:
		buffList = append(buffList, g.GetCurChallengeBuff()...)
	case spb.BattleType_Battle_CHALLENGE_Story:
		buffList = append(buffList, g.GetCurChallengeBuff()...)
	}
	return buffList
}

func (g *GamePlayer) GetCurChallengeBuff() []*proto.BattleBuff {
	db := g.GetCurChallenge()
	buffList := make([]*proto.BattleBuff, 0)
	// 关卡buff
	conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
	if conf.MazeBuffID != 0 {
		buffList = append(buffList, &proto.BattleBuff{
			Id:       conf.MazeBuffID,
			Level:    1,
			OwnerId:  4294967295,
			WaveFlag: 4294967295,
		})
	}
	// 自选buff
	buffId := g.GetCurChallengeBuffId()
	if buffId != 0 {
		buffList = append(buffList, &proto.BattleBuff{
			Id:              buffId,
			Level:           1,
			OwnerId:         0,
			TargetIndexList: []uint32{0},
			WaveFlag:        4294967295, // 失效时间
			DynamicValues:   make(map[string]float32),
		})
	}
	// TODO 添加角色buff
	return buffList
}

func (g *GamePlayer) GetChallengeInfo() *proto.ChallengeInfo {
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
	challengeInfo := &proto.ChallengeInfo{
		ChallengeId:     db.ChallengeId,                   // 挑战关卡
		Status:          proto.ChallengeStatus(db.Status), // 关卡状态
		ExtraLineupType: lineUpType,                       // 队伍type
		StoryInfo:       g.GetCurChallengeStoryInfo(),     // 挑战buff
		RoundCount:      db.RoundCount,                    // 已使用回合数
		Score:           db.ScoreOne + db.ScoreTwo,        // 第一层得分
		ScoreTwo:        db.ScoreTwo,                      // 第二层得分
	}
	return challengeInfo
}

// 记得添加自选的关卡buff
func (g *GamePlayer) GetCurChallengeStoryInfo() *proto.ChallengeStoryInfo {
	return nil
}

// 获取回合限制
func (g *GamePlayer) GetRoundsLimit() uint32 {
	status := g.GetBattleStatus()
	switch status {
	case spb.BattleType_Battle_CHALLENGE:
		db := g.GetCurChallenge()
		conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
		if conf == nil {
			return 0
		}
		return conf.ChallengeCountDown
	case spb.BattleType_Battle_CHALLENGE_Story:
		db := g.GetCurChallenge()
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
	if g.GetBattleStatus() != spb.BattleType_Battle_CHALLENGE_Story {
		return battleTargetInfoList
	}
	conf := gdconf.GetChallengeStoryMazeExtraById(db.ChallengeId)
	if conf == nil {
		return battleTargetInfoList
	}
	battleTargetInfoList[1] = &proto.BattleTargetList{
		BattleTargetList: []*proto.BattleTarget{{
			Id: 10001,
		}},
	}
	battleTargetList := make([]*proto.BattleTarget, 0)
	for _, id := range conf.BattleTargetID {
		battleTarget := &proto.BattleTarget{
			Id:       id,
			Progress: 0,
		}
		battleTargetList = append(battleTargetList, battleTarget)
	}
	battleTargetInfoList[5] = &proto.BattleTargetList{
		BattleTargetList: battleTargetList,
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
