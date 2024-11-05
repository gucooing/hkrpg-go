package player

import (
	"fmt"
	"strings"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

// msg.CommandList
func (g *GamePlayer) EnterCommand(msg Msg) {
	reqMessageTextList := text.GetTextByL(g.GetPd().GetLanguageType(), 28)
	lists := strings.Split(msg.Command, " ")
	c, err := constant.GetCommand(lists, g.GetPd().GetLanguageType())
	if err != nil {
		reqMessageTextList = err.Error()
	}
	if c != nil {
		switch x := c.(type) {
		case *constant.CommandGive:
			reqMessageTextList = g.commandGive(x)
		case *constant.CommandSet:
			reqMessageTextList = g.commandSet(x)
		case *constant.CommandRelic:
			reqMessageTextList = g.commandRelic(x)
		case *constant.CommandEquipment:
			reqMessageTextList = g.commandEquipment(x)
		case *constant.CommandAvatar:
			reqMessageTextList = g.commandAvatar(x)
		case *constant.CommandDel:
			reqMessageTextList = g.commandDel(x)
		case *constant.CommandLua:
			reqMessageTextList = g.commandLua(x)
		}
	}
	for _, reqMessageText := range strings.Split(reqMessageTextList, "\n") {
		notify := &proto.RevcMsgScNotify{
			SourceUid:   0,
			TargetUid:   g.Uid,
			MessageText: reqMessageText,
			MessageType: proto.MsgType_MSG_TYPE_CUSTOM_TEXT,
			ChatType:    proto.ChatType_CHAT_TYPE_PRIVATE,
		}
		g.Send(cmd.RevcMsgScNotify, notify)
	}
}

func (g *GamePlayer) commandGive(c *constant.CommandGive) string {
	addItem := model.NewAddItem(nil)
	if !g.GetPd().GetIsProficientPlayer() &&
		!g.GetPd().GetIsJumpMission() {
		return text.GetTextByL(g.GetPd().GetLanguageType(), 48)
	}
	if c.IsAll {
		addItem.PileItem = g.allGive()
	} else {
		if gdconf.GetItemConfigById(c.ItemId) == nil {
			return fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 32), c.ItemId)
		}
		addItem.PileItem = append(addItem.PileItem, &model.Material{
			Tid: c.ItemId,
			Num: c.ItemNum,
		})
	}

	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.AllScenePlaneEventScNotify(addItem.MaterialList)
	return fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 30), "give")
}

func (g *GamePlayer) allGive() []*model.Material {
	var pileItem []*model.Material
	itemConf := gdconf.GetItemConfig()
	avatarConf := gdconf.GetAvatarDataMap()
	// add avatar
	for _, avatar := range avatarConf {
		x := avatar.AvatarId / 1000
		if x != 1 && x != 8 {
			continue
		}
		pileItem = append(pileItem, &model.Material{
			Tid: avatar.AvatarId,
			Num: 1,
		})
	}
	// add playerIcon
	for _, item := range itemConf.AvatarPlayerIcon {
		pileItem = append(pileItem, &model.Material{
			Tid: item.ID,
			Num: 1,
		})
	}
	// add rank
	for _, rank := range itemConf.AvatarRank {
		pileItem = append(pileItem, &model.Material{
			Tid: rank.ID,
			Num: 6,
		})
	}
	// add equipment
	for _, equipment := range itemConf.Equipment {
		pileItem = append(pileItem, &model.Material{
			Tid: equipment.ID,
			Num: 1,
		})
	}
	// add item
	for _, item := range itemConf.Item {
		pileItem = append(pileItem, &model.Material{
			Tid: item.ID,
			Num: 9999999,
		})
	}
	// add disk
	for _, item := range itemConf.Disk {
		pileItem = append(pileItem, &model.Material{
			Tid: item.ID,
			Num: 1,
		})
	}
	// add relic
	for _, relic := range itemConf.Relic {
		pileItem = append(pileItem, &model.Material{
			Tid: relic.ID,
			Num: 1,
		})
	}
	return pileItem
}

