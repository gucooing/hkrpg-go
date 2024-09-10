package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

var playerMap = make(map[uint64]*spb.PlayerBasicCompBin, 0)
var playerLock sync.Mutex

func getPlayer(key uint64) *spb.PlayerBasicCompBin {
	playerLock.Lock()
	defer playerLock.Unlock()
	if playerMap[key] == nil {
		playerMap[key] = new(spb.PlayerBasicCompBin)
	}
	return playerMap[key]
}

func GetItem(bin *spb.PlayerBasicCompBin) *spb.Item {
	if bin.Item == nil {
		bin.Item = new(spb.Item)
	}
	return bin.Item
}

func GetMaterialMap(bin *spb.PlayerBasicCompBin) map[uint32]uint32 {
	db := GetItem(bin)
	if db.MaterialMap == nil {
		db.MaterialMap = make(map[uint32]uint32)
	}
	return db.MaterialMap
}

func GetEquipmentMap(bin *spb.PlayerBasicCompBin) map[uint32]*spb.Equipment {
	db := GetItem(bin)
	if db.EquipmentMap == nil {
		db.EquipmentMap = make(map[uint32]*spb.Equipment)
	}
	return db.EquipmentMap
}

func GetRelicMap(bin *spb.PlayerBasicCompBin) map[uint32]*spb.Relic {
	db := GetItem(bin)
	if db.RelicMap == nil {
		db.RelicMap = make(map[uint32]*spb.Relic)
	}
	return db.RelicMap
}

func upUniqueId(bin *spb.PlayerBasicCompBin, uniqueId uint32) {
	bin.UniqueId = alg.MaxUin32(bin.UniqueId, uniqueId)
}

func GetAvatar(bin *spb.PlayerBasicCompBin) *spb.Avatar {
	if bin.Avatar == nil {
		bin.Avatar = new(spb.Avatar)
	}
	return bin.Avatar
}

func GetLineUp(bin *spb.PlayerBasicCompBin) *spb.LineUp {
	if bin.LineUp == nil {
		bin.LineUp = &spb.LineUp{
			MainLineUp:     0,
			Mp:             5,
			LineUpList:     nil,
			BattleLineList: nil,
			StoryLineList:  nil,
		}
	}
	return bin.LineUp
}

type HandlerFunc func(payloadMsg pb.Message, conv uint64)

var handlerFuncRouteMap map[uint16]HandlerFunc

func init() {
	handlerFuncRouteMap = map[uint16]HandlerFunc{
		cmd.PlayerGetTokenScRsp:    PlayerGetTokenScRsp,
		cmd.PlayerLoginScRsp:       PlayerLoginScRsp,
		cmd.GetBagScRsp:            GetBagScRsp,
		cmd.GetAvatarDataScRsp:     GetAvatarDataScRsp,
		cmd.GetAllLineupDataScRsp:  GetAllLineupDataScRsp,
		cmd.GetBasicInfoScRsp:      GetBasicInfoScRsp,
		cmd.PlayerLoginFinishCsReq: PlayerBinClose,
	}
}

func playerBinRegisterMessage(protoObj pb.Message, cmdId uint16, conv uint64) {
	handlerFunc, ok := handlerFuncRouteMap[cmdId]
	if !ok {
		return
	}
	handlerFunc(protoObj, conv)
}

func PlayerBinClose(protoObj pb.Message, conv uint64) {
	player := getPlayer(conv)
	bin, _ := pb.Marshal(player)
	b64 := base64.StdEncoding.EncodeToString(bin)
	name := fmt.Sprintf("./%v.txt", player.Uid)
	err := ioutil.WriteFile(name, []byte(b64), 0644)
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return
	}

	fmt.Printf("[UID:%v]玩家数据保存成功\n", player.Uid)
}

func PlayerGetTokenScRsp(protoObj pb.Message, conv uint64) {
	req := protoObj.(*proto.PlayerGetTokenScRsp)
	if req == nil {
		return
	}
	player := getPlayer(conv)
	player.Uid = req.Uid
}

func PlayerLoginScRsp(protoObj pb.Message, conv uint64) {
	req := protoObj.(*proto.PlayerLoginScRsp)
	if req == nil {
		return
	}
	player := getPlayer(conv)
	basicInfo := req.BasicInfo
	if basicInfo != nil {
		player.Level = basicInfo.Level
		player.WorldLevel = basicInfo.WorldLevel
		player.Nickname = basicInfo.Nickname
		db := GetMaterialMap(player)
		db[model.Hcoin] = basicInfo.Hcoin
		db[model.Scoin] = basicInfo.Scoin
		db[model.Mcoin] = basicInfo.Mcoin
		db[model.Stamina] = basicInfo.Stamina
		db[model.Exp] = basicInfo.Exp
	}
}

