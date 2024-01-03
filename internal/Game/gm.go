package Game

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	gmpb "github.com/gucooing/hkrpg-go/protocol/gmpb"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *Game) GmGive(payloadMsg pb.Message) {
	req := payloadMsg.(*gmpb.GmGive)
	itemConf := gdconf.GetItemConfigMap()
	if req.GiveAll {
		// add avatar
		for _, avatar := range itemConf.Avatar {
			// 过滤非live角色
			if avatar.ID/1000 != 1 {
				continue
			}
			g.AddAvatar(avatar.ID)
			time.Sleep(10 * time.Millisecond)
		}
		// add playerIcon
		var playerIconList []uint32
		for _, playerIcon := range itemConf.AvatarPlayerIcon {
			playerIconList = append(playerIconList, playerIcon.ID)
		}
		g.Player.DbItem.HeadIcon = playerIconList
		// add rank
		for _, rank := range itemConf.AvatarRank {
			g.AddMaterial(rank.ID, 6)
			time.Sleep(10 * time.Millisecond)
		}
		// add equipment
		for _, equipment := range itemConf.Equipment {
			g.AddEquipment(equipment.ID)
			time.Sleep(10 * time.Millisecond)
		}
		// add item
		for _, item := range itemConf.Item {
			g.AddMaterial(item.ID, 99999)
			time.Sleep(10 * time.Millisecond)
		}
		// add relic
		for _, relic := range itemConf.Relic {
			g.AddRelic(relic.ID)
			time.Sleep(10 * time.Millisecond)
		}
	} else {
		for _, item := range itemConf.Item {
			if item.ID == req.ItemId {
				g.AddMaterial(item.ID, req.ItemCount)
				return
			}
		}
		for _, avatar := range itemConf.Avatar {
			if avatar.ID == req.ItemId {
				g.AddAvatar(avatar.ID)
				return
			}
		}
		for _, avatar := range itemConf.AvatarRank {
			if avatar.ID == req.ItemId {
				g.AddMaterial(avatar.ID, req.ItemCount)
				return
			}
		}
		for _, avatar := range itemConf.AvatarPlayerIcon {
			if avatar.ID == req.ItemId {
				g.AddHeadIcon(avatar.ID)
				return
			}
		}
		for _, equipment := range itemConf.Equipment {
			if equipment.ID == req.ItemId {
				g.AddEquipment(equipment.ID)
				return
			}
		}
		for _, relic := range itemConf.Relic {
			if relic.ID == req.ItemId {
				g.AddRelic(relic.ID)
				return
			}
		}
	}
}

func (g *Game) ScenePlaneEventScNotify(id, num uint32) {
	// 通知客户端增加了物品
	notify := &proto.ScenePlaneEventScNotify{
		GetItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	item := &proto.Item{
		ItemId:      id,
		Level:       0,
		Num:         num,
		MainAffixId: 0,
		Rank:        0,
		Promotion:   0,
		UniqueId:    0,
	}
	notify.GetItemList.ItemList = append(notify.GetItemList.ItemList, item)
	g.Send(cmd.ScenePlaneEventScNotify, notify)
}

func (g *Game) RelicScenePlaneEventScNotify(uniqueId uint32) {
	relicItme := g.Player.DbItem.RelicMap[uniqueId]
	// 通知客户端增加了物品
	notify := &proto.ScenePlaneEventScNotify{
		GetItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	item := &proto.Item{
		ItemId:      relicItme.Tid,
		Level:       relicItme.Level,
		Num:         1,
		MainAffixId: relicItme.MainAffixId,
		Rank:        0,
		Promotion:   0,
		UniqueId:    relicItme.UniqueId,
	}
	notify.GetItemList.ItemList = append(notify.GetItemList.ItemList, item)
	g.Send(cmd.ScenePlaneEventScNotify, notify)
}

func (g *Game) GmWorldLevel(payloadMsg pb.Message) {
	req := payloadMsg.(*gmpb.GmWorldLevel)

	g.Player.WorldLevel = req.WorldLevel

	// 账号状态通知
	g.PlayerPlayerSyncScNotify()
}
