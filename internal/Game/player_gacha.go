package Game

import (
	"math/rand"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetGachaInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetGachaInfoScRsp)
	rsp.GachaInfoList = make([]*proto.GachaInfo, 0)

	for _, bannerslist := range gdconf.GetBannersMap() {
		gachaInfoList := &proto.GachaInfo{
			HistoryUrl: "http://127.0.0.1:8080/api/gacha/history",      // 历史记录
			DetailUrl:  "https://www.bilibili.com/video/BV1X94y177QK/", // 卡池详情
			Featured:   bannerslist.RateUpItems5,                       // 五星up
			UpInfo:     bannerslist.RateUpItems4,                       // 四星up
			GachaId:    bannerslist.Id,
		}
		if bannerslist.GachaType == "Normal" {
			list := []uint32{1003, 1004, 1101, 1104, 1107, 1209, 1211}
			gachaInfoList.GachaCeiling = &proto.GachaCeiling{
				IsClaimed:  false,
				AvatarList: make([]*proto.GachaCeilingAvatar, 0),
				CeilingNum: 0,
			}
			for _, id := range list {
				avatarlist := &proto.GachaCeilingAvatar{
					RepeatedCnt: 0,
					AvatarId:    id,
				}
				gachaInfoList.GachaCeiling.AvatarList = append(gachaInfoList.GachaCeiling.AvatarList, avatarlist)
			}
		} else {
			gachaInfoList.BeginTime = bannerslist.BeginTime // 开始时间
			gachaInfoList.EndTime = bannerslist.EndTime     // 结束时间
		}

		rsp.GachaInfoList = append(rsp.GachaInfoList, gachaInfoList)
	}

	g.send(cmd.GetGachaInfoScRsp, rsp)
}

func (g *Game) HandleGetGachaCeilingCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.GetGachaCeilingCsReq, payloadMsg)
	req := msg.(*proto.GetGachaCeilingCsReq)

	logger.Info("", req)

	rsp := &proto.GetGachaCeilingScRsp{
		GachaType: req.GachaType,
	}
	list := []uint32{1003, 1004, 1101, 1104, 1107, 1209, 1211}
	rsp.GachaCeiling = &proto.GachaCeiling{
		IsClaimed:  false,
		AvatarList: make([]*proto.GachaCeilingAvatar, 0),
		CeilingNum: 299,
	}
	for _, id := range list {
		avatarlist := &proto.GachaCeilingAvatar{
			RepeatedCnt: 0,
			AvatarId:    id,
		}
		rsp.GachaCeiling.AvatarList = append(rsp.GachaCeiling.AvatarList, avatarlist)
	}

	g.send(cmd.GetGachaCeilingScRsp, rsp)
}

func (g *Game) DoGachaCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.DoGachaCsReq, payloadMsg)
	req := msg.(*proto.DoGachaCsReq)

	if req.GachaNum != 10 && req.GachaNum != 1 {
		return
	}

	rsp := &proto.DoGachaScRsp{
		GachaId:       req.GachaId,
		CeilingNum:    0,
		GachaItemList: make([]*proto.GachaItem, 0),
		GachaNum:      req.GachaNum,
	}
	for i := 0; i < int(req.GachaNum); i++ {
		id := g.GachaRandom(req.GachaId)
		gachaItemList := &proto.GachaItem{
			TransferItemList: nil,
			IsNew:            true,
			GachaItem:        nil,
			TokenItem:        &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		}
		itemList := &proto.Item{
			ItemId:      id,
			Level:       1,
			Num:         1,
			MainAffixId: 0,
			Rank:        1,
			Promotion:   0,
			UniqueId:    0,
		}
		gachaItemList.GachaItem = itemList
		gachaItemList.TokenItem.ItemList = append(gachaItemList.TokenItem.ItemList, itemList)

		rsp.GachaItemList = append(rsp.GachaItemList, gachaItemList)
	}

	g.send(cmd.DoGachaScRsp, rsp)
	g.UpDataPlayer()
}

