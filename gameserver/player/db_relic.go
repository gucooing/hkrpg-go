package player

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
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
	relic.RelicAffix = g.addRelicAffix(addSubAffixes, mainAffixConf.Property, relicConf.SubAffixGroup, relicAffix)

	g.GetItem().RelicMap[uniqueId] = relic
	g.RelicPlayerSyncScNotify(uniqueId)
}

func (g *GamePlayer) addRelicAffix(addSubAffixes int, mainAffixProperty string, subAffixGroup uint32, relicAffix map[uint32]*spb.RelicAffix) map[uint32]*spb.RelicAffix {
	for i := 0; i < addSubAffixes; {
		if len(relicAffix) >= 4 {
			randIndex := rand.Intn(len(relicAffix))
			randKey := uint32(0)
			for key := range relicAffix {
				if randIndex == 0 {
					randKey = key
					break
				}
				randIndex--
			}
			relicAffix[randKey].Cnt++
			i++
		} else {
			affixConf := gdconf.GetRelicSubAffixConfigById(subAffixGroup)
			if affixConf == nil {
				return nil
			}
			if affixConf.Property == mainAffixProperty {
				continue
			}
			if ra, ok := relicAffix[affixConf.AffixID]; ok {
				ra.Cnt++
			} else {
				relicAffix[affixConf.AffixID] = &spb.RelicAffix{
					AffixId: affixConf.AffixID,
					Cnt:     1,
					Step:    0,
				}
			}
			i++
		}
	}
	return relicAffix
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
