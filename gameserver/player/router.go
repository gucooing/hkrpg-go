package player

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

type HandlerFunc func(payloadMsg pb.Message)

type RouteManager struct {
	handlerFuncRouteMap map[uint16]HandlerFunc
}

func (r *RouteManager) initRoute(g *GamePlayer) {
	r.handlerFuncRouteMap = map[uint16]HandlerFunc{
		// 登录
		cmd.PlayerLoginCsReq:       g.HandlePlayerLoginCsReq,       // 玩家登录请求 第二个登录包
		cmd.PlayerLoginFinishCsReq: g.HandlePlayerLoginFinishCsReq, // 登录完成包
		// 队伍
		cmd.GetAllLineupDataCsReq:    g.HandleGetAllLineupDataCsReq,  // 获取队伍信息请求
		cmd.GetCurLineupDataCsReq:    g.HandleGetCurLineupDataCsReq,  // 获取当前上场队伍请求
		cmd.GetLineupAvatarDataCsReq: g.GetLineupAvatarDataCsReq,     // 获取队伍
		cmd.JoinLineupCsReq:          g.HandleJoinLineupCsReq,        // 更新队伍请求
		cmd.SwitchLineupIndexCsReq:   g.HandleSwitchLineupIndexCsReq, // 上场队伍更新请求
		// cmd.SwapLineupCsReq:          g.HandleSwapLineupCsReq,        // 队伍角色交换请求 // 2.5.0 遗弃
		cmd.SetLineupNameCsReq:      g.SetLineupNameCsReq,      // 修改队伍名称
		cmd.ReplaceLineupCsReq:      g.ReplaceLineupCsReq,      // 快速入队
		cmd.ChangeLineupLeaderCsReq: g.ChangeLineupLeaderCsReq, // 切换角色
		cmd.QuitLineupCsReq:         g.QuitLineupCsReq,         // 角色离队
		// 角色管理
		cmd.GetMultiPathAvatarInfoCsReq: g.GetMultiPathAvatarInfoCsReq, // 请求多命途角色基本信息
		cmd.GetAvatarDataCsReq:          g.HandleGetAvatarDataCsReq,    // 请求全部角色信息
		cmd.RankUpAvatarCsReq:           g.RankUpAvatarCsReq,           // 提高角色命座
		cmd.AvatarExpUpCsReq:            g.AvatarExpUpCsReq,            // 角色升级
		cmd.PromoteAvatarCsReq:          g.PromoteAvatarCsReq,          // 角色突破
		cmd.UnlockSkilltreeCsReq:        g.UnlockSkilltreeCsReq,        // 行迹升级
		cmd.TakePromotionRewardCsReq:    g.TakePromotionRewardCsReq,    // 领取角色突破奖励
		cmd.UnlockAvatarPathCsReq:       g.UnlockAvatarPathCsReq,       // 来自客户端的解锁命途要求
		// 光锥
		cmd.DressAvatarCsReq:      g.DressAvatarCsReq,      // 角色光锥装备
		cmd.TakeOffEquipmentCsReq: g.TakeOffEquipmentCsReq, // 卸下光锥
		cmd.ExpUpEquipmentCsReq:   g.ExpUpEquipmentCsReq,   // 光锥升级
		cmd.RankUpEquipmentCsReq:  g.RankUpEquipmentCsReq,  // 光锥叠影
		cmd.PromoteEquipmentCsReq: g.PromoteEquipmentCsReq, // 光锥突破
		// 圣遗物
		cmd.RelicRecommendCsReq:       g.RelicRecommendCsReq,       // 获取推荐圣遗物
		cmd.DressRelicAvatarCsReq:     g.DressRelicAvatarCsReq,     // 圣遗物装备
		cmd.TakeOffRelicCsReq:         g.TakeOffRelicCsReq,         // 卸下圣遗物
		cmd.ExpUpRelicCsReq:           g.ExpUpRelicCsReq,           // 圣遗物升级
		cmd.RelicAvatarRecommendCsReq: g.RelicAvatarRecommendCsReq, // 查看圣遗物推荐角色
		// 场景
		cmd.GetEnteredSceneCsReq:        g.HandleGetEnteredSceneCsReq, // 获取当前场景id
		cmd.GetSceneMapInfoCsReq:        g.HanldeGetSceneMapInfoCsReq, // 获取地图信息
		cmd.GetCurSceneInfoCsReq:        g.HandleGetCurSceneInfoCsReq, // 获取场景信息(关键包)
		cmd.SceneEntityMoveCsReq:        g.SceneEntityMoveCsReq,       // 场景实体移动
		cmd.EnterSceneCsReq:             g.EnterSceneCsReq,            // 场景传送
		cmd.GetUnlockTeleportCsReq:      g.GetUnlockTeleportCsReq,     // 获取解锁的传送点
		cmd.InteractPropCsReq:           g.InteractPropCsReq,          // 实体交互
		cmd.GroupStateChangeCsReq:       g.GroupStateChangeCsReq,      // 组状态变更
		cmd.DeployRotaterCsReq:          g.DeployRotaterCsReq,         // 设置旋转
		cmd.StartWolfBroGameCsReq:       g.StartWolfBroGameCsReq,      // 变身
		cmd.SetGroupCustomSaveDataCsReq: g.SetGroupCustomSaveDataCsReq,
		// 战斗
		cmd.SceneCastSkillCostMpCsReq:    g.SceneCastSkillCostMpCsReq,    // 技能使用
		cmd.SceneCastSkillCsReq:          g.SceneCastSkillCsReq,          // 场景开启战斗
		cmd.SetTurnFoodSwitchCsReq:       g.SetTurnFoodSwitchCsReq,       // 使用消耗品buff
		cmd.RefreshTriggerByClientCsReq:  g.RefreshTriggerByClientCsReq,  // 领域buff
		cmd.PVEBattleResultCsReq:         g.PVEBattleResultCsReq,         // PVE战斗结算
		cmd.StartCocoonStageCsReq:        g.StartCocoonStageCsReq,        // 副本/周本等
		cmd.ActivateFarmElementCsReq:     g.ActivateFarmElementCsReq,     // 虚影战斗
		cmd.ReEnterLastElementStageCsReq: g.ReEnterLastElementStageCsReq, // 虚影战斗再来一次
		cmd.DeactivateFarmElementCsReq:   g.DeactivateFarmElementCsReq,   // 虚影
		cmd.SceneEnterStageCsReq:         g.SceneEnterStageCsReq,         // 场景直接发起战斗
		cmd.GetRaidInfoCsReq:             g.GetRaidInfoCsReq,             // 获取raid
		cmd.StartRaidCsReq:               g.StartRaidCsReq,               // 拓境探游
		cmd.LeaveRaidCsReq:               g.LeaveRaidCsReq,               // 退出拓境探游
		// 模拟宇宙
		cmd.GetRogueHandbookDataCsReq:           g.GetRogueHandbookDataCsReq,           // 模拟宇宙图鉴
		cmd.GetRogueScoreRewardInfoCsReq:        g.GetRogueScoreRewardInfoCsReq,        // 获取模拟宇宙排期
		cmd.GetRogueInitialScoreCsReq:           g.GetRogueInitialScoreCsReq,           // 查询模拟宇宙当前分数
		cmd.GetRogueTalentInfoCsReq:             g.GetRogueTalentInfoCsReq,             // 获取天赋信息
		cmd.GetRogueInfoCsReq:                   g.GetRogueInfoCsReq,                   // 获取模拟宇宙
		cmd.StartRogueCsReq:                     g.StartRogueCsReq,                     // 模拟宇宙,启动!
		cmd.LeaveRogueCsReq:                     g.LeaveRogueCsReq,                     // 模拟宇宙撤离请求
		cmd.QuitRogueCsReq:                      g.QuitRogueCsReq,                      // 模拟宇宙结算请求
		cmd.HandleRogueCommonPendingActionCsReq: g.HandleRogueCommonPendingActionCsReq, // 模拟宇宙常见操作请求
		cmd.EnterRogueMapRoomCsReq:              g.EnterRogueMapRoomCsReq,              // 模拟宇宙进入下一场景
		cmd.GetRogueBuffEnhanceInfoCsReq:        g.GetRogueBuffEnhanceInfoCsReq,        // 获取模拟宇宙buff信息
		// cmd.EnhanceRogueBuffCsReq:g.EnhanceRogueBuffCsReq,// 强化buff
		cmd.GetRogueAdventureRoomInfoCsReq: g.GetRogueAdventureRoomInfoCsReq, // 模拟宇宙冒险
		// 差分宇宙
		cmd.RogueTournQueryCsReq:                  g.RogueTournQueryCsReq,                  // 获取差分宇宙信息
		cmd.RogueTournGetPermanentTalentInfoCsReq: g.RogueTournGetPermanentTalentInfoCsReq, // 获取差分宇宙灵感回路
		cmd.RogueTournStartCsReq:                  g.RogueTournStartCsReq,                  // 差分宇宙.启动!
		cmd.RogueTournGetMiscRealTimeDataCsReq:    g.RogueTournGetMiscRealTimeDataCsReq,    // 获取差分宇宙实时信息
		cmd.RogueTournEnterCsReq:                  g.RogueTournEnterCsReq,                  // 继续进度
		cmd.RogueTournSettleCsReq:                 g.RogueTournSettleCsReq,                 // 结束并结算
		cmd.RogueTournEnterRoomCsReq:              g.RogueTournEnterRoomCsReq,              // 差分宇宙进入下一场景
		// 忘却之庭
		cmd.GetChallengeGroupStatisticsCsReq: g.GetChallengeGroupStatisticsCsReq, // 获取忘却之庭状态
		cmd.GetChallengeCsReq:                g.HandleGetChallengeCsReq,          // 获取忘却之庭挑战完成信息
		cmd.StartChallengeCsReq:              g.StartChallengeCsReq,              // 忘却之庭,启动!
		cmd.GetCurChallengeCsReq:             g.GetCurChallengeCsReq,             // 获取忘却之庭状态
		cmd.LeaveChallengeCsReq:              g.LeaveChallengeCsReq,              // 退出忘却之庭
		cmd.TakeChallengeRewardCsReq:         g.TakeChallengeRewardCsReq,         // 忘却之庭领取奖励
		cmd.RestartChallengePhaseCsReq:       g.RestartChallengePhaseCsReq,       // 重新挑战忘却之庭
		// 末日之影
		cmd.StartPartialChallengeCsReq:    g.StartPartialChallengeCsReq,    // 末日幻影,二次启动!
		cmd.EnterChallengeNextPhaseCsReq:  g.EnterChallengeNextPhaseCsReq,  // 前往下一节点
		cmd.GetFriendChallengeLineupCsReq: g.GetFriendChallengeLineupCsReq, // 获取好友通关阵容
		// 背包
		cmd.GetBagCsReq:               g.HandleGetBagCsReq,         // 获取背包物品
		cmd.DestroyItemCsReq:          g.DestroyItemCsReq,          // 销毁物品
		cmd.SellItemCsReq:             g.SellItemCsReq,             // 光锥销毁
		cmd.UseItemCsReq:              g.UseItemCsReq,              // 物品使用
		cmd.ComposeItemCsReq:          g.ComposeItemCsReq,          // 合成
		cmd.ComposeSelectedRelicCsReq: g.ComposeSelectedRelicCsReq, // 遗器合成
		cmd.LockRelicCsReq:            g.LockRelicCsReq,            // 圣遗物上锁
		cmd.LockEquipmentCsReq:        g.LockEquipmentCsReq,        // 光锥上锁
		cmd.CancelCacheNotifyCsReq:    g.CancelCacheNotifyCsReq,
		// 交易
		cmd.QueryProductInfoCsReq:       g.QueryProductInfoCsReq,       // 获取交易信息
		cmd.GetShopListCsReq:            g.GetShopListCsReq,            // 获取商店物品列表
		cmd.ExchangeHcoinCsReq:          g.ExchangeHcoinCsReq,          // 梦华兑换
		cmd.ExchangeRogueRewardKeyCsReq: g.ExchangeRogueRewardKeyCsReq, // 储存沉浸器
		cmd.BuyGoodsCsReq:               g.BuyGoodsCsReq,               // 商店交易
		cmd.GetRollShopInfoCsReq:        g.GetRollShopInfoCsReq,        //
		// 好友
		cmd.GetChatEmojiListCsReq:       g.HandleGetChatEmojiListCsReq,   // 获取聊天表情
		cmd.SetDisplayAvatarCsReq:       g.SetDisplayAvatarCsReq,         // 设置展示角色
		cmd.SetAssistAvatarCsReq:        g.SetAssistAvatarCsReq,          // 设置支援角色
		cmd.GetFriendLoginInfoCsReq:     g.HandleGetFriendLoginInfoCsReq, // 获取好友信息列表
		cmd.GetFriendListInfoCsReq:      g.GetFriendListInfoCsReq,        // 获取好友信息
		cmd.GetPrivateChatHistoryCsReq:  g.GetPrivateChatHistoryCsReq,    // 获取私聊记录
		cmd.GetChatFriendHistoryCsReq:   g.GetChatFriendHistoryCsReq,     // 获取正在进行的聊天室
		cmd.SearchPlayerCsReq:           g.SearchPlayerCsReq,             // 查找玩家
		cmd.GetFriendApplyListInfoCsReq: g.GetFriendApplyListInfoCsReq,   // 获取好友申请列表
		cmd.HandleFriendCsReq:           g.HandleFriendCsReq,             // 处理好友申请
		cmd.GetPlayerDetailInfoCsReq:    g.GetPlayerDetailInfoCsReq,      // 获取玩家详细信息
		// 邮件
		cmd.MarkReadMailCsReq:       g.MarkReadMailCsReq,       // 读取邮件
		cmd.GetMailCsReq:            g.GetMailCsReq,            // 获取邮件
		cmd.DelMailCsReq:            g.DelMailCsReq,            // 删除邮件
		cmd.TakeMailAttachmentCsReq: g.TakeMailAttachmentCsReq, // 领取邮件
		// 卡池
		cmd.GetGachaInfoCsReq:          g.HandleGetGachaInfoCsReq,    // 获取卡池信息
		cmd.DoGachaCsReq:               g.DoGachaCsReq,               // 抽卡请求
		cmd.GetGachaCeilingCsReq:       g.HandleGetGachaCeilingCsReq, // 基础卡池保底达到进度请求
		cmd.ExchangeGachaCeilingCsReq:  g.ExchangeGachaCeilingCsReq,  // 300抽保底
		cmd.GetFarmStageGachaInfoCsReq: g.GetFarmStageGachaInfoCsReq, // 获取卡池刷新情况?
		// 任务
		cmd.GetDailyActiveInfoCsReq:        g.GetDailyActiveInfoCsReq, // 每日实训
		cmd.GetMainMissionCustomValueCsReq: g.GetMainMissionCustomValueCsReq,
		cmd.GetMissionEventDataCsReq:       g.GetMissionEventDataCsReq,
		cmd.GetMissionStatusCsReq:          g.HandleGetMissionStatusCsReq,  // 获取任务状态
		cmd.GetMissionDataCsReq:            g.GetMissionDataCsReq,          // 获取任务数据
		cmd.FinishTalkMissionCsReq:         g.FinishTalkMissionCsReq,       // 完成任务
		cmd.FinishCosumeItemMissionCsReq:   g.FinishCosumeItemMissionCsReq, // 完成道具提交任务
		cmd.GetVideoVersionKeyCsReq:        g.GetVideoVersionKeyCsReq,      // 获取key
		cmd.GetSecretKeyInfoCsReq:          g.GetSecretKeyInfoCsReq,        // key
		cmd.FinishItemIdCsReq:              g.FinishItemIdCsReq,            // 对话选项
		// cmd.FinishSectionIdCsReq:           g.FinishSectionIdCsReq,          // 对话完成
		cmd.UpdateTrackMainMissionIdCsReq: g.UpdateTrackMainMissionIdCsReq, //  更改当前任务
		// 活动
		cmd.PlayerReturnInfoQueryCsReq:          g.PlayerReturnInfoQueryCsReq,           // 获取回归信息
		cmd.PlayerReturnTakeRewardCsReq:         g.PlayerReturnTakeRewardCsReq,          // 领取回归横幅奖励
		cmd.PlayerReturnSignCsReq:               g.PlayerReturnSignCsReq,                // 领取回归签到奖励
		cmd.HeliobusActivityDataCsReq:           g.HeliobusActivityDataCsReq,            // 活动数据
		cmd.GetActivityScheduleConfigCsReq:      g.HandleGetActivityScheduleConfigCsReq, // 活动排期请求
		cmd.GetLoginActivityCsReq:               g.GetLoginActivityCsReq,                // 登录活动完成情况
		cmd.GetTrialActivityDataCsReq:           g.GetTrialActivityDataCsReq,            // 角色试用完成情况
		cmd.StartTrialActivityCsReq:             g.StartTrialActivityCsReq,              // 角色试用
		cmd.TakeLoginActivityRewardCsReq:        g.TakeLoginActivityRewardCsReq,         // 领取登录活动奖励
		cmd.TakeTrialActivityRewardCsReq:        g.TakeTrialActivityRewardCsReq,         // 角色试用奖励领取
		cmd.GetTreasureDungeonActivityDataCsReq: g.GetTreasureDungeonActivityDataCsReq,  // 抽象
		// 以太战线
		cmd.GetAetherDivideInfoCsReq:              g.GetAetherDivideInfoCsReq,              // 获取以太战线信息
		cmd.GetAetherDivideChallengeInfoCsReq:     g.GetAetherDivideChallengeInfoCsReq,     // 获取以太通关信息
		cmd.SetAetherDivideLineUpCsReq:            g.SetAetherDivideLineUpCsReq,            // 设置队伍
		cmd.EquipAetherDividePassiveSkillCsReq:    g.EquipAetherDividePassiveSkillCsReq,    // 装备道具
		cmd.ClearAetherDividePassiveSkillCsReq:    g.ClearAetherDividePassiveSkillCsReq,    // 卸载装备
		cmd.AetherDivideTakeChallengeRewardCsReq:  g.AetherDivideTakeChallengeRewardCsReq,  // 领取对决奖励
		cmd.StartAetherDivideChallengeBattleCsReq: g.StartAetherDivideChallengeBattleCsReq, // 开始战斗！
		cmd.StartAetherDivideSceneBattleCsReq:     g.StartAetherDivideSceneBattleCsReq,     // 场景开启战斗
		cmd.StartAetherDivideStageBattleCsReq:     g.StartAetherDivideStageBattleCsReq,     // 路人挑衅进入战斗
		cmd.LeaveAetherDivideSceneCsReq:           g.LeaveAetherDivideSceneCsReq,           // 退出以太战线
		// 练剑游戏
		cmd.GetSwordTrainingDataCsReq:   g.GetSwordTrainingDataCsReq,   // 获取练剑游戏信息
		cmd.SwordTrainingStartGameCsReq: g.SwordTrainingStartGameCsReq, // 开始练剑游戏请求
		// cmd.SwordTrainingLearnSkillCsReq:g.SwordTrainingLearnSkillCsReq,// 领悟剑招请求
		// cmd.SwordTrainingTurnActionCsReq:g.SwordTrainingTurnActionCsReq,// 开始日常训练
		// 音乐游戏
		cmd.MusicRhythmDataCsReq:       g.MusicRhythmDataCsReq,       // 获取音乐游戏信息
		cmd.MusicRhythmStartLevelCsReq: g.MusicRhythmStartLevelCsReq, // 获取音乐游戏关卡配置
		// 基础
		cmd.GetBasicInfoCsReq:              g.HandleGetBasicInfoCsReq,        // 基础信息
		cmd.GetPhoneDataCsReq:              g.HandleGetPhoneDataCsReq,        // 获取手机信息
		cmd.SetClientPausedCsReq:           g.SetClientPausedCsReq,           // 客户端暂停请求
		cmd.SyncClientResVersionCsReq:      g.SyncClientResVersionCsReq,      // 版本同步
		cmd.GetAssistHistoryCsReq:          g.HandleGetAssistHistoryCsReq,    // 漫游签证
		cmd.SetHeadIconCsReq:               g.SetHeadIconCsReq,               // 切换头像
		cmd.SetAvatarPathCsReq:             g.SetAvatarPathCsReq,             // 切换主角类型
		cmd.SetNicknameCsReq:               g.SetNicknameCsReq,               // 修改昵称请求
		cmd.SetGameplayBirthdayCsReq:       g.SetGameplayBirthdayCsReq,       // 修改生日请求
		cmd.SetSignatureCsReq:              g.SetSignatureCsReq,              // 简介修改请求
		cmd.GetPlayerBoardDataCsReq:        g.HandleGetPlayerBoardDataCsReq,  // 获取角色名片页信息
		cmd.GetTutorialCsReq:               g.GetTutorialCsReq,               // 获取新手教程状态
		cmd.GetTutorialGuideCsReq:          g.GetTutorialGuideCsReq,          // 获取教程指南
		cmd.UnlockTutorialCsReq:            g.UnlockTutorialCsReq,            // 教程解锁
		cmd.UnlockTutorialGuideCsReq:       g.UnlockTutorialGuideCsReq,       // 解锁指南
		cmd.FinishTutorialCsReq:            g.FinishTutorialCsReq,            // 完成教程
		cmd.FinishTutorialGuideCsReq:       g.FinishTutorialGuideCsReq,       // 完成指南
		cmd.SetPlayerInfoCsReq:             g.SetPlayerInfoCsReq,             // 新手设置名字
		cmd.PlayerHeartBeatCsReq:           g.HandlePlayerHeartBeatCsReq,     // 玩家ping包
		cmd.GetLevelRewardTakenListCsReq:   g.GetLevelRewardTakenListCsReq,   // 等级奖励领取情况
		cmd.GetLevelRewardCsReq:            g.GetLevelRewardCsReq,            // 领取等级奖励
		cmd.GetSpringRecoverDataCsReq:      g.GetSpringRecoverDataCsReq,      // 恢复
		cmd.SpringRecoverSingleAvatarCsReq: g.SpringRecoverSingleAvatarCsReq, // 回血锚点
		cmd.TakeBpRewardCsReq:              g.TakeBpRewardCsReq,              // 战令奖励领取
		cmd.TakeAllRewardCsReq:             g.TakeAllRewardCsReq,             // 领取全部战令奖励
		cmd.ReserveStaminaExchangeCsReq:    g.ReserveStaminaExchangeCsReq,    // 取出体力
		cmd.SelectChatBubbleCsReq:          g.SelectChatBubbleCsReq,          // 设置聊天主题
		cmd.SelectPhoneThemeCsReq:          g.SelectPhoneThemeCsReq,          // 设置手机壁纸
		cmd.PlayBackGroundMusicCsReq:       g.PlayBackGroundMusicCsReq,       // 设置车厢音乐
		cmd.GetJukeboxDataCsReq:            g.HandleGetJukeboxDataCsReq,      // 获取车厢音乐
		cmd.UnlockBackGroundMusicCsReq:     g.UnlockBackGroundMusicCsReq,     // 解锁车厢音乐
		cmd.TextJoinQueryCsReq:             g.TextJoinQueryCsReq,             // 获取自定义文本
		cmd.TextJoinSaveCsReq:              g.TextJoinSaveCsReq,              // 保存自定义文本
		cmd.TextJoinBatchSaveCsReq:         g.TextJoinBatchSaveCsReq,         // 批量保存自定义文本
		// 成就
		cmd.GetArchiveDataCsReq:        g.HandleGetArchiveDataCsReq,  // 获取收集
		cmd.GetUpdatedArchiveDataCsReq: g.GetUpdatedArchiveDataCsReq, // 更新收集
		cmd.GetQuestDataCsReq:          g.GetQuestDataCsReq,          // 获取成就信息
		// NPC
		cmd.GetFirstTalkNpcCsReq:              g.GetFirstTalkNpcCsReq,
		cmd.GetNpcTakenRewardCsReq:            g.GetNpcTakenRewardCsReq,            // NPC对话
		cmd.GetFirstTalkByPerformanceNpcCsReq: g.GetFirstTalkByPerformanceNpcCsReq, // NPC商店
		cmd.GetNpcMessageGroupCsReq:           g.GetNpcMessageGroupCsReq,           // 获取npc聊天信息
		cmd.FinishPerformSectionIdCsReq:       g.FinishPerformSectionIdCsReq,       // 完成npc聊天
		// cmd.FinishSectionIdCsReq:                 g.FinishSectionIdCsReq,                 // npc聊天任务完成
		cmd.GetNpcStatusCsReq:                    g.GetNpcStatusCsReq,                    // 获取npc聊天状态
		cmd.FinishFirstTalkByPerformanceNpcCsReq: g.FinishFirstTalkByPerformanceNpcCsReq, // 完成对话
		// 乱七八糟
		cmd.GetAuthkeyCsReq: g.GetAuthkeyCsReq,
		// cmd.ClockParkGetInfoCsReq: g.ClockParkGetInfoCsReq, // 获取皮诺康妮时钟广场信息
	}
}

