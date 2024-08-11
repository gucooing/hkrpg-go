package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	TakeTrialActivityRewardScRsp                       = 2633
	TakeTrialActivityRewardCsReq                       = 2668
	TakeMonsterResearchActivityRewardCsReq             = 2673
	LeaveTrialActivityScRsp                            = 2632
	GetLoginActivityCsReq                              = 2618
	SubmitMonsterResearchActivityMaterialCsReq         = 2640
	StartTrialActivityScRsp                            = 2667
	EnterTrialActivityStageScRsp                       = 2620
	TakeLoginActivityRewardCsReq                       = 2629
	TakeMonsterResearchActivityRewardScRsp             = 2608
	StartTrialActivityCsReq                            = 2690
	GetActivityScheduleConfigScRsp                     = 2649
	TrialActivityDataChangeScNotify                    = 2700
	GetMonsterResearchActivityDataScRsp                = 2662
	TakeLoginActivityRewardScRsp                       = 2637
	GetMonsterResearchActivityDataCsReq                = 2691
	SubmitMonsterResearchActivityMaterialScRsp         = 2648
	CurTrialActivityScNotify                           = 2694
	LeaveTrialActivityCsReq                            = 2603
	GetTrialActivityDataScRsp                          = 2688
	EnterTrialActivityStageCsReq                       = 2610
	GetActivityScheduleConfigCsReq                     = 2658
	GetLoginActivityScRsp                              = 2674
	GetTrialActivityDataCsReq                          = 2636
	GetFarmStageGachaInfoCsReq                         = 1329
	EnterAdventureScRsp                                = 1374
	GetFarmStageGachaInfoScRsp                         = 1337
	EnterAdventureCsReq                                = 1318
	StartAetherDivideStageBattleCsReq                  = 4873
	StartAetherDivideStageBattleScRsp                  = 4808
	ClearAetherDividePassiveSkillScRsp                 = 4862
	AetherDivideLineupScNotify                         = 4859
	AetherDivideFinishChallengeScNotify                = 4847
	AetherDivideSpiritInfoScNotify                     = 4877
	GetAetherDivideInfoScRsp                           = 4857
	SwitchAetherDivideLineUpSlotCsReq                  = 4840
	SetAetherDivideLineUpCsReq                         = 4846
	AetherDivideTakeChallengeRewardScRsp               = 4806
	AetherDivideRefreshEndlessCsReq                    = 4802
	AetherDivideRefreshEndlessScNotify                 = 4809
	LeaveAetherDivideSceneScRsp                        = 4837
	ClearAetherDividePassiveSkillCsReq                 = 4891
	EnterAetherDivideSceneScRsp                        = 4874
	GetAetherDivideChallengeInfoScRsp                  = 4892
	AetherDivideTakeChallengeRewardCsReq               = 4856
	EnterAetherDivideSceneCsReq                        = 4818
	SetAetherDivideLineUpScRsp                         = 4898
	LeaveAetherDivideSceneCsReq                        = 4829
	AetherDivideSkillItemScNotify                      = 4838
	AetherDivideTainerInfoScNotify                     = 4861
	StartAetherDivideSceneBattleScRsp                  = 4849
	StartAetherDivideSceneBattleCsReq                  = 4858
	EquipAetherDividePassiveSkillCsReq                 = 4883
	SwitchAetherDivideLineUpSlotScRsp                  = 4848
	GetAetherDivideChallengeInfoCsReq                  = 4824
	AetherDivideRefreshEndlessScRsp                    = 4813
	StartAetherDivideChallengeBattleScRsp              = 4841
	AetherDivideSpiritExpUpCsReq                       = 4811
	AetherDivideSpiritExpUpScRsp                       = 4821
	StartAetherDivideChallengeBattleCsReq              = 4828
	EquipAetherDividePassiveSkillScRsp                 = 4822
	GetAetherDivideInfoCsReq                           = 4889
	SaveLogisticsScRsp                                 = 4792
	TakePrestigeRewardCsReq                            = 4789
	AlleyFundsScNotify                                 = 4711
	AlleyShopLevelScNotify                             = 4721
	GetSaveLogisticsMapCsReq                           = 4738
	AlleyShipmentEventEffectsScNotify                  = 4761
	AlleyEventEffectNotify                             = 4764
	StartAlleyEventScRsp                               = 4741
	AlleyTakeEventRewardCsReq                          = 4709
	AlleyGuaranteedFundsScRsp                          = 4713
	TakePrestigeRewardScRsp                            = 4757
	AlleyPlacingGameCsReq                              = 4746
	LogisticsScoreRewardSyncInfoScNotify               = 4780
	AlleyShipUsedCountScNotify                         = 4759
	GetSaveLogisticsMapScRsp                           = 4782
	AlleyPlacingGameScRsp                              = 4798
	RefreshAlleyOrderScRsp                             = 4762
	RefreshAlleyOrderCsReq                             = 4791
	SaveLogisticsCsReq                                 = 4724
	AlleyGuaranteedFundsCsReq                          = 4702
	LogisticsGameCsReq                                 = 4729
	LogisticsGameScRsp                                 = 4737
	LogisticsDetonateStarSkiffScRsp                    = 4770
	GetAlleyInfoCsReq                                  = 4718
	PrestigeLevelUpScRsp                               = 4708
	AlleyEventChangeNotify                             = 4719
	AlleyTakeEventRewardScRsp                          = 4756
	PrestigeLevelUpCsReq                               = 4773
	GetAlleyInfoScRsp                                  = 4774
	AlleyShipUnlockScNotify                            = 4777
	AlleyOrderChangedScNotify                          = 4740
	LogisticsDetonateStarSkiffCsReq                    = 4706
	LogisticsInfoScNotify                              = 4747
	StartAlleyEventCsReq                               = 4728
	GetUpdatedArchiveDataCsReq                         = 2329
	GetArchiveDataScRsp                                = 2374
	GetArchiveDataCsReq                                = 2318
	GetUpdatedArchiveDataScRsp                         = 2337
	DressAvatarSkinScRsp                               = 311
	RankUpAvatarCsReq                                  = 398
	RankUpAvatarScRsp                                  = 383
	MarkAvatarCsReq                                    = 392
	TakeOffRelicCsReq                                  = 362
	TakePromotionRewardCsReq                           = 348
	LevelUpSpecialSkillTreeCsReq                       = 361
	UnlockSpecialSkillTreeScNotify                     = 382
	LevelUpSpecialSkillTreeScRsp                       = 338
	UnlockSkilltreeCsReq                               = 358
	DressAvatarSkinCsReq                               = 308
	TakeOffRelicScRsp                                  = 340
	DressRelicAvatarCsReq                              = 322
	DressAvatarScRsp                                   = 364
	UnlockSkilltreeScRsp                               = 349
	PromoteAvatarScRsp                                 = 341
	AvatarExpUpCsReq                                   = 329
	AddAvatarScNotify                                  = 346
	TakeOffAvatarSkinScRsp                             = 377
	TakePromotionRewardScRsp                           = 373
	GetAvatarDataScRsp                                 = 374
	PromoteAvatarCsReq                                 = 328
	DressAvatarCsReq                                   = 319
	DressRelicAvatarScRsp                              = 391
	MarkAvatarScRsp                                    = 347
	TakeOffAvatarSkinCsReq                             = 321
	GetAvatarDataCsReq                                 = 318
	UnlockAvatarSkinScNotify                           = 324
	TakeOffEquipmentScRsp                              = 357
	AvatarExpUpScRsp                                   = 337
	TakeOffEquipmentCsReq                              = 389
	SyncClientResVersionScRsp                          = 141
	QuitBattleScRsp                                    = 137
	ServerSimulateBattleFinishScNotify                 = 157
	BattleLogReportCsReq                               = 164
	PVEBattleResultScRsp                               = 174
	ReBattleAfterBattleLoseCsNotify                    = 146
	SyncClientResVersionCsReq                          = 128
	QuitBattleScNotify                                 = 119
	BattleLogReportScRsp                               = 189
	GetCurBattleInfoScRsp                              = 149
	PVEBattleResultCsReq                               = 118
	QuitBattleCsReq                                    = 129
	RebattleByClientCsNotify                           = 198
	GetCurBattleInfoCsReq                              = 158
	GetBattleCollegeDataCsReq                          = 5718
	StartBattleCollegeCsReq                            = 5737
	BattleCollegeDataChangeScNotify                    = 5729
	GetBattleCollegeDataScRsp                          = 5774
	StartBattleCollegeScRsp                            = 5758
	TakeBpRewardCsReq                                  = 3037
	TakeBpRewardScRsp                                  = 3058
	BattlePassInfoNotify                               = 3018
	BuyBpLevelScRsp                                    = 3028
	TakeAllRewardScRsp                                 = 3019
	TakeAllRewardCsReq                                 = 3041
	BuyBpLevelCsReq                                    = 3049
	MatchBoxingClubOpponentCsReq                       = 4229
	GetBoxingClubInfoCsReq                             = 4218
	GiveUpBoxingClubChallengeScRsp                     = 4241
	SetBoxingClubResonanceLineupCsReq                  = 4246
	ChooseBoxingClubResonanceScRsp                     = 4257
	BoxingClubRewardScNotify                           = 4219
	GetBoxingClubInfoScRsp                             = 4274
	StartBoxingClubBattleScRsp                         = 4249
	ChooseBoxingClubResonanceCsReq                     = 4289
	SetBoxingClubResonanceLineupScRsp                  = 4298
	ChooseBoxingClubStageOptionalBuffScRsp             = 4222
	BoxingClubChallengeUpdateScNotify                  = 4264
	StartBoxingClubBattleCsReq                         = 4258
	MatchBoxingClubOpponentScRsp                       = 4237
	GiveUpBoxingClubChallengeCsReq                     = 4228
	ChooseBoxingClubStageOptionalBuffCsReq             = 4283
	GetChallengeGroupStatisticsScRsp                   = 1762
	GetCurChallengeScRsp                               = 1789
	LeaveChallengeScRsp                                = 1749
	TakeChallengeRewardCsReq                           = 1783
	EnterChallengeNextPhaseCsReq                       = 1721
	ChallengeSettleNotify                              = 1728
	LeaveChallengeCsReq                                = 1758
	StartChallengeCsReq                                = 1729
	GetCurChallengeCsReq                               = 1764
	RestartChallengePhaseCsReq                         = 1708
	EnterChallengeNextPhaseScRsp                       = 1777
	GetChallengeGroupStatisticsCsReq                   = 1791
	StartPartialChallengeCsReq                         = 1740
	StartChallengeScRsp                                = 1737
	GetChallengeScRsp                                  = 1774
	GetChallengeCsReq                                  = 1718
	StartPartialChallengeScRsp                         = 1748
	TakeChallengeRewardScRsp                           = 1722
	ChallengeLineupNotify                              = 1757
	RestartChallengePhaseScRsp                         = 1711
	ChallengeBossPhaseSettleNotify                     = 1724
	GetLoginChatInfoCsReq                              = 3983
	SendMsgScRsp                                       = 3974
	RevcMsgScNotify                                    = 3929
	PrivateMsgOfflineUsersScNotify                     = 3937
	MarkChatEmojiScRsp                                 = 3957
	GetChatFriendHistoryCsReq                          = 3928
	GetLoginChatInfoScRsp                              = 3922
	GetChatFriendHistoryScRsp                          = 3941
	SendMsgCsReq                                       = 3918
	GetChatEmojiListCsReq                              = 3919
	MarkChatEmojiCsReq                                 = 3989
	GetChatEmojiListScRsp                              = 3964
	GetPrivateChatHistoryScRsp                         = 3949
	GetPrivateChatHistoryCsReq                         = 3958
	BatchMarkChatEmojiScRsp                            = 3998
	BatchMarkChatEmojiCsReq                            = 3946
	ChessRogueRollDiceScRsp                            = 5471
	FinishChessRogueSubStoryScRsp                      = 5407
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5548
	ChessRogueUpdateAeonModifierValueScNotify          = 5475
	GetChessRogueNousStoryInfoScRsp                    = 5466
	ChessRogueQueryBpCsReq                             = 5564
	ChessRogueGiveUpRollScRsp                          = 5455
	ChessRogueLeaveScRsp                               = 5587
	ChessRogueSkipTeachingLevelCsReq                   = 5519
	ChessRogueUpdateDiceInfoScNotify                   = 5447
	FinishChessRogueSubStoryCsReq                      = 5411
	ChessRogueNousGetRogueTalentInfoCsReq              = 5539
	ChessRogueEnterCellCsReq                           = 5493
	SyncChessRogueMainStoryFinishScNotify              = 5438
	ChessRogueRollDiceCsReq                            = 5450
	ChessRogueCheatRollScRsp                           = 5567
	ChessRogueSelectBpCsReq                            = 5406
	ChessRogueEnterNextLayerCsReq                      = 5473
	ChessRogueCellUpdateNotify                         = 5412
	ChessRogueMoveCellNotify                           = 5536
	ChessRogueQueryAeonDimensionsScRsp                 = 5576
	ChessRogueUpdateReviveInfoScNotify                 = 5420
	EnterChessRogueAeonRoomCsReq                       = 5452
	ChessRogueNousEditDiceCsReq                        = 5445
	EnhanceChessRogueBuffScRsp                         = 5505
	GetChessRogueStoryAeonTalkInfoScRsp                = 5500
	ChessRogueSelectCellScRsp                          = 5442
	EnhanceChessRogueBuffCsReq                         = 5404
	ChessRogueChangeyAeonDimensionNotify               = 5512
	ChessRogueSelectBpScRsp                            = 5526
	ChessRogueLeaveCsReq                               = 5507
	ChessRogueReviveAvatarScRsp                        = 5470
	ChessRogueConfirmRollCsReq                         = 5415
	ChessRogueEnterScRsp                               = 5590
	ChessRogueLayerAccountInfoNotify                   = 5565
	ChessRogueNousEnableRogueTalentCsReq               = 5531
	ChessRogueEnterCsReq                               = 5521
	GetChessRogueBuffEnhanceInfoScRsp                  = 5421
	ChessRogueGoAheadCsReq                             = 5504
	ChessRogueQuitCsReq                                = 5549
	ChessRogueSkipTeachingLevelScRsp                   = 5489
	GetChessRogueStoryAeonTalkInfoCsReq                = 5456
	SelectChessRogueSubStoryScRsp                      = 5517
	ChessRogueNousGetRogueTalentInfoScRsp              = 5571
	ChessRogueStartScRsp                               = 5570
	ChessRogueGiveUpCsReq                              = 5524
	GetChessRogueBuffEnhanceInfoCsReq                  = 5534
	EnterChessRogueAeonRoomScRsp                       = 5435
	SelectChessRogueNousSubStoryScRsp                  = 5535
	ChessRogueNousEditDiceScRsp                        = 5463
	ChessRogueStartCsReq                               = 5511
	ChessRogueGiveUpScRsp                              = 5481
	SelectChessRogueSubStoryCsReq                      = 5427
	ChessRoguePickAvatarScRsp                          = 5441
	ChessRogueSelectCellCsReq                          = 5513
	SyncChessRogueNousSubStoryScNotify                 = 5402
	ChessRogueQueryBpScRsp                             = 5479
	ChessRogueEnterCellScRsp                           = 5516
	GetChessRogueStoryInfoScRsp                        = 5575
	ChessRogueUpdateActionPointScNotify                = 5425
	ChessRogueNousDiceUpdateNotify                     = 5542
	GetChessRogueNousStoryInfoCsReq                    = 5480
	ChessRogueGoAheadScRsp                             = 5529
	ChessRogueQuitScRsp                                = 5588
	ChessRogueQueryCsReq                               = 5458
	FinishChessRogueNousSubStoryCsReq                  = 5579
	ChessRogueFinishCurRoomNotify                      = 5439
	ChessRogueUpdateBoardScNotify                      = 5515
	SyncChessRogueNousMainStoryScNotify                = 5419
	ChessRogueConfirmRollScRsp                         = 5428
	ChessRogueQueryScRsp                               = 5484
	ChessRogueGiveUpRollCsReq                          = 5448
	ChessRogueEnterNextLayerScRsp                      = 5582
	FinishChessRogueNousSubStoryScRsp                  = 5560
	ChessRogueUpdateUnlockLevelScNotify                = 5472
	ChessRogueUpdateMoneyInfoScNotify                  = 5594
	SelectChessRogueNousSubStoryCsReq                  = 5566
	ChessRogueReviveAvatarCsReq                        = 5518
	GetChessRogueStoryInfoCsReq                        = 5482
	ChessRogueUpdateAllowedSelectCellScNotify          = 5464
	ChessRoguePickAvatarCsReq                          = 5495
	ChessRogueNousDiceSurfaceUnlockNotify              = 5426
	ChessRogueCheatRollCsReq                           = 5444
	ChessRogueReRollDiceScRsp                          = 5555
	ChessRogueReRollDiceCsReq                          = 5424
	SyncChessRogueNousValueScNotify                    = 5437
	ChessRogueNousEnableRogueTalentScRsp               = 5449
	ChessRogueUpdateLevelBaseInfoScNotify              = 5546
	ChessRogueQueryAeonDimensionsCsReq                 = 5510
	ChessRogueQuestFinishNotify                        = 5578
	ClockParkBattleEndScNotify                         = 7291
	ClockParkStartScriptScRsp                          = 7241
	ClockParkHandleWaitOperationCsReq                  = 7289
	ClockParkGetInfoScRsp                              = 7274
	ClockParkGetInfoCsReq                              = 7218
	ClockParkQuitScriptCsReq                           = 7246
	ClockParkHandleWaitOperationScRsp                  = 7257
	ClockParkUseBuffCsReq                              = 7240
	ClockParkGetOngoingScriptInfoCsReq                 = 7219
	ClockParkStartScriptCsReq                          = 7228
	ClockParkUseBuffScRsp                              = 7248
	ClockParkGetOngoingScriptInfoScRsp                 = 7264
	ClockParkUnlockTalentScRsp                         = 7249
	ClockParkQuitScriptScRsp                           = 7298
	ClockParkFinishScriptScNotify                      = 7273
	ClockParkUnlockTalentCsReq                         = 7258
	ContentPackageGetDataCsReq                         = 7519
	ContentPackageSyncDataScNotify                     = 7506
	ContentPackageUnlockScRsp                          = 7537
	ContentPackageUnlockCsReq                          = 7545
	ContentPackageTransferScNotify                     = 7520
	ContentPackageGetDataScRsp                         = 7529
	DailyActiveInfoNotify                              = 3358
	TakeApRewardScRsp                                  = 3374
	TakeAllApRewardCsReq                               = 3349
	TakeApRewardCsReq                                  = 3318
	GetDailyActiveInfoScRsp                            = 3337
	TakeAllApRewardScRsp                               = 3328
	GetDailyActiveInfoCsReq                            = 3329
	MakeMissionDrinkCsReq                              = 6989
	DrinkMakerChallengeScRsp                           = 6981
	GetDrinkMakerDataCsReq                             = 6994
	DrinkMakerChallengeCsReq                           = 6999
	MakeDrinkCsReq                                     = 6993
	EndDrinkMakerSequenceCsReq                         = 6983
	DrinkMakerUpdateTipsNotify                         = 6990
	MakeMissionDrinkScRsp                              = 6986
	GetDrinkMakerDataScRsp                             = 6997
	EndDrinkMakerSequenceScRsp                         = 6984
	DrinkMakerDayEndScNotify                           = 6992
	MakeDrinkScRsp                                     = 6985
	EvolveBuildQueryInfoScRsp                          = 7129
	EvolveBuildStartStageScRsp                         = 7120
	EvolveBuildStartStageCsReq                         = 7137
	EvolveBuildGiveupCsReq                             = 7112
	EvolveBuildLeaveScRsp                              = 7150
	EvolveBuildShopAbilityDownCsReq                    = 7124
	EvolveBuildReRandomStageCsReq                      = 7105
	EvolveBuildStartLevelCsReq                         = 7106
	EvolveBuildShopAbilityUpScRsp                      = 7131
	EvolveBuildTakeExpRewardCsReq                      = 7148
	EvolveBuildUnlockInfoNotify                        = 7123
	EvolveBuildCoinNotify                              = 7139
	EvolveBuildFinishScNotify                          = 7101
	EvolveBuildGiveupScRsp                             = 7121
	EvolveBuildShopAbilityUpCsReq                      = 7113
	EvolveBuildShopAbilityResetCsReq                   = 7117
	EvolveBuildTakeExpRewardScRsp                      = 7144
	EvolveBuildQueryInfoCsReq                          = 7119
	EvolveBuildShopAbilityDownScRsp                    = 7116
	EvolveBuildReRandomStageScRsp                      = 7103
	EvolveBuildStartLevelScRsp                         = 7145
	EvolveBuildShopAbilityResetScRsp                   = 7149
	EvolveBuildLeaveCsReq                              = 7109
	TakeMultipleExpeditionRewardCsReq                  = 2562
	CancelExpeditionCsReq                              = 2558
	AcceptMultipleExpeditionScRsp                      = 2591
	TakeMultipleExpeditionRewardScRsp                  = 2540
	AcceptActivityExpeditionCsReq                      = 2564
	GetExpeditionDataCsReq                             = 2518
	CancelActivityExpeditionCsReq                      = 2557
	AcceptExpeditionScRsp                              = 2537
	CancelExpeditionScRsp                              = 2549
	AcceptMultipleExpeditionCsReq                      = 2522
	ExpeditionDataChangeScNotify                       = 2519
	TakeActivityExpeditionRewardScRsp                  = 2583
	TakeExpeditionRewardCsReq                          = 2528
	CancelActivityExpeditionScRsp                      = 2546
	AcceptExpeditionCsReq                              = 2529
	AcceptActivityExpeditionScRsp                      = 2589
	TakeActivityExpeditionRewardCsReq                  = 2598
	TakeExpeditionRewardScRsp                          = 2541
	GetExpeditionDataScRsp                             = 2574
	EnterFantasticStoryActivityStageScRsp              = 4958
	EnterFantasticStoryActivityStageCsReq              = 4937
	FinishChapterScNotify                              = 4929
	GetFantasticStoryActivityDataCsReq                 = 4918
	GetFantasticStoryActivityDataScRsp                 = 4974
	FantasticStoryActivityBattleEndScNotify            = 4949
	FeverTimeActivityBattleEndScNotify                 = 7157
	EnterFeverTimeActivityStageCsReq                   = 7152
	GetFeverTimeActivityDataCsReq                      = 7160
	GetFeverTimeActivityDataScRsp                      = 7156
	EnterFeverTimeActivityStageScRsp                   = 7155
	FightLeaveScNotify                                 = 30029
	FightHeartBeatCsReq                                = 30058
	FightEnterScRsp                                    = 30074
	FightHeartBeatScRsp                                = 30049
	FightKickOutScNotify                               = 30037
	FightEnterCsReq                                    = 30018
	FightSessionStopScNotify                           = 30028
	GetFightActivityDataCsReq                          = 3618
	EnterFightActivityStageCsReq                       = 3637
	GetFightActivityDataScRsp                          = 3674
	EnterFightActivityStageScRsp                       = 3658
	FightActivityDataChangeScNotify                    = 3629
	TakeFightActivityRewardCsReq                       = 3649
	TakeFightActivityRewardScRsp                       = 3628
	FightMatch3DataScRsp                               = 30174
	FightMatch3StartCountDownScNotify                  = 30129
	FightMatch3ChatScNotify                            = 30189
	FightMatch3SwapCsReq                               = 30149
	FightMatch3ChatCsReq                               = 30119
	FightMatch3TurnStartScNotify                       = 30137
	FightMatch3SwapScRsp                               = 30128
	FightMatch3ChatScRsp                               = 30164
	FightMatch3ForceUpdateNotify                       = 30157
	FightMatch3TurnEndScNotify                         = 30158
	FightMatch3DataCsReq                               = 30118
	FightMatch3OpponentDataScNotify                    = 30141
	GetPlayerDetailInfoCsReq                           = 2929
	GetFriendChallengeDetailCsReq                      = 2910
	GetPlatformPlayerInfoScRsp                         = 2939
	GetAssistListCsReq                                 = 2961
	SetFriendMarkCsReq                                 = 2942
	TakeAssistRewardCsReq                              = 2970
	GetFriendBattleRecordDetailCsReq                   = 2968
	GetFriendDevelopmentInfoScRsp                      = 2967
	GetFriendLoginInfoScRsp                            = 2975
	GetFriendLoginInfoCsReq                            = 2914
	NewAssistHistoryNotify                             = 2906
	GetAssistHistoryCsReq                              = 2909
	GetFriendRecommendListInfoScRsp                    = 2948
	ApplyFriendScRsp                                   = 2941
	SetForbidOtherApplyFriendCsReq                     = 2953
	SyncHandleFriendScNotify                           = 2957
	SetFriendRemarkNameScRsp                           = 2908
	SetFriendMarkScRsp                                 = 2945
	GetCurAssistScRsp                                  = 2913
	DeleteBlacklistScRsp                               = 2924
	SearchPlayerScRsp                                  = 2947
	SetFriendRemarkNameCsReq                           = 2973
	GetFriendListInfoCsReq                             = 2918
	TakeAssistRewardScRsp                              = 2980
	SearchPlayerCsReq                                  = 2992
	GetFriendRecommendListInfoCsReq                    = 2940
	DeleteFriendCsReq                                  = 2946
	ReportPlayerScRsp                                  = 2921
	SyncAddBlacklistScNotify                           = 2962
	GetAssistHistoryScRsp                              = 2956
	ReportPlayerCsReq                                  = 2911
	CurAssistChangedNotify                             = 2923
	HandleFriendCsReq                                  = 2964
	ApplyFriendCsReq                                   = 2928
	SetForbidOtherApplyFriendScRsp                     = 2912
	GetFriendChallengeLineupScRsp                      = 3000
	GetPlayerDetailInfoScRsp                           = 2937
	SetAssistScRsp                                     = 2959
	AddBlacklistScRsp                                  = 2991
	GetPlatformPlayerInfoCsReq                         = 2987
	GetFriendChallengeDetailScRsp                      = 2920
	GetFriendApplyListInfoCsReq                        = 2958
	GetFriendAssistListCsReq                           = 2916
	GetFriendDevelopmentInfoCsReq                      = 2990
	GetFriendChallengeLineupCsReq                      = 2988
	GetFriendApplyListInfoScRsp                        = 2949
	AddBlacklistCsReq                                  = 2922
	GetFriendListInfoScRsp                             = 2974
	GetFriendBattleRecordDetailScRsp                   = 2933
	GetCurAssistCsReq                                  = 2902
	SetAssistCsReq                                     = 2982
	GetFriendAssistListScRsp                           = 2936
	HandleFriendScRsp                                  = 2989
	DeleteBlacklistCsReq                               = 2977
	DeleteFriendScRsp                                  = 2998
	SyncDeleteFriendScNotify                           = 2983
	GetAssistListScRsp                                 = 2938
	SyncApplyFriendScNotify                            = 2919
	DoGachaScRsp                                       = 1937
	ExchangeGachaCeilingScRsp                          = 1941
	GetGachaCeilingScRsp                               = 1949
	GetGachaInfoScRsp                                  = 1974
	ExchangeGachaCeilingCsReq                          = 1928
	GetGachaCeilingCsReq                               = 1958
	DoGachaCsReq                                       = 1929
	GetGachaInfoCsReq                                  = 1918
	FinishEmotionDialoguePerformanceScRsp              = 6341
	GetHeartDialInfoScRsp                              = 6374
	HeartDialScriptChangeScNotify                      = 6319
	SubmitEmotionItemScRsp                             = 6349
	SubmitEmotionItemCsReq                             = 6358
	HeartDialTraceScriptScRsp                          = 6389
	HeartDialTraceScriptCsReq                          = 6364
	GetHeartDialInfoCsReq                              = 6318
	ChangeScriptEmotionCsReq                           = 6329
	ChangeScriptEmotionScRsp                           = 6337
	FinishEmotionDialoguePerformanceCsReq              = 6328
	HeliobusEnterBattleScRsp                           = 5873
	HeliobusSnsReadCsReq                               = 5829
	HeliobusActivityDataScRsp                          = 5874
	HeliobusSnsReadScRsp                               = 5837
	HeliobusSnsCommentCsReq                            = 5819
	HeliobusActivityDataCsReq                          = 5818
	HeliobusSnsLikeCsReq                               = 5828
	HeliobusLineupUpdateScNotify                       = 5877
	HeliobusUpgradeLevelCsReq                          = 5846
	HeliobusSnsPostCsReq                               = 5858
	HeliobusChallengeUpdateScNotify                    = 5821
	HeliobusStartRaidScRsp                             = 5811
	HeliobusUnlockSkillScNotify                        = 5883
	HeliobusSnsLikeScRsp                               = 5841
	HeliobusSnsPostScRsp                               = 5849
	HeliobusInfoChangedScNotify                        = 5857
	HeliobusEnterBattleCsReq                           = 5848
	HeliobusSelectSkillScRsp                           = 5891
	HeliobusUpgradeLevelScRsp                          = 5898
	HeliobusStartRaidCsReq                             = 5808
	HeliobusSnsUpdateScNotify                          = 5889
	HeliobusSelectSkillCsReq                           = 5822
	HeliobusSnsCommentScRsp                            = 5864
	ComposeSelectedRelicCsReq                          = 521
	ExpUpRelicScRsp                                    = 522
	ComposeItemCsReq                                   = 546
	LockRelicScRsp                                     = 562
	RankUpEquipmentCsReq                               = 519
	GetMarkItemListCsReq                               = 502
	SyncTurnFoodNotify                                 = 570
	LockEquipmentCsReq                                 = 558
	ExpUpRelicCsReq                                    = 583
	RelicAvatarRecommendScRsp                          = 542
	MarkItemCsReq                                      = 509
	GetRelicFilterPlanScRsp                            = 516
	ComposeLimitNumUpdateNotify                        = 538
	GetMarkItemListScRsp                               = 513
	DeleteRelicFilterPlanCsReq                         = 520
	AddRelicFilterPlanCsReq                            = 536
	AddEquipmentScNotify                               = 524
	ComposeSelectedRelicScRsp                          = 577
	UseItemScRsp                                       = 541
	DeleteRelicFilterPlanScRsp                         = 568
	GetRecyleTimeCsReq                                 = 592
	ComposeLimitNumCompleteNotify                      = 561
	MarkRelicFilterPlanScRsp                           = 590
	DestroyItemCsReq                                   = 582
	ComposeItemScRsp                                   = 598
	DestroyItemScRsp                                   = 559
	UseItemCsReq                                       = 528
	MarkRelicFilterPlanCsReq                           = 533
	RechargeSuccNotify                                 = 573
	AddRelicFilterPlanScRsp                            = 588
	SellItemScRsp                                      = 548
	LockRelicCsReq                                     = 591
	GetBagScRsp                                        = 574
	ExpUpEquipmentScRsp                                = 557
	MarkItemScRsp                                      = 556
	RelicFilterPlanClearNameScNotify                   = 567
	SetTurnFoodSwitchCsReq                             = 580
	ModifyRelicFilterPlanScRsp                         = 510
	GetBagCsReq                                        = 518
	DiscardRelicCsReq                                  = 539
	SetTurnFoodSwitchScRsp                             = 523
	PromoteEquipmentCsReq                              = 529
	RelicRecommendCsReq                                = 575
	SellItemCsReq                                      = 540
	GeneralVirtualItemDataNotify                       = 587
	LockEquipmentScRsp                                 = 549
	ExchangeHcoinScRsp                                 = 511
	ModifyRelicFilterPlanCsReq                         = 600
	GetRelicFilterPlanCsReq                            = 545
	ExchangeHcoinCsReq                                 = 508
	ExpUpEquipmentCsReq                                = 589
	CancelMarkItemNotify                               = 506
	GetRecyleTimeScRsp                                 = 547
	PromoteEquipmentScRsp                              = 537
	DiscardRelicScRsp                                  = 514
	RelicRecommendScRsp                                = 553
	RelicAvatarRecommendCsReq                          = 512
	RankUpEquipmentScRsp                               = 564
	GetJukeboxDataScRsp                                = 3174
	TrialBackGroundMusicScRsp                          = 3141
	UnlockBackGroundMusicCsReq                         = 3158
	PlayBackGroundMusicCsReq                           = 3129
	UnlockBackGroundMusicScRsp                         = 3149
	TrialBackGroundMusicCsReq                          = 3128
	PlayBackGroundMusicScRsp                           = 3137
	GetJukeboxDataCsReq                                = 3118
	QuitLineupCsReq                                    = 728
	SwapLineupCsReq                                    = 719
	GetAllLineupDataScRsp                              = 773
	SetLineupNameCsReq                                 = 762
	ChangeLineupLeaderCsReq                            = 798
	SetLineupNameScRsp                                 = 740
	JoinLineupCsReq                                    = 758
	ChangeLineupLeaderScRsp                            = 783
	QuitLineupScRsp                                    = 741
	ReplaceLineupScRsp                                 = 721
	GetStageLineupScRsp                                = 774
	VirtualLineupDestroyNotify                         = 708
	SwapLineupScRsp                                    = 764
	SyncLineupNotify                                   = 789
	GetLineupAvatarDataCsReq                           = 757
	GetLineupAvatarDataScRsp                           = 746
	GetCurLineupDataCsReq                              = 729
	GetCurLineupDataScRsp                              = 737
	SwitchLineupIndexCsReq                             = 722
	ReplaceLineupCsReq                                 = 711
	GetStageLineupCsReq                                = 718
	GetAllLineupDataCsReq                              = 748
	ExtraLineupDestroyNotify                           = 777
	SwitchLineupIndexScRsp                             = 791
	JoinLineupScRsp                                    = 749
	LobbySyncInfoScNotify                              = 7351
	LobbyGetInfoScRsp                                  = 7374
	LobbyQuitScRsp                                     = 7360
	LobbyCreateScRsp                                   = 7379
	LobbyQuitCsReq                                     = 7391
	LobbyInviteScRsp                                   = 7353
	LobbyGetInfoCsReq                                  = 7381
	LobbyInviteCsReq                                   = 7355
	LobbyModifyPlayerInfoCsReq                         = 7362
	LobbyJoinScRsp                                     = 7395
	LobbyBeginScRsp                                    = 7370
	LobbyJoinCsReq                                     = 7356
	LobbyKickOutCsReq                                  = 7359
	LobbyBeginCsReq                                    = 7387
	LobbyKickOutScRsp                                  = 7400
	LobbyInviteScNotify                                = 7363
	LobbyCreateCsReq                                   = 7369
	LobbyModifyPlayerInfoScRsp                         = 7371
	MarkReadMailCsReq                                  = 829
	TakeMailAttachmentScRsp                            = 841
	NewMailScNotify                                    = 819
	MarkReadMailScRsp                                  = 837
	GetMailScRsp                                       = 874
	GetMailCsReq                                       = 818
	DelMailScRsp                                       = 849
	DelMailCsReq                                       = 858
	TakeMailAttachmentCsReq                            = 828
	InteractChargerScRsp                               = 6837
	DeployRotaterCsReq                                 = 6858
	EnterMapRotationRegionScRsp                        = 6874
	GetMapRotationDataScRsp                            = 6857
	EnterMapRotationRegionCsReq                        = 6818
	RotateMapScRsp                                     = 6841
	GetMapRotationDataCsReq                            = 6889
	RotateMapCsReq                                     = 6828
	RemoveRotaterCsReq                                 = 6862
	LeaveMapRotationRegionCsReq                        = 6819
	RemoveRotaterScRsp                                 = 6840
	InteractChargerCsReq                               = 6829
	UpdateEnergyScNotify                               = 6822
	UpdateMapRotationDataScNotify                      = 6891
	DeployRotaterScRsp                                 = 6849
	LeaveMapRotationRegionScNotify                     = 6883
	ResetMapRotationRegionScRsp                        = 6898
	UpdateRotaterScNotify                              = 6848
	LeaveMapRotationRegionScRsp                        = 6864
	ResetMapRotationRegionCsReq                        = 6846
	StartMatchCsReq                                    = 7319
	StartMatchScRsp                                    = 7329
	GetCrossInfoCsReq                                  = 7320
	GetCrossInfoScRsp                                  = 7312
	MatchResultScNotify                                = 7337
	CancelMatchScRsp                                   = 7345
	CancelMatchCsReq                                   = 7306
	MatchThreeSyncDataScNotify                         = 7437
	MatchThreeGetDataScRsp                             = 7429
	MatchThreeLevelEndCsReq                            = 7406
	MatchThreeSetBirdPosScRsp                          = 7412
	MatchThreeLevelEndScRsp                            = 7445
	MatchThreeSetBirdPosCsReq                          = 7420
	MatchThreeGetDataCsReq                             = 7419
	FinishSectionIdCsReq                               = 2728
	FinishPerformSectionIdCsReq                        = 2719
	GetNpcStatusCsReq                                  = 2729
	FinishPerformSectionIdScRsp                        = 2764
	GetNpcStatusScRsp                                  = 2737
	FinishItemIdScRsp                                  = 2749
	FinishSectionIdScRsp                               = 2741
	FinishItemIdCsReq                                  = 2758
	GetNpcMessageGroupScRsp                            = 2774
	GetNpcMessageGroupCsReq                            = 2718
	GetShareDataCsReq                                  = 4129
	SubmitOrigamiItemCsReq                             = 4183
	ShareCsReq                                         = 4118
	GetMovieRacingDataScRsp                            = 4108
	CancelCacheNotifyCsReq                             = 4119
	SecurityReportCsReq                                = 4189
	TriggerVoiceScRsp                                  = 4198
	SubmitOrigamiItemScRsp                             = 4122
	GetMovieRacingDataCsReq                            = 4173
	ShareScRsp                                         = 4174
	GetGunPlayDataScRsp                                = 4124
	TriggerVoiceCsReq                                  = 4146
	GetShareDataScRsp                                  = 4137
	TakePictureCsReq                                   = 4158
	GetGunPlayDataCsReq                                = 4177
	UpdateMovieRacingDataCsReq                         = 4111
	UpdateGunPlayDataCsReq                             = 4192
	CancelCacheNotifyScRsp                             = 4164
	TakePictureScRsp                                   = 4149
	UpdateMovieRacingDataScRsp                         = 4121
	UpdateGunPlayDataScRsp                             = 4147
	SecurityReportScRsp                                = 4157
	SetMissionEventProgressCsReq                       = 1221
	FinishCosumeItemMissionScRsp                       = 1298
	GetMissionEventDataCsReq                           = 1283
	GetMainMissionCustomValueScRsp                     = 1213
	UpdateTrackMainMissionIdScRsp                      = 1270
	StartFinishMainMissionScNotify                     = 1238
	SyncTaskScRsp                                      = 1228
	AcceptMainMissionScRsp                             = 1259
	AcceptMissionEventCsReq                            = 1262
	StartFinishSubMissionScNotify                      = 1261
	GetMainMissionCustomValueCsReq                     = 1202
	InterruptMissionEventCsReq                         = 1208
	FinishTalkMissionCsReq                             = 1229
	MissionAcceptScNotify                              = 1209
	AcceptMainMissionCsReq                             = 1282
	AcceptMissionEventScRsp                            = 1240
	TeleportToMissionResetPointCsReq                   = 1292
	MissionRewardScNotify                              = 1258
	SyncTaskCsReq                                      = 1249
	GetMissionDataCsReq                                = 1218
	InterruptMissionEventScRsp                         = 1211
	FinishTalkMissionScRsp                             = 1237
	GetMissionStatusCsReq                              = 1248
	TeleportToMissionResetPointScRsp                   = 1247
	FinishCosumeItemMissionCsReq                       = 1246
	GetMissionStatusScRsp                              = 1273
	SetMissionEventProgressScRsp                       = 1277
	DailyTaskDataScNotify                              = 1241
	MissionGroupWarnScNotify                           = 1257
	GetMissionEventDataScRsp                           = 1222
	SubMissionRewardScNotify                           = 1224
	MissionEventRewardScNotify                         = 1291
	UpdateTrackMainMissionIdCsReq                      = 1206
	GetMissionDataScRsp                                = 1274
	GetMbtiReportScRsp                                 = 7090
	MonopolySelectOptionCsReq                          = 7064
	GetMbtiReportCsReq                                 = 7033
	MonopolyLikeScRsp                                  = 7020
	MonopolyGameSettleScNotify                         = 7059
	DailyFirstEnterMonopolyActivityScRsp               = 7098
	MonopolyDailySettleScNotify                        = 7036
	MonopolyConfirmRandomScRsp                         = 7048
	MonopolyRollDiceScRsp                              = 7028
	MonopolyCheatDiceCsReq                             = 7092
	MonopolyGameRaiseRatioScRsp                        = 7082
	MonopolyScrachRaffleTicketScRsp                    = 7071
	MonopolyGetRegionProgressScRsp                     = 7054
	MonopolyGuessChooseCsReq                           = 7023
	MonopolyClickCellCsReq                             = 7035
	MonopolyGameRaiseRatioCsReq                        = 7038
	MonopolyAcceptQuizCsReq                            = 7006
	MonopolyLikeCsReq                                  = 7010
	DeleteSocialEventServerCacheCsReq                  = 7084
	MonopolyGetRegionProgressCsReq                     = 7095
	GetMonopolyDailyReportScRsp                        = 7052
	MonopolyCellUpdateNotify                           = 7037
	MonopolyContentUpdateScNotify                      = 7061
	MonopolyGetRaffleTicketCsReq                       = 7001
	MonopolyEventLoadUpdateScNotify                    = 7093
	MonopolyLikeScNotify                               = 7068
	MonopolyGiveUpCurContentScRsp                      = 7024
	MonopolyGuessDrawScNotify                          = 7075
	MonopolyRollRandomScRsp                            = 7022
	MonopolyClickMbtiReportScRsp                       = 7055
	MonopolyGuessBuyInformationScRsp                   = 7014
	MonopolyClickCellScRsp                             = 7069
	GetMonopolyMbtiReportRewardCsReq                   = 7026
	MonopolyEventSelectFriendScRsp                     = 7003
	MonopolySelectOptionScRsp                          = 7089
	MonopolyEventSelectFriendCsReq                     = 7067
	MonopolyReRollRandomScRsp                          = 7062
	MonopolyAcceptQuizScRsp                            = 7070
	DailyFirstEnterMonopolyActivityCsReq               = 7046
	MonopolyGuessChooseScRsp                           = 7087
	MonopolyTakePhaseRewardScRsp                       = 7085
	DeleteSocialEventServerCacheScRsp                  = 7076
	MonopolyGiveUpCurContentCsReq                      = 7077
	GetSocialEventServerCacheCsReq                     = 7094
	MonopolyClickMbtiReportCsReq                       = 7066
	MonopolyMoveCsReq                                  = 7041
	MonopolyGetDailyInitItemScRsp                      = 7060
	MonopolyGameGachaCsReq                             = 7002
	MonopolyBuyGoodsCsReq                              = 7073
	MonopolyScrachRaffleTicketCsReq                    = 7081
	MonopolyRollRandomCsReq                            = 7083
	GetMonopolyDailyReportCsReq                        = 7030
	MonopolyGetRafflePoolInfoScRsp                     = 7065
	MonopolyTakePhaseRewardCsReq                       = 7099
	MonopolySocialEventEffectScNotify                  = 7032
	MonopolyBuyGoodsScRsp                              = 7008
	GetMonopolyInfoScRsp                               = 7074
	GetSocialEventServerCacheScRsp                     = 7078
	MonopolyGetDailyInitItemCsReq                      = 7044
	MonopolyQuizDurationChangeScNotify                 = 7053
	MonopolyCheatDiceScRsp                             = 7047
	GetMonopolyFriendRankingListScRsp                  = 7100
	MonopolyReRollRandomCsReq                          = 7091
	GetMonopolyMbtiReportRewardScRsp                   = 7086
	MonopolyActionResultScNotify                       = 7029
	MonopolyGameBingoFlipCardScRsp                     = 7056
	MonopolyConditionUpdateScNotify                    = 7097
	MonopolyGameGachaScRsp                             = 7013
	MonopolyGetRaffleTicketScRsp                       = 7034
	MonopolyTakeRaffleTicketRewardCsReq                = 7017
	MonopolyGetRafflePoolInfoCsReq                     = 7031
	MonopolySttUpdateScNotify                          = 7005
	MonopolyMoveScRsp                                  = 7019
	MonopolyGameCreateScNotify                         = 7080
	MonopolyRollDiceCsReq                              = 7049
	MonopolyTakeRaffleTicketRewardScRsp                = 7051
	GetMonopolyFriendRankingListCsReq                  = 7088
	MonopolyGuessBuyInformationCsReq                   = 7039
	MonopolyConfirmRandomCsReq                         = 7040
	MonopolyGameBingoFlipCardCsReq                     = 7009
	GetMonopolyInfoCsReq                               = 7018
	MonopolyUpgradeAssetScRsp                          = 7021
	MonopolyUpgradeAssetCsReq                          = 7011
	MultiplayerFightGiveUpCsReq                        = 1058
	MultiplayerFightGiveUpScRsp                        = 1049
	MultiplayerGetFightGateCsReq                       = 1029
	MultiplayerMatch3FinishScNotify                    = 1019
	MultiplayerFightGameStateScRsp                     = 1074
	MultiplayerFightGameStateCsReq                     = 1018
	MultiplayerFightGameFinishScNotify                 = 1041
	MultiplayerFightGameStartScNotify                  = 1028
	MultiplayerGetFightGateScRsp                       = 1037
	GetPlayerReturnMultiDropInfoCsReq                  = 4637
	GetPlayerReturnMultiDropInfoScRsp                  = 4658
	MultipleDropInfoNotify                             = 4649
	GetMultipleDropInfoCsReq                           = 4618
	GetMultipleDropInfoScRsp                           = 4674
	MultipleDropInfoScNotify                           = 4629
	SetStuffToAreaCsReq                                = 4358
	BuyNpcStuffScRsp                                   = 4337
	UpgradeAreaScRsp                                   = 4398
	MuseumDispatchFinishedScNotify                     = 4321
	GetStuffScNotify                                   = 4319
	UpgradeAreaStatScRsp                               = 4322
	MuseumTakeCollectRewardCsReq                       = 4347
	UpgradeAreaStatCsReq                               = 4383
	MuseumRandomEventQueryCsReq                        = 4348
	MuseumRandomEventSelectCsReq                       = 4308
	GetMuseumInfoCsReq                                 = 4318
	GetExhibitScNotify                                 = 4364
	SetStuffToAreaScRsp                                = 4349
	MuseumFundsChangedScNotify                         = 4362
	RemoveStuffFromAreaCsReq                           = 4328
	BuyNpcStuffCsReq                                   = 4329
	FinishCurTurnCsReq                                 = 4389
	UpgradeAreaCsReq                                   = 4346
	MuseumRandomEventQueryScRsp                        = 4373
	MuseumRandomEventSelectScRsp                       = 4311
	MuseumTargetMissionFinishNotify                    = 4324
	MuseumTargetRewardNotify                           = 4392
	GetMuseumInfoScRsp                                 = 4374
	MuseumTakeCollectRewardScRsp                       = 4361
	MuseumRandomEventStartScNotify                     = 4340
	MuseumTargetStartNotify                            = 4377
	FinishCurTurnScRsp                                 = 4357
	MuseumInfoChangedScNotify                          = 4391
	RemoveStuffFromAreaScRsp                           = 4341
	TakeOfferingRewardCsReq                            = 6923
	OfferingInfoScNotify                               = 6929
	GetOfferingInfoCsReq                               = 6934
	GetOfferingInfoScRsp                               = 6937
	SubmitOfferingItemCsReq                            = 6933
	TakeOfferingRewardScRsp                            = 6924
	SubmitOfferingItemScRsp                            = 6925
	SyncAcceptedPamMissionNotify                       = 4029
	AcceptedPamMissionExpireScRsp                      = 4074
	AcceptedPamMissionExpireCsReq                      = 4018
	UnlockPhoneThemeScNotify                           = 5141
	SelectChatBubbleScRsp                              = 5137
	GetPhoneDataScRsp                                  = 5174
	UnlockChatBubbleScNotify                           = 5158
	SelectChatBubbleCsReq                              = 5129
	SelectPhoneThemeScRsp                              = 5128
	SelectPhoneThemeCsReq                              = 5149
	GetPhoneDataCsReq                                  = 5118
	SetGenderCsReq                                     = 70
	SetRedPointStatusScNotify                          = 17
	GetBasicInfoScRsp                                  = 45
	PlayerLoginFinishScRsp                             = 25
	SetLanguageCsReq                                   = 47
	UpdateFeatureSwitchScNotify                        = 12
	UnlockAvatarPathScRsp                              = 66
	GmTalkCsReq                                        = 64
	RegionStopScNotify                                 = 62
	UpdatePlayerSettingCsReq                           = 31
	ClientObjUploadCsReq                               = 85
	AvatarPathChangedNotify                            = 60
	SetAvatarPathCsReq                                 = 52
	PlayerLogoutCsReq                                  = 29
	StaminaInfoScNotify                                = 54
	SetPlayerInfoScRsp                                 = 87
	GetAuthkeyCsReq                                    = 22
	SetPlayerInfoCsReq                                 = 23
	PlayerLoginFinishCsReq                             = 76
	GateServerScNotify                                 = 67
	ExchangeStaminaScRsp                               = 83
	GetMultiPathAvatarInfoScRsp                        = 44
	ExchangeStaminaCsReq                               = 98
	GmTalkScNotify                                     = 41
	SetGameplayBirthdayScRsp                           = 88
	UpdatePlayerSettingScRsp                           = 65
	GmTalkScRsp                                        = 89
	ClientDownloadDataScNotify                         = 53
	PlayerGetTokenScRsp                                = 49
	UpdatePsnSettingsInfoCsReq                         = 5
	GetAuthkeyScRsp                                    = 91
	ServerAnnounceNotify                               = 38
	GetLevelRewardTakenListCsReq                       = 8
	AceAntiCheaterScRsp                                = 10
	GetMultiPathAvatarInfoCsReq                        = 69
	QueryProductInfoCsReq                              = 14
	AceAntiCheaterCsReq                                = 100
	PlayerLoginCsReq                                   = 18
	GetSecretKeyInfoCsReq                              = 78
	PlayerLogoutScRsp                                  = 37
	DailyRefreshNotify                                 = 16
	SetGenderScRsp                                     = 80
	FeatureSwitchClosedScNotify                        = 3
	PlayerKickOutScNotify                              = 19
	PlayerLoginScRsp                                   = 74
	PlayerHeartBeatScRsp                               = 90
	SetLanguageScRsp                                   = 61
	GetVideoVersionKeyCsReq                            = 1
	GetBasicInfoCsReq                                  = 42
	AntiAddictScNotify                                 = 40
	ReserveStaminaExchangeScRsp                        = 95
	ReserveStaminaExchangeCsReq                        = 71
	SetNicknameCsReq                                   = 48
	PlayerGetTokenCsReq                                = 58
	MonthCardRewardNotify                              = 20
	QueryProductInfoScRsp                              = 75
	GetSecretKeyInfoScRsp                              = 84
	RetcodeNotify                                      = 68
	PlayerHeartBeatCsReq                               = 33
	GetLevelRewardCsReq                                = 21
	ClientObjUploadScRsp                               = 93
	GetLevelRewardTakenListScRsp                       = 11
	SetAvatarPathScRsp                                 = 35
	GetLevelRewardScRsp                                = 77
	GetVideoVersionKeyScRsp                            = 34
	UpdatePsnSettingsInfoScRsp                         = 26
	ClientObjDownloadDataScNotify                      = 99
	SetGameplayBirthdayCsReq                           = 36
	SetNicknameScRsp                                   = 73
	UnlockAvatarPathCsReq                              = 97
	SetAssistAvatarScRsp                               = 2846
	UnlockHeadIconScNotify                             = 2819
	SetSignatureScRsp                                  = 2889
	SetDisplayAvatarScRsp                              = 2849
	SetDisplayAvatarCsReq                              = 2858
	SetHeadIconScRsp                                   = 2837
	SetHeadIconCsReq                                   = 2829
	SetAssistAvatarCsReq                               = 2857
	SetIsDisplayAvatarInfoScRsp                        = 2841
	SetIsDisplayAvatarInfoCsReq                        = 2828
	GetPlayerBoardDataCsReq                            = 2818
	GetPlayerBoardDataScRsp                            = 2874
	SetSignatureCsReq                                  = 2864
	PlayerReturnTakeRewardCsReq                        = 4528
	PlayerReturnTakePointRewardCsReq                   = 4558
	PlayerReturnPointChangeScNotify                    = 4537
	PlayerReturnTakeRewardScRsp                        = 4541
	PlayerReturnSignCsReq                              = 4574
	PlayerReturnInfoQueryScRsp                         = 4564
	PlayerReturnSignScRsp                              = 4529
	PlayerReturnInfoQueryCsReq                         = 4519
	PlayerReturnTakePointRewardScRsp                   = 4549
	PlayerReturnForceFinishScNotify                    = 4589
	PlayerReturnStartScNotify                          = 4518
	FinishPlotCsReq                                    = 1118
	FinishPlotScRsp                                    = 1174
	TakeKilledPunkLordMonsterScoreScRsp                = 3238
	PunkLordBattleResultScNotify                       = 3211
	TakePunkLordPointRewardCsReq                       = 3246
	GetPunkLordMonsterDataScRsp                        = 3274
	TakePunkLordPointRewardScRsp                       = 3298
	GetPunkLordBattleRecordScRsp                       = 3202
	GetKilledPunkLordMonsterDataCsReq                  = 3221
	StartPunkLordRaidCsReq                             = 3229
	StartPunkLordRaidScRsp                             = 3237
	SummonPunkLordMonsterScRsp                         = 3241
	SharePunkLordMonsterCsReq                          = 3258
	PunkLordMonsterInfoScNotify                        = 3283
	PunkLordRaidTimeOutScNotify                        = 3240
	TakeKilledPunkLordMonsterScoreCsReq                = 3261
	GetKilledPunkLordMonsterDataScRsp                  = 3277
	SharePunkLordMonsterScRsp                          = 3249
	PunkLordDataChangeNotify                           = 3282
	PunkLordMonsterKilledNotify                        = 3247
	GetPunkLordMonsterDataCsReq                        = 3218
	SummonPunkLordMonsterCsReq                         = 3228
	GetPunkLordBattleRecordCsReq                       = 3259
	GetPunkLordDataCsReq                               = 3222
	GetPunkLordDataScRsp                               = 3291
	GetQuestRecordScRsp                                = 941
	FinishQuestCsReq                                   = 964
	GetQuestDataCsReq                                  = 918
	BatchGetQuestDataCsReq                             = 983
	QuestRecordScNotify                                = 919
	TakeQuestRewardCsReq                               = 929
	GetQuestDataScRsp                                  = 974
	TakeQuestOptionalRewardCsReq                       = 957
	GetQuestRecordCsReq                                = 928
	TakeQuestOptionalRewardScRsp                       = 946
	BatchGetQuestDataScRsp                             = 922
	FinishQuestScRsp                                   = 989
	TakeQuestRewardScRsp                               = 937
	GetRaidInfoCsReq                                   = 2289
	GetAllSaveRaidScRsp                                = 2262
	DelSaveRaidScNotify                                = 2240
	StartRaidScRsp                                     = 2274
	LeaveRaidCsReq                                     = 2229
	GetSaveRaidScRsp                                   = 2222
	ChallengeRaidNotify                                = 2264
	RaidInfoNotify                                     = 2258
	SetClientRaidTargetCountScRsp                      = 2298
	GetRaidInfoScRsp                                   = 2257
	SetClientRaidTargetCountCsReq                      = 2246
	RaidKickByServerScNotify                           = 2248
	StartRaidCsReq                                     = 2218
	TakeChallengeRaidRewardScRsp                       = 2219
	LeaveRaidScRsp                                     = 2237
	TakeChallengeRaidRewardCsReq                       = 2241
	GetAllSaveRaidCsReq                                = 2291
	GetSaveRaidCsReq                                   = 2283
	GetChallengeRaidInfoScRsp                          = 2228
	GetChallengeRaidInfoCsReq                          = 2249
	RaidCollectionDataScRsp                            = 6957
	RaidCollectionDataCsReq                            = 6954
	RaidCollectionDataScNotify                         = 6953
	GetSingleRedDotParamGroupCsReq                     = 5958
	GetAllRedDotDataCsReq                              = 5918
	UpdateRedDotDataScRsp                              = 5937
	UpdateRedDotDataCsReq                              = 5929
	GetSingleRedDotParamGroupScRsp                     = 5949
	GetAllRedDotDataScRsp                              = 5974
	GetReplayTokenScRsp                                = 3574
	GetPlayerReplayInfoCsReq                           = 3529
	GetReplayTokenCsReq                                = 3518
	GetPlayerReplayInfoScRsp                           = 3537
	DailyFirstMeetPamCsReq                             = 3429
	GetRndOptionCsReq                                  = 3418
	DailyFirstMeetPamScRsp                             = 3437
	GetRndOptionScRsp                                  = 3474
	PickRogueAvatarCsReq                               = 1822
	FinishAeonDialogueGroupScRsp                       = 1860
	ExchangeRogueRewardKeyScRsp                        = 1890
	SyncRogueStatusScNotify                            = 1843
	ReviveRogueAvatarCsReq                             = 1840
	SyncRogueFinishScNotify                            = 1883
	EnableRogueTalentScRsp                             = 1850
	GetRogueInitialScoreCsReq                          = 1887
	SyncRogueExploreWinScNotify                        = 1809
	EnterRogueMapRoomScRsp                             = 1823
	EnterRogueMapRoomCsReq                             = 1880
	StartRogueScRsp                                    = 1837
	GetRogueScoreRewardInfoCsReq                       = 1899
	EnhanceRogueBuffCsReq                              = 1877
	OpenRogueChestScRsp                                = 1868
	GetRogueScoreRewardInfoScRsp                       = 1885
	ReviveRogueAvatarScRsp                             = 1848
	SyncRogueSeasonFinishScNotify                      = 1856
	QuitRogueCsReq                                     = 1859
	GetRogueBuffEnhanceInfoScRsp                       = 1821
	GetRogueTalentInfoScRsp                            = 1866
	FinishAeonDialogueGroupCsReq                       = 1844
	OpenRogueChestCsReq                                = 1820
	StartRogueCsReq                                    = 1829
	GetRogueInfoScRsp                                  = 1874
	GetRogueTalentInfoCsReq                            = 1897
	SyncRogueReviveInfoScNotify                        = 1882
	EnableRogueTalentCsReq                             = 1855
	TakeRogueScoreRewardCsReq                          = 1873
	SyncRogueRewardInfoScNotify                        = 1827
	TakeRogueAeonLevelRewardCsReq                      = 1881
	SyncRogueAeonLevelUpRewardScNotify                 = 1865
	SyncRogueVirtualItemInfoScNotify                   = 1863
	EnhanceRogueBuffScRsp                              = 1824
	GetRogueInfoCsReq                                  = 1818
	QuitRogueScRsp                                     = 1802
	ExchangeRogueRewardKeyCsReq                        = 1833
	EnterRogueCsReq                                    = 1858
	SyncRogueGetItemScNotify                           = 1851
	EnterRogueScRsp                                    = 1849
	LeaveRogueScRsp                                    = 1841
	TakeRogueScoreRewardScRsp                          = 1808
	SyncRogueMapRoomScNotify                           = 1814
	SyncRogueAeonScNotify                              = 1834
	GetRogueAeonInfoCsReq                              = 1835
	SyncRogueAreaUnlockScNotify                        = 1817
	GetRogueBuffEnhanceInfoCsReq                       = 1811
	SyncRoguePickAvatarInfoScNotify                    = 1872
	GetRogueInitialScoreScRsp                          = 1839
	LeaveRogueCsReq                                    = 1828
	TakeRogueAeonLevelRewardScRsp                      = 1871
	PickRogueAvatarScRsp                               = 1891
	GetRogueAeonInfoScRsp                              = 1869
	SyncRogueCommonVirtualItemInfoScNotify             = 5645
	BuyRogueShopBuffScRsp                              = 5689
	GetRogueAdventureRoomInfoScRsp                     = 5683
	BuyRogueShopBuffCsReq                              = 5664
	RogueDoGambleScRsp                                 = 5669
	RogueWorkbenchSelectFuncCsReq                      = 5644
	RogueWorkbenchHandleFuncScRsp                      = 5651
	GetRogueHandbookDataScRsp                          = 5670
	BuyRogueShopMiracleCsReq                           = 5641
	SyncRogueHandbookDataUpdateScNotify                = 5680
	PrepareRogueAdventureRoomScRsp                     = 5629
	TakeRogueMiracleHandbookRewardScRsp                = 5687
	SetRogueExhibitionScRsp                            = 5685
	CommonRogueComponentUpdateScNotify                 = 5625
	SyncRogueCommonActionResultScNotify                = 5675
	ExchangeRogueBuffWithMiracleScRsp                  = 5648
	SyncRogueCommonDialogueDataScNotify                = 5684
	GetRogueShopMiracleInfoScRsp                       = 5658
	CommonRogueQueryScRsp                              = 5668
	GetRogueExhibitionScRsp                            = 5665
	SetRogueCollectionScRsp                            = 5654
	ExchangeRogueBuffWithMiracleCsReq                  = 5640
	GetRogueShopBuffInfoCsReq                          = 5649
	TakeRogueEventHandbookRewardScRsp                  = 5614
	GetRogueExhibitionCsReq                            = 5631
	TakeRogueEventHandbookRewardCsReq                  = 5639
	CommonRogueQueryCsReq                              = 5620
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5673
	HandleRogueCommonPendingActionScRsp                = 5610
	RogueGetGambleInfoScRsp                            = 5652
	RogueGetGambleInfoCsReq                            = 5630
	SyncRogueAdventureRoomInfoScNotify                 = 5618
	UpdateRogueAdventureRoomScoreCsReq                 = 5612
	RogueDoGambleCsReq                                 = 5635
	RogueWorkbenchGetInfoCsReq                         = 5601
	RogueNpcDisappearScRsp                             = 5646
	GetRogueShopBuffInfoScRsp                          = 5628
	SyncRogueCommonPendingActionScNotify               = 5653
	GetRogueCommonDialogueDataScRsp                    = 5667
	StopRogueAdventureRoomCsReq                        = 5677
	RogueWorkbenchGetInfoScRsp                         = 5634
	UpdateRogueAdventureRoomScoreScRsp                 = 5642
	StopRogueAdventureRoomScRsp                        = 5624
	SyncRogueCommonDialogueOptionFinishScNotify        = 5676
	SelectRogueCommonDialogueOptionScRsp               = 5632
	RogueWorkbenchHandleFuncCsReq                      = 5617
	GetRogueCommonDialogueDataCsReq                    = 5690
	SelectRogueCommonDialogueOptionCsReq               = 5603
	FinishRogueCommonDialogueCsReq                     = 5694
	TakeRogueMiracleHandbookRewardCsReq                = 5623
	PrepareRogueAdventureRoomCsReq                     = 5674
	GetRogueHandbookDataCsReq                          = 5606
	GetRogueCollectionScRsp                            = 5671
	EnhanceCommonRogueBuffCsReq                        = 5611
	BuyRogueShopMiracleScRsp                           = 5619
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5608
	SetRogueCollectionCsReq                            = 5695
	GetRogueCollectionCsReq                            = 5681
	EnhanceCommonRogueBuffScRsp                        = 5621
	CommonRogueUpdateScNotify                          = 5633
	SetRogueExhibitionCsReq                            = 5699
	GetRogueAdventureRoomInfoCsReq                     = 5698
	FinishRogueCommonDialogueScRsp                     = 5678
	RogueNpcDisappearCsReq                             = 5657
	RogueWorkbenchSelectFuncScRsp                      = 5660
	HandleRogueCommonPendingActionCsReq                = 5700
	GetRogueShopMiracleInfoCsReq                       = 5637
	TakeRogueEndlessActivityPointRewardCsReq           = 6008
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6004
	GetRogueEndlessActivityDataCsReq                   = 6002
	EnterRogueEndlessActivityStageScRsp                = 6010
	RogueEndlessActivityBattleEndScNotify              = 6006
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6009
	EnterRogueEndlessActivityStageCsReq                = 6007
	GetRogueEndlessActivityDataScRsp                   = 6005
	TakeRogueEndlessActivityPointRewardScRsp           = 6003
	RogueModifierAddNotify                             = 5329
	RogueModifierDelNotify                             = 5319
	RogueModifierSelectCellCsReq                       = 5337
	RogueModifierUpdateNotify                          = 5341
	RogueModifierStageStartNotify                      = 5364
	RogueModifierSelectCellScRsp                       = 5358
	RogueTournResetPermanentTalentScRsp                = 6087
	RogueTournGetMiscRealTimeDataScRsp                 = 6042
	RogueTournDifficultyCompNotify                     = 6100
	RogueTournEnablePermanentTalentScRsp               = 6012
	RogueTournLeaveRogueCocoonSceneScRsp               = 6022
	RogueTournGetCurRogueCocoonInfoScRsp               = 6061
	RogueTournTakeExpRewardScRsp                       = 6081
	RogueTournGetCurRogueCocoonInfoCsReq               = 6090
	RogueTournGetAllArchiveScRsp                       = 6034
	RogueTournReEnterRogueCocoonStageScRsp             = 6074
	RogueTournGetArchiveRepositoryCsReq                = 6043
	RogueTournStartCsReq                               = 6031
	RogueTournGetAllArchiveCsReq                       = 6058
	RogueTournEnterScRsp                               = 6028
	RogueTournReviveCostUpdateScNotify                 = 6011
	RogueTournLeaveCsReq                               = 6068
	RogueTournResetPermanentTalentCsReq                = 6066
	RogueTournLeaveScRsp                               = 6098
	RogueTournWeekChallengeUpdateScNotify              = 6078
	RogueTournHandBookNotify                           = 6020
	RogueTournQueryScRsp                               = 6018
	RogueTournTakeExpRewardCsReq                       = 6016
	RogueTournReviveAvatarScRsp                        = 6071
	RogueTournBattleFailSettleInfoScNotify             = 6027
	RogueTournEnablePermanentTalentCsReq               = 6057
	RogueTournClearArchiveNameScNotify                 = 6023
	RogueTournEnterLayerCsReq                          = 6072
	RogueTournConfirmSettleCsReq                       = 6091
	RogueTournGetSettleInfoCsReq                       = 6044
	RogueTournEnterRogueCocoonSceneScRsp               = 6035
	RogueTournReEnterRogueCocoonStageCsReq             = 6026
	RogueTournEnterRoomCsReq                           = 6033
	RogueTournStartScRsp                               = 6093
	RogueTournGetPermanentTalentInfoCsReq              = 6088
	RogueTournGetArchiveRepositoryScRsp                = 6052
	RogueTournGetSettleInfoScRsp                       = 6013
	RogueTournConfirmSettleScRsp                       = 6094
	RogueTournQueryCsReq                               = 6059
	RogueTournReviveAvatarCsReq                        = 6083
	RogueTournAreaUpdateScNotify                       = 6021
	RogueTournExpNotify                                = 6030
	RogueTournLeaveRogueCocoonSceneCsReq               = 6050
	RogueTournGetMiscRealTimeDataCsReq                 = 6084
	RogueTournGetPermanentTalentInfoScRsp              = 6092
	RogueTournDeleteArchiveCsReq                       = 6099
	RogueTournLevelInfoUpdateScNotify                  = 6055
	RogueTournEnterCsReq                               = 6047
	RogueTournEnterLayerScRsp                          = 6038
	RogueTournEnterRogueCocoonSceneCsReq               = 6024
	RogueTournEnterRoomScRsp                           = 6029
	RogueTournRenameArchiveScRsp                       = 6056
	RogueTournSettleScRsp                              = 6069
	RogueTournSettleCsReq                              = 6039
	RogueTournRenameArchiveCsReq                       = 6048
	RogueTournDeleteArchiveScRsp                       = 6063
	DoGachaInRollShopScRsp                             = 6905
	GetRollShopInfoCsReq                               = 6914
	TakeRollShopRewardCsReq                            = 6903
	DoGachaInRollShopCsReq                             = 6913
	GetRollShopInfoScRsp                               = 6917
	TakeRollShopRewardScRsp                            = 6904
	SceneEntityMoveCsReq                               = 1418
	SceneCastSkillScRsp                                = 1449
	EntityBindPropCsReq                                = 1470
	EnteredSceneChangeScNotify                         = 1460
	ReturnLastTownScRsp                                = 1408
	EnterSceneCsReq                                    = 1425
	SetClientPausedScRsp                               = 1487
	GameplayCounterCountDownScRsp                      = 1485
	StartTimedFarmElementScRsp                         = 1443
	GameplayCounterRecoverScRsp                        = 1486
	RecoverAllLineupCsReq                              = 1402
	SpringRecoverCsReq                                 = 1488
	GroupStateChangeCsReq                              = 1430
	DeleteSummonUnitCsReq                              = 1450
	SceneEnterStageCsReq                               = 1411
	SceneCastSkillCsReq                                = 1458
	ScenePlaneEventScNotify                            = 1417
	SetGroupCustomSaveDataScRsp                        = 1467
	EnterSectionCsReq                                  = 1492
	SpringRefreshScRsp                                 = 1440
	InteractPropScRsp                                  = 1437
	SyncEntityBuffChangeListScNotify                   = 1446
	UpdateMechanismBarScNotify                         = 1433
	SetGroupCustomSaveDataCsReq                        = 1490
	SpringRecoverSingleAvatarScRsp                     = 1468
	GetSceneMapInfoScRsp                               = 1481
	InteractPropCsReq                                  = 1429
	SetCurInteractEntityScRsp                          = 1459
	StartTimedCocoonStageScRsp                         = 1496
	SceneEntityTeleportCsReq                           = 1484
	ReEnterLastElementStageScRsp                       = 1478
	GetSpringRecoverDataScRsp                          = 1445
	RecoverAllLineupScRsp                              = 1413
	EnterSceneByServerScNotify                         = 1434
	ReEnterLastElementStageCsReq                       = 1494
	GetUnlockTeleportScRsp                             = 1454
	GroupStateChangeScNotify                           = 1435
	RefreshTriggerByClientScRsp                        = 1466
	StartCocoonStageScRsp                              = 1406
	GetCurSceneInfoCsReq                               = 1428
	GetCurSceneInfoScRsp                               = 1441
	SceneReviveAfterRebattleCsReq                      = 1477
	ReturnLastTownCsReq                                = 1473
	SetCurInteractEntityCsReq                          = 1482
	DeactivateFarmElementCsReq                         = 1414
	EnterSceneScRsp                                    = 1401
	SetSpringRecoverConfigScRsp                        = 1436
	SceneGroupRefreshScNotify                          = 1405
	SceneUpdatePositionVersionNotify                   = 1457
	LastSpringRefreshTimeNotify                        = 1448
	StartTimedFarmElementCsReq                         = 1463
	SetClientPausedCsReq                               = 1423
	HealPoolInfoNotify                                 = 1410
	UnlockedAreaMapScNotify                            = 1479
	SpringRecoverScRsp                                 = 1500
	SavePointsInfoNotify                               = 1409
	DeactivateFarmElementScRsp                         = 1475
	SceneReviveAfterRebattleScRsp                      = 1424
	RefreshTriggerByClientCsReq                        = 1497
	GroupStateChangeScRsp                              = 1452
	UpdateFloorSavedValueNotify                        = 1465
	EnterSectionScRsp                                  = 1447
	GetEnteredSceneScRsp                               = 1444
	RefreshTriggerByClientScNotify                     = 1455
	GetSpringRecoverDataCsReq                          = 1442
	ActivateFarmElementCsReq                           = 1453
	GameplayCounterCountDownCsReq                      = 1499
	ActivateFarmElementScRsp                           = 1412
	GetSceneMapInfoCsReq                               = 1451
	SceneCastSkillCostMpCsReq                          = 1498
	SetSpringRecoverConfigCsReq                        = 1416
	SyncServerSceneChangeNotify                        = 1471
	SceneCastSkillMpUpdateScNotify                     = 1422
	SceneCastSkillCostMpScRsp                          = 1483
	StartCocoonStageCsReq                              = 1456
	GetUnlockTeleportCsReq                             = 1495
	SceneEntityMoveScRsp                               = 1474
	GameplayCounterUpdateScNotify                      = 1493
	SceneEnterStageScRsp                               = 1421
	SpringRecoverSingleAvatarCsReq                     = 1420
	SpringRefreshCsReq                                 = 1462
	UnlockTeleportNotify                               = 1427
	GameplayCounterRecoverCsReq                        = 1426
	EntityBindPropScRsp                                = 1480
	DeleteSummonUnitScRsp                              = 1404
	SceneEntityMoveScNotify                            = 1489
	SceneEntityTeleportScRsp                           = 1476
	StartTimedCocoonStageCsReq                         = 1415
	GetEnteredSceneCsReq                               = 1469
	GetAllServerPrefsDataScRsp                         = 6174
	UpdateServerPrefsDataScRsp                         = 6149
	GetAllServerPrefsDataCsReq                         = 6118
	GetServerPrefsDataScRsp                            = 6137
	UpdateServerPrefsDataCsReq                         = 6158
	GetServerPrefsDataCsReq                            = 6129
	TakeCityShopRewardCsReq                            = 1558
	BuyGoodsCsReq                                      = 1529
	GetShopListCsReq                                   = 1518
	CityShopInfoScNotify                               = 1528
	TakeCityShopRewardScRsp                            = 1549
	BuyGoodsScRsp                                      = 1537
	GetShopListScRsp                                   = 1574
	SpaceZooMutateScRsp                                = 6749
	SpaceZooOpCatteryScRsp                             = 6741
	SpaceZooMutateCsReq                                = 6758
	SpaceZooBornScRsp                                  = 6737
	SpaceZooTakeScRsp                                  = 6783
	SpaceZooExchangeItemScRsp                          = 6746
	SpaceZooExchangeItemCsReq                          = 6757
	SpaceZooTakeCsReq                                  = 6798
	SpaceZooDataScRsp                                  = 6774
	SpaceZooOpCatteryCsReq                             = 6728
	SpaceZooDeleteCatScRsp                             = 6764
	SpaceZooCatUpdateNotify                            = 6789
	SpaceZooDeleteCatCsReq                             = 6719
	SpaceZooDataCsReq                                  = 6718
	SpaceZooBornCsReq                                  = 6729
	GetStarFightDataCsReq                              = 7162
	StarFightDataChangeNotify                          = 7166
	StartStarFightLevelScRsp                           = 7170
	GetStarFightDataScRsp                              = 7165
	StartStarFightLevelCsReq                           = 7167
	StoryLineTrialAvatarChangeScNotify                 = 6228
	ChangeStoryLineFinishScNotify                      = 6249
	StoryLineInfoScNotify                              = 6229
	GetStoryLineInfoScRsp                              = 6274
	GetStoryLineInfoCsReq                              = 6218
	GetStrongChallengeActivityDataScRsp                = 6674
	GetStrongChallengeActivityDataCsReq                = 6618
	EnterStrongChallengeActivityStageCsReq             = 6629
	EnterStrongChallengeActivityStageScRsp             = 6637
	StrongChallengeActivityBattleEndScNotify           = 6658
	SummonActivityBattleEndScNotify                    = 7566
	EnterSummonActivityStageScRsp                      = 7570
	EnterSummonActivityStageCsReq                      = 7567
	GetSummonActivityDataCsReq                         = 7562
	GetSummonActivityDataScRsp                         = 7565
	SwordTrainingRestoreGameCsReq                      = 7476
	GetSwordTrainingDataScRsp                          = 7456
	SwordTrainingGiveUpGameCsReq                       = 7467
	SwordTrainingLearnSkillScRsp                       = 7481
	SwordTrainingSetSkillTraceScRsp                    = 7472
	SwordTrainingExamResultConfirmCsReq                = 7488
	SwordTrainingDialogueSelectOptionScRsp             = 7459
	SwordTrainingExamResultConfirmScRsp                = 7480
	SwordTrainingFinishEndingHintScRsp                 = 7492
	SwordTrainingTakeEndingRewardCsReq                 = 7497
	SwordTrainingResumeGameCsReq                       = 7452
	SwordTrainingStoryBattleScRsp                      = 7475
	SwordTrainingFinishEndingHintCsReq                 = 7493
	SwordTrainingGameSettleScNotify                    = 7473
	SwordTrainingResumeGameScRsp                       = 7490
	EnterSwordTrainingExamScRsp                        = 7460
	SwordTrainingSelectEndingCsReq                     = 7457
	SwordTrainingDailyPhaseConfirmCsReq                = 7470
	SwordTrainingSetSkillTraceCsReq                    = 7482
	SwordTrainingRestoreGameScRsp                      = 7468
	SwordTrainingTurnActionScRsp                       = 7487
	SwordTrainingStartGameCsReq                        = 7474
	SwordTrainingUpdateRankScRsp                       = 7496
	SwordTrainingActionTurnSettleScNotify              = 7485
	SwordTrainingTurnActionCsReq                       = 7495
	SwordTrainingGetSkillInfoScRsp                     = 7465
	SwordTrainingStoryConfirmScRsp                     = 7494
	SwordTrainingDialogueSelectOptionCsReq             = 7471
	SwordTrainingSelectEndingScRsp                     = 7477
	SwordTrainingStoryBattleCsReq                      = 7478
	SwordTrainingUnlockSyncScNotify                    = 7484
	SwordTrainingRefreshPartnerAbilityCsReq            = 7500
	SwordTrainingDailyPhaseConfirmScRsp                = 7462
	EnterSwordTrainingExamCsReq                        = 7491
	SwordTrainingGiveUpGameScRsp                       = 7499
	SwordTrainingGetSkillInfoCsReq                     = 7461
	SwordTrainingLearnSkillCsReq                       = 7463
	SwordTrainingGameSyncChangeScNotify                = 7469
	SwordTrainingUpdateRankCsReq                       = 7458
	SwordTrainingRefreshPartnerAbilityScRsp            = 7451
	SwordTrainingStartGameScRsp                        = 7466
	SwordTrainingStoryConfirmCsReq                     = 7498
	GetSwordTrainingDataCsReq                          = 7479
	SwordTrainingTakeEndingRewardScRsp                 = 7483
	PlayerSyncScNotify                                 = 618
	GetNpcTakenRewardScRsp                             = 2174
	TakeTalkRewardScRsp                                = 2137
	SelectInclinationTextCsReq                         = 2119
	GetNpcTakenRewardCsReq                             = 2118
	GetFirstTalkNpcScRsp                               = 2149
	TakeTalkRewardCsReq                                = 2129
	GetFirstTalkByPerformanceNpcCsReq                  = 2189
	GetFirstTalkByPerformanceNpcScRsp                  = 2157
	SelectInclinationTextScRsp                         = 2164
	FinishFirstTalkNpcCsReq                            = 2128
	FinishFirstTalkByPerformanceNpcCsReq               = 2146
	FinishFirstTalkByPerformanceNpcScRsp               = 2198
	GetFirstTalkNpcCsReq                               = 2158
	FinishFirstTalkNpcScRsp                            = 2141
	TelevisionActivityDataChangeScNotify               = 6973
	GetTelevisionActivityDataScRsp                     = 6977
	EnterTelevisionActivityStageScRsp                  = 6963
	EnterTelevisionActivityStageCsReq                  = 6965
	GetTelevisionActivityDataCsReq                     = 6974
	TelevisionActivityBattleEndScNotify                = 6964
	TextJoinQueryCsReq                                 = 3829
	TextJoinSaveScRsp                                  = 3874
	TextJoinQueryScRsp                                 = 3837
	TextJoinBatchSaveScRsp                             = 3849
	TextJoinBatchSaveCsReq                             = 3858
	TextJoinSaveCsReq                                  = 3818
	GetTrackPhotoActivityDataCsReq                     = 7552
	SettleTrackPhotoStageCsReq                         = 7557
	GetTrackPhotoActivityDataScRsp                     = 7555
	QuitTrackPhotoStageCsReq                           = 7553
	SettleTrackPhotoStageScRsp                         = 7560
	QuitTrackPhotoStageScRsp                           = 7554
	StartTrackPhotoStageCsReq                          = 7556
	StartTrackPhotoStageScRsp                          = 7558
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3719
	TrainVisitorRewardSendNotify                       = 3749
	GetTrainVisitorRegisterScRsp                       = 3741
	TrainVisitorBehaviorFinishScRsp                    = 3774
	GetTrainVisitorBehaviorScRsp                       = 3737
	TrainVisitorBehaviorFinishCsReq                    = 3718
	GetTrainVisitorRegisterCsReq                       = 3728
	GetTrainVisitorBehaviorCsReq                       = 3729
	ShowNewSupplementVisitorScRsp                      = 3757
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3764
	TrainRefreshTimeNotify                             = 3758
	ShowNewSupplementVisitorCsReq                      = 3789
	TravelBrochureSetCustomValueScRsp                  = 6422
	TravelBrochureSetPageDescStatusScRsp               = 6462
	TravelBrochureUpdatePasterPosCsReq                 = 6489
	TravelBrochureSelectMessageScRsp                   = 6449
	TravelBrochurePageResetCsReq                       = 6440
	TravelBrochureGetDataCsReq                         = 6418
	TravelBrochureGetPasterScNotify                    = 6446
	TravelBrochureApplyPasterListScRsp                 = 6408
	TravelBrochureRemovePasterScRsp                    = 6464
	TravelBrochureGetDataScRsp                         = 6474
	TravelBrochureUpdatePasterPosScRsp                 = 6457
	TravelBrochureRemovePasterCsReq                    = 6419
	TravelBrochurePageUnlockScNotify                   = 6429
	TravelBrochureApplyPasterScRsp                     = 6441
	TravelBrochureSetCustomValueCsReq                  = 6483
	TravelBrochureApplyPasterCsReq                     = 6428
	TravelBrochureSetPageDescStatusCsReq               = 6491
	TravelBrochurePageResetScRsp                       = 6448
	TravelBrochureSelectMessageCsReq                   = 6458
	TravelBrochureApplyPasterListCsReq                 = 6473
	QuitTreasureDungeonCsReq                           = 4411
	FightTreasureDungeonMonsterCsReq                   = 4491
	InteractTreasureDungeonGridCsReq                   = 4440
	TreasureDungeonDataScNotify                        = 4418
	OpenTreasureDungeonGridScRsp                       = 4422
	GetTreasureDungeonActivityDataScRsp                = 4457
	UseTreasureDungeonItemScRsp                        = 4408
	QuitTreasureDungeonScRsp                           = 4421
	EnterTreasureDungeonCsReq                          = 4446
	OpenTreasureDungeonGridCsReq                       = 4483
	UseTreasureDungeonItemCsReq                        = 4473
	GetTreasureDungeonActivityDataCsReq                = 4489
	FightTreasureDungeonMonsterScRsp                   = 4462
	TreasureDungeonFinishScNotify                      = 4474
	EnterTreasureDungeonScRsp                          = 4498
	InteractTreasureDungeonGridScRsp                   = 4448
	UnlockTutorialScRsp                                = 1649
	UnlockTutorialGuideCsReq                           = 1628
	FinishTutorialCsReq                                = 1619
	GetTutorialGuideScRsp                              = 1637
	GetTutorialCsReq                                   = 1618
	FinishTutorialGuideCsReq                           = 1689
	GetTutorialGuideCsReq                              = 1629
	UnlockTutorialCsReq                                = 1658
	FinishTutorialGuideScRsp                           = 1657
	GetTutorialScRsp                                   = 1674
	FinishTutorialScRsp                                = 1664
	UnlockTutorialGuideScRsp                           = 1641
	TakeChapterRewardScRsp                             = 419
	SetCurWaypointCsReq                                = 429
	SetCurWaypointScRsp                                = 437
	GetChapterScRsp                                    = 449
	GetWaypointScRsp                                   = 474
	GetChapterCsReq                                    = 458
	WaypointShowNewCsNotify                            = 428
	TakeChapterRewardCsReq                             = 441
	GetWaypointCsReq                                   = 418
	GetWolfBroGameDataScRsp                            = 6564
	StartWolfBroGameCsReq                              = 6518
	GetWolfBroGameDataCsReq                            = 6519
	WolfBroGameDataChangeScNotify                      = 6589
	RestoreWolfBroGameArchiveCsReq                     = 6558
	WolfBroGameActivateBulletScRsp                     = 6591
	StartWolfBroGameScRsp                              = 6574
	QuitWolfBroGameScRsp                               = 6541
	WolfBroGameUseBulletScRsp                          = 6546
	WolfBroGameExplodeMonsterScRsp                     = 6540
	WolfBroGamePickupBulletCsReq                       = 6598
	WolfBroGameExplodeMonsterCsReq                     = 6562
	ArchiveWolfBroGameScRsp                            = 6537
	WolfBroGamePickupBulletScRsp                       = 6583
	WolfBroGameUseBulletCsReq                          = 6557
	QuitWolfBroGameCsReq                               = 6528
	ArchiveWolfBroGameCsReq                            = 6529
	WolfBroGameActivateBulletCsReq                     = 6522
	RestoreWolfBroGameArchiveScRsp                     = 6549
)

