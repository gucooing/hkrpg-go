package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	NONE                                               = 0
	SpringRecoverCsReq                                 = 1444
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3729
	FinishEmotionDialoguePerformanceScRsp              = 6343
	SelectChatBubbleScRsp                              = 5188
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5616
	SubmitMonsterResearchActivityMaterialCsReq         = 2637
	EnterAetherDivideSceneCsReq                        = 4834
	HeliobusActivityDataCsReq                          = 5834
	CommonRogueQueryScRsp                              = 5698
	GetRogueShopMiracleInfoCsReq                       = 5688
	ApplyFriendCsReq                                   = 2919
	GetHeartDialInfoScRsp                              = 6348
	ComposeLimitNumUpdateNotify                        = 518
	GetCurAssistCsReq                                  = 2924
	TakeRollShopRewardCsReq                            = 6920
	RelicRecommendScRsp                                = 592
	ChangeLineupLeaderCsReq                            = 706
	TrialActivityDataChangeScNotify                    = 2604
	DrinkMakerChallengeCsReq                           = 6985
	GetMbtiReportCsReq                                 = 7071
	AlleyGuaranteedFundsCsReq                          = 4724
	UpdateFeatureSwitchScNotify                        = 55
	FinishQuestCsReq                                   = 929
	ChessRogueRollDiceScRsp                            = 5546
	GetAuthkeyScRsp                                    = 95
	ActivateFarmElementCsReq                           = 1492
	SetAetherDivideLineUpCsReq                         = 4896
	StartAetherDivideChallengeBattleScRsp              = 4843
	MarkAvatarCsReq                                    = 341
	StartRogueScRsp                                    = 1888
	UnlockAvatarSkinScNotify                           = 301
	SellItemCsReq                                      = 537
	SwapLineupScRsp                                    = 729
	TakeAllApRewardCsReq                               = 3309
	GetChallengeGroupStatisticsCsReq                   = 1795
	RetcodeNotify                                      = 98
	FinishPerformSectionIdScRsp                        = 2729
	AetherDivideFinishChallengeScNotify                = 4828
	SyncLineupNotify                                   = 745
	UseItemScRsp                                       = 543
	GetFarmStageGachaInfoCsReq                         = 1362
	UseTreasureDungeonItemScRsp                        = 4430
	AcceptMainMissionCsReq                             = 1291
	GetWolfBroGameDataCsReq                            = 6586
	JoinLineupCsReq                                    = 702
	DelSaveRaidScNotify                                = 2237
	InteractPropScRsp                                  = 1488
	DeleteSocialEventServerCacheScRsp                  = 7015
	EnableRogueTalentCsReq                             = 1874
	EvolveBuildReRandomStageCsReq                      = 7131
	CurTrialActivityScNotify                           = 2605
	UnlockPhoneThemeScNotify                           = 5143
	SubmitEmotionItemScRsp                             = 6309
	UpdateRogueAdventureRoomScoreScRsp                 = 5666
	ClockParkStartScriptCsReq                          = 7219
	NewAssistHistoryNotify                             = 2954
	AlleyShipUsedCountScNotify                         = 4797
	RecoverAllLineupCsReq                              = 1424
	HeliobusSnsPostScRsp                               = 5809
	SetNicknameCsReq                                   = 39
	SharePunkLordMonsterCsReq                          = 3202
	GetChessRogueBuffEnhanceInfoScRsp                  = 5426
	AlleyPlacingGameCsReq                              = 4796
	TakeRogueEventHandbookRewardCsReq                  = 5689
	GetPlayerReturnMultiDropInfoScRsp                  = 4602
	MuseumTakeCollectRewardCsReq                       = 4328
	MonopolyGameCreateScNotify                         = 7025
	DailyFirstEnterMonopolyActivityCsReq               = 7096
	MonopolyGetRafflePoolInfoScRsp                     = 7020
	GetActivityScheduleConfigCsReq                     = 2602
	DeleteSummonUnitCsReq                              = 1417
	RefreshAlleyOrderScRsp                             = 4742
	FinishChessRogueSubStoryScRsp                      = 5437
	GetCurSceneInfoScRsp                               = 1443
	GetFightActivityDataScRsp                          = 3648
	QuitBattleCsReq                                    = 162
	FinishFirstTalkNpcScRsp                            = 2143
	EvolveBuildQueryInfoScRsp                          = 7149
	StartCocoonStageScRsp                              = 1454
	BattleLogReportScRsp                               = 145
	TreasureDungeonFinishScNotify                      = 4448
	ChessRogueEnterScRsp                               = 5559
	SetMissionEventProgressCsReq                       = 1256
	TreasureDungeonDataScNotify                        = 4434
	StrongChallengeActivityBattleEndScNotify           = 6602
	StartAlleyEventCsReq                               = 4719
	ChessRogueCellUpdateNotify                         = 5508
	QuestRecordScNotify                                = 986
	MonopolyCheatDiceCsReq                             = 7041
	BuyRogueShopMiracleCsReq                           = 5643
	EnterRogueEndlessActivityStageScRsp                = 6088
	StartAetherDivideStageBattleCsReq                  = 4816
	GetLoginActivityCsReq                              = 2634
	UpdateServerPrefsDataCsReq                         = 6102
	ClockParkGetInfoCsReq                              = 7234
	StartBoxingClubBattleScRsp                         = 4209
	GetFeverTimeActivityDataScRsp                      = 7154
	GetStoryLineInfoCsReq                              = 6234
	ChessRogueUpdateReviveInfoScNotify                 = 5419
	GetWaypointCsReq                                   = 434
	ReturnLastTownScRsp                                = 1430
	SubmitOfferingItemCsReq                            = 6933
	LeaveRogueScRsp                                    = 1843
	FinishTutorialGuideCsReq                           = 1645
	SpaceZooExchangeItemScRsp                          = 6796
	MonopolyAcceptQuizCsReq                            = 7054
	MonopolySocialEventEffectScNotify                  = 7046
	SetSpringRecoverConfigScRsp                        = 1435
	TravelBrochureSetPageDescStatusScRsp               = 6442
	GetAlleyInfoCsReq                                  = 4734
	GetAllSaveRaidScRsp                                = 2242
	FightTreasureDungeonMonsterScRsp                   = 4442
	FinishCosumeItemMissionCsReq                       = 1296
	WolfBroGamePickupBulletScRsp                       = 6533
	ClockParkUnlockTalentScRsp                         = 7209
	ExchangeStaminaScRsp                               = 33
	ChessRogueNousDiceSurfaceUnlockNotify              = 5413
	FinishChapterScNotify                              = 4962
	GetPunkLordBattleRecordScRsp                       = 3224
	GetShopListScRsp                                   = 1548
	LogisticsScoreRewardSyncInfoScNotify               = 4725
	GetStageLineupScRsp                                = 748
	UpdateRotaterScNotify                              = 6839
	GetTelevisionActivityDataCsReq                     = 6961
	LockRelicCsReq                                     = 595
	EnterSceneCsReq                                    = 1472
	SetClientPausedCsReq                               = 1500
	PlayerLogoutCsReq                                  = 62
	SpaceZooMutateCsReq                                = 6702
	BattlePassInfoNotify                               = 3034
	MonopolyMoveCsReq                                  = 7043
	ReserveStaminaExchangeScRsp                        = 40
	MarkItemCsReq                                      = 511
	StartFinishMainMissionScNotify                     = 1218
	FinishAeonDialogueGroupCsReq                       = 1831
	HeliobusSnsUpdateScNotify                          = 5845
	MonopolyRollDiceScRsp                              = 7019
	MonopolyGameRaiseRatioCsReq                        = 7018
	UpdateFloorSavedValueNotify                        = 1420
	WolfBroGameActivateBulletScRsp                     = 6595
	TrainVisitorBehaviorFinishScRsp                    = 3748
	LockEquipmentCsReq                                 = 502
	AetherDivideTainerInfoScNotify                     = 4861
	TextJoinBatchSaveCsReq                             = 3802
	ExpUpRelicScRsp                                    = 559
	AlleyShipUnlockScNotify                            = 4763
	DeactivateFarmElementScRsp                         = 1467
	ChessRogueEnterNextLayerScRsp                      = 5436
	SyncRogueFinishScNotify                            = 1833
	HeliobusSnsReadCsReq                               = 5862
	ChessRogueNousEditDiceCsReq                        = 5550
	LeaveTrialActivityScRsp                            = 2646
	SetGameplayBirthdayScRsp                           = 44
	GroupStateChangeScNotify                           = 1447
	MonopolyTakePhaseRewardScRsp                       = 7064
	SetSignatureCsReq                                  = 2829
	EquipAetherDividePassiveSkillCsReq                 = 4833
	EvolveBuildUnlockInfoNotify                        = 7110
	GetLevelRewardTakenListCsReq                       = 30
	SetLanguageScRsp                                   = 61
	InteractTreasureDungeonGridScRsp                   = 4439
	GetRndOptionCsReq                                  = 3434
	ChessRogueNousDiceUpdateNotify                     = 5452
	TakeQuestRewardCsReq                               = 962
	TravelBrochureSelectMessageCsReq                   = 6402
	ChessRoguePickAvatarScRsp                          = 5449
	LeaveAetherDivideSceneScRsp                        = 4888
	MuseumRandomEventStartScNotify                     = 4337
	UpdateTrackMainMissionIdCsReq                      = 1254
	RemoveStuffFromAreaScRsp                           = 4343
	SetGenderScRsp                                     = 25
	MonopolyScrachRaffleTicketCsReq                    = 7099
	BoxingClubChallengeUpdateScNotify                  = 4229
	SyncRogueCommonVirtualItemInfoScNotify             = 5673
	GetBasicInfoScRsp                                  = 73
	AetherDivideSpiritExpUpCsReq                       = 4885
	GetFirstTalkByPerformanceNpcScRsp                  = 2168
	TrainRefreshTimeNotify                             = 3702
	GetFriendDevelopmentInfoScRsp                      = 2903
	ChooseBoxingClubStageOptionalBuffScRsp             = 4259
	TakeCityShopRewardScRsp                            = 1509
	ChessRogueUpdateUnlockLevelScNotify                = 5582
	ReviveRogueAvatarCsReq                             = 1837
	SetTurnFoodSwitchScRsp                             = 600
	SceneEntityMoveScNotify                            = 1445
	LeaveMapRotationRegionCsReq                        = 6886
	TakeTrialActivityRewardScRsp                       = 2671
	ClientObjDownloadDataScNotify                      = 58
	ExchangeHcoinScRsp                                 = 585
	DressAvatarSkinCsReq                               = 330
	EvolveBuildGiveupScRsp                             = 7150
	QueryProductInfoCsReq                              = 90
	TakeChallengeRaidRewardScRsp                       = 2286
	GetReplayTokenCsReq                                = 3534
	DoGachaInRollShopCsReq                             = 6913
	ChessRogueNousEnableRogueTalentScRsp               = 5425
	EnhanceCommonRogueBuffCsReq                        = 5685
	EnterStrongChallengeActivityStageScRsp             = 6688
	WolfBroGameUseBulletCsReq                          = 6568
	GetMapRotationDataCsReq                            = 6845
	PromoteEquipmentScRsp                              = 588
	RaidCollectionDataScNotify                         = 6953
	RankUpAvatarScRsp                                  = 333
	GetMovieRacingDataScRsp                            = 4130
	SceneEntityMoveScRsp                               = 1448
	AlleyEventEffectNotify                             = 4729
	GetCurBattleInfoScRsp                              = 109
	ChangeStoryLineCsReq                               = 6288
	TriggerVoiceCsReq                                  = 4196
	GameplayCounterUpdateScNotify                      = 1478
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5542
	GiveUpBoxingClubChallengeScRsp                     = 4243
	AcceptedPamMissionExpireScRsp                      = 4048
	WolfBroGameExplodeMonsterCsReq                     = 6542
	RegionStopScNotify                                 = 42
	TakeRogueEndlessActivityPointRewardScRsp           = 6019
	GetChessRogueStoryAeonTalkInfoCsReq                = 5477
	ClockParkFinishScriptScNotify                      = 7216
	DailyTaskDataScNotify                              = 1243
	GetMonopolyFriendRankingListScRsp                  = 7004
	GetMonopolyDailyReportScRsp                        = 7021
	BuyNpcStuffScRsp                                   = 4388
	GetMissionDataCsReq                                = 1234
	GetFriendApplyListInfoCsReq                        = 2902
	GetUnlockTeleportCsReq                             = 1440
	ChessRogueFinishCurRoomNotify                      = 5422
	TextJoinQueryCsReq                                 = 3862
	RotateMapScRsp                                     = 6843
	LeaveChallengeScRsp                                = 1709
	TakeRogueAeonLevelRewardCsReq                      = 1899
	MonopolyGuessDrawScNotify                          = 7067
	MonopolyGameGachaCsReq                             = 7024
	SummonPunkLordMonsterScRsp                         = 3243
	LeaveMapRotationRegionScNotify                     = 6833
	SetGroupCustomSaveDataScRsp                        = 1403
	SwitchLineupIndexCsReq                             = 759
	FinishItemIdCsReq                                  = 2702
	MuseumRandomEventSelectCsReq                       = 4330
	PlayerGetTokenScRsp                                = 9
	RefreshTriggerByClientCsReq                        = 1432
	EnterTreasureDungeonScRsp                          = 4406
	GetQuestRecordScRsp                                = 943
	RogueModifierUpdateNotify                          = 5343
	RestoreWolfBroGameArchiveCsReq                     = 6502
	CommonRogueUpdateScNotify                          = 5671
	MissionGroupWarnScNotify                           = 1268
	GetEnteredSceneScRsp                               = 1431
	ArchiveWolfBroGameScRsp                            = 6588
	SyncTaskScRsp                                      = 1219
	PlayerReturnTakeRewardScRsp                        = 4543
	EvolveBuildLeaveScRsp                              = 7101
	PickRogueAvatarScRsp                               = 1895
	ComposeItemCsReq                                   = 596
	SetFriendMarkCsReq                                 = 2966
	SyncRogueRewardInfoScNotify                        = 1883
	EvolveBuildShopAbilityUpCsReq                      = 7133
	UnlockHeadIconScNotify                             = 2886
	SyncHandleFriendScNotify                           = 2968
	GetGachaInfoCsReq                                  = 1934
	RogueModifierSelectCellScRsp                       = 5302
	HeliobusUpgradeLevelScRsp                          = 5806
	TakeMultipleExpeditionRewardScRsp                  = 2537
	TakeLoginActivityRewardScRsp                       = 2688
	PVEBattleResultCsReq                               = 134
	MonopolyGuessChooseScRsp                           = 7065
	QuitBattleScNotify                                 = 186
	MonopolyGiveUpCurContentScRsp                      = 7001
	TakeOffRelicCsReq                                  = 342
	StartRaidScRsp                                     = 2248
	RogueModifierStageStartNotify                      = 5329
	MonopolyGetRegionProgressCsReq                     = 7040
	GetFriendBattleRecordDetailScRsp                   = 2971
	SetHeroBasicTypeCsReq                              = 91
	ComposeSelectedRelicCsReq                          = 556
	GetTutorialGuideScRsp                              = 1688
	SceneCastSkillCostMpScRsp                          = 1433
	GetFriendAssistListScRsp                           = 2935
	GetTreasureDungeonActivityDataScRsp                = 4468
	InteractChargerCsReq                               = 6862
	GetPlayerBoardDataScRsp                            = 2848
	EntityBindPropScRsp                                = 1425
	GetExpeditionDataCsReq                             = 2534
	MonopolyConfirmRandomScRsp                         = 7039
	GetMonopolyDailyReportCsReq                        = 7076
	AlleyTakeEventRewardScRsp                          = 4708
	TakeMailAttachmentScRsp                            = 843
	SetDisplayAvatarCsReq                              = 2802
	ChessRogueQueryAeonDimensionsScRsp                 = 5466
	HeliobusSnsLikeCsReq                               = 5819
	TravelBrochureSetCustomValueScRsp                  = 6459
	ScenePlaneEventScNotify                            = 1484
	MatchBoxingClubOpponentScRsp                       = 4288
	CurAssistChangedNotify                             = 3000
	ArchiveWolfBroGameCsReq                            = 6562
	PlayerReturnTakePointRewardCsReq                   = 4502
	ChessRogueSelectBpCsReq                            = 5549
	DeployRotaterCsReq                                 = 6802
	AceAntiCheaterScRsp                                = 75
	GetAllLineupDataCsReq                              = 739
	TakeFightActivityRewardCsReq                       = 3609
	GetMonopolyMbtiReportRewardScRsp                   = 7081
	MultipleDropInfoNotify                             = 4609
	GetPunkLordDataCsReq                               = 3259
	UnlockTutorialGuideScRsp                           = 1643
	WaypointShowNewCsNotify                            = 419
	AlleyShopLevelScNotify                             = 4756
	EnterSceneByServerScNotify                         = 1410
	GeneralVirtualItemDataNotify                       = 565
	DeleteFriendScRsp                                  = 2906
	GetTutorialScRsp                                   = 1648
	PlayerGetTokenCsReq                                = 2
	TravelBrochureGetDataScRsp                         = 6448
	GiveUpBoxingClubChallengeCsReq                     = 4219
	SwitchAetherDivideLineUpSlotScRsp                  = 4839
	HeliobusEnterBattleCsReq                           = 5839
	FinishCurTurnCsReq                                 = 4345
	TakeRogueScoreRewardScRsp                          = 1830
	ClockParkHandleWaitOperationScRsp                  = 7268
	GetFriendRecommendListInfoScRsp                    = 2939
	PlayerReturnSignCsReq                              = 4548
	EnterFeverTimeActivityStageScRsp                   = 7159
	GetFriendLoginInfoCsReq                            = 2990
	GetFriendListInfoScRsp                             = 2948
	BoxingClubRewardScNotify                           = 4286
	RemoveRotaterCsReq                                 = 6842
	GetQuestDataCsReq                                  = 934
	PlayerReturnInfoQueryScRsp                         = 4529
	PlayerReturnStartScNotify                          = 4534
	ChallengeLineupNotify                              = 1768
	ExchangeRogueBuffWithMiracleScRsp                  = 5639
	UpdateMechanismBarScNotify                         = 1471
	UpdatePlayerSettingCsReq                           = 7
	UpdateTrackMainMissionIdScRsp                      = 1279
	GetAssistHistoryScRsp                              = 2908
	PlayBackGroundMusicCsReq                           = 3162
	GetGunPlayDataCsReq                                = 4163
	SyncRogueExploreWinScNotify                        = 1811
	GetMonopolyFriendRankingListCsReq                  = 7044
	MonopolyGameBingoFlipCardCsReq                     = 7011
	StartBattleCollegeScRsp                            = 5702
	ChangeScriptEmotionScRsp                           = 6388
	LeaveAetherDivideSceneCsReq                        = 4862
	SyncRogueCommonPendingActionScNotify               = 5692
	MarkReadMailScRsp                                  = 888
	GetRogueDialogueEventDataScRsp                     = 1844
	MarkChatEmojiScRsp                                 = 3968
	GetServerPrefsDataCsReq                            = 6162
	GetSecretKeyInfoCsReq                              = 22
	ChessRogueSelectCellScRsp                          = 5450
	GetTrainVisitorBehaviorScRsp                       = 3788
	SelectInclinationTextScRsp                         = 2129
	SetIsDisplayAvatarInfoScRsp                        = 2843
	GetRogueInitialScoreCsReq                          = 1865
	LeaveChallengeCsReq                                = 1702
	BuyBpLevelScRsp                                    = 3019
	AcceptMissionEventScRsp                            = 1237
	GetMonopolyInfoCsReq                               = 7034
	GetQuestRecordCsReq                                = 919
	StartTimedCocoonStageCsReq                         = 1426
	EnterRogueScRsp                                    = 1809
	LastSpringRefreshTimeNotify                        = 1439
	LogisticsDetonateStarSkiffScRsp                    = 4779
	DrinkMakerDayEndScNotify                           = 6984
	TakeOffEquipmentCsReq                              = 345
	GetRogueBuffEnhanceInfoScRsp                       = 1856
	EnhanceChessRogueBuffCsReq                         = 5592
	CancelCacheNotifyScRsp                             = 4129
	SendMsgCsReq                                       = 3934
	SwitchLineupIndexScRsp                             = 795
	MonopolyMoveScRsp                                  = 7086
	StoryLineInfoScNotify                              = 6262
	ChessRogueUpdateBoardScNotify                      = 5502
	GetQuestDataScRsp                                  = 948
	BuyNpcStuffCsReq                                   = 4362
	TakeRogueScoreRewardCsReq                          = 1816
	CancelActivityExpeditionCsReq                      = 2568
	HeliobusEnterBattleScRsp                           = 5816
	GetLevelRewardTakenListScRsp                       = 85
	SetHeadIconScRsp                                   = 2888
	AcceptExpeditionCsReq                              = 2562
	HandleRogueCommonPendingActionCsReq                = 5604
	LogisticsInfoScNotify                              = 4728
	DailyFirstMeetPamScRsp                             = 3488
	ChessRogueQuestFinishNotify                        = 5565
	RechargeSuccNotify                                 = 516
	MonopolyLikeScNotify                               = 7098
	SyncRogueMapRoomScNotify                           = 1890
	SetCurWaypointCsReq                                = 462
	ChessRogueUpdateAllowedSelectCellScNotify          = 5577
	SceneEntityTeleportCsReq                           = 1412
	TravelBrochureUpdatePasterPosScRsp                 = 6468
	ChessRogueCheatRollScRsp                           = 5599
	UnlockTutorialScRsp                                = 1609
	UnlockTeleportNotify                               = 1483
	SetFriendRemarkNameScRsp                           = 2930
	GetFriendListInfoCsReq                             = 2934
	MonopolyTakeRaffleTicketRewardScRsp                = 7070
	HeartDialScriptChangeScNotify                      = 6386
	ExchangeHcoinCsReq                                 = 530
	HeliobusLineupUpdateScNotify                       = 5863
	ChessRogueQueryBpCsReq                             = 5495
	TakeChapterRewardCsReq                             = 443
	SelectPhoneThemeScRsp                              = 5119
	SyncRoguePickAvatarInfoScNotify                    = 1880
	TakeRogueMiracleHandbookRewardScRsp                = 5665
	DoGachaInRollShopScRsp                             = 6917
	GetSaveLogisticsMapCsReq                           = 4718
	EnterFeverTimeActivityStageCsReq                   = 7151
	StartChallengeScRsp                                = 1788
	PVEBattleResultScRsp                               = 148
	UpdateRedDotDataCsReq                              = 5962
	FinishPlotCsReq                                    = 1134
	BuyGoodsScRsp                                      = 1588
	BuyRogueShopBuffCsReq                              = 5629
	SubmitOrigamiItemCsReq                             = 4133
	UpdateMovieRacingDataScRsp                         = 4156
	MonopolyGameBingoFlipCardScRsp                     = 7008
	ExpeditionDataChangeScNotify                       = 2586
	GetTelevisionActivityDataScRsp                     = 6978
	FinishTutorialCsReq                                = 1686
	WolfBroGameDataChangeScNotify                      = 6545
	GroupStateChangeScRsp                              = 1421
	SyncRogueReviveInfoScNotify                        = 1891
	HeliobusUpgradeLevelCsReq                          = 5896
	HeliobusSelectSkillCsReq                           = 5859
	PlayerKickOutScNotify                              = 86
	FinishFirstTalkByPerformanceNpcScRsp               = 2106
	SyncClientResVersionCsReq                          = 119
	SceneUpdatePositionVersionNotify                   = 1468
	EnterTelevisionActivityStageScRsp                  = 6980
	TrainVisitorRewardSendNotify                       = 3709
	EndDrinkMakerSequenceScRsp                         = 6999
	GetMarkItemListCsReq                               = 524
	AetherDivideRefreshEndlessScRsp                    = 4882
	ServerSimulateBattleFinishScNotify                 = 168
	EnterTrialActivityStageScRsp                       = 2693
	ChessRogueSkipTeachingLevelScRsp                   = 5474
	EntityBindPropCsReq                                = 1479
	GetMultipleDropInfoCsReq                           = 4634
	GetDailyActiveInfoCsReq                            = 3362
	GetPhoneDataCsReq                                  = 5134
	UpdateGunPlayDataScRsp                             = 4128
	SpaceZooExchangeItemCsReq                          = 6768
	ChessRogueReRollDiceCsReq                          = 5490
	PlayerHeartBeatScRsp                               = 49
	SecurityReportScRsp                                = 4168
	RaidCollectionDataScRsp                            = 6958
	GetDrinkMakerDataScRsp                             = 6998
	ExtraLineupDestroyNotify                           = 763
	MonopolyGetRegionProgressScRsp                     = 7069
	GetPrivateChatHistoryCsReq                         = 3902
	PlayerReturnTakeRewardCsReq                        = 4519
	GetSingleRedDotParamGroupScRsp                     = 5909
	RankUpAvatarCsReq                                  = 306
	GameplayCounterCountDownScRsp                      = 1464
	GetRogueHandbookDataCsReq                          = 5654
	SceneCastSkillScRsp                                = 1409
	ChessRogueGoAheadScRsp                             = 5431
	MonopolyReRollRandomScRsp                          = 7042
	ChessRogueChangeyAeonDimensionNotify               = 5557
	MarkAvatarScRsp                                    = 328
	ChessRogueNousGetRogueTalentInfoCsReq              = 5448
	GetTutorialCsReq                                   = 1634
	GetArchiveDataScRsp                                = 2348
	SyncAddBlacklistScNotify                           = 2942
	AlleyTakeEventRewardCsReq                          = 4711
	ExchangeRogueRewardKeyCsReq                        = 1871
	ChessRogueQuitCsReq                                = 5575
	ChessRogueGiveUpCsReq                              = 5463
	QueryProductInfoScRsp                              = 67
	StartTimedFarmElementScRsp                         = 1460
	ReplaceLineupCsReq                                 = 785
	GetTrainVisitorRegisterCsReq                       = 3719
	SetBoxingClubResonanceLineupCsReq                  = 4296
	GetFriendBattleRecordDetailCsReq                   = 2998
	GetAllRedDotDataCsReq                              = 5934
	UseItemCsReq                                       = 519
	ShareScRsp                                         = 4148
	StartTrialActivityCsReq                            = 2649
	GetStrongChallengeActivityDataCsReq                = 6634
	TakeAllApRewardScRsp                               = 3319
	GetTrialActivityDataScRsp                          = 2644
	GetChessRogueBuffEnhanceInfoCsReq                  = 5522
	MonopolyGetRaffleTicketScRsp                       = 7010
	GetFriendApplyListInfoScRsp                        = 2909
	EvolveBuildTakeExpRewardCsReq                      = 7122
	GetPlayerBoardDataCsReq                            = 2834
	EvolveBuildStartLevelCsReq                         = 7148
	UnlockSkilltreeCsReq                               = 302
	TakeAllRewardCsReq                                 = 3043
	ChessRogueQueryAeonDimensionsCsReq                 = 5529
	PunkLordMonsterInfoScNotify                        = 3233
	NewMailScNotify                                    = 886
	ChessRogueStartScRsp                               = 5471
	GetRogueInfoScRsp                                  = 1848
	RogueModifierSelectCellCsReq                       = 5388
	LogisticsGameScRsp                                 = 4788
	TakePromotionRewardCsReq                           = 339
	ShowNewSupplementVisitorCsReq                      = 3745
	GetAssistListCsReq                                 = 2961
	EnterAdventureCsReq                                = 1334
	TravelBrochureSelectMessageScRsp                   = 6409
	GetAetherDivideInfoScRsp                           = 4868
	GetCurAssistScRsp                                  = 2982
	GetNpcStatusScRsp                                  = 2788
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5630
	RogueNpcDisappearScRsp                             = 5696
	RankUpEquipmentScRsp                               = 529
	GetMonsterResearchActivityDataScRsp                = 2642
	ClockParkGetOngoingScriptInfoCsReq                 = 7286
	SellItemScRsp                                      = 539
	FinishChessRogueNousSubStoryScRsp                  = 5501
	GetPlatformPlayerInfoScRsp                         = 2989
	SpaceZooTakeCsReq                                  = 6706
	AddAvatarScNotify                                  = 396
	SelectChessRogueSubStoryScRsp                      = 5536
	FinishQuestScRsp                                   = 945
	SyncChessRogueNousValueScNotify                    = 5537
	RelicRecommendCsReq                                = 567
	ChangeLineupLeaderScRsp                            = 733
	TakePictureCsReq                                   = 4102
	GetBagScRsp                                        = 548
	GetHeroBasicTypeInfoCsReq                          = 24
	GetChatFriendHistoryCsReq                          = 3919
	ClockParkUseBuffScRsp                              = 7239
	GetPunkLordMonsterDataScRsp                        = 3248
	WolfBroGameActivateBulletCsReq                     = 6559
	SavePointsInfoNotify                               = 1411
	DeleteBlacklistCsReq                               = 2963
	PrivateMsgOfflineUsersScNotify                     = 3988
	TakeActivityExpeditionRewardCsReq                  = 2506
	SubmitEmotionItemCsReq                             = 6302
	StartWolfBroGameScRsp                              = 6548
	UnlockBackGroundMusicCsReq                         = 3102
	GetPlayerReplayInfoScRsp                           = 3588
	ClockParkSyncVirtualItemScNotify                   = 7230
	ClientObjUploadCsReq                               = 64
	GetChallengeRaidInfoScRsp                          = 2219
	GetSocialEventServerCacheCsReq                     = 7005
	GetRogueScoreRewardInfoScRsp                       = 1864
	GetLoginChatInfoCsReq                              = 3933
	QuitLineupCsReq                                    = 719
	GetFirstTalkNpcCsReq                               = 2102
	GetAvatarDataCsReq                                 = 334
	GetBoxingClubInfoScRsp                             = 4248
	StartAetherDivideSceneBattleCsReq                  = 4802
	TakeChallengeRewardScRsp                           = 1759
	DressRelicAvatarCsReq                              = 359
	MonopolyEventLoadUpdateScNotify                    = 7078
	FightTreasureDungeonMonsterCsReq                   = 4495
	GetFantasticStoryActivityDataScRsp                 = 4948
	ChooseBoxingClubResonanceScRsp                     = 4268
	GetAetherDivideChallengeInfoCsReq                  = 4801
	GetRogueEndlessActivityDataScRsp                   = 6048
	GetMissionEventDataCsReq                           = 1233
	StopRogueAdventureRoomScRsp                        = 5601
	TravelBrochureRemovePasterScRsp                    = 6429
	GetBattleCollegeDataScRsp                          = 5748
	GetFriendChallengeDetailCsReq                      = 2975
	SyncRogueAdventureRoomInfoScNotify                 = 5634
	TakePrestigeRewardScRsp                            = 4768
	LockEquipmentScRsp                                 = 509
	SetLineupNameScRsp                                 = 737
	HeliobusStartRaidScRsp                             = 5885
	HeliobusChallengeUpdateScNotify                    = 5856
	EvolveBuildShopAbilityResetCsReq                   = 7144
	ResetMapRotationRegionScRsp                        = 6806
	HeliobusActivityDataScRsp                          = 5848
	GetShareDataScRsp                                  = 4188
	SpringRefreshScRsp                                 = 1437
	GetCurLineupDataScRsp                              = 788
	HeliobusSnsCommentCsReq                            = 5886
	PrestigeLevelUpCsReq                               = 4716
	TelevisionActivityBattleEndScNotify                = 6979
	SetStuffToAreaCsReq                                = 4302
	SubmitOfferingItemScRsp                            = 6937
	GetCurChallengeCsReq                               = 1729
	SetMissionEventProgressScRsp                       = 1263
	SubMissionRewardScNotify                           = 1201
	TravelBrochureApplyPasterScRsp                     = 6443
	MakeMissionDrinkCsReq                              = 6982
	StartFinishSubMissionScNotify                      = 1261
	TakeOfferingRewardCsReq                            = 6940
	TakeTalkRewardCsReq                                = 2162
	ChessRogueUpdateMoneyInfoScNotify                  = 5480
	TakeBpRewardScRsp                                  = 3002
	GetRaidInfoScRsp                                   = 2268
	MonopolyAcceptQuizScRsp                            = 7079
	TravelBrochurePageUnlockScNotify                   = 6462
	CityShopInfoScNotify                               = 1519
	SpaceZooDataScRsp                                  = 6748
	PromoteAvatarScRsp                                 = 343
	ChessRogueGiveUpRollScRsp                          = 5576
	MonopolyLikeScRsp                                  = 7093
	GetStoryLineInfoScRsp                              = 6248
	HeartDialTraceScriptCsReq                          = 6329
	SetLanguageCsReq                                   = 28
	DeactivateFarmElementCsReq                         = 1490
	ReportPlayerCsReq                                  = 2985
	SyncAcceptedPamMissionNotify                       = 4062
	DressAvatarCsReq                                   = 386
	TakeExpeditionRewardCsReq                          = 2519
	TrainVisitorBehaviorFinishCsReq                    = 3734
	TakeMonsterResearchActivityRewardCsReq             = 2616
	SpaceZooDeleteCatScRsp                             = 6729
	OpenTreasureDungeonGridScRsp                       = 4459
	GetRogueAeonInfoScRsp                              = 1827
	QuitTreasureDungeonScRsp                           = 4456
	EvolveBuildReRandomStageScRsp                      = 7106
	EnterFightActivityStageScRsp                       = 3602
	MonopolyGetDailyInitItemScRsp                      = 7050
	MonopolyClickCellScRsp                             = 7027
	RogueModifierDelNotify                             = 5386
	RaidInfoNotify                                     = 2202
	UpgradeAreaScRsp                                   = 4306
	GetRollShopInfoCsReq                               = 6901
	ChessRogueQueryScRsp                               = 5597
	BatchGetQuestDataCsReq                             = 933
	MuseumRandomEventQueryScRsp                        = 4316
	SearchPlayerScRsp                                  = 2928
	TeleportToMissionResetPointScRsp                   = 1228
	GmTalkScRsp                                        = 45
	GetCurSceneInfoCsReq                               = 1419
	CancelExpeditionCsReq                              = 2502
	SyncApplyFriendScNotify                            = 2986
	EnteredSceneChangeScNotify                         = 1450
	GetAllServerPrefsDataCsReq                         = 6134
	TravelBrochurePageResetCsReq                       = 6437
	SyncRogueVirtualItemInfoScNotify                   = 1836
	MonopolySttUpdateScNotify                          = 7077
	EnhanceRogueBuffScRsp                              = 1801
	SpaceZooCatUpdateNotify                            = 6745
	GetRogueTalentInfoScRsp                            = 1838
	QuitWolfBroGameScRsp                               = 6543
	GetWolfBroGameDataScRsp                            = 6529
	HeliobusSnsPostCsReq                               = 5802
	RogueModifierAddNotify                             = 5362
	AcceptActivityExpeditionScRsp                      = 2545
	PunkLordDataChangeNotify                           = 3291
	SetClientRaidTargetCountScRsp                      = 2206
	ChessRogueUpdateDiceInfoScNotify                   = 5526
	MonopolyRollRandomCsReq                            = 7033
	ChangeStoryLineFinishScNotify                      = 6209
	StartPunkLordRaidCsReq                             = 3262
	GetChallengeGroupStatisticsScRsp                   = 1742
	ChessRogueEnterCellCsReq                           = 5518
	GetSaveRaidCsReq                                   = 2233
	ChallengeSettleNotify                              = 1719
	RecoverAllLineupScRsp                              = 1482
	ClockParkQuitScriptScRsp                           = 7206
	DoGachaScRsp                                       = 1988
	RefreshTriggerByClientScNotify                     = 1474
	ChessRogueLeaveScRsp                               = 5531
	ReviveRogueAvatarScRsp                             = 1839
	RemoveStuffFromAreaCsReq                           = 4319
	GetFirstTalkByPerformanceNpcCsReq                  = 2145
	TravelBrochureApplyPasterListCsReq                 = 6416
	SetAssistCsReq                                     = 2991
	GetFriendDevelopmentInfoCsReq                      = 2949
	GetOfferingInfoCsReq                               = 6921
	GetFriendChallengeLineupScRsp                      = 2904
	LeaveMapRotationRegionScRsp                        = 6829
	EvolveBuildStartStageScRsp                         = 7136
	SyncRogueHandbookDataUpdateScNotify                = 5625
	SpaceZooOpCatteryScRsp                             = 6743
	SetPlayerInfoCsReq                                 = 100
	StoryLineTrialAvatarChangeScNotify                 = 6219
	WolfBroGameUseBulletScRsp                          = 6596
	GetAssistHistoryCsReq                              = 2911
	ShowNewSupplementVisitorScRsp                      = 3768
	ChessRogueConfirmRollCsReq                         = 5424
	PickRogueAvatarCsReq                               = 1859
	GetRogueInfoCsReq                                  = 1834
	RestoreWolfBroGameArchiveScRsp                     = 6509
	ChooseBoxingClubStageOptionalBuffCsReq             = 4233
	TakeRogueAeonLevelRewardScRsp                      = 1814
	EvolveBuildShopAbilityDownScRsp                    = 7145
	SetGroupCustomSaveDataCsReq                        = 1449
	GetUpdatedArchiveDataScRsp                         = 2388
	AetherDivideSpiritExpUpScRsp                       = 4856
	MonopolyTakePhaseRewardCsReq                       = 7058
	MissionEventRewardScNotify                         = 1295
	HeliobusInfoChangedScNotify                        = 5868
	UpdateMapRotationDataScNotify                      = 6895
	MonopolyReRollRandomCsReq                          = 7095
	FantasticStoryActivityBattleEndScNotify            = 4909
	GetStuffScNotify                                   = 4386
	TakePromotionRewardScRsp                           = 316
	GetFirstTalkNpcScRsp                               = 2109
	GetRogueTalentInfoCsReq                            = 1832
	MonopolyScrachRaffleTicketScRsp                    = 7014
	MonopolyGuessBuyInformationScRsp                   = 7090
	GameplayCounterRecoverCsReq                        = 1452
	StartWolfBroGameCsReq                              = 6534
	PromoteEquipmentCsReq                              = 562
	ChangeScriptEmotionCsReq                           = 6362
	GetRogueAdventureRoomInfoCsReq                     = 5606
	EnterSectionScRsp                                  = 1428
	TakeAssistRewardCsReq                              = 2979
	GetLineupAvatarDataCsReq                           = 768
	GetChessRogueStoryInfoCsReq                        = 5532
	WolfBroGameExplodeMonsterScRsp                     = 6537
	MonopolyGetDailyInitItemCsReq                      = 7031
	HeliobusUnlockSkillScNotify                        = 5833
	TakeTrialActivityRewardCsReq                       = 2698
	GetHeroBasicTypeInfoScRsp                          = 82
	TakeRogueEndlessActivityPointRewardCsReq           = 6009
	ClockParkUnlockScriptScRsp                         = 7288
	GetGunPlayDataScRsp                                = 4101
	QuitRogueScRsp                                     = 1824
	SwitchAetherDivideLineUpSlotCsReq                  = 4837
	GateServerScNotify                                 = 3
	UnlockTutorialGuideCsReq                           = 1619
	MonopolyContentUpdateScNotify                      = 7061
	MuseumTargetRewardNotify                           = 4341
	PlayerLoginFinishCsReq                             = 15
	StartRaidCsReq                                     = 2234
	GetPunkLordMonsterDataCsReq                        = 3234
	EnterSceneScRsp                                    = 1413
	ChessRogueEnterNextLayerCsReq                      = 5543
	BatchMarkChatEmojiCsReq                            = 3996
	SceneGroupRefreshScNotify                          = 1477
	ReBattleAfterBattleLoseCsNotify                    = 196
	GetBasicInfoCsReq                                  = 66
	ClockParkBattleEndScNotify                         = 7295
	TakeCityShopRewardCsReq                            = 1502
	EnterFantasticStoryActivityStageScRsp              = 4902
	FinishCurTurnScRsp                                 = 4368
	SceneEnterStageScRsp                               = 1456
	RogueEndlessActivityBattleEndScNotify              = 6002
	GetFeverTimeActivityDataCsReq                      = 7156
	MonopolyGiveUpCurContentCsReq                      = 7063
	SceneCastSkillCostMpCsReq                          = 1406
	TakeOffAvatarSkinScRsp                             = 363
	AlleyShipmentEventEffectsScNotify                  = 4761
	ClockParkUnlockTalentCsReq                         = 7202
	GetChallengeScRsp                                  = 1748
	MuseumDispatchFinishedScNotify                     = 4356
	GetGachaInfoScRsp                                  = 1948
	GetExpeditionDataScRsp                             = 2548
	TextJoinSaveScRsp                                  = 3848
	SelectChessRogueNousSubStoryScRsp                  = 5454
	InteractChargerScRsp                               = 6888
	ChessRogueEnterCsReq                               = 5456
	ExchangeRogueBuffWithMiracleCsReq                  = 5637
	HeliobusSnsCommentScRsp                            = 5829
	UpdateMovieRacingDataCsReq                         = 4185
	MonthCardRewardNotify                              = 93
	SelectPhoneThemeCsReq                              = 5109
	FinishTutorialGuideScRsp                           = 1668
	MonopolyTakeRaffleTicketRewardCsReq                = 7084
	SetCurInteractEntityScRsp                          = 1497
	PunkLordMonsterKilledNotify                        = 3228
	TakeOffRelicScRsp                                  = 337
	MakeDrinkScRsp                                     = 6997
	GetMultipleDropInfoScRsp                           = 4648
	TakeApRewardCsReq                                  = 3334
	GetUnlockTeleportScRsp                             = 1469
	MonopolyCellUpdateNotify                           = 7088
	GetMissionDataScRsp                                = 1248
	StartAetherDivideSceneBattleScRsp                  = 4809
	ChessRogueReviveAvatarCsReq                        = 5539
	QuitWolfBroGameCsReq                               = 6519
	AcceptMultipleExpeditionCsReq                      = 2559
	GetMainMissionCustomValueScRsp                     = 1282
	TextJoinQueryScRsp                                 = 3888
	DressRelicAvatarScRsp                              = 395
	LeaveRaidCsReq                                     = 2262
	HeliobusStartRaidCsReq                             = 5830
	SpringRecoverSingleAvatarScRsp                     = 1498
	DailyFirstMeetPamCsReq                             = 3462
	TakeMonsterResearchActivityRewardScRsp             = 2630
	SetFriendMarkScRsp                                 = 2973
	BatchGetQuestDataScRsp                             = 959
	FeverTimeActivityBattleEndScNotify                 = 7153
	MonopolyConditionUpdateScNotify                    = 7032
	FinishRogueDialogueGroupCsReq                      = 1804
	SyncEntityBuffChangeListScNotify                   = 1496
	SetFriendRemarkNameCsReq                           = 2916
	FinishFirstTalkNpcCsReq                            = 2119
	DeleteSocialEventServerCacheCsReq                  = 7012
	ClockParkStartScriptScRsp                          = 7243
	HeliobusSnsLikeScRsp                               = 5843
	EnterRogueMapRoomCsReq                             = 1825
	MatchBoxingClubOpponentCsReq                       = 4262
	EnterTreasureDungeonCsReq                          = 4496
	GetFantasticStoryActivityDataCsReq                 = 4934
	SpaceZooDataCsReq                                  = 6734
	TriggerVoiceScRsp                                  = 4106
	PrepareRogueAdventureRoomCsReq                     = 5648
	TravelBrochureGetPasterScNotify                    = 6496
	AceAntiCheaterCsReq                                = 4
	GetRogueShopBuffInfoCsReq                          = 5609
	SetForbidOtherApplyFriendScRsp                     = 2955
	ExpUpEquipmentScRsp                                = 568
	SetDisplayAvatarScRsp                              = 2809
	RotateMapCsReq                                     = 6819
	MonopolyEventSelectFriendCsReq                     = 7003
	HeartDialTraceScriptScRsp                          = 6345
	ChessRogueRollDiceCsReq                            = 5535
	GetTutorialGuideCsReq                              = 1662
	MultipleDropInfoScNotify                           = 4662
	TakeMultipleExpeditionRewardCsReq                  = 2542
	ChessRogueMoveCellNotify                           = 5586
	EvolveBuildGiveupCsReq                             = 7134
	DelMailCsReq                                       = 802
	ChessRogueUpdateAeonModifierValueScNotify          = 5498
	GetServerPrefsDataScRsp                            = 6188
	MonopolyBuyGoodsCsReq                              = 7016
	GetChatEmojiListScRsp                              = 3929
	SetAssistAvatarScRsp                               = 2896
	GetRogueDialogueEventDataCsReq                     = 1835
	TakePunkLordPointRewardCsReq                       = 3296
	MonopolyClickMbtiReportScRsp                       = 7074
	FinishEmotionDialoguePerformanceCsReq              = 6319
	QuitBattleScRsp                                    = 188
	EnterMapRotationRegionCsReq                        = 6834
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3786
	SelectChatBubbleCsReq                              = 5162
	MissionRewardScNotify                              = 1202
	SyncRogueSeasonFinishScNotify                      = 1808
	MuseumTakeCollectRewardScRsp                       = 4361
	UnlockedAreaMapScNotify                            = 1457
	CancelCacheNotifyCsReq                             = 4186
	GetMuseumInfoScRsp                                 = 4348
	PlayerLoginCsReq                                   = 34
	ChessRogueUpdateLevelBaseInfoScNotify              = 5432
	GetFriendLoginInfoScRsp                            = 2967
	GetChapterCsReq                                    = 402
	GetSaveLogisticsMapScRsp                           = 4791
	GetTrialActivityDataCsReq                          = 2635
	GetTrainVisitorBehaviorCsReq                       = 3762
	GetSpringRecoverDataCsReq                          = 1466
	SharePunkLordMonsterScRsp                          = 3209
	GetActivityScheduleConfigScRsp                     = 2609
	GetFriendRecommendListInfoCsReq                    = 2937
	GetHeartDialInfoCsReq                              = 6334
	GetRogueHandbookDataScRsp                          = 5679
	HandleRogueCommonPendingActionScRsp                = 5675
	ClockParkHandleWaitOperationCsReq                  = 7245
	ChessRogueNousGetRogueTalentInfoScRsp              = 5429
	DiscardRelicScRsp                                  = 590
	GetSceneMapInfoScRsp                               = 1499
	UpdatePlayerSettingScRsp                           = 20
	GetMailCsReq                                       = 834
	FeatureSwitchClosedScNotify                        = 94
	TakeQuestRewardScRsp                               = 988
	TakeKilledPunkLordMonsterScoreCsReq                = 3261
	AetherDivideSpiritInfoScNotify                     = 4863
	SaveLogisticsCsReq                                 = 4701
	GetRndOptionScRsp                                  = 3448
	SpaceZooBornScRsp                                  = 6788
	HandleFriendCsReq                                  = 2929
	GetPlayerDetailInfoScRsp                           = 2988
	RevcMsgScNotify                                    = 3962
	InterruptMissionEventCsReq                         = 1230
	TakeQuestOptionalRewardScRsp                       = 996
	LeaveTrialActivityCsReq                            = 2694
	WolfBroGamePickupBulletCsReq                       = 6506
	StaminaInfoScNotify                                = 69
	GetMissionStatusScRsp                              = 1216
	SetAetherDivideLineUpScRsp                         = 4806
	ClearAetherDividePassiveSkillCsReq                 = 4895
	FinishTalkMissionScRsp                             = 1288
	GetPunkLordBattleRecordCsReq                       = 3297
	SetClientPausedScRsp                               = 1465
	AetherDivideTakeChallengeRewardCsReq               = 4808
	ChessRogueQuitScRsp                                = 5412
	FinishSectionIdCsReq                               = 2719
	TravelBrochureGetDataCsReq                         = 6434
	FinishPerformSectionIdCsReq                        = 2786
	VirtualLineupDestroyNotify                         = 730
	TakeRollShopRewardScRsp                            = 6919
	MonopolyRollDiceCsReq                              = 7009
	GetSingleRedDotParamGroupCsReq                     = 5902
	MarkItemScRsp                                      = 508
	PlayerReturnSignScRsp                              = 4562
	ActivateFarmElementScRsp                           = 1455
	ReturnLastTownCsReq                                = 1416
	StartAlleyEventScRsp                               = 4743
	FinishTutorialScRsp                                = 1629
	PlayerLogoutScRsp                                  = 88
	GetEnteredSceneCsReq                               = 1427
	MonopolyUpgradeAssetCsReq                          = 7085
	SelectRogueDialogueEventCsReq                      = 1815
	StartChallengeCsReq                                = 1762
	GetMarkItemListScRsp                               = 582
	SetSpringRecoverConfigCsReq                        = 1451
	GetAllLineupDataScRsp                              = 716
	CancelMarkItemNotify                               = 554
	SetGameplayBirthdayCsReq                           = 35
	SyncRogueAreaUnlockScNotify                        = 1884
	DestroyItemCsReq                                   = 591
	GetChessRogueNousStoryInfoScRsp                    = 5561
	GetPunkLordDataScRsp                               = 3295
	GetDrinkMakerDataCsReq                             = 6981
	TakeFightActivityRewardScRsp                       = 3619
	SummonPunkLordMonsterCsReq                         = 3219
	StartAetherDivideStageBattleScRsp                  = 4830
	GetMonopolyMbtiReportRewardCsReq                   = 7052
	GetRecyleTimeScRsp                                 = 528
	GetNpcTakenRewardCsReq                             = 2134
	ExchangeGachaCeilingScRsp                          = 1943
	GetGachaCeilingScRsp                               = 1909
	FinishChessRogueSubStoryCsReq                      = 5405
	TrialBackGroundMusicCsReq                          = 3119
	GetRecyleTimeCsReq                                 = 541
	DailyFirstEnterMonopolyActivityScRsp               = 7006
	AddEquipmentScNotify                               = 501
	ChessRogueSelectBpScRsp                            = 5416
	GetNpcMessageGroupScRsp                            = 2748
	MonopolySelectOptionCsReq                          = 7029
	GetKilledPunkLordMonsterDataCsReq                  = 3256
	OpenRogueChestScRsp                                = 1898
	GetVideoVersionKeyCsReq                            = 13
	GetPlayerReturnMultiDropInfoCsReq                  = 4688
	GetAlleyInfoScRsp                                  = 4748
	AddBlacklistScRsp                                  = 2995
	MonopolyCheatDiceScRsp                             = 7028
	TextJoinBatchSaveScRsp                             = 3809
	ClearAetherDividePassiveSkillScRsp                 = 4842
	EndDrinkMakerSequenceCsReq                         = 7000
	TakeLoginActivityRewardCsReq                       = 2662
	GetStrongChallengeActivityDataScRsp                = 6648
	DestroyItemScRsp                                   = 597
	GetJukeboxDataCsReq                                = 3134
	SyncRogueAeonLevelUpRewardScNotify                 = 1820
	EvolveBuildQueryInfoCsReq                          = 7108
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6086
	SyncTaskCsReq                                      = 1209
	SyncChessRogueNousMainStoryScNotify                = 5487
	UpdateRogueAdventureRoomScoreCsReq                 = 5655
	FinishItemIdScRsp                                  = 2709
	MuseumRandomEventSelectScRsp                       = 4385
	ExpUpRelicCsReq                                    = 533
	GetSceneMapInfoCsReq                               = 1470
	SpaceZooMutateScRsp                                = 6709
	ComposeItemScRsp                                   = 506
	EnterChessRogueAeonRoomScRsp                       = 5496
	PlayerLoginScRsp                                   = 48
	MonopolyGameGachaScRsp                             = 7082
	GetMapRotationDataScRsp                            = 6868
	GameplayCounterCountDownCsReq                      = 1458
	ChessRogueLayerAccountInfoNotify                   = 5507
	DrinkMakerUpdateTipsNotify                         = 6983
	GetStageLineupCsReq                                = 734
	AlleyFundsScNotify                                 = 4785
	EnterStrongChallengeActivityStageCsReq             = 6662
	GetChapterScRsp                                    = 409
	MuseumFundsChangedScNotify                         = 4342
	EnhanceCommonRogueBuffScRsp                        = 5656
	ReEnterLastElementStageCsReq                       = 1405
	TakeRogueEventHandbookRewardScRsp                  = 5690
	ChessRogueUpdateActionPointScNotify                = 5469
	UpgradeAreaStatScRsp                               = 4359
	UnlockTutorialCsReq                                = 1602
	GetLevelRewardScRsp                                = 63
	AvatarExpUpScRsp                                   = 388
	MuseumTargetStartNotify                            = 4363
	AlleyEventChangeNotify                             = 4786
	TakeQuestOptionalRewardCsReq                       = 968
	GetWaypointScRsp                                   = 448
	GetMailScRsp                                       = 848
	GetReplayTokenScRsp                                = 3548
	StartBoxingClubBattleCsReq                         = 4202
	ClockParkGetInfoScRsp                              = 7248
	UpdateServerPrefsDataScRsp                         = 6109
	TakeChallengeRaidRewardCsReq                       = 2243
	SceneEntityMoveCsReq                               = 1434
	GetChatFriendHistoryScRsp                          = 3943
	TakePunkLordPointRewardScRsp                       = 3206
	ClockParkUseBuffCsReq                              = 7237
	MonopolyBuyGoodsScRsp                              = 7030
	GetMuseumInfoCsReq                                 = 4334
	SetTurnFoodSwitchCsReq                             = 525
	BattleCollegeDataChangeScNotify                    = 5762
	SetAssistAvatarCsReq                               = 2868
	GetMonsterResearchActivityDataCsReq                = 2695
	SyncRogueDialogueEventDataScNotify                 = 1813
	SetGenderCsReq                                     = 79
	EnterAdventureScRsp                                = 1348
	LogisticsGameCsReq                                 = 4762
	GetChatEmojiListCsReq                              = 3986
	ClockParkGetOngoingScriptInfoScRsp                 = 7229
	DressAvatarSkinScRsp                               = 385
	GetAssistListScRsp                                 = 2918
	AetherDivideTakeChallengeRewardScRsp               = 4854
	SetClientRaidTargetCountCsReq                      = 2296
	EnterMapRotationRegionScRsp                        = 6848
	GetMissionStatusCsReq                              = 1239
	RogueNpcDisappearCsReq                             = 5668
	GetMovieRacingDataCsReq                            = 4116
	GetNpcStatusCsReq                                  = 2762
	MonopolyClickMbtiReportCsReq                       = 7038
	GetTrainVisitorRegisterScRsp                       = 3743
	ChessRoguePickAvatarCsReq                          = 5517
	ChessRogueNousEnableRogueTalentCsReq               = 5570
	GetBattleCollegeDataCsReq                          = 5734
	InteractTreasureDungeonGridCsReq                   = 4437
	SetSignatureScRsp                                  = 2845
	SyncServerSceneChangeNotify                        = 1414
	EquipAetherDividePassiveSkillScRsp                 = 4859
	SelectChessRogueSubStoryCsReq                      = 5600
	GetChallengeRaidInfoCsReq                          = 2209
	UnlockChatBubbleScNotify                           = 5102
	HealPoolInfoNotify                                 = 1475
	SetForbidOtherApplyFriendCsReq                     = 2992
	ReEnterLastElementStageScRsp                       = 1422
	InterruptMissionEventScRsp                         = 1285
	GetPlatformPlayerInfoCsReq                         = 2965
	GetAetherDivideInfoCsReq                           = 4845
	RefreshTriggerByClientScRsp                        = 1438
	RemoveRotaterScRsp                                 = 6837
	ClientObjUploadScRsp                               = 78
	AvatarExpUpCsReq                                   = 362
	GetArchiveDataCsReq                                = 2334
	GetPlayerDetailInfoCsReq                           = 2962
	GetShopListCsReq                                   = 1534
	CancelExpeditionScRsp                              = 2509
	LockRelicScRsp                                     = 542
	FinishRogueDialogueGroupScRsp                      = 1875
	GetPrivateChatHistoryScRsp                         = 3909
	MonopolyEventSelectFriendScRsp                     = 7094
	StartPunkLordRaidScRsp                             = 3288
	SceneCastSkillCsReq                                = 1402
	ChessRogueStartCsReq                               = 5596
	EnterRogueMapRoomScRsp                             = 1900
	EvolveBuildCoinNotify                              = 7104
	SaveLogisticsScRsp                                 = 4741
	SyncRogueCommonActionResultScNotify                = 5667
	EvolveBuildStartLevelScRsp                         = 7147
	DiscardRelicCsReq                                  = 589
	AcceptedPamMissionExpireCsReq                      = 4034
	GetNpcMessageGroupCsReq                            = 2734
	HeliobusSnsReadScRsp                               = 5888
	SpringRecoverSingleAvatarCsReq                     = 1493
	MuseumTargetMissionFinishNotify                    = 4301
	QuitTreasureDungeonCsReq                           = 4485
	GetBagCsReq                                        = 534
	TakePictureScRsp                                   = 4109
	ChangeStoryLineScRsp                               = 6202
	SetRedPointStatusScNotify                          = 84
	FinishChessRogueNousSubStoryCsReq                  = 5411
	GetLoginActivityScRsp                              = 2648
	RaidCollectionDataCsReq                            = 6941
	DressAvatarScRsp                                   = 329
	ChessRogueQueryCsReq                               = 5459
	TeleportToMissionResetPointCsReq                   = 1241
	UpgradeAreaCsReq                                   = 4396
	ChessRogueNousEditDiceScRsp                        = 5482
	DailyActiveInfoNotify                              = 3302
	ExchangeGachaCeilingCsReq                          = 1919
	SceneCastSkillMpUpdateScNotify                     = 1459
	GetRollShopInfoScRsp                               = 6918
	AetherDivideSkillItemScNotify                      = 4818
	RaidKickByServerScNotify                           = 2239
	GetFriendAssistListCsReq                           = 2951
	TravelBrochureSetCustomValueCsReq                  = 6433
	DelMailScRsp                                       = 809
	MonopolyUpgradeAssetScRsp                          = 7056
	ChessRogueSkipTeachingLevelCsReq                   = 5465
	SubmitOrigamiItemScRsp                             = 4159
	GetGachaCeilingCsReq                               = 1902
	SyncClientResVersionScRsp                          = 143
	FinishFirstTalkByPerformanceNpcCsReq               = 2196
	StartCocoonStageCsReq                              = 1408
	DeleteFriendCsReq                                  = 2996
	PlayBackGroundMusicScRsp                           = 3188
	ResetMapRotationRegionCsReq                        = 6896
	ChessRogueEnterCellScRsp                           = 5540
	EvolveBuildShopAbilityUpScRsp                      = 7135
	MuseumInfoChangedScNotify                          = 4395
	MakeDrinkCsReq                                     = 6993
	StartBattleCollegeCsReq                            = 5788
	GetCurChallengeScRsp                               = 1745
	TravelBrochureApplyPasterCsReq                     = 6419
	GetTreasureDungeonActivityDataCsReq                = 4445
	MonopolyGuessChooseCsReq                           = 7100
	AcceptMultipleExpeditionScRsp                      = 2595
	GetMainMissionCustomValueCsReq                     = 1224
	FinishCosumeItemMissionScRsp                       = 1206
	MonopolySelectOptionScRsp                          = 7045
	PlayerReturnInfoQueryCsReq                         = 4586
	ClockParkUnlockScriptCsReq                         = 7262
	GetPhoneDataScRsp                                  = 5148
	GetLevelRewardCsReq                                = 56
	GetChessRogueStoryAeonTalkInfoScRsp                = 5580
	UpgradeAreaStatCsReq                               = 4333
	EnhanceRogueBuffCsReq                              = 1863
	TakeApRewardScRsp                                  = 3348
	ChessRogueQueryBpScRsp                             = 5598
	GetNpcTakenRewardScRsp                             = 2148
	GetShareDataCsReq                                  = 4162
	TakeMailAttachmentCsReq                            = 819
	EnterTelevisionActivityStageCsReq                  = 6977
	MonopolyGameSettleScNotify                         = 7097
	HandleFriendScRsp                                  = 2945
	AetherDivideRefreshEndlessScNotify                 = 4811
	SearchPlayerCsReq                                  = 2941
	MonopolyActionResultScNotify                       = 7062
	HeliobusSelectSkillScRsp                           = 5895
	MonopolyGuessBuyInformationCsReq                   = 7089
	SelectRogueDialogueEventScRsp                      = 1872
	EnterSectionCsReq                                  = 1441
	EnterChessRogueAeonRoomCsReq                       = 5589
	EvolveBuildLeaveCsReq                              = 7117
	QuitLineupScRsp                                    = 743
	DeleteBlacklistScRsp                               = 2901
	EnterFightActivityStageCsReq                       = 3688
	TakeKilledPunkLordMonsterScoreScRsp                = 3218
	ComposeLimitNumCompleteNotify                      = 561
	SpaceZooBornCsReq                                  = 6762
	GetJukeboxDataScRsp                                = 3148
	SetNicknameScRsp                                   = 16
	TakeRogueMiracleHandbookRewardCsReq                = 5700
	GameplayCounterRecoverScRsp                        = 1481
	EvolveBuildFinishScNotify                          = 7107
	GetSpringRecoverDataScRsp                          = 1473
	TakeExpeditionRewardScRsp                          = 2543
	GetVideoVersionKeyScRsp                            = 10
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6043
	SyncTurnFoodNotify                                 = 579
	FinishTalkMissionCsReq                             = 1262
	RefreshAlleyOrderCsReq                             = 4795
	PlayerReturnTakePointRewardScRsp                   = 4509
	FinishSectionIdScRsp                               = 2743
	SpaceZooDeleteCatCsReq                             = 6786
	GetFarmStageGachaInfoScRsp                         = 1388
	JoinLineupScRsp                                    = 709
	GetCurBattleInfoCsReq                              = 102
	DeployRotaterScRsp                                 = 6809
	ChooseBoxingClubResonanceCsReq                     = 4245
	SetHeroBasicTypeScRsp                              = 97
	GetKilledPunkLordMonsterDataScRsp                  = 3263
	ChallengeRaidNotify                                = 2229
	AddBlacklistCsReq                                  = 2959
	GetRogueAeonInfoCsReq                              = 1847
	AcceptExpeditionScRsp                              = 2588
	FightActivityDataChangeScNotify                    = 3662
	TrialBackGroundMusicScRsp                          = 3143
	OpenRogueChestCsReq                                = 1893
	QuitRogueCsReq                                     = 1897
	ComposeSelectedRelicScRsp                          = 563
	GetLineupAvatarDataScRsp                           = 796
	MonopolyConfirmRandomCsReq                         = 7037
	ChessRogueSelectCellCsReq                          = 5434
	AcceptMissionEventCsReq                            = 1242
	PunkLordRaidTimeOutScNotify                        = 3237
	EnableRogueTalentScRsp                             = 1817
	TakeAssistRewardScRsp                              = 2925
	GetSocialEventServerCacheScRsp                     = 7022
	GetChessRogueNousStoryInfoCsReq                    = 5415
	DrinkMakerChallengeScRsp                           = 6990
	SpaceZooOpCatteryCsReq                             = 6719
	AlleyGuaranteedFundsScRsp                          = 4782
	MarkReadMailCsReq                                  = 862
	GetRogueShopMiracleInfoScRsp                       = 5602
	GetRogueBuffEnhanceInfoCsReq                       = 1885
	SelectInclinationTextCsReq                         = 2186
	GetOfferingInfoScRsp                               = 6938
	PunkLordBattleResultScNotify                       = 3285
	BuyBpLevelCsReq                                    = 3009
	SetAssistScRsp                                     = 2997
	PlayerReturnPointChangeScNotify                    = 4588
	UpdateEnergyScNotify                               = 6859
	TravelBrochureApplyPasterListScRsp                 = 6430
	SetPlayerInfoScRsp                                 = 65
	SetIsDisplayAvatarInfoCsReq                        = 2819
	EnterAetherDivideSceneScRsp                        = 4848
	CommonRogueQueryCsReq                              = 5693
	EvolveBuildShopAbilityDownCsReq                    = 7103
	SwapLineupCsReq                                    = 786
	ChessRogueLeaveCsReq                               = 5473
	PlayerReturnForceFinishScNotify                    = 4545
	DoGachaCsReq                                       = 1962
	SpringRecoverScRsp                                 = 1404
	ChessRogueConfirmRollScRsp                         = 5523
	ReplaceLineupScRsp                                 = 756
	SubmitMonsterResearchActivityMaterialScRsp         = 2639
	BatchMarkChatEmojiScRsp                            = 3906
	TravelBrochurePageResetScRsp                       = 6439
	LogisticsDetonateStarSkiffCsReq                    = 4754
	GetFriendChallengeLineupCsReq                      = 2944
	DeleteSummonUnitScRsp                              = 1487
	GetSaveRaidScRsp                                   = 2259
	GetFightActivityDataCsReq                          = 3634
	MonopolyGetRaffleTicketCsReq                       = 7013
	MuseumRandomEventQueryCsReq                        = 4339
	EvolveBuildStartStageCsReq                         = 7141
	SyncChessRogueNousSubStoryScNotify                 = 5420
	ClockParkQuitScriptCsReq                           = 7296
	PlayerLoginFinishScRsp                             = 72
	MonopolyGameRaiseRatioScRsp                        = 7091
	EnterRogueCsReq                                    = 1802
	AetherDivideLineupScNotify                         = 4897
	SyncChessRogueMainStoryFinishScNotify              = 5573
	AcceptActivityExpeditionCsReq                      = 2529
	GetAuthkeyCsReq                                    = 59
	ExpUpEquipmentCsReq                                = 545
	MonopolyRollRandomScRsp                            = 7059
	TravelBrochureUpdatePasterPosCsReq                 = 6445
	ChessRogueCheatRollCsReq                           = 5544
	GetRogueInitialScoreScRsp                          = 1889
	GetSecretKeyInfoScRsp                              = 12
	MarkChatEmojiCsReq                                 = 3945
	ShareCsReq                                         = 4134
	SyncRogueGetItemScNotify                           = 1870
	MonopolyClickCellCsReq                             = 7047
	OpenTreasureDungeonGridCsReq                       = 4433
	MissionAcceptScNotify                              = 1211
	ChessRogueGoAheadCsReq                             = 5458
	SceneEntityTeleportScRsp                           = 1415
	GetChessRogueStoryInfoScRsp                        = 5527
	GmTalkCsReq                                        = 29
	HeroBasicTypeChangedNotify                         = 89
	UnlockSkilltreeScRsp                               = 309
	MonopolyLikeCsReq                                  = 7075
	PlayerSyncScNotify                                 = 634
	SetBoxingClubResonanceLineupScRsp                  = 4206
	GetAllRedDotDataScRsp                              = 5948
	InteractPropCsReq                                  = 1462
	SetHeadIconCsReq                                   = 2862
	ChessRogueGiveUpRollCsReq                          = 5558
	EnhanceChessRogueBuffScRsp                         = 5468
	ChessRogueGiveUpScRsp                              = 5511
	SendMsgScRsp                                       = 3948
	StartTrialActivityScRsp                            = 2603
	PromoteAvatarCsReq                                 = 319
	ReportPlayerScRsp                                  = 2956
	MonopolyQuizDurationChangeScNotify                 = 7092
	TakeTalkRewardScRsp                                = 2188
	EnterTrialActivityStageCsReq                       = 2675
	GetAllServerPrefsDataScRsp                         = 6148
	GetMbtiReportScRsp                                 = 7049
	SyncRogueStatusScNotify                            = 1860
	ApplyFriendScRsp                                   = 2943
	GetMonopolyInfoScRsp                               = 7048
	StartTimedCocoonStageScRsp                         = 1423
	GetDailyActiveInfoScRsp                            = 3388
	StartRogueCsReq                                    = 1862
	TakeAllRewardScRsp                                 = 3086
	TravelBrochureSetPageDescStatusCsReq               = 6495
	GetAllSaveRaidCsReq                                = 2295
	EnterRogueEndlessActivityStageCsReq                = 6062
	GetChallengeCsReq                                  = 1734
	TakeOfferingRewardScRsp                            = 6939
	SyncDeleteFriendScNotify                           = 2933
	GetCurLineupDataCsReq                              = 762
	ExchangeStaminaCsReq                               = 6
	BattleLogReportCsReq                               = 129
	MonopolyDailySettleScNotify                        = 7035
	BuyGoodsCsReq                                      = 1562
	FinishPlotScRsp                                    = 1148
	BuyRogueShopBuffScRsp                              = 5645
	PrestigeLevelUpScRsp                               = 4730
	ChessRogueReviveAvatarScRsp                        = 5470
	UpdateGunPlayDataCsReq                             = 4141
	StopRogueAdventureRoomCsReq                        = 5663
	GetRaidInfoCsReq                                   = 2245
	SyncRogueAeonScNotify                              = 1810
	SetStuffToAreaScRsp                                = 4309
	TakeOffAvatarSkinCsReq                             = 356
	MonopolyGetRafflePoolInfoCsReq                     = 7007
	GetRogueScoreRewardInfoCsReq                       = 1858
	AlleyOrderChangedScNotify                          = 4737
	SceneEnterStageCsReq                               = 1485
	TakePrestigeRewardCsReq                            = 4745
	EvolveBuildTakeExpRewardScRsp                      = 7121
	PrepareRogueAdventureRoomScRsp                     = 5662
	TakeBpRewardCsReq                                  = 3088
	AlleyPlacingGameScRsp                              = 4706
	TakeChallengeRewardCsReq                           = 1733
	StartTimedFarmElementCsReq                         = 1436
	GetExhibitScNotify                                 = 4329
	UpdateRedDotDataScRsp                              = 5988
	GetAetherDivideChallengeInfoScRsp                  = 4841
	GetRogueShopBuffInfoScRsp                          = 5619
	TelevisionActivityDataChangeScNotify               = 6973
	GetLoginChatInfoScRsp                              = 3959
	MakeMissionDrinkScRsp                              = 6996
	GroupStateChangeCsReq                              = 1476
	GetPlayerReplayInfoCsReq                           = 3562
	UnlockBackGroundMusicScRsp                         = 3109
	ExchangeRogueRewardKeyScRsp                        = 1849
	GetRogueAdventureRoomInfoScRsp                     = 5633
	TakeActivityExpeditionRewardScRsp                  = 2533
	GetBoxingClubInfoCsReq                             = 4234
	GetUpdatedArchiveDataCsReq                         = 2362
	ReserveStaminaExchangeCsReq                        = 14
	GetFriendChallengeDetailScRsp                      = 2993
	EvolveBuildShopAbilityResetScRsp                   = 7120
	AcceptMainMissionScRsp                             = 1297
	FinishAeonDialogueGroupScRsp                       = 1850
	TakeOffEquipmentScRsp                              = 368
	CancelActivityExpeditionScRsp                      = 2596
	TakeChapterRewardScRsp                             = 486
	SetLineupNameCsReq                                 = 742
	LeaveRaidScRsp                                     = 2288
	SpringRefreshCsReq                                 = 1442
	SetCurWaypointScRsp                                = 488
	EnterFantasticStoryActivityStageCsReq              = 4988
	GetMissionEventDataScRsp                           = 1259
	DailyRefreshNotify                                 = 51
	TravelBrochureRemovePasterCsReq                    = 6486
	GmTalkScNotify                                     = 43
	UseTreasureDungeonItemCsReq                        = 4416
	SelectChessRogueNousSubStoryCsReq                  = 5484
	LeaveRogueCsReq                                    = 1819
	PlayerHeartBeatCsReq                               = 71
	GetRogueEndlessActivityDataCsReq                   = 6034
	AetherDivideRefreshEndlessCsReq                    = 4824
	SetCurInteractEntityCsReq                          = 1491
	SpaceZooTakeScRsp                                  = 6733
	ChessRogueReRollDiceScRsp                          = 5500
	SecurityReportCsReq                                = 4145
	BuyRogueShopMiracleScRsp                           = 5686
	GetAvatarDataScRsp                                 = 348
	RankUpEquipmentCsReq                               = 586
	AntiAddictScNotify                                 = 37
	StartAetherDivideChallengeBattleCsReq              = 4819
	TextJoinSaveCsReq                                  = 3834
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
	// SyncPlayerOnlineDataNotify   = 11001
	// PlayerLoginNotify            = 11002
	// NodeToGsPlayerLogoutNotify   = 11003
	GameToGatePlayerLogoutNotify = 11004

	GmGive       = 12001
	GmWorldLevel = 12002
	DelItem      = 12003
	MaxCurAvatar = 12004
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(SpringRecoverCsReq, func() any { return new(proto.SpringRecoverCsReq) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(SubmitMonsterResearchActivityMaterialCsReq, func() any { return new(proto.SubmitMonsterResearchActivityMaterialCsReq) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(GetMbtiReportCsReq, func() any { return new(proto.GetMbtiReportCsReq) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(NewAssistHistoryNotify, func() any { return new(proto.NewAssistHistoryNotify) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(GetPlayerReturnMultiDropInfoScRsp, func() any { return new(proto.GetPlayerReturnMultiDropInfoScRsp) })
	c.regMsg(MuseumTakeCollectRewardCsReq, func() any { return new(proto.MuseumTakeCollectRewardCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(StrongChallengeActivityBattleEndScNotify, func() any { return new(proto.StrongChallengeActivityBattleEndScNotify) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(SubmitOfferingItemCsReq, func() any { return new(proto.SubmitOfferingItemCsReq) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(SpaceZooExchangeItemScRsp, func() any { return new(proto.SpaceZooExchangeItemScRsp) })
	c.regMsg(MonopolySocialEventEffectScNotify, func() any { return new(proto.MonopolySocialEventEffectScNotify) })
	c.regMsg(SetSpringRecoverConfigScRsp, func() any { return new(proto.SetSpringRecoverConfigScRsp) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(MonopolyMoveCsReq, func() any { return new(proto.MonopolyMoveCsReq) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(MonopolyGameRaiseRatioCsReq, func() any { return new(proto.MonopolyGameRaiseRatioCsReq) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(WolfBroGameActivateBulletScRsp, func() any { return new(proto.WolfBroGameActivateBulletScRsp) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(BoxingClubChallengeUpdateScNotify, func() any { return new(proto.BoxingClubChallengeUpdateScNotify) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(QueryProductInfoCsReq, func() any { return new(proto.QueryProductInfoCsReq) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
	c.regMsg(GetMapRotationDataCsReq, func() any { return new(proto.GetMapRotationDataCsReq) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(RaidCollectionDataScNotify, func() any { return new(proto.RaidCollectionDataScNotify) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(GiveUpBoxingClubChallengeScRsp, func() any { return new(proto.GiveUpBoxingClubChallengeScRsp) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(WolfBroGameExplodeMonsterCsReq, func() any { return new(proto.WolfBroGameExplodeMonsterCsReq) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(DailyTaskDataScNotify, func() any { return new(proto.DailyTaskDataScNotify) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(GetMonopolyDailyReportScRsp, func() any { return new(proto.GetMonopolyDailyReportScRsp) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(RogueModifierUpdateNotify, func() any { return new(proto.RogueModifierUpdateNotify) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(RogueModifierSelectCellScRsp, func() any { return new(proto.RogueModifierSelectCellScRsp) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(RogueModifierStageStartNotify, func() any { return new(proto.RogueModifierStageStartNotify) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(SetHeroBasicTypeCsReq, func() any { return new(proto.SetHeroBasicTypeCsReq) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(GetMonopolyDailyReportCsReq, func() any { return new(proto.GetMonopolyDailyReportCsReq) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(CurAssistChangedNotify, func() any { return new(proto.CurAssistChangedNotify) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(MultipleDropInfoNotify, func() any { return new(proto.MultipleDropInfoNotify) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(GiveUpBoxingClubChallengeCsReq, func() any { return new(proto.GiveUpBoxingClubChallengeCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(BoxingClubRewardScNotify, func() any { return new(proto.BoxingClubRewardScNotify) })
	c.regMsg(RemoveRotaterCsReq, func() any { return new(proto.RemoveRotaterCsReq) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(UpdatePlayerSettingCsReq, func() any { return new(proto.UpdatePlayerSettingCsReq) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(GetMonopolyFriendRankingListCsReq, func() any { return new(proto.GetMonopolyFriendRankingListCsReq) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(AcceptMissionEventScRsp, func() any { return new(proto.AcceptMissionEventScRsp) })
	c.regMsg(GetMonopolyInfoCsReq, func() any { return new(proto.GetMonopolyInfoCsReq) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(StartTimedCocoonStageCsReq, func() any { return new(proto.StartTimedCocoonStageCsReq) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(MonopolyMoveScRsp, func() any { return new(proto.MonopolyMoveScRsp) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(MonopolyLikeScNotify, func() any { return new(proto.MonopolyLikeScNotify) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(SyncRoguePickAvatarInfoScNotify, func() any { return new(proto.SyncRoguePickAvatarInfoScNotify) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(GetMultipleDropInfoCsReq, func() any { return new(proto.GetMultipleDropInfoCsReq) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(SpaceZooExchangeItemCsReq, func() any { return new(proto.SpaceZooExchangeItemCsReq) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(RaidCollectionDataScRsp, func() any { return new(proto.RaidCollectionDataScRsp) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(StartTimedFarmElementScRsp, func() any { return new(proto.StartTimedFarmElementScRsp) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupCsReq, func() any { return new(proto.SetBoxingClubResonanceLineupCsReq) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(GetMonsterResearchActivityDataScRsp, func() any { return new(proto.GetMonsterResearchActivityDataScRsp) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(SpaceZooTakeCsReq, func() any { return new(proto.SpaceZooTakeCsReq) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(GetHeroBasicTypeInfoCsReq, func() any { return new(proto.GetHeroBasicTypeInfoCsReq) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(ClientObjUploadCsReq, func() any { return new(proto.ClientObjUploadCsReq) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(SubmitOfferingItemScRsp, func() any { return new(proto.SubmitOfferingItemScRsp) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(TakeOfferingRewardCsReq, func() any { return new(proto.TakeOfferingRewardCsReq) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(MonopolyLikeScRsp, func() any { return new(proto.MonopolyLikeScRsp) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(TakeMonsterResearchActivityRewardCsReq, func() any { return new(proto.TakeMonsterResearchActivityRewardCsReq) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(MonopolyClickCellScRsp, func() any { return new(proto.MonopolyClickCellScRsp) })
	c.regMsg(RogueModifierDelNotify, func() any { return new(proto.RogueModifierDelNotify) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(EnteredSceneChangeScNotify, func() any { return new(proto.EnteredSceneChangeScNotify) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(MonopolySttUpdateScNotify, func() any { return new(proto.MonopolySttUpdateScNotify) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(GetOfferingInfoCsReq, func() any { return new(proto.GetOfferingInfoCsReq) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffCsReq, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(GetUpdatedArchiveDataScRsp, func() any { return new(proto.GetUpdatedArchiveDataScRsp) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(MissionEventRewardScNotify, func() any { return new(proto.MissionEventRewardScNotify) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(UpdateMapRotationDataScNotify, func() any { return new(proto.UpdateMapRotationDataScNotify) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(TakeAssistRewardCsReq, func() any { return new(proto.TakeAssistRewardCsReq) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(WolfBroGameExplodeMonsterScRsp, func() any { return new(proto.WolfBroGameExplodeMonsterScRsp) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(GetHeroBasicTypeInfoScRsp, func() any { return new(proto.GetHeroBasicTypeInfoScRsp) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(MuseumTargetRewardNotify, func() any { return new(proto.MuseumTargetRewardNotify) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(GetMultipleDropInfoScRsp, func() any { return new(proto.GetMultipleDropInfoScRsp) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(SpringRecoverSingleAvatarScRsp, func() any { return new(proto.SpringRecoverSingleAvatarScRsp) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(TakeMonsterResearchActivityRewardScRsp, func() any { return new(proto.TakeMonsterResearchActivityRewardScRsp) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(MonopolyConditionUpdateScNotify, func() any { return new(proto.MonopolyConditionUpdateScNotify) })
	c.regMsg(FinishRogueDialogueGroupCsReq, func() any { return new(proto.FinishRogueDialogueGroupCsReq) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(MatchBoxingClubOpponentCsReq, func() any { return new(proto.MatchBoxingClubOpponentCsReq) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(SetForbidOtherApplyFriendScRsp, func() any { return new(proto.SetForbidOtherApplyFriendScRsp) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(MonopolyEventSelectFriendCsReq, func() any { return new(proto.MonopolyEventSelectFriendCsReq) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(QuitBattleScRsp, func() any { return new(proto.QuitBattleScRsp) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(SyncRogueSeasonFinishScNotify, func() any { return new(proto.SyncRogueSeasonFinishScNotify) })
	c.regMsg(MuseumTakeCollectRewardScRsp, func() any { return new(proto.MuseumTakeCollectRewardScRsp) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(GetSpringRecoverDataCsReq, func() any { return new(proto.GetSpringRecoverDataCsReq) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(UpdatePlayerSettingScRsp, func() any { return new(proto.UpdatePlayerSettingScRsp) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(InterruptMissionEventCsReq, func() any { return new(proto.InterruptMissionEventCsReq) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(SelectRogueDialogueEventCsReq, func() any { return new(proto.SelectRogueDialogueEventCsReq) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(SetSpringRecoverConfigCsReq, func() any { return new(proto.SetSpringRecoverConfigCsReq) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(MonopolySelectOptionCsReq, func() any { return new(proto.MonopolySelectOptionCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(GetPlayerReturnMultiDropInfoCsReq, func() any { return new(proto.GetPlayerReturnMultiDropInfoCsReq) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(MonopolyCheatDiceScRsp, func() any { return new(proto.MonopolyCheatDiceScRsp) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(StartBoxingClubBattleCsReq, func() any { return new(proto.StartBoxingClubBattleCsReq) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(GetMonsterResearchActivityDataCsReq, func() any { return new(proto.GetMonsterResearchActivityDataCsReq) })
	c.regMsg(SyncRogueDialogueEventDataScNotify, func() any { return new(proto.SyncRogueDialogueEventDataScNotify) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(GetBattleCollegeDataCsReq, func() any { return new(proto.GetBattleCollegeDataCsReq) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(HealPoolInfoNotify, func() any { return new(proto.HealPoolInfoNotify) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(RemoveRotaterScRsp, func() any { return new(proto.RemoveRotaterScRsp) })
	c.regMsg(ClientObjUploadScRsp, func() any { return new(proto.ClientObjUploadScRsp) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(FinishRogueDialogueGroupScRsp, func() any { return new(proto.FinishRogueDialogueGroupScRsp) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(MonopolyEventSelectFriendScRsp, func() any { return new(proto.MonopolyEventSelectFriendScRsp) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(AcceptedPamMissionExpireCsReq, func() any { return new(proto.AcceptedPamMissionExpireCsReq) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(SpringRecoverSingleAvatarCsReq, func() any { return new(proto.SpringRecoverSingleAvatarCsReq) })
	c.regMsg(MuseumTargetMissionFinishNotify, func() any { return new(proto.MuseumTargetMissionFinishNotify) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(RaidCollectionDataCsReq, func() any { return new(proto.RaidCollectionDataCsReq) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(MonopolySelectOptionScRsp, func() any { return new(proto.MonopolySelectOptionScRsp) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(MonopolyActionResultScNotify, func() any { return new(proto.MonopolyActionResultScNotify) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(SelectRogueDialogueEventScRsp, func() any { return new(proto.SelectRogueDialogueEventScRsp) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(GetSpringRecoverDataScRsp, func() any { return new(proto.GetSpringRecoverDataScRsp) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(ChooseBoxingClubResonanceCsReq, func() any { return new(proto.ChooseBoxingClubResonanceCsReq) })
	c.regMsg(SetHeroBasicTypeScRsp, func() any { return new(proto.SetHeroBasicTypeScRsp) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(AcceptMissionEventCsReq, func() any { return new(proto.AcceptMissionEventCsReq) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(TakeAssistRewardScRsp, func() any { return new(proto.TakeAssistRewardScRsp) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(SpringRecoverScRsp, func() any { return new(proto.SpringRecoverScRsp) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(SubmitMonsterResearchActivityMaterialScRsp, func() any { return new(proto.SubmitMonsterResearchActivityMaterialScRsp) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(MonopolyClickCellCsReq, func() any { return new(proto.MonopolyClickCellCsReq) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(HeroBasicTypeChangedNotify, func() any { return new(proto.HeroBasicTypeChangedNotify) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(MonopolyLikeCsReq, func() any { return new(proto.MonopolyLikeCsReq) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(GetMbtiReportScRsp, func() any { return new(proto.GetMbtiReportScRsp) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(GetMonopolyInfoScRsp, func() any { return new(proto.GetMonopolyInfoScRsp) })
	c.regMsg(StartTimedCocoonStageScRsp, func() any { return new(proto.StartTimedCocoonStageScRsp) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(TakeOfferingRewardScRsp, func() any { return new(proto.TakeOfferingRewardScRsp) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(ExchangeStaminaCsReq, func() any { return new(proto.ExchangeStaminaCsReq) })
	c.regMsg(MonopolyDailySettleScNotify, func() any { return new(proto.MonopolyDailySettleScNotify) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(SyncRogueAeonScNotify, func() any { return new(proto.SyncRogueAeonScNotify) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(GetBoxingClubInfoCsReq, func() any { return new(proto.GetBoxingClubInfoCsReq) })
	c.regMsg(GetUpdatedArchiveDataCsReq, func() any { return new(proto.GetUpdatedArchiveDataCsReq) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	// server
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

	// c.regMsg(PlayerLoginNotify, func() any { return new(spb.PlayerLoginNotify) })
	// c.regMsg(NodeToGsPlayerLogoutNotify, func() any { return new(spb.NodeToGsPlayerLogoutNotify) })
	c.regMsg(GameToNodePingReq, func() any { return new(spb.GameToNodePingReq) })
	c.regMsg(GameToNodePingRsp, func() any { return new(spb.GameToNodePingRsp) })
	c.regMsg(GateToNodePingReq, func() any { return new(spb.GateToNodePingReq) })
	c.regMsg(GateToNodePingRsp, func() any { return new(spb.GateToNodePingRsp) })
}
