// 记录战斗关键数据并储存，用于战斗结算

package model

import (
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	IsBattle         bool                     // 是否开启战斗
	BattleId         uint32                   // 战斗id
	BattleAvatarList map[uint32]*BattleAvatar // 参加战斗的角色
	CocoonId         uint32                   // 关卡id
	WorldLevel       uint32                   // 关卡等级
	StageID          uint32                   // cocoon
	StageIDList      []uint32                 // cocoon
	EventId          uint32                   // 任务用的
	AetherDivideId   uint32                   // 以太战线id
	AetherAvatarList []*AetherAvatar          // 以太战线战斗角色
	Sce              *SceneCastEntity         // 参与实体
	// Skill
	SummonUnitId uint32 // 领域
	// 奖励
	AvatarExpReward uint32
	DisplayItemList []*Material
}

type ActivityInfoOnline struct {
	StageId uint32 // 关卡id
}

func (g *PlayerData) NewCurBattle() *CurBattle {
	return &CurBattle{
		BattleBackup: make(map[uint32]*BattleBackup),
	}
}

func (g *PlayerData) GetCurBattle() *CurBattle {
	db := g.GetOnlineData()
	if db.CurBattle == nil {
		db.CurBattle = g.NewCurBattle()
	}
	return db.CurBattle
}

func (g *PlayerData) GetBattleBackup() map[uint32]*BattleBackup {
	db := g.GetCurBattle()
	if db.BattleBackup == nil {
		db.BattleBackup = make(map[uint32]*BattleBackup)
	}
	return db.BattleBackup
}

func (g *PlayerData) GetBattleBackupById(battleId uint32) *BattleBackup {
	BattleBackupLock.Lock()
	defer BattleBackupLock.Unlock()
	return g.GetBattleBackup()[battleId]
}

func (g *PlayerData) AddBattleBackup(bb *BattleBackup) {
	BattleBackupLock.Lock()
	defer BattleBackupLock.Unlock()
	g.GetBattleBackup()[bb.BattleId] = bb
}

func (g *PlayerData) DelBattleBackupById(battleId uint32) {
	BattleBackupLock.Lock()
	defer BattleBackupLock.Unlock()
	delete(g.GetBattleBackup(), battleId)
}

type OnBuffMap struct {
	AvatarId  uint32 // 角色
	BuffId    uint32 // buffid
	Level     uint32 // 等级
	Count     uint32 // 使用次数
	LifeCount uint32 // 有效次数
	AddTime   uint64 // 添加时间
	LifeTime  uint32 // 有效时间
}

func (g *PlayerData) GetOnBuffMap() map[uint32]*OnBuffMap {
	db := g.GetCurBattle()
	if db.MazeBuffList == nil {
		db.MazeBuffList = make(map[uint32]*OnBuffMap)
	}
	return db.MazeBuffList
}

func (g *PlayerData) GetMazeBuffList() map[uint32]*OnBuffMap {
	db := g.GetCurBattle()
	if db.MazeBuffList == nil {
		db.MazeBuffList = make(map[uint32]*OnBuffMap)
	}
	return db.MazeBuffList
}

func (g *PlayerData) AddOnLineAvatarBuff(avatarID, buffId uint32) {
	db := g.GetMazeBuffList()
	addTime := time.Now().Unix()
	db[buffId] = &OnBuffMap{
		AvatarId: avatarID,
		BuffId:   buffId,
		Count:    1,
		AddTime:  uint64(addTime),
		LifeTime: uint32(addTime + 20),
	}
}

func (g *PlayerData) DelOnLineAvatarBuff(buffId uint32) {
	db := g.GetMazeBuffList()
	delete(db, buffId)
}

func (g *PlayerData) SceneCastSkill(battleInfo *BattleBackup, skill *gdconf.GoppMazeSkill, req *proto.SceneCastSkillCsReq) {
	battleInfo.IsBattle = skill.TriggerBattle
	sce := battleInfo.Sce
	for _, actions := range skill.ActionsList {
		switch actions.Type {
		case constant.AddMazeBuff:
			g.AddOnLineAvatarBuff(sce.AvatarId, actions.Id)
		case constant.SetMonsterDie:
			for i := 0; i < len(sce.EvenIdList); i++ {
				monsterId := sce.EvenIdList[i]
				if g.setMonsterDie(monsterId) {
					sce.EvenIdList = append(sce.EvenIdList[:i], sce.EvenIdList[i+1:]...)
					i--
				}
			}
		case constant.AddTeamPlayerHP:
		case constant.AddTeamPlayerSp:
		case constant.SummonUnit:
			battleInfo.SummonUnitId = actions.Id
		}
	}
}