func (g *GamePlayer) commandSet(c *constant.CommandSet) string {
	addItem := model.NewAddItem(nil)
	res := text.GetTextByL(g.GetPd().GetLanguageType(), 47)

	switch c.SetType {
	case constant.SetTypeWorldLevel:
		g.GetPd().SetWorldLevel(c.Sub1)
		res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 42), g.GetPd().GetWorldLevel())
	case constant.SetTypePlayerLevel:
		g.GetPd().SetPlayerLevel(c.Sub1)
		res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 43), c.Sub1)
	case constant.SetTypeJumpMission:
		if c.Sub1 == 0 {
			g.GetPd().SeIsJumpMission(false)
		} else if c.Sub1 == 1 {
			g.GetPd().SeIsJumpMission(true)
		}
		res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 44), g.GetPd().GetIsJumpMission())
	case constant.SetTypeLanguage:
		g.GetPd().SetLanguageType(c.Language)
		res = text.GetTextByL(g.GetPd().GetLanguageType(), 45)
	case constant.SetTypeMainAvatar:
		if c.Sub1 == 0 {
			g.GetPd().SetGender(spb.Gender_GenderMan)
			g.GetPd().SetMultiPathAvatar(8001)
		} else if c.Sub1 == 1 {
			g.GetPd().SetGender(spb.Gender_GenderWoman)
			g.GetPd().SetMultiPathAvatar(8002)
		}
		addItem.AllSync.AvatarList = append(addItem.AllSync.AvatarList, 8001)
		res = text.GetTextByL(g.GetPd().GetLanguageType(), 46)
	}

	g.AllPlayerSyncScNotify(addItem.AllSync)
	return res
}

func (g *GamePlayer) commandRelic(c *constant.CommandRelic) string {
	addItem := model.NewAddItem(nil)
	res := fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 30), "equipment")

	if c.IsAll {
		for id := range gdconf.GetItemRelic() {
			addItem.PileItem = append(addItem.PileItem, &model.Material{
				Tid: id,
				Num: c.Num,
			})
		}
	} else {
		var i uint32 = 0
		for ; i < c.Num; i++ {
			if uniqueId := g.GetPd().AddRelic(c.RelicId, c.Level, c.Main, c.Sub); uniqueId == 0 {
				res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 54), c.RelicId)
				break
			} else {
				addItem.AllSync.RelicList = append(addItem.AllSync.RelicList, uniqueId)
			}
		}
	}

	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	return res
}

func (g *GamePlayer) commandEquipment(c *constant.CommandEquipment) string {
	addItem := model.NewAddItem(nil)
	res := fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 30), "relic")

	if c.IsAll {
		for id := range gdconf.GetItemEquipment() {
			addItem.PileItem = append(addItem.PileItem, &model.Material{
				Tid: id,
				Num: c.Num,
			})
		}
	} else {
		var i uint32 = 0
		for ; i < c.Num; i++ {
			if uniqueId := g.GetPd().AddEquipment(c.EquipmentId, c.Level, c.Rank); uniqueId == 0 {
				res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 60), c.EquipmentId)
				break
			} else {
				addItem.AllSync.EquipmentList = append(addItem.AllSync.EquipmentList, uniqueId)
			}
		}
	}

	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	return res
}

