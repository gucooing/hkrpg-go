package player

import (
	"encoding/base64"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) HandlePlayerLoginCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PlayerLoginCsReq)
	logger.Info("[UID:%v]登录的客户端版本是:%s", g.Uid, req.ClientVersion)
	g.Platform = spb.PlatformType(req.Platform)
	rsp := new(proto.PlayerLoginScRsp)
	db := g.GetPd().GetMaterialMap()
	rsp.Stamina = db[model.Stamina] // 还有多久恢复下一个体力
	rsp.ServerTimestampMs = uint64(time.Now().UnixMilli())
	rsp.CurTimezone = 4 // 时区
	rsp.BasicInfo = &proto.PlayerBasicInfo{
		Nickname:   g.GetPd().GetNickname(),
		Level:      g.GetPd().GetLevel(),
		WorldLevel: g.GetPd().GetWorldLevel(),
		Hcoin:      db[model.Hcoin],
		Scoin:      db[model.Scoin],
		Mcoin:      db[model.Mcoin],
		Stamina:    db[model.Stamina],
		Exp:        db[model.Exp],
	}
	rsp.LoginRandom = g.LoginRandom
	g.Send(cmd.PlayerLoginScRsp, rsp)
	g.LoginReady() // 登录准备工作
}

func (g *GamePlayer) SyncClientResVersionCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SyncClientResVersionCsReq)

	rsp := new(proto.SyncClientResVersionScRsp)
	rsp.ResVersion = req.ResVersion

	g.Send(cmd.SyncClientResVersionScRsp, rsp)
}

func (g *GamePlayer) BattlePassInfoNotify() {
	// 战斗通行证信息通知
	notify := &proto.BattlePassInfoNotify{
		// TakenPremiumOptionalReward: 2251799813685246,
		// TakenPremiumExtendedReward: 127,
		// TakenPremiumReward2:        7,
		Exp: 1,
		// TakenPremiumReward1:        2,
		// CurWeekAddExpSum:           8000,
		// TakenFreeExtendedReward:    2,
		// CurBpId:                    5,
		// TakenFreeReward:            6,
		BpTierType: proto.BpTierType_BP_TIER_TYPE_PREMIUM_2,
		Level:      70,
	}
	g.Send(cmd.BattlePassInfoNotify, notify)
}

// 登录通知包
func (g *GamePlayer) LoginNotify() {
	g.Send(cmd.UpdateFeatureSwitchScNotify,
		&proto.UpdateFeatureSwitchScNotify{})
	g.DailyTaskNotify() // 每日刷新事务
	g.Send(cmd.RaidInfoNotify, &proto.RaidInfoNotify{})
	g.BattlePassInfoNotify()
	g.StaminaInfoScNotify()
	g.Send(cmd.GeneralVirtualItemDataNotify,
		&proto.GeneralVirtualItemDataNotify{})
	g.StoryLineInfoScNotify() // 故事线通知包
	g.ContentPackageSyncDataScNotify()

	// g.Send(cmd.SyncServerSceneChangeNotify, &proto.SyncServerSceneChangeNotify{})
	// g.Send(cmd.SyncTurnFoodNotify, &proto.SyncTurnFoodNotify{})
	// g.Send(cmd.ComposeLimitNumCompleteNotify, &proto.ComposeLimitNumCompleteNotify{})
	// g.Send(cmd.NewMailScNotify, &proto.NewMailScNotify{})
	// g.Send(cmd.NewAssistHistoryNotify, &proto.NewAssistHistoryNotify{})

	// g.ServerAnnounceNotify()
	// g.ClientDownloadDataScNotify()
}

// 飘窗通知
func (g *GamePlayer) ServerAnnounceNotify() {
	notify := &proto.ServerAnnounceNotify{AnnounceDataList: make([]*proto.AnnounceData, 0)}
	notify.AnnounceDataList = append(notify.AnnounceDataList, &proto.AnnounceData{
		EndTime:   4294967295,
		ConfigId:  0,
		BeginTime: 1664308800,
	})
	g.Send(cmd.ServerAnnounceNotify, notify)
}

