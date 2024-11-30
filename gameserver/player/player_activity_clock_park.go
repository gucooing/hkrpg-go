package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func ClockParkGetInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.ClockParkGetInfoScRsp{
		ParkInfos:   g.GetPd().GetClockParkInfoList(),
		CMAECALDMAN: make([]uint32, 0),
		Retcode:     0,
		Progress:    0,
		GJLJIOICNBE: 0,
		KGHFABADCHE: 0,
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
		NPONPMGNKIE: 0,
		ScriptId:    1,
		JPECEHHAMPE: "114514",
		Retcode:     0,
		AOGMMEFAIFJ: 0,
		PFBOBBMELFB: &proto.AMIGGMENHFA{
			LDDNDPHOGKK: 0,
			LHECJKAMCIH: 0,
			IKDAEHJKBPA: 0,
		},
		JOIPFMCOINI: 0,
		PNOAKGNANBO: 0,
		RogueBuffInfo: &proto.MGGJKPGEGLP{
			BuffList: make([]*proto.MIFEPBDNGGC, 0),
		},
		EJOMHILNCMC: nil,
		DMBNHOFLDFO: &proto.HJLBLONCPML{
			BEBCFIIABLI: 0,
			GFDAPLICKGC: &proto.HJLBLONCPML_LGNHIBENJDK{
				LGNHIBENJDK: true,
			},
		},
		FBELDFBDFAG: 0,
	}
	g.Send(cmd.ClockParkGetOngoingScriptInfoScRsp, rsp)
}