const (
	ServiceConnectionReq = 10000
	ServiceConnectionRsp = 10100
	GateLoginGameRsp     = 10001
	GateLoginGameReq     = 10101
	GateToGameMsgNotify  = 10002
	GameToGateMsgNotify  = 10102
	GetAllServiceGateReq = 10003
	GetAllServiceGateRsp = 10103
	MultiToNodePingReq   = 10004
	MultiToNodePingRsp   = 10104
	MuipToNodePingReq    = 10005
	MuipToNodePingRsp    = 10105
	// GetAllServiceGameReq     = 10006
	// GetAllServiceGameRsp     = 10106
	GameToNodePingReq        = 10007
	GameToNodePingRsp        = 10107
	GateGamePingReq          = 10008
	GateGamePingRsp          = 10108
	GateGamePlayerLoginReq   = 10009
	GateGamePlayerLoginRsp   = 10109
	GetToGamePlayerLogoutReq = 10010
	GetToGamePlayerLogoutRsp = 10110
	GateLoginMultiReq        = 10011
	GateLoginMultiRsp        = 10111
	GateToNodePingReq        = 10012
	GateToNodePingRsp        = 10112

	GateToGamePlayerLogoutNotify = 11000
	PlayerMsgGateToNodeNotify    = 11001
	// PlayerLoginNotify            = 11002
	// NodeToGsPlayerLogoutNotify   = 11003
	GameToGatePlayerLogoutNotify = 11004

	GmGive       = 12001
	GmWorldLevel = 12002
	DelItem      = 12003
	MaxCurAvatar = 12004
	GmMission    = 12005
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(TakeMonsterResearchActivityRewardCsReq, func() any { return new(proto.TakeMonsterResearchActivityRewardCsReq) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(SubmitMonsterResearchActivityMaterialCsReq, func() any { return new(proto.SubmitMonsterResearchActivityMaterialCsReq) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(TakeMonsterResearchActivityRewardScRsp, func() any { return new(proto.TakeMonsterResearchActivityRewardScRsp) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(GetMonsterResearchActivityDataScRsp, func() any { return new(proto.GetMonsterResearchActivityDataScRsp) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(GetMonsterResearchActivityDataCsReq, func() any { return new(proto.GetMonsterResearchActivityDataCsReq) })
	c.regMsg(SubmitMonsterResearchActivityMaterialScRsp, func() any { return new(proto.SubmitMonsterResearchActivityMaterialScRsp) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(LogisticsGameCsReq, func() any { return new(proto.LogisticsGameCsReq) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(GetUpdatedArchiveDataCsReq, func() any { return new(proto.GetUpdatedArchiveDataCsReq) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(GetUpdatedArchiveDataScRsp, func() any { return new(proto.GetUpdatedArchiveDataScRsp) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(MarkAvatarCsReq, func() any { return new(proto.MarkAvatarCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(LevelUpSpecialSkillTreeCsReq, func() any { return new(proto.LevelUpSpecialSkillTreeCsReq) })
	c.regMsg(UnlockSpecialSkillTreeScNotify, func() any { return new(proto.UnlockSpecialSkillTreeScNotify) })
	c.regMsg(LevelUpSpecialSkillTreeScRsp, func() any { return new(proto.LevelUpSpecialSkillTreeScRsp) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(MarkAvatarScRsp, func() any { return new(proto.MarkAvatarScRsp) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(QuitBattleScRsp, func() any { return new(proto.QuitBattleScRsp) })
	c.regMsg(ServerSimulateBattleFinishScNotify, func() any { return new(proto.ServerSimulateBattleFinishScNotify) })
	c.regMsg(BattleLogReportCsReq, func() any { return new(proto.BattleLogReportCsReq) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(ReBattleAfterBattleLoseCsNotify, func() any { return new(proto.ReBattleAfterBattleLoseCsNotify) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(QuitBattleScNotify, func() any { return new(proto.QuitBattleScNotify) })
	c.regMsg(BattleLogReportScRsp, func() any { return new(proto.BattleLogReportScRsp) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(QuitBattleCsReq, func() any { return new(proto.QuitBattleCsReq) })
	c.regMsg(RebattleByClientCsNotify, func() any { return new(proto.RebattleByClientCsNotify) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(GetBattleCollegeDataCsReq, func() any { return new(proto.GetBattleCollegeDataCsReq) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(MatchBoxingClubOpponentCsReq, func() any { return new(proto.MatchBoxingClubOpponentCsReq) })
	c.regMsg(GetBoxingClubInfoCsReq, func() any { return new(proto.GetBoxingClubInfoCsReq) })
	c.regMsg(GiveUpBoxingClubChallengeScRsp, func() any { return new(proto.GiveUpBoxingClubChallengeScRsp) })
	c.regMsg(SetBoxingClubResonanceLineupCsReq, func() any { return new(proto.SetBoxingClubResonanceLineupCsReq) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(BoxingClubRewardScNotify, func() any { return new(proto.BoxingClubRewardScNotify) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(ChooseBoxingClubResonanceCsReq, func() any { return new(proto.ChooseBoxingClubResonanceCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(BoxingClubChallengeUpdateScNotify, func() any { return new(proto.BoxingClubChallengeUpdateScNotify) })
	c.regMsg(StartBoxingClubBattleCsReq, func() any { return new(proto.StartBoxingClubBattleCsReq) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(GiveUpBoxingClubChallengeCsReq, func() any { return new(proto.GiveUpBoxingClubChallengeCsReq) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffCsReq, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffCsReq) })
	c.regMsg(GetChallengeGroupStatisticsScRsp, func() any { return new(proto.GetChallengeGroupStatisticsScRsp) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(EnterChallengeNextPhaseCsReq, func() any { return new(proto.EnterChallengeNextPhaseCsReq) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(RestartChallengePhaseCsReq, func() any { return new(proto.RestartChallengePhaseCsReq) })
	c.regMsg(EnterChallengeNextPhaseScRsp, func() any { return new(proto.EnterChallengeNextPhaseScRsp) })
	c.regMsg(GetChallengeGroupStatisticsCsReq, func() any { return new(proto.GetChallengeGroupStatisticsCsReq) })
	c.regMsg(StartPartialChallengeCsReq, func() any { return new(proto.StartPartialChallengeCsReq) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(StartPartialChallengeScRsp, func() any { return new(proto.StartPartialChallengeScRsp) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(RestartChallengePhaseScRsp, func() any { return new(proto.RestartChallengePhaseScRsp) })
	c.regMsg(ChallengeBossPhaseSettleNotify, func() any { return new(proto.ChallengeBossPhaseSettleNotify) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoCsReq, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoCsReq) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(ChessRogueNousEnableRogueTalentCsReq, func() any { return new(proto.ChessRogueNousEnableRogueTalentCsReq) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoScRsp, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoScRsp) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(ChessRogueNousEnableRogueTalentScRsp, func() any { return new(proto.ChessRogueNousEnableRogueTalentScRsp) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(ClockParkBattleEndScNotify, func() any { return new(proto.ClockParkBattleEndScNotify) })
	c.regMsg(ClockParkStartScriptScRsp, func() any { return new(proto.ClockParkStartScriptScRsp) })
	c.regMsg(ClockParkHandleWaitOperationCsReq, func() any { return new(proto.ClockParkHandleWaitOperationCsReq) })
	c.regMsg(ClockParkGetInfoScRsp, func() any { return new(proto.ClockParkGetInfoScRsp) })
	c.regMsg(ClockParkGetInfoCsReq, func() any { return new(proto.ClockParkGetInfoCsReq) })
	c.regMsg(ClockParkQuitScriptCsReq, func() any { return new(proto.ClockParkQuitScriptCsReq) })
	c.regMsg(ClockParkHandleWaitOperationScRsp, func() any { return new(proto.ClockParkHandleWaitOperationScRsp) })
	c.regMsg(ClockParkUseBuffCsReq, func() any { return new(proto.ClockParkUseBuffCsReq) })
	c.regMsg(ClockParkGetOngoingScriptInfoCsReq, func() any { return new(proto.ClockParkGetOngoingScriptInfoCsReq) })
	c.regMsg(ClockParkStartScriptCsReq, func() any { return new(proto.ClockParkStartScriptCsReq) })
	c.regMsg(ClockParkUseBuffScRsp, func() any { return new(proto.ClockParkUseBuffScRsp) })
	c.regMsg(ClockParkGetOngoingScriptInfoScRsp, func() any { return new(proto.ClockParkGetOngoingScriptInfoScRsp) })
	c.regMsg(ClockParkUnlockTalentScRsp, func() any { return new(proto.ClockParkUnlockTalentScRsp) })
	c.regMsg(ClockParkQuitScriptScRsp, func() any { return new(proto.ClockParkQuitScriptScRsp) })
	c.regMsg(ClockParkFinishScriptScNotify, func() any { return new(proto.ClockParkFinishScriptScNotify) })
	c.regMsg(ClockParkUnlockTalentCsReq, func() any { return new(proto.ClockParkUnlockTalentCsReq) })
	c.regMsg(ContentPackageGetDataCsReq, func() any { return new(proto.ContentPackageGetDataCsReq) })
	c.regMsg(ContentPackageSyncDataScNotify, func() any { return new(proto.ContentPackageSyncDataScNotify) })
	c.regMsg(ContentPackageUnlockScRsp, func() any { return new(proto.ContentPackageUnlockScRsp) })
	c.regMsg(ContentPackageUnlockCsReq, func() any { return new(proto.ContentPackageUnlockCsReq) })
	c.regMsg(ContentPackageTransferScNotify, func() any { return new(proto.ContentPackageTransferScNotify) })
	c.regMsg(ContentPackageGetDataScRsp, func() any { return new(proto.ContentPackageGetDataScRsp) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(DrinkMakerChallengeScRsp, func() any { return new(proto.DrinkMakerChallengeScRsp) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(DrinkMakerChallengeCsReq, func() any { return new(proto.DrinkMakerChallengeCsReq) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(DrinkMakerUpdateTipsNotify, func() any { return new(proto.DrinkMakerUpdateTipsNotify) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(DrinkMakerDayEndScNotify, func() any { return new(proto.DrinkMakerDayEndScNotify) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(EvolveBuildQueryInfoScRsp, func() any { return new(proto.EvolveBuildQueryInfoScRsp) })
	c.regMsg(EvolveBuildStartStageScRsp, func() any { return new(proto.EvolveBuildStartStageScRsp) })
	c.regMsg(EvolveBuildStartStageCsReq, func() any { return new(proto.EvolveBuildStartStageCsReq) })
	c.regMsg(EvolveBuildGiveupCsReq, func() any { return new(proto.EvolveBuildGiveupCsReq) })
	c.regMsg(EvolveBuildLeaveScRsp, func() any { return new(proto.EvolveBuildLeaveScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownCsReq, func() any { return new(proto.EvolveBuildShopAbilityDownCsReq) })
	c.regMsg(EvolveBuildReRandomStageCsReq, func() any { return new(proto.EvolveBuildReRandomStageCsReq) })
	c.regMsg(EvolveBuildStartLevelCsReq, func() any { return new(proto.EvolveBuildStartLevelCsReq) })
	c.regMsg(EvolveBuildShopAbilityUpScRsp, func() any { return new(proto.EvolveBuildShopAbilityUpScRsp) })
	c.regMsg(EvolveBuildTakeExpRewardCsReq, func() any { return new(proto.EvolveBuildTakeExpRewardCsReq) })
	c.regMsg(EvolveBuildUnlockInfoNotify, func() any { return new(proto.EvolveBuildUnlockInfoNotify) })
	c.regMsg(EvolveBuildCoinNotify, func() any { return new(proto.EvolveBuildCoinNotify) })
	c.regMsg(EvolveBuildFinishScNotify, func() any { return new(proto.EvolveBuildFinishScNotify) })
	c.regMsg(EvolveBuildGiveupScRsp, func() any { return new(proto.EvolveBuildGiveupScRsp) })
	c.regMsg(EvolveBuildShopAbilityUpCsReq, func() any { return new(proto.EvolveBuildShopAbilityUpCsReq) })
	c.regMsg(EvolveBuildShopAbilityResetCsReq, func() any { return new(proto.EvolveBuildShopAbilityResetCsReq) })
	c.regMsg(EvolveBuildTakeExpRewardScRsp, func() any { return new(proto.EvolveBuildTakeExpRewardScRsp) })
	c.regMsg(EvolveBuildQueryInfoCsReq, func() any { return new(proto.EvolveBuildQueryInfoCsReq) })
	c.regMsg(EvolveBuildShopAbilityDownScRsp, func() any { return new(proto.EvolveBuildShopAbilityDownScRsp) })
	c.regMsg(EvolveBuildReRandomStageScRsp, func() any { return new(proto.EvolveBuildReRandomStageScRsp) })
	c.regMsg(EvolveBuildStartLevelScRsp, func() any { return new(proto.EvolveBuildStartLevelScRsp) })
	c.regMsg(EvolveBuildShopAbilityResetScRsp, func() any { return new(proto.EvolveBuildShopAbilityResetScRsp) })
	c.regMsg(EvolveBuildLeaveCsReq, func() any { return new(proto.EvolveBuildLeaveCsReq) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(CancelExpeditionScRsp, func() any { return new(proto.CancelExpeditionScRsp) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(FightLeaveScNotify, func() any { return new(proto.FightLeaveScNotify) })
	c.regMsg(FightHeartBeatCsReq, func() any { return new(proto.FightHeartBeatCsReq) })
	c.regMsg(FightEnterScRsp, func() any { return new(proto.FightEnterScRsp) })
	c.regMsg(FightHeartBeatScRsp, func() any { return new(proto.FightHeartBeatScRsp) })
	c.regMsg(FightKickOutScNotify, func() any { return new(proto.FightKickOutScNotify) })
	c.regMsg(FightEnterCsReq, func() any { return new(proto.FightEnterCsReq) })
	c.regMsg(FightSessionStopScNotify, func() any { return new(proto.FightSessionStopScNotify) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(EnterFightActivityStageScRsp, func() any { return new(proto.EnterFightActivityStageScRsp) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(TakeFightActivityRewardCsReq, func() any { return new(proto.TakeFightActivityRewardCsReq) })
	c.regMsg(TakeFightActivityRewardScRsp, func() any { return new(proto.TakeFightActivityRewardScRsp) })
	c.regMsg(FightMatch3DataScRsp, func() any { return new(proto.FightMatch3DataScRsp) })
	c.regMsg(FightMatch3StartCountDownScNotify, func() any { return new(proto.FightMatch3StartCountDownScNotify) })
	c.regMsg(FightMatch3ChatScNotify, func() any { return new(proto.FightMatch3ChatScNotify) })
	c.regMsg(FightMatch3SwapCsReq, func() any { return new(proto.FightMatch3SwapCsReq) })
	c.regMsg(FightMatch3ChatCsReq, func() any { return new(proto.FightMatch3ChatCsReq) })
	c.regMsg(FightMatch3TurnStartScNotify, func() any { return new(proto.FightMatch3TurnStartScNotify) })
	c.regMsg(FightMatch3SwapScRsp, func() any { return new(proto.FightMatch3SwapScRsp) })
	c.regMsg(FightMatch3ChatScRsp, func() any { return new(proto.FightMatch3ChatScRsp) })
	c.regMsg(FightMatch3ForceUpdateNotify, func() any { return new(proto.FightMatch3ForceUpdateNotify) })
	c.regMsg(FightMatch3TurnEndScNotify, func() any { return new(proto.FightMatch3TurnEndScNotify) })
	c.regMsg(FightMatch3DataCsReq, func() any { return new(proto.FightMatch3DataCsReq) })
	c.regMsg(FightMatch3OpponentDataScNotify, func() any { return new(proto.FightMatch3OpponentDataScNotify) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(GetFriendChallengeDetailCsReq, func() any { return new(proto.GetFriendChallengeDetailCsReq) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(SetFriendMarkCsReq, func() any { return new(proto.SetFriendMarkCsReq) })
	c.regMsg(TakeAssistRewardCsReq, func() any { return new(proto.TakeAssistRewardCsReq) })
	c.regMsg(GetFriendBattleRecordDetailCsReq, func() any { return new(proto.GetFriendBattleRecordDetailCsReq) })
	c.regMsg(GetFriendDevelopmentInfoScRsp, func() any { return new(proto.GetFriendDevelopmentInfoScRsp) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(NewAssistHistoryNotify, func() any { return new(proto.NewAssistHistoryNotify) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(SetFriendMarkScRsp, func() any { return new(proto.SetFriendMarkScRsp) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(TakeAssistRewardScRsp, func() any { return new(proto.TakeAssistRewardScRsp) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(CurAssistChangedNotify, func() any { return new(proto.CurAssistChangedNotify) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(GetFriendChallengeLineupScRsp, func() any { return new(proto.GetFriendChallengeLineupScRsp) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(GetFriendChallengeDetailScRsp, func() any { return new(proto.GetFriendChallengeDetailScRsp) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(GetFriendAssistListCsReq, func() any { return new(proto.GetFriendAssistListCsReq) })
	c.regMsg(GetFriendDevelopmentInfoCsReq, func() any { return new(proto.GetFriendDevelopmentInfoCsReq) })
	c.regMsg(GetFriendChallengeLineupCsReq, func() any { return new(proto.GetFriendChallengeLineupCsReq) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(GetFriendBattleRecordDetailScRsp, func() any { return new(proto.GetFriendBattleRecordDetailScRsp) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(GetFriendAssistListScRsp, func() any { return new(proto.GetFriendAssistListScRsp) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(DeleteFriendScRsp, func() any { return new(proto.DeleteFriendScRsp) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(HeliobusLineupUpdateScNotify, func() any { return new(proto.HeliobusLineupUpdateScNotify) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(HeliobusStartRaidScRsp, func() any { return new(proto.HeliobusStartRaidScRsp) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(HeliobusStartRaidCsReq, func() any { return new(proto.HeliobusStartRaidCsReq) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(RelicAvatarRecommendScRsp, func() any { return new(proto.RelicAvatarRecommendScRsp) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(GetRelicFilterPlanScRsp, func() any { return new(proto.GetRelicFilterPlanScRsp) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(DeleteRelicFilterPlanCsReq, func() any { return new(proto.DeleteRelicFilterPlanCsReq) })
	c.regMsg(AddRelicFilterPlanCsReq, func() any { return new(proto.AddRelicFilterPlanCsReq) })
	c.regMsg(AddEquipmentScNotify, func() any { return new(proto.AddEquipmentScNotify) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(DeleteRelicFilterPlanScRsp, func() any { return new(proto.DeleteRelicFilterPlanScRsp) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(MarkRelicFilterPlanScRsp, func() any { return new(proto.MarkRelicFilterPlanScRsp) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(MarkRelicFilterPlanCsReq, func() any { return new(proto.MarkRelicFilterPlanCsReq) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(AddRelicFilterPlanScRsp, func() any { return new(proto.AddRelicFilterPlanScRsp) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(RelicFilterPlanClearNameScNotify, func() any { return new(proto.RelicFilterPlanClearNameScNotify) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(ModifyRelicFilterPlanScRsp, func() any { return new(proto.ModifyRelicFilterPlanScRsp) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(ModifyRelicFilterPlanCsReq, func() any { return new(proto.ModifyRelicFilterPlanCsReq) })
	c.regMsg(GetRelicFilterPlanCsReq, func() any { return new(proto.GetRelicFilterPlanCsReq) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(RelicAvatarRecommendCsReq, func() any { return new(proto.RelicAvatarRecommendCsReq) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(GetStageLineupCsReq, func() any { return new(proto.GetStageLineupCsReq) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(LobbySyncInfoScNotify, func() any { return new(proto.LobbySyncInfoScNotify) })
	c.regMsg(LobbyGetInfoScRsp, func() any { return new(proto.LobbyGetInfoScRsp) })
	c.regMsg(LobbyQuitScRsp, func() any { return new(proto.LobbyQuitScRsp) })
	c.regMsg(LobbyCreateScRsp, func() any { return new(proto.LobbyCreateScRsp) })
	c.regMsg(LobbyQuitCsReq, func() any { return new(proto.LobbyQuitCsReq) })
	c.regMsg(LobbyInviteScRsp, func() any { return new(proto.LobbyInviteScRsp) })
	c.regMsg(LobbyGetInfoCsReq, func() any { return new(proto.LobbyGetInfoCsReq) })
	c.regMsg(LobbyInviteCsReq, func() any { return new(proto.LobbyInviteCsReq) })
	c.regMsg(LobbyModifyPlayerInfoCsReq, func() any { return new(proto.LobbyModifyPlayerInfoCsReq) })
	c.regMsg(LobbyJoinScRsp, func() any { return new(proto.LobbyJoinScRsp) })
	c.regMsg(LobbyBeginScRsp, func() any { return new(proto.LobbyBeginScRsp) })
	c.regMsg(LobbyJoinCsReq, func() any { return new(proto.LobbyJoinCsReq) })
	c.regMsg(LobbyKickOutCsReq, func() any { return new(proto.LobbyKickOutCsReq) })
	c.regMsg(LobbyBeginCsReq, func() any { return new(proto.LobbyBeginCsReq) })
	c.regMsg(LobbyKickOutScRsp, func() any { return new(proto.LobbyKickOutScRsp) })
	c.regMsg(LobbyInviteScNotify, func() any { return new(proto.LobbyInviteScNotify) })
	c.regMsg(LobbyCreateCsReq, func() any { return new(proto.LobbyCreateCsReq) })
	c.regMsg(LobbyModifyPlayerInfoScRsp, func() any { return new(proto.LobbyModifyPlayerInfoScRsp) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(GetMapRotationDataCsReq, func() any { return new(proto.GetMapRotationDataCsReq) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(RemoveRotaterCsReq, func() any { return new(proto.RemoveRotaterCsReq) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(RemoveRotaterScRsp, func() any { return new(proto.RemoveRotaterScRsp) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(UpdateMapRotationDataScNotify, func() any { return new(proto.UpdateMapRotationDataScNotify) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(UpdateRotaterScNotify, func() any { return new(proto.UpdateRotaterScNotify) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(StartMatchCsReq, func() any { return new(proto.StartMatchCsReq) })
	c.regMsg(StartMatchScRsp, func() any { return new(proto.StartMatchScRsp) })
	c.regMsg(GetCrossInfoCsReq, func() any { return new(proto.GetCrossInfoCsReq) })
	c.regMsg(GetCrossInfoScRsp, func() any { return new(proto.GetCrossInfoScRsp) })
	c.regMsg(MatchResultScNotify, func() any { return new(proto.MatchResultScNotify) })
	c.regMsg(CancelMatchScRsp, func() any { return new(proto.CancelMatchScRsp) })
	c.regMsg(CancelMatchCsReq, func() any { return new(proto.CancelMatchCsReq) })
	c.regMsg(MatchThreeSyncDataScNotify, func() any { return new(proto.MatchThreeSyncDataScNotify) })
	c.regMsg(MatchThreeGetDataScRsp, func() any { return new(proto.MatchThreeGetDataScRsp) })
	c.regMsg(MatchThreeLevelEndCsReq, func() any { return new(proto.MatchThreeLevelEndCsReq) })
	c.regMsg(MatchThreeSetBirdPosScRsp, func() any { return new(proto.MatchThreeSetBirdPosScRsp) })
	c.regMsg(MatchThreeLevelEndScRsp, func() any { return new(proto.MatchThreeLevelEndScRsp) })
	c.regMsg(MatchThreeSetBirdPosCsReq, func() any { return new(proto.MatchThreeSetBirdPosCsReq) })
	c.regMsg(MatchThreeGetDataCsReq, func() any { return new(proto.MatchThreeGetDataCsReq) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(SubmitOrigamiItemCsReq, func() any { return new(proto.SubmitOrigamiItemCsReq) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(SubmitOrigamiItemScRsp, func() any { return new(proto.SubmitOrigamiItemScRsp) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(GetGunPlayDataScRsp, func() any { return new(proto.GetGunPlayDataScRsp) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(GetGunPlayDataCsReq, func() any { return new(proto.GetGunPlayDataCsReq) })
	c.regMsg(UpdateMovieRacingDataCsReq, func() any { return new(proto.UpdateMovieRacingDataCsReq) })
	c.regMsg(UpdateGunPlayDataCsReq, func() any { return new(proto.UpdateGunPlayDataCsReq) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(UpdateMovieRacingDataScRsp, func() any { return new(proto.UpdateMovieRacingDataScRsp) })
	c.regMsg(UpdateGunPlayDataScRsp, func() any { return new(proto.UpdateGunPlayDataScRsp) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(AcceptMissionEventCsReq, func() any { return new(proto.AcceptMissionEventCsReq) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(InterruptMissionEventCsReq, func() any { return new(proto.InterruptMissionEventCsReq) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(AcceptMissionEventScRsp, func() any { return new(proto.AcceptMissionEventScRsp) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(DailyTaskDataScNotify, func() any { return new(proto.DailyTaskDataScNotify) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(MissionEventRewardScNotify, func() any { return new(proto.MissionEventRewardScNotify) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(GetMbtiReportScRsp, func() any { return new(proto.GetMbtiReportScRsp) })
	c.regMsg(MonopolySelectOptionCsReq, func() any { return new(proto.MonopolySelectOptionCsReq) })
	c.regMsg(GetMbtiReportCsReq, func() any { return new(proto.GetMbtiReportCsReq) })
	c.regMsg(MonopolyLikeScRsp, func() any { return new(proto.MonopolyLikeScRsp) })
	c.regMsg(MonopolyGameSettleScNotify, func() any { return new(proto.MonopolyGameSettleScNotify) })
	c.regMsg(DailyFirstEnterMonopolyActivityScRsp, func() any { return new(proto.DailyFirstEnterMonopolyActivityScRsp) })
	c.regMsg(MonopolyDailySettleScNotify, func() any { return new(proto.MonopolyDailySettleScNotify) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(MonopolyGuessChooseCsReq, func() any { return new(proto.MonopolyGuessChooseCsReq) })
	c.regMsg(MonopolyClickCellCsReq, func() any { return new(proto.MonopolyClickCellCsReq) })
	c.regMsg(MonopolyGameRaiseRatioCsReq, func() any { return new(proto.MonopolyGameRaiseRatioCsReq) })
	c.regMsg(MonopolyAcceptQuizCsReq, func() any { return new(proto.MonopolyAcceptQuizCsReq) })
	c.regMsg(MonopolyLikeCsReq, func() any { return new(proto.MonopolyLikeCsReq) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(GetMonopolyDailyReportScRsp, func() any { return new(proto.GetMonopolyDailyReportScRsp) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(MonopolyEventLoadUpdateScNotify, func() any { return new(proto.MonopolyEventLoadUpdateScNotify) })
	c.regMsg(MonopolyLikeScNotify, func() any { return new(proto.MonopolyLikeScNotify) })
	c.regMsg(MonopolyGiveUpCurContentScRsp, func() any { return new(proto.MonopolyGiveUpCurContentScRsp) })
	c.regMsg(MonopolyGuessDrawScNotify, func() any { return new(proto.MonopolyGuessDrawScNotify) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(MonopolyClickMbtiReportScRsp, func() any { return new(proto.MonopolyClickMbtiReportScRsp) })
	c.regMsg(MonopolyGuessBuyInformationScRsp, func() any { return new(proto.MonopolyGuessBuyInformationScRsp) })
	c.regMsg(MonopolyClickCellScRsp, func() any { return new(proto.MonopolyClickCellScRsp) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(MonopolyEventSelectFriendScRsp, func() any { return new(proto.MonopolyEventSelectFriendScRsp) })
	c.regMsg(MonopolySelectOptionScRsp, func() any { return new(proto.MonopolySelectOptionScRsp) })
	c.regMsg(MonopolyEventSelectFriendCsReq, func() any { return new(proto.MonopolyEventSelectFriendCsReq) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(MonopolyAcceptQuizScRsp, func() any { return new(proto.MonopolyAcceptQuizScRsp) })
	c.regMsg(DailyFirstEnterMonopolyActivityCsReq, func() any { return new(proto.DailyFirstEnterMonopolyActivityCsReq) })
	c.regMsg(MonopolyGuessChooseScRsp, func() any { return new(proto.MonopolyGuessChooseScRsp) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(MonopolyGiveUpCurContentCsReq, func() any { return new(proto.MonopolyGiveUpCurContentCsReq) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(MonopolyClickMbtiReportCsReq, func() any { return new(proto.MonopolyClickMbtiReportCsReq) })
	c.regMsg(MonopolyMoveCsReq, func() any { return new(proto.MonopolyMoveCsReq) })
	c.regMsg(MonopolyGetDailyInitItemScRsp, func() any { return new(proto.MonopolyGetDailyInitItemScRsp) })
	c.regMsg(MonopolyGameGachaCsReq, func() any { return new(proto.MonopolyGameGachaCsReq) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(GetMonopolyDailyReportCsReq, func() any { return new(proto.GetMonopolyDailyReportCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(MonopolySocialEventEffectScNotify, func() any { return new(proto.MonopolySocialEventEffectScNotify) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(GetMonopolyInfoScRsp, func() any { return new(proto.GetMonopolyInfoScRsp) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(MonopolyGetDailyInitItemCsReq, func() any { return new(proto.MonopolyGetDailyInitItemCsReq) })
	c.regMsg(MonopolyQuizDurationChangeScNotify, func() any { return new(proto.MonopolyQuizDurationChangeScNotify) })
	c.regMsg(MonopolyCheatDiceScRsp, func() any { return new(proto.MonopolyCheatDiceScRsp) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(MonopolyActionResultScNotify, func() any { return new(proto.MonopolyActionResultScNotify) })
	c.regMsg(MonopolyGameBingoFlipCardScRsp, func() any { return new(proto.MonopolyGameBingoFlipCardScRsp) })
	c.regMsg(MonopolyConditionUpdateScNotify, func() any { return new(proto.MonopolyConditionUpdateScNotify) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(MonopolySttUpdateScNotify, func() any { return new(proto.MonopolySttUpdateScNotify) })
	c.regMsg(MonopolyMoveScRsp, func() any { return new(proto.MonopolyMoveScRsp) })
	c.regMsg(MonopolyGameCreateScNotify, func() any { return new(proto.MonopolyGameCreateScNotify) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(GetMonopolyFriendRankingListCsReq, func() any { return new(proto.GetMonopolyFriendRankingListCsReq) })
	c.regMsg(MonopolyGuessBuyInformationCsReq, func() any { return new(proto.MonopolyGuessBuyInformationCsReq) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(MonopolyGameBingoFlipCardCsReq, func() any { return new(proto.MonopolyGameBingoFlipCardCsReq) })
	c.regMsg(GetMonopolyInfoCsReq, func() any { return new(proto.GetMonopolyInfoCsReq) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(MultiplayerFightGiveUpCsReq, func() any { return new(proto.MultiplayerFightGiveUpCsReq) })
	c.regMsg(MultiplayerFightGiveUpScRsp, func() any { return new(proto.MultiplayerFightGiveUpScRsp) })
	c.regMsg(MultiplayerGetFightGateCsReq, func() any { return new(proto.MultiplayerGetFightGateCsReq) })
	c.regMsg(MultiplayerMatch3FinishScNotify, func() any { return new(proto.MultiplayerMatch3FinishScNotify) })
	c.regMsg(MultiplayerFightGameStateScRsp, func() any { return new(proto.MultiplayerFightGameStateScRsp) })
	c.regMsg(MultiplayerFightGameStateCsReq, func() any { return new(proto.MultiplayerFightGameStateCsReq) })
	c.regMsg(MultiplayerFightGameFinishScNotify, func() any { return new(proto.MultiplayerFightGameFinishScNotify) })
	c.regMsg(MultiplayerFightGameStartScNotify, func() any { return new(proto.MultiplayerFightGameStartScNotify) })
	c.regMsg(MultiplayerGetFightGateScRsp, func() any { return new(proto.MultiplayerGetFightGateScRsp) })
	c.regMsg(GetPlayerReturnMultiDropInfoCsReq, func() any { return new(proto.GetPlayerReturnMultiDropInfoCsReq) })
	c.regMsg(GetPlayerReturnMultiDropInfoScRsp, func() any { return new(proto.GetPlayerReturnMultiDropInfoScRsp) })
	c.regMsg(MultipleDropInfoNotify, func() any { return new(proto.MultipleDropInfoNotify) })
	c.regMsg(GetMultipleDropInfoCsReq, func() any { return new(proto.GetMultipleDropInfoCsReq) })
	c.regMsg(GetMultipleDropInfoScRsp, func() any { return new(proto.GetMultipleDropInfoScRsp) })
	c.regMsg(MultipleDropInfoScNotify, func() any { return new(proto.MultipleDropInfoScNotify) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(MuseumTakeCollectRewardCsReq, func() any { return new(proto.MuseumTakeCollectRewardCsReq) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(MuseumTargetMissionFinishNotify, func() any { return new(proto.MuseumTargetMissionFinishNotify) })
	c.regMsg(MuseumTargetRewardNotify, func() any { return new(proto.MuseumTargetRewardNotify) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(MuseumTakeCollectRewardScRsp, func() any { return new(proto.MuseumTakeCollectRewardScRsp) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(TakeOfferingRewardCsReq, func() any { return new(proto.TakeOfferingRewardCsReq) })
	c.regMsg(OfferingInfoScNotify, func() any { return new(proto.OfferingInfoScNotify) })
	c.regMsg(GetOfferingInfoCsReq, func() any { return new(proto.GetOfferingInfoCsReq) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(SubmitOfferingItemCsReq, func() any { return new(proto.SubmitOfferingItemCsReq) })
	c.regMsg(TakeOfferingRewardScRsp, func() any { return new(proto.TakeOfferingRewardScRsp) })
	c.regMsg(SubmitOfferingItemScRsp, func() any { return new(proto.SubmitOfferingItemScRsp) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(AcceptedPamMissionExpireCsReq, func() any { return new(proto.AcceptedPamMissionExpireCsReq) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(SetRedPointStatusScNotify, func() any { return new(proto.SetRedPointStatusScNotify) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(UnlockAvatarPathScRsp, func() any { return new(proto.UnlockAvatarPathScRsp) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(RegionStopScNotify, func() any { return new(proto.RegionStopScNotify) })
	c.regMsg(UpdatePlayerSettingCsReq, func() any { return new(proto.UpdatePlayerSettingCsReq) })
	c.regMsg(ClientObjUploadCsReq, func() any { return new(proto.ClientObjUploadCsReq) })
	c.regMsg(AvatarPathChangedNotify, func() any { return new(proto.AvatarPathChangedNotify) })
	c.regMsg(SetAvatarPathCsReq, func() any { return new(proto.SetAvatarPathCsReq) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(GetAuthkeyCsReq, func() any { return new(proto.GetAuthkeyCsReq) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(GetMultiPathAvatarInfoScRsp, func() any { return new(proto.GetMultiPathAvatarInfoScRsp) })
	c.regMsg(ExchangeStaminaCsReq, func() any { return new(proto.ExchangeStaminaCsReq) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(UpdatePlayerSettingScRsp, func() any { return new(proto.UpdatePlayerSettingScRsp) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(ClientDownloadDataScNotify, func() any { return new(proto.ClientDownloadDataScNotify) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(UpdatePsnSettingsInfoCsReq, func() any { return new(proto.UpdatePsnSettingsInfoCsReq) })
	c.regMsg(GetAuthkeyScRsp, func() any { return new(proto.GetAuthkeyScRsp) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(GetMultiPathAvatarInfoCsReq, func() any { return new(proto.GetMultiPathAvatarInfoCsReq) })
	c.regMsg(QueryProductInfoCsReq, func() any { return new(proto.QueryProductInfoCsReq) })
	c.regMsg(AceAntiCheaterCsReq, func() any { return new(proto.AceAntiCheaterCsReq) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(ClientObjUploadScRsp, func() any { return new(proto.ClientObjUploadScRsp) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(SetAvatarPathScRsp, func() any { return new(proto.SetAvatarPathScRsp) })
	c.regMsg(GetLevelRewardScRsp, func() any { return new(proto.GetLevelRewardScRsp) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(UpdatePsnSettingsInfoScRsp, func() any { return new(proto.UpdatePsnSettingsInfoScRsp) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(UnlockAvatarPathCsReq, func() any { return new(proto.UnlockAvatarPathCsReq) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(RaidCollectionDataScRsp, func() any { return new(proto.RaidCollectionDataScRsp) })
	c.regMsg(RaidCollectionDataCsReq, func() any { return new(proto.RaidCollectionDataCsReq) })
	c.regMsg(RaidCollectionDataScNotify, func() any { return new(proto.RaidCollectionDataScNotify) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(GetAllRedDotDataCsReq, func() any { return new(proto.GetAllRedDotDataCsReq) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(SyncRogueSeasonFinishScNotify, func() any { return new(proto.SyncRogueSeasonFinishScNotify) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SyncRogueAeonScNotify, func() any { return new(proto.SyncRogueAeonScNotify) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(SyncRoguePickAvatarInfoScNotify, func() any { return new(proto.SyncRoguePickAvatarInfoScNotify) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(RogueDoGambleScRsp, func() any { return new(proto.RogueDoGambleScRsp) })
	c.regMsg(RogueWorkbenchSelectFuncCsReq, func() any { return new(proto.RogueWorkbenchSelectFuncCsReq) })
	c.regMsg(RogueWorkbenchHandleFuncScRsp, func() any { return new(proto.RogueWorkbenchHandleFuncScRsp) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(BuyRogueShopMiracleCsReq, func() any { return new(proto.BuyRogueShopMiracleCsReq) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(SetRogueExhibitionScRsp, func() any { return new(proto.SetRogueExhibitionScRsp) })
	c.regMsg(CommonRogueComponentUpdateScNotify, func() any { return new(proto.CommonRogueComponentUpdateScNotify) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(SyncRogueCommonDialogueDataScNotify, func() any { return new(proto.SyncRogueCommonDialogueDataScNotify) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(CommonRogueQueryScRsp, func() any { return new(proto.CommonRogueQueryScRsp) })
	c.regMsg(GetRogueExhibitionScRsp, func() any { return new(proto.GetRogueExhibitionScRsp) })
	c.regMsg(SetRogueCollectionScRsp, func() any { return new(proto.SetRogueCollectionScRsp) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(GetRogueShopBuffInfoCsReq, func() any { return new(proto.GetRogueShopBuffInfoCsReq) })
	c.regMsg(TakeRogueEventHandbookRewardScRsp, func() any { return new(proto.TakeRogueEventHandbookRewardScRsp) })
	c.regMsg(GetRogueExhibitionCsReq, func() any { return new(proto.GetRogueExhibitionCsReq) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(CommonRogueQueryCsReq, func() any { return new(proto.CommonRogueQueryCsReq) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoCsReq, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoCsReq) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(RogueGetGambleInfoScRsp, func() any { return new(proto.RogueGetGambleInfoScRsp) })
	c.regMsg(RogueGetGambleInfoCsReq, func() any { return new(proto.RogueGetGambleInfoCsReq) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(RogueDoGambleCsReq, func() any { return new(proto.RogueDoGambleCsReq) })
	c.regMsg(RogueWorkbenchGetInfoCsReq, func() any { return new(proto.RogueWorkbenchGetInfoCsReq) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(GetRogueCommonDialogueDataScRsp, func() any { return new(proto.GetRogueCommonDialogueDataScRsp) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(RogueWorkbenchGetInfoScRsp, func() any { return new(proto.RogueWorkbenchGetInfoScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(SyncRogueCommonDialogueOptionFinishScNotify, func() any { return new(proto.SyncRogueCommonDialogueOptionFinishScNotify) })
	c.regMsg(SelectRogueCommonDialogueOptionScRsp, func() any { return new(proto.SelectRogueCommonDialogueOptionScRsp) })
	c.regMsg(RogueWorkbenchHandleFuncCsReq, func() any { return new(proto.RogueWorkbenchHandleFuncCsReq) })
	c.regMsg(GetRogueCommonDialogueDataCsReq, func() any { return new(proto.GetRogueCommonDialogueDataCsReq) })
	c.regMsg(SelectRogueCommonDialogueOptionCsReq, func() any { return new(proto.SelectRogueCommonDialogueOptionCsReq) })
	c.regMsg(FinishRogueCommonDialogueCsReq, func() any { return new(proto.FinishRogueCommonDialogueCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(GetRogueCollectionScRsp, func() any { return new(proto.GetRogueCollectionScRsp) })
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(SetRogueCollectionCsReq, func() any { return new(proto.SetRogueCollectionCsReq) })
	c.regMsg(GetRogueCollectionCsReq, func() any { return new(proto.GetRogueCollectionCsReq) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(CommonRogueUpdateScNotify, func() any { return new(proto.CommonRogueUpdateScNotify) })
	c.regMsg(SetRogueExhibitionCsReq, func() any { return new(proto.SetRogueExhibitionCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
	c.regMsg(FinishRogueCommonDialogueScRsp, func() any { return new(proto.FinishRogueCommonDialogueScRsp) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(RogueWorkbenchSelectFuncScRsp, func() any { return new(proto.RogueWorkbenchSelectFuncScRsp) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(GetRogueShopMiracleInfoCsReq, func() any { return new(proto.GetRogueShopMiracleInfoCsReq) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(RogueModifierDelNotify, func() any { return new(proto.RogueModifierDelNotify) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(RogueModifierUpdateNotify, func() any { return new(proto.RogueModifierUpdateNotify) })
	c.regMsg(RogueModifierStageStartNotify, func() any { return new(proto.RogueModifierStageStartNotify) })
	c.regMsg(RogueModifierSelectCellScRsp, func() any { return new(proto.RogueModifierSelectCellScRsp) })
	c.regMsg(RogueTournResetPermanentTalentScRsp, func() any { return new(proto.RogueTournResetPermanentTalentScRsp) })
	c.regMsg(RogueTournGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueTournGetMiscRealTimeDataScRsp) })
	c.regMsg(RogueTournDifficultyCompNotify, func() any { return new(proto.RogueTournDifficultyCompNotify) })
	c.regMsg(RogueTournEnablePermanentTalentScRsp, func() any { return new(proto.RogueTournEnablePermanentTalentScRsp) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneScRsp) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoScRsp, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoScRsp) })
	c.regMsg(RogueTournTakeExpRewardScRsp, func() any { return new(proto.RogueTournTakeExpRewardScRsp) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoCsReq, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoCsReq) })
	c.regMsg(RogueTournGetAllArchiveScRsp, func() any { return new(proto.RogueTournGetAllArchiveScRsp) })
	c.regMsg(RogueTournReEnterRogueCocoonStageScRsp, func() any { return new(proto.RogueTournReEnterRogueCocoonStageScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryCsReq, func() any { return new(proto.RogueTournGetArchiveRepositoryCsReq) })
	c.regMsg(RogueTournStartCsReq, func() any { return new(proto.RogueTournStartCsReq) })
	c.regMsg(RogueTournGetAllArchiveCsReq, func() any { return new(proto.RogueTournGetAllArchiveCsReq) })
	c.regMsg(RogueTournEnterScRsp, func() any { return new(proto.RogueTournEnterScRsp) })
	c.regMsg(RogueTournReviveCostUpdateScNotify, func() any { return new(proto.RogueTournReviveCostUpdateScNotify) })
	c.regMsg(RogueTournLeaveCsReq, func() any { return new(proto.RogueTournLeaveCsReq) })
	c.regMsg(RogueTournResetPermanentTalentCsReq, func() any { return new(proto.RogueTournResetPermanentTalentCsReq) })
	c.regMsg(RogueTournLeaveScRsp, func() any { return new(proto.RogueTournLeaveScRsp) })
	c.regMsg(RogueTournWeekChallengeUpdateScNotify, func() any { return new(proto.RogueTournWeekChallengeUpdateScNotify) })
	c.regMsg(RogueTournHandBookNotify, func() any { return new(proto.RogueTournHandBookNotify) })
	c.regMsg(RogueTournQueryScRsp, func() any { return new(proto.RogueTournQueryScRsp) })
	c.regMsg(RogueTournTakeExpRewardCsReq, func() any { return new(proto.RogueTournTakeExpRewardCsReq) })
	c.regMsg(RogueTournReviveAvatarScRsp, func() any { return new(proto.RogueTournReviveAvatarScRsp) })
	c.regMsg(RogueTournBattleFailSettleInfoScNotify, func() any { return new(proto.RogueTournBattleFailSettleInfoScNotify) })
	c.regMsg(RogueTournEnablePermanentTalentCsReq, func() any { return new(proto.RogueTournEnablePermanentTalentCsReq) })
	c.regMsg(RogueTournClearArchiveNameScNotify, func() any { return new(proto.RogueTournClearArchiveNameScNotify) })
	c.regMsg(RogueTournEnterLayerCsReq, func() any { return new(proto.RogueTournEnterLayerCsReq) })
	c.regMsg(RogueTournConfirmSettleCsReq, func() any { return new(proto.RogueTournConfirmSettleCsReq) })
	c.regMsg(RogueTournGetSettleInfoCsReq, func() any { return new(proto.RogueTournGetSettleInfoCsReq) })
	c.regMsg(RogueTournEnterRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournEnterRogueCocoonSceneScRsp) })
	c.regMsg(RogueTournReEnterRogueCocoonStageCsReq, func() any { return new(proto.RogueTournReEnterRogueCocoonStageCsReq) })
	c.regMsg(RogueTournEnterRoomCsReq, func() any { return new(proto.RogueTournEnterRoomCsReq) })
	c.regMsg(RogueTournStartScRsp, func() any { return new(proto.RogueTournStartScRsp) })
	c.regMsg(RogueTournGetPermanentTalentInfoCsReq, func() any { return new(proto.RogueTournGetPermanentTalentInfoCsReq) })
	c.regMsg(RogueTournGetArchiveRepositoryScRsp, func() any { return new(proto.RogueTournGetArchiveRepositoryScRsp) })
	c.regMsg(RogueTournGetSettleInfoScRsp, func() any { return new(proto.RogueTournGetSettleInfoScRsp) })
	c.regMsg(RogueTournConfirmSettleScRsp, func() any { return new(proto.RogueTournConfirmSettleScRsp) })
	c.regMsg(RogueTournQueryCsReq, func() any { return new(proto.RogueTournQueryCsReq) })
	c.regMsg(RogueTournReviveAvatarCsReq, func() any { return new(proto.RogueTournReviveAvatarCsReq) })
	c.regMsg(RogueTournAreaUpdateScNotify, func() any { return new(proto.RogueTournAreaUpdateScNotify) })
	c.regMsg(RogueTournExpNotify, func() any { return new(proto.RogueTournExpNotify) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneCsReq) })
	c.regMsg(RogueTournGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueTournGetMiscRealTimeDataCsReq) })
	c.regMsg(RogueTournGetPermanentTalentInfoScRsp, func() any { return new(proto.RogueTournGetPermanentTalentInfoScRsp) })
	c.regMsg(RogueTournDeleteArchiveCsReq, func() any { return new(proto.RogueTournDeleteArchiveCsReq) })
	c.regMsg(RogueTournLevelInfoUpdateScNotify, func() any { return new(proto.RogueTournLevelInfoUpdateScNotify) })
	c.regMsg(RogueTournEnterCsReq, func() any { return new(proto.RogueTournEnterCsReq) })
	c.regMsg(RogueTournEnterLayerScRsp, func() any { return new(proto.RogueTournEnterLayerScRsp) })
	c.regMsg(RogueTournEnterRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournEnterRogueCocoonSceneCsReq) })
	c.regMsg(RogueTournEnterRoomScRsp, func() any { return new(proto.RogueTournEnterRoomScRsp) })
	c.regMsg(RogueTournRenameArchiveScRsp, func() any { return new(proto.RogueTournRenameArchiveScRsp) })
	c.regMsg(RogueTournSettleScRsp, func() any { return new(proto.RogueTournSettleScRsp) })
	c.regMsg(RogueTournSettleCsReq, func() any { return new(proto.RogueTournSettleCsReq) })
	c.regMsg(RogueTournRenameArchiveCsReq, func() any { return new(proto.RogueTournRenameArchiveCsReq) })
	c.regMsg(RogueTournDeleteArchiveScRsp, func() any { return new(proto.RogueTournDeleteArchiveScRsp) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(TakeRollShopRewardCsReq, func() any { return new(proto.TakeRollShopRewardCsReq) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(StartTimedFarmElementScRsp, func() any { return new(proto.StartTimedFarmElementScRsp) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(SpringRecoverCsReq, func() any { return new(proto.SpringRecoverCsReq) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(SpringRefreshScRsp, func() any { return new(proto.SpringRefreshScRsp) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(SpringRecoverSingleAvatarScRsp, func() any { return new(proto.SpringRecoverSingleAvatarScRsp) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(StartTimedCocoonStageScRsp, func() any { return new(proto.StartTimedCocoonStageScRsp) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(GetSpringRecoverDataScRsp, func() any { return new(proto.GetSpringRecoverDataScRsp) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(SceneReviveAfterRebattleCsReq, func() any { return new(proto.SceneReviveAfterRebattleCsReq) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(SetSpringRecoverConfigScRsp, func() any { return new(proto.SetSpringRecoverConfigScRsp) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(LastSpringRefreshTimeNotify, func() any { return new(proto.LastSpringRefreshTimeNotify) })
	c.regMsg(StartTimedFarmElementCsReq, func() any { return new(proto.StartTimedFarmElementCsReq) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(HealPoolInfoNotify, func() any { return new(proto.HealPoolInfoNotify) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(SpringRecoverScRsp, func() any { return new(proto.SpringRecoverScRsp) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(SceneReviveAfterRebattleScRsp, func() any { return new(proto.SceneReviveAfterRebattleScRsp) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(GetSpringRecoverDataCsReq, func() any { return new(proto.GetSpringRecoverDataCsReq) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(SetSpringRecoverConfigCsReq, func() any { return new(proto.SetSpringRecoverConfigCsReq) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(SpringRecoverSingleAvatarCsReq, func() any { return new(proto.SpringRecoverSingleAvatarCsReq) })
	c.regMsg(SpringRefreshCsReq, func() any { return new(proto.SpringRefreshCsReq) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(StartTimedCocoonStageCsReq, func() any { return new(proto.StartTimedCocoonStageCsReq) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(SpaceZooTakeScRsp, func() any { return new(proto.SpaceZooTakeScRsp) })
	c.regMsg(SpaceZooExchangeItemScRsp, func() any { return new(proto.SpaceZooExchangeItemScRsp) })
	c.regMsg(SpaceZooExchangeItemCsReq, func() any { return new(proto.SpaceZooExchangeItemCsReq) })
	c.regMsg(SpaceZooTakeCsReq, func() any { return new(proto.SpaceZooTakeCsReq) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(GetStarFightDataCsReq, func() any { return new(proto.GetStarFightDataCsReq) })
	c.regMsg(StarFightDataChangeNotify, func() any { return new(proto.StarFightDataChangeNotify) })
	c.regMsg(StartStarFightLevelScRsp, func() any { return new(proto.StartStarFightLevelScRsp) })
	c.regMsg(GetStarFightDataScRsp, func() any { return new(proto.GetStarFightDataScRsp) })
	c.regMsg(StartStarFightLevelCsReq, func() any { return new(proto.StartStarFightLevelCsReq) })
	c.regMsg(StoryLineTrialAvatarChangeScNotify, func() any { return new(proto.StoryLineTrialAvatarChangeScNotify) })
	c.regMsg(ChangeStoryLineFinishScNotify, func() any { return new(proto.ChangeStoryLineFinishScNotify) })
	c.regMsg(StoryLineInfoScNotify, func() any { return new(proto.StoryLineInfoScNotify) })
	c.regMsg(GetStoryLineInfoScRsp, func() any { return new(proto.GetStoryLineInfoScRsp) })
	c.regMsg(GetStoryLineInfoCsReq, func() any { return new(proto.GetStoryLineInfoCsReq) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(StrongChallengeActivityBattleEndScNotify, func() any { return new(proto.StrongChallengeActivityBattleEndScNotify) })
	c.regMsg(SummonActivityBattleEndScNotify, func() any { return new(proto.SummonActivityBattleEndScNotify) })
	c.regMsg(EnterSummonActivityStageScRsp, func() any { return new(proto.EnterSummonActivityStageScRsp) })
	c.regMsg(EnterSummonActivityStageCsReq, func() any { return new(proto.EnterSummonActivityStageCsReq) })
	c.regMsg(GetSummonActivityDataCsReq, func() any { return new(proto.GetSummonActivityDataCsReq) })
	c.regMsg(GetSummonActivityDataScRsp, func() any { return new(proto.GetSummonActivityDataScRsp) })
	c.regMsg(SwordTrainingRestoreGameCsReq, func() any { return new(proto.SwordTrainingRestoreGameCsReq) })
	c.regMsg(GetSwordTrainingDataScRsp, func() any { return new(proto.GetSwordTrainingDataScRsp) })
	c.regMsg(SwordTrainingGiveUpGameCsReq, func() any { return new(proto.SwordTrainingGiveUpGameCsReq) })
	c.regMsg(SwordTrainingLearnSkillScRsp, func() any { return new(proto.SwordTrainingLearnSkillScRsp) })
	c.regMsg(SwordTrainingSetSkillTraceScRsp, func() any { return new(proto.SwordTrainingSetSkillTraceScRsp) })
	c.regMsg(SwordTrainingExamResultConfirmCsReq, func() any { return new(proto.SwordTrainingExamResultConfirmCsReq) })
	c.regMsg(SwordTrainingDialogueSelectOptionScRsp, func() any { return new(proto.SwordTrainingDialogueSelectOptionScRsp) })
	c.regMsg(SwordTrainingExamResultConfirmScRsp, func() any { return new(proto.SwordTrainingExamResultConfirmScRsp) })
	c.regMsg(SwordTrainingFinishEndingHintScRsp, func() any { return new(proto.SwordTrainingFinishEndingHintScRsp) })
	c.regMsg(SwordTrainingTakeEndingRewardCsReq, func() any { return new(proto.SwordTrainingTakeEndingRewardCsReq) })
	c.regMsg(SwordTrainingResumeGameCsReq, func() any { return new(proto.SwordTrainingResumeGameCsReq) })
	c.regMsg(SwordTrainingStoryBattleScRsp, func() any { return new(proto.SwordTrainingStoryBattleScRsp) })
	c.regMsg(SwordTrainingFinishEndingHintCsReq, func() any { return new(proto.SwordTrainingFinishEndingHintCsReq) })
	c.regMsg(SwordTrainingGameSettleScNotify, func() any { return new(proto.SwordTrainingGameSettleScNotify) })
	c.regMsg(SwordTrainingResumeGameScRsp, func() any { return new(proto.SwordTrainingResumeGameScRsp) })
	c.regMsg(EnterSwordTrainingExamScRsp, func() any { return new(proto.EnterSwordTrainingExamScRsp) })
	c.regMsg(SwordTrainingSelectEndingCsReq, func() any { return new(proto.SwordTrainingSelectEndingCsReq) })
	c.regMsg(SwordTrainingDailyPhaseConfirmCsReq, func() any { return new(proto.SwordTrainingDailyPhaseConfirmCsReq) })
	c.regMsg(SwordTrainingSetSkillTraceCsReq, func() any { return new(proto.SwordTrainingSetSkillTraceCsReq) })
	c.regMsg(SwordTrainingRestoreGameScRsp, func() any { return new(proto.SwordTrainingRestoreGameScRsp) })
	c.regMsg(SwordTrainingTurnActionScRsp, func() any { return new(proto.SwordTrainingTurnActionScRsp) })
	c.regMsg(SwordTrainingStartGameCsReq, func() any { return new(proto.SwordTrainingStartGameCsReq) })
	c.regMsg(SwordTrainingUpdateRankScRsp, func() any { return new(proto.SwordTrainingUpdateRankScRsp) })
	c.regMsg(SwordTrainingActionTurnSettleScNotify, func() any { return new(proto.SwordTrainingActionTurnSettleScNotify) })
	c.regMsg(SwordTrainingTurnActionCsReq, func() any { return new(proto.SwordTrainingTurnActionCsReq) })
	c.regMsg(SwordTrainingGetSkillInfoScRsp, func() any { return new(proto.SwordTrainingGetSkillInfoScRsp) })
	c.regMsg(SwordTrainingStoryConfirmScRsp, func() any { return new(proto.SwordTrainingStoryConfirmScRsp) })
	c.regMsg(SwordTrainingDialogueSelectOptionCsReq, func() any { return new(proto.SwordTrainingDialogueSelectOptionCsReq) })
	c.regMsg(SwordTrainingSelectEndingScRsp, func() any { return new(proto.SwordTrainingSelectEndingScRsp) })
	c.regMsg(SwordTrainingStoryBattleCsReq, func() any { return new(proto.SwordTrainingStoryBattleCsReq) })
	c.regMsg(SwordTrainingUnlockSyncScNotify, func() any { return new(proto.SwordTrainingUnlockSyncScNotify) })
	c.regMsg(SwordTrainingRefreshPartnerAbilityCsReq, func() any { return new(proto.SwordTrainingRefreshPartnerAbilityCsReq) })
	c.regMsg(SwordTrainingDailyPhaseConfirmScRsp, func() any { return new(proto.SwordTrainingDailyPhaseConfirmScRsp) })
	c.regMsg(EnterSwordTrainingExamCsReq, func() any { return new(proto.EnterSwordTrainingExamCsReq) })
	c.regMsg(SwordTrainingGiveUpGameScRsp, func() any { return new(proto.SwordTrainingGiveUpGameScRsp) })
	c.regMsg(SwordTrainingGetSkillInfoCsReq, func() any { return new(proto.SwordTrainingGetSkillInfoCsReq) })
	c.regMsg(SwordTrainingLearnSkillCsReq, func() any { return new(proto.SwordTrainingLearnSkillCsReq) })
	c.regMsg(SwordTrainingGameSyncChangeScNotify, func() any { return new(proto.SwordTrainingGameSyncChangeScNotify) })
	c.regMsg(SwordTrainingUpdateRankCsReq, func() any { return new(proto.SwordTrainingUpdateRankCsReq) })
	c.regMsg(SwordTrainingRefreshPartnerAbilityScRsp, func() any { return new(proto.SwordTrainingRefreshPartnerAbilityScRsp) })
	c.regMsg(SwordTrainingStartGameScRsp, func() any { return new(proto.SwordTrainingStartGameScRsp) })
	c.regMsg(SwordTrainingStoryConfirmCsReq, func() any { return new(proto.SwordTrainingStoryConfirmCsReq) })
	c.regMsg(GetSwordTrainingDataCsReq, func() any { return new(proto.GetSwordTrainingDataCsReq) })
	c.regMsg(SwordTrainingTakeEndingRewardScRsp, func() any { return new(proto.SwordTrainingTakeEndingRewardScRsp) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	c.regMsg(GetTrackPhotoActivityDataCsReq, func() any { return new(proto.GetTrackPhotoActivityDataCsReq) })
	c.regMsg(SettleTrackPhotoStageCsReq, func() any { return new(proto.SettleTrackPhotoStageCsReq) })
	c.regMsg(GetTrackPhotoActivityDataScRsp, func() any { return new(proto.GetTrackPhotoActivityDataScRsp) })
	c.regMsg(QuitTrackPhotoStageCsReq, func() any { return new(proto.QuitTrackPhotoStageCsReq) })
	c.regMsg(SettleTrackPhotoStageScRsp, func() any { return new(proto.SettleTrackPhotoStageScRsp) })
	c.regMsg(QuitTrackPhotoStageScRsp, func() any { return new(proto.QuitTrackPhotoStageScRsp) })
	c.regMsg(StartTrackPhotoStageCsReq, func() any { return new(proto.StartTrackPhotoStageCsReq) })
	c.regMsg(StartTrackPhotoStageScRsp, func() any { return new(proto.StartTrackPhotoStageScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(TrainVisitorBehaviorFinishCsReq, func() any { return new(proto.TrainVisitorBehaviorFinishCsReq) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(GetTrainVisitorBehaviorCsReq, func() any { return new(proto.GetTrainVisitorBehaviorCsReq) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(WolfBroGameActivateBulletScRsp, func() any { return new(proto.WolfBroGameActivateBulletScRsp) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(WolfBroGameExplodeMonsterScRsp, func() any { return new(proto.WolfBroGameExplodeMonsterScRsp) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(WolfBroGameExplodeMonsterCsReq, func() any { return new(proto.WolfBroGameExplodeMonsterCsReq) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	// seever
	c.regMsg(GmGive, func() any { return new(spb.GmGive) })
	c.regMsg(GmWorldLevel, func() any { return new(spb.GmWorldLevel) })
	c.regMsg(DelItem, func() any { return new(spb.DelItem) })
	c.regMsg(MaxCurAvatar, func() any { return new(spb.MaxCurAvatar) })
	c.regMsg(ServiceConnectionReq, func() any { return new(spb.ServiceConnectionReq) })
	c.regMsg(ServiceConnectionRsp, func() any { return new(spb.ServiceConnectionRsp) })
	c.regMsg(GateLoginGameRsp, func() any { return new(spb.GateLoginGameRsp) })
	c.regMsg(GateLoginGameReq, func() any { return new(spb.GateLoginGameReq) })
	c.regMsg(GateToGameMsgNotify, func() any { return new(spb.GateToGameMsgNotify) })
	c.regMsg(GameToGateMsgNotify, func() any { return new(spb.GameToGateMsgNotify) })
	c.regMsg(GetAllServiceGateReq, func() any { return new(spb.GetAllServiceGateReq) })
	c.regMsg(GetAllServiceGateRsp, func() any { return new(spb.GetAllServiceGateRsp) })
	c.regMsg(MultiToNodePingReq, func() any { return new(spb.MultiToNodePingReq) })
	c.regMsg(MultiToNodePingRsp, func() any { return new(spb.MultiToNodePingRsp) })
	c.regMsg(MuipToNodePingReq, func() any { return new(spb.MuipToNodePingReq) })
	c.regMsg(MuipToNodePingRsp, func() any { return new(spb.MuipToNodePingRsp) })
	c.regMsg(GateGamePingReq, func() any { return new(spb.GateGamePingReq) })
	c.regMsg(GateGamePingRsp, func() any { return new(spb.GateGamePingRsp) })
	c.regMsg(GateGamePlayerLoginReq, func() any { return new(spb.GateGamePlayerLoginReq) })
	c.regMsg(GateGamePlayerLoginRsp, func() any { return new(spb.GateGamePlayerLoginRsp) })
	c.regMsg(GetToGamePlayerLogoutReq, func() any { return new(spb.GetToGamePlayerLogoutReq) })
	c.regMsg(GetToGamePlayerLogoutRsp, func() any { return new(spb.GetToGamePlayerLogoutRsp) })
	c.regMsg(GateLoginMultiReq, func() any { return new(spb.GateLoginMultiReq) })
	c.regMsg(GateLoginMultiRsp, func() any { return new(spb.GateLoginMultiRsp) })
	c.regMsg(GameToGatePlayerLogoutNotify, func() any { return new(spb.GameToGatePlayerLogoutNotify) })
	c.regMsg(GateToGamePlayerLogoutNotify, func() any { return new(spb.GateToGamePlayerLogoutNotify) })
	c.regMsg(GmMission, func() any { return new(spb.GmMission) })
	c.regMsg(PlayerMsgGateToNodeNotify, func() any { return new(spb.PlayerMsgGateToNodeNotify) })
	c.regMsg(GameToNodePingReq, func() any { return new(spb.GameToNodePingReq) })
	c.regMsg(GameToNodePingRsp, func() any { return new(spb.GameToNodePingRsp) })
	c.regMsg(GateToNodePingReq, func() any { return new(spb.GateToNodePingReq) })
	c.regMsg(GateToNodePingRsp, func() any { return new(spb.GateToNodePingRsp) })
}