func (g *Game) GachaRandom(gachaId uint32) uint32 {
	var (
		array1 []uint32 // g3
		array2 []uint32 // g4
		array3 []uint32 // g5
		array4 []uint32 // a4
		array5 []uint32 // a5
	)

	if g.Player.DbGacha.GachaMap[gachaId] == nil {
		g.Player.DbGacha.GachaMap[gachaId] = &Num{
			CeilingNum:               0,
			Pity4:                    0,
			FailedFeaturedItemPulls4: false,
			FailedFeaturedItemPulls5: false,
		}
	}

	upBanners := gdconf.GetBannersMap()[gachaId]

	for _, equi := range gdconf.GetEquipmentConfigMap() {
		switch equi.Rarity {
		case "CombatPowerLightconeRarity3":
			array1 = append(array1, equi.EquipmentID)
		case "CombatPowerLightconeRarity4":
			array2 = append(array2, equi.EquipmentID)
		case "CombatPowerLightconeRarity5":
			array3 = append(array3, equi.EquipmentID)
		}
	}

	for _, avatar := range gdconf.GetAvatarDataMap() {
		// 过滤主角
		if avatar.AvatarId/100 == 80 {
			continue
		}
		switch avatar.Rarity {
		case "CombatPowerAvatarRarityType4":
			array4 = append(array4, avatar.AvatarId)
		case "CombatPowerAvatarRarityType5":
			array5 = append(array5, avatar.AvatarId)
		}
	}

	// 保底四星
	if g.Player.DbGacha.GachaMap[gachaId].Pity4 == 9 && !g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 {
		list4 := append(array2, array4...)
		idIndex := rand.Intn(len(list4))

		for _, id := range upBanners.RateUpItems4 {
			if list4[idIndex] == id {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 = false
				break
			} else {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 = true
			}
		}
		g.Player.DbGacha.GachaMap[gachaId].Pity4 = 0
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum++
		return list4[idIndex]
	}

	// 大保底四星
	if g.Player.DbGacha.GachaMap[gachaId].Pity4 == 9 && g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 {
		idIndex := rand.Intn(len(upBanners.RateUpItems4))
		g.Player.DbGacha.GachaMap[gachaId].Pity4 = 0
		g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 = false
		return upBanners.RateUpItems4[idIndex]
	}

	// 保底五星
	if g.Player.DbGacha.GachaMap[gachaId].CeilingNum == 79 && !g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 {
		list5 := append(array3, array5...)
		idIndex := rand.Intn(len(list5))
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum = 0
		for _, id := range upBanners.RateUpItems5 {
			if list5[idIndex] == id {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 = false
				break
			} else {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 = true
			}
		}
		g.Player.DbGacha.GachaMap[gachaId].Pity4++
		return list5[idIndex]
	}

	// 大保底五星
	if g.Player.DbGacha.GachaMap[gachaId].CeilingNum == 79 && g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 {
		idIndex := rand.Intn(len(upBanners.RateUpItems5))
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum = 0
		g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 = false
		return upBanners.RateUpItems5[idIndex]
	}

	// 下面是概率
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(10000) + 1
	if randomNumber > 570 {
		idIndex := rand.Intn(len(array1))
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum++
		g.Player.DbGacha.GachaMap[gachaId].Pity4++
		return array1[idIndex]
	}
	if randomNumber > 315 {
		idIndex := rand.Intn(len(array2))
		for _, id := range upBanners.RateUpItems4 {
			if array2[idIndex] == id {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 = false
				break
			} else {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 = true
			}
		}
		g.Player.DbGacha.GachaMap[gachaId].Pity4 = 0
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum++
		return array2[idIndex]
	}
	if randomNumber > 60 {
		idIndex := rand.Intn(len(array4))
		for _, id := range upBanners.RateUpItems4 {
			if array4[idIndex] == id {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 = false
				break
			} else {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls4 = true
			}
		}
		g.Player.DbGacha.GachaMap[gachaId].Pity4 = 0
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum++
		return array4[idIndex]
	}
	if randomNumber > 30 {
		idIndex := rand.Intn(len(array3))
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum = 0
		for _, id := range upBanners.RateUpItems5 {
			if array3[idIndex] == id {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 = false
				break
			} else {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 = true
			}
		}
		g.Player.DbGacha.GachaMap[gachaId].Pity4++
		return array3[idIndex]
	}
	if randomNumber > 0 {
		idIndex := rand.Intn(len(array5))
		g.Player.DbGacha.GachaMap[gachaId].CeilingNum = 0
		for _, id := range upBanners.RateUpItems5 {
			if array5[idIndex] == id {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 = false
				break
			} else {
				g.Player.DbGacha.GachaMap[gachaId].FailedFeaturedItemPulls5 = true
			}
		}
		g.Player.DbGacha.GachaMap[gachaId].Pity4++
		return array5[idIndex]
	}
	return 0
}
