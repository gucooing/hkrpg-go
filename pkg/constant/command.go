package constant

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/lua"
	"github.com/gucooing/hkrpg-go/pkg/text"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

const (
	CommAndTest    = 1
	SetWorldLevel  = 1001
	GetPlayerDb    = 1002
	Status         = 1003
	Give           = 1004
	GiveRelic      = 1005
	AddMail        = 1006
	SetJumpMission = 1007
)

var SetMap = map[string]int{
	"unlock": Unlock,
}

const (
	Unlock = iota + 1
)

type CommandAll interface {
	getCommand(list []string, l spb.LanguageType) (CommandAll, error)
}

func GetCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 1 {
		return nil, errors.New(text.GetTextByL(l, 27))
	}
	var c CommandAll
	switch list[0] {
	case "/help":
		return getHelp(l)
	case "/give":
		c = new(CommandGive)
		return c.getCommand(list[1:], l)
	case "/relic":
		c = new(CommandRelic)
		return c.getCommand(list[1:], l)
	case "/equipment":
		c = new(CommandEquipment)
		return c.getCommand(list[1:], l)
	case "/avatar":
		c = new(CommandAvatar)
		return c.getCommand(list[1:], l)
	case "/del":
		c = new(CommandDel)
		return c.getCommand(list[1:], l)
	case "/set":
		c = new(CommandSet)
		return c.getCommand(list[1:], l)
	case "/lua":
		c = new(CommandLua)
		return c.getCommand(list[1:], l)
	case "/rogue":
		return new(CommandRogue).getCommand(list[1:], l)
	case "/info":
		return new(CommandStatus).getCommand(nil, l)
	case "/mission":

	}
	return nil, errors.New(text.GetTextByL(l, 27))
}

func getHelp(l spb.LanguageType) (CommandAll, error) {
	return nil, errors.New(fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		text.GetTextByL(l, 29),
		text.GetTextByL(l, 31),
		text.GetTextByL(l, 37),
		text.GetTextByL(l, 38),
		text.GetTextByL(l, 39),
		text.GetTextByL(l, 40),
		text.GetTextByL(l, 41),
		text.GetTextByL(l, 49),
		text.GetTextByL(l, 55),
		text.GetTextByL(l, 61),
		text.GetTextByL(l, 69),
		text.GetTextByL(l, 83)))
}

type CommandGive struct {
	Uid     uint32   // 作用玩家
	Type    GiveType // 类型
	ItemId  uint32   // 物品id
	ItemNum uint32   // 物品数量
}

type GiveType string

var GiveTypeMap = map[string]GiveType{
	"":          GiveTypeNone,
	"all":       GiveTypeAll,
	"item":      GiveTypeItem,
	"relic":     GiveTypeRelic,
	"equipment": GiveTypeEquipment,
	"avatar":    GiveTypeAvatar,
	"icon":      GiveTypeIcon,
	"book":      GiveTypeBook,
	"disk":      GiveTypeDisk,
	"food":      GiveTypeFood,
	"formula":   GiveTypeFormula,
	"chat":      GiveTypeChat,
	"theme":     GiveTypeTheme,
	"mission":   GiveTypeMission,
	"gift":      GiveTypeGift,
	"pam":       GiveTypePam,
	"pet":       GiveTypePet,
}

const (
	GiveTypeNone      GiveType = ""
	GiveTypeAll       GiveType = "all"
	GiveTypeItem      GiveType = "item"
	GiveTypeRelic     GiveType = "relic"
	GiveTypeEquipment GiveType = "equipment"
	GiveTypeAvatar    GiveType = "avatar"
	GiveTypeIcon      GiveType = "icon"
	GiveTypeBook      GiveType = "book"
	GiveTypeDisk      GiveType = "disk"
	GiveTypeFood      GiveType = "food"
	GiveTypeFormula   GiveType = "formula"
	GiveTypeChat      GiveType = "chat"
	GiveTypeTheme     GiveType = "theme"
	GiveTypeMission   GiveType = "mission"
	GiveTypeGift      GiveType = "gift"
	GiveTypePam       GiveType = "pam"
	GiveTypePet       GiveType = "pet"
)

func (c *CommandGive) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) == 0 {
		return nil, errors.New(text.GetTextByL(l, 31))
	}
	comm := &CommandGive{
		ItemId:  0,
		ItemNum: 0,
	}
	var ok bool
	comm.Type, ok = GiveTypeMap[list[0]]
	if !ok {
		id := s2U32(list[0])
		if id == 0 {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 33), list[0]))
		}
		comm.ItemId = id
	}
	if len(list) < 2 {
		return comm, nil
	}
	comm.ItemNum = s2U32(list[1])
	if comm.ItemNum == 0 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 34), list[0], list[1]))
	}

	return comm, nil
}

type CommandRelic struct {
	Uid     uint32
	IsAll   bool
	RelicId uint32
	Num     uint32
	Main    uint32
	Sub     map[uint32]uint32
	Level   uint32
}

