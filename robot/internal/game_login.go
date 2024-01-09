package internal

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type Game struct {
	EntryId uint32
	Pos     *Vector
	Rot     *Vector
}
type Vector struct {
	X int32
	Y int32
	Z int32
}

func (r *RoBot) PlayerGetTokenCsReq() {
	req := &proto.PlayerGetTokenCsReq{
		AccountUid:   r.AccountUid,
		PlatformType: 3,
		Token:        r.Token,
	}

	r.send(cmd.PlayerGetTokenCsReq, req)
}

func (r *RoBot) PlayerLoginCsReq() {
	req := &proto.PlayerLoginCsReq{
		Signature:   "",
		LoginRandom: 0,
	}

	r.send(cmd.PlayerLoginCsReq, req)
}

func (r *RoBot) PlayerLoginScRsp() {
	req := new(proto.GetBagScRsp)

	r.send(cmd.GetBasicInfoCsReq, req)
	r.send(cmd.GetHeroBasicTypeInfoCsReq, req)
	r.send(cmd.GetBagCsReq, req)
	r.send(cmd.GetMarkItemListCsReq, req)
	r.send(cmd.GetPlayerBoardDataCsReq, req)
	getAvatarDataCsReq := &proto.GetAvatarDataCsReq{
		IsGetAll: true,
	}
	r.send(cmd.GetAvatarDataCsReq, getAvatarDataCsReq)
	r.send(cmd.GetCurAssistCsReq, req)
	r.send(cmd.GetAllLineupDataCsReq, req)
	r.send(cmd.GetAllRedDotDataCsReq, req)
	r.send(cmd.GetAllServerPrefsDataCsReq, req)
	r.send(cmd.GetActivityScheduleConfigCsReq, req)
	r.send(cmd.GetMissionDataCsReq, req)
	r.send(cmd.GetMissionEventDataCsReq, req)
	r.send(cmd.GetQuestDataCsReq, req)
	r.send(cmd.GetCurChallengeCsReq, req)
	r.send(cmd.GetRogueInfoCsReq, req)
	r.send(cmd.ChessRogueQueryCsReq, req)
	r.send(cmd.GetRogueDialogueEventDataCsReq, req)
	syncClientResVersionCsReq := &proto.SyncClientResVersionCsReq{
		ClientResVersion: 6057946,
	}
	r.send(cmd.SyncClientResVersionCsReq, syncClientResVersionCsReq)
	r.send(cmd.DailyFirstMeetPamCsReq, req)
	r.send(cmd.GetBattleCollegeDataCsReq, req)
	r.send(cmd.GetNpcStatusCsReq, req)
	r.send(cmd.GetMainMissionCustomValueCsReq, req)
	r.send(cmd.GetSecretKeyInfoCsReq, req)
	r.send(cmd.GetVideoVersionKeyCsReq, req)
	r.send(cmd.GetCurLineupDataCsReq, req)
	r.send(cmd.GetCurBattleInfoCsReq, req)
	r.send(cmd.GetCurSceneInfoCsReq, req)
	r.send(cmd.HeliobusActivityDataCsReq, req)
	r.send(cmd.GetEnteredSceneCsReq, req)
	r.send(cmd.GetAetherDivideInfoCsReq, req)
	r.send(cmd.PlayerLoginFinishCsReq, req)
	r.send(cmd.GetMainMissionCustomValueCsReq, req)
	r.send(cmd.GetArchiveDataCsReq, req)
	r.send(cmd.GetNpcMessageGroupCsReq, req)
	r.send(cmd.GetGachaInfoCsReq, req)
	r.send(cmd.QueryProductInfoCsReq, req)
	r.send(cmd.GetQuestRecordCsReq, req)
	r.send(cmd.GetFriendLoginInfoCsReq, req)
	r.send(cmd.GetCurAssistCsReq, req)
	r.send(cmd.GetFightActivityDataCsReq, req)
	r.send(cmd.GetMultipleDropInfoCsReq, req)
	r.send(cmd.GetPlayerReturnMultiDropInfoCsReq, req)
	r.send(cmd.GetShareDataCsReq, req)
	r.send(cmd.PlayerReturnInfoQueryCsReq, req)
	r.send(cmd.GetAlleyInfoCsReq, req)
	r.send(cmd.GetAetherDivideChallengeInfoCsReq, req)
	r.send(cmd.GetExpeditionDataCsReq, req)
	r.send(cmd.GetChatEmojiListCsReq, req)
	r.send(cmd.GetChallengeCsReq, req)
	r.send(cmd.GetLoginActivityCsReq, req)
	r.send(cmd.GetRaidInfoCsReq, req)
	r.send(cmd.GetTrialActivityDataCsReq, req)
	r.send(cmd.GetJukeboxDataCsReq, req)
	r.send(cmd.GetMuseumInfoCsReq, req)
	r.send(cmd.GetTrainVisitorRegisterCsReq, req)
	r.send(cmd.GetBoxingClubInfoCsReq, req)
	r.send(cmd.TextJoinQueryCsReq, req)
	r.send(cmd.GetSpringRecoverDataCsReq, req)
	r.send(cmd.GetDailyActiveInfoCsReq, req)
	r.send(cmd.GetLoginChatInfoCsReq, req)
	r.send(cmd.GetChatFriendHistoryCsReq, req)
	r.send(cmd.GetPhoneDataCsReq, req)
	r.send(cmd.GetAllSaveRaidCsReq, req)
	r.send(cmd.GetMailCsReq, req)
	r.send(cmd.GetShopListCsReq, req)
	setNicknameCsReq := &proto.SetNicknameCsReq{
		Nickname: r.AccountName,
	}

	r.send(cmd.SetNicknameCsReq, setNicknameCsReq)
	r.send(cmd.SetNicknameCsReq, setNicknameCsReq)
}

func (r *RoBot) PlayerHeartBeatCsReq() {
	go func() {
		for {
			if r.KcpAddr == "" {
				return
			}
			req := &proto.PlayerHeartbeatCsReq{
				ClientTimeMs: uint64(time.Now().UnixNano() / 1e6),
			}

			r.send(cmd.PlayerHeartBeatCsReq, req)

			time.Sleep(1 * time.Second)
		}
	}()
}

func (r *RoBot) PlayerHeartbeatScRsp(payloadMsg pb.Message) {
	rsp := payloadMsg.(*proto.PlayerHeartbeatScRsp)

	times := rsp.ServerTimeMs - rsp.ClientTimeMs

	logger.Debug("[UID:%v] ping: %v", r.GameUid, times)
}
