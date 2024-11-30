package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RaidConfig struct {
	RaidID                  uint32                      `json:"RaidID"`
	HardLevel               uint32                      `json:"HardLevel"`
	RaidTagList             []constant.RaidTagType      `json:"RaidTagList"`
	UnlockWorldLevel        []uint32                    `json:"UnlockWorldLevel"`
	MonsterList             []uint32                    `json:"MonsterList"`
	MonsterHideList         []uint32                    `json:"MonsterHideList"`
	LimitIDList             []uint32                    `json:"LimitIDList"`
	RecoverType             []constant.RaidRecoverType  `json:"RecoverType"`
	RewardList              []uint32                    `json:"RewardList"`
	TrialAvatarList         []uint32                    `json:"TrialAvatarList"`
	MainMissionIDList       []uint32                    `json:"MainMissionIDList"`
	EntrancePageBGImagePath string                      `json:"EntrancePageBGImagePath"`
	DamageType              []constant.AttackDamageType `json:"DamageType"`
	RaidTargetID            []uint32                    `json:"RaidTargetID"`
	Type                    constant.RaidConfigType     `json:"Type"`
	MappingInfoID           uint32                      `json:"MappingInfoID"`
	DisplayEventID          uint32                      `json:"DisplayEventID"`
	FinishEntranceID        uint32                      `json:"FinishEntranceID"`
	TeamType                constant.RaidTeamType       `json:"TeamType"`
	RecommendLevel          uint32                      `json:"RecommendLevel"`
	TeamPositionLockNum     uint32                      `json:"TeamPositionLockNum"`
	MainMissionIDBefore     uint32                      `json:"MainMissionIDBefore"`
	MainMissionIDAfter      uint32                      `json:"MainMissionIDAfter"`
	LockCaptainAvatarID     uint32                      `json:"LockCaptainAvatarID"`
	EnterType               constant.RaidEnterType      `json:"EnterType"`
	IsEntryByProp           bool                        `json:"IsEntryByProp"`
	SkipRewardOnFinish      bool                        `json:"SkipRewardOnFinish"`
	AutoObtainDamageType    bool                        `json:"AutoObtainDamageType"`
	LockCaptain             bool                        `json:"LockCaptain"`
	IsHiddenAreaMap         bool                        `json:"IsHiddenAreaMap"`
}

func (g *GameDataConfig) loadRaidConfig() {
	g.RaidConfigMap = make(map[uint32]map[uint32]*RaidConfig)
	raidConfigMap := make([]*RaidConfig, 0)
	name := "RaidConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &raidConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range raidConfigMap {
		if g.RaidConfigMap[v.RaidID] == nil {
			g.RaidConfigMap[v.RaidID] = make(map[uint32]*RaidConfig)
		}
		g.RaidConfigMap[v.RaidID][v.HardLevel] = v
	}

	logger.Info(text.GetText(17), len(g.RaidConfigMap), name)
}

func GetRaidConfig(raidID, hardLevel uint32) *RaidConfig {
	if getConf().RaidConfigMap[raidID] == nil {
		return nil
	}
	return getConf().RaidConfigMap[raidID][hardLevel]
}
