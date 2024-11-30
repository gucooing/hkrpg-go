package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type LoadingDesc struct {
	ID       uint32 `json:"ID"`       // 随机种子
	MinLevel uint32 `json:"MinLevel"` // 最小等级
	MaxLevel uint32 `json:"MaxLevel"` // 最大等级
	Weight   uint32 `json:"Weight"`   // 比重(貌似全部都是20)
	// TODO 还有部分字段没读，需要再读
}

func (g *GameDataConfig) loadLoadingDesc() {
	g.LoadingDescMap = make(map[uint32]*LoadingDesc)
	loadingDescMap := make([]*LoadingDesc, 0)
	name := "LoadingDesc.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &loadingDescMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range loadingDescMap {
		g.LoadingDescMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.LoadingDescMap), name)
}

func GetLoadingDesc() uint32 {
	var list []uint32
	idIndex := rand.Intn(len(getConf().LoadingDescMap))
	for _, id := range getConf().LoadingDescMap {
		list = append(list, id.ID)
	}
	return list[idIndex]
}
