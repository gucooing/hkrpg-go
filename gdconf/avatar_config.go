package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type AvatarData struct {
	AvatarId        uint32        `json:"AvatarID"` // 角色id
	Rarity          string        `json:"Rarity"`   // 星级
	JsonPath        string        `json:"JsonPath"`
	ManikinJsonPath string        `json:"ManikinJsonPath"`
	RankIDList      []uint32      `json:"RankIDList"` // 命座id
	DamageType      string        `json:"DamageType"`
	ExpGroup        uint32        `json:"ExpGroup"`       // 经验ID
	RewardList      []*RewardList `json:"RewardList"`     // 重复获得角色奖励
	AvatarBaseType  string        `json:"AvatarBaseType"` // 角色类型
	SkillList       []uint32      `json:"SkillList"`
	Release         bool          `json:"Release"`
}

type RewardList struct {
	ItemID  uint32 `json:"ItemID"`  // 物品id
	ItemNum uint32 `json:"ItemNum"` // 个数
}

func (g *GameDataConfig) loadAvatarData() {
	g.AvatarDataMap = make(map[uint32]*AvatarData)
	avatarDataMap := make([]*AvatarData, 0)
	name := "AvatarConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarDataMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range avatarDataMap {
		g.AvatarDataMap[v.AvatarId] = v
	}
	logger.Info(text.GetText(17), len(g.AvatarDataMap), name)
	g.loadConfigAdventureAbility()
}

func GetAvatarDataById(avatarId uint32) *AvatarData {
	return getConf().AvatarDataMap[avatarId]
}

func GetAvatarDataMap() map[uint32]*AvatarData {
	return getConf().AvatarDataMap
}

var damageTypeEnum = map[string]uint32{
	"Physical":  1000111,
	"Fire":      1000112,
	"Ice":       1000113,
	"Thunder":   1000114,
	"Wind":      1000115,
	"Quantum":   1000116,
	"Imaginary": 1000117,
}

func GetAvatarDamage(id uint32) uint32 {
	conf := getConf().AvatarDataMap[id]
	if conf == nil {
		return 0
	}
	return damageTypeEnum[conf.DamageType]
}

func GetAvatarList() []uint32 {
	var avatarList []uint32
	for _, avatar := range getConf().AvatarDataMap {
		avatarList = append(avatarList, avatar.AvatarId)
	}
	return avatarList
}