// wind
func (g *GamePlayer) ClientDownloadDataScNotify() {
	// content, _ := os.ReadFile("./data/t.lua")
	// luac := base64.StdEncoding.EncodeToString(content)
	// logger.Info("luac:%s", luac)
	content, _ := base64.StdEncoding.DecodeString("LS0g5a6a5LmJ56Gu6K6k5a+56K+d5qGG55qE5Zue6LCD5Ye95pWwDQpsb2NhbCBmdW5jdGlvbiBvbkRpYWxvZ0Nsb3NlZCgpDQogICAgLS0g5omT5byA5oyH5a6a55qEVVJMDQogICAgQ1MuVW5pdHlFbmdpbmUuQXBwbGljYXRpb24uT3BlblVSTCgiaHR0cHM6Ly9naXRodWIuY29tL2d1Y29vaW5nL2hrcnBnLWdvLXB1YmxpYyIpDQplbmQNCg0KLS0g5pi+56S66Ieq5a6a5LmJ56Gu6K6k5Y+W5raI5o+Q56S65qGG77yM5bm257uR5a6a5Zue6LCD5Ye95pWwDQpDUy5SUEcuQ2xpZW50LkNvbmZpcm1EaWFsb2dVdGlsLlNob3dDdXN0b21Pa0NhbmNlbEhpbnQoDQogICAgIjxjb2xvcj0jRkZGRjAwPuasoui/juadpeWIsGhrcnBnLWdv77yBXG48L2NvbG9yPjxjb2xvcj0jRjg5NkZDPuWFjTwvY29sb3I+PGNvbG9yPSNGMTkzRjk+6LS5PC9jb2xvcj48Y29sb3I9I0VBOTBGNj7CtzwvY29sb3I+PGNvbG9yPSNFMzhERjM+5bSpPC9jb2xvcj48Y29sb3I9I0RDOEFGMD7lnY88L2NvbG9yPjxjb2xvcj0jRDU4N0VEPjo8L2NvbG9yPjxjb2xvcj0jQ0U4NEVBPuaYnzwvY29sb3I+PGNvbG9yPSNDNzgxRTc+56m5PC9jb2xvcj48Y29sb3I9I0MwN0VFND7pk4E8L2NvbG9yPjxjb2xvcj0jQjk3QkUxPumBkzwvY29sb3I+XG48Y29sb3I9I0IyNzhERT7mnKzmnI3liqHlmajlrozlhajlhY3otLnlpoLmnpzmgqjmmK/otK3kubDlvpfliLDnmoTpgqPkuYjmgqjlt7Lnu4/ooqvpqpfkuobvvIE8L2NvbG9yPlxuPGNvbG9yPSNCMjc4REU+R2l0aHVi5byA5rqQ6aG555uuPC9jb2xvcj4iLA0KICAgIG9uRGlhbG9nQ2xvc2VkDQop")
	g.Send(cmd.ClientDownloadDataScNotify, &proto.ClientDownloadDataScNotify{
		DownloadData: &proto.ClientDownloadData{
			Version: 1,
			Time:    1935664461,
			Data:    content,
		},
	},
	)
}

func (g *GamePlayer) Dump() {
	content, _ := os.ReadFile("./data/dump.lua")
	g.Send(cmd.ClientDownloadDataScNotify, &proto.ClientDownloadDataScNotify{
		DownloadData: &proto.ClientDownloadData{
			Version: 1,
			Time:    1935664461,
			Data:    content,
		},
	})
}

// 1.检查是否有好友再redis里
// 2.任务检查
// 3.检查redis里是否有私人邮件
func (g *GamePlayer) LoginReady() { // 登录准备工作
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE) // 取消掉战斗状态
	g.GetPd().InspectionRedisAcceptApplyFriend()          // 1.检查是否有好友在redis里
	g.GetPd().NewMissionInfo()                            // 生成任务数据
	g.LoginReadyMission()                                 // 任务检查

	go g.LoginNotify() // 并发通知
}
