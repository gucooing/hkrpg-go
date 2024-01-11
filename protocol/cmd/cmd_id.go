package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	SpringRecoverCsReq                                 = 1471
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3738
	SelectChatBubbleScRsp                              = 5184
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5659
	SubmitMonsterResearchActivityMaterialCsReq         = 2620
	EnterAetherDivideSceneCsReq                        = 4883
	HeliobusActivityDataCsReq                          = 5883
	CommonRogueQueryScRsp                              = 5686
	GetRogueShopMiracleInfoCsReq                       = 5684
	ApplyFriendCsReq                                   = 2942
	ComposeLimitNumUpdateNotify                        = 506
	GetCurAssistCsReq                                  = 2977
	ChangeLineupLeaderCsReq                            = 764
	TrialActivityDataChangeScNotify                    = 2679
	AlleyGuaranteedFundsCsReq                          = 4777
	UpdateFeatureSwitchScNotify                        = 43
	FinishQuestCsReq                                   = 938
	ChessRogueRollDiceScRsp                            = 5485
	GetAuthkeyScRsp                                    = 73
	ActivateFarmElementCsReq                           = 1470
	StartAetherDivideChallengeBattleScRsp              = 4819
	SetAetherDivideLineUpCsReq                         = 4855
	StartRogueScRsp                                    = 1884
	UnlockAvatarSkinScNotify                           = 333
	SellItemCsReq                                      = 520
	SwapLineupScRsp                                    = 738
	TakeAllApRewardCsReq                               = 3360
	GetChallengeGroupStatisticsCsReq                   = 1773
	RetcodeNotify                                      = 86
	FinishPerformSectionIdScRsp                        = 2738
	AetherDivideFinishChallengeScNotify                = 4865
	SyncLineupNotify                                   = 749
	UseItemScRsp                                       = 519
	GetFarmStageGachaInfoCsReq                         = 1304
	UseTreasureDungeonItemScRsp                        = 4480
	AcceptMainMissionCsReq                             = 1224
	JoinLineupCsReq                                    = 782
	DelSaveRaidScNotify                                = 2220
	InteractPropScRsp                                  = 1484
	EnableRogueTalentCsReq                             = 1823
	CurTrialActivityScNotify                           = 2611
	UnlockPhoneThemeScNotify                           = 5119
	UpdateRogueAdventureRoomScoreScRsp                 = 5676
	NewAssistHistoryNotify                             = 2927
	AlleyShipUsedCountScNotify                         = 4778
	RecoverAllLineupCsReq                              = 1477
	HeliobusSnsPostScRsp                               = 5860
	SetNicknameCsReq                                   = 92
	SharePunkLordMonsterCsReq                          = 3282
	GetChessRogueBuffEnhanceInfoScRsp                  = 5570
	AlleyPlacingGameCsReq                              = 4755
	TakeRogueEventHandbookRewardCsReq                  = 5618
	GetPlayerReturnMultiDropInfoScRsp                  = 4682
	MuseumTakeCollectRewardCsReq                       = 4365
	GetActivityScheduleConfigCsReq                     = 2682
	DeleteSummonUnitCsReq                              = 1447
	RefreshAlleyOrderScRsp                             = 4745
	FinishChessRogueSubStoryScRsp                      = 5406
	GetCurSceneInfoScRsp                               = 1419
	GetFightActivityDataScRsp                          = 3661
	QuitBattleCsReq                                    = 104
	FinishFirstTalkNpcScRsp                            = 2119
	StartCocoonStageScRsp                              = 1427
	BattleLogReportScRsp                               = 149
	TreasureDungeonFinishScNotify                      = 4461
	ChessRogueEnterScRsp                               = 5595
	SetMissionEventProgressCsReq                       = 1246
	TreasureDungeonDataScNotify                        = 4483
	StrongChallengeActivityBattleEndScNotify           = 6682
	StartAlleyEventCsReq                               = 4742
	ChessRogueCellUpdateNotify                         = 5469
	QuestRecordScNotify                                = 997
	BuyRogueShopMiracleCsReq                           = 5619
	EnterRogueEndlessActivityStageScRsp                = 6084
	StartAetherDivideStageBattleCsReq                  = 4859
	GetLoginActivityCsReq                              = 2683
	UpdateServerPrefsDataCsReq                         = 6182
	StartBoxingClubBattleScRsp                         = 4260
	ChessRogueUpdateReviveInfoScNotify                 = 5562
	GetWaypointCsReq                                   = 483
	ReturnLastTownScRsp                                = 1480
	LeaveRogueScRsp                                    = 1819
	FinishTutorialGuideCsReq                           = 1649
	SpaceZooExchangeItemScRsp                          = 6755
	SetSpringRecoverConfigScRsp                        = 1444
	GetAlleyInfoCsReq                                  = 4783
	GetAllSaveRaidScRsp                                = 2245
	FightTreasureDungeonMonsterScRsp                   = 4445
	FinishCosumeItemMissionCsReq                       = 1255
	ExchangeStaminaScRsp                               = 40
	ChessRogueNousDiceSurfaceUnlockNotify              = 5520
	FinishChapterScNotify                              = 4904
	GetPunkLordBattleRecordScRsp                       = 3277
	GetShopListScRsp                                   = 1561
	LogisticsScoreRewardSyncInfoScNotify               = 4756
	GetStageLineupScRsp                                = 761
	LockRelicCsReq                                     = 573
	EnterSceneCsReq                                    = 1417
	SetClientPausedCsReq                               = 1422
	PlayerLogoutCsReq                                  = 4
	SpaceZooMutateCsReq                                = 6782
	BattlePassInfoNotify                               = 3083
	ReserveStaminaExchangeScRsp                        = 81
	MarkItemCsReq                                      = 553
	StartFinishMainMissionScNotify                     = 1206
	FinishAeonDialogueGroupCsReq                       = 1821
	HeliobusSnsUpdateScNotify                          = 5849
	UpdateFloorSavedValueNotify                        = 1491
	TrainVisitorBehaviorFinishScRsp                    = 3761
	LockEquipmentCsReq                                 = 582
	AetherDivideTainerInfoScNotify                     = 4887
	TextJoinBatchSaveCsReq                             = 3882
	ExpUpRelicScRsp                                    = 552
	AlleyShipUnlockScNotify                            = 4708
	DeactivateFarmElementScRsp                         = 1425
	ChessRogueEnterNextLayerScRsp                      = 5448
	SyncRogueFinishScNotify                            = 1840
	HeliobusSnsReadCsReq                               = 5804
	ChessRogueNousEditDiceCsReq                        = 5437
	LeaveTrialActivityScRsp                            = 2658
	SetGameplayBirthdayScRsp                           = 71
	GroupStateChangeScNotify                           = 1466
	SetSignatureCsReq                                  = 2838
	EquipAetherDividePassiveSkillCsReq                 = 4840
	GetLevelRewardTakenListCsReq                       = 80
	SetLanguageScRsp                                   = 87
	InteractTreasureDungeonGridScRsp                   = 4492
	GetRndOptionCsReq                                  = 3483
	ChessRogueNousDiceUpdateNotify                     = 5537
	TakeQuestRewardCsReq                               = 904
	ChessRoguePickAvatarScRsp                          = 5453
	LeaveAetherDivideSceneScRsp                        = 4884
	MuseumRandomEventStartScNotify                     = 4320
	RemoveStuffFromAreaScRsp                           = 4319
	SetGenderScRsp                                     = 56
	BoxingClubChallengeUpdateScNotify                  = 4238
	SyncRogueCommonVirtualItemInfoScNotify             = 5667
	GetBasicInfoScRsp                                  = 67
	AetherDivideSpiritExpUpCsReq                       = 4890
	GetFirstTalkByPerformanceNpcScRsp                  = 2109
	TrainRefreshTimeNotify                             = 3782
	ChooseBoxingClubStageOptionalBuffScRsp             = 4252
	TakeCityShopRewardScRsp                            = 1560
	ChessRogueUpdateUnlockLevelScNotify                = 5592
	ReviveRogueAvatarCsReq                             = 1820
	SetTurnFoodSwitchScRsp                             = 522
	SceneEntityMoveScNotify                            = 1449
	TakeTrialActivityRewardScRsp                       = 2651
	ClientObjDownloadDataScNotify                      = 10
	ExchangeHcoinScRsp                                 = 590
	DressAvatarSkinCsReq                               = 380
	QueryProductInfoCsReq                              = 16
	TakeChallengeRaidRewardScRsp                       = 2297
	GetReplayTokenCsReq                                = 3583
	ChessRogueNousEnableRogueTalentScRsp               = 5580
	EnhanceCommonRogueBuffCsReq                        = 5690
	EnterStrongChallengeActivityStageScRsp             = 6684
	PromoteEquipmentScRsp                              = 584
	RankUpAvatarScRsp                                  = 340
	SceneEntityMoveScRsp                               = 1461
	AlleyEventEffectNotify                             = 4738
	GetCurBattleInfoScRsp                              = 160
	TriggerVoiceCsReq                                  = 4155
	GameplayCounterUpdateScNotify                      = 1454
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5567
	GiveUpBoxingClubChallengeScRsp                     = 4219
	AcceptedPamMissionExpireScRsp                      = 4061
	RegionStopScNotify                                 = 45
	TakeRogueEndlessActivityPointRewardScRsp           = 6042
	GetChessRogueStoryAeonTalkInfoCsReq                = 5472
	DailyTaskDataScNotify                              = 1219
	BuyNpcStuffScRsp                                   = 4384
	GetMissionDataCsReq                                = 1283
	GetFriendApplyListInfoCsReq                        = 2982
	GetUnlockTeleportCsReq                             = 1481
	ChessRogueFinishCurRoomNotify                      = 5587
	TextJoinQueryCsReq                                 = 3804
	LeaveChallengeScRsp                                = 1760
	TakeRogueAeonLevelRewardCsReq                      = 1874
	SummonPunkLordMonsterScRsp                         = 3219
	SetGroupCustomSaveDataScRsp                        = 1413
	SwitchLineupIndexCsReq                             = 752
	FinishItemIdCsReq                                  = 2782
	MuseumRandomEventSelectCsReq                       = 4380
	PlayerGetTokenScRsp                                = 60
	RefreshTriggerByClientCsReq                        = 1407
	EnterTreasureDungeonScRsp                          = 4464
	GetQuestRecordScRsp                                = 919
	RogueModifierUpdateNotify                          = 5319
	CommonRogueUpdateScNotify                          = 5651
	MissionGroupWarnScNotify                           = 1209
	GetEnteredSceneScRsp                               = 1421
	SyncTaskScRsp                                      = 1242
	PlayerReturnTakeRewardScRsp                        = 4519
	PickRogueAvatarScRsp                               = 1873
	ComposeItemCsReq                                   = 555
	SyncRogueRewardInfoScNotify                        = 1829
	UnlockHeadIconScNotify                             = 2897
	SyncHandleFriendScNotify                           = 2909
	GetGachaInfoCsReq                                  = 1983
	RogueModifierSelectCellScRsp                       = 5382
	HeliobusUpgradeLevelScRsp                          = 5864
	TakeLoginActivityRewardScRsp                       = 2684
	PVEBattleResultCsReq                               = 183
	QuitBattleScNotify                                 = 197
	TakeOffRelicCsReq                                  = 345
	StartRaidScRsp                                     = 2261
	RogueModifierStageStartNotify                      = 5338
	SetHeroBasicTypeCsReq                              = 24
	ComposeSelectedRelicCsReq                          = 546
	GetTutorialGuideScRsp                              = 1684
	GetTreasureDungeonActivityDataScRsp                = 4409
	GetPlayerBoardDataScRsp                            = 2861
	EntityBindPropScRsp                                = 1456
	GetExpeditionDataCsReq                             = 2583
	AlleyTakeEventRewardScRsp                          = 4703
	TakeMailAttachmentScRsp                            = 819
	SetDisplayAvatarCsReq                              = 2882
	ChessRogueQueryAeonDimensionsScRsp                 = 5523
	HeliobusSnsLikeCsReq                               = 5842
	ScenePlaneEventScNotify                            = 1462
	MatchBoxingClubOpponentScRsp                       = 4284
	CurAssistChangedNotify                             = 2922
	PlayerReturnTakePointRewardCsReq                   = 4582
	ChessRogueSelectBpCsReq                            = 5597
	AceAntiCheaterScRsp                                = 89
	GetAllLineupDataCsReq                              = 792
	TakeFightActivityRewardCsReq                       = 3660
	MultipleDropInfoNotify                             = 4660
	GetPunkLordDataCsReq                               = 3252
	UnlockTutorialGuideScRsp                           = 1619
	WaypointShowNewCsNotify                            = 442
	AlleyShopLevelScNotify                             = 4746
	EnterSceneByServerScNotify                         = 1448
	GeneralVirtualItemDataNotify                       = 514
	DeleteFriendScRsp                                  = 2964
	GetTutorialScRsp                                   = 1661
	PlayerGetTokenCsReq                                = 82
	GiveUpBoxingClubChallengeCsReq                     = 4242
	SwitchAetherDivideLineUpSlotScRsp                  = 4892
	HeliobusEnterBattleCsReq                           = 5892
	FinishCurTurnCsReq                                 = 4349
	TakeRogueScoreRewardScRsp                          = 1880
	GetFriendRecommendListInfoScRsp                    = 2992
	PlayerReturnSignCsReq                              = 4561
	GetFriendLoginInfoCsReq                            = 2916
	GetFriendListInfoScRsp                             = 2961
	BoxingClubRewardScNotify                           = 4297
	GetQuestDataCsReq                                  = 983
	PlayerReturnInfoQueryScRsp                         = 4538
	PlayerReturnStartScNotify                          = 4583
	ChallengeLineupNotify                              = 1709
	ExchangeRogueBuffWithMiracleScRsp                  = 5692
	UpdateMechanismBarScNotify                         = 1451
	UpdatePlayerSettingCsReq                           = 99
	GetAssistHistoryScRsp                              = 2903
	PlayBackGroundMusicCsReq                           = 3104
	StartBattleCollegeScRsp                            = 5782
	LeaveAetherDivideSceneCsReq                        = 4804
	SyncRogueCommonPendingActionScNotify               = 5670
	MarkReadMailScRsp                                  = 884
	GetRogueDialogueEventDataScRsp                     = 1871
	MarkChatEmojiScRsp                                 = 3909
	GetServerPrefsDataCsReq                            = 6104
	GetSecretKeyInfoCsReq                              = 35
	ChessRogueSelectCellScRsp                          = 5532
	GetTrainVisitorBehaviorScRsp                       = 3784
	SelectInclinationTextScRsp                         = 2138
	SetIsDisplayAvatarInfoScRsp                        = 2819
	GetRogueInitialScoreCsReq                          = 1814
	LeaveChallengeCsReq                                = 1782
	BuyBpLevelScRsp                                    = 3042
	AcceptMissionEventScRsp                            = 1220
	GetQuestRecordCsReq                                = 942
	StartTimedCocoonStageCsReq                         = 1405
	EnterRogueScRsp                                    = 1860
	LastSpringRefreshTimeNotify                        = 1492
	LogisticsDetonateStarSkiffScRsp                    = 4793
	TakeOffEquipmentCsReq                              = 349
	GetRogueBuffEnhanceInfoScRsp                       = 1846
	EnhanceChessRogueBuffCsReq                         = 5455
	CancelCacheNotifyScRsp                             = 4138
	SendMsgCsReq                                       = 3983
	SwitchLineupIndexScRsp                             = 773
	ChessRogueUpdateBoardScNotify                      = 5499
	GetQuestDataScRsp                                  = 961
	BuyNpcStuffCsReq                                   = 4304
	TakeRogueScoreRewardCsReq                          = 1859
	CancelActivityExpeditionCsReq                      = 2509
	HeliobusEnterBattleScRsp                           = 5859
	GetLevelRewardTakenListScRsp                       = 90
	SetHeadIconScRsp                                   = 2884
	AcceptExpeditionCsReq                              = 2504
	HandleRogueCommonPendingActionCsReq                = 5679
	LogisticsInfoScNotify                              = 4765
	DailyFirstMeetPamScRsp                             = 3484
	ChessRogueQuestFinishNotify                        = 5490
	RechargeSuccNotify                                 = 559
	SyncRogueMapRoomScNotify                           = 1816
	SetCurWaypointCsReq                                = 404
	ChessRogueUpdateAllowedSelectCellScNotify          = 5444
	SceneEntityTeleportCsReq                           = 1426
	ChessRogueCheatRollScRsp                           = 5542
	UnlockTutorialScRsp                                = 1660
	UnlockTeleportNotify                               = 1429
	SetFriendRemarkNameScRsp                           = 2980
	GetFriendListInfoCsReq                             = 2983
	ExchangeHcoinCsReq                                 = 580
	HeliobusLineupUpdateScNotify                       = 5808
	ChessRogueQueryBpCsReq                             = 5558
	TakeChapterRewardCsReq                             = 419
	SelectPhoneThemeScRsp                              = 5142
	SyncRoguePickAvatarInfoScNotify                    = 1875
	TakeRogueMiracleHandbookRewardScRsp                = 5614
	GetSaveLogisticsMapCsReq                           = 4706
	StartChallengeScRsp                                = 1784
	PVEBattleResultScRsp                               = 161
	UpdateRedDotDataCsReq                              = 5904
	FinishPlotCsReq                                    = 1183
	BuyGoodsScRsp                                      = 1584
	BuyRogueShopBuffCsReq                              = 5638
	ExpeditionDataChangeScNotify                       = 2597
	FinishTutorialCsReq                                = 1697
	GroupStateChangeScRsp                              = 1412
	SyncRogueReviveInfoScNotify                        = 1824
	HeliobusUpgradeLevelCsReq                          = 5855
	HeliobusSelectSkillCsReq                           = 5852
	PlayerKickOutScNotify                              = 97
	FinishFirstTalkByPerformanceNpcScRsp               = 2164
	SyncClientResVersionCsReq                          = 142
	SceneUpdatePositionVersionNotify                   = 1409
	TrainVisitorRewardSendNotify                       = 3760
	GetMarkItemListCsReq                               = 577
	AetherDivideRefreshEndlessScRsp                    = 4894
	EnterTrialActivityStageScRsp                       = 2685
	ChessRogueSkipTeachingLevelScRsp                   = 5539
	EntityBindPropCsReq                                = 1493
	GetMultipleDropInfoCsReq                           = 4683
	GetDailyActiveInfoCsReq                            = 3304
	GetPhoneDataCsReq                                  = 5183
	SpaceZooExchangeItemCsReq                          = 6709
	ChessRogueReRollDiceCsReq                          = 5581
	PlayerHeartBeatScRsp                               = 98
	SecurityReportScRsp                                = 4109
	GetPrivateChatHistoryCsReq                         = 3982
	PlayerReturnTakeRewardCsReq                        = 4542
	GetSingleRedDotParamGroupScRsp                     = 5960
	RankUpAvatarCsReq                                  = 364
	GameplayCounterCountDownScRsp                      = 1402
	GetRogueHandbookDataCsReq                          = 5627
	SceneCastSkillScRsp                                = 1460
	ChessRogueGoAheadScRsp                             = 5525
	ChessRogueChangeyAeonDimensionNotify               = 5550
	ChessRogueNousGetRogueTalentInfoCsReq              = 5596
	GetTutorialCsReq                                   = 1683
	GetArchiveDataScRsp                                = 2361
	SyncAddBlacklistScNotify                           = 2945
	AlleyTakeEventRewardCsReq                          = 4753
	ExchangeRogueRewardKeyCsReq                        = 1851
	ChessRogueGiveUpCsReq                              = 5477
	ChessRogueQuitCsReq                                = 5415
	QueryProductInfoScRsp                              = 25
	StartTimedFarmElementScRsp                         = 1469
	ReplaceLineupCsReq                                 = 790
	GetTrainVisitorRegisterCsReq                       = 3742
	SetBoxingClubResonanceLineupCsReq                  = 4255
	GetAllRedDotDataCsReq                              = 5983
	UseItemCsReq                                       = 542
	ShareScRsp                                         = 4161
	StartTrialActivityCsReq                            = 2698
	GetStrongChallengeActivityDataCsReq                = 6683
	TakeAllApRewardScRsp                               = 3342
	GetTrialActivityDataScRsp                          = 2671
	GetChessRogueBuffEnhanceInfoCsReq                  = 5552
	GetFriendApplyListInfoScRsp                        = 2960
	GetPlayerBoardDataCsReq                            = 2883
	UnlockSkilltreeCsReq                               = 382
	TakeAllRewardCsReq                                 = 3019
	ChessRogueQueryAeonDimensionsCsReq                 = 5530
	PunkLordMonsterInfoScNotify                        = 3240
	NewMailScNotify                                    = 897
	ChessRogueStartScRsp                               = 5589
	GetRogueInfoScRsp                                  = 1861
	RogueModifierSelectCellCsReq                       = 5384
	LogisticsGameScRsp                                 = 4784
	TakePromotionRewardCsReq                           = 392
	ShowNewSupplementVisitorCsReq                      = 3749
	GetAssistListCsReq                                 = 2987
	EnterAdventureCsReq                                = 1383
	GetAetherDivideInfoScRsp                           = 4809
	GetCurAssistScRsp                                  = 2994
	GetNpcStatusScRsp                                  = 2784
	RogueNpcDisappearScRsp                             = 5655
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5680
	RankUpEquipmentScRsp                               = 538
	GetMonsterResearchActivityDataScRsp                = 2645
	SellItemScRsp                                      = 592
	FinishChessRogueNousSubStoryScRsp                  = 5441
	GetPlatformPlayerInfoScRsp                         = 2918
	SpaceZooTakeCsReq                                  = 6764
	AddAvatarScNotify                                  = 355
	SelectChessRogueSubStoryScRsp                      = 5560
	FinishQuestScRsp                                   = 949
	SyncChessRogueNousValueScNotify                    = 5522
	ChangeLineupLeaderScRsp                            = 740
	TakePictureCsReq                                   = 4182
	GetBagScRsp                                        = 561
	GetHeroBasicTypeInfoCsReq                          = 77
	GetChatFriendHistoryCsReq                          = 3942
	GetPunkLordMonsterDataScRsp                        = 3261
	SavePointsInfoNotify                               = 1453
	DeleteBlacklistCsReq                               = 2908
	PrivateMsgOfflineUsersScNotify                     = 3984
	TakeActivityExpeditionRewardCsReq                  = 2564
	UnlockBackGroundMusicCsReq                         = 3182
	GetPlayerReplayInfoScRsp                           = 3584
	ClientObjUploadCsReq                               = 2
	GetChallengeRaidInfoScRsp                          = 2242
	GetRogueScoreRewardInfoScRsp                       = 1802
	GetLoginChatInfoCsReq                              = 3940
	QuitLineupCsReq                                    = 742
	GetFirstTalkNpcCsReq                               = 2182
	GetAvatarDataCsReq                                 = 383
	GetBoxingClubInfoScRsp                             = 4261
	StartAetherDivideSceneBattleCsReq                  = 4882
	TakeChallengeRewardScRsp                           = 1752
	DressRelicAvatarCsReq                              = 352
	FightTreasureDungeonMonsterCsReq                   = 4473
	GetFantasticStoryActivityDataScRsp                 = 4961
	ChooseBoxingClubResonanceScRsp                     = 4209
	GetAetherDivideChallengeInfoCsReq                  = 4833
	GetRogueEndlessActivityDataScRsp                   = 6061
	GetMissionEventDataCsReq                           = 1240
	StopRogueAdventureRoomScRsp                        = 5633
	GetBattleCollegeDataScRsp                          = 5761
	SyncRogueAdventureRoomInfoScNotify                 = 5683
	TakePrestigeRewardScRsp                            = 4709
	LockEquipmentScRsp                                 = 560
	SetLineupNameScRsp                                 = 720
	HeliobusStartRaidScRsp                             = 5890
	HeliobusChallengeUpdateScNotify                    = 5846
	HeliobusActivityDataScRsp                          = 5861
	GetShareDataScRsp                                  = 4184
	SpringRefreshScRsp                                 = 1420
	GetCurLineupDataScRsp                              = 784
	HeliobusSnsCommentCsReq                            = 5897
	PrestigeLevelUpCsReq                               = 4759
	SetStuffToAreaCsReq                                = 4382
	GetCurChallengeCsReq                               = 1738
	SetMissionEventProgressScRsp                       = 1208
	SubMissionRewardScNotify                           = 1233
	StartFinishSubMissionScNotify                      = 1287
	TakeTalkRewardCsReq                                = 2104
	ChessRogueUpdateMoneyInfoScNotify                  = 5480
	TakeBpRewardScRsp                                  = 3082
	GetRaidInfoScRsp                                   = 2209
	CityShopInfoScNotify                               = 1542
	SpaceZooDataScRsp                                  = 6761
	PromoteAvatarScRsp                                 = 319
	ChessRogueGiveUpRollScRsp                          = 5509
	SetLanguageCsReq                                   = 65
	DeactivateFarmElementCsReq                         = 1416
	ReportPlayerCsReq                                  = 2990
	SyncAcceptedPamMissionNotify                       = 4004
	DressAvatarCsReq                                   = 397
	TakeExpeditionRewardCsReq                          = 2542
	TrainVisitorBehaviorFinishCsReq                    = 3783
	TakeMonsterResearchActivityRewardCsReq             = 2659
	SpaceZooDeleteCatScRsp                             = 6738
	OpenTreasureDungeonGridScRsp                       = 4452
	GetRogueAeonInfoScRsp                              = 1801
	QuitTreasureDungeonScRsp                           = 4446
	EnterFightActivityStageScRsp                       = 3682
	RogueModifierDelNotify                             = 5397
	RaidInfoNotify                                     = 2282
	UpgradeAreaScRsp                                   = 4364
	ChessRogueQueryScRsp                               = 5425
	MuseumRandomEventQueryScRsp                        = 4359
	SearchPlayerScRsp                                  = 2965
	TeleportToMissionResetPointScRsp                   = 1265
	GmTalkScRsp                                        = 49
	GetCurSceneInfoCsReq                               = 1442
	CancelExpeditionCsReq                              = 2582
	SyncApplyFriendScNotify                            = 2997
	EnteredSceneChangeScNotify                         = 1432
	GetAllServerPrefsDataCsReq                         = 6183
	SyncRogueVirtualItemInfoScNotify                   = 1857
	EnhanceRogueBuffScRsp                              = 1833
	SpaceZooCatUpdateNotify                            = 6749
	GetRogueTalentInfoScRsp                            = 1839
	HeliobusSnsPostCsReq                               = 5882
	RogueModifierAddNotify                             = 5304
	AcceptActivityExpeditionScRsp                      = 2549
	PunkLordDataChangeNotify                           = 3224
	SetClientRaidTargetCountScRsp                      = 2264
	ChessRogueUpdateDiceInfoScNotify                   = 5511
	StartPunkLordRaidCsReq                             = 3204
	GetChallengeGroupStatisticsScRsp                   = 1745
	ChessRogueEnterCellCsReq                           = 5483
	GetSaveRaidCsReq                                   = 2240
	ChallengeSettleNotify                              = 1742
	RecoverAllLineupScRsp                              = 1494
	DoGachaScRsp                                       = 1984
	RefreshTriggerByClientScNotify                     = 1423
	ChessRogueLeaveScRsp                               = 5429
	ReviveRogueAvatarScRsp                             = 1892
	RemoveStuffFromAreaCsReq                           = 4342
	GetFirstTalkByPerformanceNpcCsReq                  = 2149
	SetAssistCsReq                                     = 2924
	SyncRogueHandbookDataUpdateScNotify                = 5656
	SpaceZooOpCatteryScRsp                             = 6719
	SetPlayerInfoCsReq                                 = 22
	GetAssistHistoryCsReq                              = 2953
	ShowNewSupplementVisitorScRsp                      = 3709
	ChessRogueConfirmRollCsReq                         = 5476
	PickRogueAvatarCsReq                               = 1852
	GetRogueInfoCsReq                                  = 1883
	ChooseBoxingClubStageOptionalBuffCsReq             = 4240
	TakeRogueAeonLevelRewardScRsp                      = 1830
	SetGroupCustomSaveDataCsReq                        = 1498
	GetUpdatedArchiveDataScRsp                         = 2384
	AetherDivideSpiritExpUpScRsp                       = 4846
	MissionEventRewardScNotify                         = 1273
	HeliobusInfoChangedScNotify                        = 5809
	FantasticStoryActivityBattleEndScNotify            = 4960
	GetStuffScNotify                                   = 4397
	TakePromotionRewardScRsp                           = 359
	GetFirstTalkNpcScRsp                               = 2160
	GetRogueTalentInfoCsReq                            = 1807
	SceneEntityUpdateScNotify                          = 1497
	GameplayCounterRecoverCsReq                        = 1450
	PromoteEquipmentCsReq                              = 504
	GetRogueAdventureRoomInfoCsReq                     = 5664
	EnterSectionScRsp                                  = 1465
	TakeAssistRewardCsReq                              = 2993
	GetLineupAvatarDataCsReq                           = 709
	GetChessRogueStoryInfoCsReq                        = 5409
	HeliobusUnlockSkillScNotify                        = 5840
	TakeTrialActivityRewardCsReq                       = 2686
	GetHeroBasicTypeInfoScRsp                          = 94
	TakeRogueEndlessActivityPointRewardCsReq           = 6060
	QuitRogueScRsp                                     = 1877
	SwitchAetherDivideLineUpSlotCsReq                  = 4820
	GateServerScNotify                                 = 13
	UnlockTutorialGuideCsReq                           = 1642
	MuseumTargetRewardNotify                           = 4388
	PlayerLoginFinishCsReq                             = 96
	StartRaidCsReq                                     = 2283
	GetPunkLordMonsterDataCsReq                        = 3283
	EnterSceneScRsp                                    = 1472
	ChessRogueEnterNextLayerCsReq                      = 5412
	BatchMarkChatEmojiCsReq                            = 3955
	SceneGroupRefreshScNotify                          = 1431
	GetBasicInfoCsReq                                  = 76
	TakeCityShopRewardCsReq                            = 1582
	EnterFantasticStoryActivityStageScRsp              = 4982
	FinishCurTurnScRsp                                 = 4309
	SceneEnterStageScRsp                               = 1446
	RogueEndlessActivityBattleEndScNotify              = 6082
	TakeOffAvatarSkinScRsp                             = 308
	AlleyShipmentEventEffectsScNotify                  = 4787
	GetChallengeScRsp                                  = 1761
	MuseumDispatchFinishedScNotify                     = 4346
	GetGachaInfoScRsp                                  = 1961
	GetExpeditionDataScRsp                             = 2561
	TextJoinSaveScRsp                                  = 3861
	SelectChessRogueNousSubStoryScRsp                  = 5449
	ChessRogueEnterCsReq                               = 5501
	ExchangeRogueBuffWithMiracleCsReq                  = 5620
	HeliobusSnsCommentScRsp                            = 5838
	MonthCardRewardNotify                              = 85
	SelectPhoneThemeCsReq                              = 5160
	FinishTutorialGuideScRsp                           = 1609
	SetCurInteractEntityScRsp                          = 1478
	PunkLordMonsterKilledNotify                        = 3265
	TakeOffRelicScRsp                                  = 320
	GetMultipleDropInfoScRsp                           = 4661
	TakeApRewardCsReq                                  = 3383
	GetUnlockTeleportScRsp                             = 1434
	GetMissionDataScRsp                                = 1261
	StartAetherDivideSceneBattleScRsp                  = 4860
	ChessRogueReviveAvatarCsReq                        = 5510
	SceneEntityDisappearScNotify                       = 1438
	GetMainMissionCustomValueScRsp                     = 1294
	TextJoinQueryScRsp                                 = 3884
	DressRelicAvatarScRsp                              = 373
	LeaveRaidCsReq                                     = 2204
	HeliobusStartRaidCsReq                             = 5880
	SpringRecoverSingleAvatarScRsp                     = 1486
	DailyFirstMeetPamCsReq                             = 3404
	TakeMonsterResearchActivityRewardScRsp             = 2680
	FinishRogueDialogueGroupCsReq                      = 1879
	SyncEntityBuffChangeListScNotify                   = 1455
	SetFriendRemarkNameCsReq                           = 2959
	FinishFirstTalkNpcCsReq                            = 2142
	HeliobusSnsLikeScRsp                               = 5819
	EnterRogueMapRoomCsReq                             = 1856
	MatchBoxingClubOpponentCsReq                       = 4204
	EnterTreasureDungeonCsReq                          = 4455
	GetFantasticStoryActivityDataCsReq                 = 4983
	SpaceZooDataCsReq                                  = 6783
	TriggerVoiceScRsp                                  = 4164
	PrepareRogueAdventureRoomCsReq                     = 5661
	AceAntiCheaterCsReq                                = 79
	GetRogueShopBuffInfoCsReq                          = 5660
	SetForbidOtherApplyFriendScRsp                     = 2943
	ExpUpEquipmentScRsp                                = 509
	SetDisplayAvatarScRsp                              = 2860
	ChessRogueRollDiceCsReq                            = 5526
	GetTutorialGuideCsReq                              = 1604
	MultipleDropInfoScNotify                           = 4604
	ChessRogueMoveCellNotify                           = 5470
	DelMailCsReq                                       = 882
	ChessRogueUpdateAeonModifierValueScNotify          = 5506
	GetServerPrefsDataScRsp                            = 6184
	GetChatEmojiListScRsp                              = 3938
	SetAssistAvatarScRsp                               = 2855
	GetRogueDialogueEventDataCsReq                     = 1844
	TakePunkLordPointRewardCsReq                       = 3255
	QuitBattleScRsp                                    = 184
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3797
	SelectChatBubbleCsReq                              = 5104
	MissionRewardScNotify                              = 1282
	SyncRogueSeasonFinishScNotify                      = 1803
	MuseumTakeCollectRewardScRsp                       = 4387
	UnlockedAreaMapScNotify                            = 1415
	CancelCacheNotifyCsReq                             = 4197
	GetMuseumInfoScRsp                                 = 4361
	PlayerLoginCsReq                                   = 83
	ChessRogueUpdateLevelBaseInfoScNotify              = 5563
	GetFriendLoginInfoScRsp                            = 2925
	GetChapterCsReq                                    = 482
	GetSaveLogisticsMapScRsp                           = 4724
	GetTrialActivityDataCsReq                          = 2644
	GetTrainVisitorBehaviorCsReq                       = 3704
	GetSpringRecoverDataCsReq                          = 1476
	SharePunkLordMonsterScRsp                          = 3260
	GetActivityScheduleConfigScRsp                     = 2660
	GetFriendRecommendListInfoCsReq                    = 2920
	GetRogueHandbookDataScRsp                          = 5693
	HandleRogueCommonPendingActionScRsp                = 5689
	ChessRogueNousGetRogueTalentInfoScRsp              = 5573
	GetSceneMapInfoScRsp                               = 1474
	UpdatePlayerSettingScRsp                           = 91
	GetMailCsReq                                       = 883
	FeatureSwitchClosedScNotify                        = 63
	TakeQuestRewardScRsp                               = 984
	TakeKilledPunkLordMonsterScoreCsReq                = 3287
	AetherDivideSpiritInfoScNotify                     = 4808
	SaveLogisticsCsReq                                 = 4733
	GetRndOptionScRsp                                  = 3461
	SpaceZooBornScRsp                                  = 6784
	HandleFriendCsReq                                  = 2938
	GetPlayerDetailInfoScRsp                           = 2984
	RevcMsgScNotify                                    = 3904
	InterruptMissionEventCsReq                         = 1280
	TakeQuestOptionalRewardScRsp                       = 955
	LeaveTrialActivityCsReq                            = 2663
	StaminaInfoScNotify                                = 34
	GetMissionStatusScRsp                              = 1259
	SetAetherDivideLineUpScRsp                         = 4864
	ClearAetherDividePassiveSkillCsReq                 = 4873
	FinishTalkMissionScRsp                             = 1284
	GetPunkLordBattleRecordCsReq                       = 3278
	SetClientPausedScRsp                               = 1414
	AetherDivideTakeChallengeRewardCsReq               = 4803
	ChessRogueQuitScRsp                                = 5551
	FinishSectionIdCsReq                               = 2742
	FinishPerformSectionIdCsReq                        = 2797
	VirtualLineupDestroyNotify                         = 780
	GetSingleRedDotParamGroupCsReq                     = 5982
	MarkItemScRsp                                      = 503
	PlayerReturnSignScRsp                              = 4504
	ActivateFarmElementScRsp                           = 1443
	ReturnLastTownCsReq                                = 1459
	StartAlleyEventScRsp                               = 4719
	FinishTutorialScRsp                                = 1638
	PlayerLogoutScRsp                                  = 84
	GetEnteredSceneCsReq                               = 1401
	SelectRogueDialogueEventCsReq                      = 1896
	StartChallengeCsReq                                = 1704
	GetMarkItemListScRsp                               = 594
	SetSpringRecoverConfigCsReq                        = 1468
	GetAllLineupDataScRsp                              = 759
	CancelMarkItemNotify                               = 527
	SetGameplayBirthdayCsReq                           = 44
	SyncRogueAreaUnlockScNotify                        = 1862
	DestroyItemCsReq                                   = 524
	GetChessRogueNousStoryInfoScRsp                    = 5568
	GetPunkLordDataScRsp                               = 3273
	TakeFightActivityRewardScRsp                       = 3642
	SummonPunkLordMonsterCsReq                         = 3242
	StartAetherDivideStageBattleScRsp                  = 4880
	GetRecyleTimeScRsp                                 = 565
	GetNpcTakenRewardCsReq                             = 2183
	ExchangeGachaCeilingScRsp                          = 1919
	GetGachaCeilingScRsp                               = 1960
	FinishChessRogueSubStoryCsReq                      = 5401
	TrialBackGroundMusicCsReq                          = 3142
	GetRecyleTimeCsReq                                 = 588
	AddEquipmentScNotify                               = 533
	ChessRogueSelectBpScRsp                            = 5571
	GetNpcMessageGroupScRsp                            = 2761
	GetKilledPunkLordMonsterDataCsReq                  = 3246
	OpenRogueChestScRsp                                = 1886
	GetVideoVersionKeyCsReq                            = 72
	GetPlayerReturnMultiDropInfoCsReq                  = 4684
	GetAlleyInfoScRsp                                  = 4761
	AddBlacklistScRsp                                  = 2973
	TextJoinBatchSaveScRsp                             = 3860
	ClearAetherDividePassiveSkillScRsp                 = 4845
	TakeLoginActivityRewardCsReq                       = 2604
	GetStrongChallengeActivityDataScRsp                = 6661
	DestroyItemScRsp                                   = 578
	GetJukeboxDataCsReq                                = 3183
	SyncRogueAeonLevelUpRewardScNotify                 = 1891
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6097
	SyncTaskCsReq                                      = 1260
	SyncChessRogueNousMainStoryScNotify                = 5502
	UpdateRogueAdventureRoomScoreCsReq                 = 5643
	FinishItemIdScRsp                                  = 2760
	MuseumRandomEventSelectScRsp                       = 4390
	ExpUpRelicCsReq                                    = 540
	GetSceneMapInfoCsReq                               = 1495
	SpaceZooMutateScRsp                                = 6760
	ComposeItemScRsp                                   = 564
	EnterChessRogueAeonRoomScRsp                       = 5519
	PlayerLoginScRsp                                   = 61
	GameplayCounterCountDownCsReq                      = 1410
	ChessRogueLayerAccountInfoNotify                   = 5468
	GetStageLineupCsReq                                = 783
	AlleyFundsScNotify                                 = 4790
	EnterStrongChallengeActivityStageCsReq             = 6604
	GetChapterScRsp                                    = 460
	MuseumFundsChangedScNotify                         = 4345
	EnhanceCommonRogueBuffScRsp                        = 5646
	ReEnterLastElementStageCsReq                       = 1411
	TakeRogueEventHandbookRewardScRsp                  = 5616
	ChessRogueUpdateActionPointScNotify                = 5431
	UpgradeAreaStatScRsp                               = 4352
	UnlockTutorialCsReq                                = 1682
	GetLevelRewardScRsp                                = 8
	AvatarExpUpScRsp                                   = 384
	MuseumTargetStartNotify                            = 4308
	AlleyEventChangeNotify                             = 4797
	TakeQuestOptionalRewardCsReq                       = 909
	GetWaypointScRsp                                   = 461
	GetMailScRsp                                       = 861
	GetReplayTokenScRsp                                = 3561
	StartBoxingClubBattleCsReq                         = 4282
	UpdateServerPrefsDataScRsp                         = 6160
	TakeChallengeRaidRewardCsReq                       = 2219
	SceneEntityMoveCsReq                               = 1483
	GetChatFriendHistoryScRsp                          = 3919
	TakePunkLordPointRewardScRsp                       = 3264
	GetMuseumInfoCsReq                                 = 4383
	SetTurnFoodSwitchCsReq                             = 556
	BattleCollegeDataChangeScNotify                    = 5704
	SetAssistAvatarCsReq                               = 2809
	GetMonsterResearchActivityDataCsReq                = 2673
	SyncRogueDialogueEventDataScNotify                 = 1872
	SetGenderCsReq                                     = 93
	EnterAdventureScRsp                                = 1361
	LogisticsGameCsReq                                 = 4704
	GetChatEmojiListCsReq                              = 3997
	DressAvatarSkinScRsp                               = 390
	GetAssistListScRsp                                 = 2906
	AetherDivideTakeChallengeRewardScRsp               = 4827
	SetClientRaidTargetCountCsReq                      = 2255
	GetMissionStatusCsReq                              = 1292
	RogueNpcDisappearCsReq                             = 5609
	GetNpcStatusCsReq                                  = 2704
	GetTrainVisitorRegisterScRsp                       = 3719
	ChessRoguePickAvatarCsReq                          = 5434
	ChessRogueNousEnableRogueTalentCsReq               = 5433
	GetBattleCollegeDataCsReq                          = 5783
	InteractTreasureDungeonGridCsReq                   = 4420
	SetSignatureScRsp                                  = 2849
	SyncServerSceneChangeNotify                        = 1430
	EquipAetherDividePassiveSkillScRsp                 = 4852
	SelectChessRogueSubStoryCsReq                      = 5407
	GetChallengeRaidInfoCsReq                          = 2260
	UnlockChatBubbleScNotify                           = 5182
	HealPoolInfoNotify                                 = 1489
	SetForbidOtherApplyFriendCsReq                     = 2970
	ReEnterLastElementStageScRsp                       = 1435
	InterruptMissionEventScRsp                         = 1290
	GetPlatformPlayerInfoCsReq                         = 2914
	GetAetherDivideInfoCsReq                           = 4849
	RefreshTriggerByClientScRsp                        = 1439
	ClientObjUploadScRsp                               = 54
	AvatarExpUpCsReq                                   = 304
	GetArchiveDataCsReq                                = 2383
	GetPlayerDetailInfoCsReq                           = 2904
	GetShopListCsReq                                   = 1583
	CancelExpeditionScRsp                              = 2560
	LockRelicScRsp                                     = 545
	FinishRogueDialogueGroupScRsp                      = 1889
	GetPrivateChatHistoryScRsp                         = 3960
	StartPunkLordRaidScRsp                             = 3284
	SceneCastSkillCsReq                                = 1482
	ChessRogueStartCsReq                               = 5554
	EnterRogueMapRoomScRsp                             = 1822
	SaveLogisticsScRsp                                 = 4788
	SyncRogueCommonActionResultScNotify                = 5625
	AcceptedPamMissionExpireCsReq                      = 4083
	GetNpcMessageGroupCsReq                            = 2783
	HeliobusSnsReadScRsp                               = 5884
	SpringRecoverSingleAvatarCsReq                     = 1485
	MuseumTargetMissionFinishNotify                    = 4333
	QuitTreasureDungeonCsReq                           = 4490
	GetBagCsReq                                        = 583
	TakePictureScRsp                                   = 4160
	SetRedPointStatusScNotify                          = 62
	FinishChessRogueNousSubStoryCsReq                  = 5504
	GetLoginActivityScRsp                              = 2661
	DressAvatarScRsp                                   = 338
	ChessRogueQueryCsReq                               = 5564
	TeleportToMissionResetPointCsReq                   = 1288
	UpgradeAreaCsReq                                   = 4355
	ChessRogueNousEditDiceScRsp                        = 5440
	DailyActiveInfoNotify                              = 3382
	ExchangeGachaCeilingCsReq                          = 1942
	SceneCastSkillMpUpdateScNotify                     = 1452
	AetherDivideSkillItemScNotify                      = 4806
	RaidKickByServerScNotify                           = 2292
	DelMailScRsp                                       = 860
	ChessRogueSkipTeachingLevelCsReq                   = 5489
	GetGachaCeilingCsReq                               = 1982
	SyncClientResVersionScRsp                          = 119
	FinishFirstTalkByPerformanceNpcCsReq               = 2155
	StartCocoonStageCsReq                              = 1403
	DeleteFriendCsReq                                  = 2955
	PlayBackGroundMusicScRsp                           = 3184
	ChessRogueEnterCellScRsp                           = 5547
	MuseumInfoChangedScNotify                          = 4373
	StartBattleCollegeCsReq                            = 5784
	GetCurChallengeScRsp                               = 1749
	GetTreasureDungeonActivityDataCsReq                = 4449
	GetMainMissionCustomValueCsReq                     = 1277
	FinishCosumeItemMissionScRsp                       = 1264
	PlayerReturnInfoQueryCsReq                         = 4597
	GetPhoneDataScRsp                                  = 5161
	GetLevelRewardCsReq                                = 46
	GetChessRogueStoryAeonTalkInfoScRsp                = 5578
	UpgradeAreaStatCsReq                               = 4340
	EnhanceRogueBuffCsReq                              = 1808
	TakeApRewardScRsp                                  = 3361
	ChessRogueQueryBpScRsp                             = 5416
	GetNpcTakenRewardScRsp                             = 2161
	GetShareDataCsReq                                  = 4104
	TakeMailAttachmentCsReq                            = 842
	HandleFriendScRsp                                  = 2949
	AetherDivideRefreshEndlessScNotify                 = 4853
	SearchPlayerCsReq                                  = 2988
	HeliobusSelectSkillScRsp                           = 5873
	SelectRogueDialogueEventScRsp                      = 1817
	EnterSectionCsReq                                  = 1488
	EnterChessRogueAeonRoomCsReq                       = 5419
	QuitLineupScRsp                                    = 719
	DeleteBlacklistScRsp                               = 2933
	EnterFightActivityStageCsReq                       = 3684
	TakeKilledPunkLordMonsterScoreScRsp                = 3206
	ComposeLimitNumCompleteNotify                      = 587
	SpaceZooBornCsReq                                  = 6704
	GetJukeboxDataScRsp                                = 3161
	SetNicknameScRsp                                   = 59
	TakeRogueMiracleHandbookRewardCsReq                = 5622
	GameplayCounterRecoverScRsp                        = 1441
	GetSpringRecoverDataScRsp                          = 1467
	TakeExpeditionRewardScRsp                          = 2519
	GetVideoVersionKeyScRsp                            = 48
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6019
	SyncTurnFoodNotify                                 = 593
	FinishTalkMissionCsReq                             = 1204
	RefreshAlleyOrderCsReq                             = 4773
	PlayerReturnTakePointRewardScRsp                   = 4560
	FinishSectionIdScRsp                               = 2719
	SpaceZooDeleteCatCsReq                             = 6797
	GetFarmStageGachaInfoScRsp                         = 1384
	JoinLineupScRsp                                    = 760
	GetCurBattleInfoCsReq                              = 182
	ChooseBoxingClubResonanceCsReq                     = 4249
	SetHeroBasicTypeScRsp                              = 78
	GetKilledPunkLordMonsterDataScRsp                  = 3208
	ChallengeRaidNotify                                = 2238
	AddBlacklistCsReq                                  = 2952
	GetRogueAeonInfoCsReq                              = 1866
	AcceptExpeditionScRsp                              = 2584
	FightActivityDataChangeScNotify                    = 3604
	TrialBackGroundMusicScRsp                          = 3119
	OpenRogueChestCsReq                                = 1885
	QuitRogueCsReq                                     = 1878
	ComposeSelectedRelicScRsp                          = 508
	GetLineupAvatarDataScRsp                           = 755
	ChessRogueSelectCellCsReq                          = 5497
	AcceptMissionEventCsReq                            = 1245
	PunkLordRaidTimeOutScNotify                        = 3220
	EnableRogueTalentScRsp                             = 1847
	TakeAssistRewardScRsp                              = 2956
	GetChessRogueNousStoryInfoCsReq                    = 5404
	SpaceZooOpCatteryCsReq                             = 6742
	AlleyGuaranteedFundsScRsp                          = 4794
	MarkReadMailCsReq                                  = 804
	GetRogueShopMiracleInfoScRsp                       = 5682
	GetRogueBuffEnhanceInfoCsReq                       = 1890
	SelectInclinationTextCsReq                         = 2197
	PunkLordBattleResultScNotify                       = 3290
	BuyBpLevelCsReq                                    = 3060
	SetAssistScRsp                                     = 2978
	PlayerReturnPointChangeScNotify                    = 4584
	SetPlayerInfoScRsp                                 = 14
	SetIsDisplayAvatarInfoCsReq                        = 2842
	EnterAetherDivideSceneScRsp                        = 4861
	CommonRogueQueryCsReq                              = 5685
	SwapLineupCsReq                                    = 797
	ChessRogueLeaveCsReq                               = 5421
	PlayerReturnForceFinishScNotify                    = 4549
	DoGachaCsReq                                       = 1904
	SpringRecoverScRsp                                 = 1479
	ChessRogueConfirmRollScRsp                         = 5545
	ReplaceLineupScRsp                                 = 746
	SubmitMonsterResearchActivityMaterialScRsp         = 2692
	BatchMarkChatEmojiScRsp                            = 3964
	LogisticsDetonateStarSkiffCsReq                    = 4727
	DeleteSummonUnitScRsp                              = 1436
	GetSaveRaidScRsp                                   = 2252
	GetFightActivityDataCsReq                          = 3683
	MuseumRandomEventQueryCsReq                        = 4392
	SyncChessRogueNousSubStoryScNotify                 = 5496
	PlayerLoginFinishScRsp                             = 17
	EnterRogueCsReq                                    = 1882
	AetherDivideLineupScNotify                         = 4878
	SyncChessRogueMainStoryFinishScNotify              = 5579
	AcceptActivityExpeditionCsReq                      = 2538
	GetAuthkeyCsReq                                    = 52
	ExpUpEquipmentCsReq                                = 549
	ChessRogueCheatRollCsReq                           = 5427
	GetRogueInitialScoreScRsp                          = 1818
	GetSecretKeyInfoScRsp                              = 26
	MarkChatEmojiCsReq                                 = 3949
	ShareCsReq                                         = 4183
	SyncRogueGetItemScNotify                           = 1895
	OpenTreasureDungeonGridCsReq                       = 4440
	MissionAcceptScNotify                              = 1253
	ChessRogueGoAheadCsReq                             = 5559
	SceneEntityTeleportScRsp                           = 1496
	GetChessRogueStoryInfoScRsp                        = 5452
	GmTalkCsReq                                        = 38
	HeroBasicTypeChangedNotify                         = 18
	UnlockSkilltreeScRsp                               = 360
	PlayerSyncScNotify                                 = 683
	SetBoxingClubResonanceLineupScRsp                  = 4264
	GetAllRedDotDataScRsp                              = 5961
	InteractPropCsReq                                  = 1404
	SetHeadIconCsReq                                   = 2804
	ChessRogueGiveUpRollCsReq                          = 5512
	EnhanceChessRogueBuffScRsp                         = 5518
	ChessRogueGiveUpScRsp                              = 5517
	SendMsgScRsp                                       = 3961
	StartTrialActivityScRsp                            = 2613
	PromoteAvatarCsReq                                 = 342
	ReportPlayerScRsp                                  = 2946
	TakeTalkRewardScRsp                                = 2184
	EnterTrialActivityStageCsReq                       = 2689
	GetAllServerPrefsDataScRsp                         = 6161
	SyncRogueStatusScNotify                            = 1869
	ApplyFriendScRsp                                   = 2919
	StartTimedCocoonStageScRsp                         = 1437
	GetDailyActiveInfoScRsp                            = 3384
	StartRogueCsReq                                    = 1804
	TakeAllRewardScRsp                                 = 3097
	GetAllSaveRaidCsReq                                = 2273
	EnterRogueEndlessActivityStageCsReq                = 6004
	GetChallengeCsReq                                  = 1783
	SyncDeleteFriendScNotify                           = 2940
	GetCurLineupDataCsReq                              = 704
	ExchangeStaminaCsReq                               = 64
	BattleLogReportCsReq                               = 138
	BuyGoodsCsReq                                      = 1504
	FinishPlotScRsp                                    = 1161
	BuyRogueShopBuffScRsp                              = 5649
	PrestigeLevelUpScRsp                               = 4780
	ChessRogueReviveAvatarScRsp                        = 5521
	StopRogueAdventureRoomCsReq                        = 5608
	GetRaidInfoCsReq                                   = 2249
	SyncRogueAeonScNotify                              = 1848
	SetStuffToAreaScRsp                                = 4360
	TakeOffAvatarSkinCsReq                             = 346
	GetRogueScoreRewardInfoCsReq                       = 1810
	AlleyOrderChangedScNotify                          = 4720
	SceneEnterStageCsReq                               = 1490
	TakePrestigeRewardCsReq                            = 4749
	PrepareRogueAdventureRoomScRsp                     = 5604
	TakeBpRewardCsReq                                  = 3084
	AlleyPlacingGameScRsp                              = 4764
	TakeChallengeRewardCsReq                           = 1740
	StartTimedFarmElementCsReq                         = 1457
	GetExhibitScNotify                                 = 4338
	UpdateRedDotDataScRsp                              = 5984
	GetAetherDivideChallengeInfoScRsp                  = 4888
	GetRogueShopBuffInfoScRsp                          = 5642
	GetLoginChatInfoScRsp                              = 3952
	GroupStateChangeCsReq                              = 1428
	GetPlayerReplayInfoCsReq                           = 3504
	UnlockBackGroundMusicScRsp                         = 3160
	ExchangeRogueRewardKeyScRsp                        = 1898
	GetRogueAdventureRoomInfoScRsp                     = 5640
	TakeActivityExpeditionRewardScRsp                  = 2540
	GetBoxingClubInfoCsReq                             = 4283
	GetUpdatedArchiveDataCsReq                         = 2304
	ReserveStaminaExchangeCsReq                        = 30
	AcceptMainMissionScRsp                             = 1278
	FinishAeonDialogueGroupScRsp                       = 1832
	SceneEntityDieScNotify                             = 1418
	TakeOffEquipmentScRsp                              = 309
	CancelActivityExpeditionScRsp                      = 2555
	TakeChapterRewardScRsp                             = 497
	SetLineupNameCsReq                                 = 745
	LeaveRaidScRsp                                     = 2284
	SpringRefreshCsReq                                 = 1445
	SetCurWaypointScRsp                                = 484
	EnterFantasticStoryActivityStageCsReq              = 4984
	GetMissionEventDataScRsp                           = 1252
	DailyRefreshNotify                                 = 68
	GmTalkScNotify                                     = 19
	UseTreasureDungeonItemCsReq                        = 4459
	SelectChessRogueNousSubStoryCsReq                  = 5516
	LeaveRogueCsReq                                    = 1842
	PlayerHeartBeatCsReq                               = 51
	GetRogueEndlessActivityDataCsReq                   = 6083
	AetherDivideRefreshEndlessCsReq                    = 4877
	SetCurInteractEntityCsReq                          = 1424
	SpaceZooTakeScRsp                                  = 6740
	ChessRogueReRollDiceScRsp                          = 5533
	SecurityReportCsReq                                = 4149
	BuyRogueShopMiracleScRsp                           = 5697
	GetAvatarDataScRsp                                 = 361
	RankUpEquipmentCsReq                               = 597
	AntiAddictScNotify                                 = 20
	StartAetherDivideChallengeBattleCsReq              = 4842
	TextJoinSaveCsReq                                  = 3883
)

const (
	GmGive       = 11127
	GmWorldLevel = 11001
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(GetHeroBasicTypeInfoScRsp, func() any { return new(proto.GetHeroBasicTypeInfoScRsp) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartbeatCsReq) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartbeatScRsp) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(SelectRogueDialogueEventCsReq, func() any { return new(proto.SelectRogueDialogueEventCsReq) })
	c.regMsg(SelectRogueDialogueEventScRsp, func() any { return new(proto.SelectRogueDialogueEventScRsp) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(SetHeroBasicTypeCsReq, func() any { return new(proto.SetHeroBasicTypeCsReq) })
	c.regMsg(SetHeroBasicTypeScRsp, func() any { return new(proto.SetHeroBasicTypeScRsp) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(GmGive, func() any { return new(spb.GmGive) })
	c.regMsg(GmWorldLevel, func() any { return new(spb.GmWorldLevel) })
}
