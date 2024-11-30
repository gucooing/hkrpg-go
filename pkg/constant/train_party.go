package constant

type TrainPartyCardType string

const (
	TrainPartyCardTypeTrainPartyCardTypeNone TrainPartyCardType = ""
	TrainPartyCardTypeAddPassengerStat       TrainPartyCardType = "AddPassengerStat"
	TrainPartyCardTypeAddPassengerMemory     TrainPartyCardType = "AddPassengerMemory"
	TrainPartyCardTypeAddPassengerMood       TrainPartyCardType = "AddPassengerMood"
	TrainPartyCardTypeRecoverPamStamina      TrainPartyCardType = "RecoverPamStamina"
)
