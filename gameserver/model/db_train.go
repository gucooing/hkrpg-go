package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func newTrain() *spb.Train {
	return &spb.Train{
		Pam: &spb.PamInfo{
			CurPamSkin:      252000,
			UnlockedPamSkin: map[uint32]bool{252000: true},
		},
	}
}

func (g *PlayerData) GetTrain() *spb.Train {
	db := g.GetBasicBin()
	if db.Train == nil {
		db.Train = newTrain()
	}
	return db.Train
}

func (g *PlayerData) GetPamInfo() *spb.PamInfo {
	db := g.GetTrain()
	if db.Pam == nil {
		db.Pam = &spb.PamInfo{}
	}
	return db.Pam
}

func (g *PlayerData) GetCurPamSkin() uint32 {
	db := g.GetPamInfo()
	if db.CurPamSkin == 0 {
		db.CurPamSkin = 252000
	}
	return db.CurPamSkin
}

func (g *PlayerData) GetUnlockedPamSkin() map[uint32]bool {
	db := g.GetPamInfo()
	if db.UnlockedPamSkin == nil {
		db.UnlockedPamSkin = map[uint32]bool{252000: true}
	}
	return db.UnlockedPamSkin
}

func (g *PlayerData) AddUnlockedPamSkin(skin uint32) {
	db := g.GetUnlockedPamSkin()
	db[skin] = true
}

func (g *PlayerData) SetCurPamSkin(skin uint32) bool {
	db := g.GetUnlockedPamSkin()
	if !db[skin] {
		return false
	}
	pamInfo := g.GetPamInfo()
	pamInfo.CurPamSkin = skin
	return true
}

/*****************接口*******************/

func (g *PlayerData) GetVisitorInfoList() []*proto.TrainVisitorInfo {
	list := make([]*proto.TrainVisitorInfo, 0)
	for _, conf := range gdconf.GetTrainVisitorConfigMap() {
		info := &proto.TrainVisitorInfo{
			MissionId:              conf.MissionID,
			VisitorId:              conf.VisitorID,
			BKHCLMJKOIP:            make([]uint32, 0),
			ToastFinishMainMission: conf.ToastFinishMainMission,
			Status:                 proto.TrainVisitorStatus_TRAIN_VISITOR_STATUS_GET_ON,
		}
		list = append(list, info)
	}
	return list
}

func (g *PlayerData) GetPassengerInfo() *proto.TrainPartyPassengerInfo {
	info := &proto.TrainPartyPassengerInfo{
		PassengerInfoList: make([]*proto.TrainPartyPassenger, 0),
	}

	for _, v := range gdconf.GetTrainPartyPassengerConfigMap() {
		info.PassengerInfoList = append(info.PassengerInfoList, &proto.TrainPartyPassenger{
			RecordId:    0,
			DNGLGOMMFNP: false,
			BFGGMBADMHB: make([]uint32, 0),
			PassengerId: v.PassengerID,
		})
	}

	return info
}

func (g *PlayerData) GetTrainPartyGameInfo() *proto.TrainPartyGameInfo {
	info := &proto.TrainPartyGameInfo{
		TeamId:             1,
		TrainPartyItemInfo: g.GetTrainPartyItemInfo(),
		TrainPassengerInfo: g.GetTrainPassengerInfo(),
		TrainPartyGridInfo: g.GetTrainPartyGridInfo(),
		TrainActionInfo: &proto.TrainPartyActionInfo{
			QueuePosition: 1,
			TrainActionCase: &proto.TrainPartyActionInfo_TrainPartyEvent{
				TrainPartyEvent: &proto.TrainPartyEvent{
					OptionList: make([]*proto.TrainPartyOption, 0),
					EventType:  proto.TrainPartyEventType_kGamePlayStartDialogueEvent,
					EventId:    60004,
				},
			},
		},
	}

	return info
}

func (g *PlayerData) GetTrainPartyGridInfo() *proto.TrainPartyGameGridInfo {
	info := &proto.TrainPartyGameGridInfo{
		OOPDLCMLKKL: 1,
		NOFPLBABCCB: 1,
		OPDABBNHHCG: &proto.OKFGPABKEJE{
			AADBGCLODKK: 0,
			HLIDPIFAJCG: make([]uint32, 0),
			AOJBNMEECHF: 16,
			DFNCNGPDILM: 8001,
			EPALCFOJBKJ: 2,
			ECFPFKNJINA: 0,
		},
		GridList: make([]*proto.TrainPartyGameGrid, 0),
	}

	for i := 1; i < 12; i++ {
		info.GridList = append(info.GridList, &proto.TrainPartyGameGrid{
			GAEIOFOPLFN:  uint32(i),
			GridId:       1001,
			DisplayValue: 0,
			UniqueId:     uint32(i),
		})
	}

	return info
}

