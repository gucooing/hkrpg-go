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
	ChallengeCount uint32 // 波数
	// 回包
	ChallengeId     uint32
	Status          proto.ChallengeStatus
	RoundCount      uint32
	ExtraLineupType proto.ExtraLineupType
	// 缓存状态
	Pos               *spb.VectorBin
	Rot               *spb.VectorBin
	CurChallengeCount uint32 // 当前波次
	EntranceID        uint32 // 场景
	MazeGroupID1      uint32 // 区块1
	MazeGroupID2      uint32 // 区块2
	MazeGroupID3      uint32 // 区块3
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
