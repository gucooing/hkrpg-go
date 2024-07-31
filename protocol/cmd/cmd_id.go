package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	QuitBattleCsReq                                    = 191
	FinishQuestScRsp                                   = 962
	StartWolfBroGameScRsp                              = 6588
	PlayerReturnStartScNotify                          = 4561
	SelectPhoneThemeCsReq                              = 5173
	GetQuestDataCsReq                                  = 961
	GetLoginChatInfoCsReq                              = 3996
	DeleteSocialEventServerCacheScRsp                  = 7086
	SetPlayerInfoCsReq                                 = 17
	RogueTournReEnterRogueCocoonStageScRsp             = 6015
	ChallengeLineupNotify                              = 1774
	DeployRotaterCsReq                                 = 6839
	EnterAdventureScRsp                                = 1388
	UpgradeAreaCsReq                                   = 4347
	UpdateRedDotDataScRsp                              = 5920
	AetherDivideRefreshEndlessCsReq                    = 4868
	LobbyJoinCsReq                                     = 7370
	CurTrialActivityScNotify                           = 2689
	JoinLineupScRsp                                    = 773
	FinishTalkMissionScRsp                             = 1220
	GetSecretKeyInfoCsReq                              = 10
	MissionGroupWarnScNotify                           = 1274
	MonopolyGameBingoFlipCardScRsp                     = 7013
	UpgradeAreaStatScRsp                               = 4321
	RogueTournDeleteArchiveScRsp                       = 6096
	SetStuffToAreaScRsp                                = 4373
	LobbyModifyPlayerInfoCsReq                         = 7383
	LeaveRaidScRsp                                     = 2220
	UpdatePsnSettingsInfoCsReq                         = 63
	EnterAetherDivideSceneScRsp                        = 4888
	TakeOffRelicCsReq                                  = 303
	TakeAllRewardScRsp                                 = 3087
	ChessRogueQuitScRsp                                = 5575
	ChessRogueEnterNextLayerCsReq                      = 5429
	GetQuestDataScRsp                                  = 988
	LobbyJoinScRsp                                     = 7363
	PrestigeLevelUpCsReq                               = 4749
	SceneEntityMoveCsReq                               = 1461
	ClockParkHandleWaitOperationScRsp                  = 7274
	SetFriendRemarkNameCsReq                           = 2949
	HeliobusSnsCommentCsReq                            = 5887
	MonopolyGameRaiseRatioCsReq                        = 7050
	GetTutorialCsReq                                   = 1661
	RogueTournReviveAvatarCsReq                        = 6038
	StaminaInfoScNotify                                = 53
	GetSocialEventServerCacheCsReq                     = 7089
	SendMsgCsReq                                       = 3961
	SearchPlayerScRsp                                  = 2901
	ExpUpEquipmentScRsp                                = 574
	HeliobusSnsPostScRsp                               = 5873
	TakePromotionRewardCsReq                           = 356
	SetLineupNameCsReq                                 = 703
	StartBattleCollegeCsReq                            = 5720
	GetPlayerReturnMultiDropInfoScRsp                  = 4639
	EvolveBuildUnlockInfoNotify                        = 7132
	CommonRogueComponentUpdateScNotify                 = 5680
	SetHeadIconCsReq                                   = 2891
	ExchangeStaminaCsReq                               = 94
	EvolveBuildStartStageScRsp                         = 7128
	EnableRogueTalentScRsp                             = 1877
	HeliobusSnsLikeCsReq                               = 5830
	AddEquipmentScNotify                               = 537
	GetPrivateChatHistoryCsReq                         = 3939
	GetStuffScNotify                                   = 4387
	ChessRoguePickAvatarCsReq                          = 5558
	SyncEntityBuffChangeListScNotify                   = 1447
	GetPlatformPlayerInfoCsReq                         = 3000
	EvolveBuildShopAbilityDownScRsp                    = 7143
	QueryProductInfoCsReq                              = 40
	RogueGetGambleInfoCsReq                            = 5641
	ChessRogueQueryAeonDimensionsCsReq                 = 5512
	GetAssistHistoryScRsp                              = 2913
	SpaceZooMutateScRsp                                = 6773
	UnlockSkilltreeCsReq                               = 339
	LobbyGetInfoScRsp                                  = 7398
	MonopolyGuessChooseScRsp                           = 7100
	GetChallengeScRsp                                  = 1788
	ChessRogueConfirmRollCsReq                         = 5593
	HeliobusInfoChangedScNotify                        = 5874
	GetRecyleTimeCsReq                                 = 528
	PickRogueAvatarScRsp                               = 1882
	GetStageLineupCsReq                                = 761
	GetFriendListInfoScRsp                             = 2988
	ExpUpEquipmentCsReq                                = 562
	MonopolyQuizDurationChangeScNotify                 = 7005
	MonopolySttUpdateScNotify                          = 7063
	SceneEnterStageScRsp                               = 1481
	MakeDrinkCsReq                                     = 6987
	GetBasicInfoScRsp                                  = 18
	GetRollShopInfoCsReq                               = 6909
	GetMuseumInfoCsReq                                 = 4361
	GetMonopolyInfoCsReq                               = 7061
	DeleteBlacklistCsReq                               = 2912
	StartPartialChallengeCsReq                         = 1770
	MonopolyGuessBuyInformationScRsp                   = 7040
	ReturnLastTownScRsp                                = 1451
	GetUpdatedArchiveDataScRsp                         = 2320
	RogueTournExpNotify                                = 6068
	ChessRogueCheatRollCsReq                           = 5596
	ChessRogueLeaveScRsp                               = 5513
	MultiplayerFightGiveUpScRsp                        = 1073
	HeroBasicTypeChangedNotify                         = 93
	AlleyGuaranteedFundsCsReq                          = 4768
	DeleteFriendScRsp                                  = 2994
	FightTreasureDungeonMonsterScRsp                   = 4403
	RotateMapCsReq                                     = 6830
	EnterTrialActivityStageScRsp                       = 2644
	GetLevelRewardTakenListScRsp                       = 9
	GetReplayTokenCsReq                                = 3561
	TakeChallengeRewardScRsp                           = 1721
	UseItemScRsp                                       = 571
	AetherDivideLineupScNotify                         = 4855
	GetAuthkeyCsReq                                    = 21
	RogueTournGetSettleInfoScRsp                       = 6052
	GetSocialEventServerCacheScRsp                     = 7010
	RogueTournDifficultyCompNotify                     = 6079
	PromoteAvatarCsReq                                 = 330
	EvolveBuildShopAbilityResetCsReq                   = 7131
	GetSecretKeyInfoScRsp                              = 98
	GetGachaInfoCsReq                                  = 1961
	GetCurBattleInfoScRsp                              = 173
	CurAssistChangedNotify                             = 2917
	QuitLineupScRsp                                    = 771
	TakePromotionRewardScRsp                           = 349
	ClockParkStartScriptCsReq                          = 7230
	DeleteBlacklistScRsp                               = 2937
	TrialActivityDataChangeScNotify                    = 2646
	LeaveChallengeCsReq                                = 1739
	SecurityReportCsReq                                = 4162
	ChessRogueQueryBpScRsp                             = 5423
	AlleyTakeEventRewardScRsp                          = 4713
	SetIsDisplayAvatarInfoCsReq                        = 2830
	SetClientRaidTargetCountCsReq                      = 2247
	TakeRogueMiracleHandbookRewardScRsp                = 5700
	SelectChessRogueSubStoryCsReq                      = 5491
	ChessRogueRollDiceScRsp                            = 5508
	BuyRogueShopBuffCsReq                              = 5627
	SetClientPausedCsReq                               = 1417
	MonopolyAcceptQuizCsReq                            = 7085
	MonopolyGetRegionProgressScRsp                     = 7053
	PlayerLogoutScRsp                                  = 20
	InteractTreasureDungeonGridScRsp                   = 4456
	SetStuffToAreaCsReq                                = 4339
	AetherDivideTakeChallengeRewardScRsp               = 4885
	TakeChallengeRewardCsReq                           = 1796
	TakeAssistRewardCsReq                              = 2954
	LobbyCreateCsReq                                   = 7395
	WolfBroGameActivateBulletScRsp                     = 6582
	GetChallengeCsReq                                  = 1761
	BoxingClubChallengeUpdateScNotify                  = 4227
	RogueTournReviveAvatarScRsp                        = 6067
	LobbyInviteScRsp                                   = 7379
	ChessRogueUpdateDiceInfoScNotify                   = 5435
	TakeQuestRewardCsReq                               = 991
	EnhanceChessRogueBuffCsReq                         = 5544
	GameplayCounterRecoverCsReq                        = 1445
	MonopolyGiveUpCurContentScRsp                      = 7037
	MakeDrinkScRsp                                     = 6983
	GetPunkLordBattleRecordScRsp                       = 3268
	LogisticsGameScRsp                                 = 4720
	PromoteEquipmentScRsp                              = 520
	PlayerReturnPointChangeScNotify                    = 4520
	SetSignatureCsReq                                  = 2827
	SetTurnFoodSwitchCsReq                             = 526
	MonopolyMoveCsReq                                  = 7071
	PlayerLoginCsReq                                   = 61
	MonopolyEventLoadUpdateScNotify                    = 7029
	SubmitOfferingItemScRsp                            = 6923
	GetQuestRecordCsReq                                = 930
	GetBagCsReq                                        = 561
	TakeTrialActivityRewardCsReq                       = 2624
	GetPlayerReturnMultiDropInfoCsReq                  = 4620
	GetMonopolyFriendRankingListCsReq                  = 7008
	GetAetherDivideChallengeInfoCsReq                  = 4837
	FinishTutorialGuideCsReq                           = 1662
	TakeTrialActivityRewardScRsp                       = 2642
	ClientObjUploadCsReq                               = 60
	GetFriendChallengeLineupScRsp                      = 2946
	GetChallengeRaidInfoScRsp                          = 2230
	EnterTelevisionActivityStageCsReq                  = 6963
	RogueTournEnterRoomScRsp                           = 6099
	GetMailScRsp                                       = 888
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5649
	ChessRogueSelectBpScRsp                            = 5415
	ClockParkBattleEndScNotify                         = 7282
	ComposeSelectedRelicCsReq                          = 581
	InteractChargerCsReq                               = 6891
	SceneReviveAfterRebattleCsReq                      = 1412
	SetCurWaypointCsReq                                = 491
	GetMainMissionCustomValueCsReq                     = 1268
	GetStrongChallengeActivityDataCsReq                = 6661
	MonopolyEventSelectFriendCsReq                     = 7065
	PunkLordMonsterInfoScNotify                        = 3296
	SpaceZooBornCsReq                                  = 6791
	RaidInfoNotify                                     = 2239
	GetLoginActivityScRsp                              = 2688
	StartTimedFarmElementScRsp                         = 1419
	SwitchAetherDivideLineUpSlotCsReq                  = 4870
	PlayerReturnTakeRewardCsReq                        = 4530
	PlayBackGroundMusicScRsp                           = 3120
	GetTelevisionActivityDataScRsp                     = 6962
	RogueEndlessActivityBattleEndScNotify              = 6006
	EvolveBuildLeaveCsReq                              = 7127
	UpdateMovieRacingDataScRsp                         = 4181
	TakeOffAvatarSkinScRsp                             = 312
	SetBoxingClubResonanceLineupScRsp                  = 4294
	GetSaveRaidCsReq                                   = 2296
	MonopolyGameBingoFlipCardCsReq                     = 7058
	GetChessRogueNousStoryInfoScRsp                    = 5444
	RogueTournClearArchiveNameScNotify                 = 6081
	GetLoginActivityCsReq                              = 2661
	LeaveTrialActivityCsReq                            = 2657
	MonopolyGuessChooseCsReq                           = 7017
	MultiplayerFightGameStartScNotify                  = 1030
	PlayerHeartBeatScRsp                               = 16
	AcceptActivityExpeditionCsReq                      = 2527
	FightMatch3OpponentDataScNotify                    = 30171
	GetRogueInitialScoreCsReq                          = 1900
	HeliobusLineupUpdateScNotify                       = 5812
	ShowNewSupplementVisitorCsReq                      = 3762
	FinishPlotCsReq                                    = 1161
	LobbySyncInfoScNotify                              = 7390
	UnlockHeadIconScNotify                             = 2887
	CommonRogueUpdateScNotify                          = 5642
	ChessRogueNousEditDiceScRsp                        = 5550
	GetLevelRewardCsReq                                = 81
	ChooseBoxingClubResonanceCsReq                     = 4262
	FinishEmotionDialoguePerformanceScRsp              = 6371
	ClockParkGetOngoingScriptInfoScRsp                 = 7227
	EnterSectionScRsp                                  = 1401
	SetLanguageScRsp                                   = 33
	MuseumRandomEventSelectScRsp                       = 4309
	MonopolyCheatDiceScRsp                             = 7001
	GetMissionStatusScRsp                              = 1249
	SearchPlayerCsReq                                  = 2928
	GetSaveLogisticsMapCsReq                           = 4750
	ReserveStaminaExchangeScRsp                        = 72
	EvolveBuildShopAbilityUpScRsp                      = 7114
	GetFirstTalkByPerformanceNpcScRsp                  = 2174
	GetSaveLogisticsMapScRsp                           = 4732
	HandleFriendCsReq                                  = 2927
	RogueTournWeekChallengeUpdateScNotify              = 6044
	GetFeverTimeActivityDataScRsp                      = 7156
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3787
	SyncTaskScRsp                                      = 1230
	DestroyItemCsReq                                   = 532
	FightMatch3SwapCsReq                               = 30173
	RogueTournQueryScRsp                               = 6057
	MonopolyReRollRandomScRsp                          = 7003
	FinishCurTurnScRsp                                 = 4374
	MonopolyClickMbtiReportScRsp                       = 7038
	RemoveRotaterCsReq                                 = 6803
	FinishSectionIdScRsp                               = 2771
	FeverTimeActivityBattleEndScNotify                 = 7154
	AcceptExpeditionCsReq                              = 2591
	SceneEnterStageCsReq                               = 1409
	FinishItemIdCsReq                                  = 2739
	DoGachaInRollShopCsReq                             = 6907
	GetServerPrefsDataCsReq                            = 6191
	InterruptMissionEventCsReq                         = 1251
	CancelCacheNotifyCsReq                             = 4187
	HeartDialScriptChangeScNotify                      = 6387
	FinishPerformSectionIdCsReq                        = 2787
	ContentPackageUnlockCsReq                          = 7463
	GetRogueAeonInfoCsReq                              = 1899
	StartAetherDivideStageBattleCsReq                  = 4849
	SyncRogueVirtualItemInfoScNotify                   = 1823
	MonopolyGiveUpCurContentCsReq                      = 7012
	FightEnterScRsp                                    = 30088
	TakePrestigeRewardCsReq                            = 4762
	FinishRogueCommonDialogueCsReq                     = 5689
	DrinkMakerChallengeCsReq                           = 6992
	GetFriendApplyListInfoScRsp                        = 2973
	GetRogueCommonDialogueDataScRsp                    = 5665
	SyncServerSceneChangeNotify                        = 1425
	SelectPhoneThemeScRsp                              = 5130
	GetChessRogueStoryAeonTalkInfoScRsp                = 5465
	GetOfferingInfoScRsp                               = 6922
	RestoreWolfBroGameArchiveCsReq                     = 6539
	DelSaveRaidScNotify                                = 2270
	RefreshTriggerByClientScRsp                        = 1475
	MonopolyLikeScNotify                               = 7024
	ShareCsReq                                         = 4161
	ReplaceLineupCsReq                                 = 709
	GetPunkLordDataCsReq                               = 3221
	RogueTournGetPermanentTalentInfoCsReq              = 6066
	GetArchiveDataCsReq                                = 2361
	ReserveStaminaExchangeCsReq                        = 25
	ChessRogueRollDiceCsReq                            = 5563
	GetFarmStageGachaInfoScRsp                         = 1320
	ChessRogueGiveUpScRsp                              = 5479
	ChessRogueQuitCsReq                                = 5539
	HeliobusUnlockSkillScNotify                        = 5896
	HeliobusStartRaidCsReq                             = 5851
	GetRogueShopMiracleInfoScRsp                       = 5639
	GetFriendDevelopmentInfoScRsp                      = 2965
	ChessRogueNousDiceSurfaceUnlockNotify              = 5514
	GetJukeboxDataScRsp                                = 3188
	ContentPackageGetDataScRsp                         = 7467
	GetMainMissionCustomValueScRsp                     = 1295
	LogisticsScoreRewardSyncInfoScNotify               = 4726
	MarkItemScRsp                                      = 513
	GetMissionStatusCsReq                              = 1256
	SceneGroupRefreshScNotify                          = 1463
	GetFantasticStoryActivityDataScRsp                 = 4988
	ChallengeSettleNotify                              = 1730
	ExchangeGachaCeilingCsReq                          = 1930
	MonopolyRollRandomScRsp                            = 7021
	GroupStateChangeScRsp                              = 1435
	GetRogueBuffEnhanceInfoScRsp                       = 1881
	GetGunPlayDataScRsp                                = 4137
	GetTutorialScRsp                                   = 1688
	MonopolyReRollRandomCsReq                          = 7082
	RogueTournLevelInfoUpdateScNotify                  = 6013
	ChessRogueNousEnableRogueTalentCsReq               = 5457
	GetAlleyInfoCsReq                                  = 4761
	FinishChessRogueNousSubStoryScRsp                  = 5455
	SetGroupCustomSaveDataCsReq                        = 1416
	SelectChessRogueNousSubStoryScRsp                  = 5463
	PickRogueAvatarCsReq                               = 1821
	GetHeartDialInfoCsReq                              = 6361
	UpdateMechanismBarScNotify                         = 1442
	GameplayCounterRecoverScRsp                        = 1415
	GetSpringRecoverDataCsReq                          = 1490
	TakeMailAttachmentCsReq                            = 830
	SetAssistCsReq                                     = 2932
	GetRogueExhibitionScRsp                            = 5614
	CancelExpeditionScRsp                              = 2573
	ClockParkHandleWaitOperationCsReq                  = 7262
	EnterRogueMapRoomCsReq                             = 1826
	DiscardRelicCsReq                                  = 593
	SubmitOrigamiItemScRsp                             = 4121
	GetCurAssistCsReq                                  = 2968
	TravelBrochureApplyPasterListCsReq                 = 6449
	SubmitMonsterResearchActivityMaterialCsReq         = 2670
	CancelExpeditionCsReq                              = 2539
	SyncChessRogueMainStoryFinishScNotify              = 5486
	GetAllRedDotDataCsReq                              = 5961
	ChessRogueEnterCellCsReq                           = 5532
	FinishTalkMissionCsReq                             = 1291
	UpdateMapRotationDataScNotify                      = 6882
	StartWolfBroGameCsReq                              = 6561
	GetAetherDivideChallengeInfoScRsp                  = 4828
	MultiplayerFightGameFinishScNotify                 = 1071
	GetMonopolyInfoScRsp                               = 7088
	TakeChapterRewardCsReq                             = 471
	FinishChessRogueSubStoryScRsp                      = 5592
	SetHeadIconScRsp                                   = 2820
	ClockParkUnlockTalentCsReq                         = 7239
	EnhanceRogueBuffScRsp                              = 1837
	GetReplayTokenScRsp                                = 3588
	EnterTrialActivityStageCsReq                       = 2611
	GetFightActivityDataCsReq                          = 3661
	ExchangeRogueRewardKeyScRsp                        = 1816
	ReplaceLineupScRsp                                 = 781
	MultiplayerGetFightGateCsReq                       = 1091
	ChessRogueNousEnableRogueTalentScRsp               = 5599
	SceneEntityMoveScNotify                            = 1462
	StartBattleCollegeScRsp                            = 5739
	StartAetherDivideSceneBattleCsReq                  = 4839
	GetRogueAdventureRoomInfoScRsp                     = 5696
	RogueTournStartScRsp                               = 6040
	ScenePlaneEventScNotify                            = 1479
	GetHeroBasicTypeInfoScRsp                          = 95
	GetTreasureDungeonActivityDataScRsp                = 4474
	ClockParkFinishScriptScNotify                      = 7249
	SelectChessRogueNousSubStoryCsReq                  = 5495
	TakeLoginActivityRewardScRsp                       = 2620
	BuyGoodsCsReq                                      = 1591
	GetUpdatedArchiveDataCsReq                         = 2391
	StartTrialActivityCsReq                            = 2616
	NewMailScNotify                                    = 887
	GetMailCsReq                                       = 861
	GetAetherDivideInfoScRsp                           = 4874
	ChangeScriptEmotionCsReq                           = 6391
	ReBattleAfterBattleLoseCsNotify                    = 147
	BattleLogReportScRsp                               = 162
	AlleyShipmentEventEffectsScNotify                  = 4733
	SyncRogueExploreWinScNotify                        = 1858
	GetChessRogueStoryAeonTalkInfoCsReq                = 5502
	DressRelicAvatarCsReq                              = 321
	BuyBpLevelCsReq                                    = 3073
	RogueTournGetMiscRealTimeDataCsReq                 = 6076
	HandleFriendScRsp                                  = 2962
	HeartDialTraceScriptScRsp                          = 6362
	MonopolyRollDiceCsReq                              = 7073
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6010
	FinishSectionIdCsReq                               = 2730
	GetPunkLordDataScRsp                               = 3282
	ClockParkGetInfoCsReq                              = 7261
	LogisticsDetonateStarSkiffCsReq                    = 4785
	ChallengeBossPhaseSettleNotify                     = 1737
	SetHeroBasicTypeCsReq                              = 32
	QuitRogueCsReq                                     = 1855
	GetExhibitScNotify                                 = 4327
	GetMapRotationDataCsReq                            = 6862
	TakeRogueAeonLevelRewardCsReq                      = 1869
	GetMonopolyFriendRankingListScRsp                  = 7046
	GetChessRogueBuffEnhanceInfoCsReq                  = 5414
	SaveLogisticsScRsp                                 = 4728
	GetRogueTalentInfoCsReq                            = 1836
	UpdatePlayerSettingCsReq                           = 6
	MuseumRandomEventQueryCsReq                        = 4356
	GetActivityScheduleConfigCsReq                     = 2639
	EvolveBuildLeaveScRsp                              = 7106
	GiveUpBoxingClubChallengeScRsp                     = 4271
	TakeQuestOptionalRewardScRsp                       = 947
	RogueTournEnterRogueCocoonSceneScRsp               = 6027
	UpgradeAreaStatCsReq                               = 4396
	TravelBrochureUpdatePasterPosScRsp                 = 6474
	MuseumRandomEventQueryScRsp                        = 4349
	GetCurBattleInfoCsReq                              = 139
	MonopolyGetRafflePoolInfoCsReq                     = 7006
	MatchThreeSetBirdPosCsReq                          = 7428
	RogueModifierUpdateNotify                          = 5371
	GetChatEmojiListCsReq                              = 3987
	ActivateFarmElementCsReq                           = 1405
	GetLineupAvatarDataScRsp                           = 747
	GetRogueTalentInfoScRsp                            = 1875
	CancelMatchScRsp                                   = 7313
	ClearAetherDividePassiveSkillScRsp                 = 4803
	MonopolyGameSettleScNotify                         = 7055
	ChessRogueLeaveCsReq                               = 5545
	ComposeItemCsReq                                   = 547
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6009
	SetLineupNameScRsp                                 = 770
	UpdateFeatureSwitchScNotify                        = 66
	EquipAetherDividePassiveSkillScRsp                 = 4821
	RogueTournEnterCsReq                               = 6037
	UpgradeAreaScRsp                                   = 4394
	StartMatchCsReq                                    = 7345
	PlayerKickOutScNotify                              = 87
	GetRogueExhibitionCsReq                            = 5606
	MonopolyGetDailyInitItemScRsp                      = 7002
	MarkReadMailScRsp                                  = 820
	LogisticsInfoScNotify                              = 4701
	ClientObjUploadScRsp                               = 29
	DressAvatarCsReq                                   = 387
	FightHeartBeatScRsp                                = 30073
	EvolveBuildStartLevelScRsp                         = 7113
	LobbyGetInfoCsReq                                  = 7364
	AlleyFundsScNotify                                 = 4709
	ShareScRsp                                         = 4188
	HeliobusChallengeUpdateScNotify                    = 5881
	SpaceZooTakeCsReq                                  = 6794
	SceneCastSkillCostMpCsReq                          = 1494
	SetDisplayAvatarCsReq                              = 2839
	TakeMonsterResearchActivityRewardCsReq             = 2649
	InteractPropCsReq                                  = 1491
	MonopolyActionResultScNotify                       = 7091
	ChessRogueConfirmRollScRsp                         = 5555
	RestartChallengePhaseCsReq                         = 1751
	SyncChessRogueNousSubStoryScNotify                 = 5526
	ChessRogueUpdateUnlockLevelScNotify                = 5454
	DrinkMakerChallengeScRsp                           = 6996
	ChessRogueGiveUpCsReq                              = 5481
	TravelBrochureSelectMessageCsReq                   = 6439
	EvolveBuildQueryInfoScRsp                          = 7117
	LobbyQuitCsReq                                     = 7387
	AlleyShopLevelScNotify                             = 4781
	GetRogueAdventureRoomInfoCsReq                     = 5694
	StartChallengeScRsp                                = 1720
	GetEnteredSceneScRsp                               = 1452
	HeliobusEnterBattleCsReq                           = 5856
	PlayerGetTokenCsReq                                = 39
	MonopolyGameGachaScRsp                             = 7095
	MonopolyClickCellCsReq                             = 7099
	SyncRogueAeonScNotify                              = 1883
	ChooseBoxingClubStageOptionalBuffCsReq             = 4296
	RogueTournTakeExpRewardCsReq                       = 6078
	EnterChessRogueAeonRoomCsReq                       = 5433
	TakeBpRewardCsReq                                  = 3020
	GetFriendChallengeDetailCsReq                      = 2911
	TakeExpeditionRewardCsReq                          = 2530
	SharePunkLordMonsterCsReq                          = 3239
	SharePunkLordMonsterScRsp                          = 3273
	TakeRollShopRewardScRsp                            = 6914
	MultiplayerGetFightGateScRsp                       = 1020
	ChessRogueStartCsReq                               = 5452
	ChessRogueEnterScRsp                               = 5424
	ChangeLineupLeaderCsReq                            = 794
	RotateMapScRsp                                     = 6871
	TakeFightActivityRewardScRsp                       = 3630
	RevcMsgScNotify                                    = 3991
	ChessRogueSkipTeachingLevelScRsp                   = 5581
	TravelBrochureApplyPasterCsReq                     = 6430
	QuitRogueScRsp                                     = 1868
	GetFriendApplyListInfoCsReq                        = 2939
	GetMarkItemListCsReq                               = 568
	TrainVisitorBehaviorFinishCsReq                    = 3761
	ClockParkStartScriptScRsp                          = 7271
	AceAntiCheaterScRsp                                = 11
	GetMonsterResearchActivityDataScRsp                = 2603
	GetGunPlayDataCsReq                                = 4112
	ExchangeRogueRewardKeyCsReq                        = 1842
	UnlockSkilltreeScRsp                               = 373
	UnlockedAreaMapScNotify                            = 1459
	HeliobusUpgradeLevelScRsp                          = 5894
	UpdateRogueAdventureRoomScoreCsReq                 = 5666
	AlleyTakeEventRewardCsReq                          = 4758
	GetAvatarDataScRsp                                 = 388
	UpdateMovieRacingDataCsReq                         = 4109
	FinishAeonDialogueGroupScRsp                       = 1802
	ClockParkUseBuffScRsp                              = 7256
	DeleteFriendCsReq                                  = 2947
	ExchangeGachaCeilingScRsp                          = 1971
	QuitWolfBroGameScRsp                               = 6571
	GetBoxingClubInfoCsReq                             = 4261
	SceneEntityMoveScRsp                               = 1488
	BattleLogReportCsReq                               = 127
	FightActivityDataChangeScNotify                    = 3691
	GetFirstTalkNpcCsReq                               = 2139
	SetGenderScRsp                                     = 26
	ArchiveWolfBroGameScRsp                            = 6520
	GetQuestRecordScRsp                                = 971
	MonopolyTakeRaffleTicketRewardCsReq                = 7079
	UpdateTrackMainMissionIdCsReq                      = 1285
	ClockParkUseBuffCsReq                              = 7270
	StarFightDataChangeNotify                          = 7166
	TrainRefreshTimeNotify                             = 3739
	GetFriendListInfoCsReq                             = 2961
	RankUpEquipmentCsReq                               = 587
	TakeApRewardCsReq                                  = 3361
	SetClientPausedScRsp                               = 1500
	HealPoolInfoNotify                                 = 1411
	SyncRogueCommonDialogueOptionFinishScNotify        = 5686
	MonopolyGetRaffleTicketScRsp                       = 7083
	EnableRogueTalentCsReq                             = 1838
	RaidCollectionDataScRsp                            = 6942
	SyncRogueMapRoomScNotify                           = 1840
	SyncTaskCsReq                                      = 1273
	TakeChallengeRaidRewardScRsp                       = 2287
	SetTurnFoodSwitchScRsp                             = 517
	SetFriendMarkCsReq                                 = 2990
	EnterFightActivityStageScRsp                       = 3639
	MonopolyLikeCsReq                                  = 7011
	MonopolyMoveScRsp                                  = 7087
	GetFeverTimeActivityDataCsReq                      = 7155
	GetChallengeGroupStatisticsCsReq                   = 1782
	EnterSceneScRsp                                    = 1478
	LastSpringRefreshTimeNotify                        = 1456
	GetCrossInfoScRsp                                  = 7333
	SetCurInteractEntityScRsp                          = 1455
	EvolveBuildShopAbilityResetScRsp                   = 7108
	GiveUpBoxingClubChallengeCsReq                     = 4230
	SyncClientResVersionCsReq                          = 130
	SelectRogueCommonDialogueOptionScRsp               = 5634
	AcceptMissionEventCsReq                            = 1203
	RogueTournGetArchiveRepositoryScRsp                = 6036
	GetTutorialGuideScRsp                              = 1620
	ChessRogueSkipTeachingLevelCsReq                   = 5440
	StartStarFightLevelCsReq                           = 7164
	ReviveRogueAvatarScRsp                             = 1856
	GetFriendChallengeDetailScRsp                      = 2944
	EnterSceneCsReq                                    = 1480
	StartPunkLordRaidCsReq                             = 3291
	SecurityReportScRsp                                = 4174
	GetHeartDialInfoScRsp                              = 6388
	MarkChatEmojiScRsp                                 = 3974
	RogueTournEnterLayerScRsp                          = 6072
	PrepareRogueAdventureRoomCsReq                     = 5688
	GetFriendAssistListScRsp                           = 2964
	GetMissionEventDataCsReq                           = 1296
	GetPlayerBoardDataCsReq                            = 2861
	ChessRogueUpdateAllowedSelectCellScNotify          = 5566
	FinishPlotScRsp                                    = 1188
	GetSingleRedDotParamGroupCsReq                     = 5939
	SetRogueExhibitionScRsp                            = 5660
	LobbyQuitScRsp                                     = 7359
	MonopolyGetRaffleTicketCsReq                       = 7078
	RemoveStuffFromAreaCsReq                           = 4330
	SyncRogueSeasonFinishScNotify                      = 1813
	EnterStrongChallengeActivityStageScRsp             = 6620
	QuitTreasureDungeonCsReq                           = 4409
	ReviveRogueAvatarCsReq                             = 1870
	SetNicknameCsReq                                   = 56
	GetMonopolyDailyReportCsReq                        = 7041
	GetFriendBattleRecordDetailScRsp                   = 2942
	SetPlayerInfoScRsp                                 = 100
	TakeOfferingRewardCsReq                            = 6939
	SpaceZooOpCatteryScRsp                             = 6771
	MonopolyScrachRaffleTicketScRsp                    = 7025
	RefreshTriggerByClientCsReq                        = 1436
	SceneCastSkillCostMpScRsp                          = 1496
	MuseumTargetStartNotify                            = 4312
	GetLevelRewardScRsp                                = 12
	TakeMonsterResearchActivityRewardScRsp             = 2651
	GetMultipleDropInfoCsReq                           = 4661
	LockRelicScRsp                                     = 503
	SwapLineupCsReq                                    = 787
	GetNpcStatusScRsp                                  = 2720
	TravelBrochureSelectMessageScRsp                   = 6473
	GetNpcMessageGroupCsReq                            = 2761
	GetNpcTakenRewardCsReq                             = 2161
	SyncChessRogueNousValueScNotify                    = 5469
	SetRogueCollectionCsReq                            = 5672
	AetherDivideTakeChallengeRewardCsReq               = 4813
	RegionStopScNotify                                 = 3
	AlleyEventEffectNotify                             = 4727
	FightMatch3SwapScRsp                               = 30130
	EnterMapRotationRegionCsReq                        = 6861
	SetMissionEventProgressCsReq                       = 1281
	GetDrinkMakerDataCsReq                             = 6989
	TakeMailAttachmentScRsp                            = 871
	StartRogueScRsp                                    = 1820
	InterruptMissionEventScRsp                         = 1209
	ChessRogueSelectCellScRsp                          = 5442
	QuitTreasureDungeonScRsp                           = 4481
	ContentPackageGetDataCsReq                         = 7495
	AlleyShipUsedCountScNotify                         = 4755
	TakeApRewardScRsp                                  = 3388
	RemoveRotaterScRsp                                 = 6870
	PlayBackGroundMusicCsReq                           = 3191
	ChessRogueGiveUpRollCsReq                          = 5560
	GetMissionDataScRsp                                = 1288
	VirtualLineupDestroyNotify                         = 751
	SyncRogueAreaUnlockScNotify                        = 1879
	SubmitEmotionItemScRsp                             = 6373
	TakeOffAvatarSkinCsReq                             = 381
	MonopolyConfirmRandomCsReq                         = 7070
	ChessRogueQuestFinishNotify                        = 5571
	MonopolyTakeRaffleTicketRewardScRsp                = 7084
	GetRogueScoreRewardInfoScRsp                       = 1860
	TrialBackGroundMusicCsReq                          = 3130
	ReEnterLastElementStageCsReq                       = 1489
	SubmitOfferingItemCsReq                            = 6927
	GmTalkCsReq                                        = 27
	PunkLordBattleResultScNotify                       = 3209
	LobbyKickOutCsReq                                  = 7377
	GetCurAssistScRsp                                  = 2995
	EndDrinkMakerSequenceCsReq                         = 6999
	GetAllRedDotDataScRsp                              = 5988
	SetRogueCollectionScRsp                            = 5653
	WolfBroGameActivateBulletCsReq                     = 6521
	DoGachaCsReq                                       = 1991
	SetClientRaidTargetCountScRsp                      = 2294
	QuitWolfBroGameCsReq                               = 6530
	WolfBroGameUseBulletCsReq                          = 6574
	GetSingleRedDotParamGroupScRsp                     = 5973
	ChessRogueUpdateBoardScNotify                      = 5488
	WolfBroGamePickupBulletCsReq                       = 6594
	StartPunkLordRaidScRsp                             = 3220
	GetWolfBroGameDataCsReq                            = 6587
	RestoreWolfBroGameArchiveScRsp                     = 6573
	MonopolySocialEventEffectScNotify                  = 7034
	ChessRogueUpdateReviveInfoScNotify                 = 5402
	FinishRogueCommonDialogueScRsp                     = 5610
	SetCurInteractEntityCsReq                          = 1432
	GetRaidInfoScRsp                                   = 2274
	MuseumFundsChangedScNotify                         = 4303
	GetWolfBroGameDataScRsp                            = 6527
	ChessRoguePickAvatarScRsp                          = 5520
	WolfBroGameExplodeMonsterScRsp                     = 6570
	DoGachaScRsp                                       = 1920
	FinishTutorialGuideScRsp                           = 1674
	GetChatFriendHistoryScRsp                          = 3971
	FinishFirstTalkNpcScRsp                            = 2171
	WolfBroGameUseBulletScRsp                          = 6547
	DeleteSummonUnitScRsp                              = 1404
	WaypointShowNewCsNotify                            = 430
	GetWaypointCsReq                                   = 461
	GetChapterScRsp                                    = 473
	HeliobusSelectSkillScRsp                           = 5882
	MatchThreeLevelEndScRsp                            = 7413
	SetCurWaypointScRsp                                = 420
	EntityBindPropCsReq                                = 1454
	GetWaypointScRsp                                   = 488
	GetChapterCsReq                                    = 439
	RogueTournGetAllArchiveCsReq                       = 6054
	TakeChapterRewardScRsp                             = 487
	TreasureDungeonFinishScNotify                      = 4488
	TakeOfferingRewardScRsp                            = 6934
	ChessRogueSelectBpCsReq                            = 5450
	FinishTutorialCsReq                                = 1687
	ChessRogueGoAheadScRsp                             = 5574
	MonopolyGameGachaCsReq                             = 7068
	GetTutorialGuideCsReq                              = 1691
	MonopolyGameRaiseRatioScRsp                        = 7032
	UnlockTutorialGuideCsReq                           = 1630
	ChessRogueUpdateLevelBaseInfoScNotify              = 5499
	GameplayCounterUpdateScNotify                      = 1429
	PlayerLoginFinishScRsp                             = 80
	GetChallengeRaidInfoCsReq                          = 2273
	RestartChallengePhaseScRsp                         = 1709
	BattlePassInfoNotify                               = 3061
	MonopolyGetRegionProgressCsReq                     = 7072
	UnlockTutorialCsReq                                = 1639
	DailyFirstEnterMonopolyActivityScRsp               = 7094
	GetKilledPunkLordMonsterDataScRsp                  = 3212
	FightTreasureDungeonMonsterCsReq                   = 4482
	UpdateRotaterScNotify                              = 6856
	RogueTournGetAllArchiveScRsp                       = 6022
	MonopolySelectOptionScRsp                          = 7062
	MuseumTakeCollectRewardScRsp                       = 4333
	GetBattleCollegeDataScRsp                          = 5788
	EnterRogueEndlessActivityStageCsReq                = 6004
	StartBoxingClubBattleScRsp                         = 4273
	EnterTreasureDungeonScRsp                          = 4494
	PlayerLogoutCsReq                                  = 91
	MissionRewardScNotify                              = 1239
	StartAetherDivideChallengeBattleCsReq              = 4830
	RelicRecommendCsReq                                = 548
	OpenTreasureDungeonGridCsReq                       = 4496
	AetherDivideRefreshEndlessScNotify                 = 4858
	GetAllServerPrefsDataCsReq                         = 6161
	InteractTreasureDungeonGridCsReq                   = 4470
	UseTreasureDungeonItemScRsp                        = 4451
	GetKilledPunkLordMonsterDataCsReq                  = 3281
	MultiplayerFightGameStateScRsp                     = 1088
	GetTreasureDungeonActivityDataCsReq                = 4462
	SpaceZooDataScRsp                                  = 6788
	UseTreasureDungeonItemCsReq                        = 4449
	OpenTreasureDungeonGridScRsp                       = 4421
	TravelBrochureApplyPasterListScRsp                 = 6451
	GetPunkLordMonsterDataScRsp                        = 3288
	GetExpeditionDataScRsp                             = 2588
	MonopolyContentUpdateScNotify                      = 7033
	TravelBrochurePageResetCsReq                       = 6470
	TravelBrochureRemovePasterScRsp                    = 6427
	ExpUpRelicScRsp                                    = 521
	FinishPerformSectionIdScRsp                        = 2727
	GetRogueInfoScRsp                                  = 1888
	AcceptMainMissionScRsp                             = 1255
	TravelBrochureRemovePasterCsReq                    = 6487
	MonopolyConditionUpdateScNotify                    = 7036
	TeleportToMissionResetPointCsReq                   = 1228
	GetNpcStatusCsReq                                  = 2791
	StartRogueCsReq                                    = 1891
	TravelBrochureSetPageDescStatusCsReq               = 6482
	FinishChapterScNotify                              = 4991
	RankUpAvatarScRsp                                  = 396
	SetSpringRecoverConfigScRsp                        = 1464
	MonopolyScrachRaffleTicketCsReq                    = 7069
	TravelBrochureSetPageDescStatusScRsp               = 6403
	PrestigeLevelUpScRsp                               = 4751
	TravelBrochureGetPasterScNotify                    = 6447
	EquipAetherDividePassiveSkillCsReq                 = 4896
	DeactivateFarmElementScRsp                         = 1448
	ChessRogueReRollDiceScRsp                          = 5416
	GetFirstTalkNpcScRsp                               = 2173
	CancelCacheNotifyScRsp                             = 4127
	EnterFeverTimeActivityStageCsReq                   = 7158
	PrepareRogueAdventureRoomScRsp                     = 5691
	TravelBrochurePageUnlockScNotify                   = 6491
	TakeCityShopRewardCsReq                            = 1539
	GetTrainVisitorRegisterScRsp                       = 3771
	SyncRogueFinishScNotify                            = 1896
	SetRedPointStatusScNotify                          = 79
	ChessRogueReviveAvatarScRsp                        = 5522
	GetVideoVersionKeyCsReq                            = 78
	GetTrainVisitorBehaviorCsReq                       = 3791
	SubmitEmotionItemCsReq                             = 6339
	TrainVisitorBehaviorFinishScRsp                    = 3788
	GetGachaCeilingScRsp                               = 1973
	EvolveBuildReRandomStageCsReq                      = 7105
	GetTrainVisitorRegisterCsReq                       = 3730
	GetCurLineupDataScRsp                              = 720
	TrainVisitorRewardSendNotify                       = 3773
	DailyFirstMeetPamCsReq                             = 3491
	RankUpEquipmentScRsp                               = 527
	GetAssistHistoryCsReq                              = 2958
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3727
	AddBlacklistCsReq                                  = 2921
	TextJoinSaveScRsp                                  = 3888
	TextJoinBatchSaveCsReq                             = 3839
	TakeLoginActivityRewardCsReq                       = 2691
	TextJoinQueryScRsp                                 = 3820
	TakeRogueScoreRewardScRsp                          = 1851
	LockEquipmentCsReq                                 = 539
	TextJoinBatchSaveScRsp                             = 3873
	SyncHandleFriendScNotify                           = 2974
	TextJoinSaveCsReq                                  = 3861
	TakeRogueMiracleHandbookRewardCsReq                = 5617
	GetTelevisionActivityDataCsReq                     = 6969
	MatchThreeLevelEndCsReq                            = 7420
	AlleyEventChangeNotify                             = 4787
	LogisticsGameCsReq                                 = 4791
	DeleteSocialEventServerCacheCsReq                  = 7098
	EnterTelevisionActivityStageScRsp                  = 6979
	QuitBattleScRsp                                    = 120
	TelevisionActivityBattleEndScNotify                = 6974
	SpaceZooTakeScRsp                                  = 6796
	TeleportToMissionResetPointScRsp                   = 1201
	TakeOffRelicScRsp                                  = 370
	AlleyGuaranteedFundsScRsp                          = 4795
	TelevisionActivityDataChangeScNotify               = 6967
	FinishFirstTalkByPerformanceNpcScRsp               = 2194
	ContentPackageUnlockScRsp                          = 7474
	SetHeroBasicTypeScRsp                              = 55
	SelectInclinationTextCsReq                         = 2187
	PunkLordDataChangeNotify                           = 3232
	RogueTournDeleteArchiveCsReq                       = 6086
	EvolveBuildStartLevelCsReq                         = 7120
	ChessRogueReRollDiceCsReq                          = 5460
	DeactivateFarmElementCsReq                         = 1440
	AlleyOrderChangedScNotify                          = 4770
	FinishFirstTalkNpcCsReq                            = 2130
	TakeTalkRewardCsReq                                = 2191
	GetTrialActivityDataCsReq                          = 2664
	MonopolyDailySettleScNotify                        = 7064
	PlayerSyncScNotify                                 = 661
	FightMatch3DataCsReq                               = 30161
	DelMailCsReq                                       = 839
	WolfBroGamePickupBulletScRsp                       = 6596
	ChessRogueUpdateMoneyInfoScNotify                  = 5564
	RogueTournLeaveRogueCocoonSceneScRsp               = 6056
	MissionAcceptScNotify                              = 1258
	SaveLogisticsCsReq                                 = 4737
	FinishFirstTalkByPerformanceNpcCsReq               = 2147
	TakeTalkRewardScRsp                                = 2120
	SelectInclinationTextScRsp                         = 2127
	GetStrongChallengeActivityDataScRsp                = 6688
	RogueTournConfirmSettleScRsp                       = 6074
	StrongChallengeActivityBattleEndScNotify           = 6639
	StoryLineTrialAvatarChangeScNotify                 = 6230
	GetStoryLineInfoCsReq                              = 6261
	AetherDivideSkillItemScNotify                      = 4850
	SpaceZooBornScRsp                                  = 6720
	RogueTournLeaveCsReq                               = 6064
	SwapLineupScRsp                                    = 727
	SyncTurnFoodNotify                                 = 554
	GetAssistListCsReq                                 = 2933
	ChangeStoryLineFinishScNotify                      = 6273
	GetStoryLineInfoScRsp                              = 6288
	GetStarFightDataCsReq                              = 7168
	SetFriendRemarkNameScRsp                           = 2951
	ExchangeHcoinCsReq                                 = 551
	MonopolyGetDailyInitItemCsReq                      = 7052
	GetStarFightDataScRsp                              = 7161
	DeployRotaterScRsp                                 = 6873
	MonopolyGuessDrawScNotify                          = 7048
	SpaceZooExchangeItemCsReq                          = 6774
	SpaceZooCatUpdateNotify                            = 6762
	SetAetherDivideLineUpScRsp                         = 4894
	RefreshAlleyOrderCsReq                             = 4782
	GetBattleCollegeDataCsReq                          = 5761
	SellItemCsReq                                      = 570
	StartTimedCocoonStageScRsp                         = 1497
	SpaceZooDeleteCatScRsp                             = 6727
	DrinkMakerDayEndScNotify                           = 6991
	SpaceZooOpCatteryCsReq                             = 6730
	ApplyFriendCsReq                                   = 2930
	ChessRogueCellUpdateNotify                         = 5572
	GetPlayerDetailInfoScRsp                           = 2920
	SpaceZooMutateCsReq                                = 6739
	StartTimedCocoonStageCsReq                         = 1431
	ComposeItemScRsp                                   = 594
	MonopolyGuessBuyInformationCsReq                   = 7093
	HeliobusStartRaidScRsp                             = 5809
	GetShopListCsReq                                   = 1561
	GetUnlockTeleportCsReq                             = 1472
	GetAssistListScRsp                                 = 2950
	InteractChargerScRsp                               = 6820
	TravelBrochureUpdatePasterPosCsReq                 = 6462
	TakeCityShopRewardScRsp                            = 1573
	SpringRefreshScRsp                                 = 1470
	CityShopInfoScNotify                               = 1530
	GetPlatformPlayerInfoScRsp                         = 2993
	GetServerPrefsDataScRsp                            = 6120
	UpdateServerPrefsDataCsReq                         = 6139
	MatchThreeSetBirdPosScRsp                          = 7433
	EvolveBuildShopAbilityUpCsReq                      = 7107
	GetAllServerPrefsDataScRsp                         = 6188
	EnterTreasureDungeonCsReq                          = 4447
	MuseumTakeCollectRewardCsReq                       = 4301
	UpdateServerPrefsDataScRsp                         = 6173
	MatchBoxingClubOpponentScRsp                       = 4220
	MatchResultScNotify                                = 7324
	ExtraLineupDestroyNotify                           = 712
	EnterSceneByServerScNotify                         = 1483
	SceneUpdatePositionVersionNotify                   = 1474
	RogueModifierSelectCellScRsp                       = 5339
	GetFriendRecommendListInfoCsReq                    = 2970
	SpringRefreshCsReq                                 = 1403
	MatchThreeSyncDataScNotify                         = 7424
	FightSessionStopScNotify                           = 30030
	PunkLordMonsterKilledNotify                        = 3201
	GroupStateChangeScNotify                           = 1499
	StartCocoonStageCsReq                              = 1413
	EnterSectionCsReq                                  = 1428
	MarkReadMailCsReq                                  = 891
	SubmitOrigamiItemCsReq                             = 4196
	GetFriendLoginInfoScRsp                            = 2948
	RebattleByClientCsNotify                           = 194
	RecoverAllLineupScRsp                              = 1495
	LockRelicCsReq                                     = 582
	SavePointsInfoNotify                               = 1458
	GetStageLineupScRsp                                = 788
	GetSceneMapInfoCsReq                               = 1484
	ClockParkUnlockTalentScRsp                         = 7273
	RelicRecommendScRsp                                = 505
	SpaceZooExchangeItemScRsp                          = 6747
	MakeMissionDrinkCsReq                              = 6995
	AetherDivideSpiritExpUpScRsp                       = 4881
	GetLoginChatInfoScRsp                              = 3921
	StartTrialActivityScRsp                            = 2665
	GetFriendChallengeLineupCsReq                      = 2908
	GetCurSceneInfoScRsp                               = 1471
	FeatureSwitchClosedScNotify                        = 57
	SceneCastSkillCsReq                                = 1439
	UpdateFloorSavedValueNotify                        = 1414
	DressAvatarScRsp                                   = 327
	StartCocoonStageScRsp                              = 1485
	UnlockTeleportNotify                               = 1467
	ReEnterLastElementStageScRsp                       = 1410
	AlleyPlacingGameScRsp                              = 4794
	RecoverAllLineupCsReq                              = 1468
	StartRaidCsReq                                     = 2261
	BuyBpLevelScRsp                                    = 3030
	SpringRecoverScRsp                                 = 1446
	SetSpringRecoverConfigCsReq                        = 1476
	RogueDoGambleCsReq                                 = 5699
	EnterChessRogueAeonRoomScRsp                       = 5494
	SetGroupCustomSaveDataScRsp                        = 1465
	ChessRogueLayerAccountInfoNotify                   = 5492
	SpringRecoverSingleAvatarCsReq                     = 1444
	RefreshTriggerByClientScNotify                     = 1438
	AntiAddictScNotify                                 = 70
	BatchMarkChatEmojiCsReq                            = 3947
	MarkAvatarScRsp                                    = 301
	SelectChessRogueSubStoryScRsp                      = 5540
	SetAetherDivideLineUpCsReq                         = 4847
	EntityBindPropScRsp                                = 1426
	BuyGoodsScRsp                                      = 1520
	LogisticsDetonateStarSkiffScRsp                    = 4754
	RogueTournEnterRoomCsReq                           = 6094
	StartBoxingClubBattleCsReq                         = 4239
	StartTimedFarmElementCsReq                         = 1423
	WolfBroGameExplodeMonsterCsReq                     = 6503
	LeaveAetherDivideSceneScRsp                        = 4820
	GetRollShopInfoScRsp                               = 6902
	TakePrestigeRewardScRsp                            = 4774
	ClockParkGetOngoingScriptInfoCsReq                 = 7287
	SpringRecoverSingleAvatarScRsp                     = 1424
	GetPunkLordBattleRecordCsReq                       = 3255
	GameplayCounterCountDownCsReq                      = 1443
	GeneralVirtualItemDataNotify                       = 600
	GetRogueEndlessActivityDataCsReq                   = 6008
	ClockParkQuitScriptCsReq                           = 7247
	SceneEntityTeleportCsReq                           = 1498
	MultiplayerFightGameStateCsReq                     = 1061
	ChessRogueGiveUpRollScRsp                          = 5406
	GetFriendLoginInfoCsReq                            = 2940
	SyncRogueGetItemScNotify                           = 1884
	MarkAvatarCsReq                                    = 328
	EnteredSceneChangeScNotify                         = 1402
	GetSceneMapInfoScRsp                               = 1469
	GetCurSceneInfoCsReq                               = 1430
	MonopolyEventSelectFriendScRsp                     = 7057
	GameplayCounterCountDownScRsp                      = 1460
	SyncRogueCommonVirtualItemInfoScNotify             = 5618
	MonopolyAcceptQuizScRsp                            = 7054
	GetSpringRecoverDataScRsp                          = 1418
	DoGachaInRollShopScRsp                             = 6903
	EnterFantasticStoryActivityStageCsReq              = 4920
	RemoveStuffFromAreaScRsp                           = 4371
	StartFinishSubMissionScNotify                      = 1233
	GetPlayerReplayInfoScRsp                           = 3520
	SceneCastSkillMpUpdateScNotify                     = 1421
	SwitchLineupIndexCsReq                             = 721
	StartFinishMainMissionScNotify                     = 1250
	AcceptedPamMissionExpireCsReq                      = 4061
	RogueTournRenameArchiveScRsp                       = 6047
	MarkItemCsReq                                      = 558
	GetRogueInitialScoreScRsp                          = 1893
	RogueModifierSelectCellCsReq                       = 5320
	RogueTournQueryCsReq                               = 6058
	RogueTournEnablePermanentTalentScRsp               = 6061
	GetChallengeGroupStatisticsScRsp                   = 1703
	EnterAdventureCsReq                                = 1361
	ActivateFarmElementScRsp                           = 1466
	BuyNpcStuffCsReq                                   = 4391
	CommonRogueQueryScRsp                              = 5624
	RogueTournEnablePermanentTalentCsReq               = 6049
	RogueTournAreaUpdateScNotify                       = 6019
	FinishCosumeItemMissionCsReq                       = 1247
	LeaveMapRotationRegionScRsp                        = 6827
	RogueTournGetPermanentTalentInfoScRsp              = 6035
	RogueTournTakeExpRewardScRsp                       = 6080
	RogueTournStartCsReq                               = 6075
	ChessRogueChangeyAeonDimensionNotify               = 5531
	GetPlayerDetailInfoCsReq                           = 2991
	TakeAssistRewardScRsp                              = 2926
	GetMbtiReportCsReq                                 = 7042
	RogueTournConfirmSettleCsReq                       = 6026
	RogueTournRenameArchiveCsReq                       = 6090
	RogueTournEnterScRsp                               = 6031
	LobbyModifyPlayerInfoScRsp                         = 7353
	GetFriendBattleRecordDetailCsReq                   = 2924
	RogueTournGetCurRogueCocoonInfoScRsp               = 6021
	GetRogueBuffEnhanceInfoCsReq                       = 1809
	RogueTournGetSettleInfoCsReq                       = 6034
	RogueTournGetMiscRealTimeDataScRsp                 = 6020
	UnlockBackGroundMusicCsReq                         = 3139
	RogueTournRevivieCostUpdateScNotify                = 6097
	DelMailScRsp                                       = 873
	LeaveTrialActivityScRsp                            = 2634
	EvolveBuildTakeExpRewardCsReq                      = 7150
	RogueDoGambleScRsp                                 = 5607
	JoinLineupCsReq                                    = 739
	DrinkMakerUpdateTipsNotify                         = 6990
	ChessRogueStartScRsp                               = 5477
	RogueTournEnterRogueCocoonSceneCsReq               = 6092
	ChessRogueNousGetRogueTalentInfoCsReq              = 5586
	HeliobusSnsReadCsReq                               = 5891
	MonopolyBuyGoodsCsReq                              = 7049
	RogueTournLeaveRogueCocoonSceneCsReq               = 6100
	MuseumTargetMissionFinishNotify                    = 4337
	RogueTournHandBookNotify                           = 6091
	ComposeLimitNumUpdateNotify                        = 550
	RogueTournGetCurRogueCocoonInfoCsReq               = 6018
	GetShareDataScRsp                                  = 4120
	GetMonsterResearchActivityDataCsReq                = 2682
	GmTalkScRsp                                        = 62
	ChessRogueQueryCsReq                               = 5517
	GetMapRotationDataScRsp                            = 6874
	EvolveBuildStartStageCsReq                         = 7124
	GetLineupAvatarDataCsReq                           = 774
	GetChessRogueStoryInfoCsReq                        = 5542
	BoxingClubRewardScNotify                           = 4287
	AetherDivideSpiritInfoScNotify                     = 4812
	LeaveMapRotationRegionCsReq                        = 6887
	GetArchiveDataScRsp                                = 2388
	PunkLordRaidTimeOutScNotify                        = 3270
	RogueModifierAddNotify                             = 5391
	RogueModifierDelNotify                             = 5387
	RogueModifierStageStartNotify                      = 5327
	EnterRogueEndlessActivityStageScRsp                = 6005
	DailyFirstMeetPamScRsp                             = 3420
	UpdatePsnSettingsInfoScRsp                         = 45
	TakeAllApRewardScRsp                               = 3330
	GetMbtiReportScRsp                                 = 7016
	GetBasicInfoCsReq                                  = 90
	BuyNpcStuffScRsp                                   = 4320
	GetBagScRsp                                        = 588
	TakeRogueEndlessActivityPointRewardScRsp           = 6007
	TakeBpRewardScRsp                                  = 3039
	ChangeScriptEmotionScRsp                           = 6320
	ExchangeRogueBuffWithMiracleScRsp                  = 5656
	RogueWorkbenchHandleFuncScRsp                      = 5684
	SyncRogueHandbookDataUpdateScNotify                = 5626
	TakeRogueEventHandbookRewardScRsp                  = 5640
	HandleRogueCommonPendingActionCsReq                = 5646
	DeleteSummonUnitCsReq                              = 1477
	MakeMissionDrinkScRsp                              = 6993
	SyncRogueCommonActionResultScNotify                = 5648
	EnhanceCommonRogueBuffCsReq                        = 5609
	ChessRogueSelectCellCsReq                          = 5591
	GetPlayerBoardDataScRsp                            = 2888
	RogueNpcDisappearScRsp                             = 5647
	RogueWorkbenchHandleFuncCsReq                      = 5679
	PlayerGetTokenScRsp                                = 73
	GetAllSaveRaidCsReq                                = 2282
	GetRogueCommonDialogueDataCsReq                    = 5616
	SelectRogueCommonDialogueOptionCsReq               = 5657
	GetRogueShopBuffInfoCsReq                          = 5673
	RogueGetGambleInfoScRsp                            = 5635
	ContentPackageSyncDataScNotify                     = 7470
	GetRogueShopBuffInfoScRsp                          = 5630
	UnlockBackGroundMusicScRsp                         = 3173
	TakeOffEquipmentCsReq                              = 362
	RogueWorkbenchSelectFuncCsReq                      = 5678
	ServerSimulateBattleFinishScNotify                 = 174
	EnterChallengeNextPhaseScRsp                       = 1712
	UpdateGunPlayDataScRsp                             = 4101
	GetRndOptionCsReq                                  = 3461
	RogueWorkbenchSelectFuncScRsp                      = 5683
	MonopolyCheatDiceCsReq                             = 7028
	StopRogueAdventureRoomScRsp                        = 5637
	EvolveBuildFinishScNotify                          = 7140
	TakeOffEquipmentScRsp                              = 374
	DressRelicAvatarScRsp                              = 382
	BuyRogueShopBuffScRsp                              = 5662
	RaidCollectionDataScNotify                         = 6947
	TravelBrochureSetCustomValueScRsp                  = 6421
	HandleRogueCommonPendingActionScRsp                = 5611
	ChessRogueMoveCellNotify                           = 5582
	ReportPlayerCsReq                                  = 2909
	BuyRogueShopMiracleCsReq                           = 5671
	AcceptedPamMissionExpireScRsp                      = 4088
	CommonRogueQueryCsReq                              = 5644
	ExchangeRogueBuffWithMiracleCsReq                  = 5670
	MonopolyLikeScRsp                                  = 7044
	BuyRogueShopMiracleScRsp                           = 5687
	GetRogueCollectionScRsp                            = 5625
	GetRogueHandbookDataCsReq                          = 5685
	GetJukeboxDataCsReq                                = 3161
	RogueNpcDisappearCsReq                             = 5674
	ComposeLimitNumCompleteNotify                      = 533
	TakePunkLordPointRewardCsReq                       = 3247
	StopRogueAdventureRoomCsReq                        = 5612
	AetherDivideFinishChallengeScNotify                = 4801
	SetRogueExhibitionCsReq                            = 5643
	GetRogueShopMiracleInfoCsReq                       = 5620
	PVEBattleResultCsReq                               = 161
	SyncRogueAdventureRoomInfoScNotify                 = 5661
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5498
	StartAetherDivideStageBattleScRsp                  = 4851
	ChessRogueNousEditDiceCsReq                        = 5464
	GetRogueCollectionCsReq                            = 5669
	FinishAeonDialogueGroupCsReq                       = 1852
	SyncRogueStatusScNotify                            = 1819
	ChessRogueNousDiceUpdateNotify                     = 5418
	LeaveAetherDivideSceneCsReq                        = 4891
	UpdateTrackMainMissionIdScRsp                      = 1254
	SyncRogueAeonLevelUpRewardScNotify                 = 1814
	GetRogueScoreRewardInfoCsReq                       = 1843
	RogueTournBattleFailSettleInfoScNotify             = 6060
	SceneReviveAfterRebattleScRsp                      = 1437
	TakeRogueAeonLevelRewardScRsp                      = 1825
	SyncRogueReviveInfoScNotify                        = 1832
	RefreshAlleyOrderScRsp                             = 4703
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5651
	TravelBrochurePageResetScRsp                       = 6456
	SetBoxingClubResonanceLineupCsReq                  = 4247
	GetRecyleTimeScRsp                                 = 501
	EnterRogueCsReq                                    = 1839
	OpenRogueChestScRsp                                = 1824
	ShowNewSupplementVisitorScRsp                      = 3774
	SyncRogueRewardInfoScNotify                        = 1867
	LeaveRogueScRsp                                    = 1871
	ExchangeStaminaScRsp                               = 96
	LeaveRaidCsReq                                     = 2291
	TakeExpeditionRewardScRsp                          = 2571
	TravelBrochureGetDataScRsp                         = 6488
	RogueTournGetArchiveRepositoryCsReq                = 6071
	GetRogueInfoCsReq                                  = 1861
	OpenRogueChestCsReq                                = 1844
	HeliobusSelectSkillCsReq                           = 5821
	PlayerReturnInfoQueryScRsp                         = 4527
	ChessRogueNousGetRogueTalentInfoScRsp              = 5577
	TakeActivityExpeditionRewardCsReq                  = 2594
	RankUpAvatarCsReq                                  = 394
	ServerAnnounceNotify                               = 50
	SummonPunkLordMonsterScRsp                         = 3271
	GetTrainVisitorBehaviorScRsp                       = 3720
	CancelMarkItemNotify                               = 585
	TakeMultipleExpeditionRewardScRsp                  = 2570
	GetRndOptionScRsp                                  = 3488
	GetPlayerReplayInfoCsReq                           = 3591
	UpdateRedDotDataCsReq                              = 5991
	SceneEntityTeleportScRsp                           = 1486
	RaidCollectionDataCsReq                            = 6949
	SyncChessRogueNousMainStoryScNotify                = 5537
	ExpUpRelicCsReq                                    = 596
	GetBoxingClubInfoScRsp                             = 4288
	MonopolyUpgradeAssetScRsp                          = 7081
	AddBlacklistScRsp                                  = 2982
	EnterRogueScRsp                                    = 1873
	AetherDivideTainerInfoScNotify                     = 4833
	GetCurChallengeCsReq                               = 1727
	MatchThreeGetDataCsReq                             = 7445
	EndDrinkMakerSequenceScRsp                         = 6994
	SyncApplyFriendScNotify                            = 2987
	GetRaidInfoCsReq                                   = 2262
	StartMatchScRsp                                    = 7317
	RaidKickByServerScNotify                           = 2256
	ChessRogueEnterCsReq                               = 5557
	EvolveBuildReRandomStageScRsp                      = 7129
	CancelActivityExpeditionCsReq                      = 2574
	ChangeLineupLeaderScRsp                            = 796
	EnhanceChessRogueBuffScRsp                         = 5410
	FightMatch3ForceUpdateNotify                       = 30174
	SellItemScRsp                                      = 556
	MonopolyCellUpdateNotify                           = 7020
	GetSaveRaidScRsp                                   = 2221
	BattleCollegeDataChangeScNotify                    = 5791
	GetMissionEventDataScRsp                           = 1221
	HeliobusSnsReadScRsp                               = 5820
	SetAssistAvatarScRsp                               = 2847
	MuseumDispatchFinishedScNotify                     = 4381
	GetAllSaveRaidScRsp                                = 2203
	HeliobusActivityDataCsReq                          = 5861
	TakeChallengeRaidRewardCsReq                       = 2271
	FightMatch3ChatCsReq                               = 30187
	GetUnlockTeleportScRsp                             = 1453
	GetExpeditionDataCsReq                             = 2561
	MatchBoxingClubOpponentCsReq                       = 4291
	StartRaidScRsp                                     = 2288
	UnlockTutorialScRsp                                = 1673
	BatchGetQuestDataCsReq                             = 996
	EnterRogueMapRoomScRsp                             = 1817
	TakeQuestOptionalRewardCsReq                       = 974
	DestroyItemScRsp                                   = 555
	ExchangeHcoinScRsp                                 = 509
	FightKickOutScNotify                               = 30020
	LobbyCreateScRsp                                   = 7367
	QueryProductInfoScRsp                              = 48
	GetCrossInfoCsReq                                  = 7328
	EnterFeverTimeActivityStageScRsp                   = 7151
	FinishQuestCsReq                                   = 927
	SetGenderCsReq                                     = 54
	BatchGetQuestDataScRsp                             = 921
	TakeAllApRewardCsReq                               = 3373
	TakeQuestRewardScRsp                               = 920
	MonopolyClickCellScRsp                             = 7007
	FightMatch3DataScRsp                               = 30188
	PlayerLoginScRsp                                   = 88
	QuestRecordScNotify                                = 987
	GetFarmStageGachaInfoCsReq                         = 1391
	TakeKilledPunkLordMonsterScoreCsReq                = 3233
	SetForbidOtherApplyFriendCsReq                     = 2905
	GroupStateChangeCsReq                              = 1441
	UpdateRogueAdventureRoomScoreScRsp                 = 5690
	HeliobusSnsCommentScRsp                            = 5827
	SetAssistScRsp                                     = 2955
	WolfBroGameDataChangeScNotify                      = 6562
	TakePictureScRsp                                   = 4173
	TakePunkLordPointRewardScRsp                       = 3294
	PromoteAvatarScRsp                                 = 371
	SetFriendMarkScRsp                                 = 2918
	QuitLineupCsReq                                    = 730
	FightMatch3TurnEndScNotify                         = 30139
	TakeRogueEventHandbookRewardCsReq                  = 5693
	TakeAllRewardCsReq                                 = 3071
	HeliobusSnsPostCsReq                               = 5839
	AlleyPlacingGameCsReq                              = 4747
	SummonPunkLordMonsterCsReq                         = 3230
	UnlockPhoneThemeScNotify                           = 5171
	GetAlleyInfoScRsp                                  = 4788
	TakeKilledPunkLordMonsterScoreScRsp                = 3250
	UpdateEnergyScNotify                               = 6821
	SyncRogueCommonDialogueDataScNotify                = 5698
	LobbyInviteScNotify                                = 7357
	ReturnLastTownCsReq                                = 1449
	ClientObjDownloadDataScNotify                      = 43
	TravelBrochureGetDataCsReq                         = 6461
	LobbyBeginCsReq                                    = 7374
	ComposeSelectedRelicScRsp                          = 512
	GetMarkItemListScRsp                               = 595
	FightMatch3TurnStartScNotify                       = 30120
	PlayerLoginFinishCsReq                             = 86
	TreasureDungeonDataScNotify                        = 4461
	ChessRogueUpdateActionPointScNotify                = 5419
	PlayerReturnSignCsReq                              = 4588
	EnterFantasticStoryActivityStageScRsp              = 4939
	TriggerVoiceScRsp                                  = 4194
	FightMatch3StartCountDownScNotify                  = 30191
	MuseumRandomEventSelectCsReq                       = 4351
	StartAlleyEventScRsp                               = 4771
	PlayerReturnTakeRewardScRsp                        = 4571
	PlayerReturnInfoQueryCsReq                         = 4587
	GetNpcTakenRewardScRsp                             = 2188
	FinishTutorialScRsp                                = 1627
	ChooseBoxingClubStageOptionalBuffScRsp             = 4221
	GetAuthkeyScRsp                                    = 82
	EvolveBuildQueryInfoCsReq                          = 7145
	GetChessRogueNousStoryInfoCsReq                    = 5431
	SyncLineupNotify                                   = 762
	GetFirstTalkByPerformanceNpcCsReq                  = 2162
	CancelActivityExpeditionScRsp                      = 2547
	SetSignatureScRsp                                  = 2862
	AcceptMultipleExpeditionCsReq                      = 2521
	FightEnterCsReq                                    = 30061
	SetAssistAvatarCsReq                               = 2874
	SpringRecoverCsReq                                 = 1408
	GetRogueHandbookDataScRsp                          = 5654
	RogueTournResetPermanentTalentScRsp                = 6050
	EnterMapRotationRegionScRsp                        = 6888
	DailyFirstEnterMonopolyActivityCsReq               = 7047
	GateServerScNotify                                 = 65
	SetIsDisplayAvatarInfoScRsp                        = 2871
	StartStarFightLevelScRsp                           = 7165
	ClientDownloadDataScNotify                         = 5
	FantasticStoryActivityBattleEndScNotify            = 4973
	GetVideoVersionKeyScRsp                            = 83
	GetShareDataCsReq                                  = 4191
	MonthCardRewardNotify                              = 44
	SyncAcceptedPamMissionNotify                       = 4091
	DressAvatarSkinCsReq                               = 351
	GetMuseumInfoScRsp                                 = 4388
	ChessRogueGoAheadCsReq                             = 5472
	SetLanguageCsReq                                   = 1
	GetFriendRecommendListInfoScRsp                    = 2956
	HeliobusSnsLikeScRsp                               = 5871
	ResetMapRotationRegionScRsp                        = 6894
	FightHeartBeatCsReq                                = 30039
	PlayerHeartBeatCsReq                               = 42
	GetMultipleDropInfoScRsp                           = 4688
	MuseumRandomEventStartScNotify                     = 4370
	TakeRollShopRewardCsReq                            = 6919
	GetDailyActiveInfoScRsp                            = 3320
	SetGameplayBirthdayScRsp                           = 8
	ChessRogueCheatRollScRsp                           = 5456
	TakeActivityExpeditionRewardScRsp                  = 2596
	ExpeditionDataChangeScNotify                       = 2587
	GetHeroBasicTypeInfoCsReq                          = 68
	SetMissionEventProgressScRsp                       = 1212
	RetcodeNotify                                      = 24
	FinishChessRogueSubStoryCsReq                      = 5405
	RogueTournSettleCsReq                              = 6030
	UpdatePlayerSettingScRsp                           = 14
	SyncRoguePickAvatarInfoScNotify                    = 1892
	AceAntiCheaterCsReq                                = 46
	ChooseBoxingClubResonanceScRsp                     = 4274
	GetAvatarDataCsReq                                 = 361
	GetDrinkMakerDataScRsp                             = 6982
	GetRogueAeonInfoScRsp                              = 1807
	GmTalkScNotify                                     = 71
	SetNicknameScRsp                                   = 49
	GetRogueEndlessActivityDataScRsp                   = 6001
	SyncClientResVersionScRsp                          = 171
	GetDailyActiveInfoCsReq                            = 3391
	EvolveBuildGiveupCsReq                             = 7133
	DailyRefreshNotify                                 = 76
	SetGameplayBirthdayCsReq                           = 64
	GetLevelRewardTakenListCsReq                       = 51
	StartAetherDivideChallengeBattleScRsp              = 4871
	AetherDivideSpiritExpUpCsReq                       = 4809
	EnhanceRogueBuffCsReq                              = 1812
	OfferingInfoScNotify                               = 6935
	GetFriendAssistListCsReq                           = 2976
	AcceptExpeditionScRsp                              = 2520
	GetGachaInfoScRsp                                  = 1988
	GetPhoneDataScRsp                                  = 5188
	BatchMarkChatEmojiScRsp                            = 3994
	RogueTournSettleScRsp                              = 6043
	SpaceZooDataCsReq                                  = 6761
	SelectChatBubbleScRsp                              = 5120
	GetPhoneDataCsReq                                  = 5161
	AlleyShipUnlockScNotify                            = 4712
	GetOfferingInfoCsReq                               = 6929
	UseItemCsReq                                       = 530
	ClearAetherDividePassiveSkillCsReq                 = 4882
	ChessRogueUpdateAeonModifierValueScNotify          = 5594
	PVEBattleResultScRsp                               = 188
	PlayerReturnTakePointRewardCsReq                   = 4539
	GetPunkLordMonsterDataCsReq                        = 3261
	SyncRogueCommonPendingActionScNotify               = 5605
	GetCurLineupDataCsReq                              = 791
	CancelMatchCsReq                                   = 7320
	DailyTaskDataScNotify                              = 1271
	LeaveRogueCsReq                                    = 1830
	EnterFightActivityStageCsReq                       = 3620
	GetActivityScheduleConfigScRsp                     = 2673
	UnlockTutorialGuideScRsp                           = 1671
	FightMatch3ChatScNotify                            = 30162
	InteractPropScRsp                                  = 1420
	PlayerReturnSignScRsp                              = 4591
	EvolveBuildCoinNotify                              = 7136
	GetChatFriendHistoryCsReq                          = 3930
	DiscardRelicScRsp                                  = 540
	MonopolyClickMbtiReportCsReq                       = 7075
	MuseumInfoChangedScNotify                          = 4382
	LeaveChallengeScRsp                                = 1773
	SetDisplayAvatarScRsp                              = 2873
	MarkChatEmojiCsReq                                 = 3962
	FinishCurTurnCsReq                                 = 4362
	TextJoinQueryCsReq                                 = 3891
	UnlockChatBubbleScNotify                           = 5139
	TakeRogueScoreRewardCsReq                          = 1849
	GetAllLineupDataScRsp                              = 749
	SpaceZooDeleteCatCsReq                             = 6787
	PlayerReturnTakePointRewardScRsp                   = 4573
	RogueTournLeaveScRsp                               = 6059
	AvatarExpUpCsReq                                   = 391
	MultipleDropInfoScNotify                           = 4691
	DailyActiveInfoNotify                              = 3339
	SendMsgScRsp                                       = 3988
	MultipleDropInfoNotify                             = 4673
	AcceptMainMissionCsReq                             = 1232
	EvolveBuildTakeExpRewardScRsp                      = 7115
	MultiplayerFightGiveUpCsReq                        = 1039
	ReportPlayerScRsp                                  = 2981
	MultiplayerMatch3FinishScNotify                    = 1087
	TakeFightActivityRewardCsReq                       = 3673
	GetFantasticStoryActivityDataCsReq                 = 4961
	GetMonopolyMbtiReportRewardScRsp                   = 7015
	ClockParkGetInfoScRsp                              = 7288
	TravelBrochureSetCustomValueCsReq                  = 6496
	AddAvatarScNotify                                  = 347
	FinishChessRogueNousSubStoryCsReq                  = 5535
	GetMonopolyMbtiReportRewardCsReq                   = 7045
	MonopolyTakePhaseRewardScRsp                       = 7060
	GetNpcMessageGroupScRsp                            = 2788
	TriggerVoiceCsReq                                  = 4147
	ApplyFriendScRsp                                   = 2971
	ChessRogueQueryScRsp                               = 5448
	TrialBackGroundMusicScRsp                          = 3171
	GetChessRogueBuffEnhanceInfoScRsp                  = 5524
	MonopolyUpgradeAssetCsReq                          = 7009
	MonopolyRollDiceScRsp                              = 7030
	GetFriendDevelopmentInfoCsReq                      = 2916
	EvolveBuildShopAbilityDownCsReq                    = 7148
	StoryLineInfoScNotify                              = 6291
	FinishItemIdScRsp                                  = 2773
	AetherDivideRefreshEndlessScRsp                    = 4895
	MonopolyRollRandomCsReq                            = 7096
	EnterAetherDivideSceneCsReq                        = 4861
	SyncDeleteFriendScNotify                           = 2996
	EnhanceCommonRogueBuffScRsp                        = 5681
	StartChallengeCsReq                                = 1791
	GetMonopolyDailyReportScRsp                        = 7035
	TakeRogueEndlessActivityPointRewardCsReq           = 6002
	GetCurChallengeScRsp                               = 1762
	MonopolyBuyGoodsScRsp                              = 7051
	ChessRogueQueryBpCsReq                             = 5475
	NewAssistHistoryNotify                             = 2985
	GetAllLineupDataCsReq                              = 756
	MonopolyConfirmRandomScRsp                         = 7056
	StartAetherDivideSceneBattleScRsp                  = 4873
	RogueTournResetPermanentTalentCsReq                = 6011
	MonopolyTakePhaseRewardCsReq                       = 7043
	MonopolyGetRafflePoolInfoScRsp                     = 7014
	QuitBattleScNotify                                 = 187
	MonopolyGameCreateScNotify                         = 7026
	StartPartialChallengeScRsp                         = 1756
	GetShopListScRsp                                   = 1588
	SubMissionRewardScNotify                           = 1237
	AcceptActivityExpeditionScRsp                      = 2562
	RogueTournReEnterRogueCocoonStageCsReq             = 6083
	EnterStrongChallengeActivityStageCsReq             = 6691
	GetMissionDataCsReq                                = 1261
	SelectChatBubbleCsReq                              = 5191
	GetAetherDivideInfoCsReq                           = 4862
	AvatarExpUpScRsp                                   = 320
	ChessRogueFinishCurRoomNotify                      = 5426
	MissionEventRewardScNotify                         = 1282
	PrivateMsgOfflineUsersScNotify                     = 3920
	AcceptMissionEventScRsp                            = 1270
	ChessRogueEnterCellScRsp                           = 5562
	EvolveBuildGiveupScRsp                             = 7103
	GetChatEmojiListScRsp                              = 3927
	HeliobusSnsUpdateScNotify                          = 5862
	LobbyBeginScRsp                                    = 7378
	TravelBrochureApplyPasterScRsp                     = 6471
	EnterChallengeNextPhaseCsReq                       = 1781
	PlayerReturnForceFinishScNotify                    = 4562
	AcceptMultipleExpeditionScRsp                      = 2582
	RogueTournEnterLayerCsReq                          = 6095
	TakePictureCsReq                                   = 4139
	UpdateGunPlayDataCsReq                             = 4128
	ChessRogueReviveAvatarCsReq                        = 5570
	SetForbidOtherApplyFriendScRsp                     = 2966
	SceneCastSkillScRsp                                = 1473
	GetMovieRacingDataScRsp                            = 4151
	GetChessRogueStoryInfoScRsp                        = 5515
	MatchThreeGetDataScRsp                             = 7417
	ClockParkQuitScriptScRsp                           = 7294
	LobbyInviteCsReq                                   = 7355
	DressAvatarSkinScRsp                               = 309
	StartAlleyEventCsReq                               = 4730
	HeliobusActivityDataScRsp                          = 5888
	MuseumTargetRewardNotify                           = 4328
	FightLeaveScNotify                                 = 30091
	FightMatch3ChatScRsp                               = 30127
	GetMovieRacingDataCsReq                            = 4149
	ResetMapRotationRegionCsReq                        = 6847
	LobbyKickOutScRsp                                  = 7356
	ChessRogueEnterNextLayerScRsp                      = 5598
	SwitchLineupIndexScRsp                             = 782
	GetPrivateChatHistoryScRsp                         = 3973
	HeliobusEnterBattleScRsp                           = 5849
	LeaveMapRotationRegionScNotify                     = 6896
	ChallengeRaidNotify                                = 2227
	TakeMultipleExpeditionRewardCsReq                  = 2503
	PromoteEquipmentCsReq                              = 591
	SwitchAetherDivideLineUpSlotScRsp                  = 4856
	LockEquipmentScRsp                                 = 573
	RechargeSuccNotify                                 = 549
	HeliobusUpgradeLevelCsReq                          = 5847
	MonopolySelectOptionCsReq                          = 7027
	FinishCosumeItemMissionScRsp                       = 1294
	SubmitMonsterResearchActivityMaterialScRsp         = 2656
	GetTrialActivityDataScRsp                          = 2608
	HeartDialTraceScriptCsReq                          = 6327
	GetGachaCeilingCsReq                               = 1939
	SyncAddBlacklistScNotify                           = 2903
	UnlockAvatarSkinScNotify                           = 337
	ArchiveWolfBroGameCsReq                            = 6591
	ChessRogueQueryAeonDimensionsScRsp                 = 5590
	FinishEmotionDialoguePerformanceCsReq              = 6330
	GetFightActivityDataScRsp                          = 3688
	GetEnteredSceneCsReq                               = 1407
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
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(RogueTournReEnterRogueCocoonStageScRsp, func() any { return new(proto.RogueTournReEnterRogueCocoonStageScRsp) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(RogueTournDeleteArchiveScRsp, func() any { return new(proto.RogueTournDeleteArchiveScRsp) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(RogueTournReviveAvatarCsReq, func() any { return new(proto.RogueTournReviveAvatarCsReq) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(CommonRogueComponentUpdateScNotify, func() any { return new(proto.CommonRogueComponentUpdateScNotify) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(StartPartialChallengeCsReq, func() any { return new(proto.StartPartialChallengeCsReq) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(RogueTournExpNotify, func() any { return new(proto.RogueTournExpNotify) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(HeroBasicTypeChangedNotify, func() any { return new(proto.HeroBasicTypeChangedNotify) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(DeleteFriendScRsp, func() any { return new(proto.DeleteFriendScRsp) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(GetAuthkeyCsReq, func() any { return new(proto.GetAuthkeyCsReq) })
	c.regMsg(RogueTournGetSettleInfoScRsp, func() any { return new(proto.RogueTournGetSettleInfoScRsp) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(RogueTournDifficultyCompNotify, func() any { return new(proto.RogueTournDifficultyCompNotify) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(RogueTournReviveAvatarScRsp, func() any { return new(proto.RogueTournReviveAvatarScRsp) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(GetFriendChallengeLineupScRsp, func() any { return new(proto.GetFriendChallengeLineupScRsp) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(RogueTournEnterRoomScRsp, func() any { return new(proto.RogueTournEnterRoomScRsp) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(StartTimedFarmElementScRsp, func() any { return new(proto.StartTimedFarmElementScRsp) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(UpdateMovieRacingDataScRsp, func() any { return new(proto.UpdateMovieRacingDataScRsp) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(RogueTournClearArchiveNameScNotify, func() any { return new(proto.RogueTournClearArchiveNameScNotify) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(CommonRogueUpdateScNotify, func() any { return new(proto.CommonRogueUpdateScNotify) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(RogueTournWeekChallengeUpdateScNotify, func() any { return new(proto.RogueTournWeekChallengeUpdateScNotify) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(FightMatch3SwapCsReq, func() any { return new(proto.FightMatch3SwapCsReq) })
	c.regMsg(RogueTournQueryScRsp, func() any { return new(proto.RogueTournQueryScRsp) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(FightEnterScRsp, func() any { return new(proto.FightEnterScRsp) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(FinishRogueCommonDialogueCsReq, func() any { return new(proto.FinishRogueCommonDialogueCsReq) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(GetRogueCommonDialogueDataScRsp, func() any { return new(proto.GetRogueCommonDialogueDataScRsp) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(RogueTournGetPermanentTalentInfoCsReq, func() any { return new(proto.RogueTournGetPermanentTalentInfoCsReq) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(GetFriendDevelopmentInfoScRsp, func() any { return new(proto.GetFriendDevelopmentInfoScRsp) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(ContentPackageGetDataScRsp, func() any { return new(proto.ContentPackageGetDataScRsp) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(RogueTournLevelInfoUpdateScNotify, func() any { return new(proto.RogueTournLevelInfoUpdateScNotify) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(GetSpringRecoverDataCsReq, func() any { return new(proto.GetSpringRecoverDataCsReq) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(GetRogueExhibitionScRsp, func() any { return new(proto.GetRogueExhibitionScRsp) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(SubmitOrigamiItemScRsp, func() any { return new(proto.SubmitOrigamiItemScRsp) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(SubmitMonsterResearchActivityMaterialCsReq, func() any { return new(proto.SubmitMonsterResearchActivityMaterialCsReq) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(RogueTournStartScRsp, func() any { return new(proto.RogueTournStartScRsp) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(GetHeroBasicTypeInfoScRsp, func() any { return new(proto.GetHeroBasicTypeInfoScRsp) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(RogueTournGetMiscRealTimeDataCsReq, func() any { return new(proto.RogueTournGetMiscRealTimeDataCsReq) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(ChallengeBossPhaseSettleNotify, func() any { return new(proto.ChallengeBossPhaseSettleNotify) })
	c.regMsg(SetHeroBasicTypeCsReq, func() any { return new(proto.SetHeroBasicTypeCsReq) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(RogueTournEnterRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournEnterRogueCocoonSceneScRsp) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(RogueTournEnterCsReq, func() any { return new(proto.RogueTournEnterCsReq) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(GetRogueExhibitionCsReq, func() any { return new(proto.GetRogueExhibitionCsReq) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(FightHeartBeatScRsp, func() any { return new(proto.FightHeartBeatScRsp) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(TakeMonsterResearchActivityRewardCsReq, func() any { return new(proto.TakeMonsterResearchActivityRewardCsReq) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(RogueTournTakeExpRewardCsReq, func() any { return new(proto.RogueTournTakeExpRewardCsReq) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(GetFriendChallengeDetailCsReq, func() any { return new(proto.GetFriendChallengeDetailCsReq) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(TakeFightActivityRewardScRsp, func() any { return new(proto.TakeFightActivityRewardScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(GetMonsterResearchActivityDataScRsp, func() any { return new(proto.GetMonsterResearchActivityDataScRsp) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(UpdateMovieRacingDataCsReq, func() any { return new(proto.UpdateMovieRacingDataCsReq) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(DeleteFriendCsReq, func() any { return new(proto.DeleteFriendCsReq) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(HealPoolInfoNotify, func() any { return new(proto.HealPoolInfoNotify) })
	c.regMsg(SyncRogueCommonDialogueOptionFinishScNotify, func() any { return new(proto.SyncRogueCommonDialogueOptionFinishScNotify) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(SetFriendMarkCsReq, func() any { return new(proto.SetFriendMarkCsReq) })
	c.regMsg(EnterFightActivityStageScRsp, func() any { return new(proto.EnterFightActivityStageScRsp) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(GetChallengeGroupStatisticsCsReq, func() any { return new(proto.GetChallengeGroupStatisticsCsReq) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(SelectRogueCommonDialogueOptionScRsp, func() any { return new(proto.SelectRogueCommonDialogueOptionScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryScRsp, func() any { return new(proto.RogueTournGetArchiveRepositoryScRsp) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(GetFriendChallengeDetailScRsp, func() any { return new(proto.GetFriendChallengeDetailScRsp) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(RogueTournEnterLayerScRsp, func() any { return new(proto.RogueTournEnterLayerScRsp) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(GetFriendAssistListScRsp, func() any { return new(proto.GetFriendAssistListScRsp) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(GetFriendBattleRecordDetailScRsp, func() any { return new(proto.GetFriendBattleRecordDetailScRsp) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(GetLevelRewardScRsp, func() any { return new(proto.GetLevelRewardScRsp) })
	c.regMsg(TakeMonsterResearchActivityRewardScRsp, func() any { return new(proto.TakeMonsterResearchActivityRewardScRsp) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(FightMatch3SwapScRsp, func() any { return new(proto.FightMatch3SwapScRsp) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(ContentPackageGetDataCsReq, func() any { return new(proto.ContentPackageGetDataCsReq) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(FinishRogueCommonDialogueScRsp, func() any { return new(proto.FinishRogueCommonDialogueScRsp) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(RogueTournGetAllArchiveCsReq, func() any { return new(proto.RogueTournGetAllArchiveCsReq) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(RogueTournGetAllArchiveScRsp, func() any { return new(proto.RogueTournGetAllArchiveScRsp) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(SetSpringRecoverConfigScRsp, func() any { return new(proto.SetSpringRecoverConfigScRsp) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(SetRedPointStatusScNotify, func() any { return new(proto.SetRedPointStatusScNotify) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(ContentPackageUnlockScRsp, func() any { return new(proto.ContentPackageUnlockScRsp) })
	c.regMsg(SetHeroBasicTypeScRsp, func() any { return new(proto.SetHeroBasicTypeScRsp) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(RogueTournDeleteArchiveCsReq, func() any { return new(proto.RogueTournDeleteArchiveCsReq) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(FightMatch3DataCsReq, func() any { return new(proto.FightMatch3DataCsReq) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneScRsp, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneScRsp) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(RogueTournConfirmSettleScRsp, func() any { return new(proto.RogueTournConfirmSettleScRsp) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(RogueTournLeaveCsReq, func() any { return new(proto.RogueTournLeaveCsReq) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(StartTimedCocoonStageScRsp, func() any { return new(proto.StartTimedCocoonStageScRsp) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(StartTimedCocoonStageCsReq, func() any { return new(proto.StartTimedCocoonStageCsReq) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(FightSessionStopScNotify, func() any { return new(proto.FightSessionStopScNotify) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(SubmitOrigamiItemCsReq, func() any { return new(proto.SubmitOrigamiItemCsReq) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(GetFriendChallengeLineupCsReq, func() any { return new(proto.GetFriendChallengeLineupCsReq) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(SpringRecoverScRsp, func() any { return new(proto.SpringRecoverScRsp) })
	c.regMsg(SetSpringRecoverConfigCsReq, func() any { return new(proto.SetSpringRecoverConfigCsReq) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(SpringRecoverSingleAvatarCsReq, func() any { return new(proto.SpringRecoverSingleAvatarCsReq) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(MarkAvatarScRsp, func() any { return new(proto.MarkAvatarScRsp) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(RogueTournEnterRoomCsReq, func() any { return new(proto.RogueTournEnterRoomCsReq) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(SpringRecoverSingleAvatarScRsp, func() any { return new(proto.SpringRecoverSingleAvatarScRsp) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(MarkAvatarCsReq, func() any { return new(proto.MarkAvatarCsReq) })
	c.regMsg(EnteredSceneChangeScNotify, func() any { return new(proto.EnteredSceneChangeScNotify) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(GetSpringRecoverDataScRsp, func() any { return new(proto.GetSpringRecoverDataScRsp) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(RogueTournRenameArchiveScRsp, func() any { return new(proto.RogueTournRenameArchiveScRsp) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(RogueTournQueryCsReq, func() any { return new(proto.RogueTournQueryCsReq) })
	c.regMsg(RogueTournEnablePermanentTalentScRsp, func() any { return new(proto.RogueTournEnablePermanentTalentScRsp) })
	c.regMsg(GetChallengeGroupStatisticsScRsp, func() any { return new(proto.GetChallengeGroupStatisticsScRsp) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(RogueTournEnablePermanentTalentCsReq, func() any { return new(proto.RogueTournEnablePermanentTalentCsReq) })
	c.regMsg(RogueTournAreaUpdateScNotify, func() any { return new(proto.RogueTournAreaUpdateScNotify) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(RogueTournGetPermanentTalentInfoScRsp, func() any { return new(proto.RogueTournGetPermanentTalentInfoScRsp) })
	c.regMsg(RogueTournTakeExpRewardScRsp, func() any { return new(proto.RogueTournTakeExpRewardScRsp) })
	c.regMsg(RogueTournStartCsReq, func() any { return new(proto.RogueTournStartCsReq) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(RogueTournConfirmSettleCsReq, func() any { return new(proto.RogueTournConfirmSettleCsReq) })
	c.regMsg(RogueTournRenameArchiveCsReq, func() any { return new(proto.RogueTournRenameArchiveCsReq) })
	c.regMsg(RogueTournEnterScRsp, func() any { return new(proto.RogueTournEnterScRsp) })
	c.regMsg(GetFriendBattleRecordDetailCsReq, func() any { return new(proto.GetFriendBattleRecordDetailCsReq) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoScRsp, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(RogueTournGetSettleInfoCsReq, func() any { return new(proto.RogueTournGetSettleInfoCsReq) })
	c.regMsg(RogueTournGetMiscRealTimeDataScRsp, func() any { return new(proto.RogueTournGetMiscRealTimeDataScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(RogueTournEnterRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournEnterRogueCocoonSceneCsReq) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(RogueTournLeaveRogueCocoonSceneCsReq, func() any { return new(proto.RogueTournLeaveRogueCocoonSceneCsReq) })
	c.regMsg(RogueTournHandBookNotify, func() any { return new(proto.RogueTournHandBookNotify) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(RogueTournGetCurRogueCocoonInfoCsReq, func() any { return new(proto.RogueTournGetCurRogueCocoonInfoCsReq) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(GetMonsterResearchActivityDataCsReq, func() any { return new(proto.GetMonsterResearchActivityDataCsReq) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(RogueWorkbenchHandleFuncScRsp, func() any { return new(proto.RogueWorkbenchHandleFuncScRsp) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(RogueWorkbenchHandleFuncCsReq, func() any { return new(proto.RogueWorkbenchHandleFuncCsReq) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(GetRogueCommonDialogueDataCsReq, func() any { return new(proto.GetRogueCommonDialogueDataCsReq) })
	c.regMsg(SelectRogueCommonDialogueOptionCsReq, func() any { return new(proto.SelectRogueCommonDialogueOptionCsReq) })
	c.regMsg(GetRogueShopBuffInfoCsReq, func() any { return new(proto.GetRogueShopBuffInfoCsReq) })
	c.regMsg(ContentPackageSyncDataScNotify, func() any { return new(proto.ContentPackageSyncDataScNotify) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(RogueWorkbenchSelectFuncCsReq, func() any { return new(proto.RogueWorkbenchSelectFuncCsReq) })
	c.regMsg(EnterChallengeNextPhaseScRsp, func() any { return new(proto.EnterChallengeNextPhaseScRsp) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(RogueWorkbenchSelectFuncScRsp, func() any { return new(proto.RogueWorkbenchSelectFuncScRsp) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(EvolveBuildFinishScNotify, func() any { return new(proto.EvolveBuildFinishScNotify) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(BuyRogueShopMiracleCsReq, func() any { return new(proto.BuyRogueShopMiracleCsReq) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(GetRogueCollectionScRsp, func() any { return new(proto.GetRogueCollectionScRsp) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(GetRogueShopMiracleInfoCsReq, func() any { return new(proto.GetRogueShopMiracleInfoCsReq) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(GetRogueCollectionCsReq, func() any { return new(proto.GetRogueCollectionCsReq) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(RogueTournBattleFailSettleInfoScNotify, func() any { return new(proto.RogueTournBattleFailSettleInfoScNotify) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(RogueTournGetArchiveRepositoryCsReq, func() any { return new(proto.RogueTournGetArchiveRepositoryCsReq) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(FightMatch3ForceUpdateNotify, func() any { return new(proto.FightMatch3ForceUpdateNotify) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(FightMatch3ChatCsReq, func() any { return new(proto.FightMatch3ChatCsReq) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(FightKickOutScNotify, func() any { return new(proto.FightKickOutScNotify) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(FightMatch3DataScRsp, func() any { return new(proto.FightMatch3DataScRsp) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(SetFriendMarkScRsp, func() any { return new(proto.SetFriendMarkScRsp) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(FightMatch3TurnEndScNotify, func() any { return new(proto.FightMatch3TurnEndScNotify) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(SyncRogueCommonDialogueDataScNotify, func() any { return new(proto.SyncRogueCommonDialogueDataScNotify) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(FightMatch3TurnStartScNotify, func() any { return new(proto.FightMatch3TurnStartScNotify) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(FightMatch3StartCountDownScNotify, func() any { return new(proto.FightMatch3StartCountDownScNotify) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(GetAuthkeyScRsp, func() any { return new(proto.GetAuthkeyScRsp) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(FightEnterCsReq, func() any { return new(proto.FightEnterCsReq) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(SpringRecoverCsReq, func() any { return new(proto.SpringRecoverCsReq) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(RogueTournResetPermanentTalentScRsp, func() any { return new(proto.RogueTournResetPermanentTalentScRsp) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(ClientDownloadDataScNotify, func() any { return new(proto.ClientDownloadDataScNotify) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(FightHeartBeatCsReq, func() any { return new(proto.FightHeartBeatCsReq) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(GetHeroBasicTypeInfoCsReq, func() any { return new(proto.GetHeroBasicTypeInfoCsReq) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(RogueTournSettleCsReq, func() any { return new(proto.RogueTournSettleCsReq) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(GetFriendAssistListCsReq, func() any { return new(proto.GetFriendAssistListCsReq) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(RogueTournSettleScRsp, func() any { return new(proto.RogueTournSettleScRsp) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(DailyTaskDataScNotify, func() any { return new(proto.DailyTaskDataScNotify) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(FightMatch3ChatScNotify, func() any { return new(proto.FightMatch3ChatScNotify) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(RogueTournLeaveScRsp, func() any { return new(proto.RogueTournLeaveScRsp) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(TakeFightActivityRewardCsReq, func() any { return new(proto.TakeFightActivityRewardCsReq) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(GetFriendDevelopmentInfoCsReq, func() any { return new(proto.GetFriendDevelopmentInfoCsReq) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(RogueTournResetPermanentTalentCsReq, func() any { return new(proto.RogueTournResetPermanentTalentCsReq) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(StartPartialChallengeScRsp, func() any { return new(proto.StartPartialChallengeScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(RogueTournReEnterRogueCocoonStageCsReq, func() any { return new(proto.RogueTournReEnterRogueCocoonStageCsReq) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(EnterChallengeNextPhaseCsReq, func() any { return new(proto.EnterChallengeNextPhaseCsReq) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(RogueTournEnterLayerCsReq, func() any { return new(proto.RogueTournEnterLayerCsReq) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(SetForbidOtherApplyFriendScRsp, func() any { return new(proto.SetForbidOtherApplyFriendScRsp) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(FightLeaveScNotify, func() any { return new(proto.FightLeaveScNotify) })
	c.regMsg(FightMatch3ChatScRsp, func() any { return new(proto.FightMatch3ChatScRsp) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(SubmitMonsterResearchActivityMaterialScRsp, func() any { return new(proto.SubmitMonsterResearchActivityMaterialScRsp) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(RogueGetGambleInfoCsReq, func() any { return new(proto.RogueGetGambleInfoCsReq) })
	c.regMsg(RogueGetGambleInfoScRsp, func() any { return new(proto.RogueGetGambleInfoScRsp) })
	c.regMsg(RogueDoGambleCsReq, func() any { return new(proto.RogueDoGambleCsReq) })
	c.regMsg(RogueDoGambleScRsp, func() any { return new(proto.RogueDoGambleScRsp) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
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
	c.regMsg(GmMission, func() any { return new(spb.GmMission) })
	c.regMsg(PlayerMsgGateToNodeNotify, func() any { return new(spb.PlayerMsgGateToNodeNotify) })
	c.regMsg(GameToNodePingReq, func() any { return new(spb.GameToNodePingReq) })
	c.regMsg(GameToNodePingRsp, func() any { return new(spb.GameToNodePingRsp) })
	c.regMsg(GateToNodePingReq, func() any { return new(spb.GateToNodePingReq) })
	c.regMsg(GateToNodePingRsp, func() any { return new(spb.GateToNodePingRsp) })
}
