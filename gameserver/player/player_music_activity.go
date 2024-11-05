package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) MusicRhythmDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.MusicRhythmDataScRsp{
		UnlockSongList:  nil,
		CurSongId:       0,
		UnlockTrackList: nil,
		UnlockPhaseList: nil,
		ShowHint:        false,
		MusicGroup:      nil,
		MusicLevel:      nil,
		CurLevelId:      0,
		Retcode:         0,
	}
	g.Send(cmd.MusicRhythmDataScRsp, rsp)
}

func (g *GamePlayer) MusicRhythmStartLevelCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.MusicRhythmStartLevelCsReq)
	rsp := &proto.MusicRhythmStartLevelScRsp{
		Retcode:     0,
		LevelId:     req.LevelId,
		NJOONPFKHGE: "1",
	}
	g.Send(cmd.MusicRhythmStartLevelScRsp, rsp)
}
