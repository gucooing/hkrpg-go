package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type AllPlayerSync struct {
	IsBasic       bool     // 基本信息
	AvatarList    []uint32 // 角色列表
	MaterialList  []uint32 // 物品id列表
	EquipmentList []uint32 // 光锥列表
	RelicList     []uint32 // 圣遗物列表
}

// 角色状态改变时需要发送通知
func (g *GamePlayer) PlayerPlayerSyncScNotify() {
	db := g.GetMaterialMap()
	notify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   g.GetNickname(),
			Level:      g.GetLevel(),
			Exp:        db[Exp],
			Hcoin:      db[Hcoin],
			Scoin:      db[Scoin],
			Mcoin:      db[Mcoin],
			Stamina:    db[Stamina],
			WorldLevel: g.GetWorldLevel(),
		},
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) AddTrailblazerExp(num uint32) {
	material := g.GetMaterialMap()
	db := g.GetBasicBin()
	material[Exp] += num
	level, exp, worldLevel := gdconf.GetPlayerLevelConfigByLevel(material[Exp], g.GetLevel(), g.GetWorldLevel())
	material[Exp] = exp
	db.Level = level
	db.WorldLevel = worldLevel
	g.PlayerPlayerSyncScNotify()
}

func (g *GamePlayer) SetPlayerInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetPlayerInfoCsReq, payloadMsg)
	req := msg.(*proto.SetPlayerInfoCsReq)

	g.SetNickname(req.Nickname)

	if req.IsModify {
		if req.Gender == proto.Gender_GenderWoman {
			g.Send(cmd.HeroBasicTypeChangedNotify, &proto.HeroBasicTypeChangedNotify{CurBasicType: proto.HeroBasicType_GirlWarrior})
			db := g.GetAvatar()
			db.Gender = spb.Gender_GenderWoman
			db.CurMainAvatar = spb.HeroBasicType_GirlWarrior
			g.AvatarPlayerSyncScNotify(8001)
		}
		g.CreateCharacterSubMission()
	}
	rsp := &proto.SetPlayerInfoScRsp{
		Retcode:      0,
		CurBasicType: proto.HeroBasicType(g.GetAvatar().CurMainAvatar),
		EGCIPLNFHGD:  0,
		IsModify:     req.IsModify,
	}
	g.PlayerPlayerSyncScNotify() // 角色信息通知
	g.Send(cmd.SetPlayerInfoScRsp, rsp)
}

func (g *GamePlayer) AllPlayerSyncScNotify(allSync *AllPlayerSync) {
	if allSync == nil {
		return
	}
	notify := &proto.PlayerSyncScNotify{
		AvatarSync:   &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		MaterialList: make([]*proto.Material, 0),
	}
	db := g.GetMaterialMap()
	// 添加账户基本信息
	if allSync.IsBasic {
		notify.BasicInfo = &proto.PlayerBasicInfo{
			Nickname:   g.GetNickname(),
			Level:      g.GetLevel(),
			Exp:        db[Exp],
			Hcoin:      db[Hcoin],
			Scoin:      db[Scoin],
			Mcoin:      db[Mcoin],
			Stamina:    db[Stamina],
			WorldLevel: g.GetWorldLevel(),
		}
	}
	// 添加角色
	if allSync.AvatarList != nil {
		for _, avatarId := range allSync.AvatarList {
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, g.GetProtoAvatarById(avatarId))
		}
	}
	// 添加物品
	if allSync.MaterialList != nil {
		for _, materialId := range allSync.MaterialList {
			notify.MaterialList = append(notify.MaterialList, &proto.Material{
				Tid: materialId,
				Num: db[materialId],
			})
		}
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}
