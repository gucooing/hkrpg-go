package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func MusicRhythmDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.MusicRhythmDataScRsp{
		Retcode: 0,
		// MusicLevel:      g.GetPd().GetMusicRhythmLevelList(),
		// UnlockPhaseList: nil,
		// CurLevelId:      0,
		// CurSongId:       0,
		// MusicGroup:      nil,
		// ShowHint:        false,
		// UnlockSongList:  nil,
		// UnlockTrackList: nil,
	}
	g.Send(cmd.MusicRhythmDataScRsp, rsp)
}
