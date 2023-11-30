package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

const (
	LeaveChallengeCsReq                         = 1739
	PlayerReturnSignScRsp                       = 4551
	MultipleDropInfoNotify                      = 4642
	SyncRogueMiracleScNotify                    = 1820
	ChessRogueFinishCurRoomNotify               = 5422
	ReplaceRogueMiracleDisplayScNotify          = 1891
	GetRogueAeonInfoCsReq                       = 1853
	RollChessRogueBuffScRsp                     = 5511
	ChessRogueReviveAvatarCsReq                 = 5417
	DailyFirstMeetPamCsReq                      = 3451
	FinishQuestScRsp                            = 909
	RogueChallengeActivityBuffChooseScNotify    = 2641
	GetWaypointScRsp                            = 468
	EnhanceChessRogueBuffCsReq                  = 5414
	CancelCacheNotifyScRsp                      = 4172
	EnterSceneCsReq                             = 1458
	TakeQuestRewardCsReq                        = 951
	GetTrialActivityDataCsReq                   = 2664
	MissionRewardScNotify                       = 1239
	RepairChessRogueMiracleCsReq                = 5484
	ChessRogueEnterNextLayerCsReq               = 5545
	StartTrialActivityCsReq                     = 2687
	FinishTalkMissionScRsp                      = 1244
	SwapRogueBuffCsReq                          = 1900
	GetLoginActivityScRsp                       = 2668
	UpdateMechanismBarScNotify                  = 1426
	TakeBpRewardScRsp                           = 3039
	TextJoinSaveScRsp                           = 3868
	SelectInclinationTextCsReq                  = 2167
	GameplayCounterCountDownCsReq               = 1495
	StartCocoonStageScRsp                       = 1441
	EnterSceneByServerScNotify                  = 1448
	GetRogueInfoCsReq                           = 1852
	UnlockTutorialGuideScRsp                    = 1676
	UnlockTutorialGuideCsReq                    = 1623
	GetVideoVersionKeyCsReq                     = 65
	GmTalkScRsp                                 = 9
	GetAssistListCsReq                          = 3000
	GetChapterScRsp                             = 442
	GetNpcMessageGroupScRsp                     = 2768
	ReplaceLineupScRsp                          = 728
	GetStuffScNotify                            = 4367
	BattleCollegeDataChangeScNotify             = 5751
	SyncAcceptedPamMissionNotify                = 4051
	UpdateServerPrefsDataCsReq                  = 6139
	TakeAllApRewardScRsp                        = 3323
	GetAetherDivideChallengeInfoScRsp           = 4802
	GetArchiveDataScRsp                         = 2368
	HeliobusEnterBattleCsReq                    = 5829
	SyncRogueMiracleInfoScNotify                = 1850
	GetMissionDataScRsp                         = 1268
	GetAllServerPrefsDataCsReq                  = 6152
	GetSingleRedDotParamGroupScRsp              = 5942
	UpgradeAreaScRsp                            = 4386
	DestroyItemCsReq                            = 522
	LockEquipmentCsReq                          = 539
	ExchangeStaminaScRsp                        = 8
	GetRogueUnlockDataCsReq                     = 5605
	SetAssistAvatarCsReq                        = 2870
	BuyGoodsScRsp                               = 1544
	SyncChessRogueBuffDropInfoScNotify          = 5470
	EnhanceCommonRogueBuffCsReq                 = 5696
	EnterSceneScRsp                             = 1465
	RemoveStuffFromAreaCsReq                    = 4323
	RogueChallengeRefreshAssistListCsReq        = 2616
	ChessRogueQueryBpCsReq                      = 5457
	SellItemCsReq                               = 504
	GameplayCounterRecoverScRsp                 = 1463
	SceneEntityMoveScRsp                        = 1468
	PlayerReturnInfoQueryCsReq                  = 4567
	TriggerHealVoiceScRsp                       = 4176
	GetRogueShopBuffInfoCsReq                   = 5642
	GetRogueDialogueEventDataScRsp              = 1859
	GetDailyActiveInfoCsReq                     = 3351
	PlayerLoginScRsp                            = 68
	TakePromotionRewardCsReq                    = 329
	ReEnterLastElementStageScRsp                = 1494
	FinishAeonDialogueGroupCsReq                = 1893
	PlayerGetTokenCsReq                         = 39
	SaveLogisticsCsReq                          = 4730
	PunkLordMonsterKilledNotify                 = 3279
	InteractPropScRsp                           = 1444
	ComposeLimitNumUpdateNotify                 = 537
	GetPlayerReturnMultiDropInfoScRsp           = 4639
	SceneEnterStageCsReq                        = 1496
	GetChessRogueBuffEnhanceInfoScRsp           = 5460
	SetSignatureCsReq                           = 2872
	SetHeadIconScRsp                            = 2844
	GetPrivateChatHistoryScRsp                  = 3942
	FinishPerformSectionIdScRsp                 = 2772
	LeaveChallengeScRsp                         = 1742
	AlleyShipUnlockScNotify                     = 4732
	StartChallengeCsReq                         = 1751
	HeliobusActivityDataScRsp                   = 5868
	TrialBackGroundMusicCsReq                   = 3123
	PromoteAvatarScRsp                          = 376
	LeaveTrialActivityScRsp                     = 2685
	EnterChessRogueAeonRoomScRsp                = 5463
	SelectRogueMiracleCsReq                     = 1841
	TakeQuestOptionalRewardCsReq                = 970
	GetTreasureDungeonActivityDataCsReq         = 4409
	TakeLoginActivityRewardCsReq                = 2651
	GroupStateChangeCsReq                       = 1406
	DelMailCsReq                                = 839
	SetHeadIconCsReq                            = 2851
	GetMultipleDropInfoCsReq                    = 4652
	QuitBattleScRsp                             = 144
	GetMailScRsp                                = 868
	SwitchLineupIndexScRsp                      = 792
	SetClientPausedScRsp                        = 1424
	LeaveRaidCsReq                              = 2251
	ReserveStaminaExchangeCsReq                 = 1
	DeleteFriendScRsp                           = 2986
	GetJukeboxDataScRsp                         = 3168
	UpgradeAreaStatCsReq                        = 4308
	UnlockTutorialCsReq                         = 1639
	FinishFirstTalkNpcCsReq                     = 2123
	FightTreasureDungeonMonsterScRsp            = 4403
	GetSpringRecoverDataScRsp                   = 1499
	GetTrainVisitorBehaviorScRsp                = 3744
	EnterTreasureDungeonCsReq                   = 4482
	BuyRogueShopMiracleCsReq                    = 5676
	AetherDivideSpiritExpUpScRsp                = 4828
	TreasureDungeonDataScNotify                 = 4452
	GetPunkLordMonsterDataCsReq                 = 3252
	RogueChallengeActivityDataScRsp             = 2628
	GetFightActivityDataScRsp                   = 3668
	GetGachaCeilingScRsp                        = 1942
	PlayerReturnInfoQueryScRsp                  = 4572
	SyncServerSceneChangeNotify                 = 1401
	TakePromotionRewardScRsp                    = 374
	UpdateFloorSavedValueNotify                 = 1418
	UseItemScRsp                                = 576
	SetSpringRecoverConfigScRsp                 = 1464
	GetAllLineupDataCsReq                       = 729
	NewAssistHistoryNotify                      = 2941
	FinishTutorialGuideCsReq                    = 1609
	TrainRefreshTimeNotify                      = 3739
	MuseumTargetMissionFinishNotify             = 4330
	GetLineupAvatarDataCsReq                    = 770
	GetFantasticStoryActivityDataCsReq          = 4952
	SetClientPausedCsReq                        = 1416
	ChessRogueMoveCellNotify                    = 5559
	UpdateFeatureSwitchScNotify                 = 62
	SyncRogueAeonLevelUpRewardScNotify          = 1818
	TakeRogueScoreRewardScRsp                   = 1819
	HealPoolInfoNotify                          = 1484
	BuyBpLevelCsReq                             = 3042
	GetChallengeScRsp                           = 1768
	StartRogueCsReq                             = 1851
	ExchangeHcoinScRsp                          = 596
	DeleteSummonUnitCsReq                       = 1478
	MuseumRandomEventSelectScRsp                = 4396
	TakeChapterRewardScRsp                      = 467
	UpdatePlayerSettingCsReq                    = 27
	GetStageLineupScRsp                         = 768
	GetMainMissionCustomValueScRsp              = 1205
	RogueModifierSelectCellScRsp                = 5339
	ReportPlayerCsReq                           = 2996
	BatchMarkChatEmojiCsReq                     = 3982
	TakeTrialActivityRewardCsReq                = 2675
	LogisticsInfoScNotify                       = 4779
	SyncAddBlacklistScNotify                    = 2903
	SceneEntityMoveCsReq                        = 1452
	AddBlacklistCsReq                           = 2947
	SceneEntityDieScNotify                      = 1489
	TakePunkLordPointRewardCsReq                = 3282
	RollChessRogueBuffCsReq                     = 5565
	AddAvatarScNotify                           = 382
	SelectChatBubbleCsReq                       = 5151
	HeliobusSnsCommentScRsp                     = 5872
	EnableRogueTalentCsReq                      = 1888
	TakeMonsterResearchActivityRewardScRsp      = 2619
	TakeRogueEventHandbookRewardCsReq           = 1849
	AetherDivideTainerInfoScNotify              = 4900
	TakePunkLordPointRewardScRsp                = 3286
	GetLineupAvatarDataScRsp                    = 782
	TakeApRewardScRsp                           = 3368
	ReforgeChessRogueBuffScRsp                  = 5415
	SavePointsInfoNotify                        = 1438
	StartBoxingClubBattleScRsp                  = 4242
	GetMonsterResearchActivityDataCsReq         = 2692
	SpringRecoverSingleAvatarScRsp              = 1475
	BuyGoodsCsReq                               = 1551
	MonthCardRewardNotify                       = 56
	ShowNewSupplementVisitorScRsp               = 3770
	GetFirstTalkNpcCsReq                        = 2139
	FinishPlotScRsp                             = 1168
	AddRogueMiracleScNotify                     = 1836
	RecoverAllLineupScRsp                       = 1405
	SyncRogueCommonItemDisplayScNotify          = 5603
	SetGenderScRsp                              = 73
	GetFriendLoginInfoScRsp                     = 2955
	RogueNpcDisappearScRsp                      = 5682
	EnterActivityBattleStageCsReq               = 2669
	PunkLordMonsterInfoScNotify                 = 3208
	StartRogueScRsp                             = 1844
	TakePictureScRsp                            = 4142
	SaveLogisticsScRsp                          = 4702
	GetChessRogueStoryInfoCsReq                 = 5580
	TakeRogueMiracleHandbookRewardCsReq         = 1831
	GetSaveRaidScRsp                            = 2247
	SetSpringRecoverConfigCsReq                 = 1435
	GetUnlockTeleportCsReq                      = 1450
	MuseumTargetRewardNotify                    = 4302
	NewMailScNotify                             = 867
	SceneEnterStageScRsp                        = 1428
	CityShopInfoScNotify                        = 1523
	ChessRogueSelectCellCsReq                   = 5466
	ReforgeRogueBuffCsReq                       = 1899
	ChessRogueGiveUpCsReq                       = 5462
	GetRaidInfoCsReq                            = 2209
	GetMissionStatusCsReq                       = 1229
	QuitLineupCsReq                             = 723
	BuyRogueShopBuffCsReq                       = 5672
	MatchBoxingClubOpponentCsReq                = 4251
	TeleportToMissionResetPointScRsp            = 1279
	LogisticsDetonateStarSkiffScRsp             = 4769
	AcceptedPamMissionExpireScRsp               = 4068
	FinishCurTurnCsReq                          = 4309
	SetClientRaidTargetCountScRsp               = 2286
	SyncRogueSeasonFinishScNotify               = 1866
	ChessRoguePickAvatarScRsp                   = 5532
	SetTurnFoodSwitchCsReq                      = 573
	StopRogueAdventureRoomCsReq                 = 5632
	StartAetherDivideSceneBattleCsReq           = 4839
	TakeRogueScoreRewardCsReq                   = 1874
	JoinLineupScRsp                             = 742
	SelectRogueBuffCsReq                        = 1872
	ClientDownloadDataScNotify                  = 14
	AetherDivideSpiritInfoScNotify              = 4832
	QuitTreasureDungeonScRsp                    = 4428
	SceneCastSkillMpUpdateScNotify              = 1447
	SyncRogueReviveInfoScNotify                 = 1822
	SendMsgScRsp                                = 3968
	PlayerKickOutScNotify                       = 67
	AetherDivideTakeChallengeRewardCsReq        = 4866
	GetNpcTakenRewardScRsp                      = 2168
	SyncRogueGetItemScNotify                    = 1880
	SwitchLineupIndexCsReq                      = 747
	StopRogueAdventureRoomScRsp                 = 5630
	GiveUpBoxingClubChallengeCsReq              = 4223
	SelectRogueBonusCsReq                       = 1814
	RechargeSuccNotify                          = 574
	SyncRogueCommonBuffDisplayScNotify          = 5647
	StaminaInfoScNotify                         = 12
	UnlockSkilltreeCsReq                        = 339
	GetAllRedDotDataScRsp                       = 5968
	ComposeItemCsReq                            = 582
	TakeTrialActivityRewardScRsp                = 2626
	GetRogueDialogueEventDataCsReq              = 1864
	GetAllServerPrefsDataScRsp                  = 6168
	AddRogueBuffScNotify                        = 1803
	TrialBackGroundMusicScRsp                   = 3176
	ChooseBoxingClubStageOptionalBuffScRsp      = 4247
	PromoteEquipmentScRsp                       = 544
	GetRecyleTimeCsReq                          = 502
	ClientObjDownloadDataScNotify               = 95
	TakeRogueAeonLevelRewardCsReq               = 1898
	SetCurWaypointCsReq                         = 451
	SelectChessRogueBonusCsReq                  = 5527
	AddBlacklistScRsp                           = 2992
	ChessRogueUpdateBoardScNotify               = 5533
	GetPunkLordMonsterDataScRsp                 = 3268
	OpenTreasureDungeonGridCsReq                = 4408
	TakeMailAttachmentScRsp                     = 876
	GetGachaInfoCsReq                           = 1952
	GetRogueShopBuffInfoScRsp                   = 5623
	GetAllSaveRaidScRsp                         = 2203
	SelectPhoneThemeScRsp                       = 5123
	ExchangeHcoinCsReq                          = 519
	AcceptedPamMissionExpireCsReq               = 4052
	GetSaveLogisticsMapScRsp                    = 4722
	SetLineupNameScRsp                          = 704
	GetUnlockTeleportScRsp                      = 1412
	DressAvatarCsReq                            = 367
	GetMissionDataCsReq                         = 1252
	GetLevelRewardTakenListCsReq                = 19
	ChessRogueUpdateActionPointScNotify         = 5571
	TakeOffAvatarSkinCsReq                      = 328
	ReEnterLastElementStageCsReq                = 1436
	ChessRogueEnterCsReq                        = 5478
	TakeRogueEndlessActivityPointRewardCsReq    = 6042
	BuyNpcStuffCsReq                            = 4351
	FinishCosumeItemMissionScRsp                = 1286
	MuseumTargetStartNotify                     = 4332
	SyncClientResVersionCsReq                   = 123
	QuitBattleScNotify                          = 167
	TrialActivityDataChangeScNotify             = 2613
	AcceptMainMissionScRsp                      = 1246
	AceAntiCheaterScRsp                         = 84
	ReportPlayerScRsp                           = 2928
	SubMissionRewardScNotify                    = 1230
	MuseumRandomEventQueryScRsp                 = 4374
	SpringRecoverScRsp                          = 1413
	GetMuseumInfoCsReq                          = 4352
	ChessRoguePickAvatarCsReq                   = 5408
	GameplayCounterRecoverCsReq                 = 1449
	GetActivityScheduleConfigCsReq              = 2639
	SelectInclinationTextScRsp                  = 2172
	SyncRogueHandbookDataUpdateScNotify         = 5673
	GetActivityScheduleConfigScRsp              = 2642
	EnterChessRogueAeonRoomCsReq                = 5445
	ExchangeGachaCeilingScRsp                   = 1976
	GetFriendListInfoCsReq                      = 2952
	BoxingClubRewardScNotify                    = 4267
	TakeRogueEventHandbookRewardScRsp           = 1863
	SyncRogueMiracleSelectInfoScNotify          = 1886
	SetSignatureScRsp                           = 2809
	GetBoxingClubInfoCsReq                      = 4252
	GetReplayTokenScRsp                         = 3568
	StartRaidCsReq                              = 2252
	GetTreasureDungeonActivityDataScRsp         = 4470
	LeaveAetherDivideSceneScRsp                 = 4844
	SyncRogueAeonScNotify                       = 1848
	TakeActivityExpeditionRewardCsReq           = 2586
	ActivateFarmElementCsReq                    = 1414
	GetTrainVisitorRegisterCsReq                = 3723
	GroupStateChangeScNotify                    = 1453
	GetFriendLoginInfoCsReq                     = 2981
	ChessRogueReviveAvatarScRsp                 = 5517
	GetRogueBuffEnhanceInfoScRsp                = 1828
	SpringRecoverSingleAvatarCsReq              = 1456
	TakeAllRewardCsReq                          = 3076
	TakeExpeditionRewardCsReq                   = 2523
	AlleyGuaranteedFundsCsReq                   = 4777
	SyncRogueBuffSwapInfoScNotify               = 1879
	GetRaidInfoScRsp                            = 2270
	SceneGroupRefreshScNotify                   = 1460
	SetGroupCustomSaveDataCsReq                 = 1487
	GetServerPrefsDataScRsp                     = 6144
	MissionAcceptScNotify                       = 1238
	ExpeditionDataChangeScNotify                = 2567
	GetCurAssistCsReq                           = 2977
	SetNicknameCsReq                            = 29
	SyncDeleteFriendScNotify                    = 2908
	GiveUpBoxingClubChallengeScRsp              = 4276
	PlayerGetTokenScRsp                         = 42
	BoxingClubChallengeUpdateScNotify           = 4272
	AceAntiCheaterCsReq                         = 13
	SetPlayerInfoCsReq                          = 16
	MuseumTakeCollectRewardCsReq                = 4379
	GetRogueInitialScoreCsReq                   = 1824
	EnterTrialActivityStageCsReq                = 2684
	GetAuthkeyCsReq                             = 47
	SceneEntityDisappearScNotify                = 1472
	QuitLineupScRsp                             = 776
	TrainVisitorBehaviorFinishCsReq             = 3752
	QuestRecordScNotify                         = 967
	GetChatFriendHistoryScRsp                   = 3976
	SyncRogueAdventureRoomInfoScNotify          = 5652
	SceneEntityTeleportScRsp                    = 1454
	FinishCosumeItemMissionCsReq                = 1282
	GetEnhanceCommonRogueBuffInfoScRsp          = 5619
	TakeChallengeTargetRewardCsReq              = 1782
	EnterFantasticStoryActivityStageScRsp       = 4939
	DeleteBlacklistCsReq                        = 2932
	LeaveRogueScRsp                             = 1876
	GetRogueEndlessActivityDataScRsp            = 6068
	InteractPropCsReq                           = 1451
	MatchBoxingClubOpponentScRsp                = 4244
	UpdateServerPrefsDataScRsp                  = 6142
	TakeOffRelicCsReq                           = 303
	PlayerLoginFinishCsReq                      = 54
	SetRedPointStatusScNotify                   = 90
	GetChapterCsReq                             = 439
	RogueChallengeRefreshAssistListScRsp        = 2624
	InteractTreasureDungeonGridCsReq            = 4404
	GetFriendRecommendListInfoCsReq             = 2904
	GetExpeditionDataScRsp                      = 2568
	EntityBindPropScRsp                         = 1473
	SyncRogueRewardInfoScNotify                 = 1871
	TradeChessRogueMiracleScRsp                 = 5476
	ChessRogueGoAheadScRsp                      = 5550
	ChessRogueUpdateAllowedSelectCellScNotify   = 5442
	GetBattleCollegeDataScRsp                   = 5768
	RogueChallengeActivityBuffChooseCsReq       = 2632
	SetDisplayAvatarScRsp                       = 2842
	SyncRogueBuffSelectInfoScNotify             = 1867
	GetRogueAdventureRoomInfoCsReq              = 5686
	SpringRecoverCsReq                          = 1459
	ReforgeRogueBuffScRsp                       = 1835
	PlayerLoginFinishScRsp                      = 58
	UpdateChessRogueMiracleScNotify             = 5485
	WaypointShowNewCsNotify                     = 423
	RepairRogueMiracleScRsp                     = 5637
	SetForbidOtherApplyFriendScRsp              = 2962
	CancelActivityExpeditionScRsp               = 2582
	AlleyTakeEventRewardCsReq                   = 4738
	ChessRogueCheatRollScRsp                    = 5459
	GetLevelRewardTakenListScRsp                = 96
	DailyActiveInfoNotify                       = 3339
	GetAssistHistoryCsReq                       = 2938
	TakePictureCsReq                            = 4139
	HeliobusSnsCommentCsReq                     = 5867
	MarkChatEmojiScRsp                          = 3970
	HeliobusSnsPostScRsp                        = 5842
	RevcMsgScNotify                             = 3951
	GetPlayerDetailInfoCsReq                    = 2951
	StartTrialActivityScRsp                     = 2643
	HeliobusSnsReadScRsp                        = 5844
	ChessRogueGiveUpRollScRsp                   = 5477
	ReviveRogueAvatarCsReq                      = 1804
	UnlockChatBubbleScNotify                    = 5139
	GmTalkScNotify                              = 76
	ChessRogueLeaveCsReq                        = 5585
	BuyNpcStuffScRsp                            = 4344
	ChessRogueUpdateMoneyInfoScNotify           = 5486
	RogueModifierUpdateNotify                   = 5376
	GetAetherDivideInfoScRsp                    = 4870
	SelectChessRogueBonusScRsp                  = 5574
	SyncChessRogueBonusSelectInfoScNotify       = 5418
	PlayerReturnTakeRewardCsReq                 = 4523
	EquipAetherDividePassiveSkillCsReq          = 4808
	PlayerLoginCsReq                            = 52
	TakeFightActivityRewardCsReq                = 3642
	SetLineupNameCsReq                          = 703
	GetRogueHandbookDataCsReq                   = 1812
	SceneUpdatePositionVersionNotify            = 1470
	SceneCastSkillScRsp                         = 1442
	GetTutorialCsReq                            = 1652
	HandleFriendCsReq                           = 2972
	FinishChessRogueSubStoryScRsp               = 5561
	BatchMarkChatEmojiScRsp                     = 3986
	GroupStateChangeScRsp                       = 1491
	ExchangeGachaCeilingCsReq                   = 1923
	GetMonsterResearchActivityDataScRsp         = 2603
	FinishPlotCsReq                             = 1152
	MultipleDropInfoScNotify                    = 4651
	GetPhoneDataCsReq                           = 5152
	PrivateMsgOfflineUsersScNotify              = 3944
	GetChessRogueBuffEnhanceInfoCsReq           = 5430
	GetFirstTalkByPerformanceNpcCsReq           = 2109
	SyncTaskScRsp                               = 1223
	ChessRogueGiveUpScRsp                       = 5453
	GetRogueScoreRewardInfoCsReq                = 1895
	GetRndOptionScRsp                           = 3468
	SyncRogueHandbookMiracleUnlockScNotify      = 5602
	SetClientRaidTargetCountCsReq               = 2282
	GetPlayerReplayInfoScRsp                    = 3544
	ChessRogueReRollDiceCsReq                   = 5516
	AvatarExpUpCsReq                            = 351
	PunkLordRaidTimeOutScNotify                 = 3204
	GetUpdatedArchiveDataCsReq                  = 2351
	UnlockPhoneThemeScNotify                    = 5176
	DropRogueBuffCsReq                          = 5646
	SetStuffToAreaCsReq                         = 4339
	GetBasicInfoCsReq                           = 25
	GetPunkLordDataScRsp                        = 3292
	SetMissionEventProgressCsReq                = 1228
	ChessRogueLeaveScRsp                        = 5410
	StartPunkLordRaidScRsp                      = 3244
	SetFriendRemarkNameScRsp                    = 2919
	ClearAetherDividePassiveSkillCsReq          = 4892
	SyncChessRogueBuffSelectInfoScNotify        = 5520
	SelectChessRogueMiracleScRsp                = 5406
	RepairRogueMiracleCsReq                     = 5700
	DailyRefreshNotify                          = 35
	SetCurWaypointScRsp                         = 444
	GetRogueInitialScoreScRsp                   = 1889
	SelectChessRogueSubStoryCsReq               = 5425
	SwitchAetherDivideLineUpSlotScRsp           = 4829
	SyncChessRogueMiracleInfoScNotify           = 5496
	SelectRogueDialogueEventCsReq               = 1854
	SecurityReportCsReq                         = 4109
	SharePunkLordMonsterScRsp                   = 3242
	FightActivityDataChangeScNotify             = 3651
	GetCurSceneInfoCsReq                        = 1423
	PrepareRogueAdventureRoomCsReq              = 5668
	GetCurSceneInfoScRsp                        = 1476
	SelectChatBubbleScRsp                       = 5144
	RaidKickByServerScNotify                    = 2229
	SetAssistCsReq                              = 2922
	GetPlayerReplayInfoCsReq                    = 3551
	GetFriendApplyListInfoScRsp                 = 2942
	GetGachaCeilingCsReq                        = 1939
	FinishSectionIdCsReq                        = 2723
	AlleyOrderChangedScNotify                   = 4704
	SyncTaskCsReq                               = 1242
	ReviveRogueAvatarScRsp                      = 1829
	SetStuffToAreaScRsp                         = 4342
	TakeMailAttachmentCsReq                     = 823
	SyncRogueMapRoomScNotify                    = 1881
	TakeCityShopRewardCsReq                     = 1539
	SummonPunkLordMonsterScRsp                  = 3276
	SyncRogueBuffReforgeInfoScNotify            = 1825
	GetBagScRsp                                 = 568
	PlayerHeartBeatScRsp                        = 87
	ReserveStaminaExchangeScRsp                 = 50
	SelectChessRogueBuffCsReq                   = 5595
	GetDailyActiveInfoScRsp                     = 3344
	GetEnhanceCommonRogueBuffInfoCsReq          = 5674
	ReturnLastTownCsReq                         = 1474
	AetherDivideRefreshEndlessScRsp             = 4805
	DropChessRogueBuffCsReq                     = 5456
	TakeChallengeRaidRewardCsReq                = 2276
	FinishItemIdCsReq                           = 2739
	TeleportToMissionResetPointCsReq            = 1202
	PVEBattleResultCsReq                        = 152
	ChessRogueUpdateReviveInfoScNotify          = 5575
	TreasureDungeonFinishScNotify               = 4468
	LeaveRaidScRsp                              = 2244
	TakeActivityExpeditionRewardScRsp           = 2508
	ShowNewSupplementVisitorCsReq               = 3709
	GetMailCsReq                                = 852
	ChessRogueQueryAeonDimensionsScRsp          = 5487
	StartBoxingClubBattleCsReq                  = 4239
	UpgradeAreaStatScRsp                        = 4347
	EnteredSceneChangeScNotify                  = 1407
	QuitBattleCsReq                             = 151
	MarkItemScRsp                               = 566
	TakeQuestRewardScRsp                        = 944
	RefreshTriggerByClientScNotify              = 1488
	EnterRogueEndlessActivityStageCsReq         = 6051
	EnterActivityBattleStageScRsp               = 2673
	SyncRogueDialogueEventDataScNotify          = 1865
	PVEBattleResultScRsp                        = 168
	SyncChessRogueMiracleRepairInfoScNotify     = 5494
	GetLoginActivityCsReq                       = 2652
	PrepareRogueAdventureRoomScRsp              = 5651
	ChessRogueSelectCellScRsp                   = 5452
	MissionEventRewardScNotify                  = 1292
	GetRecyleTimeScRsp                          = 579
	TradeRogueMiracleCsReq                      = 1833
	AddChessRogueBuffScNotify                   = 5505
	TakeChallengeRewardScRsp                    = 1747
	GetRogueBuffEnhanceInfoCsReq                = 1896
	ExchangeRogueRewardKeyCsReq                 = 1826
	FinishSectionIdScRsp                        = 2776
	GetPrivateChatHistoryCsReq                  = 3939
	RankUpAvatarCsReq                           = 386
	UseItemCsReq                                = 523
	GetAuthkeyScRsp                             = 92
	TakeRogueEndlessActivityPointRewardScRsp    = 6023
	GetChessRogueStoryInfoScRsp                 = 5547
	RemoveRogueMiracleScNotify                  = 1894
	ApplyFriendScRsp                            = 2976
	ActivateFarmElementScRsp                    = 1462
	GetTutorialScRsp                            = 1668
	ShareScRsp                                  = 4168
	ChooseBoxingClubResonanceCsReq              = 4209
	GetRogueUnlockDataScRsp                     = 5638
	GetAlleyInfoScRsp                           = 4768
	AetherDivideRefreshEndlessCsReq             = 4877
	ChessRogueConfirmRollScRsp                  = 5572
	RankUpEquipmentCsReq                        = 567
	CurTrialActivityScNotify                    = 2636
	GetRogueShopMiracleInfoCsReq                = 5644
	UnlockBackGroundMusicScRsp                  = 3142
	GetChatFriendHistoryCsReq                   = 3923
	SceneEntityUpdateScNotify                   = 1467
	ClientObjUploadCsReq                        = 57
	GetShareDataCsReq                           = 4151
	GetFightActivityDataCsReq                   = 3652
	DelSaveRaidScNotify                         = 2204
	GetKilledPunkLordMonsterDataScRsp           = 3232
	FinishTutorialScRsp                         = 1672
	TrainVisitorBehaviorFinishScRsp             = 3768
	ChessRogueUpdateLevelBaseInfoScNotify       = 5589
	PlayerReturnSignCsReq                       = 4568
	ComposeLimitNumCompleteNotify               = 600
	FantasticStoryActivityBattleEndScNotify     = 4942
	EnterRogueEndlessActivityStageScRsp         = 6044
	EnterSectionScRsp                           = 1479
	StartBattleCollegeCsReq                     = 5744
	SendMsgCsReq                                = 3952
	EnhanceRogueBuffScRsp                       = 1830
	LockRelicCsReq                              = 592
	MarkChatEmojiCsReq                          = 3909
	SetGroupCustomSaveDataScRsp                 = 1443
	SetNicknameScRsp                            = 74
	DropRogueBuffScRsp                          = 5677
	DailyTaskDataScNotify                       = 1276
	GetNpcStatusCsReq                           = 2751
	StartRaidScRsp                              = 2268
	GetLevelRewardCsReq                         = 28
	TakeChallengeTargetRewardScRsp              = 1786
	GetRogueHandbookDataScRsp                   = 1827
	GetAllSaveRaidCsReq                         = 2292
	TakeLoginActivityRewardScRsp                = 2644
	QuitTreasureDungeonCsReq                    = 4496
	PlayerReturnPointChangeScNotify             = 4544
	DailyFirstMeetPamScRsp                      = 3444
	SetGenderCsReq                              = 69
	GetEnteredSceneScRsp                        = 1493
	HeliobusSnsPostCsReq                        = 5839
	FinishChapterScNotify                       = 4951
	QueryProductInfoScRsp                       = 55
	GetQuestDataScRsp                           = 968
	GetSaveLogisticsMapCsReq                    = 4737
	QuitRogueScRsp                              = 1877
	GetRogueAeonInfoScRsp                       = 1897
	SetFriendRemarkNameCsReq                    = 2974
	ChessRogueRollDiceCsReq                     = 5448
	ChessRogueRollDiceScRsp                     = 5560
	InterruptMissionEventCsReq                  = 1219
	EnhanceCommonRogueBuffScRsp                 = 5628
	DressAvatarScRsp                            = 372
	SelectChessRogueMiracleCsReq                = 5483
	AntiAddictScNotify                          = 4
	SelectChessRogueSubStoryScRsp               = 5588
	TakeAssistRewardScRsp                       = 2973
	ChessRogueCheatRollCsReq                    = 5428
	GetServerPrefsDataCsReq                     = 6151
	ServerAnnounceNotify                        = 37
	OpenTreasureDungeonGridScRsp                = 4447
	GetNpcTakenRewardCsReq                      = 2152
	GetStageLineupCsReq                         = 752
	FinishChessRogueSubStoryCsReq               = 5539
	RogueChallengeActivityBuffChooseScRsp       = 2630
	StartFinishMainMissionScNotify              = 1237
	SpringRefreshCsReq                          = 1403
	ChessRogueEnterCellScRsp                    = 5405
	GetQuestRecordCsReq                         = 923
	ChangeLineupLeaderCsReq                     = 786
	HeliobusSnsReadCsReq                        = 5851
	SetTurnFoodSwitchScRsp                      = 516
	GetRogueEndlessActivityDataCsReq            = 6052
	AetherDivideFinishChallengeScNotify         = 4879
	ChessRogueQueryScRsp                        = 5569
	GetBattleCollegeDataCsReq                   = 5752
	FinishTalkMissionCsReq                      = 1251
	FinishPerformSectionIdCsReq                 = 2767
	SetAetherDivideLineUpScRsp                  = 4886
	TakeTrainVisitorUntakenBehaviorRewardScRsp  = 3772
	RankUpAvatarScRsp                           = 308
	LockRelicScRsp                              = 503
	PlayerSyncScNotify                          = 652
	RankUpEquipmentScRsp                        = 572
	StartBattleCollegeScRsp                     = 5739
	AlleyEventEffectNotify                      = 4772
	SyncRogueMiracleRepairInfoScNotify          = 5679
	TakeAllRewardScRsp                          = 3067
	SyncRoguePickAvatarInfoScNotify             = 1861
	ChessRogueSelectBpScRsp                     = 5536
	ChessRogueSelectBpCsReq                     = 5590
	TextJoinBatchSaveCsReq                      = 3839
	LogisticsDetonateStarSkiffCsReq             = 4741
	TriggerVoiceCsReq                           = 4182
	GetExpeditionDataCsReq                      = 2552
	SyncRogueCommonMiracleDisplayScNotify       = 5692
	TradeRogueMiracleScRsp                      = 1885
	ChessRogueReRollDiceScRsp                   = 5409
	GetFirstTalkByPerformanceNpcScRsp           = 2170
	GetFriendRecommendListInfoScRsp             = 2929
	PunkLordDataChangeNotify                    = 3222
	RogueModifierAddNotify                      = 5351
	GetAssistHistoryScRsp                       = 2966
	GetAetherDivideInfoCsReq                    = 4809
	AetherDivideRefreshEndlessScNotify          = 4838
	DestroyItemScRsp                            = 546
	PromoteEquipmentCsReq                       = 551
	EnterSectionCsReq                           = 1402
	HeliobusChallengeUpdateScNotify             = 5828
	ChessRogueQueryCsReq                        = 5455
	AetherDivideLineupScNotify                  = 4846
	SpringRefreshScRsp                          = 1404
	AlleyShipUsedCountScNotify                  = 4746
	UseTreasureDungeonItemScRsp                 = 4419
	SelectPhoneThemeCsReq                       = 5142
	GateServerScNotify                          = 43
	GetSingleRedDotParamGroupCsReq              = 5939
	TextJoinBatchSaveScRsp                      = 3842
	SyncClientResVersionScRsp                   = 176
	PlayerReturnTakePointRewardCsReq            = 4539
	GetCurChallengeCsReq                        = 1772
	GetRogueAdventureRoomInfoScRsp              = 5608
	SyncRogueFinishScNotify                     = 1808
	ExchangeRogueBuffWithMiracleScRsp           = 5629
	ChessRogueUpdateDiceInfoScNotify            = 5503
	GetRogueScoreRewardInfoScRsp                = 1857
	FightTreasureDungeonMonsterCsReq            = 4492
	AlleyGuaranteedFundsScRsp                   = 4705
	TakeRogueMiracleHandbookRewardScRsp         = 1860
	GetTrialActivityDataScRsp                   = 2659
	SellItemScRsp                               = 529
	SetIsDisplayAvatarInfoCsReq                 = 2823
	TakeOffRelicScRsp                           = 304
	GetChatEmojiListScRsp                       = 3972
	GetAlleyInfoCsReq                           = 4752
	GetSecretKeyInfoScRsp                       = 45
	SetPlayerInfoScRsp                          = 24
	GetCurChallengeScRsp                        = 1709
	TakeOffAvatarSkinScRsp                      = 332
	LogisticsScoreRewardSyncInfoScNotify        = 4773
	TakeKilledPunkLordMonsterScoreCsReq         = 3300
	GetFriendListInfoScRsp                      = 2968
	SyncChessRogueMiracleTradeInfoScNotify      = 5507
	RemoveChessRogueBuffScNotify                = 5443
	ChessRogueStartCsReq                        = 5473
	DressRelicAvatarScRsp                       = 392
	HeliobusUpgradeLevelScRsp                   = 5886
	GetMultipleDropInfoScRsp                    = 4668
	GetNpcStatusScRsp                           = 2744
	ChessRogueEnterCellCsReq                    = 5510
	SetAssistScRsp                              = 2946
	QuitRogueCsReq                              = 1846
	RetcodeNotify                               = 75
	SecurityReportScRsp                         = 4170
	DoGachaScRsp                                = 1944
	SyncRogueEventHandbookScNotify              = 1806
	EnterTreasureDungeonScRsp                   = 4486
	GetBoxingClubInfoScRsp                      = 4268
	GetRndOptionCsReq                           = 3452
	StartAetherDivideStageBattleCsReq           = 4874
	ChessRogueGoAheadCsReq                      = 5567
	TakeChallengeRaidRewardScRsp                = 2267
	HeliobusActivityDataCsReq                   = 5852
	TakePrestigeRewardCsReq                     = 4709
	HeliobusUpgradeLevelCsReq                   = 5882
	UnlockSkilltreeScRsp                        = 342
	EnterAetherDivideSceneCsReq                 = 4852
	GetTrainVisitorRegisterScRsp                = 3776
	TextJoinQueryCsReq                          = 3851
	SelectRogueBuffScRsp                        = 1809
	AcceptExpeditionScRsp                       = 2544
	PlayerReturnTakePointRewardScRsp            = 4542
	SyncRogueBuffDropInfoScNotify               = 5622
	UpgradeAreaCsReq                            = 4382
	SwitchAetherDivideLineUpSlotCsReq           = 4804
	MarkItemCsReq                               = 538
	AcceptMissionEventCsReq                     = 1203
	ExchangeRogueBuffWithMiracleCsReq           = 5604
	GetAllLineupDataScRsp                       = 774
	GetPunkLordDataCsReq                        = 3247
	PlayerLogoutCsReq                           = 51
	TakeBpRewardCsReq                           = 3044
	SelectRogueDialogueEventScRsp               = 1858
	SetBoxingClubResonanceLineupCsReq           = 4282
	CancelMarkItemNotify                        = 541
	TakeChallengeRewardCsReq                    = 1708
	GetShopListScRsp                            = 1568
	FinishTutorialGuideScRsp                    = 1670
	ChessRogueStartScRsp                        = 5548
	EnterRogueMapRoomCsReq                      = 1873
	GetPunkLordBattleRecordCsReq                = 3246
	ChessRogueUpdateAeonModifierValueScNotify   = 5525
	GetLoginChatInfoScRsp                       = 3947
	GetTrainVisitorBehaviorCsReq                = 3751
	EnterTrialActivityStageScRsp                = 2656
	RogueNpcDisappearCsReq                      = 5670
	PickRogueAvatarScRsp                        = 1892
	GetMissionEventDataScRsp                    = 1247
	GetRogueTalentInfoScRsp                     = 1821
	BattleLogReportCsReq                        = 172
	UpdatePlayerSettingScRsp                    = 18
	StartFinishSubMissionScNotify               = 1300
	SwapLineupScRsp                             = 772
	StartChallengeScRsp                         = 1744
	FinishFirstTalkNpcScRsp                     = 2176
	FeatureSwitchClosedScNotify                 = 33
	GameplayCounterCountDownScRsp               = 1457
	SyncEntityBuffChangeListScNotify            = 1482
	StartAetherDivideSceneBattleScRsp           = 4842
	TakeRogueEndlessActivityAllBonusRewardScRsp = 6067
	MuseumRandomEventSelectCsReq                = 4319
	FinishTutorialCsReq                         = 1667
	AlleyTakeEventRewardScRsp                   = 4766
	PunkLordBattleResultScNotify                = 3296
	BuyRogueShopBuffScRsp                       = 5609
	TakeOffEquipmentScRsp                       = 370
	GetChessRogueStoryAeonTalkInfoCsReq         = 5500
	LastSpringRefreshTimeNotify                 = 1429
	CurAssistChangedNotify                      = 2916
	GetFriendApplyListInfoCsReq                 = 2939
	OpenRogueChestCsReq                         = 1856
	TakeOffEquipmentCsReq                       = 309
	DeleteFriendCsReq                           = 2982
	GetAvatarDataCsReq                          = 352
	ExpUpRelicScRsp                             = 547
	ChallengeLineupNotify                       = 1770
	HeliobusEnterBattleScRsp                    = 5874
	GetSpringRecoverDataCsReq                   = 1425
	GetRogueTalentInfoCsReq                     = 1810
	BuyBpLevelScRsp                             = 3023
	AvatarExpUpScRsp                            = 344
	ComposeSelectedRelicCsReq                   = 528
	ChangeLineupLeaderScRsp                     = 708
	GetVideoVersionKeyScRsp                     = 48
	DelMailScRsp                                = 842
	AetherDivideSpiritExpUpCsReq                = 4896
	SyncRogueAreaUnlockScNotify                 = 1890
	GetEnteredSceneCsReq                        = 1497
	GetQuestDataCsReq                           = 952
	SyncLineupNotify                            = 709
	ChessRogueLayerAccountInfoNotify            = 5558
	EnterAetherDivideSceneScRsp                 = 4868
	AcceptMissionEventScRsp                     = 1204
	EnhanceChessRogueBuffScRsp                  = 5479
	GetLevelRewardScRsp                         = 32
	DressRelicAvatarCsReq                       = 347
	SceneEntityTeleportCsReq                    = 1445
	FinishFirstTalkByPerformanceNpcScRsp        = 2186
	AlleyPlacingGameScRsp                       = 4786
	PlayerReturnStartScNotify                   = 4552
	GetPhoneDataScRsp                           = 5168
	EnterRogueMapRoomScRsp                      = 1816
	ChessRogueQuitScRsp                         = 5512
	ComposeSelectedRelicScRsp                   = 532
	GetPunkLordBattleRecordScRsp                = 3277
	GetSaveRaidCsReq                            = 2208
	GetLoginChatInfoCsReq                       = 3908
	SetForbidOtherApplyFriendCsReq              = 2914
	CancelExpeditionScRsp                       = 2542
	SetGameplayBirthdayCsReq                    = 64
	SetCurInteractEntityScRsp                   = 1446
	ChessRogueGiveUpRollCsReq                   = 5538
	InterruptMissionEventScRsp                  = 1296
	GetTutorialGuideCsReq                       = 1651
	GetCurLineupDataCsReq                       = 751
	EnterFantasticStoryActivityStageCsReq       = 4944
	GetCurBattleInfoCsReq                       = 139
	TriggerVoiceScRsp                           = 4186
	DoGachaCsReq                                = 1951
	GetPlayerReturnMultiDropInfoCsReq           = 4644
	GetAllRedDotDataCsReq                       = 5952
	GetMainMissionCustomValueCsReq              = 1277
	HeliobusSelectSkillCsReq                    = 5847
	GetSceneMapInfoCsReq                        = 1480
	FinishCurTurnScRsp                          = 4370
	EquipAetherDividePassiveSkillScRsp          = 4847
	DeleteSummonUnitScRsp                       = 1411
	ChallengeSettleNotify                       = 1723
	SyncRogueStatusScNotify                     = 1834
	PromoteAvatarCsReq                          = 323
	BattlePassInfoNotify                        = 3052
	CancelActivityExpeditionCsReq               = 2570
	SetIsDisplayAvatarInfoScRsp                 = 2876
	PickRogueAvatarCsReq                        = 1847
	GetRogueShopMiracleInfoScRsp                = 5639
	StartCocoonStageCsReq                       = 1466
	GetMuseumInfoScRsp                          = 4368
	AetherDivideTakeChallengeRewardScRsp        = 4841
	EntityBindPropCsReq                         = 1469
	SetMissionEventProgressScRsp                = 1232
	GetShopListCsReq                            = 1552
	GetArchiveDataCsReq                         = 2352
	SelectRogueBonusScRsp                       = 1862
	PrestigeLevelUpCsReq                        = 4774
	MarkReadMailCsReq                           = 851
	SyncRogueVirtualItemInfoScNotify            = 1817
	GameplayCounterUpdateScNotify               = 1431
	ChessRogueQuestFinishNotify                 = 5573
	SyncChessRogueMainStoryFinishScNotify       = 5537
	AcceptMainMissionCsReq                      = 1222
	DropChessRogueBuffScRsp                     = 5521
	ChooseBoxingClubResonanceScRsp              = 4270
	TextJoinQueryScRsp                          = 3844
	MuseumFundsChangedScNotify                  = 4303
	GetCurAssistScRsp                           = 2905
	HeliobusLineupUpdateScNotify                = 5832
	GetHeroBasicTypeInfoCsReq                   = 77
	TriggerHealVoiceCsReq                       = 4123
	RogueModifierDelNotify                      = 5367
	ExchangeRogueRewardKeyScRsp                 = 1887
	GetMarkItemListScRsp                        = 505
	GetChallengeRaidInfoScRsp                   = 2223
	DressAvatarSkinCsReq                        = 319
	DressAvatarSkinScRsp                        = 396
	HeliobusInfoChangedScNotify                 = 5870
	GetAvatarDataScRsp                          = 368
	ExpUpEquipmentCsReq                         = 509
	SceneCastSkillCsReq                         = 1439
	SyncChessRogueBuffReforgeInfoScNotify       = 5458
	EnterAdventureScRsp                         = 1368
	SyncRogueMiracleHandbookScNotify            = 1883
	UseTreasureDungeonItemCsReq                 = 4474
	StartAetherDivideStageBattleScRsp           = 4819
	CancelCacheNotifyCsReq                      = 4167
	RemoveStuffFromAreaScRsp                    = 4376
	AcceptExpeditionCsReq                       = 2551
	LogisticsGameScRsp                          = 4744
	UpdateRedDotDataScRsp                       = 5944
	ClearAetherDividePassiveSkillScRsp          = 4803
	GetMissionStatusScRsp                       = 1274
	AcceptActivityExpeditionCsReq               = 2572
	AetherDivideSkillItemScNotify               = 4837
	UnlockBackGroundMusicCsReq                  = 3139
	RefreshTriggerByClientCsReq                 = 1410
	SwapLineupCsReq                             = 767
	StartAetherDivideChallengeBattleCsReq       = 4823
	FinishRogueDialogueGroupCsReq               = 1813
	SyncTurnFoodNotify                          = 569
	PlayerReturnTakeRewardScRsp                 = 4576
	MuseumRandomEventStartScNotify              = 4304
	GetAssistListScRsp                          = 2937
	ComposeItemScRsp                            = 586
	MarkReadMailScRsp                           = 844
	SwapRogueBuffScRsp                          = 1837
	RemoveRogueBuffScNotify                     = 1802
	TakeRogueEndlessActivityAllBonusRewardCsReq = 6076
	TakeCityShopRewardScRsp                     = 1542
	RefreshTriggerByClientScRsp                 = 1421
	RogueModifierStageStartNotify               = 5372
	SetGameplayBirthdayScRsp                    = 59
	RogueModifierSelectCellCsReq                = 5344
	RepairChessRogueMiracleScRsp                = 5570
	PlayerReturnForceFinishScNotify             = 4509
	JoinLineupCsReq                             = 739
	HeliobusUnlockSkillScNotify                 = 5808
	ChessRogueConfirmRollCsReq                  = 5475
	MuseumRandomEventQueryCsReq                 = 4329
	SyncRogueBonusSelectInfoScNotify            = 1855
	PlayerHeartBeatCsReq                        = 26
	UpdateRedDotDataCsReq                       = 5951
	GetPlatformPlayerInfoScRsp                  = 2989
	StartPunkLordRaidCsReq                      = 3251
	ChooseBoxingClubStageOptionalBuffCsReq      = 4208
	GetPlatformPlayerInfoCsReq                  = 2924
	ExpUpRelicCsReq                             = 508
	GetNpcMessageGroupCsReq                     = 2752
	StartAlleyEventCsReq                        = 4723
	GetBasicInfoScRsp                           = 99
	GetGachaInfoScRsp                           = 1968
	SelectRogueMiracleScRsp                     = 1869
	RegionStopScNotify                          = 3
	FinishQuestCsReq                            = 972
	GetPlayerBoardDataScRsp                     = 2868
	MuseumDispatchFinishedScNotify              = 4328
	ChessRogueUpdateUnlockLevelScNotify         = 5449
	RogueEndlessActivityBattleEndScNotify       = 6039
	GetCurLineupDataScRsp                       = 744
	ExchangeStaminaCsReq                        = 86
	RaidInfoNotify                              = 2239
	SummonPunkLordMonsterCsReq                  = 3223
	TakeMonsterResearchActivityRewardCsReq      = 2674
	ChessRogueEnterScRsp                        = 5444
	GetChallengeRaidInfoCsReq                   = 2242
	SetAssistAvatarScRsp                        = 2882
	SetLanguageCsReq                            = 79
	GetHeroBasicTypeInfoScRsp                   = 5
	GetCurBattleInfoScRsp                       = 142
	TakeExpeditionRewardScRsp                   = 2576
	TakeAllApRewardCsReq                        = 3342
	TakePrestigeRewardScRsp                     = 4770
	EnterRogueScRsp                             = 1842
	SceneEntityMoveScNotify                     = 1409
	ChessRogueQueryAeonDimensionsCsReq          = 5482
	TradeChessRogueMiracleCsReq                 = 5513
	RollRogueBuffCsReq                          = 1870
	GmTalkCsReq                                 = 72
	UpdateRogueMiracleScNotify                  = 1845
	EnterFightActivityStageCsReq                = 3644
	UnlockTutorialScRsp                         = 1642
	HeliobusSnsLikeScRsp                        = 5876
	LeaveTrialActivityCsReq                     = 2633
	FinishAeonDialogueGroupScRsp                = 1807
	ClientObjUploadScRsp                        = 31
	VirtualLineupDestroyNotify                  = 719
	EnableRogueTalentScRsp                      = 1878
	GetSceneMapInfoScRsp                        = 1498
	TakeTrainVisitorUntakenBehaviorRewardCsReq  = 3767
	GetJukeboxDataCsReq                         = 3152
	SetHeroBasicTypeScRsp                       = 46
	GetFantasticStoryActivityDataScRsp          = 4968
	TakeChapterRewardCsReq                      = 476
	GetExhibitScNotify                          = 4372
	GetFirstTalkNpcScRsp                        = 2142
	TrainVisitorRewardSendNotify                = 3742
	AcceptActivityExpeditionScRsp               = 2509
	SelectChessRogueBuffScRsp                   = 5469
	SetAetherDivideLineUpCsReq                  = 4882
	BuyRogueShopMiracleScRsp                    = 5667
	PlayBackGroundMusicScRsp                    = 3144
	TextJoinSaveCsReq                           = 3852
	SyncApplyFriendScNotify                     = 2967
	StartAlleyEventScRsp                        = 4776
	SearchPlayerScRsp                           = 2979
	SyncHandleFriendScNotify                    = 2970
	BattleLogReportScRsp                        = 109
	MuseumInfoChangedScNotify                   = 4392
	OpenRogueChestScRsp                         = 1875
	HeliobusStartRaidCsReq                      = 5819
	GetMissionEventDataCsReq                    = 1208
	TakeApRewardCsReq                           = 3352
	TakeTalkRewardCsReq                         = 2151
	LeaveAetherDivideSceneCsReq                 = 4851
	RollRogueBuffScRsp                          = 1882
	EnterAdventureCsReq                         = 1352
	ExpUpEquipmentScRsp                         = 570
	MuseumTakeCollectRewardScRsp                = 4400
	SyncChessRogueMiracleSelectInfoScNotify     = 5518
	HeliobusSnsLikeCsReq                        = 5823
	MissionGroupWarnScNotify                    = 1270
	TakeRogueAeonLevelRewardScRsp               = 1801
	SetBoxingClubResonanceLineupScRsp           = 4286
	SyncRogueMiracleTradeInfoScNotify           = 1843
	GetRogueInfoScRsp                           = 1868
	SubmitMonsterResearchActivityMaterialCsReq  = 2604
	HeliobusSnsUpdateScNotify                   = 5809
	HeroBasicTypeChangedNotify                  = 89
	GetMarkItemListCsReq                        = 577
	HandleFriendScRsp                           = 2909
	ShareCsReq                                  = 4152
	GetSecretKeyInfoCsReq                       = 94
	HeliobusStartRaidScRsp                      = 5896
	GetChatEmojiListCsReq                       = 3967
	TakeQuestOptionalRewardScRsp                = 982
	FinishItemIdScRsp                           = 2742
	ScenePlaneEventScNotify                     = 1490
	ChessRogueQueryBpScRsp                      = 5490
	FinishFirstTalkByPerformanceNpcCsReq        = 2182
	EnterFightActivityStageScRsp                = 3639
	GetWaypointCsReq                            = 452
	ChallengeRaidNotify                         = 2272
	AlleyShipmentEventEffectsScNotify           = 4800
	TakeAssistRewardCsReq                       = 2969
	SearchPlayerCsReq                           = 2902
	ReplaceChessRogueMiracleDisplayScNotify     = 5407
	SetCurInteractEntityCsReq                   = 1422
	CancelExpeditionCsReq                       = 2539
	RogueChallengeBattleResultScNotify          = 2689
	StartAetherDivideChallengeBattleScRsp       = 4876
	PrestigeLevelUpScRsp                        = 4719
	GetAetherDivideChallengeInfoCsReq           = 4830
	DeleteBlacklistScRsp                        = 2930
	GetBagCsReq                                 = 552
	EnterRogueCsReq                             = 1839
	GetPlayerDetailInfoScRsp                    = 2944
	AlleyPlacingGameCsReq                       = 4782
	SetLanguageScRsp                            = 100
	ReturnLastTownScRsp                         = 1419
	LogisticsGameCsReq                          = 4751
	GetShareDataScRsp                           = 4144
	LeaveRogueCsReq                             = 1823
	GetChessRogueStoryAeonTalkInfoScRsp         = 5544
	SyncRogueUnlockDataUpdateScNotify           = 5666
	RogueChallengeActivityDataCsReq             = 2696
	ChessRogueChangeyAeonDimensionNotify        = 5438
	GetTutorialGuideScRsp                       = 1644
	PlayBackGroundMusicCsReq                    = 3151
	SetHeroBasicTypeCsReq                       = 22
	AddEquipmentScNotify                        = 530
	LockEquipmentScRsp                          = 542
	ReforgeChessRogueBuffCsReq                  = 5411
	ChessRogueCellUpdateNotify                  = 5592
	SharePunkLordMonsterCsReq                   = 3239
	RecoverAllLineupCsReq                       = 1477
	FinishRogueDialogueGroupScRsp               = 1884
	ChessRogueQuitCsReq                         = 5509
	EnhanceRogueBuffCsReq                       = 1832
	GetQuestRecordScRsp                         = 976
	TakeTalkRewardScRsp                         = 2144
	GetKilledPunkLordMonsterDataCsReq           = 3228
	SetDisplayAvatarCsReq                       = 2839
	InteractTreasureDungeonGridScRsp            = 4429
	ChessRogueEnterNextLayerScRsp               = 5594
	TakeFightActivityRewardScRsp                = 3623
	AlleyFundsScNotify                          = 4796
	HeliobusSelectSkillScRsp                    = 5892
	GetUpdatedArchiveDataScRsp                  = 2344
	TakeKilledPunkLordMonsterScoreScRsp         = 3237
	ApplyFriendCsReq                            = 2923
	SubmitMonsterResearchActivityMaterialScRsp  = 2629
	ReplaceLineupCsReq                          = 796
	QueryProductInfoCsReq                       = 81
	GetPlayerBoardDataCsReq                     = 2852
	UnlockAvatarSkinScNotify                    = 330
	AlleyEventChangeNotify                      = 4767
	GetReplayTokenCsReq                         = 3552
	GetChallengeCsReq                           = 1752
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(AddRogueBuffScNotify, func() any { return new(proto.AddRogueBuffScNotify) })
	c.regMsg(AddRogueMiracleScNotify, func() any { return new(proto.AddRogueMiracleScNotify) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
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
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
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
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(RollRogueBuffScRsp, func() any { return new(proto.RollRogueBuffScRsp) })
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
	c.regMsg(SelectRogueBuffCsReq, func() any { return new(proto.SelectRogueBuffCsReq) })
	c.regMsg(SelectRogueBuffScRsp, func() any { return new(proto.SelectRogueBuffScRsp) })
	c.regMsg(SelectRogueDialogueEventCsReq, func() any { return new(proto.SelectRogueDialogueEventCsReq) })
	c.regMsg(SelectRogueDialogueEventScRsp, func() any { return new(proto.SelectRogueDialogueEventScRsp) })
	c.regMsg(SelectRogueMiracleCsReq, func() any { return new(proto.SelectRogueMiracleCsReq) })
	c.regMsg(SelectRogueMiracleScRsp, func() any { return new(proto.SelectRogueMiracleScRsp) })
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
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(SyncRogueBuffSelectInfoScNotify, func() any { return new(proto.SyncRogueBuffSelectInfoScNotify) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SyncRogueMiracleSelectInfoScNotify, func() any { return new(proto.SyncRogueMiracleSelectInfoScNotify) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
}
