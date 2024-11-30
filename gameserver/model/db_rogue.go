package model

import (
	"math/rand"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	QuestRogue      = 101 // 模拟宇宙
	RogueDlc        = 200 // 模拟宇宙：寰宇蝗灾
	RogueNous       = 201 // 模拟宇宙：黄金与机械
	QuestRogueTourn = 301 // 模拟宇宙：差分宇宙
)

// Default Probability
const (
	RogueBuffType      = 900  // 各属性
	AddRogueBuffType   = 1900 // 属性增加概率
	RogueBuffRarityOne = 6000 // 白
	RogueBuffRarityTwo = 3000 // 蓝
	AddRogueBuffRarity = 1000 // 品质增加概率
)

type RogueInfoOnline struct { // 模拟宇宙临时数据
	RogueBuffByType    map[uint32]*RogueBuffByType
	RogueBuffRarityOne int32 // 白的概率
	RogueBuffRarityTwo int32 // 蓝的概率
}

type RogueBuffByType struct {
	Weight          int32                                           // 权重
	RogueBuffRarity map[constant.RogueBuffCategory]*RogueBuffRarity // 稀有度
}

type RogueBuffRarity struct {
	Rarity   constant.RogueBuffCategory // 稀有度
	BuffList []uint32                   // buff列表
}

func (g *PlayerData) GetCurRogue() *spb.CurRogue {
	db := g.GetBattle()
	return db.CurRogue
}

func (g *PlayerData) GetQueuePositionNum() uint32 {
	db := g.GetCurRogue()
	if db == nil {
		return 0
	}
	defer func() {
		db.QueuePosition++
	}()
	return db.QueuePosition
}

func (g *PlayerData) GetRogueActionByQueuePosition(queuePosition uint32) *spb.RogueAction {
	curRogue := g.GetCurRogue()
	if curRogue == nil {
		return nil
	}
	if curRogue.Action == nil {
		curRogue.Action = map[uint32]*spb.RogueAction{}
	}
	return curRogue.Action[queuePosition]
}

func (g *PlayerData) DelRogueActionByQueuePosition(queuePosition uint32) {
	curRogue := g.GetCurRogue()
	if curRogue == nil {
		return
	}
	if curRogue.Action == nil {
		curRogue.Action = map[uint32]*spb.RogueAction{}
	}
	delete(curRogue.Action, queuePosition)
}

func (g *PlayerData) AddRogueActionBonusSelect(bonusIdList []uint32) {
	curRogue := g.GetCurRogue()
	if curRogue == nil {
		return
	}
	if curRogue.Action == nil {
		curRogue.Action = map[uint32]*spb.RogueAction{}
	}
	bonusIdMap := make(map[uint32]bool)
	for _, bonusId := range bonusIdList {
		bonusIdMap[bonusId] = true
	}
	if len(bonusIdMap) == 0 {
		return
	}
	queuePosition := g.GetQueuePositionNum()
	info := &spb.RogueAction{
		Action: &spb.RogueAction_BonusSelect{BonusSelect: &spb.BonusSelect{BonusIdMap: bonusIdMap}},
	}
	curRogue.Action[queuePosition] = info
}

func (g *PlayerData) AddRogueActionBuffSelect(buffIdList []uint32) {
	curRogue := g.GetCurRogue()
	if curRogue == nil {
		return
	}
	if curRogue.Action == nil {
		curRogue.Action = map[uint32]*spb.RogueAction{}
	}
	buffMap := make(map[uint32]*spb.RogueBuff)
	for _, buffId := range buffIdList {
		buffMap[buffId] = &spb.RogueBuff{
			BuffId:    buffId,
			BuffLevel: 1,
		}
	}
	if len(buffMap) == 0 {
		return
	}
	queuePosition := g.GetQueuePositionNum()
	info := &spb.RogueAction{
		Action: &spb.RogueAction_BuffSelect{BuffSelect: &spb.BuffSelect{
			SourceHintId:     7,
			RollBuffCost:     nil,
			CanRoll:          true,
			RollBuffCount:    1,
			RollBuffMaxCount: 1,
			SourceCurCount:   1,
			BuffMap:          buffMap,
		}},
	}
	curRogue.Action[queuePosition] = info
}

