package player

import (
	"fmt"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

type commHandlerFunc func(g *GamePlayer, parameter []string) string

var commandMap = map[string]commHandlerFunc{
	"world_level": setWorldLevel,
	"give":        give,
	"give_relic":  giveRelic,
	"jump_ission": jumpMission,
}

func (g *GamePlayer) EnterCommand(msg Msg) string {
	var rspSt string
	if len(msg.CommandList) <= 0 {
		rspSt = "Command Not enough parameters"
	} else {
		commFunc, ok := commandMap[msg.CommandList[0]]
		if !ok {
			rspSt = "Unknown command"
		} else {
			rspSt = commFunc(g, msg.CommandList[1:])
		}
	}
	return fmt.Sprintf("[UID%v]执行指令:%s|响应:%s", g.Uid, msg.CommandList, rspSt)
}

// 设置世界等级
func setWorldLevel(g *GamePlayer, parameter []string) string {
	if len(parameter) < 1 {
		return "Command Not enough parameters"
	}
	g.GetPd().SetWorldLevel(alg.S2U32(parameter[0]))
	g.AllPlayerSyncScNotify(&model.AllPlayerSync{IsBasic: true})
	return fmt.Sprintf("set world_level ok")
}

func give(g *GamePlayer, parameter []string) string {
	if len(parameter) < 3 {
		return "Command Not enough parameters"
	}
	var pileItem []*model.Material
	allSync := &model.AllPlayerSync{
		IsBasic:       true,
		AvatarList:    make([]uint32, 0),
		MaterialList:  make([]uint32, 0),
		EquipmentList: make([]uint32, 0),
		RelicList:     make([]uint32, 0),
	}
	all := alg.S2U32(parameter[0])
	if all == 1 {
		pileItem = append(pileItem, g.allGive(allSync)...)
	} else {
		pileItem = append(pileItem, &model.Material{
			Tid: alg.S2U32(parameter[1]),
			Num: alg.S2U32(parameter[2]),
		})

	}
	g.GetPd().AddItem(pileItem, allSync)
	g.AllPlayerSyncScNotify(allSync)
	g.AllScenePlaneEventScNotify(pileItem)
	return fmt.Sprintf("ok")
}

func (g *GamePlayer) allGive(allSync *model.AllPlayerSync) []*model.Material {
	var pileItem []*model.Material
	itemConf := gdconf.GetItemConfigMap()
	avatarConf := gdconf.GetAvatarDataMap()
	// add avatar
	for _, avatar := range avatarConf {
		x := avatar.AvatarId / 1000
		if x != 1 && x != 8 {
			continue
		}
		allSync.AvatarList = append(allSync.AvatarList, avatar.AvatarId)
		g.GetPd().AddAvatar(avatar.AvatarId)
	}
	// add playerIcon
	var playerIconList []uint32
	for _, playerIcon := range itemConf.AvatarPlayerIcon {
		playerIconList = append(playerIconList, playerIcon.ID)
	}
	g.GetPd().GetItem().HeadIcon = playerIconList
	// add rank
	for _, rank := range itemConf.AvatarRank {
		pileItem = append(pileItem, &model.Material{
			Tid: rank.ID,
			Num: 6,
		})
	}
	// add equipment
	for _, equipment := range itemConf.Equipment {
		uniqueId := g.GetPd().AddEquipment(equipment.ID)
		allSync.EquipmentList = append(allSync.EquipmentList, uniqueId)
	}
	// add item
	for _, item := range itemConf.Item {
		pileItem = append(pileItem, &model.Material{
			Tid: item.ID,
			Num: 999999999,
		})
	}
	// add relic
	for _, relic := range itemConf.Relic {
		uniqueId := g.GetPd().AddRelic(relic.ID, 0, nil)
		allSync.RelicList = append(allSync.RelicList, uniqueId)
	}
	return pileItem
}

func giveRelic(g *GamePlayer, parameter []string) string {
	if len(parameter) < 5 {
		return "Command Not enough parameters"
	}
	allSync := &model.AllPlayerSync{
		RelicList: make([]uint32, 0),
	}
	all := alg.S2U32(parameter[0])
	if all == 1 {
		itemConf := gdconf.GetItemConfigMap()
		for _, relic := range itemConf.Relic {
			uniqueId := g.GetPd().AddRelic(relic.ID, 0, nil)
			allSync.RelicList = append(allSync.RelicList, uniqueId)
		}
	} else {
		id := alg.S2U32(parameter[1])
		num := alg.S2I(parameter[2])
		main := alg.S2U32(parameter[3])
		sub := parameter[4]
		for i := 0; i < num; i++ {
			uniqueId := g.GetPd().AddRelic(id, main, alg.GetRelicSub(sub))
			allSync.RelicList = append(allSync.RelicList, uniqueId)
		}
	}
	g.AllPlayerSyncScNotify(allSync)
	return fmt.Sprintf("ok")
}

func jumpMission(g *GamePlayer, parameter []string) string {
	if len(parameter) < 1 {
		return "Command Not enough parameters"
	}
	db := g.GetPd().GetBasicBin()
	is := alg.S2U32(parameter[0])
	if is == 0 {
		db.IsJumpMission = false
	} else {
		db.IsJumpMission = true
	}
	g.playerKickOutScNotify()
	return fmt.Sprintf("ok")
}

/**********************************分割线*******************************/

// 清空背包
func (g *GamePlayer) DelItem(payloadMsg pb.Message) {
	db := g.GetPd().GetItem()
	db = &spb.Item{
		RelicMap:     make(map[uint32]*spb.Relic),
		EquipmentMap: make(map[uint32]*spb.Equipment),
		MaterialMap:  make(map[uint32]uint32),
		HeadIcon:     make([]uint32, 0),
	}
	db.MaterialMap[11] = 240
}

// 角色一键满级
func (g *GamePlayer) GmMaxCurAvatar(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.MaxCurAvatar)
	allSync := &model.AllPlayerSync{AvatarList: make([]uint32, 0)}
	if req.All {
		bin := g.GetPd().GetAvatar()
		if bin == nil {
			return
		}
		for _, db := range bin.AvatarList {
			g.SetAvatarMaxByDb(db)
			allSync.AvatarList = append(allSync.AvatarList, db.AvatarId)
		}
	} else {
		var db *spb.AvatarBin
		db = g.GetPd().GetAvatarBinById(req.AvatarId)
		if db == nil {
			db = g.GetPd().GetCurAvatar()
		}
		allSync.AvatarList = append(allSync.AvatarList, db.AvatarId)
		g.SetAvatarMaxByDb(db)
	}
	g.AllPlayerSyncScNotify(allSync)
}

