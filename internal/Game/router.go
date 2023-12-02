package Game

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
)

func (g *Game) RegisterMessage(cmdId uint16, payloadMsg []byte /*payloadMsg pb.Message*/) {
	switch cmdId {
	case cmd.GetBasicInfoCsReq:
		g.HandleGetBasicInfoCsReq(payloadMsg)
	case cmd.GetPlayerBoardDataCsReq:
		g.HandleGetPlayerBoardDataCsReq(payloadMsg)
	case cmd.GetHeroBasicTypeInfoCsReq:
		g.HandleGetHeroBasicTypeInfoCsReq(payloadMsg)
	case cmd.GetBagCsReq:
		g.HandleGetBagCsReq(payloadMsg)
	case cmd.GetAvatarDataCsReq:
		g.HandleGetAvatarDataCsReq(payloadMsg)
	case cmd.GetActivityScheduleConfigCsReq:
		g.HandleGetActivityScheduleConfigCsReq(payloadMsg)
	case cmd.GetCurChallengeCsReq:
		g.GetCurChallengeCsReq(payloadMsg)
	case cmd.GetMissionStatusCsReq:
		g.HandleGetMissionStatusCsReq(payloadMsg)
	case cmd.GetEnteredSceneCsReq:
		g.HandleGetEnteredSceneCsReq(payloadMsg)
	case cmd.GetArchiveDataCsReq:
		g.HandleGetArchiveDataCsReq(payloadMsg) // 获取存档
	case cmd.GetGachaInfoCsReq:
		g.HandleGetGachaInfoCsReq(payloadMsg) // 获取卡池信息
	case cmd.QueryProductInfoCsReq:
		g.HandleQueryProductInfoCsReq(payloadMsg)
	case cmd.GetFriendLoginInfoCsReq:
		g.HandleGetFriendLoginInfoCsReq(payloadMsg) // 获取好友登录信息
	case cmd.GetRogueHandbookDataCsReq:
		g.HandleGetRogueHandbookDataCsReq(payloadMsg) // 获取帮助手册
	case cmd.GetChatEmojiListCsReq:
		g.HandleGetChatEmojiListCsReq(payloadMsg) // 获取聊天表情
	case cmd.GetChallengeCsReq:
		g.HandleGetChallengeCsReq(payloadMsg) // 获取挑战id列表
	case cmd.GetJukeboxDataCsReq:
		g.HandleGetJukeboxDataCsReq(payloadMsg) // 点歌？
	case cmd.GetPhoneDataCsReq:
		g.HandleGetPhoneDataCsReq(payloadMsg) // 获取手机信息?
	// 登录
	case cmd.PlayerGetTokenCsReq:
		g.HandlePlayerGetTokenCsReq(payloadMsg) // 获取玩家token请求 第一个登录包
	case cmd.PlayerLoginCsReq:
		g.HandlePlayerLoginCsReq(payloadMsg) // 玩家登录请求 第二个登录包
	case cmd.PlayerLoginFinishCsReq:
		g.HandlePlayerLoginFinishCsReq(payloadMsg) // 登录完成包
	case cmd.PlayerLogoutCsReq:
		// TODO 退出登录
	// 队伍
	case cmd.GetAllLineupDataCsReq:
		g.HandleGetAllLineupDataCsReq(payloadMsg) // 获取队伍信息请求
	case cmd.GetCurLineupDataCsReq:
		g.HandleGetCurLineupDataCsReq(payloadMsg) // 获取当前上场队伍请求
	case cmd.JoinLineupCsReq:
		g.HandleJoinLineupCsReq(payloadMsg) // 更新队伍请求
	case cmd.SwitchLineupIndexCsReq:
		g.HandleSwitchLineupIndexCsReq(payloadMsg) // 上场队伍更新请求
	case cmd.SwapLineupCsReq:
		g.HandleSwapLineupCsReq(payloadMsg) // 队伍角色交换请求
	case cmd.SetLineupNameCsReq:
		g.SetLineupNameCsReq(payloadMsg) // 修改队伍名称
	case cmd.ReplaceLineupCsReq:
		g.ReplaceLineupCsReq(payloadMsg) // 快速入队
	case cmd.ChangeLineupLeaderCsReq:
		g.ChangeLineupLeaderCsReq(payloadMsg) // 切换角色
	// 场景
	case cmd.GetRogueInfoCsReq:
		g.GetRogueInfoCsReq(payloadMsg) // 获取副本库
	case cmd.GetCurSceneInfoCsReq:
		g.HandleGetCurSceneInfoCsReq(payloadMsg) // 获取场景信息(关键包)
	case cmd.SceneEntityMoveCsReq:
		g.SceneEntityMoveCsReq() // 场景实体移动
	case cmd.GetRogueScoreRewardInfoCsReq:
		g.GetRogueScoreRewardInfoCsReq()
	// 交易
	case cmd.GetShopListCsReq:
		g.GetShopListCsReq() // 获取商店物品列表
	// 社交
	case cmd.GetMailCsReq:
		g.GetMailCsReq() // 获取邮件
	// 基础
	case cmd.SetClientPausedCsReq:
		g.SetClientPausedCsReq() // 客户端暂停请求
	case cmd.PlayerHeartBeatCsReq:
		g.HandlePlayerHeartBeatCsReq(payloadMsg) // 心跳包
	case cmd.SyncClientResVersionCsReq:
		g.SyncClientResVersionCsReq(payloadMsg) // 版本同步
	case cmd.SetHeadIconCsReq:
		g.SetHeadIconCsReq(payloadMsg) // 切换头像
	case cmd.SetHeroBasicTypeCsReq:
		g.SetHeroBasicTypeCsReq(payloadMsg) // 切换主角类型
	// 乱七八糟
	case cmd.GetFirstTalkNpcCsReq:
		g.GetFirstTalkNpcCsReq()
	default:
		logger.Error("C --> S error router: %v", cmdId)
	}
	return
}