func (g *PlayerData) GetRogueBuffList() map[uint32]*spb.RogueBuff {
	db := g.GetCurRogue()
	if db == nil {
		return nil
	}
	if db.BuffList == nil {
		db.BuffList = make(map[uint32]*spb.RogueBuff)
	}
	return db.BuffList
}

func (g *PlayerData) GetRogueBuffById(id uint32) *spb.RogueBuff {
	db := g.GetRogueBuffList()
	return db[id]
}

func (g *PlayerData) AddRogueBuff(buffMap map[uint32]*spb.RogueBuff) {
	db := g.GetRogueBuffList()
	for buffId, info := range buffMap {
		if gdconf.GetBuffByIdAndLevel(buffId, info.BuffLevel) != nil {
			db[buffId] = &spb.RogueBuff{
				BuffId:    buffId,
				BuffLevel: info.BuffLevel,
				AddTime:   uint64(time.Now().UnixNano() / 1e6),
			}
		}
	}
}

/**************************************************Buff获取概率计算*******************************************/

func (g *PlayerData) GetRogueInfoOnline() *RogueInfoOnline {
	db := g.GetCurBattle()
	if db.RogueInfoOnline == nil {
		db.RogueInfoOnline = &RogueInfoOnline{}
	}
	return db.RogueInfoOnline
}

func (g *PlayerData) NewGetRogueBuffByType() {
	db := g.GetRogueInfoOnline()
	db.RogueBuffRarityOne = RogueBuffRarityOne
	db.RogueBuffRarityTwo = RogueBuffRarityTwo
	rogueBuffByTypeList := make(map[uint32]*RogueBuffByType, 0)
	buffTypeList := gdconf.GetRogueBuffByType()
	if buffTypeList != nil {
		for typeId, rogueBuffByType := range buffTypeList {
			if typeId == 100 { // 过滤基础
				continue
			}
			rogueBuffRarityList := make(map[constant.RogueBuffCategory]*RogueBuffRarity)
			for rarityId, buffListConf := range rogueBuffByType {
				buffList := make([]uint32, 0)
				for _, buff := range buffListConf {
					// 此处加个判断特殊祝福就行了
					conf := gdconf.GetBuffByIdAndLevel(buff, 1)
					if conf.ActivityModuleID != 0 { // && conf.ActivityModuleID != g.GetCurRogue().RogueActivityModuleID {
						continue
					}
					buffList = append(buffList, buff)
				}
				rogueBuffRarityList[rarityId] = &RogueBuffRarity{
					Rarity:   rarityId,
					BuffList: buffList,
				}
			}
			rogueBuffByTypeList[typeId] = &RogueBuffByType{
				Weight:          RogueBuffType,
				RogueBuffRarity: rogueBuffRarityList,
			}
		}
	}

	db.RogueBuffByType = rogueBuffByTypeList
}

func (g *PlayerData) GetRogueBuffByType() map[uint32]*RogueBuffByType {
	db := g.GetRogueInfoOnline()
	if db.RogueBuffByType == nil {
		g.NewGetRogueBuffByType()
	}
	return db.RogueBuffByType
}

