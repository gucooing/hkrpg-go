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

type MapCameraType string

const (
	MapCameraTypeBase       MapCameraType = ""
	MapCameraTypeMaze       MapCameraType = "Maze"
	MapCameraTypeRoom       MapCameraType = "Room"
	MapCameraTypeLargeSpace MapCameraType = "LargeSpace"
)

type LevelFeatureType string

const (
	LevelFeatureTypeUnknown         LevelFeatureType = ""
	LevelFeatureTypeRotatableRegion LevelFeatureType = "RotatableRegion"
	LevelFeatureTypeHeartDial       LevelFeatureType = "HeartDial"
	LevelFeatureTypeEraFlipper      LevelFeatureType = "EraFlipper"
)

type LevelDimensionCategory string

const (
	LevelDimensionCategoryMain                LevelDimensionCategory = ""
	LevelDimensionCategoryStoryLine           LevelDimensionCategory = "StoryLine"
	LevelDimensionCategoryActivityEarlyAccess LevelDimensionCategory = "ActivityEarlyAccess"
	LevelDimensionCategoryActivityFirstAccess LevelDimensionCategory = "ActivityFirstAccess"
	LevelDimensionCategoryCustom              LevelDimensionCategory = "Custom"
)

type CompareType string

const (
	CompareTypeUnknow       CompareType = ""
	CompareTypeGreater      CompareType = "Greater"
	CompareTypeGreaterEqual CompareType = "GreaterEqual"
	CompareTypeNotEqual     CompareType = "NotEqual"
	CompareTypeEqual        CompareType = "Equal"
	CompareTypeLessEqual    CompareType = "LessEqual"
	CompareTypeLess         CompareType = "Less"
)
