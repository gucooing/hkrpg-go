package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

const (
	SwordTrainingGameSyncChangeScNotify                = 7465
	DeleteRelicFilterPlanScRsp                         = 532
	SwordTrainingDialogueSelectOptionCsReq             = 7473
	SetAssistCsReq                                     = 2910
	GetSaveLogisticsMapScRsp                           = 4710
	ClockParkBattleEndScNotify                         = 7233
	RaidCollectionEnterNextRaidCsReq                   = 6950
	StartRogueCsReq                                    = 1803
	EnterRogueCsReq                                    = 1839
	TravelBrochureGetDataScRsp                         = 6420
	GetTutorialCsReq                                   = 1659
	MuseumDispatchFinishedScNotify                     = 4373
	HeliobusSnsPostCsReq                               = 5839
	StartRaidCsReq                                     = 2259
	RogueTournLeaveRogueCocoonSceneCsReq               = 6069
	AddAvatarScNotify                                  = 330
	TravelBrochureApplyPasterListScRsp                 = 6433
	HandleFriendCsReq                                  = 2916
	SellItemScRsp                                      = 525
	TakeMaterialSubmitActivityRewardScRsp              = 2633
	SetGameplayBirthdayScRsp                           = 72
	TakeLoginActivityRewardCsReq                       = 2603
	RogueMagicSetAutoDressInMagicUnitCsReq             = 7777
	GetMultipleDropInfoCsReq                           = 4659
	PlayerHeartBeatScRsp                               = 31
	TakeQuestRewardScRsp                               = 946
	StartTrackPhotoStageCsReq                          = 7551
	FightTreasureDungeonMonsterCsReq                   = 4479
	TakeBpRewardCsReq                                  = 3046
	HeliobusUnlockSkillScNotify                        = 5848
	ExtraLineupDestroyNotify                           = 783
	MultipleDropInfoScNotify                           = 4603
	MusicRhythmSaveSongConfigDataScRsp                 = 7578
	GetFriendLoginInfoCsReq                            = 2914
	MusicRhythmUnlockTrackScNotify                     = 7579
	MonthCardRewardNotify                              = 60
	ChessRogueUpdateDiceInfoScNotify                   = 5438
	ContentPackageUnlockScRsp                          = 7542
	RankUpEquipmentCsReq                               = 580
	ExpUpRelicScRsp                                    = 590
	ClientObjUploadScRsp                               = 67
	TrainPartyMoveScNotify                             = 8039
	SaveLogisticsCsReq                                 = 4799
	GetGachaCeilingCsReq                               = 1939
	SetGenderScRsp                                     = 4
	FinishRogueCommonDialogueScRsp                     = 5627
	MarkAvatarScRsp                                    = 393
	GetChallengeCsReq                                  = 1759
	RestoreWolfBroGameArchiveScRsp                     = 6553
	GetMuseumInfoCsReq                                 = 4359
	BuyNpcStuffCsReq                                   = 4303
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6003
	SetAetherDivideLineUpScRsp                         = 4875
	QuitBattleScRsp                                    = 146
	RogueGetGambleInfoScRsp                            = 5601
	RogueTournDeleteArchiveCsReq                       = 6023
	EnterSectionScRsp                                  = 1493
	ChessRogueNousEnableRogueTalentScRsp               = 5510
	SetGameplayBirthdayCsReq                           = 78
	PlayerReturnForceFinishScNotify                    = 4547
	TrainPartyUpdatePosEnvCsReq                        = 8010
	StartBoxingClubBattleCsReq                         = 4239
	SubMissionRewardScNotify                           = 1299
	PlayerHeartBeatCsReq                               = 96
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5629
	RogueTournLevelInfoUpdateScNotify                  = 6083
	GetMultiPathAvatarInfoScRsp                        = 87
	FinishItemIdScRsp                                  = 2753
	SelectRogueCommonDialogueOptionCsReq               = 5682
	TakeMultipleActivityExpeditionRewardScRsp          = 2529
	SwitchLineupIndexCsReq                             = 790
	TakeMultipleActivityExpeditionRewardCsReq          = 2525
	GetSocialEventServerCacheCsReq                     = 7084
	GetChallengeRecommendLineupListScRsp               = 2437
	FinishChessRogueSubStoryCsReq                      = 5550
	SetCurInteractEntityScRsp                          = 1470
	StartAetherDivideSceneBattleScRsp                  = 4853
	SyncHandleFriendScNotify                           = 2974
	GetFriendAssistListCsReq                           = 2956
	GetAllServerPrefsDataScRsp                         = 6120
	GetArchiveDataCsReq                                = 2359
	FightMatch3OpponentDataScNotify                    = 30137
	GetPunkLordBattleRecordScRsp                       = 3236
	MonopolyLikeScRsp                                  = 7060
	RetcodeNotify                                      = 32
	ClockParkStartScriptCsReq                          = 7240
	SetFriendRemarkNameScRsp                           = 2933
	RemoveRotaterCsReq                                 = 6819
	TakePromotionRewardScRsp                           = 329
	ChessRogueEnterNextLayerCsReq                      = 5568
	MonopolyRollDiceScRsp                              = 7034
	AetherDivideLineupScNotify                         = 4870
	FightMatch3ChatScNotify                            = 30147
	GetMapRotationDataCsReq                            = 6847
	GetTutorialScRsp                                   = 1620
	FightMatch3SwapCsReq                               = 30153
	TravelBrochureApplyPasterListCsReq                 = 6429
	EvolveBuildTakeExpRewardCsReq                      = 7105
	SetMultipleAvatarPathsCsReq                        = 100
	RogueWorkbenchGetInfoScRsp                         = 5635
	SyncRogueRewardInfoScNotify                        = 1889
	GetFeverTimeActivityDataScRsp                      = 7151
	SyncApplyFriendScNotify                            = 2980
	MatchThreeGetDataScRsp                             = 7416
	MissionRewardScNotify                              = 1239
	SyncRoguePickAvatarInfoScNotify                    = 1828
	RogueTournGetPermanentTalentInfoScRsp              = 6085
	AlleyGuaranteedFundsScRsp                          = 4765
	TakeOffEquipmentCsReq                              = 347
	SpaceZooExchangeItemCsReq                          = 6774
	GetFirstTalkByPerformanceNpcCsReq                  = 2147
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3780
	SetAssistScRsp                                     = 2970
	MissionGroupWarnScNotify                           = 1274
	GetMonopolyFriendRankingListScRsp                  = 7024
	GetRogueBuffEnhanceInfoScRsp                       = 1873
	SwordTrainingUnlockSyncScNotify                    = 7452
	TrainPartyUnlockBuildAreaCsReq                     = 8083
	FinishTutorialCsReq                                = 1680
	ChessRogueNousEditDiceCsReq                        = 5453
	MonopolyGuessBuyInformationCsReq                   = 7063
	FinishTutorialGuideScRsp                           = 1674
	RogueTournConfirmSettleCsReq                       = 6082
	SetFriendMarkScRsp                                 = 2966
	FeatureSwitchClosedScNotify                        = 82
	OfferingInfoScNotify                               = 6924
	SummonPunkLordMonsterCsReq                         = 3234
	DailyFirstMeetPamScRsp                             = 3446
	AlleyShipUnlockScNotify                            = 4783
	GetDrinkMakerDataCsReq                             = 6985
	TravelBrochureApplyPasterScRsp                     = 6437
	EnterStrongChallengeActivityStageCsReq             = 6603
	TravelBrochureSetPageDescStatusScRsp               = 6419
	UpdatePsnSettingsInfoCsReq                         = 52
	StarFightDataChangeNotify                          = 7161
	ComposeSelectedRelicCsReq                          = 573
	RogueMagicAreaUpdateScNotify                       = 7775
	RogueTournExpNotify                                = 6032
	GetFirstTalkNpcScRsp                               = 2153
	SyncChessRogueNousMainStoryScNotify                = 5521
	GetPlayerDetailInfoCsReq                           = 2903
	MonopolySttUpdateScNotify                          = 7052
	MonopolyCellUpdateNotify                           = 7046
	StartAetherDivideChallengeBattleCsReq              = 4834
	ContentPackageUnlockCsReq                          = 7524
	AvatarExpUpCsReq                                   = 303
	SetAvatarPathScRsp                                 = 41
	TakeOffAvatarSkinCsReq                             = 373
	GetGachaInfoCsReq                                  = 1959
	LockRelicScRsp                                     = 519
	FightFestScoreUpdateNotify                         = 7292
	GetCrossInfoCsReq                                  = 7337
	ClientDownloadDataScNotify                         = 15
	GetMaterialSubmitActivityDataScRsp                 = 2619
	GetServerPrefsDataScRsp                            = 6146
	RogueArcadeRestartCsReq                            = 7692
	ChessRogueSelectCellScRsp                          = 5517
	BuyRogueShopMiracleCsReq                           = 5637
	TravelBrochurePageResetScRsp                       = 6425
	LogisticsGameScRsp                                 = 4746
	TrialBackGroundMusicCsReq                          = 3134
	GetAllLineupDataScRsp                              = 729
	LobbyJoinCsReq                                     = 7362
	UpdateGunPlayDataScRsp                             = 4193
	QueryProductInfoScRsp                              = 23
	ChessRogueCheatRollScRsp                           = 5573
	HeartDialTraceScriptCsReq                          = 6316
	SetCurInteractEntityCsReq                          = 1410
	TakePrestigeRewardScRsp                            = 4774
	RogueMagicQueryCsReq                               = 7733
	UnlockTeleportNotify                               = 1489
	TrainPartyLeaveScRsp                               = 8033
	TravelBrochureRemovePasterScRsp                    = 6416
	SyncTaskCsReq                                      = 1253
	GetRogueExhibitionCsReq                            = 5692
	BuyGoodsScRsp                                      = 1546
	ChessRogueUpdateActionPointScNotify                = 5481
	RaidCollectionDataScRsp                            = 6948
	TrialActivityDataChangeScNotify                    = 2624
	UpdateGroupPropertyScRsp                           = 1485
	MuseumRandomEventQueryScRsp                        = 4329
	AcceptMainMissionCsReq                             = 1210
	LeaveRaidCsReq                                     = 2203
	StartTimedCocoonStageCsReq                         = 1449
	ChessRogueQueryCsReq                               = 5587
	SetPlayerInfoCsReq                                 = 88
	TravelBrochureRemovePasterCsReq                    = 6480
	LeaveTrialActivityCsReq                            = 2682
	GetFriendRecommendListInfoScRsp                    = 2925
	EvolveBuildShopAbilityDownScRsp                    = 7120
	GetMonopolyFriendRankingListCsReq                  = 7072
	StartAetherDivideSceneBattleCsReq                  = 4839
	TakeChallengeRaidRewardCsReq                       = 2237
	SetSpringRecoverConfigCsReq                        = 1456
	MusicRhythmUnlockSongSfxScNotify                   = 7571
	TakeChallengeRaidRewardScRsp                       = 2280
	FightKickOutScNotify                               = 30046
	VirtualLineupDestroyNotify                         = 733
	SceneCastSkillCsReq                                = 1439
	ChessRogueLeaveCsReq                               = 5401
	PlayerReturnStartScNotify                          = 4559
	RogueTournEnterLayerCsReq                          = 6049
	GetGunPlayDataCsReq                                = 4183
	GetFriendListInfoScRsp                             = 2920
	ChangeLineupLeaderScRsp                            = 748
	TakeRollShopRewardCsReq                            = 6912
	AcceptExpeditionScRsp                              = 2546
	SetStuffToAreaScRsp                                = 4353
	StartAetherDivideStageBattleCsReq                  = 4829
	LobbyCreateCsReq                                   = 7365
	MonopolyGameGachaScRsp                             = 7065
	BuyRogueShopMiracleScRsp                           = 5680
	RogueTournGetArchiveRepositoryCsReq                = 6079
	GetHeartDialInfoCsReq                              = 6359
	GetSecretKeyInfoCsReq                              = 27
	TakeOffRelicScRsp                                  = 361
	ClockParkGetInfoCsReq                              = 7215
	GetVideoVersionKeyScRsp                            = 35
	SwordTrainingDailyPhaseConfirmScRsp                = 7490
	SpringRecoverSingleAvatarCsReq                     = 1460
	MonopolyActionResultScNotify                       = 7003
	GetRogueShopMiracleInfoCsReq                       = 5646
	GetExpeditionDataScRsp                             = 2520
	MonopolyClickMbtiReportCsReq                       = 7008
	RogueTournTakeExpRewardCsReq                       = 6063
	HeliobusSnsUpdateScNotify                          = 5847
	StartCocoonStageCsReq                              = 1413
	GetSwordTrainingDataCsReq                          = 7466
	TakeTrialActivityRewardCsReq                       = 2632
	GetLoginChatInfoScRsp                              = 3990
	UnlockChatBubbleScNotify                           = 5139
	GetAetherDivideInfoScRsp                           = 4874
	TravelBrochureGetDataCsReq                         = 6459
	SyncRogueStatusScNotify                            = 1821
	HeliobusSelectSkillCsReq                           = 5890
	HeliobusSnsCommentCsReq                            = 5880
	QuitLineupScRsp                                    = 737
	AlleyTakeEventRewardCsReq                          = 4726
	CancelExpeditionCsReq                              = 2539
	SwordTrainingRestoreGameScRsp                      = 7451
	ExchangeRogueRewardKeyScRsp                        = 1831
	UpdateServerPrefsDataScRsp                         = 6153
	MultiplayerMatch3FinishScNotify                    = 1080
	GetStarFightDataScRsp                              = 7165
	BattleCollegeDataChangeScNotify                    = 5703
	EvolveBuildShopAbilityResetCsReq                   = 7118
	StartAetherDivideChallengeBattleScRsp              = 4837
	ChooseBoxingClubStageOptionalBuffScRsp             = 4290
	SpaceZooDeleteCatCsReq                             = 6780
	QuitTrackPhotoStageScRsp                           = 7557
	MarkRelicFilterPlanCsReq                           = 596
	RegionStopScNotify                                 = 19
	GetChapterCsReq                                    = 439
	BattleLogReportCsReq                               = 116
	MonopolyConditionUpdateScNotify                    = 7011
	UpdateFeatureSwitchScNotify                        = 45
	GetAssistListCsReq                                 = 2964
	RogueWorkbenchSelectFuncScRsp                      = 5644
	TakeActivityExpeditionRewardCsReq                  = 2575
	UnlockAvatarPathCsReq                              = 11
	GetFriendRecommendListInfoCsReq                    = 2961
	RecallPetScRsp                                     = 7637
	ChessRogueUpdateReviveInfoScNotify                 = 5520
	MonopolyRollRandomCsReq                            = 7048
	RogueMagicGetMiscRealTimeDataScRsp                 = 7788
	ChessRogueSkipTeachingLevelCsReq                   = 5459
	EnterRogueScRsp                                    = 1853
	GetShareDataScRsp                                  = 4146
	SelectRogueCommonDialogueOptionScRsp               = 5662
	RogueTournEnablePermanentTalentCsReq               = 6043
	MonopolyGetRaffleTicketScRsp                       = 7035
	EnableRogueTalentScRsp                             = 1895
	GetRogueCommonDialogueDataCsReq                    = 5631
	GetSceneMapInfoScRsp                               = 1481
	ComposeSelectedRelicScRsp                          = 583
	MakeMissionDrinkScRsp                              = 6989
	ClockParkHandleWaitOperationScRsp                  = 7228
	TakeFightActivityRewardScRsp                       = 3634
	GetSaveRaidCsReq                                   = 2248
	BoxingClubRewardScNotify                           = 4280
	MatchThreeSyncDataScNotify                         = 7442
	DailyTaskDataScNotify                              = 1237
	UpdateRedDotDataScRsp                              = 5946
	AlleyPlacingGameCsReq                              = 4730
	DeleteSummonUnitScRsp                              = 1405
	EnterFeverTimeActivityStageScRsp                   = 7155
	MatchResultScNotify                                = 7342
	GetRogueInfoScRsp                                  = 1820
	SyncRogueCommonDialogueDataScNotify                = 5638
	ChessRogueQueryAeonDimensionsCsReq                 = 5512
	MarkChatEmojiCsReq                                 = 3947
	SetFriendRemarkNameCsReq                           = 2929
	FinishChessRogueSubStoryScRsp                      = 5565
	GetAuthkeyCsReq                                    = 90
	GetRogueAdventureRoomInfoScRsp                     = 5648
	RogueTournReviveAvatarScRsp                        = 6035
	GetSaveLogisticsMapCsReq                           = 4794
	GetMonopolyInfoScRsp                               = 7020
	TravelBrochureSetPageDescStatusCsReq               = 6479
	MakeMissionDrinkCsReq                              = 6984
	ExpUpEquipmentScRsp                                = 574
	GetChapterScRsp                                    = 453
	MuseumRandomEventSelectScRsp                       = 4351
	GroupStateChangeScNotify                           = 1441
	GetUpdatedArchiveDataCsReq                         = 2303
	EndDrinkMakerSequenceCsReq                         = 6992
	PrepareRogueAdventureRoomCsReq                     = 5620
	ChessRogueUpdateAllowedSelectCellScNotify          = 5592
	MonopolySelectOptionCsReq                          = 7016
	PlayerReturnSignCsReq                              = 4520
	TrainPartyBuildDiyCsReq                            = 8048
	HandleRogueCommonPendingActionCsReq                = 5624
	MonopolyReRollRandomCsReq                          = 7079
	GetStageLineupScRsp                                = 720
	TravelBrochureApplyPasterCsReq                     = 6434
	SetClientPausedCsReq                               = 1488
	RecoverAllLineupScRsp                              = 1465
	GmTalkCsReq                                        = 16
	RogueWorkbenchHandleFuncCsReq                      = 5607
	SetCurWaypointCsReq                                = 403
	RogueArcadeGetInfoScRsp                            = 7673
	ClockParkUnlockTalentScRsp                         = 7237
	HeliobusActivityDataScRsp                          = 5820
	DeployRotaterCsReq                                 = 6839
	GetAuthkeyScRsp                                    = 79
	ScenePlaneEventScNotify                            = 1407
	EnterTelevisionActivityStageCsReq                  = 6970
	SpringRecoverCsReq                                 = 1472
	RaidCollectionDataScNotify                         = 6955
	SyncRogueAeonLevelUpRewardScNotify                 = 1898
	SwordTrainingSelectEndingCsReq                     = 7467
	ChessRogueQuitScRsp                                = 5595
	FinishCurTurnScRsp                                 = 4374
	GetMuseumInfoScRsp                                 = 4320
	AcceptedPamMissionExpireCsReq                      = 4059
	ChessRogueQueryAeonDimensionsScRsp                 = 5594
	SwordTrainingResumeGameCsReq                       = 7453
	SpringRecoverScRsp                                 = 1424
	MonopolySelectOptionScRsp                          = 7047
	GetBagScRsp                                        = 520
	TakeActivityExpeditionRewardScRsp                  = 2548
	SetForbidOtherApplyFriendScRsp                     = 2945
	StartRaidScRsp                                     = 2220
	RogueTournEnterLayerScRsp                          = 6053
	SceneCastSkillScRsp                                = 1453
	GetShopListCsReq                                   = 1559
	SpaceZooDataScRsp                                  = 6720
	AcceptActivityExpeditionCsReq                      = 2516
	UpgradeAreaScRsp                                   = 4375
	PlayerGetTokenScRsp                                = 53
	RogueMagicGetTalentInfoCsReq                       = 7736
	MultiplayerFightGiveUpScRsp                        = 1053
	BatchGetQuestDataCsReq                             = 948
	GetRaidInfoScRsp                                   = 2274
	DoGachaScRsp                                       = 1946
	DiscardRelicScRsp                                  = 514
	TakeLoginActivityRewardScRsp                       = 2646
	TakeApRewardScRsp                                  = 3320
	GetNpcMessageGroupCsReq                            = 2759
	MonopolyGiveUpCurContentScRsp                      = 7099
	WolfBroGameUseBulletScRsp                          = 6530
	DeleteSummonUnitCsReq                              = 1495
	AvatarExpUpScRsp                                   = 346
	NewAssistHistoryNotify                             = 2977
	MonopolyReRollRandomScRsp                          = 7019
	TakeTalkRewardCsReq                                = 2103
	LogisticsInfoScNotify                              = 4793
	ChessRogueGiveUpRollScRsp                          = 5522
	GetBagCsReq                                        = 559
	SearchPlayerCsReq                                  = 2909
	PunkLordBattleResultScNotify                       = 3251
	EnterSceneScRsp                                    = 1422
	DressAvatarSkinCsReq                               = 333
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5406
	MonopolyTakePhaseRewardCsReq                       = 7042
	RogueMagicLeaveScRsp                               = 7753
	HealPoolInfoNotify                                 = 1412
	DressAvatarCsReq                                   = 380
	TakeApRewardCsReq                                  = 3359
	GetFriendBattleRecordDetailCsReq                   = 2932
	GetChallengeRecommendLineupListCsReq               = 2442
	ActivityRaidPlacingGameScRsp                       = 4743
	SyncRogueCommonDialogueOptionFinishScNotify        = 5655
	ActivateFarmElementCsReq                           = 1415
	GetJukeboxDataScRsp                                = 3120
	CancelMatchCsReq                                   = 7312
	MonopolyAcceptQuizCsReq                            = 7077
	MultiplayerFightGameStateCsReq                     = 1059
	MuseumTargetRewardNotify                           = 4309
	SelectChessRogueNousSubStoryScRsp                  = 5496
	SettleTrackPhotoStageScRsp                         = 7558
	SyncRogueMapRoomScNotify                           = 1814
	SpaceZooCatUpdateNotify                            = 6747
	GetPlayerReturnMultiDropInfoCsReq                  = 4646
	SyncRogueAreaUnlockScNotify                        = 1807
	MonopolyTakeRaffleTicketRewardScRsp                = 7018
	GetChatEmojiListScRsp                              = 3916
	EvolveBuildShopAbilityDownCsReq                    = 7144
	StartRogueScRsp                                    = 1846
	SaveLogisticsScRsp                                 = 4709
	RogueMagicGetMiscRealTimeDataCsReq                 = 7704
	RecoverAllLineupCsReq                              = 1436
	WolfBroGameActivateBulletScRsp                     = 6579
	GetAlleyInfoCsReq                                  = 4759
	DestroyItemCsReq                                   = 510
	RogueMagicBattleFailSettleInfoScNotify             = 7719
	LogisticsScoreRewardSyncInfoScNotify               = 4704
	DifficultyAdjustmentUpdateDataScRsp                = 4170
	GetLevelRewardScRsp                                = 83
	TakeChapterRewardScRsp                             = 480
	LeaveMapRotationRegionScRsp                        = 6816
	ClientObjUploadCsReq                               = 40
	EvolveBuildQueryInfoScRsp                          = 7116
	BuyRogueShopBuffScRsp                              = 5647
	AetherDivideTakeChallengeRewardScRsp               = 4877
	TakeRogueEndlessActivityPointRewardCsReq           = 6002
	TakeRollShopRewardScRsp                            = 6903
	GetSpringRecoverDataScRsp                          = 1466
	TrainPartyGamePlayStartCsReq                       = 8009
	GetCurChallengeScRsp                               = 1747
	SyncRogueAdventureRoomInfoScNotify                 = 5659
	MarkReadMailScRsp                                  = 846
	GetRogueCollectionCsReq                            = 5681
	SearchPlayerScRsp                                  = 2993
	MonopolyQuizDurationChangeScNotify                 = 7015
	MonopolyGetRegionProgressScRsp                     = 7068
	TakeAllApRewardCsReq                               = 3353
	MultipleDropInfoNotify                             = 4653
	AddBlacklistScRsp                                  = 2979
	ChallengeRaidNotify                                = 2216
	MarkItemScRsp                                      = 513
	LobbyInviteScNotify                                = 7389
	PlayerLoginFinishCsReq                             = 55
	GetAllServerPrefsDataCsReq                         = 6159
	EvolveBuildStartLevelScRsp                         = 7124
	OpenRogueChestScRsp                                = 1832
	GetRogueExhibitionScRsp                            = 5698
	BatchMarkChatEmojiScRsp                            = 3975
	ClockParkGetOngoingScriptInfoCsReq                 = 7208
	PlayerReturnTakePointRewardCsReq                   = 4539
	TakeRogueAeonLevelRewardScRsp                      = 1850
	ChessRogueSelectBpScRsp                            = 5561
	ChessRogueConfirmRollCsReq                         = 5580
	GroupStateChangeScRsp                              = 1401
	GetPhoneDataCsReq                                  = 5159
	GetChatFriendHistoryCsReq                          = 3934
	ExpUpRelicCsReq                                    = 548
	SyncRogueVirtualItemInfoScNotify                   = 1886
	AlleyFundsScNotify                                 = 4751
	SpaceZooMutateScRsp                                = 6753
	SwordTrainingMarkEndingViewedCsReq                 = 7495
	StartChallengeCsReq                                = 1703
	MonopolyBuyGoodsScRsp                              = 7033
	TakeKilledPunkLordMonsterScoreScRsp                = 3294
	FinishTalkMissionScRsp                             = 1246
	ReviveRogueAvatarScRsp                             = 1825
	FinishFirstTalkNpcScRsp                            = 2137
	CancelExpeditionScRsp                              = 2553
	GetLoginActivityScRsp                              = 2620
	RotateMapCsReq                                     = 6834
	FinishSectionIdCsReq                               = 2734
	AlleyEventChangeNotify                             = 4780
	ServerSimulateBattleFinishScNotify                 = 174
	StartFinishMainMissionScNotify                     = 1294
	GetPunkLordDataCsReq                               = 3290
	LeaveRaidScRsp                                     = 2246
	GetGachaInfoScRsp                                  = 1920
	PlayBackGroundMusicScRsp                           = 3146
	GetPunkLordDataScRsp                               = 3279
	RogueArcadeLeaveScRsp                              = 7674
	TakeRogueMiracleHandbookRewardCsReq                = 5688
	DeactivateFarmElementCsReq                         = 1414
	EnhanceCommonRogueBuffCsReq                        = 5651
	ExchangeRogueBuffWithMiracleScRsp                  = 5625
	UnlockBackGroundMusicScRsp                         = 3153
	RogueTournGetMiscRealTimeDataScRsp                 = 6076
	ReEnterLastElementStageCsReq                       = 1484
	SpaceZooBornCsReq                                  = 6703
	GetAetherDivideChallengeInfoScRsp                  = 4809
	ChessRogueGiveUpCsReq                              = 5476
	RogueTournClearArchiveNameScNotify                 = 6086
	GetLoginActivityCsReq                              = 2659
	LeaveMapRotationRegionScNotify                     = 6848
	MonopolyTakeRaffleTicketRewardCsReq                = 7007
	GetRogueTalentInfoScRsp                            = 1808
	RogueTournBattleFailSettleInfoScNotify             = 6048
	SwordTrainingStoryBattleScRsp                      = 7498
	RebattleByClientCsNotify                           = 175
	GetFriendApplyListInfoScRsp                        = 2953
	MuseumTargetMissionFinishNotify                    = 4399
	MusicRhythmUnlockSongNotify                        = 7589
	SpaceZooBornScRsp                                  = 6746
	SyncChessRogueNousValueScNotify                    = 5405
	MuseumTargetStartNotify                            = 4383
	GetEnteredSceneScRsp                               = 1487
	CancelActivityExpeditionScRsp                      = 2530
	DoGachaInRollShopCsReq                             = 6915
	GetMbtiReportScRsp                                 = 7031
	UpdateRogueAdventureRoomScoreScRsp                 = 5602
	SyncRogueFinishScNotify                            = 1848
	GetNpcTakenRewardScRsp                             = 2120
	StartMatchCsReq                                    = 7315
	FightActivityDataChangeScNotify                    = 3603
	HeliobusSnsLikeCsReq                               = 5834
	RogueTournRenameArchiveScRsp                       = 6061
	RankUpAvatarScRsp                                  = 348
	EnterMapRotationRegionScRsp                        = 6820
	UpdatePlayerSettingScRsp                           = 98
	PlayerReturnTakePointRewardScRsp                   = 4553
	SwordTrainingDailyPhaseConfirmCsReq                = 7487
	AlleyShipUsedCountScNotify                         = 4770
	GetLineupAvatarDataCsReq                           = 774
	SetSignatureScRsp                                  = 2847
	SetRogueExhibitionCsReq                            = 5642
	ClockParkStartScriptScRsp                          = 7223
	SetAssistAvatarCsReq                               = 2874
	SwordTrainingGiveUpGameCsReq                       = 7468
	ChessRogueSelectBpCsReq                            = 5448
	SetIsDisplayAvatarInfoCsReq                        = 2834
	GetWolfBroGameDataScRsp                            = 6516
	FantasticStoryActivityBattleEndScNotify            = 4953
	GetOfferingInfoScRsp                               = 6928
	PromoteEquipmentScRsp                              = 546
	MuseumRandomEventQueryCsReq                        = 4325
	TrainPartyUseCardScRsp                             = 8046
	GetChessRogueBuffEnhanceInfoCsReq                  = 5562
	GetPlayerReplayInfoCsReq                           = 3503
	RogueTournDeleteArchiveScRsp                       = 6058
	GetShopListScRsp                                   = 1520
	SetLineupNameCsReq                                 = 719
	UpdateGunPlayDataCsReq                             = 4109
	GetMissionStatusScRsp                              = 1229
	FinishAeonDialogueGroupCsReq                       = 1887
	SyncServerSceneChangeNotify                        = 1450
	RogueMagicEnterRoomCsReq                           = 7780
	TrainPartyBuildStartStepScRsp                      = 8075
	TrainPartyBuildingUpdateNotify                     = 8079
	PrestigeLevelUpCsReq                               = 4729
	ChessRoguePickAvatarCsReq                          = 5585
	TakeAllApRewardScRsp                               = 3334
	CancelMarkItemNotify                               = 577
	StartFinishSubMissionScNotify                      = 1264
	TeleportToMissionResetPointScRsp                   = 1293
	RogueArcadeStartCsReq                              = 7665
	TrainPartyGetDataScRsp                             = 8020
	GetMainMissionCustomValueScRsp                     = 1265
	GetAetherDivideChallengeInfoCsReq                  = 4899
	RogueTournResetPermanentTalentScRsp                = 6087
	ChessRogueNousEditDiceScRsp                        = 5599
	SetRogueCollectionScRsp                            = 5668
	EvolveBuildShopAbilityUpScRsp                      = 7133
	TrainPartyGamePlaySettleNotify                     = 8064
	RogueTournWeekChallengeUpdateScNotify              = 6088
	SyncRogueReviveInfoScNotify                        = 1810
	RogueModifierDelNotify                             = 5380
	EnterTreasureDungeonCsReq                          = 4430
	EvolveBuildUnlockInfoNotify                        = 7126
	GetNpcStatusCsReq                                  = 2703
	UpdateTrackMainMissionIdScRsp                      = 1291
	ChessRogueRollDiceScRsp                            = 5532
	DeleteFriendCsReq                                  = 2930
	GetKilledPunkLordMonsterDataCsReq                  = 3273
	TravelBrochureGetPasterScNotify                    = 6430
	ChessRogueQueryBpScRsp                             = 5535
	GetChessRogueStoryInfoCsReq                        = 5509
	PromoteEquipmentCsReq                              = 503
	AetherDivideFinishChallengeScNotify                = 4893
	HeliobusEnterBattleScRsp                           = 5829
	MonopolyGameSettleScNotify                         = 7070
	GetMbtiReportCsReq                                 = 7096
	MonopolyClickCellCsReq                             = 7041
	GetRogueShopBuffInfoCsReq                          = 5653
	GetCurSceneInfoCsReq                               = 1434
	SetClientRaidTargetCountScRsp                      = 2275
	RogueTournLeaveRogueCocoonSceneScRsp               = 6047
	OpenTreasureDungeonGridScRsp                       = 4490
	StartWolfBroGameScRsp                              = 6520
	MusicRhythmSaveSongConfigDataCsReq                 = 7599
	SubmitOfferingItemCsReq                            = 6935
	DestroyItemScRsp                                   = 570
	DifficultyAdjustmentGetDataScRsp                   = 4194
	SwordTrainingSetSkillTraceCsReq                    = 7500
	CancelCacheNotifyCsReq                             = 4180
	FightMatch3TurnStartScNotify                       = 30146
	RogueMagicReviveCostUpdateScNotify                 = 7761
	SceneCastSkillMpUpdateScNotify                     = 1490
	SwordTrainingExamResultConfirmCsReq                = 7479
	ChessRogueMoveCellNotify                           = 5547
	MusicRhythmDataScRsp                               = 7574
	GetRogueBuffEnhanceInfoCsReq                       = 1851
	TravelBrochureUpdatePasterPosCsReq                 = 6447
	ChessRogueNousGetRogueTalentInfoCsReq              = 5560
	SceneEnterStageScRsp                               = 1473
	TakeAssistRewardCsReq                              = 2991
	QuitLineupCsReq                                    = 734
	FinishedMissionScNotify                            = 1204
	ChessRogueChangeyAeonDimensionNotify               = 5458
	RogueTournGetSettleInfoScRsp                       = 6022
	FinishTutorialScRsp                                = 1616
	GetStrongChallengeActivityDataCsReq                = 6659
	SyncTurnFoodNotify                                 = 591
	TakeQuestRewardCsReq                               = 903
	DressAvatarSkinScRsp                               = 351
	ComposeLimitNumCompleteNotify                      = 564
	UpdatePlayerSettingCsReq                           = 92
	PlayerLoginScRsp                                   = 20
	ChessRogueGiveUpRollCsReq                          = 5424
	MonopolyContentUpdateScNotify                      = 7064
	FeverTimeActivityBattleEndScNotify                 = 7159
	PlayerReturnTakeRewardCsReq                        = 4534
	GetFriendAssistListScRsp                           = 2978
	ServerAnnounceNotify                               = 94
	EntityBindPropCsReq                                = 1491
	GetGameStateServiceConfigScRsp                     = 97
	GetNpcTakenRewardCsReq                             = 2159
	SelectInclinationTextCsReq                         = 2180
	GetPamSkinDataScRsp                                = 8128
	GetAllLineupDataCsReq                              = 725
	DelSaveRaidScNotify                                = 2261
	TakeRogueEventHandbookRewardScRsp                  = 5614
	GetRogueAeonInfoScRsp                              = 1858
	SpaceZooOpCatteryScRsp                             = 6737
	GetTrainVisitorRegisterScRsp                       = 3737
	GetMonopolyInfoCsReq                               = 7059
	EnterChessRogueAeonRoomScRsp                       = 5576
	ReserveStaminaExchangeScRsp                        = 69
	RogueTournGetAllArchiveScRsp                       = 6019
	StartCocoonStageScRsp                              = 1477
	EnterTrialActivityStageScRsp                       = 2660
	EvolveBuildFinishScNotify                          = 7135
	GetPunkLordMonsterDataScRsp                        = 3220
	TrainPartyBuildDiyScRsp                            = 8090
	ChessRogueCellUpdateNotify                         = 5436
	RefreshAlleyOrderCsReq                             = 4779
	MultiplayerFightGiveUpCsReq                        = 1039
	RogueMagicStoryInfoUpdateScNotify                  = 7714
	UpdateGroupPropertyCsReq                           = 1428
	InteractChargerScRsp                               = 6846
	TextJoinQueryScRsp                                 = 3846
	ResetMapRotationRegionScRsp                        = 6875
	SyncRogueCommonPendingActionScNotify               = 5615
	EnterSectionCsReq                                  = 1409
	GetFeverTimeActivityDataCsReq                      = 7158
	RogueTournHandBookNotify                           = 6012
	ClockParkUseBuffCsReq                              = 7220
	SyncRogueHandbookDataUpdateScNotify                = 5604
	RogueTournGetCurRogueCocoonInfoCsReq               = 6045
	SpaceZooDataCsReq                                  = 6759
	WolfBroGameActivateBulletCsReq                     = 6590
	HeliobusUpgradeLevelScRsp                          = 5875
	GetMarkItemListCsReq                               = 536
	MonopolyCheatDiceScRsp                             = 7093
	RogueArcadeLeaveCsReq                              = 7662
	RogueMagicAutoDressInUnitCsReq                     = 7743
	AetherDivideRefreshEndlessScNotify                 = 4826
	ExchangeGachaCeilingScRsp                          = 1937
	SetIsDisplayAvatarInfoScRsp                        = 2837
	ChessRoguePickAvatarScRsp                          = 5425
	SyncRogueCommonActionResultScNotify                = 5623
	SetTurnFoodSwitchScRsp                             = 588
	MarkReadMailCsReq                                  = 803
	GetMultipleDropInfoScRsp                           = 4620
	InteractTreasureDungeonGridScRsp                   = 4425
	PromoteAvatarScRsp                                 = 337
	UpdateEnergyScNotify                               = 6890
	TravelBrochureUpdatePasterPosScRsp                 = 6474
	RankUpAvatarCsReq                                  = 375
	QuitTreasureDungeonCsReq                           = 4451
	GetExpeditionDataCsReq                             = 2559
	FightMatch3ForceUpdateNotify                       = 30174
	SubmitEmotionItemScRsp                             = 6353
	StartBoxingClubBattleScRsp                         = 4253
	GetLevelRewardTakenListCsReq                       = 33
	SetGenderCsReq                                     = 91
	UpdateFloorSavedValueNotify                        = 1498
	HeliobusUpgradeLevelCsReq                          = 5830
	GetStageLineupCsReq                                = 759
	MuseumInfoChangedScNotify                          = 4379
	GetAllSaveRaidCsReq                                = 2279
	SubmitMaterialSubmitActivityMaterialScRsp          = 2625
	TravelBrochureSetCustomValueCsReq                  = 6448
	SceneReviveAfterRebattleCsReq                      = 1483
	TeleportToMissionResetPointCsReq                   = 1209
	RogueMagicQueryScRsp                               = 7751
	AlleyTakeEventRewardScRsp                          = 4713
	RogueDoGambleCsReq                                 = 5641
	TelevisionActivityBattleEndScNotify                = 6963
	MonopolyGuessDrawScNotify                          = 7023
	GetRogueHandbookDataScRsp                          = 5691
	TakeCityShopRewardCsReq                            = 1539
	PlayBackGroundMusicCsReq                           = 3103
	MonopolyDailySettleScNotify                        = 7078
	UpgradeAreaStatScRsp                               = 4390
	SetAvatarPathCsReq                                 = 1
	EnterRogueMapRoomCsReq                             = 1804
	RogueMagicSettleCsReq                              = 7734
	GetMarkItemListScRsp                               = 565
	ChangeStoryLineFinishScNotify                      = 6253
	TextJoinQueryCsReq                                 = 3803
	DelMailCsReq                                       = 839
	RogueNpcDisappearCsReq                             = 5674
	SetGroupCustomSaveDataCsReq                        = 1431
	ClockParkQuitScriptScRsp                           = 7234
	ComposeItemScRsp                                   = 575
	AcceptMultipleExpeditionCsReq                      = 2590
	RogueTournDifficultyCompNotify                     = 6025
	TravelBrochurePageResetCsReq                       = 6461
	EnterTrialActivityStageCsReq                       = 2612
	SwitchAetherDivideLineUpSlotCsReq                  = 4861
	GetBoxingClubInfoCsReq                             = 4259
	RogueTournGetMiscRealTimeDataCsReq                 = 6041
	ChessRogueQueryScRsp                               = 5443
	MatchThreeSetBirdPosCsReq                          = 7437
	SyncAcceptedPamMissionNotify                       = 4003
	GetPamSkinDataCsReq                                = 8125
	SharePunkLordMonsterScRsp                          = 3253
	QuitBattleCsReq                                    = 103
	TakeRogueScoreRewardCsReq                          = 1829
	RogueTournGetSettleInfoCsReq                       = 6034
	GetMovieRacingDataCsReq                            = 4129
	OpenRogueChestCsReq                                = 1860
	EnableRogueTalentCsReq                             = 1900
	SetFriendMarkCsReq                                 = 2902
	RelicRecommendScRsp                                = 2416
	WolfBroGameExplodeMonsterCsReq                     = 6519
	ReportPlayerScRsp                                  = 2973
	GetSingleRedDotParamGroupCsReq                     = 5939
	QuitWolfBroGameCsReq                               = 6534
	SetAetherDivideLineUpCsReq                         = 4830
	GetOfferingInfoCsReq                               = 6925
	AcceptMissionEventCsReq                            = 1219
	DrinkMakerUpdateTipsNotify                         = 6991
	ApplyFriendScRsp                                   = 2937
	GetGachaCeilingScRsp                               = 1953
	GetQuestDataCsReq                                  = 959
	GetAllSaveRaidScRsp                                = 2219
	ExpeditionDataChangeScNotify                       = 2580
	StartPunkLordRaidCsReq                             = 3203
	StrongChallengeActivityBattleEndScNotify           = 6639
	MonopolyGetRaffleTicketCsReq                       = 7022
	GetEnteredSceneCsReq                               = 1458
	SwordTrainingGameSettleScNotify                    = 7476
	TrainPartyAddBuildDynamicBuffCsReq                 = 8026
	GetDrinkMakerDataScRsp                             = 6988
	MonopolyGetDailyInitItemCsReq                      = 7087
	StartFightFestCsReq                                = 7262
	HeartDialTraceScriptScRsp                          = 6347
	GetPhoneDataScRsp                                  = 5120
	DressAvatarScRsp                                   = 316
	ClockParkFinishScriptScNotify                      = 7231
	MonopolyEventLoadUpdateScNotify                    = 7067
	RogueMagicSettleScRsp                              = 7737
	LockEquipmentCsReq                                 = 539
	TakeRogueAeonLevelRewardCsReq                      = 1881
	FinishRogueCommonDialogueCsReq                     = 5684
	UpdateMechanismBarScNotify                         = 1496
	AddBlacklistCsReq                                  = 2990
	RelicFilterPlanClearNameScNotify                   = 571
	FightMatch3ChatCsReq                               = 30180
	GetFarmStageGachaInfoCsReq                         = 1303
	TakeOffRelicCsReq                                  = 319
	MultiplayerFightGameStateScRsp                     = 1020
	MuseumRandomEventStartScNotify                     = 4361
	AceAntiCheaterScRsp                                = 12
	TakeAllRewardScRsp                                 = 3080
	HeliobusStartRaidCsReq                             = 5833
	RogueTournSettleCsReq                              = 6046
	GetPunkLordBattleRecordCsReq                       = 3270
	GetCurAssistScRsp                                  = 2965
	SetTurnFoodSwitchCsReq                             = 504
	GetMissionEventDataScRsp                           = 1290
	RogueMagicUnitReforgeScRsp                         = 7770
	EnterStrongChallengeActivityStageScRsp             = 6646
	EvolveBuildGiveupScRsp                             = 7123
	TakeOffAvatarSkinScRsp                             = 383
	SyncClientResVersionScRsp                          = 137
	MonopolyGameRaiseRatioCsReq                        = 7094
	FinishPlotCsReq                                    = 1159
	LogisticsDetonateStarSkiffScRsp                    = 4791
	GetTrainVisitorBehaviorCsReq                       = 3703
	StartBattleCollegeCsReq                            = 5746
	UnlockTutorialGuideCsReq                           = 1634
	LeaveChallengeCsReq                                = 1739
	ModifyRelicFilterPlanCsReq                         = 524
	BatchMarkChatEmojiCsReq                            = 3930
	JoinLineupScRsp                                    = 753
	SelectPamSkinScRsp                                 = 8130
	SpaceZooExchangeItemScRsp                          = 6730
	GetAlleyInfoScRsp                                  = 4720
	GetCurChallengeCsReq                               = 1716
	QuitRogueCsReq                                     = 1870
	RogueEndlessActivityBattleEndScNotify              = 6001
	ChessRogueUpdateMoneyInfoScNotify                  = 5477
	TakeMailAttachmentCsReq                            = 834
	RogueTournStartScRsp                               = 6080
	RemoveRotaterScRsp                                 = 6861
	TakePunkLordPointRewardCsReq                       = 3230
	SavePointsInfoNotify                               = 1426
	EvolveBuildLeaveCsReq                              = 7108
	FinishQuestScRsp                                   = 947
	LeaveMapRotationRegionCsReq                        = 6880
	RogueTournGetArchiveRepositoryScRsp                = 6072
	EnterTreasureDungeonScRsp                          = 4475
	MarkRelicFilterPlanScRsp                           = 531
	LobbyKickOutScRsp                                  = 7380
	PlayerSyncScNotify                                 = 659
	UseItemScRsp                                       = 537
	SetRogueExhibitionScRsp                            = 5640
	RogueTournReviveAvatarCsReq                        = 6100
	RogueTournQueryScRsp                               = 6029
	UnlockAvatarPathScRsp                              = 8
	SetAssistAvatarScRsp                               = 2830
	MonopolyGetDailyInitItemScRsp                      = 7044
	PlayerReturnInfoQueryScRsp                         = 4516
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6007
	SwordTrainingLearnSkillCsReq                       = 7489
	GetPrivateChatHistoryScRsp                         = 3953
	GetMapRotationDataScRsp                            = 6874
	GetFriendApplyListInfoCsReq                        = 2939
	PlayerLoginCsReq                                   = 59
	EvolveBuildReRandomStageCsReq                      = 7128
	CancelActivityExpeditionCsReq                      = 2574
	GetServerPrefsDataCsReq                            = 6103
	SetNicknameCsReq                                   = 25
	RelicAvatarRecommendScRsp                          = 2424
	ExchangeStaminaCsReq                               = 75
	GetUnlockTeleportCsReq                             = 1469
	DoGachaInRollShopScRsp                             = 6910
	GetReplayTokenScRsp                                = 3520
	GetRollShopInfoCsReq                               = 6905
	GameplayCounterCountDownCsReq                      = 1442
	SwordTrainingSetSkillTraceScRsp                    = 7461
	ChessRogueCheatRollCsReq                           = 5544
	GetChatEmojiListCsReq                              = 3980
	GetPrivateChatHistoryCsReq                         = 3939
	PrivateMsgOfflineUsersScNotify                     = 3946
	SwapLineupScRsp                                    = 716
	PromoteAvatarCsReq                                 = 334
	DailyFirstMeetPamCsReq                             = 3403
	FinishCurTurnCsReq                                 = 4347
	AddEquipmentScNotify                               = 599
	MultiplayerGetFightGateCsReq                       = 1003
	SwordTrainingMarkEndingViewedScRsp                 = 7464
	GetFriendDevelopmentInfoScRsp                      = 2971
	EnterSwordTrainingExamScRsp                        = 7499
	MonopolyClickMbtiReportScRsp                       = 7100
	TakePictureCsReq                                   = 4139
	HandleRogueCommonPendingActionScRsp                = 5612
	TakeOffEquipmentScRsp                              = 374
	InteractTreasureDungeonGridCsReq                   = 4461
	ExchangeHcoinCsReq                                 = 533
	TrainVisitorRewardSendNotify                       = 3753
	GetTutorialGuideScRsp                              = 1646
	GetWaypointScRsp                                   = 420
	MonopolyGameCreateScNotify                         = 7004
	EnterAdventureScRsp                                = 1320
	HeliobusEnterBattleCsReq                           = 5825
	SyncClientResVersionCsReq                          = 134
	ExchangeGachaCeilingCsReq                          = 1934
	TakePromotionRewardCsReq                           = 325
	DrinkMakerChallengeCsReq                           = 6982
	FinishPerformSectionIdCsReq                        = 2780
	CurPetChangedScNotify                              = 7640
	RogueTournQueryCsReq                               = 6055
	OpenTreasureDungeonGridCsReq                       = 4448
	DressRelicAvatarCsReq                              = 390
	RelicAvatarRecommendCsReq                          = 2412
	WolfBroGamePickupBulletCsReq                       = 6575
	UnlockAvatarSkinScNotify                           = 399
	TakeAssistRewardScRsp                              = 2904
	SelectPhoneThemeScRsp                              = 5134
	SetDisplayAvatarCsReq                              = 2839
	TakePictureScRsp                                   = 4153
	MonopolyMoveCsReq                                  = 7037
	ChessRogueLayerAccountInfoNotify                   = 5468
	DelMailScRsp                                       = 853
	DeleteSocialEventServerCacheScRsp                  = 7055
	SceneReviveAfterRebattleScRsp                      = 1499
	GetStarFightDataCsReq                              = 7166
	FinishFirstTalkByPerformanceNpcCsReq               = 2130
	EnhanceChessRogueBuffCsReq                         = 5471
	MatchThreeSetBirdPosScRsp                          = 7440
	ReEnterLastElementStageScRsp                       = 1427
	ContentPackageTransferScNotify                     = 7537
	RogueTournLeaveCsReq                               = 6016
	SwordTrainingTurnActionCsReq                       = 7474
	GetLoginChatInfoCsReq                              = 3948
	AcceptMainMissionScRsp                             = 1270
	RecallPetCsReq                                     = 7642
	GetRaidInfoCsReq                                   = 2247
	CommonRogueComponentUpdateScNotify                 = 5606
	TrainPartyEnterScRsp                               = 8025
	TrainRefreshTimeNotify                             = 3739
	SyncRogueAeonScNotify                              = 1835
	RogueTournEnterCsReq                               = 6026
	SceneCastSkillCostMpScRsp                          = 1448
	EnterSceneByServerScNotify                         = 1435
	EnterSummonActivityStageCsReq                      = 7569
	SelectChessRogueNousSubStoryCsReq                  = 5472
	GetRelicFilterPlanScRsp                            = 556
	LobbyModifyPlayerInfoCsReq                         = 7390
	MonopolyConfirmRandomCsReq                         = 7061
	AetherDivideTainerInfoScNotify                     = 4864
	AlleyShipmentEventEffectsScNotify                  = 4764
	SetLineupNameScRsp                                 = 761
	RogueDoGambleScRsp                                 = 5658
	RogueMagicEnterLayerCsReq                          = 7747
	RogueTournResetPermanentTalentCsReq                = 6075
	QuitWolfBroGameScRsp                               = 6537
	DeleteSocialEventServerCacheCsReq                  = 7038
	AceAntiCheaterCsReq                                = 24
	CommonRogueUpdateScNotify                          = 5696
	EnhanceChessRogueBuffScRsp                         = 5457
	RogueMagicUnitComposeCsReq                         = 7764
	StartTimedFarmElementCsReq                         = 1486
	GetAvatarDataCsReq                                 = 359
	GetStoryLineInfoCsReq                              = 6259
	TakeMultipleExpeditionRewardScRsp                  = 2561
	GetFriendLoginInfoScRsp                            = 2923
	ReportPlayerCsReq                                  = 2951
	DeployRotaterScRsp                                 = 6853
	GetRndOptionCsReq                                  = 3459
	TakeRogueScoreRewardScRsp                          = 1833
	PVEBattleResultCsReq                               = 159
	ChessRogueNousGetRogueTalentInfoScRsp              = 5505
	FinishFirstTalkNpcCsReq                            = 2134
	WolfBroGamePickupBulletScRsp                       = 6548
	SceneEnterStageCsReq                               = 1451
	SetRogueCollectionCsReq                            = 5669
	RogueTournEnterRogueCocoonSceneScRsp               = 6071
	MultiplayerFightGameStartScNotify                  = 1034
	GetQuestDataScRsp                                  = 920
	RaidCollectionEnterNextRaidScRsp                   = 6952
	SyncEntityBuffChangeListScNotify                   = 1430
	UnlockHeadIconScNotify                             = 2880
	GetQuestRecordScRsp                                = 937
	RogueTournTakeExpRewardScRsp                       = 6057
	LeaveRogueScRsp                                    = 1837
	TakeRogueMiracleHandbookRewardScRsp                = 5643
	PickRogueAvatarCsReq                               = 1890
	EvolveBuildShopAbilityResetScRsp                   = 7132
	TakeFightActivityRewardCsReq                       = 3653
	FightHeartBeatScRsp                                = 30053
	DeleteBlacklistCsReq                               = 2983
	RotateMapScRsp                                     = 6837
	MatchThreeGetDataCsReq                             = 7415
	RestartChallengePhaseScRsp                         = 1751
	GetChatFriendHistoryScRsp                          = 3937
	SummonPetCsReq                                     = 7612
	StoryLineInfoScNotify                              = 6203
	GetCurSceneInfoScRsp                               = 1437
	UpdateMovieRacingDataScRsp                         = 4173
	SelectChatBubbleCsReq                              = 5103
	UpdateRedDotDataCsReq                              = 5903
	GetHeartDialInfoScRsp                              = 6320
	FinishChapterScNotify                              = 4903
	SummonPetScRsp                                     = 7624
	PunkLordMonsterKilledNotify                        = 3293
	TakeQuestOptionalRewardCsReq                       = 974
	FightMatch3DataCsReq                               = 30159
	SyncLineupNotify                                   = 747
	EvolveBuildQueryInfoCsReq                          = 7115
	GetPetDataScRsp                                    = 7616
	TextJoinBatchSaveScRsp                             = 3853
	GetTrackPhotoActivityDataCsReq                     = 7556
	FightFestUnlockSkillNotify                         = 7287
	SwitchLineupIndexScRsp                             = 779
	GetFirstTalkByPerformanceNpcScRsp                  = 2174
	EnterAetherDivideSceneCsReq                        = 4859
	GetNpcStatusScRsp                                  = 2746
	WolfBroGameUseBulletCsReq                          = 6574
	ChangeLineupLeaderCsReq                            = 775
	GetGunPlayDataScRsp                                = 4199
	GetLevelRewardCsReq                                = 73
	ReBattleAfterBattleLoseCsNotify                    = 130
	RogueTournSettleScRsp                              = 6040
	UnlockedAreaMapScNotify                            = 1497
	TelevisionActivityDataChangeScNotify               = 6975
	PlayerLogoutScRsp                                  = 46
	HeliobusStartRaidScRsp                             = 5851
	RogueMagicUnitReforgeCsReq                         = 7710
	GetSummonActivityDataCsReq                         = 7566
	WolfBroGameDataChangeScNotify                      = 6547
	AetherDivideRefreshEndlessScRsp                    = 4865
	ClockParkUnlockTalentCsReq                         = 7242
	GetMailCsReq                                       = 859
	GetSecretKeyInfoScRsp                              = 38
	EquipAetherDividePassiveSkillScRsp                 = 4890
	ReplaceLineupScRsp                                 = 773
	LobbyCreateScRsp                                   = 7366
	MonopolyGameBingoFlipCardScRsp                     = 7013
	EnhanceRogueBuffCsReq                              = 1883
	GetRogueAeonInfoCsReq                              = 1841
	StopRogueAdventureRoomCsReq                        = 5683
	StartTrialActivityCsReq                            = 2631
	MuseumFundsChangedScNotify                         = 4319
	MarkAvatarCsReq                                    = 309
	GetRogueHandbookDataCsReq                          = 5677
	FightMatch3TurnEndScNotify                         = 30139
	GetFarmStageGachaInfoScRsp                         = 1346
	ChessRogueStartScRsp                               = 5567
	PunkLordMonsterInfoScNotify                        = 3248
	SceneEntityMoveCsReq                               = 1459
	FightMatch3ChatScRsp                               = 30116
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5633
	RogueTournLeaveScRsp                               = 6078
	GetAssistHistoryScRsp                              = 2913
	GetWaypointCsReq                                   = 459
	GetRndOptionScRsp                                  = 3420
	GetRogueScoreRewardInfoScRsp                       = 1840
	SceneEntityTeleportCsReq                           = 1438
	RestoreWolfBroGameArchiveCsReq                     = 6539
	MonopolyGameGachaCsReq                             = 7036
	MonopolyLikeScNotify                               = 7032
	RefreshAlleyOrderScRsp                             = 4719
	RogueTournGetCurRogueCocoonInfoScRsp               = 6073
	StartStarFightLevelScRsp                           = 7168
	FightMatch3SwapScRsp                               = 30134
	LogisticsDetonateStarSkiffCsReq                    = 4777
	SceneGroupRefreshScNotify                          = 1452
	MonopolyGiveUpCurContentCsReq                      = 7083
	MonopolyMoveScRsp                                  = 7080
	SelectPhoneThemeCsReq                              = 5153
	GetJukeboxDataCsReq                                = 3159
	InteractChargerCsReq                               = 6803
	MonopolyGameRaiseRatioScRsp                        = 7010
	GetLevelRewardTakenListScRsp                       = 51
	ExchangeRogueRewardKeyCsReq                        = 1896
	FightMatch3DataScRsp                               = 30120
	LobbyJoinScRsp                                     = 7374
	InterruptMissionEventScRsp                         = 1251
	MultiplayerGetFightGateScRsp                       = 1046
	AlleyEventEffectNotify                             = 4716
	GetSceneMapInfoCsReq                               = 1418
	TrainVisitorBehaviorFinishCsReq                    = 3759
	FinishEmotionDialoguePerformanceCsReq              = 6334
	ReserveStaminaExchangeCsReq                        = 50
	RogueTournEnterRogueCocoonSceneCsReq               = 6056
	ChessRogueGoAheadCsReq                             = 5539
	StopRogueAdventureRoomScRsp                        = 5699
	ArchiveWolfBroGameScRsp                            = 6546
	SubmitOrigamiItemScRsp                             = 4190
	ChangeScriptEmotionCsReq                           = 6303
	PrestigeLevelUpScRsp                               = 4733
	MonopolyEventSelectFriendCsReq                     = 7071
	GiveUpBoxingClubChallengeCsReq                     = 4234
	UnlockBackGroundMusicCsReq                         = 3139
	StartTimedCocoonStageScRsp                         = 1454
	RogueDebugMessageScNotify                          = 5611
	GetRogueCollectionScRsp                            = 5650
	LeaveChallengeScRsp                                = 1753
	ChessRogueEnterNextLayerScRsp                      = 5404
	FightFestUpdateChallengeRecordNotify               = 7290
	SyncChessRogueMainStoryFinishScNotify              = 5593
	HeliobusSnsReadScRsp                               = 5846
	TriggerVoiceScRsp                                  = 4175
	GetAvatarDataScRsp                                 = 320
	GetBattleCollegeDataScRsp                          = 5720
	AcceptMultipleExpeditionScRsp                      = 2579
	MakeDrinkScRsp                                     = 6990
	EvolveBuildStartStageCsReq                         = 7142
	TrainPartyLeaveCsReq                               = 8029
	SendMsgCsReq                                       = 3959
	ReturnLastTownCsReq                                = 1429
	SwordTrainingResumeGameScRsp                       = 7497
	ClockParkGetOngoingScriptInfoScRsp                 = 7230
	TakeChallengeRewardCsReq                           = 1748
	FightEnterCsReq                                    = 30059
	MonopolyGetRafflePoolInfoScRsp                     = 7098
	MusicRhythmFinishLevelCsReq                        = 7587
	LeaveTrialActivityScRsp                            = 2662
	GetRogueEndlessActivityDataCsReq                   = 6006
	EnterMapRotationRegionCsReq                        = 6859
	MonopolyRollDiceCsReq                              = 7053
	UpgradeAreaStatCsReq                               = 4348
	TravelBrochurePageUnlockScNotify                   = 6403
	SetCurWaypointScRsp                                = 446
	EnterTelevisionActivityStageScRsp                  = 6972
	FinishTutorialGuideCsReq                           = 1647
	LockRelicCsReq                                     = 579
	EnterFantasticStoryActivityStageCsReq              = 4946
	GetFantasticStoryActivityDataScRsp                 = 4920
	GetMissionEventDataCsReq                           = 1248
	GetRogueCommonDialogueDataScRsp                    = 5671
	SpringRefreshScRsp                                 = 1461
	FinishItemIdCsReq                                  = 2739
	ChessRogueReRollDiceScRsp                          = 5494
	MonopolyClickCellScRsp                             = 7058
	MultiplayerFightGameFinishScNotify                 = 1037
	MissionEventRewardScNotify                         = 1279
	RogueMagicReviveAvatarCsReq                        = 7725
	SpaceZooMutateCsReq                                = 6739
	MonopolyGetRafflePoolInfoCsReq                     = 7092
	EquipAetherDividePassiveSkillCsReq                 = 4848
	StartFightFestScRsp                                = 7274
	PlayerReturnTakeRewardScRsp                        = 4537
	GetRogueInitialScoreScRsp                          = 1863
	UnlockTutorialScRsp                                = 1653
	StartAlleyEventScRsp                               = 4737
	SetLanguageCsReq                                   = 93
	SpringRecoverSingleAvatarScRsp                     = 1432
	GetChessRogueBuffEnhanceInfoScRsp                  = 5540
	VirtualLineupTrialAvatarChangeScNotify             = 799
	LobbyBeginScRsp                                    = 7387
	RogueMagicUnitComposeScRsp                         = 7794
	SetGroupCustomSaveDataScRsp                        = 1471
	TreasureDungeonFinishScNotify                      = 4420
	EnterAetherDivideSceneScRsp                        = 4820
	UnlockSkilltreeScRsp                               = 353
	GetRogueInitialScoreCsReq                          = 1843
	SpaceZooTakeScRsp                                  = 6748
	MuseumTakeCollectRewardCsReq                       = 4393
	GetShareDataCsReq                                  = 4103
	PunkLordRaidTimeOutScNotify                        = 3261
	GetRogueShopMiracleInfoScRsp                       = 5639
	ChessRogueReviveAvatarScRsp                        = 5484
	GetTelevisionActivityDataScRsp                     = 6968
	MonopolyGameBingoFlipCardCsReq                     = 7026
	RogueMagicScepterTakeOffUnitScRsp                  = 7793
	RogueMagicReviveAvatarScRsp                        = 7729
	SetMissionEventProgressScRsp                       = 1283
	MusicRhythmFinishLevelScRsp                        = 7575
	MusicRhythmStartLevelCsReq                         = 7582
	RogueWorkbenchGetInfoCsReq                         = 5622
	AlleyPlacingGameScRsp                              = 4775
	RogueTournGetPermanentTalentInfoCsReq              = 6036
	SceneUpdatePositionVersionNotify                   = 1474
	ActivityRaidPlacingGameCsReq                       = 4788
	LeaveRogueCsReq                                    = 1834
	ExchangeRogueBuffWithMiracleCsReq                  = 5661
	GetAetherDivideInfoCsReq                           = 4847
	CancelMatchScRsp                                   = 7324
	StoryLineTrialAvatarChangeScNotify                 = 6234
	EnterAdventureCsReq                                = 1359
	SellItemCsReq                                      = 561
	RogueMagicStartScRsp                               = 7720
	SpaceZooOpCatteryCsReq                             = 6734
	TakeChallengeRewardScRsp                           = 1790
	PlayerReturnInfoQueryCsReq                         = 4580
	EnterChallengeNextPhaseCsReq                       = 1773
	CancelCacheNotifyScRsp                             = 4116
	MonopolyLikeCsReq                                  = 7012
	DailyFirstEnterMonopolyActivityScRsp               = 7075
	RogueMagicEnableTalentScRsp                        = 7713
	ChessRogueFinishCurRoomNotify                      = 5489
	SyncChessRogueNousSubStoryScNotify                 = 5411
	ChessRogueNousDiceSurfaceUnlockNotify              = 5460
	LobbyModifyPlayerInfoScRsp                         = 7373
	GetMultiPathAvatarInfoCsReq                        = 58
	StaminaInfoScNotify                                = 68
	BattlePassInfoNotify                               = 3059
	TrainPartyHandlePendingActionScRsp                 = 8074
	GetPlayerDetailInfoScRsp                           = 2946
	GetPlayerReplayInfoScRsp                           = 3546
	GetChallengeGroupStatisticsScRsp                   = 1719
	RogueGetGambleInfoCsReq                            = 5676
	RogueMagicGetTalentInfoScRsp                       = 7765
	HeliobusSnsPostScRsp                               = 5853
	MonopolyUpgradeAssetScRsp                          = 7073
	GetUnlockTeleportScRsp                             = 1468
	GetWolfBroGameDataCsReq                            = 6580
	InteractPropScRsp                                  = 1446
	InterruptMissionEventCsReq                         = 1233
	UnlockPamSkinScNotify                              = 8132
	TakeRogueEventHandbookRewardCsReq                  = 5663
	RogueTournReEnterRogueCocoonStageCsReq             = 6024
	TakeCityShopRewardScRsp                            = 1553
	DrinkMakerChallengeScRsp                           = 6998
	RogueWorkbenchSelectFuncCsReq                      = 5687
	RogueWorkbenchHandleFuncScRsp                      = 5618
	SpringRefreshCsReq                                 = 1419
	GroupStateChangeCsReq                              = 1476
	GetCurLineupDataCsReq                              = 703
	FightTreasureDungeonMonsterScRsp                   = 4419
	RogueMagicEnterScRsp                               = 7746
	StartStarFightLevelCsReq                           = 7169
	TrainPartyUnlockBuildAreaScRsp                     = 8099
	LobbyBeginCsReq                                    = 7392
	AlleyGuaranteedFundsCsReq                          = 4736
	ChooseBoxingClubStageOptionalBuffCsReq             = 4248
	GetTrainVisitorBehaviorScRsp                       = 3746
	LobbyQuitCsReq                                     = 7384
	GetPunkLordMonsterDataCsReq                        = 3259
	DeleteBlacklistScRsp                               = 2999
	RechargeSuccNotify                                 = 529
	MonopolyRollRandomScRsp                            = 7090
	QuestRecordScNotify                                = 980
	GetReplayTokenCsReq                                = 3559
	RogueTournAreaUpdateScNotify                       = 6039
	HeliobusSelectSkillScRsp                           = 5879
	UpdateMovieRacingDataCsReq                         = 4151
	FinishAeonDialogueGroupScRsp                       = 1844
	ChessRogueSelectCellCsReq                          = 5542
	MonopolyGuessChooseScRsp                           = 7043
	RelicRecommendCsReq                                = 2415
	GetRogueAdventureRoomInfoCsReq                     = 5675
	LobbySyncInfoScNotify                              = 7385
	SetClientPausedScRsp                               = 1443
	AlleyShopLevelScNotify                             = 4773
	GetFriendDevelopmentInfoCsReq                      = 2931
	GetRelicFilterPlanCsReq                            = 566
	SetStuffToAreaCsReq                                = 4339
	SetNicknameScRsp                                   = 29
	ClientObjDownloadDataScNotify                      = 42
	RogueTournReEnterRogueCocoonStageScRsp             = 6033
	HeliobusInfoChangedScNotify                        = 5874
	GetPlayerBoardDataScRsp                            = 2820
	ExpUpEquipmentCsReq                                = 547
	DifficultyAdjustmentUpdateDataCsReq                = 4110
	EnterFantasticStoryActivityStageScRsp              = 4939
	TrainPartySettleNotify                             = 8034
	ChessRogueReviveAvatarCsReq                        = 5597
	LobbyQuitScRsp                                     = 7399
	LobbyGetInfoCsReq                                  = 7383
	UpgradeAreaCsReq                                   = 4330
	RogueNpcDisappearScRsp                             = 5630
	MonopolyGuessChooseCsReq                           = 7088
	MonopolyUpgradeAssetCsReq                          = 7051
	GetRogueInfoCsReq                                  = 1859
	RogueMagicEnterRoomScRsp                           = 7716
	SpaceZooDeleteCatScRsp                             = 6716
	StartTrialActivityScRsp                            = 2671
	ClockParkGetInfoScRsp                              = 7216
	SetDisplayAvatarScRsp                              = 2853
	SwordTrainingStartGameScRsp                        = 7470
	ShowNewSupplementVisitorCsReq                      = 3747
	EnterChallengeNextPhaseScRsp                       = 1783
	ContentPackageGetDataCsReq                         = 7515
	SetClientRaidTargetCountCsReq                      = 2230
	ReturnLastTownScRsp                                = 1433
	SwordTrainingStoryBattleCsReq                      = 7454
	TrainPartyBuildStartStepCsReq                      = 8030
	SetBoxingClubResonanceLineupScRsp                  = 4275
	LobbyKickOutCsReq                                  = 7358
	QuitTrackPhotoStageCsReq                           = 7560
	ChessRogueEnterScRsp                               = 5575
	GetMonopolyDailyReportScRsp                        = 7001
	GetVideoVersionKeyCsReq                            = 22
	HandleFriendScRsp                                  = 2947
	UnlockSkilltreeCsReq                               = 339
	BuyBpLevelScRsp                                    = 3034
	HeliobusSnsLikeScRsp                               = 5837
	ClearAetherDividePassiveSkillScRsp                 = 4819
	UpdatePsnSettingsInfoScRsp                         = 57
	AetherDivideSpiritExpUpCsReq                       = 4851
	TakeTalkRewardScRsp                                = 2146
	TextJoinSaveCsReq                                  = 3859
	RogueTournRenameArchiveCsReq                       = 6017
	TreasureDungeonDataScNotify                        = 4459
	SetSignatureCsReq                                  = 2816
	GameplayCounterCountDownScRsp                      = 1440
	GetMailScRsp                                       = 820
	TakePrestigeRewardCsReq                            = 4747
	GetFantasticStoryActivityDataCsReq                 = 4959
	RemoveStuffFromAreaScRsp                           = 4337
	ClockParkUseBuffScRsp                              = 7205
	ReplaceLineupCsReq                                 = 751
	MonopolySocialEventEffectScNotify                  = 7062
	BoxingClubChallengeUpdateScNotify                  = 4216
	EvolveBuildLeaveScRsp                              = 7130
	EnterRogueEndlessActivityStageScRsp                = 6008
	SelectChessRogueSubStoryScRsp                      = 5452
	ChessRogueStartCsReq                               = 5446
	GetFriendChallengeDetailCsReq                      = 2912
	MarkItemCsReq                                      = 526
	TravelBrochureSetCustomValueScRsp                  = 6490
	MakeDrinkCsReq                                     = 6995
	GetFightActivityDataScRsp                          = 3620
	WaypointShowNewCsNotify                            = 434
	GmTalkScRsp                                        = 47
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3716
	AetherDivideSpiritExpUpScRsp                       = 4873
	ChessRogueGoAheadScRsp                             = 5417
	SwordTrainingDialogueSelectOptionScRsp             = 7458
	AetherDivideRefreshEndlessCsReq                    = 4836
	GetDailyActiveInfoCsReq                            = 3303
	TakePunkLordPointRewardScRsp                       = 3275
	RaidCollectionDataCsReq                            = 6945
	ChallengeBossPhaseSettleNotify                     = 1799
	RefreshTriggerByClientScRsp                        = 1408
	RefreshTriggerByClientScNotify                     = 1500
	RogueArcadeRestartScRsp                            = 7687
	SyncDeleteFriendScNotify                           = 2948
	MonopolyScrachRaffleTicketScRsp                    = 7050
	SwordTrainingGiveUpGameScRsp                       = 7482
	GetTrialActivityDataCsReq                          = 2678
	HeliobusSnsReadCsReq                               = 5803
	ReviveRogueAvatarCsReq                             = 1861
	RefreshTriggerByClientCsReq                        = 1411
	GetFriendChallengeDetailScRsp                      = 2960
	RogueModifierSelectCellCsReq                       = 5346
	GameplayCounterUpdateScNotify                      = 1467
	BatchGetQuestDataScRsp                             = 990
	StartAlleyEventCsReq                               = 4734
	TakeExpeditionRewardScRsp                          = 2537
	AddRelicFilterPlanCsReq                            = 578
	TrainPartyHandlePendingActionCsReq                 = 8047
	FightSessionStopScNotify                           = 30034
	LobbyInviteCsReq                                   = 7378
	GetTrainVisitorRegisterCsReq                       = 3734
	AvatarPathChangedNotify                            = 44
	ChallengeLineupNotify                              = 1774
	FinishChessRogueNousSubStoryCsReq                  = 5447
	MonopolyTakePhaseRewardScRsp                       = 7040
	GetArchiveDataScRsp                                = 2320
	GetRecyleTimeScRsp                                 = 593
	SummonActivityBattleEndScNotify                    = 7561
	QuitBattleScNotify                                 = 180
	RogueMagicScepterDressInUnitCsReq                  = 7783
	BuyGoodsCsReq                                      = 1503
	GetActivityScheduleConfigCsReq                     = 2639
	RogueModifierUpdateNotify                          = 5337
	RogueModifierAddNotify                             = 5303
	SpaceZooTakeCsReq                                  = 6775
	TextJoinBatchSaveCsReq                             = 3839
	HeliobusChallengeUpdateScNotify                    = 5873
	TakeMailAttachmentScRsp                            = 837
	SetMissionEventProgressCsReq                       = 1273
	TrainPartyUseCardCsReq                             = 8003
	CurTrialActivityScNotify                           = 2684
	GiveUpBoxingClubChallengeScRsp                     = 4237
	HeliobusSnsCommentScRsp                            = 5816
	EnterRogueMapRoomScRsp                             = 1888
	RevcMsgScNotify                                    = 3903
	SecurityReportScRsp                                = 4174
	UpdateRotaterScNotify                              = 6825
	RaidKickByServerScNotify                           = 2225
	ChessRogueConfirmRollScRsp                         = 5450
	DeleteFriendScRsp                                  = 2975
	FinishCosumeItemMissionScRsp                       = 1275
	UpdateRogueAdventureRoomScoreCsReq                 = 5645
	GetChallengeScRsp                                  = 1720
	ChessRogueEnterCellCsReq                           = 5462
	StartMatchScRsp                                    = 7316
	CommonRogueQueryCsReq                              = 5660
	RogueTournEnterScRsp                               = 6066
	DressRelicAvatarScRsp                              = 379
	GetFightFestDataScRsp                              = 7266
	MatchThreeLevelEndCsReq                            = 7412
	TrainPartyAddBuildDynamicBuffScRsp                 = 8013
	UnlockTutorialCsReq                                = 1639
	CurAssistChangedNotify                             = 2988
	LobbyGetInfoScRsp                                  = 7394
	FinishPerformSectionIdScRsp                        = 2716
	GetBasicInfoCsReq                                  = 2
	ChessRogueQueryBpCsReq                             = 5420
	GetPlatformPlayerInfoScRsp                         = 2963
	SyncRogueSeasonFinishScNotify                      = 1813
	RogueTournGetAllArchiveCsReq                       = 6020
	PlayerReturnSignScRsp                              = 4503
	GetSpringRecoverDataCsReq                          = 1402
	TrainPartyGamePlayStartScRsp                       = 8093
	PlayerLogoutCsReq                                  = 3
	GameplayCounterRecoverCsReq                        = 1457
	ContentPackageSyncDataScNotify                     = 7512
	EnterFeverTimeActivityStageCsReq                   = 7156
	GetTrackPhotoActivityDataScRsp                     = 7555
	TravelBrochureSelectMessageCsReq                   = 6439
	FinishCosumeItemMissionCsReq                       = 1230
	GetFriendBattleRecordDetailScRsp                   = 2996
	MusicRhythmDataCsReq                               = 7590
	StartPunkLordRaidScRsp                             = 3246
	GetBoxingClubInfoScRsp                             = 4220
	SetSpringRecoverConfigScRsp                        = 1478
	GetFightFestDataCsReq                              = 7265
	SummonPunkLordMonsterScRsp                         = 3237
	SharePunkLordMonsterCsReq                          = 3239
	EvolveBuildStartStageScRsp                         = 7137
	SceneEntityTeleportScRsp                           = 1455
	ComposeItemCsReq                                   = 530
	GetQuestRecordCsReq                                = 934
	ActivateFarmElementScRsp                           = 1445
	FinishQuestCsReq                                   = 916
	ChessRogueQuestFinishNotify                        = 5524
	MonopolyCheatDiceCsReq                             = 7009
	RogueMagicScepterTakeOffUnitCsReq                  = 7709
	LastSpringRefreshTimeNotify                        = 1425
	MatchThreeLevelEndScRsp                            = 7424
	FightLeaveScNotify                                 = 30003
	GetRecyleTimeCsReq                                 = 509
	StartPartialChallengeCsReq                         = 1761
	EvolveBuildStartLevelCsReq                         = 7112
	ShareScRsp                                         = 4120
	GetMissionStatusCsReq                              = 1225
	AcceptExpeditionCsReq                              = 2503
	GetFriendChallengeLineupScRsp                      = 2924
	ComposeLimitNumUpdateNotify                        = 594
	BuyNpcStuffScRsp                                   = 4346
	MuseumRandomEventSelectCsReq                       = 4333
	UseTreasureDungeonItemCsReq                        = 4429
	GetAssistListScRsp                                 = 2994
	TrialBackGroundMusicScRsp                          = 3137
	ClearAetherDividePassiveSkillCsReq                 = 4879
	PVEBattleResultScRsp                               = 120
	CityShopInfoScNotify                               = 1534
	GetAssistHistoryCsReq                              = 2926
	GetMovieRacingDataScRsp                            = 4133
	UseItemCsReq                                       = 534
	EnhanceRogueBuffScRsp                              = 1899
	LogisticsGameCsReq                                 = 4703
	GetMainMissionCustomValueCsReq                     = 1236
	HeliobusLineupUpdateScNotify                       = 5883
	ChangeScriptEmotionScRsp                           = 6346
	DailyFirstEnterMonopolyActivityCsReq               = 7030
	QueryProductInfoCsReq                              = 14
	GetAllRedDotDataScRsp                              = 5920
	RogueArcadeGetInfoCsReq                            = 7690
	SwitchAetherDivideLineUpSlotScRsp                  = 4825
	GetChallengeRaidInfoScRsp                          = 2234
	EnhanceCommonRogueBuffScRsp                        = 5673
	GetRollShopInfoScRsp                               = 6908
	EnterChessRogueAeonRoomCsReq                       = 5533
	UpdateTrackMainMissionIdCsReq                      = 1277
	MonopolyBuyGoodsCsReq                              = 7029
	GetMonopolyMbtiReportRewardScRsp                   = 7017
	AntiAddictScNotify                                 = 61
	GetPlayerReturnMultiDropInfoScRsp                  = 4639
	EvolveBuildShopAbilityUpCsReq                      = 7139
	ApplyFriendCsReq                                   = 2934
	EnterSwordTrainingExamCsReq                        = 7484
	AetherDivideSpiritInfoScNotify                     = 4883
	BuyBpLevelCsReq                                    = 3053
	ChessRogueReRollDiceCsReq                          = 5558
	MusicRhythmStartLevelScRsp                         = 7591
	PickRogueAvatarScRsp                               = 1879
	SceneCastSkillCostMpCsReq                          = 1475
	GetSummonActivityDataScRsp                         = 7565
	ChessRogueQuitCsReq                                = 5473
	LobbyInviteScRsp                                   = 7357
	ModifyRelicFilterPlanScRsp                         = 512
	GetPlayerBoardDataCsReq                            = 2859
	SwordTrainingSelectEndingScRsp                     = 7491
	StartWolfBroGameCsReq                              = 6559
	GetUpdatedArchiveDataScRsp                         = 2346
	ClockParkQuitScriptCsReq                           = 7207
	MusicRhythmMaxDifficultyLevelsUnlockNotify         = 7580
	RogueArcadeStartScRsp                              = 7666
	GetRogueEndlessActivityDataScRsp                   = 6005
	EnterSummonActivityStageScRsp                      = 7568
	GetMissionDataScRsp                                = 1220
	SwordTrainingStoryConfirmCsReq                     = 7455
	SettleTrackPhotoStageCsReq                         = 7559
	BattleLogReportScRsp                               = 147
	ChessRogueUpdateLevelBaseInfoScNotify              = 5531
	GetChessRogueStoryAeonTalkInfoScRsp                = 5566
	RogueMagicAutoDressInMagicUnitChangeScNotify       = 7723
	StartPartialChallengeScRsp                         = 1725
	CommonRogueQueryScRsp                              = 5632
	GetFriendChallengeLineupCsReq                      = 2972
	RankUpEquipmentScRsp                               = 516
	GetCrossInfoScRsp                                  = 7340
	TakeMaterialSubmitActivityRewardCsReq              = 2629
	EnterFightActivityStageCsReq                       = 3646
	EvolveBuildGiveupCsReq                             = 7140
	GetCurLineupDataScRsp                              = 746
	GetBasicInfoScRsp                                  = 66
	StartTrackPhotoStageScRsp                          = 7552
	SubmitMaterialSubmitActivityMaterialCsReq          = 2661
	GetTreasureDungeonActivityDataScRsp                = 4474
	FinishEmotionDialoguePerformanceScRsp              = 6337
	SubmitOrigamiItemCsReq                             = 4148
	SceneEntityMoveScRsp                               = 1420
	ChessRogueUpdateUnlockLevelScNotify                = 5596
	SwordTrainingTurnActionScRsp                       = 7492
	UnlockPhoneThemeScNotify                           = 5137
	ChessRogueUpdateAeonModifierValueScNotify          = 5464
	AcceptMissionEventScRsp                            = 1261
	AcceptedPamMissionExpireScRsp                      = 4020
	ChessRogueNousDiceUpdateNotify                     = 5490
	GetDailyActiveInfoScRsp                            = 3346
	PrepareRogueAdventureRoomScRsp                     = 5603
	TakeMultipleExpeditionRewardCsReq                  = 2519
	SubmitOfferingItemScRsp                            = 6930
	RestartChallengePhaseCsReq                         = 1733
	MarkChatEmojiScRsp                                 = 3974
	SetForbidOtherApplyFriendCsReq                     = 2915
	ShareCsReq                                         = 4159
	SetRedPointStatusScNotify                          = 7
	TrainPartyTakeBuildLevelAwardScRsp                 = 8065
	GetActivityScheduleConfigScRsp                     = 2653
	ChessRogueGiveUpScRsp                              = 5461
	WolfBroGameExplodeMonsterScRsp                     = 6561
	SwapLineupCsReq                                    = 780
	RaidInfoNotify                                     = 2239
	GetRogueScoreRewardInfoCsReq                       = 1842
	ContentPackageGetDataScRsp                         = 7516
	DrinkMakerDayEndScNotify                           = 6996
	GetSocialEventServerCacheScRsp                     = 7027
	GetChallengeRaidInfoCsReq                          = 2253
	RogueTournConfirmSettleScRsp                       = 6042
	SendMsgScRsp                                       = 3920
	SelectChessRogueSubStoryCsReq                      = 5500
	GetTelevisionActivityDataCsReq                     = 6965
	TriggerVoiceCsReq                                  = 4130
	ClockParkHandleWaitOperationCsReq                  = 7235
	ExchangeHcoinScRsp                                 = 551
	LeaveAetherDivideSceneCsReq                        = 4803
	TakeOfferingRewardCsReq                            = 6932
	SetPlayerInfoScRsp                                 = 43
	GetMaterialSubmitActivityDataCsReq                 = 2679
	ChessRogueUpdateBoardScNotify                      = 5553
	FightEnterScRsp                                    = 30020
	StartBattleCollegeScRsp                            = 5739
	MonopolyConfirmRandomScRsp                         = 7025
	GameplayCounterRecoverScRsp                        = 1417
	GetSwordTrainingDataScRsp                          = 7462
	GetLineupAvatarDataScRsp                           = 730
	GetGameStateServiceConfigCsReq                     = 5
	GetStoryLineInfoScRsp                              = 6220
	GetChallengeGroupStatisticsCsReq                   = 1779
	DoGachaCsReq                                       = 1903
	EndDrinkMakerSequenceScRsp                         = 6983
	AetherDivideTakeChallengeRewardCsReq               = 4813
	GetPetDataCsReq                                    = 7615
	SecurityReportCsReq                                = 4147
	SyncAddBlacklistScNotify                           = 2919
	MonopolyGetRegionProgressCsReq                     = 7069
	GetChessRogueNousStoryInfoScRsp                    = 5538
	AddRelicFilterPlanScRsp                            = 572
	JoinLineupCsReq                                    = 739
	GetCurBattleInfoScRsp                              = 153
	PunkLordDataChangeNotify                           = 3210
	GetAllRedDotDataCsReq                              = 5959
	GetExhibitScNotify                                 = 4316
	RemoveStuffFromAreaCsReq                           = 4334
	ChessRogueLeaveScRsp                               = 5506
	FightMatch3StartCountDownScNotify                  = 30103
	RogueMagicEnterCsReq                               = 7703
	RogueTournReviveCostUpdateScNotify                 = 6037
	ChooseBoxingClubResonanceCsReq                     = 4247
	AlleyOrderChangedScNotify                          = 4761
	EvolveBuildCoinNotify                              = 7119
	GmTalkScNotify                                     = 37
	RogueMagicEnableTalentCsReq                        = 7726
	LeaveAetherDivideSceneScRsp                        = 4846
	SwordTrainingStartGameCsReq                        = 7494
	HeliobusActivityDataCsReq                          = 5859
	GetMonopolyMbtiReportRewardCsReq                   = 7057
	PlayerLoginFinishScRsp                             = 6
	GetTrialActivityDataScRsp                          = 2672
	MatchBoxingClubOpponentCsReq                       = 4203
	SyncTaskScRsp                                      = 1234
	RogueMagicStartCsReq                               = 7759
	SwordTrainingStoryConfirmScRsp                     = 7481
	GetFirstTalkNpcCsReq                               = 2139
	EvolveBuildReRandomStageScRsp                      = 7107
	SetMultipleAvatarPathsScRsp                        = 95
	GetRogueTalentInfoCsReq                            = 1811
	RogueMagicSetAutoDressInMagicUnitScRsp             = 7791
	StartAetherDivideStageBattleScRsp                  = 4833
	SwordTrainingRestoreGameCsReq                      = 7459
	RogueMagicAutoDressInUnitScRsp                     = 7763
	ChooseBoxingClubResonanceScRsp                     = 4274
	MonopolyScrachRaffleTicketCsReq                    = 7081
	TakeTrialActivityRewardScRsp                       = 2696
	GetNpcMessageGroupScRsp                            = 2720
	QuitRogueScRsp                                     = 1836
	RogueTournStartCsReq                               = 6081
	RogueTournEnterRoomCsReq                           = 6030
	GetFightActivityDataCsReq                          = 3659
	DifficultyAdjustmentGetDataCsReq                   = 4164
	TakeExpeditionRewardCsReq                          = 2534
	EntityBindPropScRsp                                = 1404
	TrainPartySyncUpdateScNotify                       = 8080
	AcceptActivityExpeditionScRsp                      = 2547
	GetRogueShopBuffInfoScRsp                          = 5634
	SceneEntityMoveScNotify                            = 1447
	SyncRogueGetItemScNotify                           = 1818
	SetBoxingClubResonanceLineupCsReq                  = 4230
	StartTimedFarmElementScRsp                         = 1421
	SetLanguageScRsp                                   = 64
	TrainVisitorBehaviorFinishScRsp                    = 3720
	GetChessRogueStoryAeonTalkInfoCsReq                = 5556
	GetCurAssistCsReq                                  = 2936
	FinishSectionIdScRsp                               = 2737
	ChessRogueEnterCellScRsp                           = 5431
	GateServerScNotify                                 = 71
	StartChallengeScRsp                                = 1746
	GetPlatformPlayerInfoCsReq                         = 2943
	DailyRefreshNotify                                 = 56
	SyncRogueExploreWinScNotify                        = 1826
	ChessRogueSkipTeachingLevelScRsp                   = 5557
	SetHeadIconScRsp                                   = 2846
	FinishPlotScRsp                                    = 1120
	ChallengeSettleNotify                              = 1734
	FinishChessRogueNousSubStoryScRsp                  = 5591
	TakeOfferingRewardScRsp                            = 6923
	SelectPamSkinCsReq                                 = 8135
	PlayerReturnPointChangeScNotify                    = 4546
	TakeBpRewardScRsp                                  = 3039
	MatchBoxingClubOpponentScRsp                       = 4246
	RogueModifierStageStartNotify                      = 5316
	SwordTrainingActionTurnSettleScNotify              = 7471
	TakeChapterRewardCsReq                             = 437
	DeleteRelicFilterPlanCsReq                         = 560
	PlayerKickOutScNotify                              = 80
	TrainPartyTakeBuildLevelAwardCsReq                 = 8036
	GetMissionDataCsReq                                = 1259
	DailyActiveInfoNotify                              = 3339
	QuitTreasureDungeonScRsp                           = 4473
	GetTreasureDungeonActivityDataCsReq                = 4447
	GetMonopolyDailyReportCsReq                        = 7076
	GetKilledPunkLordMonsterDataScRsp                  = 3283
	GeneralVirtualItemDataNotify                       = 543
	PlayerGetTokenCsReq                                = 39
	SyncRogueCommonVirtualItemInfoScNotify             = 5666
	EnteredSceneChangeScNotify                         = 1444
	SwordTrainingExamResultConfirmScRsp                = 7488
	SelectInclinationTextScRsp                         = 2116
	TakeAllRewardCsReq                                 = 3037
	GetChessRogueNousStoryInfoCsReq                    = 5486
	EvolveBuildTakeExpRewardScRsp                      = 7131
	DiscardRelicCsReq                                  = 563
	EnterFightActivityStageScRsp                       = 3639
	GetBattleCollegeDataCsReq                          = 5759
	ChessRogueNousEnableRogueTalentCsReq               = 5581
	LockEquipmentScRsp                                 = 553
	ResetMapRotationRegionCsReq                        = 6830
	ExchangeStaminaScRsp                               = 48
	GetSingleRedDotParamGroupScRsp                     = 5953
	EnterRogueEndlessActivityStageCsReq                = 6009
	TrainPartyGetDataCsReq                             = 8059
	TakeRogueEndlessActivityPointRewardScRsp           = 6010
	TakeQuestOptionalRewardScRsp                       = 930
	FinishFirstTalkByPerformanceNpcScRsp               = 2175
	ChessRogueEnterCsReq                               = 5502
	SubmitEmotionItemCsReq                             = 6339
	MonopolyAcceptQuizScRsp                            = 7091
	TextJoinSaveScRsp                                  = 3820
	GetFriendListInfoCsReq                             = 2959
	HeartDialScriptChangeScNotify                      = 6380
	MuseumTakeCollectRewardScRsp                       = 4364
	SelectChatBubbleScRsp                              = 5146
	RogueMagicScepterDressInUnitScRsp                  = 7799
	UpdateServerPrefsDataCsReq                         = 6139
	ChessRogueRollDiceCsReq                            = 5414
	GetTutorialGuideCsReq                              = 1603
	TakeKilledPunkLordMonsterScoreCsReq                = 3264
	UseTreasureDungeonItemScRsp                        = 4433
	AetherDivideSkillItemScNotify                      = 4894
	NewMailScNotify                                    = 880
	FinishTalkMissionCsReq                             = 1203
	MonopolyEventSelectFriendScRsp                     = 7082
	RogueTournEnterRoomScRsp                           = 6044
	DeactivateFarmElementScRsp                         = 1423
	FightHeartBeatCsReq                                = 30039
	RogueTournEnablePermanentTalentScRsp               = 6094
	GetStuffScNotify                                   = 4380
	UnlockTutorialGuideScRsp                           = 1637
	MissionAcceptScNotify                              = 1226
	SetHeadIconCsReq                                   = 2803
	TrainPartyUpdatePosEnvScRsp                        = 8070
	BuyRogueShopBuffCsReq                              = 5616
	RogueMagicEnterLayerScRsp                          = 7774
	GetChessRogueStoryInfoScRsp                        = 5407
	UpdateMapRotationDataScNotify                      = 6879
	ShowNewSupplementVisitorScRsp                      = 3774
	GetCurBattleInfoCsReq                              = 139
	GetStrongChallengeActivityDataScRsp                = 6620
	RogueMagicLeaveCsReq                               = 7739
	InteractPropCsReq                                  = 1403
	SwordTrainingLearnSkillScRsp                       = 7483
	MonopolyGuessBuyInformationScRsp                   = 7014
	RogueMagicLevelInfoUpdateScNotify                  = 7730
	RogueModifierSelectCellScRsp                       = 5339
	EnterSceneCsReq                                    = 1406
	TrainPartyEnterCsReq                               = 8061
	TravelBrochureSelectMessageScRsp                   = 6453
	ArchiveWolfBroGameCsReq                            = 6503
	FightFestUpdateCoinNotify                          = 7273
	GetSaveRaidScRsp                                   = 2290
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(SwordTrainingGameSyncChangeScNotify, func() any { return new(proto.SwordTrainingGameSyncChangeScNotify) })
	c.regMsg(DeleteRelicFilterPlanScRsp, func() any { return new(proto.DeleteRelicFilterPlanScRsp) })
	c.regMsg(SwordTrainingDialogueSelectOptionCsReq, func() any { return new(proto.SwordTrainingDialogueSelectOptionCsReq) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(ClockParkBattleEndScNotify, func() any { return new(proto.ClockParkBattleEndScNotify) })
	c.regMsg(RaidCollectionEnterNextRaidCsReq, func() any { return new(proto.RaidCollectionEnterNextRaidCsReq) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneCsReq) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(TakeMaterialSubmitActivityRewardScRsp, func() any { return new(proto.TakeMaterialSubmitActivityRewardScRsp) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(RogueMagicSetAutoDressInMagicUnitCsReq, func() any { return new(proto.RogueMagicSetAutoDressInMagicUnitCsReq) })
	c.regMsg(GetMultipleDropInfoCsReq, func() any { return new(proto.GetMultipleDropInfoCsReq) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(StartTrackPhotoStageCsReq, func() any { return new(proto.StartTrackPhotoStageCsReq) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(MultipleDropInfoScNotify, func() any { return new(proto.MultipleDropInfoScNotify) })
	c.regMsg(MusicRhythmSaveSongConfigDataScRsp, func() any { return new(proto.MusicRhythmSaveSongConfigDataScRsp) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(MusicRhythmUnlockTrackScNotify, func() any { return new(proto.MusicRhythmUnlockTrackScNotify) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(ContentPackageUnlockScRsp, func() any { return new(proto.ContentPackageUnlockScRsp) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(ClientObjUploadScRsp, func() any { return new(proto.ClientObjUploadScRsp) })
	c.regMsg(TrainPartyMoveScNotify, func() any { return new(proto.TrainPartyMoveScNotify) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(FinishRogueCommonDialogueScRsp, func() any { return new(proto.FinishRogueCommonDialogueScRsp) })
	c.regMsg(MarkAvatarScRsp, func() any { return new(proto.MarkAvatarScRsp) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(QuitBattleScRsp, func() any { return new(proto.QuitBattleScRsp) })
	c.regMsg(RogueGetGambleInfoScRsp, func() any { return new(proto.RogueGetGambleInfoScRsp) })
	c.regMsg(RogueTournDeleteArchiveCsReq, func() any { return new(proto.RogueTournDeleteArchiveCsReq) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(ChessRogueNousEnableRogueTalentScRsp, func() any { return new(proto.ChessRogueNousEnableRogueTalentScRsp) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(TrainPartyUpdatePosEnvCsReq, func() any { return new(proto.TrainPartyUpdatePosEnvCsReq) })
	c.regMsg(StartBoxingClubBattleCsReq, func() any { return new(proto.StartBoxingClubBattleCsReq) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoCsReq, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoCsReq) })
	c.regMsg(RogueTournLevelInfoUpdateScNotify, func() any { return new(proto.RogueTournLevelInfoUpdateScNotify) })
	c.regMsg(GetMultiPathAvatarInfoScRsp, func() any { return new(proto.GetMultiPathAvatarInfoScRsp) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(SelectRogueCommonDialogueOptionCsReq, func() any { return new(proto.SelectRogueCommonDialogueOptionCsReq) })
	c.regMsg(TakeMultipleActivityExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleActivityExpeditionRewardScRsp) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(TakeMultipleActivityExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleActivityExpeditionRewardCsReq) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(GetChallengeRecommendLineupListScRsp, func() any { return new(proto.GetChallengeRecommendLineupListScRsp) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(GetFriendAssistListCsReq, func() any { return new(proto.GetFriendAssistListCsReq) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(FightMatch3OpponentDataScNotify, func() any { return new(proto.FightMatch3OpponentDataScNotify) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(MonopolyLikeScRsp, func() any { return new(proto.MonopolyLikeScRsp) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(ClockParkStartScriptCsReq, func() any { return new(proto.ClockParkStartScriptCsReq) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(RemoveRotaterCsReq, func() any { return new(proto.RemoveRotaterCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(FightMatch3ChatScNotify, func() any { return new(proto.FightMatch3ChatScNotify) })
	c.regMsg(GetMapRotationDataCsReq, func() any { return new(proto.GetMapRotationDataCsReq) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(FightMatch3SwapCsReq, func() any { return new(proto.FightMatch3SwapCsReq) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(EvolveBuildTakeExpRewardCsReq, func() any { return new(proto.EvolveBuildTakeExpRewardCsReq) })
	c.regMsg(SetMultipleAvatarPathsCsReq, func() any { return new(proto.SetMultipleAvatarPathsCsReq) })
	c.regMsg(RogueWorkbenchGetInfoScRsp, func() any { return new(proto.RogueWorkbenchGetInfoScRsp) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(MatchThreeGetDataScRsp, func() any { return new(proto.MatchThreeGetDataScRsp) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(SyncRoguePickAvatarInfoScNotify, func() any { return new(proto.SyncRoguePickAvatarInfoScNotify) })
	c.regMsg(RogueTournGetPermanentTalentInfoScRsp, func() any { return new(proto.RogueTournGetPermanentTalentInfoScRsp) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(SpaceZooExchangeItemCsReq, func() any { return new(proto.SpaceZooExchangeItemCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(SwordTrainingUnlockSyncScNotify, func() any { return new(proto.SwordTrainingUnlockSyncScNotify) })
	c.regMsg(TrainPartyUnlockBuildAreaCsReq, func() any { return new(proto.TrainPartyUnlockBuildAreaCsReq) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(MonopolyGuessBuyInformationCsReq, func() any { return new(proto.MonopolyGuessBuyInformationCsReq) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(RogueTournConfirmSettleCsReq, func() any { return new(proto.RogueTournConfirmSettleCsReq) })
	c.regMsg(SetFriendMarkScRsp, func() any { return new(proto.SetFriendMarkScRsp) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(OfferingInfoScNotify, func() any { return new(proto.OfferingInfoScNotify) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(UpdatePsnSettingsInfoCsReq, func() any { return new(proto.UpdatePsnSettingsInfoCsReq) })
	c.regMsg(StarFightDataChangeNotify, func() any { return new(proto.StarFightDataChangeNotify) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(RogueMagicAreaUpdateScNotify, func() any { return new(proto.RogueMagicAreaUpdateScNotify) })
	c.regMsg(RogueTournExpNotify, func() any { return new(proto.RogueTournExpNotify) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(MonopolySttUpdateScNotify, func() any { return new(proto.MonopolySttUpdateScNotify) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(ContentPackageUnlockCsReq, func() any { return new(proto.ContentPackageUnlockCsReq) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(SetAvatarPathScRsp, func() any { return new(proto.SetAvatarPathScRsp) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(FightFestScoreUpdateNotify, func() any { return new(proto.FightFestScoreUpdateNotify) })
	c.regMsg(GetCrossInfoCsReq, func() any { return new(proto.GetCrossInfoCsReq) })
	c.regMsg(ClientDownloadDataScNotify, func() any { return new(proto.ClientDownloadDataScNotify) })
	c.regMsg(GetMaterialSubmitActivityDataScRsp, func() any { return new(proto.GetMaterialSubmitActivityDataScRsp) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(RogueArcadeRestartCsReq, func() any { return new(proto.RogueArcadeRestartCsReq) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(BuyRogueShopMiracleCsReq, func() any { return new(proto.BuyRogueShopMiracleCsReq) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(LobbyJoinCsReq, func() any { return new(proto.LobbyJoinCsReq) })
	c.regMsg(UpdateGunPlayDataScRsp, func() any { return new(proto.UpdateGunPlayDataScRsp) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(RogueMagicQueryCsReq, func() any { return new(proto.RogueMagicQueryCsReq) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(TrainPartyLeaveScRsp, func() any { return new(proto.TrainPartyLeaveScRsp) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(GetRogueExhibitionCsReq, func() any { return new(proto.GetRogueExhibitionCsReq) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(RaidCollectionDataScRsp, func() any { return new(proto.RaidCollectionDataScRsp) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(UpdateGroupPropertyScRsp, func() any { return new(proto.UpdateGroupPropertyScRsp) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(StartTimedCocoonStageCsReq, func() any { return new(proto.StartTimedCocoonStageCsReq) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownScRsp, func() any { return new(proto.EvolveBuildShopAbilityDownScRsp) })
	c.regMsg(GetMonopolyFriendRankingListCsReq, func() any { return new(proto.GetMonopolyFriendRankingListCsReq) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(SetSpringRecoverConfigCsReq, func() any { return new(proto.SetSpringRecoverConfigCsReq) })
	c.regMsg(MusicRhythmUnlockSongSfxScNotify, func() any { return new(proto.MusicRhythmUnlockSongSfxScNotify) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(FightKickOutScNotify, func() any { return new(proto.FightKickOutScNotify) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(RogueTournEnterLayerCsReq, func() any { return new(proto.RogueTournEnterLayerCsReq) })
	c.regMsg(GetGunPlayDataCsReq, func() any { return new(proto.GetGunPlayDataCsReq) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(TakeRollShopRewardCsReq, func() any { return new(proto.TakeRollShopRewardCsReq) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(LobbyCreateCsReq, func() any { return new(proto.LobbyCreateCsReq) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryCsReq, func() any { return new(proto.RogueTournGetArchiveRepositoryCsReq) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(ClockParkGetInfoCsReq, func() any { return new(proto.ClockParkGetInfoCsReq) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(SwordTrainingDailyPhaseConfirmScRsp, func() any { return new(proto.SwordTrainingDailyPhaseConfirmScRsp) })
	c.regMsg(SpringRecoverSingleAvatarCsReq, func() any { return new(proto.SpringRecoverSingleAvatarCsReq) })
	c.regMsg(MonopolyActionResultScNotify, func() any { return new(proto.MonopolyActionResultScNotify) })
	c.regMsg(GetRogueShopMiracleInfoCsReq, func() any { return new(proto.GetRogueShopMiracleInfoCsReq) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(MonopolyClickMbtiReportCsReq, func() any { return new(proto.MonopolyClickMbtiReportCsReq) })
	c.regMsg(RogueTournTakeExpRewardCsReq, func() any { return new(proto.RogueTournTakeExpRewardCsReq) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(GetSwordTrainingDataCsReq, func() any { return new(proto.GetSwordTrainingDataCsReq) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(SwordTrainingRestoreGameScRsp, func() any { return new(proto.SwordTrainingRestoreGameScRsp) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(MultiplayerMatch3FinishScNotify, func() any { return new(proto.MultiplayerMatch3FinishScNotify) })
	c.regMsg(GetStarFightDataScRsp, func() any { return new(proto.GetStarFightDataScRsp) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(EvolveBuildShopAbilityResetCsReq, func() any { return new(proto.EvolveBuildShopAbilityResetCsReq) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(QuitTrackPhotoStageScRsp, func() any { return new(proto.QuitTrackPhotoStageScRsp) })
	c.regMsg(MarkRelicFilterPlanCsReq, func() any { return new(proto.MarkRelicFilterPlanCsReq) })
	c.regMsg(RegionStopScNotify, func() any { return new(proto.RegionStopScNotify) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(BattleLogReportCsReq, func() any { return new(proto.BattleLogReportCsReq) })
	c.regMsg(MonopolyConditionUpdateScNotify, func() any { return new(proto.MonopolyConditionUpdateScNotify) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(RogueWorkbenchSelectFuncScRsp, func() any { return new(proto.RogueWorkbenchSelectFuncScRsp) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(UnlockAvatarPathCsReq, func() any { return new(proto.UnlockAvatarPathCsReq) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(RecallPetScRsp, func() any { return new(proto.RecallPetScRsp) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(RogueMagicGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueMagicGetMiscRealTimeDataScRsp) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(SelectRogueCommonDialogueOptionScRsp, func() any { return new(proto.SelectRogueCommonDialogueOptionScRsp) })
	c.regMsg(RogueTournEnablePermanentTalentCsReq, func() any { return new(proto.RogueTournEnablePermanentTalentCsReq) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(GetRogueCommonDialogueDataCsReq, func() any { return new(proto.GetRogueCommonDialogueDataCsReq) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(ClockParkHandleWaitOperationScRsp, func() any { return new(proto.ClockParkHandleWaitOperationScRsp) })
	c.regMsg(TakeFightActivityRewardScRsp, func() any { return new(proto.TakeFightActivityRewardScRsp) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(BoxingClubRewardScNotify, func() any { return new(proto.BoxingClubRewardScNotify) })
	c.regMsg(MatchThreeSyncDataScNotify, func() any { return new(proto.MatchThreeSyncDataScNotify) })
	c.regMsg(DailyTaskDataScNotify, func() any { return new(proto.DailyTaskDataScNotify) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(MatchResultScNotify, func() any { return new(proto.MatchResultScNotify) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(SyncRogueCommonDialogueDataScNotify, func() any { return new(proto.SyncRogueCommonDialogueDataScNotify) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(GetAuthkeyCsReq, func() any { return new(proto.GetAuthkeyCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(RogueTournReviveAvatarScRsp, func() any { return new(proto.RogueTournReviveAvatarScRsp) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(GetMonopolyInfoScRsp, func() any { return new(proto.GetMonopolyInfoScRsp) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(GetUpdatedArchiveDataCsReq, func() any { return new(proto.GetUpdatedArchiveDataCsReq) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(MonopolySelectOptionCsReq, func() any { return new(proto.MonopolySelectOptionCsReq) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(TrainPartyBuildDiyCsReq, func() any { return new(proto.TrainPartyBuildDiyCsReq) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(RogueWorkbenchHandleFuncCsReq, func() any { return new(proto.RogueWorkbenchHandleFuncCsReq) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(RogueArcadeGetInfoScRsp, func() any { return new(proto.RogueArcadeGetInfoScRsp) })
	c.regMsg(ClockParkUnlockTalentScRsp, func() any { return new(proto.ClockParkUnlockTalentScRsp) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(GetAuthkeyScRsp, func() any { return new(proto.GetAuthkeyScRsp) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(SpringRecoverCsReq, func() any { return new(proto.SpringRecoverCsReq) })
	c.regMsg(RaidCollectionDataScNotify, func() any { return new(proto.RaidCollectionDataScNotify) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(SwordTrainingSelectEndingCsReq, func() any { return new(proto.SwordTrainingSelectEndingCsReq) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(AcceptedPamMissionExpireCsReq, func() any { return new(proto.AcceptedPamMissionExpireCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(SwordTrainingResumeGameCsReq, func() any { return new(proto.SwordTrainingResumeGameCsReq) })
	c.regMsg(SpringRecoverScRsp, func() any { return new(proto.SpringRecoverScRsp) })
	c.regMsg(MonopolySelectOptionScRsp, func() any { return new(proto.MonopolySelectOptionScRsp) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(SetForbidOtherApplyFriendScRsp, func() any { return new(proto.SetForbidOtherApplyFriendScRsp) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(RogueTournEnterLayerScRsp, func() any { return new(proto.RogueTournEnterLayerScRsp) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(RogueMagicGetTalentInfoCsReq, func() any { return new(proto.RogueMagicGetTalentInfoCsReq) })
	c.regMsg(MultiplayerFightGiveUpScRsp, func() any { return new(proto.MultiplayerFightGiveUpScRsp) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(MonopolyGiveUpCurContentScRsp, func() any { return new(proto.MonopolyGiveUpCurContentScRsp) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(NewAssistHistoryNotify, func() any { return new(proto.NewAssistHistoryNotify) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(RogueMagicLeaveScRsp, func() any { return new(proto.RogueMagicLeaveScRsp) })
	c.regMsg(HealPoolInfoNotify, func() any { return new(proto.HealPoolInfoNotify) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(GetFriendBattleRecordDetailCsReq, func() any { return new(proto.GetFriendBattleRecordDetailCsReq) })
	c.regMsg(GetChallengeRecommendLineupListCsReq, func() any { return new(proto.GetChallengeRecommendLineupListCsReq) })
	c.regMsg(ActivityRaidPlacingGameScRsp, func() any { return new(proto.ActivityRaidPlacingGameScRsp) })
	c.regMsg(SyncRogueCommonDialogueOptionFinishScNotify, func() any { return new(proto.SyncRogueCommonDialogueOptionFinishScNotify) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(CancelMatchCsReq, func() any { return new(proto.CancelMatchCsReq) })
	c.regMsg(MonopolyAcceptQuizCsReq, func() any { return new(proto.MonopolyAcceptQuizCsReq) })
	c.regMsg(MultiplayerFightGameStateCsReq, func() any { return new(proto.MultiplayerFightGameStateCsReq) })
	c.regMsg(MuseumTargetRewardNotify, func() any { return new(proto.MuseumTargetRewardNotify) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(SettleTrackPhotoStageScRsp, func() any { return new(proto.SettleTrackPhotoStageScRsp) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(GetPlayerReturnMultiDropInfoCsReq, func() any { return new(proto.GetPlayerReturnMultiDropInfoCsReq) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownCsReq, func() any { return new(proto.EvolveBuildShopAbilityDownCsReq) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(RogueMagicGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueMagicGetMiscRealTimeDataCsReq) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(WolfBroGameActivateBulletScRsp, func() any { return new(proto.WolfBroGameActivateBulletScRsp) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(RogueMagicBattleFailSettleInfoScNotify, func() any { return new(proto.RogueMagicBattleFailSettleInfoScNotify) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(DifficultyAdjustmentUpdateDataScRsp, func() any { return new(proto.DifficultyAdjustmentUpdateDataScRsp) })
	c.regMsg(GetLevelRewardScRsp, func() any { return new(proto.GetLevelRewardScRsp) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(ClientObjUploadCsReq, func() any { return new(proto.ClientObjUploadCsReq) })
	c.regMsg(EvolveBuildQueryInfoScRsp, func() any { return new(proto.EvolveBuildQueryInfoScRsp) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(GetSpringRecoverDataScRsp, func() any { return new(proto.GetSpringRecoverDataScRsp) })
	c.regMsg(TrainPartyGamePlayStartCsReq, func() any { return new(proto.TrainPartyGamePlayStartCsReq) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(GetRogueCollectionCsReq, func() any { return new(proto.GetRogueCollectionCsReq) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(MonopolyQuizDurationChangeScNotify, func() any { return new(proto.MonopolyQuizDurationChangeScNotify) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(MultipleDropInfoNotify, func() any { return new(proto.MultipleDropInfoNotify) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(LobbyInviteScNotify, func() any { return new(proto.LobbyInviteScNotify) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(EvolveBuildStartLevelScRsp, func() any { return new(proto.EvolveBuildStartLevelScRsp) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(GetRogueExhibitionScRsp, func() any { return new(proto.GetRogueExhibitionScRsp) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(ClockParkGetOngoingScriptInfoCsReq, func() any { return new(proto.ClockParkGetOngoingScriptInfoCsReq) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(SwordTrainingMarkEndingViewedCsReq, func() any { return new(proto.SwordTrainingMarkEndingViewedCsReq) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(CancelExpeditionScRsp, func() any { return new(proto.CancelExpeditionScRsp) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(ServerSimulateBattleFinishScNotify, func() any { return new(proto.ServerSimulateBattleFinishScNotify) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(RogueArcadeLeaveScRsp, func() any { return new(proto.RogueArcadeLeaveScRsp) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(RogueTournGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueTournGetMiscRealTimeDataScRsp) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(RogueTournClearArchiveNameScNotify, func() any { return new(proto.RogueTournClearArchiveNameScNotify) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(RogueTournBattleFailSettleInfoScNotify, func() any { return new(proto.RogueTournBattleFailSettleInfoScNotify) })
	c.regMsg(SwordTrainingStoryBattleScRsp, func() any { return new(proto.SwordTrainingStoryBattleScRsp) })
	c.regMsg(RebattleByClientCsNotify, func() any { return new(proto.RebattleByClientCsNotify) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(MuseumTargetMissionFinishNotify, func() any { return new(proto.MuseumTargetMissionFinishNotify) })
	c.regMsg(MusicRhythmUnlockSongNotify, func() any { return new(proto.MusicRhythmUnlockSongNotify) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(GetMbtiReportScRsp, func() any { return new(proto.GetMbtiReportScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(StartMatchCsReq, func() any { return new(proto.StartMatchCsReq) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(RogueTournRenameArchiveScRsp, func() any { return new(proto.RogueTournRenameArchiveScRsp) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(UpdatePlayerSettingScRsp, func() any { return new(proto.UpdatePlayerSettingScRsp) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(SwordTrainingDailyPhaseConfirmCsReq, func() any { return new(proto.SwordTrainingDailyPhaseConfirmCsReq) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(SetRogueExhibitionCsReq, func() any { return new(proto.SetRogueExhibitionCsReq) })
	c.regMsg(ClockParkStartScriptScRsp, func() any { return new(proto.ClockParkStartScriptScRsp) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(SwordTrainingGiveUpGameCsReq, func() any { return new(proto.SwordTrainingGiveUpGameCsReq) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(TrainPartyUseCardScRsp, func() any { return new(proto.TrainPartyUseCardScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(RogueTournDeleteArchiveScRsp, func() any { return new(proto.RogueTournDeleteArchiveScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(UpdateGunPlayDataCsReq, func() any { return new(proto.UpdateGunPlayDataCsReq) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(RogueMagicEnterRoomCsReq, func() any { return new(proto.RogueMagicEnterRoomCsReq) })
	c.regMsg(TrainPartyBuildStartStepScRsp, func() any { return new(proto.TrainPartyBuildStartStepScRsp) })
	c.regMsg(TrainPartyBuildingUpdateNotify, func() any { return new(proto.TrainPartyBuildingUpdateNotify) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(RogueArcadeStartCsReq, func() any { return new(proto.RogueArcadeStartCsReq) })
	c.regMsg(TrainPartyGetDataScRsp, func() any { return new(proto.TrainPartyGetDataScRsp) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(RogueTournResetPermanentTalentScRsp, func() any { return new(proto.RogueTournResetPermanentTalentScRsp) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(SetRogueCollectionScRsp, func() any { return new(proto.SetRogueCollectionScRsp) })
	c.regMsg(EvolveBuildShopAbilityUpScRsp, func() any { return new(proto.EvolveBuildShopAbilityUpScRsp) })
	c.regMsg(TrainPartyGamePlaySettleNotify, func() any { return new(proto.TrainPartyGamePlaySettleNotify) })
	c.regMsg(RogueTournWeekChallengeUpdateScNotify, func() any { return new(proto.RogueTournWeekChallengeUpdateScNotify) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(RogueModifierDelNotify, func() any { return new(proto.RogueModifierDelNotify) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(EvolveBuildUnlockInfoNotify, func() any { return new(proto.EvolveBuildUnlockInfoNotify) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(MonopolyGameSettleScNotify, func() any { return new(proto.MonopolyGameSettleScNotify) })
	c.regMsg(GetMbtiReportCsReq, func() any { return new(proto.GetMbtiReportCsReq) })
	c.regMsg(MonopolyClickCellCsReq, func() any { return new(proto.MonopolyClickCellCsReq) })
	c.regMsg(GetRogueShopBuffInfoCsReq, func() any { return new(proto.GetRogueShopBuffInfoCsReq) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneScRsp) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(MusicRhythmSaveSongConfigDataCsReq, func() any { return new(proto.MusicRhythmSaveSongConfigDataCsReq) })
	c.regMsg(SubmitOfferingItemCsReq, func() any { return new(proto.SubmitOfferingItemCsReq) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(DifficultyAdjustmentGetDataScRsp, func() any { return new(proto.DifficultyAdjustmentGetDataScRsp) })
	c.regMsg(SwordTrainingSetSkillTraceCsReq, func() any { return new(proto.SwordTrainingSetSkillTraceCsReq) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(FightMatch3TurnStartScNotify, func() any { return new(proto.FightMatch3TurnStartScNotify) })
	c.regMsg(RogueMagicReviveCostUpdateScNotify, func() any { return new(proto.RogueMagicReviveCostUpdateScNotify) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(SwordTrainingExamResultConfirmCsReq, func() any { return new(proto.SwordTrainingExamResultConfirmCsReq) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(MusicRhythmDataScRsp, func() any { return new(proto.MusicRhythmDataScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoCsReq, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoCsReq) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(TakeAssistRewardCsReq, func() any { return new(proto.TakeAssistRewardCsReq) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(FinishedMissionScNotify, func() any { return new(proto.FinishedMissionScNotify) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(RogueTournGetSettleInfoScRsp, func() any { return new(proto.RogueTournGetSettleInfoScRsp) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(UpdatePlayerSettingCsReq, func() any { return new(proto.UpdatePlayerSettingCsReq) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(GetFriendAssistListScRsp, func() any { return new(proto.GetFriendAssistListScRsp) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(GetGameStateServiceConfigScRsp, func() any { return new(proto.GetGameStateServiceConfigScRsp) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(GetPamSkinDataScRsp, func() any { return new(proto.GetPamSkinDataScRsp) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(TakeRogueEventHandbookRewardScRsp, func() any { return new(proto.TakeRogueEventHandbookRewardScRsp) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(GetMonopolyInfoCsReq, func() any { return new(proto.GetMonopolyInfoCsReq) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(RogueTournGetAllArchiveScRsp, func() any { return new(proto.RogueTournGetAllArchiveScRsp) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(EvolveBuildFinishScNotify, func() any { return new(proto.EvolveBuildFinishScNotify) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(TrainPartyBuildDiyScRsp, func() any { return new(proto.TrainPartyBuildDiyScRsp) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(MultiplayerFightGiveUpCsReq, func() any { return new(proto.MultiplayerFightGiveUpCsReq) })
	c.regMsg(RogueMagicStoryInfoUpdateScNotify, func() any { return new(proto.RogueMagicStoryInfoUpdateScNotify) })
	c.regMsg(UpdateGroupPropertyCsReq, func() any { return new(proto.UpdateGroupPropertyCsReq) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(RogueTournHandBookNotify, func() any { return new(proto.RogueTournHandBookNotify) })
	c.regMsg(ClockParkUseBuffCsReq, func() any { return new(proto.ClockParkUseBuffCsReq) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoCsReq, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoCsReq) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(MonopolyCheatDiceScRsp, func() any { return new(proto.MonopolyCheatDiceScRsp) })
	c.regMsg(RogueArcadeLeaveCsReq, func() any { return new(proto.RogueArcadeLeaveCsReq) })
	c.regMsg(RogueMagicAutoDressInUnitCsReq, func() any { return new(proto.RogueMagicAutoDressInUnitCsReq) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(GetMultipleDropInfoScRsp, func() any { return new(proto.GetMultipleDropInfoScRsp) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(FightMatch3ForceUpdateNotify, func() any { return new(proto.FightMatch3ForceUpdateNotify) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(GetStageLineupCsReq, func() any { return new(proto.GetStageLineupCsReq) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(SubmitMaterialSubmitActivityMaterialScRsp, func() any { return new(proto.SubmitMaterialSubmitActivityMaterialScRsp) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(SceneReviveAfterRebattleCsReq, func() any { return new(proto.SceneReviveAfterRebattleCsReq) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(RogueMagicQueryScRsp, func() any { return new(proto.RogueMagicQueryScRsp) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(RogueDoGambleCsReq, func() any { return new(proto.RogueDoGambleCsReq) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(MonopolyGuessDrawScNotify, func() any { return new(proto.MonopolyGuessDrawScNotify) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(MonopolyDailySettleScNotify, func() any { return new(proto.MonopolyDailySettleScNotify) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(SetAvatarPathCsReq, func() any { return new(proto.SetAvatarPathCsReq) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(RogueMagicSettleCsReq, func() any { return new(proto.RogueMagicSettleCsReq) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(ChangeStoryLineFinishScNotify, func() any { return new(proto.ChangeStoryLineFinishScNotify) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(ClockParkQuitScriptScRsp, func() any { return new(proto.ClockParkQuitScriptScRsp) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(RogueTournDifficultyCompNotify, func() any { return new(proto.RogueTournDifficultyCompNotify) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(GetBoxingClubInfoCsReq, func() any { return new(proto.GetBoxingClubInfoCsReq) })
	c.regMsg(RogueTournGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueTournGetMiscRealTimeDataCsReq) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(MatchThreeSetBirdPosCsReq, func() any { return new(proto.MatchThreeSetBirdPosCsReq) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(GetPamSkinDataCsReq, func() any { return new(proto.GetPamSkinDataCsReq) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(QuitBattleCsReq, func() any { return new(proto.QuitBattleCsReq) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(RogueTournGetSettleInfoCsReq, func() any { return new(proto.RogueTournGetSettleInfoCsReq) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(SetFriendMarkCsReq, func() any { return new(proto.SetFriendMarkCsReq) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(WolfBroGameExplodeMonsterCsReq, func() any { return new(proto.WolfBroGameExplodeMonsterCsReq) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(GetOfferingInfoCsReq, func() any { return new(proto.GetOfferingInfoCsReq) })
	c.regMsg(AcceptMissionEventCsReq, func() any { return new(proto.AcceptMissionEventCsReq) })
	c.regMsg(DrinkMakerUpdateTipsNotify, func() any { return new(proto.DrinkMakerUpdateTipsNotify) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(StrongChallengeActivityBattleEndScNotify, func() any { return new(proto.StrongChallengeActivityBattleEndScNotify) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(SwordTrainingGameSettleScNotify, func() any { return new(proto.SwordTrainingGameSettleScNotify) })
	c.regMsg(TrainPartyAddBuildDynamicBuffCsReq, func() any { return new(proto.TrainPartyAddBuildDynamicBuffCsReq) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(MonopolyGetDailyInitItemCsReq, func() any { return new(proto.MonopolyGetDailyInitItemCsReq) })
	c.regMsg(StartFightFestCsReq, func() any { return new(proto.StartFightFestCsReq) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(ClockParkFinishScriptScNotify, func() any { return new(proto.ClockParkFinishScriptScNotify) })
	c.regMsg(MonopolyEventLoadUpdateScNotify, func() any { return new(proto.MonopolyEventLoadUpdateScNotify) })
	c.regMsg(RogueMagicSettleScRsp, func() any { return new(proto.RogueMagicSettleScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
	c.regMsg(FinishRogueCommonDialogueCsReq, func() any { return new(proto.FinishRogueCommonDialogueCsReq) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(RelicFilterPlanClearNameScNotify, func() any { return new(proto.RelicFilterPlanClearNameScNotify) })
	c.regMsg(FightMatch3ChatCsReq, func() any { return new(proto.FightMatch3ChatCsReq) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(MultiplayerFightGameStateScRsp, func() any { return new(proto.MultiplayerFightGameStateScRsp) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(HeliobusStartRaidCsReq, func() any { return new(proto.HeliobusStartRaidCsReq) })
	c.regMsg(RogueTournSettleCsReq, func() any { return new(proto.RogueTournSettleCsReq) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(RogueMagicUnitReforgeScRsp, func() any { return new(proto.RogueMagicUnitReforgeScRsp) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(EvolveBuildGiveupScRsp, func() any { return new(proto.EvolveBuildGiveupScRsp) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(MonopolyGameRaiseRatioCsReq, func() any { return new(proto.MonopolyGameRaiseRatioCsReq) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(GetTrainVisitorBehaviorCsReq, func() any { return new(proto.GetTrainVisitorBehaviorCsReq) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(ModifyRelicFilterPlanCsReq, func() any { return new(proto.ModifyRelicFilterPlanCsReq) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(SelectPamSkinScRsp, func() any { return new(proto.SelectPamSkinScRsp) })
	c.regMsg(SpaceZooExchangeItemScRsp, func() any { return new(proto.SpaceZooExchangeItemScRsp) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(RogueTournStartScRsp, func() any { return new(proto.RogueTournStartScRsp) })
	c.regMsg(RemoveRotaterScRsp, func() any { return new(proto.RemoveRotaterScRsp) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(EvolveBuildLeaveCsReq, func() any { return new(proto.EvolveBuildLeaveCsReq) })
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(RogueTournGetArchiveRepositoryScRsp, func() any { return new(proto.RogueTournGetArchiveRepositoryScRsp) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(MarkRelicFilterPlanScRsp, func() any { return new(proto.MarkRelicFilterPlanScRsp) })
	c.regMsg(LobbyKickOutScRsp, func() any { return new(proto.LobbyKickOutScRsp) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(SetRogueExhibitionScRsp, func() any { return new(proto.SetRogueExhibitionScRsp) })
	c.regMsg(RogueTournReviveAvatarCsReq, func() any { return new(proto.RogueTournReviveAvatarCsReq) })
	c.regMsg(RogueTournQueryScRsp, func() any { return new(proto.RogueTournQueryScRsp) })
	c.regMsg(UnlockAvatarPathScRsp, func() any { return new(proto.UnlockAvatarPathScRsp) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(MonopolyGetDailyInitItemScRsp, func() any { return new(proto.MonopolyGetDailyInitItemScRsp) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(SwordTrainingLearnSkillCsReq, func() any { return new(proto.SwordTrainingLearnSkillCsReq) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(EvolveBuildReRandomStageCsReq, func() any { return new(proto.EvolveBuildReRandomStageCsReq) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(RelicAvatarRecommendScRsp, func() any { return new(proto.RelicAvatarRecommendScRsp) })
	c.regMsg(ExchangeStaminaCsReq, func() any { return new(proto.ExchangeStaminaCsReq) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(SwordTrainingSetSkillTraceScRsp, func() any { return new(proto.SwordTrainingSetSkillTraceScRsp) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(AddEquipmentScNotify, func() any { return new(proto.AddEquipmentScNotify) })
	c.regMsg(MultiplayerGetFightGateCsReq, func() any { return new(proto.MultiplayerGetFightGateCsReq) })
	c.regMsg(SwordTrainingMarkEndingViewedScRsp, func() any { return new(proto.SwordTrainingMarkEndingViewedScRsp) })
	c.regMsg(GetFriendDevelopmentInfoScRsp, func() any { return new(proto.GetFriendDevelopmentInfoScRsp) })
	c.regMsg(EnterSwordTrainingExamScRsp, func() any { return new(proto.EnterSwordTrainingExamScRsp) })
	c.regMsg(MonopolyClickMbtiReportScRsp, func() any { return new(proto.MonopolyClickMbtiReportScRsp) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(MonopolyGameCreateScNotify, func() any { return new(proto.MonopolyGameCreateScNotify) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(DrinkMakerChallengeCsReq, func() any { return new(proto.DrinkMakerChallengeCsReq) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(CurPetChangedScNotify, func() any { return new(proto.CurPetChangedScNotify) })
	c.regMsg(RogueTournQueryCsReq, func() any { return new(proto.RogueTournQueryCsReq) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(RelicAvatarRecommendCsReq, func() any { return new(proto.RelicAvatarRecommendCsReq) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(TakeAssistRewardScRsp, func() any { return new(proto.TakeAssistRewardScRsp) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(MonopolyMoveCsReq, func() any { return new(proto.MonopolyMoveCsReq) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(SceneReviveAfterRebattleScRsp, func() any { return new(proto.SceneReviveAfterRebattleScRsp) })
	c.regMsg(GetStarFightDataCsReq, func() any { return new(proto.GetStarFightDataCsReq) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(MatchThreeSetBirdPosScRsp, func() any { return new(proto.MatchThreeSetBirdPosScRsp) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(ContentPackageTransferScNotify, func() any { return new(proto.ContentPackageTransferScNotify) })
	c.regMsg(RogueTournLeaveCsReq, func() any { return new(proto.RogueTournLeaveCsReq) })
	c.regMsg(SwordTrainingTurnActionCsReq, func() any { return new(proto.SwordTrainingTurnActionCsReq) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(RecallPetCsReq, func() any { return new(proto.RecallPetCsReq) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(CommonRogueComponentUpdateScNotify, func() any { return new(proto.CommonRogueComponentUpdateScNotify) })
	c.regMsg(TrainPartyEnterScRsp, func() any { return new(proto.TrainPartyEnterScRsp) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(SyncRogueAeonScNotify, func() any { return new(proto.SyncRogueAeonScNotify) })
	c.regMsg(RogueTournEnterCsReq, func() any { return new(proto.RogueTournEnterCsReq) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(EnterSummonActivityStageCsReq, func() any { return new(proto.EnterSummonActivityStageCsReq) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(GetRelicFilterPlanScRsp, func() any { return new(proto.GetRelicFilterPlanScRsp) })
	c.regMsg(LobbyModifyPlayerInfoCsReq, func() any { return new(proto.LobbyModifyPlayerInfoCsReq) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(RogueDoGambleScRsp, func() any { return new(proto.RogueDoGambleScRsp) })
	c.regMsg(RogueMagicEnterLayerCsReq, func() any { return new(proto.RogueMagicEnterLayerCsReq) })
	c.regMsg(RogueTournResetPermanentTalentCsReq, func() any { return new(proto.RogueTournResetPermanentTalentCsReq) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(AceAntiCheaterCsReq, func() any { return new(proto.AceAntiCheaterCsReq) })
	c.regMsg(CommonRogueUpdateScNotify, func() any { return new(proto.CommonRogueUpdateScNotify) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(RogueMagicUnitComposeCsReq, func() any { return new(proto.RogueMagicUnitComposeCsReq) })
	c.regMsg(StartTimedFarmElementCsReq, func() any { return new(proto.StartTimedFarmElementCsReq) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(GetStoryLineInfoCsReq, func() any { return new(proto.GetStoryLineInfoCsReq) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoScRsp, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoScRsp) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(SetRogueCollectionCsReq, func() any { return new(proto.SetRogueCollectionCsReq) })
	c.regMsg(RogueTournEnterRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournEnterRogueCocoonSceneScRsp) })
	c.regMsg(MultiplayerFightGameStartScNotify, func() any { return new(proto.MultiplayerFightGameStartScNotify) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(RaidCollectionEnterNextRaidScRsp, func() any { return new(proto.RaidCollectionEnterNextRaidScRsp) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(RogueTournTakeExpRewardScRsp, func() any { return new(proto.RogueTournTakeExpRewardScRsp) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(EvolveBuildShopAbilityResetScRsp, func() any { return new(proto.EvolveBuildShopAbilityResetScRsp) })
	c.regMsg(TakeFightActivityRewardCsReq, func() any { return new(proto.TakeFightActivityRewardCsReq) })
	c.regMsg(FightHeartBeatScRsp, func() any { return new(proto.FightHeartBeatScRsp) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(MatchThreeGetDataCsReq, func() any { return new(proto.MatchThreeGetDataCsReq) })
	c.regMsg(RestartChallengePhaseScRsp, func() any { return new(proto.RestartChallengePhaseScRsp) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(SummonPetCsReq, func() any { return new(proto.SummonPetCsReq) })
	c.regMsg(StoryLineInfoScNotify, func() any { return new(proto.StoryLineInfoScNotify) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(UpdateMovieRacingDataScRsp, func() any { return new(proto.UpdateMovieRacingDataScRsp) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(SummonPetScRsp, func() any { return new(proto.SummonPetScRsp) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(FightMatch3DataCsReq, func() any { return new(proto.FightMatch3DataCsReq) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(EvolveBuildQueryInfoCsReq, func() any { return new(proto.EvolveBuildQueryInfoCsReq) })
	c.regMsg(GetPetDataScRsp, func() any { return new(proto.GetPetDataScRsp) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(GetTrackPhotoActivityDataCsReq, func() any { return new(proto.GetTrackPhotoActivityDataCsReq) })
	c.regMsg(FightFestUnlockSkillNotify, func() any { return new(proto.FightFestUnlockSkillNotify) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(GetGunPlayDataScRsp, func() any { return new(proto.GetGunPlayDataScRsp) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(ReBattleAfterBattleLoseCsNotify, func() any { return new(proto.ReBattleAfterBattleLoseCsNotify) })
	c.regMsg(RogueTournSettleScRsp, func() any { return new(proto.RogueTournSettleScRsp) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(HeliobusStartRaidScRsp, func() any { return new(proto.HeliobusStartRaidScRsp) })
	c.regMsg(RogueMagicUnitReforgeCsReq, func() any { return new(proto.RogueMagicUnitReforgeCsReq) })
	c.regMsg(GetSummonActivityDataCsReq, func() any { return new(proto.GetSummonActivityDataCsReq) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(ClockParkUnlockTalentCsReq, func() any { return new(proto.ClockParkUnlockTalentCsReq) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(LobbyCreateScRsp, func() any { return new(proto.LobbyCreateScRsp) })
	c.regMsg(MonopolyGameBingoFlipCardScRsp, func() any { return new(proto.MonopolyGameBingoFlipCardScRsp) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(MarkAvatarCsReq, func() any { return new(proto.MarkAvatarCsReq) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(FightMatch3TurnEndScNotify, func() any { return new(proto.FightMatch3TurnEndScNotify) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(FightMatch3ChatScRsp, func() any { return new(proto.FightMatch3ChatScRsp) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(RogueTournLeaveScRsp, func() any { return new(proto.RogueTournLeaveScRsp) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(MonopolyGameGachaCsReq, func() any { return new(proto.MonopolyGameGachaCsReq) })
	c.regMsg(MonopolyLikeScNotify, func() any { return new(proto.MonopolyLikeScNotify) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoScRsp, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoScRsp) })
	c.regMsg(StartStarFightLevelScRsp, func() any { return new(proto.StartStarFightLevelScRsp) })
	c.regMsg(FightMatch3SwapScRsp, func() any { return new(proto.FightMatch3SwapScRsp) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(MonopolyGiveUpCurContentCsReq, func() any { return new(proto.MonopolyGiveUpCurContentCsReq) })
	c.regMsg(MonopolyMoveScRsp, func() any { return new(proto.MonopolyMoveScRsp) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(FightMatch3DataScRsp, func() any { return new(proto.FightMatch3DataScRsp) })
	c.regMsg(LobbyJoinScRsp, func() any { return new(proto.LobbyJoinScRsp) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(MultiplayerGetFightGateScRsp, func() any { return new(proto.MultiplayerGetFightGateScRsp) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(TrainVisitorBehaviorFinishCsReq, func() any { return new(proto.TrainVisitorBehaviorFinishCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(RogueTournEnterRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournEnterRogueCocoonSceneCsReq) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(SubmitOrigamiItemScRsp, func() any { return new(proto.SubmitOrigamiItemScRsp) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(MonopolyEventSelectFriendCsReq, func() any { return new(proto.MonopolyEventSelectFriendCsReq) })
	c.regMsg(GiveUpBoxingClubChallengeCsReq, func() any { return new(proto.GiveUpBoxingClubChallengeCsReq) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(StartTimedCocoonStageScRsp, func() any { return new(proto.StartTimedCocoonStageScRsp) })
	c.regMsg(RogueDebugMessageScNotify, func() any { return new(proto.RogueDebugMessageScNotify) })
	c.regMsg(GetRogueCollectionScRsp, func() any { return new(proto.GetRogueCollectionScRsp) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(FightFestUpdateChallengeRecordNotify, func() any { return new(proto.FightFestUpdateChallengeRecordNotify) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(EvolveBuildStartStageCsReq, func() any { return new(proto.EvolveBuildStartStageCsReq) })
	c.regMsg(TrainPartyLeaveCsReq, func() any { return new(proto.TrainPartyLeaveCsReq) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(SwordTrainingResumeGameScRsp, func() any { return new(proto.SwordTrainingResumeGameScRsp) })
	c.regMsg(ClockParkGetOngoingScriptInfoScRsp, func() any { return new(proto.ClockParkGetOngoingScriptInfoScRsp) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(FightEnterCsReq, func() any { return new(proto.FightEnterCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(MusicRhythmFinishLevelCsReq, func() any { return new(proto.MusicRhythmFinishLevelCsReq) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(GetRogueCommonDialogueDataScRsp, func() any { return new(proto.GetRogueCommonDialogueDataScRsp) })
	c.regMsg(SpringRefreshScRsp, func() any { return new(proto.SpringRefreshScRsp) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(MonopolyClickCellScRsp, func() any { return new(proto.MonopolyClickCellScRsp) })
	c.regMsg(MultiplayerFightGameFinishScNotify, func() any { return new(proto.MultiplayerFightGameFinishScNotify) })
	c.regMsg(MissionEventRewardScNotify, func() any { return new(proto.MissionEventRewardScNotify) })
	c.regMsg(RogueMagicReviveAvatarCsReq, func() any { return new(proto.RogueMagicReviveAvatarCsReq) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(StartFightFestScRsp, func() any { return new(proto.StartFightFestScRsp) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(SpringRecoverSingleAvatarScRsp, func() any { return new(proto.SpringRecoverSingleAvatarScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(VirtualLineupTrialAvatarChangeScNotify, func() any { return new(proto.VirtualLineupTrialAvatarChangeScNotify) })
	c.regMsg(LobbyBeginScRsp, func() any { return new(proto.LobbyBeginScRsp) })
	c.regMsg(RogueMagicUnitComposeScRsp, func() any { return new(proto.RogueMagicUnitComposeScRsp) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(SpaceZooTakeScRsp, func() any { return new(proto.SpaceZooTakeScRsp) })
	c.regMsg(MuseumTakeCollectRewardCsReq, func() any { return new(proto.MuseumTakeCollectRewardCsReq) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(MonopolyGameBingoFlipCardCsReq, func() any { return new(proto.MonopolyGameBingoFlipCardCsReq) })
	c.regMsg(RogueMagicScepterTakeOffUnitScRsp, func() any { return new(proto.RogueMagicScepterTakeOffUnitScRsp) })
	c.regMsg(RogueMagicReviveAvatarScRsp, func() any { return new(proto.RogueMagicReviveAvatarScRsp) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(MusicRhythmFinishLevelScRsp, func() any { return new(proto.MusicRhythmFinishLevelScRsp) })
	c.regMsg(MusicRhythmStartLevelCsReq, func() any { return new(proto.MusicRhythmStartLevelCsReq) })
	c.regMsg(RogueWorkbenchGetInfoCsReq, func() any { return new(proto.RogueWorkbenchGetInfoCsReq) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(RogueTournGetPermanentTalentInfoCsReq, func() any { return new(proto.RogueTournGetPermanentTalentInfoCsReq) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(ActivityRaidPlacingGameCsReq, func() any { return new(proto.ActivityRaidPlacingGameCsReq) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(CancelMatchScRsp, func() any { return new(proto.CancelMatchScRsp) })
	c.regMsg(StoryLineTrialAvatarChangeScNotify, func() any { return new(proto.StoryLineTrialAvatarChangeScNotify) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(RogueMagicStartScRsp, func() any { return new(proto.RogueMagicStartScRsp) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(EnterChallengeNextPhaseCsReq, func() any { return new(proto.EnterChallengeNextPhaseCsReq) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(MonopolyLikeCsReq, func() any { return new(proto.MonopolyLikeCsReq) })
	c.regMsg(DailyFirstEnterMonopolyActivityScRsp, func() any { return new(proto.DailyFirstEnterMonopolyActivityScRsp) })
	c.regMsg(RogueMagicEnableTalentScRsp, func() any { return new(proto.RogueMagicEnableTalentScRsp) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(LobbyModifyPlayerInfoScRsp, func() any { return new(proto.LobbyModifyPlayerInfoScRsp) })
	c.regMsg(GetMultiPathAvatarInfoCsReq, func() any { return new(proto.GetMultiPathAvatarInfoCsReq) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(TrainPartyHandlePendingActionScRsp, func() any { return new(proto.TrainPartyHandlePendingActionScRsp) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(GetChallengeGroupStatisticsScRsp, func() any { return new(proto.GetChallengeGroupStatisticsScRsp) })
	c.regMsg(RogueGetGambleInfoCsReq, func() any { return new(proto.RogueGetGambleInfoCsReq) })
	c.regMsg(RogueMagicGetTalentInfoScRsp, func() any { return new(proto.RogueMagicGetTalentInfoScRsp) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(InterruptMissionEventCsReq, func() any { return new(proto.InterruptMissionEventCsReq) })
	c.regMsg(UnlockPamSkinScNotify, func() any { return new(proto.UnlockPamSkinScNotify) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(RogueTournReEnterRogueCocoonStageCsReq, func() any { return new(proto.RogueTournReEnterRogueCocoonStageCsReq) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(DrinkMakerChallengeScRsp, func() any { return new(proto.DrinkMakerChallengeScRsp) })
	c.regMsg(RogueWorkbenchSelectFuncCsReq, func() any { return new(proto.RogueWorkbenchSelectFuncCsReq) })
	c.regMsg(RogueWorkbenchHandleFuncScRsp, func() any { return new(proto.RogueWorkbenchHandleFuncScRsp) })
	c.regMsg(SpringRefreshCsReq, func() any { return new(proto.SpringRefreshCsReq) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(RogueMagicEnterScRsp, func() any { return new(proto.RogueMagicEnterScRsp) })
	c.regMsg(StartStarFightLevelCsReq, func() any { return new(proto.StartStarFightLevelCsReq) })
	c.regMsg(TrainPartyUnlockBuildAreaScRsp, func() any { return new(proto.TrainPartyUnlockBuildAreaScRsp) })
	c.regMsg(LobbyBeginCsReq, func() any { return new(proto.LobbyBeginCsReq) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffCsReq, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffCsReq) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(LobbyQuitCsReq, func() any { return new(proto.LobbyQuitCsReq) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(RogueTournAreaUpdateScNotify, func() any { return new(proto.RogueTournAreaUpdateScNotify) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(UpdateMovieRacingDataCsReq, func() any { return new(proto.UpdateMovieRacingDataCsReq) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(MonopolyGuessChooseScRsp, func() any { return new(proto.MonopolyGuessChooseScRsp) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
	c.regMsg(LobbySyncInfoScNotify, func() any { return new(proto.LobbySyncInfoScNotify) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(GetFriendDevelopmentInfoCsReq, func() any { return new(proto.GetFriendDevelopmentInfoCsReq) })
	c.regMsg(GetRelicFilterPlanCsReq, func() any { return new(proto.GetRelicFilterPlanCsReq) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(RogueTournReEnterRogueCocoonStageScRsp, func() any { return new(proto.RogueTournReEnterRogueCocoonStageScRsp) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(DifficultyAdjustmentUpdateDataCsReq, func() any { return new(proto.DifficultyAdjustmentUpdateDataCsReq) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(TrainPartySettleNotify, func() any { return new(proto.TrainPartySettleNotify) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(LobbyQuitScRsp, func() any { return new(proto.LobbyQuitScRsp) })
	c.regMsg(LobbyGetInfoCsReq, func() any { return new(proto.LobbyGetInfoCsReq) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(MonopolyGuessChooseCsReq, func() any { return new(proto.MonopolyGuessChooseCsReq) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(RogueMagicEnterRoomScRsp, func() any { return new(proto.RogueMagicEnterRoomScRsp) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(ClockParkGetInfoScRsp, func() any { return new(proto.ClockParkGetInfoScRsp) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(SwordTrainingStartGameScRsp, func() any { return new(proto.SwordTrainingStartGameScRsp) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(EnterChallengeNextPhaseScRsp, func() any { return new(proto.EnterChallengeNextPhaseScRsp) })
	c.regMsg(ContentPackageGetDataCsReq, func() any { return new(proto.ContentPackageGetDataCsReq) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(SwordTrainingStoryBattleCsReq, func() any { return new(proto.SwordTrainingStoryBattleCsReq) })
	c.regMsg(TrainPartyBuildStartStepCsReq, func() any { return new(proto.TrainPartyBuildStartStepCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(LobbyKickOutCsReq, func() any { return new(proto.LobbyKickOutCsReq) })
	c.regMsg(QuitTrackPhotoStageCsReq, func() any { return new(proto.QuitTrackPhotoStageCsReq) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(GetMonopolyDailyReportScRsp, func() any { return new(proto.GetMonopolyDailyReportScRsp) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(UpdatePsnSettingsInfoScRsp, func() any { return new(proto.UpdatePsnSettingsInfoScRsp) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	c.regMsg(RogueTournRenameArchiveCsReq, func() any { return new(proto.RogueTournRenameArchiveCsReq) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(ClockParkUseBuffScRsp, func() any { return new(proto.ClockParkUseBuffScRsp) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(MonopolySocialEventEffectScNotify, func() any { return new(proto.MonopolySocialEventEffectScNotify) })
	c.regMsg(BoxingClubChallengeUpdateScNotify, func() any { return new(proto.BoxingClubChallengeUpdateScNotify) })
	c.regMsg(EvolveBuildLeaveScRsp, func() any { return new(proto.EvolveBuildLeaveScRsp) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(GetFriendChallengeDetailCsReq, func() any { return new(proto.GetFriendChallengeDetailCsReq) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(SwordTrainingDialogueSelectOptionScRsp, func() any { return new(proto.SwordTrainingDialogueSelectOptionScRsp) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(RaidCollectionDataCsReq, func() any { return new(proto.RaidCollectionDataCsReq) })
	c.regMsg(ChallengeBossPhaseSettleNotify, func() any { return new(proto.ChallengeBossPhaseSettleNotify) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(RogueArcadeRestartScRsp, func() any { return new(proto.RogueArcadeRestartScRsp) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(SwordTrainingGiveUpGameScRsp, func() any { return new(proto.SwordTrainingGiveUpGameScRsp) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(GetFriendChallengeDetailScRsp, func() any { return new(proto.GetFriendChallengeDetailScRsp) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(AddRelicFilterPlanCsReq, func() any { return new(proto.AddRelicFilterPlanCsReq) })
	c.regMsg(TrainPartyHandlePendingActionCsReq, func() any { return new(proto.TrainPartyHandlePendingActionCsReq) })
	c.regMsg(FightSessionStopScNotify, func() any { return new(proto.FightSessionStopScNotify) })
	c.regMsg(LobbyInviteCsReq, func() any { return new(proto.LobbyInviteCsReq) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(AvatarPathChangedNotify, func() any { return new(proto.AvatarPathChangedNotify) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(SummonActivityBattleEndScNotify, func() any { return new(proto.SummonActivityBattleEndScNotify) })
	c.regMsg(QuitBattleScNotify, func() any { return new(proto.QuitBattleScNotify) })
	c.regMsg(RogueMagicScepterDressInUnitCsReq, func() any { return new(proto.RogueMagicScepterDressInUnitCsReq) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(RogueModifierUpdateNotify, func() any { return new(proto.RogueModifierUpdateNotify) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(SpaceZooTakeCsReq, func() any { return new(proto.SpaceZooTakeCsReq) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(TrainPartyUseCardCsReq, func() any { return new(proto.TrainPartyUseCardCsReq) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(GiveUpBoxingClubChallengeScRsp, func() any { return new(proto.GiveUpBoxingClubChallengeScRsp) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(UpdateRotaterScNotify, func() any { return new(proto.UpdateRotaterScNotify) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(DeleteFriendScRsp, func() any { return new(proto.DeleteFriendScRsp) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(StartMatchScRsp, func() any { return new(proto.StartMatchScRsp) })
	c.regMsg(CommonRogueQueryCsReq, func() any { return new(proto.CommonRogueQueryCsReq) })
	c.regMsg(RogueTournEnterScRsp, func() any { return new(proto.RogueTournEnterScRsp) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(GetFightFestDataScRsp, func() any { return new(proto.GetFightFestDataScRsp) })
	c.regMsg(MatchThreeLevelEndCsReq, func() any { return new(proto.MatchThreeLevelEndCsReq) })
	c.regMsg(TrainPartyAddBuildDynamicBuffScRsp, func() any { return new(proto.TrainPartyAddBuildDynamicBuffScRsp) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(CurAssistChangedNotify, func() any { return new(proto.CurAssistChangedNotify) })
	c.regMsg(LobbyGetInfoScRsp, func() any { return new(proto.LobbyGetInfoScRsp) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(SyncRogueSeasonFinishScNotify, func() any { return new(proto.SyncRogueSeasonFinishScNotify) })
	c.regMsg(RogueTournGetAllArchiveCsReq, func() any { return new(proto.RogueTournGetAllArchiveCsReq) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(GetSpringRecoverDataCsReq, func() any { return new(proto.GetSpringRecoverDataCsReq) })
	c.regMsg(TrainPartyGamePlayStartScRsp, func() any { return new(proto.TrainPartyGamePlayStartScRsp) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(ContentPackageSyncDataScNotify, func() any { return new(proto.ContentPackageSyncDataScNotify) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(GetTrackPhotoActivityDataScRsp, func() any { return new(proto.GetTrackPhotoActivityDataScRsp) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(GetFriendBattleRecordDetailScRsp, func() any { return new(proto.GetFriendBattleRecordDetailScRsp) })
	c.regMsg(MusicRhythmDataCsReq, func() any { return new(proto.MusicRhythmDataCsReq) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(SetSpringRecoverConfigScRsp, func() any { return new(proto.SetSpringRecoverConfigScRsp) })
	c.regMsg(GetFightFestDataCsReq, func() any { return new(proto.GetFightFestDataCsReq) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(EvolveBuildStartStageScRsp, func() any { return new(proto.EvolveBuildStartStageScRsp) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(RogueMagicScepterTakeOffUnitCsReq, func() any { return new(proto.RogueMagicScepterTakeOffUnitCsReq) })
	c.regMsg(LastSpringRefreshTimeNotify, func() any { return new(proto.LastSpringRefreshTimeNotify) })
	c.regMsg(MatchThreeLevelEndScRsp, func() any { return new(proto.MatchThreeLevelEndScRsp) })
	c.regMsg(FightLeaveScNotify, func() any { return new(proto.FightLeaveScNotify) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(StartPartialChallengeCsReq, func() any { return new(proto.StartPartialChallengeCsReq) })
	c.regMsg(EvolveBuildStartLevelCsReq, func() any { return new(proto.EvolveBuildStartLevelCsReq) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(GetFriendChallengeLineupScRsp, func() any { return new(proto.GetFriendChallengeLineupScRsp) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(LogisticsGameCsReq, func() any { return new(proto.LogisticsGameCsReq) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(HeliobusLineupUpdateScNotify, func() any { return new(proto.HeliobusLineupUpdateScNotify) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(DailyFirstEnterMonopolyActivityCsReq, func() any { return new(proto.DailyFirstEnterMonopolyActivityCsReq) })
	c.regMsg(QueryProductInfoCsReq, func() any { return new(proto.QueryProductInfoCsReq) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(RogueArcadeGetInfoCsReq, func() any { return new(proto.RogueArcadeGetInfoCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(GetPlayerReturnMultiDropInfoScRsp, func() any { return new(proto.GetPlayerReturnMultiDropInfoScRsp) })
	c.regMsg(EvolveBuildShopAbilityUpCsReq, func() any { return new(proto.EvolveBuildShopAbilityUpCsReq) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(EnterSwordTrainingExamCsReq, func() any { return new(proto.EnterSwordTrainingExamCsReq) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(MusicRhythmStartLevelScRsp, func() any { return new(proto.MusicRhythmStartLevelScRsp) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(GetSummonActivityDataScRsp, func() any { return new(proto.GetSummonActivityDataScRsp) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(LobbyInviteScRsp, func() any { return new(proto.LobbyInviteScRsp) })
	c.regMsg(ModifyRelicFilterPlanScRsp, func() any { return new(proto.ModifyRelicFilterPlanScRsp) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(SwordTrainingSelectEndingScRsp, func() any { return new(proto.SwordTrainingSelectEndingScRsp) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(GetUpdatedArchiveDataScRsp, func() any { return new(proto.GetUpdatedArchiveDataScRsp) })
	c.regMsg(ClockParkQuitScriptCsReq, func() any { return new(proto.ClockParkQuitScriptCsReq) })
	c.regMsg(MusicRhythmMaxDifficultyLevelsUnlockNotify, func() any { return new(proto.MusicRhythmMaxDifficultyLevelsUnlockNotify) })
	c.regMsg(RogueArcadeStartScRsp, func() any { return new(proto.RogueArcadeStartScRsp) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(EnterSummonActivityStageScRsp, func() any { return new(proto.EnterSummonActivityStageScRsp) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(SwordTrainingStoryConfirmCsReq, func() any { return new(proto.SwordTrainingStoryConfirmCsReq) })
	c.regMsg(SettleTrackPhotoStageCsReq, func() any { return new(proto.SettleTrackPhotoStageCsReq) })
	c.regMsg(BattleLogReportScRsp, func() any { return new(proto.BattleLogReportScRsp) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(RogueMagicAutoDressInMagicUnitChangeScNotify, func() any { return new(proto.RogueMagicAutoDressInMagicUnitChangeScNotify) })
	c.regMsg(StartPartialChallengeScRsp, func() any { return new(proto.StartPartialChallengeScRsp) })
	c.regMsg(CommonRogueQueryScRsp, func() any { return new(proto.CommonRogueQueryScRsp) })
	c.regMsg(GetFriendChallengeLineupCsReq, func() any { return new(proto.GetFriendChallengeLineupCsReq) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(GetCrossInfoScRsp, func() any { return new(proto.GetCrossInfoScRsp) })
	c.regMsg(TakeMaterialSubmitActivityRewardCsReq, func() any { return new(proto.TakeMaterialSubmitActivityRewardCsReq) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(EvolveBuildGiveupCsReq, func() any { return new(proto.EvolveBuildGiveupCsReq) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(StartTrackPhotoStageScRsp, func() any { return new(proto.StartTrackPhotoStageScRsp) })
	c.regMsg(SubmitMaterialSubmitActivityMaterialCsReq, func() any { return new(proto.SubmitMaterialSubmitActivityMaterialCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(SubmitOrigamiItemCsReq, func() any { return new(proto.SubmitOrigamiItemCsReq) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(SwordTrainingTurnActionScRsp, func() any { return new(proto.SwordTrainingTurnActionScRsp) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(AcceptMissionEventScRsp, func() any { return new(proto.AcceptMissionEventScRsp) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(SubmitOfferingItemScRsp, func() any { return new(proto.SubmitOfferingItemScRsp) })
	c.regMsg(RestartChallengePhaseCsReq, func() any { return new(proto.RestartChallengePhaseCsReq) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(SetRedPointStatusScNotify, func() any { return new(proto.SetRedPointStatusScNotify) })
	c.regMsg(TrainPartyTakeBuildLevelAwardScRsp, func() any { return new(proto.TrainPartyTakeBuildLevelAwardScRsp) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(WolfBroGameExplodeMonsterScRsp, func() any { return new(proto.WolfBroGameExplodeMonsterScRsp) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(ContentPackageGetDataScRsp, func() any { return new(proto.ContentPackageGetDataScRsp) })
	c.regMsg(DrinkMakerDayEndScNotify, func() any { return new(proto.DrinkMakerDayEndScNotify) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(RogueTournConfirmSettleScRsp, func() any { return new(proto.RogueTournConfirmSettleScRsp) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(ClockParkHandleWaitOperationCsReq, func() any { return new(proto.ClockParkHandleWaitOperationCsReq) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(TakeOfferingRewardCsReq, func() any { return new(proto.TakeOfferingRewardCsReq) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(GetMaterialSubmitActivityDataCsReq, func() any { return new(proto.GetMaterialSubmitActivityDataCsReq) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(FightEnterScRsp, func() any { return new(proto.FightEnterScRsp) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(GetSwordTrainingDataScRsp, func() any { return new(proto.GetSwordTrainingDataScRsp) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(GetGameStateServiceConfigCsReq, func() any { return new(proto.GetGameStateServiceConfigCsReq) })
	c.regMsg(GetStoryLineInfoScRsp, func() any { return new(proto.GetStoryLineInfoScRsp) })
	c.regMsg(GetChallengeGroupStatisticsCsReq, func() any { return new(proto.GetChallengeGroupStatisticsCsReq) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(GetPetDataCsReq, func() any { return new(proto.GetPetDataCsReq) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(AddRelicFilterPlanScRsp, func() any { return new(proto.AddRelicFilterPlanScRsp) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(GetAllRedDotDataCsReq, func() any { return new(proto.GetAllRedDotDataCsReq) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(FightMatch3StartCountDownScNotify, func() any { return new(proto.FightMatch3StartCountDownScNotify) })
	c.regMsg(RogueMagicEnterCsReq, func() any { return new(proto.RogueMagicEnterCsReq) })
	c.regMsg(RogueTournReviveCostUpdateScNotify, func() any { return new(proto.RogueTournReviveCostUpdateScNotify) })
	c.regMsg(ChooseBoxingClubResonanceCsReq, func() any { return new(proto.ChooseBoxingClubResonanceCsReq) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(EvolveBuildCoinNotify, func() any { return new(proto.EvolveBuildCoinNotify) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(RogueMagicEnableTalentCsReq, func() any { return new(proto.RogueMagicEnableTalentCsReq) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(SwordTrainingStartGameCsReq, func() any { return new(proto.SwordTrainingStartGameCsReq) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(MatchBoxingClubOpponentCsReq, func() any { return new(proto.MatchBoxingClubOpponentCsReq) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(RogueMagicStartCsReq, func() any { return new(proto.RogueMagicStartCsReq) })
	c.regMsg(SwordTrainingStoryConfirmScRsp, func() any { return new(proto.SwordTrainingStoryConfirmScRsp) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(EvolveBuildReRandomStageScRsp, func() any { return new(proto.EvolveBuildReRandomStageScRsp) })
	c.regMsg(SetMultipleAvatarPathsScRsp, func() any { return new(proto.SetMultipleAvatarPathsScRsp) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(RogueMagicSetAutoDressInMagicUnitScRsp, func() any { return new(proto.RogueMagicSetAutoDressInMagicUnitScRsp) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(SwordTrainingRestoreGameCsReq, func() any { return new(proto.SwordTrainingRestoreGameCsReq) })
	c.regMsg(RogueMagicAutoDressInUnitScRsp, func() any { return new(proto.RogueMagicAutoDressInUnitScRsp) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(RogueTournStartCsReq, func() any { return new(proto.RogueTournStartCsReq) })
	c.regMsg(RogueTournEnterRoomCsReq, func() any { return new(proto.RogueTournEnterRoomCsReq) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(DifficultyAdjustmentGetDataCsReq, func() any { return new(proto.DifficultyAdjustmentGetDataCsReq) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(TrainPartySyncUpdateScNotify, func() any { return new(proto.TrainPartySyncUpdateScNotify) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(SetBoxingClubResonanceLineupCsReq, func() any { return new(proto.SetBoxingClubResonanceLineupCsReq) })
	c.regMsg(StartTimedFarmElementScRsp, func() any { return new(proto.StartTimedFarmElementScRsp) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(TakeOfferingRewardScRsp, func() any { return new(proto.TakeOfferingRewardScRsp) })
	c.regMsg(SelectPamSkinCsReq, func() any { return new(proto.SelectPamSkinCsReq) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(RogueModifierStageStartNotify, func() any { return new(proto.RogueModifierStageStartNotify) })
	c.regMsg(SwordTrainingActionTurnSettleScNotify, func() any { return new(proto.SwordTrainingActionTurnSettleScNotify) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(DeleteRelicFilterPlanCsReq, func() any { return new(proto.DeleteRelicFilterPlanCsReq) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(TrainPartyTakeBuildLevelAwardCsReq, func() any { return new(proto.TrainPartyTakeBuildLevelAwardCsReq) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(GetMonopolyDailyReportCsReq, func() any { return new(proto.GetMonopolyDailyReportCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(EnteredSceneChangeScNotify, func() any { return new(proto.EnteredSceneChangeScNotify) })
	c.regMsg(SwordTrainingExamResultConfirmScRsp, func() any { return new(proto.SwordTrainingExamResultConfirmScRsp) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(EvolveBuildTakeExpRewardScRsp, func() any { return new(proto.EvolveBuildTakeExpRewardScRsp) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(EnterFightActivityStageScRsp, func() any { return new(proto.EnterFightActivityStageScRsp) })
	c.regMsg(GetBattleCollegeDataCsReq, func() any { return new(proto.GetBattleCollegeDataCsReq) })
	c.regMsg(ChessRogueNousEnableRogueTalentCsReq, func() any { return new(proto.ChessRogueNousEnableRogueTalentCsReq) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(TrainPartyGetDataCsReq, func() any { return new(proto.TrainPartyGetDataCsReq) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(MonopolyAcceptQuizScRsp, func() any { return new(proto.MonopolyAcceptQuizScRsp) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(MuseumTakeCollectRewardScRsp, func() any { return new(proto.MuseumTakeCollectRewardScRsp) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(RogueMagicScepterDressInUnitScRsp, func() any { return new(proto.RogueMagicScepterDressInUnitScRsp) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(MonopolyEventSelectFriendScRsp, func() any { return new(proto.MonopolyEventSelectFriendScRsp) })
	c.regMsg(RogueTournEnterRoomScRsp, func() any { return new(proto.RogueTournEnterRoomScRsp) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(FightHeartBeatCsReq, func() any { return new(proto.FightHeartBeatCsReq) })
	c.regMsg(RogueTournEnablePermanentTalentScRsp, func() any { return new(proto.RogueTournEnablePermanentTalentScRsp) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(TrainPartyUpdatePosEnvScRsp, func() any { return new(proto.TrainPartyUpdatePosEnvScRsp) })
	c.regMsg(RogueMagicEnterLayerScRsp, func() any { return new(proto.RogueMagicEnterLayerScRsp) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(UpdateMapRotationDataScNotify, func() any { return new(proto.UpdateMapRotationDataScNotify) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(RogueMagicLeaveCsReq, func() any { return new(proto.RogueMagicLeaveCsReq) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(SwordTrainingLearnSkillScRsp, func() any { return new(proto.SwordTrainingLearnSkillScRsp) })
	c.regMsg(MonopolyGuessBuyInformationScRsp, func() any { return new(proto.MonopolyGuessBuyInformationScRsp) })
	c.regMsg(RogueMagicLevelInfoUpdateScNotify, func() any { return new(proto.RogueMagicLevelInfoUpdateScNotify) })
	c.regMsg(RogueModifierSelectCellScRsp, func() any { return new(proto.RogueModifierSelectCellScRsp) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(TrainPartyEnterCsReq, func() any { return new(proto.TrainPartyEnterCsReq) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(FightFestUpdateCoinNotify, func() any { return new(proto.FightFestUpdateCoinNotify) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
}