func (g *PlayerData) GetRogueBuff() uint32 {
	db := g.GetRogueInfoOnline()
	rogueBuffByTypeList := db.RogueBuffByType
	var totalWeight int32 = 0
	for id, rogueBuffByType := range rogueBuffByTypeList {
		if rogueBuffByType.RogueBuffRarity == nil || len(rogueBuffByType.RogueBuffRarity) == 0 {
			continue
		}
		if id == 0 {
			rogueBuffByType.Weight += AddRogueBuffType
		}
		totalWeight += rogueBuffByType.Weight
	}
	if totalWeight == 0 {
		return 600000
	}
	randomWeight := rand.Int31n(totalWeight)
	for _, rogueBuffByType := range rogueBuffByTypeList {
		if rogueBuffByType.RogueBuffRarity == nil || len(rogueBuffByType.RogueBuffRarity) == 0 {
			continue
		}
		if randomWeight <= rogueBuffByType.Weight {
			// 已选定命途属性
			var rarityTotalWeight int32 = 0
			for _, rogueBuffRarity := range rogueBuffByType.RogueBuffRarity {
				var weight int32 = 0
				switch rogueBuffRarity.Rarity {
				case constant.RogueBuffCategoryNone:
					weight = db.RogueBuffRarityOne
				case constant.RogueBuffCategoryCommon:
					weight = db.RogueBuffRarityTwo
				default:
					continue
				}
				rarityTotalWeight += weight
			}
			if rarityTotalWeight == 0 {
				return 600000
			}
			rarityRandomWeight := rand.Int31n(rarityTotalWeight)
			for _, rogueBuffRarity := range rogueBuffByType.RogueBuffRarity {
				if rogueBuffRarity.BuffList == nil || len(rogueBuffRarity.BuffList) == 0 {
					continue
				}
				var weight int32 = 0
				switch rogueBuffRarity.Rarity {
				case constant.RogueBuffCategoryNone:
					weight = db.RogueBuffRarityOne
				case constant.RogueBuffCategoryCommon:
					weight = db.RogueBuffRarityTwo
				default:
					continue
				}
				if rarityRandomWeight <= weight {
					// 已选定稀有属性
					idIndex := rand.Intn(len(rogueBuffRarity.BuffList))
					return rogueBuffRarity.BuffList[idIndex]
				}
				randomWeight -= weight
			}
		}
		randomWeight -= rogueBuffByType.Weight
	}
	return 600000
}

/****************************************************功能***************************************************/

func (g *PlayerData) GetGameAeonInfo() (info *proto.GameAeonInfo) {
	info = &proto.GameAeonInfo{
		IsUnlocked:             true,
		UnlockedAeonEnhanceNum: 3,
		GameAeonId:             0,
	}
	rogue := g.GetCurRogue()
	if rogue == nil {
		return
	}
	info.GameAeonId = rogue.AeonId
	return
}

func (g *PlayerData) GetRogueMap() *proto.RogueMapInfo {
	rogue := g.GetCurRogue()
	if rogue == nil {
		return nil
	}
	questRogue := rogue.GetQuestRogue()
	roomMap := &proto.RogueMapInfo{
		MapId:     questRogue.RogueMapId,
		AreaId:    rogue.CurAreaId,
		CurSiteId: questRogue.CurSiteId, // 当前id
		CurRoomId: g.GetCurQuestRogueRoomId(),
		RoomList:  make([]*proto.RogueRoom, 0),
	}
	for id, rogueScene := range questRogue.RogueRoomMap {
		roomList := &proto.RogueRoom{
			SiteId:    id,
			RoomId:    rogueScene.RoomId,
			CurStatus: proto.RogueRoomStatus(rogueScene.RoomStatus),
		}
		roomMap.RoomList = append(roomMap.RoomList, roomList)
	}

	return roomMap
}

func (g *PlayerData) GetRogueLineupInfo() *proto.RogueLineupInfo {
	info := &proto.RogueLineupInfo{
		BaseAvatarIdList: make([]uint32, 0),
		ReviveInfo:       nil,
	}

	lineup := g.GetBattleLineUpById(Rogue)
	if lineup.AvatarIdList != nil {
		for _, avatar := range lineup.AvatarIdList {
			if avatar.AvatarId == 0 {
				continue
			}
			info.BaseAvatarIdList = append(info.BaseAvatarIdList, avatar.AvatarId)
		}
	}

	return info
}

func (g *PlayerData) GetRogueVirtualItem() *proto.RogueVirtualItem {
	info := &proto.RogueVirtualItem{
		RogueMoney: g.GetMaterialById(Cf),
	}

	return info
}

func (g *PlayerData) GetGameMiracleInfo() *proto.GameMiracleInfo {
	info := &proto.GameMiracleInfo{
		GameMiracleInfo: &proto.RogueMiracleInfo{
			MiracleList: make([]*proto.RogueMiracle, 0),
		},
	}

	return info
}

func (g *PlayerData) GetCurRogueBuff() []*proto.BattleBuff {
	buffList := make([]*proto.BattleBuff, 0)
	db := g.GetRogueBuffList()
	for _, buff := range db {
		buffList = append(buffList, &proto.BattleBuff{
			Id:              buff.BuffId,
			Level:           buff.BuffLevel,
			OwnerIndex:      4294967295,
			WaveFlag:        4294967295,
			TargetIndexList: make([]uint32, 0),
			DynamicValues:   make(map[string]float32),
		})
	}

	return buffList
}

