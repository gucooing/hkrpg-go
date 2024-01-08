// 记录战斗关键数据并储存，用于战斗结算和重连重新开启战斗

package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type BattleState struct {
	BattleType     spb.BattleType
	ChallengeState *ChallengeState // ChallengeState
	BuffList       []uint32        // 进入战斗需要添加的buff
}
type ChallengeState struct {
	ChallengeCount     uint32   // 波数
	CurChallengeCount  uint32   // 当前波次
	ChallengeTargetID  []uint32 // 满星条件
	ChallengeCountDown uint32
	// 回包
	ChallengeId     uint32
	Status          proto.ChallengeStatus
	RoundCount      uint32
	ExtraLineupType proto.ExtraLineupType
	// 缓存状态
	Pos                *spb.VectorBin
	Rot                *spb.VectorBin
	NPCMonsterPos      *spb.VectorBin
	NPCMonsterRot      *spb.VectorBin
	PlaneID            uint32
	FloorID            uint32
	EntranceID         uint32
	CurChallengeBattle map[uint32]*CurChallengeBattle
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

func (g *Game) GetBattleState() *BattleState {
	return g.Player.BattleState
}

func (g *Game) GetChallengeState() *ChallengeState {
	return g.Player.BattleState.ChallengeState
}

func (g *Game) GetBattle() *spb.Battle {
	if g.PlayerPb.Battle == nil {
		g.PlayerPb.Battle = &spb.Battle{}
	}
	if g.PlayerPb.Battle.Challenge == nil {
		g.PlayerPb.Battle.Challenge = &spb.Challenge{
			ChallengeList:       make(map[uint32]uint32),
			ChallengeRewardList: make(map[uint64]uint32),
		}
	}
	return g.PlayerPb.Battle
}

func (g *Game) GetChallenge() *spb.Challenge {
	battle := g.GetBattle()
	if battle.Challenge.ChallengeList == nil {
		battle.Challenge.ChallengeList = make(map[uint32]uint32)
	}
	if battle.Challenge.ChallengeRewardList == nil {
		battle.Challenge.ChallengeRewardList = make(map[uint64]uint32)
	}
	return battle.Challenge
}
