package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) DressRelicAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DressRelicAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressRelicAvatarCsReq)

	g.DressRelicAvatarPlayerSyncScNotify(req.BaseAvatarId, req.ParamList)
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.DressAvatarScRsp, rsp)
}

func (g *Game) DressRelicAvatarPlayerSyncScNotify(avatarId uint32, paramList []*proto.RelicParam) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		RelicList:  make([]*proto.Relic, 0),
	}
	avatardb := g.Player.DbAvatar.Avatar[avatarId]
	// 是否已被装备
	for _, relic := range paramList {
		relicdb := g.Player.DbItem.RelicMap[relic.RelicUniqueId]
		if relicdb == nil {
			return
		}
		if relicdb.BaseAvatarId != 0 {
			// 进入交换
			avatardbs := g.Player.DbAvatar.Avatar[relicdb.BaseAvatarId]
			if avatardb.EquipRelic[relic.Slot] == 0 {
				delete(g.Player.DbAvatar.Avatar[relicdb.BaseAvatarId].EquipRelic, relic.Slot)
			} else {
				g.Player.DbAvatar.Avatar[relicdb.BaseAvatarId].EquipRelic[relic.Slot] = avatardb.EquipRelic[relic.Slot]
				g.Player.DbItem.RelicMap[avatardb.EquipRelic[relic.Slot]].BaseAvatarId = avatardbs.AvatarId

				relicList := g.GetRelic(avatardb.EquipRelic[relic.Slot])
				notify.RelicList = append(notify.RelicList, relicList)
			}
			avatar := g.GetAvatar(avatardbs.AvatarId)
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)
		}

		if avatardb.EquipRelic[relic.Slot] != 0 {
			if g.Player.DbItem.RelicMap[avatardb.EquipRelic[relic.Slot]].BaseAvatarId == avatarId {
				g.Player.DbItem.RelicMap[avatardb.EquipRelic[relic.Slot]].BaseAvatarId = 0
				relicList := g.GetRelic(avatardb.EquipRelic[relic.Slot])
				notify.RelicList = append(notify.RelicList, relicList)
			}
		}
		g.Player.DbItem.RelicMap[relic.RelicUniqueId].BaseAvatarId = avatarId
		g.Player.DbAvatar.Avatar[avatarId].EquipRelic[relic.Slot] = relic.RelicUniqueId

		relicList := g.GetRelic(relic.RelicUniqueId)
		notify.RelicList = append(notify.RelicList, relicList)
	}

	avatar := g.GetAvatar(avatarId)
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) ExpUpRelicCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExpUpRelicCsReq, payloadMsg)
	req := msg.(*proto.ExpUpRelicCsReq)
	if req.RelicUniqueId == 0 {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	// var relicList []uint32   // 需要删除的relicList
	// var pileItem []*Material // 需要删除的升级材料
	// var delScoin uint32      // 扣除的信用点
	// var addExp uint32        // 增加的经验

	// 从背包获取需要升级的圣遗物
	dbRelic := g.Player.DbItem.RelicMap[req.RelicUniqueId]
	if dbRelic == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}
	// 获取需要升级圣遗物的配置信息
	relicConf := gdconf.GetRelicById(strconv.Itoa(int(dbRelic.Tid)))
	if relicConf == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}
}
