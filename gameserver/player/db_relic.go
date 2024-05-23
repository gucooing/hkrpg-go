package player

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) AddRelic(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relicConf := gdconf.GetRelicById(strconv.Itoa(int(tid)))
	mainAffixConf := gdconf.GetRelicMainAffixConfigById(relicConf.MainAffixGroup)

	relic := &spb.Relic{
		Tid:               tid,
		UniqueId:          uniqueId,
		Exp:               0,
		Level:             0,
		MainAffixId:       mainAffixConf.AffixID,
		RelicAffix:        make(map[uint32]*spb.RelicAffix),
		BaseAvatarId:      0,
		IsProtected:       false,
		MainAffixProperty: mainAffixConf.Property,
	}

	baseSubAffixes := math.Min(math.Max(float64(relicConf.Type-2), 0), 3)
	addSubAffixes := rand.Intn(2) + int(baseSubAffixes)
	relicAffix := make(map[uint32]*spb.RelicAffix)
	g.addRelicAffix(&addRelicAffix{
		addSubAffixes:     addSubAffixes,
		mainAffixProperty: mainAffixConf.Property,
		subAffixGroup:     relicConf.SubAffixGroup,
		relicAffix:        relicAffix,
	})
	relic.RelicAffix = relicAffix

	g.GetItem().RelicMap[uniqueId] = relic
	g.RelicPlayerSyncScNotify(uniqueId)
}

func (g *GamePlayer) AddBtRelic(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relicConf := gdconf.GetRelicById(strconv.Itoa(int(tid)))
	mainAffixConf := gdconf.GetRelicMainAffixConfigById(relicConf.MainAffixGroup)

	relic := &spb.Relic{
		Tid:               tid,
		UniqueId:          uniqueId,
		Exp:               0,
		Level:             0,
		MainAffixId:       mainAffixConf.AffixID,
		RelicAffix:        make(map[uint32]*spb.RelicAffix),
		BaseAvatarId:      0,
		IsProtected:       false,
		MainAffixProperty: mainAffixConf.Property,
	}

	relicAffix := make(map[uint32]*spb.RelicAffix)
	g.addRelicAffix(&addRelicAffix{
		addSubAffixes:     400,
		mainAffixProperty: mainAffixConf.Property,
		subAffixGroup:     relicConf.SubAffixGroup,
		relicAffix:        relicAffix,
	})
	relic.RelicAffix = relicAffix

	g.GetItem().RelicMap[uniqueId] = relic
	g.RelicPlayerSyncScNotify(uniqueId)
}

type addRelicAffix struct {
	addSubAffixes     int                        // 添加词条数
	mainAffixProperty string                     // 主词条效果
	subAffixGroup     uint32                     // 副词条随机库id
	relicAffix        map[uint32]*spb.RelicAffix // 副词条内存
}

func (g *GamePlayer) addRelicAffix(str *addRelicAffix) {
	for i := 0; i < str.addSubAffixes; {
		if len(str.relicAffix) >= 4 {
			randIndex := rand.Intn(len(str.relicAffix))
			randKey := uint32(0)
			for key := range str.relicAffix {
				if randIndex == 0 {
					randKey = key
					break
				}
				randIndex--
			}
			str.relicAffix[randKey].Cnt++
			i++
		} else {
			affixConf := gdconf.GetRelicSubAffixConfigById(str.subAffixGroup)
			if affixConf == nil {
				return
			}
			if affixConf.Property == str.mainAffixProperty {
				continue
			}
			if ra, ok := str.relicAffix[affixConf.AffixID]; ok {
				ra.Cnt++
			} else {
				str.relicAffix[affixConf.AffixID] = &spb.RelicAffix{
					AffixId: affixConf.AffixID,
					Cnt:     1,
					Step:    0,
				}
			}
			i++
		}
	}
}

func (g *GamePlayer) getRelicDbById(uniqueId uint32) *spb.Relic {
	if relicDb, ok := g.GetItem().RelicMap[uniqueId]; !ok {
		return nil
	} else {
		return relicDb
	}
}

/*********************************************接口方法******************************************/

func (g *GamePlayer) GetProtoRelicById(uniqueId uint32) *proto.Relic {
	if relicDb, ok := g.GetItem().RelicMap[uniqueId]; !ok {
		return nil
	} else {
		relic := &proto.Relic{
			Tid:          relicDb.Tid,
			SubAffixList: make([]*proto.RelicAffix, 0),
			BaseAvatarId: relicDb.BaseAvatarId,
			UniqueId:     relicDb.UniqueId,
			Level:        relicDb.Level,
			IsProtected:  relicDb.IsProtected,
			MainAffixId:  relicDb.MainAffixId,
			Exp:          relicDb.Exp,
		}
		for _, subAffixList := range relicDb.RelicAffix {
			relicAffix := &proto.RelicAffix{
				AffixId: subAffixList.AffixId,
				Cnt:     subAffixList.Cnt,
				Step:    subAffixList.Step,
			}
			relic.SubAffixList = append(relic.SubAffixList, relicAffix)
		}

		return relic
	}
}

func (g *GamePlayer) GetProtoBattleRelicById(uniqueId uint32) *proto.BattleRelic {
	if relicDb, ok := g.GetItem().RelicMap[uniqueId]; !ok {
		return nil
	} else {
		relic := &proto.BattleRelic{
			Id:           relicDb.Tid,
			SubAffixList: make([]*proto.RelicAffix, 0),
			UniqueId:     relicDb.UniqueId,
			Level:        relicDb.Level,
			MainAffixId:  relicDb.MainAffixId,
		}
		for _, subAffixList := range relicDb.RelicAffix {
			relicAffix := &proto.RelicAffix{
				AffixId: subAffixList.AffixId,
				Cnt:     subAffixList.Cnt,
				Step:    subAffixList.Step,
			}
			relic.SubAffixList = append(relic.SubAffixList, relicAffix)
		}

		return relic
	}
}

func (g *GamePlayer) RelicScenePlaneEventScNotify(uniqueId uint32) {
	relicItme := g.GetProtoRelicById(uniqueId)
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