func GetBagScRsp(protoObj pb.Message, conv uint64) {
	req := protoObj.(*proto.GetBagScRsp)
	if req == nil {
		return
	}
	player := getPlayer(conv)
	materialList := GetMaterialMap(player)
	for _, v := range req.MaterialList {
		materialList[v.Tid] = v.Num
	}
	equipmentMap := GetEquipmentMap(player)
	for _, v := range req.EquipmentList {
		equipmentMap[v.UniqueId] = &spb.Equipment{
			Tid:          v.Tid,
			UniqueId:     v.UniqueId,
			Exp:          v.Exp,
			Level:        v.Level,
			Promotion:    v.Promotion,
			BaseAvatarId: v.DressAvatarId,
			IsProtected:  v.IsProtected,
			Rank:         v.Rank,
		}
		upUniqueId(player, v.UniqueId)
	}
	relicMap := GetRelicMap(player)
	for _, v := range req.RelicList {
		info := &spb.Relic{
			Tid:          v.Tid,
			UniqueId:     v.UniqueId,
			Exp:          v.Exp,
			Level:        v.Level,
			MainAffixId:  v.MainAffixId,
			RelicAffix:   make(map[uint32]*spb.RelicAffix),
			BaseAvatarId: v.DressAvatarId,
			IsProtected:  v.IsProtected,
		}
		for _, sub := range v.SubAffixList {
			info.RelicAffix[sub.AffixId] = &spb.RelicAffix{
				AffixId: sub.AffixId,
				Cnt:     sub.Cnt,
				Step:    sub.Step,
			}
		}
		relicMap[v.UniqueId] = info
		upUniqueId(player, v.UniqueId)
	}
}

func GetBasicInfoScRsp(protoObj pb.Message, conv uint64) {
	req := protoObj.(*proto.GetBasicInfoScRsp)
	if req == nil {
		return
	}
	player := getPlayer(conv)
	pbAva := GetAvatar(player)
	pbAva.Gender = spb.Gender(req.Gender)
}

func GetAvatarDataScRsp(protoObj pb.Message, conv uint64) {
	req := protoObj.(*proto.GetAvatarDataScRsp)
	if req == nil {
		return
	}
	player := getPlayer(conv)
	pbAva := GetAvatar(player)
	if req.IsGetAll {
		if pbAva.AvatarList == nil {
			pbAva.AvatarList = make(map[uint32]*spb.AvatarBin)
		}
		for _, v := range req.AvatarList {
			info := &spb.AvatarBin{
				AvatarId:          v.BaseAvatarId,
				Exp:               v.Exp,
				Level:             v.Level,
				AvatarType:        3,
				FirstMetTimeStamp: v.FirstMetTimeStamp,
				PromoteLevel:      v.Promotion,
				TakenRewards:      v.HasTakenPromotionRewardList,
				Hp:                12000,
				SpBar: &spb.AvatarSpBarInfo{
					CurSp: 12000,
					MaxSp: 12000,
				},
				IsMultiPath:             false,
				CurPath:                 v.BaseAvatarId,
				MultiPathAvatarInfoList: make(map[uint32]*spb.MultiPathAvatarInfo),
			}
			multi := &spb.MultiPathAvatarInfo{
				AvatarId:          v.BaseAvatarId,
				Rank:              v.Rank,
				SkilltreeList:     make([]*spb.AvatarSkillBin, 0),
				EquipmentUniqueId: v.EquipmentUniqueId,
				EquipRelic:        make(map[uint32]uint32),
			}
			for _, relic := range v.EquipRelicList {
				multi.EquipRelic[relic.Type] = relic.RelicUniqueId
			}
			for _, skill := range v.SkilltreeList {
				multi.SkilltreeList = append(multi.SkilltreeList, &spb.AvatarSkillBin{
					PointId: skill.PointId,
					Level:   skill.Level,
				})
			}
			info.MultiPathAvatarInfoList[v.BaseAvatarId] = multi
			pbAva.AvatarList[v.BaseAvatarId] = info
		}
	}
}

func GetAllLineupDataScRsp(protoObj pb.Message, conv uint64) {
	req := protoObj.(*proto.GetAllLineupDataScRsp)
	if req == nil {
		return
	}
	player := getPlayer(conv)
	db := GetLineUp(player)
	db.MainLineUp = req.CurIndex
	if db.LineUpList == nil {
		db.LineUpList = make(map[uint32]*spb.Line)
	}
	for _, v := range req.LineupList {
		if v.ExtraLineupType != proto.ExtraLineupType_LINEUP_NONE {
			continue
		}
		info := &spb.Line{
			Name:         v.Name,
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			Index:        v.Index,
			LeaderSlot:   v.LeaderSlot,
			LineType:     spb.ExtraLineupType(v.ExtraLineupType),
		}
		for _, ava := range v.AvatarList {
			info.AvatarIdList[ava.Slot] = &spb.LineAvatarList{
				Slot:     ava.Slot,
				AvatarId: ava.Id,
			}
		}
		db.LineUpList[v.Index] = info
	}
}