func (c *CommandRelic) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 1 {
		return nil, errors.New(text.GetTextByL(l, 49))
	}
	comm := &CommandRelic{
		IsAll:   false,
		RelicId: 0,
		Num:     1,
		Main:    0,
		Sub:     nil,
		Level:   1,
	}
	if list[0] == "all" {
		comm.IsAll = true
		if len(list) >= 2 {
			comm.Level = s2U32(list[1])
		}
		return comm, nil
	}
	if comm.RelicId = s2U32(list[0]); comm.RelicId == 0 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 50), list[0]))
	}
	if len(list) < 2 {
		return comm, nil
	}
	if comm.Num = s2U32(list[1]); comm.Num < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 51), list[0], list[1]))
	}
	if len(list) < 3 {
		return comm, nil
	}
	if comm.Level = s2U32(list[2]); comm.Level < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 52), list[0], list[1], list[2]))
	}
	if len(list) < 4 {
		return comm, nil
	}
	if comm.Main = s2U32(list[3]); comm.Main < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 53), list[0], list[1], list[2], list[3]))
	}
	if len(list) < 5 {
		return comm, nil
	}
	comm.Sub = GetRelicSub(list[5])
	return comm, nil
}

type CommandEquipment struct {
	Uid         uint32
	IsAll       bool
	EquipmentId uint32
	Num         uint32
	Level       uint32
	Rank        uint32
}

func (c *CommandEquipment) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 1 {
		return nil, errors.New(text.GetTextByL(l, 55))
	}
	comm := &CommandEquipment{
		IsAll:       false,
		EquipmentId: 0,
		Num:         0,
		Level:       1,
		Rank:        1,
	}
	if list[0] == "all" {
		comm.IsAll = true
	} else {
		if comm.EquipmentId = s2U32(list[0]); comm.EquipmentId == 0 {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 56), list[0]))
		}
	}
	if len(list) < 2 {
		return comm, nil
	}
	if comm.Num = s2U32(list[1]); comm.Num < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 57), list[0], list[1]))
	}
	if len(list) < 3 {
		return comm, nil
	}
	if comm.Level = s2U32(list[2]); comm.Level < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 58), list[0], list[1], list[2]))
	}
	if len(list) < 4 {
		return comm, nil
	}
	if comm.Rank = s2U32(list[3]); comm.Rank < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 59), list[0], list[1], list[2], list[3]))
	}
	return comm, nil
}

type CommandDel struct {
	Uid     uint32
	DelType DelType
	IsAll   bool
	Id      uint32
	Num     uint32
}

type DelType = string

const (
	DelTypeUnknown   DelType = ""
	DelTypeItem      DelType = "item"
	DelTypeRelic     DelType = "relic"
	DelTypeEquipment DelType = "equipment"
)

func isValidDelType(d DelType) bool {
	switch d {
	case DelTypeItem, DelTypeRelic, DelTypeEquipment:
		return true
	default:
		return false
	}
}

func (c *CommandDel) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 1 {
		return nil, errors.New(text.GetTextByL(l, 61))
	}
	comm := &CommandDel{
		IsAll:   false,
		DelType: DelTypeUnknown,
		Id:      0,
		Num:     0,
	}
	if list[0] == "all" {
		comm.IsAll = true
		return comm, nil
	}
	comm.DelType = list[0]
	if comm.Id = s2U32(list[0]); comm.Id != 0 {
		comm.DelType = DelTypeItem
		if len(list) < 2 {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 66), list[0], "?"))
		}
		if comm.Num = s2U32(list[1]); comm.Num < 1 {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 66), list[0], list[1]))
		}
		return comm, nil
	}
	if !isValidDelType(comm.DelType) {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 65), comm.DelType))
	}
	if len(list) < 2 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 67), list[0], "?"))
	}
	if list[1] == "all" {
		comm.IsAll = true
		return comm, nil
	}
	if comm.Id = s2U32(list[1]); comm.Id < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 67), list[0], list[1]))
	}

	return comm, nil
}

type CommandAvatar struct {
	Uid      uint32
	IsAll    bool
	IsMax    bool
	Type     CommAvatarType
	AvatarId uint32 // 角色
	Rank     uint32 // 命座
	Level    uint32 // 等级
	Skill    bool
}

type CommAvatarType = string

const (
	CommAvatarTypeAdd   CommAvatarType = "add"
	CommAvatarTypeDel   CommAvatarType = "del"
	CommAvatarTypeBuild CommAvatarType = "build"
)

