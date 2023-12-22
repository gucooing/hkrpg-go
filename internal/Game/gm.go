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
	if req.GiveAll {
		for _, avatar := range gdconf.GetAvatarList() {
			if avatar == 0 {
				continue
			}
			// 过滤主角
			if avatar/100 == 80 {
				continue
			}
			g.AddAvatar(avatar)
			g.AddMaterial(avatar+10000, 6)
		}
		for _, equipment := range gdconf.GetEquipmentList() {
			if equipment == 0 {
				continue
			}
			g.AddEquipment(equipment)
			time.Sleep(10 * time.Millisecond)
		}
		for _, item := range gdconf.GetItemList() {
			if item == 0 {
				continue
			}
			g.AddMaterial(item, 99999)
			time.Sleep(10 * time.Millisecond)
		}
	} else {
		for _, item := range gdconf.GetItemList() {
			if item == 0 {
				continue
			}
			if item == req.ItemId {
				g.AddMaterial(item, req.ItemCount)
			}
		}
		for _, avatar := range gdconf.GetAvatarList() {
			if avatar == 0 {
				continue
			}
			if avatar == req.ItemId {
				g.AddAvatar(avatar)
				return
			}
		}
		for _, equipment := range gdconf.GetEquipmentList() {
			if equipment == 0 {
				continue
			}
			if equipment == req.ItemId {
				g.AddEquipment(equipment)
				return
			}
		}
		// 特殊物品(不再EquipmentList表中的物品)
		if req.ItemId/10000 == 1 {
			g.AddMaterial(req.ItemId, req.ItemCount)
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

func (g *Game) GmWorldLevel(payloadMsg pb.Message) {
	req := payloadMsg.(*gmpb.GmWorldLevel)

	g.Player.WorldLevel = req.WorldLevel

	// 账号状态通知
	g.PlayerPlayerSyncScNotify()
}
