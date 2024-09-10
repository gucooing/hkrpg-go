package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

const (
	GetTrialActivityDataCsReq                          = 2653
	GetLoginActivityScRsp                              = 2606
	StartTrialActivityCsReq                            = 2603
	GetTrialActivityDataScRsp                          = 2658
	EnterTrialActivityStageCsReq                       = 2659
	SubmitMonsterResearchActivityMaterialCsReq         = 2625
	EnterTrialActivityStageScRsp                       = 2607
	GetMonsterResearchActivityDataScRsp                = 2698
	TakeMonsterResearchActivityRewardCsReq             = 2617
	StartTrialActivityScRsp                            = 2626
	TakeLoginActivityRewardCsReq                       = 2611
	TakeMonsterResearchActivityRewardScRsp             = 2692
	TrialActivityDataChangeScNotify                    = 2647
	GetActivityScheduleConfigCsReq                     = 2656
	GetMonsterResearchActivityDataCsReq                = 2697
	TakeTrialActivityRewardCsReq                       = 2641
	LeaveTrialActivityCsReq                            = 2613
	TakeLoginActivityRewardScRsp                       = 2643
	LeaveTrialActivityScRsp                            = 2614
	CurTrialActivityScNotify                           = 2622
	GetLoginActivityCsReq                              = 2668
	SubmitMonsterResearchActivityMaterialScRsp         = 2624
	TakeTrialActivityRewardScRsp                       = 2631
	GetActivityScheduleConfigScRsp                     = 2693
	EnterAdventureScRsp                                = 1306
	GetFarmStageGachaInfoCsReq                         = 1311
	EnterAdventureCsReq                                = 1368
	GetFarmStageGachaInfoScRsp                         = 1343
	EnterAetherDivideSceneCsReq                        = 4868
	AetherDivideRefreshEndlessScRsp                    = 4876
	AetherDivideTakeChallengeRewardCsReq               = 4845
	StartAetherDivideSceneBattleCsReq                  = 4856
	GetAetherDivideInfoScRsp                           = 4854
	AetherDivideLineupScNotify                         = 4823
	GetAetherDivideChallengeInfoScRsp                  = 4808
	GetAetherDivideChallengeInfoCsReq                  = 4888
	SwitchAetherDivideLineUpSlotScRsp                  = 4824
	AetherDivideSpiritExpUpScRsp                       = 4828
	SetAetherDivideLineUpScRsp                         = 4848
	GetAetherDivideInfoCsReq                           = 4899
	StartAetherDivideStageBattleScRsp                  = 4892
	ClearAetherDividePassiveSkillScRsp                 = 4898
	StartAetherDivideChallengeBattleScRsp              = 4829
	StartAetherDivideChallengeBattleCsReq              = 4839
	AetherDivideSkillItemScNotify                      = 4816
	AetherDivideTainerInfoScNotify                     = 4832
	StartAetherDivideSceneBattleScRsp                  = 4893
	SwitchAetherDivideLineUpSlotCsReq                  = 4825
	EquipAetherDividePassiveSkillCsReq                 = 4882
	AetherDivideFinishChallengeScNotify                = 4881
	AetherDivideRefreshEndlessCsReq                    = 4878
	ClearAetherDividePassiveSkillCsReq                 = 4897
	LeaveAetherDivideSceneCsReq                        = 4811
	EnterAetherDivideSceneScRsp                        = 4806
	SetAetherDivideLineUpCsReq                         = 4865
	AetherDivideTakeChallengeRewardScRsp               = 4866
	AetherDivideRefreshEndlessScNotify                 = 4849
	EquipAetherDividePassiveSkillScRsp                 = 4834
	AetherDivideSpiritInfoScNotify                     = 4809
	AetherDivideSpiritExpUpCsReq                       = 4890
	StartAetherDivideStageBattleCsReq                  = 4817
	LeaveAetherDivideSceneScRsp                        = 4843
	AlleyShipUsedCountScNotify                         = 4723
	AlleyGuaranteedFundsCsReq                          = 4778
	TakePrestigeRewardCsReq                            = 4799
	RefreshAlleyOrderCsReq                             = 4797
	AlleyEventEffectNotify                             = 4733
	LogisticsGameCsReq                                 = 4711
	AlleyPlacingGameScRsp                              = 4748
	AlleyEventChangeNotify                             = 4751
	GetSaveLogisticsMapCsReq                           = 4716
	LogisticsDetonateStarSkiffCsReq                    = 4766
	AlleyFundsScNotify                                 = 4790
	PrestigeLevelUpCsReq                               = 4717
	LogisticsDetonateStarSkiffScRsp                    = 4746
	AlleyTakeEventRewardScRsp                          = 4745
	AlleyGuaranteedFundsScRsp                          = 4776
	AlleyShipUnlockScNotify                            = 4709
	AlleyTakeEventRewardCsReq                          = 4749
	TakePrestigeRewardScRsp                            = 4754
	LogisticsGameScRsp                                 = 4743
	StartAlleyEventCsReq                               = 4739
	LogisticsInfoScNotify                              = 4781
	AlleyShipmentEventEffectsScNotify                  = 4732
	LogisticsScoreRewardSyncInfoScNotify               = 4712
	AlleyOrderChangedScNotify                          = 4725
	ActivityRaidPlacingGameCsReq                       = 4730
	StartAlleyEventScRsp                               = 4729
	PrestigeLevelUpScRsp                               = 4792
	AlleyShopLevelScNotify                             = 4728
	GetAlleyInfoCsReq                                  = 4768
	GetSaveLogisticsMapScRsp                           = 4784
	RefreshAlleyOrderScRsp                             = 4798
	GetAlleyInfoScRsp                                  = 4706
	SaveLogisticsScRsp                                 = 4708
	AlleyPlacingGameCsReq                              = 4765
	SaveLogisticsCsReq                                 = 4788
	ActivityRaidPlacingGameScRsp                       = 4795
	GetArchiveDataScRsp                                = 2306
	GetUpdatedArchiveDataScRsp                         = 2343
	GetArchiveDataCsReq                                = 2368
	GetUpdatedArchiveDataCsReq                         = 2311
	TakeOffEquipmentCsReq                              = 399
	DressRelicAvatarCsReq                              = 334
	AvatarExpUpScRsp                                   = 343
	TakeOffEquipmentScRsp                              = 354
	AddAvatarScNotify                                  = 365
	MarkAvatarScRsp                                    = 381
	AvatarExpUpCsReq                                   = 311
	RankUpAvatarScRsp                                  = 382
	PromoteAvatarCsReq                                 = 339
	DressAvatarScRsp                                   = 333
	UnlockSkilltreeScRsp                               = 393
	DressAvatarSkinScRsp                               = 390
	PromoteAvatarScRsp                                 = 329
	TakeOffRelicScRsp                                  = 325
	RankUpAvatarCsReq                                  = 348
	DressRelicAvatarScRsp                              = 397
	TakeOffAvatarSkinCsReq                             = 328
	DressAvatarSkinCsReq                               = 392
	TakeOffRelicCsReq                                  = 398
	MarkAvatarCsReq                                    = 308
	TakePromotionRewardScRsp                           = 317
	TakeOffAvatarSkinScRsp                             = 309
	GetAvatarDataScRsp                                 = 306
	TakePromotionRewardCsReq                           = 324
	GetAvatarDataCsReq                                 = 368
	UnlockAvatarSkinScNotify                           = 388
	DressAvatarCsReq                                   = 351
	UnlockSkilltreeCsReq                               = 356
	SyncClientResVersionCsReq                          = 139
	GetCurBattleInfoScRsp                              = 193
	ReBattleAfterBattleLoseCsNotify                    = 165
	ServerSimulateBattleFinishScNotify                 = 154
	QuitBattleScRsp                                    = 143
	QuitBattleCsReq                                    = 111
	GetCurBattleInfoCsReq                              = 156
	PVEBattleResultCsReq                               = 168
	BattleLogReportScRsp                               = 199
	QuitBattleScNotify                                 = 151
	RebattleByClientCsNotify                           = 148
	BattleLogReportCsReq                               = 133
	SyncClientResVersionScRsp                          = 129
	PVEBattleResultScRsp                               = 106
	GetBattleCollegeDataScRsp                          = 5706
	GetBattleCollegeDataCsReq                          = 5768
	StartBattleCollegeScRsp                            = 5756
	BattleCollegeDataChangeScNotify                    = 5711
	StartBattleCollegeCsReq                            = 5743
	TakeAllRewardCsReq                                 = 3029
	BattlePassInfoNotify                               = 3068
	TakeAllRewardScRsp                                 = 3051
	TakeBpRewardCsReq                                  = 3043
	BuyBpLevelScRsp                                    = 3039
	BuyBpLevelCsReq                                    = 3093
	TakeBpRewardScRsp                                  = 3056
	BoxingClubRewardScNotify                           = 4251
	ChooseBoxingClubResonanceScRsp                     = 4254
	ChooseBoxingClubStageOptionalBuffCsReq             = 4282
	GetBoxingClubInfoScRsp                             = 4206
	SetBoxingClubResonanceLineupScRsp                  = 4248
	GiveUpBoxingClubChallengeScRsp                     = 4229
	GetBoxingClubInfoCsReq                             = 4268
	MatchBoxingClubOpponentCsReq                       = 4211
	StartBoxingClubBattleScRsp                         = 4293
	ChooseBoxingClubStageOptionalBuffScRsp             = 4234
	ChooseBoxingClubResonanceCsReq                     = 4299
	BoxingClubChallengeUpdateScNotify                  = 4233
	MatchBoxingClubOpponentScRsp                       = 4243
	StartBoxingClubBattleCsReq                         = 4256
	GiveUpBoxingClubChallengeCsReq                     = 4239
	SetBoxingClubResonanceLineupCsReq                  = 4265
	GetCurChallengeCsReq                               = 1733
	TakeChallengeRewardScRsp                           = 1734
	ChallengeSettleNotify                              = 1739
	StartChallengeCsReq                                = 1711
	GetChallengeGroupStatisticsCsReq                   = 1797
	GetChallengeCsReq                                  = 1768
	TakeChallengeRewardCsReq                           = 1782
	StartPartialChallengeCsReq                         = 1725
	StartChallengeScRsp                                = 1743
	EnterChallengeNextPhaseScRsp                       = 1709
	LeaveChallengeScRsp                                = 1793
	GetCurChallengeScRsp                               = 1799
	GetChallengeScRsp                                  = 1706
	RestartChallengePhaseCsReq                         = 1792
	ChallengeLineupNotify                              = 1754
	StartPartialChallengeScRsp                         = 1724
	EnterChallengeNextPhaseCsReq                       = 1728
	RestartChallengePhaseScRsp                         = 1790
	LeaveChallengeCsReq                                = 1756
	GetChallengeGroupStatisticsScRsp                   = 1798
	ChallengeBossPhaseSettleNotify                     = 1788
	GetChatFriendHistoryScRsp                          = 3929
	GetChatFriendHistoryCsReq                          = 3939
	SendMsgScRsp                                       = 3906
	GetChatEmojiListCsReq                              = 3951
	SendMsgCsReq                                       = 3968
	PrivateMsgOfflineUsersScNotify                     = 3943
	GetLoginChatInfoScRsp                              = 3934
	RevcMsgScNotify                                    = 3911
	GetChatEmojiListScRsp                              = 3933
	BatchMarkChatEmojiCsReq                            = 3965
	MarkChatEmojiCsReq                                 = 3999
	MarkChatEmojiScRsp                                 = 3954
	GetLoginChatInfoCsReq                              = 3982
	GetPrivateChatHistoryScRsp                         = 3993
	GetPrivateChatHistoryCsReq                         = 3956
	BatchMarkChatEmojiScRsp                            = 3948
	ChessRoguePickAvatarCsReq                          = 5446
	EnhanceChessRogueBuffCsReq                         = 5458
	GetChessRogueBuffEnhanceInfoScRsp                  = 5476
	ChessRogueNousGetRogueTalentInfoScRsp              = 5482
	SelectChessRogueNousSubStoryCsReq                  = 5427
	FinishChessRogueSubStoryScRsp                      = 5491
	GetChessRogueBuffEnhanceInfoCsReq                  = 5555
	ChessRogueQuitScRsp                                = 5444
	ChessRogueSelectBpScRsp                            = 5566
	ChessRogueQuestFinishNotify                        = 5460
	ChessRogueStartCsReq                               = 5514
	ChessRogueRollDiceCsReq                            = 5529
	ChessRogueEnterScRsp                               = 5421
	ChessRogueCheatRollScRsp                           = 5499
	ChessRogueReRollDiceCsReq                          = 5474
	ChessRogueSkipTeachingLevelScRsp                   = 5531
	ChessRogueLayerAccountInfoNotify                   = 5591
	ChessRogueReviveAvatarScRsp                        = 5481
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5463
	ChessRogueUpdateUnlockLevelScNotify                = 5430
	GetChessRogueStoryInfoScRsp                        = 5462
	ChessRogueMoveCellNotify                           = 5419
	ChessRogueLeaveScRsp                               = 5447
	FinishChessRogueSubStoryCsReq                      = 5418
	ChessRogueQueryBpCsReq                             = 5401
	ChessRogueConfirmRollCsReq                         = 5523
	ChessRogueNousEnableRogueTalentScRsp               = 5576
	GetChessRogueNousStoryInfoScRsp                    = 5557
	ChessRogueGiveUpScRsp                              = 5456
	ChessRogueUpdateReviveInfoScNotify                 = 5434
	ChessRogueSelectCellCsReq                          = 5549
	ChessRogueUpdateAllowedSelectCellScNotify          = 5426
	ChessRogueUpdateBoardScNotify                      = 5518
	ChessRogueNousEditDiceScRsp                        = 5565
	GetChessRogueNousStoryInfoCsReq                    = 5501
	GetChessRogueStoryAeonTalkInfoScRsp                = 5544
	ChessRogueConfirmRollScRsp                         = 5596
	FinishChessRogueNousSubStoryScRsp                  = 5577
	EnhanceChessRogueBuffScRsp                         = 5552
	SelectChessRogueNousSubStoryScRsp                  = 5521
	ChessRogueQueryBpScRsp                             = 5588
	ChessRogueReRollDiceScRsp                          = 5586
	ChessRogueReviveAvatarCsReq                        = 5569
	GetChessRogueStoryAeonTalkInfoCsReq                = 5417
	ChessRogueGoAheadCsReq                             = 5405
	ChessRogueUpdateLevelBaseInfoScNotify              = 5431
	ChessRogueQueryCsReq                               = 5493
	ChessRogueEnterCellScRsp                           = 5433
	EnterChessRogueAeonRoomScRsp                       = 5589
	SyncChessRogueNousMainStoryScNotify                = 5455
	ChessRogueGoAheadScRsp                             = 5556
	ChessRogueEnterNextLayerScRsp                      = 5492
	ChessRogueChangeyAeonDimensionNotify               = 5425
	ChessRogueStartScRsp                               = 5600
	ChessRogueLeaveCsReq                               = 5445
	FinishChessRogueNousSubStoryCsReq                  = 5409
	ChessRogueGiveUpRollScRsp                          = 5489
	ChessRogueUpdateAeonModifierValueScNotify          = 5542
	ChessRogueNousEditDiceCsReq                        = 5459
	ChessRoguePickAvatarScRsp                          = 5450
	ChessRogueNousDiceUpdateNotify                     = 5508
	ChessRogueSelectCellScRsp                          = 5532
	ChessRogueCheatRollCsReq                           = 5408
	GetChessRogueStoryInfoCsReq                        = 5497
	SyncChessRogueMainStoryFinishScNotify              = 5480
	ChessRogueGiveUpCsReq                              = 5585
	ChessRogueQuitCsReq                                = 5506
	ChessRogueSelectBpCsReq                            = 5486
	EnterChessRogueAeonRoomCsReq                       = 5520
	ChessRogueEnterCsReq                               = 5567
	ChessRogueUpdateDiceInfoScNotify                   = 5461
	ChessRogueQueryAeonDimensionsScRsp                 = 5536
	ChessRogueRollDiceScRsp                            = 5539
	SyncChessRogueNousSubStoryScNotify                 = 5484
	ChessRogueQueryScRsp                               = 5507
	ChessRogueQueryAeonDimensionsCsReq                 = 5413
	ChessRogueUpdateActionPointScNotify                = 5487
	ChessRogueEnterNextLayerCsReq                      = 5568
	SelectChessRogueSubStoryScRsp                      = 5587
	ChessRogueGiveUpRollCsReq                          = 5545
	SyncChessRogueNousValueScNotify                    = 5451
	ChessRogueNousEnableRogueTalentCsReq               = 5571
	ChessRogueNousGetRogueTalentInfoCsReq              = 5592
	ChessRogueFinishCurRoomNotify                      = 5467
	SelectChessRogueSubStoryCsReq                      = 5406
	ChessRogueSkipTeachingLevelCsReq                   = 5477
	ChessRogueUpdateMoneyInfoScNotify                  = 5442
	ChessRogueNousDiceSurfaceUnlockNotify              = 5453
	ChessRogueCellUpdateNotify                         = 5498
	ChessRogueEnterCellCsReq                           = 5490
	ClockParkHandleWaitOperationScRsp                  = 7245
	ClockParkQuitScriptCsReq                           = 7206
	ClockParkQuitScriptScRsp                           = 7212
	ClockParkStartScriptCsReq                          = 7223
	ClockParkGetOngoingScriptInfoCsReq                 = 7234
	ClockParkGetInfoScRsp                              = 7224
	ClockParkUnlockTalentScRsp                         = 7208
	ClockParkFinishScriptScNotify                      = 7225
	ClockParkBattleEndScNotify                         = 7248
	ClockParkGetInfoCsReq                              = 7207
	ClockParkUseBuffScRsp                              = 7235
	ClockParkHandleWaitOperationCsReq                  = 7236
	ClockParkStartScriptScRsp                          = 7204
	ClockParkGetOngoingScriptInfoScRsp                 = 7220
	ClockParkUnlockTalentCsReq                         = 7211
	ClockParkUseBuffCsReq                              = 7203
	ContentPackageGetDataScRsp                         = 7524
	ContentPackageUnlockCsReq                          = 7517
	ContentPackageGetDataCsReq                         = 7507
	ContentPackageTransferScNotify                     = 7508
	ContentPackageUnlockScRsp                          = 7511
	ContentPackageSyncDataScNotify                     = 7515
	TakeApRewardScRsp                                  = 3306
	TakeAllApRewardCsReq                               = 3393
	DailyActiveInfoNotify                              = 3356
	TakeAllApRewardScRsp                               = 3339
	TakeApRewardCsReq                                  = 3368
	GetDailyActiveInfoScRsp                            = 3343
	GetDailyActiveInfoCsReq                            = 3311
	MakeMissionDrinkScRsp                              = 6987
	MakeMissionDrinkCsReq                              = 6996
	DrinkMakerDayEndScNotify                           = 6988
	MakeDrinkCsReq                                     = 6992
	GetDrinkMakerDataCsReq                             = 6999
	EndDrinkMakerSequenceScRsp                         = 6982
	MakeDrinkScRsp                                     = 7000
	DrinkMakerChallengeCsReq                           = 6993
	DrinkMakerUpdateTipsNotify                         = 6981
	GetDrinkMakerDataScRsp                             = 6983
	EndDrinkMakerSequenceCsReq                         = 6984
	DrinkMakerChallengeScRsp                           = 6990
	EvolveBuildShopAbilityUpScRsp                      = 7148
	EvolveBuildShopAbilityUpCsReq                      = 7105
	EvolveBuildCoinNotify                              = 7127
	EvolveBuildLeaveScRsp                              = 7120
	EvolveBuildStartStageCsReq                         = 7111
	EvolveBuildUnlockInfoNotify                        = 7149
	EvolveBuildStartLevelScRsp                         = 7117
	EvolveBuildReRandomStageCsReq                      = 7145
	EvolveBuildShopAbilityResetCsReq                   = 7133
	EvolveBuildLeaveCsReq                              = 7134
	EvolveBuildTakeExpRewardCsReq                      = 7135
	EvolveBuildQueryInfoCsReq                          = 7107
	EvolveBuildReRandomStageScRsp                      = 7106
	EvolveBuildTakeExpRewardScRsp                      = 7125
	EvolveBuildShopAbilityDownScRsp                    = 7103
	EvolveBuildStartLevelCsReq                         = 7115
	EvolveBuildFinishScNotify                          = 7136
	EvolveBuildStartStageScRsp                         = 7108
	EvolveBuildGiveupCsReq                             = 7123
	EvolveBuildQueryInfoScRsp                          = 7124
	EvolveBuildGiveupScRsp                             = 7104
	EvolveBuildShopAbilityDownCsReq                    = 7121
	EvolveBuildShopAbilityResetScRsp                   = 7143
	TakeExpeditionRewardCsReq                          = 2539
	TakeActivityExpeditionRewardScRsp                  = 2582
	CancelActivityExpeditionCsReq                      = 2554
	AcceptMultipleExpeditionCsReq                      = 2534
	TakeMultipleExpeditionRewardScRsp                  = 2525
	AcceptActivityExpeditionCsReq                      = 2533
	GetExpeditionDataScRsp                             = 2506
	CancelExpeditionCsReq                              = 2556
	AcceptExpeditionCsReq                              = 2511
	AcceptExpeditionScRsp                              = 2543
	CancelActivityExpeditionScRsp                      = 2565
	TakeActivityExpeditionRewardCsReq                  = 2548
	AcceptActivityExpeditionScRsp                      = 2599
	GetExpeditionDataCsReq                             = 2568
	ExpeditionDataChangeScNotify                       = 2551
	AcceptMultipleExpeditionScRsp                      = 2597
	TakeMultipleExpeditionRewardCsReq                  = 2598
	CancelExpeditionScRsp                              = 2593
	TakeExpeditionRewardScRsp                          = 2529
	EnterFantasticStoryActivityStageCsReq              = 4943
	FantasticStoryActivityBattleEndScNotify            = 4993
	GetFantasticStoryActivityDataCsReq                 = 4968
	EnterFantasticStoryActivityStageScRsp              = 4956
	GetFantasticStoryActivityDataScRsp                 = 4906
	FinishChapterScNotify                              = 4911
	GetFeverTimeActivityDataCsReq                      = 7160
	EnterFeverTimeActivityStageScRsp                   = 7154
	GetFeverTimeActivityDataScRsp                      = 7152
	EnterFeverTimeActivityStageCsReq                   = 7151
	FeverTimeActivityBattleEndScNotify                 = 7159
	FightHeartBeatScRsp                                = 30093
	FightSessionStopScNotify                           = 30039
	FightEnterCsReq                                    = 30068
	FightHeartBeatCsReq                                = 30056
	FightKickOutScNotify                               = 30043
	FightEnterScRsp                                    = 30006
	FightLeaveScNotify                                 = 30011
	EnterFightActivityStageCsReq                       = 3643
	TakeFightActivityRewardScRsp                       = 3639
	GetFightActivityDataScRsp                          = 3606
	GetFightActivityDataCsReq                          = 3668
	FightActivityDataChangeScNotify                    = 3611
	EnterFightActivityStageScRsp                       = 3656
	TakeFightActivityRewardCsReq                       = 3693
	GetFightFestDataScRsp                              = 7274
	FightFestUpdateCoinNotify                          = 7254
	FightFestScoreUpdateNotify                         = 7261
	StartFightFestScRsp                                = 7267
	GetFightFestDataCsReq                              = 7257
	FightFestUpdateChallengeRecordNotify               = 7273
	StartFightFestCsReq                                = 7265
	FightFestUnlockSkillNotify                         = 7258
	FightMatch3ForceUpdateNotify                       = 30154
	FightMatch3SwapCsReq                               = 30193
	FightMatch3StartCountDownScNotify                  = 30111
	FightMatch3ChatScNotify                            = 30199
	FightMatch3DataCsReq                               = 30168
	FightMatch3ChatCsReq                               = 30151
	FightMatch3TurnStartScNotify                       = 30143
	FightMatch3SwapScRsp                               = 30139
	FightMatch3DataScRsp                               = 30106
	FightMatch3OpponentDataScNotify                    = 30129
	FightMatch3TurnEndScNotify                         = 30156
	FightMatch3ChatScRsp                               = 30133
	ApplyFriendCsReq                                   = 2939
	GetFriendLoginInfoCsReq                            = 2969
	GetPlayerDetailInfoScRsp                           = 2943
	SearchPlayerCsReq                                  = 2908
	GetCurAssistScRsp                                  = 2976
	SyncDeleteFriendScNotify                           = 2982
	DeleteBlacklistCsReq                               = 2909
	GetFriendBattleRecordDetailCsReq                   = 2941
	GetPlatformPlayerInfoCsReq                         = 2995
	GetFriendLoginInfoScRsp                            = 2967
	NewAssistHistoryNotify                             = 2966
	SetFriendRemarkNameCsReq                           = 2917
	DeleteFriendCsReq                                  = 2965
	GetPlatformPlayerInfoScRsp                         = 2989
	SearchPlayerScRsp                                  = 2981
	SetAssistScRsp                                     = 2923
	GetFriendApplyListInfoScRsp                        = 2993
	TakeAssistRewardScRsp                              = 2912
	SetAssistCsReq                                     = 2984
	GetAssistHistoryCsReq                              = 2949
	SetFriendRemarkNameScRsp                           = 2992
	GetAssistHistoryScRsp                              = 2945
	DeleteFriendScRsp                                  = 2948
	GetAssistListScRsp                                 = 2916
	CurAssistChangedNotify                             = 2930
	TakeAssistRewardCsReq                              = 2946
	GetFriendListInfoScRsp                             = 2906
	GetFriendListInfoCsReq                             = 2968
	GetFriendAssistListScRsp                           = 2953
	GetFriendDevelopmentInfoCsReq                      = 2903
	SyncHandleFriendScNotify                           = 2954
	GetFriendApplyListInfoCsReq                        = 2956
	GetFriendAssistListCsReq                           = 3000
	GetFriendChallengeLineupCsReq                      = 2958
	HandleFriendCsReq                                  = 2933
	ApplyFriendScRsp                                   = 2929
	SetFriendMarkCsReq                                 = 2940
	GetFriendChallengeDetailCsReq                      = 2959
	GetAssistListCsReq                                 = 2932
	SyncApplyFriendScNotify                            = 2951
	GetFriendDevelopmentInfoScRsp                      = 2926
	HandleFriendScRsp                                  = 2999
	GetFriendChallengeDetailScRsp                      = 2907
	GetPlayerDetailInfoCsReq                           = 2911
	SyncAddBlacklistScNotify                           = 2998
	ReportPlayerScRsp                                  = 2928
	AddBlacklistCsReq                                  = 2934
	GetCurAssistCsReq                                  = 2978
	GetFriendChallengeLineupScRsp                      = 2947
	GetFriendBattleRecordDetailScRsp                   = 2931
	DeleteBlacklistScRsp                               = 2988
	GetFriendRecommendListInfoScRsp                    = 2924
	AddBlacklistScRsp                                  = 2997
	SetForbidOtherApplyFriendCsReq                     = 2915
	GetFriendRecommendListInfoCsReq                    = 2925
	SetForbidOtherApplyFriendScRsp                     = 2975
	SetFriendMarkScRsp                                 = 2985
	ReportPlayerCsReq                                  = 2990
	ExchangeGachaCeilingCsReq                          = 1939
	GetGachaInfoCsReq                                  = 1968
	GetGachaCeilingScRsp                               = 1993
	DoGachaScRsp                                       = 1943
	DoGachaCsReq                                       = 1911
	GetGachaInfoScRsp                                  = 1906
	ExchangeGachaCeilingScRsp                          = 1929
	GetGachaCeilingCsReq                               = 1956
	HeartDialScriptChangeScNotify                      = 6351
	ChangeScriptEmotionScRsp                           = 6343
	HeartDialTraceScriptScRsp                          = 6399
	GetHeartDialInfoCsReq                              = 6368
	GetHeartDialInfoScRsp                              = 6306
	HeartDialTraceScriptCsReq                          = 6333
	ChangeScriptEmotionCsReq                           = 6311
	FinishEmotionDialoguePerformanceCsReq              = 6339
	SubmitEmotionItemCsReq                             = 6356
	SubmitEmotionItemScRsp                             = 6393
	FinishEmotionDialoguePerformanceScRsp              = 6329
	HeliobusEnterBattleCsReq                           = 5824
	HeliobusStartRaidCsReq                             = 5892
	HeliobusEnterBattleScRsp                           = 5817
	HeliobusSnsCommentScRsp                            = 5833
	HeliobusInfoChangedScNotify                        = 5854
	HeliobusSnsLikeScRsp                               = 5829
	HeliobusUnlockSkillScNotify                        = 5882
	HeliobusSnsPostCsReq                               = 5856
	HeliobusActivityDataScRsp                          = 5806
	HeliobusUpgradeLevelCsReq                          = 5865
	HeliobusSnsLikeCsReq                               = 5839
	HeliobusActivityDataCsReq                          = 5868
	HeliobusSnsReadCsReq                               = 5811
	HeliobusSnsUpdateScNotify                          = 5899
	HeliobusSnsReadScRsp                               = 5843
	HeliobusSnsPostScRsp                               = 5893
	HeliobusUpgradeLevelScRsp                          = 5848
	HeliobusSnsCommentCsReq                            = 5851
	HeliobusSelectSkillCsReq                           = 5834
	HeliobusLineupUpdateScNotify                       = 5809
	HeliobusChallengeUpdateScNotify                    = 5828
	HeliobusStartRaidScRsp                             = 5890
	HeliobusSelectSkillScRsp                           = 5897
	DestroyItemScRsp                                   = 523
	LockEquipmentScRsp                                 = 593
	ExchangeHcoinScRsp                                 = 590
	GetMarkItemListCsReq                               = 578
	GetMarkItemListScRsp                               = 576
	AddEquipmentScNotify                               = 588
	DeleteRelicFilterPlanCsReq                         = 507
	MarkItemCsReq                                      = 549
	GetRecyleTimeScRsp                                 = 581
	ComposeItemScRsp                                   = 548
	ExpUpRelicCsReq                                    = 582
	GetBagScRsp                                        = 506
	AddRelicFilterPlanScRsp                            = 558
	ExpUpEquipmentScRsp                                = 554
	SetTurnFoodSwitchScRsp                             = 530
	ExpUpEquipmentCsReq                                = 599
	PromoteEquipmentScRsp                              = 543
	ExchangeHcoinCsReq                                 = 592
	RankUpEquipmentScRsp                               = 533
	LockEquipmentCsReq                                 = 556
	CancelMarkItemNotify                               = 566
	MarkItemScRsp                                      = 545
	ComposeSelectedRelicCsReq                          = 528
	SetTurnFoodSwitchCsReq                             = 512
	DeleteRelicFilterPlanScRsp                         = 541
	GeneralVirtualItemDataNotify                       = 595
	GetRecyleTimeCsReq                                 = 508
	ModifyRelicFilterPlanCsReq                         = 547
	UseItemScRsp                                       = 529
	UseItemCsReq                                       = 539
	DiscardRelicScRsp                                  = 569
	RechargeSuccNotify                                 = 517
	AddRelicFilterPlanCsReq                            = 553
	RankUpEquipmentCsReq                               = 551
	ComposeLimitNumUpdateNotify                        = 516
	SellItemScRsp                                      = 524
	GetBagCsReq                                        = 568
	ModifyRelicFilterPlanScRsp                         = 559
	PromoteEquipmentCsReq                              = 511
	LockRelicScRsp                                     = 598
	MarkRelicFilterPlanCsReq                           = 531
	DiscardRelicCsReq                                  = 589
	ComposeItemCsReq                                   = 565
	SyncTurnFoodNotify                                 = 546
	ComposeLimitNumCompleteNotify                      = 532
	LockRelicCsReq                                     = 597
	ExpUpRelicScRsp                                    = 534
	SellItemCsReq                                      = 525
	GetRelicFilterPlanScRsp                            = 600
	DestroyItemCsReq                                   = 584
	GetRelicFilterPlanCsReq                            = 585
	MarkRelicFilterPlanScRsp                           = 503
	RelicFilterPlanClearNameScNotify                   = 526
	ComposeSelectedRelicScRsp                          = 509
	PlayBackGroundMusicCsReq                           = 3111
	TrialBackGroundMusicCsReq                          = 3139
	GetJukeboxDataScRsp                                = 3106
	UnlockBackGroundMusicScRsp                         = 3193
	TrialBackGroundMusicScRsp                          = 3129
	PlayBackGroundMusicScRsp                           = 3143
	GetJukeboxDataCsReq                                = 3168
	UnlockBackGroundMusicCsReq                         = 3156
	SetLineupNameCsReq                                 = 798
	QuitLineupScRsp                                    = 729
	GetCurLineupDataScRsp                              = 743
	GetAllLineupDataCsReq                              = 724
	SetLineupNameScRsp                                 = 725
	JoinLineupScRsp                                    = 793
	ChangeLineupLeaderScRsp                            = 782
	SwitchLineupIndexCsReq                             = 734
	SwapLineupCsReq                                    = 751
	SwapLineupScRsp                                    = 733
	ReplaceLineupScRsp                                 = 728
	GetAllLineupDataScRsp                              = 717
	GetLineupAvatarDataCsReq                           = 754
	GetStageLineupScRsp                                = 706
	GetLineupAvatarDataScRsp                           = 765
	JoinLineupCsReq                                    = 756
	GetStageLineupCsReq                                = 768
	ReplaceLineupCsReq                                 = 790
	SyncLineupNotify                                   = 799
	GetCurLineupDataCsReq                              = 711
	SwitchLineupIndexScRsp                             = 797
	ExtraLineupDestroyNotify                           = 709
	ChangeLineupLeaderCsReq                            = 748
	QuitLineupCsReq                                    = 739
	VirtualLineupDestroyNotify                         = 792
	LobbyQuitCsReq                                     = 7362
	LobbyInviteScRsp                                   = 7356
	LobbyBeginCsReq                                    = 7361
	LobbyJoinScRsp                                     = 7367
	LobbyJoinCsReq                                     = 7365
	LobbyModifyPlayerInfoScRsp                         = 7354
	LobbyKickOutCsReq                                  = 7384
	LobbyGetInfoScRsp                                  = 7371
	LobbyKickOutScRsp                                  = 7370
	LobbyGetInfoCsReq                                  = 7398
	LobbyBeginScRsp                                    = 7358
	LobbyCreateScRsp                                   = 7374
	LobbyCreateCsReq                                   = 7357
	LobbySyncInfoScNotify                              = 7386
	LobbyInviteCsReq                                   = 7395
	LobbyModifyPlayerInfoCsReq                         = 7373
	LobbyQuitScRsp                                     = 7389
	LobbyInviteScNotify                                = 7355
	GetMailScRsp                                       = 806
	MarkReadMailCsReq                                  = 811
	GetMailCsReq                                       = 868
	MarkReadMailScRsp                                  = 843
	NewMailScNotify                                    = 851
	DelMailScRsp                                       = 893
	TakeMailAttachmentScRsp                            = 829
	DelMailCsReq                                       = 856
	TakeMailAttachmentCsReq                            = 839
	ResetMapRotationRegionCsReq                        = 6865
	UpdateRotaterScNotify                              = 6824
	DeployRotaterScRsp                                 = 6893
	RotateMapCsReq                                     = 6839
	InteractChargerScRsp                               = 6843
	GetMapRotationDataCsReq                            = 6899
	InteractChargerCsReq                               = 6811
	UpdateMapRotationDataScNotify                      = 6897
	LeaveMapRotationRegionScRsp                        = 6833
	UpdateEnergyScNotify                               = 6834
	LeaveMapRotationRegionScNotify                     = 6882
	EnterMapRotationRegionScRsp                        = 6806
	GetMapRotationDataScRsp                            = 6854
	LeaveMapRotationRegionCsReq                        = 6851
	ResetMapRotationRegionScRsp                        = 6848
	RotateMapScRsp                                     = 6829
	DeployRotaterCsReq                                 = 6856
	RemoveRotaterScRsp                                 = 6825
	EnterMapRotationRegionCsReq                        = 6868
	RemoveRotaterCsReq                                 = 6898
	MatchResultScNotify                                = 7311
	StartMatchScRsp                                    = 7324
	CancelMatchScRsp                                   = 7317
	StartMatchCsReq                                    = 7307
	GetCrossInfoCsReq                                  = 7308
	CancelMatchCsReq                                   = 7315
	GetCrossInfoScRsp                                  = 7323
	MatchThreeLevelEndCsReq                            = 7415
	MatchThreeGetDataCsReq                             = 7407
	MatchThreeSetBirdPosCsReq                          = 7408
	MatchThreeLevelEndScRsp                            = 7417
	MatchThreeSetBirdPosScRsp                          = 7423
	MatchThreeGetDataScRsp                             = 7424
	MatchThreeSyncDataScNotify                         = 7411
	FinishItemIdScRsp                                  = 2793
	FinishPerformSectionIdScRsp                        = 2733
	FinishSectionIdScRsp                               = 2729
	FinishSectionIdCsReq                               = 2739
	GetNpcMessageGroupScRsp                            = 2706
	FinishPerformSectionIdCsReq                        = 2751
	FinishItemIdCsReq                                  = 2756
	GetNpcStatusScRsp                                  = 2743
	GetNpcMessageGroupCsReq                            = 2768
	GetNpcStatusCsReq                                  = 2711
	DifficultyAdjustmentUpdateDataScRsp                = 4123
	SubmitOrigamiItemCsReq                             = 4182
	DifficultyAdjustmentUpdateDataCsReq                = 4184
	SecurityReportScRsp                                = 4154
	GetGunPlayDataScRsp                                = 4188
	DifficultyAdjustmentGetDataCsReq                   = 4132
	GetMovieRacingDataCsReq                            = 4117
	GetShareDataScRsp                                  = 4143
	DifficultyAdjustmentGetDataScRsp                   = 4116
	SecurityReportCsReq                                = 4199
	CancelCacheNotifyCsReq                             = 4151
	ShareCsReq                                         = 4168
	GetShareDataCsReq                                  = 4111
	TakePictureScRsp                                   = 4193
	CancelCacheNotifyScRsp                             = 4133
	UpdateGunPlayDataScRsp                             = 4181
	GetGunPlayDataCsReq                                = 4109
	TakePictureCsReq                                   = 4156
	GetMovieRacingDataScRsp                            = 4192
	UpdateMovieRacingDataScRsp                         = 4128
	ShareScRsp                                         = 4106
	TriggerVoiceScRsp                                  = 4148
	UpdateGunPlayDataCsReq                             = 4108
	UpdateMovieRacingDataCsReq                         = 4190
	SubmitOrigamiItemScRsp                             = 4134
	TriggerVoiceCsReq                                  = 4165
	DailyTaskDataScNotify                              = 1229
	AcceptMissionEventScRsp                            = 1225
	SyncTaskCsReq                                      = 1293
	MissionAcceptScNotify                              = 1249
	SetMissionEventProgressScRsp                       = 1209
	SetMissionEventProgressCsReq                       = 1228
	StartFinishSubMissionScNotify                      = 1232
	MissionGroupWarnScNotify                           = 1254
	GetMissionStatusScRsp                              = 1217
	GetMissionDataCsReq                                = 1268
	FinishCosumeItemMissionCsReq                       = 1265
	AcceptMainMissionScRsp                             = 1223
	GetMissionStatusCsReq                              = 1224
	GetMissionEventDataScRsp                           = 1234
	FinishTalkMissionScRsp                             = 1243
	SyncTaskScRsp                                      = 1239
	FinishCosumeItemMissionScRsp                       = 1248
	AcceptMainMissionCsReq                             = 1284
	GetMainMissionCustomValueScRsp                     = 1276
	StartFinishMainMissionScNotify                     = 1216
	SubMissionRewardScNotify                           = 1288
	InterruptMissionEventCsReq                         = 1292
	GetMissionDataScRsp                                = 1206
	FinishTalkMissionCsReq                             = 1211
	GetMissionEventDataCsReq                           = 1282
	TeleportToMissionResetPointScRsp                   = 1281
	GetMainMissionCustomValueCsReq                     = 1278
	MissionEventRewardScNotify                         = 1297
	MissionRewardScNotify                              = 1256
	UpdateTrackMainMissionIdScRsp                      = 1246
	AcceptMissionEventCsReq                            = 1298
	UpdateTrackMainMissionIdCsReq                      = 1266
	TeleportToMissionResetPointCsReq                   = 1208
	InterruptMissionEventScRsp                         = 1290
	MonopolyScrachRaffleTicketScRsp                    = 7050
	MonopolyMoveScRsp                                  = 7051
	GetMonopolyFriendRankingListCsReq                  = 7058
	MonopolyReRollRandomScRsp                          = 7098
	MonopolyGameGachaCsReq                             = 7078
	GetMonopolyFriendRankingListScRsp                  = 7047
	MonopolyGuessDrawScNotify                          = 7067
	MonopolyGetRegionProgressScRsp                     = 7005
	GetMonopolyInfoCsReq                               = 7068
	MonopolyBuyGoodsCsReq                              = 7017
	MonopolyTakePhaseRewardCsReq                       = 7061
	MonopolyGameRaiseRatioCsReq                        = 7016
	MonopolySelectOptionCsReq                          = 7033
	MonopolyGetRaffleTicketCsReq                       = 7042
	MonopolyCellUpdateNotify                           = 7043
	GetMonopolyMbtiReportRewardCsReq                   = 7062
	DailyFirstEnterMonopolyActivityScRsp               = 7048
	DeleteSocialEventServerCacheScRsp                  = 7073
	MonopolyGuessBuyInformationCsReq                   = 7089
	MonopolyGiveUpCurContentCsReq                      = 7009
	MonopolyActionResultScNotify                       = 7011
	MonopolyMoveCsReq                                  = 7029
	MonopolyUpgradeAssetCsReq                          = 7090
	MonopolyLikeCsReq                                  = 7059
	MonopolyClickMbtiReportCsReq                       = 7038
	MonopolyGameRaiseRatioScRsp                        = 7084
	MonopolyCheatDiceCsReq                             = 7008
	GetMonopolyMbtiReportRewardScRsp                   = 7001
	MonopolyGameBingoFlipCardScRsp                     = 7045
	MonopolyRollRandomCsReq                            = 7082
	MonopolyGuessChooseScRsp                           = 7095
	MonopolyGiveUpCurContentScRsp                      = 7088
	MonopolyRollDiceCsReq                              = 7093
	GetMonopolyInfoScRsp                               = 7006
	MonopolyConfirmRandomCsReq                         = 7025
	MonopolyDailySettleScNotify                        = 7053
	GetMonopolyDailyReportScRsp                        = 7020
	GetSocialEventServerCacheCsReq                     = 7022
	MonopolyGameSettleScNotify                         = 7023
	MonopolyEventLoadUpdateScNotify                    = 7052
	MonopolyAcceptQuizCsReq                            = 7066
	MonopolyContentUpdateScNotify                      = 7032
	MonopolyGetDailyInitItemCsReq                      = 7077
	MonopolyGetRaffleTicketScRsp                       = 7083
	MonopolyReRollRandomCsReq                          = 7097
	MonopolyLikeScNotify                               = 7041
	MonopolyGameGachaScRsp                             = 7076
	GetSocialEventServerCacheScRsp                     = 7021
	MonopolyGetRegionProgressCsReq                     = 7072
	MonopolySocialEventEffectScNotify                  = 7014
	MonopolyGetRafflePoolInfoCsReq                     = 7064
	MonopolyRollRandomScRsp                            = 7034
	MonopolyGetRafflePoolInfoScRsp                     = 7070
	MonopolyLikeScRsp                                  = 7007
	MonopolyClickCellScRsp                             = 7027
	MonopolyUpgradeAssetScRsp                          = 7028
	MonopolyGameBingoFlipCardCsReq                     = 7049
	MonopolyClickMbtiReportScRsp                       = 7080
	MonopolyGetDailyInitItemScRsp                      = 7055
	MonopolyEventSelectFriendCsReq                     = 7026
	MonopolyTakeRaffleTicketRewardScRsp                = 7036
	MonopolyRollDiceScRsp                              = 7039
	MonopolyCheatDiceScRsp                             = 7081
	GetMbtiReportScRsp                                 = 7003
	MonopolyConditionUpdateScNotify                    = 7063
	MonopolySelectOptionScRsp                          = 7099
	MonopolyGameCreateScNotify                         = 7012
	MonopolyGuessBuyInformationScRsp                   = 7069
	MonopolyQuizDurationChangeScNotify                 = 7015
	MonopolyTakePhaseRewardScRsp                       = 7079
	GetMbtiReportCsReq                                 = 7031
	MonopolyBuyGoodsScRsp                              = 7092
	MonopolyAcceptQuizScRsp                            = 7046
	GetMonopolyDailyReportCsReq                        = 7091
	MonopolySttUpdateScNotify                          = 7002
	MonopolyTakeRaffleTicketRewardCsReq                = 7096
	MonopolyGuessChooseCsReq                           = 7030
	MonopolyEventSelectFriendScRsp                     = 7013
	DailyFirstEnterMonopolyActivityCsReq               = 7065
	MonopolyConfirmRandomScRsp                         = 7024
	DeleteSocialEventServerCacheCsReq                  = 7071
	MonopolyScrachRaffleTicketCsReq                    = 7010
	MonopolyClickCellCsReq                             = 7057
	MultiplayerFightGiveUpCsReq                        = 1056
	MultiplayerGetFightGateScRsp                       = 1043
	MultiplayerFightGameFinishScNotify                 = 1029
	MultiplayerFightGameStateCsReq                     = 1068
	MultiplayerFightGameStateScRsp                     = 1006
	MultiplayerFightGameStartScNotify                  = 1039
	MultiplayerMatch3FinishScNotify                    = 1051
	MultiplayerFightGiveUpScRsp                        = 1093
	MultiplayerGetFightGateCsReq                       = 1011
	MultipleDropInfoNotify                             = 4693
	GetMultipleDropInfoCsReq                           = 4668
	GetMultipleDropInfoScRsp                           = 4606
	MultipleDropInfoScNotify                           = 4611
	GetPlayerReturnMultiDropInfoCsReq                  = 4643
	GetPlayerReturnMultiDropInfoScRsp                  = 4656
	UpgradeAreaScRsp                                   = 4348
	MuseumDispatchFinishedScNotify                     = 4328
	UpgradeAreaCsReq                                   = 4365
	MuseumTargetMissionFinishNotify                    = 4388
	MuseumRandomEventQueryCsReq                        = 4324
	GetExhibitScNotify                                 = 4333
	BuyNpcStuffScRsp                                   = 4343
	MuseumRandomEventQueryScRsp                        = 4317
	SetStuffToAreaCsReq                                = 4356
	RemoveStuffFromAreaCsReq                           = 4339
	MuseumTargetRewardNotify                           = 4308
	MuseumRandomEventStartScNotify                     = 4325
	MuseumRandomEventSelectScRsp                       = 4390
	FinishCurTurnScRsp                                 = 4354
	MuseumFundsChangedScNotify                         = 4398
	MuseumTakeCollectRewardScRsp                       = 4332
	GetMuseumInfoCsReq                                 = 4368
	GetMuseumInfoScRsp                                 = 4306
	SetStuffToAreaScRsp                                = 4393
	UpgradeAreaStatScRsp                               = 4334
	MuseumTakeCollectRewardCsReq                       = 4381
	UpgradeAreaStatCsReq                               = 4382
	MuseumTargetStartNotify                            = 4309
	BuyNpcStuffCsReq                                   = 4311
	FinishCurTurnCsReq                                 = 4399
	GetStuffScNotify                                   = 4351
	MuseumRandomEventSelectCsReq                       = 4392
	RemoveStuffFromAreaScRsp                           = 4329
	MuseumInfoChangedScNotify                          = 4397
	GetOfferingInfoScRsp                               = 6923
	TakeOfferingRewardCsReq                            = 6924
	TakeOfferingRewardScRsp                            = 6922
	SubmitOfferingItemScRsp                            = 6940
	SubmitOfferingItemCsReq                            = 6932
	OfferingInfoScNotify                               = 6936
	GetOfferingInfoCsReq                               = 6939
	AcceptedPamMissionExpireCsReq                      = 4068
	AcceptedPamMissionExpireScRsp                      = 4006
	SyncAcceptedPamMissionNotify                       = 4011
	RecallPetScRsp                                     = 7608
	RecallPetCsReq                                     = 7611
	CurPetChangedScNotify                              = 7623
	GetPetDataCsReq                                    = 7607
	GetPetDataScRsp                                    = 7624
	SummonPetScRsp                                     = 7617
	SummonPetCsReq                                     = 7615
	UnlockChatBubbleScNotify                           = 5156
	GetPhoneDataCsReq                                  = 5168
	GetPhoneDataScRsp                                  = 5106
	SelectChatBubbleScRsp                              = 5143
	UnlockPhoneThemeScNotify                           = 5129
	SelectPhoneThemeScRsp                              = 5139
	SelectPhoneThemeCsReq                              = 5193
	SelectChatBubbleCsReq                              = 5111
	GetVideoVersionKeyCsReq                            = 42
	ExchangeStaminaCsReq                               = 48
	PlayerLoginFinishScRsp                             = 86
	SetPlayerInfoScRsp                                 = 95
	GetBasicInfoScRsp                                  = 85
	PlayerGetTokenScRsp                                = 93
	ExchangeStaminaScRsp                               = 82
	AceAntiCheaterScRsp                                = 59
	UnlockAvatarPathCsReq                              = 63
	UpdatePlayerSettingScRsp                           = 70
	SetMultipleAvatarPathsCsReq                        = 80
	PlayerLogoutScRsp                                  = 43
	GetLevelRewardScRsp                                = 9
	ReserveStaminaExchangeCsReq                        = 50
	GetMultiPathAvatarInfoScRsp                        = 77
	GmTalkScRsp                                        = 99
	RetcodeNotify                                      = 41
	ClientDownloadDataScNotify                         = 15
	PlayerLoginScRsp                                   = 6
	FeatureSwitchClosedScNotify                        = 13
	SetLanguageCsReq                                   = 81
	GetLevelRewardCsReq                                = 28
	AceAntiCheaterCsReq                                = 47
	StaminaInfoScNotify                                = 5
	PlayerHeartBeatScRsp                               = 3
	SetNicknameScRsp                                   = 17
	AntiAddictScNotify                                 = 25
	UpdatePsnSettingsInfoCsReq                         = 2
	GetVideoVersionKeyScRsp                            = 83
	ClientObjDownloadDataScNotify                      = 61
	PlayerLogoutCsReq                                  = 11
	SetNicknameCsReq                                   = 24
	PlayerLoginCsReq                                   = 68
	DailyRefreshNotify                                 = 100
	ReserveStaminaExchangeScRsp                        = 72
	GetAuthkeyScRsp                                    = 97
	GetBasicInfoCsReq                                  = 40
	SetRedPointStatusScNotify                          = 96
	SetLanguageScRsp                                   = 32
	UpdateFeatureSwitchScNotify                        = 75
	SetPlayerInfoCsReq                                 = 30
	ServerAnnounceNotify                               = 16
	QueryProductInfoScRsp                              = 67
	UnlockAvatarPathScRsp                              = 38
	GmTalkScNotify                                     = 29
	SetAvatarPathCsReq                                 = 20
	GetMultiPathAvatarInfoCsReq                        = 27
	AvatarPathChangedNotify                            = 55
	GetLevelRewardTakenListCsReq                       = 92
	MonthCardRewardNotify                              = 7
	PlayerHeartBeatCsReq                               = 31
	RegionStopScNotify                                 = 98
	PlayerLoginFinishCsReq                             = 73
	SetAvatarPathScRsp                                 = 57
	GetSecretKeyInfoCsReq                              = 21
	ClientObjUploadCsReq                               = 79
	SetGenderScRsp                                     = 12
	UpdatePsnSettingsInfoScRsp                         = 62
	PlayerGetTokenCsReq                                = 56
	ClientObjUploadScRsp                               = 52
	SetMultipleAvatarPathsScRsp                        = 4
	SetGenderCsReq                                     = 46
	GmTalkCsReq                                        = 33
	GetAuthkeyCsReq                                    = 34
	QueryProductInfoCsReq                              = 69
	SetGameplayBirthdayScRsp                           = 58
	UpdatePlayerSettingCsReq                           = 64
	SetGameplayBirthdayCsReq                           = 53
	GateServerScNotify                                 = 26
	GetSecretKeyInfoScRsp                              = 71
	PlayerKickOutScNotify                              = 51
	GetLevelRewardTakenListScRsp                       = 90
	SetHeadIconScRsp                                   = 2843
	SetIsDisplayAvatarInfoCsReq                        = 2839
	GetPlayerBoardDataScRsp                            = 2806
	SetAssistAvatarScRsp                               = 2865
	SetSignatureCsReq                                  = 2833
	SetHeadIconCsReq                                   = 2811
	SetDisplayAvatarScRsp                              = 2893
	UnlockHeadIconScNotify                             = 2851
	SetSignatureScRsp                                  = 2899
	GetPlayerBoardDataCsReq                            = 2868
	SetAssistAvatarCsReq                               = 2854
	SetDisplayAvatarCsReq                              = 2856
	SetIsDisplayAvatarInfoScRsp                        = 2829
	PlayerReturnSignScRsp                              = 4511
	PlayerReturnTakePointRewardScRsp                   = 4593
	PlayerReturnStartScNotify                          = 4568
	PlayerReturnInfoQueryCsReq                         = 4551
	PlayerReturnTakeRewardCsReq                        = 4539
	PlayerReturnSignCsReq                              = 4506
	PlayerReturnForceFinishScNotify                    = 4599
	PlayerReturnTakeRewardScRsp                        = 4529
	PlayerReturnTakePointRewardCsReq                   = 4556
	PlayerReturnPointChangeScNotify                    = 4543
	PlayerReturnInfoQueryScRsp                         = 4533
	FinishPlotScRsp                                    = 1106
	FinishPlotCsReq                                    = 1168
	GetPunkLordBattleRecordScRsp                       = 3278
	PunkLordRaidTimeOutScNotify                        = 3225
	SharePunkLordMonsterCsReq                          = 3256
	SharePunkLordMonsterScRsp                          = 3293
	PunkLordDataChangeNotify                           = 3284
	SummonPunkLordMonsterScRsp                         = 3229
	GetKilledPunkLordMonsterDataCsReq                  = 3228
	TakePunkLordPointRewardScRsp                       = 3248
	GetPunkLordDataScRsp                               = 3297
	StartPunkLordRaidCsReq                             = 3211
	GetPunkLordDataCsReq                               = 3234
	PunkLordMonsterInfoScNotify                        = 3282
	SummonPunkLordMonsterCsReq                         = 3239
	PunkLordBattleResultScNotify                       = 3290
	GetPunkLordMonsterDataScRsp                        = 3206
	TakeKilledPunkLordMonsterScoreCsReq                = 3232
	StartPunkLordRaidScRsp                             = 3243
	TakeKilledPunkLordMonsterScoreScRsp                = 3216
	PunkLordMonsterKilledNotify                        = 3281
	GetPunkLordBattleRecordCsReq                       = 3223
	TakePunkLordPointRewardCsReq                       = 3265
	GetKilledPunkLordMonsterDataScRsp                  = 3209
	GetPunkLordMonsterDataCsReq                        = 3268
	BatchGetQuestDataScRsp                             = 934
	GetQuestDataCsReq                                  = 968
	TakeQuestOptionalRewardScRsp                       = 965
	QuestRecordScNotify                                = 951
	FinishQuestScRsp                                   = 999
	BatchGetQuestDataCsReq                             = 982
	GetQuestRecordCsReq                                = 939
	TakeQuestOptionalRewardCsReq                       = 954
	GetQuestDataScRsp                                  = 906
	TakeQuestRewardCsReq                               = 911
	FinishQuestCsReq                                   = 933
	TakeQuestRewardScRsp                               = 943
	GetQuestRecordScRsp                                = 929
	StartRaidScRsp                                     = 2206
	GetChallengeRaidInfoScRsp                          = 2239
	StartRaidCsReq                                     = 2268
	SetClientRaidTargetCountCsReq                      = 2265
	RaidKickByServerScNotify                           = 2224
	GetSaveRaidScRsp                                   = 2234
	GetAllSaveRaidCsReq                                = 2297
	ChallengeRaidNotify                                = 2233
	RaidInfoNotify                                     = 2256
	LeaveRaidCsReq                                     = 2211
	GetAllSaveRaidScRsp                                = 2298
	GetRaidInfoScRsp                                   = 2254
	TakeChallengeRaidRewardCsReq                       = 2229
	GetRaidInfoCsReq                                   = 2299
	DelSaveRaidScNotify                                = 2225
	TakeChallengeRaidRewardScRsp                       = 2251
	SetClientRaidTargetCountScRsp                      = 2248
	GetSaveRaidCsReq                                   = 2282
	LeaveRaidScRsp                                     = 2243
	GetChallengeRaidInfoCsReq                          = 2293
	RaidCollectionEnterNextRaidCsReq                   = 6960
	RaidCollectionEnterNextRaidScRsp                   = 6944
	RaidCollectionDataScNotify                         = 6952
	RaidCollectionDataScRsp                            = 6943
	RaidCollectionDataCsReq                            = 6959
	GetChallengeRecommendLineupListScRsp               = 2408
	GetChallengeRecommendLineupListCsReq               = 2411
	RelicAvatarRecommendScRsp                          = 2417
	RelicRecommendCsReq                                = 2407
	RelicAvatarRecommendCsReq                          = 2415
	RelicRecommendScRsp                                = 2424
	UpdateRedDotDataCsReq                              = 5911
	GetSingleRedDotParamGroupCsReq                     = 5956
	GetSingleRedDotParamGroupScRsp                     = 5993
	GetAllRedDotDataCsReq                              = 5968
	UpdateRedDotDataScRsp                              = 5943
	GetAllRedDotDataScRsp                              = 5906
	GetReplayTokenScRsp                                = 3506
	GetReplayTokenCsReq                                = 3568
	GetPlayerReplayInfoCsReq                           = 3511
	GetPlayerReplayInfoScRsp                           = 3543
	DailyFirstMeetPamScRsp                             = 3443
	GetRndOptionCsReq                                  = 3468
	GetRndOptionScRsp                                  = 3406
	DailyFirstMeetPamCsReq                             = 3411
	GetRogueInfoCsReq                                  = 1868
	GetRogueBuffEnhanceInfoCsReq                       = 1890
	ExchangeRogueRewardKeyScRsp                        = 1803
	SyncRogueExploreWinScNotify                        = 1849
	SyncRogueStatusScNotify                            = 1887
	LeaveRogueScRsp                                    = 1829
	GetRogueInfoScRsp                                  = 1806
	TakeRogueScoreRewardCsReq                          = 1817
	SyncRogueAeonLevelUpRewardScNotify                 = 1870
	FinishAeonDialogueGroupScRsp                       = 1855
	EnableRogueTalentCsReq                             = 1880
	OpenRogueChestScRsp                                = 1841
	GetRogueScoreRewardInfoCsReq                       = 1861
	SyncRogueRewardInfoScNotify                        = 1818
	ReviveRogueAvatarCsReq                             = 1825
	GetRogueInitialScoreCsReq                          = 1895
	EnterRogueCsReq                                    = 1856
	FinishAeonDialogueGroupCsReq                       = 1877
	GetRogueTalentInfoScRsp                            = 1838
	TakeRogueAeonLevelRewardCsReq                      = 1810
	QuitRogueScRsp                                     = 1878
	LeaveRogueCsReq                                    = 1839
	TakeRogueAeonLevelRewardScRsp                      = 1850
	GetRogueAeonInfoScRsp                              = 1827
	GetRogueScoreRewardInfoScRsp                       = 1879
	SyncRogueAreaUnlockScNotify                        = 1896
	EnhanceRogueBuffScRsp                              = 1888
	EnterRogueMapRoomCsReq                             = 1812
	ReviveRogueAvatarScRsp                             = 1824
	PickRogueAvatarCsReq                               = 1834
	SyncRogueAeonScNotify                              = 1883
	SyncRogueVirtualItemInfoScNotify                   = 1819
	OpenRogueChestCsReq                                = 1807
	EnterRogueScRsp                                    = 1893
	SyncRogueGetItemScNotify                           = 1836
	SyncRoguePickAvatarInfoScNotify                    = 1835
	EnhanceRogueBuffCsReq                              = 1809
	TakeRogueScoreRewardScRsp                          = 1892
	ExchangeRogueRewardKeyCsReq                        = 1831
	GetRogueBuffEnhanceInfoScRsp                       = 1828
	EnableRogueTalentScRsp                             = 1804
	GetRogueAeonInfoCsReq                              = 1857
	QuitRogueCsReq                                     = 1823
	SyncRogueReviveInfoScNotify                        = 1884
	GetRogueTalentInfoCsReq                            = 1863
	SyncRogueMapRoomScNotify                           = 1869
	SyncRogueFinishScNotify                            = 1882
	SyncRogueSeasonFinishScNotify                      = 1845
	PickRogueAvatarScRsp                               = 1897
	GetRogueInitialScoreScRsp                          = 1889
	StartRogueCsReq                                    = 1811
	EnterRogueMapRoomScRsp                             = 1830
	StartRogueScRsp                                    = 1843
	RogueArcadeLeaveCsReq                              = 7665
	RogueArcadeStartCsReq                              = 7657
	RogueArcadeRestartCsReq                            = 7661
	RogueArcadeGetInfoCsReq                            = 7673
	RogueArcadeGetInfoScRsp                            = 7654
	RogueArcadeLeaveScRsp                              = 7667
	RogueArcadeStartScRsp                              = 7674
	RogueArcadeRestartScRsp                            = 7658
	HandleRogueCommonPendingActionScRsp                = 5659
	CommonRogueUpdateScNotify                          = 5631
	RogueWorkbenchSelectFuncCsReq                      = 5677
	TakeRogueMiracleHandbookRewardCsReq                = 5630
	RogueWorkbenchSelectFuncScRsp                      = 5655
	RogueWorkbenchHandleFuncCsReq                      = 5696
	SyncRogueCommonPendingActionScNotify               = 5615
	GetRogueCollectionCsReq                            = 5610
	GetRogueShopBuffInfoScRsp                          = 5639
	PrepareRogueAdventureRoomCsReq                     = 5606
	FinishRogueCommonDialogueScRsp                     = 5621
	RogueWorkbenchGetInfoCsReq                         = 5642
	TakeRogueEventHandbookRewardScRsp                  = 5669
	RogueWorkbenchHandleFuncScRsp                      = 5636
	RogueGetGambleInfoCsReq                            = 5691
	SetRogueCollectionCsReq                            = 5672
	RogueNpcDisappearCsReq                             = 5654
	SyncRogueCommonDialogueDataScNotify                = 5671
	SetRogueExhibitionCsReq                            = 5661
	GetRogueAdventureRoomInfoScRsp                     = 5682
	GetRogueShopMiracleInfoCsReq                       = 5643
	BuyRogueShopMiracleCsReq                           = 5629
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5692
	CommonRogueQueryCsReq                              = 5607
	RogueWorkbenchGetInfoScRsp                         = 5683
	RogueGetGambleInfoScRsp                            = 5620
	HandleRogueCommonPendingActionCsReq                = 5647
	SyncRogueHandbookDataUpdateScNotify                = 5612
	GetRogueCommonDialogueDataCsReq                    = 5603
	RogueNpcDisappearScRsp                             = 5665
	CommonRogueComponentUpdateScNotify                 = 5686
	GetRogueExhibitionScRsp                            = 5670
	BuyRogueShopMiracleScRsp                           = 5651
	UpdateRogueAdventureRoomScoreCsReq                 = 5675
	GetRogueCommonDialogueDataScRsp                    = 5626
	GetRogueHandbookDataScRsp                          = 5646
	PrepareRogueAdventureRoomScRsp                     = 5611
	SetRogueExhibitionScRsp                            = 5679
	StopRogueAdventureRoomCsReq                        = 5609
	GetRogueExhibitionCsReq                            = 5664
	GetRogueCollectionScRsp                            = 5650
	GetRogueHandbookDataCsReq                          = 5666
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5617
	SyncRogueAdventureRoomInfoScNotify                 = 5668
	SelectRogueCommonDialogueOptionCsReq               = 5613
	ExchangeRogueBuffWithMiracleScRsp                  = 5624
	ExchangeRogueBuffWithMiracleCsReq                  = 5625
	UpdateRogueAdventureRoomScoreScRsp                 = 5640
	GetRogueShopMiracleInfoScRsp                       = 5656
	SyncRogueCommonVirtualItemInfoScNotify             = 5685
	SyncRogueCommonDialogueOptionFinishScNotify        = 5673
	SelectRogueCommonDialogueOptionScRsp               = 5614
	EnhanceCommonRogueBuffCsReq                        = 5690
	CommonRogueQueryScRsp                              = 5641
	FinishRogueCommonDialogueCsReq                     = 5622
	GetRogueShopBuffInfoCsReq                          = 5693
	StopRogueAdventureRoomScRsp                        = 5688
	BuyRogueShopBuffCsReq                              = 5633
	SetRogueCollectionScRsp                            = 5605
	RogueDoGambleCsReq                                 = 5657
	SyncRogueCommonActionResultScNotify                = 5667
	EnhanceCommonRogueBuffScRsp                        = 5628
	BuyRogueShopBuffScRsp                              = 5699
	TakeRogueEventHandbookRewardCsReq                  = 5689
	GetRogueAdventureRoomInfoCsReq                     = 5648
	RogueDoGambleScRsp                                 = 5627
	TakeRogueMiracleHandbookRewardScRsp                = 5695
	TakeRogueEndlessActivityPointRewardCsReq           = 6006
	EnterRogueEndlessActivityStageScRsp                = 6010
	RogueEndlessActivityBattleEndScNotify              = 6002
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6005
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6003
	EnterRogueEndlessActivityStageCsReq                = 6009
	GetRogueEndlessActivityDataScRsp                   = 6004
	GetRogueEndlessActivityDataCsReq                   = 6001
	TakeRogueEndlessActivityPointRewardScRsp           = 6007
	RogueModifierSelectCellCsReq                       = 5343
	RogueModifierAddNotify                             = 5311
	RogueModifierSelectCellScRsp                       = 5356
	RogueModifierStageStartNotify                      = 5333
	RogueModifierUpdateNotify                          = 5329
	RogueModifierDelNotify                             = 5351
	RogueTournSettleScRsp                              = 6043
	RogueTournLeaveScRsp                               = 6078
	RogueTournEnterLayerScRsp                          = 6058
	RogueTournStartCsReq                               = 6083
	RogueTournLeaveRogueCocoonSceneCsReq               = 6064
	RogueTournLeaveRogueCocoonSceneScRsp               = 6086
	RogueTournDeleteArchiveCsReq                       = 6016
	RogueTournLevelInfoUpdateScNotify                  = 6038
	RogueTournExpNotify                                = 6017
	RogueTournGetMiscRealTimeDataScRsp                 = 6013
	RogueTournReviveAvatarScRsp                        = 6042
	RogueTournReviveAvatarCsReq                        = 6081
	RogueTournGetAllArchiveCsReq                       = 6068
	RogueTournQueryScRsp                               = 6040
	RogueTournDeleteArchiveScRsp                       = 6035
	RogueTournReEnterRogueCocoonStageCsReq             = 6063
	RogueTournTakeExpRewardCsReq                       = 6027
	RogueTournEnterRogueCocoonSceneCsReq               = 6095
	RogueTournGetCurRogueCocoonInfoScRsp               = 6041
	RogueTournGetArchiveRepositoryScRsp                = 6031
	RogueTournRenameArchiveScRsp                       = 6088
	RogueTournResetPermanentTalentScRsp                = 6096
	RogueTournQueryCsReq                               = 6034
	RogueTournWeekChallengeUpdateScNotify              = 6032
	RogueTournGetCurRogueCocoonInfoCsReq               = 6051
	RogueTournRenameArchiveCsReq                       = 6044
	RogueTournTakeExpRewardScRsp                       = 6049
	RogueTournEnterCsReq                               = 6099
	RogueTournConfirmSettleScRsp                       = 6024
	RogueTournGetMiscRealTimeDataCsReq                 = 6079
	RogueTournEnterLayerCsReq                          = 6092
	RogueTournGetAllArchiveScRsp                       = 6018
	RogueTournReviveCostUpdateScNotify                 = 6087
	RogueTournDifficultyCompNotify                     = 6055
	RogueTournGetArchiveRepositoryCsReq                = 6082
	RogueTournEnablePermanentTalentScRsp               = 6053
	RogueTournGetPermanentTalentInfoCsReq              = 6026
	RogueTournGetPermanentTalentInfoScRsp              = 6085
	RogueTournHandBookNotify                           = 6056
	RogueTournAreaUpdateScNotify                       = 6066
	RogueTournEnablePermanentTalentCsReq               = 6022
	RogueTournEnterRoomScRsp                           = 6019
	RogueTournLeaveCsReq                               = 6059
	RogueTournEnterRoomCsReq                           = 6089
	RogueTournSettleCsReq                              = 6100
	RogueTournConfirmSettleCsReq                       = 6023
	RogueTournGetSettleInfoScRsp                       = 6036
	RogueTournReEnterRogueCocoonStageScRsp             = 6050
	RogueTournEnterRogueCocoonSceneScRsp               = 6025
	RogueTournEnterScRsp                               = 6033
	RogueTournClearArchiveNameScNotify                 = 6057
	RogueTournGetSettleInfoCsReq                       = 6077
	RogueTournStartScRsp                               = 6039
	RogueTournBattleFailSettleInfoScNotify             = 6020
	RogueTournResetPermanentTalentCsReq                = 6021
	GetRollShopInfoCsReq                               = 6919
	TakeRollShopRewardScRsp                            = 6902
	DoGachaInRollShopScRsp                             = 6920
	DoGachaInRollShopCsReq                             = 6912
	TakeRollShopRewardCsReq                            = 6904
	GetRollShopInfoScRsp                               = 6903
	InteractPropCsReq                                  = 1411
	ReturnLastTownScRsp                                = 1492
	SetSpringRecoverConfigCsReq                        = 1500
	SyncServerSceneChangeNotify                        = 1450
	SceneUpdatePositionVersionNotify                   = 1454
	GroupStateChangeScRsp                              = 1420
	SpringRecoverCsReq                                 = 1458
	GetEnteredSceneScRsp                               = 1477
	ScenePlaneEventScNotify                            = 1496
	StartTimedCocoonStageCsReq                         = 1444
	ActivateFarmElementScRsp                           = 1475
	SetGroupCustomSaveDataScRsp                        = 1426
	SceneEntityTeleportScRsp                           = 1473
	SceneCastSkillCsReq                                = 1456
	SceneEntityMoveScNotify                            = 1499
	StartTimedFarmElementScRsp                         = 1487
	GetEnteredSceneCsReq                               = 1427
	SceneReviveAfterRebattleCsReq                      = 1409
	UnlockTeleportNotify                               = 1418
	SceneCastSkillCostMpCsReq                          = 1448
	SyncEntityBuffChangeListScNotify                   = 1465
	GetSceneMapInfoScRsp                               = 1410
	SceneCastSkillMpUpdateScNotify                     = 1434
	GameplayCounterCountDownCsReq                      = 1461
	GetSceneMapInfoCsReq                               = 1436
	ActivateFarmElementCsReq                           = 1415
	EnterSceneByServerScNotify                         = 1483
	SetGroupCustomSaveDataCsReq                        = 1403
	SetClientPausedScRsp                               = 1495
	GroupStateChangeScNotify                           = 1457
	EnterSectionCsReq                                  = 1408
	SavePointsInfoNotify                               = 1449
	DeleteSummonUnitCsReq                              = 1404
	DeactivateFarmElementCsReq                         = 1469
	SceneCastSkillScRsp                                = 1493
	RefreshTriggerByClientScRsp                        = 1438
	ReEnterLastElementStageCsReq                       = 1422
	RefreshTriggerByClientScNotify                     = 1480
	SceneEntityTeleportCsReq                           = 1471
	SceneGroupRefreshScNotify                          = 1402
	EntityBindPropScRsp                                = 1412
	UpdateFloorSavedValueNotify                        = 1470
	SetCurInteractEntityCsReq                          = 1484
	GetCurSceneInfoCsReq                               = 1439
	DeactivateFarmElementScRsp                         = 1467
	StartCocoonStageCsReq                              = 1445
	HealPoolInfoNotify                                 = 1459
	SpringRecoverSingleAvatarScRsp                     = 1441
	GameplayCounterUpdateScNotify                      = 1452
	SetCurInteractEntityScRsp                          = 1423
	RecoverAllLineupScRsp                              = 1476
	UpdateMechanismBarScNotify                         = 1431
	ReturnLastTownCsReq                                = 1417
	SceneCastSkillCostMpScRsp                          = 1482
	EnteredSceneChangeScNotify                         = 1455
	EnterSectionScRsp                                  = 1481
	GetUnlockTeleportCsReq                             = 1472
	SetClientPausedCsReq                               = 1430
	SpringRecoverScRsp                                 = 1447
	SetSpringRecoverConfigScRsp                        = 1453
	StartCocoonStageScRsp                              = 1466
	GetSpringRecoverDataScRsp                          = 1485
	SceneEntityMoveScRsp                               = 1406
	GameplayCounterCountDownScRsp                      = 1479
	GroupStateChangeCsReq                              = 1491
	ReEnterLastElementStageScRsp                       = 1421
	SceneEntityMoveCsReq                               = 1468
	GameplayCounterRecoverScRsp                        = 1401
	StartTimedFarmElementCsReq                         = 1419
	StartTimedCocoonStageScRsp                         = 1474
	EnterSceneScRsp                                    = 1442
	LastSpringRefreshTimeNotify                        = 1424
	GetCurSceneInfoScRsp                               = 1429
	SceneEnterStageScRsp                               = 1428
	SceneReviveAfterRebattleScRsp                      = 1488
	EnterSceneCsReq                                    = 1486
	InteractPropScRsp                                  = 1443
	EntityBindPropCsReq                                = 1446
	SpringRecoverSingleAvatarCsReq                     = 1407
	RefreshTriggerByClientCsReq                        = 1463
	SpringRefreshScRsp                                 = 1425
	SpringRefreshCsReq                                 = 1498
	GameplayCounterRecoverCsReq                        = 1462
	DeleteSummonUnitScRsp                              = 1437
	SceneEnterStageCsReq                               = 1490
	GetUnlockTeleportScRsp                             = 1405
	UnlockedAreaMapScNotify                            = 1494
	RecoverAllLineupCsReq                              = 1478
	GetSpringRecoverDataCsReq                          = 1440
	GetServerPrefsDataCsReq                            = 6111
	UpdateServerPrefsDataCsReq                         = 6156
	UpdateServerPrefsDataScRsp                         = 6193
	GetAllServerPrefsDataCsReq                         = 6168
	GetServerPrefsDataScRsp                            = 6143
	GetAllServerPrefsDataScRsp                         = 6106
	TakeCityShopRewardCsReq                            = 1556
	CityShopInfoScNotify                               = 1539
	BuyGoodsScRsp                                      = 1543
	GetShopListCsReq                                   = 1568
	GetShopListScRsp                                   = 1506
	BuyGoodsCsReq                                      = 1511
	TakeCityShopRewardScRsp                            = 1593
	SpaceZooExchangeItemCsReq                          = 6754
	SpaceZooDeleteCatCsReq                             = 6751
	SpaceZooTakeCsReq                                  = 6748
	SpaceZooTakeScRsp                                  = 6782
	SpaceZooMutateCsReq                                = 6756
	SpaceZooMutateScRsp                                = 6793
	SpaceZooBornScRsp                                  = 6743
	SpaceZooDataCsReq                                  = 6768
	SpaceZooOpCatteryCsReq                             = 6739
	SpaceZooBornCsReq                                  = 6711
	SpaceZooExchangeItemScRsp                          = 6765
	SpaceZooOpCatteryScRsp                             = 6729
	SpaceZooDataScRsp                                  = 6706
	SpaceZooCatUpdateNotify                            = 6799
	SpaceZooDeleteCatScRsp                             = 6733
	StartStarFightLevelScRsp                           = 7170
	GetStarFightDataScRsp                              = 7164
	GetStarFightDataCsReq                              = 7161
	StartStarFightLevelCsReq                           = 7169
	StarFightDataChangeNotify                          = 7162
	StoryLineTrialAvatarChangeScNotify                 = 6239
	GetStoryLineInfoScRsp                              = 6206
	StoryLineInfoScNotify                              = 6211
	ChangeStoryLineFinishScNotify                      = 6293
	GetStoryLineInfoCsReq                              = 6268
	GetStrongChallengeActivityDataCsReq                = 6668
	StrongChallengeActivityBattleEndScNotify           = 6656
	EnterStrongChallengeActivityStageScRsp             = 6643
	EnterStrongChallengeActivityStageCsReq             = 6611
	GetStrongChallengeActivityDataScRsp                = 6606
	GetSummonActivityDataScRsp                         = 7564
	GetSummonActivityDataCsReq                         = 7561
	EnterSummonActivityStageCsReq                      = 7569
	SummonActivityBattleEndScNotify                    = 7562
	EnterSummonActivityStageScRsp                      = 7570
	SwordTrainingSetSkillTraceScRsp                    = 7494
	SwordTrainingGameSyncChangeScNotify                = 7457
	SwordTrainingMarkEndingViewedScRsp                 = 7468
	SwordTrainingExamResultConfirmScRsp                = 7478
	SwordTrainingUnlockSyncScNotify                    = 7481
	SwordTrainingStoryConfirmScRsp                     = 7475
	SwordTrainingTurnActionCsReq                       = 7467
	EnterSwordTrainingExamCsReq                        = 7462
	SwordTrainingTurnActionScRsp                       = 7461
	SwordTrainingRestoreGameScRsp                      = 7496
	SwordTrainingLearnSkillCsReq                       = 7455
	EnterSwordTrainingExamScRsp                        = 7489
	SwordTrainingExamResultConfirmCsReq                = 7464
	SwordTrainingRestoreGameCsReq                      = 7491
	GetSwordTrainingDataCsReq                          = 7474
	SwordTrainingDialogueSelectOptionScRsp             = 7484
	SwordTrainingGiveUpGameScRsp                       = 7493
	SwordTrainingResumeGameScRsp                       = 7497
	SwordTrainingGameSettleScNotify                    = 7499
	SwordTrainingMarkEndingViewedCsReq                 = 7459
	SwordTrainingDialogueSelectOptionCsReq             = 7454
	SwordTrainingStartGameScRsp                        = 7453
	SwordTrainingStoryBattleScRsp                      = 7452
	SwordTrainingSetSkillTraceCsReq                    = 7460
	SwordTrainingLearnSkillScRsp                       = 7498
	SwordTrainingActionTurnSettleScNotify              = 7479
	SwordTrainingDailyPhaseConfirmCsReq                = 7458
	SwordTrainingSelectEndingScRsp                     = 7488
	SwordTrainingDailyPhaseConfirmScRsp                = 7473
	SwordTrainingStoryBattleCsReq                      = 7490
	SwordTrainingGiveUpGameCsReq                       = 7483
	SwordTrainingStartGameCsReq                        = 7471
	SwordTrainingSelectEndingCsReq                     = 7463
	GetSwordTrainingDataScRsp                          = 7465
	SwordTrainingStoryConfirmCsReq                     = 7485
	SwordTrainingResumeGameCsReq                       = 7469
	PlayerSyncScNotify                                 = 668
	FinishFirstTalkByPerformanceNpcScRsp               = 2148
	FinishFirstTalkNpcCsReq                            = 2139
	TakeTalkRewardScRsp                                = 2143
	FinishFirstTalkNpcScRsp                            = 2129
	SelectInclinationTextCsReq                         = 2151
	GetFirstTalkNpcScRsp                               = 2193
	SelectInclinationTextScRsp                         = 2133
	GetFirstTalkNpcCsReq                               = 2156
	FinishFirstTalkByPerformanceNpcCsReq               = 2165
	GetNpcTakenRewardScRsp                             = 2106
	GetFirstTalkByPerformanceNpcScRsp                  = 2154
	GetFirstTalkByPerformanceNpcCsReq                  = 2199
	TakeTalkRewardCsReq                                = 2111
	GetNpcTakenRewardCsReq                             = 2168
	GetTelevisionActivityDataScRsp                     = 6963
	GetTelevisionActivityDataCsReq                     = 6979
	TelevisionActivityDataChangeScNotify               = 6972
	TelevisionActivityBattleEndScNotify                = 6962
	EnterTelevisionActivityStageCsReq                  = 6980
	EnterTelevisionActivityStageScRsp                  = 6964
	TextJoinQueryCsReq                                 = 3811
	TextJoinQueryScRsp                                 = 3843
	TextJoinSaveScRsp                                  = 3806
	TextJoinBatchSaveScRsp                             = 3893
	TextJoinBatchSaveCsReq                             = 3856
	TextJoinSaveCsReq                                  = 3868
	StartTrackPhotoStageScRsp                          = 7556
	SettleTrackPhotoStageCsReq                         = 7559
	QuitTrackPhotoStageCsReq                           = 7557
	GetTrackPhotoActivityDataCsReq                     = 7551
	SettleTrackPhotoStageScRsp                         = 7560
	StartTrackPhotoStageCsReq                          = 7552
	QuitTrackPhotoStageScRsp                           = 7553
	GetTrackPhotoActivityDataScRsp                     = 7554
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3751
	GetTrainVisitorBehaviorCsReq                       = 3711
	GetTrainVisitorBehaviorScRsp                       = 3743
	GetTrainVisitorRegisterScRsp                       = 3729
	TrainVisitorRewardSendNotify                       = 3793
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3733
	TrainVisitorBehaviorFinishScRsp                    = 3706
	ShowNewSupplementVisitorScRsp                      = 3754
	ShowNewSupplementVisitorCsReq                      = 3799
	TrainRefreshTimeNotify                             = 3756
	TrainVisitorBehaviorFinishCsReq                    = 3768
	GetTrainVisitorRegisterCsReq                       = 3739
	TravelBrochureApplyPasterCsReq                     = 6439
	TravelBrochureSetPageDescStatusCsReq               = 6497
	TravelBrochurePageResetScRsp                       = 6424
	TravelBrochurePageResetCsReq                       = 6425
	TravelBrochureUpdatePasterPosCsReq                 = 6499
	TravelBrochureSetCustomValueCsReq                  = 6482
	TravelBrochureGetPasterScNotify                    = 6465
	TravelBrochurePageUnlockScNotify                   = 6411
	TravelBrochureSetPageDescStatusScRsp               = 6498
	TravelBrochureSelectMessageCsReq                   = 6456
	TravelBrochureRemovePasterScRsp                    = 6433
	TravelBrochureApplyPasterScRsp                     = 6429
	TravelBrochureSetCustomValueScRsp                  = 6434
	TravelBrochureGetDataCsReq                         = 6468
	TravelBrochureSelectMessageScRsp                   = 6493
	TravelBrochureUpdatePasterPosScRsp                 = 6454
	TravelBrochureApplyPasterListCsReq                 = 6417
	TravelBrochureApplyPasterListScRsp                 = 6492
	TravelBrochureGetDataScRsp                         = 6406
	TravelBrochureRemovePasterCsReq                    = 6451
	FightTreasureDungeonMonsterScRsp                   = 4498
	UseTreasureDungeonItemScRsp                        = 4492
	UseTreasureDungeonItemCsReq                        = 4417
	QuitTreasureDungeonScRsp                           = 4428
	QuitTreasureDungeonCsReq                           = 4490
	OpenTreasureDungeonGridScRsp                       = 4434
	TreasureDungeonDataScNotify                        = 4468
	InteractTreasureDungeonGridScRsp                   = 4424
	OpenTreasureDungeonGridCsReq                       = 4482
	InteractTreasureDungeonGridCsReq                   = 4425
	FightTreasureDungeonMonsterCsReq                   = 4497
	GetTreasureDungeonActivityDataCsReq                = 4499
	GetTreasureDungeonActivityDataScRsp                = 4454
	EnterTreasureDungeonScRsp                          = 4448
	TreasureDungeonFinishScNotify                      = 4406
	EnterTreasureDungeonCsReq                          = 4465
	GetTutorialCsReq                                   = 1668
	UnlockTutorialCsReq                                = 1656
	FinishTutorialGuideCsReq                           = 1699
	GetTutorialGuideCsReq                              = 1611
	FinishTutorialCsReq                                = 1651
	UnlockTutorialGuideScRsp                           = 1629
	FinishTutorialScRsp                                = 1633
	UnlockTutorialScRsp                                = 1693
	UnlockTutorialGuideCsReq                           = 1639
	GetTutorialScRsp                                   = 1606
	GetTutorialGuideScRsp                              = 1643
	FinishTutorialGuideScRsp                           = 1654
	SetCurWaypointScRsp                                = 443
	TakeChapterRewardScRsp                             = 451
	GetChapterScRsp                                    = 493
	WaypointShowNewCsNotify                            = 439
	SetCurWaypointCsReq                                = 411
	TakeChapterRewardCsReq                             = 429
	GetWaypointCsReq                                   = 468
	GetWaypointScRsp                                   = 406
	GetChapterCsReq                                    = 456
	RestoreWolfBroGameArchiveScRsp                     = 6593
	WolfBroGamePickupBulletScRsp                       = 6582
	WolfBroGameDataChangeScNotify                      = 6599
	RestoreWolfBroGameArchiveCsReq                     = 6556
	GetWolfBroGameDataScRsp                            = 6533
	ArchiveWolfBroGameScRsp                            = 6543
	StartWolfBroGameScRsp                              = 6506
	ArchiveWolfBroGameCsReq                            = 6511
	QuitWolfBroGameCsReq                               = 6539
	WolfBroGameExplodeMonsterCsReq                     = 6598
	WolfBroGameActivateBulletCsReq                     = 6534
	QuitWolfBroGameScRsp                               = 6529
	GetWolfBroGameDataCsReq                            = 6551
	WolfBroGameUseBulletScRsp                          = 6565
	WolfBroGamePickupBulletCsReq                       = 6548
	StartWolfBroGameCsReq                              = 6568
	WolfBroGameActivateBulletScRsp                     = 6597
	WolfBroGameExplodeMonsterScRsp                     = 6525
	WolfBroGameUseBulletCsReq                          = 6554
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(SubmitMonsterResearchActivityMaterialCsReq, func() any { return new(proto.SubmitMonsterResearchActivityMaterialCsReq) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(GetMonsterResearchActivityDataScRsp, func() any { return new(proto.GetMonsterResearchActivityDataScRsp) })
	c.regMsg(TakeMonsterResearchActivityRewardCsReq, func() any { return new(proto.TakeMonsterResearchActivityRewardCsReq) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(TakeMonsterResearchActivityRewardScRsp, func() any { return new(proto.TakeMonsterResearchActivityRewardScRsp) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(GetMonsterResearchActivityDataCsReq, func() any { return new(proto.GetMonsterResearchActivityDataCsReq) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(SubmitMonsterResearchActivityMaterialScRsp, func() any { return new(proto.SubmitMonsterResearchActivityMaterialScRsp) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(LogisticsGameCsReq, func() any { return new(proto.LogisticsGameCsReq) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(ActivityRaidPlacingGameCsReq, func() any { return new(proto.ActivityRaidPlacingGameCsReq) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(ActivityRaidPlacingGameScRsp, func() any { return new(proto.ActivityRaidPlacingGameScRsp) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(GetUpdatedArchiveDataScRsp, func() any { return new(proto.GetUpdatedArchiveDataScRsp) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(GetUpdatedArchiveDataCsReq, func() any { return new(proto.GetUpdatedArchiveDataCsReq) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(MarkAvatarScRsp, func() any { return new(proto.MarkAvatarScRsp) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(MarkAvatarCsReq, func() any { return new(proto.MarkAvatarCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(ReBattleAfterBattleLoseCsNotify, func() any { return new(proto.ReBattleAfterBattleLoseCsNotify) })
	c.regMsg(ServerSimulateBattleFinishScNotify, func() any { return new(proto.ServerSimulateBattleFinishScNotify) })
	c.regMsg(QuitBattleScRsp, func() any { return new(proto.QuitBattleScRsp) })
	c.regMsg(QuitBattleCsReq, func() any { return new(proto.QuitBattleCsReq) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(BattleLogReportScRsp, func() any { return new(proto.BattleLogReportScRsp) })
	c.regMsg(QuitBattleScNotify, func() any { return new(proto.QuitBattleScNotify) })
	c.regMsg(RebattleByClientCsNotify, func() any { return new(proto.RebattleByClientCsNotify) })
	c.regMsg(BattleLogReportCsReq, func() any { return new(proto.BattleLogReportCsReq) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(GetBattleCollegeDataCsReq, func() any { return new(proto.GetBattleCollegeDataCsReq) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(BoxingClubRewardScNotify, func() any { return new(proto.BoxingClubRewardScNotify) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffCsReq, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffCsReq) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(GiveUpBoxingClubChallengeScRsp, func() any { return new(proto.GiveUpBoxingClubChallengeScRsp) })
	c.regMsg(GetBoxingClubInfoCsReq, func() any { return new(proto.GetBoxingClubInfoCsReq) })
	c.regMsg(MatchBoxingClubOpponentCsReq, func() any { return new(proto.MatchBoxingClubOpponentCsReq) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(ChooseBoxingClubResonanceCsReq, func() any { return new(proto.ChooseBoxingClubResonanceCsReq) })
	c.regMsg(BoxingClubChallengeUpdateScNotify, func() any { return new(proto.BoxingClubChallengeUpdateScNotify) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(StartBoxingClubBattleCsReq, func() any { return new(proto.StartBoxingClubBattleCsReq) })
	c.regMsg(GiveUpBoxingClubChallengeCsReq, func() any { return new(proto.GiveUpBoxingClubChallengeCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupCsReq, func() any { return new(proto.SetBoxingClubResonanceLineupCsReq) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(GetChallengeGroupStatisticsCsReq, func() any { return new(proto.GetChallengeGroupStatisticsCsReq) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(StartPartialChallengeCsReq, func() any { return new(proto.StartPartialChallengeCsReq) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(EnterChallengeNextPhaseScRsp, func() any { return new(proto.EnterChallengeNextPhaseScRsp) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(RestartChallengePhaseCsReq, func() any { return new(proto.RestartChallengePhaseCsReq) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(StartPartialChallengeScRsp, func() any { return new(proto.StartPartialChallengeScRsp) })
	c.regMsg(EnterChallengeNextPhaseCsReq, func() any { return new(proto.EnterChallengeNextPhaseCsReq) })
	c.regMsg(RestartChallengePhaseScRsp, func() any { return new(proto.RestartChallengePhaseScRsp) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(GetChallengeGroupStatisticsScRsp, func() any { return new(proto.GetChallengeGroupStatisticsScRsp) })
	c.regMsg(ChallengeBossPhaseSettleNotify, func() any { return new(proto.ChallengeBossPhaseSettleNotify) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoScRsp, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoScRsp) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(ChessRogueNousEnableRogueTalentScRsp, func() any { return new(proto.ChessRogueNousEnableRogueTalentScRsp) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(ChessRogueNousEnableRogueTalentCsReq, func() any { return new(proto.ChessRogueNousEnableRogueTalentCsReq) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoCsReq, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoCsReq) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(ClockParkHandleWaitOperationScRsp, func() any { return new(proto.ClockParkHandleWaitOperationScRsp) })
	c.regMsg(ClockParkQuitScriptCsReq, func() any { return new(proto.ClockParkQuitScriptCsReq) })
	c.regMsg(ClockParkQuitScriptScRsp, func() any { return new(proto.ClockParkQuitScriptScRsp) })
	c.regMsg(ClockParkStartScriptCsReq, func() any { return new(proto.ClockParkStartScriptCsReq) })
	c.regMsg(ClockParkGetOngoingScriptInfoCsReq, func() any { return new(proto.ClockParkGetOngoingScriptInfoCsReq) })
	c.regMsg(ClockParkGetInfoScRsp, func() any { return new(proto.ClockParkGetInfoScRsp) })
	c.regMsg(ClockParkUnlockTalentScRsp, func() any { return new(proto.ClockParkUnlockTalentScRsp) })
	c.regMsg(ClockParkFinishScriptScNotify, func() any { return new(proto.ClockParkFinishScriptScNotify) })
	c.regMsg(ClockParkBattleEndScNotify, func() any { return new(proto.ClockParkBattleEndScNotify) })
	c.regMsg(ClockParkGetInfoCsReq, func() any { return new(proto.ClockParkGetInfoCsReq) })
	c.regMsg(ClockParkUseBuffScRsp, func() any { return new(proto.ClockParkUseBuffScRsp) })
	c.regMsg(ClockParkHandleWaitOperationCsReq, func() any { return new(proto.ClockParkHandleWaitOperationCsReq) })
	c.regMsg(ClockParkStartScriptScRsp, func() any { return new(proto.ClockParkStartScriptScRsp) })
	c.regMsg(ClockParkGetOngoingScriptInfoScRsp, func() any { return new(proto.ClockParkGetOngoingScriptInfoScRsp) })
	c.regMsg(ClockParkUnlockTalentCsReq, func() any { return new(proto.ClockParkUnlockTalentCsReq) })
	c.regMsg(ClockParkUseBuffCsReq, func() any { return new(proto.ClockParkUseBuffCsReq) })
	c.regMsg(ContentPackageGetDataScRsp, func() any { return new(proto.ContentPackageGetDataScRsp) })
	c.regMsg(ContentPackageUnlockCsReq, func() any { return new(proto.ContentPackageUnlockCsReq) })
	c.regMsg(ContentPackageGetDataCsReq, func() any { return new(proto.ContentPackageGetDataCsReq) })
	c.regMsg(ContentPackageTransferScNotify, func() any { return new(proto.ContentPackageTransferScNotify) })
	c.regMsg(ContentPackageUnlockScRsp, func() any { return new(proto.ContentPackageUnlockScRsp) })
	c.regMsg(ContentPackageSyncDataScNotify, func() any { return new(proto.ContentPackageSyncDataScNotify) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(DrinkMakerDayEndScNotify, func() any { return new(proto.DrinkMakerDayEndScNotify) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(DrinkMakerChallengeCsReq, func() any { return new(proto.DrinkMakerChallengeCsReq) })
	c.regMsg(DrinkMakerUpdateTipsNotify, func() any { return new(proto.DrinkMakerUpdateTipsNotify) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(DrinkMakerChallengeScRsp, func() any { return new(proto.DrinkMakerChallengeScRsp) })
	c.regMsg(EvolveBuildShopAbilityUpScRsp, func() any { return new(proto.EvolveBuildShopAbilityUpScRsp) })
	c.regMsg(EvolveBuildShopAbilityUpCsReq, func() any { return new(proto.EvolveBuildShopAbilityUpCsReq) })
	c.regMsg(EvolveBuildCoinNotify, func() any { return new(proto.EvolveBuildCoinNotify) })
	c.regMsg(EvolveBuildLeaveScRsp, func() any { return new(proto.EvolveBuildLeaveScRsp) })
	c.regMsg(EvolveBuildStartStageCsReq, func() any { return new(proto.EvolveBuildStartStageCsReq) })
	c.regMsg(EvolveBuildUnlockInfoNotify, func() any { return new(proto.EvolveBuildUnlockInfoNotify) })
	c.regMsg(EvolveBuildStartLevelScRsp, func() any { return new(proto.EvolveBuildStartLevelScRsp) })
	c.regMsg(EvolveBuildReRandomStageCsReq, func() any { return new(proto.EvolveBuildReRandomStageCsReq) })
	c.regMsg(EvolveBuildShopAbilityResetCsReq, func() any { return new(proto.EvolveBuildShopAbilityResetCsReq) })
	c.regMsg(EvolveBuildLeaveCsReq, func() any { return new(proto.EvolveBuildLeaveCsReq) })
	c.regMsg(EvolveBuildTakeExpRewardCsReq, func() any { return new(proto.EvolveBuildTakeExpRewardCsReq) })
	c.regMsg(EvolveBuildQueryInfoCsReq, func() any { return new(proto.EvolveBuildQueryInfoCsReq) })
	c.regMsg(EvolveBuildReRandomStageScRsp, func() any { return new(proto.EvolveBuildReRandomStageScRsp) })
	c.regMsg(EvolveBuildTakeExpRewardScRsp, func() any { return new(proto.EvolveBuildTakeExpRewardScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownScRsp, func() any { return new(proto.EvolveBuildShopAbilityDownScRsp) })
	c.regMsg(EvolveBuildStartLevelCsReq, func() any { return new(proto.EvolveBuildStartLevelCsReq) })
	c.regMsg(EvolveBuildFinishScNotify, func() any { return new(proto.EvolveBuildFinishScNotify) })
	c.regMsg(EvolveBuildStartStageScRsp, func() any { return new(proto.EvolveBuildStartStageScRsp) })
	c.regMsg(EvolveBuildGiveupCsReq, func() any { return new(proto.EvolveBuildGiveupCsReq) })
	c.regMsg(EvolveBuildQueryInfoScRsp, func() any { return new(proto.EvolveBuildQueryInfoScRsp) })
	c.regMsg(EvolveBuildGiveupScRsp, func() any { return new(proto.EvolveBuildGiveupScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownCsReq, func() any { return new(proto.EvolveBuildShopAbilityDownCsReq) })
	c.regMsg(EvolveBuildShopAbilityResetScRsp, func() any { return new(proto.EvolveBuildShopAbilityResetScRsp) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(CancelExpeditionScRsp, func() any { return new(proto.CancelExpeditionScRsp) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(FightHeartBeatScRsp, func() any { return new(proto.FightHeartBeatScRsp) })
	c.regMsg(FightSessionStopScNotify, func() any { return new(proto.FightSessionStopScNotify) })
	c.regMsg(FightEnterCsReq, func() any { return new(proto.FightEnterCsReq) })
	c.regMsg(FightHeartBeatCsReq, func() any { return new(proto.FightHeartBeatCsReq) })
	c.regMsg(FightKickOutScNotify, func() any { return new(proto.FightKickOutScNotify) })
	c.regMsg(FightEnterScRsp, func() any { return new(proto.FightEnterScRsp) })
	c.regMsg(FightLeaveScNotify, func() any { return new(proto.FightLeaveScNotify) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(TakeFightActivityRewardScRsp, func() any { return new(proto.TakeFightActivityRewardScRsp) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(EnterFightActivityStageScRsp, func() any { return new(proto.EnterFightActivityStageScRsp) })
	c.regMsg(TakeFightActivityRewardCsReq, func() any { return new(proto.TakeFightActivityRewardCsReq) })
	c.regMsg(GetFightFestDataScRsp, func() any { return new(proto.GetFightFestDataScRsp) })
	c.regMsg(FightFestUpdateCoinNotify, func() any { return new(proto.FightFestUpdateCoinNotify) })
	c.regMsg(FightFestScoreUpdateNotify, func() any { return new(proto.FightFestScoreUpdateNotify) })
	c.regMsg(StartFightFestScRsp, func() any { return new(proto.StartFightFestScRsp) })
	c.regMsg(GetFightFestDataCsReq, func() any { return new(proto.GetFightFestDataCsReq) })
	c.regMsg(FightFestUpdateChallengeRecordNotify, func() any { return new(proto.FightFestUpdateChallengeRecordNotify) })
	c.regMsg(StartFightFestCsReq, func() any { return new(proto.StartFightFestCsReq) })
	c.regMsg(FightFestUnlockSkillNotify, func() any { return new(proto.FightFestUnlockSkillNotify) })
	c.regMsg(FightMatch3ForceUpdateNotify, func() any { return new(proto.FightMatch3ForceUpdateNotify) })
	c.regMsg(FightMatch3SwapCsReq, func() any { return new(proto.FightMatch3SwapCsReq) })
	c.regMsg(FightMatch3StartCountDownScNotify, func() any { return new(proto.FightMatch3StartCountDownScNotify) })
	c.regMsg(FightMatch3ChatScNotify, func() any { return new(proto.FightMatch3ChatScNotify) })
	c.regMsg(FightMatch3DataCsReq, func() any { return new(proto.FightMatch3DataCsReq) })
	c.regMsg(FightMatch3ChatCsReq, func() any { return new(proto.FightMatch3ChatCsReq) })
	c.regMsg(FightMatch3TurnStartScNotify, func() any { return new(proto.FightMatch3TurnStartScNotify) })
	c.regMsg(FightMatch3SwapScRsp, func() any { return new(proto.FightMatch3SwapScRsp) })
	c.regMsg(FightMatch3DataScRsp, func() any { return new(proto.FightMatch3DataScRsp) })
	c.regMsg(FightMatch3OpponentDataScNotify, func() any { return new(proto.FightMatch3OpponentDataScNotify) })
	c.regMsg(FightMatch3TurnEndScNotify, func() any { return new(proto.FightMatch3TurnEndScNotify) })
	c.regMsg(FightMatch3ChatScRsp, func() any { return new(proto.FightMatch3ChatScRsp) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(GetFriendBattleRecordDetailCsReq, func() any { return new(proto.GetFriendBattleRecordDetailCsReq) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(NewAssistHistoryNotify, func() any { return new(proto.NewAssistHistoryNotify) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(TakeAssistRewardScRsp, func() any { return new(proto.TakeAssistRewardScRsp) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(DeleteFriendScRsp, func() any { return new(proto.DeleteFriendScRsp) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(CurAssistChangedNotify, func() any { return new(proto.CurAssistChangedNotify) })
	c.regMsg(TakeAssistRewardCsReq, func() any { return new(proto.TakeAssistRewardCsReq) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(GetFriendAssistListScRsp, func() any { return new(proto.GetFriendAssistListScRsp) })
	c.regMsg(GetFriendDevelopmentInfoCsReq, func() any { return new(proto.GetFriendDevelopmentInfoCsReq) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(GetFriendAssistListCsReq, func() any { return new(proto.GetFriendAssistListCsReq) })
	c.regMsg(GetFriendChallengeLineupCsReq, func() any { return new(proto.GetFriendChallengeLineupCsReq) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(SetFriendMarkCsReq, func() any { return new(proto.SetFriendMarkCsReq) })
	c.regMsg(GetFriendChallengeDetailCsReq, func() any { return new(proto.GetFriendChallengeDetailCsReq) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(GetFriendDevelopmentInfoScRsp, func() any { return new(proto.GetFriendDevelopmentInfoScRsp) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(GetFriendChallengeDetailScRsp, func() any { return new(proto.GetFriendChallengeDetailScRsp) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(GetFriendChallengeLineupScRsp, func() any { return new(proto.GetFriendChallengeLineupScRsp) })
	c.regMsg(GetFriendBattleRecordDetailScRsp, func() any { return new(proto.GetFriendBattleRecordDetailScRsp) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(SetForbidOtherApplyFriendScRsp, func() any { return new(proto.SetForbidOtherApplyFriendScRsp) })
	c.regMsg(SetFriendMarkScRsp, func() any { return new(proto.SetFriendMarkScRsp) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(HeliobusStartRaidCsReq, func() any { return new(proto.HeliobusStartRaidCsReq) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(HeliobusLineupUpdateScNotify, func() any { return new(proto.HeliobusLineupUpdateScNotify) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(HeliobusStartRaidScRsp, func() any { return new(proto.HeliobusStartRaidScRsp) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(AddEquipmentScNotify, func() any { return new(proto.AddEquipmentScNotify) })
	c.regMsg(DeleteRelicFilterPlanCsReq, func() any { return new(proto.DeleteRelicFilterPlanCsReq) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(AddRelicFilterPlanScRsp, func() any { return new(proto.AddRelicFilterPlanScRsp) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(DeleteRelicFilterPlanScRsp, func() any { return new(proto.DeleteRelicFilterPlanScRsp) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(ModifyRelicFilterPlanCsReq, func() any { return new(proto.ModifyRelicFilterPlanCsReq) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(AddRelicFilterPlanCsReq, func() any { return new(proto.AddRelicFilterPlanCsReq) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(ModifyRelicFilterPlanScRsp, func() any { return new(proto.ModifyRelicFilterPlanScRsp) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(MarkRelicFilterPlanCsReq, func() any { return new(proto.MarkRelicFilterPlanCsReq) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(GetRelicFilterPlanScRsp, func() any { return new(proto.GetRelicFilterPlanScRsp) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(GetRelicFilterPlanCsReq, func() any { return new(proto.GetRelicFilterPlanCsReq) })
	c.regMsg(MarkRelicFilterPlanScRsp, func() any { return new(proto.MarkRelicFilterPlanScRsp) })
	c.regMsg(RelicFilterPlanClearNameScNotify, func() any { return new(proto.RelicFilterPlanClearNameScNotify) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(GetStageLineupCsReq, func() any { return new(proto.GetStageLineupCsReq) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(LobbyQuitCsReq, func() any { return new(proto.LobbyQuitCsReq) })
	c.regMsg(LobbyInviteScRsp, func() any { return new(proto.LobbyInviteScRsp) })
	c.regMsg(LobbyBeginCsReq, func() any { return new(proto.LobbyBeginCsReq) })
	c.regMsg(LobbyJoinScRsp, func() any { return new(proto.LobbyJoinScRsp) })
	c.regMsg(LobbyJoinCsReq, func() any { return new(proto.LobbyJoinCsReq) })
	c.regMsg(LobbyModifyPlayerInfoScRsp, func() any { return new(proto.LobbyModifyPlayerInfoScRsp) })
	c.regMsg(LobbyKickOutCsReq, func() any { return new(proto.LobbyKickOutCsReq) })
	c.regMsg(LobbyGetInfoScRsp, func() any { return new(proto.LobbyGetInfoScRsp) })
	c.regMsg(LobbyKickOutScRsp, func() any { return new(proto.LobbyKickOutScRsp) })
	c.regMsg(LobbyGetInfoCsReq, func() any { return new(proto.LobbyGetInfoCsReq) })
	c.regMsg(LobbyBeginScRsp, func() any { return new(proto.LobbyBeginScRsp) })
	c.regMsg(LobbyCreateScRsp, func() any { return new(proto.LobbyCreateScRsp) })
	c.regMsg(LobbyCreateCsReq, func() any { return new(proto.LobbyCreateCsReq) })
	c.regMsg(LobbySyncInfoScNotify, func() any { return new(proto.LobbySyncInfoScNotify) })
	c.regMsg(LobbyInviteCsReq, func() any { return new(proto.LobbyInviteCsReq) })
	c.regMsg(LobbyModifyPlayerInfoCsReq, func() any { return new(proto.LobbyModifyPlayerInfoCsReq) })
	c.regMsg(LobbyQuitScRsp, func() any { return new(proto.LobbyQuitScRsp) })
	c.regMsg(LobbyInviteScNotify, func() any { return new(proto.LobbyInviteScNotify) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(UpdateRotaterScNotify, func() any { return new(proto.UpdateRotaterScNotify) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(GetMapRotationDataCsReq, func() any { return new(proto.GetMapRotationDataCsReq) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(UpdateMapRotationDataScNotify, func() any { return new(proto.UpdateMapRotationDataScNotify) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(RemoveRotaterScRsp, func() any { return new(proto.RemoveRotaterScRsp) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(RemoveRotaterCsReq, func() any { return new(proto.RemoveRotaterCsReq) })
	c.regMsg(MatchResultScNotify, func() any { return new(proto.MatchResultScNotify) })
	c.regMsg(StartMatchScRsp, func() any { return new(proto.StartMatchScRsp) })
	c.regMsg(CancelMatchScRsp, func() any { return new(proto.CancelMatchScRsp) })
	c.regMsg(StartMatchCsReq, func() any { return new(proto.StartMatchCsReq) })
	c.regMsg(GetCrossInfoCsReq, func() any { return new(proto.GetCrossInfoCsReq) })
	c.regMsg(CancelMatchCsReq, func() any { return new(proto.CancelMatchCsReq) })
	c.regMsg(GetCrossInfoScRsp, func() any { return new(proto.GetCrossInfoScRsp) })
	c.regMsg(MatchThreeLevelEndCsReq, func() any { return new(proto.MatchThreeLevelEndCsReq) })
	c.regMsg(MatchThreeGetDataCsReq, func() any { return new(proto.MatchThreeGetDataCsReq) })
	c.regMsg(MatchThreeSetBirdPosCsReq, func() any { return new(proto.MatchThreeSetBirdPosCsReq) })
	c.regMsg(MatchThreeLevelEndScRsp, func() any { return new(proto.MatchThreeLevelEndScRsp) })
	c.regMsg(MatchThreeSetBirdPosScRsp, func() any { return new(proto.MatchThreeSetBirdPosScRsp) })
	c.regMsg(MatchThreeGetDataScRsp, func() any { return new(proto.MatchThreeGetDataScRsp) })
	c.regMsg(MatchThreeSyncDataScNotify, func() any { return new(proto.MatchThreeSyncDataScNotify) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(DifficultyAdjustmentUpdateDataScRsp, func() any { return new(proto.DifficultyAdjustmentUpdateDataScRsp) })
	c.regMsg(SubmitOrigamiItemCsReq, func() any { return new(proto.SubmitOrigamiItemCsReq) })
	c.regMsg(DifficultyAdjustmentUpdateDataCsReq, func() any { return new(proto.DifficultyAdjustmentUpdateDataCsReq) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(GetGunPlayDataScRsp, func() any { return new(proto.GetGunPlayDataScRsp) })
	c.regMsg(DifficultyAdjustmentGetDataCsReq, func() any { return new(proto.DifficultyAdjustmentGetDataCsReq) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(DifficultyAdjustmentGetDataScRsp, func() any { return new(proto.DifficultyAdjustmentGetDataScRsp) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(UpdateGunPlayDataScRsp, func() any { return new(proto.UpdateGunPlayDataScRsp) })
	c.regMsg(GetGunPlayDataCsReq, func() any { return new(proto.GetGunPlayDataCsReq) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(UpdateMovieRacingDataScRsp, func() any { return new(proto.UpdateMovieRacingDataScRsp) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(UpdateGunPlayDataCsReq, func() any { return new(proto.UpdateGunPlayDataCsReq) })
	c.regMsg(UpdateMovieRacingDataCsReq, func() any { return new(proto.UpdateMovieRacingDataCsReq) })
	c.regMsg(SubmitOrigamiItemScRsp, func() any { return new(proto.SubmitOrigamiItemScRsp) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(DailyTaskDataScNotify, func() any { return new(proto.DailyTaskDataScNotify) })
	c.regMsg(AcceptMissionEventScRsp, func() any { return new(proto.AcceptMissionEventScRsp) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(InterruptMissionEventCsReq, func() any { return new(proto.InterruptMissionEventCsReq) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(MissionEventRewardScNotify, func() any { return new(proto.MissionEventRewardScNotify) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
	c.regMsg(AcceptMissionEventCsReq, func() any { return new(proto.AcceptMissionEventCsReq) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(MonopolyMoveScRsp, func() any { return new(proto.MonopolyMoveScRsp) })
	c.regMsg(GetMonopolyFriendRankingListCsReq, func() any { return new(proto.GetMonopolyFriendRankingListCsReq) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(MonopolyGameGachaCsReq, func() any { return new(proto.MonopolyGameGachaCsReq) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(MonopolyGuessDrawScNotify, func() any { return new(proto.MonopolyGuessDrawScNotify) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(GetMonopolyInfoCsReq, func() any { return new(proto.GetMonopolyInfoCsReq) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(MonopolyGameRaiseRatioCsReq, func() any { return new(proto.MonopolyGameRaiseRatioCsReq) })
	c.regMsg(MonopolySelectOptionCsReq, func() any { return new(proto.MonopolySelectOptionCsReq) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(DailyFirstEnterMonopolyActivityScRsp, func() any { return new(proto.DailyFirstEnterMonopolyActivityScRsp) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(MonopolyGuessBuyInformationCsReq, func() any { return new(proto.MonopolyGuessBuyInformationCsReq) })
	c.regMsg(MonopolyGiveUpCurContentCsReq, func() any { return new(proto.MonopolyGiveUpCurContentCsReq) })
	c.regMsg(MonopolyActionResultScNotify, func() any { return new(proto.MonopolyActionResultScNotify) })
	c.regMsg(MonopolyMoveCsReq, func() any { return new(proto.MonopolyMoveCsReq) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(MonopolyLikeCsReq, func() any { return new(proto.MonopolyLikeCsReq) })
	c.regMsg(MonopolyClickMbtiReportCsReq, func() any { return new(proto.MonopolyClickMbtiReportCsReq) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(MonopolyGameBingoFlipCardScRsp, func() any { return new(proto.MonopolyGameBingoFlipCardScRsp) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(MonopolyGuessChooseScRsp, func() any { return new(proto.MonopolyGuessChooseScRsp) })
	c.regMsg(MonopolyGiveUpCurContentScRsp, func() any { return new(proto.MonopolyGiveUpCurContentScRsp) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(GetMonopolyInfoScRsp, func() any { return new(proto.GetMonopolyInfoScRsp) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(MonopolyDailySettleScNotify, func() any { return new(proto.MonopolyDailySettleScNotify) })
	c.regMsg(GetMonopolyDailyReportScRsp, func() any { return new(proto.GetMonopolyDailyReportScRsp) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(MonopolyGameSettleScNotify, func() any { return new(proto.MonopolyGameSettleScNotify) })
	c.regMsg(MonopolyEventLoadUpdateScNotify, func() any { return new(proto.MonopolyEventLoadUpdateScNotify) })
	c.regMsg(MonopolyAcceptQuizCsReq, func() any { return new(proto.MonopolyAcceptQuizCsReq) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(MonopolyGetDailyInitItemCsReq, func() any { return new(proto.MonopolyGetDailyInitItemCsReq) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(MonopolyLikeScNotify, func() any { return new(proto.MonopolyLikeScNotify) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(MonopolySocialEventEffectScNotify, func() any { return new(proto.MonopolySocialEventEffectScNotify) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(MonopolyLikeScRsp, func() any { return new(proto.MonopolyLikeScRsp) })
	c.regMsg(MonopolyClickCellScRsp, func() any { return new(proto.MonopolyClickCellScRsp) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(MonopolyGameBingoFlipCardCsReq, func() any { return new(proto.MonopolyGameBingoFlipCardCsReq) })
	c.regMsg(MonopolyClickMbtiReportScRsp, func() any { return new(proto.MonopolyClickMbtiReportScRsp) })
	c.regMsg(MonopolyGetDailyInitItemScRsp, func() any { return new(proto.MonopolyGetDailyInitItemScRsp) })
	c.regMsg(MonopolyEventSelectFriendCsReq, func() any { return new(proto.MonopolyEventSelectFriendCsReq) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(MonopolyCheatDiceScRsp, func() any { return new(proto.MonopolyCheatDiceScRsp) })
	c.regMsg(GetMbtiReportScRsp, func() any { return new(proto.GetMbtiReportScRsp) })
	c.regMsg(MonopolyConditionUpdateScNotify, func() any { return new(proto.MonopolyConditionUpdateScNotify) })
	c.regMsg(MonopolySelectOptionScRsp, func() any { return new(proto.MonopolySelectOptionScRsp) })
	c.regMsg(MonopolyGameCreateScNotify, func() any { return new(proto.MonopolyGameCreateScNotify) })
	c.regMsg(MonopolyGuessBuyInformationScRsp, func() any { return new(proto.MonopolyGuessBuyInformationScRsp) })
	c.regMsg(MonopolyQuizDurationChangeScNotify, func() any { return new(proto.MonopolyQuizDurationChangeScNotify) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(GetMbtiReportCsReq, func() any { return new(proto.GetMbtiReportCsReq) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(MonopolyAcceptQuizScRsp, func() any { return new(proto.MonopolyAcceptQuizScRsp) })
	c.regMsg(GetMonopolyDailyReportCsReq, func() any { return new(proto.GetMonopolyDailyReportCsReq) })
	c.regMsg(MonopolySttUpdateScNotify, func() any { return new(proto.MonopolySttUpdateScNotify) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(MonopolyGuessChooseCsReq, func() any { return new(proto.MonopolyGuessChooseCsReq) })
	c.regMsg(MonopolyEventSelectFriendScRsp, func() any { return new(proto.MonopolyEventSelectFriendScRsp) })
	c.regMsg(DailyFirstEnterMonopolyActivityCsReq, func() any { return new(proto.DailyFirstEnterMonopolyActivityCsReq) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(MonopolyClickCellCsReq, func() any { return new(proto.MonopolyClickCellCsReq) })
	c.regMsg(MultiplayerFightGiveUpCsReq, func() any { return new(proto.MultiplayerFightGiveUpCsReq) })
	c.regMsg(MultiplayerGetFightGateScRsp, func() any { return new(proto.MultiplayerGetFightGateScRsp) })
	c.regMsg(MultiplayerFightGameFinishScNotify, func() any { return new(proto.MultiplayerFightGameFinishScNotify) })
	c.regMsg(MultiplayerFightGameStateCsReq, func() any { return new(proto.MultiplayerFightGameStateCsReq) })
	c.regMsg(MultiplayerFightGameStateScRsp, func() any { return new(proto.MultiplayerFightGameStateScRsp) })
	c.regMsg(MultiplayerFightGameStartScNotify, func() any { return new(proto.MultiplayerFightGameStartScNotify) })
	c.regMsg(MultiplayerMatch3FinishScNotify, func() any { return new(proto.MultiplayerMatch3FinishScNotify) })
	c.regMsg(MultiplayerFightGiveUpScRsp, func() any { return new(proto.MultiplayerFightGiveUpScRsp) })
	c.regMsg(MultiplayerGetFightGateCsReq, func() any { return new(proto.MultiplayerGetFightGateCsReq) })
	c.regMsg(MultipleDropInfoNotify, func() any { return new(proto.MultipleDropInfoNotify) })
	c.regMsg(GetMultipleDropInfoCsReq, func() any { return new(proto.GetMultipleDropInfoCsReq) })
	c.regMsg(GetMultipleDropInfoScRsp, func() any { return new(proto.GetMultipleDropInfoScRsp) })
	c.regMsg(MultipleDropInfoScNotify, func() any { return new(proto.MultipleDropInfoScNotify) })
	c.regMsg(GetPlayerReturnMultiDropInfoCsReq, func() any { return new(proto.GetPlayerReturnMultiDropInfoCsReq) })
	c.regMsg(GetPlayerReturnMultiDropInfoScRsp, func() any { return new(proto.GetPlayerReturnMultiDropInfoScRsp) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(MuseumTargetMissionFinishNotify, func() any { return new(proto.MuseumTargetMissionFinishNotify) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(MuseumTargetRewardNotify, func() any { return new(proto.MuseumTargetRewardNotify) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(MuseumTakeCollectRewardScRsp, func() any { return new(proto.MuseumTakeCollectRewardScRsp) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(MuseumTakeCollectRewardCsReq, func() any { return new(proto.MuseumTakeCollectRewardCsReq) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(TakeOfferingRewardCsReq, func() any { return new(proto.TakeOfferingRewardCsReq) })
	c.regMsg(TakeOfferingRewardScRsp, func() any { return new(proto.TakeOfferingRewardScRsp) })
	c.regMsg(SubmitOfferingItemScRsp, func() any { return new(proto.SubmitOfferingItemScRsp) })
	c.regMsg(SubmitOfferingItemCsReq, func() any { return new(proto.SubmitOfferingItemCsReq) })
	c.regMsg(OfferingInfoScNotify, func() any { return new(proto.OfferingInfoScNotify) })
	c.regMsg(GetOfferingInfoCsReq, func() any { return new(proto.GetOfferingInfoCsReq) })
	c.regMsg(AcceptedPamMissionExpireCsReq, func() any { return new(proto.AcceptedPamMissionExpireCsReq) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(RecallPetScRsp, func() any { return new(proto.RecallPetScRsp) })
	c.regMsg(RecallPetCsReq, func() any { return new(proto.RecallPetCsReq) })
	c.regMsg(CurPetChangedScNotify, func() any { return new(proto.CurPetChangedScNotify) })
	c.regMsg(GetPetDataCsReq, func() any { return new(proto.GetPetDataCsReq) })
	c.regMsg(GetPetDataScRsp, func() any { return new(proto.GetPetDataScRsp) })
	c.regMsg(SummonPetScRsp, func() any { return new(proto.SummonPetScRsp) })
	c.regMsg(SummonPetCsReq, func() any { return new(proto.SummonPetCsReq) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(ExchangeStaminaCsReq, func() any { return new(proto.ExchangeStaminaCsReq) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(UnlockAvatarPathCsReq, func() any { return new(proto.UnlockAvatarPathCsReq) })
	c.regMsg(UpdatePlayerSettingScRsp, func() any { return new(proto.UpdatePlayerSettingScRsp) })
	c.regMsg(SetMultipleAvatarPathsCsReq, func() any { return new(proto.SetMultipleAvatarPathsCsReq) })
	c.regMsg(GetLevelRewardScRsp, func() any { return new(proto.GetLevelRewardScRsp) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(GetMultiPathAvatarInfoScRsp, func() any { return new(proto.GetMultiPathAvatarInfoScRsp) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(ClientDownloadDataScNotify, func() any { return new(proto.ClientDownloadDataScNotify) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(AceAntiCheaterCsReq, func() any { return new(proto.AceAntiCheaterCsReq) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(UpdatePsnSettingsInfoCsReq, func() any { return new(proto.UpdatePsnSettingsInfoCsReq) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(GetAuthkeyScRsp, func() any { return new(proto.GetAuthkeyScRsp) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(SetRedPointStatusScNotify, func() any { return new(proto.SetRedPointStatusScNotify) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(UnlockAvatarPathScRsp, func() any { return new(proto.UnlockAvatarPathScRsp) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(SetAvatarPathCsReq, func() any { return new(proto.SetAvatarPathCsReq) })
	c.regMsg(GetMultiPathAvatarInfoCsReq, func() any { return new(proto.GetMultiPathAvatarInfoCsReq) })
	c.regMsg(AvatarPathChangedNotify, func() any { return new(proto.AvatarPathChangedNotify) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(RegionStopScNotify, func() any { return new(proto.RegionStopScNotify) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(SetAvatarPathScRsp, func() any { return new(proto.SetAvatarPathScRsp) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(ClientObjUploadCsReq, func() any { return new(proto.ClientObjUploadCsReq) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(UpdatePsnSettingsInfoScRsp, func() any { return new(proto.UpdatePsnSettingsInfoScRsp) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(ClientObjUploadScRsp, func() any { return new(proto.ClientObjUploadScRsp) })
	c.regMsg(SetMultipleAvatarPathsScRsp, func() any { return new(proto.SetMultipleAvatarPathsScRsp) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(GetAuthkeyCsReq, func() any { return new(proto.GetAuthkeyCsReq) })
	c.regMsg(QueryProductInfoCsReq, func() any { return new(proto.QueryProductInfoCsReq) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(UpdatePlayerSettingCsReq, func() any { return new(proto.UpdatePlayerSettingCsReq) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(RaidCollectionEnterNextRaidCsReq, func() any { return new(proto.RaidCollectionEnterNextRaidCsReq) })
	c.regMsg(RaidCollectionEnterNextRaidScRsp, func() any { return new(proto.RaidCollectionEnterNextRaidScRsp) })
	c.regMsg(RaidCollectionDataScNotify, func() any { return new(proto.RaidCollectionDataScNotify) })
	c.regMsg(RaidCollectionDataScRsp, func() any { return new(proto.RaidCollectionDataScRsp) })
	c.regMsg(RaidCollectionDataCsReq, func() any { return new(proto.RaidCollectionDataCsReq) })
	c.regMsg(GetChallengeRecommendLineupListScRsp, func() any { return new(proto.GetChallengeRecommendLineupListScRsp) })
	c.regMsg(GetChallengeRecommendLineupListCsReq, func() any { return new(proto.GetChallengeRecommendLineupListCsReq) })
	c.regMsg(RelicAvatarRecommendScRsp, func() any { return new(proto.RelicAvatarRecommendScRsp) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(RelicAvatarRecommendCsReq, func() any { return new(proto.RelicAvatarRecommendCsReq) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(GetAllRedDotDataCsReq, func() any { return new(proto.GetAllRedDotDataCsReq) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(SyncRogueAeonScNotify, func() any { return new(proto.SyncRogueAeonScNotify) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(SyncRoguePickAvatarInfoScNotify, func() any { return new(proto.SyncRoguePickAvatarInfoScNotify) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(SyncRogueSeasonFinishScNotify, func() any { return new(proto.SyncRogueSeasonFinishScNotify) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(RogueArcadeLeaveCsReq, func() any { return new(proto.RogueArcadeLeaveCsReq) })
	c.regMsg(RogueArcadeStartCsReq, func() any { return new(proto.RogueArcadeStartCsReq) })
	c.regMsg(RogueArcadeRestartCsReq, func() any { return new(proto.RogueArcadeRestartCsReq) })
	c.regMsg(RogueArcadeGetInfoCsReq, func() any { return new(proto.RogueArcadeGetInfoCsReq) })
	c.regMsg(RogueArcadeGetInfoScRsp, func() any { return new(proto.RogueArcadeGetInfoScRsp) })
	c.regMsg(RogueArcadeLeaveScRsp, func() any { return new(proto.RogueArcadeLeaveScRsp) })
	c.regMsg(RogueArcadeStartScRsp, func() any { return new(proto.RogueArcadeStartScRsp) })
	c.regMsg(RogueArcadeRestartScRsp, func() any { return new(proto.RogueArcadeRestartScRsp) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(CommonRogueUpdateScNotify, func() any { return new(proto.CommonRogueUpdateScNotify) })
	c.regMsg(RogueWorkbenchSelectFuncCsReq, func() any { return new(proto.RogueWorkbenchSelectFuncCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(RogueWorkbenchSelectFuncScRsp, func() any { return new(proto.RogueWorkbenchSelectFuncScRsp) })
	c.regMsg(RogueWorkbenchHandleFuncCsReq, func() any { return new(proto.RogueWorkbenchHandleFuncCsReq) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(GetRogueCollectionCsReq, func() any { return new(proto.GetRogueCollectionCsReq) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(FinishRogueCommonDialogueScRsp, func() any { return new(proto.FinishRogueCommonDialogueScRsp) })
	c.regMsg(RogueWorkbenchGetInfoCsReq, func() any { return new(proto.RogueWorkbenchGetInfoCsReq) })
	c.regMsg(TakeRogueEventHandbookRewardScRsp, func() any { return new(proto.TakeRogueEventHandbookRewardScRsp) })
	c.regMsg(RogueWorkbenchHandleFuncScRsp, func() any { return new(proto.RogueWorkbenchHandleFuncScRsp) })
	c.regMsg(RogueGetGambleInfoCsReq, func() any { return new(proto.RogueGetGambleInfoCsReq) })
	c.regMsg(SetRogueCollectionCsReq, func() any { return new(proto.SetRogueCollectionCsReq) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(SyncRogueCommonDialogueDataScNotify, func() any { return new(proto.SyncRogueCommonDialogueDataScNotify) })
	c.regMsg(SetRogueExhibitionCsReq, func() any { return new(proto.SetRogueExhibitionCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(GetRogueShopMiracleInfoCsReq, func() any { return new(proto.GetRogueShopMiracleInfoCsReq) })
	c.regMsg(BuyRogueShopMiracleCsReq, func() any { return new(proto.BuyRogueShopMiracleCsReq) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(CommonRogueQueryCsReq, func() any { return new(proto.CommonRogueQueryCsReq) })
	c.regMsg(RogueWorkbenchGetInfoScRsp, func() any { return new(proto.RogueWorkbenchGetInfoScRsp) })
	c.regMsg(RogueGetGambleInfoScRsp, func() any { return new(proto.RogueGetGambleInfoScRsp) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(GetRogueCommonDialogueDataCsReq, func() any { return new(proto.GetRogueCommonDialogueDataCsReq) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(CommonRogueComponentUpdateScNotify, func() any { return new(proto.CommonRogueComponentUpdateScNotify) })
	c.regMsg(GetRogueExhibitionScRsp, func() any { return new(proto.GetRogueExhibitionScRsp) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(GetRogueCommonDialogueDataScRsp, func() any { return new(proto.GetRogueCommonDialogueDataScRsp) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(SetRogueExhibitionScRsp, func() any { return new(proto.SetRogueExhibitionScRsp) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(GetRogueExhibitionCsReq, func() any { return new(proto.GetRogueExhibitionCsReq) })
	c.regMsg(GetRogueCollectionScRsp, func() any { return new(proto.GetRogueCollectionScRsp) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoCsReq, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoCsReq) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(SelectRogueCommonDialogueOptionCsReq, func() any { return new(proto.SelectRogueCommonDialogueOptionCsReq) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(SyncRogueCommonDialogueOptionFinishScNotify, func() any { return new(proto.SyncRogueCommonDialogueOptionFinishScNotify) })
	c.regMsg(SelectRogueCommonDialogueOptionScRsp, func() any { return new(proto.SelectRogueCommonDialogueOptionScRsp) })
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(CommonRogueQueryScRsp, func() any { return new(proto.CommonRogueQueryScRsp) })
	c.regMsg(FinishRogueCommonDialogueCsReq, func() any { return new(proto.FinishRogueCommonDialogueCsReq) })
	c.regMsg(GetRogueShopBuffInfoCsReq, func() any { return new(proto.GetRogueShopBuffInfoCsReq) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(SetRogueCollectionScRsp, func() any { return new(proto.SetRogueCollectionScRsp) })
	c.regMsg(RogueDoGambleCsReq, func() any { return new(proto.RogueDoGambleCsReq) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
	c.regMsg(RogueDoGambleScRsp, func() any { return new(proto.RogueDoGambleScRsp) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(RogueModifierSelectCellScRsp, func() any { return new(proto.RogueModifierSelectCellScRsp) })
	c.regMsg(RogueModifierStageStartNotify, func() any { return new(proto.RogueModifierStageStartNotify) })
	c.regMsg(RogueModifierUpdateNotify, func() any { return new(proto.RogueModifierUpdateNotify) })
	c.regMsg(RogueModifierDelNotify, func() any { return new(proto.RogueModifierDelNotify) })
	c.regMsg(RogueTournSettleScRsp, func() any { return new(proto.RogueTournSettleScRsp) })
	c.regMsg(RogueTournLeaveScRsp, func() any { return new(proto.RogueTournLeaveScRsp) })
	c.regMsg(RogueTournEnterLayerScRsp, func() any { return new(proto.RogueTournEnterLayerScRsp) })
	c.regMsg(RogueTournStartCsReq, func() any { return new(proto.RogueTournStartCsReq) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneCsReq) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneScRsp) })
	c.regMsg(RogueTournDeleteArchiveCsReq, func() any { return new(proto.RogueTournDeleteArchiveCsReq) })
	c.regMsg(RogueTournLevelInfoUpdateScNotify, func() any { return new(proto.RogueTournLevelInfoUpdateScNotify) })
	c.regMsg(RogueTournExpNotify, func() any { return new(proto.RogueTournExpNotify) })
	c.regMsg(RogueTournGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueTournGetMiscRealTimeDataScRsp) })
	c.regMsg(RogueTournReviveAvatarScRsp, func() any { return new(proto.RogueTournReviveAvatarScRsp) })
	c.regMsg(RogueTournReviveAvatarCsReq, func() any { return new(proto.RogueTournReviveAvatarCsReq) })
	c.regMsg(RogueTournGetAllArchiveCsReq, func() any { return new(proto.RogueTournGetAllArchiveCsReq) })
	c.regMsg(RogueTournQueryScRsp, func() any { return new(proto.RogueTournQueryScRsp) })
	c.regMsg(RogueTournDeleteArchiveScRsp, func() any { return new(proto.RogueTournDeleteArchiveScRsp) })
	c.regMsg(RogueTournReEnterRogueCocoonStageCsReq, func() any { return new(proto.RogueTournReEnterRogueCocoonStageCsReq) })
	c.regMsg(RogueTournTakeExpRewardCsReq, func() any { return new(proto.RogueTournTakeExpRewardCsReq) })
	c.regMsg(RogueTournEnterRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournEnterRogueCocoonSceneCsReq) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoScRsp, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryScRsp, func() any { return new(proto.RogueTournGetArchiveRepositoryScRsp) })
	c.regMsg(RogueTournRenameArchiveScRsp, func() any { return new(proto.RogueTournRenameArchiveScRsp) })
	c.regMsg(RogueTournResetPermanentTalentScRsp, func() any { return new(proto.RogueTournResetPermanentTalentScRsp) })
	c.regMsg(RogueTournQueryCsReq, func() any { return new(proto.RogueTournQueryCsReq) })
	c.regMsg(RogueTournWeekChallengeUpdateScNotify, func() any { return new(proto.RogueTournWeekChallengeUpdateScNotify) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoCsReq, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoCsReq) })
	c.regMsg(RogueTournRenameArchiveCsReq, func() any { return new(proto.RogueTournRenameArchiveCsReq) })
	c.regMsg(RogueTournTakeExpRewardScRsp, func() any { return new(proto.RogueTournTakeExpRewardScRsp) })
	c.regMsg(RogueTournEnterCsReq, func() any { return new(proto.RogueTournEnterCsReq) })
	c.regMsg(RogueTournConfirmSettleScRsp, func() any { return new(proto.RogueTournConfirmSettleScRsp) })
	c.regMsg(RogueTournGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueTournGetMiscRealTimeDataCsReq) })
	c.regMsg(RogueTournEnterLayerCsReq, func() any { return new(proto.RogueTournEnterLayerCsReq) })
	c.regMsg(RogueTournGetAllArchiveScRsp, func() any { return new(proto.RogueTournGetAllArchiveScRsp) })
	c.regMsg(RogueTournReviveCostUpdateScNotify, func() any { return new(proto.RogueTournReviveCostUpdateScNotify) })
	c.regMsg(RogueTournDifficultyCompNotify, func() any { return new(proto.RogueTournDifficultyCompNotify) })
	c.regMsg(RogueTournGetArchiveRepositoryCsReq, func() any { return new(proto.RogueTournGetArchiveRepositoryCsReq) })
	c.regMsg(RogueTournEnablePermanentTalentScRsp, func() any { return new(proto.RogueTournEnablePermanentTalentScRsp) })
	c.regMsg(RogueTournGetPermanentTalentInfoCsReq, func() any { return new(proto.RogueTournGetPermanentTalentInfoCsReq) })
	c.regMsg(RogueTournGetPermanentTalentInfoScRsp, func() any { return new(proto.RogueTournGetPermanentTalentInfoScRsp) })
	c.regMsg(RogueTournHandBookNotify, func() any { return new(proto.RogueTournHandBookNotify) })
	c.regMsg(RogueTournAreaUpdateScNotify, func() any { return new(proto.RogueTournAreaUpdateScNotify) })
	c.regMsg(RogueTournEnablePermanentTalentCsReq, func() any { return new(proto.RogueTournEnablePermanentTalentCsReq) })
	c.regMsg(RogueTournEnterRoomScRsp, func() any { return new(proto.RogueTournEnterRoomScRsp) })
	c.regMsg(RogueTournLeaveCsReq, func() any { return new(proto.RogueTournLeaveCsReq) })
	c.regMsg(RogueTournEnterRoomCsReq, func() any { return new(proto.RogueTournEnterRoomCsReq) })
	c.regMsg(RogueTournSettleCsReq, func() any { return new(proto.RogueTournSettleCsReq) })
	c.regMsg(RogueTournConfirmSettleCsReq, func() any { return new(proto.RogueTournConfirmSettleCsReq) })
	c.regMsg(RogueTournGetSettleInfoScRsp, func() any { return new(proto.RogueTournGetSettleInfoScRsp) })
	c.regMsg(RogueTournReEnterRogueCocoonStageScRsp, func() any { return new(proto.RogueTournReEnterRogueCocoonStageScRsp) })
	c.regMsg(RogueTournEnterRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournEnterRogueCocoonSceneScRsp) })
	c.regMsg(RogueTournEnterScRsp, func() any { return new(proto.RogueTournEnterScRsp) })
	c.regMsg(RogueTournClearArchiveNameScNotify, func() any { return new(proto.RogueTournClearArchiveNameScNotify) })
	c.regMsg(RogueTournGetSettleInfoCsReq, func() any { return new(proto.RogueTournGetSettleInfoCsReq) })
	c.regMsg(RogueTournStartScRsp, func() any { return new(proto.RogueTournStartScRsp) })
	c.regMsg(RogueTournBattleFailSettleInfoScNotify, func() any { return new(proto.RogueTournBattleFailSettleInfoScNotify) })
	c.regMsg(RogueTournResetPermanentTalentCsReq, func() any { return new(proto.RogueTournResetPermanentTalentCsReq) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(TakeRollShopRewardCsReq, func() any { return new(proto.TakeRollShopRewardCsReq) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(SetSpringRecoverConfigCsReq, func() any { return new(proto.SetSpringRecoverConfigCsReq) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(SpringRecoverCsReq, func() any { return new(proto.SpringRecoverCsReq) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(StartTimedCocoonStageCsReq, func() any { return new(proto.StartTimedCocoonStageCsReq) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(StartTimedFarmElementScRsp, func() any { return new(proto.StartTimedFarmElementScRsp) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(SceneReviveAfterRebattleCsReq, func() any { return new(proto.SceneReviveAfterRebattleCsReq) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(HealPoolInfoNotify, func() any { return new(proto.HealPoolInfoNotify) })
	c.regMsg(SpringRecoverSingleAvatarScRsp, func() any { return new(proto.SpringRecoverSingleAvatarScRsp) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(SpringRecoverScRsp, func() any { return new(proto.SpringRecoverScRsp) })
	c.regMsg(SetSpringRecoverConfigScRsp, func() any { return new(proto.SetSpringRecoverConfigScRsp) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(GetSpringRecoverDataScRsp, func() any { return new(proto.GetSpringRecoverDataScRsp) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(StartTimedFarmElementCsReq, func() any { return new(proto.StartTimedFarmElementCsReq) })
	c.regMsg(StartTimedCocoonStageScRsp, func() any { return new(proto.StartTimedCocoonStageScRsp) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(LastSpringRefreshTimeNotify, func() any { return new(proto.LastSpringRefreshTimeNotify) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(SceneReviveAfterRebattleScRsp, func() any { return new(proto.SceneReviveAfterRebattleScRsp) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(SpringRecoverSingleAvatarCsReq, func() any { return new(proto.SpringRecoverSingleAvatarCsReq) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(SpringRefreshScRsp, func() any { return new(proto.SpringRefreshScRsp) })
	c.regMsg(SpringRefreshCsReq, func() any { return new(proto.SpringRefreshCsReq) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(GetSpringRecoverDataCsReq, func() any { return new(proto.GetSpringRecoverDataCsReq) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(SpaceZooExchangeItemCsReq, func() any { return new(proto.SpaceZooExchangeItemCsReq) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(SpaceZooTakeCsReq, func() any { return new(proto.SpaceZooTakeCsReq) })
	c.regMsg(SpaceZooTakeScRsp, func() any { return new(proto.SpaceZooTakeScRsp) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(SpaceZooExchangeItemScRsp, func() any { return new(proto.SpaceZooExchangeItemScRsp) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(StartStarFightLevelScRsp, func() any { return new(proto.StartStarFightLevelScRsp) })
	c.regMsg(GetStarFightDataScRsp, func() any { return new(proto.GetStarFightDataScRsp) })
	c.regMsg(GetStarFightDataCsReq, func() any { return new(proto.GetStarFightDataCsReq) })
	c.regMsg(StartStarFightLevelCsReq, func() any { return new(proto.StartStarFightLevelCsReq) })
	c.regMsg(StarFightDataChangeNotify, func() any { return new(proto.StarFightDataChangeNotify) })
	c.regMsg(StoryLineTrialAvatarChangeScNotify, func() any { return new(proto.StoryLineTrialAvatarChangeScNotify) })
	c.regMsg(GetStoryLineInfoScRsp, func() any { return new(proto.GetStoryLineInfoScRsp) })
	c.regMsg(StoryLineInfoScNotify, func() any { return new(proto.StoryLineInfoScNotify) })
	c.regMsg(ChangeStoryLineFinishScNotify, func() any { return new(proto.ChangeStoryLineFinishScNotify) })
	c.regMsg(GetStoryLineInfoCsReq, func() any { return new(proto.GetStoryLineInfoCsReq) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(StrongChallengeActivityBattleEndScNotify, func() any { return new(proto.StrongChallengeActivityBattleEndScNotify) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(GetSummonActivityDataScRsp, func() any { return new(proto.GetSummonActivityDataScRsp) })
	c.regMsg(GetSummonActivityDataCsReq, func() any { return new(proto.GetSummonActivityDataCsReq) })
	c.regMsg(EnterSummonActivityStageCsReq, func() any { return new(proto.EnterSummonActivityStageCsReq) })
	c.regMsg(SummonActivityBattleEndScNotify, func() any { return new(proto.SummonActivityBattleEndScNotify) })
	c.regMsg(EnterSummonActivityStageScRsp, func() any { return new(proto.EnterSummonActivityStageScRsp) })
	c.regMsg(SwordTrainingSetSkillTraceScRsp, func() any { return new(proto.SwordTrainingSetSkillTraceScRsp) })
	c.regMsg(SwordTrainingGameSyncChangeScNotify, func() any { return new(proto.SwordTrainingGameSyncChangeScNotify) })
	c.regMsg(SwordTrainingMarkEndingViewedScRsp, func() any { return new(proto.SwordTrainingMarkEndingViewedScRsp) })
	c.regMsg(SwordTrainingExamResultConfirmScRsp, func() any { return new(proto.SwordTrainingExamResultConfirmScRsp) })
	c.regMsg(SwordTrainingUnlockSyncScNotify, func() any { return new(proto.SwordTrainingUnlockSyncScNotify) })
	c.regMsg(SwordTrainingStoryConfirmScRsp, func() any { return new(proto.SwordTrainingStoryConfirmScRsp) })
	c.regMsg(SwordTrainingTurnActionCsReq, func() any { return new(proto.SwordTrainingTurnActionCsReq) })
	c.regMsg(EnterSwordTrainingExamCsReq, func() any { return new(proto.EnterSwordTrainingExamCsReq) })
	c.regMsg(SwordTrainingTurnActionScRsp, func() any { return new(proto.SwordTrainingTurnActionScRsp) })
	c.regMsg(SwordTrainingRestoreGameScRsp, func() any { return new(proto.SwordTrainingRestoreGameScRsp) })
	c.regMsg(SwordTrainingLearnSkillCsReq, func() any { return new(proto.SwordTrainingLearnSkillCsReq) })
	c.regMsg(EnterSwordTrainingExamScRsp, func() any { return new(proto.EnterSwordTrainingExamScRsp) })
	c.regMsg(SwordTrainingExamResultConfirmCsReq, func() any { return new(proto.SwordTrainingExamResultConfirmCsReq) })
	c.regMsg(SwordTrainingRestoreGameCsReq, func() any { return new(proto.SwordTrainingRestoreGameCsReq) })
	c.regMsg(GetSwordTrainingDataCsReq, func() any { return new(proto.GetSwordTrainingDataCsReq) })
	c.regMsg(SwordTrainingDialogueSelectOptionScRsp, func() any { return new(proto.SwordTrainingDialogueSelectOptionScRsp) })
	c.regMsg(SwordTrainingGiveUpGameScRsp, func() any { return new(proto.SwordTrainingGiveUpGameScRsp) })
	c.regMsg(SwordTrainingResumeGameScRsp, func() any { return new(proto.SwordTrainingResumeGameScRsp) })
	c.regMsg(SwordTrainingGameSettleScNotify, func() any { return new(proto.SwordTrainingGameSettleScNotify) })
	c.regMsg(SwordTrainingMarkEndingViewedCsReq, func() any { return new(proto.SwordTrainingMarkEndingViewedCsReq) })
	c.regMsg(SwordTrainingDialogueSelectOptionCsReq, func() any { return new(proto.SwordTrainingDialogueSelectOptionCsReq) })
	c.regMsg(SwordTrainingStartGameScRsp, func() any { return new(proto.SwordTrainingStartGameScRsp) })
	c.regMsg(SwordTrainingStoryBattleScRsp, func() any { return new(proto.SwordTrainingStoryBattleScRsp) })
	c.regMsg(SwordTrainingSetSkillTraceCsReq, func() any { return new(proto.SwordTrainingSetSkillTraceCsReq) })
	c.regMsg(SwordTrainingLearnSkillScRsp, func() any { return new(proto.SwordTrainingLearnSkillScRsp) })
	c.regMsg(SwordTrainingActionTurnSettleScNotify, func() any { return new(proto.SwordTrainingActionTurnSettleScNotify) })
	c.regMsg(SwordTrainingDailyPhaseConfirmCsReq, func() any { return new(proto.SwordTrainingDailyPhaseConfirmCsReq) })
	c.regMsg(SwordTrainingSelectEndingScRsp, func() any { return new(proto.SwordTrainingSelectEndingScRsp) })
	c.regMsg(SwordTrainingDailyPhaseConfirmScRsp, func() any { return new(proto.SwordTrainingDailyPhaseConfirmScRsp) })
	c.regMsg(SwordTrainingStoryBattleCsReq, func() any { return new(proto.SwordTrainingStoryBattleCsReq) })
	c.regMsg(SwordTrainingGiveUpGameCsReq, func() any { return new(proto.SwordTrainingGiveUpGameCsReq) })
	c.regMsg(SwordTrainingStartGameCsReq, func() any { return new(proto.SwordTrainingStartGameCsReq) })
	c.regMsg(SwordTrainingSelectEndingCsReq, func() any { return new(proto.SwordTrainingSelectEndingCsReq) })
	c.regMsg(GetSwordTrainingDataScRsp, func() any { return new(proto.GetSwordTrainingDataScRsp) })
	c.regMsg(SwordTrainingStoryConfirmCsReq, func() any { return new(proto.SwordTrainingStoryConfirmCsReq) })
	c.regMsg(SwordTrainingResumeGameCsReq, func() any { return new(proto.SwordTrainingResumeGameCsReq) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	c.regMsg(StartTrackPhotoStageScRsp, func() any { return new(proto.StartTrackPhotoStageScRsp) })
	c.regMsg(SettleTrackPhotoStageCsReq, func() any { return new(proto.SettleTrackPhotoStageCsReq) })
	c.regMsg(QuitTrackPhotoStageCsReq, func() any { return new(proto.QuitTrackPhotoStageCsReq) })
	c.regMsg(GetTrackPhotoActivityDataCsReq, func() any { return new(proto.GetTrackPhotoActivityDataCsReq) })
	c.regMsg(SettleTrackPhotoStageScRsp, func() any { return new(proto.SettleTrackPhotoStageScRsp) })
	c.regMsg(StartTrackPhotoStageCsReq, func() any { return new(proto.StartTrackPhotoStageCsReq) })
	c.regMsg(QuitTrackPhotoStageScRsp, func() any { return new(proto.QuitTrackPhotoStageScRsp) })
	c.regMsg(GetTrackPhotoActivityDataScRsp, func() any { return new(proto.GetTrackPhotoActivityDataScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(GetTrainVisitorBehaviorCsReq, func() any { return new(proto.GetTrainVisitorBehaviorCsReq) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(TrainVisitorBehaviorFinishCsReq, func() any { return new(proto.TrainVisitorBehaviorFinishCsReq) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(WolfBroGameExplodeMonsterCsReq, func() any { return new(proto.WolfBroGameExplodeMonsterCsReq) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(WolfBroGameActivateBulletScRsp, func() any { return new(proto.WolfBroGameActivateBulletScRsp) })
	c.regMsg(WolfBroGameExplodeMonsterScRsp, func() any { return new(proto.WolfBroGameExplodeMonsterScRsp) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
}
