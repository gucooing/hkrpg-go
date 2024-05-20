package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 添加物品
func (g *GamePlayer) GmGive(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GmGive)
	if req.PlayerUid == 0 {
		return
	}
	itemConf := gdconf.GetItemConfigMap()
	if req.GiveAll {
		var pileItem []*Material
		// add avatar
		for _, avatar := range itemConf.Avatar {
			// 过滤非live角色
			if avatar.ID/1000 != 1 {
				continue
			}
			g.AddAvatar(avatar.ID)
		}
		// add playerIcon
		var playerIconList []uint32
		for _, playerIcon := range itemConf.AvatarPlayerIcon {
			playerIconList = append(playerIconList, playerIcon.ID)
		}
		g.GetItem().HeadIcon = playerIconList
		// add rank
		for _, rank := range itemConf.AvatarRank {
			pileItem = append(pileItem, &Material{
				Tid: rank.ID,
				Num: 6,
			})
		}
		// add equipment
		for _, equipment := range itemConf.Equipment {
			g.AddEquipment(equipment.ID)
		}
		// add item
		for _, item := range itemConf.Item {
			pileItem = append(pileItem, &Material{
				Tid: item.ID,
				Num: 99999,
			})
		}
		// add relic
		for _, relic := range itemConf.Relic {
			g.AddRelic(relic.ID)
		}
		g.AddMaterial(pileItem)
		// g.ScenePlaneEventScNotify(pileItem)
	} else {
		var pileItem []*Material
		for _, item := range itemConf.Item {
			if item.ID == req.ItemId {
				pileItem = append(pileItem, &Material{
					Tid: item.ID,
					Num: req.ItemCount,
				})
				g.AddMaterial(pileItem)
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
				pileItem = append(pileItem, &Material{
					Tid: avatar.ID,
					Num: req.ItemCount,
				})
				g.AddMaterial(pileItem)
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

// 设置世界等级
func (g *GamePlayer) GmWorldLevel(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GmWorldLevel)
	g.SetWorldLevel(req.WorldLevel)
	// 账号状态通知
	g.PlayerPlayerSyncScNotify()
}

// 清空背包
func (g *GamePlayer) DelItem(payloadMsg pb.Message) {
	g.BasicBin.Item = &spb.Item{
		RelicMap:     make(map[uint32]*spb.Relic),
		EquipmentMap: make(map[uint32]*spb.Equipment),
		MaterialMap:  make(map[uint32]uint32),
		HeadIcon:     make([]uint32, 0),
	}
	g.BasicBin.Item.MaterialMap[11] = 240
}

func (g *GamePlayer) GmMaxCurAvatar(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.MaxCurAvatar)
	if req.All {
		bin := g.GetAvatar()
		if bin == nil {
			return
		}
		for _, db := range bin.AvatarList {
			g.SetAvatarMaxByDb(db)
		}
	} else {
		var db *spb.AvatarBin
		db = g.GetAvatarBinById(req.AvatarId)
		if db == nil {
			db = g.GetCurAvatar()
		}
		g.SetAvatarMaxByDb(db)
	}
}

func (g *GamePlayer) SetAvatarMaxByDb(db *spb.AvatarBin) {
	if db == nil {
		return
	}
	db.Level = 80                              // 80级
	db.PromoteLevel = 6                        // 突破等级
	db.Rank = 6                                // 六命
	db.Hp = 10000                              // 满血
	db.SpBar.CurSp = 10000                     // 满能量
	g.SetAvatarMakSkillByAvatarId(db.AvatarId) // 技能满级
	// 通知角色信息
	g.AvatarPlayerSyncScNotify(db.AvatarId)
}
