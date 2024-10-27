package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

const (
	MuseumTargetRewardNotify                           = 4362
	GetGameStateServiceConfigCsReq                     = 94
	SelectPhoneThemeScRsp                              = 5133
	ModifyRelicFilterPlanScRsp                         = 558
	UseTreasureDungeonItemCsReq                        = 4404
	ChessRogueNousGetRogueTalentInfoCsReq              = 5496
	DressAvatarScRsp                                   = 338
	GetPlayerReturnMultiDropInfoScRsp                  = 4679
	LeaveRogueScRsp                                    = 1812
	SpaceZooDeleteCatCsReq                             = 6728
	TakeAssistRewardCsReq                              = 2902
	ChooseBoxingClubStageOptionalBuffScRsp             = 4205
	ClockParkGetInfoCsReq                              = 7229
	SwordTrainingUnlockSyncScNotify                    = 7488
	FinishTutorialGuideCsReq                           = 1678
	WaypointShowNewCsNotify                            = 433
	StartTimedFarmElementCsReq                         = 1461
	TakeOffRelicCsReq                                  = 320
	SetAssistAvatarScRsp                               = 2889
	BuyRogueShopBuffCsReq                              = 5638
	LeaveMapRotationRegionScNotify                     = 6866
	GetPlayerDetailInfoCsReq                           = 2983
	GetMultiPathAvatarInfoCsReq                        = 68
	TakePrestigeRewardScRsp                            = 4756
	MonopolyRollDiceScRsp                              = 7033
	AvatarExpUpCsReq                                   = 383
	GetLevelRewardTakenListScRsp                       = 52
	StartTrackPhotoStageCsReq                          = 7560
	SubmitOrigamiItemCsReq                             = 4166
	MultiplayerMatch3FinishScNotify                    = 1028
	ChessRogueRollDiceScRsp                            = 5423
	RogueMagicSetAutoDressInMagicUnitScRsp             = 7702
	StartRaidCsReq                                     = 2298
	LobbyJoinScRsp                                     = 7352
	MuseumTakeCollectRewardScRsp                       = 4323
	TakeApRewardCsReq                                  = 3398
	SpaceZooBornScRsp                                  = 6742
	HandleRogueCommonPendingActionScRsp                = 5658
	ContentPackageGetDataCsReq                         = 7529
	UnlockAvatarSkinScNotify                           = 400
	RogueTournDeleteArchiveScRsp                       = 6089
	FinishChessRogueNousSubStoryCsReq                  = 5526
	SwordTrainingStartGameScRsp                        = 7465
	UpgradeAreaScRsp                                   = 4322
	TakeOffRelicScRsp                                  = 350
	AetherDivideTakeChallengeRewardCsReq               = 4855
	DiscardRelicCsReq                                  = 515
	PlayerLoginFinishCsReq                             = 90
	ReBattleAfterBattleLoseCsNotify                    = 189
	GetPlatformPlayerInfoCsReq                         = 2925
	GetNpcStatusCsReq                                  = 2783
	MuseumFundsChangedScNotify                         = 4320
	ChessRogueUpdateBoardScNotify                      = 5406
	MonopolyConfirmRandomScRsp                         = 7031
	AlleyEventEffectNotify                             = 4738
	MusicRhythmSaveSongConfigDataScRsp                 = 7585
	TakeChapterRewardCsReq                             = 412
	SetPlayerInfoScRsp                                 = 25
	EvolveBuildCoinNotify                              = 7128
	ChessRogueEnterNextLayerCsReq                      = 5506
	PunkLordMonsterInfoScNotify                        = 3266
	CurAssistChangedNotify                             = 2957
	PrepareRogueAdventureRoomCsReq                     = 5671
	RetcodeNotify                                      = 76
	DoGachaInRollShopCsReq                             = 6913
	MonopolyTakeRaffleTicketRewardScRsp                = 7026
	LobbyJoinCsReq                                     = 7392
	ClockParkQuitScriptScRsp                           = 7225
	StartRaidScRsp                                     = 2271
	SetMultipleAvatarPathsScRsp                        = 7
	TrainRefreshTimeNotify                             = 3779
	ChangeStoryLineFinishScNotify                      = 6277
	HeliobusSnsLikeScRsp                               = 5812
	GetAllSaveRaidCsReq                                = 2245
	RogueMagicEnterCsReq                               = 7783
	StartAetherDivideSceneBattleCsReq                  = 4879
	GetTrialActivityDataCsReq                          = 2687
	GetSocialEventServerCacheCsReq                     = 7086
	HeliobusSnsReadCsReq                               = 5883
	AddAvatarScNotify                                  = 389
	TakeBpRewardCsReq                                  = 3042
	StartAetherDivideChallengeBattleScRsp              = 4812
	GetChapterCsReq                                    = 479
	GetSingleRedDotParamGroupScRsp                     = 5977
	TakeApRewardScRsp                                  = 3371
	SyncEntityBuffChangeListScNotify                   = 1489
	FinishSectionIdScRsp                               = 2712
	AlleyGuaranteedFundsCsReq                          = 4788
	SwordTrainingTurnActionScRsp                       = 7494
	DailyFirstEnterMonopolyActivityCsReq               = 7089
	JoinLineupScRsp                                    = 777
	FinishTutorialGuideScRsp                           = 1656
	GetCrossInfoCsReq                                  = 7310
	SetBoxingClubResonanceLineupCsReq                  = 4289
	MonopolyUpgradeAssetCsReq                          = 7052
	MultiplayerFightGiveUpCsReq                        = 1079
	HeliobusSnsCommentCsReq                            = 5828
	RogueEndlessActivityBattleEndScNotify              = 6010
	TakeOffEquipmentScRsp                              = 356
	AddRelicFilterPlanScRsp                            = 554
	ContentPackageUnlockCsReq                          = 7502
	GetFarmStageGachaInfoScRsp                         = 1342
	StartAetherDivideStageBattleCsReq                  = 4804
	PunkLordBattleResultScNotify                       = 3252
	GetCurAssistCsReq                                  = 2988
	RogueModifierSelectCellCsReq                       = 5342
	ChessRogueGoAheadCsReq                             = 5431
	GetCurLineupDataCsReq                              = 783
	AlleyEventChangeNotify                             = 4728
	GetLineupAvatarDataCsReq                           = 756
	TravelBrochureSelectMessageCsReq                   = 6479
	MonopolyGetRaffleTicketCsReq                       = 7072
	AetherDivideLineupScNotify                         = 4814
	ExchangeStaminaCsReq                               = 22
	SwordTrainingGameSyncChangeScNotify                = 7479
	GetBoxingClubInfoCsReq                             = 4298
	GetChallengeRaidInfoScRsp                          = 2233
	GetCrossInfoScRsp                                  = 7324
	MarkAvatarScRsp                                    = 306
	GetArchiveDataCsReq                                = 2398
	MonopolySelectOptionScRsp                          = 7078
	GetStuffScNotify                                   = 4328
	TakeMultipleExpeditionRewardScRsp                  = 2550
	RogueTournGetAllArchiveScRsp                       = 6081
	TakeBpRewardScRsp                                  = 3079
	EquipAetherDividePassiveSkillCsReq                 = 4866
	BuyGoodsCsReq                                      = 1583
	GetFarmStageGachaInfoCsReq                         = 1383
	RogueModifierSelectCellScRsp                       = 5379
	RogueTournGetPermanentTalentInfoCsReq              = 6100
	RogueTournEnterRoomCsReq                           = 6079
	UpdateTrackMainMissionIdScRsp                      = 1202
	PunkLordRaidTimeOutScNotify                        = 3250
	EnterMapRotationRegionCsReq                        = 6898
	SetFriendRemarkNameScRsp                           = 2960
	GetFriendChallengeLineupCsReq                      = 2954
	MatchThreeSyncDataScNotify                         = 7444
	TakeRollShopRewardScRsp                            = 6915
	RemoveStuffFromAreaCsReq                           = 4333
	GetHeartDialInfoCsReq                              = 6398
	RogueTournLeaveRogueCocoonSceneCsReq               = 6029
	GetFriendChallengeLineupScRsp                      = 2969
	BattleLogReportCsReq                               = 138
	RogueTournConfirmSettleScRsp                       = 6049
	PrestigeLevelUpCsReq                               = 4704
	DifficultyAdjustmentGetDataCsReq                   = 4123
	TravelBrochureGetDataCsReq                         = 6498
	RogueMagicScepterTakeOffUnitScRsp                  = 7706
	HeliobusStartRaidCsReq                             = 5860
	RaidCollectionDataCsReq                            = 6954
	ClockParkGetInfoScRsp                              = 7237
	SwordTrainingSelectEndingCsReq                     = 7489
	MatchThreeLevelEndCsReq                            = 7442
	SetSignatureCsReq                                  = 2838
	SubMissionRewardScNotify                           = 1300
	GetFriendChallengeDetailCsReq                      = 2958
	EnterRogueEndlessActivityStageScRsp                = 6003
	SyncChessRogueNousMainStoryScNotify                = 5584
	TrainVisitorRewardSendNotify                       = 3777
	GetReplayTokenCsReq                                = 3598
	StartAlleyEventCsReq                               = 4733
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6009
	EnterFeverTimeActivityStageCsReq                   = 7152
	GetChessRogueBuffEnhanceInfoCsReq                  = 5501
	StartPartialChallengeScRsp                         = 1731
	SetBoxingClubResonanceLineupScRsp                  = 4222
	FeatureSwitchClosedScNotify                        = 84
	EnterSummonActivityStageScRsp                      = 7563
	AcceptActivityExpeditionCsReq                      = 2538
	ServerSimulateBattleFinishScNotify                 = 156
	SubmitOfferingItemScRsp                            = 6930
	AlleyPlacingGameScRsp                              = 4722
	GetUpdatedArchiveDataScRsp                         = 2342
	StrongChallengeActivityBattleEndScNotify           = 6679
	GetMarkItemListScRsp                               = 593
	UpdateMovieRacingDataCsReq                         = 4152
	MonopolyGameBingoFlipCardScRsp                     = 7055
	TravelBrochureApplyPasterListScRsp                 = 6460
	PlayerReturnTakeRewardScRsp                        = 4512
	TakeActivityExpeditionRewardScRsp                  = 2566
	ReportPlayerScRsp                                  = 2911
	SyncRogueCommonDialogueOptionFinishScNotify        = 5690
	SyncDeleteFriendScNotify                           = 2966
	AlleyShopLevelScNotify                             = 4711
	SwordTrainingDailyPhaseConfirmCsReq                = 7460
	GetMailCsReq                                       = 898
	UnlockTeleportNotify                               = 1444
	GetChessRogueNousStoryInfoCsReq                    = 5410
	BatchGetQuestDataCsReq                             = 966
	GetGunPlayDataCsReq                                = 4110
	GetChallengeScRsp                                  = 1771
	GetLoginActivityScRsp                              = 2671
	GetMissionDataScRsp                                = 1271
	GetFightActivityDataScRsp                          = 3671
	ClearAetherDividePassiveSkillScRsp                 = 4820
	GetFriendListInfoCsReq                             = 2998
	RogueMagicReviveAvatarCsReq                        = 7731
	GetQuestDataCsReq                                  = 998
	GetAssistHistoryCsReq                              = 2963
	SetNicknameScRsp                                   = 4
	UpdateTrackMainMissionIdCsReq                      = 1236
	ClientObjUploadCsReq                               = 99
	SearchPlayerCsReq                                  = 2962
	GetMuseumInfoCsReq                                 = 4398
	LeaveAetherDivideSceneScRsp                        = 4842
	TravelBrochurePageUnlockScNotify                   = 6483
	GetFriendApplyListInfoScRsp                        = 2977
	RankUpEquipmentCsReq                               = 528
	TakeRogueMiracleHandbookRewardScRsp                = 5625
	DeactivateFarmElementCsReq                         = 1419
	GetMissionStatusScRsp                              = 1204
	SceneUpdatePositionVersionNotify                   = 1456
	StartAetherDivideChallengeBattleCsReq              = 4833
	DressAvatarCsReq                                   = 328
	GetNpcTakenRewardScRsp                             = 2171
	ChessRogueNousEditDiceScRsp                        = 5470
	RaidCollectionEnterNextRaidCsReq                   = 6950
	ChessRogueUpdateDiceInfoScNotify                   = 5430
	LogisticsGameCsReq                                 = 4783
	GetChatEmojiListScRsp                              = 3938
	RogueMagicGetMiscRealTimeDataScRsp                 = 7757
	ChessRogueChangeyAeonDimensionNotify               = 5550
	MonopolyMoveCsReq                                  = 7012
	RogueMagicGetTalentInfoScRsp                       = 7793
	GetStoryLineInfoScRsp                              = 6271
	MonopolyTakePhaseRewardScRsp                       = 7099
	SetStuffToAreaScRsp                                = 4377
	MonopolyQuizDurationChangeScNotify                 = 7040
	ChessRogueNousEnableRogueTalentScRsp               = 5494
	GetRogueInfoScRsp                                  = 1871
	ClockParkGetOngoingScriptInfoScRsp                 = 7201
	EnhanceChessRogueBuffCsReq                         = 5530
	FinishFirstTalkNpcScRsp                            = 2112
	RogueMagicUnitReforgeScRsp                         = 7714
	SetGenderCsReq                                     = 2
	ChessRogueSelectBpCsReq                            = 5454
	SetCurWaypointCsReq                                = 483
	CancelMatchScRsp                                   = 7302
	TextJoinBatchSaveScRsp                             = 3877
	TakeOffAvatarSkinCsReq                             = 311
	BatchMarkChatEmojiCsReq                            = 3989
	EnterRogueEndlessActivityStageCsReq                = 6007
	PlayerReturnInfoQueryCsReq                         = 4528
	GetDailyActiveInfoScRsp                            = 3342
	LobbyKickOutScRsp                                  = 7351
	SwordTrainingStoryConfirmCsReq                     = 7480
	RogueMagicEnterRoomCsReq                           = 7728
	GetRogueExhibitionCsReq                            = 5624
	EvolveBuildTakeExpRewardCsReq                      = 7130
	MuseumTakeCollectRewardCsReq                       = 4306
	QuitBattleScRsp                                    = 142
	MultiplayerGetFightGateScRsp                       = 1042
	SyncRogueSeasonFinishScNotify                      = 1855
	MarkRelicFilterPlanScRsp                           = 549
	RogueTournDifficultyCompNotify                     = 6082
	MonopolyGuessDrawScNotify                          = 7092
	TakeMaterialSubmitActivityRewardCsReq              = 2604
	DeleteFriendScRsp                                  = 2922
	PlayerKickOutScNotify                              = 28
	RogueWorkbenchHandleFuncCsReq                      = 5647
	SetStuffToAreaCsReq                                = 4379
	RogueTournEnterLayerScRsp                          = 6062
	GetFirstTalkByPerformanceNpcCsReq                  = 2178
	MonopolyRollDiceCsReq                              = 7077
	ChessRogueNousGetRogueTalentInfoScRsp              = 5517
	StopRogueAdventureRoomScRsp                        = 5700
	ResetMapRotationRegionScRsp                        = 6822
	StartBoxingClubBattleScRsp                         = 4277
	QuitWolfBroGameScRsp                               = 6512
	GetRecyleTimeCsReq                                 = 562
	AddBlacklistCsReq                                  = 2905
	RogueMagicScepterDressInUnitCsReq                  = 7710
	UnlockSkilltreeCsReq                               = 379
	StartMatchScRsp                                    = 7337
	MatchThreeSetBirdPosScRsp                          = 7424
	RaidCollectionEnterNextRaidScRsp                   = 6948
	SpaceZooDeleteCatScRsp                             = 6738
	GetSaveRaidScRsp                                   = 2205
	ChessRogueGiveUpScRsp                              = 5556
	MakeDrinkCsReq                                     = 6993
	FantasticStoryActivityBattleEndScNotify            = 4977
	LobbyInviteCsReq                                   = 7399
	ChessRogueConfirmRollCsReq                         = 5420
	CancelActivityExpeditionScRsp                      = 2589
	ChallengeSettleNotify                              = 1733
	QuitLineupScRsp                                    = 712
	QuitRogueScRsp                                     = 1888
	ChessRogueLeaveScRsp                               = 5544
	ExchangeRogueBuffWithMiracleCsReq                  = 5650
	SyncRogueCommonActionResultScNotify                = 5692
	SpringRefreshScRsp                                 = 1450
	SetAssistScRsp                                     = 2914
	ChangeScriptEmotionCsReq                           = 6383
	SummonPetScRsp                                     = 7602
	TakeAssistRewardScRsp                              = 2965
	AetherDivideRefreshEndlessCsReq                    = 4888
	SyncChessRogueMainStoryFinishScNotify              = 5427
	GetNpcMessageGroupCsReq                            = 2798
	TakeChallengeRewardCsReq                           = 1766
	GetNpcMessageGroupScRsp                            = 2771
	ChessRogueEnterNextLayerScRsp                      = 5543
	ChessRogueConfirmRollScRsp                         = 5413
	AcceptMultipleExpeditionScRsp                      = 2545
	ClockParkFinishScriptScNotify                      = 7212
	GetTelevisionActivityDataScRsp                     = 6977
	EnterTelevisionActivityStageScRsp                  = 6968
	PlayerReturnTakeRewardCsReq                        = 4533
	TakePromotionRewardScRsp                           = 304
	GetDrinkMakerDataCsReq                             = 6994
	ExtraLineupDestroyNotify                           = 710
	MuseumTargetStartNotify                            = 4310
	PVEBattleResultScRsp                               = 171
	GetRogueInitialScoreScRsp                          = 1815
	EnterSwordTrainingExamCsReq                        = 7475
	SelectChessRogueNousSubStoryScRsp                  = 5412
	GetMultiPathAvatarInfoScRsp                        = 46
	MarkReadMailScRsp                                  = 842
	RefreshTriggerByClientScRsp                        = 1451
	AetherDivideRefreshEndlessScNotify                 = 4863
	ComposeItemScRsp                                   = 522
	TakeKilledPunkLordMonsterScoreScRsp                = 3285
	SpaceZooOpCatteryScRsp                             = 6712
	GetAllServerPrefsDataScRsp                         = 6171
	SyncLineupNotify                                   = 778
	MonopolyGetRafflePoolInfoScRsp                     = 7096
	AetherDivideTakeChallengeRewardScRsp               = 4836
	ClockParkUnlockTalentScRsp                         = 7210
	SetAvatarPathCsReq                                 = 37
	GetRndOptionCsReq                                  = 3498
	SetRogueExhibitionCsReq                            = 5609
	HeliobusSnsUpdateScNotify                          = 5878
	PlayerReturnForceFinishScNotify                    = 4578
	ChessRogueEnterCellCsReq                           = 5591
	GetFantasticStoryActivityDataScRsp                 = 4971
	ChessRogueLayerAccountInfoNotify                   = 5466
	TakeAllRewardCsReq                                 = 3012
	GetMonopolyInfoScRsp                               = 7071
	SwordTrainingMarkEndingViewedCsReq                 = 7472
	EnterRogueMapRoomCsReq                             = 1865
	GetAssistListScRsp                                 = 2985
	RelicFilterPlanClearNameScNotify                   = 503
	PlayerSyncScNotify                                 = 698
	GetPunkLordMonsterDataCsReq                        = 3298
	GetFeverTimeActivityDataScRsp                      = 7160
	ExchangeGachaCeilingScRsp                          = 1912
	UpdateRedDotDataScRsp                              = 5942
	SyncRogueAeonLevelUpRewardScNotify                 = 1896
	ClockParkStartScriptCsReq                          = 7224
	GetVideoVersionKeyScRsp                            = 74
	GetFriendBattleRecordDetailCsReq                   = 2976
	MusicRhythmDataCsReq                               = 7591
	SwitchLineupIndexCsReq                             = 705
	RogueTournResetPermanentTalentScRsp                = 6021
	GetSceneMapInfoScRsp                               = 1497
	AcceptMissionEventScRsp                            = 1250
	UpdateMovieRacingDataScRsp                         = 4111
	ChooseBoxingClubStageOptionalBuffCsReq             = 4266
	PlayerReturnSignScRsp                              = 4583
	AvatarPathChangedNotify                            = 41
	RegionStopScNotify                                 = 20
	RemoveRotaterCsReq                                 = 6820
	RaidInfoNotify                                     = 2279
	MonopolyGameGachaCsReq                             = 7088
	SyncRogueVirtualItemInfoScNotify                   = 1861
	HeliobusSnsReadScRsp                               = 5842
	HandleRogueCommonPendingActionCsReq                = 5669
	AcceptExpeditionCsReq                              = 2583
	UpdatePlayerSettingScRsp                           = 96
	EvolveBuildShopAbilityDownCsReq                    = 7143
	GmTalkScNotify                                     = 12
	MonopolyGetDailyInitItemScRsp                      = 7041
	MonopolyMoveScRsp                                  = 7028
	UpdateEnergyScNotify                               = 6805
	TravelBrochureGetDataScRsp                         = 6471
	MonopolyEventLoadUpdateScNotify                    = 7043
	GetGachaInfoScRsp                                  = 1971
	SelectChessRogueSubStoryScRsp                      = 5524
	SetTurnFoodSwitchScRsp                             = 557
	GetGachaCeilingScRsp                               = 1977
	GetFriendLoginInfoScRsp                            = 2992
	WolfBroGameUseBulletScRsp                          = 6589
	GetFantasticStoryActivityDataCsReq                 = 4998
	UnlockBackGroundMusicCsReq                         = 3179
	TravelBrochureApplyPasterScRsp                     = 6412
	StartRogueCsReq                                    = 1883
	FinishChessRogueNousSubStoryScRsp                  = 5513
	SyncTaskScRsp                                      = 1233
	PlayBackGroundMusicCsReq                           = 3183
	EvolveBuildFinishScNotify                          = 7114
	GeneralVirtualItemDataNotify                       = 525
	GetDrinkMakerDataScRsp                             = 6997
	StopRogueAdventureRoomCsReq                        = 5610
	ChangeScriptEmotionScRsp                           = 6342
	SubmitMaterialSubmitActivityMaterialScRsp          = 2631
	FinishTutorialCsReq                                = 1628
	SyncRogueAreaUnlockScNotify                        = 1847
	StartBattleCollegeCsReq                            = 5742
	RogueMagicUnitReforgeCsReq                         = 7716
	RogueModifierUpdateNotify                          = 5312
	FightHeartBeatScRsp                                = 30077
	GetRelicFilterPlanScRsp                            = 535
	HeliobusActivityDataScRsp                          = 5871
	RogueGetGambleInfoScRsp                            = 5637
	GetLoginChatInfoCsReq                              = 3966
	GetLoginChatInfoScRsp                              = 3905
	AetherDivideRefreshEndlessScRsp                    = 4893
	SwordTrainingTurnActionCsReq                       = 7452
	GetSaveLogisticsMapCsReq                           = 4785
	QuitRogueCsReq                                     = 1814
	GetMonopolyDailyReportCsReq                        = 7017
	GetChallengeGroupStatisticsScRsp                   = 1720
	GetMissionEventDataCsReq                           = 1266
	BuyNpcStuffCsReq                                   = 4383
	PromoteEquipmentCsReq                              = 583
	GetFriendApplyListInfoCsReq                        = 2979
	GetAetherDivideChallengeInfoScRsp                  = 4862
	GetActivityScheduleConfigCsReq                     = 2679
	FightActivityDataChangeScNotify                    = 3683
	GetChatFriendHistoryCsReq                          = 3933
	GetOfferingInfoScRsp                               = 6937
	MatchThreeLevelEndScRsp                            = 7402
	FinishEmotionDialoguePerformanceScRsp              = 6312
	StartPartialChallengeCsReq                         = 1750
	GetGachaInfoCsReq                                  = 1998
	MonopolyClickMbtiReportScRsp                       = 7081
	DrinkMakerUpdateTipsNotify                         = 6999
	HeliobusActivityDataCsReq                          = 5898
	GetPlayerReturnMultiDropInfoCsReq                  = 4642
	AlleyTakeEventRewardCsReq                          = 4763
	SwapLineupScRsp                                    = 738
	EnterChessRogueAeonRoomCsReq                       = 5585
	FinishChessRogueSubStoryCsReq                      = 5566
	SyncClientResVersionCsReq                          = 133
	GetChallengeGroupStatisticsCsReq                   = 1745
	GetFightActivityDataCsReq                          = 3698
	MusicRhythmDataScRsp                               = 7577
	GetPrivateChatHistoryCsReq                         = 3979
	MonopolyClickCellCsReq                             = 7075
	ChessRogueNousEditDiceCsReq                        = 5554
	ChessRogueQueryAeonDimensionsCsReq                 = 5402
	FightMatch3DataCsReq                               = 30198
	HeartDialTraceScriptCsReq                          = 6338
	UnlockBackGroundMusicScRsp                         = 3177
	GetMonopolyInfoCsReq                               = 7098
	ClearAetherDividePassiveSkillCsReq                 = 4845
	DifficultyAdjustmentGetDataScRsp                   = 4185
	MakeDrinkScRsp                                     = 6990
	ReserveStaminaExchangeCsReq                        = 27
	FightFestUpdateChallengeRecordNotify               = 7274
	ChessRogueEnterCsReq                               = 5564
	DestroyItemCsReq                                   = 516
	SwitchAetherDivideLineUpSlotScRsp                  = 4831
	MonopolyGetRaffleTicketScRsp                       = 7074
	SpringRecoverScRsp                                 = 1469
	SetCurInteractEntityScRsp                          = 1414
	ChessRogueUpdateActionPointScNotify                = 5575
	LobbyGetInfoScRsp                                  = 7393
	GetMbtiReportScRsp                                 = 7049
	LastSpringRefreshTimeNotify                        = 1431
	GetAssistListCsReq                                 = 2923
	GetAvatarDataScRsp                                 = 371
	FightMatch3ForceUpdateNotify                       = 30156
	ClockParkBattleEndScNotify                         = 7205
	HeliobusUpgradeLevelCsReq                          = 5889
	RogueMagicAreaUpdateScNotify                       = 7722
	SceneCastSkillCostMpScRsp                          = 1466
	GetAllRedDotDataScRsp                              = 5971
	GetStrongChallengeActivityDataScRsp                = 6671
	SubmitEmotionItemCsReq                             = 6379
	SetAssistAvatarCsReq                               = 2856
	LeaveTrialActivityCsReq                            = 2684
	LockRelicScRsp                                     = 520
	ModifyRelicFilterPlanCsReq                         = 569
	RogueMagicQueryCsReq                               = 7760
	ChangeLineupLeaderCsReq                            = 722
	EvolveBuildQueryInfoCsReq                          = 7129
	SelectChessRogueNousSubStoryCsReq                  = 5574
	SyncRogueExploreWinScNotify                        = 1863
	TravelBrochurePageResetScRsp                       = 6431
	DoGachaInRollShopScRsp                             = 6910
	ContentPackageTransferScNotify                     = 7510
	FightMatch3TurnStartScNotify                       = 30142
	TakeChapterRewardScRsp                             = 428
	TakePunkLordPointRewardScRsp                       = 3222
	MarkItemScRsp                                      = 555
	SetFriendRemarkNameCsReq                           = 2904
	GetSingleRedDotParamGroupCsReq                     = 5979
	RogueTournRenameArchiveScRsp                       = 6057
	ClockParkQuitScriptCsReq                           = 7246
	GetMapRotationDataScRsp                            = 6856
	MonopolyGuessBuyInformationCsReq                   = 7015
	StartPunkLordRaidScRsp                             = 3242
	EvolveBuildTakeExpRewardScRsp                      = 7112
	PlayerLoginFinishScRsp                             = 73
	StartFinishMainMissionScNotify                     = 1285
	TakeAllApRewardCsReq                               = 3377
	EvolveBuildGiveupCsReq                             = 7124
	GetCurBattleInfoScRsp                              = 177
	GetWolfBroGameDataCsReq                            = 6528
	LobbyQuitScRsp                                     = 7398
	GetMissionDataCsReq                                = 1298
	MusicRhythmUnlockTrackScNotify                     = 7598
	GiveUpBoxingClubChallengeScRsp                     = 4212
	StartAlleyEventScRsp                               = 4712
	MonopolyGameRaiseRatioCsReq                        = 7085
	BattleCollegeDataChangeScNotify                    = 5783
	GetEnteredSceneCsReq                               = 1468
	MonopolyLikeScRsp                                  = 7067
	FinishTutorialScRsp                                = 1638
	ClientDownloadDataScNotify                         = 40
	SwordTrainingSelectEndingScRsp                     = 7470
	MusicRhythmStartLevelCsReq                         = 7599
	SaveLogisticsCsReq                                 = 4800
	ReEnterLastElementStageCsReq                       = 1486
	GetRogueShopMiracleInfoScRsp                       = 5679
	FinishAeonDialogueGroupScRsp                       = 1841
	MultiplayerFightGameStartScNotify                  = 1033
	WolfBroGameActivateBulletCsReq                     = 6505
	MonopolySelectOptionCsReq                          = 7038
	RogueTournLeaveScRsp                               = 6050
	PlayerGetTokenScRsp                                = 77
	SetForbidOtherApplyFriendScRsp                     = 2995
	MonopolyAcceptQuizScRsp                            = 7002
	MonopolyGuessBuyInformationScRsp                   = 7019
	DeployRotaterScRsp                                 = 6877
	EnterFantasticStoryActivityStageCsReq              = 4942
	TextJoinQueryScRsp                                 = 3842
	WolfBroGameUseBulletCsReq                          = 6556
	GetWolfBroGameDataScRsp                            = 6538
	EnterStrongChallengeActivityStageScRsp             = 6642
	TravelBrochureApplyPasterListCsReq                 = 6404
	FightTreasureDungeonMonsterScRsp                   = 4420
	MonopolyReRollRandomScRsp                          = 7020
	WolfBroGameExplodeMonsterScRsp                     = 6550
	RogueTournEnterLayerCsReq                          = 6038
	DoGachaScRsp                                       = 1942
	QuitWolfBroGameCsReq                               = 6533
	SetAvatarPathScRsp                                 = 75
	GetAllLineupDataCsReq                              = 731
	ArchiveWolfBroGameCsReq                            = 6583
	GetBagCsReq                                        = 598
	TakeOfferingRewardScRsp                            = 6935
	SyncApplyFriendScNotify                            = 2928
	WolfBroGamePickupBulletCsReq                       = 6522
	EnterAetherDivideSceneCsReq                        = 4898
	ExpUpEquipmentCsReq                                = 578
	GetChessRogueStoryAeonTalkInfoCsReq                = 5476
	TravelBrochureSetPageDescStatusCsReq               = 6445
	ExpUpRelicScRsp                                    = 505
	RogueGetGambleInfoCsReq                            = 5617
	GetPrivateChatHistoryScRsp                         = 3977
	WolfBroGameActivateBulletScRsp                     = 6545
	DelMailCsReq                                       = 879
	WolfBroGameExplodeMonsterCsReq                     = 6520
	WolfBroGameDataChangeScNotify                      = 6578
	RogueWorkbenchSelectFuncScRsp                      = 5641
	SwordTrainingActionTurnSettleScNotify              = 7469
	ChessRogueQueryBpCsReq                             = 5509
	StartRogueScRsp                                    = 1842
	ClockParkHandleWaitOperationCsReq                  = 7214
	EvolveBuildUnlockInfoNotify                        = 7103
	GetRogueShopBuffInfoScRsp                          = 5633
	AlleyOrderChangedScNotify                          = 4750
	SyncRogueHandbookDataUpdateScNotify                = 5665
	HeliobusSnsCommentScRsp                            = 5838
	MonopolyGuessChooseCsReq                           = 7057
	GetActivityScheduleConfigScRsp                     = 2677
	SharePunkLordMonsterScRsp                          = 3277
	HeliobusChallengeUpdateScNotify                    = 5811
	SetIsDisplayAvatarInfoCsReq                        = 2833
	SetMultipleAvatarPathsCsReq                        = 81
	DrinkMakerDayEndScNotify                           = 6986
	GetTutorialScRsp                                   = 1671
	BuyRogueShopMiracleScRsp                           = 5628
	GetTutorialGuideScRsp                              = 1642
	SetClientRaidTargetCountScRsp                      = 2222
	StartTimedFarmElementScRsp                         = 1480
	UnlockTutorialGuideScRsp                           = 1612
	SettleTrackPhotoStageScRsp                         = 7553
	FinishQuestCsReq                                   = 938
	UnlockTutorialScRsp                                = 1677
	GetMultipleDropInfoScRsp                           = 4671
	MarkItemCsReq                                      = 563
	GetTutorialCsReq                                   = 1698
	OpenTreasureDungeonGridCsReq                       = 4466
	GetAlleyInfoScRsp                                  = 4771
	GetGachaCeilingCsReq                               = 1979
	LogisticsDetonateStarSkiffCsReq                    = 4736
	FinishCurTurnScRsp                                 = 4356
	UseTreasureDungeonItemScRsp                        = 4460
	HeliobusEnterBattleCsReq                           = 5831
	GetAuthkeyCsReq                                    = 5
	TakeFightActivityRewardScRsp                       = 3633
	GetDailyActiveInfoCsReq                            = 3383
	TakeRogueScoreRewardScRsp                          = 1860
	TreasureDungeonDataScNotify                        = 4498
	QuitBattleScNotify                                 = 128
	JoinLineupCsReq                                    = 779
	TreasureDungeonFinishScNotify                      = 4471
	QuitTreasureDungeonCsReq                           = 4452
	DrinkMakerChallengeScRsp                           = 6987
	GetExpeditionDataCsReq                             = 2598
	QuitTreasureDungeonScRsp                           = 4411
	InteractTreasureDungeonGridCsReq                   = 4450
	SetPlayerInfoCsReq                                 = 57
	RogueMagicLeaveCsReq                               = 7779
	GetTreasureDungeonActivityDataCsReq                = 4478
	RestoreWolfBroGameArchiveScRsp                     = 6577
	EnterTreasureDungeonCsReq                          = 4489
	TakePictureCsReq                                   = 4179
	FightTreasureDungeonMonsterCsReq                   = 4445
	TakePictureScRsp                                   = 4177
	RestoreWolfBroGameArchiveCsReq                     = 6579
	TravelBrochureUpdatePasterPosScRsp                 = 6456
	ArchiveWolfBroGameScRsp                            = 6542
	GetMovieRacingDataCsReq                            = 4104
	MarkChatEmojiScRsp                                 = 3956
	TravelBrochureRemovePasterCsReq                    = 6428
	TravelBrochureApplyPasterCsReq                     = 6433
	GetServerPrefsDataScRsp                            = 6142
	GetMainMissionCustomValueScRsp                     = 1293
	GetTrainVisitorBehaviorCsReq                       = 3783
	RogueWorkbenchHandleFuncScRsp                      = 5626
	GetAetherDivideInfoScRsp                           = 4856
	UpgradeAreaStatCsReq                               = 4366
	TravelBrochurePageResetCsReq                       = 6450
	DeployRotaterCsReq                                 = 6879
	RelicRecommendScRsp                                = 2437
	TravelBrochureSetCustomValueCsReq                  = 6466
	ContentPackageSyncDataScNotify                     = 7542
	SetDisplayAvatarCsReq                              = 2879
	MonopolyClickCellScRsp                             = 7068
	FinishSectionIdCsReq                               = 2733
	EvolveBuildStartLevelScRsp                         = 7102
	GameplayCounterCountDownScRsp                      = 1499
	SwordTrainingStoryBattleCsReq                      = 7476
	PlayerReturnStartScNotify                          = 4598
	TravelBrochureSetPageDescStatusScRsp               = 6420
	TravelBrochureGetPasterScNotify                    = 6489
	ChessRogueUpdateUnlockLevelScNotify                = 5438
	SetForbidOtherApplyFriendCsReq                     = 2940
	GetHeartDialInfoScRsp                              = 6371
	GetTrainVisitorRegisterScRsp                       = 3712
	TakeTalkRewardCsReq                                = 2183
	ChessRogueCellUpdateNotify                         = 5405
	TrainVisitorBehaviorFinishScRsp                    = 3771
	GetTrainVisitorRegisterCsReq                       = 3733
	GetShareDataScRsp                                  = 4142
	MultiplayerFightGameStateCsReq                     = 1098
	GetKilledPunkLordMonsterDataScRsp                  = 3210
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3728
	GetTrainVisitorBehaviorScRsp                       = 3742
	SceneEntityTeleportScRsp                           = 1490
	SetLineupNameCsReq                                 = 720
	TelevisionActivityBattleEndScNotify                = 6975
	ActivityRaidPlacingGameCsReq                       = 4757
	ReplaceLineupScRsp                                 = 711
	MusicRhythmSaveSongConfigDataCsReq                 = 7584
	RogueTournEnterCsReq                               = 6066
	LobbySyncInfoScNotify                              = 7364
	RogueTournGetMiscRealTimeDataScRsp                 = 6042
	TakePunkLordPointRewardCsReq                       = 3289
	GetNpcTakenRewardCsReq                             = 2198
	GetChessRogueBuffEnhanceInfoScRsp                  = 5552
	SyncRogueGetItemScNotify                           = 1826
	RogueTournResetPermanentTalentCsReq                = 6084
	FinishPerformSectionIdCsReq                        = 2728
	GetTrackPhotoActivityDataCsReq                     = 7552
	DeleteBlacklistScRsp                               = 3000
	MonopolyEventSelectFriendScRsp                     = 7084
	UpdateGunPlayDataScRsp                             = 4106
	GetTrackPhotoActivityDataScRsp                     = 7558
	LeaveAetherDivideSceneCsReq                        = 4883
	StartTrackPhotoStageScRsp                          = 7554
	UnlockTutorialCsReq                                = 1679
	LobbyInviteScNotify                                = 7377
	SettleTrackPhotoStageCsReq                         = 7557
	MonopolyGiveUpCurContentCsReq                      = 7010
	FinishItemIdCsReq                                  = 2779
	ChessRogueQuestFinishNotify                        = 5451
	PlayerGetTokenCsReq                                = 79
	MonopolySocialEventEffectScNotify                  = 7032
	CommonRogueQueryScRsp                              = 5676
	QuitTrackPhotoStageScRsp                           = 7556
	RecallPetCsReq                                     = 7644
	TextJoinSaveScRsp                                  = 3871
	UpdateRogueAdventureRoomScoreScRsp                 = 5653
	AcceptMissionEventCsReq                            = 1220
	TextJoinSaveCsReq                                  = 3898
	RogueTournEnterScRsp                               = 6059
	GetReplayTokenScRsp                                = 3571
	TextJoinBatchSaveCsReq                             = 3879
	EnterFantasticStoryActivityStageScRsp              = 4979
	RotateMapCsReq                                     = 6833
	GetCurBattleInfoCsReq                              = 179
	RogueTournEnablePermanentTalentScRsp               = 6026
	SpaceZooBornCsReq                                  = 6783
	RefreshAlleyOrderCsReq                             = 4745
	GetTelevisionActivityDataCsReq                     = 6974
	FinishCurTurnCsReq                                 = 4378
	GetQuestDataScRsp                                  = 971
	RogueMagicReviveCostUpdateScNotify                 = 7750
	FinishFirstTalkByPerformanceNpcCsReq               = 2189
	GetPetDataCsReq                                    = 7629
	PrepareRogueAdventureRoomScRsp                     = 5683
	SelectPhoneThemeCsReq                              = 5177
	CancelMarkItemNotify                               = 536
	StartCocoonStageCsReq                              = 1455
	SelectInclinationTextScRsp                         = 2138
	GetFriendLoginInfoCsReq                            = 2919
	RogueMagicScepterDressInUnitScRsp                  = 7800
	GetPhoneDataScRsp                                  = 5171
	GetSaveRaidCsReq                                   = 2266
	GetRogueCommonDialogueDataCsReq                    = 5649
	SetTurnFoodSwitchCsReq                             = 565
	SetRogueCollectionScRsp                            = 5648
	UpdateRotaterScNotify                              = 6831
	ServerAnnounceNotify                               = 85
	SelectInclinationTextCsReq                         = 2128
	UnlockPhoneThemeScNotify                           = 5112
	ChessRogueSkipTeachingLevelCsReq                   = 5562
	RogueArcadeLeaveScRsp                              = 7652
	GetMovieRacingDataScRsp                            = 4160
	FinishFirstTalkByPerformanceNpcScRsp               = 2122
	MonopolyGetDailyInitItemCsReq                      = 7046
	SwordTrainingExamResultConfirmCsReq                = 7484
	ChessRogueQueryBpScRsp                             = 5510
	SwordTrainingDialogueSelectOptionCsReq             = 7482
	MonopolyEventSelectFriendCsReq                     = 7003
	SetRedPointStatusScNotify                          = 47
	GetRndOptionScRsp                                  = 3471
	SwordTrainingResumeGameScRsp                       = 7495
	CancelCacheNotifyCsReq                             = 4128
	ComposeItemCsReq                                   = 589
	RelicAvatarRecommendCsReq                          = 2442
	ChessRogueUpdateAllowedSelectCellScNotify          = 5416
	SwordTrainingStoryConfirmScRsp                     = 7462
	SwordTrainingLearnSkillCsReq                       = 7477
	SwordTrainingGiveUpGameCsReq                       = 7456
	RogueMagicGetMiscRealTimeDataCsReq                 = 7765
	MusicRhythmFinishLevelCsReq                        = 7578
	SyncClientResVersionScRsp                          = 112
	SwordTrainingDialogueSelectOptionScRsp             = 7463
	GetSwordTrainingDataCsReq                          = 7487
	SwordTrainingResumeGameCsReq                       = 7500
	GetSceneMapInfoCsReq                               = 1426
	GetAetherDivideInfoCsReq                           = 4878
	GetMonopolyDailyReportScRsp                        = 7037
	SwordTrainingLearnSkillScRsp                       = 7455
	TakeActivityExpeditionRewardCsReq                  = 2522
	RogueMagicLeaveScRsp                               = 7777
	SetMissionEventProgressCsReq                       = 1211
	SwordTrainingSetSkillTraceCsReq                    = 7481
	TeleportToMissionResetPointScRsp                   = 1206
	SwordTrainingGiveUpGameScRsp                       = 7466
	RevcMsgScNotify                                    = 3983
	RaidKickByServerScNotify                           = 2231
	MuseumRandomEventQueryCsReq                        = 4331
	DressRelicAvatarScRsp                              = 345
	MonopolyConditionUpdateScNotify                    = 7021
	SwordTrainingGameSettleScNotify                    = 7453
	GetRogueInitialScoreCsReq                          = 1825
	SwordTrainingSetSkillTraceScRsp                    = 7490
	SwordTrainingMarkEndingViewedScRsp                 = 7485
	SwordTrainingDailyPhaseConfirmScRsp                = 7474
	EnterSwordTrainingExamScRsp                        = 7498
	EnterChessRogueAeonRoomScRsp                       = 5401
	GetSummonActivityDataScRsp                         = 7568
	GetJukeboxDataCsReq                                = 3198
	BoxingClubChallengeUpdateScNotify                  = 4238
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5660
	GetSummonActivityDataCsReq                         = 7562
	GetSpringRecoverDataCsReq                          = 1453
	FightMatch3ChatScRsp                               = 30138
	StaminaInfoScNotify                                = 48
	SelectRogueCommonDialogueOptionScRsp               = 5632
	SummonActivityBattleEndScNotify                    = 7570
	SaveLogisticsScRsp                                 = 4762
	OpenRogueChestCsReq                                = 1867
	EnterAetherDivideSceneScRsp                        = 4871
	EnterStrongChallengeActivityStageCsReq             = 6683
	StoryLineTrialAvatarChangeScNotify                 = 6233
	StoryLineInfoScNotify                              = 6283
	GetMaterialSubmitActivityDataCsReq                 = 2645
	RogueTournReEnterRogueCocoonStageCsReq             = 6092
	GmTalkCsReq                                        = 38
	SetGroupCustomSaveDataScRsp                        = 1403
	GetRogueTalentInfoScRsp                            = 1851
	MatchThreeGetDataScRsp                             = 7437
	AddBlacklistScRsp                                  = 2945
	LockRelicCsReq                                     = 545
	MissionEventRewardScNotify                         = 1245
	RogueMagicAutoDressInMagicUnitChangeScNotify       = 7792
	ScenePlaneEventScNotify                            = 1447
	StartStarFightLevelScRsp                           = 7163
	GetStarFightDataScRsp                              = 7168
	StarFightDataChangeNotify                          = 7170
	GetMissionStatusCsReq                              = 1231
	ChessRogueEnterCellScRsp                           = 5571
	LobbyModifyPlayerInfoScRsp                         = 7382
	StartStarFightLevelCsReq                           = 7167
	GetStarFightDataCsReq                              = 7162
	RogueNpcDisappearCsReq                             = 5656
	GetMissionEventDataScRsp                           = 1205
	HeliobusSnsPostCsReq                               = 5879
	RogueTournGetCurRogueCocoonInfoScRsp               = 6036
	AlleyShipmentEventEffectsScNotify                  = 4723
	EnterAdventureCsReq                                = 1398
	SpaceZooOpCatteryCsReq                             = 6733
	SetSignatureScRsp                                  = 2878
	MonopolyGetRegionProgressCsReq                     = 7059
	SpaceZooDataScRsp                                  = 6771
	SpaceZooDataCsReq                                  = 6798
	MonopolyGameRaiseRatioScRsp                        = 7016
	ChessRogueNousDiceSurfaceUnlockNotify              = 5569
	SyncChessRogueNousValueScNotify                    = 5589
	SpaceZooTakeScRsp                                  = 6766
	GetRogueBuffEnhanceInfoScRsp                       = 1811
	FightMatch3ChatCsReq                               = 30128
	SpringRecoverSingleAvatarScRsp                     = 1476
	SpaceZooExchangeItemCsReq                          = 6756
	LobbyKickOutCsReq                                  = 7363
	TelevisionActivityDataChangeScNotify               = 6973
	BuyGoodsScRsp                                      = 1542
	TakeCityShopRewardScRsp                            = 1577
	GetShopListCsReq                                   = 1598
	RogueMagicUnitComposeScRsp                         = 7785
	GetShopListScRsp                                   = 1571
	TakeCityShopRewardCsReq                            = 1579
	GetServerPrefsDataCsReq                            = 6183
	UpdateServerPrefsDataCsReq                         = 6179
	GiveUpBoxingClubChallengeCsReq                     = 4233
	EnterRogueScRsp                                    = 1877
	UnlockAvatarPathScRsp                              = 51
	UpdateServerPrefsDataScRsp                         = 6177
	DestroyItemScRsp                                   = 514
	CancelCacheNotifyScRsp                             = 4138
	SyncTaskCsReq                                      = 1277
	EnterSectionCsReq                                  = 1462
	GetRogueEndlessActivityDataScRsp                   = 6008
	MonopolyAcceptQuizCsReq                            = 7036
	EnterSceneScRsp                                    = 1472
	SpaceZooExchangeItemScRsp                          = 6789
	SceneCastSkillCostMpCsReq                          = 1422
	TravelBrochureUpdatePasterPosCsReq                 = 6478
	SetClientRaidTargetCountCsReq                      = 2289
	InteractPropScRsp                                  = 1442
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3738
	TakeMultipleExpeditionRewardCsReq                  = 2520
	ReEnterLastElementStageScRsp                       = 1429
	UpdateMechanismBarScNotify                         = 1439
	PrivateMsgOfflineUsersScNotify                     = 3942
	DeleteSummonUnitCsReq                              = 1407
	RogueTournWeekChallengeUpdateScNotify              = 6013
	SpringRefreshCsReq                                 = 1420
	GetCurSceneInfoScRsp                               = 1412
	GetRogueAeonInfoScRsp                              = 1868
	MuseumDispatchFinishedScNotify                     = 4311
	EnterSceneCsReq                                    = 1473
	GetMonopolyFriendRankingListScRsp                  = 7069
	EnterChallengeNextPhaseCsReq                       = 1711
	EntityBindPropScRsp                                = 1465
	RogueTournReviveAvatarCsReq                        = 6041
	MultiplayerFightGameStateScRsp                     = 1071
	RogueMagicSettleCsReq                              = 7733
	MusicRhythmStartLevelScRsp                         = 7571
	ChessRogueNousDiceUpdateNotify                     = 5537
	StartTimedCocoonStageCsReq                         = 1413
	FightMatch3StartCountDownScNotify                  = 30183
	FinishCosumeItemMissionScRsp                       = 1222
	ChessRogueMoveCellNotify                           = 5533
	SceneEntityMoveCsReq                               = 1498
	SceneCastSkillMpUpdateScNotify                     = 1405
	GetStoryLineInfoCsReq                              = 6298
	GameplayCounterCountDownCsReq                      = 1409
	GameplayCounterRecoverCsReq                        = 1482
	ExpeditionDataChangeScNotify                       = 2528
	EvolveBuildStartLevelCsReq                         = 7142
	SpringRecoverCsReq                                 = 1454
	QuitBattleCsReq                                    = 183
	FinishCosumeItemMissionCsReq                       = 1289
	SetClientPausedCsReq                               = 1457
	ComposeSelectedRelicCsReq                          = 511
	EnterSectionScRsp                                  = 1406
	GetAvatarDataCsReq                                 = 398
	HeliobusUpgradeLevelScRsp                          = 5822
	SpaceZooMutateCsReq                                = 6779
	GroupStateChangeCsReq                              = 1417
	GetMarkItemListCsReq                               = 588
	GetFriendAssistListCsReq                           = 2935
	SceneEntityTeleportCsReq                           = 1401
	GetChallengeCsReq                                  = 1798
	SavePointsInfoNotify                               = 1463
	StartTimedCocoonStageScRsp                         = 1491
	GroupStateChangeScRsp                              = 1437
	HeliobusUnlockSkillScNotify                        = 5866
	MuseumInfoChangedScNotify                          = 4345
	ExchangeStaminaScRsp                               = 66
	SwordTrainingStoryBattleScRsp                      = 7491
	SetSpringRecoverConfigScRsp                        = 1487
	SpringRecoverSingleAvatarCsReq                     = 1467
	RogueMagicStartScRsp                               = 7771
	GetChessRogueStoryInfoScRsp                        = 5444
	MonopolyTakeRaffleTicketRewardCsReq                = 7047
	InteractPropCsReq                                  = 1483
	ApplyFriendCsReq                                   = 2933
	SceneEntityMoveScRsp                               = 1471
	MusicRhythmUnlockSongSfxScNotify                   = 7580
	SetClientPausedScRsp                               = 1425
	PlayerLogoutCsReq                                  = 83
	SetCurInteractEntityCsReq                          = 1416
	GetExhibitScNotify                                 = 4338
	GetSecretKeyInfoScRsp                              = 1
	RefreshTriggerByClientScNotify                     = 1481
	RecoverAllLineupCsReq                              = 1488
	DeactivateFarmElementScRsp                         = 1492
	MarkRelicFilterPlanCsReq                           = 539
	GetQuestRecordCsReq                                = 933
	GetMailScRsp                                       = 871
	ActivateFarmElementScRsp                           = 1495
	PlayerReturnTakePointRewardCsReq                   = 4579
	SceneEnterStageCsReq                               = 1452
	GameplayCounterRecoverScRsp                        = 1418
	ReturnLastTownCsReq                                = 1404
	SetSpringRecoverConfigCsReq                        = 1435
	EntityBindPropCsReq                                = 1402
	ActivateFarmElementCsReq                           = 1440
	ChessRogueQueryCsReq                               = 5418
	GetCurAssistScRsp                                  = 2993
	MuseumRandomEventSelectCsReq                       = 4360
	ReturnLastTownScRsp                                = 1460
	SyncServerSceneChangeNotify                        = 1427
	SceneGroupRefreshScNotify                          = 1430
	HealPoolInfoNotify                                 = 1458
	ComposeLimitNumUpdateNotify                        = 585
	UpdateFloorSavedValueNotify                        = 1496
	SetGroupCustomSaveDataCsReq                        = 1449
	RogueTournLeaveCsReq                               = 6043
	DressRelicAvatarCsReq                              = 305
	ChessRoguePickAvatarCsReq                          = 5424
	PlayerReturnSignCsReq                              = 4571
	TakeOffEquipmentCsReq                              = 378
	FinishRogueCommonDialogueCsReq                     = 5686
	SceneCastSkillScRsp                                = 1477
	ChessRogueLeaveCsReq                               = 5484
	SceneEntityMoveScNotify                            = 1478
	GroupStateChangeScNotify                           = 1475
	EnteredSceneChangeScNotify                         = 1441
	GetRogueAdventureRoomInfoScRsp                     = 5666
	GetUnlockTeleportScRsp                             = 1448
	MonopolyActionResultScNotify                       = 7083
	LeaveTrialActivityScRsp                            = 2632
	GetNpcStatusScRsp                                  = 2742
	TakeRollShopRewardCsReq                            = 6908
	ChessRogueSelectCellCsReq                          = 5578
	GetRollShopInfoCsReq                               = 6914
	GetRollShopInfoScRsp                               = 6917
	SetAetherDivideLineUpCsReq                         = 4889
	ClientObjDownloadDataScNotify                      = 9
	TextJoinQueryCsReq                                 = 3883
	DifficultyAdjustmentUpdateDataCsReq                = 4116
	MonopolyLikeScNotify                               = 7076
	VirtualLineupTrialAvatarChangeScNotify             = 800
	RestartChallengePhaseScRsp                         = 1752
	RogueTournReviveCostUpdateScNotify                 = 6077
	PrestigeLevelUpScRsp                               = 4760
	RogueTournGetSettleInfoCsReq                       = 6075
	SecurityReportCsReq                                = 4178
	ExpUpEquipmentScRsp                                = 556
	RogueTournHandBookNotify                           = 6070
	RogueTournExpNotify                                = 6088
	RogueTournAreaUpdateScNotify                       = 6020
	SetLineupNameScRsp                                 = 750
	ComposeLimitNumCompleteNotify                      = 523
	RogueTournGetCurRogueCocoonInfoCsReq               = 6055
	RaidCollectionDataScRsp                            = 6957
	MarkReadMailCsReq                                  = 883
	RogueTournTakeExpRewardScRsp                       = 6052
	DoGachaCsReq                                       = 1983
	HandleFriendScRsp                                  = 2978
	UseItemCsReq                                       = 533
	RogueTournGetAllArchiveCsReq                       = 6032
	MissionAcceptScNotify                              = 1263
	RogueTournStartScRsp                               = 6046
	EnhanceCommonRogueBuffScRsp                        = 5611
	RogueTournReEnterRogueCocoonStageScRsp             = 6098
	RogueTournSettleCsReq                              = 6094
	MultiplayerFightGiveUpScRsp                        = 1077
	PickRogueAvatarCsReq                               = 1805
	GetChessRogueStoryInfoCsReq                        = 5464
	GetMaterialSubmitActivityDataScRsp                 = 2620
	UpdatePsnSettingsInfoScRsp                         = 82
	LobbyBeginCsReq                                    = 7394
	NewMailScNotify                                    = 828
	SubmitEmotionItemScRsp                             = 6377
	GetChessRogueStoryAeonTalkInfoScRsp                = 5450
	LeaveChallengeScRsp                                = 1777
	ChooseBoxingClubResonanceScRsp                     = 4256
	FinishItemIdScRsp                                  = 2777
	RogueTournLeaveRogueCocoonSceneScRsp               = 6074
	SpaceZooCatUpdateNotify                            = 6778
	EquipAetherDividePassiveSkillScRsp                 = 4805
	FightKickOutScNotify                               = 30042
	StartChallengeCsReq                                = 1783
	RogueTournDeleteArchiveCsReq                       = 6060
	TrainVisitorBehaviorFinishCsReq                    = 3798
	RogueTournClearArchiveNameScNotify                 = 6011
	ChessRogueStartScRsp                               = 5467
	RogueTournEnablePermanentTalentCsReq               = 6048
	EnterTreasureDungeonScRsp                          = 4422
	MatchThreeSetBirdPosCsReq                          = 7410
	TakeAllRewardScRsp                                 = 3028
	GetRelicFilterPlanCsReq                            = 564
	EnableRogueTalentCsReq                             = 1881
	FinishPlotScRsp                                    = 1171
	RogueWorkbenchGetInfoScRsp                         = 5674
	HeliobusSelectSkillScRsp                           = 5845
	RogueTournLevelInfoUpdateScNotify                  = 6015
	UpdatePlayerSettingCsReq                           = 24
	FightMatch3ChatScNotify                            = 30178
	GetFriendDevelopmentInfoCsReq                      = 2949
	RogueTournSettleScRsp                              = 6030
	RogueTournTakeExpRewardCsReq                       = 6012
	RogueTournGetPermanentTalentInfoScRsp              = 6016
	WolfBroGamePickupBulletScRsp                       = 6566
	SetMissionEventProgressScRsp                       = 1210
	SyncRogueStatusScNotify                            = 1880
	RogueTournGetMiscRealTimeDataCsReq                 = 6097
	InteractChargerScRsp                               = 6842
	FightLeaveScNotify                                 = 30083
	SyncRogueAeonScNotify                              = 1874
	RogueModifierAddNotify                             = 5383
	RogueModifierDelNotify                             = 5328
	CityShopInfoScNotify                               = 1533
	MonopolyGetRegionProgressScRsp                     = 7048
	MarkChatEmojiCsReq                                 = 3978
	DelMailScRsp                                       = 877
	FeverTimeActivityBattleEndScNotify                 = 7157
	RogueMagicUnitComposeCsReq                         = 7723
	RogueMagicScepterTakeOffUnitCsReq                  = 7762
	RogueMagicReviveAvatarScRsp                        = 7704
	RogueMagicAutoDressInUnitScRsp                     = 7715
	FightEnterCsReq                                    = 30098
	RogueMagicEnterLayerScRsp                          = 7756
	RogueMagicEnterLayerCsReq                          = 7778
	RogueMagicLevelInfoUpdateScNotify                  = 7789
	TeleportToMissionResetPointCsReq                   = 1262
	RogueMagicSettleScRsp                              = 7712
	GetBasicInfoCsReq                                  = 53
	EnterMapRotationRegionScRsp                        = 6871
	RogueMagicStartCsReq                               = 7798
	CommonRogueQueryCsReq                              = 5667
	PlayerLoginScRsp                                   = 71
	SwordTrainingStartGameCsReq                        = 7493
	ChessRogueFinishCurRoomNotify                      = 5409
	RogueMagicEnableTalentScRsp                        = 7755
	RogueMagicQueryScRsp                               = 7752
	MonopolyClickMbtiReportCsReq                       = 7051
	EvolveBuildStartStageCsReq                         = 7144
	SyncHandleFriendScNotify                           = 2956
	RogueMagicBattleFailSettleInfoScNotify             = 7720
	TakeRogueEndlessActivityPointRewardScRsp           = 6005
	GetRogueEndlessActivityDataCsReq                   = 6002
	MonopolyBuyGoodsScRsp                              = 7060
	GetRogueAeonInfoCsReq                              = 1875
	TakeRogueEndlessActivityPointRewardCsReq           = 6004
	SecurityReportScRsp                                = 4156
	GetEnteredSceneScRsp                               = 1446
	FinishRogueCommonDialogueScRsp                     = 5629
	UpdateRogueAdventureRoomScoreCsReq                 = 5695
	ChessRogueQuitScRsp                                = 5597
	PunkLordDataChangeNotify                           = 3216
	StartChallengeScRsp                                = 1742
	UseItemScRsp                                       = 512
	CommonRogueComponentUpdateScNotify                 = 5673
	SelectRogueCommonDialogueOptionCsReq               = 5684
	CancelExpeditionScRsp                              = 2577
	RogueWorkbenchGetInfoCsReq                         = 5672
	GetRogueAdventureRoomInfoCsReq                     = 5622
	RogueDebugMessageScNotify                          = 5621
	SetGameplayBirthdayScRsp                           = 54
	GetWaypointCsReq                                   = 498
	GetSaveLogisticsMapScRsp                           = 4716
	GetFightFestDataCsReq                              = 7279
	HeliobusEnterBattleScRsp                           = 5804
	ClockParkStartScriptScRsp                          = 7232
	RogueTournEnterRogueCocoonSceneCsReq               = 6035
	LobbyCreateScRsp                                   = 7387
	GetRogueExhibitionScRsp                            = 5696
	GetFirstTalkByPerformanceNpcScRsp                  = 2156
	RogueMagicEnableTalentCsReq                        = 7763
	MatchThreeGetDataCsReq                             = 7429
	RogueDoGambleScRsp                                 = 5668
	GetRogueHandbookDataScRsp                          = 5602
	FinishQuestScRsp                                   = 978
	UpdateRedDotDataCsReq                              = 5983
	TakeMultipleActivityExpeditionRewardCsReq          = 2531
	GetLevelRewardTakenListCsReq                       = 60
	TakeQuestRewardCsReq                               = 983
	ChessRogueNousEnableRogueTalentCsReq               = 5447
	TakeExpeditionRewardCsReq                          = 2533
	GetRogueCommonDialogueDataScRsp                    = 5603
	GetRogueCollectionCsReq                            = 5697
	TakeRogueMiracleHandbookRewardCsReq                = 5657
	GetArchiveDataScRsp                                = 2371
	PromoteEquipmentScRsp                              = 542
	BuyRogueShopBuffScRsp                              = 5678
	EnhanceCommonRogueBuffCsReq                        = 5652
	RogueArcadeLeaveCsReq                              = 7692
	FinishChessRogueSubStoryScRsp                      = 5587
	RecallPetScRsp                                     = 7610
	RogueWorkbenchSelectFuncCsReq                      = 5646
	ChessRogueUpdateAeonModifierValueScNotify          = 5559
	GetRogueHandbookDataCsReq                          = 5636
	DeleteSocialEventServerCacheCsReq                  = 7001
	GetRogueShopBuffInfoCsReq                          = 5677
	LogisticsDetonateStarSkiffScRsp                    = 4702
	SyncRogueCommonPendingActionScNotify               = 5640
	LogisticsScoreRewardSyncInfoScNotify               = 4765
	GetRogueCollectionScRsp                            = 5627
	TakeMultipleActivityExpeditionRewardScRsp          = 2504
	MatchResultScNotify                                = 7344
	OpenRogueChestScRsp                                = 1876
	GetStageLineupCsReq                                = 798
	DailyActiveInfoNotify                              = 3379
	ChessRogueGiveUpCsReq                              = 5461
	DifficultyAdjustmentUpdateDataScRsp                = 4114
	AlleyTakeEventRewardScRsp                          = 4755
	SetRogueCollectionCsReq                            = 5659
	TakeRogueEventHandbookRewardScRsp                  = 5619
	SyncRogueCommonDialogueDataScNotify                = 5601
	DelSaveRaidScNotify                                = 2250
	SummonPetCsReq                                     = 7642
	SpaceZooTakeCsReq                                  = 6722
	RogueDoGambleCsReq                                 = 5675
	SyncRogueAdventureRoomInfoScNotify                 = 5698
	SyncRogueCommonVirtualItemInfoScNotify             = 5664
	EnterTelevisionActivityStageCsReq                  = 6970
	MonopolyCellUpdateNotify                           = 7042
	SetCurWaypointScRsp                                = 442
	RogueArcadeRestartCsReq                            = 7694
	DressAvatarSkinCsReq                               = 360
	RogueArcadeStartCsReq                              = 7679
	GetFriendChallengeDetailScRsp                      = 2967
	ReplaceLineupCsReq                                 = 752
	RogueArcadeRestartScRsp                            = 7660
	GetChallengeRecommendLineupListCsReq               = 2444
	RogueArcadeStartScRsp                              = 7687
	LeaveRogueCsReq                                    = 1833
	ChessRogueSelectCellScRsp                          = 5497
	TrialBackGroundMusicCsReq                          = 3133
	RogueTournGetArchiveRepositoryCsReq                = 6087
	GetMonopolyMbtiReportRewardCsReq                   = 7082
	ChallengeLineupNotify                              = 1756
	GetAssistHistoryScRsp                              = 2955
	EvolveBuildReRandomStageCsReq                      = 7149
	FinishAeonDialogueGroupCsReq                       = 1846
	SyncRogueMapRoomScNotify                           = 1819
	MatchBoxingClubOpponentCsReq                       = 4283
	HeliobusStartRaidScRsp                             = 5852
	SyncRoguePickAvatarInfoScNotify                    = 1870
	GetRogueScoreRewardInfoScRsp                       = 1899
	SceneCastSkillCsReq                                = 1479
	EvolveBuildQueryInfoScRsp                          = 7137
	ReviveRogueAvatarScRsp                             = 1831
	GetBagScRsp                                        = 571
	TakeOffAvatarSkinScRsp                             = 310
	ExchangeRogueRewardKeyCsReq                        = 1839
	RogueModifierStageStartNotify                      = 5338
	SwordTrainingRestoreGameCsReq                      = 7486
	LobbyInviteScRsp                                   = 7396
	HeartDialTraceScriptScRsp                          = 6378
	SyncRogueFinishScNotify                            = 1866
	ChessRogueEnterScRsp                               = 5551
	QuitTrackPhotoStageCsReq                           = 7555
	GetRogueBuffEnhanceInfoCsReq                       = 1852
	GetRogueScoreRewardInfoCsReq                       = 1809
	LeaveMapRotationRegionScRsp                        = 6838
	ExchangeRogueRewardKeyScRsp                        = 1849
	EnterRogueCsReq                                    = 1879
	GetFriendListInfoScRsp                             = 2971
	EnterRogueMapRoomScRsp                             = 1857
	ExchangeHcoinCsReq                                 = 560
	ReviveRogueAvatarCsReq                             = 1850
	LeaveRaidScRsp                                     = 2242
	SubmitMaterialSubmitActivityMaterialCsReq          = 2650
	FightSessionStopScNotify                           = 30033
	RotateMapScRsp                                     = 6812
	GetRogueTalentInfoCsReq                            = 1821
	RechargeSuccNotify                                 = 504
	HeliobusSnsLikeCsReq                               = 5833
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6006
	PlayerReturnTakePointRewardScRsp                   = 4577
	PickRogueAvatarScRsp                               = 1845
	GetRogueInfoCsReq                                  = 1898
	GetStrongChallengeActivityDataCsReq                = 6698
	EnableRogueTalentScRsp                             = 1807
	RogueArcadeGetInfoCsReq                            = 7674
	EnterFeverTimeActivityStageScRsp                   = 7158
	HeliobusLineupUpdateScNotify                       = 5810
	GetMainMissionCustomValueCsReq                     = 1288
	GetLevelRewardCsReq                                = 11
	ExchangeGachaCeilingCsReq                          = 1933
	GetWaypointScRsp                                   = 471
	EvolveBuildShopAbilityUpScRsp                      = 7105
	EnhanceRogueBuffCsReq                              = 1810
	DailyFirstMeetPamScRsp                             = 3442
	DailyFirstMeetPamCsReq                             = 3483
	GetPlayerReplayInfoScRsp                           = 3542
	GetPlayerReplayInfoCsReq                           = 3583
	TakeQuestOptionalRewardCsReq                       = 956
	EvolveBuildShopAbilityUpCsReq                      = 7127
	MonopolyScrachRaffleTicketScRsp                    = 7027
	SyncRogueReviveInfoScNotify                        = 1816
	GetChallengeRecommendLineupListScRsp               = 2410
	GetFightFestDataScRsp                              = 7287
	GetLoginActivityCsReq                              = 2698
	GetFriendBattleRecordDetailScRsp                   = 2939
	RaidCollectionDataScNotify                         = 6953
	SetNicknameCsReq                                   = 31
	GetChallengeRaidInfoCsReq                          = 2277
	LeaveRaidCsReq                                     = 2283
	TakeFightActivityRewardCsReq                       = 3677
	MonopolyScrachRaffleTicketCsReq                    = 7097
	FightMatch3OpponentDataScNotify                    = 30112
	GetRaidInfoScRsp                                   = 2256
	AetherDivideTainerInfoScNotify                     = 4823
	TakeChallengeRaidRewardCsReq                       = 2212
	GetAllSaveRaidScRsp                                = 2220
	ChallengeRaidNotify                                = 2238
	SetAssistCsReq                                     = 2916
	QuitLineupCsReq                                    = 733
	DeleteSummonUnitScRsp                              = 1494
	ActivityRaidPlacingGameScRsp                       = 4725
	AvatarExpUpScRsp                                   = 342
	ChessRoguePickAvatarScRsp                          = 5521
	GetQuestRecordScRsp                                = 912
	DressAvatarSkinScRsp                               = 352
	GetPlatformPlayerInfoScRsp                         = 2915
	MakeMissionDrinkScRsp                              = 6989
	GetAllRedDotDataCsReq                              = 5998
	CommonRogueUpdateScNotify                          = 5639
	GetChapterScRsp                                    = 477
	AlleyShipUsedCountScNotify                         = 4714
	MonopolyGetRafflePoolInfoCsReq                     = 7024
	TakeQuestOptionalRewardScRsp                       = 989
	AcceptExpeditionScRsp                              = 2542
	LockEquipmentScRsp                                 = 577
	UpgradeAreaStatScRsp                               = 4305
	StartBoxingClubBattleCsReq                         = 4279
	EnterSceneByServerScNotify                         = 1474
	BuyRogueShopMiracleCsReq                           = 5612
	RogueTournEnterRoomScRsp                           = 6067
	QuestRecordScNotify                                = 928
	GetTutorialGuideCsReq                              = 1683
	SummonPunkLordMonsterScRsp                         = 3212
	RogueTournRenameArchiveCsReq                       = 6063
	ShowNewSupplementVisitorCsReq                      = 3778
	PromoteAvatarCsReq                                 = 333
	VirtualLineupDestroyNotify                         = 760
	TakeRogueEventHandbookRewardCsReq                  = 5615
	MarkAvatarCsReq                                    = 362
	GetPunkLordDataScRsp                               = 3245
	ChessRogueCheatRollCsReq                           = 5446
	TakeKilledPunkLordMonsterScoreCsReq                = 3223
	GetPunkLordBattleRecordScRsp                       = 3288
	MonopolyConfirmRandomCsReq                         = 7050
	ChessRogueQueryScRsp                               = 5547
	GetRogueShopMiracleInfoCsReq                       = 5642
	GetPunkLordBattleRecordCsReq                       = 3214
	SetFriendMarkCsReq                                 = 2953
	ChessRogueSkipTeachingLevelScRsp                   = 5472
	GetExpeditionDataScRsp                             = 2571
	SetLanguageScRsp                                   = 23
	EvolveBuildShopAbilityDownScRsp                    = 7115
	RemoveRotaterScRsp                                 = 6850
	MonopolyLikeCsReq                                  = 7058
	SummonPunkLordMonsterCsReq                         = 3233
	UpdateGunPlayDataCsReq                             = 4162
	FinishPlotCsReq                                    = 1198
	PlayerReturnInfoQueryScRsp                         = 4538
	FinishTalkMissionScRsp                             = 1242
	RelicRecommendCsReq                                = 2429
	TakeChallengeRewardScRsp                           = 1705
	StartMatchCsReq                                    = 7329
	GetUnlockTeleportCsReq                             = 1459
	MakeMissionDrinkCsReq                              = 6992
	MonopolyContentUpdateScNotify                      = 7023
	DeleteBlacklistCsReq                               = 2910
	MonopolyGiveUpCurContentScRsp                      = 7100
	GetTreasureDungeonActivityDataScRsp                = 4456
	SyncChessRogueNousSubStoryScNotify                 = 5452
	AcceptMainMissionScRsp                             = 1214
	SearchPlayerScRsp                                  = 2906
	FinishTalkMissionCsReq                             = 1283
	GetCurSceneInfoCsReq                               = 1433
	SceneEnterStageScRsp                               = 1411
	PunkLordMonsterKilledNotify                        = 3206
	EvolveBuildGiveupScRsp                             = 7132
	ChangeLineupLeaderScRsp                            = 766
	UnlockHeadIconScNotify                             = 2828
	SetDisplayAvatarScRsp                              = 2877
	AceAntiCheaterScRsp                                = 58
	ContentPackageGetDataScRsp                         = 7537
	TakeTrialActivityRewardCsReq                       = 2676
	SetHeadIconCsReq                                   = 2883
	SetHeadIconScRsp                                   = 2842
	SharePunkLordMonsterCsReq                          = 3279
	GetAlleyInfoCsReq                                  = 4798
	GetPunkLordDataCsReq                               = 3205
	UnlockAvatarPathCsReq                              = 21
	MonopolyGameSettleScNotify                         = 7014
	SelectChatBubbleCsReq                              = 5183
	GetPlayerBoardDataCsReq                            = 2898
	ClientObjUploadScRsp                               = 43
	StartCocoonStageScRsp                              = 1436
	SetGenderScRsp                                     = 65
	GetGameStateServiceConfigScRsp                     = 34
	AetherDivideSpiritExpUpCsReq                       = 4852
	EvolveBuildReRandomStageScRsp                      = 7146
	RestartChallengePhaseCsReq                         = 1760
	AlleyGuaranteedFundsScRsp                          = 4793
	UpdatePsnSettingsInfoCsReq                         = 30
	LockEquipmentCsReq                                 = 579
	TakePrestigeRewardCsReq                            = 4778
	SyncTurnFoodNotify                                 = 502
	ReserveStaminaExchangeScRsp                        = 59
	ChessRogueGiveUpRollScRsp                          = 5505
	GetAllServerPrefsDataCsReq                         = 6198
	BoxingClubRewardScNotify                           = 4228
	EvolveBuildStartStageScRsp                         = 7110
	GetSecretKeyInfoCsReq                              = 29
	RogueTournBattleFailSettleInfoScNotify             = 6068
	PlayerHeartBeatCsReq                               = 39
	FinishEmotionDialoguePerformanceCsReq              = 6333
	SetLanguageCsReq                                   = 6
	EnterAdventureScRsp                                = 1371
	GetChatFriendHistoryScRsp                          = 3912
	QueryProductInfoCsReq                              = 19
	TakeRogueScoreRewardCsReq                          = 1804
	GetAuthkeyScRsp                                    = 45
	PlayerHeartBeatScRsp                               = 49
	GetPlayerBoardDataScRsp                            = 2871
	SyncAddBlacklistScNotify                           = 2920
	QueryProductInfoScRsp                              = 92
	FightMatch3SwapCsReq                               = 30177
	GateServerScNotify                                 = 3
	RogueTournStartCsReq                               = 6072
	UpdateFeatureSwitchScNotify                        = 95
	TakeTalkRewardScRsp                                = 2142
	GetChessRogueNousStoryInfoScRsp                    = 5493
	GmTalkScRsp                                        = 78
	EvolveBuildShopAbilityResetCsReq                   = 7106
	EnterTrialActivityStageCsReq                       = 2658
	SendMsgScRsp                                       = 3971
	DeleteSocialEventServerCacheScRsp                  = 7090
	GetPlayerDetailInfoScRsp                           = 2942
	GetPunkLordMonsterDataScRsp                        = 3271
	MonopolyRollRandomCsReq                            = 7066
	GetRaidInfoCsReq                                   = 2278
	AlleyFundsScNotify                                 = 4752
	ChallengeBossPhaseSettleNotify                     = 1800
	FightHeartBeatCsReq                                = 30079
	MonopolyGameBingoFlipCardCsReq                     = 7063
	AceAntiCheaterCsReq                                = 69
	SwitchLineupIndexScRsp                             = 745
	GetOfferingInfoCsReq                               = 6934
	DailyRefreshNotify                                 = 35
	DailyFirstEnterMonopolyActivityScRsp               = 7022
	GetBoxingClubInfoScRsp                             = 4271
	ApplyFriendScRsp                                   = 2912
	RogueMagicStoryInfoUpdateScNotify                  = 7719
	EndDrinkMakerSequenceCsReq                         = 6988
	LobbyGetInfoCsReq                                  = 7355
	GetSwordTrainingDataScRsp                          = 7492
	InteractTreasureDungeonGridScRsp                   = 4431
	GetBattleCollegeDataScRsp                          = 5771
	LobbyQuitCsReq                                     = 7375
	FightMatch3TurnEndScNotify                         = 30179
	BuyNpcStuffScRsp                                   = 4342
	PlayerLogoutScRsp                                  = 42
	MatchBoxingClubOpponentScRsp                       = 4242
	SendMsgCsReq                                       = 3998
	SetIsDisplayAvatarInfoScRsp                        = 2812
	ChessRogueStartCsReq                               = 5583
	MonopolySttUpdateScNotify                          = 7030
	PlayerLoginCsReq                                   = 98
	MonthCardRewardNotify                              = 67
	AlleyShipUnlockScNotify                            = 4710
	SubmitOfferingItemCsReq                            = 6933
	GetMuseumInfoScRsp                                 = 4371
	RogueMagicGetTalentInfoCsReq                       = 7788
	RankUpAvatarScRsp                                  = 366
	BattlePassInfoNotify                               = 3098
	RogueTournQueryCsReq                               = 6014
	GetLevelRewardScRsp                                = 10
	GetVideoVersionKeyCsReq                            = 72
	GetTrialActivityDataScRsp                          = 2654
	SetGameplayBirthdayCsReq                           = 87
	AntiAddictScNotify                                 = 50
	AcceptMainMissionCsReq                             = 1216
	AetherDivideSpiritInfoScNotify                     = 4810
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5604
	TakeLoginActivityRewardCsReq                       = 2683
	ChessRogueReviveAvatarScRsp                        = 5507
	SyncRogueRewardInfoScNotify                        = 1844
	SetFriendMarkScRsp                                 = 2964
	StartPunkLordRaidCsReq                             = 3283
	UnlockTutorialGuideCsReq                           = 1633
	SetRogueExhibitionScRsp                            = 5699
	StartTrialActivityCsReq                            = 2649
	GetChatEmojiListCsReq                              = 3928
	TravelBrochureSetCustomValueScRsp                  = 6405
	MultipleDropInfoScNotify                           = 4683
	ExchangeRogueBuffWithMiracleScRsp                  = 5631
	BatchMarkChatEmojiScRsp                            = 3922
	AcceptMultipleExpeditionCsReq                      = 2505
	GetBasicInfoScRsp                                  = 64
	LobbyBeginScRsp                                    = 7360
	GetFeverTimeActivityDataCsReq                      = 7153
	EnterChallengeNextPhaseScRsp                       = 1710
	EnhanceChessRogueBuffScRsp                         = 5576
	UnlockChatBubbleScNotify                           = 5179
	BattleLogReportScRsp                               = 178
	RogueTournEnterRogueCocoonSceneScRsp               = 6025
	GetFriendRecommendListInfoCsReq                    = 2950
	AetherDivideSkillItemScNotify                      = 4885
	SelectChatBubbleScRsp                              = 5142
	TakeLoginActivityRewardScRsp                       = 2642
	GetJukeboxDataScRsp                                = 3171
	GetFirstTalkNpcCsReq                               = 2179
	TriggerVoiceScRsp                                  = 4122
	SellItemCsReq                                      = 550
	GetPetDataScRsp                                    = 7637
	HeliobusSnsPostScRsp                               = 5877
	GetMonopolyFriendRankingListCsReq                  = 7054
	CurPetChangedScNotify                              = 7624
	SyncAcceptedPamMissionNotify                       = 4083
	TakeRogueAeonLevelRewardScRsp                      = 1827
	AddEquipmentScNotify                               = 600
	DeleteRelicFilterPlanCsReq                         = 567
	AcceptedPamMissionExpireCsReq                      = 4098
	ExpUpRelicCsReq                                    = 566
	AcceptedPamMissionExpireScRsp                      = 4071
	ResetMapRotationRegionCsReq                        = 6889
	StartWolfBroGameScRsp                              = 6571
	ChessRogueUpdateReviveInfoScNotify                 = 5422
	LogisticsInfoScNotify                              = 4706
	EnterSummonActivityStageCsReq                      = 7567
	PVEBattleResultCsReq                               = 198
	TravelBrochureSelectMessageScRsp                   = 6477
	ChessRogueReviveAvatarCsReq                        = 5515
	TakeOfferingRewardCsReq                            = 6928
	OfferingInfoScNotify                               = 6932
	DeleteRelicFilterPlanScRsp                         = 576
	MusicRhythmMaxDifficultyLevelsUnlockNotify         = 7575
	BuyBpLevelCsReq                                    = 3077
	TakeChallengeRaidRewardScRsp                       = 2228
	MusicRhythmUnlockSongNotify                        = 7573
	MusicRhythmFinishLevelScRsp                        = 7587
	RemoveStuffFromAreaScRsp                           = 4312
	GetMonopolyMbtiReportRewardScRsp                   = 7018
	FightEnterScRsp                                    = 30071
	RefreshTriggerByClientCsReq                        = 1421
	ClockParkUnlockTalentCsReq                         = 7244
	TakeQuestRewardScRsp                               = 942
	GetKilledPunkLordMonsterDataCsReq                  = 3211
	DrinkMakerChallengeCsReq                           = 6983
	MuseumTargetMissionFinishNotify                    = 4400
	MuseumRandomEventSelectScRsp                       = 4352
	TakeTrialActivityRewardScRsp                       = 2639
	SwordTrainingRestoreGameScRsp                      = 7459
	GetUpdatedArchiveDataCsReq                         = 2383
	GetAllLineupDataScRsp                              = 704
	OpenTreasureDungeonGridScRsp                       = 4405
	RogueTournGetArchiveRepositoryScRsp                = 6039
	UpgradeAreaCsReq                                   = 4389
	ClockParkHandleWaitOperationScRsp                  = 7249
	TakeAllApRewardScRsp                               = 3333
	PromoteAvatarScRsp                                 = 312
	AcceptActivityExpeditionScRsp                      = 2578
	CurTrialActivityScNotify                           = 2686
	SpaceZooMutateScRsp                                = 6777
	AetherDivideFinishChallengeScNotify                = 4806
	ComposeSelectedRelicScRsp                          = 510
	GetPhoneDataCsReq                                  = 5198
	TrialActivityDataChangeScNotify                    = 2669
	FinishFirstTalkNpcCsReq                            = 2133
	EnterTrialActivityStageScRsp                       = 2667
	StartAetherDivideSceneBattleScRsp                  = 4877
	TakeMailAttachmentCsReq                            = 833
	MuseumRandomEventStartScNotify                     = 4350
	GameplayCounterUpdateScNotify                      = 1443
	EvolveBuildLeaveCsReq                              = 7113
	SwapLineupCsReq                                    = 728
	SwitchAetherDivideLineUpSlotCsReq                  = 4850
	MultipleDropInfoNotify                             = 4677
	FightFestScoreUpdateNotify                         = 7294
	DeleteFriendCsReq                                  = 2989
	GetCurChallengeCsReq                               = 1738
	StartBattleCollegeScRsp                            = 5779
	GetMultipleDropInfoCsReq                           = 4698
	SellItemScRsp                                      = 531
	MultiplayerGetFightGateCsReq                       = 1083
	UnlockedAreaMapScNotify                            = 1434
	ChessRogueGoAheadScRsp                             = 5592
	GetRecyleTimeScRsp                                 = 506
	MonopolyBuyGoodsCsReq                              = 7004
	RogueMagicEnterRoomScRsp                           = 7738
	ChooseBoxingClubResonanceCsReq                     = 4278
	CancelExpeditionCsReq                              = 2579
	MonopolyCheatDiceScRsp                             = 7006
	GetStageLineupScRsp                                = 771
	ShareScRsp                                         = 4171
	RogueTournGetSettleInfoScRsp                       = 6065
	GetFriendAssistListScRsp                           = 2987
	RebattleByClientCsNotify                           = 122
	MonopolyCheatDiceCsReq                             = 7062
	ClockParkUseBuffCsReq                              = 7215
	MonopolyGameGachaScRsp                             = 7093
	GetMbtiReportCsReq                                 = 7039
	GetFriendRecommendListInfoScRsp                    = 2931
	MonopolyDailySettleScNotify                        = 7087
	MonopolyTakePhaseRewardCsReq                       = 7009
	SwordTrainingExamResultConfirmScRsp                = 7473
	GetCurChallengeScRsp                               = 1778
	DailyTaskDataScNotify                              = 1212
	RecoverAllLineupScRsp                              = 1493
	ChessRogueUpdateLevelBaseInfoScNotify              = 5553
	RogueTournConfirmSettleCsReq                       = 6073
	ContentPackageUnlockScRsp                          = 7544
	RogueNpcDisappearScRsp                             = 5689
	EnhanceRogueBuffScRsp                              = 1900
	FinishChapterScNotify                              = 4983
	InteractChargerCsReq                               = 6883
	PlayBackGroundMusicScRsp                           = 3142
	BuyBpLevelScRsp                                    = 3033
	PlayerReturnPointChangeScNotify                    = 4542
	SceneReviveAfterRebattleScRsp                      = 1500
	RogueTournQueryScRsp                               = 6033
	MonopolyUpgradeAssetScRsp                          = 7011
	NewAssistHistoryNotify                             = 2936
	HeliobusInfoChangedScNotify                        = 5856
	HandleFriendCsReq                                  = 2938
	GetShareDataCsReq                                  = 4183
	ChessRogueCheatRollScRsp                           = 5475
	LobbyModifyPlayerInfoCsReq                         = 7374
	MonopolyGuessChooseScRsp                           = 7025
	TravelBrochureRemovePasterScRsp                    = 6438
	StartFightFestCsReq                                = 7292
	MonopolyReRollRandomCsReq                          = 7045
	GetBattleCollegeDataCsReq                          = 5798
	EvolveBuildShopAbilityResetScRsp                   = 7116
	GetSocialEventServerCacheScRsp                     = 7029
	RogueMagicEnterScRsp                               = 7742
	MonopolyGameCreateScNotify                         = 7065
	FightMatch3SwapScRsp                               = 30133
	InterruptMissionEventScRsp                         = 1252
	StartTrialActivityScRsp                            = 2603
	ChessRogueRollDiceCsReq                            = 5403
	StartFightFestScRsp                                = 7252
	MissionRewardScNotify                              = 1279
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5598
	RelicAvatarRecommendScRsp                          = 2402
	ChessRogueReRollDiceCsReq                          = 5590
	ChessRogueUpdateMoneyInfoScNotify                  = 5557
	ChessRogueQueryAeonDimensionsScRsp                 = 5421
	HeartDialScriptChangeScNotify                      = 6328
	ClockParkUseBuffScRsp                              = 7230
	StartFinishSubMissionScNotify                      = 1223
	GetAetherDivideChallengeInfoCsReq                  = 4900
	MissionGroupWarnScNotify                           = 1256
	FightFestUpdateCoinNotify                          = 7282
	GetGunPlayDataScRsp                                = 4200
	RankUpAvatarCsReq                                  = 322
	SceneReviveAfterRebattleCsReq                      = 1410
	EndDrinkMakerSequenceScRsp                         = 6995
	GetFirstTalkNpcScRsp                               = 2177
	TriggerVoiceCsReq                                  = 4189
	GetFriendDevelopmentInfoScRsp                      = 2903
	MonopolyRollRandomScRsp                            = 7005
	ShareCsReq                                         = 4198
	AddRelicFilterPlanCsReq                            = 587
	SubmitOrigamiItemScRsp                             = 4105
	GetMapRotationDataCsReq                            = 6878
	FinishPerformSectionIdScRsp                        = 2738
	FightMatch3DataScRsp                               = 30171
	RefreshAlleyOrderScRsp                             = 4720
	TrialBackGroundMusicScRsp                          = 3112
	CancelMatchCsReq                                   = 7342
	LeaveMapRotationRegionCsReq                        = 6828
	StartWolfBroGameCsReq                              = 6598
	ChessRogueGiveUpRollCsReq                          = 5480
	RogueMagicSetAutoDressInMagicUnitCsReq             = 7736
	TakeMailAttachmentScRsp                            = 812
	RogueMagicAutoDressInUnitCsReq                     = 7725
	InterruptMissionEventCsReq                         = 1260
	ShowNewSupplementVisitorScRsp                      = 3756
	LeaveChallengeCsReq                                = 1779
	LobbyCreateCsReq                                   = 7379
	EvolveBuildLeaveScRsp                              = 7101
	SelectChessRogueSubStoryCsReq                      = 5573
	RogueArcadeGetInfoScRsp                            = 7682
	GetSpringRecoverDataScRsp                          = 1464
	TakePromotionRewardCsReq                           = 331
	GetLineupAvatarDataScRsp                           = 789
	GetCurLineupDataScRsp                              = 742
	HeliobusSelectSkillCsReq                           = 5805
	RogueTournReviveAvatarScRsp                        = 6037
	DiscardRelicScRsp                                  = 519
	MultiplayerFightGameFinishScNotify                 = 1012
	UpdateMapRotationDataScNotify                      = 6845
	ChessRogueQuitCsReq                                = 5563
	ChessRogueReRollDiceScRsp                          = 5593
	CancelActivityExpeditionCsReq                      = 2556
	ExchangeHcoinScRsp                                 = 552
	EnterFightActivityStageScRsp                       = 3679
	RankUpEquipmentScRsp                               = 538
	BatchGetQuestDataScRsp                             = 905
	UnlockSkilltreeScRsp                               = 377
	SetAetherDivideLineUpScRsp                         = 4822
	AetherDivideSpiritExpUpScRsp                       = 4811
	ClockParkGetOngoingScriptInfoCsReq                 = 7213
	TakeMaterialSubmitActivityRewardScRsp              = 2660
	TakeExpeditionRewardScRsp                          = 2512
	EnterFightActivityStageCsReq                       = 3642
	StartAetherDivideStageBattleScRsp                  = 4860
	LogisticsGameScRsp                                 = 4742
	ChessRogueSelectBpScRsp                            = 5546
	ReportPlayerCsReq                                  = 2952
	FightFestUnlockSkillNotify                         = 7260
	MuseumRandomEventQueryScRsp                        = 4304
	AlleyPlacingGameCsReq                              = 4789
	TakeRogueAeonLevelRewardCsReq                      = 1897
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(MuseumTargetRewardNotify, func() any { return new(proto.MuseumTargetRewardNotify) })
	c.regMsg(GetGameStateServiceConfigCsReq, func() any { return new(proto.GetGameStateServiceConfigCsReq) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(ModifyRelicFilterPlanScRsp, func() any { return new(proto.ModifyRelicFilterPlanScRsp) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoCsReq, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoCsReq) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(GetPlayerReturnMultiDropInfoScRsp, func() any { return new(proto.GetPlayerReturnMultiDropInfoScRsp) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(TakeAssistRewardCsReq, func() any { return new(proto.TakeAssistRewardCsReq) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(ClockParkGetInfoCsReq, func() any { return new(proto.ClockParkGetInfoCsReq) })
	c.regMsg(SwordTrainingUnlockSyncScNotify, func() any { return new(proto.SwordTrainingUnlockSyncScNotify) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(StartTimedFarmElementCsReq, func() any { return new(proto.StartTimedFarmElementCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(GetMultiPathAvatarInfoCsReq, func() any { return new(proto.GetMultiPathAvatarInfoCsReq) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(StartTrackPhotoStageCsReq, func() any { return new(proto.StartTrackPhotoStageCsReq) })
	c.regMsg(SubmitOrigamiItemCsReq, func() any { return new(proto.SubmitOrigamiItemCsReq) })
	c.regMsg(MultiplayerMatch3FinishScNotify, func() any { return new(proto.MultiplayerMatch3FinishScNotify) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(RogueMagicSetAutoDressInMagicUnitScRsp, func() any { return new(proto.RogueMagicSetAutoDressInMagicUnitScRsp) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(LobbyJoinScRsp, func() any { return new(proto.LobbyJoinScRsp) })
	c.regMsg(MuseumTakeCollectRewardScRsp, func() any { return new(proto.MuseumTakeCollectRewardScRsp) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(ContentPackageGetDataCsReq, func() any { return new(proto.ContentPackageGetDataCsReq) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(RogueTournDeleteArchiveScRsp, func() any { return new(proto.RogueTournDeleteArchiveScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(SwordTrainingStartGameScRsp, func() any { return new(proto.SwordTrainingStartGameScRsp) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(ReBattleAfterBattleLoseCsNotify, func() any { return new(proto.ReBattleAfterBattleLoseCsNotify) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(MusicRhythmSaveSongConfigDataScRsp, func() any { return new(proto.MusicRhythmSaveSongConfigDataScRsp) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(EvolveBuildCoinNotify, func() any { return new(proto.EvolveBuildCoinNotify) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(CurAssistChangedNotify, func() any { return new(proto.CurAssistChangedNotify) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(LobbyJoinCsReq, func() any { return new(proto.LobbyJoinCsReq) })
	c.regMsg(ClockParkQuitScriptScRsp, func() any { return new(proto.ClockParkQuitScriptScRsp) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(SetMultipleAvatarPathsScRsp, func() any { return new(proto.SetMultipleAvatarPathsScRsp) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(ChangeStoryLineFinishScNotify, func() any { return new(proto.ChangeStoryLineFinishScNotify) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(RogueMagicEnterCsReq, func() any { return new(proto.RogueMagicEnterCsReq) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(SwordTrainingTurnActionScRsp, func() any { return new(proto.SwordTrainingTurnActionScRsp) })
	c.regMsg(DailyFirstEnterMonopolyActivityCsReq, func() any { return new(proto.DailyFirstEnterMonopolyActivityCsReq) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(GetCrossInfoCsReq, func() any { return new(proto.GetCrossInfoCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupCsReq, func() any { return new(proto.SetBoxingClubResonanceLineupCsReq) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(MultiplayerFightGiveUpCsReq, func() any { return new(proto.MultiplayerFightGiveUpCsReq) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(AddRelicFilterPlanScRsp, func() any { return new(proto.AddRelicFilterPlanScRsp) })
	c.regMsg(ContentPackageUnlockCsReq, func() any { return new(proto.ContentPackageUnlockCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(ExchangeStaminaCsReq, func() any { return new(proto.ExchangeStaminaCsReq) })
	c.regMsg(SwordTrainingGameSyncChangeScNotify, func() any { return new(proto.SwordTrainingGameSyncChangeScNotify) })
	c.regMsg(GetBoxingClubInfoCsReq, func() any { return new(proto.GetBoxingClubInfoCsReq) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(GetCrossInfoScRsp, func() any { return new(proto.GetCrossInfoScRsp) })
	c.regMsg(MarkAvatarScRsp, func() any { return new(proto.MarkAvatarScRsp) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(MonopolySelectOptionScRsp, func() any { return new(proto.MonopolySelectOptionScRsp) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(RogueTournGetAllArchiveScRsp, func() any { return new(proto.RogueTournGetAllArchiveScRsp) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(RogueModifierSelectCellScRsp, func() any { return new(proto.RogueModifierSelectCellScRsp) })
	c.regMsg(RogueTournGetPermanentTalentInfoCsReq, func() any { return new(proto.RogueTournGetPermanentTalentInfoCsReq) })
	c.regMsg(RogueTournEnterRoomCsReq, func() any { return new(proto.RogueTournEnterRoomCsReq) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(GetFriendChallengeLineupCsReq, func() any { return new(proto.GetFriendChallengeLineupCsReq) })
	c.regMsg(MatchThreeSyncDataScNotify, func() any { return new(proto.MatchThreeSyncDataScNotify) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneCsReq) })
	c.regMsg(GetFriendChallengeLineupScRsp, func() any { return new(proto.GetFriendChallengeLineupScRsp) })
	c.regMsg(BattleLogReportCsReq, func() any { return new(proto.BattleLogReportCsReq) })
	c.regMsg(RogueTournConfirmSettleScRsp, func() any { return new(proto.RogueTournConfirmSettleScRsp) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(DifficultyAdjustmentGetDataCsReq, func() any { return new(proto.DifficultyAdjustmentGetDataCsReq) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(RogueMagicScepterTakeOffUnitScRsp, func() any { return new(proto.RogueMagicScepterTakeOffUnitScRsp) })
	c.regMsg(HeliobusStartRaidCsReq, func() any { return new(proto.HeliobusStartRaidCsReq) })
	c.regMsg(RaidCollectionDataCsReq, func() any { return new(proto.RaidCollectionDataCsReq) })
	c.regMsg(ClockParkGetInfoScRsp, func() any { return new(proto.ClockParkGetInfoScRsp) })
	c.regMsg(SwordTrainingSelectEndingCsReq, func() any { return new(proto.SwordTrainingSelectEndingCsReq) })
	c.regMsg(MatchThreeLevelEndCsReq, func() any { return new(proto.MatchThreeLevelEndCsReq) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(GetFriendChallengeDetailCsReq, func() any { return new(proto.GetFriendChallengeDetailCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(StartPartialChallengeScRsp, func() any { return new(proto.StartPartialChallengeScRsp) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(EnterSummonActivityStageScRsp, func() any { return new(proto.EnterSummonActivityStageScRsp) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(ServerSimulateBattleFinishScNotify, func() any { return new(proto.ServerSimulateBattleFinishScNotify) })
	c.regMsg(SubmitOfferingItemScRsp, func() any { return new(proto.SubmitOfferingItemScRsp) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(GetUpdatedArchiveDataScRsp, func() any { return new(proto.GetUpdatedArchiveDataScRsp) })
	c.regMsg(StrongChallengeActivityBattleEndScNotify, func() any { return new(proto.StrongChallengeActivityBattleEndScNotify) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(UpdateMovieRacingDataCsReq, func() any { return new(proto.UpdateMovieRacingDataCsReq) })
	c.regMsg(MonopolyGameBingoFlipCardScRsp, func() any { return new(proto.MonopolyGameBingoFlipCardScRsp) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(SyncRogueCommonDialogueOptionFinishScNotify, func() any { return new(proto.SyncRogueCommonDialogueOptionFinishScNotify) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(SwordTrainingDailyPhaseConfirmCsReq, func() any { return new(proto.SwordTrainingDailyPhaseConfirmCsReq) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(GetGunPlayDataCsReq, func() any { return new(proto.GetGunPlayDataCsReq) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(RogueMagicReviveAvatarCsReq, func() any { return new(proto.RogueMagicReviveAvatarCsReq) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(ClientObjUploadCsReq, func() any { return new(proto.ClientObjUploadCsReq) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(RaidCollectionEnterNextRaidCsReq, func() any { return new(proto.RaidCollectionEnterNextRaidCsReq) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(LogisticsGameCsReq, func() any { return new(proto.LogisticsGameCsReq) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(RogueMagicGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueMagicGetMiscRealTimeDataScRsp) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(MonopolyMoveCsReq, func() any { return new(proto.MonopolyMoveCsReq) })
	c.regMsg(RogueMagicGetTalentInfoScRsp, func() any { return new(proto.RogueMagicGetTalentInfoScRsp) })
	c.regMsg(GetStoryLineInfoScRsp, func() any { return new(proto.GetStoryLineInfoScRsp) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(MonopolyQuizDurationChangeScNotify, func() any { return new(proto.MonopolyQuizDurationChangeScNotify) })
	c.regMsg(ChessRogueNousEnableRogueTalentScRsp, func() any { return new(proto.ChessRogueNousEnableRogueTalentScRsp) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(ClockParkGetOngoingScriptInfoScRsp, func() any { return new(proto.ClockParkGetOngoingScriptInfoScRsp) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(RogueMagicUnitReforgeScRsp, func() any { return new(proto.RogueMagicUnitReforgeScRsp) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(CancelMatchScRsp, func() any { return new(proto.CancelMatchScRsp) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(LobbyKickOutScRsp, func() any { return new(proto.LobbyKickOutScRsp) })
	c.regMsg(SwordTrainingStoryConfirmCsReq, func() any { return new(proto.SwordTrainingStoryConfirmCsReq) })
	c.regMsg(RogueMagicEnterRoomCsReq, func() any { return new(proto.RogueMagicEnterRoomCsReq) })
	c.regMsg(GetRogueExhibitionCsReq, func() any { return new(proto.GetRogueExhibitionCsReq) })
	c.regMsg(EvolveBuildTakeExpRewardCsReq, func() any { return new(proto.EvolveBuildTakeExpRewardCsReq) })
	c.regMsg(MuseumTakeCollectRewardCsReq, func() any { return new(proto.MuseumTakeCollectRewardCsReq) })
	c.regMsg(QuitBattleScRsp, func() any { return new(proto.QuitBattleScRsp) })
	c.regMsg(MultiplayerGetFightGateScRsp, func() any { return new(proto.MultiplayerGetFightGateScRsp) })
	c.regMsg(SyncRogueSeasonFinishScNotify, func() any { return new(proto.SyncRogueSeasonFinishScNotify) })
	c.regMsg(MarkRelicFilterPlanScRsp, func() any { return new(proto.MarkRelicFilterPlanScRsp) })
	c.regMsg(RogueTournDifficultyCompNotify, func() any { return new(proto.RogueTournDifficultyCompNotify) })
	c.regMsg(MonopolyGuessDrawScNotify, func() any { return new(proto.MonopolyGuessDrawScNotify) })
	c.regMsg(TakeMaterialSubmitActivityRewardCsReq, func() any { return new(proto.TakeMaterialSubmitActivityRewardCsReq) })
	c.regMsg(DeleteFriendScRsp, func() any { return new(proto.DeleteFriendScRsp) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(RogueWorkbenchHandleFuncCsReq, func() any { return new(proto.RogueWorkbenchHandleFuncCsReq) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(RogueTournEnterLayerScRsp, func() any { return new(proto.RogueTournEnterLayerScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(ChessRogueNousGetRogueTalentInfoScRsp, func() any { return new(proto.ChessRogueNousGetRogueTalentInfoScRsp) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(RogueMagicScepterDressInUnitCsReq, func() any { return new(proto.RogueMagicScepterDressInUnitCsReq) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(StartMatchScRsp, func() any { return new(proto.StartMatchScRsp) })
	c.regMsg(MatchThreeSetBirdPosScRsp, func() any { return new(proto.MatchThreeSetBirdPosScRsp) })
	c.regMsg(RaidCollectionEnterNextRaidScRsp, func() any { return new(proto.RaidCollectionEnterNextRaidScRsp) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(LobbyInviteCsReq, func() any { return new(proto.LobbyInviteCsReq) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(SpringRefreshScRsp, func() any { return new(proto.SpringRefreshScRsp) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(SummonPetScRsp, func() any { return new(proto.SummonPetScRsp) })
	c.regMsg(TakeAssistRewardScRsp, func() any { return new(proto.TakeAssistRewardScRsp) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(ClockParkFinishScriptScNotify, func() any { return new(proto.ClockParkFinishScriptScNotify) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(EnterSwordTrainingExamCsReq, func() any { return new(proto.EnterSwordTrainingExamCsReq) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(GetMultiPathAvatarInfoScRsp, func() any { return new(proto.GetMultiPathAvatarInfoScRsp) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(ClockParkUnlockTalentScRsp, func() any { return new(proto.ClockParkUnlockTalentScRsp) })
	c.regMsg(SetAvatarPathCsReq, func() any { return new(proto.SetAvatarPathCsReq) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(SetRogueExhibitionCsReq, func() any { return new(proto.SetRogueExhibitionCsReq) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(GetMonopolyInfoScRsp, func() any { return new(proto.GetMonopolyInfoScRsp) })
	c.regMsg(SwordTrainingMarkEndingViewedCsReq, func() any { return new(proto.SwordTrainingMarkEndingViewedCsReq) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(RelicFilterPlanClearNameScNotify, func() any { return new(proto.RelicFilterPlanClearNameScNotify) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(ClockParkStartScriptCsReq, func() any { return new(proto.ClockParkStartScriptCsReq) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(GetFriendBattleRecordDetailCsReq, func() any { return new(proto.GetFriendBattleRecordDetailCsReq) })
	c.regMsg(MusicRhythmDataCsReq, func() any { return new(proto.MusicRhythmDataCsReq) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(RogueTournResetPermanentTalentScRsp, func() any { return new(proto.RogueTournResetPermanentTalentScRsp) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(AcceptMissionEventScRsp, func() any { return new(proto.AcceptMissionEventScRsp) })
	c.regMsg(UpdateMovieRacingDataScRsp, func() any { return new(proto.UpdateMovieRacingDataScRsp) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffCsReq, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffCsReq) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(AvatarPathChangedNotify, func() any { return new(proto.AvatarPathChangedNotify) })
	c.regMsg(RegionStopScNotify, func() any { return new(proto.RegionStopScNotify) })
	c.regMsg(RemoveRotaterCsReq, func() any { return new(proto.RemoveRotaterCsReq) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(MonopolyGameGachaCsReq, func() any { return new(proto.MonopolyGameGachaCsReq) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(UpdatePlayerSettingScRsp, func() any { return new(proto.UpdatePlayerSettingScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownCsReq, func() any { return new(proto.EvolveBuildShopAbilityDownCsReq) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(MonopolyGetDailyInitItemScRsp, func() any { return new(proto.MonopolyGetDailyInitItemScRsp) })
	c.regMsg(MonopolyMoveScRsp, func() any { return new(proto.MonopolyMoveScRsp) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(MonopolyEventLoadUpdateScNotify, func() any { return new(proto.MonopolyEventLoadUpdateScNotify) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(EvolveBuildFinishScNotify, func() any { return new(proto.EvolveBuildFinishScNotify) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(SubmitMaterialSubmitActivityMaterialScRsp, func() any { return new(proto.SubmitMaterialSubmitActivityMaterialScRsp) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(RogueMagicUnitReforgeCsReq, func() any { return new(proto.RogueMagicUnitReforgeCsReq) })
	c.regMsg(RogueModifierUpdateNotify, func() any { return new(proto.RogueModifierUpdateNotify) })
	c.regMsg(FightHeartBeatScRsp, func() any { return new(proto.FightHeartBeatScRsp) })
	c.regMsg(GetRelicFilterPlanScRsp, func() any { return new(proto.GetRelicFilterPlanScRsp) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(RogueGetGambleInfoScRsp, func() any { return new(proto.RogueGetGambleInfoScRsp) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(SwordTrainingTurnActionCsReq, func() any { return new(proto.SwordTrainingTurnActionCsReq) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(GetMonopolyDailyReportCsReq, func() any { return new(proto.GetMonopolyDailyReportCsReq) })
	c.regMsg(GetChallengeGroupStatisticsScRsp, func() any { return new(proto.GetChallengeGroupStatisticsScRsp) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(MatchThreeLevelEndScRsp, func() any { return new(proto.MatchThreeLevelEndScRsp) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(StartPartialChallengeCsReq, func() any { return new(proto.StartPartialChallengeCsReq) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(MonopolyClickMbtiReportScRsp, func() any { return new(proto.MonopolyClickMbtiReportScRsp) })
	c.regMsg(DrinkMakerUpdateTipsNotify, func() any { return new(proto.DrinkMakerUpdateTipsNotify) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(GetPlayerReturnMultiDropInfoCsReq, func() any { return new(proto.GetPlayerReturnMultiDropInfoCsReq) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(GetChallengeGroupStatisticsCsReq, func() any { return new(proto.GetChallengeGroupStatisticsCsReq) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(MusicRhythmDataScRsp, func() any { return new(proto.MusicRhythmDataScRsp) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(MonopolyClickCellCsReq, func() any { return new(proto.MonopolyClickCellCsReq) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(FightMatch3DataCsReq, func() any { return new(proto.FightMatch3DataCsReq) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(GetMonopolyInfoCsReq, func() any { return new(proto.GetMonopolyInfoCsReq) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(DifficultyAdjustmentGetDataScRsp, func() any { return new(proto.DifficultyAdjustmentGetDataScRsp) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(FightFestUpdateChallengeRecordNotify, func() any { return new(proto.FightFestUpdateChallengeRecordNotify) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(SpringRecoverScRsp, func() any { return new(proto.SpringRecoverScRsp) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(LobbyGetInfoScRsp, func() any { return new(proto.LobbyGetInfoScRsp) })
	c.regMsg(GetMbtiReportScRsp, func() any { return new(proto.GetMbtiReportScRsp) })
	c.regMsg(LastSpringRefreshTimeNotify, func() any { return new(proto.LastSpringRefreshTimeNotify) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(FightMatch3ForceUpdateNotify, func() any { return new(proto.FightMatch3ForceUpdateNotify) })
	c.regMsg(ClockParkBattleEndScNotify, func() any { return new(proto.ClockParkBattleEndScNotify) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(RogueMagicAreaUpdateScNotify, func() any { return new(proto.RogueMagicAreaUpdateScNotify) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(ModifyRelicFilterPlanCsReq, func() any { return new(proto.ModifyRelicFilterPlanCsReq) })
	c.regMsg(RogueMagicQueryCsReq, func() any { return new(proto.RogueMagicQueryCsReq) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(EvolveBuildQueryInfoCsReq, func() any { return new(proto.EvolveBuildQueryInfoCsReq) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(ContentPackageTransferScNotify, func() any { return new(proto.ContentPackageTransferScNotify) })
	c.regMsg(FightMatch3TurnStartScNotify, func() any { return new(proto.FightMatch3TurnStartScNotify) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(RogueTournRenameArchiveScRsp, func() any { return new(proto.RogueTournRenameArchiveScRsp) })
	c.regMsg(ClockParkQuitScriptCsReq, func() any { return new(proto.ClockParkQuitScriptCsReq) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(MonopolyGuessBuyInformationCsReq, func() any { return new(proto.MonopolyGuessBuyInformationCsReq) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(EvolveBuildTakeExpRewardScRsp, func() any { return new(proto.EvolveBuildTakeExpRewardScRsp) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(EvolveBuildGiveupCsReq, func() any { return new(proto.EvolveBuildGiveupCsReq) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(LobbyQuitScRsp, func() any { return new(proto.LobbyQuitScRsp) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(MusicRhythmUnlockTrackScNotify, func() any { return new(proto.MusicRhythmUnlockTrackScNotify) })
	c.regMsg(GiveUpBoxingClubChallengeScRsp, func() any { return new(proto.GiveUpBoxingClubChallengeScRsp) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(MonopolyGameRaiseRatioCsReq, func() any { return new(proto.MonopolyGameRaiseRatioCsReq) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(MonopolyLikeScRsp, func() any { return new(proto.MonopolyLikeScRsp) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(ClientDownloadDataScNotify, func() any { return new(proto.ClientDownloadDataScNotify) })
	c.regMsg(SwordTrainingSelectEndingScRsp, func() any { return new(proto.SwordTrainingSelectEndingScRsp) })
	c.regMsg(MusicRhythmStartLevelCsReq, func() any { return new(proto.MusicRhythmStartLevelCsReq) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(MultiplayerFightGameStartScNotify, func() any { return new(proto.MultiplayerFightGameStartScNotify) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(MonopolySelectOptionCsReq, func() any { return new(proto.MonopolySelectOptionCsReq) })
	c.regMsg(RogueTournLeaveScRsp, func() any { return new(proto.RogueTournLeaveScRsp) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(SetForbidOtherApplyFriendScRsp, func() any { return new(proto.SetForbidOtherApplyFriendScRsp) })
	c.regMsg(MonopolyAcceptQuizScRsp, func() any { return new(proto.MonopolyAcceptQuizScRsp) })
	c.regMsg(MonopolyGuessBuyInformationScRsp, func() any { return new(proto.MonopolyGuessBuyInformationScRsp) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(WolfBroGameExplodeMonsterScRsp, func() any { return new(proto.WolfBroGameExplodeMonsterScRsp) })
	c.regMsg(RogueTournEnterLayerCsReq, func() any { return new(proto.RogueTournEnterLayerCsReq) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(SetAvatarPathScRsp, func() any { return new(proto.SetAvatarPathScRsp) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(TakeOfferingRewardScRsp, func() any { return new(proto.TakeOfferingRewardScRsp) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(RogueGetGambleInfoCsReq, func() any { return new(proto.RogueGetGambleInfoCsReq) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(WolfBroGameActivateBulletScRsp, func() any { return new(proto.WolfBroGameActivateBulletScRsp) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(WolfBroGameExplodeMonsterCsReq, func() any { return new(proto.WolfBroGameExplodeMonsterCsReq) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(RogueWorkbenchSelectFuncScRsp, func() any { return new(proto.RogueWorkbenchSelectFuncScRsp) })
	c.regMsg(SwordTrainingActionTurnSettleScNotify, func() any { return new(proto.SwordTrainingActionTurnSettleScNotify) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(ClockParkHandleWaitOperationCsReq, func() any { return new(proto.ClockParkHandleWaitOperationCsReq) })
	c.regMsg(EvolveBuildUnlockInfoNotify, func() any { return new(proto.EvolveBuildUnlockInfoNotify) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(MonopolyGuessChooseCsReq, func() any { return new(proto.MonopolyGuessChooseCsReq) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(SetMultipleAvatarPathsCsReq, func() any { return new(proto.SetMultipleAvatarPathsCsReq) })
	c.regMsg(DrinkMakerDayEndScNotify, func() any { return new(proto.DrinkMakerDayEndScNotify) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(StartTimedFarmElementScRsp, func() any { return new(proto.StartTimedFarmElementScRsp) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(SettleTrackPhotoStageScRsp, func() any { return new(proto.SettleTrackPhotoStageScRsp) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(GetMultipleDropInfoScRsp, func() any { return new(proto.GetMultipleDropInfoScRsp) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(GetAuthkeyCsReq, func() any { return new(proto.GetAuthkeyCsReq) })
	c.regMsg(TakeFightActivityRewardScRsp, func() any { return new(proto.TakeFightActivityRewardScRsp) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(QuitBattleScNotify, func() any { return new(proto.QuitBattleScNotify) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(DrinkMakerChallengeScRsp, func() any { return new(proto.DrinkMakerChallengeScRsp) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(RogueMagicLeaveCsReq, func() any { return new(proto.RogueMagicLeaveCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(GetTrainVisitorBehaviorCsReq, func() any { return new(proto.GetTrainVisitorBehaviorCsReq) })
	c.regMsg(RogueWorkbenchHandleFuncScRsp, func() any { return new(proto.RogueWorkbenchHandleFuncScRsp) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(ContentPackageSyncDataScNotify, func() any { return new(proto.ContentPackageSyncDataScNotify) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(MonopolyClickCellScRsp, func() any { return new(proto.MonopolyClickCellScRsp) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(EvolveBuildStartLevelScRsp, func() any { return new(proto.EvolveBuildStartLevelScRsp) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(SwordTrainingStoryBattleCsReq, func() any { return new(proto.SwordTrainingStoryBattleCsReq) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(MultiplayerFightGameStateCsReq, func() any { return new(proto.MultiplayerFightGameStateCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(ActivityRaidPlacingGameCsReq, func() any { return new(proto.ActivityRaidPlacingGameCsReq) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(MusicRhythmSaveSongConfigDataCsReq, func() any { return new(proto.MusicRhythmSaveSongConfigDataCsReq) })
	c.regMsg(RogueTournEnterCsReq, func() any { return new(proto.RogueTournEnterCsReq) })
	c.regMsg(LobbySyncInfoScNotify, func() any { return new(proto.LobbySyncInfoScNotify) })
	c.regMsg(RogueTournGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueTournGetMiscRealTimeDataScRsp) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(RogueTournResetPermanentTalentCsReq, func() any { return new(proto.RogueTournResetPermanentTalentCsReq) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(GetTrackPhotoActivityDataCsReq, func() any { return new(proto.GetTrackPhotoActivityDataCsReq) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(MonopolyEventSelectFriendScRsp, func() any { return new(proto.MonopolyEventSelectFriendScRsp) })
	c.regMsg(UpdateGunPlayDataScRsp, func() any { return new(proto.UpdateGunPlayDataScRsp) })
	c.regMsg(GetTrackPhotoActivityDataScRsp, func() any { return new(proto.GetTrackPhotoActivityDataScRsp) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(StartTrackPhotoStageScRsp, func() any { return new(proto.StartTrackPhotoStageScRsp) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(LobbyInviteScNotify, func() any { return new(proto.LobbyInviteScNotify) })
	c.regMsg(SettleTrackPhotoStageCsReq, func() any { return new(proto.SettleTrackPhotoStageCsReq) })
	c.regMsg(MonopolyGiveUpCurContentCsReq, func() any { return new(proto.MonopolyGiveUpCurContentCsReq) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(MonopolySocialEventEffectScNotify, func() any { return new(proto.MonopolySocialEventEffectScNotify) })
	c.regMsg(CommonRogueQueryScRsp, func() any { return new(proto.CommonRogueQueryScRsp) })
	c.regMsg(QuitTrackPhotoStageScRsp, func() any { return new(proto.QuitTrackPhotoStageScRsp) })
	c.regMsg(RecallPetCsReq, func() any { return new(proto.RecallPetCsReq) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(AcceptMissionEventCsReq, func() any { return new(proto.AcceptMissionEventCsReq) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	c.regMsg(RogueTournEnterScRsp, func() any { return new(proto.RogueTournEnterScRsp) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(RogueTournEnablePermanentTalentScRsp, func() any { return new(proto.RogueTournEnablePermanentTalentScRsp) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(RogueMagicReviveCostUpdateScNotify, func() any { return new(proto.RogueMagicReviveCostUpdateScNotify) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(GetPetDataCsReq, func() any { return new(proto.GetPetDataCsReq) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(RogueMagicScepterDressInUnitScRsp, func() any { return new(proto.RogueMagicScepterDressInUnitScRsp) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(GetRogueCommonDialogueDataCsReq, func() any { return new(proto.GetRogueCommonDialogueDataCsReq) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(SetRogueCollectionScRsp, func() any { return new(proto.SetRogueCollectionScRsp) })
	c.regMsg(UpdateRotaterScNotify, func() any { return new(proto.UpdateRotaterScNotify) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(RogueArcadeLeaveScRsp, func() any { return new(proto.RogueArcadeLeaveScRsp) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(MonopolyGetDailyInitItemCsReq, func() any { return new(proto.MonopolyGetDailyInitItemCsReq) })
	c.regMsg(SwordTrainingExamResultConfirmCsReq, func() any { return new(proto.SwordTrainingExamResultConfirmCsReq) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(SwordTrainingDialogueSelectOptionCsReq, func() any { return new(proto.SwordTrainingDialogueSelectOptionCsReq) })
	c.regMsg(MonopolyEventSelectFriendCsReq, func() any { return new(proto.MonopolyEventSelectFriendCsReq) })
	c.regMsg(SetRedPointStatusScNotify, func() any { return new(proto.SetRedPointStatusScNotify) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(SwordTrainingResumeGameScRsp, func() any { return new(proto.SwordTrainingResumeGameScRsp) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(RelicAvatarRecommendCsReq, func() any { return new(proto.RelicAvatarRecommendCsReq) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(SwordTrainingStoryConfirmScRsp, func() any { return new(proto.SwordTrainingStoryConfirmScRsp) })
	c.regMsg(SwordTrainingLearnSkillCsReq, func() any { return new(proto.SwordTrainingLearnSkillCsReq) })
	c.regMsg(SwordTrainingGiveUpGameCsReq, func() any { return new(proto.SwordTrainingGiveUpGameCsReq) })
	c.regMsg(RogueMagicGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueMagicGetMiscRealTimeDataCsReq) })
	c.regMsg(MusicRhythmFinishLevelCsReq, func() any { return new(proto.MusicRhythmFinishLevelCsReq) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(SwordTrainingDialogueSelectOptionScRsp, func() any { return new(proto.SwordTrainingDialogueSelectOptionScRsp) })
	c.regMsg(GetSwordTrainingDataCsReq, func() any { return new(proto.GetSwordTrainingDataCsReq) })
	c.regMsg(SwordTrainingResumeGameCsReq, func() any { return new(proto.SwordTrainingResumeGameCsReq) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(GetMonopolyDailyReportScRsp, func() any { return new(proto.GetMonopolyDailyReportScRsp) })
	c.regMsg(SwordTrainingLearnSkillScRsp, func() any { return new(proto.SwordTrainingLearnSkillScRsp) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(RogueMagicLeaveScRsp, func() any { return new(proto.RogueMagicLeaveScRsp) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(SwordTrainingSetSkillTraceCsReq, func() any { return new(proto.SwordTrainingSetSkillTraceCsReq) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(SwordTrainingGiveUpGameScRsp, func() any { return new(proto.SwordTrainingGiveUpGameScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(MonopolyConditionUpdateScNotify, func() any { return new(proto.MonopolyConditionUpdateScNotify) })
	c.regMsg(SwordTrainingGameSettleScNotify, func() any { return new(proto.SwordTrainingGameSettleScNotify) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(SwordTrainingSetSkillTraceScRsp, func() any { return new(proto.SwordTrainingSetSkillTraceScRsp) })
	c.regMsg(SwordTrainingMarkEndingViewedScRsp, func() any { return new(proto.SwordTrainingMarkEndingViewedScRsp) })
	c.regMsg(SwordTrainingDailyPhaseConfirmScRsp, func() any { return new(proto.SwordTrainingDailyPhaseConfirmScRsp) })
	c.regMsg(EnterSwordTrainingExamScRsp, func() any { return new(proto.EnterSwordTrainingExamScRsp) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(GetSummonActivityDataScRsp, func() any { return new(proto.GetSummonActivityDataScRsp) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(BoxingClubChallengeUpdateScNotify, func() any { return new(proto.BoxingClubChallengeUpdateScNotify) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(GetSummonActivityDataCsReq, func() any { return new(proto.GetSummonActivityDataCsReq) })
	c.regMsg(GetSpringRecoverDataCsReq, func() any { return new(proto.GetSpringRecoverDataCsReq) })
	c.regMsg(FightMatch3ChatScRsp, func() any { return new(proto.FightMatch3ChatScRsp) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(SelectRogueCommonDialogueOptionScRsp, func() any { return new(proto.SelectRogueCommonDialogueOptionScRsp) })
	c.regMsg(SummonActivityBattleEndScNotify, func() any { return new(proto.SummonActivityBattleEndScNotify) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(StoryLineTrialAvatarChangeScNotify, func() any { return new(proto.StoryLineTrialAvatarChangeScNotify) })
	c.regMsg(StoryLineInfoScNotify, func() any { return new(proto.StoryLineInfoScNotify) })
	c.regMsg(GetMaterialSubmitActivityDataCsReq, func() any { return new(proto.GetMaterialSubmitActivityDataCsReq) })
	c.regMsg(RogueTournReEnterRogueCocoonStageCsReq, func() any { return new(proto.RogueTournReEnterRogueCocoonStageCsReq) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(MatchThreeGetDataScRsp, func() any { return new(proto.MatchThreeGetDataScRsp) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(MissionEventRewardScNotify, func() any { return new(proto.MissionEventRewardScNotify) })
	c.regMsg(RogueMagicAutoDressInMagicUnitChangeScNotify, func() any { return new(proto.RogueMagicAutoDressInMagicUnitChangeScNotify) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(StartStarFightLevelScRsp, func() any { return new(proto.StartStarFightLevelScRsp) })
	c.regMsg(GetStarFightDataScRsp, func() any { return new(proto.GetStarFightDataScRsp) })
	c.regMsg(StarFightDataChangeNotify, func() any { return new(proto.StarFightDataChangeNotify) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(LobbyModifyPlayerInfoScRsp, func() any { return new(proto.LobbyModifyPlayerInfoScRsp) })
	c.regMsg(StartStarFightLevelCsReq, func() any { return new(proto.StartStarFightLevelCsReq) })
	c.regMsg(GetStarFightDataCsReq, func() any { return new(proto.GetStarFightDataCsReq) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoScRsp, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoScRsp) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(SpaceZooTakeScRsp, func() any { return new(proto.SpaceZooTakeScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(FightMatch3ChatCsReq, func() any { return new(proto.FightMatch3ChatCsReq) })
	c.regMsg(SpringRecoverSingleAvatarScRsp, func() any { return new(proto.SpringRecoverSingleAvatarScRsp) })
	c.regMsg(SpaceZooExchangeItemCsReq, func() any { return new(proto.SpaceZooExchangeItemCsReq) })
	c.regMsg(LobbyKickOutCsReq, func() any { return new(proto.LobbyKickOutCsReq) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(RogueMagicUnitComposeScRsp, func() any { return new(proto.RogueMagicUnitComposeScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(GiveUpBoxingClubChallengeCsReq, func() any { return new(proto.GiveUpBoxingClubChallengeCsReq) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(UnlockAvatarPathScRsp, func() any { return new(proto.UnlockAvatarPathScRsp) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(MonopolyAcceptQuizCsReq, func() any { return new(proto.MonopolyAcceptQuizCsReq) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(SpaceZooExchangeItemScRsp, func() any { return new(proto.SpaceZooExchangeItemScRsp) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(RogueTournWeekChallengeUpdateScNotify, func() any { return new(proto.RogueTournWeekChallengeUpdateScNotify) })
	c.regMsg(SpringRefreshCsReq, func() any { return new(proto.SpringRefreshCsReq) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(EnterChallengeNextPhaseCsReq, func() any { return new(proto.EnterChallengeNextPhaseCsReq) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(RogueTournReviveAvatarCsReq, func() any { return new(proto.RogueTournReviveAvatarCsReq) })
	c.regMsg(MultiplayerFightGameStateScRsp, func() any { return new(proto.MultiplayerFightGameStateScRsp) })
	c.regMsg(RogueMagicSettleCsReq, func() any { return new(proto.RogueMagicSettleCsReq) })
	c.regMsg(MusicRhythmStartLevelScRsp, func() any { return new(proto.MusicRhythmStartLevelScRsp) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(StartTimedCocoonStageCsReq, func() any { return new(proto.StartTimedCocoonStageCsReq) })
	c.regMsg(FightMatch3StartCountDownScNotify, func() any { return new(proto.FightMatch3StartCountDownScNotify) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(GetStoryLineInfoCsReq, func() any { return new(proto.GetStoryLineInfoCsReq) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(EvolveBuildStartLevelCsReq, func() any { return new(proto.EvolveBuildStartLevelCsReq) })
	c.regMsg(SpringRecoverCsReq, func() any { return new(proto.SpringRecoverCsReq) })
	c.regMsg(QuitBattleCsReq, func() any { return new(proto.QuitBattleCsReq) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(GetFriendAssistListCsReq, func() any { return new(proto.GetFriendAssistListCsReq) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(StartTimedCocoonStageScRsp, func() any { return new(proto.StartTimedCocoonStageScRsp) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(SwordTrainingStoryBattleScRsp, func() any { return new(proto.SwordTrainingStoryBattleScRsp) })
	c.regMsg(SetSpringRecoverConfigScRsp, func() any { return new(proto.SetSpringRecoverConfigScRsp) })
	c.regMsg(SpringRecoverSingleAvatarCsReq, func() any { return new(proto.SpringRecoverSingleAvatarCsReq) })
	c.regMsg(RogueMagicStartScRsp, func() any { return new(proto.RogueMagicStartScRsp) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(MusicRhythmUnlockSongSfxScNotify, func() any { return new(proto.MusicRhythmUnlockSongSfxScNotify) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(MarkRelicFilterPlanCsReq, func() any { return new(proto.MarkRelicFilterPlanCsReq) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(SetSpringRecoverConfigCsReq, func() any { return new(proto.SetSpringRecoverConfigCsReq) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(HealPoolInfoNotify, func() any { return new(proto.HealPoolInfoNotify) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(RogueTournLeaveCsReq, func() any { return new(proto.RogueTournLeaveCsReq) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(FinishRogueCommonDialogueCsReq, func() any { return new(proto.FinishRogueCommonDialogueCsReq) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(EnteredSceneChangeScNotify, func() any { return new(proto.EnteredSceneChangeScNotify) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(MonopolyActionResultScNotify, func() any { return new(proto.MonopolyActionResultScNotify) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(TakeRollShopRewardCsReq, func() any { return new(proto.TakeRollShopRewardCsReq) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(DifficultyAdjustmentUpdateDataCsReq, func() any { return new(proto.DifficultyAdjustmentUpdateDataCsReq) })
	c.regMsg(MonopolyLikeScNotify, func() any { return new(proto.MonopolyLikeScNotify) })
	c.regMsg(VirtualLineupTrialAvatarChangeScNotify, func() any { return new(proto.VirtualLineupTrialAvatarChangeScNotify) })
	c.regMsg(RestartChallengePhaseScRsp, func() any { return new(proto.RestartChallengePhaseScRsp) })
	c.regMsg(RogueTournReviveCostUpdateScNotify, func() any { return new(proto.RogueTournReviveCostUpdateScNotify) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(RogueTournGetSettleInfoCsReq, func() any { return new(proto.RogueTournGetSettleInfoCsReq) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(RogueTournHandBookNotify, func() any { return new(proto.RogueTournHandBookNotify) })
	c.regMsg(RogueTournExpNotify, func() any { return new(proto.RogueTournExpNotify) })
	c.regMsg(RogueTournAreaUpdateScNotify, func() any { return new(proto.RogueTournAreaUpdateScNotify) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoCsReq, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoCsReq) })
	c.regMsg(RaidCollectionDataScRsp, func() any { return new(proto.RaidCollectionDataScRsp) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(RogueTournTakeExpRewardScRsp, func() any { return new(proto.RogueTournTakeExpRewardScRsp) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(RogueTournGetAllArchiveCsReq, func() any { return new(proto.RogueTournGetAllArchiveCsReq) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(RogueTournStartScRsp, func() any { return new(proto.RogueTournStartScRsp) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(RogueTournReEnterRogueCocoonStageScRsp, func() any { return new(proto.RogueTournReEnterRogueCocoonStageScRsp) })
	c.regMsg(RogueTournSettleCsReq, func() any { return new(proto.RogueTournSettleCsReq) })
	c.regMsg(MultiplayerFightGiveUpScRsp, func() any { return new(proto.MultiplayerFightGiveUpScRsp) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(GetMaterialSubmitActivityDataScRsp, func() any { return new(proto.GetMaterialSubmitActivityDataScRsp) })
	c.regMsg(UpdatePsnSettingsInfoScRsp, func() any { return new(proto.UpdatePsnSettingsInfoScRsp) })
	c.regMsg(LobbyBeginCsReq, func() any { return new(proto.LobbyBeginCsReq) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneScRsp) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(FightKickOutScNotify, func() any { return new(proto.FightKickOutScNotify) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(RogueTournDeleteArchiveCsReq, func() any { return new(proto.RogueTournDeleteArchiveCsReq) })
	c.regMsg(TrainVisitorBehaviorFinishCsReq, func() any { return new(proto.TrainVisitorBehaviorFinishCsReq) })
	c.regMsg(RogueTournClearArchiveNameScNotify, func() any { return new(proto.RogueTournClearArchiveNameScNotify) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(RogueTournEnablePermanentTalentCsReq, func() any { return new(proto.RogueTournEnablePermanentTalentCsReq) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(MatchThreeSetBirdPosCsReq, func() any { return new(proto.MatchThreeSetBirdPosCsReq) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(GetRelicFilterPlanCsReq, func() any { return new(proto.GetRelicFilterPlanCsReq) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(RogueWorkbenchGetInfoScRsp, func() any { return new(proto.RogueWorkbenchGetInfoScRsp) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(RogueTournLevelInfoUpdateScNotify, func() any { return new(proto.RogueTournLevelInfoUpdateScNotify) })
	c.regMsg(UpdatePlayerSettingCsReq, func() any { return new(proto.UpdatePlayerSettingCsReq) })
	c.regMsg(FightMatch3ChatScNotify, func() any { return new(proto.FightMatch3ChatScNotify) })
	c.regMsg(GetFriendDevelopmentInfoCsReq, func() any { return new(proto.GetFriendDevelopmentInfoCsReq) })
	c.regMsg(RogueTournSettleScRsp, func() any { return new(proto.RogueTournSettleScRsp) })
	c.regMsg(RogueTournTakeExpRewardCsReq, func() any { return new(proto.RogueTournTakeExpRewardCsReq) })
	c.regMsg(RogueTournGetPermanentTalentInfoScRsp, func() any { return new(proto.RogueTournGetPermanentTalentInfoScRsp) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(RogueTournGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueTournGetMiscRealTimeDataCsReq) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(FightLeaveScNotify, func() any { return new(proto.FightLeaveScNotify) })
	c.regMsg(SyncRogueAeonScNotify, func() any { return new(proto.SyncRogueAeonScNotify) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(RogueModifierDelNotify, func() any { return new(proto.RogueModifierDelNotify) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(RogueMagicUnitComposeCsReq, func() any { return new(proto.RogueMagicUnitComposeCsReq) })
	c.regMsg(RogueMagicScepterTakeOffUnitCsReq, func() any { return new(proto.RogueMagicScepterTakeOffUnitCsReq) })
	c.regMsg(RogueMagicReviveAvatarScRsp, func() any { return new(proto.RogueMagicReviveAvatarScRsp) })
	c.regMsg(RogueMagicAutoDressInUnitScRsp, func() any { return new(proto.RogueMagicAutoDressInUnitScRsp) })
	c.regMsg(FightEnterCsReq, func() any { return new(proto.FightEnterCsReq) })
	c.regMsg(RogueMagicEnterLayerScRsp, func() any { return new(proto.RogueMagicEnterLayerScRsp) })
	c.regMsg(RogueMagicEnterLayerCsReq, func() any { return new(proto.RogueMagicEnterLayerCsReq) })
	c.regMsg(RogueMagicLevelInfoUpdateScNotify, func() any { return new(proto.RogueMagicLevelInfoUpdateScNotify) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(RogueMagicSettleScRsp, func() any { return new(proto.RogueMagicSettleScRsp) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(RogueMagicStartCsReq, func() any { return new(proto.RogueMagicStartCsReq) })
	c.regMsg(CommonRogueQueryCsReq, func() any { return new(proto.CommonRogueQueryCsReq) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(SwordTrainingStartGameCsReq, func() any { return new(proto.SwordTrainingStartGameCsReq) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(RogueMagicEnableTalentScRsp, func() any { return new(proto.RogueMagicEnableTalentScRsp) })
	c.regMsg(RogueMagicQueryScRsp, func() any { return new(proto.RogueMagicQueryScRsp) })
	c.regMsg(MonopolyClickMbtiReportCsReq, func() any { return new(proto.MonopolyClickMbtiReportCsReq) })
	c.regMsg(EvolveBuildStartStageCsReq, func() any { return new(proto.EvolveBuildStartStageCsReq) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(RogueMagicBattleFailSettleInfoScNotify, func() any { return new(proto.RogueMagicBattleFailSettleInfoScNotify) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(FinishRogueCommonDialogueScRsp, func() any { return new(proto.FinishRogueCommonDialogueScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(CommonRogueComponentUpdateScNotify, func() any { return new(proto.CommonRogueComponentUpdateScNotify) })
	c.regMsg(SelectRogueCommonDialogueOptionCsReq, func() any { return new(proto.SelectRogueCommonDialogueOptionCsReq) })
	c.regMsg(CancelExpeditionScRsp, func() any { return new(proto.CancelExpeditionScRsp) })
	c.regMsg(RogueWorkbenchGetInfoCsReq, func() any { return new(proto.RogueWorkbenchGetInfoCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
	c.regMsg(RogueDebugMessageScNotify, func() any { return new(proto.RogueDebugMessageScNotify) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(GetFightFestDataCsReq, func() any { return new(proto.GetFightFestDataCsReq) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(ClockParkStartScriptScRsp, func() any { return new(proto.ClockParkStartScriptScRsp) })
	c.regMsg(RogueTournEnterRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournEnterRogueCocoonSceneCsReq) })
	c.regMsg(LobbyCreateScRsp, func() any { return new(proto.LobbyCreateScRsp) })
	c.regMsg(GetRogueExhibitionScRsp, func() any { return new(proto.GetRogueExhibitionScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(RogueMagicEnableTalentCsReq, func() any { return new(proto.RogueMagicEnableTalentCsReq) })
	c.regMsg(MatchThreeGetDataCsReq, func() any { return new(proto.MatchThreeGetDataCsReq) })
	c.regMsg(RogueDoGambleScRsp, func() any { return new(proto.RogueDoGambleScRsp) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(TakeMultipleActivityExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleActivityExpeditionRewardCsReq) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(ChessRogueNousEnableRogueTalentCsReq, func() any { return new(proto.ChessRogueNousEnableRogueTalentCsReq) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(GetRogueCommonDialogueDataScRsp, func() any { return new(proto.GetRogueCommonDialogueDataScRsp) })
	c.regMsg(GetRogueCollectionCsReq, func() any { return new(proto.GetRogueCollectionCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(RogueArcadeLeaveCsReq, func() any { return new(proto.RogueArcadeLeaveCsReq) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(RecallPetScRsp, func() any { return new(proto.RecallPetScRsp) })
	c.regMsg(RogueWorkbenchSelectFuncCsReq, func() any { return new(proto.RogueWorkbenchSelectFuncCsReq) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(GetRogueShopBuffInfoCsReq, func() any { return new(proto.GetRogueShopBuffInfoCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(GetRogueCollectionScRsp, func() any { return new(proto.GetRogueCollectionScRsp) })
	c.regMsg(TakeMultipleActivityExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleActivityExpeditionRewardScRsp) })
	c.regMsg(MatchResultScNotify, func() any { return new(proto.MatchResultScNotify) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(GetStageLineupCsReq, func() any { return new(proto.GetStageLineupCsReq) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(DifficultyAdjustmentUpdateDataScRsp, func() any { return new(proto.DifficultyAdjustmentUpdateDataScRsp) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(SetRogueCollectionCsReq, func() any { return new(proto.SetRogueCollectionCsReq) })
	c.regMsg(TakeRogueEventHandbookRewardScRsp, func() any { return new(proto.TakeRogueEventHandbookRewardScRsp) })
	c.regMsg(SyncRogueCommonDialogueDataScNotify, func() any { return new(proto.SyncRogueCommonDialogueDataScNotify) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(SummonPetCsReq, func() any { return new(proto.SummonPetCsReq) })
	c.regMsg(SpaceZooTakeCsReq, func() any { return new(proto.SpaceZooTakeCsReq) })
	c.regMsg(RogueDoGambleCsReq, func() any { return new(proto.RogueDoGambleCsReq) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(RogueArcadeRestartCsReq, func() any { return new(proto.RogueArcadeRestartCsReq) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(RogueArcadeStartCsReq, func() any { return new(proto.RogueArcadeStartCsReq) })
	c.regMsg(GetFriendChallengeDetailScRsp, func() any { return new(proto.GetFriendChallengeDetailScRsp) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(RogueArcadeRestartScRsp, func() any { return new(proto.RogueArcadeRestartScRsp) })
	c.regMsg(GetChallengeRecommendLineupListCsReq, func() any { return new(proto.GetChallengeRecommendLineupListCsReq) })
	c.regMsg(RogueArcadeStartScRsp, func() any { return new(proto.RogueArcadeStartScRsp) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(RogueTournGetArchiveRepositoryCsReq, func() any { return new(proto.RogueTournGetArchiveRepositoryCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(EvolveBuildReRandomStageCsReq, func() any { return new(proto.EvolveBuildReRandomStageCsReq) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(MatchBoxingClubOpponentCsReq, func() any { return new(proto.MatchBoxingClubOpponentCsReq) })
	c.regMsg(HeliobusStartRaidScRsp, func() any { return new(proto.HeliobusStartRaidScRsp) })
	c.regMsg(SyncRoguePickAvatarInfoScNotify, func() any { return new(proto.SyncRoguePickAvatarInfoScNotify) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(EvolveBuildQueryInfoScRsp, func() any { return new(proto.EvolveBuildQueryInfoScRsp) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(RogueModifierStageStartNotify, func() any { return new(proto.RogueModifierStageStartNotify) })
	c.regMsg(SwordTrainingRestoreGameCsReq, func() any { return new(proto.SwordTrainingRestoreGameCsReq) })
	c.regMsg(LobbyInviteScRsp, func() any { return new(proto.LobbyInviteScRsp) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(QuitTrackPhotoStageCsReq, func() any { return new(proto.QuitTrackPhotoStageCsReq) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(SubmitMaterialSubmitActivityMaterialCsReq, func() any { return new(proto.SubmitMaterialSubmitActivityMaterialCsReq) })
	c.regMsg(FightSessionStopScNotify, func() any { return new(proto.FightSessionStopScNotify) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(RogueArcadeGetInfoCsReq, func() any { return new(proto.RogueArcadeGetInfoCsReq) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(HeliobusLineupUpdateScNotify, func() any { return new(proto.HeliobusLineupUpdateScNotify) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(EvolveBuildShopAbilityUpScRsp, func() any { return new(proto.EvolveBuildShopAbilityUpScRsp) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(EvolveBuildShopAbilityUpCsReq, func() any { return new(proto.EvolveBuildShopAbilityUpCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(GetChallengeRecommendLineupListScRsp, func() any { return new(proto.GetChallengeRecommendLineupListScRsp) })
	c.regMsg(GetFightFestDataScRsp, func() any { return new(proto.GetFightFestDataScRsp) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(GetFriendBattleRecordDetailScRsp, func() any { return new(proto.GetFriendBattleRecordDetailScRsp) })
	c.regMsg(RaidCollectionDataScNotify, func() any { return new(proto.RaidCollectionDataScNotify) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(TakeFightActivityRewardCsReq, func() any { return new(proto.TakeFightActivityRewardCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(FightMatch3OpponentDataScNotify, func() any { return new(proto.FightMatch3OpponentDataScNotify) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(ActivityRaidPlacingGameScRsp, func() any { return new(proto.ActivityRaidPlacingGameScRsp) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(GetAllRedDotDataCsReq, func() any { return new(proto.GetAllRedDotDataCsReq) })
	c.regMsg(CommonRogueUpdateScNotify, func() any { return new(proto.CommonRogueUpdateScNotify) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(StartBoxingClubBattleCsReq, func() any { return new(proto.StartBoxingClubBattleCsReq) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(BuyRogueShopMiracleCsReq, func() any { return new(proto.BuyRogueShopMiracleCsReq) })
	c.regMsg(RogueTournEnterRoomScRsp, func() any { return new(proto.RogueTournEnterRoomScRsp) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(RogueTournRenameArchiveCsReq, func() any { return new(proto.RogueTournRenameArchiveCsReq) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(MarkAvatarCsReq, func() any { return new(proto.MarkAvatarCsReq) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(GetRogueShopMiracleInfoCsReq, func() any { return new(proto.GetRogueShopMiracleInfoCsReq) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(SetFriendMarkCsReq, func() any { return new(proto.SetFriendMarkCsReq) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(EvolveBuildShopAbilityDownScRsp, func() any { return new(proto.EvolveBuildShopAbilityDownScRsp) })
	c.regMsg(RemoveRotaterScRsp, func() any { return new(proto.RemoveRotaterScRsp) })
	c.regMsg(MonopolyLikeCsReq, func() any { return new(proto.MonopolyLikeCsReq) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(UpdateGunPlayDataCsReq, func() any { return new(proto.UpdateGunPlayDataCsReq) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(StartMatchCsReq, func() any { return new(proto.StartMatchCsReq) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(MonopolyGiveUpCurContentScRsp, func() any { return new(proto.MonopolyGiveUpCurContentScRsp) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(EvolveBuildGiveupScRsp, func() any { return new(proto.EvolveBuildGiveupScRsp) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(ContentPackageGetDataScRsp, func() any { return new(proto.ContentPackageGetDataScRsp) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(UnlockAvatarPathCsReq, func() any { return new(proto.UnlockAvatarPathCsReq) })
	c.regMsg(MonopolyGameSettleScNotify, func() any { return new(proto.MonopolyGameSettleScNotify) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(ClientObjUploadScRsp, func() any { return new(proto.ClientObjUploadScRsp) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(GetGameStateServiceConfigScRsp, func() any { return new(proto.GetGameStateServiceConfigScRsp) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(EvolveBuildReRandomStageScRsp, func() any { return new(proto.EvolveBuildReRandomStageScRsp) })
	c.regMsg(RestartChallengePhaseCsReq, func() any { return new(proto.RestartChallengePhaseCsReq) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(UpdatePsnSettingsInfoCsReq, func() any { return new(proto.UpdatePsnSettingsInfoCsReq) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(BoxingClubRewardScNotify, func() any { return new(proto.BoxingClubRewardScNotify) })
	c.regMsg(EvolveBuildStartStageScRsp, func() any { return new(proto.EvolveBuildStartStageScRsp) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(RogueTournBattleFailSettleInfoScNotify, func() any { return new(proto.RogueTournBattleFailSettleInfoScNotify) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(QueryProductInfoCsReq, func() any { return new(proto.QueryProductInfoCsReq) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(GetAuthkeyScRsp, func() any { return new(proto.GetAuthkeyScRsp) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(FightMatch3SwapCsReq, func() any { return new(proto.FightMatch3SwapCsReq) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(RogueTournStartCsReq, func() any { return new(proto.RogueTournStartCsReq) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(EvolveBuildShopAbilityResetCsReq, func() any { return new(proto.EvolveBuildShopAbilityResetCsReq) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(ChallengeBossPhaseSettleNotify, func() any { return new(proto.ChallengeBossPhaseSettleNotify) })
	c.regMsg(FightHeartBeatCsReq, func() any { return new(proto.FightHeartBeatCsReq) })
	c.regMsg(MonopolyGameBingoFlipCardCsReq, func() any { return new(proto.MonopolyGameBingoFlipCardCsReq) })
	c.regMsg(AceAntiCheaterCsReq, func() any { return new(proto.AceAntiCheaterCsReq) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(GetOfferingInfoCsReq, func() any { return new(proto.GetOfferingInfoCsReq) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(DailyFirstEnterMonopolyActivityScRsp, func() any { return new(proto.DailyFirstEnterMonopolyActivityScRsp) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(RogueMagicStoryInfoUpdateScNotify, func() any { return new(proto.RogueMagicStoryInfoUpdateScNotify) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(LobbyGetInfoCsReq, func() any { return new(proto.LobbyGetInfoCsReq) })
	c.regMsg(GetSwordTrainingDataScRsp, func() any { return new(proto.GetSwordTrainingDataScRsp) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(LobbyQuitCsReq, func() any { return new(proto.LobbyQuitCsReq) })
	c.regMsg(FightMatch3TurnEndScNotify, func() any { return new(proto.FightMatch3TurnEndScNotify) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(MonopolySttUpdateScNotify, func() any { return new(proto.MonopolySttUpdateScNotify) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(SubmitOfferingItemCsReq, func() any { return new(proto.SubmitOfferingItemCsReq) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(RogueMagicGetTalentInfoCsReq, func() any { return new(proto.RogueMagicGetTalentInfoCsReq) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(RogueTournQueryCsReq, func() any { return new(proto.RogueTournQueryCsReq) })
	c.regMsg(GetLevelRewardScRsp, func() any { return new(proto.GetLevelRewardScRsp) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoCsReq, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoCsReq) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(SetFriendMarkScRsp, func() any { return new(proto.SetFriendMarkScRsp) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(SetRogueExhibitionScRsp, func() any { return new(proto.SetRogueExhibitionScRsp) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(MultipleDropInfoScNotify, func() any { return new(proto.MultipleDropInfoScNotify) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(LobbyBeginScRsp, func() any { return new(proto.LobbyBeginScRsp) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(EnterChallengeNextPhaseScRsp, func() any { return new(proto.EnterChallengeNextPhaseScRsp) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(BattleLogReportScRsp, func() any { return new(proto.BattleLogReportScRsp) })
	c.regMsg(RogueTournEnterRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournEnterRogueCocoonSceneScRsp) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(GetPetDataScRsp, func() any { return new(proto.GetPetDataScRsp) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(GetMonopolyFriendRankingListCsReq, func() any { return new(proto.GetMonopolyFriendRankingListCsReq) })
	c.regMsg(CurPetChangedScNotify, func() any { return new(proto.CurPetChangedScNotify) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(AddEquipmentScNotify, func() any { return new(proto.AddEquipmentScNotify) })
	c.regMsg(DeleteRelicFilterPlanCsReq, func() any { return new(proto.DeleteRelicFilterPlanCsReq) })
	c.regMsg(AcceptedPamMissionExpireCsReq, func() any { return new(proto.AcceptedPamMissionExpireCsReq) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(EnterSummonActivityStageCsReq, func() any { return new(proto.EnterSummonActivityStageCsReq) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(TakeOfferingRewardCsReq, func() any { return new(proto.TakeOfferingRewardCsReq) })
	c.regMsg(OfferingInfoScNotify, func() any { return new(proto.OfferingInfoScNotify) })
	c.regMsg(DeleteRelicFilterPlanScRsp, func() any { return new(proto.DeleteRelicFilterPlanScRsp) })
	c.regMsg(MusicRhythmMaxDifficultyLevelsUnlockNotify, func() any { return new(proto.MusicRhythmMaxDifficultyLevelsUnlockNotify) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(MusicRhythmUnlockSongNotify, func() any { return new(proto.MusicRhythmUnlockSongNotify) })
	c.regMsg(MusicRhythmFinishLevelScRsp, func() any { return new(proto.MusicRhythmFinishLevelScRsp) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(FightEnterScRsp, func() any { return new(proto.FightEnterScRsp) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(ClockParkUnlockTalentCsReq, func() any { return new(proto.ClockParkUnlockTalentCsReq) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(DrinkMakerChallengeCsReq, func() any { return new(proto.DrinkMakerChallengeCsReq) })
	c.regMsg(MuseumTargetMissionFinishNotify, func() any { return new(proto.MuseumTargetMissionFinishNotify) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(SwordTrainingRestoreGameScRsp, func() any { return new(proto.SwordTrainingRestoreGameScRsp) })
	c.regMsg(GetUpdatedArchiveDataCsReq, func() any { return new(proto.GetUpdatedArchiveDataCsReq) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryScRsp, func() any { return new(proto.RogueTournGetArchiveRepositoryScRsp) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(ClockParkHandleWaitOperationScRsp, func() any { return new(proto.ClockParkHandleWaitOperationScRsp) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(EvolveBuildLeaveCsReq, func() any { return new(proto.EvolveBuildLeaveCsReq) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(MultipleDropInfoNotify, func() any { return new(proto.MultipleDropInfoNotify) })
	c.regMsg(FightFestScoreUpdateNotify, func() any { return new(proto.FightFestScoreUpdateNotify) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(GetMultipleDropInfoCsReq, func() any { return new(proto.GetMultipleDropInfoCsReq) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(MultiplayerGetFightGateCsReq, func() any { return new(proto.MultiplayerGetFightGateCsReq) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(RogueMagicEnterRoomScRsp, func() any { return new(proto.RogueMagicEnterRoomScRsp) })
	c.regMsg(ChooseBoxingClubResonanceCsReq, func() any { return new(proto.ChooseBoxingClubResonanceCsReq) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(MonopolyCheatDiceScRsp, func() any { return new(proto.MonopolyCheatDiceScRsp) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(RogueTournGetSettleInfoScRsp, func() any { return new(proto.RogueTournGetSettleInfoScRsp) })
	c.regMsg(GetFriendAssistListScRsp, func() any { return new(proto.GetFriendAssistListScRsp) })
	c.regMsg(RebattleByClientCsNotify, func() any { return new(proto.RebattleByClientCsNotify) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(ClockParkUseBuffCsReq, func() any { return new(proto.ClockParkUseBuffCsReq) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(GetMbtiReportCsReq, func() any { return new(proto.GetMbtiReportCsReq) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(MonopolyDailySettleScNotify, func() any { return new(proto.MonopolyDailySettleScNotify) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(SwordTrainingExamResultConfirmScRsp, func() any { return new(proto.SwordTrainingExamResultConfirmScRsp) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(DailyTaskDataScNotify, func() any { return new(proto.DailyTaskDataScNotify) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(RogueTournConfirmSettleCsReq, func() any { return new(proto.RogueTournConfirmSettleCsReq) })
	c.regMsg(ContentPackageUnlockScRsp, func() any { return new(proto.ContentPackageUnlockScRsp) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(SceneReviveAfterRebattleScRsp, func() any { return new(proto.SceneReviveAfterRebattleScRsp) })
	c.regMsg(RogueTournQueryScRsp, func() any { return new(proto.RogueTournQueryScRsp) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(NewAssistHistoryNotify, func() any { return new(proto.NewAssistHistoryNotify) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(LobbyModifyPlayerInfoCsReq, func() any { return new(proto.LobbyModifyPlayerInfoCsReq) })
	c.regMsg(MonopolyGuessChooseScRsp, func() any { return new(proto.MonopolyGuessChooseScRsp) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(StartFightFestCsReq, func() any { return new(proto.StartFightFestCsReq) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(GetBattleCollegeDataCsReq, func() any { return new(proto.GetBattleCollegeDataCsReq) })
	c.regMsg(EvolveBuildShopAbilityResetScRsp, func() any { return new(proto.EvolveBuildShopAbilityResetScRsp) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(RogueMagicEnterScRsp, func() any { return new(proto.RogueMagicEnterScRsp) })
	c.regMsg(MonopolyGameCreateScNotify, func() any { return new(proto.MonopolyGameCreateScNotify) })
	c.regMsg(FightMatch3SwapScRsp, func() any { return new(proto.FightMatch3SwapScRsp) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(StartFightFestScRsp, func() any { return new(proto.StartFightFestScRsp) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(RelicAvatarRecommendScRsp, func() any { return new(proto.RelicAvatarRecommendScRsp) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(ClockParkUseBuffScRsp, func() any { return new(proto.ClockParkUseBuffScRsp) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(FightFestUpdateCoinNotify, func() any { return new(proto.FightFestUpdateCoinNotify) })
	c.regMsg(GetGunPlayDataScRsp, func() any { return new(proto.GetGunPlayDataScRsp) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(SceneReviveAfterRebattleCsReq, func() any { return new(proto.SceneReviveAfterRebattleCsReq) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(GetFriendDevelopmentInfoScRsp, func() any { return new(proto.GetFriendDevelopmentInfoScRsp) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(AddRelicFilterPlanCsReq, func() any { return new(proto.AddRelicFilterPlanCsReq) })
	c.regMsg(SubmitOrigamiItemScRsp, func() any { return new(proto.SubmitOrigamiItemScRsp) })
	c.regMsg(GetMapRotationDataCsReq, func() any { return new(proto.GetMapRotationDataCsReq) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(FightMatch3DataScRsp, func() any { return new(proto.FightMatch3DataScRsp) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(CancelMatchCsReq, func() any { return new(proto.CancelMatchCsReq) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(RogueMagicSetAutoDressInMagicUnitCsReq, func() any { return new(proto.RogueMagicSetAutoDressInMagicUnitCsReq) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(RogueMagicAutoDressInUnitCsReq, func() any { return new(proto.RogueMagicAutoDressInUnitCsReq) })
	c.regMsg(InterruptMissionEventCsReq, func() any { return new(proto.InterruptMissionEventCsReq) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(LobbyCreateCsReq, func() any { return new(proto.LobbyCreateCsReq) })
	c.regMsg(EvolveBuildLeaveScRsp, func() any { return new(proto.EvolveBuildLeaveScRsp) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(RogueArcadeGetInfoScRsp, func() any { return new(proto.RogueArcadeGetInfoScRsp) })
	c.regMsg(GetSpringRecoverDataScRsp, func() any { return new(proto.GetSpringRecoverDataScRsp) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(RogueTournReviveAvatarScRsp, func() any { return new(proto.RogueTournReviveAvatarScRsp) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(MultiplayerFightGameFinishScNotify, func() any { return new(proto.MultiplayerFightGameFinishScNotify) })
	c.regMsg(UpdateMapRotationDataScNotify, func() any { return new(proto.UpdateMapRotationDataScNotify) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(EnterFightActivityStageScRsp, func() any { return new(proto.EnterFightActivityStageScRsp) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(ClockParkGetOngoingScriptInfoCsReq, func() any { return new(proto.ClockParkGetOngoingScriptInfoCsReq) })
	c.regMsg(TakeMaterialSubmitActivityRewardScRsp, func() any { return new(proto.TakeMaterialSubmitActivityRewardScRsp) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(FightFestUnlockSkillNotify, func() any { return new(proto.FightFestUnlockSkillNotify) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
}