func NewRouteManager(g *GamePlayer) (r *RouteManager) {
	r = new(RouteManager)
	r.initRoute(g)
	return r
}

func (g *GamePlayer) registerMessage(cmdId uint16, payloadMsg pb.Message) {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GAMESERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("uid: %v", g.Uid)
			return
		}
	}()
	if g.Uid == LogMsgPlayer {
		LogMsgRecv(cmdId, payloadMsg)
	}
	handlerFunc, ok := g.RouteManager.handlerFuncRouteMap[cmdId]
	if !ok {
		if g.Uid == LogMsgPlayer {
			logger.Warn("[UID:%v]C --> S no route for msg, cmdId: %s", g.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId))
		}
		return
	}
	handlerFunc(payloadMsg)
	return
}

// 收包
func (g *GamePlayer) RecvMsg() {
	for {
		select {
		case recvMsg, ok := <-g.RecvChan:
			if !ok {
				return
			}
			switch recvMsg.MsgType {
			case Client:
				g.registerMessage(recvMsg.CmdId, recvMsg.PlayerMsg)
			case GmReq:
				g.EnterCommand(recvMsg)
			case DailyTask:
				g.DailyTaskNotify()
			}
		}
	}
}

// 发包
func (g *GamePlayer) SendMsg(cmdId uint16, playerMsg pb.Message) {
	if g.IsClosed {
		return
	}
	g.SendChan <- Msg{
		CmdId:     cmdId,
		MsgType:   Server,
		PlayerMsg: playerMsg,
	}
}
