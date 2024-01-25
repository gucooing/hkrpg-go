package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/gameserver/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetPlayerDate() {
	var err error
	playerData := new(PlayerData)

	dbPlayer := db.DBASE.QueryAccountUidByFieldPlayer(g.Uid)
	if dbPlayer.PlayerDataPb == nil {
		logger.Info("新账号登录，进入初始化流程")
		playerDataPb := g.NewPlayer(g.Uid)
		g.Player = playerData
		// 保存账号数据
		dbPlayer.AccountUid = g.Uid
		dbPlayer.PlayerDataPb, err = pb.Marshal(playerDataPb)
		if err != nil {
			logger.Error("pb marshal error: %v", err)
		}

		err = db.DBASE.AddDatePlayerFieldByFieldName(dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败")
			return
		}
	} else {
		g.PlayerPb = new(spb.PlayerBasicCompBin)

		err = pb.Unmarshal(dbPlayer.PlayerDataPb, g.PlayerPb)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			return
		}
	}

	if g.Player == nil {
		g.Player = &PlayerData{
			Battle: make(map[uint32]*Battle),
			BattleState: &BattleState{
				ChallengeState: &ChallengeState{},
			},
		}
	}
}

func (g *GamePlayer) HandlePlayerLoginCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PlayerLoginCsReq, payloadMsg)
	req := msg.(*proto.PlayerLoginCsReq)

	logger.Info("登录的系统是:%s", req.SystemVersion)

	rsp := new(proto.PlayerLoginScRsp)
	rsp.Stamina = g.GetItem().MaterialMap[11]
	rsp.ServerTimestampMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.CurTimezone = 8 // 时区
	rsp.BasicInfo = &proto.PlayerBasicInfo{
		Nickname:   g.PlayerPb.Nickname,
		Level:      g.PlayerPb.Level,
		Exp:        g.PlayerPb.Exp,
		Hcoin:      g.GetItem().MaterialMap[1],
		Scoin:      g.GetItem().MaterialMap[2],
		Mcoin:      g.GetItem().MaterialMap[3],
		Stamina:    g.GetItem().MaterialMap[11],
		WorldLevel: g.PlayerPb.WorldLevel,
	}
	// 开启数据定时保存
	// go g.AutoUpDataPlayer()

	g.StaminaInfoScNotify()
	g.Send(cmd.PlayerLoginScRsp, rsp)
}

func (g *GamePlayer) SyncClientResVersionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SyncClientResVersionCsReq, payloadMsg)
	req := msg.(*proto.SyncClientResVersionCsReq)

	rsp := new(proto.SyncClientResVersionScRsp)
	rsp.ClientResVersion = req.ClientResVersion

	g.Send(cmd.SyncClientResVersionScRsp, rsp)
}

// 账号离线
func (g *GamePlayer) PlayerLogoutCsReq() {
	g.KickPlayer()
}