func (g *GamePlayer) SetAvatarMaxByDb(db *spb.AvatarBin) {
	if db == nil {
		return
	}
	db.Level = 80          // 80级
	db.PromoteLevel = 6    // 突破等级
	db.Hp = 10000          // 满血
	db.SpBar.CurSp = 10000 // 满能量
	for _, info := range db.MultiPathAvatarInfoList {
		info.Rank = 6                              // 六命
		for _, skill := range info.SkilltreeList { // 技能满级
			conf := gdconf.GetAvatarSkilltreeBySkillId(skill.PointId, 1)
			if conf == nil {
				continue
			}
			skill.Level = conf.MaxLevel
		}
	}
}

func (g *GamePlayer) RecoverLine() {
	db := g.GetPd().GetCurLineUp()
	for _, a := range db.AvatarIdList {
		bin := g.GetPd().GetAvatarById(a.AvatarId)
		if bin != nil {
			bin.Hp = 10000
			bin.SpBar.CurSp = 10000
		}
	}
	g.SyncLineupNotify(db)
}

func (g *GamePlayer) GmMission(req *spb.GmMission) {
	if req.FinishAll {
		g.FinishAllMission()
		g.FinishAllTutorial()
		return
	}
}

func (g *GamePlayer) FinishAllMission() {
	db := g.GetPd().GetMainMission()
	db.SubMissionList = make(map[uint32]*spb.MissionInfo)
	db.MainMissionList = make(map[uint32]*spb.MissionInfo)
	for id, info := range gdconf.GetSubMainMission() {
		if db.FinishSubMissionList == nil {
			db.FinishSubMissionList = make(map[uint32]*spb.MissionInfo)
		}
		db.FinishSubMissionList[id] = &spb.MissionInfo{
			MissionId: id,
			Progress:  info.Progress,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
	}
	for id := range gdconf.GetGoppMainMission() {
		if db.FinishMainMissionList == nil {
			db.FinishMainMissionList = make(map[uint32]*spb.MissionInfo)
		}
		db.FinishMainMissionList[id] = &spb.MissionInfo{
			MissionId: id,
			Progress:  1,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
	}
}

func (g *GamePlayer) FinishAllTutorial() {
	tDb := g.GetPd().GetTutorial()
	for id := range gdconf.GetTutorialData() {
		tDb[id] = &spb.TutorialInfo{
			Id:     id,
			Status: spb.TutorialStatus_TUTORIAL_FINISH,
		}
	}
	gDb := g.GetPd().GetTutorialGuide()
	for id := range gdconf.GetTutorialGuideGroupMap() {
		gDb[id] = &spb.TutorialInfo{
			Id:     id,
			Status: spb.TutorialStatus_TUTORIAL_FINISH,
		}
	}
}
