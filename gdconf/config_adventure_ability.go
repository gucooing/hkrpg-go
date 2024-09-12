package gdconf

import (
	"os"
	"regexp"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ConfigAdventureAbility struct {
	LocalPlayer map[uint32]*LocalPlayerAbility
}

type LocalPlayerAbility struct {
	AbilityList []AbilityList `json:"AbilityList"`
}

type AbilityList struct {
	OnAbort    []OnAbort   `json:"OnAbort,omitempty"`
	Name       string      `json:"Name"`
	TargetInfo TargetInfo  `json:"TargetInfo"`
	OnStart    []*TaskInfo `json:"OnStart"`
}

type OnAbort struct {
	Type         string     `json:"$type"`
	TargetType   TargetType `json:"TargetType"`
	Duration     float64    `json:"Duration"`
	BlurX        float64    `json:"BlurX"`
	BlurY        float64    `json:"BlurY"`
	BlurRadius   float64    `json:"BlurRadius"`
	Iteration    int        `json:"Iteration"`
	BlurStart    float64    `json:"BlurStart"`
	BlurFeather  float64    `json:"BlurFeather"`
	TaskEnabled  bool       `json:"TaskEnabled"`
	IsClientOnly bool       `json:"IsClientOnly"`
}
type TargetType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type TargetInfo struct {
	TargetType     string `json:"TargetType"`
	MaxTargetCount int    `json:"MaxTargetCount"`
}

type TaskInfo struct {
	Type                       string        `json:"$type"`
	ID                         uint32        `json:"ID"`
	ModifierName               string        `json:"ModifierName"`
	TriggerBattle              bool          `json:"TriggerBattle"`
	IsClientOnly               bool          `json:"IsClientOnly"`
	TaskEnabled                bool          `json:"TaskEnabled"`
	LifeTime                   *DynamicFloat `json:"LifeTime"`
	OnAttack                   []*TaskInfo   `json:"OnAttack"`
	OnBattle                   []*TaskInfo   `json:"OnBattle"`
	SuccessTaskList            []*TaskInfo   `json:"SuccessTaskList"`
	OnProjectileHit            []*TaskInfo   `json:"OnProjectileHit"`
	OnProjectileLifetimeFinish []*TaskInfo   `json:"OnProjectileLifetimeFinish"`
}

type DynamicFloat struct {
	IsDynamic bool `json:"IsDynamic"`
	// FixedValue
}

func (g *GameDataConfig) loadConfigAdventureAbility() {
	g.ConfigAdventureAbility = &ConfigAdventureAbility{
		LocalPlayer: make(map[uint32]*LocalPlayerAbility),
	}
	re := regexp.MustCompile(`Avatar_(.*?)_Config`)
	for _, avatarData := range g.AvatarDataMap {
		matches := re.FindAllStringSubmatch(avatarData.ManikinJsonPath, -1)
		if len(matches) < 1 {
			continue
		}
		localPlayerAbility := new(LocalPlayerAbility)
		playerElementsFilePath := g.pathPrefix + "/Config/ConfigAdventureAbility/LocalPlayer/LocalPlayer_" + matches[0][1] + "_Ability.json"
		playerElementsFile, err := os.ReadFile(playerElementsFilePath)
		if err != nil {
			logger.Error("open file error: %v", err)
			continue
		}

		err = hjson.Unmarshal(playerElementsFile, &localPlayerAbility)
		if err != nil {
			logger.Error("parse file error: %v", err)
			continue
		}
		g.ConfigAdventureAbility.LocalPlayer[avatarData.AvatarId] = localPlayerAbility
	}

	logger.Info("load %v LocalPlayerAbility", len(g.ConfigAdventureAbility.LocalPlayer))
}

func GetAvatarAbilityMap() map[uint32]*LocalPlayerAbility {
	return CONF.ConfigAdventureAbility.LocalPlayer
}

func GetAvatarAbility(avatarId uint32) *LocalPlayerAbility {
	return CONF.ConfigAdventureAbility.LocalPlayer[avatarId]
}
