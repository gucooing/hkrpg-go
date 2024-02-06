// 记录战斗关键数据并储存，用于战斗结算和重连重新开启战斗

package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type BattleState struct {
	BattleType         spb.BattleType
	ChallengeState     *ChallengeState
	TrialActivityState *TrialActivityState
	BuffList           []uint32 // 进入战斗需要添加的buff
	RogueState         *RogueState
}
type ChallengeState struct {
	// 回包
	ChallengeId     uint32                // 关卡id
	Status          proto.ChallengeStatus // 状态
	RoundCount      uint32                // 已使用的回合数
	ExtraLineupType proto.ExtraLineupType // 哪个队伍
	ChallengeScore  uint32                // 活动分数
	// 缓存状态
	Pos                *spb.VectorBin
	Rot                *spb.VectorBin
	NPCMonsterPos      *spb.VectorBin
	NPCMonsterRot      *spb.VectorBin
	PlaneID            uint32
	FloorID            uint32
	EntranceID         uint32
	CurChallengeBattle map[uint32]*CurChallengeBattle // 每一波关卡配置
	SceneBuffList      []uint32                       // 场景buff
	AvatarBuffList     []uint32                       // 角色buff
	MonsterEntityMap   []uint32                       // 当前战斗实体id
	// 下面是普通
	ChallengeCount     uint32   // 波数
	CurChallengeCount  uint32   // 当前波次
	ChallengeTargetID  []uint32 // 满星条件
	ChallengeCountDown uint32   // 总回合限制
	// 下面是活动
	StoryBuffOne uint32 // 第一个buff
	StoryBuffTwo uint32 // 第二个buff
	ScoreOne     uint32
	ScoreTwo     uint32
}
type TrialActivityState struct {
	AvatarDemoId   uint32
	NPCMonsterPos  *spb.VectorBin
	NPCMonsterRot  *spb.VectorBin
	PlaneID        uint32
	FloorID        uint32
	EntranceID     uint32
	AvatarBuffList []uint32 // 角色buff

	NPCMonsterID uint32
	EventID      uint32
	GroupID      uint32
	ConfigID     uint32
}

type CurChallengeBattle struct {
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
	BuffNum        uint32
	AvatarBuffList []uint32
	BuffList       map[uint32]*RogueBuff
	AvatarEntity   map[uint32]*AvatarEntity
	MonsterEntity  map[uint32]*MonsterEntity
	Battle         map[uint32]*RogueBattle
}
type RogueBattle struct {
	monsterEntityMap []uint32 // 实体列表
}
type RogueBuff struct {
	Level     uint32
	AddTimeMs uint64
}

func (g *GamePlayer) GetBattleState() *BattleState {
	if g.Player.BattleState == nil {
		g.Player.BattleState = &BattleState{
			BattleType:         0,
			ChallengeState:     &ChallengeState{},
			BuffList:           make([]uint32, 0),
			TrialActivityState: &TrialActivityState{},
			RogueState:         &RogueState{},
		}
	}
	return g.Player.BattleState
}

func (g *GamePlayer) NewChallengeState() *ChallengeState {
	g.GetBattleState().ChallengeState = &ChallengeState{}
	return g.GetBattleState().ChallengeState
}

func (g *GamePlayer) GetChallengeState() *ChallengeState {
	if g.GetBattleState().ChallengeState == nil {
		g.GetBattleState().ChallengeState = &ChallengeState{}
	}
	return g.GetBattleState().ChallengeState
}

func (g *GamePlayer) GetTrialActivityState() *TrialActivityState {
	if g.GetBattleState().TrialActivityState == nil {
		g.GetBattleState().TrialActivityState = &TrialActivityState{}
	}
	return g.GetBattleState().TrialActivityState
}

func (g *GamePlayer) GetBattle() *spb.Battle {
	if g.PlayerPb.Battle == nil {
		g.PlayerPb.Battle = &spb.Battle{}
	}
	if g.PlayerPb.Battle.Challenge == nil {
		g.PlayerPb.Battle.Challenge = &spb.Challenge{
			ChallengeList:       make(map[uint32]*spb.ChallengeList),
			ChallengeRewardList: make(map[uint64]uint32),
		}
	}
	return g.PlayerPb.Battle
}

func (g *GamePlayer) GetChallenge() *spb.Challenge {
	battle := g.GetBattle()
	if battle.Challenge.ChallengeList == nil {
		battle.Challenge.ChallengeList = make(map[uint32]*spb.ChallengeList)
	}
	if battle.Challenge.ChallengeRewardList == nil {
		battle.Challenge.ChallengeRewardList = make(map[uint64]uint32)
	}
	return battle.Challenge
}

func (g *GamePlayer) GetChallengeById(id uint32) *spb.ChallengeList {
	battle := g.GetChallenge()
	if battle.ChallengeList[id] == nil {
		battle.ChallengeList[id] = &spb.ChallengeList{}
	}
	return battle.ChallengeList[id]
}

func (g *GamePlayer) GetDbRogue() *spb.Rogue {
	if g.GetBattle().Rogue == nil {
		g.GetBattle().Rogue = &spb.Rogue{
			RogueArea: make(map[uint32]*spb.RogueArea),
		}

		g.GetBattle().Rogue.RogueArea[100] = &spb.RogueArea{
			AreaId:          100,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_UNLOCK,
		}
		g.GetBattle().Rogue.RogueArea[110] = &spb.RogueArea{
			AreaId:          110,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_UNLOCK,
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

func (g *GamePlayer) GetDbRogueArea(areaId uint32) *spb.RogueArea {
	rogue := g.GetDbRogue()
	if rogue.RogueArea[areaId] == nil {
		rogue.RogueArea[areaId] = &spb.RogueArea{
			AreaId:          areaId,
			RogueAreaStatus: spb.RogueAreaStatus_RogueAreaStatus_ROGUE_AREA_STATUS_LOCK,
		}
	}

	return rogue.RogueArea[areaId]
}

func (g *GamePlayer) NewRogueState() {
	g.GetBattleState().RogueState = &RogueState{
		BuffNum:        0,
		AvatarBuffList: make([]uint32, 0),
		AvatarEntity:   make(map[uint32]*AvatarEntity),
		MonsterEntity:  make(map[uint32]*MonsterEntity),
		Battle:         make(map[uint32]*RogueBattle),
		BuffList:       make(map[uint32]*RogueBuff),
	}
}

func (g *GamePlayer) GetRogue() *RogueState {
	if g.GetBattleState().RogueState == nil {
		g.GetBattleState().RogueState = &RogueState{
			AvatarBuffList: make([]uint32, 0),
			AvatarEntity:   make(map[uint32]*AvatarEntity),
			MonsterEntity:  make(map[uint32]*MonsterEntity),
			Battle:         make(map[uint32]*RogueBattle),
			BuffList:       make(map[uint32]*RogueBuff),
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
