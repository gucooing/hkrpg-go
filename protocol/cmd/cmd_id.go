package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	NONE                                               = 0
	SpringRecoverCsReq                                 = 1443
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3770
	FinishEmotionDialoguePerformanceScRsp              = 6326
	SelectChatBubbleScRsp                              = 5145
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5637
	SubmitMonsterResearchActivityMaterialCsReq         = 2691
	EnterAetherDivideSceneCsReq                        = 4810
	HeliobusActivityDataCsReq                          = 5810
	CommonRogueQueryScRsp                              = 5622
	GetRogueShopMiracleInfoCsReq                       = 5645
	ApplyFriendCsReq                                   = 2977
	GetHeartDialInfoScRsp                              = 6333
	ComposeLimitNumUpdateNotify                        = 567
	GetCurAssistCsReq                                  = 2932
	TakeRollShopRewardCsReq                            = 6911
	RelicRecommendScRsp                                = 523
	ChangeLineupLeaderCsReq                            = 795
	TrialActivityDataChangeScNotify                    = 2668
	DrinkMakerChallengeCsReq                           = 6988
	GetMbtiReportCsReq                                 = 7046
	AlleyGuaranteedFundsCsReq                          = 4732
	UpdateFeatureSwitchScNotify                        = 31
	FinishQuestCsReq                                   = 970
	ChessRogueRollDiceScRsp                            = 5444
	GetAuthkeyScRsp                                    = 99
	ActivateFarmElementCsReq                           = 1423
	SetAetherDivideLineUpCsReq                         = 4858
	StartAetherDivideChallengeBattleScRsp              = 4826
	StartRogueScRsp                                    = 1845
	UnlockAvatarSkinScNotify                           = 365
	SellItemCsReq                                      = 591
	SwapLineupScRsp                                    = 770
	TakeAllApRewardCsReq                               = 3360
	GetChallengeGroupStatisticsCsReq                   = 1799
	RetcodeNotify                                      = 22
	FinishPerformSectionIdScRsp                        = 2770
	AetherDivideFinishChallengeScNotify                = 4885
	SyncLineupNotify                                   = 703
	UseItemScRsp                                       = 526
	GetFarmStageGachaInfoCsReq                         = 1342
	UseTreasureDungeonItemScRsp                        = 4494
	AcceptMainMissionCsReq                             = 1238
	GetWolfBroGameDataCsReq                            = 6600
	JoinLineupCsReq                                    = 716
	DelSaveRaidScNotify                                = 2291
	InteractPropScRsp                                  = 1445
	DeleteSocialEventServerCacheScRsp                  = 7079
	EnableRogueTalentCsReq                             = 1815
	CurTrialActivityScNotify                           = 2628
	UnlockPhoneThemeScNotify                           = 5126
	SubmitEmotionItemScRsp                             = 6360
	UpdateRogueAdventureRoomScoreScRsp                 = 5629
	NewAssistHistoryNotify                             = 2987
	AlleyShipUsedCountScNotify                         = 4712
	RecoverAllLineupCsReq                              = 1432
	HeliobusSnsPostScRsp                               = 5860
	SetNicknameCsReq                                   = 66
	SharePunkLordMonsterCsReq                          = 3216
	GetChessRogueBuffEnhanceInfoScRsp                  = 5418
	AlleyPlacingGameCsReq                              = 4758
	TakeRogueEventHandbookRewardCsReq                  = 5649
	GetPlayerReturnMultiDropInfoScRsp                  = 4616
	MuseumTakeCollectRewardCsReq                       = 4385
	MonopolyGameCreateScNotify                         = 7018
	MonopolyGetRafflePoolInfoScRsp                     = 7088
	DailyFirstEnterMonopolyActivityCsReq               = 7058
	GetActivityScheduleConfigCsReq                     = 2616
	DeleteSummonUnitCsReq                              = 1441
	RefreshAlleyOrderScRsp                             = 4776
	FinishChessRogueSubStoryScRsp                      = 5478
	GetCurSceneInfoScRsp                               = 1426
	GetFightActivityDataScRsp                          = 3633
	QuitBattleCsReq                                    = 142
	FinishFirstTalkNpcScRsp                            = 2126
	StartCocoonStageScRsp                              = 1487
	BattleLogReportScRsp                               = 103
	TreasureDungeonFinishScNotify                      = 4433
	ChessRogueEnterScRsp                               = 5582
	SetMissionEventProgressCsReq                       = 1208
	TreasureDungeonDataScNotify                        = 4410
	StrongChallengeActivityBattleEndScNotify           = 6616
	StartAlleyEventCsReq                               = 4777
	ChessRogueCellUpdateNotify                         = 5477
	QuestRecordScNotify                                = 1000
	MonopolyCheatDiceCsReq                             = 7069
	BuyRogueShopMiracleCsReq                           = 5626
	EnterRogueEndlessActivityStageScRsp                = 6045
	StartAetherDivideStageBattleCsReq                  = 4837
	GetLoginActivityCsReq                              = 2610
	UpdateServerPrefsDataCsReq                         = 6116
	StartBoxingClubBattleScRsp                         = 4260
	GetFeverTimeActivityDataScRsp                      = 7151
	GetStoryLineInfoCsReq                              = 6210
	ChessRogueUpdateReviveInfoScNotify                 = 5554
	GetWaypointCsReq                                   = 410
	ReturnLastTownScRsp                                = 1494
	SubmitOfferingItemCsReq                            = 6921
	LeaveRogueScRsp                                    = 1826
	FinishTutorialGuideCsReq                           = 1603
	SpaceZooExchangeItemScRsp                          = 6758
	MonopolyAcceptQuizCsReq                            = 7087
	MonopolySocialEventEffectScNotify                  = 7025
	SetSpringRecoverConfigScRsp                        = 1448
	TravelBrochureSetPageDescStatusScRsp               = 6476
	GetAlleyInfoCsReq                                  = 4710
	GetAllSaveRaidScRsp                                = 2276
	FightTreasureDungeonMonsterScRsp                   = 4476
	FinishCosumeItemMissionCsReq                       = 1258
	WolfBroGamePickupBulletScRsp                       = 6536
	ExchangeStaminaScRsp                               = 36
	ChessRogueNousDiceSurfaceUnlockNotify              = 5590
	FinishChapterScNotify                              = 4942
	GetPunkLordBattleRecordScRsp                       = 3232
	GetShopListScRsp                                   = 1533
	LogisticsScoreRewardSyncInfoScNotify               = 4718
	GetStageLineupScRsp                                = 733
	GetTelevisionActivityDataCsReq                     = 6979
	LockRelicCsReq                                     = 599
	EnterSceneCsReq                                    = 1419
	SetClientPausedCsReq                               = 1497
	PlayerLogoutCsReq                                  = 42
	SpaceZooMutateCsReq                                = 6716
	BattlePassInfoNotify                               = 3010
	MonopolyMoveCsReq                                  = 7026
	ReserveStaminaExchangeScRsp                        = 20
	MarkItemCsReq                                      = 501
	StartFinishMainMissionScNotify                     = 1267
	FinishAeonDialogueGroupCsReq                       = 1864
	HeliobusSnsUpdateScNotify                          = 5803
	MonopolyGameRaiseRatioCsReq                        = 7067
	MonopolyRollDiceScRsp                              = 7077
	UpdateFloorSavedValueNotify                        = 1488
	WolfBroGameActivateBulletScRsp                     = 6599
	TrainVisitorBehaviorFinishScRsp                    = 3733
	LockEquipmentCsReq                                 = 516
	AetherDivideTainerInfoScNotify                     = 4883
	TextJoinBatchSaveCsReq                             = 3816
	ExpUpRelicScRsp                                    = 502
	AlleyShipUnlockScNotify                            = 4706
	DeactivateFarmElementScRsp                         = 1444
	ChessRogueEnterNextLayerScRsp                      = 5437
	SyncRogueFinishScNotify                            = 1836
	HeliobusSnsReadCsReq                               = 5842
	ChessRogueNousEditDiceCsReq                        = 5450
	LeaveTrialActivityScRsp                            = 2625
	SetGameplayBirthdayScRsp                           = 43
	GroupStateChangeScNotify                           = 1417
	MonopolyTakePhaseRewardScRsp                       = 7052
	SetSignatureCsReq                                  = 2870
	EquipAetherDividePassiveSkillCsReq                 = 4836
	SetLanguageScRsp                                   = 83
	GetLevelRewardTakenListCsReq                       = 94
	InteractTreasureDungeonGridScRsp                   = 4466
	GetRndOptionCsReq                                  = 3410
	ChessRogueNousDiceUpdateNotify                     = 5465
	TakeQuestRewardCsReq                               = 942
	TravelBrochureSelectMessageCsReq                   = 6416
	ChessRoguePickAvatarScRsp                          = 5595
	LeaveAetherDivideSceneScRsp                        = 4845
	MuseumRandomEventStartScNotify                     = 4391
	UpdateTrackMainMissionIdCsReq                      = 1287
	RemoveStuffFromAreaScRsp                           = 4326
	SetGenderScRsp                                     = 18
	MonopolyScrachRaffleTicketCsReq                    = 7034
	BoxingClubChallengeUpdateScNotify                  = 4270
	SyncRogueCommonVirtualItemInfoScNotify             = 5624
	GetBasicInfoScRsp                                  = 24
	AetherDivideSpiritExpUpCsReq                       = 4896
	GetFirstTalkByPerformanceNpcScRsp                  = 2111
	TrainRefreshTimeNotify                             = 3716
	ChooseBoxingClubStageOptionalBuffScRsp             = 4202
	TakeCityShopRewardScRsp                            = 1560
	ChessRogueUpdateUnlockLevelScNotify                = 5572
	ReviveRogueAvatarCsReq                             = 1891
	SetTurnFoodSwitchScRsp                             = 597
	SceneEntityMoveScNotify                            = 1403
	LeaveMapRotationRegionCsReq                        = 6900
	TakeTrialActivityRewardScRsp                       = 2646
	ClientObjDownloadDataScNotify                      = 63
	ExchangeHcoinScRsp                                 = 596
	DressAvatarSkinCsReq                               = 394
	QueryProductInfoCsReq                              = 39
	TakeChallengeRaidRewardScRsp                       = 2300
	GetReplayTokenCsReq                                = 3510
	DoGachaInRollShopCsReq                             = 6901
	ChessRogueNousEnableRogueTalentScRsp               = 5562
	EnhanceCommonRogueBuffCsReq                        = 5696
	EnterStrongChallengeActivityStageScRsp             = 6645
	WolfBroGameUseBulletCsReq                          = 6511
	GetMapRotationDataCsReq                            = 6803
	PromoteEquipmentScRsp                              = 545
	RaidCollectionDataScNotify                         = 6941
	RankUpAvatarScRsp                                  = 336
	GetMovieRacingDataScRsp                            = 4194
	SceneEntityMoveScRsp                               = 1433
	AlleyEventEffectNotify                             = 4770
	GetCurBattleInfoScRsp                              = 160
	ChangeStoryLineCsReq                               = 6245
	TriggerVoiceCsReq                                  = 4158
	GameplayCounterUpdateScNotify                      = 1404
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5570
	GiveUpBoxingClubChallengeScRsp                     = 4226
	AcceptedPamMissionExpireScRsp                      = 4033
	RegionStopScNotify                                 = 76
	TakeRogueEndlessActivityPointRewardScRsp           = 6077
	GetChessRogueStoryAeonTalkInfoCsReq                = 5524
	DailyTaskDataScNotify                              = 1226
	GetMonopolyFriendRankingListScRsp                  = 7068
	GetMonopolyDailyReportScRsp                        = 7040
	BuyNpcStuffScRsp                                   = 4345
	GetMissionDataCsReq                                = 1210
	GetFriendApplyListInfoCsReq                        = 2916
	GetUnlockTeleportCsReq                             = 1420
	ChessRogueFinishCurRoomNotify                      = 5402
	TextJoinQueryCsReq                                 = 3842
	RotateMapScRsp                                     = 6826
	LeaveChallengeScRsp                                = 1760
	TakeRogueAeonLevelRewardCsReq                      = 1834
	MonopolyGuessDrawScNotify                          = 7044
	MonopolyGameGachaCsReq                             = 7032
	SummonPunkLordMonsterScRsp                         = 3226
	LeaveMapRotationRegionScNotify                     = 6836
	SetGroupCustomSaveDataScRsp                        = 1480
	SwitchLineupIndexCsReq                             = 702
	FinishItemIdCsReq                                  = 2716
	MuseumRandomEventSelectCsReq                       = 4394
	PlayerGetTokenScRsp                                = 60
	RefreshTriggerByClientCsReq                        = 1461
	EnterTreasureDungeonScRsp                          = 4495
	GetQuestRecordScRsp                                = 926
	RogueModifierUpdateNotify                          = 5326
	RestoreWolfBroGameArchiveCsReq                     = 6516
	CommonRogueUpdateScNotify                          = 5646
	MissionGroupWarnScNotify                           = 1211
	GetEnteredSceneScRsp                               = 1464
	ArchiveWolfBroGameScRsp                            = 6545
	SyncTaskScRsp                                      = 1277
	PlayerReturnTakeRewardScRsp                        = 4526
	PickRogueAvatarScRsp                               = 1899
	ComposeItemCsReq                                   = 558
	SetFriendMarkCsReq                                 = 2929
	SyncRogueRewardInfoScNotify                        = 1882
	UnlockHeadIconScNotify                             = 2900
	SyncHandleFriendScNotify                           = 2911
	GetGachaInfoCsReq                                  = 1910
	RogueModifierSelectCellScRsp                       = 5316
	HeliobusUpgradeLevelScRsp                          = 5895
	TakeMultipleExpeditionRewardScRsp                  = 2591
	TakeLoginActivityRewardScRsp                       = 2645
	PVEBattleResultCsReq                               = 110
	MonopolyGuessChooseScRsp                           = 7047
	QuitBattleScNotify                                 = 200
	MonopolyGiveUpCurContentScRsp                      = 7065
	TakeOffRelicCsReq                                  = 376
	StartRaidScRsp                                     = 2233
	RogueModifierStageStartNotify                      = 5370
	MonopolyGetRegionProgressCsReq                     = 7020
	SetHeroBasicTypeCsReq                              = 38
	ComposeSelectedRelicCsReq                          = 508
	GetTutorialGuideScRsp                              = 1645
	SceneCastSkillCostMpScRsp                          = 1436
	GetFriendAssistListScRsp                           = 2948
	GetTreasureDungeonActivityDataScRsp                = 4411
	InteractChargerCsReq                               = 6842
	GetPlayerBoardDataScRsp                            = 2833
	EntityBindPropScRsp                                = 1418
	GetExpeditionDataCsReq                             = 2510
	MonopolyConfirmRandomScRsp                         = 7066
	GetMonopolyDailyReportCsReq                        = 7050
	AlleyTakeEventRewardScRsp                          = 4707
	TakeMailAttachmentScRsp                            = 826
	SetDisplayAvatarCsReq                              = 2816
	ChessRogueQueryAeonDimensionsScRsp                 = 5404
	HeliobusSnsLikeCsReq                               = 5877
	TravelBrochureSetCustomValueScRsp                  = 6402
	ScenePlaneEventScNotify                            = 1471
	MatchBoxingClubOpponentScRsp                       = 4245
	CurAssistChangedNotify                             = 2997
	ArchiveWolfBroGameCsReq                            = 6542
	PlayerReturnTakePointRewardCsReq                   = 4516
	ChessRogueSelectBpCsReq                            = 5443
	DeployRotaterCsReq                                 = 6816
	AceAntiCheaterScRsp                                = 84
	GetAllLineupDataCsReq                              = 766
	TakeFightActivityRewardCsReq                       = 3660
	GetMonopolyMbtiReportRewardScRsp                   = 7014
	MultipleDropInfoNotify                             = 4660
	GetPunkLordDataCsReq                               = 3202
	UnlockTutorialGuideScRsp                           = 1626
	WaypointShowNewCsNotify                            = 477
	AlleyShopLevelScNotify                             = 4708
	EnterSceneByServerScNotify                         = 1459
	GeneralVirtualItemDataNotify                       = 547
	DeleteFriendScRsp                                  = 2995
	GetTutorialScRsp                                   = 1633
	PlayerGetTokenCsReq                                = 16
	TravelBrochureGetDataScRsp                         = 6433
	GiveUpBoxingClubChallengeCsReq                     = 4277
	SwitchAetherDivideLineUpSlotScRsp                  = 4866
	HeliobusEnterBattleCsReq                           = 5866
	FinishCurTurnCsReq                                 = 4303
	TakeRogueScoreRewardScRsp                          = 1894
	GetFriendRecommendListInfoScRsp                    = 2966
	PlayerReturnSignCsReq                              = 4533
	EnterFeverTimeActivityStageScRsp                   = 7158
	GetFriendLoginInfoCsReq                            = 2939
	GetFriendListInfoScRsp                             = 2933
	BoxingClubRewardScNotify                           = 4300
	RemoveRotaterCsReq                                 = 6876
	GetQuestDataCsReq                                  = 910
	PlayerReturnInfoQueryScRsp                         = 4570
	PlayerReturnStartScNotify                          = 4510
	ChallengeLineupNotify                              = 1711
	ExchangeRogueBuffWithMiracleScRsp                  = 5666
	UpdateMechanismBarScNotify                         = 1446
	UpdatePlayerSettingCsReq                           = 98
	UpdateTrackMainMissionIdScRsp                      = 1286
	GetAssistHistoryScRsp                              = 2907
	PlayBackGroundMusicCsReq                           = 3142
	SyncRogueExploreWinScNotify                        = 1801
	MonopolyGameBingoFlipCardCsReq                     = 7001
	GetMonopolyFriendRankingListCsReq                  = 7043
	StartBattleCollegeScRsp                            = 5716
	ChangeScriptEmotionScRsp                           = 6345
	LeaveAetherDivideSceneCsReq                        = 4842
	SyncRogueCommonPendingActionScNotify               = 5623
	MarkReadMailScRsp                                  = 845
	GetRogueDialogueEventDataScRsp                     = 1843
	MarkChatEmojiScRsp                                 = 3911
	GetServerPrefsDataCsReq                            = 6142
	GetSecretKeyInfoCsReq                              = 54
	ChessRogueSelectCellScRsp                          = 5475
	GetTrainVisitorBehaviorScRsp                       = 3745
	SelectInclinationTextScRsp                         = 2170
	SetIsDisplayAvatarInfoScRsp                        = 2826
	GetRogueInitialScoreCsReq                          = 1847
	LeaveChallengeCsReq                                = 1716
	BuyBpLevelScRsp                                    = 3077
	AcceptMissionEventScRsp                            = 1291
	GetMonopolyInfoCsReq                               = 7010
	GetQuestRecordCsReq                                = 977
	StartTimedCocoonStageCsReq                         = 1490
	EnterRogueScRsp                                    = 1860
	LastSpringRefreshTimeNotify                        = 1466
	LogisticsDetonateStarSkiffScRsp                    = 4786
	DrinkMakerDayEndScNotify                           = 6996
	TakeOffEquipmentCsReq                              = 303
	GetRogueBuffEnhanceInfoScRsp                       = 1808
	EnhanceChessRogueBuffCsReq                         = 5545
	CancelCacheNotifyScRsp                             = 4170
	SendMsgCsReq                                       = 3910
	SwitchLineupIndexScRsp                             = 799
	MonopolyMoveScRsp                                  = 7100
	StoryLineInfoScNotify                              = 6242
	ChessRogueUpdateBoardScNotify                      = 5523
	GetQuestDataScRsp                                  = 933
	BuyNpcStuffCsReq                                   = 4342
	TakeRogueScoreRewardCsReq                          = 1837
	CancelActivityExpeditionCsReq                      = 2511
	HeliobusEnterBattleScRsp                           = 5837
	GetLevelRewardTakenListScRsp                       = 96
	SetHeadIconScRsp                                   = 2845
	AcceptExpeditionCsReq                              = 2542
	HandleRogueCommonPendingActionCsReq                = 5668
	LogisticsInfoScNotify                              = 4785
	DailyFirstMeetPamScRsp                             = 3445
	ChessRogueQuestFinishNotify                        = 5598
	RechargeSuccNotify                                 = 537
	MonopolyLikeScNotify                               = 7022
	SyncRogueMapRoomScNotify                           = 1839
	SetCurWaypointCsReq                                = 442
	ChessRogueUpdateAllowedSelectCellScNotify          = 5494
	SceneEntityTeleportCsReq                           = 1456
	TravelBrochureUpdatePasterPosScRsp                 = 6411
	ChessRogueCheatRollScRsp                           = 5529
	UnlockTutorialScRsp                                = 1660
	UnlockTeleportNotify                               = 1482
	SetFriendRemarkNameScRsp                           = 2994
	GetFriendListInfoCsReq                             = 2910
	MonopolyTakeRaffleTicketRewardScRsp                = 7005
	HeartDialScriptChangeScNotify                      = 6400
	ExchangeHcoinCsReq                                 = 594
	HeliobusLineupUpdateScNotify                       = 5806
	ChessRogueQueryBpCsReq                             = 5482
	TakeChapterRewardCsReq                             = 426
	SelectPhoneThemeScRsp                              = 5177
	SyncRoguePickAvatarInfoScNotify                    = 1853
	TakeRogueMiracleHandbookRewardScRsp                = 5647
	DoGachaInRollShopScRsp                             = 6907
	GetSaveLogisticsMapCsReq                           = 4767
	EnterFeverTimeActivityStageCsReq                   = 7155
	StartChallengeScRsp                                = 1745
	PVEBattleResultScRsp                               = 133
	UpdateRedDotDataCsReq                              = 5942
	FinishPlotCsReq                                    = 1110
	BuyGoodsScRsp                                      = 1545
	BuyRogueShopBuffCsReq                              = 5670
	SubmitOrigamiItemCsReq                             = 4136
	UpdateMovieRacingDataScRsp                         = 4108
	MonopolyGameBingoFlipCardScRsp                     = 7007
	ExpeditionDataChangeScNotify                       = 2600
	GetTelevisionActivityDataScRsp                     = 6977
	FinishTutorialCsReq                                = 1700
	WolfBroGameDataChangeScNotify                      = 6503
	GroupStateChangeScRsp                              = 1440
	SyncRogueReviveInfoScNotify                        = 1838
	HeliobusUpgradeLevelCsReq                          = 5858
	HeliobusSelectSkillCsReq                           = 5802
	PlayerKickOutScNotify                              = 100
	FinishFirstTalkByPerformanceNpcScRsp               = 2195
	SyncClientResVersionCsReq                          = 177
	SceneUpdatePositionVersionNotify                   = 1411
	EnterTelevisionActivityStageScRsp                  = 6971
	TrainVisitorRewardSendNotify                       = 3760
	EndDrinkMakerSequenceScRsp                         = 7000
	GetMarkItemListCsReq                               = 532
	AetherDivideRefreshEndlessScRsp                    = 4872
	ServerSimulateBattleFinishScNotify                 = 111
	EnterTrialActivityStageScRsp                       = 2689
	ChessRogueSkipTeachingLevelScRsp                   = 5540
	EntityBindPropCsReq                                = 1486
	GetMultipleDropInfoCsReq                           = 4610
	GetDailyActiveInfoCsReq                            = 3342
	GetPhoneDataCsReq                                  = 5110
	SpaceZooExchangeItemCsReq                          = 6711
	ChessRogueReRollDiceCsReq                          = 5435
	PlayerHeartBeatScRsp                               = 73
	SecurityReportScRsp                                = 4111
	RaidCollectionDataScRsp                            = 6957
	GetDrinkMakerDataScRsp                             = 6997
	ExtraLineupDestroyNotify                           = 706
	MonopolyGetRegionProgressScRsp                     = 7081
	GetPrivateChatHistoryCsReq                         = 3916
	PlayerReturnTakeRewardCsReq                        = 4577
	GetSingleRedDotParamGroupScRsp                     = 5960
	RankUpAvatarCsReq                                  = 395
	GameplayCounterCountDownScRsp                      = 1452
	GetRogueHandbookDataCsReq                          = 5687
	SceneCastSkillScRsp                                = 1460
	ChessRogueGoAheadScRsp                             = 5539
	MonopolyReRollRandomScRsp                          = 7076
	ChessRogueChangeyAeonDimensionNotify               = 5492
	ChessRogueNousGetRogueTalentInfoCsReq              = 5537
	GetTutorialCsReq                                   = 1610
	GetArchiveDataScRsp                                = 2333
	SyncAddBlacklistScNotify                           = 2976
	AlleyTakeEventRewardCsReq                          = 4701
	ExchangeRogueRewardKeyCsReq                        = 1846
	ChessRogueGiveUpCsReq                              = 5485
	ChessRogueQuitCsReq                                = 5569
	QueryProductInfoScRsp                              = 44
	StartTimedFarmElementScRsp                         = 1451
	ReplaceLineupCsReq                                 = 796
	GetTrainVisitorRegisterCsReq                       = 3777
	SetBoxingClubResonanceLineupCsReq                  = 4258
	GetAllRedDotDataCsReq                              = 5910
	UseItemCsReq                                       = 577
	ShareScRsp                                         = 4133
	StartTrialActivityCsReq                            = 2673
	GetStrongChallengeActivityDataCsReq                = 6610
	TakeAllApRewardScRsp                               = 3377
	GetTrialActivityDataScRsp                          = 2643
	GetChessRogueBuffEnhanceInfoCsReq                  = 5411
	MonopolyGetRaffleTicketScRsp                       = 7059
	GetFriendApplyListInfoScRsp                        = 2960
	GetPlayerBoardDataCsReq                            = 2810
	UnlockSkilltreeCsReq                               = 316
	TakeAllRewardCsReq                                 = 3026
	ChessRogueQueryAeonDimensionsCsReq                 = 5424
	PunkLordMonsterInfoScNotify                        = 3236
	NewMailScNotify                                    = 900
	ChessRogueStartScRsp                               = 5567
	GetRogueInfoScRsp                                  = 1833
	RogueModifierSelectCellCsReq                       = 5345
	LogisticsGameScRsp                                 = 4745
	TakePromotionRewardCsReq                           = 366
	ShowNewSupplementVisitorCsReq                      = 3703
	GetAssistListCsReq                                 = 2983
	EnterAdventureCsReq                                = 1310
	TravelBrochureSelectMessageScRsp                   = 6460
	GetAetherDivideInfoScRsp                           = 4811
	GetCurAssistScRsp                                  = 2972
	GetNpcStatusScRsp                                  = 2745
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5694
	RogueNpcDisappearScRsp                             = 5658
	RankUpEquipmentScRsp                               = 570
	GetMonsterResearchActivityDataScRsp                = 2676
	SellItemScRsp                                      = 566
	FinishChessRogueNousSubStoryScRsp                  = 5438
	GetPlatformPlayerInfoScRsp                         = 2949
	SpaceZooTakeCsReq                                  = 6795
	AddAvatarScNotify                                  = 358
	SelectChessRogueSubStoryScRsp                      = 5446
	FinishQuestScRsp                                   = 903
	SyncChessRogueNousValueScNotify                    = 5486
	RelicRecommendCsReq                                = 544
	ChangeLineupLeaderScRsp                            = 736
	TakePictureCsReq                                   = 4116
	GetBagScRsp                                        = 533
	GetHeroBasicTypeInfoCsReq                          = 32
	GetChatFriendHistoryCsReq                          = 3977
	GetPunkLordMonsterDataScRsp                        = 3233
	WolfBroGameActivateBulletCsReq                     = 6502
	SavePointsInfoNotify                               = 1401
	DeleteBlacklistCsReq                               = 2906
	PrivateMsgOfflineUsersScNotify                     = 3945
	TakeActivityExpeditionRewardCsReq                  = 2595
	SubmitEmotionItemCsReq                             = 6316
	StartWolfBroGameScRsp                              = 6533
	UnlockBackGroundMusicCsReq                         = 3116
	GetPlayerReplayInfoScRsp                           = 3545
	ClientObjUploadCsReq                               = 52
	GetChallengeRaidInfoScRsp                          = 2277
	GetSocialEventServerCacheCsReq                     = 7028
	GetRogueScoreRewardInfoScRsp                       = 1852
	GetLoginChatInfoCsReq                              = 3936
	QuitLineupCsReq                                    = 777
	GetFirstTalkNpcCsReq                               = 2116
	GetAvatarDataCsReq                                 = 310
	GetBoxingClubInfoScRsp                             = 4233
	StartAetherDivideSceneBattleCsReq                  = 4816
	TakeChallengeRewardScRsp                           = 1702
	DressRelicAvatarCsReq                              = 302
	MonopolyEventLoadUpdateScNotify                    = 7004
	FightTreasureDungeonMonsterCsReq                   = 4499
	GetFantasticStoryActivityDataScRsp                 = 4933
	ChooseBoxingClubResonanceScRsp                     = 4211
	GetAetherDivideChallengeInfoCsReq                  = 4865
	GetRogueEndlessActivityDataScRsp                   = 6033
	GetMissionEventDataCsReq                           = 1236
	StopRogueAdventureRoomScRsp                        = 5665
	TravelBrochureRemovePasterScRsp                    = 6470
	GetBattleCollegeDataScRsp                          = 5733
	SyncRogueAdventureRoomInfoScNotify                 = 5610
	TakePrestigeRewardScRsp                            = 4711
	LockEquipmentScRsp                                 = 560
	SetLineupNameScRsp                                 = 791
	HeliobusStartRaidScRsp                             = 5896
	HeliobusChallengeUpdateScNotify                    = 5808
	ResetMapRotationRegionScRsp                        = 6895
	HeliobusActivityDataScRsp                          = 5833
	GetShareDataScRsp                                  = 4145
	SpringRefreshScRsp                                 = 1491
	GetCurLineupDataScRsp                              = 745
	HeliobusSnsCommentCsReq                            = 5900
	PrestigeLevelUpCsReq                               = 4737
	TelevisionActivityBattleEndScNotify                = 6980
	SetStuffToAreaCsReq                                = 4316
	SubmitOfferingItemScRsp                            = 6927
	GetCurChallengeCsReq                               = 1770
	SetMissionEventProgressScRsp                       = 1206
	SubMissionRewardScNotify                           = 1265
	TravelBrochureApplyPasterScRsp                     = 6426
	MakeMissionDrinkCsReq                              = 6989
	StartFinishSubMissionScNotify                      = 1283
	TakeOfferingRewardCsReq                            = 6931
	TakeTalkRewardCsReq                                = 2142
	ChessRogueUpdateMoneyInfoScNotify                  = 5592
	TakeBpRewardScRsp                                  = 3016
	GetRaidInfoScRsp                                   = 2211
	MonopolyAcceptQuizScRsp                            = 7086
	TravelBrochurePageUnlockScNotify                   = 6442
	CityShopInfoScNotify                               = 1577
	SpaceZooDataScRsp                                  = 6733
	PromoteAvatarScRsp                                 = 326
	ChessRogueGiveUpRollScRsp                          = 5504
	MonopolyLikeScRsp                                  = 7089
	GetStoryLineInfoScRsp                              = 6233
	HeartDialTraceScriptCsReq                          = 6370
	SetLanguageCsReq                                   = 85
	DeactivateFarmElementCsReq                         = 1439
	ReportPlayerCsReq                                  = 2996
	SyncAcceptedPamMissionNotify                       = 4042
	DressAvatarCsReq                                   = 400
	TakeExpeditionRewardCsReq                          = 2577
	TrainVisitorBehaviorFinishCsReq                    = 3710
	TakeMonsterResearchActivityRewardCsReq             = 2637
	SpaceZooDeleteCatScRsp                             = 6770
	OpenTreasureDungeonGridScRsp                       = 4402
	GetRogueAeonInfoScRsp                              = 1855
	QuitTreasureDungeonScRsp                           = 4408
	EnterFightActivityStageScRsp                       = 3616
	MonopolyGetDailyInitItemScRsp                      = 7030
	MonopolyClickCellScRsp                             = 7055
	RogueModifierDelNotify                             = 5400
	RaidInfoNotify                                     = 2216
	UpgradeAreaScRsp                                   = 4395
	GetRollShopInfoCsReq                               = 6919
	ChessRogueQueryScRsp                               = 5593
	BatchGetQuestDataCsReq                             = 936
	MuseumRandomEventQueryScRsp                        = 4337
	SearchPlayerScRsp                                  = 2985
	TeleportToMissionResetPointScRsp                   = 1285
	GmTalkScRsp                                        = 3
	GetCurSceneInfoCsReq                               = 1477
	CancelExpeditionCsReq                              = 2516
	SyncApplyFriendScNotify                            = 3000
	EnteredSceneChangeScNotify                         = 1430
	GetAllServerPrefsDataCsReq                         = 6110
	TravelBrochurePageResetCsReq                       = 6491
	SyncRogueVirtualItemInfoScNotify                   = 1809
	MonopolySttUpdateScNotify                          = 7075
	EnhanceRogueBuffScRsp                              = 1865
	SpaceZooCatUpdateNotify                            = 6703
	GetRogueTalentInfoScRsp                            = 1821
	QuitWolfBroGameScRsp                               = 6526
	GetWolfBroGameDataScRsp                            = 6570
	HeliobusSnsPostCsReq                               = 5816
	RogueModifierAddNotify                             = 5342
	AcceptActivityExpeditionScRsp                      = 2503
	PunkLordDataChangeNotify                           = 3238
	SetClientRaidTargetCountScRsp                      = 2295
	ChessRogueUpdateDiceInfoScNotify                   = 5530
	MonopolyRollRandomCsReq                            = 7036
	ChangeStoryLineFinishScNotify                      = 6260
	StartPunkLordRaidCsReq                             = 3242
	GetChallengeGroupStatisticsScRsp                   = 1776
	ChessRogueEnterCellCsReq                           = 5551
	GetSaveRaidCsReq                                   = 2236
	ChallengeSettleNotify                              = 1777
	RecoverAllLineupScRsp                              = 1472
	DoGachaScRsp                                       = 1945
	RefreshTriggerByClientScNotify                     = 1415
	ChessRogueLeaveScRsp                               = 5432
	ReviveRogueAvatarScRsp                             = 1866
	RemoveStuffFromAreaCsReq                           = 4377
	GetFirstTalkByPerformanceNpcCsReq                  = 2103
	TravelBrochureApplyPasterListCsReq                 = 6437
	SetAssistCsReq                                     = 2938
	GetOfferingInfoCsReq                               = 6939
	LeaveMapRotationRegionScRsp                        = 6870
	SyncRogueHandbookDataUpdateScNotify                = 5618
	SpaceZooOpCatteryScRsp                             = 6726
	SetPlayerInfoCsReq                                 = 97
	StoryLineTrialAvatarChangeScNotify                 = 6277
	WolfBroGameUseBulletScRsp                          = 6558
	GetAssistHistoryCsReq                              = 2901
	ShowNewSupplementVisitorScRsp                      = 3711
	ChessRogueConfirmRollCsReq                         = 5519
	PickRogueAvatarCsReq                               = 1802
	GetRogueInfoCsReq                                  = 1810
	RestoreWolfBroGameArchiveScRsp                     = 6560
	ChooseBoxingClubStageOptionalBuffCsReq             = 4236
	TakeRogueAeonLevelRewardScRsp                      = 1878
	SetGroupCustomSaveDataCsReq                        = 1473
	GetUpdatedArchiveDataScRsp                         = 2345
	AetherDivideSpiritExpUpScRsp                       = 4808
	MonopolyTakePhaseRewardCsReq                       = 7063
	MissionEventRewardScNotify                         = 1299
	HeliobusInfoChangedScNotify                        = 5811
	UpdateMapRotationDataScNotify                      = 6899
	MonopolyReRollRandomCsReq                          = 7099
	FantasticStoryActivityBattleEndScNotify            = 4960
	GetStuffScNotify                                   = 4400
	TakePromotionRewardScRsp                           = 337
	GetFirstTalkNpcScRsp                               = 2160
	GetRogueTalentInfoCsReq                            = 1861
	SceneEntityUpdateScNotify                          = 1500
	MonopolyScrachRaffleTicketScRsp                    = 7078
	MonopolyGuessBuyInformationScRsp                   = 7039
	GameplayCounterRecoverCsReq                        = 1427
	StartWolfBroGameCsReq                              = 6510
	PromoteEquipmentCsReq                              = 542
	ChangeScriptEmotionCsReq                           = 6342
	GetRogueAdventureRoomInfoCsReq                     = 5695
	EnterSectionScRsp                                  = 1485
	TakeAssistRewardCsReq                              = 2986
	GetLineupAvatarDataCsReq                           = 711
	GetChessRogueStoryInfoCsReq                        = 5463
	MonopolyGetDailyInitItemCsReq                      = 7064
	HeliobusUnlockSkillScNotify                        = 5836
	TakeTrialActivityRewardCsReq                       = 2622
	GetHeroBasicTypeInfoScRsp                          = 72
	TakeRogueEndlessActivityPointRewardCsReq           = 6060
	QuitRogueScRsp                                     = 1832
	SwitchAetherDivideLineUpSlotCsReq                  = 4891
	GateServerScNotify                                 = 80
	UnlockTutorialGuideCsReq                           = 1677
	MonopolyContentUpdateScNotify                      = 7083
	MuseumTargetRewardNotify                           = 4369
	PlayerLoginFinishCsReq                             = 79
	StartRaidCsReq                                     = 2210
	GetPunkLordMonsterDataCsReq                        = 3210
	EnterSceneScRsp                                    = 1474
	ChessRogueEnterNextLayerCsReq                      = 5490
	BatchMarkChatEmojiCsReq                            = 3958
	SceneGroupRefreshScNotify                          = 1475
	GetBasicInfoCsReq                                  = 29
	TakeCityShopRewardCsReq                            = 1516
	EnterFantasticStoryActivityStageScRsp              = 4916
	FinishCurTurnScRsp                                 = 4311
	SceneEnterStageScRsp                               = 1408
	RogueEndlessActivityBattleEndScNotify              = 6016
	GetFeverTimeActivityDataCsReq                      = 7153
	MonopolyGiveUpCurContentCsReq                      = 7006
	SceneCastSkillCostMpCsReq                          = 1495
	TakeOffAvatarSkinScRsp                             = 306
	AlleyShipmentEventEffectsScNotify                  = 4783
	GetChallengeScRsp                                  = 1733
	MuseumDispatchFinishedScNotify                     = 4308
	GetGachaInfoScRsp                                  = 1933
	GetExpeditionDataScRsp                             = 2533
	TextJoinSaveScRsp                                  = 3833
	SelectChessRogueNousSubStoryScRsp                  = 5511
	InteractChargerScRsp                               = 6845
	ChessRogueEnterCsReq                               = 5469
	ExchangeRogueBuffWithMiracleCsReq                  = 5691
	HeliobusSnsCommentScRsp                            = 5870
	UpdateMovieRacingDataCsReq                         = 4196
	MonthCardRewardNotify                              = 89
	SelectPhoneThemeCsReq                              = 5160
	FinishTutorialGuideScRsp                           = 1611
	MonopolyTakeRaffleTicketRewardCsReq                = 7071
	SetCurInteractEntityScRsp                          = 1412
	PunkLordMonsterKilledNotify                        = 3285
	TakeOffRelicScRsp                                  = 391
	MakeDrinkScRsp                                     = 6987
	GetMultipleDropInfoScRsp                           = 4633
	TakeApRewardCsReq                                  = 3310
	GetUnlockTeleportScRsp                             = 1481
	MonopolyCellUpdateNotify                           = 7045
	GetMissionDataScRsp                                = 1233
	StartAetherDivideSceneBattleScRsp                  = 4860
	ChessRogueReviveAvatarCsReq                        = 5412
	SceneEntityDisappearScNotify                       = 1470
	QuitWolfBroGameCsReq                               = 6577
	AcceptMultipleExpeditionCsReq                      = 2502
	GetMainMissionCustomValueScRsp                     = 1272
	TextJoinQueryScRsp                                 = 3845
	DressRelicAvatarScRsp                              = 399
	LeaveRaidCsReq                                     = 2242
	HeliobusStartRaidCsReq                             = 5894
	SpringRecoverSingleAvatarScRsp                     = 1422
	DailyFirstMeetPamCsReq                             = 3442
	TakeMonsterResearchActivityRewardScRsp             = 2694
	SetFriendMarkScRsp                                 = 2924
	BatchGetQuestDataScRsp                             = 902
	FeverTimeActivityBattleEndScNotify                 = 7159
	MonopolyConditionUpdateScNotify                    = 7061
	FinishRogueDialogueGroupCsReq                      = 1868
	SyncEntityBuffChangeListScNotify                   = 1458
	SetFriendRemarkNameCsReq                           = 2937
	FinishFirstTalkNpcCsReq                            = 2177
	DeleteSocialEventServerCacheCsReq                  = 7056
	HeliobusSnsLikeScRsp                               = 5826
	EnterRogueMapRoomCsReq                             = 1818
	MatchBoxingClubOpponentCsReq                       = 4242
	EnterTreasureDungeonCsReq                          = 4458
	GetFantasticStoryActivityDataCsReq                 = 4910
	SpaceZooDataCsReq                                  = 6710
	TriggerVoiceScRsp                                  = 4195
	PrepareRogueAdventureRoomCsReq                     = 5633
	TravelBrochureGetPasterScNotify                    = 6458
	AceAntiCheaterCsReq                                = 68
	GetRogueShopBuffInfoCsReq                          = 5660
	SetForbidOtherApplyFriendScRsp                     = 2931
	ExpUpEquipmentScRsp                                = 511
	SetDisplayAvatarScRsp                              = 2860
	RotateMapCsReq                                     = 6877
	MonopolyEventSelectFriendCsReq                     = 7080
	HeartDialTraceScriptScRsp                          = 6303
	ChessRogueRollDiceCsReq                            = 5415
	GetTutorialGuideCsReq                              = 1642
	MultipleDropInfoScNotify                           = 4642
	TakeMultipleExpeditionRewardCsReq                  = 2576
	ChessRogueMoveCellNotify                           = 5422
	DelMailCsReq                                       = 816
	ChessRogueUpdateAeonModifierValueScNotify          = 5497
	GetServerPrefsDataScRsp                            = 6145
	MonopolyBuyGoodsCsReq                              = 7037
	GetChatEmojiListScRsp                              = 3970
	SetAssistAvatarScRsp                               = 2858
	GetRogueDialogueEventDataCsReq                     = 1848
	TakePunkLordPointRewardCsReq                       = 3258
	FinishEmotionDialoguePerformanceCsReq              = 6377
	QuitBattleScRsp                                    = 145
	EnterMapRotationRegionCsReq                        = 6810
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3800
	SelectChatBubbleCsReq                              = 5142
	MissionRewardScNotify                              = 1216
	SyncRogueSeasonFinishScNotify                      = 1807
	MuseumTakeCollectRewardScRsp                       = 4383
	UnlockedAreaMapScNotify                            = 1413
	CancelCacheNotifyCsReq                             = 4200
	GetMuseumInfoScRsp                                 = 4333
	PlayerLoginCsReq                                   = 10
	ChessRogueUpdateLevelBaseInfoScNotify              = 5451
	GetFriendLoginInfoScRsp                            = 2944
	GetChapterCsReq                                    = 416
	GetSaveLogisticsMapScRsp                           = 4738
	GetTrialActivityDataCsReq                          = 2648
	GetTrainVisitorBehaviorCsReq                       = 3742
	GetSpringRecoverDataCsReq                          = 1429
	SharePunkLordMonsterScRsp                          = 3260
	GetActivityScheduleConfigScRsp                     = 2660
	GetFriendRecommendListInfoCsReq                    = 2991
	GetHeartDialInfoCsReq                              = 6310
	GetRogueHandbookDataScRsp                          = 5686
	HandleRogueCommonPendingActionScRsp                = 5684
	ChessRogueNousGetRogueTalentInfoScRsp              = 5558
	DiscardRelicScRsp                                  = 539
	GetSceneMapInfoScRsp                               = 1434
	UpdatePlayerSettingScRsp                           = 88
	GetMailCsReq                                       = 810
	FeatureSwitchClosedScNotify                        = 62
	TakeQuestRewardScRsp                               = 945
	TakeKilledPunkLordMonsterScoreCsReq                = 3283
	AetherDivideSpiritInfoScNotify                     = 4806
	SaveLogisticsCsReq                                 = 4765
	GetRndOptionScRsp                                  = 3433
	SpaceZooBornScRsp                                  = 6745
	HandleFriendCsReq                                  = 2970
	GetPlayerDetailInfoScRsp                           = 2945
	RevcMsgScNotify                                    = 3942
	InterruptMissionEventCsReq                         = 1294
	TakeQuestOptionalRewardScRsp                       = 958
	LeaveTrialActivityCsReq                            = 2662
	WolfBroGamePickupBulletCsReq                       = 6595
	StaminaInfoScNotify                                = 81
	GetMissionStatusScRsp                              = 1237
	SetAetherDivideLineUpScRsp                         = 4895
	ClearAetherDividePassiveSkillCsReq                 = 4899
	FinishTalkMissionScRsp                             = 1245
	GetPunkLordBattleRecordCsReq                       = 3212
	SetClientPausedScRsp                               = 1447
	AetherDivideTakeChallengeRewardCsReq               = 4807
	ChessRogueQuitScRsp                                = 5473
	FinishSectionIdCsReq                               = 2777
	TravelBrochureGetDataCsReq                         = 6410
	FinishPerformSectionIdCsReq                        = 2800
	VirtualLineupDestroyNotify                         = 794
	TakeRollShopRewardScRsp                            = 6920
	MonopolyRollDiceCsReq                              = 7060
	GetSingleRedDotParamGroupCsReq                     = 5916
	MarkItemScRsp                                      = 507
	PlayerReturnSignScRsp                              = 4542
	ActivateFarmElementScRsp                           = 1431
	ReturnLastTownCsReq                                = 1437
	StartAlleyEventScRsp                               = 4726
	FinishTutorialScRsp                                = 1670
	PlayerLogoutScRsp                                  = 45
	GetEnteredSceneCsReq                               = 1455
	MonopolyUpgradeAssetCsReq                          = 7096
	SelectRogueDialogueEventCsReq                      = 1879
	StartChallengeCsReq                                = 1742
	GetMarkItemListScRsp                               = 572
	SetSpringRecoverConfigCsReq                        = 1493
	GetAllLineupDataScRsp                              = 737
	CancelMarkItemNotify                               = 587
	SetGameplayBirthdayCsReq                           = 48
	SyncRogueAreaUnlockScNotify                        = 1871
	DestroyItemCsReq                                   = 538
	GetChessRogueNousStoryInfoScRsp                    = 5505
	GetPunkLordDataScRsp                               = 3299
	GetDrinkMakerDataCsReq                             = 6999
	TakeFightActivityRewardScRsp                       = 3677
	SummonPunkLordMonsterCsReq                         = 3277
	StartAetherDivideStageBattleScRsp                  = 4894
	GetMonopolyMbtiReportRewardCsReq                   = 7027
	GetRecyleTimeScRsp                                 = 585
	GetNpcTakenRewardCsReq                             = 2110
	ExchangeGachaCeilingScRsp                          = 1926
	GetGachaCeilingScRsp                               = 1960
	FinishChessRogueSubStoryCsReq                      = 5591
	TrialBackGroundMusicCsReq                          = 3177
	GetRecyleTimeCsReq                                 = 569
	DailyFirstEnterMonopolyActivityScRsp               = 7095
	AddEquipmentScNotify                               = 565
	ChessRogueSelectBpScRsp                            = 5474
	GetNpcMessageGroupScRsp                            = 2733
	MonopolySelectOptionCsReq                          = 7070
	GetKilledPunkLordMonsterDataCsReq                  = 3208
	OpenRogueChestScRsp                                = 1822
	GetVideoVersionKeyCsReq                            = 74
	GetPlayerReturnMultiDropInfoCsReq                  = 4645
	GetAlleyInfoScRsp                                  = 4733
	AddBlacklistScRsp                                  = 2999
	MonopolyCheatDiceScRsp                             = 7085
	TextJoinBatchSaveScRsp                             = 3860
	ClearAetherDividePassiveSkillScRsp                 = 4876
	EndDrinkMakerSequenceCsReq                         = 6991
	TakeLoginActivityRewardCsReq                       = 2642
	GetStrongChallengeActivityDataScRsp                = 6633
	DestroyItemScRsp                                   = 512
	GetJukeboxDataCsReq                                = 3110
	SyncRogueAeonLevelUpRewardScNotify                 = 1888
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6100
	SyncTaskCsReq                                      = 1260
	SyncChessRogueNousMainStoryScNotify                = 5542
	UpdateRogueAdventureRoomScoreCsReq                 = 5631
	FinishItemIdScRsp                                  = 2760
	MuseumRandomEventSelectScRsp                       = 4396
	ExpUpRelicCsReq                                    = 536
	GetSceneMapInfoCsReq                               = 1405
	SpaceZooMutateScRsp                                = 6760
	ComposeItemScRsp                                   = 595
	EnterChessRogueAeonRoomScRsp                       = 5426
	PlayerLoginScRsp                                   = 33
	MonopolyGameGachaScRsp                             = 7072
	GetMapRotationDataScRsp                            = 6811
	GameplayCounterCountDownCsReq                      = 1463
	ChessRogueLayerAccountInfoNotify                   = 5597
	DrinkMakerUpdateTipsNotify                         = 6982
	GetStageLineupCsReq                                = 710
	AlleyFundsScNotify                                 = 4796
	EnterStrongChallengeActivityStageCsReq             = 6642
	GetChapterScRsp                                    = 460
	MuseumFundsChangedScNotify                         = 4376
	EnhanceCommonRogueBuffScRsp                        = 5608
	ReEnterLastElementStageCsReq                       = 1428
	TakeRogueEventHandbookRewardScRsp                  = 5639
	ChessRogueUpdateActionPointScNotify                = 5414
	UpgradeAreaStatScRsp                               = 4302
	UnlockTutorialCsReq                                = 1616
	GetLevelRewardScRsp                                = 6
	AvatarExpUpScRsp                                   = 345
	MuseumTargetStartNotify                            = 4306
	AlleyEventChangeNotify                             = 4800
	TakeQuestOptionalRewardCsReq                       = 911
	GetWaypointScRsp                                   = 433
	GetMailScRsp                                       = 833
	GetReplayTokenScRsp                                = 3533
	StartBoxingClubBattleCsReq                         = 4216
	UpdateServerPrefsDataScRsp                         = 6160
	TakeChallengeRaidRewardCsReq                       = 2226
	SceneEntityMoveCsReq                               = 1410
	GetChatFriendHistoryScRsp                          = 3926
	TakePunkLordPointRewardScRsp                       = 3295
	MonopolyBuyGoodsScRsp                              = 7094
	GetMuseumInfoCsReq                                 = 4310
	SetTurnFoodSwitchCsReq                             = 518
	BattleCollegeDataChangeScNotify                    = 5742
	SetAssistAvatarCsReq                               = 2811
	GetMonsterResearchActivityDataCsReq                = 2699
	SyncRogueDialogueEventDataScNotify                 = 1874
	SetGenderCsReq                                     = 86
	EnterAdventureScRsp                                = 1333
	LogisticsGameCsReq                                 = 4742
	GetChatEmojiListCsReq                              = 4000
	DressAvatarSkinScRsp                               = 396
	GetAssistListScRsp                                 = 2967
	AetherDivideTakeChallengeRewardScRsp               = 4887
	SetClientRaidTargetCountCsReq                      = 2258
	EnterMapRotationRegionScRsp                        = 6833
	GetMissionStatusCsReq                              = 1266
	RogueNpcDisappearCsReq                             = 5611
	GetMovieRacingDataCsReq                            = 4137
	GetNpcStatusCsReq                                  = 2742
	GetTrainVisitorRegisterScRsp                       = 3726
	ChessRoguePickAvatarCsReq                          = 5461
	ChessRogueNousEnableRogueTalentCsReq               = 5518
	GetBattleCollegeDataCsReq                          = 5710
	InteractTreasureDungeonGridCsReq                   = 4491
	SetSignatureScRsp                                  = 2803
	SyncServerSceneChangeNotify                        = 1478
	EquipAetherDividePassiveSkillScRsp                 = 4802
	SelectChessRogueSubStoryCsReq                      = 5436
	GetChallengeRaidInfoCsReq                          = 2260
	UnlockChatBubbleScNotify                           = 5116
	HealPoolInfoNotify                                 = 1484
	SetForbidOtherApplyFriendCsReq                     = 2923
	ReEnterLastElementStageScRsp                       = 1454
	InterruptMissionEventScRsp                         = 1296
	GetPlatformPlayerInfoCsReq                         = 2947
	GetAetherDivideInfoCsReq                           = 4803
	RefreshTriggerByClientScRsp                        = 1421
	RemoveRotaterScRsp                                 = 6891
	ClientObjUploadScRsp                               = 4
	AvatarExpUpCsReq                                   = 342
	GetArchiveDataCsReq                                = 2310
	GetPlayerDetailInfoCsReq                           = 2942
	GetShopListCsReq                                   = 1510
	CancelExpeditionScRsp                              = 2560
	LockRelicScRsp                                     = 576
	FinishRogueDialogueGroupScRsp                      = 1884
	GetPrivateChatHistoryScRsp                         = 3960
	MonopolyEventSelectFriendScRsp                     = 7062
	StartPunkLordRaidScRsp                             = 3245
	SceneCastSkillCsReq                                = 1416
	ChessRogueStartCsReq                               = 5423
	EnterRogueMapRoomScRsp                             = 1897
	SaveLogisticsScRsp                                 = 4769
	SyncRogueCommonActionResultScNotify                = 5644
	DiscardRelicCsReq                                  = 549
	AcceptedPamMissionExpireCsReq                      = 4010
	GetNpcMessageGroupCsReq                            = 2710
	HeliobusSnsReadScRsp                               = 5845
	SpringRecoverSingleAvatarCsReq                     = 1489
	MuseumTargetMissionFinishNotify                    = 4365
	QuitTreasureDungeonCsReq                           = 4496
	GetBagCsReq                                        = 510
	TakePictureScRsp                                   = 4160
	ChangeStoryLineScRsp                               = 6216
	SetRedPointStatusScNotify                          = 71
	FinishChessRogueNousSubStoryCsReq                  = 5587
	GetLoginActivityScRsp                              = 2633
	RaidCollectionDataCsReq                            = 6959
	DressAvatarScRsp                                   = 370
	ChessRogueQueryCsReq                               = 5466
	TeleportToMissionResetPointCsReq                   = 1269
	UpgradeAreaCsReq                                   = 4358
	ChessRogueNousEditDiceScRsp                        = 5416
	DailyActiveInfoNotify                              = 3316
	ExchangeGachaCeilingCsReq                          = 1977
	SceneCastSkillMpUpdateScNotify                     = 1402
	GetRollShopInfoScRsp                               = 6917
	AetherDivideSkillItemScNotify                      = 4867
	RaidKickByServerScNotify                           = 2266
	GetFriendAssistListCsReq                           = 2993
	TravelBrochureSetCustomValueCsReq                  = 6436
	DelMailScRsp                                       = 860
	MonopolyUpgradeAssetScRsp                          = 7008
	ChessRogueSkipTeachingLevelCsReq                   = 5448
	SubmitOrigamiItemScRsp                             = 4102
	GetGachaCeilingCsReq                               = 1916
	SyncClientResVersionScRsp                          = 126
	FinishFirstTalkByPerformanceNpcCsReq               = 2158
	StartCocoonStageCsReq                              = 1407
	DeleteFriendCsReq                                  = 2958
	PlayBackGroundMusicScRsp                           = 3145
	ResetMapRotationRegionCsReq                        = 6858
	ChessRogueEnterCellScRsp                           = 5533
	MuseumInfoChangedScNotify                          = 4399
	MakeDrinkCsReq                                     = 6981
	StartBattleCollegeCsReq                            = 5745
	GetCurChallengeScRsp                               = 1703
	TravelBrochureApplyPasterCsReq                     = 6477
	GetTreasureDungeonActivityDataCsReq                = 4403
	MonopolyGuessChooseCsReq                           = 7097
	AcceptMultipleExpeditionScRsp                      = 2599
	GetMainMissionCustomValueCsReq                     = 1232
	FinishCosumeItemMissionScRsp                       = 1295
	MonopolySelectOptionScRsp                          = 7003
	PlayerReturnInfoQueryCsReq                         = 4600
	GetPhoneDataScRsp                                  = 5133
	GetLevelRewardCsReq                                = 8
	GetChessRogueStoryAeonTalkInfoScRsp                = 5513
	UpgradeAreaStatCsReq                               = 4336
	EnhanceRogueBuffCsReq                              = 1806
	TakeApRewardScRsp                                  = 3333
	ChessRogueQueryBpScRsp                             = 5599
	GetNpcTakenRewardScRsp                             = 2133
	GetShareDataCsReq                                  = 4142
	TakeMailAttachmentCsReq                            = 877
	EnterTelevisionActivityStageCsReq                  = 6967
	MonopolyGameSettleScNotify                         = 7012
	HandleFriendScRsp                                  = 2903
	AetherDivideRefreshEndlessScNotify                 = 4801
	SearchPlayerCsReq                                  = 2969
	MonopolyActionResultScNotify                       = 7042
	HeliobusSelectSkillScRsp                           = 5899
	MonopolyGuessBuyInformationCsReq                   = 7049
	SelectRogueDialogueEventScRsp                      = 1819
	EnterSectionCsReq                                  = 1469
	EnterChessRogueAeonRoomCsReq                       = 5401
	QuitLineupScRsp                                    = 726
	DeleteBlacklistScRsp                               = 2965
	EnterFightActivityStageCsReq                       = 3645
	TakeKilledPunkLordMonsterScoreScRsp                = 3267
	ComposeLimitNumCompleteNotify                      = 583
	SpaceZooBornCsReq                                  = 6742
	GetJukeboxDataScRsp                                = 3133
	SetNicknameScRsp                                   = 37
	TakeRogueMiracleHandbookRewardCsReq                = 5697
	GameplayCounterRecoverScRsp                        = 1414
	GetSpringRecoverDataScRsp                          = 1424
	TakeExpeditionRewardScRsp                          = 2526
	GetVideoVersionKeyScRsp                            = 59
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6026
	SyncTurnFoodNotify                                 = 586
	FinishTalkMissionCsReq                             = 1242
	RefreshAlleyOrderCsReq                             = 4799
	PlayerReturnTakePointRewardScRsp                   = 4560
	FinishSectionIdScRsp                               = 2726
	SpaceZooDeleteCatCsReq                             = 6800
	GetFarmStageGachaInfoScRsp                         = 1345
	JoinLineupScRsp                                    = 760
	GetCurBattleInfoCsReq                              = 116
	DeployRotaterScRsp                                 = 6860
	ChooseBoxingClubResonanceCsReq                     = 4203
	SetHeroBasicTypeScRsp                              = 12
	GetKilledPunkLordMonsterDataScRsp                  = 3206
	ChallengeRaidNotify                                = 2270
	AddBlacklistCsReq                                  = 2902
	GetRogueAeonInfoCsReq                              = 1817
	AcceptExpeditionScRsp                              = 2545
	FightActivityDataChangeScNotify                    = 3642
	TrialBackGroundMusicScRsp                          = 3126
	OpenRogueChestCsReq                                = 1889
	QuitRogueCsReq                                     = 1812
	ComposeSelectedRelicScRsp                          = 506
	GetLineupAvatarDataScRsp                           = 758
	MonopolyConfirmRandomCsReq                         = 7091
	ChessRogueSelectCellCsReq                          = 5534
	AcceptMissionEventCsReq                            = 1276
	PunkLordRaidTimeOutScNotify                        = 3291
	EnableRogueTalentScRsp                             = 1841
	TakeAssistRewardScRsp                              = 2918
	GetSocialEventServerCacheScRsp                     = 7054
	GetChessRogueNousStoryInfoCsReq                    = 5455
	DrinkMakerChallengeScRsp                           = 6998
	SpaceZooOpCatteryCsReq                             = 6777
	AlleyGuaranteedFundsScRsp                          = 4772
	MarkReadMailCsReq                                  = 842
	GetRogueShopMiracleInfoScRsp                       = 5616
	GetRogueBuffEnhanceInfoCsReq                       = 1896
	SelectInclinationTextCsReq                         = 2200
	GetOfferingInfoScRsp                               = 6937
	PunkLordBattleResultScNotify                       = 3296
	BuyBpLevelCsReq                                    = 3060
	SetAssistScRsp                                     = 2912
	PlayerReturnPointChangeScNotify                    = 4545
	UpdateEnergyScNotify                               = 6802
	TravelBrochureApplyPasterListScRsp                 = 6494
	SetPlayerInfoScRsp                                 = 47
	SetIsDisplayAvatarInfoCsReq                        = 2877
	EnterAetherDivideSceneScRsp                        = 4833
	CommonRogueQueryCsReq                              = 5689
	SwapLineupCsReq                                    = 800
	ChessRogueLeaveCsReq                               = 5493
	PlayerReturnForceFinishScNotify                    = 4503
	DoGachaCsReq                                       = 1942
	SpringRecoverScRsp                                 = 1468
	ChessRogueConfirmRollScRsp                         = 5532
	ReplaceLineupScRsp                                 = 708
	SubmitMonsterResearchActivityMaterialScRsp         = 2666
	BatchMarkChatEmojiScRsp                            = 3995
	TravelBrochurePageResetScRsp                       = 6466
	LogisticsDetonateStarSkiffCsReq                    = 4787
	DeleteSummonUnitScRsp                              = 1435
	GetSaveRaidScRsp                                   = 2202
	GetFightActivityDataCsReq                          = 3610
	MonopolyGetRaffleTicketCsReq                       = 7074
	MuseumRandomEventQueryCsReq                        = 4366
	SyncChessRogueNousSubStoryScNotify                 = 5568
	PlayerLoginFinishScRsp                             = 19
	MonopolyGameRaiseRatioScRsp                        = 7038
	EnterRogueCsReq                                    = 1816
	AetherDivideLineupScNotify                         = 4812
	SyncChessRogueMainStoryFinishScNotify              = 5496
	AcceptActivityExpeditionCsReq                      = 2570
	GetAuthkeyCsReq                                    = 2
	ExpUpEquipmentCsReq                                = 503
	MonopolyRollRandomScRsp                            = 7002
	TravelBrochureUpdatePasterPosCsReq                 = 6403
	ChessRogueCheatRollCsReq                           = 5413
	GetRogueInitialScoreScRsp                          = 1849
	GetSecretKeyInfoScRsp                              = 56
	MarkChatEmojiCsReq                                 = 3903
	ShareCsReq                                         = 4110
	SyncRogueGetItemScNotify                           = 1805
	MonopolyClickCellCsReq                             = 7017
	OpenTreasureDungeonGridCsReq                       = 4436
	MissionAcceptScNotify                              = 1201
	ChessRogueGoAheadCsReq                             = 5549
	SceneEntityTeleportScRsp                           = 1479
	GetChessRogueStoryInfoScRsp                        = 5441
	GmTalkCsReq                                        = 70
	HeroBasicTypeChangedNotify                         = 49
	UnlockSkilltreeScRsp                               = 360
	MonopolyLikeCsReq                                  = 7084
	PlayerSyncScNotify                                 = 610
	SetBoxingClubResonanceLineupScRsp                  = 4295
	GetAllRedDotDataScRsp                              = 5933
	InteractPropCsReq                                  = 1442
	SetHeadIconCsReq                                   = 2842
	ChessRogueGiveUpRollCsReq                          = 5428
	EnhanceChessRogueBuffScRsp                         = 5525
	ChessRogueGiveUpScRsp                              = 5453
	SendMsgScRsp                                       = 3933
	StartTrialActivityScRsp                            = 2680
	PromoteAvatarCsReq                                 = 377
	ReportPlayerScRsp                                  = 2908
	MonopolyQuizDurationChangeScNotify                 = 7023
	TakeTalkRewardScRsp                                = 2145
	EnterTrialActivityStageCsReq                       = 2684
	GetAllServerPrefsDataScRsp                         = 6133
	GetMbtiReportScRsp                                 = 7073
	SyncRogueStatusScNotify                            = 1851
	ApplyFriendScRsp                                   = 2926
	GetMonopolyInfoScRsp                               = 7033
	StartTimedCocoonStageScRsp                         = 1457
	GetDailyActiveInfoScRsp                            = 3345
	StartRogueCsReq                                    = 1842
	TakeAllRewardScRsp                                 = 3100
	TravelBrochureSetPageDescStatusCsReq               = 6499
	GetAllSaveRaidCsReq                                = 2299
	EnterRogueEndlessActivityStageCsReq                = 6042
	GetChallengeCsReq                                  = 1710
	TakeOfferingRewardScRsp                            = 6940
	SyncDeleteFriendScNotify                           = 2936
	GetCurLineupDataCsReq                              = 742
	ExchangeStaminaCsReq                               = 95
	BattleLogReportCsReq                               = 170
	MonopolyDailySettleScNotify                        = 7048
	BuyGoodsCsReq                                      = 1542
	FinishPlotScRsp                                    = 1133
	BuyRogueShopBuffScRsp                              = 5603
	PrestigeLevelUpScRsp                               = 4794
	ChessRogueReviveAvatarScRsp                        = 5476
	StopRogueAdventureRoomCsReq                        = 5606
	GetRaidInfoCsReq                                   = 2203
	SyncRogueAeonScNotify                              = 1859
	SetStuffToAreaScRsp                                = 4360
	TakeOffAvatarSkinCsReq                             = 308
	MonopolyGetRafflePoolInfoCsReq                     = 7098
	GetRogueScoreRewardInfoCsReq                       = 1863
	AlleyOrderChangedScNotify                          = 4791
	SceneEnterStageCsReq                               = 1496
	TakePrestigeRewardCsReq                            = 4703
	PrepareRogueAdventureRoomScRsp                     = 5642
	TakeBpRewardCsReq                                  = 3045
	AlleyPlacingGameScRsp                              = 4795
	TakeChallengeRewardCsReq                           = 1736
	StartTimedFarmElementCsReq                         = 1409
	GetExhibitScNotify                                 = 4370
	UpdateRedDotDataScRsp                              = 5945
	GetAetherDivideChallengeInfoScRsp                  = 4869
	GetRogueShopBuffInfoScRsp                          = 5677
	TelevisionActivityDataChangeScNotify               = 6961
	GetLoginChatInfoScRsp                              = 3902
	MakeMissionDrinkScRsp                              = 6990
	GroupStateChangeCsReq                              = 1450
	GetPlayerReplayInfoCsReq                           = 3542
	UnlockBackGroundMusicScRsp                         = 3160
	ExchangeRogueRewardKeyScRsp                        = 1873
	GetRogueAdventureRoomInfoScRsp                     = 5636
	TakeActivityExpeditionRewardScRsp                  = 2536
	GetBoxingClubInfoCsReq                             = 4210
	GetUpdatedArchiveDataCsReq                         = 2342
	ReserveStaminaExchangeCsReq                        = 78
	AcceptMainMissionScRsp                             = 1212
	FinishAeonDialogueGroupScRsp                       = 1830
	SceneEntityDieScNotify                             = 1449
	TakeOffEquipmentScRsp                              = 311
	CancelActivityExpeditionScRsp                      = 2558
	TakeChapterRewardScRsp                             = 500
	SetLineupNameCsReq                                 = 776
	LeaveRaidScRsp                                     = 2245
	SpringRefreshCsReq                                 = 1476
	SetCurWaypointScRsp                                = 445
	EnterFantasticStoryActivityStageCsReq              = 4945
	GetMissionEventDataScRsp                           = 1202
	DailyRefreshNotify                                 = 93
	TravelBrochureRemovePasterCsReq                    = 6500
	GmTalkScNotify                                     = 26
	UseTreasureDungeonItemCsReq                        = 4437
	SelectChessRogueNousSubStoryCsReq                  = 5498
	LeaveRogueCsReq                                    = 1877
	PlayerHeartBeatCsReq                               = 46
	GetRogueEndlessActivityDataCsReq                   = 6010
	AetherDivideRefreshEndlessCsReq                    = 4832
	SetCurInteractEntityCsReq                          = 1438
	SpaceZooTakeScRsp                                  = 6736
	ChessRogueReRollDiceScRsp                          = 5430
	SecurityReportCsReq                                = 4103
	BuyRogueShopMiracleScRsp                           = 5700
	GetAvatarDataScRsp                                 = 333
	RankUpEquipmentCsReq                               = 600
	AntiAddictScNotify                                 = 91
	StartAetherDivideChallengeBattleCsReq              = 4877
	TextJoinSaveCsReq                                  = 3810
)