func (c *CommandAvatar) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 2 {
		return nil, errors.New(text.GetTextByL(l, 69))
	}
	comm := &CommandAvatar{
		IsAll:    false,
		IsMax:    false,
		Skill:    false,
		Type:     list[0],
		AvatarId: 0,
		Rank:     0,
		Level:    1,
	}
	switch comm.Type {
	case CommAvatarTypeAdd, CommAvatarTypeDel, CommAvatarTypeBuild:
	default:
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 70), list[0]))
	}
	if list[1] == "all" {
		comm.IsAll = true
	} else {
		if comm.AvatarId = s2U32(list[1]); comm.AvatarId < 1 {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 71), list[0], list[1]))
		}
	}
	if len(list) < 3 {
		return comm, nil
	}
	if list[2] == "max" {
		comm.IsMax = true
		return comm, nil
	}
	if comm.Level = s2U32(list[2]); comm.Level < 1 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 72), list[0], list[1], list[2]))
	}
	if len(list) < 4 {
		return comm, nil
	}
	if comm.Rank = s2U32(list[3]); comm.Rank < 0 {
		comm.Rank = 0
	} else if comm.Rank > 6 {
		comm.Rank = 6
	}
	if len(list) < 5 {
		return comm, nil
	}
	if s2U32(list[4]) > 0 {
		comm.Skill = true
	}

	return comm, nil
}

type SetType = string

const (
	SetTypeWorldLevel  SetType = "WorldLevel"  // 设置世界等级
	SetTypePlayerLevel SetType = "PlayerLevel" // 设置账号等级
	SetTypeJumpMission SetType = "JumpMission" // 设置是否关闭任务进程
	SetTypeLanguage    SetType = "Language"    // 设置语言
	SetTypeMainAvatar  SetType = "MainAvatar"  // 设置主角性别
)

type CommandSet struct {
	Uid      uint32           // 作用玩家
	SetType  SetType          // 设置类型
	Sub1     uint32           // 附加信息
	Language spb.LanguageType // 语言设置附加信息
}

func (c *CommandSet) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 2 {
		return nil, errors.New(text.GetTextByL(l, 35))
	}
	comm := &CommandSet{
		SetType:  list[0],
		Sub1:     0,
		Language: spb.LanguageType_LANGUAGE_SC,
	}
	c2 := list[1]
	switch comm.SetType {
	case SetTypeWorldLevel:
		if wl := s2U32(c2); wl < 0 || wl > 6 {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 37), wl))
		} else {
			comm.Sub1 = wl
		}
	case SetTypePlayerLevel:
		if pl := s2U32(c2); pl < 1 {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 38), pl))
		} else {
			comm.Sub1 = pl
		}
	case SetTypeJumpMission:
		if bl := s2U32(c2); bl == 0 {
			comm.Sub1 = bl
		} else if bl == 1 {
			comm.Sub1 = bl
		} else {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 39), bl))
		}
	case SetTypeLanguage:
		l = text.GetLanguageTypeByS(c2)
		if l == spb.LanguageType_LANGUAGE_NONE {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 40), c2))
		}
		comm.Language = l
	case SetTypeMainAvatar:
		if ma := s2U32(c2); ma == 0 {
			comm.Sub1 = ma
		} else if ma == 1 {
			comm.Sub1 = ma
		} else {
			return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 41), ma))
		}
	default:
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 36), list[0]))
	}
	return comm, nil
}

type CommandLua struct {
	Uid  uint32
	Data []byte
}

func (c *CommandLua) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 1 {
		return nil, errors.New(text.GetTextByL(l, 80))
	}
	data := lua.GetLaLua(list[0])
	if data == nil {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 79), list[0]))
	}
	return &CommandLua{Uid: c.Uid, Data: data}, nil
}

type CommandRogue struct {
	Uid  uint32
	Type uint32 // 类型
	Set  int    // 设置类型
}

var RogueType = map[string]uint32{
	"all":      RogueTypeAll,
	"handbook": RogueTypeHandbook,
	"quest":    RogueTypeQuest,
}

const (
	RogueTypeAll      = 1
	RogueTypeHandbook = 2
	RogueTypeQuest    = 101
)

func (c *CommandRogue) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	if len(list) < 2 {
		return nil, errors.New(text.GetTextByL(l, 83))
	}
	comm := &CommandRogue{}
	comm.Type = RogueType[list[0]]
	if comm.Type == 0 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 84), list[0]))
	}
	comm.Set = SetMap[list[1]]
	if comm.Set == 0 {
		return nil, errors.New(fmt.Sprintf(text.GetTextByL(l, 85), list[0], list[1]))
	}

	return comm, nil
}

type CommandStatus struct{}

func (c *CommandStatus) getCommand(list []string, l spb.LanguageType) (CommandAll, error) {
	return c, nil
}

func s2U32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

func GetRelicSub(sub string) map[uint32]uint32 {
	re := regexp.MustCompile(`\[(\d+):(\d+)\]`)
	matches := re.FindAllStringSubmatch(sub, -1)

	result := make(map[uint32]uint32)
	for _, match := range matches {
		key := s2U32(match[1])
		value := s2U32(match[2])
		result[key] = value
	}
	return result
}
