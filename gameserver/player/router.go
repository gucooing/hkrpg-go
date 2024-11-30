package player

import (
	"encoding/base64"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

type HandlerFunc func(g *GamePlayer, payloadMsg pb.Message)

var handlerFuncRouteMap = map[uint16]HandlerFunc{
	// 登录
	cmd.PlayerLoginCsReq:       HandlePlayerLoginCsReq,       // 玩家登录请求 第二个登录包
	cmd.PlayerLoginFinishCsReq: HandlePlayerLoginFinishCsReq, // 登录完成包
	// 队伍
	cmd.GetAllLineupDataCsReq:    HandleGetAllLineupDataCsReq,  // 获取队伍信息请求
	cmd.GetCurLineupDataCsReq:    HandleGetCurLineupDataCsReq,  // 获取当前上场队伍请求
	cmd.GetLineupAvatarDataCsReq: GetLineupAvatarDataCsReq,     // 获取队伍
	cmd.JoinLineupCsReq:          HandleJoinLineupCsReq,        // 更新队伍请求
	cmd.SwitchLineupIndexCsReq:   HandleSwitchLineupIndexCsReq, // 上场队伍更新请求
	// cmd.SwapLineupCsReq:          HandleSwapLineupCsReq,        // 队伍角色交换请求 // 2.5.0 遗弃
	cmd.SetLineupNameCsReq:      SetLineupNameCsReq,      // 修改队伍名称
	cmd.ReplaceLineupCsReq:      ReplaceLineupCsReq,      // 快速入队
	cmd.ChangeLineupLeaderCsReq: ChangeLineupLeaderCsReq, // 切换角色
	cmd.QuitLineupCsReq:         QuitLineupCsReq,         // 角色离队 // 疑似 2.5.0 遗弃
	cmd.TriggerVoiceCsReq:       TriggerVoiceCsReq,       // 语音触发
	// 角色管理
	cmd.GetMultiPathAvatarInfoCsReq: GetMultiPathAvatarInfoCsReq, // 请求多命途角色基本信息
	cmd.GetAvatarDataCsReq:          HandleGetAvatarDataCsReq,    // 请求全部角色信息
	cmd.RankUpAvatarCsReq:           RankUpAvatarCsReq,           // 提高角色命座
	cmd.AvatarExpUpCsReq:            AvatarExpUpCsReq,            // 角色升级
	cmd.PromoteAvatarCsReq:          PromoteAvatarCsReq,          // 角色突破
	cmd.UnlockSkilltreeCsReq:        UnlockSkilltreeCsReq,        // 行迹升级
	cmd.TakePromotionRewardCsReq:    TakePromotionRewardCsReq,    // 领取角色突破奖励
	cmd.UnlockAvatarPathCsReq:       UnlockAvatarPathCsReq,       // 来自客户端的解锁命途要求
	cmd.SetAvatarPathCsReq:          SetAvatarPathCsReq,          // 命途切换
	// 光锥
	cmd.DressAvatarCsReq:      DressAvatarCsReq,      // 角色光锥装备
	cmd.TakeOffEquipmentCsReq: TakeOffEquipmentCsReq, // 卸下光锥
	cmd.ExpUpEquipmentCsReq:   ExpUpEquipmentCsReq,   // 光锥升级
	cmd.RankUpEquipmentCsReq:  RankUpEquipmentCsReq,  // 光锥叠影
	cmd.PromoteEquipmentCsReq: PromoteEquipmentCsReq, // 光锥突破
	// 圣遗物
	cmd.RelicRecommendCsReq:       RelicRecommendCsReq,       // 获取推荐圣遗物
	cmd.DressRelicAvatarCsReq:     DressRelicAvatarCsReq,     // 圣遗物装备
	cmd.TakeOffRelicCsReq:         TakeOffRelicCsReq,         // 卸下圣遗物
	cmd.ExpUpRelicCsReq:           ExpUpRelicCsReq,           // 圣遗物升级
	cmd.RelicAvatarRecommendCsReq: RelicAvatarRecommendCsReq, // 查看圣遗物推荐角色
	// 场景
	cmd.GetEnteredSceneCsReq:        HandleGetEnteredSceneCsReq,  // 获取当前场景id
	cmd.GetSceneMapInfoCsReq:        HanldeGetSceneMapInfoCsReq,  // 获取地图信息
	cmd.GetCurSceneInfoCsReq:        HandleGetCurSceneInfoCsReq,  // 获取场景信息(关键包)
	cmd.SceneEntityMoveCsReq:        SceneEntityMoveCsReq,        // 场景实体移动
	cmd.EnterSceneCsReq:             EnterSceneCsReq,             // 场景传送
	cmd.GetUnlockTeleportCsReq:      GetUnlockTeleportCsReq,      // 获取解锁的传送点
	cmd.InteractPropCsReq:           InteractPropCsReq,           // 实体交互
	cmd.GroupStateChangeCsReq:       GroupStateChangeCsReq,       // 组状态变更
	cmd.DeployRotaterCsReq:          DeployRotaterCsReq,          // 设置旋转
	cmd.StartWolfBroGameCsReq:       StartWolfBroGameCsReq,       // 变身
	cmd.SetGroupCustomSaveDataCsReq: SetGroupCustomSaveDataCsReq, // 组状态？
	cmd.GetPetDataCsReq:             GetPetDataCsReq,             // 获取🐖信息
	cmd.SummonPetCsReq:              SummonPetCsReq,              // 召唤🐖
	cmd.RecallPetCsReq:              RecallPetCsReq,              // 删除🐖
	// 列车
	cmd.GetPamSkinDataCsReq:          GetPamSkinDataCsReq,          // 获取帕姆服装
	cmd.SelectPamSkinCsReq:           SelectPamSkinCsReq,           // 切换帕姆服装
	cmd.TrainPartyGetDataCsReq:       TrainPartyGetDataCsReq,       // 获取列车派对信息
	cmd.GetTrainVisitorRegisterCsReq: GetTrainVisitorRegisterCsReq, // 获取车厢访客
	cmd.TrainPartyEnterCsReq:         TrainPartyEnterCsReq,         //
	// 战斗
	cmd.SceneCastSkillCostMpCsReq:    SceneCastSkillCostMpCsReq,    // 技能使用
	cmd.SceneCastSkillCsReq:          SceneCastSkillCsReq,          // 场景开启战斗
	cmd.SetTurnFoodSwitchCsReq:       SetTurnFoodSwitchCsReq,       // 使用消耗品buff
	cmd.RefreshTriggerByClientCsReq:  RefreshTriggerByClientCsReq,  // 领域buff
	cmd.PVEBattleResultCsReq:         PVEBattleResultCsReq,         // PVE战斗结算
	cmd.StartCocoonStageCsReq:        StartCocoonStageCsReq,        // 副本/周本等
	cmd.ActivateFarmElementCsReq:     ActivateFarmElementCsReq,     // 虚影战斗
	cmd.ReEnterLastElementStageCsReq: ReEnterLastElementStageCsReq, // 虚影战斗再来一次
	cmd.DeactivateFarmElementCsReq:   DeactivateFarmElementCsReq,   // 虚影
	cmd.SceneEnterStageCsReq:         SceneEnterStageCsReq,         // 场景直接发起战斗
	cmd.GetRaidInfoCsReq:             GetRaidInfoCsReq,             // 获取raid
	cmd.StartRaidCsReq:               StartRaidCsReq,               // 拓境探游
	cmd.LeaveRaidCsReq:               LeaveRaidCsReq,               // 退出拓境探游
	// 模拟宇宙公共方法 Rogue
	cmd.GetRogueHandbookDataCsReq:           GetRogueHandbookDataCsReq,           // 模拟宇宙图鉴
	cmd.CommonRogueQueryCsReq:               CommonRogueQueryCsReq,               // 模拟宇宙其他信息获取
	cmd.TakeRogueEventHandbookRewardCsReq:   TakeRogueEventHandbookRewardCsReq,   // 模拟宇宙图鉴事件奖励领取
	cmd.TakeRogueMiracleHandbookRewardCsReq: TakeRogueMiracleHandbookRewardCsReq, // 模拟宇宙图鉴奇物奖励领取
	// 模拟宇宙 QuestRogue
	cmd.GetRogueScoreRewardInfoCsReq:        GetRogueScoreRewardInfoCsReq,        // 获取模拟宇宙得分
	cmd.GetRogueInitialScoreCsReq:           GetRogueInitialScoreCsReq,           // 查询模拟宇宙当前分数
	cmd.TakeRogueScoreRewardCsReq:           TakeRogueScoreRewardCsReq,           // 模拟宇宙奖励领取
	cmd.GetRogueTalentInfoCsReq:             GetRogueTalentInfoCsReq,             // 获取模拟宇宙技能树
	cmd.GetRogueInfoCsReq:                   GetRogueInfoCsReq,                   // 获取模拟宇宙
	cmd.StartRogueCsReq:                     StartRogueCsReq,                     // 模拟宇宙,启动!
	cmd.LeaveRogueCsReq:                     LeaveRogueCsReq,                     // 模拟宇宙撤离请求
	cmd.QuitRogueCsReq:                      QuitRogueCsReq,                      // 模拟宇宙结算请求
	cmd.HandleRogueCommonPendingActionCsReq: HandleRogueCommonPendingActionCsReq, // 模拟宇宙常见操作请求
	cmd.EnterRogueMapRoomCsReq:              EnterRogueMapRoomCsReq,              // 模拟宇宙进入下一场景
	cmd.GetRogueBuffEnhanceInfoCsReq:        GetRogueBuffEnhanceInfoCsReq,        // 获取模拟宇宙buff信息
	// cmd.EnhanceRogueBuffCsReq:EnhanceRogueBuffCsReq,// 强化buff
	cmd.GetRogueAdventureRoomInfoCsReq: GetRogueAdventureRoomInfoCsReq, // 模拟宇宙冒险
	// 差分宇宙
	cmd.RogueTournQueryCsReq:                  RogueTournQueryCsReq,                  // 获取差分宇宙信息
	cmd.RogueTournGetPermanentTalentInfoCsReq: RogueTournGetPermanentTalentInfoCsReq, // 获取差分宇宙灵感回路
	cmd.RogueTournStartCsReq:                  RogueTournStartCsReq,                  // 差分宇宙.启动!
	cmd.RogueTournGetMiscRealTimeDataCsReq:    RogueTournGetMiscRealTimeDataCsReq,    // 获取差分宇宙实时信息
	cmd.RogueTournEnterCsReq:                  RogueTournEnterCsReq,                  // 继续进度
	cmd.RogueTournSettleCsReq:                 RogueTournSettleCsReq,                 // 结束并结算
	cmd.RogueTournEnterRoomCsReq:              RogueTournEnterRoomCsReq,              // 差分宇宙进入下一场景
	// 忘却之庭
	cmd.GetChallengeGroupStatisticsCsReq: GetChallengeGroupStatisticsCsReq, // 获取忘却之庭状态
	cmd.GetChallengeCsReq:                HandleGetChallengeCsReq,          // 获取忘却之庭挑战完成信息
	cmd.StartChallengeCsReq:              StartChallengeCsReq,              // 忘却之庭,启动!
	cmd.GetCurChallengeCsReq:             GetCurChallengeCsReq,             // 获取忘却之庭状态
	cmd.LeaveChallengeCsReq:              LeaveChallengeCsReq,              // 退出忘却之庭
	cmd.TakeChallengeRewardCsReq:         TakeChallengeRewardCsReq,         // 忘却之庭领取奖励
	cmd.RestartChallengePhaseCsReq:       RestartChallengePhaseCsReq,       // 重新挑战忘却之庭
	// 末日之影
	cmd.StartPartialChallengeCsReq:    StartPartialChallengeCsReq,    // 末日幻影,二次启动!
	cmd.EnterChallengeNextPhaseCsReq:  EnterChallengeNextPhaseCsReq,  // 前往下一节点
	cmd.GetFriendChallengeLineupCsReq: GetFriendChallengeLineupCsReq, // 获取好友通关阵容
	// 背包
	cmd.GetBagCsReq:               HandleGetBagCsReq,         // 获取背包物品
	cmd.DestroyItemCsReq:          DestroyItemCsReq,          // 销毁物品
	cmd.SellItemCsReq:             SellItemCsReq,             // 光锥销毁
	cmd.UseItemCsReq:              UseItemCsReq,              // 物品使用
	cmd.ComposeItemCsReq:          ComposeItemCsReq,          // 合成
	cmd.ComposeSelectedRelicCsReq: ComposeSelectedRelicCsReq, // 遗器合成
	cmd.LockRelicCsReq:            LockRelicCsReq,            // 圣遗物上锁
	cmd.LockEquipmentCsReq:        LockEquipmentCsReq,        // 光锥上锁
	cmd.DiscardRelicCsReq:         DiscardRelicCsReq,         // 删除遗器
	cmd.CancelCacheNotifyCsReq:    CancelCacheNotifyCsReq,
	// 交易
	cmd.QueryProductInfoCsReq:       QueryProductInfoCsReq,       // 获取交易信息
	cmd.GetShopListCsReq:            GetShopListCsReq,            // 获取商店物品列表
	cmd.ExchangeHcoinCsReq:          ExchangeHcoinCsReq,          // 梦华兑换
	cmd.ExchangeRogueRewardKeyCsReq: ExchangeRogueRewardKeyCsReq, // 储存沉浸器
	cmd.BuyGoodsCsReq:               BuyGoodsCsReq,               // 商店交易
	cmd.TakeCityShopRewardCsReq:     TakeCityShopRewardCsReq,     // 商店等级奖励领取
	cmd.GetRollShopInfoCsReq:        GetRollShopInfoCsReq,        //
	// 好友
	cmd.GetChatEmojiListCsReq:       HandleGetChatEmojiListCsReq,   // 获取聊天表情
	cmd.SetDisplayAvatarCsReq:       SetDisplayAvatarCsReq,         // 设置展示角色
	cmd.SetAssistAvatarCsReq:        SetAssistAvatarCsReq,          // 设置支援角色
	cmd.GetFriendLoginInfoCsReq:     HandleGetFriendLoginInfoCsReq, // 获取好友信息列表
	cmd.GetFriendListInfoCsReq:      GetFriendListInfoCsReq,        // 获取好友信息
	cmd.GetPrivateChatHistoryCsReq:  GetPrivateChatHistoryCsReq,    // 获取私聊记录
	cmd.GetChatFriendHistoryCsReq:   GetChatFriendHistoryCsReq,     // 获取正在进行的聊天室
	cmd.SearchPlayerCsReq:           SearchPlayerCsReq,             // 查找玩家
	cmd.GetFriendApplyListInfoCsReq: GetFriendApplyListInfoCsReq,   // 获取好友申请列表
	cmd.HandleFriendCsReq:           HandleFriendCsReq,             // 处理好友申请
	cmd.GetPlayerDetailInfoCsReq:    GetPlayerDetailInfoCsReq,      // 获取玩家详细信息
	// 邮件
	cmd.MarkReadMailCsReq:       MarkReadMailCsReq,       // 读取邮件
	cmd.GetMailCsReq:            GetMailCsReq,            // 获取邮件
	cmd.DelMailCsReq:            DelMailCsReq,            // 删除邮件
	cmd.TakeMailAttachmentCsReq: TakeMailAttachmentCsReq, // 领取邮件
	// 卡池
	cmd.GetGachaInfoCsReq:          HandleGetGachaInfoCsReq,    // 获取卡池信息
	cmd.DoGachaCsReq:               DoGachaCsReq,               // 抽卡请求
	cmd.GetGachaCeilingCsReq:       HandleGetGachaCeilingCsReq, // 基础卡池保底达到进度请求
	cmd.ExchangeGachaCeilingCsReq:  ExchangeGachaCeilingCsReq,  // 300抽保底
	cmd.GetFarmStageGachaInfoCsReq: GetFarmStageGachaInfoCsReq, // 获取卡池刷新情况?
	// 任务
	cmd.GetDailyActiveInfoCsReq:        GetDailyActiveInfoCsReq, // 每日实训
	cmd.GetMainMissionCustomValueCsReq: GetMainMissionCustomValueCsReq,
	cmd.GetMissionEventDataCsReq:       GetMissionEventDataCsReq,
	cmd.GetMissionStatusCsReq:          HandleGetMissionStatusCsReq,  // 获取任务状态
	cmd.GetMissionDataCsReq:            GetMissionDataCsReq,          // 获取任务数据
	cmd.FinishTalkMissionCsReq:         FinishTalkMissionCsReq,       // 完成任务
	cmd.FinishCosumeItemMissionCsReq:   FinishCosumeItemMissionCsReq, // 完成道具提交任务
	cmd.GetVideoVersionKeyCsReq:        GetVideoVersionKeyCsReq,      // 获取key
	cmd.GetSecretKeyInfoCsReq:          GetSecretKeyInfoCsReq,        // key
	cmd.FinishItemIdCsReq:              FinishItemIdCsReq,            // 对话选项
	// cmd.FinishSectionIdCsReq:           FinishSectionIdCsReq,          // 对话完成
	cmd.UpdateTrackMainMissionIdCsReq: UpdateTrackMainMissionIdCsReq, //  更改当前任务
	// 活动
	cmd.PlayerReturnInfoQueryCsReq:          PlayerReturnInfoQueryCsReq,           // 获取回归信息
	cmd.PlayerReturnTakeRewardCsReq:         PlayerReturnTakeRewardCsReq,          // 领取回归横幅奖励
	cmd.PlayerReturnSignCsReq:               PlayerReturnSignCsReq,                // 领取回归签到奖励
	cmd.HeliobusActivityDataCsReq:           HeliobusActivityDataCsReq,            // 活动数据
	cmd.GetActivityScheduleConfigCsReq:      HandleGetActivityScheduleConfigCsReq, // 活动排期请求
	cmd.GetLoginActivityCsReq:               GetLoginActivityCsReq,                // 登录活动完成情况
	cmd.GetTrialActivityDataCsReq:           GetTrialActivityDataCsReq,            // 角色试用完成情况
	cmd.StartTrialActivityCsReq:             StartTrialActivityCsReq,              // 角色试用
	cmd.TakeLoginActivityRewardCsReq:        TakeLoginActivityRewardCsReq,         // 领取登录活动奖励
	cmd.TakeTrialActivityRewardCsReq:        TakeTrialActivityRewardCsReq,         // 角色试用奖励领取
	cmd.GetTreasureDungeonActivityDataCsReq: GetTreasureDungeonActivityDataCsReq,  // 抽象
	// 下面是联机
	cmd.GetCrossInfoCsReq: GetCrossInfoCsReq, // 联机信息
	cmd.LobbyGetInfoCsReq: LobbyGetInfoCsReq, // 获取联机大厅
	// 以太战线
	cmd.GetAetherDivideInfoCsReq:              GetAetherDivideInfoCsReq,              // 获取以太战线信息
	cmd.GetAetherDivideChallengeInfoCsReq:     GetAetherDivideChallengeInfoCsReq,     // 获取以太通关信息
	cmd.SetAetherDivideLineUpCsReq:            SetAetherDivideLineUpCsReq,            // 设置队伍
	cmd.EquipAetherDividePassiveSkillCsReq:    EquipAetherDividePassiveSkillCsReq,    // 装备道具
	cmd.ClearAetherDividePassiveSkillCsReq:    ClearAetherDividePassiveSkillCsReq,    // 卸载装备
	cmd.AetherDivideTakeChallengeRewardCsReq:  AetherDivideTakeChallengeRewardCsReq,  // 领取对决奖励
	cmd.StartAetherDivideChallengeBattleCsReq: StartAetherDivideChallengeBattleCsReq, // 开始战斗！
	cmd.StartAetherDivideSceneBattleCsReq:     StartAetherDivideSceneBattleCsReq,     // 场景开启战斗
	cmd.StartAetherDivideStageBattleCsReq:     StartAetherDivideStageBattleCsReq,     // 路人挑衅进入战斗
	cmd.LeaveAetherDivideSceneCsReq:           LeaveAetherDivideSceneCsReq,           // 退出以太战线
	// 美梦往事
	cmd.ClockParkGetInfoCsReq:              ClockParkGetInfoCsReq,              // 获取美梦往事信息
	cmd.ClockParkStartScriptCsReq:          ClockParkStartScriptCsReq,          // 开始拍戏
	cmd.ClockParkGetOngoingScriptInfoCsReq: ClockParkGetOngoingScriptInfoCsReq, // 获取拍戏信息
	// 练剑游戏
	cmd.GetSwordTrainingDataCsReq:   GetSwordTrainingDataCsReq,   // 获取练剑游戏信息
	cmd.SwordTrainingStartGameCsReq: SwordTrainingStartGameCsReq, // 开始练剑游戏请求
	// cmd.SwordTrainingLearnSkillCsReq:SwordTrainingLearnSkillCsReq,// 领悟剑招请求
	// cmd.SwordTrainingTurnActionCsReq:SwordTrainingTurnActionCsReq,// 开始日常训练
	// 折纸小鸟
	cmd.MatchThreeGetDataCsReq:    MatchThreeGetDataCsReq,    // 请求折纸小鸟信息
	cmd.MatchThreeLevelEndCsReq:   MatchThreeLevelEndCsReq,   // 单人折纸小鸟结算请求
	cmd.MatchThreeSetBirdPosCsReq: MatchThreeSetBirdPosCsReq, // 摆放小鸟
	// 基础
	cmd.GetBasicInfoCsReq:              HandleGetBasicInfoCsReq,        // 基础信息
	cmd.GetPhoneDataCsReq:              HandleGetPhoneDataCsReq,        // 获取手机信息
	cmd.SetClientPausedCsReq:           SetClientPausedCsReq,           // 客户端暂停请求
	cmd.SyncClientResVersionCsReq:      SyncClientResVersionCsReq,      // 版本同步
	cmd.GetAssistHistoryCsReq:          HandleGetAssistHistoryCsReq,    // 漫游签证
	cmd.SetHeadIconCsReq:               SetHeadIconCsReq,               // 切换头像
	cmd.SetNicknameCsReq:               SetNicknameCsReq,               // 修改昵称请求
	cmd.SetGameplayBirthdayCsReq:       SetGameplayBirthdayCsReq,       // 修改生日请求
	cmd.SetSignatureCsReq:              SetSignatureCsReq,              // 简介修改请求
	cmd.GetPlayerBoardDataCsReq:        HandleGetPlayerBoardDataCsReq,  // 获取角色名片页信息
	cmd.GetTutorialCsReq:               GetTutorialCsReq,               // 获取新手教程状态
	cmd.GetTutorialGuideCsReq:          GetTutorialGuideCsReq,          // 获取教程指南
	cmd.UnlockTutorialCsReq:            UnlockTutorialCsReq,            // 教程解锁
	cmd.UnlockTutorialGuideCsReq:       UnlockTutorialGuideCsReq,       // 解锁指南
	cmd.FinishTutorialCsReq:            FinishTutorialCsReq,            // 完成教程
	cmd.FinishTutorialGuideCsReq:       FinishTutorialGuideCsReq,       // 完成指南
	cmd.SetPlayerInfoCsReq:             SetPlayerInfoCsReq,             // 新手设置名字
	cmd.PlayerHeartBeatCsReq:           HandlePlayerHeartBeatCsReq,     // 玩家ping包
	cmd.GetLevelRewardTakenListCsReq:   GetLevelRewardTakenListCsReq,   // 等级奖励领取情况
	cmd.GetLevelRewardCsReq:            GetLevelRewardCsReq,            // 领取等级奖励
	cmd.GetSpringRecoverDataCsReq:      GetSpringRecoverDataCsReq,      // 恢复
	cmd.SpringRecoverSingleAvatarCsReq: SpringRecoverSingleAvatarCsReq, // 回血锚点
	cmd.TakeBpRewardCsReq:              TakeBpRewardCsReq,              // 战令奖励领取
	cmd.TakeAllRewardCsReq:             TakeAllRewardCsReq,             // 领取全部战令奖励
	cmd.ReserveStaminaExchangeCsReq:    ReserveStaminaExchangeCsReq,    // 取出体力
	cmd.SelectChatBubbleCsReq:          SelectChatBubbleCsReq,          // 设置聊天主题
	cmd.SelectPhoneThemeCsReq:          SelectPhoneThemeCsReq,          // 设置手机壁纸
	cmd.PlayBackGroundMusicCsReq:       PlayBackGroundMusicCsReq,       // 设置车厢音乐
	cmd.GetJukeboxDataCsReq:            HandleGetJukeboxDataCsReq,      // 获取车厢音乐
	cmd.UnlockBackGroundMusicCsReq:     UnlockBackGroundMusicCsReq,     // 解锁车厢音乐
	cmd.TextJoinQueryCsReq:             TextJoinQueryCsReq,             // 获取自定义文本
	cmd.TextJoinSaveCsReq:              TextJoinSaveCsReq,              // 保存自定义文本
	cmd.TextJoinBatchSaveCsReq:         TextJoinBatchSaveCsReq,         // 批量保存自定义文本
	// 成就
	cmd.GetArchiveDataCsReq:        HandleGetArchiveDataCsReq,  // 获取收集
	cmd.GetUpdatedArchiveDataCsReq: GetUpdatedArchiveDataCsReq, // 更新收集
	cmd.GetQuestDataCsReq:          GetQuestDataCsReq,          // 获取成就信息
	// NPC
	cmd.GetFirstTalkNpcCsReq:              GetFirstTalkNpcCsReq,
	cmd.GetNpcTakenRewardCsReq:            GetNpcTakenRewardCsReq,            // NPC对话
	cmd.GetFirstTalkByPerformanceNpcCsReq: GetFirstTalkByPerformanceNpcCsReq, // NPC商店
	cmd.GetNpcMessageGroupCsReq:           GetNpcMessageGroupCsReq,           // 获取npc聊天信息
	cmd.FinishPerformSectionIdCsReq:       FinishPerformSectionIdCsReq,       // 完成npc聊天
	// cmd.FinishSectionIdCsReq:                 FinishSectionIdCsReq,                 // npc聊天任务完成
	cmd.GetNpcStatusCsReq:                    GetNpcStatusCsReq,                    // 获取npc聊天状态
	cmd.FinishFirstTalkByPerformanceNpcCsReq: FinishFirstTalkByPerformanceNpcCsReq, // 完成对话
	// 乱七八糟
	cmd.GetAuthkeyCsReq: GetAuthkeyCsReq,
	// cmd.ClockParkGetInfoCsReq: ClockParkGetInfoCsReq, // 获取皮诺康妮时钟广场信息
}

func (g *GamePlayer) registerMessage(cmdId uint16, payloadMsg pb.Message) {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			bin, _ := pb.Marshal(payloadMsg)
			logger.Error("@LogTag(player_panic_%v)@ cmdId:%s b64:%s json:%s\nerr:%s\nstack:%s", g.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId),
				base64.StdEncoding.EncodeToString(bin), protojson.Format(payloadMsg), err, logger.Stack())
			return
		}
	}()
	if g.Uid == LogMsgPlayer {
		g.logPlayerMsg(cmdId, payloadMsg, false)
	}
	handlerFunc, ok := handlerFuncRouteMap[cmdId]
	if !ok {
		if g.Uid == LogMsgPlayer {
			logger.Error("@LogTag(player_no_route_%v)@C --> S no route for msg, cmdId: %s", g.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId))
		}
		return
	}
	handlerFunc(g, payloadMsg)
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
	g.ToSendChan(Msg{
		CmdId:     cmdId,
		MsgType:   Server,
		PlayerMsg: playerMsg,
	})
}

func (g *GamePlayer) logPlayerMsg(cmdId uint16, payloadMsg pb.Message, server bool) {
	name := cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId)
	if BlackCmd != nil &&
		!BlackCmd[name] {
		var b string
		if server {
			b = "C --> S"
		} else {
			b = "S --> C"
		}
		data := protojson.Format(payloadMsg)
		logger.Debug("@LogTag(player_msg_%v)@%s cmd: %s msg: \n%s\n", g.Uid, b, name, data)
	}
}