const (
	ServiceConnectionReq     = 10000
	ServiceConnectionRsp     = 10100
	GateLoginGameRsp         = 10001
	GateLoginGameReq         = 10101
	GateToGameMsgNotify      = 10002
	GameToGateMsgNotify      = 10102
	GetAllServiceGateReq     = 10003
	GetAllServiceGateRsp     = 10103
	PlayerLogoutReq          = 10004
	PlayerLogoutRsp          = 10104
	GetAllServiceReq         = 10005
	GetAllServiceRsp         = 10105
	GetAllServiceGameReq     = 10006
	GetAllServiceGameRsp     = 10106
	GameToNodePingReq        = 10007
	GameToNodePingRsp        = 10107
	GateGamePingReq          = 10008
	GateGamePingRsp          = 10108
	GateGamePlayerLoginReq   = 10009
	GateGamePlayerLoginRsp   = 10109
	GetToGamePlayerLogoutReq = 10010
	GetToGamePlayerLogoutRsp = 10110

	PlayerLogoutNotify           = 11000
	SyncPlayerOnlineDataNotify   = 11001
	PlayerLoginNotify            = 11002
	NodeToGsPlayerLogoutNotify   = 11003
	GameToGatePlayerLogoutNotify = 11004

	GmGive       = 12001
	GmWorldLevel = 12002
	DelItem      = 12003
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(CommonRogueQueryScRsp, func() any { return new(proto.CommonRogueQueryScRsp) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(CommonRogueUpdateScNotify, func() any { return new(proto.CommonRogueUpdateScNotify) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(SetFriendMarkCsReq, func() any { return new(proto.SetFriendMarkCsReq) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(SetHeroBasicTypeCsReq, func() any { return new(proto.SetHeroBasicTypeCsReq) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(GetHeroBasicTypeInfoScRsp, func() any { return new(proto.GetHeroBasicTypeInfoScRsp) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(SetFriendMarkScRsp, func() any { return new(proto.SetFriendMarkScRsp) })
	c.regMsg(FinishRogueDialogueGroupCsReq, func() any { return new(proto.FinishRogueDialogueGroupCsReq) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(SelectRogueDialogueEventCsReq, func() any { return new(proto.SelectRogueDialogueEventCsReq) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(SyncRogueDialogueEventDataScNotify, func() any { return new(proto.SyncRogueDialogueEventDataScNotify) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(SelectRogueDialogueEventScRsp, func() any { return new(proto.SelectRogueDialogueEventScRsp) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(SetHeroBasicTypeScRsp, func() any { return new(proto.SetHeroBasicTypeScRsp) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	// server
	c.regMsg(GmGive, func() any { return new(spb.GmGive) })
	c.regMsg(GmWorldLevel, func() any { return new(spb.GmWorldLevel) })
	c.regMsg(DelItem, func() any { return new(spb.DelItem) })
	c.regMsg(ServiceConnectionReq, func() any { return new(spb.ServiceConnectionReq) })
	c.regMsg(ServiceConnectionRsp, func() any { return new(spb.ServiceConnectionRsp) })
	c.regMsg(GateLoginGameRsp, func() any { return new(spb.GateLoginGameRsp) })
	c.regMsg(GateLoginGameReq, func() any { return new(spb.GateLoginGameReq) })
	c.regMsg(GateToGameMsgNotify, func() any { return new(spb.GateToGameMsgNotify) })
	c.regMsg(GameToGateMsgNotify, func() any { return new(spb.GameToGateMsgNotify) })
	c.regMsg(GetAllServiceGateReq, func() any { return new(spb.GetAllServiceGateReq) })
	c.regMsg(GetAllServiceGateRsp, func() any { return new(spb.GetAllServiceGateRsp) })
	c.regMsg(PlayerLogoutReq, func() any { return new(spb.PlayerLogoutReq) })
	c.regMsg(PlayerLogoutRsp, func() any { return new(spb.PlayerLogoutRsp) })
	c.regMsg(GetAllServiceReq, func() any { return new(spb.GetAllServiceReq) })
	c.regMsg(GetAllServiceRsp, func() any { return new(spb.GetAllServiceRsp) })
	c.regMsg(GetAllServiceGameReq, func() any { return new(spb.GetAllServiceGameReq) })
	c.regMsg(GetAllServiceGameRsp, func() any { return new(spb.GetAllServiceGameRsp) })
	c.regMsg(PlayerLogoutNotify, func() any { return new(spb.PlayerLogoutNotify) })
	c.regMsg(SyncPlayerOnlineDataNotify, func() any { return new(spb.SyncPlayerOnlineDataNotify) })
	c.regMsg(GateGamePingReq, func() any { return new(spb.GateGamePingReq) })
	c.regMsg(GateGamePingRsp, func() any { return new(spb.GateGamePingRsp) })
	c.regMsg(GateGamePlayerLoginReq, func() any { return new(spb.GateGamePlayerLoginReq) })
	c.regMsg(GateGamePlayerLoginRsp, func() any { return new(spb.GateGamePlayerLoginRsp) })
	c.regMsg(GetToGamePlayerLogoutReq, func() any { return new(spb.GetToGamePlayerLogoutReq) })
	c.regMsg(GetToGamePlayerLogoutRsp, func() any { return new(spb.GetToGamePlayerLogoutRsp) })
	c.regMsg(GameToGatePlayerLogoutNotify, func() any { return new(spb.GameToGatePlayerLogoutNotify) })

	c.regMsg(PlayerLoginNotify, func() any { return new(spb.PlayerLoginNotify) })
	c.regMsg(NodeToGsPlayerLogoutNotify, func() any { return new(spb.NodeToGsPlayerLogoutNotify) })
	c.regMsg(GameToNodePingReq, func() any { return new(spb.GameToNodePingReq) })
	c.regMsg(GameToNodePingRsp, func() any { return new(spb.GameToNodePingRsp) })

}
