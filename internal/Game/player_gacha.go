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
			BeginTime:  bannerslist.BeginTime,                          // 开始时间
			EndTime:    bannerslist.EndTime,                            // 结束时间
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
	rand.Seed(time.Now().UnixNano())

	if req.GachaNum != 10 && req.GachaNum != 1 {
		return
	}
	avatarList := gdconf.GetAvatarList()
	equipmentList := gdconf.GetEquipmentList()

	list := append(avatarList, equipmentList...)

	rsp := &proto.DoGachaScRsp{
		GachaId:       req.GachaId,
		CeilingNum:    0,
		GachaItemList: make([]*proto.GachaItem, 0),
		GachaNum:      req.GachaNum,
	}
	for i := 0; i < int(req.GachaNum); i++ {
		idIndex := rand.Intn(len(list))
		id := list[idIndex]
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
}
