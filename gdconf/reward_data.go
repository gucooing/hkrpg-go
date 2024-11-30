package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RewardData struct {
	RewardID uint32   `json:"RewardID"`
	Hcoin    uint32   `json:"Hcoin"`
	Items    []*Items `json:"-"`
	ItemID_1 uint32   `json:"ItemID_1"`
	Count_1  uint32   `json:"Count_1"`
	ItemID_2 uint32   `json:"ItemID_2"`
	Count_2  uint32   `json:"Count_2"`
	ItemID_3 uint32   `json:"ItemID_3"`
	Count_3  uint32   `json:"Count_3"`
	ItemID_4 uint32   `json:"ItemID_4"`
	Count_4  uint32   `json:"Count_4"`
	ItemID_5 uint32   `json:"ItemID_5"`
	Count_5  uint32   `json:"Count_5"`
	ItemID_6 uint32   `json:"ItemID_6"`
	Count_6  uint32   `json:"Count_6"`
}

type Items struct {
	ItemID uint32
	Count  uint32
}

func (g *GameDataConfig) loadRewardData() {
	g.RewardDataMap = make(map[uint32]*RewardData)
	rewardDataMap := make([]*RewardData, 0)
	name := "RewardData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rewardDataMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, data := range rewardDataMap {
		data.Items = make([]*Items, 0)
		if data.ItemID_1 != 0 {
			data.Items = append(data.Items, &Items{
				ItemID: data.ItemID_1,
				Count:  data.Count_1,
			})
		}
		if data.ItemID_2 != 0 {
			data.Items = append(data.Items, &Items{
				ItemID: data.ItemID_2,
				Count:  data.Count_2,
			})
		}
		if data.ItemID_3 != 0 {
			data.Items = append(data.Items, &Items{
				ItemID: data.ItemID_3,
				Count:  data.Count_3,
			})
		}
		if data.ItemID_4 != 0 {
			data.Items = append(data.Items, &Items{
				ItemID: data.ItemID_4,
				Count:  data.Count_4,
			})
		}
		if data.ItemID_5 != 0 {
			data.Items = append(data.Items, &Items{
				ItemID: data.ItemID_5,
				Count:  data.Count_5,
			})
		}
		if data.ItemID_6 != 0 {
			data.Items = append(data.Items, &Items{
				ItemID: data.ItemID_6,
				Count:  data.Count_6,
			})
		}
		g.RewardDataMap[data.RewardID] = data
	}

	logger.Info(text.GetText(17), len(g.RewardDataMap), name)
}

func GetRewardDataById(id uint32) *RewardData {
	return getConf().RewardDataMap[id]
}
