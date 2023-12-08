package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetAvatarDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAvatarDataScRsp)
	rsp.IsGetAll = true
	rsp.AvatarList = make([]*proto.Avatar, 0)

	for _, a := range g.Player.DbAvatar.Avatar {
		avatarList := new(proto.Avatar)
		avatarList.FirstMetTimestamp = a.FirstMetTimestamp
		avatarList.BaseAvatarId = a.AvatarId
		avatarList.Promotion = a.Promotion
		avatarList.Rank = a.Rank
		avatarList.Level = a.Level
		avatarList.Exp = a.Exp
		avatarList.SkilltreeList = GetKilltreeList(a.AvatarId, 1)
		rsp.AvatarList = append(rsp.AvatarList, avatarList)
	}

	g.send(cmd.GetAvatarDataScRsp, rsp)
}

func (g *Game) RankUpAvatarCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.RankUpAvatarCsReq, payloadMsg)
	req := msg.(*proto.RankUpAvatarCsReq)

	g.Player.DbAvatar.Avatar[req.BaseAvatarId].Rank++
	g.SubtractMaterial(req.BaseAvatarId+10000, 1)
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)

	rsp := new(proto.GetChallengeScRsp)
	g.send(cmd.RankUpAvatarScRsp, rsp)
}
