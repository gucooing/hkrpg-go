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
type TargetType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
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
type TargetInfo struct {
	TargetType     string `json:"TargetType"`
	MaxTargetCount int    `json:"MaxTargetCount"`
}
type Normalized struct {
	X            float64     `json:"x"`
	Y            float64     `json:"y"`
	Z            float64     `json:"z"`
	Normalized   *Normalized `json:"normalized"`
	Magnitude    float64     `json:"magnitude"`
	SqrMagnitude float64     `json:"sqrMagnitude"`
}
type TargetValue struct {
	X            float64    `json:"x"`
	Y            float64    `json:"y"`
	Z            float64    `json:"z"`
	Normalized   Normalized `json:"normalized"`
	Magnitude    float64    `json:"magnitude"`
	SqrMagnitude float64    `json:"sqrMagnitude"`
}
type DampChangeParam struct {
	TargetValue       TargetValue `json:"TargetValue"`
	Time              float64     `json:"Time"`
	ChangeCurvePath   string      `json:"ChangeCurvePath"`
	RecoveryTime      float64     `json:"RecoveryTime"`
	RecoveryCurvePath string      `json:"RecoveryCurvePath"`
}
type Freelook3RdConfig struct {
	DampChangeParam DampChangeParam `json:"DampChangeParam"`
}
type CameraConfig struct {
	Freelook3RdConfig Freelook3RdConfig `json:"Freelook3rdConfig"`
}
type Predicate struct {
	Type         string     `json:"$type"`
	TargetType   TargetType `json:"TargetType"`
	AnimZone     string     `json:"AnimZone"`
	TaskEnabled  bool       `json:"TaskEnabled"`
	IsClientOnly bool       `json:"IsClientOnly"`
}
type PerformerType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type NormalizedTimeEnd struct {
	FixedValue FixedValue `json:"fixedValue"`
}
type AttachOffset struct {
	Y            float64    `json:"y"`
	Normalized   Normalized `json:"normalized"`
	Magnitude    float64    `json:"magnitude"`
	SqrMagnitude float64    `json:"sqrMagnitude"`
}
type LinearPitchAngle struct {
}
type Projectile struct {
	ColliderTemplate               string           `json:"ColliderTemplate"`
	Behavior                       string           `json:"Behavior"`
	FlySpeed                       int              `json:"FlySpeed"`
	EnableRayCast                  bool             `json:"EnableRayCast"`
	FlyTime                        int              `json:"FlyTime"`
	Gravity                        float64          `json:"Gravity"`
	MaxLifeTime                    float64          `json:"MaxLifeTime"`
	AttachPoint                    string           `json:"AttachPoint"`
	AttachOffset                   AttachOffset     `json:"AttachOffset"`
	TargetAttachPoint              string           `json:"TargetAttachPoint"`
	FlyEffect                      string           `json:"FlyEffect"`
	HitEffect                      string           `json:"HitEffect"`
	FlyEffectFadeOut               bool             `json:"FlyEffectFadeOut"`
	LinearPitchAngle               LinearPitchAngle `json:"LinearPitchAngle"`
	BoomerangEccentricity          float64          `json:"BoomerangEccentricity"`
	TriggerHitCallback             bool             `json:"TriggerHitCallback"`
	StaticProjectileCanPassAirWall bool             `json:"StaticProjectileCanPassAirWall"`
	ParabolaUseWorldSpaceUp        bool             `json:"ParabolaUseWorldSpaceUp"`
}
type AttackTargetType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type AttackRootTargetType struct {
	Type  string `json:"$type"`
	Alias string `json:"Alias"`
}
type Offset struct {
	Y            float64 `json:"y"`
	Magnitude    float64 `json:"magnitude"`
	SqrMagnitude float64 `json:"sqrMagnitude"`
}
type AttackDetectConfig struct {
	Type      string  `json:"$type"`
	MaxRadius float64 `json:"MaxRadius"`
	FanAngle  int     `json:"FanAngle"`
	Hight     int     `json:"Hight"`
	Offset    Offset  `json:"Offset"`
}
type HitConfig struct {
	HitAnimation string `json:"HitAnimation"`
}
type OnProjectileHit struct {
	Type                  string               `json:"$type"`
	AttackTargetType      AttackTargetType     `json:"AttackTargetType"`
	AttackRootTargetType  AttackRootTargetType `json:"AttackRootTargetType"`
	TriggerBattle         bool                 `json:"TriggerBattle"`
	TriggerBattleDelay    float64              `json:"TriggerBattleDelay"`
	AttackDetectConfig    AttackDetectConfig   `json:"AttackDetectConfig"`
	HitConfig             HitConfig            `json:"HitConfig"`
	IncludeProps          bool                 `json:"IncludeProps"`
	AttackDetectCollision bool                 `json:"AttackDetectCollision"`
	TaskEnabled           bool                 `json:"TaskEnabled"`
	IsClientOnly          bool                 `json:"IsClientOnly"`
}
type OnProjectileLifetimeFinish struct {
	Type                  string               `json:"$type"`
	AttackTargetType      AttackTargetType     `json:"AttackTargetType"`
	AttackRootTargetType  AttackRootTargetType `json:"AttackRootTargetType"`
	TriggerBattle         bool                 `json:"TriggerBattle"`
	TriggerBattleDelay    float64              `json:"TriggerBattleDelay"`
	AttackDetectConfig    AttackDetectConfig   `json:"AttackDetectConfig"`
	HitConfig             HitConfig            `json:"HitConfig"`
	IncludeProps          bool                 `json:"IncludeProps"`
	AttackDetectCollision bool                 `json:"AttackDetectCollision"`
	TaskEnabled           bool                 `json:"TaskEnabled"`
	IsClientOnly          bool                 `json:"IsClientOnly"`
}
type SuccessTaskList struct {
	Type                       string                       `json:"$type"`
	TargetType                 any                          `json:"TargetType"`
	AnimLogicState             string                       `json:"AnimLogicState,omitempty"`
	AnimStateName              string                       `json:"AnimStateName,omitempty"`
	TransitionDuration         float64                      `json:"TransitionDuration,omitempty"`
	StopWhenHitOthers          bool                         `json:"StopWhenHitOthers,omitempty"`
	TaskEnabled                bool                         `json:"TaskEnabled"`
	IsClientOnly               bool                         `json:"IsClientOnly"`
	PerformerType              PerformerType                `json:"PerformerType,omitempty"`
	ToTargetRatio              int                          `json:"ToTargetRatio,omitempty"`
	Duration                   any                          `json:"Duration,omitempty"`
	NormalizedTimeEnd          NormalizedTimeEnd            `json:"NormalizedTimeEnd,omitempty"`
	Projectile                 Projectile                   `json:"Projectile,omitempty"`
	OnProjectileHit            []OnProjectileHit            `json:"OnProjectileHit,omitempty"`
	OnProjectileLifetimeFinish []OnProjectileLifetimeFinish `json:"OnProjectileLifetimeFinish,omitempty"`
	Active                     bool                         `json:"Active,omitempty"`
	BlurX                      float64                      `json:"BlurX,omitempty"`
	BlurY                      float64                      `json:"BlurY,omitempty"`
	BlurRadius                 float64                      `json:"BlurRadius,omitempty"`
	Iteration                  int                          `json:"Iteration,omitempty"`
	BlurStart                  float64                      `json:"BlurStart,omitempty"`
	BlurFeather                float64                      `json:"BlurFeather,omitempty"`
}
type FailedTaskList struct {
	Type                         string                       `json:"$type"`
	TargetType                   TargetType                   `json:"TargetType"`
	AnimLogicState               string                       `json:"AnimLogicState,omitempty"`
	AnimStateName                string                       `json:"AnimStateName,omitempty"`
	TransitionDuration           float64                      `json:"TransitionDuration,omitempty"`
	StopWhenHitOthers            bool                         `json:"StopWhenHitOthers,omitempty"`
	TaskEnabled                  bool                         `json:"TaskEnabled"`
	IsClientOnly                 bool                         `json:"IsClientOnly"`
	NormalizedTimeEnd            any                          `json:"NormalizedTimeEnd,omitempty"`
	Active                       bool                         `json:"Active,omitempty"`
	Duration                     any                          `json:"Duration,omitempty"`
	BlurX                        float64                      `json:"BlurX,omitempty"`
	BlurY                        float64                      `json:"BlurY,omitempty"`
	BlurRadius                   float64                      `json:"BlurRadius,omitempty"`
	Iteration                    int                          `json:"Iteration,omitempty"`
	BlurStart                    float64                      `json:"BlurStart,omitempty"`
	BlurFeather                  float64                      `json:"BlurFeather,omitempty"`
	Projectile                   Projectile                   `json:"Projectile,omitempty"`
	OnProjectileHit              []OnProjectileHit            `json:"OnProjectileHit,omitempty"`
	OnProjectileLifetimeFinish   []OnProjectileLifetimeFinish `json:"OnProjectileLifetimeFinish,omitempty"`
	PosTargetType                PosTargetType                `json:"PosTargetType,omitempty"`
	EffectPath                   string                       `json:"EffectPath,omitempty"`
	Scale                        Scale                        `json:"Scale,omitempty"`
	ParamEntityUsage             string                       `json:"ParamEntityUsage,omitempty"`
	TowardMaxPitchAngle          uint32                       `json:"TowardMaxPitchAngle,omitempty"`
	MaxMutexCount                uint32                       `json:"MaxMutexCount,omitempty"`
	TargetAlias                  TargetAlias                  `json:"TargetAlias,omitempty"`
	ForceStart                   bool                         `json:"ForceStart,omitempty"`
	NormalizedTransitionDuration float64                      `json:"NormalizedTransitionDuration,omitempty"`
}
type OnStart struct {
	Type            string            `json:"$type"`
	CameraConfig    CameraConfig      `json:"CameraConfig,omitempty"`
	TaskEnabled     bool              `json:"TaskEnabled"`
	IsClientOnly    bool              `json:"IsClientOnly,omitempty"`
	Predicate       Predicate         `json:"Predicate,omitempty"`
	SuccessTaskList []SuccessTaskList `json:"SuccessTaskList,omitempty"`
	FailedTaskList  []FailedTaskList  `json:"FailedTaskList,omitempty"`
}
type AbilityList struct {
	OnAbort    []OnAbort  `json:"OnAbort,omitempty"`
	Name       string     `json:"Name"`
	TargetInfo TargetInfo `json:"TargetInfo"`
	OnStart    []OnStart  `json:"OnStart"`
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
