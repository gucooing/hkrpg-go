// 记录战斗关键数据并储存，用于战斗结算

package player

import (
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

var BattleBackupLock sync.Mutex // 战斗列表互斥锁

type CurBattle struct {
	BattleBackup    map[uint32]*BattleBackup // 正在进行的战斗[战斗id]战斗细节
	RogueInfoOnline *RogueInfoOnline         // 模拟宇宙临时数据
	// AvatarBuff         map[uint32]*OnBuffMap // 角色在线buff
	ActivityInfoOnline *ActivityInfoOnline   // 角色试用在线数据
	MazeBuffList       map[uint32]*OnBuffMap // 所有buff
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
	AvatarExpReward uint32
	DisplayItemList []*Material
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

type OnBuffMap struct {
	AvatarId  uint32 // 角色
	BuffId    uint32 // buffid
	Level     uint32
	Count     uint32 // 使用次数
	LifeCount uint32 // 有效次数
	AddTime   uint64 // 添加时间
	LifeTime  uint32 // 有效时间
}

func (g *GamePlayer) GetOnBuffMap() map[uint32]*OnBuffMap {
	db := g.GetCurBattle()
	if db.MazeBuffList == nil {
		db.MazeBuffList = make(map[uint32]*OnBuffMap)
	}
	return db.MazeBuffList
}

func (g *GamePlayer) GetMazeBuffList() map[uint32]*OnBuffMap {
	db := g.GetCurBattle()
	if db.MazeBuffList == nil {
		db.MazeBuffList = make(map[uint32]*OnBuffMap)
	}
	return db.MazeBuffList
}

func (g *GamePlayer) AddOnLineAvatarBuff(avatarID, buffId uint32) {
	db := g.GetMazeBuffList()
	addTime := time.Now().Unix()
	db[buffId] = &OnBuffMap{
		AvatarId: avatarID,
		BuffId:   buffId,
		Count:    1,
		AddTime:  uint64(addTime),
		LifeTime: uint32(addTime + 15),
	}
}

func (g *GamePlayer) DelOnLineAvatarBuff(buffId uint32) {
	db := g.GetMazeBuffList()
	delete(db, buffId)
}

func (g *GamePlayer) sceneCastSkill(sce *SceneCastEntity, skill *gdconf.GoppMazeSkill, req *proto.SceneCastSkillCsReq) bool {
	battle := skill.TriggerBattle
	for _, actions := range skill.ActionsList {
		switch actions.Type {
		case constant.AddMazeBuff:
			g.AddOnLineAvatarBuff(sce.AvatarId, actions.Id)
		case constant.SetMonsterDie:
			for i := 0; i < len(sce.MonsterIdList); i++ {
				monsterId := sce.MonsterIdList[i]
				conf := gdconf.GetNPCMonsterId(monsterId)
				if conf == nil || conf.GetMonsterRank() < 3 {
					sce.MonsterIdList = append(sce.MonsterIdList[:i], sce.MonsterIdList[i+1:]...)
					i--
				}
			}
		case constant.AddTeamPlayerHP:
		case constant.AddTeamPlayerSp:
		case constant.SummonUnit:
		}
	}
	return battle
}

func (g *GamePlayer) DelMp(avatarId, castEntityId uint32) {
	isDelMp := false
	confAvatar := gdconf.GetAdventurePlayerByAvatarId(avatarId)
	if confAvatar == nil {
		return
	}
	for _, mazeSkillId := range confAvatar.MazeSkillIdList {
		confBuff := gdconf.GetAvatarMazeBuffById(mazeSkillId, 1)
		if confBuff == nil {
			continue
		}
		if confBuff.InBattleBindingKey != "" {
			isDelMp = true
		}
	}

	if isDelMp {
		g.DelLineUpMp(1)
		g.Send(cmd.SceneCastSkillMpUpdateScNotify, &proto.SceneCastSkillMpUpdateScNotify{
			CastEntityId: castEntityId,
			Mp:           g.GetLineUpMp(),
		})
		g.SyncLineupNotify(g.GetBattleLineUp())
	}
	return
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

type SceneCastEntity struct {
	IsAvatar            bool     // 是否有玩家参与
	MonsterEntityIdList []uint32 // 怪物实体id
	MonsterIdList       []uint32 // 怪物id
	PropEntityIdList    []uint32 // 物品实体id
	PropIdList          []uint32 // 物品id
	AvatarId            uint32   // 角色id
	AvatarEntityId      uint32   // 角色实体id
}

func (g *GamePlayer) GetMem(isMem []uint32, sce *SceneCastEntity) {
	for _, id := range isMem {
		entity := g.GetEntityById(id)
		if entity == nil {
			continue
		}
		switch entity.(type) {
		case *AvatarEntity:
			sce.IsAvatar = true
			sce.AvatarId = entity.(*AvatarEntity).AvatarId
			sce.AvatarEntityId = id
		case *MonsterEntity:
			if sce.MonsterEntityIdList == nil {
				sce.MonsterEntityIdList = make([]uint32, 0)
			}
			if sce.MonsterIdList == nil {
				sce.MonsterIdList = make([]uint32, 0)
			}
			sce.MonsterEntityIdList = append(sce.MonsterEntityIdList, id)
			sce.MonsterIdList = append(sce.MonsterIdList, entity.(*MonsterEntity).EventID)
		case *PropEntity:
			if sce.PropEntityIdList == nil {
				sce.PropEntityIdList = make([]uint32, 0)
			}
			if sce.PropIdList == nil {
				sce.PropIdList = make([]uint32, 0)
			}
			sce.PropEntityIdList = append(sce.PropEntityIdList, id)
			sce.PropIdList = append(sce.PropIdList, entity.(*PropEntity).PropId)
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
	itemConf := gdconf.GetItemConfigMap()
	for _, displayItem := range conf.DisplayItemList {
		if itemConf.Equipment[displayItem.ItemID] != nil {
			uniqueId := g.AddEquipment(displayItem.ItemID)
			allSync.EquipmentList = append(allSync.EquipmentList, uniqueId)
			itemList = append(itemList, g.GetEquipmentItem(uniqueId))
			continue
		}
		if itemConf.Relic[displayItem.ItemID] != nil {
			uniqueId := g.AddRelic(displayItem.ItemID)
			allSync.RelicList = append(allSync.RelicList, uniqueId)
			itemList = append(itemList, g.GetRelicItem(uniqueId))
			continue
		}
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
func (g *GamePlayer) GetSceneBattleInfo(eventList, stageIDList []uint32, avatarMap map[uint32]*BattleAvatar, worldLevel, stageId uint32) (*proto.SceneBattleInfo, *BattleBackup) {
	if (len(eventList) == 0 && len(stageIDList) == 0) ||
		len(avatarMap) == 0 {
		return nil, nil
	}
	// 记录此次战斗
	battleId := g.GetBattleIdGuid()
	battleBackup := &BattleBackup{
		BattleId:         battleId,
		BattleAvatarList: avatarMap,
		monsterEntity:    make([]uint32, 0),
	}
	var monsterWaveList []*proto.SceneMonsterWave
	if len(eventList) != 0 {
		monsterWaveList, stageId = g.GetSceneMonsterWave(eventList, worldLevel, battleBackup)
	}
	if len(stageIDList) != 0 {
		for id, stage := range stageIDList {
			bin := g.GetSceneMonsterWaveByStageID(stage, worldLevel, uint32(id+1), battleBackup)
			monsterWaveList = append(monsterWaveList, bin...)
			if id == 0 && stageId == 0 {
				stageId = stage
			}
		}
	}

	info := &proto.SceneBattleInfo{
		LogicRandomSeed:  gdconf.GetLoadingDesc(),
		BattleAvatarList: g.GetProtoBattleAvatar(avatarMap),
		BattleEvent:      make([]*proto.BattleEventBattleInfo, 0),
		BuffList:         g.GetBattleBuff(avatarMap),
		HEAMIJGFDMO:      nil,
		StageId:          stageId,
		FNLHAHFIGNC:      nil,
		HKOOBMMLGME:      nil,
		BattleTargetInfo: g.GetBattleTargetInfo(),
		MOJLNCNDIOB:      false,
		RoundsLimit:      g.GetRoundsLimit(),
		MonsterWaveList:  monsterWaveList,
		WorldLevel:       worldLevel,
		BOJHPNCAKOP:      0,
		BattleId:         battleId,
	}

	return info, battleBackup
}

func (g *GamePlayer) GetSceneMonsterWave(eventList []uint32, worldLevel uint32, battleBackup *BattleBackup) ([]*proto.SceneMonsterWave, uint32) {
	mWList := make([]*proto.SceneMonsterWave, 0)
	var stageID uint32 = 0
	for id, meid := range eventList {
		stage := gdconf.GetPlaneEventById(meid, worldLevel)
		if stage == nil {
			continue
		}
		bin := g.GetSceneMonsterWaveByStageID(stage.StageID, worldLevel, uint32(id+1), battleBackup)
		mWList = append(mWList, bin...)
		if id == 0 {
			stageID = stage.StageID // 阶段id
		}
	}
	return mWList, stageID
}

func (g *GamePlayer) GetSceneMonsterWaveByStageID(stageID, worldLevel, waveId uint32, battleBackup *BattleBackup) []*proto.SceneMonsterWave {
	mWList := make([]*proto.SceneMonsterWave, 0)
	stageConfig := gdconf.GetStageConfigById(stageID)
	if stageConfig == nil {
		logger.Warn("[UID:%v]get SceneMonsterWave error stageID:%v", g.Uid, stageID)
		return nil
	}
	for _, monsterListMap := range stageConfig.MonsterList {
		monsterWaveList := &proto.SceneMonsterWave{
			BattleStageId: stageID,
			BattleWaveId:  waveId,
			DropList:      make([]*proto.ItemList, 0),
			MonsterList:   make([]*proto.SceneMonster, 0),
			MonsterParam:  &proto.SceneMonsterWaveParam{},
		}
		for _, monsterId := range monsterListMap {
			sceneMonster := &proto.SceneMonster{
				MonsterId: monsterId,
			}
			monsterWaveList.MonsterList = append(monsterWaveList.MonsterList, sceneMonster)
			// 添加战利品
			monsterDrop := gdconf.GetMonsterDrop(monsterId, worldLevel)
			itemList := make([]*proto.Item, 0)
			if monsterDrop != nil {
				battleBackup.AvatarExpReward += monsterDrop.AvatarExpReward
				for _, item := range monsterDrop.DisplayItemList {
					var num uint32 = 4
					if item.ItemID == 2 {
						num = 50
					}
					itemList = append(itemList, &proto.Item{
						ItemId: item.ItemID,
						Num:    num,
					})
					battleBackup.DisplayItemList = append(battleBackup.DisplayItemList, &Material{
						Tid: item.ItemID,
						Num: num,
					})
				}
			}
			monsterWaveList.DropList = append(monsterWaveList.DropList, &proto.ItemList{
				ItemList: itemList,
			})
		}
		mWList = append(mWList, monsterWaveList)
	}
	return mWList
}

// 根据战斗情况添加buff
func (g *GamePlayer) GetBattleBuff(avatarMap map[uint32]*BattleAvatar) []*proto.BattleBuff {
	buffList := make([]*proto.BattleBuff, 0)
	mazeBufflist := g.GetMazeBuffList()
	// var targetIndex = 0
	for _, buff := range mazeBufflist {
		if buff.AvatarId != 0 { // add avatarBuff
			avatarInfo := avatarMap[buff.AvatarId]
			buffList = append(buffList, &proto.BattleBuff{
				Id:              buff.BuffId,
				Level:           1,
				OwnerIndex:      avatarInfo.Index,
				WaveFlag:        4294967295,
				TargetIndexList: []uint32{avatarInfo.Index},
				DynamicValues:   make(map[string]float32),
			})
			g.DelOnLineAvatarBuff(buff.BuffId)
		}
	}

	// 默认buff
	// buffList = append(buffList, &proto.BattleBuff{
	// 	Id:              1000113,
	// 	Level:           1,
	// 	OwnerIndex:      4294967295,
	// 	TargetIndexList: []uint32{0},
	// 	DynamicValues: map[string]float32{
	// 		"SkillIndex": 1,
	// 	},
	// })
	// 添加物品buff

	for index, buff := range mazeBufflist {
		if buff.Count < buff.LifeCount {
			buff.Count++
			buffList = append(buffList, &proto.BattleBuff{
				Id:         buff.BuffId,
				Level:      buff.Level,
				OwnerIndex: buff.LifeTime,
				WaveFlag:   4294967295,
			})
		} else {
			delete(mazeBufflist, index)
		}
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
