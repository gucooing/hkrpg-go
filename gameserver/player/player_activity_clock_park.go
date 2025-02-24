package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func ClockParkGetInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.ClockParkGetInfoScRsp{
		ParkInfos: g.GetPd().GetClockParkInfoList(),
		Retcode:   0,
		Progress:  0,
	}
	g.Send(cmd.ClockParkGetInfoScRsp, rsp)
}

func ClockParkStartScriptCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ClockParkStartScriptCsReq)
	rsp := &proto.ClockParkStartScriptScRsp{
		Retcode:  0,
		ScriptId: req.ScriptId,
	}
	g.Send(cmd.ClockParkStartScriptScRsp, rsp)
}

func ClockParkGetOngoingScriptInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.ClockParkGetOngoingScriptInfoScRsp{
		ScriptId: 1,
		Retcode:  0,
	}
	g.Send(cmd.ClockParkGetOngoingScriptInfoScRsp, rsp)
}
