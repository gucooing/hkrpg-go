package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) RegisterMessage(cmdId uint16, payloadMsg []byte /*payloadMsg pb.Message*/) {
	switch cmdId {
	case cmd.GetBasicInfoCsReq:
		g.HandleGetBasicInfoCsReq()
	case cmd.GetEnteredSceneCsReq:
		g.HandleGetEnteredSceneCsReq(payloadMsg)
	case cmd.QueryProductInfoCsReq:
		g.HandleQueryProductInfoCsReq(payloadMsg)
	case cmd.GetRogueHandbookDataCsReq:
		g.HandleGetRogueHandbookDataCsReq(payloadMsg) // 获取帮助手册
	case cmd.GetChatEmojiListCsReq:
		g.HandleGetChatEmojiListCsReq(payloadMsg) // 获取聊天表情
	case cmd.GetJukeboxDataCsReq:
		g.HandleGetJukeboxDataCsReq(payloadMsg) // 点歌？
	case cmd.GetPhoneDataCsReq:
		g.HandleGetPhoneDataCsReq(payloadMsg) // 获取手机信息?
	case cmd.TextJoinQueryCsReq:
		g.TextJoinQueryCsReq() //
	// 登录
	case cmd.PlayerLoginCsReq:
		g.HandlePlayerLoginCsReq(payloadMsg) // 玩家登录请求 第二个登录包
	case cmd.PlayerLoginFinishCsReq:
		g.HandlePlayerLoginFinishCsReq(payloadMsg) // 登录完成包
	case cmd.PlayerLogoutCsReq:
		g.PlayerLogoutCsReq() // 客户端退出游戏通知
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
	case cmd.QuitLineupCsReq:
		g.QuitLineupCsReq(payloadMsg) // 角色离队
	// 角色管理
	case cmd.GetHeroBasicTypeInfoCsReq:
		g.HandleGetHeroBasicTypeInfoCsReq(payloadMsg) // 请求主角基本信息
	case cmd.GetAvatarDataCsReq:
		g.HandleGetAvatarDataCsReq(payloadMsg) // 请求全部角色信息
	case cmd.RankUpAvatarCsReq:
		g.RankUpAvatarCsReq(payloadMsg) // 提高角色命座
	case cmd.AvatarExpUpCsReq:
		g.AvatarExpUpCsReq(payloadMsg) // 角色升级
	case cmd.PromoteAvatarCsReq:
		g.PromoteAvatarCsReq(payloadMsg) // 角色突破
	case cmd.UnlockSkilltreeCsReq:
		g.UnlockSkilltreeCsReq(payloadMsg) // 行迹升级
	case cmd.TakePromotionRewardCsReq:
		g.TakePromotionRewardCsReq(payloadMsg) // 领取角色突破奖励
	// 光锥
	case cmd.DressAvatarCsReq:
		g.DressAvatarCsReq(payloadMsg) // 角色光锥装备
	case cmd.ExpUpEquipmentCsReq:
		g.ExpUpEquipmentCsReq(payloadMsg) // 光锥升级
	case cmd.RankUpEquipmentCsReq:
		g.RankUpEquipmentCsReq(payloadMsg) // 光锥叠影
	case cmd.PromoteEquipmentCsReq:
		g.PromoteEquipmentCsReq(payloadMsg) // 光锥突破
	// 圣遗物
	case cmd.DressRelicAvatarCsReq:
		g.DressRelicAvatarCsReq(payloadMsg) // 圣遗物装备
	case cmd.ExpUpRelicCsReq:
		g.ExpUpRelicCsReq(payloadMsg) // 圣遗物升级
	// 场景
	case cmd.GetSceneMapInfoCsReq:
		g.HanldeGetSceneMapInfoCsReq(payloadMsg) // 获取地图信息
	case cmd.GetCurSceneInfoCsReq:
		g.HandleGetCurSceneInfoCsReq(payloadMsg) // 获取场景信息(关键包)
	case cmd.SceneEntityMoveCsReq:
		g.SceneEntityMoveCsReq(payloadMsg) // 场景实体移动
	case cmd.EnterSceneCsReq:
		g.EnterSceneCsReq(payloadMsg) // 场景传送
	case cmd.GetUnlockTeleportCsReq:
		g.GetUnlockTeleportCsReq() // 获取解锁的传送点
	case cmd.InteractPropCsReq:
		g.InteractPropCsReq(payloadMsg) // 实体交互
	// 战斗
	case cmd.SceneCastSkillCsReq:
		g.SceneCastSkillCsReq(payloadMsg) // 场景开启战斗
	case cmd.PVEBattleResultCsReq:
		g.PVEBattleResultCsReq(payloadMsg) // PVE战斗结算
	case cmd.GetRogueInfoCsReq:
		g.GetRogueInfoCsReq(payloadMsg) // 获取模拟宇宙
	case cmd.GetRogueScoreRewardInfoCsReq:
		g.GetRogueScoreRewardInfoCsReq()
	case cmd.StartRogueCsReq:
		g.StartRogueCsReq(payloadMsg) // 模拟宇宙,启动!
	case cmd.LeaveRogueCsReq:
		g.LeaveRogueCsReq(payloadMsg) // 模拟宇宙撤离请求
	case cmd.QuitRogueCsReq:
		g.QuitRogueCsReq(payloadMsg) // 模拟宇宙结算请求
	case cmd.GetRogueTalentInfoCsReq:
		g.GetRogueTalentInfoCsReq() // 获取天赋信息
	case cmd.StartCocoonStageCsReq:
		g.StartCocoonStageCsReq(payloadMsg) // 副本/周本等
	case cmd.GetChallengeCsReq:
		g.HandleGetChallengeCsReq(payloadMsg) // 获取忘却之庭挑战完成信息
	case cmd.StartChallengeCsReq:
		g.StartChallengeCsReq(payloadMsg) // 忘却之庭,启动!
	case cmd.GetCurChallengeCsReq:
		g.GetCurChallengeCsReq(payloadMsg) // 获取忘却之庭状态
	case cmd.LeaveChallengeCsReq:
		g.LeaveChallengeCsReq() // 退出忘却之庭
	// 背包
	case cmd.GetBagCsReq:
		g.HandleGetBagCsReq(payloadMsg) // 获取背包物品
	// 交易
	case cmd.GetShopListCsReq:
		g.GetShopListCsReq(payloadMsg) // 获取商店物品列表
	case cmd.ExchangeHcoinCsReq:
		g.ExchangeHcoinCsReq(payloadMsg) // 梦华兑换
	case cmd.ExchangeRogueRewardKeyCsReq:
		g.ExchangeRogueRewardKeyCsReq(payloadMsg) // 储存沉浸器
	case cmd.BuyGoodsCsReq:
		g.BuyGoodsCsReq(payloadMsg) // 商店交易
	// 社交
	case cmd.GetMailCsReq:
		g.GetMailCsReq() // 获取邮件
	// 卡池
	case cmd.GetGachaInfoCsReq:
		g.HandleGetGachaInfoCsReq(payloadMsg) // 获取卡池信息
	case cmd.DoGachaCsReq:
		g.DoGachaCsReq(payloadMsg) // 抽卡请求
	case cmd.GetGachaCeilingCsReq:
		g.HandleGetGachaCeilingCsReq(payloadMsg) // 基础卡池保底达到进度请求
	// 任务
	case cmd.GetQuestDataCsReq:
		g.GetQuestDataCsReq(payloadMsg) // 获取任务信息
	case cmd.GetMissionStatusCsReq:
		g.HandleGetMissionStatusCsReq(payloadMsg)
	// 活动
	case cmd.GetActivityScheduleConfigCsReq:
		g.HandleGetActivityScheduleConfigCsReq(payloadMsg) // 活动排期请求
	case cmd.GetDailyActiveInfoCsReq:
		g.GetDailyActiveInfoCsReq(payloadMsg) // 每日任务
	case cmd.GetLoginActivityCsReq:
		g.GetLoginActivityCsReq() // 登录活动完成情况
	case cmd.GetTrialActivityDataCsReq:
		g.GetTrialActivityDataCsReq() // 角色试用完成情况
	case cmd.StartTrialActivityCsReq:
		g.StartTrialActivityCsReq(payloadMsg) // 角色试用
	case cmd.TakeLoginActivityRewardCsReq:
		g.TakeLoginActivityRewardCsReq(payloadMsg) // 领取登录活动奖励
	case cmd.TakeTrialActivityRewardCsReq:
		g.TakeTrialActivityRewardCsReq(payloadMsg) // 角色试用奖励领取
	// 基础
	case cmd.SetClientPausedCsReq:
		g.SetClientPausedCsReq() // 客户端暂停请求
	case cmd.PlayerHeartBeatCsReq:
		g.HandlePlayerHeartBeatCsReq(payloadMsg) // 心跳包
	case cmd.SyncClientResVersionCsReq:
		g.SyncClientResVersionCsReq(payloadMsg) // 版本同步
	case cmd.GetAssistHistoryCsReq:
		g.HandleGetAssistHistoryCsReq() // 漫游签证
	case cmd.SetHeadIconCsReq:
		g.SetHeadIconCsReq(payloadMsg) // 切换头像
	case cmd.SetHeroBasicTypeCsReq:
		g.SetHeroBasicTypeCsReq(payloadMsg) // 切换主角类型
	case cmd.SetNicknameCsReq:
		g.SetNicknameCsReq(payloadMsg) // 修改昵称请求
	case cmd.SetGameplayBirthdayCsReq:
		g.SetGameplayBirthdayCsReq(payloadMsg) // 修改生日请求
	case cmd.SetSignatureCsReq:
		g.SetSignatureCsReq(payloadMsg) // 简介修改请求
	case cmd.GetPlayerBoardDataCsReq:
		g.HandleGetPlayerBoardDataCsReq(payloadMsg) // 获取角色名片页信息
	case cmd.GetFarmStageGachaInfoCsReq:
		g.GetFarmStageGachaInfoCsReq(payloadMsg) // 获取怪物刷新情况
	// 好友
	case cmd.GetFriendLoginInfoCsReq:
		g.HandleGetFriendLoginInfoCsReq(payloadMsg) // 获取好友信息列表
	case cmd.GetFriendListInfoCsReq:
		g.GetFriendListInfoCsReq() // 获取好友信息
	case cmd.GetPrivateChatHistoryCsReq:
		g.GetPrivateChatHistoryCsReq(payloadMsg) // 获取私聊记录
	case cmd.SendMsgCsReq:
		g.SendMsgCsReq(payloadMsg) // 发送聊天信息
	// 成就
	case cmd.GetArchiveDataCsReq:
		g.HandleGetArchiveDataCsReq() // 获取收集
	case cmd.GetUpdatedArchiveDataCsReq:
		g.GetUpdatedArchiveDataCsReq() // 更新收集
	// NPC
	case cmd.GetFirstTalkNpcCsReq:
		g.GetFirstTalkNpcCsReq()
	case cmd.GetNpcTakenRewardCsReq:
		g.GetNpcTakenRewardCsReq(payloadMsg) // NPC对话
	case cmd.GetFirstTalkByPerformanceNpcCsReq:
		g.GetFirstTalkByPerformanceNpcCsReq(payloadMsg) // NPC商店
	default:
		logger.Debug("C --> S error router: %v", cmdId)
	}
	return
}

func (g *GamePlayer) GMRegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {
	case cmd.GmGive:
		g.GmGive(payloadMsg) // 获取物品
	case cmd.GmWorldLevel:
		g.GmWorldLevel(payloadMsg) // 设置世界等级
	}
}

func (g *GamePlayer) GateRegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLoginReq:
		g.PlayerLoginReq(payloadMsg) // gate玩家登录通知
	case cmd.PlayerToGameByGateReq:
		g.PlayerToGameByGateReq(payloadMsg)
	}
}
