// 记录战斗关键数据并储存，用于战斗结算和重连重新开启战斗

package player

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type BattleState struct {
	BattleType         spb.BattleType
	ChallengeState     *ChallengeState
	TrialActivityState *TrialActivityState
	BuffList           []uint32 // 进入战斗需要添加的buff
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
	EventID            uint32                         // 当前战斗实体id
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
	EventID          uint32                // 怪物群实体id
	LogicRandomSeed  uint32                // 逻辑随机种子
	RoundsLimit      uint32                // 回合限制
	StaminaCost      uint32                // 扣除体力
	DisplayItemList  []*Material           // 奖励物品
	BuffList         []*proto.BattleBuff   // Buff列表
	BattleAvatarList []*proto.BattleAvatar // 战斗角色列表
}

type Rogue struct {
}

func (g *GamePlayer) GetBattleState() *BattleState {
	if g.Player.BattleState == nil {
		g.Player.BattleState = &BattleState{
			BattleType:         0,
			ChallengeState:     &ChallengeState{},
			BuffList:           make([]uint32, 0),
			TrialActivityState: &TrialActivityState{},
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