func (g *GamePlayer) commandDel(c *constant.CommandDel) string {
	allSync := model.NewAllPlayerSync()
	res := fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 64), c.DelType)
	db := g.GetPd().GetItem()

	switch c.DelType {
	case constant.DelTypeUnknown:
		if c.IsAll {
			g.GetPd().GetBasicBin().Item = model.NewItem()
			res = text.GetTextByL(g.GetPd().GetLanguageType(), 62)
		}
	case constant.DelTypeEquipment:
		if c.IsAll {
			db.EquipmentMap = make(map[uint32]*spb.Equipment)
			res = text.GetTextByL(g.GetPd().GetLanguageType(), 62)
		} else {
			if db.EquipmentMap[c.Id] != nil {
				delete(db.EquipmentMap, c.Id)
				allSync.DelEquipmentList = append(allSync.DelEquipmentList, c.Id)
				res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 63), string(c.DelType)+string(c.Id))
			}
		}
	case constant.DelTypeRelic:
		if c.IsAll {
			db.RelicMap = make(map[uint32]*spb.Relic)
			res = text.GetTextByL(g.GetPd().GetLanguageType(), 62)
		} else {
			if db.RelicMap[c.Id] != nil {
				delete(db.RelicMap, c.Id)
				allSync.DelRelicList = append(allSync.DelRelicList, c.Id)
				res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 63), string(c.DelType)+string(c.Id))
			}
		}
	case constant.DelTypeItem:
		if c.IsAll {
			db.MaterialMap = make(map[uint32]uint32)
			res = text.GetTextByL(g.GetPd().GetLanguageType(), 62)
		} else {
			if n := db.MaterialMap[c.Id]; n >= c.Num {
				db.MaterialMap[c.Id] = n - c.Num
			} else {
				db.MaterialMap[c.Id] = 0
			}
			res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 68), c.Id, db.MaterialMap[c.Id])
			allSync.MaterialList = append(allSync.MaterialList, c.Id)
		}
	}

	g.AllPlayerSyncScNotify(allSync)
	return res
}

func (g *GamePlayer) commandAvatar(c *constant.CommandAvatar) string {
	allSync := model.NewAllPlayerSync()
	res := fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 30), "avatar")

	switch c.Type {
	case constant.CommAvatarTypeAdd:
		if c.IsAll {
			allSync.AvatarList = gdconf.GetAvatarList()
		} else {
			if gdconf.GetAvatarDataById(c.AvatarId) == nil {
				res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 74), c.AvatarId)
			} else {
				allSync.AvatarList = append(allSync.AvatarList, c.AvatarId)
			}
		}
	case constant.CommAvatarTypeBuild:
		if c.IsAll {
			for id := range g.GetPd().GetAvatarList() {
				allSync.AvatarList = append(allSync.AvatarList, id)
			}
		} else {
			if g.GetPd().GetAvatarBinById(c.AvatarId) == nil {
				res = fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 73), c.AvatarId)
			} else {
				allSync.AvatarList = append(allSync.AvatarList, c.AvatarId)
			}
		}
	case constant.CommAvatarTypeDel:
		res = text.GetTextByL(g.GetPd().GetLanguageType(), 75)
	}

	for _, id := range allSync.AvatarList {
		db := g.GetPd().GetAvatarBinById(id)
		if db == nil {
			g.GetPd().AddAvatar(id)
			db = g.GetPd().GetAvatarBinById(id)
		}
		if db == nil {
			continue
		}
		// 提升等级
		if c.IsMax {
			g.GetPd().SetAvatarLevel(db, 80)
			g.GetPd().SetAvatarMultiPath(db, true, 6)
		} else {
			g.GetPd().SetAvatarLevel(db, c.Level)
			g.GetPd().SetAvatarMultiPath(db, c.Skill, c.Rank)
		}
		db.Hp = 10000
		db.SpBar.CurSp = db.SpBar.MaxSp
	}

	g.AllPlayerSyncScNotify(allSync)
	return res
}

func (g *GamePlayer) commandLua(c *constant.CommandLua) string {
	g.ClientDownloadDataScNotify(c.Data)
	return fmt.Sprintf(text.GetTextByL(g.GetPd().GetLanguageType(), 30), "lua")
}

/**********************************分割线*******************************/

func (g *GamePlayer) RecoverLine() {
	db := g.GetPd().GetCurLineUp()
	for _, a := range db.AvatarIdList {
		bin := g.GetPd().GetAvatarBinById(a.AvatarId)
		if bin != nil {
			bin.Hp = 10000
			bin.SpBar.CurSp = 10000
		}
	}
	g.SyncLineupNotify(db)
}
