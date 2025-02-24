package player

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func HandleGetJukeboxDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	db := g.GetPd().GetPhoneData()
	rsp := &proto.GetJukeboxDataScRsp{
		CurrentMusicId:    db.CurrentMusicId,
		Retcode:           0,
		UnlockedMusicList: make([]*proto.MusicData, 0),
	}
	musicMap := g.GetPd().GetMusicInfoMap()
	for id, v := range musicMap {
		conf := gdconf.GetBackGroundMusicById(v.MusicId)
		if conf == nil {
			delete(musicMap, id)
			continue
		}
		musicList := &proto.MusicData{
			GroupId:  conf.GroupID,
			IsPlayed: true,
			Id:       conf.ID,
		}
		rsp.UnlockedMusicList = append(rsp.UnlockedMusicList, musicList)
	}
	g.Send(cmd.GetJukeboxDataScRsp, rsp)
}

func UnlockBackGroundMusicCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UnlockBackGroundMusicCsReq)

	rsp := &proto.UnlockBackGroundMusicScRsp{
		Retcode:           0,
		UnlockedMusicList: make([]*proto.MusicData, 0),
		UnlockedIds:       make([]uint32, 0),
	}
	for _, unlockId := range req.UnlockIds {
		if g.GetPd().GetMaterialById(unlockId) == 0 {
			continue
		}
		g.GetPd().SetMaterialById(unlockId, 0)
		conf := gdconf.GetBackGroundMusicById(unlockId)
		if conf == nil {
			continue
		}
		g.GetPd().AddMusicInfo(unlockId)
		musicList := &proto.MusicData{
			GroupId:  conf.GroupID,
			Id:       conf.ID,
			IsPlayed: true,
		}
		rsp.UnlockedMusicList = append(rsp.UnlockedMusicList, musicList)
	}
	g.Send(cmd.UnlockBackGroundMusicScRsp, rsp)
}

func PlayBackGroundMusicCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PlayBackGroundMusicCsReq)
	db := g.GetPd().GetPhoneData()
	db.CurrentMusicId = req.PlayMusicId
	rsp := &proto.PlayBackGroundMusicScRsp{
		PlayMusicId:    db.CurrentMusicId,
		CurrentMusicId: db.CurrentMusicId,
		Retcode:        0,
	}

	g.Send(cmd.PlayBackGroundMusicScRsp, rsp)
}

func GetPamSkinDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetPamSkinDataScRsp{
		Retcode:           0,
		UnlockedPamSkinId: make([]uint32, 0),
		CurPamSkinId:      g.GetPd().GetCurPamSkin(),
	}

	for id := range g.GetPd().GetUnlockedPamSkin() {
		rsp.UnlockedPamSkinId = append(rsp.UnlockedPamSkinId, id)
	}

	g.Send(cmd.GetPamSkinDataScRsp, rsp)
}

func SelectPamSkinCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SelectPamSkinCsReq)

	rsp := &proto.SelectPamSkinScRsp{
		SelectPamSkinId: req.PamSkinId,
	}
	if !g.GetPd().SetCurPamSkin(req.PamSkinId) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_NOT_EXIST)
	}
	rsp.CurPamSkinId = g.GetPd().GetCurPamSkin()

	g.Send(cmd.SelectPamSkinScRsp, rsp)
}

func TrainPartyGetDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.TrainPartyGetDataScRsp{
		TrainPartyData: &proto.TrainPartyData{
			PassengerInfo: g.GetPd().GetPassengerInfo(),
			// TrainPartyGameInfo: g.GetPd().GetTrainPartyGameInfo(),
			TrainPartyGameInfo: &proto.TrainPartyGameInfo{
				TrainPartyItemInfo: nil,
				TrainPassengerInfo: nil,
				TeamId:             0,
				TrainActionInfo:    nil,
				TrainPartyGridInfo: nil,
			},
			TrainPartyInfo: g.GetPd().GetTrainPartyInfo(),
			CMGMGNOMJFN:    0,
			LAGHAPIKBID:    0,
			RecordId:       6,
		},
		Retcode: 0,
	}
	g.Send(cmd.TrainPartyGetDataScRsp, rsp)
}

func GetTrainVisitorRegisterCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetTrainVisitorRegisterScRsp{
		Retcode:         0,
		VisitorInfoList: g.GetPd().GetVisitorInfoList(),
	}

	g.Send(cmd.GetTrainVisitorRegisterScRsp, rsp)
}

func TrainPartyEnterCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.TrainPartyEnterScRsp{}
	defer g.Send(cmd.TrainPartyEnterScRsp, rsp)
	g.EnterSceneByServerScNotify(1000201, 0, 0, 0)
}

func TrainPartyLeaveCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.TrainPartyLeaveScRsp{}
	defer g.Send(cmd.TrainPartyLeaveScRsp, rsp)
}