// 获取初始action
func (g *PlayerData) GetFirstRogueAction() *proto.RogueCommonPendingAction {
	curRogue := g.GetCurRogue()
	if curRogue == nil || curRogue.Action == nil {
		return nil
	}
	fa := curRogue.Action[0]
	return g.GetRogueCommonPendingAction(0, fa)
}

// 获取action
func (g *PlayerData) GetRogueCommonPendingAction(queuePosition uint32, action *spb.RogueAction) *proto.RogueCommonPendingAction {
	if action == nil || action.Action == nil {
		return nil
	}
	switch a := action.Action.(type) {
	case *spb.RogueAction_BonusSelect:
		return g.getRogueActionBonusSelectInfo(queuePosition, a.BonusSelect)
	case *spb.RogueAction_BuffSelect:
		return g.getRogueActionBuffSelectInfo(queuePosition, a.BuffSelect)
	default:
		logger.Error(text.GetText(101), a)
	}
	return nil
}

// 获取action——Bonus
func (g *PlayerData) getRogueActionBonusSelectInfo(queuePosition uint32, a *spb.BonusSelect) *proto.RogueCommonPendingAction {
	action := &proto.RogueCommonPendingAction{
		QueuePosition: queuePosition,
		RogueAction: &proto.RogueAction{
			PendingAction: &proto.RogueAction_BonusSelectInfo{
				BonusSelectInfo: &proto.RogueBonusSelectInfo{
					BonusIdList: alg.Uin32KeyTList(a.BonusIdMap),
				},
			},
		},
	}

	return action
}

// 获取action——Buff
func (g *PlayerData) getRogueActionBuffSelectInfo(queuePosition uint32, a *spb.BuffSelect) *proto.RogueCommonPendingAction {
	action := &proto.RogueCommonPendingAction{
		QueuePosition: queuePosition,
		RogueAction: &proto.RogueAction{
			PendingAction: &proto.RogueAction_BuffSelectInfo{
				BuffSelectInfo: &proto.RogueCommonBuffSelectInfo{
					CanRoll:          a.CanRoll,
					RollBuffMaxCount: a.RollBuffMaxCount,
					SourceCurCount:   a.SourceCurCount,
					SourceHintId:     a.SourceHintId,
					SourceTotalCount: 1,
					SelectBuffList:   GetRogueCommonBuff(a.BuffMap),

					FirstBuffTypeList:        nil,
					RollBuffCount:            0,
					SourceType:               0,
					RollBuffFreeCount:        0,
					HandbookUnlockBuffIdList: nil,
					CertainSelectBuffId:      0,
					RollBuffCostData:         nil,
				},
			},
		},
	}

	return action
}

func GetRogueCommonBuff(buffMap map[uint32]*spb.RogueBuff) []*proto.RogueCommonBuff {
	list := make([]*proto.RogueCommonBuff, 0)
	for _, info := range buffMap {
		list = append(list, &proto.RogueCommonBuff{
			BuffId:    info.BuffId,
			BuffLevel: info.BuffLevel,
		})
	}
	return list
}

func (g *PlayerData) GetRogueBuffInfo() *proto.RogueBuffInfo {
	info := &proto.RogueBuffInfo{
		MazeBuffList: make([]*proto.RogueBuff, 0),
	}
	for _, in := range g.GetRogueBuffList() {
		info.MazeBuffList = append(info.MazeBuffList, &proto.RogueBuff{
			Level:  in.BuffLevel,
			BuffId: in.BuffId,
		})
	}
	return info
}

func (g *PlayerData) GetRogueBuffInfoById(buffId uint32) *proto.BuffInfo {
	db := g.GetRogueBuffById(buffId)
	if db == nil {
		return nil
	}
	info := &proto.BuffInfo{
		LifeTime:  -1,
		AddTimeMs: db.AddTime,
		Count:     4294967295,
		BuffId:    db.BuffId,
		Level:     db.BuffLevel,
	}
	return info
}
