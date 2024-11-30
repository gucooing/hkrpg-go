package constant

type RogueTournRoomType uint32

const (
	RogueTournRoomTypeUnknown   RogueTournRoomType = 0
	RogueTournRoomTypeBoss      RogueTournRoomType = 1
	RogueTournRoomTypeElite     RogueTournRoomType = 2
	RogueTournRoomTypeBattle    RogueTournRoomType = 3
	RogueTournRoomTypeEncounter RogueTournRoomType = 4
	RogueTournRoomTypeEvent     RogueTournRoomType = 5
	RogueTournRoomTypeCoin      RogueTournRoomType = 6
	RogueTournRoomTypeShop      RogueTournRoomType = 7
	RogueTournRoomTypeReward    RogueTournRoomType = 8
	RogueTournRoomTypeAdventure RogueTournRoomType = 9
	RogueTournRoomTypeRespite   RogueTournRoomType = 10
	RogueTournRoomTypeReforge   RogueTournRoomType = 11
	RogueTournRoomTypeHidden    RogueTournRoomType = 12
)

type RogueBuffCategory = string

const (
	RogueBuffCategoryNone      RogueBuffCategory = ""
	RogueBuffCategoryCommon    RogueBuffCategory = "Common"
	RogueBuffCategoryRare      RogueBuffCategory = "Rare"
	RogueBuffCategoryLegendary RogueBuffCategory = "Legendary"
)

type RogueBuffAeonType = string

const (
	RogueBuffAeonTypeNormal                 RogueBuffAeonType = ""
	RogueBuffAeonTypeBattleEventBuff        RogueBuffAeonType = "BattleEventBuff"
	RogueBuffAeonTypeBattleEventBuffEnhance RogueBuffAeonType = "BattleEventBuffEnhance"
	RogueBuffAeonTypeBattleEventBuffCross   RogueBuffAeonType = "BattleEventBuffCross"
)
