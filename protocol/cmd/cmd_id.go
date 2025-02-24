package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

const (
	TakeLoginActivityRewardScRsp                       = 2627
	GetMaterialSubmitActivityDataCsReq                 = 2675
	TakeLoginActivityRewardCsReq                       = 2684
	TakeMaterialSubmitActivityRewardCsReq              = 2653
	TakeTrialActivityRewardScRsp                       = 2699
	TakeTrialActivityRewardCsReq                       = 2672
	SubmitMaterialSubmitActivityMaterialScRsp          = 2607
	TakeMaterialSubmitActivityRewardScRsp              = 2690
	StartTrialActivityCsReq                            = 2603
	EnterTrialActivityStageCsReq                       = 2645
	EnterTrialActivityStageScRsp                       = 2664
	LeaveTrialActivityCsReq                            = 2661
	GetLoginActivityCsReq                              = 2636
	SubmitMaterialSubmitActivityMaterialCsReq          = 2626
	CurTrialActivityScNotify                           = 2633
	StartTrialActivityScRsp                            = 2640
	LeaveTrialActivityScRsp                            = 2630
	GetActivityScheduleConfigCsReq                     = 2667
	GetTrialActivityDataScRsp                          = 2692
	GetTrialActivityDataCsReq                          = 2694
	GetLoginActivityScRsp                              = 2695
	GetActivityScheduleConfigScRsp                     = 2628
	TrialActivityDataChangeScNotify                    = 2660
	GetMaterialSubmitActivityDataScRsp                 = 2619
	EnterAdventureCsReq                                = 1336
	GetFarmStageGachaInfoCsReq                         = 1384
	QuickStartCocoonStageScRsp                         = 1328
	QuickStartFarmElementCsReq                         = 1352
	QuickStartCocoonStageCsReq                         = 1367
	GetFarmStageGachaInfoScRsp                         = 1327
	QuickStartFarmElementScRsp                         = 1374
	EnterAdventureScRsp                                = 1395
	AetherDivideSpiritExpUpCsReq                       = 4859
	SwitchAetherDivideLineUpSlotCsReq                  = 4826
	GetAetherDivideInfoCsReq                           = 4834
	GetAetherDivideChallengeInfoCsReq                  = 4829
	EnterAetherDivideSceneScRsp                        = 4895
	SetAetherDivideLineUpCsReq                         = 4846
	AetherDivideSpiritExpUpScRsp                       = 4812
	AetherDivideSkillItemScNotify                      = 4858
	LeaveAetherDivideSceneScRsp                        = 4827
	AetherDivideRefreshEndlessCsReq                    = 4835
	AetherDivideSpiritInfoScNotify                     = 4878
	EquipAetherDividePassiveSkillCsReq                 = 4896
	StartAetherDivideChallengeBattleCsReq              = 4852
	StartAetherDivideStageBattleCsReq                  = 4853
	GetAetherDivideChallengeInfoScRsp                  = 4850
	AetherDivideRefreshEndlessScNotify                 = 4817
	ClearAetherDividePassiveSkillCsReq                 = 4875
	StartAetherDivideSceneBattleScRsp                  = 4828
	AetherDivideTainerInfoScNotify                     = 4888
	AetherDivideFinishChallengeScNotify                = 4848
	ClearAetherDividePassiveSkillScRsp                 = 4819
	AetherDivideRefreshEndlessScRsp                    = 4873
	AetherDivideTakeChallengeRewardCsReq               = 4813
	AetherDivideLineupScNotify                         = 4876
	StartAetherDivideSceneBattleCsReq                  = 4867
	StartAetherDivideChallengeBattleScRsp              = 4874
	LeaveAetherDivideSceneCsReq                        = 4884
	SetAetherDivideLineUpScRsp                         = 4825
	AetherDivideTakeChallengeRewardScRsp               = 4821
	SwitchAetherDivideLineUpSlotScRsp                  = 4807
	GetAetherDivideInfoScRsp                           = 4843
	EquipAetherDividePassiveSkillScRsp                 = 4805
	StartAetherDivideStageBattleScRsp                  = 4890
	EnterAetherDivideSceneCsReq                        = 4836
	GetAlleyInfoScRsp                                  = 4795
	AlleyShipUsedCountScNotify                         = 4776
	LogisticsScoreRewardSyncInfoScNotify               = 4708
	AlleyTakeEventRewardScRsp                          = 4713
	GetSaveLogisticsMapCsReq                           = 4758
	GetSaveLogisticsMapScRsp                           = 4710
	StartAlleyEventScRsp                               = 4774
	LogisticsInfoScNotify                              = 4748
	AlleyPlacingGameCsReq                              = 4746
	ActivityRaidPlacingGameScRsp                       = 4737
	RefreshAlleyOrderScRsp                             = 4719
	PrestigeLevelUpCsReq                               = 4753
	AlleyShipmentEventEffectsScNotify                  = 4788
	AlleyPlacingGameScRsp                              = 4725
	AlleyEventEffectNotify                             = 4793
	RefreshAlleyOrderCsReq                             = 4775
	PrestigeLevelUpScRsp                               = 4790
	ActivityRaidPlacingGameCsReq                       = 4718
	StartAlleyEventCsReq                               = 4752
	AlleyGuaranteedFundsCsReq                          = 4735
	LogisticsGameCsReq                                 = 4784
	SaveLogisticsScRsp                                 = 4750
	AlleyGuaranteedFundsScRsp                          = 4773
	TakePrestigeRewardScRsp                            = 4743
	AlleyShipUnlockScNotify                            = 4778
	GetAlleyInfoCsReq                                  = 4736
	AlleyFundsScNotify                                 = 4759
	AlleyOrderChangedScNotify                          = 4726
	AlleyShopLevelScNotify                             = 4712
	TakePrestigeRewardCsReq                            = 4734
	LogisticsDetonateStarSkiffCsReq                    = 4721
	AlleyTakeEventRewardCsReq                          = 4717
	SaveLogisticsCsReq                                 = 4729
	LogisticsDetonateStarSkiffScRsp                    = 4765
	AlleyEventChangeNotify                             = 4724
	LogisticsGameScRsp                                 = 4727
	GetUpdatedArchiveDataCsReq                         = 2384
	GetUpdatedArchiveDataScRsp                         = 2327
	GetArchiveDataCsReq                                = 2336
	GetArchiveDataScRsp                                = 2395
	TakeOffEquipmentScRsp                              = 343
	TakeOffAvatarSkinScRsp                             = 378
	GetAvatarDataCsReq                                 = 336
	PromoteAvatarScRsp                                 = 374
	MarkAvatarCsReq                                    = 350
	TakeOffRelicCsReq                                  = 319
	DressAvatarCsReq                                   = 324
	PromoteAvatarCsReq                                 = 352
	GrowthTargetAvatarChangedScNotify                  = 376
	RankUpAvatarCsReq                                  = 325
	DressAvatarSkinCsReq                               = 390
	TakePromotionRewardScRsp                           = 353
	TakeOffRelicScRsp                                  = 326
	AddAvatarScNotify                                  = 346
	AddMultiPathAvatarScNotify                         = 388
	UnlockSkilltreeScRsp                               = 328
	TakeOffAvatarSkinCsReq                             = 312
	AvatarExpUpScRsp                                   = 327
	DressAvatarSkinScRsp                               = 359
	DressRelicAvatarCsReq                              = 305
	UnlockAvatarSkinScNotify                           = 329
	DressAvatarScRsp                                   = 393
	SetGrowthTargetAvatarCsReq                         = 358
	AvatarExpUpCsReq                                   = 384
	DressRelicAvatarScRsp                              = 375
	TakePromotionRewardCsReq                           = 307
	MarkAvatarScRsp                                    = 348
	TakeOffEquipmentCsReq                              = 334
	UnlockSkilltreeCsReq                               = 367
	GetAvatarDataScRsp                                 = 395
	RankUpAvatarScRsp                                  = 396
	SetGrowthTargetAvatarScRsp                         = 310
	QuitBattleScNotify                                 = 124
	GetCurBattleInfoCsReq                              = 167
	PVEBattleResultScRsp                               = 195
	PVEBattleResultCsReq                               = 136
	SyncClientResVersionCsReq                          = 152
	QuitBattleCsReq                                    = 184
	ServerSimulateBattleFinishScNotify                 = 143
	QuitBattleScRsp                                    = 127
	BattleLogReportScRsp                               = 134
	BattleLogReportCsReq                               = 193
	ReBattleAfterBattleLoseCsNotify                    = 146
	SyncClientResVersionScRsp                          = 174
	GetCurBattleInfoScRsp                              = 128
	RebattleByClientCsNotify                           = 125
	BattleCollegeDataChangeScNotify                    = 5784
	StartBattleCollegeCsReq                            = 5727
	GetBattleCollegeDataCsReq                          = 5736
	GetBattleCollegeDataScRsp                          = 5795
	StartBattleCollegeScRsp                            = 5767
	BuyBpLevelCsReq                                    = 3028
	BuyBpLevelScRsp                                    = 3052
	TakeAllRewardScRsp                                 = 3024
	BattlePassInfoNotify                               = 3036
	TakeBpRewardCsReq                                  = 3027
	TakeBpRewardScRsp                                  = 3067
	TakeAllRewardCsReq                                 = 3074
	SetBoxingClubResonanceLineupScRsp                  = 4225
	BoxingClubRewardScNotify                           = 4224
	GiveUpBoxingClubChallengeCsReq                     = 4252
	MatchBoxingClubOpponentScRsp                       = 4227
	SetBoxingClubResonanceLineupCsReq                  = 4246
	ChooseBoxingClubStageOptionalBuffScRsp             = 4205
	BoxingClubChallengeUpdateScNotify                  = 4293
	ChooseBoxingClubStageOptionalBuffCsReq             = 4296
	ChooseBoxingClubResonanceCsReq                     = 4234
	MatchBoxingClubOpponentCsReq                       = 4284
	StartBoxingClubBattleScRsp                         = 4228
	GiveUpBoxingClubChallengeScRsp                     = 4274
	ChooseBoxingClubResonanceScRsp                     = 4243
	GetBoxingClubInfoScRsp                             = 4295
	StartBoxingClubBattleCsReq                         = 4267
	GetBoxingClubInfoCsReq                             = 4236
	GetCurChallengeScRsp                               = 1734
	StartChallengeCsReq                                = 1784
	LeaveChallengeScRsp                                = 1728
	ChallengeSettleNotify                              = 1752
	StartChallengeScRsp                                = 1727
	RestartChallengePhaseCsReq                         = 1790
	GetCurChallengeCsReq                               = 1793
	GetChallengeGroupStatisticsScRsp                   = 1719
	TakeChallengeRewardCsReq                           = 1796
	ChallengeBossPhaseSettleNotify                     = 1729
	StartPartialChallengeScRsp                         = 1707
	EnterChallengeNextPhaseCsReq                       = 1712
	EnterChallengeNextPhaseScRsp                       = 1778
	StartPartialChallengeCsReq                         = 1726
	GetChallengeScRsp                                  = 1795
	ChallengeLineupNotify                              = 1743
	TakeChallengeRewardScRsp                           = 1705
	GetChallengeGroupStatisticsCsReq                   = 1775
	RestartChallengePhaseScRsp                         = 1759
	LeaveChallengeCsReq                                = 1767
	GetChallengeCsReq                                  = 1736
	GetChatEmojiListCsReq                              = 3924
	RevcMsgScNotify                                    = 3984
	MarkChatEmojiScRsp                                 = 3943
	GetLoginChatInfoCsReq                              = 3996
	BatchMarkChatEmojiScRsp                            = 3925
	GetPrivateChatHistoryScRsp                         = 3928
	MarkChatEmojiCsReq                                 = 3934
	SendMsgScRsp                                       = 3995
	GetPrivateChatHistoryCsReq                         = 3967
	SendMsgCsReq                                       = 3936
	BatchMarkChatEmojiCsReq                            = 3946
	GetChatFriendHistoryCsReq                          = 3952
	PrivateMsgOfflineUsersScNotify                     = 3927
	GetChatEmojiListScRsp                              = 3993
	GetLoginChatInfoScRsp                              = 3905
	GetChatFriendHistoryScRsp                          = 3974
	ChessRogueGiveUpRollScRsp                          = 5509
	ChessRoguePickAvatarScRsp                          = 5406
	ChessRogueQuitCsReq                                = 5563
	ChessRogueRollDiceCsReq                            = 5485
	ChessRogueReviveAvatarScRsp                        = 5434
	ChessRogueCheatRollCsReq                           = 5595
	ChessRogueNousEditDiceCsReq                        = 5472
	ChessRogueGiveUpScRsp                              = 5433
	SelectChessRogueSubStoryCsReq                      = 5496
	ChessRogueNousDiceUpdateNotify                     = 5585
	GetChessRogueStoryAeonTalkInfoCsReq                = 5483
	ChessRogueQueryCsReq                               = 5492
	ChessRogueUpdateUnlockLevelScNotify                = 5533
	GetChessRogueNousStoryInfoScRsp                    = 5567
	ChessRogueConfirmRollScRsp                         = 5514
	FinishChessRogueNousSubStoryCsReq                  = 5519
	ChessRogueQueryBpScRsp                             = 5431
	ChessRogueCellUpdateNotify                         = 5442
	ChessRogueLeaveCsReq                               = 5558
	GetChessRogueStoryAeonTalkInfoScRsp                = 5553
	ChessRogueQueryAeonDimensionsCsReq                 = 5518
	ChessRogueQueryAeonDimensionsScRsp                 = 5463
	ChessRogueUpdateActionPointScNotify                = 5543
	ChessRogueCheatRollScRsp                           = 5411
	SelectChessRogueNousSubStoryScRsp                  = 5529
	ChessRogueGoAheadCsReq                             = 5520
	ChessRogueUpdateReviveInfoScNotify                 = 5576
	EnhanceChessRogueBuffScRsp                         = 5541
	ChessRogueUpdateLevelBaseInfoScNotify              = 5469
	SelectChessRogueSubStoryScRsp                      = 5405
	ChessRogueEnterCsReq                               = 5441
	ChessRogueNousGetRogueTalentInfoCsReq              = 5513
	ChessRogueReRollDiceScRsp                          = 5477
	GetChessRogueBuffEnhanceInfoScRsp                  = 5414
	FinishChessRogueNousSubStoryScRsp                  = 5584
	FinishChessRogueSubStoryScRsp                      = 5479
	ChessRogueChangeyAeonDimensionNotify               = 5437
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5454
	FinishChessRogueSubStoryCsReq                      = 5516
	ChessRoguePickAvatarCsReq                          = 5580
	ChessRogueSelectCellScRsp                          = 5419
	GetChessRogueNousStoryInfoCsReq                    = 5422
	ChessRogueQuitScRsp                                = 5435
	ChessRogueLayerAccountInfoNotify                   = 5586
	ChessRogueUpdateBoardScNotify                      = 5490
	ChessRogueNousEnableRogueTalentScRsp               = 5596
	GetChessRogueStoryInfoCsReq                        = 5557
	ChessRogueGiveUpRollCsReq                          = 5445
	ChessRogueNousEnableRogueTalentCsReq               = 5493
	ChessRogueEnterNextLayerScRsp                      = 5583
	GetChessRogueBuffEnhanceInfoCsReq                  = 5457
	ChessRogueEnterCellCsReq                           = 5581
	ChessRogueReviveAvatarCsReq                        = 5473
	ChessRogueRollDiceScRsp                            = 5425
	ChessRogueQueryBpCsReq                             = 5545
	EnterChessRogueAeonRoomScRsp                       = 5552
	EnterChessRogueAeonRoomCsReq                       = 5482
	ChessRogueSelectCellCsReq                          = 5506
	ChessRogueSkipTeachingLevelCsReq                   = 5528
	ChessRogueStartCsReq                               = 5452
	ChessRogueGoAheadScRsp                             = 5578
	ChessRogueConfirmRollCsReq                         = 5591
	ChessRogueLeaveScRsp                               = 5562
	ChessRogueNousEditDiceScRsp                        = 5443
	ChessRogueUpdateDiceInfoScNotify                   = 5536
	ChessRogueUpdateAllowedSelectCellScNotify          = 5521
	ChessRogueNousGetRogueTalentInfoScRsp              = 5432
	ChessRogueEnterScRsp                               = 5510
	ChessRogueFinishCurRoomNotify                      = 5564
	ChessRogueEnterCellScRsp                           = 5556
	SyncChessRogueNousMainStoryScNotify                = 5497
	ChessRogueSkipTeachingLevelScRsp                   = 5577
	EnhanceChessRogueBuffCsReq                         = 5550
	SyncChessRogueNousValueScNotify                    = 5589
	ChessRogueStartScRsp                               = 5582
	GetChessRogueStoryInfoScRsp                        = 5530
	ChessRogueNousDiceSurfaceUnlockNotify              = 5568
	ChessRogueEnterNextLayerCsReq                      = 5440
	ChessRogueUpdateAeonModifierValueScNotify          = 5438
	SelectChessRogueNousSubStoryCsReq                  = 5409
	SyncChessRogueNousSubStoryScNotify                 = 5498
	ChessRogueMoveCellNotify                           = 5456
	ChessRogueSelectBpCsReq                            = 5572
	ChessRogueQueryScRsp                               = 5565
	SyncChessRogueMainStoryFinishScNotify              = 5429
	ChessRogueSelectBpScRsp                            = 5465
	ChessRogueQuestFinishNotify                        = 5539
	ChessRogueReRollDiceCsReq                          = 5423
	ChessRogueGiveUpCsReq                              = 5460
	ChessRogueUpdateMoneyInfoScNotify                  = 5501
	ClockParkGetOngoingScriptInfoScRsp                 = 7246
	ClockParkStartScriptCsReq                          = 7209
	ClockParkBattleEndScNotify                         = 7204
	ClockParkUnlockTalentScRsp                         = 7249
	ClockParkUnlockTalentCsReq                         = 7215
	ClockParkQuitScriptScRsp                           = 7248
	ClockParkUseBuffCsReq                              = 7211
	ClockParkHandleWaitOperationScRsp                  = 7228
	ClockParkStartScriptScRsp                          = 7243
	ClockParkGetInfoScRsp                              = 7250
	ClockParkFinishScriptScNotify                      = 7235
	ClockParkGetInfoCsReq                              = 7212
	ClockParkUseBuffScRsp                              = 7206
	ClockParkGetOngoingScriptInfoCsReq                 = 7247
	ClockParkQuitScriptCsReq                           = 7210
	ClockParkHandleWaitOperationCsReq                  = 7245
	ContentPackageGetDataScRsp                         = 7550
	ContentPackageUnlockScRsp                          = 7515
	ContentPackageGetDataCsReq                         = 7512
	ContentPackageSyncDataScNotify                     = 7530
	ContentPackageUnlockCsReq                          = 7523
	ContentPackageTransferScNotify                     = 7549
	TakeApRewardCsReq                                  = 3336
	DailyActiveInfoNotify                              = 3367
	TakeApRewardScRsp                                  = 3395
	TakeAllApRewardScRsp                               = 3352
	TakeAllApRewardCsReq                               = 3328
	GetDailyActiveInfoCsReq                            = 3384
	GetDailyActiveInfoScRsp                            = 3327
	EndDrinkMakerSequenceCsReq                         = 6997
	EndDrinkMakerSequenceScRsp                         = 6986
	DrinkMakerChallengeScRsp                           = 6992
	MakeMissionDrinkCsReq                              = 6985
	DrinkMakerChallengeCsReq                           = 6988
	MakeDrinkCsReq                                     = 6987
	DrinkMakerUpdateTipsNotify                         = 6996
	GetDrinkMakerDataScRsp                             = 6995
	DrinkMakerDayEndScNotify                           = 6982
	GetDrinkMakerDataCsReq                             = 7000
	MakeDrinkScRsp                                     = 6983
	MakeMissionDrinkScRsp                              = 6990
	EnterEraFlipperRegionCsReq                         = 6559
	EraFlipperDataChangeScNotify                       = 6597
	ResetEraFlipperDataCsReq                           = 6565
	ChangeEraFlipperDataScRsp                          = 6573
	ResetEraFlipperDataScRsp                           = 6599
	ChangeEraFlipperDataCsReq                          = 6580
	GetEraFlipperDataCsReq                             = 6562
	GetEraFlipperDataScRsp                             = 6600
	EnterEraFlipperRegionScRsp                         = 6593
	EvolveBuildReRandomStageCsReq                      = 7128
	EvolveBuildStartStageScRsp                         = 7149
	EvolveBuildShopAbilityDownScRsp                    = 7111
	EvolveBuildShopAbilityDownCsReq                    = 7140
	EvolveBuildStartLevelScRsp                         = 7123
	EvolveBuildShopAbilityUpCsReq                      = 7117
	EvolveBuildStartLevelCsReq                         = 7130
	EvolveBuildFinishScNotify                          = 7145
	EvolveBuildQueryInfoCsReq                          = 7112
	EvolveBuildCoinNotify                              = 7114
	EvolveBuildTakeExpRewardScRsp                      = 7135
	EvolveBuildLeaveCsReq                              = 7147
	EvolveBuildStartStageCsReq                         = 7115
	EvolveBuildGiveupCsReq                             = 7109
	EvolveBuildShopAbilityUpScRsp                      = 7104
	EvolveBuildTakeExpRewardCsReq                      = 7106
	EvolveBuildReRandomStageScRsp                      = 7110
	EvolveBuildUnlockInfoNotify                        = 7142
	EvolveBuildShopAbilityResetCsReq                   = 7118
	EvolveBuildGiveupScRsp                             = 7143
	EvolveBuildQueryInfoScRsp                          = 7150
	EvolveBuildLeaveScRsp                              = 7146
	EvolveBuildShopAbilityResetScRsp                   = 7122
	AcceptMultipleExpeditionCsReq                      = 2505
	AcceptActivityExpeditionScRsp                      = 2534
	CancelActivityExpeditionScRsp                      = 2546
	TakeMultipleExpeditionRewardCsReq                  = 2519
	CancelExpeditionScRsp                              = 2528
	ExpeditionDataChangeScNotify                       = 2524
	TakeMultipleActivityExpeditionRewardCsReq          = 2507
	TakeMultipleActivityExpeditionRewardScRsp          = 2553
	AcceptActivityExpeditionCsReq                      = 2593
	GetExpeditionDataScRsp                             = 2595
	AcceptMultipleExpeditionScRsp                      = 2575
	TakeExpeditionRewardScRsp                          = 2574
	TakeActivityExpeditionRewardCsReq                  = 2525
	GetExpeditionDataCsReq                             = 2536
	TakeExpeditionRewardCsReq                          = 2552
	TakeActivityExpeditionRewardScRsp                  = 2596
	CancelActivityExpeditionCsReq                      = 2543
	TakeMultipleExpeditionRewardScRsp                  = 2526
	AcceptExpeditionScRsp                              = 2527
	CancelExpeditionCsReq                              = 2567
	AcceptExpeditionCsReq                              = 2584
	FinishChapterScNotify                              = 4984
	EnterFantasticStoryActivityStageScRsp              = 4967
	FantasticStoryActivityBattleEndScNotify            = 4928
	GetFantasticStoryActivityDataScRsp                 = 4995
	EnterFantasticStoryActivityStageCsReq              = 4927
	GetFantasticStoryActivityDataCsReq                 = 4936
	GetFeverTimeActivityDataScRsp                      = 7151
	FeverTimeActivityBattleEndScNotify                 = 7153
	EnterFeverTimeActivityStageScRsp                   = 7160
	EnterFeverTimeActivityStageCsReq                   = 7159
	GetFeverTimeActivityDataCsReq                      = 7158
	FightKickOutScNotify                               = 30027
	FightEnterScRsp                                    = 30095
	FightEnterCsReq                                    = 30036
	FightSessionStopScNotify                           = 30052
	FightLeaveScNotify                                 = 30084
	FightHeartBeatScRsp                                = 30028
	FightHeartBeatCsReq                                = 30067
	EnterFightActivityStageScRsp                       = 3667
	EnterFightActivityStageCsReq                       = 3627
	FightActivityDataChangeScNotify                    = 3684
	TakeFightActivityRewardCsReq                       = 3628
	TakeFightActivityRewardScRsp                       = 3652
	GetFightActivityDataScRsp                          = 3695
	GetFightActivityDataCsReq                          = 3636
	FightFestUpdateCoinNotify                          = 7293
	StartFightFestCsReq                                = 7280
	GetFightFestDataCsReq                              = 7262
	FightFestScoreUpdateNotify                         = 7265
	StartFightFestScRsp                                = 7273
	FightFestUnlockSkillNotify                         = 7299
	FightFestUpdateChallengeRecordNotify               = 7259
	GetFightFestDataScRsp                              = 7300
	FightMatch3SwapCsReq                               = 30128
	FightMatch3ChatScNotify                            = 30134
	FightMatch3SwapScRsp                               = 30152
	FightMatch3DataScRsp                               = 30195
	FightMatch3TurnEndScNotify                         = 30167
	FightMatch3StartCountDownScNotify                  = 30184
	FightMatch3DataCsReq                               = 30136
	FightMatch3ForceUpdateNotify                       = 30143
	FightMatch3ChatCsReq                               = 30124
	FightMatch3OpponentDataScNotify                    = 30174
	FightMatch3ChatScRsp                               = 30193
	FightMatch3TurnStartScNotify                       = 30127
	GetFriendDevelopmentInfoScRsp                      = 2940
	SetAssistScRsp                                     = 2976
	GetAssistListScRsp                                 = 2958
	DeleteFriendCsReq                                  = 2946
	SetFriendRemarkNameCsReq                           = 2953
	SetFriendMarkCsReq                                 = 2922
	GetFriendRecommendListInfoCsReq                    = 2926
	TakeAssistRewardCsReq                              = 2965
	GetAssistHistoryScRsp                              = 2913
	GetFriendBattleRecordDetailScRsp                   = 2999
	TakeAssistRewardScRsp                              = 2908
	GetFriendChallengeDetailScRsp                      = 2964
	ApplyFriendCsReq                                   = 2952
	GetFriendChallengeLineupCsReq                      = 2992
	GetFriendAssistListScRsp                           = 2994
	GetFriendListInfoCsReq                             = 2936
	SetFriendMarkScRsp                                 = 3000
	GetCurAssistScRsp                                  = 2973
	GetFriendChallengeDetailCsReq                      = 2945
	GetFriendLoginInfoCsReq                            = 2979
	AddBlacklistScRsp                                  = 2975
	SyncApplyFriendScNotify                            = 2924
	HandleFriendCsReq                                  = 2993
	DeleteFriendScRsp                                  = 2925
	GetPlayerDetailInfoCsReq                           = 2984
	SearchPlayerScRsp                                  = 2948
	ReportPlayerCsReq                                  = 2959
	DeleteBlacklistCsReq                               = 2978
	HandleFriendScRsp                                  = 2934
	NewAssistHistoryNotify                             = 2921
	GetFriendDevelopmentInfoCsReq                      = 2903
	SetAssistCsReq                                     = 2910
	SyncAddBlacklistScNotify                           = 2919
	CurAssistChangedNotify                             = 2918
	SyncHandleFriendScNotify                           = 2943
	ApplyFriendScRsp                                   = 2974
	DeleteBlacklistScRsp                               = 2929
	SearchPlayerCsReq                                  = 2950
	GetFriendRecommendListInfoScRsp                    = 2907
	GetAssistHistoryCsReq                              = 2917
	SetForbidOtherApplyFriendScRsp                     = 2955
	GetFriendChallengeLineupScRsp                      = 2960
	GetFriendBattleRecordDetailCsReq                   = 2972
	ReportPlayerScRsp                                  = 2912
	GetPlayerDetailInfoScRsp                           = 2927
	GetFriendAssistListCsReq                           = 2941
	GetPlatformPlayerInfoCsReq                         = 2937
	SyncDeleteFriendScNotify                           = 2996
	GetPlatformPlayerInfoScRsp                         = 2911
	GetAssistListCsReq                                 = 2988
	GetFriendListInfoScRsp                             = 2995
	GetFriendApplyListInfoCsReq                        = 2967
	AddBlacklistCsReq                                  = 2905
	SetFriendRemarkNameScRsp                           = 2990
	GetCurAssistCsReq                                  = 2935
	GetFriendLoginInfoScRsp                            = 2981
	SetForbidOtherApplyFriendCsReq                     = 2939
	GetFriendApplyListInfoScRsp                        = 2928
	GetGachaCeilingCsReq                               = 1967
	ExchangeGachaCeilingCsReq                          = 1952
	GetGachaInfoScRsp                                  = 1995
	GetGachaCeilingScRsp                               = 1928
	DoGachaScRsp                                       = 1927
	ExchangeGachaCeilingScRsp                          = 1974
	DoGachaCsReq                                       = 1984
	GetGachaInfoCsReq                                  = 1936
	FinishEmotionDialoguePerformanceScRsp              = 6374
	SubmitEmotionItemScRsp                             = 6328
	HeartDialScriptChangeScNotify                      = 6324
	HeartDialTraceScriptScRsp                          = 6334
	GetHeartDialInfoScRsp                              = 6395
	ChangeScriptEmotionScRsp                           = 6327
	GetHeartDialInfoCsReq                              = 6336
	HeartDialTraceScriptCsReq                          = 6393
	ChangeScriptEmotionCsReq                           = 6384
	FinishEmotionDialoguePerformanceCsReq              = 6352
	SubmitEmotionItemCsReq                             = 6367
	HeliobusEnterBattleCsReq                           = 5807
	HeliobusInfoChangedScNotify                        = 5843
	HeliobusChallengeUpdateScNotify                    = 5812
	HeliobusActivityDataCsReq                          = 5836
	HeliobusSnsReadScRsp                               = 5827
	HeliobusActivityDataScRsp                          = 5895
	HeliobusSnsUpdateScNotify                          = 5834
	HeliobusSnsReadCsReq                               = 5884
	HeliobusSnsCommentCsReq                            = 5824
	HeliobusSelectSkillScRsp                           = 5875
	HeliobusSnsCommentScRsp                            = 5893
	HeliobusUpgradeLevelCsReq                          = 5846
	HeliobusUnlockSkillScNotify                        = 5896
	HeliobusStartRaidScRsp                             = 5859
	HeliobusSnsLikeScRsp                               = 5874
	HeliobusUpgradeLevelScRsp                          = 5825
	HeliobusSnsPostCsReq                               = 5867
	HeliobusStartRaidCsReq                             = 5890
	HeliobusSnsLikeCsReq                               = 5852
	HeliobusEnterBattleScRsp                           = 5853
	HeliobusSnsPostScRsp                               = 5828
	HeliobusLineupUpdateScNotify                       = 5878
	HeliobusSelectSkillCsReq                           = 5805
	AddEquipmentScNotify                               = 529
	SetTurnFoodSwitchCsReq                             = 508
	GetMarkItemListCsReq                               = 535
	DeleteRelicFilterPlanScRsp                         = 572
	PromoteEquipmentCsReq                              = 584
	RelicReforgeConfirmCsReq                           = 533
	RelicReforgeCsReq                                  = 561
	MarkItemCsReq                                      = 517
	ComposeLimitNumUpdateNotify                        = 558
	GetRecyleTimeCsReq                                 = 550
	LockRelicScRsp                                     = 519
	LockRelicCsReq                                     = 575
	GetBagCsReq                                        = 536
	DestroyItemScRsp                                   = 576
	GetBagScRsp                                        = 595
	LockEquipmentCsReq                                 = 567
	SetTurnFoodSwitchScRsp                             = 518
	ModifyRelicFilterPlanScRsp                         = 545
	DeleteRelicFilterPlanCsReq                         = 564
	ComposeLimitNumCompleteNotify                      = 588
	GetRecyleTimeScRsp                                 = 548
	GetRelicFilterPlanScRsp                            = 541
	GetRelicFilterPlanCsReq                            = 600
	MarkRelicFilterPlanScRsp                           = 503
	ExpUpEquipmentCsReq                                = 534
	DestroyItemCsReq                                   = 510
	DiscardRelicScRsp                                  = 579
	ComposeItemScRsp                                   = 525
	ExpUpRelicScRsp                                    = 505
	GetMarkItemListScRsp                               = 573
	ExpUpEquipmentScRsp                                = 543
	SellItemScRsp                                      = 507
	RelicReforgeConfirmScRsp                           = 597
	ExchangeHcoinCsReq                                 = 590
	AddRelicFilterPlanScRsp                            = 592
	MarkRelicFilterPlanCsReq                           = 599
	ComposeItemCsReq                                   = 546
	RelicReforgeScRsp                                  = 530
	UseItemCsReq                                       = 552
	CancelMarkItemNotify                               = 521
	SyncTurnFoodNotify                                 = 565
	RankUpEquipmentScRsp                               = 593
	ComposeSelectedRelicScRsp                          = 578
	ModifyRelicFilterPlanCsReq                         = 560
	SellItemCsReq                                      = 526
	RelicFilterPlanClearNameScNotify                   = 540
	ComposeSelectedRelicCsReq                          = 512
	RechargeSuccNotify                                 = 553
	DiscardRelicCsReq                                  = 511
	LockEquipmentScRsp                                 = 528
	RankUpEquipmentCsReq                               = 524
	MarkItemScRsp                                      = 513
	ExpUpRelicCsReq                                    = 596
	UseItemScRsp                                       = 574
	AddRelicFilterPlanCsReq                            = 594
	GeneralVirtualItemDataNotify                       = 537
	PromoteEquipmentScRsp                              = 527
	ExchangeHcoinScRsp                                 = 559
	UnlockBackGroundMusicCsReq                         = 3167
	GetJukeboxDataCsReq                                = 3136
	TrialBackGroundMusicCsReq                          = 3152
	TrialBackGroundMusicScRsp                          = 3174
	GetJukeboxDataScRsp                                = 3195
	PlayBackGroundMusicCsReq                           = 3184
	PlayBackGroundMusicScRsp                           = 3127
	UnlockBackGroundMusicScRsp                         = 3128
	QuitLineupCsReq                                    = 752
	GetLineupAvatarDataScRsp                           = 746
	SwapLineupScRsp                                    = 793
	SwitchLineupIndexScRsp                             = 775
	GetLineupAvatarDataCsReq                           = 743
	JoinLineupScRsp                                    = 728
	SetLineupNameScRsp                                 = 726
	GetCurLineupDataCsReq                              = 784
	ChangeLineupLeaderCsReq                            = 725
	ChangeLineupLeaderScRsp                            = 796
	SwitchLineupIndexCsReq                             = 705
	QuitLineupScRsp                                    = 774
	GetStageLineupCsReq                                = 736
	ReplaceLineupScRsp                                 = 712
	GetStageLineupScRsp                                = 795
	SetLineupNameCsReq                                 = 719
	JoinLineupCsReq                                    = 767
	ReplaceLineupCsReq                                 = 759
	SwapLineupCsReq                                    = 724
	ExtraLineupDestroyNotify                           = 778
	VirtualLineupDestroyNotify                         = 790
	GetAllLineupDataCsReq                              = 707
	SyncLineupNotify                                   = 734
	GetAllLineupDataScRsp                              = 753
	GetCurLineupDataScRsp                              = 727
	VirtualLineupTrialAvatarChangeScNotify             = 729
	LobbyBeginCsReq                                    = 7365
	LobbyInviteScRsp                                   = 7360
	LobbySyncInfoScNotify                              = 7395
	LobbyInviteCsReq                                   = 7378
	LobbyKickOutScRsp                                  = 7396
	LobbyQuitCsReq                                     = 7398
	LobbyQuitScRsp                                     = 7363
	LobbyCreateScRsp                                   = 7400
	LobbyKickOutCsReq                                  = 7397
	LobbyModifyPlayerInfoScRsp                         = 7393
	LobbyGetInfoCsReq                                  = 7354
	LobbyGetInfoScRsp                                  = 7390
	LobbyModifyPlayerInfoCsReq                         = 7359
	LobbyBeginScRsp                                    = 7399
	LobbyInviteScNotify                                = 7367
	LobbyJoinCsReq                                     = 7380
	LobbyJoinScRsp                                     = 7373
	LobbyCreateCsReq                                   = 7362
	GetMailScRsp                                       = 895
	TakeMailAttachmentCsReq                            = 852
	DelMailScRsp                                       = 828
	TakeMailAttachmentScRsp                            = 874
	MarkReadMailScRsp                                  = 827
	NewMailScNotify                                    = 824
	GetMailCsReq                                       = 836
	DelMailCsReq                                       = 867
	MarkReadMailCsReq                                  = 884
	InteractChargerCsReq                               = 6884
	ResetMapRotationRegionScRsp                        = 6825
	LeaveMapRotationRegionScRsp                        = 6893
	RotateMapScRsp                                     = 6874
	UpdateEnergyScNotify                               = 6805
	DeployRotaterScRsp                                 = 6828
	EnterMapRotationRegionCsReq                        = 6836
	RemoveRotaterScRsp                                 = 6826
	GetMapRotationDataCsReq                            = 6834
	UpdateRotaterScNotify                              = 6807
	ResetMapRotationRegionCsReq                        = 6846
	LeaveMapRotationRegionScNotify                     = 6896
	EnterMapRotationRegionScRsp                        = 6895
	UpdateMapRotationDataScNotify                      = 6875
	RemoveRotaterCsReq                                 = 6819
	DeployRotaterCsReq                                 = 6867
	RotateMapCsReq                                     = 6852
	GetMapRotationDataScRsp                            = 6843
	LeaveMapRotationRegionCsReq                        = 6824
	InteractChargerScRsp                               = 6827
	UpdateMarkChestCsReq                               = 8187
	MarkChestChangedScNotify                           = 8197
	GetMarkChestCsReq                                  = 8200
	UpdateMarkChestScRsp                               = 8183
	GetMarkChestScRsp                                  = 8195
	CancelMatchCsReq                                   = 7330
	GetCrossInfoCsReq                                  = 7349
	GetCrossInfoScRsp                                  = 7309
	StartMatchCsReq                                    = 7312
	MatchResultScNotify                                = 7315
	StartMatchScRsp                                    = 7350
	CancelMatchScRsp                                   = 7323
	MatchThreeSetBirdPosScRsp                          = 7409
	MatchThreeLevelEndScRsp                            = 7423
	MatchThreeGetDataScRsp                             = 7450
	MatchThreeSyncDataScNotify                         = 7415
	MatchThreeSetBirdPosCsReq                          = 7449
	MatchThreeLevelEndCsReq                            = 7430
	MatchThreeGetDataCsReq                             = 7412
	GetNpcStatusCsReq                                  = 2784
	FinishItemIdScRsp                                  = 2728
	GetNpcMessageGroupScRsp                            = 2795
	GetMissionMessageInfoCsReq                         = 2734
	FinishSectionIdScRsp                               = 2774
	FinishItemIdCsReq                                  = 2767
	GetNpcStatusScRsp                                  = 2727
	FinishPerformSectionIdCsReq                        = 2724
	GetNpcMessageGroupCsReq                            = 2736
	FinishPerformSectionIdScRsp                        = 2793
	FinishSectionIdCsReq                               = 2752
	GetMissionMessageInfoScRsp                         = 2743
	TriggerVoiceCsReq                                  = 4146
	CancelCacheNotifyCsReq                             = 4124
	SecurityReportCsReq                                = 4134
	DifficultyAdjustmentGetDataCsReq                   = 4188
	MazeKillDirectCsReq                                = 4135
	SubmitOrigamiItemCsReq                             = 4196
	MazeKillDirectScRsp                                = 4173
	UpdateGunPlayDataScRsp                             = 4148
	GetGunPlayDataCsReq                                = 4178
	ShareScRsp                                         = 4195
	DifficultyAdjustmentUpdateDataScRsp                = 4176
	GetMovieRacingDataScRsp                            = 4190
	GetMovieRacingDataCsReq                            = 4153
	SecurityReportScRsp                                = 4143
	GetGunPlayDataScRsp                                = 4129
	GetShareDataCsReq                                  = 4184
	TakePictureCsReq                                   = 4167
	TriggerVoiceScRsp                                  = 4125
	UpdateMovieRacingDataScRsp                         = 4112
	SubmitOrigamiItemScRsp                             = 4105
	ShareCsReq                                         = 4136
	TakePictureScRsp                                   = 4128
	UpdateMovieRacingDataCsReq                         = 4159
	GetShareDataScRsp                                  = 4127
	DifficultyAdjustmentUpdateDataCsReq                = 4110
	DifficultyAdjustmentGetDataScRsp                   = 4158
	UpdateGunPlayDataCsReq                             = 4150
	CancelCacheNotifyScRsp                             = 4193
	TeleportToMissionResetPointScRsp                   = 1248
	StartFinishMainMissionScNotify                     = 1258
	AcceptMainMissionCsReq                             = 1210
	AcceptMissionEventCsReq                            = 1219
	InterruptMissionEventScRsp                         = 1259
	FinishedMissionScNotify                            = 1208
	UpdateTrackMainMissionIdScRsp                      = 1265
	InterruptMissionEventCsReq                         = 1290
	SyncTaskCsReq                                      = 1228
	SetMissionEventProgressScRsp                       = 1278
	SubMissionRewardScNotify                           = 1229
	GetMissionStatusCsReq                              = 1207
	SyncTaskScRsp                                      = 1252
	GetMainMissionCustomValueScRsp                     = 1273
	FinishTalkMissionScRsp                             = 1227
	UpdateTrackMainMissionIdCsReq                      = 1221
	MissionRewardScNotify                              = 1267
	GetMainMissionCustomValueCsReq                     = 1235
	GetMissionDataScRsp                                = 1295
	FinishTalkMissionCsReq                             = 1284
	StartFinishSubMissionScNotify                      = 1288
	GetMissionEventDataCsReq                           = 1296
	MissionAcceptScNotify                              = 1217
	MissionGroupWarnScNotify                           = 1243
	MissionEventRewardScNotify                         = 1275
	AcceptMissionEventScRsp                            = 1226
	GetMissionStatusScRsp                              = 1253
	SetMissionEventProgressCsReq                       = 1212
	TeleportToMissionResetPointCsReq                   = 1250
	GetMissionDataCsReq                                = 1236
	GetMissionEventDataScRsp                           = 1205
	FinishCosumeItemMissionScRsp                       = 1225
	FinishCosumeItemMissionCsReq                       = 1246
	AcceptMainMissionScRsp                             = 1276
	MonopolyGuessChooseScRsp                           = 7037
	MonopolyActionResultScNotify                       = 7084
	DailyFirstEnterMonopolyActivityScRsp               = 7025
	GetMonopolyInfoScRsp                               = 7095
	MonopolyEventLoadUpdateScNotify                    = 7001
	GetMonopolyInfoCsReq                               = 7036
	MonopolyBuyGoodsScRsp                              = 7090
	GetMbtiReportCsReq                                 = 7099
	MonopolySttUpdateScNotify                          = 7032
	MonopolyGetDailyInitItemCsReq                      = 7063
	MonopolyGetRegionProgressCsReq                     = 7044
	MonopolyRollDiceCsReq                              = 7028
	MonopolyGiveUpCurContentScRsp                      = 7029
	MonopolyAcceptQuizCsReq                            = 7021
	GetMonopolyFriendRankingListScRsp                  = 7060
	MonopolyQuizDurationChangeScNotify                 = 7039
	MonopolyConfirmRandomScRsp                         = 7007
	MonopolyMoveScRsp                                  = 7024
	MonopolyCheatDiceCsReq                             = 7050
	MonopolyEventSelectFriendCsReq                     = 7040
	MonopolyClickMbtiReportCsReq                       = 7085
	MonopolyUpgradeAssetScRsp                          = 7012
	MonopolyBuyGoodsCsReq                              = 7053
	MonopolyConfirmRandomCsReq                         = 7026
	MonopolyGetRaffleTicketCsReq                       = 7016
	MonopolyConditionUpdateScNotify                    = 7015
	MonopolyGiveUpCurContentCsReq                      = 7078
	DailyFirstEnterMonopolyActivityCsReq               = 7046
	GetSocialEventServerCacheCsReq                     = 7033
	DeleteSocialEventServerCacheScRsp                  = 7068
	DeleteSocialEventServerCacheCsReq                  = 7083
	MonopolyContentUpdateScNotify                      = 7088
	GetMonopolyDailyReportCsReq                        = 7038
	GetMonopolyMbtiReportRewardCsReq                   = 7082
	MonopolyTakePhaseRewardScRsp                       = 7020
	MonopolyUpgradeAssetCsReq                          = 7059
	MonopolyGetRafflePoolInfoScRsp                     = 7056
	MonopolyGetDailyInitItemScRsp                      = 7089
	GetSocialEventServerCacheScRsp                     = 7097
	MonopolyGuessChooseCsReq                           = 7018
	MonopolyGameRaiseRatioCsReq                        = 7058
	MonopolyDailySettleScNotify                        = 7094
	MonopolyScrachRaffleTicketCsReq                    = 7098
	MonopolySelectOptionCsReq                          = 7093
	MonopolyGameGachaCsReq                             = 7035
	MonopolyReRollRandomScRsp                          = 7019
	MonopolyTakePhaseRewardCsReq                       = 7009
	MonopolyCellUpdateNotify                           = 7027
	MonopolyGameBingoFlipCardCsReq                     = 7017
	MonopolyGuessBuyInformationScRsp                   = 7079
	MonopolyRollRandomCsReq                            = 7096
	MonopolyLikeCsReq                                  = 7045
	MonopolyMoveCsReq                                  = 7074
	MonopolyGameSettleScNotify                         = 7076
	MonopolyRollDiceScRsp                              = 7052
	MonopolyAcceptQuizScRsp                            = 7065
	MonopolyGuessDrawScNotify                          = 7081
	MonopolyLikeScRsp                                  = 7064
	MonopolyTakeRaffleTicketRewardCsReq                = 7049
	MonopolyScrachRaffleTicketScRsp                    = 7091
	MonopolyTakeRaffleTicketRewardScRsp                = 7070
	MonopolyReRollRandomCsReq                          = 7075
	MonopolyClickMbtiReportScRsp                       = 7062
	MonopolyClickCellCsReq                             = 7086
	MonopolyGameBingoFlipCardScRsp                     = 7013
	MonopolyCheatDiceScRsp                             = 7048
	MonopolyLikeScNotify                               = 7072
	GetMonopolyFriendRankingListCsReq                  = 7092
	MonopolyEventSelectFriendScRsp                     = 7061
	MonopolyGetRegionProgressScRsp                     = 7023
	MonopolyGameGachaScRsp                             = 7073
	GetMonopolyDailyReportScRsp                        = 7087
	MonopolyGetRafflePoolInfoCsReq                     = 7069
	MonopolyGetRaffleTicketScRsp                       = 7047
	MonopolySelectOptionScRsp                          = 7034
	GetMonopolyMbtiReportRewardScRsp                   = 7080
	MonopolyRollRandomScRsp                            = 7005
	MonopolyGuessBuyInformationCsReq                   = 7011
	MonopolySocialEventEffectScNotify                  = 7030
	MonopolyClickCellScRsp                             = 7066
	MonopolyGameCreateScNotify                         = 7008
	GetMbtiReportScRsp                                 = 7003
	MonopolyGameRaiseRatioScRsp                        = 7010
	MultiplayerMatch3FinishScNotify                    = 1024
	MultiplayerFightGameStateCsReq                     = 1036
	MultiplayerFightGameStateScRsp                     = 1095
	MultiplayerGetFightGateCsReq                       = 1084
	MultiplayerFightGameStartScNotify                  = 1052
	MultiplayerFightGameFinishScNotify                 = 1074
	MultiplayerGetFightGateScRsp                       = 1027
	MultiplayerFightGiveUpCsReq                        = 1067
	MultiplayerFightGiveUpScRsp                        = 1028
	GetMultipleDropInfoCsReq                           = 4636
	GetPlayerReturnMultiDropInfoScRsp                  = 4667
	MultipleDropInfoNotify                             = 4628
	GetPlayerReturnMultiDropInfoCsReq                  = 4627
	MultipleDropInfoScNotify                           = 4684
	GetMultipleDropInfoScRsp                           = 4695
	SetStuffToAreaCsReq                                = 4367
	RemoveStuffFromAreaCsReq                           = 4352
	FinishCurTurnCsReq                                 = 4334
	BuyNpcStuffScRsp                                   = 4327
	MuseumRandomEventQueryScRsp                        = 4353
	MuseumRandomEventSelectScRsp                       = 4359
	MuseumTakeCollectRewardScRsp                       = 4388
	MuseumTargetStartNotify                            = 4378
	GetMuseumInfoCsReq                                 = 4336
	UpgradeAreaStatScRsp                               = 4305
	MuseumInfoChangedScNotify                          = 4375
	GetMuseumInfoScRsp                                 = 4395
	RemoveStuffFromAreaScRsp                           = 4374
	MuseumFundsChangedScNotify                         = 4319
	UpgradeAreaCsReq                                   = 4346
	MuseumDispatchFinishedScNotify                     = 4312
	GetStuffScNotify                                   = 4324
	MuseumRandomEventStartScNotify                     = 4326
	MuseumRandomEventQueryCsReq                        = 4307
	GetExhibitScNotify                                 = 4393
	FinishCurTurnScRsp                                 = 4343
	MuseumRandomEventSelectCsReq                       = 4390
	BuyNpcStuffCsReq                                   = 4384
	UpgradeAreaScRsp                                   = 4325
	UpgradeAreaStatCsReq                               = 4396
	MuseumTargetMissionFinishNotify                    = 4329
	MuseumTargetRewardNotify                           = 4350
	MuseumTakeCollectRewardCsReq                       = 4348
	SetStuffToAreaScRsp                                = 4328
	MusicRhythmStartLevelScRsp                         = 7593
	MusicRhythmFinishLevelCsReq                        = 7591
	MusicRhythmStartLevelCsReq                         = 7576
	MusicRhythmDataCsReq                               = 7573
	MusicRhythmFinishLevelScRsp                        = 7600
	MusicRhythmDataScRsp                               = 7585
	MusicRhythmUnlockSongNotify                        = 7597
	MusicRhythmSaveSongConfigDataCsReq                 = 7596
	MusicRhythmMaxDifficultyLevelsUnlockNotify         = 7580
	MusicRhythmUnlockTrackScNotify                     = 7577
	MusicRhythmSaveSongConfigDataScRsp                 = 7592
	MusicRhythmUnlockSongSfxScNotify                   = 7599
	SubmitOfferingItemScRsp                            = 6923
	TakeOfferingRewardScRsp                            = 6926
	TakeOfferingRewardCsReq                            = 6937
	GetOfferingInfoCsReq                               = 6940
	GetOfferingInfoScRsp                               = 6935
	SubmitOfferingItemCsReq                            = 6927
	OfferingInfoScNotify                               = 6925
	SyncAcceptedPamMissionNotify                       = 4084
	AcceptedPamMissionExpireCsReq                      = 4036
	AcceptedPamMissionExpireScRsp                      = 4095
	GetPamSkinDataScRsp                                = 8135
	SelectPamSkinScRsp                                 = 8123
	SelectPamSkinCsReq                                 = 8127
	UnlockPamSkinScNotify                              = 8137
	GetPamSkinDataCsReq                                = 8140
	SummonPetCsReq                                     = 7616
	RecallPetScRsp                                     = 7607
	GetPetDataCsReq                                    = 7624
	CurPetChangedScNotify                              = 7606
	RecallPetCsReq                                     = 7610
	SummonPetScRsp                                     = 7603
	GetPetDataScRsp                                    = 7623
	SelectPhoneThemeScRsp                              = 5152
	UnlockPhoneThemeScNotify                           = 5174
	SelectPhoneThemeCsReq                              = 5128
	SelectChatBubbleScRsp                              = 5127
	UnlockChatBubbleScNotify                           = 5167
	GetPhoneDataCsReq                                  = 5136
	SelectChatBubbleCsReq                              = 5184
	GetPhoneDataScRsp                                  = 5195
	PlayerLogoutCsReq                                  = 84
	SetAvatarPathCsReq                                 = 87
	GetBasicInfoScRsp                                  = 100
	UpdatePsnSettingsInfoCsReq                         = 32
	SetMultipleAvatarPathsCsReq                        = 62
	SetLanguageScRsp                                   = 88
	ServerAnnounceNotify                               = 58
	PlayerGetTokenScRsp                                = 28
	SetMultipleAvatarPathsScRsp                        = 2
	ReserveStaminaExchangeCsReq                        = 91
	GetVideoVersionKeyScRsp                            = 47
	MonthCardRewardNotify                              = 64
	PlayerHeartBeatScRsp                               = 3
	ClientObjDownloadDataScNotify                      = 9
	GetSecretKeyInfoScRsp                              = 83
	ExchangeStaminaScRsp                               = 96
	UpdatePlayerSettingCsReq                           = 69
	QueryProductInfoCsReq                              = 79
	AntiAddictScNotify                                 = 26
	GmTalkScNotify                                     = 74
	AceAntiCheaterScRsp                                = 45
	FeatureSwitchClosedScNotify                        = 61
	UpdatePlayerSettingScRsp                           = 56
	GetVideoVersionKeyCsReq                            = 16
	SetLanguageCsReq                                   = 48
	GetLevelRewardTakenListScRsp                       = 59
	SetGenderScRsp                                     = 8
	PlayerGetTokenCsReq                                = 67
	UnlockAvatarPathScRsp                              = 85
	GetSecretKeyInfoCsReq                              = 97
	SetNicknameCsReq                                   = 7
	GetLevelRewardCsReq                                = 12
	UpdateFeatureSwitchScNotify                        = 55
	ClientDownloadDataScNotify                         = 39
	GetMultiPathAvatarInfoCsReq                        = 66
	ClientObjUploadScRsp                               = 1
	GetAuthkeyCsReq                                    = 5
	ExchangeStaminaCsReq                               = 25
	PlayerLoginCsReq                                   = 36
	GetAuthkeyScRsp                                    = 75
	ClientObjUploadCsReq                               = 20
	GetGameStateServiceConfigCsReq                     = 6
	ReserveStaminaExchangeScRsp                        = 44
	GmTalkCsReq                                        = 93
	GetLevelRewardTakenListCsReq                       = 90
	RegionStopScNotify                                 = 19
	SetPlayerInfoCsReq                                 = 18
	RetcodeNotify                                      = 72
	SetRedPointStatusScNotify                          = 49
	PlayerLogoutScRsp                                  = 27
	SetNicknameScRsp                                   = 53
	SetAvatarPathScRsp                                 = 86
	SetPlayerInfoScRsp                                 = 37
	PlayerLoginScRsp                                   = 95
	GateServerScNotify                                 = 40
	GetGameStateServiceConfigScRsp                     = 71
	AvatarPathChangedNotify                            = 89
	GetLevelRewardScRsp                                = 78
	AceAntiCheaterCsReq                                = 60
	QueryProductInfoScRsp                              = 81
	PlayerKickOutScNotify                              = 24
	SetGameplayBirthdayCsReq                           = 94
	UnlockAvatarPathCsReq                              = 15
	PlayerLoginFinishCsReq                             = 68
	GetBasicInfoCsReq                                  = 22
	UpdatePsnSettingsInfoScRsp                         = 82
	GmTalkScRsp                                        = 34
	PlayerHeartBeatCsReq                               = 99
	StaminaInfoScNotify                                = 23
	DailyRefreshNotify                                 = 41
	PlayerLoginFinishScRsp                             = 57
	GetMultiPathAvatarInfoScRsp                        = 63
	SetGenderCsReq                                     = 65
	SetGameplayBirthdayScRsp                           = 92
	SetHeadIconScRsp                                   = 2827
	GetPlayerBoardDataCsReq                            = 2836
	SetAssistAvatarCsReq                               = 2843
	GetPlayerBoardDataScRsp                            = 2895
	SetDisplayAvatarScRsp                              = 2828
	SetSignatureScRsp                                  = 2834
	SetHeadIconCsReq                                   = 2884
	UnlockHeadIconScNotify                             = 2824
	SetIsDisplayAvatarInfoCsReq                        = 2852
	SetAssistAvatarScRsp                               = 2846
	SetSignatureCsReq                                  = 2893
	SetDisplayAvatarCsReq                              = 2867
	SetIsDisplayAvatarInfoScRsp                        = 2874
	PlayerReturnTakeRewardCsReq                        = 4552
	PlayerReturnSignCsReq                              = 4595
	PlayerReturnStartScNotify                          = 4536
	PlayerReturnTakeRewardScRsp                        = 4574
	PlayerReturnTakePointRewardScRsp                   = 4528
	PlayerReturnTakePointRewardCsReq                   = 4567
	PlayerReturnInfoQueryCsReq                         = 4524
	PlayerReturnInfoQueryScRsp                         = 4593
	PlayerReturnPointChangeScNotify                    = 4527
	PlayerReturnForceFinishScNotify                    = 4534
	PlayerReturnSignScRsp                              = 4584
	FinishPlotScRsp                                    = 1195
	FinishPlotCsReq                                    = 1136
	StartPunkLordRaidCsReq                             = 3284
	GetKilledPunkLordMonsterDataScRsp                  = 3278
	TakeKilledPunkLordMonsterScoreCsReq                = 3288
	StartPunkLordRaidScRsp                             = 3227
	TakeKilledPunkLordMonsterScoreScRsp                = 3258
	GetKilledPunkLordMonsterDataCsReq                  = 3212
	SharePunkLordMonsterScRsp                          = 3228
	TakePunkLordPointRewardCsReq                       = 3246
	TakePunkLordPointRewardScRsp                       = 3225
	PunkLordBattleResultScNotify                       = 3259
	SummonPunkLordMonsterCsReq                         = 3252
	PunkLordRaidTimeOutScNotify                        = 3226
	SharePunkLordMonsterCsReq                          = 3267
	PunkLordMonsterKilledNotify                        = 3248
	GetPunkLordBattleRecordCsReq                       = 3276
	PunkLordMonsterInfoScNotify                        = 3296
	GetPunkLordMonsterDataScRsp                        = 3295
	SummonPunkLordMonsterScRsp                         = 3274
	GetPunkLordMonsterDataCsReq                        = 3236
	GetPunkLordBattleRecordScRsp                       = 3235
	GetPunkLordDataCsReq                               = 3205
	GetPunkLordDataScRsp                               = 3275
	PunkLordDataChangeNotify                           = 3210
	TakeQuestOptionalRewardScRsp                       = 946
	QuestRecordScNotify                                = 924
	BatchGetQuestDataScRsp                             = 905
	FinishQuestScRsp                                   = 934
	BatchGetQuestDataCsReq                             = 996
	GetQuestDataCsReq                                  = 936
	GetQuestRecordScRsp                                = 974
	GetQuestRecordCsReq                                = 952
	TakeQuestOptionalRewardCsReq                       = 943
	TakeQuestRewardCsReq                               = 984
	GetQuestDataScRsp                                  = 995
	FinishQuestCsReq                                   = 993
	TakeQuestRewardScRsp                               = 927
	GetRaidInfoCsReq                                   = 2234
	GetSaveRaidCsReq                                   = 2296
	GetChallengeRaidInfoScRsp                          = 2252
	GetAllSaveRaidScRsp                                = 2219
	RaidInfoNotify                                     = 2267
	DelSaveRaidScNotify                                = 2226
	SetClientRaidTargetCountScRsp                      = 2225
	SetClientRaidTargetCountCsReq                      = 2246
	LeaveRaidCsReq                                     = 2284
	StartRaidScRsp                                     = 2295
	GetRaidInfoScRsp                                   = 2243
	LeaveRaidScRsp                                     = 2227
	TakeChallengeRaidRewardScRsp                       = 2224
	RaidKickByServerScNotify                           = 2207
	GetAllSaveRaidCsReq                                = 2275
	StartRaidCsReq                                     = 2236
	GetChallengeRaidInfoCsReq                          = 2228
	TakeChallengeRaidRewardCsReq                       = 2274
	GetSaveRaidScRsp                                   = 2205
	ChallengeRaidNotify                                = 2293
	RaidCollectionDataScRsp                            = 6955
	RaidCollectionDataScNotify                         = 6947
	RaidCollectionEnterNextRaidCsReq                   = 6943
	RaidCollectionEnterNextRaidScRsp                   = 6957
	RaidCollectionDataCsReq                            = 6960
	GetBigDataRecommendScRsp                           = 2443
	GetChallengeRecommendLineupListScRsp               = 2449
	RelicAvatarRecommendScRsp                          = 2423
	RelicRecommendScRsp                                = 2450
	GetChallengeRecommendLineupListCsReq               = 2415
	RelicAvatarRecommendCsReq                          = 2430
	RelicRecommendCsReq                                = 2412
	GetBigDataRecommendCsReq                           = 2409
	UpdateRedDotDataCsReq                              = 5984
	GetSingleRedDotParamGroupScRsp                     = 5928
	GetAllRedDotDataCsReq                              = 5936
	GetAllRedDotDataScRsp                              = 5995
	UpdateRedDotDataScRsp                              = 5927
	GetSingleRedDotParamGroupCsReq                     = 5967
	RelicSmartWearGetPlanCsReq                         = 8270
	RelicSmartWearDeletePlanCsReq                      = 8255
	RelicSmartWearGetPlanScRsp                         = 8265
	RelicSmartWearAddPlanScRsp                         = 8253
	RelicSmartWearUpdatePlanScRsp                      = 8256
	RelicSmartWearUpdatePlanCsReq                      = 8267
	RelicSmartWearAddPlanCsReq                         = 8257
	RelicSmartWearDeletePlanScRsp                      = 8260
	GetPlayerReplayInfoScRsp                           = 3527
	GetReplayTokenCsReq                                = 3536
	GetReplayTokenScRsp                                = 3595
	GetPlayerReplayInfoCsReq                           = 3584
	DailyFirstMeetPamScRsp                             = 3427
	GetRndOptionCsReq                                  = 3436
	DailyFirstMeetPamCsReq                             = 3484
	GetRndOptionScRsp                                  = 3495
	SyncRogueStatusScNotify                            = 1854
	GetRogueScoreRewardInfoScRsp                       = 1820
	GetRogueScoreRewardInfoCsReq                       = 1809
	ReviveRogueAvatarCsReq                             = 1826
	StartRogueScRsp                                    = 1827
	SyncRogueGetItemScNotify                           = 1870
	QuitRogueCsReq                                     = 1876
	TakeRogueScoreRewardScRsp                          = 1890
	FinishAeonDialogueGroupScRsp                       = 1889
	ExchangeRogueRewardKeyCsReq                        = 1899
	GetRogueInfoCsReq                                  = 1836
	TakeRogueAeonLevelRewardCsReq                      = 1898
	LeaveRogueScRsp                                    = 1874
	GetRogueTalentInfoScRsp                            = 1885
	EnableRogueTalentScRsp                             = 1802
	SyncRogueFinishScNotify                            = 1896
	SyncRogueAeonScNotify                              = 1847
	SyncRogueExploreWinScNotify                        = 1817
	GetRogueInitialScoreCsReq                          = 1837
	QuitRogueScRsp                                     = 1835
	GetRogueBuffEnhanceInfoScRsp                       = 1812
	OpenRogueChestScRsp                                = 1872
	EnhanceRogueBuffCsReq                              = 1878
	EnterRogueMapRoomScRsp                             = 1818
	SyncRogueAeonLevelUpRewardScNotify                 = 1856
	ExchangeRogueRewardKeyScRsp                        = 1803
	LeaveRogueCsReq                                    = 1852
	PickRogueAvatarCsReq                               = 1805
	TakeRogueAeonLevelRewardScRsp                      = 1891
	StartRogueCsReq                                    = 1884
	TakeRogueScoreRewardCsReq                          = 1853
	EnhanceRogueBuffScRsp                              = 1829
	EnterRogueCsReq                                    = 1867
	EnterRogueScRsp                                    = 1828
	SyncRogueReviveInfoScNotify                        = 1810
	GetRogueAeonInfoScRsp                              = 1866
	SyncRogueSeasonFinishScNotify                      = 1813
	OpenRogueChestCsReq                                = 1864
	GetRogueAeonInfoCsReq                              = 1886
	SyncRogueAreaUnlockScNotify                        = 1849
	GetRogueInfoScRsp                                  = 1895
	EnableRogueTalentCsReq                             = 1862
	GetRogueBuffEnhanceInfoCsReq                       = 1859
	SyncRogueVirtualItemInfoScNotify                   = 1877
	SyncRogueRewardInfoScNotify                        = 1814
	FinishAeonDialogueGroupCsReq                       = 1863
	SyncRogueMapRoomScNotify                           = 1879
	SyncRoguePickAvatarInfoScNotify                    = 1851
	EnterRogueMapRoomCsReq                             = 1808
	PickRogueAvatarScRsp                               = 1875
	ReviveRogueAvatarScRsp                             = 1807
	GetRogueTalentInfoCsReq                            = 1815
	GetRogueInitialScoreScRsp                          = 1811
	RogueArcadeRestartScRsp                            = 7699
	RogueArcadeLeaveScRsp                              = 7673
	RogueArcadeStartScRsp                              = 7700
	RogueArcadeGetInfoCsReq                            = 7659
	RogueArcadeLeaveCsReq                              = 7680
	RogueArcadeStartCsReq                              = 7662
	RogueArcadeGetInfoScRsp                            = 7693
	RogueArcadeRestartCsReq                            = 7665
	CommonRogueQueryCsReq                              = 5664
	TakeRogueEventHandbookRewardScRsp                  = 5679
	BuyRogueShopMiracleScRsp                           = 5624
	SyncRogueCommonPendingActionScNotify               = 5639
	RogueNpcDisappearScRsp                             = 5646
	TakeRogueMiracleHandbookRewardCsReq                = 5618
	UpdateRogueAdventureRoomScoreCsReq                 = 5655
	PrepareRogueAdventureRoomScRsp                     = 5684
	GetRogueExhibitionScRsp                            = 5656
	CommonRogueComponentUpdateScNotify                 = 5657
	SyncRogueCommonVirtualItemInfoScNotify             = 5700
	SyncRogueCommonDialogueDataScNotify                = 5683
	GetRogueCommonDialogueDataCsReq                    = 5603
	RogueWorkbenchHandleFuncScRsp                      = 5670
	EnhanceCommonRogueBuffCsReq                        = 5659
	SelectRogueCommonDialogueOptionCsReq               = 5661
	RogueWorkbenchGetInfoCsReq                         = 5616
	RogueDebugMessageScNotify                          = 5615
	BuyRogueShopBuffCsReq                              = 5693
	UpdateRogueAdventureRoomScoreScRsp                 = 5622
	SetRogueExhibitionCsReq                            = 5609
	RogueGetGambleInfoScRsp                            = 5687
	GetRogueHandbookDataCsReq                          = 5621
	RogueWorkbenchGetInfoScRsp                         = 5647
	GetRogueAdventureRoomInfoCsReq                     = 5625
	GetRogueShopMiracleInfoScRsp                       = 5667
	RogueWorkbenchSelectFuncCsReq                      = 5663
	RogueWorkbenchHandleFuncCsReq                      = 5649
	HandleRogueCommonPendingActionScRsp                = 5645
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5653
	SelectRogueCommonDialogueOptionScRsp               = 5630
	HandleRogueCommonPendingActionCsReq                = 5660
	ExchangeRogueBuffWithMiracleScRsp                  = 5607
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5690
	GetRogueHandbookDataScRsp                          = 5665
	SyncRogueCommonDialogueOptionFinishScNotify        = 5668
	RogueGetGambleInfoCsReq                            = 5638
	GetRogueExhibitionCsReq                            = 5669
	FinishRogueCommonDialogueScRsp                     = 5697
	SyncRogueHandbookDataUpdateScNotify                = 5608
	CommonRogueQueryScRsp                              = 5672
	RogueNpcDisappearCsReq                             = 5643
	RogueDoGambleCsReq                                 = 5686
	GetRogueAdventureRoomInfoScRsp                     = 5696
	BuyRogueShopBuffScRsp                              = 5634
	SetRogueCollectionCsReq                            = 5644
	SyncRogueAdventureRoomInfoScNotify                 = 5636
	RogueDoGambleScRsp                                 = 5666
	SetRogueExhibitionScRsp                            = 5620
	PrepareRogueAdventureRoomCsReq                     = 5695
	StopRogueAdventureRoomScRsp                        = 5629
	StopRogueAdventureRoomCsReq                        = 5678
	EnhanceCommonRogueBuffScRsp                        = 5612
	BuyRogueShopMiracleCsReq                           = 5674
	GetRogueCollectionCsReq                            = 5698
	GetRogueShopBuffInfoScRsp                          = 5652
	GetRogueCollectionScRsp                            = 5691
	GetRogueCommonDialogueDataScRsp                    = 5640
	RogueWorkbenchSelectFuncScRsp                      = 5689
	SyncRogueCommonActionResultScNotify                = 5681
	SetRogueCollectionScRsp                            = 5623
	GetRogueShopBuffInfoCsReq                          = 5628
	TakeRogueMiracleHandbookRewardScRsp                = 5637
	FinishRogueCommonDialogueCsReq                     = 5633
	TakeRogueEventHandbookRewardCsReq                  = 5611
	ExchangeRogueBuffWithMiracleCsReq                  = 5626
	GetRogueShopMiracleInfoCsReq                       = 5627
	CommonRogueUpdateScNotify                          = 5699
	TakeRogueEndlessActivityPointRewardScRsp           = 6004
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6007
	RogueEndlessActivityBattleEndScNotify              = 6001
	GetRogueEndlessActivityDataCsReq                   = 6009
	EnterRogueEndlessActivityStageScRsp                = 6008
	GetRogueEndlessActivityDataScRsp                   = 6010
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6002
	EnterRogueEndlessActivityStageCsReq                = 6003
	TakeRogueEndlessActivityPointRewardCsReq           = 6005
	RogueMagicUnitReforgeCsReq                         = 7710
	RogueMagicLeaveCsReq                               = 7767
	RogueMagicSetAutoDressInMagicUnitScRsp             = 7765
	RogueMagicUnitReforgeScRsp                         = 7776
	RogueMagicEnterRoomScRsp                           = 7793
	RogueMagicLevelInfoUpdateScNotify                  = 7746
	RogueMagicUnitComposeCsReq                         = 7788
	RogueMagicReviveAvatarCsReq                        = 7707
	RogueMagicSetAutoDressInMagicUnitCsReq             = 7721
	RogueMagicQueryScRsp                               = 7759
	RogueMagicEnableTalentCsReq                        = 7717
	RogueMagicEnterCsReq                               = 7784
	RogueMagicGetTalentInfoCsReq                       = 7735
	RogueMagicStartScRsp                               = 7795
	RogueMagicGetTalentInfoScRsp                       = 7773
	RogueMagicScepterTakeOffUnitCsReq                  = 7750
	RogueMagicScepterDressInUnitCsReq                  = 7778
	RogueMagicScepterDressInUnitScRsp                  = 7729
	RogueMagicAreaUpdateScNotify                       = 7725
	RogueMagicAutoDressInMagicUnitChangeScNotify       = 7781
	RogueMagicSettleCsReq                              = 7752
	RogueMagicScepterTakeOffUnitScRsp                  = 7748
	RogueMagicQueryCsReq                               = 7790
	RogueMagicEnterLayerCsReq                          = 7734
	RogueMagicEnterLayerScRsp                          = 7743
	RogueMagicEnterRoomCsReq                           = 7724
	RogueMagicAutoDressInUnitCsReq                     = 7737
	RogueMagicEnterScRsp                               = 7727
	RogueMagicUnitComposeScRsp                         = 7758
	RogueMagicEnableTalentScRsp                        = 7713
	RogueMagicStartCsReq                               = 7736
	RogueMagicGetMiscRealTimeDataCsReq                 = 7708
	RogueMagicReviveCostUpdateScNotify                 = 7726
	RogueMagicStoryInfoUpdateScNotify                  = 7779
	RogueMagicLeaveScRsp                               = 7728
	RogueMagicBattleFailSettleInfoScNotify             = 7719
	RogueMagicSettleScRsp                              = 7774
	RogueMagicGetMiscRealTimeDataScRsp                 = 7718
	RogueMagicAutoDressInUnitScRsp                     = 7711
	RogueMagicReviveAvatarScRsp                        = 7753
	RogueModifierAddNotify                             = 5384
	RogueModifierUpdateNotify                          = 5374
	RogueModifierSelectCellScRsp                       = 5367
	RogueModifierStageStartNotify                      = 5393
	RogueModifierDelNotify                             = 5324
	RogueModifierSelectCellCsReq                       = 5327
	RogueTournDeleteArchiveScRsp                       = 6077
	RogueTournEnterRoomCsReq                           = 6044
	RogueTournStartScRsp                               = 6027
	RogueTournEnablePermanentTalentScRsp               = 6045
	RogueTournGetArchiveRepositoryCsReq                = 6043
	RogueTournReEnterRogueCocoonStageCsReq             = 6082
	RogueTournResetPermanentTalentCsReq                = 6020
	RogueTournGetSettleInfoScRsp                       = 6055
	RogueTournClearArchiveNameScNotify                 = 6068
	RogueTournGetAllArchiveScRsp                       = 6022
	RogueTournLeaveRogueCocoonSceneCsReq               = 6028
	RogueTournEnterRogueCocoonSceneCsReq               = 6057
	RogueTournQueryCsReq                               = 6053
	RogueTournDifficultyCompNotify                     = 6086
	RogueTournTakeExpRewardCsReq                       = 6023
	RogueTournResetPermanentTalentScRsp                = 6015
	RogueTournGetArchiveRepositoryScRsp                = 6026
	RogueTournEnterScRsp                               = 6069
	RogueTournTakeExpRewardScRsp                       = 6078
	RogueTournGetCurRogueCocoonInfoScRsp               = 6092
	RogueTournGetSettleInfoCsReq                       = 6060
	RogueTournHandBookNotify                           = 6032
	RogueTournLeaveScRsp                               = 6097
	RogueTournWeekChallengeUpdateScNotify              = 6050
	RogueTournReEnterRogueCocoonStageScRsp             = 6083
	RogueTournBattleFailSettleInfoScNotify             = 6018
	RogueTournGetMiscRealTimeDataCsReq                 = 6099
	RogueTournGetAllArchiveCsReq                       = 6037
	RogueTournEnterRoomScRsp                           = 6036
	RogueTournGetMiscRealTimeDataScRsp                 = 6040
	RogueTournSettleScRsp                              = 6074
	RogueTournEnterCsReq                               = 6100
	RogueTournLevelInfoUpdateScNotify                  = 6029
	RogueTournEnablePermanentTalentCsReq               = 6089
	RogueTournSettleCsReq                              = 6075
	RogueTournGetCurRogueCocoonInfoCsReq               = 6094
	RogueTournGetPermanentTalentInfoCsReq              = 6058
	RogueTournGetPermanentTalentInfoScRsp              = 6046
	RogueTournExpNotify                                = 6051
	RogueTournReviveAvatarScRsp                        = 6079
	RogueTournConfirmSettleScRsp                       = 6013
	RogueTournConfirmSettleCsReq                       = 6085
	RogueTournDeleteArchiveCsReq                       = 6070
	RogueTournLeaveRogueCocoonSceneScRsp               = 6021
	RogueTournEnterRogueCocoonSceneScRsp               = 6065
	RogueTournQueryScRsp                               = 6093
	RogueTournLeaveCsReq                               = 6047
	RogueTournRenameArchiveScRsp                       = 6031
	RogueTournEnterLayerScRsp                          = 6038
	RogueTournStartCsReq                               = 6059
	RogueTournReviveAvatarCsReq                        = 6063
	RogueTournRenameArchiveCsReq                       = 6090
	RogueTournAreaUpdateScNotify                       = 6080
	RogueTournReviveCostUpdateScNotify                 = 6039
	RogueTournEnterLayerCsReq                          = 6095
	TakeRollShopRewardCsReq                            = 6917
	DoGachaInRollShopScRsp                             = 6903
	DoGachaInRollShopCsReq                             = 6907
	TakeRollShopRewardScRsp                            = 6906
	GetRollShopInfoScRsp                               = 6915
	GetRollShopInfoCsReq                               = 6920
	EnterSectionScRsp                                  = 1448
	GameplayCounterRecoverCsReq                        = 1482
	StartCocoonStageCsReq                              = 1413
	GameplayCounterCountDownCsReq                      = 1409
	SceneEntityTeleportCsReq                           = 1483
	RefreshTriggerByClientCsReq                        = 1415
	RefreshTriggerByClientScNotify                     = 1462
	SceneEntityMoveCsReq                               = 1436
	GetCurSceneInfoScRsp                               = 1474
	ChangePropTimelineInfoCsReq                        = 1424
	GameplayCounterCountDownScRsp                      = 1420
	InteractPropCsReq                                  = 1484
	GetUnlockTeleportScRsp                             = 1423
	GetSceneMapInfoScRsp                               = 1498
	RecoverAllLineupScRsp                              = 1473
	UnlockTeleportNotify                               = 1414
	SceneCastSkillCostMpCsReq                          = 1425
	SceneCastSkillMpUpdateScNotify                     = 1405
	ScenePlaneEventScNotify                            = 1449
	SetCurInteractEntityCsReq                          = 1410
	SceneEntityMoveScRsp                               = 1495
	SetClientPausedCsReq                               = 1418
	SceneCastSkillScRsp                                = 1428
	SceneGroupRefreshScNotify                          = 1432
	ActivateFarmElementScRsp                           = 1455
	SceneCastSkillCsReq                                = 1467
	DeleteSummonUnitScRsp                              = 1406
	SceneEntityMoveScNotify                            = 1434
	UpdateFloorSavedValueNotify                        = 1456
	UnlockedAreaMapScNotify                            = 1471
	GameplayCounterUpdateScNotify                      = 1401
	ReEnterLastElementStageScRsp                       = 1497
	SavePointsInfoNotify                               = 1417
	EnterSectionCsReq                                  = 1450
	SceneEntityTeleportScRsp                           = 1468
	SceneEnterStageScRsp                               = 1412
	SetGroupCustomSaveDataScRsp                        = 1440
	UpdateMechanismBarScNotify                         = 1499
	ReturnLastTownScRsp                                = 1490
	SetClientPausedScRsp                               = 1437
	ReturnLastTownCsReq                                = 1453
	GroupStateChangeCsReq                              = 1438
	ReEnterLastElementStageCsReq                       = 1433
	StartCocoonStageScRsp                              = 1421
	EntityBindPropScRsp                                = 1408
	EnterSceneScRsp                                    = 1416
	GetSceneMapInfoCsReq                               = 1470
	InteractPropScRsp                                  = 1427
	SyncEntityBuffChangeListScNotify                   = 1446
	UpdateGroupPropertyScRsp                           = 1442
	SpringRefreshCsReq                                 = 1419
	SceneReviveAfterRebattleCsReq                      = 1478
	SetCurInteractEntityScRsp                          = 1476
	SetGroupCustomSaveDataCsReq                        = 1403
	UpdateGroupPropertyCsReq                           = 1451
	SyncServerSceneChangeNotify                        = 1491
	RecoverAllLineupCsReq                              = 1435
	GroupStateChangeScNotify                           = 1486
	SceneReviveAfterRebattleScRsp                      = 1429
	DeleteSummonUnitCsReq                              = 1402
	EnterSceneByServerScNotify                         = 1447
	GroupStateChangeScRsp                              = 1487
	SceneCastSkillCostMpScRsp                          = 1496
	SpringRefreshScRsp                                 = 1426
	SceneEnterStageCsReq                               = 1459
	OpenChestScNotify                                  = 1469
	GetCurSceneInfoCsReq                               = 1452
	DeactivateFarmElementCsReq                         = 1479
	RefreshTriggerByClientScRsp                        = 1485
	GetEnteredSceneCsReq                               = 1466
	DeactivateFarmElementScRsp                         = 1481
	ActivateFarmElementCsReq                           = 1439
	EnterSceneCsReq                                    = 1457
	GetUnlockTeleportCsReq                             = 1444
	ChangePropTimelineInfoScRsp                        = 1493
	EnteredSceneChangeScNotify                         = 1489
	LastSpringRefreshTimeNotify                        = 1407
	GameplayCounterRecoverScRsp                        = 1480
	EntityBindPropCsReq                                = 1465
	SceneUpdatePositionVersionNotify                   = 1443
	GetEnteredSceneScRsp                               = 1463
	GetAllServerPrefsDataCsReq                         = 6136
	GetAllServerPrefsDataScRsp                         = 6195
	GetServerPrefsDataScRsp                            = 6127
	UpdateServerPrefsDataCsReq                         = 6167
	UpdateServerPrefsDataScRsp                         = 6128
	GetServerPrefsDataCsReq                            = 6184
	GetShopListCsReq                                   = 1536
	TakeCityShopRewardCsReq                            = 1567
	TakeCityShopRewardScRsp                            = 1528
	CityShopInfoScNotify                               = 1552
	BuyGoodsCsReq                                      = 1584
	BuyGoodsScRsp                                      = 1527
	GetShopListScRsp                                   = 1595
	SpaceZooExchangeItemCsReq                          = 6743
	SpaceZooDeleteCatCsReq                             = 6724
	SpaceZooExchangeItemScRsp                          = 6746
	SpaceZooOpCatteryScRsp                             = 6774
	SpaceZooOpCatteryCsReq                             = 6752
	SpaceZooMutateScRsp                                = 6728
	SpaceZooDeleteCatScRsp                             = 6793
	SpaceZooCatUpdateNotify                            = 6734
	SpaceZooBornScRsp                                  = 6727
	SpaceZooTakeCsReq                                  = 6725
	SpaceZooDataScRsp                                  = 6795
	SpaceZooBornCsReq                                  = 6784
	SpaceZooDataCsReq                                  = 6736
	SpaceZooMutateCsReq                                = 6767
	SpaceZooTakeScRsp                                  = 6796
	GetStarFightDataCsReq                              = 7169
	StarFightDataChangeNotify                          = 7161
	StartStarFightLevelCsReq                           = 7163
	GetStarFightDataScRsp                              = 7170
	StartStarFightLevelScRsp                           = 7168
	GetStoryLineInfoCsReq                              = 6236
	StoryLineTrialAvatarChangeScNotify                 = 6252
	ChangeStoryLineFinishScNotify                      = 6228
	StoryLineInfoScNotify                              = 6284
	GetStoryLineInfoScRsp                              = 6295
	GetStrongChallengeActivityDataCsReq                = 6636
	EnterStrongChallengeActivityStageScRsp             = 6627
	StrongChallengeActivityBattleEndScNotify           = 6667
	GetStrongChallengeActivityDataScRsp                = 6695
	EnterStrongChallengeActivityStageCsReq             = 6684
	GetSummonActivityDataCsReq                         = 7569
	EnterSummonActivityStageCsReq                      = 7563
	SummonActivityBattleEndScNotify                    = 7561
	GetSummonActivityDataScRsp                         = 7570
	EnterSummonActivityStageScRsp                      = 7568
	SwitchHandDataCsReq                                = 8120
	SwitchHandFinishScRsp                              = 8106
	SwitchHandCoinUpdateCsReq                          = 8102
	SwitchHandDataScRsp                                = 8115
	SwitchHandFinishCsReq                              = 8117
	SwitchHandUpdateScRsp                              = 8110
	SwitchHandStartScRsp                               = 8103
	SwitchHandStartCsReq                               = 8107
	SwitchHandUpdateCsReq                              = 8105
	SwitchHandCoinUpdateScRsp                          = 8108
	SwordTrainingGiveUpGameScRsp                       = 7472
	EnterSwordTrainingExamScRsp                        = 7463
	SwordTrainingGiveUpGameCsReq                       = 7468
	SwordTrainingResumeGameScRsp                       = 7486
	SwordTrainingSelectEndingCsReq                     = 7479
	SwordTrainingGameSyncChangeScNotify                = 7462
	SwordTrainingDialogueSelectOptionScRsp             = 7497
	SwordTrainingStoryBattleScRsp                      = 7451
	SwordTrainingStoryConfirmCsReq                     = 7456
	SwordTrainingGameSettleScNotify                    = 7492
	SwordTrainingDailyPhaseConfirmScRsp                = 7459
	SwordTrainingDailyPhaseConfirmCsReq                = 7499
	SwordTrainingExamResultConfirmCsReq                = 7488
	SwordTrainingTurnActionCsReq                       = 7473
	SwordTrainingMarkEndingViewedCsReq                 = 7457
	SwordTrainingStoryBattleCsReq                      = 7481
	SwordTrainingStartGameScRsp                        = 7461
	SwordTrainingRestoreGameCsReq                      = 7484
	SwordTrainingExamResultConfirmScRsp                = 7483
	SwordTrainingSetSkillTraceCsReq                    = 7489
	EnterSwordTrainingExamCsReq                        = 7498
	SwordTrainingActionTurnSettleScNotify              = 7487
	GetSwordTrainingDataCsReq                          = 7500
	GetSwordTrainingDataScRsp                          = 7480
	SwordTrainingLearnSkillCsReq                       = 7467
	SwordTrainingTurnActionScRsp                       = 7465
	SwordTrainingLearnSkillScRsp                       = 7454
	SwordTrainingStartGameCsReq                        = 7490
	SwordTrainingDialogueSelectOptionCsReq             = 7493
	SwordTrainingSelectEndingScRsp                     = 7475
	SwordTrainingUnlockSyncScNotify                    = 7455
	SwordTrainingStoryConfirmScRsp                     = 7485
	SwordTrainingMarkEndingViewedScRsp                 = 7476
	SwordTrainingSetSkillTraceScRsp                    = 7477
	SwordTrainingResumeGameCsReq                       = 7453
	SwordTrainingRestoreGameScRsp                      = 7469
	PlayerSyncScNotify                                 = 636
	TakeTalkRewardScRsp                                = 2127
	FinishFirstTalkNpcScRsp                            = 2174
	GetNpcTakenRewardScRsp                             = 2195
	SelectInclinationTextScRsp                         = 2193
	TakeTalkRewardCsReq                                = 2184
	GetFirstTalkNpcScRsp                               = 2128
	GetFirstTalkByPerformanceNpcScRsp                  = 2143
	GetFirstTalkByPerformanceNpcCsReq                  = 2134
	FinishFirstTalkByPerformanceNpcScRsp               = 2125
	GetFirstTalkNpcCsReq                               = 2167
	SelectInclinationTextCsReq                         = 2124
	FinishFirstTalkNpcCsReq                            = 2152
	FinishFirstTalkByPerformanceNpcCsReq               = 2146
	GetNpcTakenRewardCsReq                             = 2136
	TarotBookModifyEnergyScNotify                      = 8142
	TarotBookFinishInteractionCsReq                    = 8148
	TarotBookFinishInteractionScRsp                    = 8152
	TarotBookGetDataScRsp                              = 8155
	TarotBookOpenPackScRsp                             = 8143
	TarotBookGetDataCsReq                              = 8160
	TarotBookUnlockStoryCsReq                          = 8157
	TarotBookFinishStoryCsReq                          = 8145
	TarotBookOpenPackCsReq                             = 8147
	TarotBookUnlockStoryScRsp                          = 8146
	TarotBookFinishStoryScRsp                          = 8150
	TelevisionActivityBattleEndScNotify                = 6966
	GetTelevisionActivityDataCsReq                     = 6980
	GetTelevisionActivityDataScRsp                     = 6975
	TelevisionActivityDataChangeScNotify               = 6967
	EnterTelevisionActivityStageScRsp                  = 6977
	EnterTelevisionActivityStageCsReq                  = 6963
	TextJoinBatchSaveCsReq                             = 3867
	TextJoinQueryScRsp                                 = 3827
	TextJoinSaveCsReq                                  = 3836
	TextJoinSaveScRsp                                  = 3895
	TextJoinBatchSaveScRsp                             = 3828
	TextJoinQueryCsReq                                 = 3884
	GetTrackPhotoActivityDataCsReq                     = 7559
	GetTrackPhotoActivityDataScRsp                     = 7560
	QuitTrackPhotoStageScRsp                           = 7552
	StartTrackPhotoStageCsReq                          = 7551
	SettleTrackPhotoStageCsReq                         = 7553
	SettleTrackPhotoStageScRsp                         = 7558
	QuitTrackPhotoStageCsReq                           = 7554
	StartTrackPhotoStageScRsp                          = 7555
	TrainPartyBuildStartStepScRsp                      = 8025
	TrainPartyGamePlayStartCsReq                       = 8050
	TrainPartyAddBuildDynamicBuffScRsp                 = 8013
	TrainPartyTakeBuildLevelAwardScRsp                 = 8073
	TrainPartyBuildStartStepCsReq                      = 8046
	TrainPartyUpdatePosEnvCsReq                        = 8010
	TrainPartyGetDataCsReq                             = 8036
	TrainPartyBuildingUpdateNotify                     = 8075
	TrainPartyBuildDiyScRsp                            = 8005
	TrainPartyEnterCsReq                               = 8026
	TrainPartyUseCardCsReq                             = 8084
	TrainPartyMoveScNotify                             = 8067
	TrainPartySettleNotify                             = 8052
	TrainPartyEnterScRsp                               = 8007
	TrainPartyTakeBuildLevelAwardCsReq                 = 8035
	TrainPartySyncUpdateScNotify                       = 8024
	TrainPartyLeaveCsReq                               = 8053
	TrainPartyGamePlayStartScRsp                       = 8048
	TrainPartyBuildDiyCsReq                            = 8096
	TrainPartyUseCardScRsp                             = 8027
	TrainPartyHandlePendingActionScRsp                 = 8043
	TrainPartyGamePlaySettleNotify                     = 8088
	TrainPartyUpdatePosEnvScRsp                        = 8076
	TrainPartyHandlePendingActionCsReq                 = 8034
	TrainPartyAddBuildDynamicBuffCsReq                 = 8017
	TrainPartyLeaveScRsp                               = 8090
	TrainPartyGetDataScRsp                             = 8095
	TrainVisitorRewardSendNotify                       = 3728
	TrainVisitorBehaviorFinishCsReq                    = 3736
	ShowNewSupplementVisitorCsReq                      = 3734
	GetTrainVisitorRegisterScRsp                       = 3774
	TrainVisitorBehaviorFinishScRsp                    = 3795
	GetTrainVisitorRegisterCsReq                       = 3752
	TrainRefreshTimeNotify                             = 3767
	GetTrainVisitorBehaviorCsReq                       = 3784
	GetTrainVisitorBehaviorScRsp                       = 3727
	ShowNewSupplementVisitorScRsp                      = 3743
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3724
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3793
	TravelBrochureSetPageDescStatusCsReq               = 6475
	TravelBrochureSetCustomValueCsReq                  = 6496
	TravelBrochureGetDataCsReq                         = 6436
	TravelBrochurePageResetScRsp                       = 6407
	TravelBrochureApplyPasterScRsp                     = 6474
	TravelBrochureSetPageDescStatusScRsp               = 6419
	TravelBrochureApplyPasterListScRsp                 = 6490
	TravelBrochurePageResetCsReq                       = 6426
	TravelBrochurePageUnlockScNotify                   = 6484
	TravelBrochureApplyPasterListCsReq                 = 6453
	TravelBrochureSelectMessageCsReq                   = 6467
	TravelBrochureUpdatePasterPosCsReq                 = 6434
	TravelBrochureSelectMessageScRsp                   = 6428
	TravelBrochureSetCustomValueScRsp                  = 6405
	TravelBrochureRemovePasterCsReq                    = 6424
	TravelBrochureGetPasterScNotify                    = 6446
	TravelBrochureUpdatePasterPosScRsp                 = 6443
	TravelBrochureGetDataScRsp                         = 6495
	TravelBrochureApplyPasterCsReq                     = 6452
	TravelBrochureRemovePasterScRsp                    = 6493
	GetTreasureDungeonActivityDataCsReq                = 4434
	UseTreasureDungeonItemScRsp                        = 4490
	EnterTreasureDungeonScRsp                          = 4425
	FightTreasureDungeonMonsterCsReq                   = 4475
	UseTreasureDungeonItemCsReq                        = 4453
	QuitTreasureDungeonScRsp                           = 4412
	QuitTreasureDungeonCsReq                           = 4459
	OpenTreasureDungeonGridCsReq                       = 4496
	GetTreasureDungeonActivityDataScRsp                = 4443
	InteractTreasureDungeonGridScRsp                   = 4407
	FightTreasureDungeonMonsterScRsp                   = 4419
	TreasureDungeonFinishScNotify                      = 4495
	OpenTreasureDungeonGridScRsp                       = 4405
	InteractTreasureDungeonGridCsReq                   = 4426
	TreasureDungeonDataScNotify                        = 4436
	EnterTreasureDungeonCsReq                          = 4446
	UnlockTutorialGuideScRsp                           = 1674
	UnlockTutorialCsReq                                = 1667
	FinishTutorialCsReq                                = 1624
	UnlockTutorialGuideCsReq                           = 1652
	GetTutorialCsReq                                   = 1636
	GetTutorialGuideScRsp                              = 1627
	GetTutorialGuideCsReq                              = 1684
	FinishTutorialScRsp                                = 1693
	FinishTutorialGuideCsReq                           = 1634
	UnlockTutorialScRsp                                = 1628
	GetTutorialScRsp                                   = 1695
	FinishTutorialGuideScRsp                           = 1643
	TakeChapterRewardScRsp                             = 424
	GetChapterScRsp                                    = 428
	SetCurWaypointCsReq                                = 484
	GetWaypointCsReq                                   = 436
	GetWaypointScRsp                                   = 495
	TakeChapterRewardCsReq                             = 474
	WaypointShowNewCsNotify                            = 452
	SetCurWaypointScRsp                                = 427
	GetChapterCsReq                                    = 467
	GetWolfBroGameDataScRsp                            = 6546
	WolfBroGameUseBulletCsReq                          = 6528
	WolfBroGameActivateBulletScRsp                     = 6504
	QuitWolfBroGameScRsp                               = 6543
	StartWolfBroGameScRsp                              = 6550
	WolfBroGamePickupBulletScRsp                       = 6513
	WolfBroGameUseBulletScRsp                          = 6510
	ArchiveWolfBroGameScRsp                            = 6523
	WolfBroGameActivateBulletCsReq                     = 6517
	WolfBroGameExplodeMonsterCsReq                     = 6540
	WolfBroGameDataChangeScNotify                      = 6545
	WolfBroGameExplodeMonsterScRsp                     = 6511
	ArchiveWolfBroGameCsReq                            = 6530
	StartWolfBroGameCsReq                              = 6512
	GetWolfBroGameDataCsReq                            = 6547
	RestoreWolfBroGameArchiveScRsp                     = 6549
	WolfBroGamePickupBulletCsReq                       = 6548
	RestoreWolfBroGameArchiveCsReq                     = 6515
	QuitWolfBroGameCsReq                               = 6509
	WorldUnlockScRsp                                   = 7627
	WorldUnlockCsReq                                   = 7626
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(GetMaterialSubmitActivityDataCsReq, func() any { return new(proto.GetMaterialSubmitActivityDataCsReq) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(TakeMaterialSubmitActivityRewardCsReq, func() any { return new(proto.TakeMaterialSubmitActivityRewardCsReq) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(SubmitMaterialSubmitActivityMaterialScRsp, func() any { return new(proto.SubmitMaterialSubmitActivityMaterialScRsp) })
	c.regMsg(TakeMaterialSubmitActivityRewardScRsp, func() any { return new(proto.TakeMaterialSubmitActivityRewardScRsp) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(SubmitMaterialSubmitActivityMaterialCsReq, func() any { return new(proto.SubmitMaterialSubmitActivityMaterialCsReq) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(GetMaterialSubmitActivityDataScRsp, func() any { return new(proto.GetMaterialSubmitActivityDataScRsp) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(QuickStartCocoonStageScRsp, func() any { return new(proto.QuickStartCocoonStageScRsp) })
	c.regMsg(QuickStartFarmElementCsReq, func() any { return new(proto.QuickStartFarmElementCsReq) })
	c.regMsg(QuickStartCocoonStageCsReq, func() any { return new(proto.QuickStartCocoonStageCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(QuickStartFarmElementScRsp, func() any { return new(proto.QuickStartFarmElementScRsp) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(ActivityRaidPlacingGameScRsp, func() any { return new(proto.ActivityRaidPlacingGameScRsp) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(ActivityRaidPlacingGameCsReq, func() any { return new(proto.ActivityRaidPlacingGameCsReq) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(LogisticsGameCsReq, func() any { return new(proto.LogisticsGameCsReq) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(GetUpdatedArchiveDataCsReq, func() any { return new(proto.GetUpdatedArchiveDataCsReq) })
	c.regMsg(GetUpdatedArchiveDataScRsp, func() any { return new(proto.GetUpdatedArchiveDataScRsp) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(MarkAvatarCsReq, func() any { return new(proto.MarkAvatarCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(GrowthTargetAvatarChangedScNotify, func() any { return new(proto.GrowthTargetAvatarChangedScNotify) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(AddMultiPathAvatarScNotify, func() any { return new(proto.AddMultiPathAvatarScNotify) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(SetGrowthTargetAvatarCsReq, func() any { return new(proto.SetGrowthTargetAvatarCsReq) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(MarkAvatarScRsp, func() any { return new(proto.MarkAvatarScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(SetGrowthTargetAvatarScRsp, func() any { return new(proto.SetGrowthTargetAvatarScRsp) })
	c.regMsg(QuitBattleScNotify, func() any { return new(proto.QuitBattleScNotify) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(QuitBattleCsReq, func() any { return new(proto.QuitBattleCsReq) })
	c.regMsg(ServerSimulateBattleFinishScNotify, func() any { return new(proto.ServerSimulateBattleFinishScNotify) })
	c.regMsg(QuitBattleScRsp, func() any { return new(proto.QuitBattleScRsp) })
	c.regMsg(BattleLogReportScRsp, func() any { return new(proto.BattleLogReportScRsp) })
	c.regMsg(BattleLogReportCsReq, func() any { return new(proto.BattleLogReportCsReq) })
	c.regMsg(ReBattleAfterBattleLoseCsNotify, func() any { return new(proto.ReBattleAfterBattleLoseCsNotify) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(RebattleByClientCsNotify, func() any { return new(proto.RebattleByClientCsNotify) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(GetBattleCollegeDataCsReq, func() any { return new(proto.GetBattleCollegeDataCsReq) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(BoxingClubRewardScNotify, func() any { return new(proto.BoxingClubRewardScNotify) })
	c.regMsg(GiveUpBoxingClubChallengeCsReq, func() any { return new(proto.GiveUpBoxingClubChallengeCsReq) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(SetBoxingClubResonanceLineupCsReq, func() any { return new(proto.SetBoxingClubResonanceLineupCsReq) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(BoxingClubChallengeUpdateScNotify, func() any { return new(proto.BoxingClubChallengeUpdateScNotify) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffCsReq, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffCsReq) })
	c.regMsg(ChooseBoxingClubResonanceCsReq, func() any { return new(proto.ChooseBoxingClubResonanceCsReq) })
	c.regMsg(MatchBoxingClubOpponentCsReq, func() any { return new(proto.MatchBoxingClubOpponentCsReq) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(GiveUpBoxingClubChallengeScRsp, func() any { return new(proto.GiveUpBoxingClubChallengeScRsp) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(StartBoxingClubBattleCsReq, func() any { return new(proto.StartBoxingClubBattleCsReq) })
	c.regMsg(GetBoxingClubInfoCsReq, func() any { return new(proto.GetBoxingClubInfoCsReq) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(RestartChallengePhaseCsReq, func() any { return new(proto.RestartChallengePhaseCsReq) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(GetChallengeGroupStatisticsScRsp, func() any { return new(proto.GetChallengeGroupStatisticsScRsp) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(ChallengeBossPhaseSettleNotify, func() any { return new(proto.ChallengeBossPhaseSettleNotify) })
	c.regMsg(StartPartialChallengeScRsp, func() any { return new(proto.StartPartialChallengeScRsp) })
	c.regMsg(EnterChallengeNextPhaseCsReq, func() any { return new(proto.EnterChallengeNextPhaseCsReq) })
	c.regMsg(EnterChallengeNextPhaseScRsp, func() any { return new(proto.EnterChallengeNextPhaseScRsp) })
	c.regMsg(StartPartialChallengeCsReq, func() any { return new(proto.StartPartialChallengeCsReq) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(GetChallengeGroupStatisticsCsReq, func() any { return new(proto.GetChallengeGroupStatisticsCsReq) })
	c.regMsg(RestartChallengePhaseScRsp, func() any { return new(proto.RestartChallengePhaseScRsp) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoCsReq, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoCsReq) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(ChessRogueNousEnableRogueTalentScRsp, func() any { return new(proto.ChessRogueNousEnableRogueTalentScRsp) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(ChessRogueNousEnableRogueTalentCsReq, func() any { return new(proto.ChessRogueNousEnableRogueTalentCsReq) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoScRsp, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoScRsp) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(ClockParkGetOngoingScriptInfoScRsp, func() any { return new(proto.ClockParkGetOngoingScriptInfoScRsp) })
	c.regMsg(ClockParkStartScriptCsReq, func() any { return new(proto.ClockParkStartScriptCsReq) })
	c.regMsg(ClockParkBattleEndScNotify, func() any { return new(proto.ClockParkBattleEndScNotify) })
	c.regMsg(ClockParkUnlockTalentScRsp, func() any { return new(proto.ClockParkUnlockTalentScRsp) })
	c.regMsg(ClockParkUnlockTalentCsReq, func() any { return new(proto.ClockParkUnlockTalentCsReq) })
	c.regMsg(ClockParkQuitScriptScRsp, func() any { return new(proto.ClockParkQuitScriptScRsp) })
	c.regMsg(ClockParkUseBuffCsReq, func() any { return new(proto.ClockParkUseBuffCsReq) })
	c.regMsg(ClockParkHandleWaitOperationScRsp, func() any { return new(proto.ClockParkHandleWaitOperationScRsp) })
	c.regMsg(ClockParkStartScriptScRsp, func() any { return new(proto.ClockParkStartScriptScRsp) })
	c.regMsg(ClockParkGetInfoScRsp, func() any { return new(proto.ClockParkGetInfoScRsp) })
	c.regMsg(ClockParkFinishScriptScNotify, func() any { return new(proto.ClockParkFinishScriptScNotify) })
	c.regMsg(ClockParkGetInfoCsReq, func() any { return new(proto.ClockParkGetInfoCsReq) })
	c.regMsg(ClockParkUseBuffScRsp, func() any { return new(proto.ClockParkUseBuffScRsp) })
	c.regMsg(ClockParkGetOngoingScriptInfoCsReq, func() any { return new(proto.ClockParkGetOngoingScriptInfoCsReq) })
	c.regMsg(ClockParkQuitScriptCsReq, func() any { return new(proto.ClockParkQuitScriptCsReq) })
	c.regMsg(ClockParkHandleWaitOperationCsReq, func() any { return new(proto.ClockParkHandleWaitOperationCsReq) })
	c.regMsg(ContentPackageGetDataScRsp, func() any { return new(proto.ContentPackageGetDataScRsp) })
	c.regMsg(ContentPackageUnlockScRsp, func() any { return new(proto.ContentPackageUnlockScRsp) })
	c.regMsg(ContentPackageGetDataCsReq, func() any { return new(proto.ContentPackageGetDataCsReq) })
	c.regMsg(ContentPackageSyncDataScNotify, func() any { return new(proto.ContentPackageSyncDataScNotify) })
	c.regMsg(ContentPackageUnlockCsReq, func() any { return new(proto.ContentPackageUnlockCsReq) })
	c.regMsg(ContentPackageTransferScNotify, func() any { return new(proto.ContentPackageTransferScNotify) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(DrinkMakerChallengeScRsp, func() any { return new(proto.DrinkMakerChallengeScRsp) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(DrinkMakerChallengeCsReq, func() any { return new(proto.DrinkMakerChallengeCsReq) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(DrinkMakerUpdateTipsNotify, func() any { return new(proto.DrinkMakerUpdateTipsNotify) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(DrinkMakerDayEndScNotify, func() any { return new(proto.DrinkMakerDayEndScNotify) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(EnterEraFlipperRegionCsReq, func() any { return new(proto.EnterEraFlipperRegionCsReq) })
	c.regMsg(EraFlipperDataChangeScNotify, func() any { return new(proto.EraFlipperDataChangeScNotify) })
	c.regMsg(ResetEraFlipperDataCsReq, func() any { return new(proto.ResetEraFlipperDataCsReq) })
	c.regMsg(ChangeEraFlipperDataScRsp, func() any { return new(proto.ChangeEraFlipperDataScRsp) })
	c.regMsg(ResetEraFlipperDataScRsp, func() any { return new(proto.ResetEraFlipperDataScRsp) })
	c.regMsg(ChangeEraFlipperDataCsReq, func() any { return new(proto.ChangeEraFlipperDataCsReq) })
	c.regMsg(GetEraFlipperDataCsReq, func() any { return new(proto.GetEraFlipperDataCsReq) })
	c.regMsg(GetEraFlipperDataScRsp, func() any { return new(proto.GetEraFlipperDataScRsp) })
	c.regMsg(EnterEraFlipperRegionScRsp, func() any { return new(proto.EnterEraFlipperRegionScRsp) })
	c.regMsg(EvolveBuildReRandomStageCsReq, func() any { return new(proto.EvolveBuildReRandomStageCsReq) })
	c.regMsg(EvolveBuildStartStageScRsp, func() any { return new(proto.EvolveBuildStartStageScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownScRsp, func() any { return new(proto.EvolveBuildShopAbilityDownScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownCsReq, func() any { return new(proto.EvolveBuildShopAbilityDownCsReq) })
	c.regMsg(EvolveBuildStartLevelScRsp, func() any { return new(proto.EvolveBuildStartLevelScRsp) })
	c.regMsg(EvolveBuildShopAbilityUpCsReq, func() any { return new(proto.EvolveBuildShopAbilityUpCsReq) })
	c.regMsg(EvolveBuildStartLevelCsReq, func() any { return new(proto.EvolveBuildStartLevelCsReq) })
	c.regMsg(EvolveBuildFinishScNotify, func() any { return new(proto.EvolveBuildFinishScNotify) })
	c.regMsg(EvolveBuildQueryInfoCsReq, func() any { return new(proto.EvolveBuildQueryInfoCsReq) })
	c.regMsg(EvolveBuildCoinNotify, func() any { return new(proto.EvolveBuildCoinNotify) })
	c.regMsg(EvolveBuildTakeExpRewardScRsp, func() any { return new(proto.EvolveBuildTakeExpRewardScRsp) })
	c.regMsg(EvolveBuildLeaveCsReq, func() any { return new(proto.EvolveBuildLeaveCsReq) })
	c.regMsg(EvolveBuildStartStageCsReq, func() any { return new(proto.EvolveBuildStartStageCsReq) })
	c.regMsg(EvolveBuildGiveupCsReq, func() any { return new(proto.EvolveBuildGiveupCsReq) })
	c.regMsg(EvolveBuildShopAbilityUpScRsp, func() any { return new(proto.EvolveBuildShopAbilityUpScRsp) })
	c.regMsg(EvolveBuildTakeExpRewardCsReq, func() any { return new(proto.EvolveBuildTakeExpRewardCsReq) })
	c.regMsg(EvolveBuildReRandomStageScRsp, func() any { return new(proto.EvolveBuildReRandomStageScRsp) })
	c.regMsg(EvolveBuildUnlockInfoNotify, func() any { return new(proto.EvolveBuildUnlockInfoNotify) })
	c.regMsg(EvolveBuildShopAbilityResetCsReq, func() any { return new(proto.EvolveBuildShopAbilityResetCsReq) })
	c.regMsg(EvolveBuildGiveupScRsp, func() any { return new(proto.EvolveBuildGiveupScRsp) })
	c.regMsg(EvolveBuildQueryInfoScRsp, func() any { return new(proto.EvolveBuildQueryInfoScRsp) })
	c.regMsg(EvolveBuildLeaveScRsp, func() any { return new(proto.EvolveBuildLeaveScRsp) })
	c.regMsg(EvolveBuildShopAbilityResetScRsp, func() any { return new(proto.EvolveBuildShopAbilityResetScRsp) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(CancelExpeditionScRsp, func() any { return new(proto.CancelExpeditionScRsp) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(TakeMultipleActivityExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleActivityExpeditionRewardCsReq) })
	c.regMsg(TakeMultipleActivityExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleActivityExpeditionRewardScRsp) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(FightKickOutScNotify, func() any { return new(proto.FightKickOutScNotify) })
	c.regMsg(FightEnterScRsp, func() any { return new(proto.FightEnterScRsp) })
	c.regMsg(FightEnterCsReq, func() any { return new(proto.FightEnterCsReq) })
	c.regMsg(FightSessionStopScNotify, func() any { return new(proto.FightSessionStopScNotify) })
	c.regMsg(FightLeaveScNotify, func() any { return new(proto.FightLeaveScNotify) })
	c.regMsg(FightHeartBeatScRsp, func() any { return new(proto.FightHeartBeatScRsp) })
	c.regMsg(FightHeartBeatCsReq, func() any { return new(proto.FightHeartBeatCsReq) })
	c.regMsg(EnterFightActivityStageScRsp, func() any { return new(proto.EnterFightActivityStageScRsp) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(TakeFightActivityRewardCsReq, func() any { return new(proto.TakeFightActivityRewardCsReq) })
	c.regMsg(TakeFightActivityRewardScRsp, func() any { return new(proto.TakeFightActivityRewardScRsp) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(FightFestUpdateCoinNotify, func() any { return new(proto.FightFestUpdateCoinNotify) })
	c.regMsg(StartFightFestCsReq, func() any { return new(proto.StartFightFestCsReq) })
	c.regMsg(GetFightFestDataCsReq, func() any { return new(proto.GetFightFestDataCsReq) })
	c.regMsg(FightFestScoreUpdateNotify, func() any { return new(proto.FightFestScoreUpdateNotify) })
	c.regMsg(StartFightFestScRsp, func() any { return new(proto.StartFightFestScRsp) })
	c.regMsg(FightFestUnlockSkillNotify, func() any { return new(proto.FightFestUnlockSkillNotify) })
	c.regMsg(FightFestUpdateChallengeRecordNotify, func() any { return new(proto.FightFestUpdateChallengeRecordNotify) })
	c.regMsg(GetFightFestDataScRsp, func() any { return new(proto.GetFightFestDataScRsp) })
	c.regMsg(FightMatch3SwapCsReq, func() any { return new(proto.FightMatch3SwapCsReq) })
	c.regMsg(FightMatch3ChatScNotify, func() any { return new(proto.FightMatch3ChatScNotify) })
	c.regMsg(FightMatch3SwapScRsp, func() any { return new(proto.FightMatch3SwapScRsp) })
	c.regMsg(FightMatch3DataScRsp, func() any { return new(proto.FightMatch3DataScRsp) })
	c.regMsg(FightMatch3TurnEndScNotify, func() any { return new(proto.FightMatch3TurnEndScNotify) })
	c.regMsg(FightMatch3StartCountDownScNotify, func() any { return new(proto.FightMatch3StartCountDownScNotify) })
	c.regMsg(FightMatch3DataCsReq, func() any { return new(proto.FightMatch3DataCsReq) })
	c.regMsg(FightMatch3ForceUpdateNotify, func() any { return new(proto.FightMatch3ForceUpdateNotify) })
	c.regMsg(FightMatch3ChatCsReq, func() any { return new(proto.FightMatch3ChatCsReq) })
	c.regMsg(FightMatch3OpponentDataScNotify, func() any { return new(proto.FightMatch3OpponentDataScNotify) })
	c.regMsg(FightMatch3ChatScRsp, func() any { return new(proto.FightMatch3ChatScRsp) })
	c.regMsg(FightMatch3TurnStartScNotify, func() any { return new(proto.FightMatch3TurnStartScNotify) })
	c.regMsg(GetFriendDevelopmentInfoScRsp, func() any { return new(proto.GetFriendDevelopmentInfoScRsp) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(SetFriendMarkCsReq, func() any { return new(proto.SetFriendMarkCsReq) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(TakeAssistRewardCsReq, func() any { return new(proto.TakeAssistRewardCsReq) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(GetFriendBattleRecordDetailScRsp, func() any { return new(proto.GetFriendBattleRecordDetailScRsp) })
	c.regMsg(TakeAssistRewardScRsp, func() any { return new(proto.TakeAssistRewardScRsp) })
	c.regMsg(GetFriendChallengeDetailScRsp, func() any { return new(proto.GetFriendChallengeDetailScRsp) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(GetFriendChallengeLineupCsReq, func() any { return new(proto.GetFriendChallengeLineupCsReq) })
	c.regMsg(GetFriendAssistListScRsp, func() any { return new(proto.GetFriendAssistListScRsp) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(SetFriendMarkScRsp, func() any { return new(proto.SetFriendMarkScRsp) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(GetFriendChallengeDetailCsReq, func() any { return new(proto.GetFriendChallengeDetailCsReq) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(DeleteFriendScRsp, func() any { return new(proto.DeleteFriendScRsp) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(NewAssistHistoryNotify, func() any { return new(proto.NewAssistHistoryNotify) })
	c.regMsg(GetFriendDevelopmentInfoCsReq, func() any { return new(proto.GetFriendDevelopmentInfoCsReq) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(CurAssistChangedNotify, func() any { return new(proto.CurAssistChangedNotify) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(SetForbidOtherApplyFriendScRsp, func() any { return new(proto.SetForbidOtherApplyFriendScRsp) })
	c.regMsg(GetFriendChallengeLineupScRsp, func() any { return new(proto.GetFriendChallengeLineupScRsp) })
	c.regMsg(GetFriendBattleRecordDetailCsReq, func() any { return new(proto.GetFriendBattleRecordDetailCsReq) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(GetFriendAssistListCsReq, func() any { return new(proto.GetFriendAssistListCsReq) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(HeliobusStartRaidScRsp, func() any { return new(proto.HeliobusStartRaidScRsp) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(HeliobusStartRaidCsReq, func() any { return new(proto.HeliobusStartRaidCsReq) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(HeliobusLineupUpdateScNotify, func() any { return new(proto.HeliobusLineupUpdateScNotify) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(AddEquipmentScNotify, func() any { return new(proto.AddEquipmentScNotify) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(DeleteRelicFilterPlanScRsp, func() any { return new(proto.DeleteRelicFilterPlanScRsp) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(RelicReforgeConfirmCsReq, func() any { return new(proto.RelicReforgeConfirmCsReq) })
	c.regMsg(RelicReforgeCsReq, func() any { return new(proto.RelicReforgeCsReq) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(ModifyRelicFilterPlanScRsp, func() any { return new(proto.ModifyRelicFilterPlanScRsp) })
	c.regMsg(DeleteRelicFilterPlanCsReq, func() any { return new(proto.DeleteRelicFilterPlanCsReq) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(GetRelicFilterPlanScRsp, func() any { return new(proto.GetRelicFilterPlanScRsp) })
	c.regMsg(GetRelicFilterPlanCsReq, func() any { return new(proto.GetRelicFilterPlanCsReq) })
	c.regMsg(MarkRelicFilterPlanScRsp, func() any { return new(proto.MarkRelicFilterPlanScRsp) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(RelicReforgeConfirmScRsp, func() any { return new(proto.RelicReforgeConfirmScRsp) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(AddRelicFilterPlanScRsp, func() any { return new(proto.AddRelicFilterPlanScRsp) })
	c.regMsg(MarkRelicFilterPlanCsReq, func() any { return new(proto.MarkRelicFilterPlanCsReq) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(RelicReforgeScRsp, func() any { return new(proto.RelicReforgeScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(ModifyRelicFilterPlanCsReq, func() any { return new(proto.ModifyRelicFilterPlanCsReq) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(RelicFilterPlanClearNameScNotify, func() any { return new(proto.RelicFilterPlanClearNameScNotify) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(AddRelicFilterPlanCsReq, func() any { return new(proto.AddRelicFilterPlanCsReq) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(GetStageLineupCsReq, func() any { return new(proto.GetStageLineupCsReq) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(VirtualLineupTrialAvatarChangeScNotify, func() any { return new(proto.VirtualLineupTrialAvatarChangeScNotify) })
	c.regMsg(LobbyBeginCsReq, func() any { return new(proto.LobbyBeginCsReq) })
	c.regMsg(LobbyInviteScRsp, func() any { return new(proto.LobbyInviteScRsp) })
	c.regMsg(LobbySyncInfoScNotify, func() any { return new(proto.LobbySyncInfoScNotify) })
	c.regMsg(LobbyInviteCsReq, func() any { return new(proto.LobbyInviteCsReq) })
	c.regMsg(LobbyKickOutScRsp, func() any { return new(proto.LobbyKickOutScRsp) })
	c.regMsg(LobbyQuitCsReq, func() any { return new(proto.LobbyQuitCsReq) })
	c.regMsg(LobbyQuitScRsp, func() any { return new(proto.LobbyQuitScRsp) })
	c.regMsg(LobbyCreateScRsp, func() any { return new(proto.LobbyCreateScRsp) })
	c.regMsg(LobbyKickOutCsReq, func() any { return new(proto.LobbyKickOutCsReq) })
	c.regMsg(LobbyModifyPlayerInfoScRsp, func() any { return new(proto.LobbyModifyPlayerInfoScRsp) })
	c.regMsg(LobbyGetInfoCsReq, func() any { return new(proto.LobbyGetInfoCsReq) })
	c.regMsg(LobbyGetInfoScRsp, func() any { return new(proto.LobbyGetInfoScRsp) })
	c.regMsg(LobbyModifyPlayerInfoCsReq, func() any { return new(proto.LobbyModifyPlayerInfoCsReq) })
	c.regMsg(LobbyBeginScRsp, func() any { return new(proto.LobbyBeginScRsp) })
	c.regMsg(LobbyInviteScNotify, func() any { return new(proto.LobbyInviteScNotify) })
	c.regMsg(LobbyJoinCsReq, func() any { return new(proto.LobbyJoinCsReq) })
	c.regMsg(LobbyJoinScRsp, func() any { return new(proto.LobbyJoinScRsp) })
	c.regMsg(LobbyCreateCsReq, func() any { return new(proto.LobbyCreateCsReq) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(RemoveRotaterScRsp, func() any { return new(proto.RemoveRotaterScRsp) })
	c.regMsg(GetMapRotationDataCsReq, func() any { return new(proto.GetMapRotationDataCsReq) })
	c.regMsg(UpdateRotaterScNotify, func() any { return new(proto.UpdateRotaterScNotify) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(UpdateMapRotationDataScNotify, func() any { return new(proto.UpdateMapRotationDataScNotify) })
	c.regMsg(RemoveRotaterCsReq, func() any { return new(proto.RemoveRotaterCsReq) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(UpdateMarkChestCsReq, func() any { return new(proto.UpdateMarkChestCsReq) })
	c.regMsg(MarkChestChangedScNotify, func() any { return new(proto.MarkChestChangedScNotify) })
	c.regMsg(GetMarkChestCsReq, func() any { return new(proto.GetMarkChestCsReq) })
	c.regMsg(UpdateMarkChestScRsp, func() any { return new(proto.UpdateMarkChestScRsp) })
	c.regMsg(GetMarkChestScRsp, func() any { return new(proto.GetMarkChestScRsp) })
	c.regMsg(CancelMatchCsReq, func() any { return new(proto.CancelMatchCsReq) })
	c.regMsg(GetCrossInfoCsReq, func() any { return new(proto.GetCrossInfoCsReq) })
	c.regMsg(GetCrossInfoScRsp, func() any { return new(proto.GetCrossInfoScRsp) })
	c.regMsg(StartMatchCsReq, func() any { return new(proto.StartMatchCsReq) })
	c.regMsg(MatchResultScNotify, func() any { return new(proto.MatchResultScNotify) })
	c.regMsg(StartMatchScRsp, func() any { return new(proto.StartMatchScRsp) })
	c.regMsg(CancelMatchScRsp, func() any { return new(proto.CancelMatchScRsp) })
	c.regMsg(MatchThreeSetBirdPosScRsp, func() any { return new(proto.MatchThreeSetBirdPosScRsp) })
	c.regMsg(MatchThreeLevelEndScRsp, func() any { return new(proto.MatchThreeLevelEndScRsp) })
	c.regMsg(MatchThreeGetDataScRsp, func() any { return new(proto.MatchThreeGetDataScRsp) })
	c.regMsg(MatchThreeSyncDataScNotify, func() any { return new(proto.MatchThreeSyncDataScNotify) })
	c.regMsg(MatchThreeSetBirdPosCsReq, func() any { return new(proto.MatchThreeSetBirdPosCsReq) })
	c.regMsg(MatchThreeLevelEndCsReq, func() any { return new(proto.MatchThreeLevelEndCsReq) })
	c.regMsg(MatchThreeGetDataCsReq, func() any { return new(proto.MatchThreeGetDataCsReq) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(GetMissionMessageInfoCsReq, func() any { return new(proto.GetMissionMessageInfoCsReq) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(GetMissionMessageInfoScRsp, func() any { return new(proto.GetMissionMessageInfoScRsp) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(DifficultyAdjustmentGetDataCsReq, func() any { return new(proto.DifficultyAdjustmentGetDataCsReq) })
	c.regMsg(MazeKillDirectCsReq, func() any { return new(proto.MazeKillDirectCsReq) })
	c.regMsg(SubmitOrigamiItemCsReq, func() any { return new(proto.SubmitOrigamiItemCsReq) })
	c.regMsg(MazeKillDirectScRsp, func() any { return new(proto.MazeKillDirectScRsp) })
	c.regMsg(UpdateGunPlayDataScRsp, func() any { return new(proto.UpdateGunPlayDataScRsp) })
	c.regMsg(GetGunPlayDataCsReq, func() any { return new(proto.GetGunPlayDataCsReq) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(DifficultyAdjustmentUpdateDataScRsp, func() any { return new(proto.DifficultyAdjustmentUpdateDataScRsp) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(GetGunPlayDataScRsp, func() any { return new(proto.GetGunPlayDataScRsp) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(UpdateMovieRacingDataScRsp, func() any { return new(proto.UpdateMovieRacingDataScRsp) })
	c.regMsg(SubmitOrigamiItemScRsp, func() any { return new(proto.SubmitOrigamiItemScRsp) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(UpdateMovieRacingDataCsReq, func() any { return new(proto.UpdateMovieRacingDataCsReq) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(DifficultyAdjustmentUpdateDataCsReq, func() any { return new(proto.DifficultyAdjustmentUpdateDataCsReq) })
	c.regMsg(DifficultyAdjustmentGetDataScRsp, func() any { return new(proto.DifficultyAdjustmentGetDataScRsp) })
	c.regMsg(UpdateGunPlayDataCsReq, func() any { return new(proto.UpdateGunPlayDataCsReq) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(AcceptMissionEventCsReq, func() any { return new(proto.AcceptMissionEventCsReq) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(FinishedMissionScNotify, func() any { return new(proto.FinishedMissionScNotify) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
	c.regMsg(InterruptMissionEventCsReq, func() any { return new(proto.InterruptMissionEventCsReq) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(MissionEventRewardScNotify, func() any { return new(proto.MissionEventRewardScNotify) })
	c.regMsg(AcceptMissionEventScRsp, func() any { return new(proto.AcceptMissionEventScRsp) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(MonopolyGuessChooseScRsp, func() any { return new(proto.MonopolyGuessChooseScRsp) })
	c.regMsg(MonopolyActionResultScNotify, func() any { return new(proto.MonopolyActionResultScNotify) })
	c.regMsg(DailyFirstEnterMonopolyActivityScRsp, func() any { return new(proto.DailyFirstEnterMonopolyActivityScRsp) })
	c.regMsg(GetMonopolyInfoScRsp, func() any { return new(proto.GetMonopolyInfoScRsp) })
	c.regMsg(MonopolyEventLoadUpdateScNotify, func() any { return new(proto.MonopolyEventLoadUpdateScNotify) })
	c.regMsg(GetMonopolyInfoCsReq, func() any { return new(proto.GetMonopolyInfoCsReq) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(GetMbtiReportCsReq, func() any { return new(proto.GetMbtiReportCsReq) })
	c.regMsg(MonopolySttUpdateScNotify, func() any { return new(proto.MonopolySttUpdateScNotify) })
	c.regMsg(MonopolyGetDailyInitItemCsReq, func() any { return new(proto.MonopolyGetDailyInitItemCsReq) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(MonopolyGiveUpCurContentScRsp, func() any { return new(proto.MonopolyGiveUpCurContentScRsp) })
	c.regMsg(MonopolyAcceptQuizCsReq, func() any { return new(proto.MonopolyAcceptQuizCsReq) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(MonopolyQuizDurationChangeScNotify, func() any { return new(proto.MonopolyQuizDurationChangeScNotify) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(MonopolyMoveScRsp, func() any { return new(proto.MonopolyMoveScRsp) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(MonopolyEventSelectFriendCsReq, func() any { return new(proto.MonopolyEventSelectFriendCsReq) })
	c.regMsg(MonopolyClickMbtiReportCsReq, func() any { return new(proto.MonopolyClickMbtiReportCsReq) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(MonopolyConditionUpdateScNotify, func() any { return new(proto.MonopolyConditionUpdateScNotify) })
	c.regMsg(MonopolyGiveUpCurContentCsReq, func() any { return new(proto.MonopolyGiveUpCurContentCsReq) })
	c.regMsg(DailyFirstEnterMonopolyActivityCsReq, func() any { return new(proto.DailyFirstEnterMonopolyActivityCsReq) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(GetMonopolyDailyReportCsReq, func() any { return new(proto.GetMonopolyDailyReportCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(MonopolyGetDailyInitItemScRsp, func() any { return new(proto.MonopolyGetDailyInitItemScRsp) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(MonopolyGuessChooseCsReq, func() any { return new(proto.MonopolyGuessChooseCsReq) })
	c.regMsg(MonopolyGameRaiseRatioCsReq, func() any { return new(proto.MonopolyGameRaiseRatioCsReq) })
	c.regMsg(MonopolyDailySettleScNotify, func() any { return new(proto.MonopolyDailySettleScNotify) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(MonopolySelectOptionCsReq, func() any { return new(proto.MonopolySelectOptionCsReq) })
	c.regMsg(MonopolyGameGachaCsReq, func() any { return new(proto.MonopolyGameGachaCsReq) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(MonopolyGameBingoFlipCardCsReq, func() any { return new(proto.MonopolyGameBingoFlipCardCsReq) })
	c.regMsg(MonopolyGuessBuyInformationScRsp, func() any { return new(proto.MonopolyGuessBuyInformationScRsp) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(MonopolyLikeCsReq, func() any { return new(proto.MonopolyLikeCsReq) })
	c.regMsg(MonopolyMoveCsReq, func() any { return new(proto.MonopolyMoveCsReq) })
	c.regMsg(MonopolyGameSettleScNotify, func() any { return new(proto.MonopolyGameSettleScNotify) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(MonopolyAcceptQuizScRsp, func() any { return new(proto.MonopolyAcceptQuizScRsp) })
	c.regMsg(MonopolyGuessDrawScNotify, func() any { return new(proto.MonopolyGuessDrawScNotify) })
	c.regMsg(MonopolyLikeScRsp, func() any { return new(proto.MonopolyLikeScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(MonopolyClickMbtiReportScRsp, func() any { return new(proto.MonopolyClickMbtiReportScRsp) })
	c.regMsg(MonopolyClickCellCsReq, func() any { return new(proto.MonopolyClickCellCsReq) })
	c.regMsg(MonopolyGameBingoFlipCardScRsp, func() any { return new(proto.MonopolyGameBingoFlipCardScRsp) })
	c.regMsg(MonopolyCheatDiceScRsp, func() any { return new(proto.MonopolyCheatDiceScRsp) })
	c.regMsg(MonopolyLikeScNotify, func() any { return new(proto.MonopolyLikeScNotify) })
	c.regMsg(GetMonopolyFriendRankingListCsReq, func() any { return new(proto.GetMonopolyFriendRankingListCsReq) })
	c.regMsg(MonopolyEventSelectFriendScRsp, func() any { return new(proto.MonopolyEventSelectFriendScRsp) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(GetMonopolyDailyReportScRsp, func() any { return new(proto.GetMonopolyDailyReportScRsp) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(MonopolySelectOptionScRsp, func() any { return new(proto.MonopolySelectOptionScRsp) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(MonopolyGuessBuyInformationCsReq, func() any { return new(proto.MonopolyGuessBuyInformationCsReq) })
	c.regMsg(MonopolySocialEventEffectScNotify, func() any { return new(proto.MonopolySocialEventEffectScNotify) })
	c.regMsg(MonopolyClickCellScRsp, func() any { return new(proto.MonopolyClickCellScRsp) })
	c.regMsg(MonopolyGameCreateScNotify, func() any { return new(proto.MonopolyGameCreateScNotify) })
	c.regMsg(GetMbtiReportScRsp, func() any { return new(proto.GetMbtiReportScRsp) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(MultiplayerMatch3FinishScNotify, func() any { return new(proto.MultiplayerMatch3FinishScNotify) })
	c.regMsg(MultiplayerFightGameStateCsReq, func() any { return new(proto.MultiplayerFightGameStateCsReq) })
	c.regMsg(MultiplayerFightGameStateScRsp, func() any { return new(proto.MultiplayerFightGameStateScRsp) })
	c.regMsg(MultiplayerGetFightGateCsReq, func() any { return new(proto.MultiplayerGetFightGateCsReq) })
	c.regMsg(MultiplayerFightGameStartScNotify, func() any { return new(proto.MultiplayerFightGameStartScNotify) })
	c.regMsg(MultiplayerFightGameFinishScNotify, func() any { return new(proto.MultiplayerFightGameFinishScNotify) })
	c.regMsg(MultiplayerGetFightGateScRsp, func() any { return new(proto.MultiplayerGetFightGateScRsp) })
	c.regMsg(MultiplayerFightGiveUpCsReq, func() any { return new(proto.MultiplayerFightGiveUpCsReq) })
	c.regMsg(MultiplayerFightGiveUpScRsp, func() any { return new(proto.MultiplayerFightGiveUpScRsp) })
	c.regMsg(GetMultipleDropInfoCsReq, func() any { return new(proto.GetMultipleDropInfoCsReq) })
	c.regMsg(GetPlayerReturnMultiDropInfoScRsp, func() any { return new(proto.GetPlayerReturnMultiDropInfoScRsp) })
	c.regMsg(MultipleDropInfoNotify, func() any { return new(proto.MultipleDropInfoNotify) })
	c.regMsg(GetPlayerReturnMultiDropInfoCsReq, func() any { return new(proto.GetPlayerReturnMultiDropInfoCsReq) })
	c.regMsg(MultipleDropInfoScNotify, func() any { return new(proto.MultipleDropInfoScNotify) })
	c.regMsg(GetMultipleDropInfoScRsp, func() any { return new(proto.GetMultipleDropInfoScRsp) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(MuseumTakeCollectRewardScRsp, func() any { return new(proto.MuseumTakeCollectRewardScRsp) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(MuseumTargetMissionFinishNotify, func() any { return new(proto.MuseumTargetMissionFinishNotify) })
	c.regMsg(MuseumTargetRewardNotify, func() any { return new(proto.MuseumTargetRewardNotify) })
	c.regMsg(MuseumTakeCollectRewardCsReq, func() any { return new(proto.MuseumTakeCollectRewardCsReq) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(MusicRhythmStartLevelScRsp, func() any { return new(proto.MusicRhythmStartLevelScRsp) })
	c.regMsg(MusicRhythmFinishLevelCsReq, func() any { return new(proto.MusicRhythmFinishLevelCsReq) })
	c.regMsg(MusicRhythmStartLevelCsReq, func() any { return new(proto.MusicRhythmStartLevelCsReq) })
	c.regMsg(MusicRhythmDataCsReq, func() any { return new(proto.MusicRhythmDataCsReq) })
	c.regMsg(MusicRhythmFinishLevelScRsp, func() any { return new(proto.MusicRhythmFinishLevelScRsp) })
	c.regMsg(MusicRhythmDataScRsp, func() any { return new(proto.MusicRhythmDataScRsp) })
	c.regMsg(MusicRhythmUnlockSongNotify, func() any { return new(proto.MusicRhythmUnlockSongNotify) })
	c.regMsg(MusicRhythmSaveSongConfigDataCsReq, func() any { return new(proto.MusicRhythmSaveSongConfigDataCsReq) })
	c.regMsg(MusicRhythmMaxDifficultyLevelsUnlockNotify, func() any { return new(proto.MusicRhythmMaxDifficultyLevelsUnlockNotify) })
	c.regMsg(MusicRhythmUnlockTrackScNotify, func() any { return new(proto.MusicRhythmUnlockTrackScNotify) })
	c.regMsg(MusicRhythmSaveSongConfigDataScRsp, func() any { return new(proto.MusicRhythmSaveSongConfigDataScRsp) })
	c.regMsg(MusicRhythmUnlockSongSfxScNotify, func() any { return new(proto.MusicRhythmUnlockSongSfxScNotify) })
	c.regMsg(SubmitOfferingItemScRsp, func() any { return new(proto.SubmitOfferingItemScRsp) })
	c.regMsg(TakeOfferingRewardScRsp, func() any { return new(proto.TakeOfferingRewardScRsp) })
	c.regMsg(TakeOfferingRewardCsReq, func() any { return new(proto.TakeOfferingRewardCsReq) })
	c.regMsg(GetOfferingInfoCsReq, func() any { return new(proto.GetOfferingInfoCsReq) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(SubmitOfferingItemCsReq, func() any { return new(proto.SubmitOfferingItemCsReq) })
	c.regMsg(OfferingInfoScNotify, func() any { return new(proto.OfferingInfoScNotify) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(AcceptedPamMissionExpireCsReq, func() any { return new(proto.AcceptedPamMissionExpireCsReq) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(GetPamSkinDataScRsp, func() any { return new(proto.GetPamSkinDataScRsp) })
	c.regMsg(SelectPamSkinScRsp, func() any { return new(proto.SelectPamSkinScRsp) })
	c.regMsg(SelectPamSkinCsReq, func() any { return new(proto.SelectPamSkinCsReq) })
	c.regMsg(UnlockPamSkinScNotify, func() any { return new(proto.UnlockPamSkinScNotify) })
	c.regMsg(GetPamSkinDataCsReq, func() any { return new(proto.GetPamSkinDataCsReq) })
	c.regMsg(SummonPetCsReq, func() any { return new(proto.SummonPetCsReq) })
	c.regMsg(RecallPetScRsp, func() any { return new(proto.RecallPetScRsp) })
	c.regMsg(GetPetDataCsReq, func() any { return new(proto.GetPetDataCsReq) })
	c.regMsg(CurPetChangedScNotify, func() any { return new(proto.CurPetChangedScNotify) })
	c.regMsg(RecallPetCsReq, func() any { return new(proto.RecallPetCsReq) })
	c.regMsg(SummonPetScRsp, func() any { return new(proto.SummonPetScRsp) })
	c.regMsg(GetPetDataScRsp, func() any { return new(proto.GetPetDataScRsp) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(SetAvatarPathCsReq, func() any { return new(proto.SetAvatarPathCsReq) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(UpdatePsnSettingsInfoCsReq, func() any { return new(proto.UpdatePsnSettingsInfoCsReq) })
	c.regMsg(SetMultipleAvatarPathsCsReq, func() any { return new(proto.SetMultipleAvatarPathsCsReq) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(SetMultipleAvatarPathsScRsp, func() any { return new(proto.SetMultipleAvatarPathsScRsp) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(UpdatePlayerSettingCsReq, func() any { return new(proto.UpdatePlayerSettingCsReq) })
	c.regMsg(QueryProductInfoCsReq, func() any { return new(proto.QueryProductInfoCsReq) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(UpdatePlayerSettingScRsp, func() any { return new(proto.UpdatePlayerSettingScRsp) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(UnlockAvatarPathScRsp, func() any { return new(proto.UnlockAvatarPathScRsp) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(ClientDownloadDataScNotify, func() any { return new(proto.ClientDownloadDataScNotify) })
	c.regMsg(GetMultiPathAvatarInfoCsReq, func() any { return new(proto.GetMultiPathAvatarInfoCsReq) })
	c.regMsg(ClientObjUploadScRsp, func() any { return new(proto.ClientObjUploadScRsp) })
	c.regMsg(GetAuthkeyCsReq, func() any { return new(proto.GetAuthkeyCsReq) })
	c.regMsg(ExchangeStaminaCsReq, func() any { return new(proto.ExchangeStaminaCsReq) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(GetAuthkeyScRsp, func() any { return new(proto.GetAuthkeyScRsp) })
	c.regMsg(ClientObjUploadCsReq, func() any { return new(proto.ClientObjUploadCsReq) })
	c.regMsg(GetGameStateServiceConfigCsReq, func() any { return new(proto.GetGameStateServiceConfigCsReq) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(RegionStopScNotify, func() any { return new(proto.RegionStopScNotify) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(SetRedPointStatusScNotify, func() any { return new(proto.SetRedPointStatusScNotify) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(SetAvatarPathScRsp, func() any { return new(proto.SetAvatarPathScRsp) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(GetGameStateServiceConfigScRsp, func() any { return new(proto.GetGameStateServiceConfigScRsp) })
	c.regMsg(AvatarPathChangedNotify, func() any { return new(proto.AvatarPathChangedNotify) })
	c.regMsg(GetLevelRewardScRsp, func() any { return new(proto.GetLevelRewardScRsp) })
	c.regMsg(AceAntiCheaterCsReq, func() any { return new(proto.AceAntiCheaterCsReq) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(UnlockAvatarPathCsReq, func() any { return new(proto.UnlockAvatarPathCsReq) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(UpdatePsnSettingsInfoScRsp, func() any { return new(proto.UpdatePsnSettingsInfoScRsp) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(GetMultiPathAvatarInfoScRsp, func() any { return new(proto.GetMultiPathAvatarInfoScRsp) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(RaidCollectionDataScRsp, func() any { return new(proto.RaidCollectionDataScRsp) })
	c.regMsg(RaidCollectionDataScNotify, func() any { return new(proto.RaidCollectionDataScNotify) })
	c.regMsg(RaidCollectionEnterNextRaidCsReq, func() any { return new(proto.RaidCollectionEnterNextRaidCsReq) })
	c.regMsg(RaidCollectionEnterNextRaidScRsp, func() any { return new(proto.RaidCollectionEnterNextRaidScRsp) })
	c.regMsg(RaidCollectionDataCsReq, func() any { return new(proto.RaidCollectionDataCsReq) })
	c.regMsg(GetBigDataRecommendScRsp, func() any { return new(proto.GetBigDataRecommendScRsp) })
	c.regMsg(GetChallengeRecommendLineupListScRsp, func() any { return new(proto.GetChallengeRecommendLineupListScRsp) })
	c.regMsg(RelicAvatarRecommendScRsp, func() any { return new(proto.RelicAvatarRecommendScRsp) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(GetChallengeRecommendLineupListCsReq, func() any { return new(proto.GetChallengeRecommendLineupListCsReq) })
	c.regMsg(RelicAvatarRecommendCsReq, func() any { return new(proto.RelicAvatarRecommendCsReq) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(GetBigDataRecommendCsReq, func() any { return new(proto.GetBigDataRecommendCsReq) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(GetAllRedDotDataCsReq, func() any { return new(proto.GetAllRedDotDataCsReq) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(RelicSmartWearGetPlanCsReq, func() any { return new(proto.RelicSmartWearGetPlanCsReq) })
	c.regMsg(RelicSmartWearDeletePlanCsReq, func() any { return new(proto.RelicSmartWearDeletePlanCsReq) })
	c.regMsg(RelicSmartWearGetPlanScRsp, func() any { return new(proto.RelicSmartWearGetPlanScRsp) })
	c.regMsg(RelicSmartWearAddPlanScRsp, func() any { return new(proto.RelicSmartWearAddPlanScRsp) })
	c.regMsg(RelicSmartWearUpdatePlanScRsp, func() any { return new(proto.RelicSmartWearUpdatePlanScRsp) })
	c.regMsg(RelicSmartWearUpdatePlanCsReq, func() any { return new(proto.RelicSmartWearUpdatePlanCsReq) })
	c.regMsg(RelicSmartWearAddPlanCsReq, func() any { return new(proto.RelicSmartWearAddPlanCsReq) })
	c.regMsg(RelicSmartWearDeletePlanScRsp, func() any { return new(proto.RelicSmartWearDeletePlanScRsp) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(SyncRogueAeonScNotify, func() any { return new(proto.SyncRogueAeonScNotify) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(SyncRogueSeasonFinishScNotify, func() any { return new(proto.SyncRogueSeasonFinishScNotify) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SyncRoguePickAvatarInfoScNotify, func() any { return new(proto.SyncRoguePickAvatarInfoScNotify) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(RogueArcadeRestartScRsp, func() any { return new(proto.RogueArcadeRestartScRsp) })
	c.regMsg(RogueArcadeLeaveScRsp, func() any { return new(proto.RogueArcadeLeaveScRsp) })
	c.regMsg(RogueArcadeStartScRsp, func() any { return new(proto.RogueArcadeStartScRsp) })
	c.regMsg(RogueArcadeGetInfoCsReq, func() any { return new(proto.RogueArcadeGetInfoCsReq) })
	c.regMsg(RogueArcadeLeaveCsReq, func() any { return new(proto.RogueArcadeLeaveCsReq) })
	c.regMsg(RogueArcadeStartCsReq, func() any { return new(proto.RogueArcadeStartCsReq) })
	c.regMsg(RogueArcadeGetInfoScRsp, func() any { return new(proto.RogueArcadeGetInfoScRsp) })
	c.regMsg(RogueArcadeRestartCsReq, func() any { return new(proto.RogueArcadeRestartCsReq) })
	c.regMsg(CommonRogueQueryCsReq, func() any { return new(proto.CommonRogueQueryCsReq) })
	c.regMsg(TakeRogueEventHandbookRewardScRsp, func() any { return new(proto.TakeRogueEventHandbookRewardScRsp) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(GetRogueExhibitionScRsp, func() any { return new(proto.GetRogueExhibitionScRsp) })
	c.regMsg(CommonRogueComponentUpdateScNotify, func() any { return new(proto.CommonRogueComponentUpdateScNotify) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(SyncRogueCommonDialogueDataScNotify, func() any { return new(proto.SyncRogueCommonDialogueDataScNotify) })
	c.regMsg(GetRogueCommonDialogueDataCsReq, func() any { return new(proto.GetRogueCommonDialogueDataCsReq) })
	c.regMsg(RogueWorkbenchHandleFuncScRsp, func() any { return new(proto.RogueWorkbenchHandleFuncScRsp) })
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(SelectRogueCommonDialogueOptionCsReq, func() any { return new(proto.SelectRogueCommonDialogueOptionCsReq) })
	c.regMsg(RogueWorkbenchGetInfoCsReq, func() any { return new(proto.RogueWorkbenchGetInfoCsReq) })
	c.regMsg(RogueDebugMessageScNotify, func() any { return new(proto.RogueDebugMessageScNotify) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(SetRogueExhibitionCsReq, func() any { return new(proto.SetRogueExhibitionCsReq) })
	c.regMsg(RogueGetGambleInfoScRsp, func() any { return new(proto.RogueGetGambleInfoScRsp) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(RogueWorkbenchGetInfoScRsp, func() any { return new(proto.RogueWorkbenchGetInfoScRsp) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(RogueWorkbenchSelectFuncCsReq, func() any { return new(proto.RogueWorkbenchSelectFuncCsReq) })
	c.regMsg(RogueWorkbenchHandleFuncCsReq, func() any { return new(proto.RogueWorkbenchHandleFuncCsReq) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoCsReq, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoCsReq) })
	c.regMsg(SelectRogueCommonDialogueOptionScRsp, func() any { return new(proto.SelectRogueCommonDialogueOptionScRsp) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(SyncRogueCommonDialogueOptionFinishScNotify, func() any { return new(proto.SyncRogueCommonDialogueOptionFinishScNotify) })
	c.regMsg(RogueGetGambleInfoCsReq, func() any { return new(proto.RogueGetGambleInfoCsReq) })
	c.regMsg(GetRogueExhibitionCsReq, func() any { return new(proto.GetRogueExhibitionCsReq) })
	c.regMsg(FinishRogueCommonDialogueScRsp, func() any { return new(proto.FinishRogueCommonDialogueScRsp) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(CommonRogueQueryScRsp, func() any { return new(proto.CommonRogueQueryScRsp) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(RogueDoGambleCsReq, func() any { return new(proto.RogueDoGambleCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(SetRogueCollectionCsReq, func() any { return new(proto.SetRogueCollectionCsReq) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(RogueDoGambleScRsp, func() any { return new(proto.RogueDoGambleScRsp) })
	c.regMsg(SetRogueExhibitionScRsp, func() any { return new(proto.SetRogueExhibitionScRsp) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(BuyRogueShopMiracleCsReq, func() any { return new(proto.BuyRogueShopMiracleCsReq) })
	c.regMsg(GetRogueCollectionCsReq, func() any { return new(proto.GetRogueCollectionCsReq) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(GetRogueCollectionScRsp, func() any { return new(proto.GetRogueCollectionScRsp) })
	c.regMsg(GetRogueCommonDialogueDataScRsp, func() any { return new(proto.GetRogueCommonDialogueDataScRsp) })
	c.regMsg(RogueWorkbenchSelectFuncScRsp, func() any { return new(proto.RogueWorkbenchSelectFuncScRsp) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(SetRogueCollectionScRsp, func() any { return new(proto.SetRogueCollectionScRsp) })
	c.regMsg(GetRogueShopBuffInfoCsReq, func() any { return new(proto.GetRogueShopBuffInfoCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(FinishRogueCommonDialogueCsReq, func() any { return new(proto.FinishRogueCommonDialogueCsReq) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(GetRogueShopMiracleInfoCsReq, func() any { return new(proto.GetRogueShopMiracleInfoCsReq) })
	c.regMsg(CommonRogueUpdateScNotify, func() any { return new(proto.CommonRogueUpdateScNotify) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(RogueMagicUnitReforgeCsReq, func() any { return new(proto.RogueMagicUnitReforgeCsReq) })
	c.regMsg(RogueMagicLeaveCsReq, func() any { return new(proto.RogueMagicLeaveCsReq) })
	c.regMsg(RogueMagicSetAutoDressInMagicUnitScRsp, func() any { return new(proto.RogueMagicSetAutoDressInMagicUnitScRsp) })
	c.regMsg(RogueMagicUnitReforgeScRsp, func() any { return new(proto.RogueMagicUnitReforgeScRsp) })
	c.regMsg(RogueMagicEnterRoomScRsp, func() any { return new(proto.RogueMagicEnterRoomScRsp) })
	c.regMsg(RogueMagicLevelInfoUpdateScNotify, func() any { return new(proto.RogueMagicLevelInfoUpdateScNotify) })
	c.regMsg(RogueMagicUnitComposeCsReq, func() any { return new(proto.RogueMagicUnitComposeCsReq) })
	c.regMsg(RogueMagicReviveAvatarCsReq, func() any { return new(proto.RogueMagicReviveAvatarCsReq) })
	c.regMsg(RogueMagicSetAutoDressInMagicUnitCsReq, func() any { return new(proto.RogueMagicSetAutoDressInMagicUnitCsReq) })
	c.regMsg(RogueMagicQueryScRsp, func() any { return new(proto.RogueMagicQueryScRsp) })
	c.regMsg(RogueMagicEnableTalentCsReq, func() any { return new(proto.RogueMagicEnableTalentCsReq) })
	c.regMsg(RogueMagicEnterCsReq, func() any { return new(proto.RogueMagicEnterCsReq) })
	c.regMsg(RogueMagicGetTalentInfoCsReq, func() any { return new(proto.RogueMagicGetTalentInfoCsReq) })
	c.regMsg(RogueMagicStartScRsp, func() any { return new(proto.RogueMagicStartScRsp) })
	c.regMsg(RogueMagicGetTalentInfoScRsp, func() any { return new(proto.RogueMagicGetTalentInfoScRsp) })
	c.regMsg(RogueMagicScepterTakeOffUnitCsReq, func() any { return new(proto.RogueMagicScepterTakeOffUnitCsReq) })
	c.regMsg(RogueMagicScepterDressInUnitCsReq, func() any { return new(proto.RogueMagicScepterDressInUnitCsReq) })
	c.regMsg(RogueMagicScepterDressInUnitScRsp, func() any { return new(proto.RogueMagicScepterDressInUnitScRsp) })
	c.regMsg(RogueMagicAreaUpdateScNotify, func() any { return new(proto.RogueMagicAreaUpdateScNotify) })
	c.regMsg(RogueMagicAutoDressInMagicUnitChangeScNotify, func() any { return new(proto.RogueMagicAutoDressInMagicUnitChangeScNotify) })
	c.regMsg(RogueMagicSettleCsReq, func() any { return new(proto.RogueMagicSettleCsReq) })
	c.regMsg(RogueMagicScepterTakeOffUnitScRsp, func() any { return new(proto.RogueMagicScepterTakeOffUnitScRsp) })
	c.regMsg(RogueMagicQueryCsReq, func() any { return new(proto.RogueMagicQueryCsReq) })
	c.regMsg(RogueMagicEnterLayerCsReq, func() any { return new(proto.RogueMagicEnterLayerCsReq) })
	c.regMsg(RogueMagicEnterLayerScRsp, func() any { return new(proto.RogueMagicEnterLayerScRsp) })
	c.regMsg(RogueMagicEnterRoomCsReq, func() any { return new(proto.RogueMagicEnterRoomCsReq) })
	c.regMsg(RogueMagicAutoDressInUnitCsReq, func() any { return new(proto.RogueMagicAutoDressInUnitCsReq) })
	c.regMsg(RogueMagicEnterScRsp, func() any { return new(proto.RogueMagicEnterScRsp) })
	c.regMsg(RogueMagicUnitComposeScRsp, func() any { return new(proto.RogueMagicUnitComposeScRsp) })
	c.regMsg(RogueMagicEnableTalentScRsp, func() any { return new(proto.RogueMagicEnableTalentScRsp) })
	c.regMsg(RogueMagicStartCsReq, func() any { return new(proto.RogueMagicStartCsReq) })
	c.regMsg(RogueMagicGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueMagicGetMiscRealTimeDataCsReq) })
	c.regMsg(RogueMagicReviveCostUpdateScNotify, func() any { return new(proto.RogueMagicReviveCostUpdateScNotify) })
	c.regMsg(RogueMagicStoryInfoUpdateScNotify, func() any { return new(proto.RogueMagicStoryInfoUpdateScNotify) })
	c.regMsg(RogueMagicLeaveScRsp, func() any { return new(proto.RogueMagicLeaveScRsp) })
	c.regMsg(RogueMagicBattleFailSettleInfoScNotify, func() any { return new(proto.RogueMagicBattleFailSettleInfoScNotify) })
	c.regMsg(RogueMagicSettleScRsp, func() any { return new(proto.RogueMagicSettleScRsp) })
	c.regMsg(RogueMagicGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueMagicGetMiscRealTimeDataScRsp) })
	c.regMsg(RogueMagicAutoDressInUnitScRsp, func() any { return new(proto.RogueMagicAutoDressInUnitScRsp) })
	c.regMsg(RogueMagicReviveAvatarScRsp, func() any { return new(proto.RogueMagicReviveAvatarScRsp) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(RogueModifierUpdateNotify, func() any { return new(proto.RogueModifierUpdateNotify) })
	c.regMsg(RogueModifierSelectCellScRsp, func() any { return new(proto.RogueModifierSelectCellScRsp) })
	c.regMsg(RogueModifierStageStartNotify, func() any { return new(proto.RogueModifierStageStartNotify) })
	c.regMsg(RogueModifierDelNotify, func() any { return new(proto.RogueModifierDelNotify) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(RogueTournDeleteArchiveScRsp, func() any { return new(proto.RogueTournDeleteArchiveScRsp) })
	c.regMsg(RogueTournEnterRoomCsReq, func() any { return new(proto.RogueTournEnterRoomCsReq) })
	c.regMsg(RogueTournStartScRsp, func() any { return new(proto.RogueTournStartScRsp) })
	c.regMsg(RogueTournEnablePermanentTalentScRsp, func() any { return new(proto.RogueTournEnablePermanentTalentScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryCsReq, func() any { return new(proto.RogueTournGetArchiveRepositoryCsReq) })
	c.regMsg(RogueTournReEnterRogueCocoonStageCsReq, func() any { return new(proto.RogueTournReEnterRogueCocoonStageCsReq) })
	c.regMsg(RogueTournResetPermanentTalentCsReq, func() any { return new(proto.RogueTournResetPermanentTalentCsReq) })
	c.regMsg(RogueTournGetSettleInfoScRsp, func() any { return new(proto.RogueTournGetSettleInfoScRsp) })
	c.regMsg(RogueTournClearArchiveNameScNotify, func() any { return new(proto.RogueTournClearArchiveNameScNotify) })
	c.regMsg(RogueTournGetAllArchiveScRsp, func() any { return new(proto.RogueTournGetAllArchiveScRsp) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneCsReq) })
	c.regMsg(RogueTournEnterRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournEnterRogueCocoonSceneCsReq) })
	c.regMsg(RogueTournQueryCsReq, func() any { return new(proto.RogueTournQueryCsReq) })
	c.regMsg(RogueTournDifficultyCompNotify, func() any { return new(proto.RogueTournDifficultyCompNotify) })
	c.regMsg(RogueTournTakeExpRewardCsReq, func() any { return new(proto.RogueTournTakeExpRewardCsReq) })
	c.regMsg(RogueTournResetPermanentTalentScRsp, func() any { return new(proto.RogueTournResetPermanentTalentScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryScRsp, func() any { return new(proto.RogueTournGetArchiveRepositoryScRsp) })
	c.regMsg(RogueTournEnterScRsp, func() any { return new(proto.RogueTournEnterScRsp) })
	c.regMsg(RogueTournTakeExpRewardScRsp, func() any { return new(proto.RogueTournTakeExpRewardScRsp) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoScRsp, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoScRsp) })
	c.regMsg(RogueTournGetSettleInfoCsReq, func() any { return new(proto.RogueTournGetSettleInfoCsReq) })
	c.regMsg(RogueTournHandBookNotify, func() any { return new(proto.RogueTournHandBookNotify) })
	c.regMsg(RogueTournLeaveScRsp, func() any { return new(proto.RogueTournLeaveScRsp) })
	c.regMsg(RogueTournWeekChallengeUpdateScNotify, func() any { return new(proto.RogueTournWeekChallengeUpdateScNotify) })
	c.regMsg(RogueTournReEnterRogueCocoonStageScRsp, func() any { return new(proto.RogueTournReEnterRogueCocoonStageScRsp) })
	c.regMsg(RogueTournBattleFailSettleInfoScNotify, func() any { return new(proto.RogueTournBattleFailSettleInfoScNotify) })
	c.regMsg(RogueTournGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueTournGetMiscRealTimeDataCsReq) })
	c.regMsg(RogueTournGetAllArchiveCsReq, func() any { return new(proto.RogueTournGetAllArchiveCsReq) })
	c.regMsg(RogueTournEnterRoomScRsp, func() any { return new(proto.RogueTournEnterRoomScRsp) })
	c.regMsg(RogueTournGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueTournGetMiscRealTimeDataScRsp) })
	c.regMsg(RogueTournSettleScRsp, func() any { return new(proto.RogueTournSettleScRsp) })
	c.regMsg(RogueTournEnterCsReq, func() any { return new(proto.RogueTournEnterCsReq) })
	c.regMsg(RogueTournLevelInfoUpdateScNotify, func() any { return new(proto.RogueTournLevelInfoUpdateScNotify) })
	c.regMsg(RogueTournEnablePermanentTalentCsReq, func() any { return new(proto.RogueTournEnablePermanentTalentCsReq) })
	c.regMsg(RogueTournSettleCsReq, func() any { return new(proto.RogueTournSettleCsReq) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoCsReq, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoCsReq) })
	c.regMsg(RogueTournGetPermanentTalentInfoCsReq, func() any { return new(proto.RogueTournGetPermanentTalentInfoCsReq) })
	c.regMsg(RogueTournGetPermanentTalentInfoScRsp, func() any { return new(proto.RogueTournGetPermanentTalentInfoScRsp) })
	c.regMsg(RogueTournExpNotify, func() any { return new(proto.RogueTournExpNotify) })
	c.regMsg(RogueTournReviveAvatarScRsp, func() any { return new(proto.RogueTournReviveAvatarScRsp) })
	c.regMsg(RogueTournConfirmSettleScRsp, func() any { return new(proto.RogueTournConfirmSettleScRsp) })
	c.regMsg(RogueTournConfirmSettleCsReq, func() any { return new(proto.RogueTournConfirmSettleCsReq) })
	c.regMsg(RogueTournDeleteArchiveCsReq, func() any { return new(proto.RogueTournDeleteArchiveCsReq) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneScRsp) })
	c.regMsg(RogueTournEnterRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournEnterRogueCocoonSceneScRsp) })
	c.regMsg(RogueTournQueryScRsp, func() any { return new(proto.RogueTournQueryScRsp) })
	c.regMsg(RogueTournLeaveCsReq, func() any { return new(proto.RogueTournLeaveCsReq) })
	c.regMsg(RogueTournRenameArchiveScRsp, func() any { return new(proto.RogueTournRenameArchiveScRsp) })
	c.regMsg(RogueTournEnterLayerScRsp, func() any { return new(proto.RogueTournEnterLayerScRsp) })
	c.regMsg(RogueTournStartCsReq, func() any { return new(proto.RogueTournStartCsReq) })
	c.regMsg(RogueTournReviveAvatarCsReq, func() any { return new(proto.RogueTournReviveAvatarCsReq) })
	c.regMsg(RogueTournRenameArchiveCsReq, func() any { return new(proto.RogueTournRenameArchiveCsReq) })
	c.regMsg(RogueTournAreaUpdateScNotify, func() any { return new(proto.RogueTournAreaUpdateScNotify) })
	c.regMsg(RogueTournReviveCostUpdateScNotify, func() any { return new(proto.RogueTournReviveCostUpdateScNotify) })
	c.regMsg(RogueTournEnterLayerCsReq, func() any { return new(proto.RogueTournEnterLayerCsReq) })
	c.regMsg(TakeRollShopRewardCsReq, func() any { return new(proto.TakeRollShopRewardCsReq) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(ChangePropTimelineInfoCsReq, func() any { return new(proto.ChangePropTimelineInfoCsReq) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(UpdateGroupPropertyScRsp, func() any { return new(proto.UpdateGroupPropertyScRsp) })
	c.regMsg(SpringRefreshCsReq, func() any { return new(proto.SpringRefreshCsReq) })
	c.regMsg(SceneReviveAfterRebattleCsReq, func() any { return new(proto.SceneReviveAfterRebattleCsReq) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(UpdateGroupPropertyCsReq, func() any { return new(proto.UpdateGroupPropertyCsReq) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(SceneReviveAfterRebattleScRsp, func() any { return new(proto.SceneReviveAfterRebattleScRsp) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(SpringRefreshScRsp, func() any { return new(proto.SpringRefreshScRsp) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(OpenChestScNotify, func() any { return new(proto.OpenChestScNotify) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(ChangePropTimelineInfoScRsp, func() any { return new(proto.ChangePropTimelineInfoScRsp) })
	c.regMsg(EnteredSceneChangeScNotify, func() any { return new(proto.EnteredSceneChangeScNotify) })
	c.regMsg(LastSpringRefreshTimeNotify, func() any { return new(proto.LastSpringRefreshTimeNotify) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(SpaceZooExchangeItemCsReq, func() any { return new(proto.SpaceZooExchangeItemCsReq) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(SpaceZooExchangeItemScRsp, func() any { return new(proto.SpaceZooExchangeItemScRsp) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(SpaceZooTakeCsReq, func() any { return new(proto.SpaceZooTakeCsReq) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(SpaceZooTakeScRsp, func() any { return new(proto.SpaceZooTakeScRsp) })
	c.regMsg(GetStarFightDataCsReq, func() any { return new(proto.GetStarFightDataCsReq) })
	c.regMsg(StarFightDataChangeNotify, func() any { return new(proto.StarFightDataChangeNotify) })
	c.regMsg(StartStarFightLevelCsReq, func() any { return new(proto.StartStarFightLevelCsReq) })
	c.regMsg(GetStarFightDataScRsp, func() any { return new(proto.GetStarFightDataScRsp) })
	c.regMsg(StartStarFightLevelScRsp, func() any { return new(proto.StartStarFightLevelScRsp) })
	c.regMsg(GetStoryLineInfoCsReq, func() any { return new(proto.GetStoryLineInfoCsReq) })
	c.regMsg(StoryLineTrialAvatarChangeScNotify, func() any { return new(proto.StoryLineTrialAvatarChangeScNotify) })
	c.regMsg(ChangeStoryLineFinishScNotify, func() any { return new(proto.ChangeStoryLineFinishScNotify) })
	c.regMsg(StoryLineInfoScNotify, func() any { return new(proto.StoryLineInfoScNotify) })
	c.regMsg(GetStoryLineInfoScRsp, func() any { return new(proto.GetStoryLineInfoScRsp) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(StrongChallengeActivityBattleEndScNotify, func() any { return new(proto.StrongChallengeActivityBattleEndScNotify) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(GetSummonActivityDataCsReq, func() any { return new(proto.GetSummonActivityDataCsReq) })
	c.regMsg(EnterSummonActivityStageCsReq, func() any { return new(proto.EnterSummonActivityStageCsReq) })
	c.regMsg(SummonActivityBattleEndScNotify, func() any { return new(proto.SummonActivityBattleEndScNotify) })
	c.regMsg(GetSummonActivityDataScRsp, func() any { return new(proto.GetSummonActivityDataScRsp) })
	c.regMsg(EnterSummonActivityStageScRsp, func() any { return new(proto.EnterSummonActivityStageScRsp) })
	c.regMsg(SwitchHandDataCsReq, func() any { return new(proto.SwitchHandDataCsReq) })
	c.regMsg(SwitchHandFinishScRsp, func() any { return new(proto.SwitchHandFinishScRsp) })
	c.regMsg(SwitchHandCoinUpdateCsReq, func() any { return new(proto.SwitchHandCoinUpdateCsReq) })
	c.regMsg(SwitchHandDataScRsp, func() any { return new(proto.SwitchHandDataScRsp) })
	c.regMsg(SwitchHandFinishCsReq, func() any { return new(proto.SwitchHandFinishCsReq) })
	c.regMsg(SwitchHandUpdateScRsp, func() any { return new(proto.SwitchHandUpdateScRsp) })
	c.regMsg(SwitchHandStartScRsp, func() any { return new(proto.SwitchHandStartScRsp) })
	c.regMsg(SwitchHandStartCsReq, func() any { return new(proto.SwitchHandStartCsReq) })
	c.regMsg(SwitchHandUpdateCsReq, func() any { return new(proto.SwitchHandUpdateCsReq) })
	c.regMsg(SwitchHandCoinUpdateScRsp, func() any { return new(proto.SwitchHandCoinUpdateScRsp) })
	c.regMsg(SwordTrainingGiveUpGameScRsp, func() any { return new(proto.SwordTrainingGiveUpGameScRsp) })
	c.regMsg(EnterSwordTrainingExamScRsp, func() any { return new(proto.EnterSwordTrainingExamScRsp) })
	c.regMsg(SwordTrainingGiveUpGameCsReq, func() any { return new(proto.SwordTrainingGiveUpGameCsReq) })
	c.regMsg(SwordTrainingResumeGameScRsp, func() any { return new(proto.SwordTrainingResumeGameScRsp) })
	c.regMsg(SwordTrainingSelectEndingCsReq, func() any { return new(proto.SwordTrainingSelectEndingCsReq) })
	c.regMsg(SwordTrainingGameSyncChangeScNotify, func() any { return new(proto.SwordTrainingGameSyncChangeScNotify) })
	c.regMsg(SwordTrainingDialogueSelectOptionScRsp, func() any { return new(proto.SwordTrainingDialogueSelectOptionScRsp) })
	c.regMsg(SwordTrainingStoryBattleScRsp, func() any { return new(proto.SwordTrainingStoryBattleScRsp) })
	c.regMsg(SwordTrainingStoryConfirmCsReq, func() any { return new(proto.SwordTrainingStoryConfirmCsReq) })
	c.regMsg(SwordTrainingGameSettleScNotify, func() any { return new(proto.SwordTrainingGameSettleScNotify) })
	c.regMsg(SwordTrainingDailyPhaseConfirmScRsp, func() any { return new(proto.SwordTrainingDailyPhaseConfirmScRsp) })
	c.regMsg(SwordTrainingDailyPhaseConfirmCsReq, func() any { return new(proto.SwordTrainingDailyPhaseConfirmCsReq) })
	c.regMsg(SwordTrainingExamResultConfirmCsReq, func() any { return new(proto.SwordTrainingExamResultConfirmCsReq) })
	c.regMsg(SwordTrainingTurnActionCsReq, func() any { return new(proto.SwordTrainingTurnActionCsReq) })
	c.regMsg(SwordTrainingMarkEndingViewedCsReq, func() any { return new(proto.SwordTrainingMarkEndingViewedCsReq) })
	c.regMsg(SwordTrainingStoryBattleCsReq, func() any { return new(proto.SwordTrainingStoryBattleCsReq) })
	c.regMsg(SwordTrainingStartGameScRsp, func() any { return new(proto.SwordTrainingStartGameScRsp) })
	c.regMsg(SwordTrainingRestoreGameCsReq, func() any { return new(proto.SwordTrainingRestoreGameCsReq) })
	c.regMsg(SwordTrainingExamResultConfirmScRsp, func() any { return new(proto.SwordTrainingExamResultConfirmScRsp) })
	c.regMsg(SwordTrainingSetSkillTraceCsReq, func() any { return new(proto.SwordTrainingSetSkillTraceCsReq) })
	c.regMsg(EnterSwordTrainingExamCsReq, func() any { return new(proto.EnterSwordTrainingExamCsReq) })
	c.regMsg(SwordTrainingActionTurnSettleScNotify, func() any { return new(proto.SwordTrainingActionTurnSettleScNotify) })
	c.regMsg(GetSwordTrainingDataCsReq, func() any { return new(proto.GetSwordTrainingDataCsReq) })
	c.regMsg(GetSwordTrainingDataScRsp, func() any { return new(proto.GetSwordTrainingDataScRsp) })
	c.regMsg(SwordTrainingLearnSkillCsReq, func() any { return new(proto.SwordTrainingLearnSkillCsReq) })
	c.regMsg(SwordTrainingTurnActionScRsp, func() any { return new(proto.SwordTrainingTurnActionScRsp) })
	c.regMsg(SwordTrainingLearnSkillScRsp, func() any { return new(proto.SwordTrainingLearnSkillScRsp) })
	c.regMsg(SwordTrainingStartGameCsReq, func() any { return new(proto.SwordTrainingStartGameCsReq) })
	c.regMsg(SwordTrainingDialogueSelectOptionCsReq, func() any { return new(proto.SwordTrainingDialogueSelectOptionCsReq) })
	c.regMsg(SwordTrainingSelectEndingScRsp, func() any { return new(proto.SwordTrainingSelectEndingScRsp) })
	c.regMsg(SwordTrainingUnlockSyncScNotify, func() any { return new(proto.SwordTrainingUnlockSyncScNotify) })
	c.regMsg(SwordTrainingStoryConfirmScRsp, func() any { return new(proto.SwordTrainingStoryConfirmScRsp) })
	c.regMsg(SwordTrainingMarkEndingViewedScRsp, func() any { return new(proto.SwordTrainingMarkEndingViewedScRsp) })
	c.regMsg(SwordTrainingSetSkillTraceScRsp, func() any { return new(proto.SwordTrainingSetSkillTraceScRsp) })
	c.regMsg(SwordTrainingResumeGameCsReq, func() any { return new(proto.SwordTrainingResumeGameCsReq) })
	c.regMsg(SwordTrainingRestoreGameScRsp, func() any { return new(proto.SwordTrainingRestoreGameScRsp) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(TarotBookModifyEnergyScNotify, func() any { return new(proto.TarotBookModifyEnergyScNotify) })
	c.regMsg(TarotBookFinishInteractionCsReq, func() any { return new(proto.TarotBookFinishInteractionCsReq) })
	c.regMsg(TarotBookFinishInteractionScRsp, func() any { return new(proto.TarotBookFinishInteractionScRsp) })
	c.regMsg(TarotBookGetDataScRsp, func() any { return new(proto.TarotBookGetDataScRsp) })
	c.regMsg(TarotBookOpenPackScRsp, func() any { return new(proto.TarotBookOpenPackScRsp) })
	c.regMsg(TarotBookGetDataCsReq, func() any { return new(proto.TarotBookGetDataCsReq) })
	c.regMsg(TarotBookUnlockStoryCsReq, func() any { return new(proto.TarotBookUnlockStoryCsReq) })
	c.regMsg(TarotBookFinishStoryCsReq, func() any { return new(proto.TarotBookFinishStoryCsReq) })
	c.regMsg(TarotBookOpenPackCsReq, func() any { return new(proto.TarotBookOpenPackCsReq) })
	c.regMsg(TarotBookUnlockStoryScRsp, func() any { return new(proto.TarotBookUnlockStoryScRsp) })
	c.regMsg(TarotBookFinishStoryScRsp, func() any { return new(proto.TarotBookFinishStoryScRsp) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(GetTrackPhotoActivityDataCsReq, func() any { return new(proto.GetTrackPhotoActivityDataCsReq) })
	c.regMsg(GetTrackPhotoActivityDataScRsp, func() any { return new(proto.GetTrackPhotoActivityDataScRsp) })
	c.regMsg(QuitTrackPhotoStageScRsp, func() any { return new(proto.QuitTrackPhotoStageScRsp) })
	c.regMsg(StartTrackPhotoStageCsReq, func() any { return new(proto.StartTrackPhotoStageCsReq) })
	c.regMsg(SettleTrackPhotoStageCsReq, func() any { return new(proto.SettleTrackPhotoStageCsReq) })
	c.regMsg(SettleTrackPhotoStageScRsp, func() any { return new(proto.SettleTrackPhotoStageScRsp) })
	c.regMsg(QuitTrackPhotoStageCsReq, func() any { return new(proto.QuitTrackPhotoStageCsReq) })
	c.regMsg(StartTrackPhotoStageScRsp, func() any { return new(proto.StartTrackPhotoStageScRsp) })
	c.regMsg(TrainPartyBuildStartStepScRsp, func() any { return new(proto.TrainPartyBuildStartStepScRsp) })
	c.regMsg(TrainPartyGamePlayStartCsReq, func() any { return new(proto.TrainPartyGamePlayStartCsReq) })
	c.regMsg(TrainPartyAddBuildDynamicBuffScRsp, func() any { return new(proto.TrainPartyAddBuildDynamicBuffScRsp) })
	c.regMsg(TrainPartyTakeBuildLevelAwardScRsp, func() any { return new(proto.TrainPartyTakeBuildLevelAwardScRsp) })
	c.regMsg(TrainPartyBuildStartStepCsReq, func() any { return new(proto.TrainPartyBuildStartStepCsReq) })
	c.regMsg(TrainPartyGetDataCsReq, func() any { return new(proto.TrainPartyGetDataCsReq) })
	c.regMsg(TrainPartyBuildingUpdateNotify, func() any { return new(proto.TrainPartyBuildingUpdateNotify) })
	c.regMsg(TrainPartyBuildDiyScRsp, func() any { return new(proto.TrainPartyBuildDiyScRsp) })
	c.regMsg(TrainPartyEnterCsReq, func() any { return new(proto.TrainPartyEnterCsReq) })
	c.regMsg(TrainPartyUseCardCsReq, func() any { return new(proto.TrainPartyUseCardCsReq) })
	c.regMsg(TrainPartyMoveScNotify, func() any { return new(proto.TrainPartyMoveScNotify) })
	c.regMsg(TrainPartySettleNotify, func() any { return new(proto.TrainPartySettleNotify) })
	c.regMsg(TrainPartyEnterScRsp, func() any { return new(proto.TrainPartyEnterScRsp) })
	c.regMsg(TrainPartyTakeBuildLevelAwardCsReq, func() any { return new(proto.TrainPartyTakeBuildLevelAwardCsReq) })
	c.regMsg(TrainPartySyncUpdateScNotify, func() any { return new(proto.TrainPartySyncUpdateScNotify) })
	c.regMsg(TrainPartyLeaveCsReq, func() any { return new(proto.TrainPartyLeaveCsReq) })
	c.regMsg(TrainPartyGamePlayStartScRsp, func() any { return new(proto.TrainPartyGamePlayStartScRsp) })
	c.regMsg(TrainPartyBuildDiyCsReq, func() any { return new(proto.TrainPartyBuildDiyCsReq) })
	c.regMsg(TrainPartyUseCardScRsp, func() any { return new(proto.TrainPartyUseCardScRsp) })
	c.regMsg(TrainPartyHandlePendingActionScRsp, func() any { return new(proto.TrainPartyHandlePendingActionScRsp) })
	c.regMsg(TrainPartyGamePlaySettleNotify, func() any { return new(proto.TrainPartyGamePlaySettleNotify) })
	c.regMsg(TrainPartyHandlePendingActionCsReq, func() any { return new(proto.TrainPartyHandlePendingActionCsReq) })
	c.regMsg(TrainPartyAddBuildDynamicBuffCsReq, func() any { return new(proto.TrainPartyAddBuildDynamicBuffCsReq) })
	c.regMsg(TrainPartyLeaveScRsp, func() any { return new(proto.TrainPartyLeaveScRsp) })
	c.regMsg(TrainPartyGetDataScRsp, func() any { return new(proto.TrainPartyGetDataScRsp) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(TrainVisitorBehaviorFinishCsReq, func() any { return new(proto.TrainVisitorBehaviorFinishCsReq) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(GetTrainVisitorBehaviorCsReq, func() any { return new(proto.GetTrainVisitorBehaviorCsReq) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
	c.regMsg(WolfBroGameActivateBulletScRsp, func() any { return new(proto.WolfBroGameActivateBulletScRsp) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(WolfBroGameExplodeMonsterCsReq, func() any { return new(proto.WolfBroGameExplodeMonsterCsReq) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(WolfBroGameExplodeMonsterScRsp, func() any { return new(proto.WolfBroGameExplodeMonsterScRsp) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(WorldUnlockScRsp, func() any { return new(proto.WorldUnlockScRsp) })
	c.regMsg(WorldUnlockCsReq, func() any { return new(proto.WorldUnlockCsReq) })
}
