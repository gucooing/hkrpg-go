package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ItemUseBuffData struct {
	UseDataID                uint32                    `json:"UseDataID"`
	ConsumeType              uint32                    `json:"ConsumeType"` // 道具类型 1,2:战斗使用的buff ,3:场景buff 4,5:立即回复道具
	ConsumeTag               []constant.ItemFoodUseTag `json:"ConsumeTag"`  // 增益效果
	UseTargetType            string                    `json:"UseTargetType"`
	MazeBuffID               uint32                    `json:"MazeBuffID"`
	MazeBuffParam            []uint32                  `json:"MazeBuffParam"`
	MazeBuffID2              uint32                    `json:"MazeBuffID2"`
	MazeBuffParam2           []uint32                  `json:"MazeBuffParam2"`
	UseMultipleMax           uint32                    `json:"UseMultipleMax"`
	IsCheckHP                bool                      `json:"IsCheckHP"` // 检查hp
	UseEffect                string                    `json:"UseEffect"`
	PreviewHPRecoveryPercent float64                   `json:"PreviewHPRecoveryPercent"` // 回血百分比
	PreviewHPRecoveryValue   float64                   `json:"PreviewHPRecoveryValue"`   // 回血数值
	PreviewSkillPoint        float64                   `json:"PreviewSkillPoint"`        // 秘技回复点数
	IsShowItemDesc           bool                      `json:"IsShowItemDesc"`
	ActivityCount            uint32                    `json:"ActivityCount"`
}

func (g *GameDataConfig) loadItemUseBuffData() {
	g.ItemUseBuffDataMap = make(map[uint32]*ItemUseBuffData)
	itemUseBuffData := make([]*ItemUseBuffData, 0)
	name := "ItemUseBuffData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &itemUseBuffData)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range itemUseBuffData {
		g.ItemUseBuffDataMap[v.UseDataID] = v
	}

	logger.Info(text.GetText(17), len(g.ItemUseBuffDataMap), name)
}

func GetItemUseBuffDataById(id uint32) *ItemUseBuffData {
	return getConf().ItemUseBuffDataMap[id]
}
