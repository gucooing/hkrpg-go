package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	PreviewSkillPoint        uint32                    `json:"PreviewSkillPoint"`        // 秘技回复点数
	IsShowItemDesc           bool                      `json:"IsShowItemDesc"`
	ActivityCount            uint32                    `json:"ActivityCount"`
}

func (g *GameDataConfig) loadItemUseBuffData() {
	g.ItemUseBuffDataMap = make(map[uint32]*ItemUseBuffData)
	itemUseBuffData := make([]*ItemUseBuffData, 0)
	playerElementsFilePath := g.excelPrefix + "ItemUseBuffData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &itemUseBuffData)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range itemUseBuffData {
		g.ItemUseBuffDataMap[v.UseDataID] = v
	}
	logger.Info("load %v ItemUseBuffData", len(g.ItemUseBuffDataMap))
}

func GetItemUseBuffDataById(id uint32) *ItemUseBuffData {
	return CONF.ItemUseBuffDataMap[id]
}