func (g *PlayerData) setMonsterDie(eventID uint32) bool {
	if eventID == 20134107 { // 扑满特殊处理
		return false
	}
	worldLevel := g.GetWorldLevel()
	isDie := true
	stage := gdconf.GetPlaneEventById(eventID, worldLevel)
	if stage == nil {
		return true
	}
	stageConfig := gdconf.GetStageConfigById(stage.StageID)
	if stageConfig == nil {
		return true
	}
	switch stageConfig.StageType {
	case "Challenge":
		return false
	}
	for _, monsterListMap := range stageConfig.MonsterList {
		for _, monsterId := range monsterListMap {
			conf := gdconf.GetNPCMonsterId(monsterId)
			if conf != nil && conf.GetMonsterRank() >= 3 {
				isDie = false
			}
		}
	}
	return isDie
}

func (g *PlayerData) DelMp(avatarId uint32) bool {
	isDelMp := false
	confAvatar := gdconf.GetAdventurePlayerByAvatarId(avatarId)
	if confAvatar == nil {
		return false
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
	return isDelMp
}

func NewBattle() *spb.Battle {
	return &spb.Battle{
		BattleType: 0,
		Rogue:      nil,
		Challenge:  nil,
		Rain:       nil,
	}
}

func (g *PlayerData) GetBattle() *spb.Battle {
	db := g.GetBasicBin()
	if db.Battle == nil {
		db.Battle = NewBattle()
	}
	return db.Battle
}

func (g *PlayerData) GetBattleStatus() spb.BattleType {
	db := g.GetBattle()
	return db.BattleType
}

func (g *PlayerData) SetBattleStatus(status spb.BattleType) {
	db := g.GetBattle()
	db.BattleType = status
}

/*************忘却之庭*************/

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
	storyInfo := req.GetPlayerInfo()
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

type SceneCastEntity struct {
	IsAvatar            bool     // 是否有玩家参与
	MonsterEntityIdList []uint32 // 怪物实体id
	EvenIdList          []uint32 // 怪物id
	PropEntityIdList    []uint32 // 物品实体id
	PropIdList          []uint32 // 物品id
	AvatarId            uint32   // 角色id
	AvatarEntityId      uint32   // 角色实体id
}

func (g *PlayerData) GetMem(isMem []uint32, sce *SceneCastEntity) {
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
			if sce.EvenIdList == nil {
				sce.EvenIdList = make([]uint32, 0)
			}
			sce.MonsterEntityIdList = append(sce.MonsterEntityIdList, id)
			sce.EvenIdList = append(sce.EvenIdList, entity.(*MonsterEntity).EventID)
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

/*************奖励*************/

func (g *PlayerData) GetBattleDropData(mappingInfoID uint32, addPileItem []*Material, worldLevel uint32, allSync *AllPlayerSync) []*proto.Item {
	conf := gdconf.GetMappingInfoById(mappingInfoID, worldLevel)
	if conf == nil {
		return nil
	}
	itemList := make([]*proto.Item, 0)
	itemConfMap := gdconf.GetItemConfigMap()
	for _, displayItem := range conf.DisplayItemList {
		itemConf := itemConfMap.Item[displayItem.ItemID]
		if itemConf == nil {
			continue
		}
		switch itemConf.ItemSubType {
		case constant.ItemSubTypeRelicSetShowOnly:
			for _, id := range itemConf.CustomDataList {
				relicConf := gdconf.GetRelicBySetID(id, itemConf.Rarity)
				if relicConf != nil {
					uniqueId := g.AddRelic(relicConf.ID, 0, nil)
					itemList = append(itemList, g.GetRelicItem(uniqueId))
					allSync.RelicList = append(allSync.RelicList, uniqueId)
				}
			}
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
	}
	return itemList
}

/****************************************************功能***************************************************/

// 场景战斗
func (g *PlayerData) GetSceneBattleInfo(battleBackup *BattleBackup) *proto.SceneBattleInfo {
	if ((battleBackup.Sce == nil || len(battleBackup.Sce.EvenIdList) == 0) && len(battleBackup.StageIDList) == 0) ||
		len(battleBackup.BattleAvatarList) == 0 {
		logger.Warn("异常战斗请求")
		return nil
	}
	// 记录此次战斗
	battleBackup.BattleId = g.GetBattleIdGuid()
	var monsterWaveList []*proto.SceneMonsterWave
	if battleBackup.Sce != nil && len(battleBackup.Sce.EvenIdList) != 0 {
		monsterWaveList, battleBackup.StageID = g.GetSceneMonsterWave(battleBackup.Sce.EvenIdList, battleBackup.WorldLevel, battleBackup)
	}
	if len(battleBackup.StageIDList) != 0 {
		for id, stage := range battleBackup.StageIDList {
			bin := g.GetSceneMonsterWaveByStageID(stage, battleBackup.WorldLevel, uint32(id+1), battleBackup)
			monsterWaveList = append(monsterWaveList, bin...)
			if id == 0 && battleBackup.StageID == 0 {
				battleBackup.StageID = stage
			}
		}
	}

	info := &proto.SceneBattleInfo{
		LogicRandomSeed:  gdconf.GetLoadingDesc(),
		BattleAvatarList: g.GetProtoBattleAvatar(battleBackup.BattleAvatarList),
		BattleEvent:      make([]*proto.BattleEventBattleInfo, 0),
		BuffList:         g.GetBattleBuff(battleBackup.BattleAvatarList),
		HEAMIJGFDMO:      nil,
		StageId:          battleBackup.StageID,
		FNLHAHFIGNC:      nil,
		HKOOBMMLGME:      nil,
		BattleTargetInfo: g.GetBattleTargetInfo(),
		MOJLNCNDIOB:      false,
		RoundsLimit:      g.GetRoundsLimit(),
		MonsterWaveList:  monsterWaveList,
		WorldLevel:       battleBackup.WorldLevel,
		BOJHPNCAKOP:      0,
		BattleId:         battleBackup.BattleId,
	}

	return info
}

func (g *PlayerData) GetSceneMonsterWave(monsterIdList []uint32, worldLevel uint32, battleBackup *BattleBackup) ([]*proto.SceneMonsterWave, uint32) {
	mWList := make([]*proto.SceneMonsterWave, 0)
	var stageID uint32 = 0
	for id, meid := range monsterIdList {
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

func (g *PlayerData) GetSceneMonsterWaveByStageID(stageID, worldLevel, waveId uint32, battleBackup *BattleBackup) []*proto.SceneMonsterWave {
	mWList := make([]*proto.SceneMonsterWave, 0)
	stageConfig := gdconf.GetStageConfigById(stageID)
	if stageConfig == nil {
		logger.Warn("[UID:%v]get SceneMonsterWave error stageID:%v", g.GetBasicBin().Uid, stageID)
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
func (g *PlayerData) GetBattleBuff(avatarMap map[uint32]*BattleAvatar) []*proto.BattleBuff {
	buffList := make([]*proto.BattleBuff, 0)
	mazeBufflist := g.GetMazeBuffList()
	summonUnitInfo := g.GetSummonUnitInfo()
	// var targetIndex = 0
	for _, buff := range mazeBufflist {
		if buff.AvatarId != 0 { // add avatarBuff
			avatarInfo := avatarMap[buff.AvatarId]
			if avatarInfo != nil {
				buffList = append(buffList, &proto.BattleBuff{
					Id:              buff.BuffId,
					Level:           1,
					OwnerIndex:      avatarInfo.Index,
					WaveFlag:        4294967295,
					TargetIndexList: []uint32{avatarInfo.Index},
					DynamicValues:   make(map[string]float32),
				})
			}
			g.DelOnLineAvatarBuff(buff.BuffId)
			continue
		}
	}
	// add SummonUnitBuff
	var waveFlagMap = make(map[uint32]uint32)
	for _, avatarInfo := range avatarMap {
		if avatarInfo.AvatarId == summonUnitInfo.AvatarId {
			for _, buff := range summonUnitInfo.BuffList {
				waveFlagMap[buff.BuffId]++
				buffList = append(buffList, &proto.BattleBuff{
					Id:         buff.BuffId,
					Level:      1,
					OwnerIndex: avatarInfo.Index,
					WaveFlag:   waveFlagMap[buff.BuffId],
				})
			}
		}
	}
	g.DelSummonUnitInfo()

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
		PlayerInfo:      g.GetCurChallengeStoryInfo(),     // 挑战buff
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

// 获取回合限制
func (g *PlayerData) GetRoundsLimit() uint32 {
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
func (g *PlayerData) GetBattleTargetInfo() map[uint32]*proto.BattleTargetList {
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

func (g *PlayerData) GetChallengeReward(allSync *AllPlayerSync) *proto.ItemList {
	itemList := &proto.ItemList{
		ItemList: make([]*proto.Item, 0),
	}
	db := g.GetCurChallenge()
	conf := gdconf.GetChallengeMazeConfigById(db.ChallengeId)
	if conf == nil {
		return itemList
	}
	if db.IsWin {
		pile, item := GetRewardData(conf.RewardID)
		itemList.ItemList = item
		g.AddItem(pile, allSync)
	}

	return itemList
}
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
		LineupList:  g.GetChallengeLineupList(db.LineupList),
		DKFHAHHJILF: 0,
		Level:       db.Floor,
		BuffTwo:     db.BuffTwo,
		BuffOne:     db.BuffOne,
		ScoreId:     db.ScoreOne + db.ScoreTwo,
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
