// 记录战斗关键数据并储存，用于战斗结算和重连重新开启战斗
package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

type BattleState struct {
	ChallengeState *ChallengeState // ChallengeState
}
type ChallengeState struct {
	ChallengeId     uint32
	Status          proto.ChallengeStatus
	RoundCount      uint32
	ExtraLineupType proto.ExtraLineupType

	Type         uint32
	EntranceID   uint32
	BuffID       uint32
	MazeGroupID1 uint32
	MazeGroupID2 uint32
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

func (g *Game) GetChallengeState() *ChallengeState {
	return g.Player.BattleState.ChallengeState
}
