package constant

type RaidTagType = string

const (
	RaidTagTypeNone       RaidTagType = "None"
	RaidTagTypeNormalRaid RaidTagType = "NormalRaid"
	RaidTagTypeHardRaid   RaidTagType = "HardRaid"
	RaidTagTypeHardest    RaidTagType = "Hardest"
)

type RaidRecoverType = string

const (
	RaidRecoverTypeUnknown       RaidRecoverType = "Unknown"
	RaidRecoverTypeRecoverHp     RaidRecoverType = "RecoverHp"
	RaidRecoverTypeRecoverMp     RaidRecoverType = "RecoverMp"
	RaidRecoverTypeResetMp       RaidRecoverType = "ResetMp"
	RaidRecoverTypeRecoverSp     RaidRecoverType = "RecoverSp"
	RaidRecoverTypeResetSp       RaidRecoverType = "ResetSp"
	RaidRecoverTypeRecoverHalfSp RaidRecoverType = "RecoverHalfSp"
)

type AttackDamageType = string

const (
	AttackDamageTypeUnknow    AttackDamageType = "Unknow"
	AttackDamageTypePhysical  AttackDamageType = "Physical"
	AttackDamageTypeFire      AttackDamageType = "Fire"
	AttackDamageTypeIce       AttackDamageType = "Ice"
	AttackDamageTypeThunder   AttackDamageType = "Thunder"
	AttackDamageTypeWind      AttackDamageType = "Wind"
	AttackDamageTypeQuantum   AttackDamageType = "Quantum"
	AttackDamageTypeImaginary AttackDamageType = "Imaginary"
	AttackDamageTypeHeal      AttackDamageType = "Heal"
	AttackDamageTypeAllType   AttackDamageType = "AllType"
)

type RaidConfigType = string

const (
	RaidConfigTypeNone                   RaidConfigType = "None"
	RaidConfigTypeMission                RaidConfigType = "Mission"
	RaidConfigTypeRelic                  RaidConfigType = "Relic"
	RaidConfigTypeChallenge              RaidConfigType = "Challenge"
	RaidConfigTypeTreasureChallenge      RaidConfigType = "TreasureChallenge"
	RaidConfigTypePunkLord               RaidConfigType = "PunkLord"
	RaidConfigTypeSaveMission            RaidConfigType = "SaveMission"
	RaidConfigTypeTrial                  RaidConfigType = "Trial"
	RaidConfigTypeEquilibriumTrial       RaidConfigType = "EquilibriumTrial"
	RaidConfigTypeHeliobus               RaidConfigType = "Heliobus"
	RaidConfigTypeActivityRaidCollection RaidConfigType = "ActivityRaidCollection"
	RaidConfigTypeClockPark              RaidConfigType = "ClockPark"
)

type RaidTeamType = string

const (
	RaidTeamTypeNone           RaidTeamType = "None"
	RaidTeamTypePlayer         RaidTeamType = "Player"
	RaidTeamTypeTrial          RaidTeamType = "Trial"
	RaidTeamTypeTrialAndPlayer RaidTeamType = "TrialAndPlayer"
	RaidTeamTypeTrialOnly      RaidTeamType = "TrialOnly"
	RaidTeamTypeTrialOrPlayer  RaidTeamType = "TrialOrPlayer"
)

type RaidEnterType = string

const (
	RaidEnterTypeDefault     RaidEnterType = "Default"
	RaidEnterTypeSkipUI      RaidEnterType = "SkipUI"
	RaidEnterTypePerformance RaidEnterType = "Performance"
)
