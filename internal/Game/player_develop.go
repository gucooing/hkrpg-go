package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) ExpUpEquipmentCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.ExpUpEquipmentCsReq, payloadMsg)
	req := msg.(*proto.ExpUpEquipmentCsReq)

	var equipmentList []uint32 // 需要删除的equipmentList
	var delScoin uint32        // 扣除的信用点
	var addExp uint32          // 增加的经验

	// TODO 需要注意的是，还有特殊升级物件(读表了再加),还有需要注意的是，升级后不能超过此阶段最大等级(也是后面再写，今天摆烂)

	// 从背包获取需要升级的光锥
	upEquipment := g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId]
	// 获取需要升级光锥的配置信息
	equConf := gdconf.GetEquipmentConfigById(strconv.Itoa(int(upEquipment.Tid)))

	// 遍历用来升级的光锥
	for _, equipment := range req.ItemCostList.ItemList {
		// 一个扣250信用点
		delScoin += 250
		equipmentList = append(equipmentList, equipment.GetEquipmentUniqueId())
		// 获取光锥配置
		equipmentconfig := gdconf.GetEquipmentConfigById(strconv.Itoa(int(g.Player.DbItem.EquipmentMap[equipment.GetEquipmentUniqueId()].Tid)))
		// 获取能添加多少经验
		addExp += equipmentconfig.ExpProvide
	}

	// 计算添加后有多少经验
	exp := addExp + upEquipment.Exp

	// 获取能升级到的等级和升级后经验
	level, exp := gdconf.GetEquipmentExpByLevel(equConf.ExpType, exp, upEquipment.Level)
	if level == 0 && exp == 0 {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.send(cmd.ExpUpEquipmentScRsp, rsp)
	}

	// 扣除本次升级需要的信用点
	g.Player.DbItem.MaterialMap[2].Num -= delScoin
	// 更新需要升级的光锥状态
	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Level = level
	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Exp = exp

	// 删除用来升级的光锥
	g.DelPlayerSyncScNotify(equipmentList)
	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()
	// 通知升级后光锥消息
	g.EquipmentPlayerSyncScNotify(upEquipment.Tid, req.EquipmentUniqueId)
	rsp := &proto.ExpUpEquipmentScRsp{}
	g.send(cmd.ExpUpEquipmentScRsp, rsp)
}

func (g *Game) PlayerPlayerSyncScNotify() {
	notify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   g.Player.NickName,
			Level:      g.Player.Level,
			Exp:        g.Player.Exp,
			Stamina:    g.Player.Stamina,
			Mcoin:      0,
			Hcoin:      g.Player.DbItem.MaterialMap[1].Num,
			Scoin:      g.Player.DbItem.MaterialMap[2].Num,
			WorldLevel: g.Player.WorldLevel,
		},
	}

	g.send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) DelPlayerSyncScNotify(equipmentList []uint32) {
	for _, equipment := range equipmentList {
		delete(g.Player.DbItem.EquipmentMap, equipment)
	}

	notify := &proto.PlayerSyncScNotify{DelEquipmentList: equipmentList}
	g.send(cmd.PlayerSyncScNotify, notify)
}
