// 记录战斗关键数据并储存，用于战斗结算和重连重新开启战斗
package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

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

type Challenge struct {
	Type       uint32
	Lineup     map[uint32]*ChallengeLineUp
	EntranceID uint32
	BuffID     uint32
	//
	ChallengeId uint32
	Status      proto.ChallengeStatus
	RoundCount  uint32
}
type ChallengeLineUp struct {
	GroupID         uint32
	ExtraLineupType proto.ExtraLineupType
	Slots           []*ChallengeSlots
}
type ChallengeSlots struct {
	Slot     uint32
	AvatarId uint32
	Type     proto.AvatarType
}

func NewChallenge(data *PlayerData) *PlayerData {
	challenge := new(Challenge)
	challenge.Lineup = make(map[uint32]*ChallengeLineUp)
	data.Challenge = challenge
	return data
}

func (g *Game) GetChallengeLineUp() *proto.LineupInfo {
	index := g.Player.Challenge.Type
	lineUp := g.Player.Challenge.Lineup[index]
	lineupList := &proto.LineupInfo{
		IsVirtual:       false,
		LeaderSlot:      0,
		AvatarList:      make([]*proto.LineupAvatar, 0),
		Index:           index,
		ExtraLineupType: lineUp.ExtraLineupType,
		MaxMp:           5,
		Mp:              5,
		PlaneId:         0,
	}
	for _, slots := range lineUp.Slots {
		if slots.AvatarId == 0 {
			continue
		}
		avatar := g.PlayerPb.Avatar.Avatar[slots.AvatarId]
		lineupAvatar := &proto.LineupAvatar{
			AvatarType: slots.Type,
			Slot:       slots.Slot,
			Satiety:    0,
			Hp:         avatar.Hp,
			Id:         slots.AvatarId,
			SpBar: &proto.SpBarInfo{
				CurSp: avatar.SpBar.CurSp,
				MaxSp: avatar.SpBar.MaxSp,
			},
		}
		lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
	}

	return lineupList
}
