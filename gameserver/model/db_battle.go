// 记录战斗关键数据并储存，用于战斗结算

package model

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

type CurBattle struct {
	BattleBackup    map[uint32]*BattleBackup // 正在进行的战斗[战斗id]战斗细节
	RogueInfoOnline *RogueInfoOnline         // 模拟宇宙临时数据
	MazeBuffList    map[uint32]*OnBuffMap    // 所有buff
	FarmElementMap  map[uint32]uint32        // [id]world level // 虚影等级设置
}

type BattleBackup struct {
	IsBattle         bool                     // 是否开启战斗
	IsQuick          bool                     // 是否快速挑战
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
	IsFarmElement    bool                     // 是否虚影
	FarmElementID    uint32                   // 虚影Id
	// Skill
	// 奖励
	AvatarExpReward uint32
	AddItem         *AddItem
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
	return g.GetBattleBackup()[battleId]
}

func (g *PlayerData) AddBattleBackup(bb *BattleBackup) {
	g.GetBattleBackup()[bb.BattleId] = bb
}

func (g *PlayerData) DelBattleBackupById(battleId uint32) {
	delete(g.GetBattleBackup(), battleId)
}

func (g *PlayerData) GetFarmElementWorldLevel(stageId uint32) uint32 {
	db := g.GetCurBattle()
	if db.FarmElementMap == nil {
		db.FarmElementMap = make(map[uint32]uint32)
	}
	return db.FarmElementMap[stageId]
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

func (g *PlayerData) SetMonsterDie(eventID uint32) bool {
	switch eventID {
	case 20241170, 20134107: // 扑满特殊处理
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
		Challenge:  nil,
		Rain:       nil,
		QuestRogue: nil,
		TournRogue: nil,
		CurRogue:   nil,
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

type SceneCastEntity struct {
	IsAvatar            bool     // 是否有玩家参与
	MonsterEntityIdList []uint32 // 怪物实体id
	EvenIdList          []uint32 // 怪物id
	PropEntityIdList    []uint32 // 物品实体id
	PropIdList          []uint32 // 物品id
	AvatarId            uint32   // 角色id
	AvatarEntityId      uint32   // 角色实体id
}

func (g *PlayerData) GetMem(isMem []uint32, battleBackup *BattleBackup) {
	if battleBackup.Sce == nil {
		battleBackup.Sce = new(SceneCastEntity)
	}
	sce := battleBackup.Sce
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
			monster := entity.(*MonsterEntity)
			sce.EvenIdList = append(sce.EvenIdList, monster.EventID)
			sce.MonsterEntityIdList = append(sce.MonsterEntityIdList, monster.EntityId)
			if monster.PurposeType == "FarmElement" {
				battleBackup.IsFarmElement = true
				battleBackup.FarmElementID = monster.FarmElementID
			}
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

/*************奖励*************/

func (g *PlayerData) GetBattleDropData(mappingInfoID uint32, battleBin *BattleBackup) {
	conf := gdconf.GetMappingInfoById(mappingInfoID, battleBin.WorldLevel)
	if conf != nil {
		battleBin.AddItem = NewAddItem(battleBin.AddItem)
		itemConfMap := gdconf.GetItemConfig()
		for _, displayItem := range conf.DisplayItemList {
			itemConf := itemConfMap.Item[displayItem.ItemID]
			if itemConf != nil {
				switch itemConf.ItemSubType {
				case constant.ItemSubTypeRelicSetShowOnly:
					for _, id := range itemConf.CustomDataList {
						relicConf := gdconf.GetRelicBySetID(id, itemConf.Rarity)
						if relicConf != nil {
							battleBin.AddItem.PileItem = append(battleBin.AddItem.PileItem, &Material{
								Tid: relicConf.ID,
								Num: 1,
							})
						}
					}
					continue
				}
				itemNum := displayItem.ItemNum
				if displayItem.ItemID == Scoin {
					itemNum = 1500 + battleBin.WorldLevel*300
				}
				if itemNum == 0 {
					itemNum = 1 + battleBin.WorldLevel
				}
				battleBin.AddItem.PileItem = append(battleBin.AddItem.PileItem, &Material{
					Tid: displayItem.ItemID,
					Num: itemNum,
				})
				continue
			} else if itemConfMap.Equipment[displayItem.ItemID] != nil {
				battleBin.AddItem.PileItem = append(battleBin.AddItem.PileItem, &Material{
					Tid: displayItem.ItemID,
					Num: 1,
				})
			}
		}
	}
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
		StageId:          battleBackup.StageID,
		BattleTargetInfo: g.GetBattleTargetInfo(),
		RoundsLimit:      g.GetRoundsLimit(),
		MonsterWaveList:  monsterWaveList,
		WorldLevel:       battleBackup.WorldLevel,
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
	battleBackup.AddItem = NewAddItem(battleBackup.AddItem)
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
					battleBackup.AddItem.PileItem = append(battleBackup.AddItem.PileItem, &Material{
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
				info := &proto.BattleBuff{
					Id:              buff.BuffId,
					Level:           1,
					OwnerIndex:      avatarInfo.Index,
					WaveFlag:        4294967295,
					TargetIndexList: []uint32{avatarInfo.Index},
					DynamicValues:   make(map[string]float32),
				}
				buffList = append(buffList, info)
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
		if avatarInfo.IsCur {
			id := gdconf.GetAvatarDamage(avatarInfo.BaseAvatarId)
			if id != 0 {
				buffList = append(buffList, &proto.BattleBuff{
					Id:              id,
					Level:           1,
					OwnerIndex:      avatarInfo.Index,
					WaveFlag:        4294967295,
					TargetIndexList: []uint32{avatarInfo.Index},
					DynamicValues: map[string]float32{
						"SkillIndex": 1,
					},
				})
			}
		}
	}
	g.DelSummonUnitInfo()

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
	case spb.BattleType_Battle_QUSET_ROGUE:
		buffList = append(buffList, g.GetCurRogueBuff()...)
	}
	return buffList
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

func (g *PlayerData) GetBattleLoadScene(entryId uint32, pos, rot *proto.Vector, lineUp *spb.Line) *proto.SceneInfo {
	planeID, floorID, ok := gdconf.GetPFlaneID(entryId)
	if !ok {
		// TODO log
		return nil
	}
	leaderEntityId := g.GetNextGameObjectGuid()
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		WorldId:            gdconf.GetWorldId(planeID),
		LeaderEntityId:     leaderEntityId,
		FloorId:            floorID,
		GameModeType:       gdconf.GetPlaneType(planeID),
		PlaneId:            planeID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		FloorSavedData:     g.GetFloorSavedData(entryId),
		EntityBuffInfoList: make([]*proto.EntityBuffInfo, 0),
	}
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	g.GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)

	return scene
}
