package constant

type ItemMainType = string

const (
	ItemMainTypeUnknown    ItemMainType = "Unknown"
	ItemMainTypeVirtual    ItemMainType = "Virtual"
	ItemMainTypeAvatarCard ItemMainType = "AvatarCard"
	ItemMainTypeEquipment  ItemMainType = "Equipment"
	ItemMainTypeRelic      ItemMainType = "Relic"
	ItemMainTypeUsable     ItemMainType = "Usable"
	ItemMainTypeMaterial   ItemMainType = "Material"
	ItemMainTypeMission    ItemMainType = "Mission"
	ItemMainTypeDisplay    ItemMainType = "Display"
	ItemMainTypePet        ItemMainType = "Pet"
)

type ItemSubType = string

const (
	ItemSubTypeUnknown               ItemSubType = "Unknown"
	ItemSubTypeVirtual               ItemSubType = "Virtual"
	ItemSubTypeGameplayCounter       ItemSubType = "GameplayCounter"
	ItemSubTypeAvatarCard            ItemSubType = "AvatarCard"
	ItemSubTypeEquipment             ItemSubType = "Equipment"
	ItemSubTypeRelic                 ItemSubType = "Relic"
	ItemSubTypeGift                  ItemSubType = "Gift"
	ItemSubTypeFood                  ItemSubType = "Food"
	ItemSubTypeForceOpitonalGift     ItemSubType = "ForceOpitonalGift"
	ItemSubTypeBook                  ItemSubType = "Book"
	ItemSubTypeHeadIcon              ItemSubType = "HeadIcon"
	ItemSubTypeMusicAlbum            ItemSubType = "MusicAlbum"
	ItemSubTypeFormula               ItemSubType = "Formula"
	ItemSubTypeChatBubble            ItemSubType = "ChatBubble"
	ItemSubTypeAvatarSkin            ItemSubType = "AvatarSkin"
	ItemSubTypePhoneTheme            ItemSubType = "PhoneTheme"
	ItemSubTypeTravelBrochurePaster  ItemSubType = "TravelBrochurePaster"
	ItemSubTypeChessRogueDiceSurface ItemSubType = "ChessRogueDiceSurface"
	ItemSubTypeRogueMedal            ItemSubType = "RogueMedal"
	ItemSubTypeMaterial              ItemSubType = "Material"
	ItemSubTypeEidolon               ItemSubType = "Eidolon"
	ItemSubTypeMuseumExhibit         ItemSubType = "MuseumExhibit"
	ItemSubTypeMuseumStuff           ItemSubType = "MuseumStuff"
	ItemSubTypeAetherSkill           ItemSubType = "AetherSkill"
	ItemSubTypeAetherSpirit          ItemSubType = "AetherSpirit"
	ItemSubTypeMission               ItemSubType = "Mission"
	ItemSubTypeRelicSetShowOnly      ItemSubType = "RelicSetShowOnly"
	ItemSubTypeRelicRarityShowOnly   ItemSubType = "RelicRarityShowOnly"
)

type FormulaType = string

const (
	FormulaTypeUnknown       FormulaType = "Unknown"
	FormulaTypeNormal        FormulaType = "Normal"
	FormulaTypeSepcial       FormulaType = "Sepcial"
	FormulaTypeSelectedRelic FormulaType = "SelectedRelic"
)

type LimitType = string

const (
	LimitTypeNull              LimitType = "Null"
	LimitTypeLevel             LimitType = "Level"
	LimitTypeMainMission       LimitType = "MainMission"
	LimitTypeEventMission      LimitType = "EventMission"
	LimitTypeWorldLevel        LimitType = "WorldLevel"
	LimitTypePreGoods          LimitType = "PreGoods"
	LimitTypeHasNoRefreshGoods LimitType = "HasNoRefreshGoods"
	LimitTypeSubMission        LimitType = "SubMission"
)

type FuncType = string

const (
	FuncTypeUnknown FuncType = "Unknown"
	FuncTypeCompose FuncType = "Compose"
	FuncTypeReplace FuncType = "Replace"
)

type ItemFoodUseTag = string

const (
	ItemFoodUseTagHP ItemFoodUseTag = "HP"
	ItemFoodUseTagBP ItemFoodUseTag = "BP"
	ItemFoodUseTagSP ItemFoodUseTag = "SP"
)
