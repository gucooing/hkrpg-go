package constant

type EntranceShowType string

const (
	EntranceShowTypeNone EntranceShowType = ""
	InBPSchedule         EntranceShowType = "InBPSchedule"
	HasActivity          EntranceShowType = "HasActivity"
)

type LogicOperationType string

const (
	LogicOperationTypeUnknow LogicOperationType = ""
	LogicOperationTypeAnd    LogicOperationType = "And"
	LogicOperationTypeOr     LogicOperationType = "Or"
	LogicOperationTypeNot    LogicOperationType = "Not"
)

type LevelGroupMissionType string

const (
	LevelGroupMissionTypeMainMission  LevelGroupMissionType = ""
	LevelGroupMissionTypeSubMission   LevelGroupMissionType = "SubMission"
	LevelGroupMissionTypeEventMission LevelGroupMissionType = "EventMission"
)

type LevelGroupMissionPhase string

const (
	LevelGroupMissionPhaseAccept LevelGroupMissionPhase = ""
	LevelGroupMissionPhaseFinish LevelGroupMissionPhase = "Finish"
	LevelGroupMissionPhaseCancel LevelGroupMissionPhase = "Cancel"
)

type ConditionType string

const (
	ConditionTypeNone                  ConditionType = ""
	ConditionTypeFinishMainMission     ConditionType = "FinishMainMission"
	ConditionTypePlayerLevel           ConditionType = "PlayerLevel"
	ConditionTypeWorldLevel            ConditionType = "WorldLevel"
	ConditionTypeFinishChallenge       ConditionType = "FinishChallenge"
	ConditionTypeNotInPlaneType        ConditionType = "NotInPlaneType"
	ConditionTypeAvatarLevel           ConditionType = "AvatarLevel"
	ConditionTypeFinishSubMission      ConditionType = "FinishSubMission"
	ConditionTypeFinishQuest           ConditionType = "FinishQuest"
	ConditionTypeMaxPlayerLevel        ConditionType = "MaxPlayerLevel"
	ConditionTypeQuestClose            ConditionType = "QuestClose"
	ConditionTypeCanUseFoodInRogue     ConditionType = "CanUseFoodInRogue"
	ConditionTypeBetweenSubMission     ConditionType = "BetweenSubMission"
	ConditionTypeInStoryLine           ConditionType = "InStoryLine"
	ConditionTypeReleaseContentPackage ConditionType = "ReleaseContentPackage"
	ConditionTypeSubMissionTaken       ConditionType = "SubMissionTaken"
	ConditionTypeHasItemMainType       ConditionType = "HasItemMainType"
)