func (g *PlayerData) GetTrainPartyItemInfo() *proto.TrainPartyGameItemInfo {
	info := &proto.TrainPartyGameItemInfo{
		ALPDNPDDPJC: 100,
		ECLBPHPMGIN: false,
		TrainPartyCardInfo: &proto.TrainPartyGameCardInfo{
			TrainPartyCardInfo: make([]*proto.TrainPartyGameCard, 0),
		},
	}

	for _, v := range gdconf.GetTrainPartyCardConfigMap() {
		info.TrainPartyCardInfo.TrainPartyCardInfo = append(info.TrainPartyCardInfo.TrainPartyCardInfo,
			&proto.TrainPartyGameCard{
				UniqueId:    g.GetNextGameObjectGuid(),
				HMKMKBELCLG: make([]uint32, 0),
				NAGKACACHGD: 1,
				CardId:      v.CardID,
			})
	}

	return info
}

func (g *PlayerData) GetTrainPassengerInfo() *proto.TrainPartyGamePassengerInfo {
	info := &proto.TrainPartyGamePassengerInfo{
		CurPassengerId:       1004,
		MtRankId:             73,
		PassengerList:        make([]*proto.TrainPartyGamePassenger, 0),
		NAGKACACHGD:          0,
		AetherSkillList:      make([]*proto.TrainPartyGameSkill, 0),
		PassengerDiaryIdList: []uint32{101},
	}

	for _, v := range gdconf.GetTrainPartyPassengerConfigMap() {
		info.PassengerList = append(info.PassengerList, &proto.TrainPartyGamePassenger{
			DIKPBINAOOH: 100,
			PassengerId: v.PassengerID,
			IHLDLJGDCBL: nil,
		})
	}

	return info
}

func (g *PlayerData) GetTrainPartyInfo() *proto.TrainPartyInfo {
	info := &proto.TrainPartyInfo{
		AreaList:       g.GetTrainPartyAreaList(),
		CoinCost:       1000,
		TrainPartyRank: 4,
		NNBHLDDNLDE:    make([]*proto.HHOKBPHNFNE, 0),
		CFJKBJHNIJM:    0,
		TrainPartyTag:  1,
		KCFEECACMOD:    make([]uint32, 0),
		CDJFDJIAING:    0,
		DynamicIdList:  make([]uint32, 0),
		HIMCAOKDNMP:    0,
	}

	return info
}

func (g *PlayerData) GetTrainPartyAreaList() []*proto.TrainPartyArea {
	list := make([]*proto.TrainPartyArea, 0)
	for _, agc := range gdconf.GetTrainPartyAreaGoalConfigMap() {
		area := &proto.TrainPartyArea{
			AreaId:   agc.AreaID, // 区域
			Progress: 100,        // 进度
			AreaStepInfo: &proto.AreaStepInfo{
				AreaStepList: make([]*proto.BuildAreaStep, 0), // 该区域主物品信息
				AreaGlobalId: agc.ID,                          // 当前区域的组Id
			},
			StaticPropIdList: make([]uint32, 0), // 解锁的主物品 Prop List
			StepIdList:       make([]uint32, 0), // 解锁的主物品 List
			DynamicInfo:      nil,               // 该区域解锁的摆放物

			NOPJINFMFEI: make([]uint32, 0),
		}
		for _, groupId := range agc.StepGroupList {
			stepList := gdconf.GetTrainPartyStepConfigByGroupId(groupId)
			for _, step := range stepList {
				// TODO 53开始通过前置区域解锁
				// if step.CoinCost > 100000000 {
				// 	continue
				// }
				area.AreaStepInfo.AreaStepList = append(area.AreaStepInfo.AreaStepList, &proto.BuildAreaStep{
					Status:      proto.BuildGoalStep_BuildGoalStepFinish,
					StepId:      step.ID,
					MJALJMGLEFP: 0,
				})
				area.StepIdList = append(area.StepIdList, step.ID)
				area.StaticPropIdList = append(area.StaticPropIdList, step.StaticPropIDList...)
			}
		}
		list = append(list, area)
	}

	return list
}
