package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MainMission struct {
	MainMissionID        uint32       `json:"MainMissionID"`
	Type                 string       `json:"Type"` // 类型
	DisplayPriority      uint32       `json:"DisplayPriority"`
	NextMainMissionList  []uint32     `json:"NextMainMissionList"` // 下一个主线任务列表
	TakeOperation        string       `json:"TakeOperation"`
	BeginOperation       string       `json:"BeginOperation"`
	TakeParam            []*TakeParam `json:"TakeParam"`            // 接取条件
	BeginParam           []*TakeParam `json:"BeginParam"`           // 开始条件
	NextTrackMainMission uint32       `json:"NextTrackMainMission"` // 下一个主线任务
	TrackWeight          uint32       `json:"TrackWeight"`
	RewardID             uint32       `json:"RewardID"`
	DisplayRewardID      uint32       `json:"DisplayRewardID"`
	ChapterID            uint32       `json:"ChapterID"`
	SubRewardList        []uint32     `json:"SubRewardList"`
}

type TakeParam struct {
	Type  constant.MissionBeginType `json:"Type"`
	Value uint32                    `json:"Value"`
}

func (g *GameDataConfig) loadMainMission() {
	g.MainMissionMap = make(map[uint32]*MainMission)
	mainMissionMap := make([]*MainMission, 0)
	playerElementsFilePath := g.excelPrefix + "MainMission.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &mainMissionMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range mainMissionMap {
		g.MainMissionMap[v.MainMissionID] = v
	}

	logger.Info("load %v MainMission", len(g.MainMissionMap))
}

func GetMainMission() map[uint32]*MainMission {
	return CONF.MainMissionMap
}

func GetMainMissionById(id uint32) *MainMission {
	return CONF.MainMissionMap[id]
}
