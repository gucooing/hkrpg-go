package cmd

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	PlayerTypeNone                                     = 0
	PlayerGetTokenCsReq                                = 2
	GateServerScNotify                                 = 3
	AceAntiCheaterCsReq                                = 4
	ExchangeStaminaCsReq                               = 6
	UpdatePlayerSettingCsReq                           = 7
	PlayerGetTokenScRsp                                = 9
	GetVideoVersionKeyScRsp                            = 10
	GetSecretKeyInfoScRsp                              = 12
	GetVideoVersionKeyCsReq                            = 13
	ReserveStaminaExchangeCsReq                        = 14
	PlayerLoginFinishCsReq                             = 15
	SetNicknameScRsp                                   = 16
	ServerAnnounceNotify                               = 18
	UpdatePlayerSettingScRsp                           = 20
	GetSecretKeyInfoCsReq                              = 22
	GetHeroBasicTypeInfoCsReq                          = 24
	SetGenderScRsp                                     = 25
	SetLanguageCsReq                                   = 28
	GmTalkCsReq                                        = 29
	GetLevelRewardTakenListCsReq                       = 30
	ExchangeStaminaScRsp                               = 33
	PlayerLoginCsReq                                   = 34
	SetGameplayBirthdayCsReq                           = 35
	AntiAddictScNotify                                 = 37
	SetNicknameCsReq                                   = 39
	ReserveStaminaExchangeScRsp                        = 40
	RegionStopScNotify                                 = 42
	GmTalkScNotify                                     = 43
	SetGameplayBirthdayScRsp                           = 44
	GmTalkScRsp                                        = 45
	PlayerLoginScRsp                                   = 48
	PlayerHeartBeatScRsp                               = 49
	DailyRefreshNotify                                 = 51
	UpdateFeatureSwitchScNotify                        = 55
	GetLevelRewardCsReq                                = 56
	ClientObjDownloadDataScNotify                      = 58
	GetAuthkeyCsReq                                    = 59
	SetLanguageScRsp                                   = 61
	PlayerLogoutCsReq                                  = 62
	GetLevelRewardScRsp                                = 63
	ClientObjUploadCsReq                               = 64
	SetPlayerInfoScRsp                                 = 65
	GetBasicInfoCsReq                                  = 66
	QueryProductInfoScRsp                              = 67
	StaminaInfoScNotify                                = 69
	PlayerHeartBeatCsReq                               = 71
	PlayerLoginFinishScRsp                             = 72
	GetBasicInfoScRsp                                  = 73
	AceAntiCheaterScRsp                                = 75
	ClientObjUploadScRsp                               = 78
	SetGenderCsReq                                     = 79
	GetHeroBasicTypeInfoScRsp                          = 82
	SetRedPointStatusScNotify                          = 84
	GetLevelRewardTakenListScRsp                       = 85
	PlayerKickOutScNotify                              = 86
	PlayerLogoutScRsp                                  = 88
	HeroBasicTypeChangedNotify                         = 89
	QueryProductInfoCsReq                              = 90
	SetHeroBasicTypeCsReq                              = 91
	ClientDownloadDataScNotify                         = 92
	MonthCardRewardNotify                              = 93
	FeatureSwitchClosedScNotify                        = 94
	GetAuthkeyScRsp                                    = 95
	SetHeroBasicTypeScRsp                              = 97
	RetcodeNotify                                      = 98
	SetPlayerInfoCsReq                                 = 100
	GetCurBattleInfoCsReq                              = 102
	GetCurBattleInfoScRsp                              = 109
	SyncClientResVersionCsReq                          = 119
	BattleLogReportCsReq                               = 129
	PVEBattleResultCsReq                               = 134
	SyncClientResVersionScRsp                          = 143
	BattleLogReportScRsp                               = 145
	PVEBattleResultScRsp                               = 148
	QuitBattleCsReq                                    = 162
	ServerSimulateBattleFinishScNotify                 = 168
	QuitBattleScNotify                                 = 186
	QuitBattleScRsp                                    = 188
	ReBattleAfterBattleLoseCsNotify                    = 196
	UnlockAvatarSkinScNotify                           = 301
	UnlockSkilltreeCsReq                               = 302
	RankUpAvatarCsReq                                  = 306
	UnlockSkilltreeScRsp                               = 309
	TakePromotionRewardScRsp                           = 316
	PromoteAvatarCsReq                                 = 319
	MarkAvatarScRsp                                    = 328
	DressAvatarScRsp                                   = 329
	DressAvatarSkinCsReq                               = 330
	RankUpAvatarScRsp                                  = 333
	GetAvatarDataCsReq                                 = 334
	TakeOffRelicScRsp                                  = 337
	TakePromotionRewardCsReq                           = 339
	MarkAvatarCsReq                                    = 341
	TakeOffRelicCsReq                                  = 342
	PromoteAvatarScRsp                                 = 343
	TakeOffEquipmentCsReq                              = 345
	GetAvatarDataScRsp                                 = 348
	TakeOffAvatarSkinCsReq                             = 356
	DressRelicAvatarCsReq                              = 359
	AvatarExpUpCsReq                                   = 362
	TakeOffAvatarSkinScRsp                             = 363
	TakeOffEquipmentScRsp                              = 368
	DressAvatarSkinScRsp                               = 385
	DressAvatarCsReq                                   = 386
	AvatarExpUpScRsp                                   = 388
	DressRelicAvatarScRsp                              = 395
	AddAvatarScNotify                                  = 396
	GetChapterCsReq                                    = 402
	GetChapterScRsp                                    = 409
	WaypointShowNewCsNotify                            = 419
	GetWaypointCsReq                                   = 434
	TakeChapterRewardCsReq                             = 443
	GetWaypointScRsp                                   = 448
	SetCurWaypointCsReq                                = 462
	TakeChapterRewardScRsp                             = 486
	SetCurWaypointScRsp                                = 488
	AddEquipmentScNotify                               = 501
	LockEquipmentCsReq                                 = 502
	ComposeItemScRsp                                   = 506
	MarkItemScRsp                                      = 508
	LockEquipmentScRsp                                 = 509
	MarkItemCsReq                                      = 511
	RechargeSuccNotify                                 = 516
	ComposeLimitNumUpdateNotify                        = 518
	UseItemCsReq                                       = 519
	GetMarkItemListCsReq                               = 524
	SetTurnFoodSwitchCsReq                             = 525
	GetRecyleTimeScRsp                                 = 528
	RankUpEquipmentScRsp                               = 529
	ExchangeHcoinCsReq                                 = 530
	ExpUpRelicCsReq                                    = 533
	GetBagCsReq                                        = 534
	SellItemCsReq                                      = 537
	SellItemScRsp                                      = 539
	GetRecyleTimeCsReq                                 = 541
	LockRelicScRsp                                     = 542
	UseItemScRsp                                       = 543
	ExpUpEquipmentCsReq                                = 545
	GetBagScRsp                                        = 548
	CancelMarkItemNotify                               = 554
	ComposeSelectedRelicCsReq                          = 556
	ExpUpRelicScRsp                                    = 559
	ComposeLimitNumCompleteNotify                      = 561
	PromoteEquipmentCsReq                              = 562
	ComposeSelectedRelicScRsp                          = 563
	GeneralVirtualItemDataNotify                       = 565
	RelicRecommendCsReq                                = 567
	ExpUpEquipmentScRsp                                = 568
	SyncTurnFoodNotify                                 = 579
	GetMarkItemListScRsp                               = 582
	ExchangeHcoinScRsp                                 = 585
	RankUpEquipmentCsReq                               = 586
	PromoteEquipmentScRsp                              = 588
	DiscardRelicCsReq                                  = 589
	DiscardRelicScRsp                                  = 590
	DestroyItemCsReq                                   = 591
	RelicRecommendScRsp                                = 592
	LockRelicCsReq                                     = 595
	ComposeItemCsReq                                   = 596
	DestroyItemScRsp                                   = 597
	SetTurnFoodSwitchScRsp                             = 600
	PlayerSyncScNotify                                 = 634
	JoinLineupCsReq                                    = 702
	ChangeLineupLeaderCsReq                            = 706
	JoinLineupScRsp                                    = 709
	GetAllLineupDataScRsp                              = 716
	QuitLineupCsReq                                    = 719
	SwapLineupScRsp                                    = 729
	VirtualLineupDestroyNotify                         = 730
	ChangeLineupLeaderScRsp                            = 733
	GetStageLineupCsReq                                = 734
	SetLineupNameScRsp                                 = 737
	GetAllLineupDataCsReq                              = 739
	SetLineupNameCsReq                                 = 742
	QuitLineupScRsp                                    = 743
	SyncLineupNotify                                   = 745
	GetStageLineupScRsp                                = 748
	ReplaceLineupScRsp                                 = 756
	SwitchLineupIndexCsReq                             = 759
	GetCurLineupDataCsReq                              = 762
	ExtraLineupDestroyNotify                           = 763
	GetLineupAvatarDataCsReq                           = 768
	ReplaceLineupCsReq                                 = 785
	SwapLineupCsReq                                    = 786
	GetCurLineupDataScRsp                              = 788
	SwitchLineupIndexScRsp                             = 795
	GetLineupAvatarDataScRsp                           = 796
	DelMailCsReq                                       = 802
	DelMailScRsp                                       = 809
	TakeMailAttachmentCsReq                            = 819
	GetMailCsReq                                       = 834
	TakeMailAttachmentScRsp                            = 843
	GetMailScRsp                                       = 848
	MarkReadMailCsReq                                  = 862
	NewMailScNotify                                    = 886
	MarkReadMailScRsp                                  = 888
	GetQuestRecordCsReq                                = 919
	FinishQuestCsReq                                   = 929
	BatchGetQuestDataCsReq                             = 933
	GetQuestDataCsReq                                  = 934
	GetQuestRecordScRsp                                = 943
	FinishQuestScRsp                                   = 945
	GetQuestDataScRsp                                  = 948
	BatchGetQuestDataScRsp                             = 959
	TakeQuestRewardCsReq                               = 962
	TakeQuestOptionalRewardCsReq                       = 968
	QuestRecordScNotify                                = 986
	TakeQuestRewardScRsp                               = 988
	TakeQuestOptionalRewardScRsp                       = 996
	FinishPlotCsReq                                    = 1134
	FinishPlotScRsp                                    = 1148
	SubMissionRewardScNotify                           = 1201
	MissionRewardScNotify                              = 1202
	FinishCosumeItemMissionScRsp                       = 1206
	SyncTaskCsReq                                      = 1209
	MissionAcceptScNotify                              = 1211
	GetMissionStatusScRsp                              = 1216
	StartFinishMainMissionScNotify                     = 1218
	SyncTaskScRsp                                      = 1219
	GetMainMissionCustomValueCsReq                     = 1224
	TeleportToMissionResetPointScRsp                   = 1228
	InterruptMissionEventCsReq                         = 1230
	GetMissionEventDataCsReq                           = 1233
	GetMissionDataCsReq                                = 1234
	AcceptMissionEventScRsp                            = 1237
	GetMissionStatusCsReq                              = 1239
	TeleportToMissionResetPointCsReq                   = 1241
	AcceptMissionEventCsReq                            = 1242
	DailyTaskDataScNotify                              = 1243
	GetMissionDataScRsp                                = 1248
	UpdateTrackMainMissionIdCsReq                      = 1254
	SetMissionEventProgressCsReq                       = 1256
	GetMissionEventDataScRsp                           = 1259
	StartFinishSubMissionScNotify                      = 1261
	FinishTalkMissionCsReq                             = 1262
	SetMissionEventProgressScRsp                       = 1263
	MissionGroupWarnScNotify                           = 1268
	UpdateTrackMainMissionIdScRsp                      = 1279
	GetMainMissionCustomValueScRsp                     = 1282
	InterruptMissionEventScRsp                         = 1285
	FinishTalkMissionScRsp                             = 1288
	AcceptMainMissionCsReq                             = 1291
	MissionEventRewardScNotify                         = 1295
	FinishCosumeItemMissionCsReq                       = 1296
	AcceptMainMissionScRsp                             = 1297
	EnterAdventureCsReq                                = 1334
	EnterAdventureScRsp                                = 1348
	GetFarmStageGachaInfoCsReq                         = 1362
	GetFarmStageGachaInfoScRsp                         = 1388
	SceneCastSkillCsReq                                = 1402
	SetGroupCustomSaveDataScRsp                        = 1403
	SpringRecoverScRsp                                 = 1404
	ReEnterLastElementStageCsReq                       = 1405
	SceneCastSkillCostMpCsReq                          = 1406
	StartCocoonStageCsReq                              = 1408
	SceneCastSkillScRsp                                = 1409
	EnterSceneByServerScNotify                         = 1410
	SavePointsInfoNotify                               = 1411
	SceneEntityTeleportCsReq                           = 1412
	EnterSceneScRsp                                    = 1413
	SyncServerSceneChangeNotify                        = 1414
	SceneEntityTeleportScRsp                           = 1415
	ReturnLastTownCsReq                                = 1416
	DeleteSummonUnitCsReq                              = 1417
	GetCurSceneInfoCsReq                               = 1419
	UpdateFloorSavedValueNotify                        = 1420
	GroupStateChangeScRsp                              = 1421
	ReEnterLastElementStageScRsp                       = 1422
	StartTimedCocoonStageScRsp                         = 1423
	RecoverAllLineupCsReq                              = 1424
	EntityBindPropScRsp                                = 1425
	StartTimedCocoonStageCsReq                         = 1426
	GetEnteredSceneCsReq                               = 1427
	EnterSectionScRsp                                  = 1428
	ReturnLastTownScRsp                                = 1430
	GetEnteredSceneScRsp                               = 1431
	RefreshTriggerByClientCsReq                        = 1432
	SceneCastSkillCostMpScRsp                          = 1433
	SceneEntityMoveCsReq                               = 1434
	SetSpringRecoverConfigScRsp                        = 1435
	StartTimedFarmElementCsReq                         = 1436
	SpringRefreshScRsp                                 = 1437
	RefreshTriggerByClientScRsp                        = 1438
	LastSpringRefreshTimeNotify                        = 1439
	GetUnlockTeleportCsReq                             = 1440
	EnterSectionCsReq                                  = 1441
	SpringRefreshCsReq                                 = 1442
	GetCurSceneInfoScRsp                               = 1443
	SpringRecoverCsReq                                 = 1444
	SceneEntityMoveScNotify                            = 1445
	GroupStateChangeScNotify                           = 1447
	SceneEntityMoveScRsp                               = 1448
	SetGroupCustomSaveDataCsReq                        = 1449
	EnteredSceneChangeScNotify                         = 1450
	SetSpringRecoverConfigCsReq                        = 1451
	GameplayCounterRecoverCsReq                        = 1452
	StartCocoonStageScRsp                              = 1454
	ActivateFarmElementScRsp                           = 1455
	SceneEnterStageScRsp                               = 1456
	UnlockedAreaMapScNotify                            = 1457
	GameplayCounterCountDownCsReq                      = 1458
	SceneCastSkillMpUpdateScNotify                     = 1459
	StartTimedFarmElementScRsp                         = 1460
	InteractPropCsReq                                  = 1462
	GameplayCounterCountDownScRsp                      = 1464
	SetClientPausedScRsp                               = 1465
	GetSpringRecoverDataCsReq                          = 1466
	DeactivateFarmElementScRsp                         = 1467
	SceneUpdatePositionVersionNotify                   = 1468
	GetUnlockTeleportScRsp                             = 1469
	GetSceneMapInfoCsReq                               = 1470
	UpdateMechanismBarScNotify                         = 1471
	EnterSceneCsReq                                    = 1472
	GetSpringRecoverDataScRsp                          = 1473
	RefreshTriggerByClientScNotify                     = 1474
	HealPoolInfoNotify                                 = 1475
	GroupStateChangeCsReq                              = 1476
	SceneGroupRefreshScNotify                          = 1477
	GameplayCounterUpdateScNotify                      = 1478
	EntityBindPropCsReq                                = 1479
	GameplayCounterRecoverScRsp                        = 1481
	RecoverAllLineupScRsp                              = 1482
	UnlockTeleportNotify                               = 1483
	ScenePlaneEventScNotify                            = 1484
	SceneEnterStageCsReq                               = 1485
	DeleteSummonUnitScRsp                              = 1487
	InteractPropScRsp                                  = 1488
	DeactivateFarmElementCsReq                         = 1490
	SetCurInteractEntityCsReq                          = 1491
	ActivateFarmElementCsReq                           = 1492
	SpringRecoverSingleAvatarCsReq                     = 1493
	SyncEntityBuffChangeListScNotify                   = 1496
	SetCurInteractEntityScRsp                          = 1497
	SpringRecoverSingleAvatarScRsp                     = 1498
	GetSceneMapInfoScRsp                               = 1499
	SetClientPausedCsReq                               = 1500
	TakeCityShopRewardCsReq                            = 1502
	TakeCityShopRewardScRsp                            = 1509
	CityShopInfoScNotify                               = 1519
	GetShopListCsReq                                   = 1534
	GetShopListScRsp                                   = 1548
	BuyGoodsCsReq                                      = 1562
	BuyGoodsScRsp                                      = 1588
	UnlockTutorialCsReq                                = 1602
	UnlockTutorialScRsp                                = 1609
	UnlockTutorialGuideCsReq                           = 1619
	FinishTutorialScRsp                                = 1629
	GetTutorialCsReq                                   = 1634
	UnlockTutorialGuideScRsp                           = 1643
	FinishTutorialGuideCsReq                           = 1645
	GetTutorialScRsp                                   = 1648
	GetTutorialGuideCsReq                              = 1662
	FinishTutorialGuideScRsp                           = 1668
	FinishTutorialCsReq                                = 1686
	GetTutorialGuideScRsp                              = 1688
	LeaveChallengeCsReq                                = 1702
	LeaveChallengeScRsp                                = 1709
	ChallengeSettleNotify                              = 1719
	GetCurChallengeCsReq                               = 1729
	TakeChallengeRewardCsReq                           = 1733
	GetChallengeCsReq                                  = 1734
	GetChallengeGroupStatisticsScRsp                   = 1742
	GetCurChallengeScRsp                               = 1745
	GetChallengeScRsp                                  = 1748
	TakeChallengeRewardScRsp                           = 1759
	StartChallengeCsReq                                = 1762
	ChallengeLineupNotify                              = 1768
	StartChallengeScRsp                                = 1788
	GetChallengeGroupStatisticsCsReq                   = 1795
	EnhanceRogueBuffScRsp                              = 1801
	EnterRogueCsReq                                    = 1802
	FinishRogueDialogueGroupCsReq                      = 1804
	SyncRogueSeasonFinishScNotify                      = 1808
	EnterRogueScRsp                                    = 1809
	SyncRogueAeonScNotify                              = 1810
	SyncRogueExploreWinScNotify                        = 1811
	SyncRogueDialogueEventDataScNotify                 = 1813
	TakeRogueAeonLevelRewardScRsp                      = 1814
	SelectRogueDialogueEventCsReq                      = 1815
	TakeRogueScoreRewardCsReq                          = 1816
	EnableRogueTalentScRsp                             = 1817
	LeaveRogueCsReq                                    = 1819
	SyncRogueAeonLevelUpRewardScNotify                 = 1820
	QuitRogueScRsp                                     = 1824
	EnterRogueMapRoomCsReq                             = 1825
	GetRogueAeonInfoScRsp                              = 1827
	TakeRogueScoreRewardScRsp                          = 1830
	FinishAeonDialogueGroupCsReq                       = 1831
	GetRogueTalentInfoCsReq                            = 1832
	SyncRogueFinishScNotify                            = 1833
	GetRogueInfoCsReq                                  = 1834
	GetRogueDialogueEventDataCsReq                     = 1835
	SyncRogueVirtualItemInfoScNotify                   = 1836
	ReviveRogueAvatarCsReq                             = 1837
	GetRogueTalentInfoScRsp                            = 1838
	ReviveRogueAvatarScRsp                             = 1839
	LeaveRogueScRsp                                    = 1843
	GetRogueDialogueEventDataScRsp                     = 1844
	GetRogueAeonInfoCsReq                              = 1847
	GetRogueInfoScRsp                                  = 1848
	ExchangeRogueRewardKeyScRsp                        = 1849
	FinishAeonDialogueGroupScRsp                       = 1850
	GetRogueBuffEnhanceInfoScRsp                       = 1856
	GetRogueScoreRewardInfoCsReq                       = 1858
	PickRogueAvatarCsReq                               = 1859
	SyncRogueStatusScNotify                            = 1860
	StartRogueCsReq                                    = 1862
	EnhanceRogueBuffCsReq                              = 1863
	GetRogueScoreRewardInfoScRsp                       = 1864
	GetRogueInitialScoreCsReq                          = 1865
	SyncRogueGetItemScNotify                           = 1870
	ExchangeRogueRewardKeyCsReq                        = 1871
	SelectRogueDialogueEventScRsp                      = 1872
	EnableRogueTalentCsReq                             = 1874
	FinishRogueDialogueGroupScRsp                      = 1875
	SyncRoguePickAvatarInfoScNotify                    = 1880
	SyncRogueRewardInfoScNotify                        = 1883
	SyncRogueAreaUnlockScNotify                        = 1884
	GetRogueBuffEnhanceInfoCsReq                       = 1885
	StartRogueScRsp                                    = 1888
	GetRogueInitialScoreScRsp                          = 1889
	SyncRogueMapRoomScNotify                           = 1890
	SyncRogueReviveInfoScNotify                        = 1891
	OpenRogueChestCsReq                                = 1893
	PickRogueAvatarScRsp                               = 1895
	QuitRogueCsReq                                     = 1897
	OpenRogueChestScRsp                                = 1898
	TakeRogueAeonLevelRewardCsReq                      = 1899
	EnterRogueMapRoomScRsp                             = 1900
	GetGachaCeilingCsReq                               = 1902
	GetGachaCeilingScRsp                               = 1909
	ExchangeGachaCeilingCsReq                          = 1919
	GetGachaInfoCsReq                                  = 1934
	ExchangeGachaCeilingScRsp                          = 1943
	GetGachaInfoScRsp                                  = 1948
	DoGachaCsReq                                       = 1962
	DoGachaScRsp                                       = 1988
	GetFirstTalkNpcCsReq                               = 2102
	FinishFirstTalkByPerformanceNpcScRsp               = 2106
	GetFirstTalkNpcScRsp                               = 2109
	FinishFirstTalkNpcCsReq                            = 2119
	SelectInclinationTextScRsp                         = 2129
	GetNpcTakenRewardCsReq                             = 2134
	FinishFirstTalkNpcScRsp                            = 2143
	GetFirstTalkByPerformanceNpcCsReq                  = 2145
	GetNpcTakenRewardScRsp                             = 2148
	TakeTalkRewardCsReq                                = 2162
	GetFirstTalkByPerformanceNpcScRsp                  = 2168
	SelectInclinationTextCsReq                         = 2186
	TakeTalkRewardScRsp                                = 2188
	FinishFirstTalkByPerformanceNpcCsReq               = 2196
	RaidInfoNotify                                     = 2202
	SetClientRaidTargetCountScRsp                      = 2206
	GetChallengeRaidInfoCsReq                          = 2209
	GetChallengeRaidInfoScRsp                          = 2219
	ChallengeRaidNotify                                = 2229
	GetSaveRaidCsReq                                   = 2233
	StartRaidCsReq                                     = 2234
	DelSaveRaidScNotify                                = 2237
	RaidKickByServerScNotify                           = 2239
	GetAllSaveRaidScRsp                                = 2242
	TakeChallengeRaidRewardCsReq                       = 2243
	GetRaidInfoCsReq                                   = 2245
	StartRaidScRsp                                     = 2248
	GetSaveRaidScRsp                                   = 2259
	LeaveRaidCsReq                                     = 2262
	GetRaidInfoScRsp                                   = 2268
	TakeChallengeRaidRewardScRsp                       = 2286
	LeaveRaidScRsp                                     = 2288
	GetAllSaveRaidCsReq                                = 2295
	SetClientRaidTargetCountCsReq                      = 2296
	GetArchiveDataCsReq                                = 2334
	GetArchiveDataScRsp                                = 2348
	GetUpdatedArchiveDataCsReq                         = 2362
	GetUpdatedArchiveDataScRsp                         = 2388
	CancelExpeditionCsReq                              = 2502
	TakeActivityExpeditionRewardCsReq                  = 2506
	CancelExpeditionScRsp                              = 2509
	TakeExpeditionRewardCsReq                          = 2519
	AcceptActivityExpeditionCsReq                      = 2529
	TakeActivityExpeditionRewardScRsp                  = 2533
	GetExpeditionDataCsReq                             = 2534
	TakeMultipleExpeditionRewardScRsp                  = 2537
	TakeMultipleExpeditionRewardCsReq                  = 2542
	TakeExpeditionRewardScRsp                          = 2543
	AcceptActivityExpeditionScRsp                      = 2545
	GetExpeditionDataScRsp                             = 2548
	AcceptMultipleExpeditionCsReq                      = 2559
	AcceptExpeditionCsReq                              = 2562
	CancelActivityExpeditionCsReq                      = 2568
	ExpeditionDataChangeScNotify                       = 2586
	AcceptExpeditionScRsp                              = 2588
	AcceptMultipleExpeditionScRsp                      = 2595
	CancelActivityExpeditionScRsp                      = 2596
	GetActivityScheduleConfigCsReq                     = 2602
	StartTrialActivityScRsp                            = 2603
	TrialActivityDataChangeScNotify                    = 2604
	CurTrialActivityScNotify                           = 2605
	GetActivityScheduleConfigScRsp                     = 2609
	TakeMonsterResearchActivityRewardCsReq             = 2616
	TakeMonsterResearchActivityRewardScRsp             = 2630
	GetLoginActivityCsReq                              = 2634
	GetTrialActivityDataCsReq                          = 2635
	SubmitMonsterResearchActivityMaterialCsReq         = 2637
	SubmitMonsterResearchActivityMaterialScRsp         = 2639
	GetMonsterResearchActivityDataScRsp                = 2642
	GetTrialActivityDataScRsp                          = 2644
	LeaveTrialActivityScRsp                            = 2646
	GetLoginActivityScRsp                              = 2648
	StartTrialActivityCsReq                            = 2649
	TakeLoginActivityRewardCsReq                       = 2662
	TakeTrialActivityRewardScRsp                       = 2671
	EnterTrialActivityStageCsReq                       = 2675
	TakeLoginActivityRewardScRsp                       = 2688
	EnterTrialActivityStageScRsp                       = 2693
	LeaveTrialActivityCsReq                            = 2694
	GetMonsterResearchActivityDataCsReq                = 2695
	TakeTrialActivityRewardCsReq                       = 2698
	FinishItemIdCsReq                                  = 2702
	FinishItemIdScRsp                                  = 2709
	FinishSectionIdCsReq                               = 2719
	FinishPerformSectionIdScRsp                        = 2729
	GetNpcMessageGroupCsReq                            = 2734
	FinishSectionIdScRsp                               = 2743
	GetNpcMessageGroupScRsp                            = 2748
	GetNpcStatusCsReq                                  = 2762
	FinishPerformSectionIdCsReq                        = 2786
	GetNpcStatusScRsp                                  = 2788
	SetDisplayAvatarCsReq                              = 2802
	SetDisplayAvatarScRsp                              = 2809
	SetIsDisplayAvatarInfoCsReq                        = 2819
	SetSignatureCsReq                                  = 2829
	GetPlayerBoardDataCsReq                            = 2834
	SetIsDisplayAvatarInfoScRsp                        = 2843
	SetSignatureScRsp                                  = 2845
	GetPlayerBoardDataScRsp                            = 2848
	SetHeadIconCsReq                                   = 2862
	SetAssistAvatarCsReq                               = 2868
	UnlockHeadIconScNotify                             = 2886
	SetHeadIconScRsp                                   = 2888
	SetAssistAvatarScRsp                               = 2896
	DeleteBlacklistScRsp                               = 2901
	GetFriendApplyListInfoCsReq                        = 2902
	GetFriendDevelopmentInfoScRsp                      = 2903
	GetFriendChallengeLineupScRsp                      = 2904
	DeleteFriendScRsp                                  = 2906
	GetAssistHistoryScRsp                              = 2908
	GetFriendApplyListInfoScRsp                        = 2909
	GetAssistHistoryCsReq                              = 2911
	SetFriendRemarkNameCsReq                           = 2916
	GetAssistListScRsp                                 = 2918
	ApplyFriendCsReq                                   = 2919
	GetCurAssistCsReq                                  = 2924
	TakeAssistRewardScRsp                              = 2925
	SearchPlayerScRsp                                  = 2928
	HandleFriendCsReq                                  = 2929
	SetFriendRemarkNameScRsp                           = 2930
	SyncDeleteFriendScNotify                           = 2933
	GetFriendListInfoCsReq                             = 2934
	GetFriendAssistListScRsp                           = 2935
	GetFriendRecommendListInfoCsReq                    = 2937
	GetFriendRecommendListInfoScRsp                    = 2939
	SearchPlayerCsReq                                  = 2941
	SyncAddBlacklistScNotify                           = 2942
	ApplyFriendScRsp                                   = 2943
	GetFriendChallengeLineupCsReq                      = 2944
	HandleFriendScRsp                                  = 2945
	GetFriendListInfoScRsp                             = 2948
	GetFriendDevelopmentInfoCsReq                      = 2949
	GetFriendAssistListCsReq                           = 2951
	NewAssistHistoryNotify                             = 2954
	SetForbidOtherApplyFriendScRsp                     = 2955
	ReportPlayerScRsp                                  = 2956
	AddBlacklistCsReq                                  = 2959
	GetAssistListCsReq                                 = 2961
	GetPlayerDetailInfoCsReq                           = 2962
	DeleteBlacklistCsReq                               = 2963
	GetPlatformPlayerInfoCsReq                         = 2965
	SetFriendMarkCsReq                                 = 2966
	GetFriendLoginInfoScRsp                            = 2967
	SyncHandleFriendScNotify                           = 2968
	GetFriendBattleRecordDetailScRsp                   = 2971
	SetFriendMarkScRsp                                 = 2973
	GetFriendChallengeDetailCsReq                      = 2975
	TakeAssistRewardCsReq                              = 2979
	GetCurAssistScRsp                                  = 2982
	ReportPlayerCsReq                                  = 2985
	SyncApplyFriendScNotify                            = 2986
	GetPlayerDetailInfoScRsp                           = 2988
	GetPlatformPlayerInfoScRsp                         = 2989
	GetFriendLoginInfoCsReq                            = 2990
	SetAssistCsReq                                     = 2991
	SetForbidOtherApplyFriendCsReq                     = 2992
	GetFriendChallengeDetailScRsp                      = 2993
	AddBlacklistScRsp                                  = 2995
	DeleteFriendCsReq                                  = 2996
	SetAssistScRsp                                     = 2997
	GetFriendBattleRecordDetailCsReq                   = 2998
	CurAssistChangedNotify                             = 3000
	TakeBpRewardScRsp                                  = 3002
	BuyBpLevelCsReq                                    = 3009
	BuyBpLevelScRsp                                    = 3019
	BattlePassInfoNotify                               = 3034
	TakeAllRewardCsReq                                 = 3043
	TakeAllRewardScRsp                                 = 3086
	TakeBpRewardCsReq                                  = 3088
	UnlockBackGroundMusicCsReq                         = 3102
	UnlockBackGroundMusicScRsp                         = 3109
	TrialBackGroundMusicCsReq                          = 3119
	GetJukeboxDataCsReq                                = 3134
	TrialBackGroundMusicScRsp                          = 3143
	GetJukeboxDataScRsp                                = 3148
	PlayBackGroundMusicCsReq                           = 3162
	PlayBackGroundMusicScRsp                           = 3188
	SharePunkLordMonsterCsReq                          = 3202
	TakePunkLordPointRewardScRsp                       = 3206
	SharePunkLordMonsterScRsp                          = 3209
	TakeKilledPunkLordMonsterScoreScRsp                = 3218
	SummonPunkLordMonsterCsReq                         = 3219
	GetPunkLordBattleRecordScRsp                       = 3224
	PunkLordMonsterKilledNotify                        = 3228
	PunkLordMonsterInfoScNotify                        = 3233
	GetPunkLordMonsterDataCsReq                        = 3234
	PunkLordRaidTimeOutScNotify                        = 3237
	SummonPunkLordMonsterScRsp                         = 3243
	GetPunkLordMonsterDataScRsp                        = 3248
	GetKilledPunkLordMonsterDataCsReq                  = 3256
	GetPunkLordDataCsReq                               = 3259
	TakeKilledPunkLordMonsterScoreCsReq                = 3261
	StartPunkLordRaidCsReq                             = 3262
	GetKilledPunkLordMonsterDataScRsp                  = 3263
	PunkLordBattleResultScNotify                       = 3285
	StartPunkLordRaidScRsp                             = 3288
	PunkLordDataChangeNotify                           = 3291
	GetPunkLordDataScRsp                               = 3295
	TakePunkLordPointRewardCsReq                       = 3296
	GetPunkLordBattleRecordCsReq                       = 3297
	DailyActiveInfoNotify                              = 3302
	TakeAllApRewardCsReq                               = 3309
	TakeAllApRewardScRsp                               = 3319
	TakeApRewardCsReq                                  = 3334
	TakeApRewardScRsp                                  = 3348
	GetDailyActiveInfoCsReq                            = 3362
	GetDailyActiveInfoScRsp                            = 3388
	GetRndOptionCsReq                                  = 3434
	GetRndOptionScRsp                                  = 3448
	DailyFirstMeetPamCsReq                             = 3462
	DailyFirstMeetPamScRsp                             = 3488
	GetReplayTokenCsReq                                = 3534
	GetReplayTokenScRsp                                = 3548
	GetPlayerReplayInfoCsReq                           = 3562
	GetPlayerReplayInfoScRsp                           = 3588
	EnterFightActivityStageScRsp                       = 3602
	TakeFightActivityRewardCsReq                       = 3609
	TakeFightActivityRewardScRsp                       = 3619
	GetFightActivityDataCsReq                          = 3634
	GetFightActivityDataScRsp                          = 3648
	FightActivityDataChangeScNotify                    = 3662
	EnterFightActivityStageCsReq                       = 3688
	TrainRefreshTimeNotify                             = 3702
	TrainVisitorRewardSendNotify                       = 3709
	GetTrainVisitorRegisterCsReq                       = 3719
	TakeTrainVisitorUntakenBehaviorRewardScRsp         = 3729
	TrainVisitorBehaviorFinishCsReq                    = 3734
	GetTrainVisitorRegisterScRsp                       = 3743
	ShowNewSupplementVisitorCsReq                      = 3745
	TrainVisitorBehaviorFinishScRsp                    = 3748
	GetTrainVisitorBehaviorCsReq                       = 3762
	ShowNewSupplementVisitorScRsp                      = 3768
	TakeTrainVisitorUntakenBehaviorRewardCsReq         = 3786
	GetTrainVisitorBehaviorScRsp                       = 3788
	TextJoinBatchSaveCsReq                             = 3802
	TextJoinBatchSaveScRsp                             = 3809
	TextJoinSaveCsReq                                  = 3834
	TextJoinSaveScRsp                                  = 3848
	TextJoinQueryCsReq                                 = 3862
	TextJoinQueryScRsp                                 = 3888
	GetPrivateChatHistoryCsReq                         = 3902
	BatchMarkChatEmojiScRsp                            = 3906
	GetPrivateChatHistoryScRsp                         = 3909
	GetChatFriendHistoryCsReq                          = 3919
	GetChatEmojiListScRsp                              = 3929
	GetLoginChatInfoCsReq                              = 3933
	SendMsgCsReq                                       = 3934
	GetChatFriendHistoryScRsp                          = 3943
	MarkChatEmojiCsReq                                 = 3945
	SendMsgScRsp                                       = 3948
	GetLoginChatInfoScRsp                              = 3959
	RevcMsgScNotify                                    = 3962
	MarkChatEmojiScRsp                                 = 3968
	GetChatEmojiListCsReq                              = 3986
	PrivateMsgOfflineUsersScNotify                     = 3988
	BatchMarkChatEmojiCsReq                            = 3996
	AcceptedPamMissionExpireCsReq                      = 4034
	AcceptedPamMissionExpireScRsp                      = 4048
	SyncAcceptedPamMissionNotify                       = 4062
	GetGunPlayDataScRsp                                = 4101
	TakePictureCsReq                                   = 4102
	TriggerVoiceScRsp                                  = 4106
	TakePictureScRsp                                   = 4109
	GetMovieRacingDataCsReq                            = 4116
	UpdateGunPlayDataScRsp                             = 4128
	CancelCacheNotifyScRsp                             = 4129
	GetMovieRacingDataScRsp                            = 4130
	SubmitOrigamiItemCsReq                             = 4133
	ShareCsReq                                         = 4134
	UpdateGunPlayDataCsReq                             = 4141
	SecurityReportCsReq                                = 4145
	ShareScRsp                                         = 4148
	UpdateMovieRacingDataScRsp                         = 4156
	SubmitOrigamiItemScRsp                             = 4159
	GetShareDataCsReq                                  = 4162
	GetGunPlayDataCsReq                                = 4163
	SecurityReportScRsp                                = 4168
	UpdateMovieRacingDataCsReq                         = 4185
	CancelCacheNotifyCsReq                             = 4186
	GetShareDataScRsp                                  = 4188
	TriggerVoiceCsReq                                  = 4196
	StartBoxingClubBattleCsReq                         = 4202
	SetBoxingClubResonanceLineupScRsp                  = 4206
	StartBoxingClubBattleScRsp                         = 4209
	GiveUpBoxingClubChallengeCsReq                     = 4219
	BoxingClubChallengeUpdateScNotify                  = 4229
	ChooseBoxingClubStageOptionalBuffCsReq             = 4233
	GetBoxingClubInfoCsReq                             = 4234
	GiveUpBoxingClubChallengeScRsp                     = 4243
	ChooseBoxingClubResonanceCsReq                     = 4245
	GetBoxingClubInfoScRsp                             = 4248
	ChooseBoxingClubStageOptionalBuffScRsp             = 4259
	MatchBoxingClubOpponentCsReq                       = 4262
	ChooseBoxingClubResonanceScRsp                     = 4268
	BoxingClubRewardScNotify                           = 4286
	MatchBoxingClubOpponentScRsp                       = 4288
	SetBoxingClubResonanceLineupCsReq                  = 4296
	MuseumTargetMissionFinishNotify                    = 4301
	SetStuffToAreaCsReq                                = 4302
	UpgradeAreaScRsp                                   = 4306
	SetStuffToAreaScRsp                                = 4309
	MuseumRandomEventQueryScRsp                        = 4316
	RemoveStuffFromAreaCsReq                           = 4319
	MuseumTakeCollectRewardCsReq                       = 4328
	GetExhibitScNotify                                 = 4329
	MuseumRandomEventSelectCsReq                       = 4330
	UpgradeAreaStatCsReq                               = 4333
	GetMuseumInfoCsReq                                 = 4334
	MuseumRandomEventStartScNotify                     = 4337
	MuseumRandomEventQueryCsReq                        = 4339
	MuseumTargetRewardNotify                           = 4341
	MuseumFundsChangedScNotify                         = 4342
	RemoveStuffFromAreaScRsp                           = 4343
	FinishCurTurnCsReq                                 = 4345
	GetMuseumInfoScRsp                                 = 4348
	MuseumDispatchFinishedScNotify                     = 4356
	UpgradeAreaStatScRsp                               = 4359
	MuseumTakeCollectRewardScRsp                       = 4361
	BuyNpcStuffCsReq                                   = 4362
	MuseumTargetStartNotify                            = 4363
	FinishCurTurnScRsp                                 = 4368
	MuseumRandomEventSelectScRsp                       = 4385
	GetStuffScNotify                                   = 4386
	BuyNpcStuffScRsp                                   = 4388
	MuseumInfoChangedScNotify                          = 4395
	UpgradeAreaCsReq                                   = 4396
	EnterTreasureDungeonScRsp                          = 4406
	UseTreasureDungeonItemCsReq                        = 4416
	UseTreasureDungeonItemScRsp                        = 4430
	OpenTreasureDungeonGridCsReq                       = 4433
	TreasureDungeonDataScNotify                        = 4434
	InteractTreasureDungeonGridCsReq                   = 4437
	InteractTreasureDungeonGridScRsp                   = 4439
	FightTreasureDungeonMonsterScRsp                   = 4442
	GetTreasureDungeonActivityDataCsReq                = 4445
	TreasureDungeonFinishScNotify                      = 4448
	QuitTreasureDungeonScRsp                           = 4456
	OpenTreasureDungeonGridScRsp                       = 4459
	GetTreasureDungeonActivityDataScRsp                = 4468
	QuitTreasureDungeonCsReq                           = 4485
	FightTreasureDungeonMonsterCsReq                   = 4495
	EnterTreasureDungeonCsReq                          = 4496
	PlayerReturnTakePointRewardCsReq                   = 4502
	PlayerReturnTakePointRewardScRsp                   = 4509
	PlayerReturnTakeRewardCsReq                        = 4519
	PlayerReturnInfoQueryScRsp                         = 4529
	PlayerReturnStartScNotify                          = 4534
	PlayerReturnTakeRewardScRsp                        = 4543
	PlayerReturnForceFinishScNotify                    = 4545
	PlayerReturnSignCsReq                              = 4548
	PlayerReturnSignScRsp                              = 4562
	PlayerReturnInfoQueryCsReq                         = 4586
	PlayerReturnPointChangeScNotify                    = 4588
	GetPlayerReturnMultiDropInfoScRsp                  = 4602
	MultipleDropInfoNotify                             = 4609
	GetMultipleDropInfoCsReq                           = 4634
	GetMultipleDropInfoScRsp                           = 4648
	MultipleDropInfoScNotify                           = 4662
	GetPlayerReturnMultiDropInfoCsReq                  = 4688
	SaveLogisticsCsReq                                 = 4701
	AlleyPlacingGameScRsp                              = 4706
	AlleyTakeEventRewardScRsp                          = 4708
	AlleyTakeEventRewardCsReq                          = 4711
	PrestigeLevelUpCsReq                               = 4716
	GetSaveLogisticsMapCsReq                           = 4718
	StartAlleyEventCsReq                               = 4719
	AlleyGuaranteedFundsCsReq                          = 4724
	LogisticsScoreRewardSyncInfoScNotify               = 4725
	LogisticsInfoScNotify                              = 4728
	AlleyEventEffectNotify                             = 4729
	PrestigeLevelUpScRsp                               = 4730
	GetAlleyInfoCsReq                                  = 4734
	AlleyOrderChangedScNotify                          = 4737
	SaveLogisticsScRsp                                 = 4741
	RefreshAlleyOrderScRsp                             = 4742
	StartAlleyEventScRsp                               = 4743
	TakePrestigeRewardCsReq                            = 4745
	GetAlleyInfoScRsp                                  = 4748
	LogisticsDetonateStarSkiffCsReq                    = 4754
	AlleyShopLevelScNotify                             = 4756
	AlleyShipmentEventEffectsScNotify                  = 4761
	LogisticsGameCsReq                                 = 4762
	AlleyShipUnlockScNotify                            = 4763
	TakePrestigeRewardScRsp                            = 4768
	LogisticsDetonateStarSkiffScRsp                    = 4779
	AlleyGuaranteedFundsScRsp                          = 4782
	AlleyFundsScNotify                                 = 4785
	AlleyEventChangeNotify                             = 4786
	LogisticsGameScRsp                                 = 4788
	GetSaveLogisticsMapScRsp                           = 4791
	RefreshAlleyOrderCsReq                             = 4795
	AlleyPlacingGameCsReq                              = 4796
	AlleyShipUsedCountScNotify                         = 4797
	GetAetherDivideChallengeInfoCsReq                  = 4801
	StartAetherDivideSceneBattleCsReq                  = 4802
	SetAetherDivideLineUpScRsp                         = 4806
	AetherDivideTakeChallengeRewardCsReq               = 4808
	StartAetherDivideSceneBattleScRsp                  = 4809
	AetherDivideRefreshEndlessScNotify                 = 4811
	StartAetherDivideStageBattleCsReq                  = 4816
	AetherDivideSkillItemScNotify                      = 4818
	StartAetherDivideChallengeBattleCsReq              = 4819
	AetherDivideRefreshEndlessCsReq                    = 4824
	AetherDivideFinishChallengeScNotify                = 4828
	StartAetherDivideStageBattleScRsp                  = 4830
	EquipAetherDividePassiveSkillCsReq                 = 4833
	EnterAetherDivideSceneCsReq                        = 4834
	SwitchAetherDivideLineUpSlotCsReq                  = 4837
	SwitchAetherDivideLineUpSlotScRsp                  = 4839
	GetAetherDivideChallengeInfoScRsp                  = 4841
	ClearAetherDividePassiveSkillScRsp                 = 4842
	StartAetherDivideChallengeBattleScRsp              = 4843
	GetAetherDivideInfoCsReq                           = 4845
	EnterAetherDivideSceneScRsp                        = 4848
	AetherDivideTakeChallengeRewardScRsp               = 4854
	AetherDivideSpiritExpUpScRsp                       = 4856
	EquipAetherDividePassiveSkillScRsp                 = 4859
	AetherDivideTainerInfoScNotify                     = 4861
	LeaveAetherDivideSceneCsReq                        = 4862
	AetherDivideSpiritInfoScNotify                     = 4863
	GetAetherDivideInfoScRsp                           = 4868
	AetherDivideRefreshEndlessScRsp                    = 4882
	AetherDivideSpiritExpUpCsReq                       = 4885
	LeaveAetherDivideSceneScRsp                        = 4888
	ClearAetherDividePassiveSkillCsReq                 = 4895
	SetAetherDivideLineUpCsReq                         = 4896
	AetherDivideLineupScNotify                         = 4897
	EnterFantasticStoryActivityStageScRsp              = 4902
	FantasticStoryActivityBattleEndScNotify            = 4909
	GetFantasticStoryActivityDataCsReq                 = 4934
	GetFantasticStoryActivityDataScRsp                 = 4948
	FinishChapterScNotify                              = 4962
	EnterFantasticStoryActivityStageCsReq              = 4988
	UnlockChatBubbleScNotify                           = 5102
	SelectPhoneThemeCsReq                              = 5109
	SelectPhoneThemeScRsp                              = 5119
	GetPhoneDataCsReq                                  = 5134
	UnlockPhoneThemeScNotify                           = 5143
	GetPhoneDataScRsp                                  = 5148
	SelectChatBubbleCsReq                              = 5162
	SelectChatBubbleScRsp                              = 5188
	RogueModifierSelectCellScRsp                       = 5302
	RogueModifierStageStartNotify                      = 5329
	RogueModifierUpdateNotify                          = 5343
	RogueModifierAddNotify                             = 5362
	RogueModifierDelNotify                             = 5386
	RogueModifierSelectCellCsReq                       = 5388
	FinishChessRogueSubStoryCsReq                      = 5405
	FinishChessRogueNousSubStoryCsReq                  = 5411
	ChessRogueQuitScRsp                                = 5412
	ChessRogueNousDiceSurfaceUnlockNotify              = 5413
	GetChessRogueNousStoryInfoCsReq                    = 5415
	ChessRogueSelectBpScRsp                            = 5416
	ChessRogueUpdateReviveInfoScNotify                 = 5419
	SyncChessRogueNousSubStoryScNotify                 = 5420
	ChessRogueFinishCurRoomNotify                      = 5422
	ChessRogueConfirmRollCsReq                         = 5424
	ChessRogueNousEnableRogueTalentScRsp               = 5425
	GetChessRogueBuffEnhanceInfoScRsp                  = 5426
	ChessRogueNousGetRogueTalentInfoScRsp              = 5429
	ChessRogueGoAheadScRsp                             = 5431
	ChessRogueUpdateLevelBaseInfoScNotify              = 5432
	ChessRogueSelectCellCsReq                          = 5434
	ChessRogueEnterNextLayerScRsp                      = 5436
	FinishChessRogueSubStoryScRsp                      = 5437
	ChessRogueNousGetRogueTalentInfoCsReq              = 5448
	ChessRoguePickAvatarScRsp                          = 5449
	ChessRogueSelectCellScRsp                          = 5450
	ChessRogueNousDiceUpdateNotify                     = 5452
	SelectChessRogueNousSubStoryScRsp                  = 5454
	ChessRogueEnterCsReq                               = 5456
	ChessRogueGoAheadCsReq                             = 5458
	ChessRogueQueryCsReq                               = 5459
	ChessRogueGiveUpCsReq                              = 5463
	ChessRogueSkipTeachingLevelCsReq                   = 5465
	ChessRogueQueryAeonDimensionsScRsp                 = 5466
	EnhanceChessRogueBuffScRsp                         = 5468
	ChessRogueUpdateActionPointScNotify                = 5469
	ChessRogueReviveAvatarScRsp                        = 5470
	ChessRogueStartScRsp                               = 5471
	ChessRogueLeaveCsReq                               = 5473
	ChessRogueSkipTeachingLevelScRsp                   = 5474
	GetChessRogueStoryAeonTalkInfoCsReq                = 5477
	ChessRogueUpdateMoneyInfoScNotify                  = 5480
	ChessRogueNousEditDiceScRsp                        = 5482
	SelectChessRogueNousSubStoryCsReq                  = 5484
	SyncChessRogueNousMainStoryScNotify                = 5487
	ChessRogueReRollDiceCsReq                          = 5490
	ChessRogueQueryBpCsReq                             = 5495
	EnterChessRogueAeonRoomScRsp                       = 5496
	ChessRogueUpdateAeonModifierValueScNotify          = 5498
	ChessRogueReRollDiceScRsp                          = 5500
	FinishChessRogueNousSubStoryScRsp                  = 5501
	ChessRogueUpdateBoardScNotify                      = 5502
	ChessRogueLayerAccountInfoNotify                   = 5507
	ChessRogueCellUpdateNotify                         = 5508
	ChessRogueGiveUpScRsp                              = 5511
	ChessRoguePickAvatarCsReq                          = 5517
	ChessRogueEnterCellCsReq                           = 5518
	GetChessRogueBuffEnhanceInfoCsReq                  = 5522
	ChessRogueConfirmRollScRsp                         = 5523
	ChessRogueUpdateDiceInfoScNotify                   = 5526
	GetChessRogueStoryInfoScRsp                        = 5527
	ChessRogueQueryAeonDimensionsCsReq                 = 5529
	ChessRogueLeaveScRsp                               = 5531
	GetChessRogueStoryInfoCsReq                        = 5532
	ChessRogueRollDiceCsReq                            = 5535
	SelectChessRogueSubStoryScRsp                      = 5536
	SyncChessRogueNousValueScNotify                    = 5537
	ChessRogueReviveAvatarCsReq                        = 5539
	ChessRogueEnterCellScRsp                           = 5540
	ChessRogueUpdateDicePassiveAccumulateValueScNotify = 5542
	ChessRogueEnterNextLayerCsReq                      = 5543
	ChessRogueCheatRollCsReq                           = 5544
	ChessRogueRollDiceScRsp                            = 5546
	ChessRogueSelectBpCsReq                            = 5549
	ChessRogueNousEditDiceCsReq                        = 5550
	ChessRogueChangeyAeonDimensionNotify               = 5557
	ChessRogueGiveUpRollCsReq                          = 5558
	ChessRogueEnterScRsp                               = 5559
	GetChessRogueNousStoryInfoScRsp                    = 5561
	ChessRogueQuestFinishNotify                        = 5565
	ChessRogueNousEnableRogueTalentCsReq               = 5570
	SyncChessRogueMainStoryFinishScNotify              = 5573
	ChessRogueQuitCsReq                                = 5575
	ChessRogueGiveUpRollScRsp                          = 5576
	ChessRogueUpdateAllowedSelectCellScNotify          = 5577
	GetChessRogueStoryAeonTalkInfoScRsp                = 5580
	ChessRogueUpdateUnlockLevelScNotify                = 5582
	ChessRogueMoveCellNotify                           = 5586
	EnterChessRogueAeonRoomCsReq                       = 5589
	EnhanceChessRogueBuffCsReq                         = 5592
	ChessRogueStartCsReq                               = 5596
	ChessRogueQueryScRsp                               = 5597
	ChessRogueQueryBpScRsp                             = 5598
	ChessRogueCheatRollScRsp                           = 5599
	SelectChessRogueSubStoryCsReq                      = 5600
	StopRogueAdventureRoomScRsp                        = 5601
	GetRogueShopMiracleInfoScRsp                       = 5602
	HandleRogueCommonPendingActionCsReq                = 5604
	GetRogueAdventureRoomInfoCsReq                     = 5606
	GetRogueShopBuffInfoCsReq                          = 5609
	GetEnhanceCommonRogueBuffInfoCsReq                 = 5616
	GetRogueShopBuffInfoScRsp                          = 5619
	SyncRogueHandbookDataUpdateScNotify                = 5625
	BuyRogueShopBuffCsReq                              = 5629
	GetEnhanceCommonRogueBuffInfoScRsp                 = 5630
	GetRogueAdventureRoomInfoScRsp                     = 5633
	SyncRogueAdventureRoomInfoScNotify                 = 5634
	ExchangeRogueBuffWithMiracleCsReq                  = 5637
	ExchangeRogueBuffWithMiracleScRsp                  = 5639
	BuyRogueShopMiracleCsReq                           = 5643
	BuyRogueShopBuffScRsp                              = 5645
	PrepareRogueAdventureRoomCsReq                     = 5648
	GetRogueHandbookDataCsReq                          = 5654
	UpdateRogueAdventureRoomScoreCsReq                 = 5655
	EnhanceCommonRogueBuffScRsp                        = 5656
	PrepareRogueAdventureRoomScRsp                     = 5662
	StopRogueAdventureRoomCsReq                        = 5663
	TakeRogueMiracleHandbookRewardScRsp                = 5665
	UpdateRogueAdventureRoomScoreScRsp                 = 5666
	SyncRogueCommonActionResultScNotify                = 5667
	RogueNpcDisappearCsReq                             = 5668
	CommonRogueUpdateScNotify                          = 5671
	SyncRogueCommonVirtualItemInfoScNotify             = 5673
	HandleRogueCommonPendingActionScRsp                = 5675
	GetRogueHandbookDataScRsp                          = 5679
	EnhanceCommonRogueBuffCsReq                        = 5685
	BuyRogueShopMiracleScRsp                           = 5686
	GetRogueShopMiracleInfoCsReq                       = 5688
	TakeRogueEventHandbookRewardCsReq                  = 5689
	TakeRogueEventHandbookRewardScRsp                  = 5690
	SyncRogueCommonPendingActionScNotify               = 5692
	CommonRogueQueryCsReq                              = 5693
	RogueNpcDisappearScRsp                             = 5696
	CommonRogueQueryScRsp                              = 5698
	TakeRogueMiracleHandbookRewardCsReq                = 5700
	StartBattleCollegeScRsp                            = 5702
	GetBattleCollegeDataCsReq                          = 5734
	GetBattleCollegeDataScRsp                          = 5748
	BattleCollegeDataChangeScNotify                    = 5762
	StartBattleCollegeCsReq                            = 5788
	HeliobusSnsPostCsReq                               = 5802
	HeliobusUpgradeLevelScRsp                          = 5806
	HeliobusSnsPostScRsp                               = 5809
	HeliobusEnterBattleScRsp                           = 5816
	HeliobusSnsLikeCsReq                               = 5819
	HeliobusSnsCommentScRsp                            = 5829
	HeliobusStartRaidCsReq                             = 5830
	HeliobusUnlockSkillScNotify                        = 5833
	HeliobusActivityDataCsReq                          = 5834
	HeliobusEnterBattleCsReq                           = 5839
	HeliobusSnsLikeScRsp                               = 5843
	HeliobusSnsUpdateScNotify                          = 5845
	HeliobusActivityDataScRsp                          = 5848
	HeliobusChallengeUpdateScNotify                    = 5856
	HeliobusSelectSkillCsReq                           = 5859
	HeliobusSnsReadCsReq                               = 5862
	HeliobusLineupUpdateScNotify                       = 5863
	HeliobusInfoChangedScNotify                        = 5868
	HeliobusStartRaidScRsp                             = 5885
	HeliobusSnsCommentCsReq                            = 5886
	HeliobusSnsReadScRsp                               = 5888
	HeliobusSelectSkillScRsp                           = 5895
	HeliobusUpgradeLevelCsReq                          = 5896
	GetSingleRedDotParamGroupCsReq                     = 5902
	GetSingleRedDotParamGroupScRsp                     = 5909
	GetAllRedDotDataCsReq                              = 5934
	GetAllRedDotDataScRsp                              = 5948
	UpdateRedDotDataCsReq                              = 5962
	UpdateRedDotDataScRsp                              = 5988
	RogueEndlessActivityBattleEndScNotify              = 6002
	TakeRogueEndlessActivityPointRewardCsReq           = 6009
	TakeRogueEndlessActivityPointRewardScRsp           = 6019
	GetRogueEndlessActivityDataCsReq                   = 6034
	TakeRogueEndlessActivityAllBonusRewardCsReq        = 6043
	GetRogueEndlessActivityDataScRsp                   = 6048
	EnterRogueEndlessActivityStageCsReq                = 6062
	TakeRogueEndlessActivityAllBonusRewardScRsp        = 6086
	EnterRogueEndlessActivityStageScRsp                = 6088
	UpdateServerPrefsDataCsReq                         = 6102
	UpdateServerPrefsDataScRsp                         = 6109
	GetAllServerPrefsDataCsReq                         = 6134
	GetAllServerPrefsDataScRsp                         = 6148
	GetServerPrefsDataCsReq                            = 6162
	GetServerPrefsDataScRsp                            = 6188
	ChangeStoryLineScRsp                               = 6202
	ChangeStoryLineFinishScNotify                      = 6209
	StoryLineTrialAvatarChangeScNotify                 = 6219
	GetStoryLineInfoCsReq                              = 6234
	GetStoryLineInfoScRsp                              = 6248
	StoryLineInfoScNotify                              = 6262
	ChangeStoryLineCsReq                               = 6288
	SubmitEmotionItemCsReq                             = 6302
	SubmitEmotionItemScRsp                             = 6309
	FinishEmotionDialoguePerformanceCsReq              = 6319
	HeartDialTraceScriptCsReq                          = 6329
	GetHeartDialInfoCsReq                              = 6334
	FinishEmotionDialoguePerformanceScRsp              = 6343
	HeartDialTraceScriptScRsp                          = 6345
	GetHeartDialInfoScRsp                              = 6348
	ChangeScriptEmotionCsReq                           = 6362
	HeartDialScriptChangeScNotify                      = 6386
	ChangeScriptEmotionScRsp                           = 6388
	TravelBrochureSelectMessageCsReq                   = 6402
	TravelBrochureSelectMessageScRsp                   = 6409
	TravelBrochureApplyPasterListCsReq                 = 6416
	TravelBrochureApplyPasterCsReq                     = 6419
	TravelBrochureRemovePasterScRsp                    = 6429
	TravelBrochureApplyPasterListScRsp                 = 6430
	TravelBrochureSetCustomValueCsReq                  = 6433
	TravelBrochureGetDataCsReq                         = 6434
	TravelBrochurePageResetCsReq                       = 6437
	TravelBrochurePageResetScRsp                       = 6439
	TravelBrochureSetPageDescStatusScRsp               = 6442
	TravelBrochureApplyPasterScRsp                     = 6443
	TravelBrochureUpdatePasterPosCsReq                 = 6445
	TravelBrochureGetDataScRsp                         = 6448
	TravelBrochureSetCustomValueScRsp                  = 6459
	TravelBrochurePageUnlockScNotify                   = 6462
	TravelBrochureUpdatePasterPosScRsp                 = 6468
	TravelBrochureRemovePasterCsReq                    = 6486
	TravelBrochureSetPageDescStatusCsReq               = 6495
	TravelBrochureGetPasterScNotify                    = 6496
	RestoreWolfBroGameArchiveCsReq                     = 6502
	WolfBroGamePickupBulletCsReq                       = 6506
	RestoreWolfBroGameArchiveScRsp                     = 6509
	QuitWolfBroGameCsReq                               = 6519
	GetWolfBroGameDataScRsp                            = 6529
	WolfBroGamePickupBulletScRsp                       = 6533
	StartWolfBroGameCsReq                              = 6534
	WolfBroGameExplodeMonsterScRsp                     = 6537
	WolfBroGameExplodeMonsterCsReq                     = 6542
	QuitWolfBroGameScRsp                               = 6543
	WolfBroGameDataChangeScNotify                      = 6545
	StartWolfBroGameScRsp                              = 6548
	WolfBroGameActivateBulletCsReq                     = 6559
	ArchiveWolfBroGameCsReq                            = 6562
	WolfBroGameUseBulletCsReq                          = 6568
	GetWolfBroGameDataCsReq                            = 6586
	ArchiveWolfBroGameScRsp                            = 6588
	WolfBroGameActivateBulletScRsp                     = 6595
	WolfBroGameUseBulletScRsp                          = 6596
	StrongChallengeActivityBattleEndScNotify           = 6602
	GetStrongChallengeActivityDataCsReq                = 6634
	GetStrongChallengeActivityDataScRsp                = 6648
	EnterStrongChallengeActivityStageCsReq             = 6662
	EnterStrongChallengeActivityStageScRsp             = 6688
	SpaceZooMutateCsReq                                = 6702
	SpaceZooTakeCsReq                                  = 6706
	SpaceZooMutateScRsp                                = 6709
	SpaceZooOpCatteryCsReq                             = 6719
	SpaceZooDeleteCatScRsp                             = 6729
	SpaceZooTakeScRsp                                  = 6733
	SpaceZooDataCsReq                                  = 6734
	SpaceZooOpCatteryScRsp                             = 6743
	SpaceZooCatUpdateNotify                            = 6745
	SpaceZooDataScRsp                                  = 6748
	SpaceZooBornCsReq                                  = 6762
	SpaceZooExchangeItemCsReq                          = 6768
	SpaceZooDeleteCatCsReq                             = 6786
	SpaceZooBornScRsp                                  = 6788
	SpaceZooExchangeItemScRsp                          = 6796
	DeployRotaterCsReq                                 = 6802
	ResetMapRotationRegionScRsp                        = 6806
	DeployRotaterScRsp                                 = 6809
	RotateMapCsReq                                     = 6819
	LeaveMapRotationRegionScRsp                        = 6829
	LeaveMapRotationRegionScNotify                     = 6833
	EnterMapRotationRegionCsReq                        = 6834
	RemoveRotaterScRsp                                 = 6837
	UpdateRotaterScNotify                              = 6839
	RemoveRotaterCsReq                                 = 6842
	RotateMapScRsp                                     = 6843
	GetMapRotationDataCsReq                            = 6845
	EnterMapRotationRegionScRsp                        = 6848
	UpdateEnergyScNotify                               = 6859
	InteractChargerCsReq                               = 6862
	GetMapRotationDataScRsp                            = 6868
	LeaveMapRotationRegionCsReq                        = 6886
	InteractChargerScRsp                               = 6888
	UpdateMapRotationDataScNotify                      = 6895
	ResetMapRotationRegionCsReq                        = 6896
	GetRollShopInfoCsReq                               = 6901
	DoGachaInRollShopCsReq                             = 6913
	DoGachaInRollShopScRsp                             = 6917
	GetRollShopInfoScRsp                               = 6918
	TakeRollShopRewardScRsp                            = 6919
	TakeRollShopRewardCsReq                            = 6920
	GetOfferingInfoCsReq                               = 6921
	SubmitOfferingItemCsReq                            = 6933
	SubmitOfferingItemScRsp                            = 6937
	GetOfferingInfoScRsp                               = 6938
	TakeOfferingRewardScRsp                            = 6939
	TakeOfferingRewardCsReq                            = 6940
	RaidCollectionDataCsReq                            = 6941
	RaidCollectionDataScNotify                         = 6953
	RaidCollectionDataScRsp                            = 6958
	GetTelevisionActivityDataCsReq                     = 6961
	TelevisionActivityDataChangeScNotify               = 6973
	EnterTelevisionActivityStageCsReq                  = 6977
	GetTelevisionActivityDataScRsp                     = 6978
	TelevisionActivityBattleEndScNotify                = 6979
	EnterTelevisionActivityStageScRsp                  = 6980
	GetDrinkMakerDataCsReq                             = 6981
	MakeMissionDrinkCsReq                              = 6982
	DrinkMakerUpdateTipsNotify                         = 6983
	DrinkMakerDayEndScNotify                           = 6984
	DrinkMakerChallengeCsReq                           = 6985
	DrinkMakerChallengeScRsp                           = 6990
	MakeDrinkCsReq                                     = 6993
	MakeMissionDrinkScRsp                              = 6996
	MakeDrinkScRsp                                     = 6997
	GetDrinkMakerDataScRsp                             = 6998
	EndDrinkMakerSequenceScRsp                         = 6999
	EndDrinkMakerSequenceCsReq                         = 7000
	MonopolyGiveUpCurContentScRsp                      = 7001
	MonopolyEventSelectFriendCsReq                     = 7003
	GetMonopolyFriendRankingListScRsp                  = 7004
	GetSocialEventServerCacheCsReq                     = 7005
	DailyFirstEnterMonopolyActivityScRsp               = 7006
	MonopolyGetRafflePoolInfoCsReq                     = 7007
	MonopolyGameBingoFlipCardScRsp                     = 7008
	MonopolyRollDiceCsReq                              = 7009
	MonopolyGetRaffleTicketScRsp                       = 7010
	MonopolyGameBingoFlipCardCsReq                     = 7011
	DeleteSocialEventServerCacheCsReq                  = 7012
	MonopolyGetRaffleTicketCsReq                       = 7013
	MonopolyScrachRaffleTicketScRsp                    = 7014
	DeleteSocialEventServerCacheScRsp                  = 7015
	MonopolyBuyGoodsCsReq                              = 7016
	MonopolyGameRaiseRatioCsReq                        = 7018
	MonopolyRollDiceScRsp                              = 7019
	MonopolyGetRafflePoolInfoScRsp                     = 7020
	GetMonopolyDailyReportScRsp                        = 7021
	GetSocialEventServerCacheScRsp                     = 7022
	MonopolyGameGachaCsReq                             = 7024
	MonopolyGameCreateScNotify                         = 7025
	MonopolyClickCellScRsp                             = 7027
	MonopolyCheatDiceScRsp                             = 7028
	MonopolySelectOptionCsReq                          = 7029
	MonopolyBuyGoodsScRsp                              = 7030
	MonopolyGetDailyInitItemCsReq                      = 7031
	MonopolyConditionUpdateScNotify                    = 7032
	MonopolyRollRandomCsReq                            = 7033
	GetMonopolyInfoCsReq                               = 7034
	MonopolyDailySettleScNotify                        = 7035
	MonopolyConfirmRandomCsReq                         = 7037
	MonopolyClickMbtiReportCsReq                       = 7038
	MonopolyConfirmRandomScRsp                         = 7039
	MonopolyGetRegionProgressCsReq                     = 7040
	MonopolyCheatDiceCsReq                             = 7041
	MonopolyReRollRandomScRsp                          = 7042
	MonopolyMoveCsReq                                  = 7043
	GetMonopolyFriendRankingListCsReq                  = 7044
	MonopolySelectOptionScRsp                          = 7045
	MonopolySocialEventEffectScNotify                  = 7046
	MonopolyClickCellCsReq                             = 7047
	GetMonopolyInfoScRsp                               = 7048
	GetMbtiReportScRsp                                 = 7049
	MonopolyGetDailyInitItemScRsp                      = 7050
	GetMonopolyMbtiReportRewardCsReq                   = 7052
	MonopolyAcceptQuizCsReq                            = 7054
	MonopolyUpgradeAssetScRsp                          = 7056
	MonopolyTakePhaseRewardCsReq                       = 7058
	MonopolyRollRandomScRsp                            = 7059
	MonopolyContentUpdateScNotify                      = 7061
	MonopolyActionResultScNotify                       = 7062
	MonopolyGiveUpCurContentCsReq                      = 7063
	MonopolyTakePhaseRewardScRsp                       = 7064
	MonopolyGuessChooseScRsp                           = 7065
	MonopolyGuessDrawScNotify                          = 7067
	MonopolyGetRegionProgressScRsp                     = 7069
	MonopolyTakeRaffleTicketRewardScRsp                = 7070
	GetMbtiReportCsReq                                 = 7071
	MonopolyClickMbtiReportScRsp                       = 7074
	MonopolyLikeCsReq                                  = 7075
	GetMonopolyDailyReportCsReq                        = 7076
	MonopolySttUpdateScNotify                          = 7077
	MonopolyEventLoadUpdateScNotify                    = 7078
	MonopolyAcceptQuizScRsp                            = 7079
	GetMonopolyMbtiReportRewardScRsp                   = 7081
	MonopolyGameGachaScRsp                             = 7082
	MonopolyTakeRaffleTicketRewardCsReq                = 7084
	MonopolyUpgradeAssetCsReq                          = 7085
	MonopolyMoveScRsp                                  = 7086
	MonopolyCellUpdateNotify                           = 7088
	MonopolyGuessBuyInformationCsReq                   = 7089
	MonopolyGuessBuyInformationScRsp                   = 7090
	MonopolyGameRaiseRatioScRsp                        = 7091
	MonopolyQuizDurationChangeScNotify                 = 7092
	MonopolyLikeScRsp                                  = 7093
	MonopolyEventSelectFriendScRsp                     = 7094
	MonopolyReRollRandomCsReq                          = 7095
	DailyFirstEnterMonopolyActivityCsReq               = 7096
	MonopolyGameSettleScNotify                         = 7097
	MonopolyLikeScNotify                               = 7098
	MonopolyScrachRaffleTicketCsReq                    = 7099
	MonopolyGuessChooseCsReq                           = 7100
	EvolveBuildLeaveScRsp                              = 7101
	EvolveBuildShopAbilityDownCsReq                    = 7103
	EvolveBuildCoinNotify                              = 7104
	EvolveBuildReRandomStageScRsp                      = 7106
	EvolveBuildFinishScNotify                          = 7107
	EvolveBuildQueryInfoCsReq                          = 7108
	EvolveBuildUnlockInfoNotify                        = 7110
	EvolveBuildLeaveCsReq                              = 7117
	EvolveBuildShopAbilityResetScRsp                   = 7120
	EvolveBuildTakeExpRewardScRsp                      = 7121
	EvolveBuildTakeExpRewardCsReq                      = 7122
	EvolveBuildReRandomStageCsReq                      = 7131
	EvolveBuildShopAbilityUpCsReq                      = 7133
	EvolveBuildGiveupCsReq                             = 7134
	EvolveBuildShopAbilityUpScRsp                      = 7135
	EvolveBuildStartStageScRsp                         = 7136
	EvolveBuildStartStageCsReq                         = 7141
	EvolveBuildShopAbilityResetCsReq                   = 7144
	EvolveBuildShopAbilityDownScRsp                    = 7145
	EvolveBuildStartLevelScRsp                         = 7147
	EvolveBuildStartLevelCsReq                         = 7148
	EvolveBuildQueryInfoScRsp                          = 7149
	EvolveBuildGiveupScRsp                             = 7150
	EnterFeverTimeActivityStageCsReq                   = 7151
	FeverTimeActivityBattleEndScNotify                 = 7153
	GetFeverTimeActivityDataScRsp                      = 7154
	GetFeverTimeActivityDataCsReq                      = 7156
	EnterFeverTimeActivityStageScRsp                   = 7159
	ClockParkUnlockTalentCsReq                         = 7202
	ClockParkQuitScriptScRsp                           = 7206
	ClockParkUnlockTalentScRsp                         = 7209
	ClockParkFinishScriptScNotify                      = 7216
	ClockParkStartScriptCsReq                          = 7219
	ClockParkGetOngoingScriptInfoScRsp                 = 7229
	ClockParkSyncVirtualItemScNotify                   = 7230
	ClockParkGetInfoCsReq                              = 7234
	ClockParkUseBuffCsReq                              = 7237
	ClockParkUseBuffScRsp                              = 7239
	ClockParkStartScriptScRsp                          = 7243
	ClockParkHandleWaitOperationCsReq                  = 7245
	ClockParkGetInfoScRsp                              = 7248
	ClockParkUnlockScriptCsReq                         = 7262
	ClockParkHandleWaitOperationScRsp                  = 7268
	ClockParkGetOngoingScriptInfoCsReq                 = 7286
	ClockParkUnlockScriptScRsp                         = 7288
	ClockParkBattleEndScNotify                         = 7295
	ClockParkQuitScriptCsReq                           = 7296
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
	c.regMsg(EnhanceCommonRogueBuffCsReq, func() any { return new(proto.EnhanceCommonRogueBuffCsReq) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(ArchiveWolfBroGameScRsp, func() any { return new(proto.ArchiveWolfBroGameScRsp) })
	c.regMsg(SpringRecoverSingleAvatarScRsp, func() any { return new(proto.SpringRecoverSingleAvatarScRsp) })
	c.regMsg(ChessRogueCheatRollCsReq, func() any { return new(proto.ChessRogueCheatRollCsReq) })
	c.regMsg(AddBlacklistScRsp, func() any { return new(proto.AddBlacklistScRsp) })
	c.regMsg(SetClientPausedCsReq, func() any { return new(proto.SetClientPausedCsReq) })
	c.regMsg(TeleportToMissionResetPointCsReq, func() any { return new(proto.TeleportToMissionResetPointCsReq) })
	c.regMsg(TravelBrochureApplyPasterScRsp, func() any { return new(proto.TravelBrochureApplyPasterScRsp) })
	c.regMsg(StopRogueAdventureRoomScRsp, func() any { return new(proto.StopRogueAdventureRoomScRsp) })
	c.regMsg(StartFinishMainMissionScNotify, func() any { return new(proto.StartFinishMainMissionScNotify) })
	c.regMsg(EnterTelevisionActivityStageCsReq, func() any { return new(proto.EnterTelevisionActivityStageCsReq) })
	c.regMsg(BatchGetQuestDataScRsp, func() any { return new(proto.BatchGetQuestDataScRsp) })
	c.regMsg(SetGameplayBirthdayCsReq, func() any { return new(proto.SetGameplayBirthdayCsReq) })
	c.regMsg(GetMainMissionCustomValueCsReq, func() any { return new(proto.GetMainMissionCustomValueCsReq) })
	c.regMsg(SubmitEmotionItemCsReq, func() any { return new(proto.SubmitEmotionItemCsReq) })
	c.regMsg(SyncRogueHandbookDataUpdateScNotify, func() any { return new(proto.SyncRogueHandbookDataUpdateScNotify) })
	c.regMsg(TakeRogueAeonLevelRewardCsReq, func() any { return new(proto.TakeRogueAeonLevelRewardCsReq) })
	c.regMsg(DailyActiveInfoNotify, func() any { return new(proto.DailyActiveInfoNotify) })
	c.regMsg(MonopolyGetRaffleTicketCsReq, func() any { return new(proto.MonopolyGetRaffleTicketCsReq) })
	c.regMsg(StartRaidCsReq, func() any { return new(proto.StartRaidCsReq) })
	c.regMsg(TakeBpRewardScRsp, func() any { return new(proto.TakeBpRewardScRsp) })
	c.regMsg(MonopolyConfirmRandomScRsp, func() any { return new(proto.MonopolyConfirmRandomScRsp) })
	c.regMsg(EquipAetherDividePassiveSkillCsReq, func() any { return new(proto.EquipAetherDividePassiveSkillCsReq) })
	c.regMsg(CancelExpeditionCsReq, func() any { return new(proto.CancelExpeditionCsReq) })
	c.regMsg(CancelActivityExpeditionCsReq, func() any { return new(proto.CancelActivityExpeditionCsReq) })
	c.regMsg(TakePictureScRsp, func() any { return new(proto.TakePictureScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillScRsp, func() any { return new(proto.ClearAetherDividePassiveSkillScRsp) })
	c.regMsg(SecurityReportScRsp, func() any { return new(proto.SecurityReportScRsp) })
	c.regMsg(PrepareRogueAdventureRoomCsReq, func() any { return new(proto.PrepareRogueAdventureRoomCsReq) })
	c.regMsg(SyncRogueAreaUnlockScNotify, func() any { return new(proto.SyncRogueAreaUnlockScNotify) })
	c.regMsg(ChessRogueEnterScRsp, func() any { return new(proto.ChessRogueEnterScRsp) })
	c.regMsg(HeliobusUpgradeLevelCsReq, func() any { return new(proto.HeliobusUpgradeLevelCsReq) })
	c.regMsg(LeaveAetherDivideSceneCsReq, func() any { return new(proto.LeaveAetherDivideSceneCsReq) })
	c.regMsg(ExchangeHcoinScRsp, func() any { return new(proto.ExchangeHcoinScRsp) })
	c.regMsg(HeliobusSnsCommentCsReq, func() any { return new(proto.HeliobusSnsCommentCsReq) })
	c.regMsg(GetNpcTakenRewardScRsp, func() any { return new(proto.GetNpcTakenRewardScRsp) })
	c.regMsg(HeliobusEnterBattleCsReq, func() any { return new(proto.HeliobusEnterBattleCsReq) })
	c.regMsg(GetHeroBasicTypeInfoCsReq, func() any { return new(proto.GetHeroBasicTypeInfoCsReq) })
	c.regMsg(FinishChapterScNotify, func() any { return new(proto.FinishChapterScNotify) })
	c.regMsg(GetMovieRacingDataScRsp, func() any { return new(proto.GetMovieRacingDataScRsp) })
	c.regMsg(FinishQuestScRsp, func() any { return new(proto.FinishQuestScRsp) })
	c.regMsg(GetOfferingInfoScRsp, func() any { return new(proto.GetOfferingInfoScRsp) })
	c.regMsg(GetMonopolyMbtiReportRewardScRsp, func() any { return new(proto.GetMonopolyMbtiReportRewardScRsp) })
	c.regMsg(StartChallengeScRsp, func() any { return new(proto.StartChallengeScRsp) })
	c.regMsg(EnteredSceneChangeScNotify, func() any { return new(proto.EnteredSceneChangeScNotify) })
	c.regMsg(SyncRogueReviveInfoScNotify, func() any { return new(proto.SyncRogueReviveInfoScNotify) })
	c.regMsg(TravelBrochureApplyPasterListCsReq, func() any { return new(proto.TravelBrochureApplyPasterListCsReq) })
	c.regMsg(MatchBoxingClubOpponentCsReq, func() any { return new(proto.MatchBoxingClubOpponentCsReq) })
	c.regMsg(QueryProductInfoCsReq, func() any { return new(proto.QueryProductInfoCsReq) })
	c.regMsg(ComposeLimitNumUpdateNotify, func() any { return new(proto.ComposeLimitNumUpdateNotify) })
	c.regMsg(SetIsDisplayAvatarInfoScRsp, func() any { return new(proto.SetIsDisplayAvatarInfoScRsp) })
	c.regMsg(AetherDivideSpiritInfoScNotify, func() any { return new(proto.AetherDivideSpiritInfoScNotify) })
	c.regMsg(GetRogueEndlessActivityDataScRsp, func() any { return new(proto.GetRogueEndlessActivityDataScRsp) })
	c.regMsg(ChessRogueGoAheadScRsp, func() any { return new(proto.ChessRogueGoAheadScRsp) })
	c.regMsg(UpgradeAreaCsReq, func() any { return new(proto.UpgradeAreaCsReq) })
	c.regMsg(SellItemCsReq, func() any { return new(proto.SellItemCsReq) })
	c.regMsg(GetRogueHandbookDataScRsp, func() any { return new(proto.GetRogueHandbookDataScRsp) })
	c.regMsg(ChessRogueQueryCsReq, func() any { return new(proto.ChessRogueQueryCsReq) })
	c.regMsg(LeaveRaidCsReq, func() any { return new(proto.LeaveRaidCsReq) })
	c.regMsg(ExpUpEquipmentCsReq, func() any { return new(proto.ExpUpEquipmentCsReq) })
	c.regMsg(TrialBackGroundMusicCsReq, func() any { return new(proto.TrialBackGroundMusicCsReq) })
	c.regMsg(GetHeartDialInfoCsReq, func() any { return new(proto.GetHeartDialInfoCsReq) })
	c.regMsg(GetBagScRsp, func() any { return new(proto.GetBagScRsp) })
	c.regMsg(SubmitOfferingItemScRsp, func() any { return new(proto.SubmitOfferingItemScRsp) })
	c.regMsg(GetChessRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetChessRogueBuffEnhanceInfoCsReq) })
	c.regMsg(EquipAetherDividePassiveSkillScRsp, func() any { return new(proto.EquipAetherDividePassiveSkillScRsp) })
	c.regMsg(EnterTelevisionActivityStageScRsp, func() any { return new(proto.EnterTelevisionActivityStageScRsp) })
	c.regMsg(AcceptMainMissionScRsp, func() any { return new(proto.AcceptMainMissionScRsp) })
	c.regMsg(GetChallengeCsReq, func() any { return new(proto.GetChallengeCsReq) })
	c.regMsg(UpdateRogueAdventureRoomScoreCsReq, func() any { return new(proto.UpdateRogueAdventureRoomScoreCsReq) })
	c.regMsg(PromoteAvatarScRsp, func() any { return new(proto.PromoteAvatarScRsp) })
	c.regMsg(SceneUpdatePositionVersionNotify, func() any { return new(proto.SceneUpdatePositionVersionNotify) })
	c.regMsg(RemoveStuffFromAreaCsReq, func() any { return new(proto.RemoveStuffFromAreaCsReq) })
	c.regMsg(FightTreasureDungeonMonsterCsReq, func() any { return new(proto.FightTreasureDungeonMonsterCsReq) })
	c.regMsg(SyncApplyFriendScNotify, func() any { return new(proto.SyncApplyFriendScNotify) })
	c.regMsg(GetPlayerBoardDataCsReq, func() any { return new(proto.GetPlayerBoardDataCsReq) })
	c.regMsg(MarkChatEmojiScRsp, func() any { return new(proto.MarkChatEmojiScRsp) })
	c.regMsg(SceneEntityMoveScNotify, func() any { return new(proto.SceneEntityMoveScNotify) })
	c.regMsg(ExchangeStaminaScRsp, func() any { return new(proto.ExchangeStaminaScRsp) })
	c.regMsg(DailyFirstMeetPamCsReq, func() any { return new(proto.DailyFirstMeetPamCsReq) })
	c.regMsg(GetEnteredSceneScRsp, func() any { return new(proto.GetEnteredSceneScRsp) })
	c.regMsg(TextJoinSaveScRsp, func() any { return new(proto.TextJoinSaveScRsp) })
	c.regMsg(TakePunkLordPointRewardScRsp, func() any { return new(proto.TakePunkLordPointRewardScRsp) })
	c.regMsg(GetChapterScRsp, func() any { return new(proto.GetChapterScRsp) })
	c.regMsg(MonopolyEventSelectFriendScRsp, func() any { return new(proto.MonopolyEventSelectFriendScRsp) })
	c.regMsg(SyncRogueStatusScNotify, func() any { return new(proto.SyncRogueStatusScNotify) })
	c.regMsg(GetRogueHandbookDataCsReq, func() any { return new(proto.GetRogueHandbookDataCsReq) })
	c.regMsg(QuitLineupScRsp, func() any { return new(proto.QuitLineupScRsp) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(GetMarkItemListCsReq, func() any { return new(proto.GetMarkItemListCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
	c.regMsg(PlayerReturnInfoQueryCsReq, func() any { return new(proto.PlayerReturnInfoQueryCsReq) })
	c.regMsg(BoxingClubChallengeUpdateScNotify, func() any { return new(proto.BoxingClubChallengeUpdateScNotify) })
	c.regMsg(GetPlayerReturnMultiDropInfoScRsp, func() any { return new(proto.GetPlayerReturnMultiDropInfoScRsp) })
	c.regMsg(SceneEnterStageScRsp, func() any { return new(proto.SceneEnterStageScRsp) })
	c.regMsg(MonopolyReRollRandomCsReq, func() any { return new(proto.MonopolyReRollRandomCsReq) })
	c.regMsg(FinishPlotScRsp, func() any { return new(proto.FinishPlotScRsp) })
	c.regMsg(GetMonopolyDailyReportCsReq, func() any { return new(proto.GetMonopolyDailyReportCsReq) })
	c.regMsg(DeleteBlacklistScRsp, func() any { return new(proto.DeleteBlacklistScRsp) })
	c.regMsg(TextJoinBatchSaveScRsp, func() any { return new(proto.TextJoinBatchSaveScRsp) })
	c.regMsg(StopRogueAdventureRoomCsReq, func() any { return new(proto.StopRogueAdventureRoomCsReq) })
	c.regMsg(InterruptMissionEventScRsp, func() any { return new(proto.InterruptMissionEventScRsp) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(ChooseBoxingClubResonanceScRsp, func() any { return new(proto.ChooseBoxingClubResonanceScRsp) })
	c.regMsg(SyncRogueCommonVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueCommonVirtualItemInfoScNotify) })
	c.regMsg(HandleFriendScRsp, func() any { return new(proto.HandleFriendScRsp) })
	c.regMsg(MonopolyContentUpdateScNotify, func() any { return new(proto.MonopolyContentUpdateScNotify) })
	c.regMsg(ChessRogueUpdateMoneyInfoScNotify, func() any { return new(proto.ChessRogueUpdateMoneyInfoScNotify) })
	c.regMsg(MonopolyClickCellCsReq, func() any { return new(proto.MonopolyClickCellCsReq) })
	c.regMsg(WolfBroGameUseBulletScRsp, func() any { return new(proto.WolfBroGameUseBulletScRsp) })
	c.regMsg(HeroBasicTypeChangedNotify, func() any { return new(proto.HeroBasicTypeChangedNotify) })
	c.regMsg(GetRaidInfoScRsp, func() any { return new(proto.GetRaidInfoScRsp) })
	c.regMsg(SetAssistAvatarScRsp, func() any { return new(proto.SetAssistAvatarScRsp) })
	c.regMsg(ExchangeRogueRewardKeyScRsp, func() any { return new(proto.ExchangeRogueRewardKeyScRsp) })
	c.regMsg(SelectInclinationTextScRsp, func() any { return new(proto.SelectInclinationTextScRsp) })
	c.regMsg(TravelBrochureUpdatePasterPosScRsp, func() any { return new(proto.TravelBrochureUpdatePasterPosScRsp) })
	c.regMsg(GetCurBattleInfoCsReq, func() any { return new(proto.GetCurBattleInfoCsReq) })
	c.regMsg(SpaceZooExchangeItemCsReq, func() any { return new(proto.SpaceZooExchangeItemCsReq) })
	c.regMsg(SelectRogueDialogueEventCsReq, func() any { return new(proto.SelectRogueDialogueEventCsReq) })
	c.regMsg(ChessRogueUpdateActionPointScNotify, func() any { return new(proto.ChessRogueUpdateActionPointScNotify) })
	c.regMsg(EnterStrongChallengeActivityStageCsReq, func() any { return new(proto.EnterStrongChallengeActivityStageCsReq) })
	c.regMsg(GetFeverTimeActivityDataScRsp, func() any { return new(proto.GetFeverTimeActivityDataScRsp) })
	c.regMsg(MonopolyBuyGoodsScRsp, func() any { return new(proto.MonopolyBuyGoodsScRsp) })
	c.regMsg(HeliobusSelectSkillCsReq, func() any { return new(proto.HeliobusSelectSkillCsReq) })
	c.regMsg(FinishCosumeItemMissionScRsp, func() any { return new(proto.FinishCosumeItemMissionScRsp) })
	c.regMsg(MakeDrinkCsReq, func() any { return new(proto.MakeDrinkCsReq) })
	c.regMsg(PlayerHeartBeatCsReq, func() any { return new(proto.PlayerHeartBeatCsReq) })
	c.regMsg(SwitchLineupIndexCsReq, func() any { return new(proto.SwitchLineupIndexCsReq) })
	c.regMsg(RemoveStuffFromAreaScRsp, func() any { return new(proto.RemoveStuffFromAreaScRsp) })
	c.regMsg(SelectChessRogueNousSubStoryScRsp, func() any { return new(proto.SelectChessRogueNousSubStoryScRsp) })
	c.regMsg(GetShareDataScRsp, func() any { return new(proto.GetShareDataScRsp) })
	c.regMsg(MuseumTargetRewardNotify, func() any { return new(proto.MuseumTargetRewardNotify) })
	c.regMsg(ChallengeSettleNotify, func() any { return new(proto.ChallengeSettleNotify) })
	c.regMsg(EnterAetherDivideSceneScRsp, func() any { return new(proto.EnterAetherDivideSceneScRsp) })
	c.regMsg(GetPrivateChatHistoryCsReq, func() any { return new(proto.GetPrivateChatHistoryCsReq) })
	c.regMsg(StartRogueCsReq, func() any { return new(proto.StartRogueCsReq) })
	c.regMsg(AetherDivideFinishChallengeScNotify, func() any { return new(proto.AetherDivideFinishChallengeScNotify) })
	c.regMsg(GetFantasticStoryActivityDataScRsp, func() any { return new(proto.GetFantasticStoryActivityDataScRsp) })
	c.regMsg(AlleyTakeEventRewardScRsp, func() any { return new(proto.AlleyTakeEventRewardScRsp) })
	c.regMsg(SpaceZooTakeCsReq, func() any { return new(proto.SpaceZooTakeCsReq) })
	c.regMsg(MakeMissionDrinkScRsp, func() any { return new(proto.MakeMissionDrinkScRsp) })
	c.regMsg(DoGachaCsReq, func() any { return new(proto.DoGachaCsReq) })
	c.regMsg(GetPlayerReturnMultiDropInfoCsReq, func() any { return new(proto.GetPlayerReturnMultiDropInfoCsReq) })
	c.regMsg(GetMbtiReportScRsp, func() any { return new(proto.GetMbtiReportScRsp) })
	c.regMsg(SetMissionEventProgressScRsp, func() any { return new(proto.SetMissionEventProgressScRsp) })
	c.regMsg(ChessRogueReviveAvatarScRsp, func() any { return new(proto.ChessRogueReviveAvatarScRsp) })
	c.regMsg(GetAllSaveRaidCsReq, func() any { return new(proto.GetAllSaveRaidCsReq) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(MonopolyMoveScRsp, func() any { return new(proto.MonopolyMoveScRsp) })
	c.regMsg(FinishTutorialScRsp, func() any { return new(proto.FinishTutorialScRsp) })
	c.regMsg(EnterTrialActivityStageScRsp, func() any { return new(proto.EnterTrialActivityStageScRsp) })
	c.regMsg(ChessRogueCellUpdateNotify, func() any { return new(proto.ChessRogueCellUpdateNotify) })
	c.regMsg(GetPlayerReplayInfoScRsp, func() any { return new(proto.GetPlayerReplayInfoScRsp) })
	c.regMsg(FinishTutorialGuideCsReq, func() any { return new(proto.FinishTutorialGuideCsReq) })
	c.regMsg(FantasticStoryActivityBattleEndScNotify, func() any { return new(proto.FantasticStoryActivityBattleEndScNotify) })
	c.regMsg(UpdateMechanismBarScNotify, func() any { return new(proto.UpdateMechanismBarScNotify) })
	c.regMsg(TravelBrochureGetDataScRsp, func() any { return new(proto.TravelBrochureGetDataScRsp) })
	c.regMsg(ApplyFriendCsReq, func() any { return new(proto.ApplyFriendCsReq) })
	c.regMsg(BuyNpcStuffScRsp, func() any { return new(proto.BuyNpcStuffScRsp) })
	c.regMsg(UnlockTeleportNotify, func() any { return new(proto.UnlockTeleportNotify) })
	c.regMsg(GetAllLineupDataScRsp, func() any { return new(proto.GetAllLineupDataScRsp) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(AddBlacklistCsReq, func() any { return new(proto.AddBlacklistCsReq) })
	c.regMsg(UpdateRogueAdventureRoomScoreScRsp, func() any { return new(proto.UpdateRogueAdventureRoomScoreScRsp) })
	c.regMsg(ClientObjDownloadDataScNotify, func() any { return new(proto.ClientObjDownloadDataScNotify) })
	c.regMsg(UseTreasureDungeonItemCsReq, func() any { return new(proto.UseTreasureDungeonItemCsReq) })
	c.regMsg(GetLineupAvatarDataScRsp, func() any { return new(proto.GetLineupAvatarDataScRsp) })
	c.regMsg(GmTalkScNotify, func() any { return new(proto.GmTalkScNotify) })
	c.regMsg(FinishChessRogueSubStoryScRsp, func() any { return new(proto.FinishChessRogueSubStoryScRsp) })
	c.regMsg(HandleRogueCommonPendingActionCsReq, func() any { return new(proto.HandleRogueCommonPendingActionCsReq) })
	c.regMsg(SearchPlayerScRsp, func() any { return new(proto.SearchPlayerScRsp) })
	c.regMsg(MonopolyActionResultScNotify, func() any { return new(proto.MonopolyActionResultScNotify) })
	c.regMsg(UpdateRedDotDataCsReq, func() any { return new(proto.UpdateRedDotDataCsReq) })
	c.regMsg(AvatarExpUpCsReq, func() any { return new(proto.AvatarExpUpCsReq) })
	c.regMsg(ChessRogueEnterNextLayerCsReq, func() any { return new(proto.ChessRogueEnterNextLayerCsReq) })
	c.regMsg(GetFriendApplyListInfoCsReq, func() any { return new(proto.GetFriendApplyListInfoCsReq) })
	c.regMsg(OpenRogueChestScRsp, func() any { return new(proto.OpenRogueChestScRsp) })
	c.regMsg(TravelBrochureRemovePasterCsReq, func() any { return new(proto.TravelBrochureRemovePasterCsReq) })
	c.regMsg(BattlePassInfoNotify, func() any { return new(proto.BattlePassInfoNotify) })
	c.regMsg(MonopolyRollDiceScRsp, func() any { return new(proto.MonopolyRollDiceScRsp) })
	c.regMsg(TakeAssistRewardCsReq, func() any { return new(proto.TakeAssistRewardCsReq) })
	c.regMsg(FinishPerformSectionIdCsReq, func() any { return new(proto.FinishPerformSectionIdCsReq) })
	c.regMsg(AlleyPlacingGameScRsp, func() any { return new(proto.AlleyPlacingGameScRsp) })
	c.regMsg(MonopolyGameRaiseRatioCsReq, func() any { return new(proto.MonopolyGameRaiseRatioCsReq) })
	c.regMsg(TriggerVoiceScRsp, func() any { return new(proto.TriggerVoiceScRsp) })
	c.regMsg(PVEBattleResultScRsp, func() any { return new(proto.PVEBattleResultScRsp) })
	c.regMsg(ComposeLimitNumCompleteNotify, func() any { return new(proto.ComposeLimitNumCompleteNotify) })
	c.regMsg(StartAetherDivideStageBattleCsReq, func() any { return new(proto.StartAetherDivideStageBattleCsReq) })
	c.regMsg(SendMsgScRsp, func() any { return new(proto.SendMsgScRsp) })
	c.regMsg(ChessRogueConfirmRollCsReq, func() any { return new(proto.ChessRogueConfirmRollCsReq) })
	c.regMsg(WolfBroGameExplodeMonsterScRsp, func() any { return new(proto.WolfBroGameExplodeMonsterScRsp) })
	c.regMsg(GetAvatarDataScRsp, func() any { return new(proto.GetAvatarDataScRsp) })
	c.regMsg(GetPlayerDetailInfoScRsp, func() any { return new(proto.GetPlayerDetailInfoScRsp) })
	c.regMsg(GetShareDataCsReq, func() any { return new(proto.GetShareDataCsReq) })
	c.regMsg(SetCurWaypointCsReq, func() any { return new(proto.SetCurWaypointCsReq) })
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(GetLevelRewardTakenListScRsp, func() any { return new(proto.GetLevelRewardTakenListScRsp) })
	c.regMsg(LogisticsGameScRsp, func() any { return new(proto.LogisticsGameScRsp) })
	c.regMsg(LogisticsDetonateStarSkiffCsReq, func() any { return new(proto.LogisticsDetonateStarSkiffCsReq) })
	c.regMsg(EntityBindPropCsReq, func() any { return new(proto.EntityBindPropCsReq) })
	c.regMsg(GetSocialEventServerCacheCsReq, func() any { return new(proto.GetSocialEventServerCacheCsReq) })
	c.regMsg(SharePunkLordMonsterScRsp, func() any { return new(proto.SharePunkLordMonsterScRsp) })
	c.regMsg(ChessRogueUpdateDicePassiveAccumulateValueScNotify, func() any { return new(proto.ChessRogueUpdateDicePassiveAccumulateValueScNotify) })
	c.regMsg(EnterChessRogueAeonRoomScRsp, func() any { return new(proto.EnterChessRogueAeonRoomScRsp) })
	c.regMsg(ChessRogueGiveUpScRsp, func() any { return new(proto.ChessRogueGiveUpScRsp) })
	c.regMsg(RankUpAvatarScRsp, func() any { return new(proto.RankUpAvatarScRsp) })
	c.regMsg(ComposeItemScRsp, func() any { return new(proto.ComposeItemScRsp) })
	c.regMsg(SetClientRaidTargetCountScRsp, func() any { return new(proto.SetClientRaidTargetCountScRsp) })
	c.regMsg(EnterTreasureDungeonCsReq, func() any { return new(proto.EnterTreasureDungeonCsReq) })
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(SyncAcceptedPamMissionNotify, func() any { return new(proto.SyncAcceptedPamMissionNotify) })
	c.regMsg(TakeAllRewardScRsp, func() any { return new(proto.TakeAllRewardScRsp) })
	c.regMsg(RotateMapCsReq, func() any { return new(proto.RotateMapCsReq) })
	c.regMsg(TakeRogueScoreRewardScRsp, func() any { return new(proto.TakeRogueScoreRewardScRsp) })
	c.regMsg(GetBoxingClubInfoScRsp, func() any { return new(proto.GetBoxingClubInfoScRsp) })
	c.regMsg(ReEnterLastElementStageCsReq, func() any { return new(proto.ReEnterLastElementStageCsReq) })
	c.regMsg(TakePrestigeRewardScRsp, func() any { return new(proto.TakePrestigeRewardScRsp) })
	c.regMsg(PrestigeLevelUpScRsp, func() any { return new(proto.PrestigeLevelUpScRsp) })
	c.regMsg(SetStuffToAreaCsReq, func() any { return new(proto.SetStuffToAreaCsReq) })
	c.regMsg(ChessRogueUpdateReviveInfoScNotify, func() any { return new(proto.ChessRogueUpdateReviveInfoScNotify) })
	c.regMsg(RemoveRotaterCsReq, func() any { return new(proto.RemoveRotaterCsReq) })
	c.regMsg(QuitTreasureDungeonCsReq, func() any { return new(proto.QuitTreasureDungeonCsReq) })
	c.regMsg(AlleyEventEffectNotify, func() any { return new(proto.AlleyEventEffectNotify) })
	c.regMsg(ExpUpRelicScRsp, func() any { return new(proto.ExpUpRelicScRsp) })
	c.regMsg(HandleFriendCsReq, func() any { return new(proto.HandleFriendCsReq) })
	c.regMsg(MuseumFundsChangedScNotify, func() any { return new(proto.MuseumFundsChangedScNotify) })
	c.regMsg(TelevisionActivityDataChangeScNotify, func() any { return new(proto.TelevisionActivityDataChangeScNotify) })
	c.regMsg(GetSceneMapInfoScRsp, func() any { return new(proto.GetSceneMapInfoScRsp) })
	c.regMsg(GetMissionEventDataScRsp, func() any { return new(proto.GetMissionEventDataScRsp) })
	c.regMsg(TakeMultipleExpeditionRewardCsReq, func() any { return new(proto.TakeMultipleExpeditionRewardCsReq) })
	c.regMsg(GetNpcTakenRewardCsReq, func() any { return new(proto.GetNpcTakenRewardCsReq) })
	c.regMsg(RecoverAllLineupScRsp, func() any { return new(proto.RecoverAllLineupScRsp) })
	c.regMsg(MuseumTargetStartNotify, func() any { return new(proto.MuseumTargetStartNotify) })
	c.regMsg(GetMovieRacingDataCsReq, func() any { return new(proto.GetMovieRacingDataCsReq) })
	c.regMsg(SyncRogueAdventureRoomInfoScNotify, func() any { return new(proto.SyncRogueAdventureRoomInfoScNotify) })
	c.regMsg(MatchBoxingClubOpponentScRsp, func() any { return new(proto.MatchBoxingClubOpponentScRsp) })
	c.regMsg(StartBattleCollegeScRsp, func() any { return new(proto.StartBattleCollegeScRsp) })
	c.regMsg(GetLoginActivityCsReq, func() any { return new(proto.GetLoginActivityCsReq) })
	c.regMsg(ChessRogueStartCsReq, func() any { return new(proto.ChessRogueStartCsReq) })
	c.regMsg(EnterTrialActivityStageCsReq, func() any { return new(proto.EnterTrialActivityStageCsReq) })
	c.regMsg(UnlockedAreaMapScNotify, func() any { return new(proto.UnlockedAreaMapScNotify) })
	c.regMsg(SetGroupCustomSaveDataCsReq, func() any { return new(proto.SetGroupCustomSaveDataCsReq) })
	c.regMsg(ChessRogueSkipTeachingLevelCsReq, func() any { return new(proto.ChessRogueSkipTeachingLevelCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(PlayerLoginScRsp, func() any { return new(proto.PlayerLoginScRsp) })
	c.regMsg(RestoreWolfBroGameArchiveScRsp, func() any { return new(proto.RestoreWolfBroGameArchiveScRsp) })
	c.regMsg(LeaveMapRotationRegionScRsp, func() any { return new(proto.LeaveMapRotationRegionScRsp) })
	c.regMsg(SyncChessRogueMainStoryFinishScNotify, func() any { return new(proto.SyncChessRogueMainStoryFinishScNotify) })
	c.regMsg(StartTimedCocoonStageCsReq, func() any { return new(proto.StartTimedCocoonStageCsReq) })
	c.regMsg(GiveUpBoxingClubChallengeScRsp, func() any { return new(proto.GiveUpBoxingClubChallengeScRsp) })
	c.regMsg(ExchangeGachaCeilingCsReq, func() any { return new(proto.ExchangeGachaCeilingCsReq) })
	c.regMsg(VirtualLineupDestroyNotify, func() any { return new(proto.VirtualLineupDestroyNotify) })
	c.regMsg(HeartDialScriptChangeScNotify, func() any { return new(proto.HeartDialScriptChangeScNotify) })
	c.regMsg(SetHeadIconCsReq, func() any { return new(proto.SetHeadIconCsReq) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(GetSocialEventServerCacheScRsp, func() any { return new(proto.GetSocialEventServerCacheScRsp) })
	c.regMsg(AcceptActivityExpeditionScRsp, func() any { return new(proto.AcceptActivityExpeditionScRsp) })
	c.regMsg(PrivateMsgOfflineUsersScNotify, func() any { return new(proto.PrivateMsgOfflineUsersScNotify) })
	c.regMsg(BuyGoodsCsReq, func() any { return new(proto.BuyGoodsCsReq) })
	c.regMsg(GetMonopolyInfoCsReq, func() any { return new(proto.GetMonopolyInfoCsReq) })
	c.regMsg(StartCocoonStageScRsp, func() any { return new(proto.StartCocoonStageScRsp) })
	c.regMsg(StartBattleCollegeCsReq, func() any { return new(proto.StartBattleCollegeCsReq) })
	c.regMsg(EnterAdventureCsReq, func() any { return new(proto.EnterAdventureCsReq) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(MonopolyRollRandomCsReq, func() any { return new(proto.MonopolyRollRandomCsReq) })
	c.regMsg(MonopolyUpgradeAssetScRsp, func() any { return new(proto.MonopolyUpgradeAssetScRsp) })
	c.regMsg(GetBoxingClubInfoCsReq, func() any { return new(proto.GetBoxingClubInfoCsReq) })
	c.regMsg(TakeAllRewardCsReq, func() any { return new(proto.TakeAllRewardCsReq) })
	c.regMsg(EnterRogueCsReq, func() any { return new(proto.EnterRogueCsReq) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardCsReq) })
	c.regMsg(LogisticsDetonateStarSkiffScRsp, func() any { return new(proto.LogisticsDetonateStarSkiffScRsp) })
	c.regMsg(WaypointShowNewCsNotify, func() any { return new(proto.WaypointShowNewCsNotify) })
	c.regMsg(DeactivateFarmElementCsReq, func() any { return new(proto.DeactivateFarmElementCsReq) })
	c.regMsg(TakeLoginActivityRewardScRsp, func() any { return new(proto.TakeLoginActivityRewardScRsp) })
	c.regMsg(SpaceZooDeleteCatCsReq, func() any { return new(proto.SpaceZooDeleteCatCsReq) })
	c.regMsg(RogueModifierSelectCellScRsp, func() any { return new(proto.RogueModifierSelectCellScRsp) })
	c.regMsg(ExchangeRogueRewardKeyCsReq, func() any { return new(proto.ExchangeRogueRewardKeyCsReq) })
	c.regMsg(HeliobusSnsPostScRsp, func() any { return new(proto.HeliobusSnsPostScRsp) })
	c.regMsg(StartRogueScRsp, func() any { return new(proto.StartRogueScRsp) })
	c.regMsg(AetherDivideSkillItemScNotify, func() any { return new(proto.AetherDivideSkillItemScNotify) })
	c.regMsg(SetClientRaidTargetCountCsReq, func() any { return new(proto.SetClientRaidTargetCountCsReq) })
	c.regMsg(SpringRecoverScRsp, func() any { return new(proto.SpringRecoverScRsp) })
	c.regMsg(PlayerHeartBeatScRsp, func() any { return new(proto.PlayerHeartBeatScRsp) })
	c.regMsg(GetMissionStatusScRsp, func() any { return new(proto.GetMissionStatusScRsp) })
	c.regMsg(SetFriendRemarkNameCsReq, func() any { return new(proto.SetFriendRemarkNameCsReq) })
	c.regMsg(EnterAdventureScRsp, func() any { return new(proto.EnterAdventureScRsp) })
	c.regMsg(InteractChargerScRsp, func() any { return new(proto.InteractChargerScRsp) })
	c.regMsg(MonopolyTakePhaseRewardScRsp, func() any { return new(proto.MonopolyTakePhaseRewardScRsp) })
	c.regMsg(MonopolySocialEventEffectScNotify, func() any { return new(proto.MonopolySocialEventEffectScNotify) })
	c.regMsg(TakeExpeditionRewardCsReq, func() any { return new(proto.TakeExpeditionRewardCsReq) })
	c.regMsg(StaminaInfoScNotify, func() any { return new(proto.StaminaInfoScNotify) })
	c.regMsg(MarkItemScRsp, func() any { return new(proto.MarkItemScRsp) })
	c.regMsg(GetDrinkMakerDataScRsp, func() any { return new(proto.GetDrinkMakerDataScRsp) })
	c.regMsg(TakeTalkRewardCsReq, func() any { return new(proto.TakeTalkRewardCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(GetSpringRecoverDataScRsp, func() any { return new(proto.GetSpringRecoverDataScRsp) })
	c.regMsg(TakeRogueEndlessActivityPointRewardCsReq, func() any { return new(proto.TakeRogueEndlessActivityPointRewardCsReq) })
	c.regMsg(EndDrinkMakerSequenceCsReq, func() any { return new(proto.EndDrinkMakerSequenceCsReq) })
	c.regMsg(HandleRogueCommonPendingActionScRsp, func() any { return new(proto.HandleRogueCommonPendingActionScRsp) })
	c.regMsg(SetAssistAvatarCsReq, func() any { return new(proto.SetAssistAvatarCsReq) })
	c.regMsg(DeleteSocialEventServerCacheScRsp, func() any { return new(proto.DeleteSocialEventServerCacheScRsp) })
	c.regMsg(EnhanceRogueBuffCsReq, func() any { return new(proto.EnhanceRogueBuffCsReq) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(MuseumRandomEventSelectScRsp, func() any { return new(proto.MuseumRandomEventSelectScRsp) })
	c.regMsg(BatchMarkChatEmojiScRsp, func() any { return new(proto.BatchMarkChatEmojiScRsp) })
	c.regMsg(GetTutorialScRsp, func() any { return new(proto.GetTutorialScRsp) })
	c.regMsg(FeverTimeActivityBattleEndScNotify, func() any { return new(proto.FeverTimeActivityBattleEndScNotify) })
	c.regMsg(MonopolyCheatDiceCsReq, func() any { return new(proto.MonopolyCheatDiceCsReq) })
	c.regMsg(GmTalkCsReq, func() any { return new(proto.GmTalkCsReq) })
	c.regMsg(RechargeSuccNotify, func() any { return new(proto.RechargeSuccNotify) })
	c.regMsg(GetRogueAeonInfoCsReq, func() any { return new(proto.GetRogueAeonInfoCsReq) })
	c.regMsg(EnterSceneScRsp, func() any { return new(proto.EnterSceneScRsp) })
	c.regMsg(GetCurAssistCsReq, func() any { return new(proto.GetCurAssistCsReq) })
	c.regMsg(GetOfferingInfoCsReq, func() any { return new(proto.GetOfferingInfoCsReq) })
	c.regMsg(FinishPerformSectionIdScRsp, func() any { return new(proto.FinishPerformSectionIdScRsp) })
	c.regMsg(RecoverAllLineupCsReq, func() any { return new(proto.RecoverAllLineupCsReq) })
	c.regMsg(RetcodeNotify, func() any { return new(proto.RetcodeNotify) })
	c.regMsg(GetMbtiReportCsReq, func() any { return new(proto.GetMbtiReportCsReq) })
	c.regMsg(TakeRollShopRewardScRsp, func() any { return new(proto.TakeRollShopRewardScRsp) })
	c.regMsg(FinishEmotionDialoguePerformanceScRsp, func() any { return new(proto.FinishEmotionDialoguePerformanceScRsp) })
	c.regMsg(GetPlayerBoardDataScRsp, func() any { return new(proto.GetPlayerBoardDataScRsp) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(GetRogueEndlessActivityDataCsReq, func() any { return new(proto.GetRogueEndlessActivityDataCsReq) })
	c.regMsg(GetTelevisionActivityDataScRsp, func() any { return new(proto.GetTelevisionActivityDataScRsp) })
	c.regMsg(WolfBroGameUseBulletCsReq, func() any { return new(proto.WolfBroGameUseBulletCsReq) })
	c.regMsg(GetLineupAvatarDataCsReq, func() any { return new(proto.GetLineupAvatarDataCsReq) })
	c.regMsg(ChessRogueGiveUpCsReq, func() any { return new(proto.ChessRogueGiveUpCsReq) })
	c.regMsg(SelectInclinationTextCsReq, func() any { return new(proto.SelectInclinationTextCsReq) })
	c.regMsg(SelectChatBubbleScRsp, func() any { return new(proto.SelectChatBubbleScRsp) })
	c.regMsg(RotateMapScRsp, func() any { return new(proto.RotateMapScRsp) })
	c.regMsg(GetPhoneDataScRsp, func() any { return new(proto.GetPhoneDataScRsp) })
	c.regMsg(DailyTaskDataScNotify, func() any { return new(proto.DailyTaskDataScNotify) })
	c.regMsg(RogueModifierUpdateNotify, func() any { return new(proto.RogueModifierUpdateNotify) })
	c.regMsg(SelectChatBubbleCsReq, func() any { return new(proto.SelectChatBubbleCsReq) })
	c.regMsg(EndDrinkMakerSequenceScRsp, func() any { return new(proto.EndDrinkMakerSequenceScRsp) })
	c.regMsg(CancelMarkItemNotify, func() any { return new(proto.CancelMarkItemNotify) })
	c.regMsg(PlayerReturnSignCsReq, func() any { return new(proto.PlayerReturnSignCsReq) })
	c.regMsg(GameplayCounterCountDownScRsp, func() any { return new(proto.GameplayCounterCountDownScRsp) })
	c.regMsg(MuseumTakeCollectRewardScRsp, func() any { return new(proto.MuseumTakeCollectRewardScRsp) })
	c.regMsg(GmTalkScRsp, func() any { return new(proto.GmTalkScRsp) })
	c.regMsg(MarkReadMailScRsp, func() any { return new(proto.MarkReadMailScRsp) })
	c.regMsg(StartRaidScRsp, func() any { return new(proto.StartRaidScRsp) })
	c.regMsg(ChessRogueUpdateUnlockLevelScNotify, func() any { return new(proto.ChessRogueUpdateUnlockLevelScNotify) })
	c.regMsg(StartChallengeCsReq, func() any { return new(proto.StartChallengeCsReq) })
	c.regMsg(ScenePlaneEventScNotify, func() any { return new(proto.ScenePlaneEventScNotify) })
	c.regMsg(SetSignatureScRsp, func() any { return new(proto.SetSignatureScRsp) })
	c.regMsg(ChessRogueGiveUpRollCsReq, func() any { return new(proto.ChessRogueGiveUpRollCsReq) })
	c.regMsg(GetWolfBroGameDataScRsp, func() any { return new(proto.GetWolfBroGameDataScRsp) })
	c.regMsg(MonopolyGetRegionProgressCsReq, func() any { return new(proto.MonopolyGetRegionProgressCsReq) })
	c.regMsg(GetCurBattleInfoScRsp, func() any { return new(proto.GetCurBattleInfoScRsp) })
	c.regMsg(GetChessRogueStoryInfoCsReq, func() any { return new(proto.GetChessRogueStoryInfoCsReq) })
	c.regMsg(GetRollShopInfoCsReq, func() any { return new(proto.GetRollShopInfoCsReq) })
	c.regMsg(SyncRogueSeasonFinishScNotify, func() any { return new(proto.SyncRogueSeasonFinishScNotify) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardScRsp, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardScRsp) })
	c.regMsg(MuseumTakeCollectRewardCsReq, func() any { return new(proto.MuseumTakeCollectRewardCsReq) })
	c.regMsg(TakeQuestRewardCsReq, func() any { return new(proto.TakeQuestRewardCsReq) })
	c.regMsg(GetMailScRsp, func() any { return new(proto.GetMailScRsp) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(LeaveRogueScRsp, func() any { return new(proto.LeaveRogueScRsp) })
	c.regMsg(ChessRogueQuitScRsp, func() any { return new(proto.ChessRogueQuitScRsp) })
	c.regMsg(RaidInfoNotify, func() any { return new(proto.RaidInfoNotify) })
	c.regMsg(InteractPropCsReq, func() any { return new(proto.InteractPropCsReq) })
	c.regMsg(EnhanceCommonRogueBuffScRsp, func() any { return new(proto.EnhanceCommonRogueBuffScRsp) })
	c.regMsg(GetChessRogueStoryInfoScRsp, func() any { return new(proto.GetChessRogueStoryInfoScRsp) })
	c.regMsg(QuestRecordScNotify, func() any { return new(proto.QuestRecordScNotify) })
	c.regMsg(GetAllServerPrefsDataScRsp, func() any { return new(proto.GetAllServerPrefsDataScRsp) })
	c.regMsg(LeaveRogueCsReq, func() any { return new(proto.LeaveRogueCsReq) })
	c.regMsg(LockRelicScRsp, func() any { return new(proto.LockRelicScRsp) })
	c.regMsg(ChessRogueRollDiceCsReq, func() any { return new(proto.ChessRogueRollDiceCsReq) })
	c.regMsg(GetMonsterResearchActivityDataScRsp, func() any { return new(proto.GetMonsterResearchActivityDataScRsp) })
	c.regMsg(ChessRogueNousEditDiceCsReq, func() any { return new(proto.ChessRogueNousEditDiceCsReq) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(GetChessRogueNousStoryInfoCsReq, func() any { return new(proto.GetChessRogueNousStoryInfoCsReq) })
	c.regMsg(AlleyEventChangeNotify, func() any { return new(proto.AlleyEventChangeNotify) })
	c.regMsg(GetFriendListInfoCsReq, func() any { return new(proto.GetFriendListInfoCsReq) })
	c.regMsg(GetCurChallengeScRsp, func() any { return new(proto.GetCurChallengeScRsp) })
	c.regMsg(UpdateServerPrefsDataCsReq, func() any { return new(proto.UpdateServerPrefsDataCsReq) })
	c.regMsg(GetTrainVisitorRegisterScRsp, func() any { return new(proto.GetTrainVisitorRegisterScRsp) })
	c.regMsg(PlayerReturnTakePointRewardScRsp, func() any { return new(proto.PlayerReturnTakePointRewardScRsp) })
	c.regMsg(GetTutorialGuideScRsp, func() any { return new(proto.GetTutorialGuideScRsp) })
	c.regMsg(ChessRogueEnterCsReq, func() any { return new(proto.ChessRogueEnterCsReq) })
	c.regMsg(AetherDivideRefreshEndlessScNotify, func() any { return new(proto.AetherDivideRefreshEndlessScNotify) })
	c.regMsg(SyncRogueDialogueEventDataScNotify, func() any { return new(proto.SyncRogueDialogueEventDataScNotify) })
	c.regMsg(GetKilledPunkLordMonsterDataScRsp, func() any { return new(proto.GetKilledPunkLordMonsterDataScRsp) })
	c.regMsg(ChessRogueEnterCellScRsp, func() any { return new(proto.ChessRogueEnterCellScRsp) })
	c.regMsg(GetGachaInfoScRsp, func() any { return new(proto.GetGachaInfoScRsp) })
	c.regMsg(GetRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetRogueBuffEnhanceInfoScRsp) })
	c.regMsg(GetAetherDivideInfoScRsp, func() any { return new(proto.GetAetherDivideInfoScRsp) })
	c.regMsg(BuyBpLevelCsReq, func() any { return new(proto.BuyBpLevelCsReq) })
	c.regMsg(GetAssistHistoryScRsp, func() any { return new(proto.GetAssistHistoryScRsp) })
	c.regMsg(GetSaveLogisticsMapScRsp, func() any { return new(proto.GetSaveLogisticsMapScRsp) })
	c.regMsg(FinishItemIdScRsp, func() any { return new(proto.FinishItemIdScRsp) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(TakeRogueEndlessActivityAllBonusRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityAllBonusRewardScRsp) })
	c.regMsg(HeartDialTraceScriptCsReq, func() any { return new(proto.HeartDialTraceScriptCsReq) })
	c.regMsg(SceneCastSkillCsReq, func() any { return new(proto.SceneCastSkillCsReq) })
	c.regMsg(GetGachaCeilingCsReq, func() any { return new(proto.GetGachaCeilingCsReq) })
	c.regMsg(GetMuseumInfoCsReq, func() any { return new(proto.GetMuseumInfoCsReq) })
	c.regMsg(SyncClientResVersionCsReq, func() any { return new(proto.SyncClientResVersionCsReq) })
	c.regMsg(MakeDrinkScRsp, func() any { return new(proto.MakeDrinkScRsp) })
	c.regMsg(DeleteSummonUnitScRsp, func() any { return new(proto.DeleteSummonUnitScRsp) })
	c.regMsg(UpdateEnergyScNotify, func() any { return new(proto.UpdateEnergyScNotify) })
	c.regMsg(HeliobusSnsCommentScRsp, func() any { return new(proto.HeliobusSnsCommentScRsp) })
	c.regMsg(GetChallengeRaidInfoScRsp, func() any { return new(proto.GetChallengeRaidInfoScRsp) })
	c.regMsg(GetNpcStatusScRsp, func() any { return new(proto.GetNpcStatusScRsp) })
	c.regMsg(SceneCastSkillCostMpScRsp, func() any { return new(proto.SceneCastSkillCostMpScRsp) })
	c.regMsg(ChessRogueUpdateBoardScNotify, func() any { return new(proto.ChessRogueUpdateBoardScNotify) })
	c.regMsg(EnterMapRotationRegionCsReq, func() any { return new(proto.EnterMapRotationRegionCsReq) })
	c.regMsg(GetAssistListScRsp, func() any { return new(proto.GetAssistListScRsp) })
	c.regMsg(TakeRogueEndlessActivityPointRewardScRsp, func() any { return new(proto.TakeRogueEndlessActivityPointRewardScRsp) })
	c.regMsg(SetAetherDivideLineUpScRsp, func() any { return new(proto.SetAetherDivideLineUpScRsp) })
	c.regMsg(RogueModifierAddNotify, func() any { return new(proto.RogueModifierAddNotify) })
	c.regMsg(GetWaypointCsReq, func() any { return new(proto.GetWaypointCsReq) })
	c.regMsg(StartTimedCocoonStageScRsp, func() any { return new(proto.StartTimedCocoonStageScRsp) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
	c.regMsg(PunkLordMonsterInfoScNotify, func() any { return new(proto.PunkLordMonsterInfoScNotify) })
	c.regMsg(GetFeverTimeActivityDataCsReq, func() any { return new(proto.GetFeverTimeActivityDataCsReq) })
	c.regMsg(FinishCurTurnCsReq, func() any { return new(proto.FinishCurTurnCsReq) })
	c.regMsg(GetCurAssistScRsp, func() any { return new(proto.GetCurAssistScRsp) })
	c.regMsg(SetSpringRecoverConfigCsReq, func() any { return new(proto.SetSpringRecoverConfigCsReq) })
	c.regMsg(TravelBrochureSetCustomValueScRsp, func() any { return new(proto.TravelBrochureSetCustomValueScRsp) })
	c.regMsg(GetArchiveDataScRsp, func() any { return new(proto.GetArchiveDataScRsp) })
	c.regMsg(SetFriendRemarkNameScRsp, func() any { return new(proto.SetFriendRemarkNameScRsp) })
	c.regMsg(LeaveChallengeScRsp, func() any { return new(proto.LeaveChallengeScRsp) })
	c.regMsg(TakeOfferingRewardScRsp, func() any { return new(proto.TakeOfferingRewardScRsp) })
	c.regMsg(PrepareRogueAdventureRoomScRsp, func() any { return new(proto.PrepareRogueAdventureRoomScRsp) })
	c.regMsg(HeliobusSnsReadCsReq, func() any { return new(proto.HeliobusSnsReadCsReq) })
	c.regMsg(EnterFantasticStoryActivityStageScRsp, func() any { return new(proto.EnterFantasticStoryActivityStageScRsp) })
	c.regMsg(TakeTalkRewardScRsp, func() any { return new(proto.TakeTalkRewardScRsp) })
	c.regMsg(AlleyPlacingGameCsReq, func() any { return new(proto.AlleyPlacingGameCsReq) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(GetReplayTokenCsReq, func() any { return new(proto.GetReplayTokenCsReq) })
	c.regMsg(DelMailScRsp, func() any { return new(proto.DelMailScRsp) })
	c.regMsg(StartFinishSubMissionScNotify, func() any { return new(proto.StartFinishSubMissionScNotify) })
	c.regMsg(TakeRogueAeonLevelRewardScRsp, func() any { return new(proto.TakeRogueAeonLevelRewardScRsp) })
	c.regMsg(AetherDivideRefreshEndlessCsReq, func() any { return new(proto.AetherDivideRefreshEndlessCsReq) })
	c.regMsg(GetRaidInfoCsReq, func() any { return new(proto.GetRaidInfoCsReq) })
	c.regMsg(MarkItemCsReq, func() any { return new(proto.MarkItemCsReq) })
	c.regMsg(HeliobusSnsPostCsReq, func() any { return new(proto.HeliobusSnsPostCsReq) })
	c.regMsg(MuseumRandomEventQueryCsReq, func() any { return new(proto.MuseumRandomEventQueryCsReq) })
	c.regMsg(MonopolyBuyGoodsCsReq, func() any { return new(proto.MonopolyBuyGoodsCsReq) })
	c.regMsg(MonopolyGetRegionProgressScRsp, func() any { return new(proto.MonopolyGetRegionProgressScRsp) })
	c.regMsg(SyncRogueCommonPendingActionScNotify, func() any { return new(proto.SyncRogueCommonPendingActionScNotify) })
	c.regMsg(SyncRogueCommonActionResultScNotify, func() any { return new(proto.SyncRogueCommonActionResultScNotify) })
	c.regMsg(SetLanguageScRsp, func() any { return new(proto.SetLanguageScRsp) })
	c.regMsg(GetFriendLoginInfoCsReq, func() any { return new(proto.GetFriendLoginInfoCsReq) })
	c.regMsg(GetSceneMapInfoCsReq, func() any { return new(proto.GetSceneMapInfoCsReq) })
	c.regMsg(StartTimedFarmElementScRsp, func() any { return new(proto.StartTimedFarmElementScRsp) })
	c.regMsg(FinishChessRogueNousSubStoryScRsp, func() any { return new(proto.FinishChessRogueNousSubStoryScRsp) })
	c.regMsg(SyncChessRogueNousSubStoryScNotify, func() any { return new(proto.SyncChessRogueNousSubStoryScNotify) })
	c.regMsg(StartAlleyEventCsReq, func() any { return new(proto.StartAlleyEventCsReq) })
	c.regMsg(StartAetherDivideStageBattleScRsp, func() any { return new(proto.StartAetherDivideStageBattleScRsp) })
	c.regMsg(RaidKickByServerScNotify, func() any { return new(proto.RaidKickByServerScNotify) })
	c.regMsg(GetChessRogueBuffEnhanceInfoScRsp, func() any { return new(proto.GetChessRogueBuffEnhanceInfoScRsp) })
	c.regMsg(SyncRogueRewardInfoScNotify, func() any { return new(proto.SyncRogueRewardInfoScNotify) })
	c.regMsg(JoinLineupCsReq, func() any { return new(proto.JoinLineupCsReq) })
	c.regMsg(TakeQuestOptionalRewardCsReq, func() any { return new(proto.TakeQuestOptionalRewardCsReq) })
	c.regMsg(ChessRoguePickAvatarScRsp, func() any { return new(proto.ChessRoguePickAvatarScRsp) })
	c.regMsg(ChessRogueCheatRollScRsp, func() any { return new(proto.ChessRogueCheatRollScRsp) })
	c.regMsg(MuseumRandomEventQueryScRsp, func() any { return new(proto.MuseumRandomEventQueryScRsp) })
	c.regMsg(BoxingClubRewardScNotify, func() any { return new(proto.BoxingClubRewardScNotify) })
	c.regMsg(GetHeartDialInfoScRsp, func() any { return new(proto.GetHeartDialInfoScRsp) })
	c.regMsg(ChessRogueNousDiceSurfaceUnlockNotify, func() any { return new(proto.ChessRogueNousDiceSurfaceUnlockNotify) })
	c.regMsg(GetMonopolyFriendRankingListCsReq, func() any { return new(proto.GetMonopolyFriendRankingListCsReq) })
	c.regMsg(SellItemScRsp, func() any { return new(proto.SellItemScRsp) })
	c.regMsg(HeliobusEnterBattleScRsp, func() any { return new(proto.HeliobusEnterBattleScRsp) })
	c.regMsg(TakeApRewardCsReq, func() any { return new(proto.TakeApRewardCsReq) })
	c.regMsg(MonopolySttUpdateScNotify, func() any { return new(proto.MonopolySttUpdateScNotify) })
	c.regMsg(SceneEntityMoveCsReq, func() any { return new(proto.SceneEntityMoveCsReq) })
	c.regMsg(UpdateMapRotationDataScNotify, func() any { return new(proto.UpdateMapRotationDataScNotify) })
	c.regMsg(ReportPlayerScRsp, func() any { return new(proto.ReportPlayerScRsp) })
	c.regMsg(GetShopListCsReq, func() any { return new(proto.GetShopListCsReq) })
	c.regMsg(GetPunkLordBattleRecordScRsp, func() any { return new(proto.GetPunkLordBattleRecordScRsp) })
	c.regMsg(SyncRogueGetItemScNotify, func() any { return new(proto.SyncRogueGetItemScNotify) })
	c.regMsg(QueryProductInfoScRsp, func() any { return new(proto.QueryProductInfoScRsp) })
	c.regMsg(GetRogueTalentInfoScRsp, func() any { return new(proto.GetRogueTalentInfoScRsp) })
	c.regMsg(MonthCardRewardNotify, func() any { return new(proto.MonthCardRewardNotify) })
	c.regMsg(CancelCacheNotifyScRsp, func() any { return new(proto.CancelCacheNotifyScRsp) })
	c.regMsg(GameplayCounterCountDownCsReq, func() any { return new(proto.GameplayCounterCountDownCsReq) })
	c.regMsg(GetAlleyInfoCsReq, func() any { return new(proto.GetAlleyInfoCsReq) })
	c.regMsg(MarkReadMailCsReq, func() any { return new(proto.MarkReadMailCsReq) })
	c.regMsg(SpaceZooBornScRsp, func() any { return new(proto.SpaceZooBornScRsp) })
	c.regMsg(AetherDivideLineupScNotify, func() any { return new(proto.AetherDivideLineupScNotify) })
	c.regMsg(DressAvatarSkinScRsp, func() any { return new(proto.DressAvatarSkinScRsp) })
	c.regMsg(ExchangeStaminaCsReq, func() any { return new(proto.ExchangeStaminaCsReq) })
	c.regMsg(PVEBattleResultCsReq, func() any { return new(proto.PVEBattleResultCsReq) })
	c.regMsg(SpaceZooDataScRsp, func() any { return new(proto.SpaceZooDataScRsp) })
	c.regMsg(LockEquipmentCsReq, func() any { return new(proto.LockEquipmentCsReq) })
	c.regMsg(SubmitEmotionItemScRsp, func() any { return new(proto.SubmitEmotionItemScRsp) })
	c.regMsg(BuyRogueShopBuffScRsp, func() any { return new(proto.BuyRogueShopBuffScRsp) })
	c.regMsg(PlayerReturnTakeRewardCsReq, func() any { return new(proto.PlayerReturnTakeRewardCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoScRsp, func() any { return new(proto.GetAetherDivideChallengeInfoScRsp) })
	c.regMsg(EnterSceneByServerScNotify, func() any { return new(proto.EnterSceneByServerScNotify) })
	c.regMsg(SetHeroBasicTypeScRsp, func() any { return new(proto.SetHeroBasicTypeScRsp) })
	c.regMsg(AcceptMultipleExpeditionScRsp, func() any { return new(proto.AcceptMultipleExpeditionScRsp) })
	c.regMsg(GetPunkLordBattleRecordCsReq, func() any { return new(proto.GetPunkLordBattleRecordCsReq) })
	c.regMsg(DressAvatarScRsp, func() any { return new(proto.DressAvatarScRsp) })
	c.regMsg(FinishTutorialCsReq, func() any { return new(proto.FinishTutorialCsReq) })
	c.regMsg(StartAetherDivideSceneBattleCsReq, func() any { return new(proto.StartAetherDivideSceneBattleCsReq) })
	c.regMsg(SetPlayerInfoScRsp, func() any { return new(proto.SetPlayerInfoScRsp) })
	c.regMsg(TakeCityShopRewardScRsp, func() any { return new(proto.TakeCityShopRewardScRsp) })
	c.regMsg(EnterAetherDivideSceneCsReq, func() any { return new(proto.EnterAetherDivideSceneCsReq) })
	c.regMsg(SummonPunkLordMonsterScRsp, func() any { return new(proto.SummonPunkLordMonsterScRsp) })
	c.regMsg(HeliobusSnsUpdateScNotify, func() any { return new(proto.HeliobusSnsUpdateScNotify) })
	c.regMsg(ChessRogueLeaveScRsp, func() any { return new(proto.ChessRogueLeaveScRsp) })
	c.regMsg(HeliobusActivityDataScRsp, func() any { return new(proto.HeliobusActivityDataScRsp) })
	c.regMsg(AcceptedPamMissionExpireCsReq, func() any { return new(proto.AcceptedPamMissionExpireCsReq) })
	c.regMsg(ChessRogueFinishCurRoomNotify, func() any { return new(proto.ChessRogueFinishCurRoomNotify) })
	c.regMsg(GetRogueBuffEnhanceInfoCsReq, func() any { return new(proto.GetRogueBuffEnhanceInfoCsReq) })
	c.regMsg(UpdatePlayerSettingScRsp, func() any { return new(proto.UpdatePlayerSettingScRsp) })
	c.regMsg(GameplayCounterRecoverScRsp, func() any { return new(proto.GameplayCounterRecoverScRsp) })
	c.regMsg(MonopolyLikeScNotify, func() any { return new(proto.MonopolyLikeScNotify) })
	c.regMsg(MarkChatEmojiCsReq, func() any { return new(proto.MarkChatEmojiCsReq) })
	c.regMsg(SetHeadIconScRsp, func() any { return new(proto.SetHeadIconScRsp) })
	c.regMsg(MonopolyGameGachaScRsp, func() any { return new(proto.MonopolyGameGachaScRsp) })
	c.regMsg(AntiAddictScNotify, func() any { return new(proto.AntiAddictScNotify) })
	c.regMsg(ChessRogueQueryAeonDimensionsCsReq, func() any { return new(proto.ChessRogueQueryAeonDimensionsCsReq) })
	c.regMsg(AceAntiCheaterScRsp, func() any { return new(proto.AceAntiCheaterScRsp) })
	c.regMsg(GetCurSceneInfoScRsp, func() any { return new(proto.GetCurSceneInfoScRsp) })
	c.regMsg(ChallengeRaidNotify, func() any { return new(proto.ChallengeRaidNotify) })
	c.regMsg(FinishQuestCsReq, func() any { return new(proto.FinishQuestCsReq) })
	c.regMsg(WolfBroGamePickupBulletScRsp, func() any { return new(proto.WolfBroGamePickupBulletScRsp) })
	c.regMsg(UnlockTutorialScRsp, func() any { return new(proto.UnlockTutorialScRsp) })
	c.regMsg(EnterFeverTimeActivityStageCsReq, func() any { return new(proto.EnterFeverTimeActivityStageCsReq) })
	c.regMsg(StrongChallengeActivityBattleEndScNotify, func() any { return new(proto.StrongChallengeActivityBattleEndScNotify) })
	c.regMsg(ChessRogueNousEditDiceScRsp, func() any { return new(proto.ChessRogueNousEditDiceScRsp) })
	c.regMsg(MonopolyRollDiceCsReq, func() any { return new(proto.MonopolyRollDiceCsReq) })
	c.regMsg(ReturnLastTownCsReq, func() any { return new(proto.ReturnLastTownCsReq) })
	c.regMsg(FinishCurTurnScRsp, func() any { return new(proto.FinishCurTurnScRsp) })
	c.regMsg(GetTrainVisitorBehaviorScRsp, func() any { return new(proto.GetTrainVisitorBehaviorScRsp) })
	c.regMsg(TreasureDungeonDataScNotify, func() any { return new(proto.TreasureDungeonDataScNotify) })
	c.regMsg(GetChatFriendHistoryScRsp, func() any { return new(proto.GetChatFriendHistoryScRsp) })
	c.regMsg(GetCurLineupDataCsReq, func() any { return new(proto.GetCurLineupDataCsReq) })
	c.regMsg(GetAetherDivideInfoCsReq, func() any { return new(proto.GetAetherDivideInfoCsReq) })
	c.regMsg(GetUpdatedArchiveDataCsReq, func() any { return new(proto.GetUpdatedArchiveDataCsReq) })
	c.regMsg(SetSpringRecoverConfigScRsp, func() any { return new(proto.SetSpringRecoverConfigScRsp) })
	c.regMsg(SetForbidOtherApplyFriendScRsp, func() any { return new(proto.SetForbidOtherApplyFriendScRsp) })
	c.regMsg(GetRollShopInfoScRsp, func() any { return new(proto.GetRollShopInfoScRsp) })
	c.regMsg(GetCurSceneInfoCsReq, func() any { return new(proto.GetCurSceneInfoCsReq) })
	c.regMsg(MuseumDispatchFinishedScNotify, func() any { return new(proto.MuseumDispatchFinishedScNotify) })
	c.regMsg(GetRogueInitialScoreScRsp, func() any { return new(proto.GetRogueInitialScoreScRsp) })
	c.regMsg(AetherDivideRefreshEndlessScRsp, func() any { return new(proto.AetherDivideRefreshEndlessScRsp) })
	c.regMsg(PrestigeLevelUpCsReq, func() any { return new(proto.PrestigeLevelUpCsReq) })
	c.regMsg(GetChallengeRaidInfoCsReq, func() any { return new(proto.GetChallengeRaidInfoCsReq) })
	c.regMsg(QuitRogueScRsp, func() any { return new(proto.QuitRogueScRsp) })
	c.regMsg(EnterChessRogueAeonRoomCsReq, func() any { return new(proto.EnterChessRogueAeonRoomCsReq) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(CancelCacheNotifyCsReq, func() any { return new(proto.CancelCacheNotifyCsReq) })
	c.regMsg(GetMailCsReq, func() any { return new(proto.GetMailCsReq) })
	c.regMsg(HeliobusUnlockSkillScNotify, func() any { return new(proto.HeliobusUnlockSkillScNotify) })
	c.regMsg(SubmitMonsterResearchActivityMaterialScRsp, func() any { return new(proto.SubmitMonsterResearchActivityMaterialScRsp) })
	c.regMsg(ChessRogueRollDiceScRsp, func() any { return new(proto.ChessRogueRollDiceScRsp) })
	c.regMsg(GetJukeboxDataScRsp, func() any { return new(proto.GetJukeboxDataScRsp) })
	c.regMsg(ArchiveWolfBroGameCsReq, func() any { return new(proto.ArchiveWolfBroGameCsReq) })
	c.regMsg(ChangeScriptEmotionCsReq, func() any { return new(proto.ChangeScriptEmotionCsReq) })
	c.regMsg(SubMissionRewardScNotify, func() any { return new(proto.SubMissionRewardScNotify) })
	c.regMsg(SwitchLineupIndexScRsp, func() any { return new(proto.SwitchLineupIndexScRsp) })
	c.regMsg(ChessRogueStartScRsp, func() any { return new(proto.ChessRogueStartScRsp) })
	c.regMsg(EnterMapRotationRegionScRsp, func() any { return new(proto.EnterMapRotationRegionScRsp) })
	c.regMsg(EnterRogueEndlessActivityStageScRsp, func() any { return new(proto.EnterRogueEndlessActivityStageScRsp) })
	c.regMsg(TravelBrochureSetPageDescStatusScRsp, func() any { return new(proto.TravelBrochureSetPageDescStatusScRsp) })
	c.regMsg(WolfBroGameExplodeMonsterCsReq, func() any { return new(proto.WolfBroGameExplodeMonsterCsReq) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffScRsp, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffScRsp) })
	c.regMsg(MonopolyTakePhaseRewardCsReq, func() any { return new(proto.MonopolyTakePhaseRewardCsReq) })
	c.regMsg(SubmitMonsterResearchActivityMaterialCsReq, func() any { return new(proto.SubmitMonsterResearchActivityMaterialCsReq) })
	c.regMsg(TakePunkLordPointRewardCsReq, func() any { return new(proto.TakePunkLordPointRewardCsReq) })
	c.regMsg(SetCurInteractEntityScRsp, func() any { return new(proto.SetCurInteractEntityScRsp) })
	c.regMsg(SaveLogisticsCsReq, func() any { return new(proto.SaveLogisticsCsReq) })
	c.regMsg(SwapLineupScRsp, func() any { return new(proto.SwapLineupScRsp) })
	c.regMsg(FinishTalkMissionScRsp, func() any { return new(proto.FinishTalkMissionScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardScRsp, func() any { return new(proto.MonopolyTakeRaffleTicketRewardScRsp) })
	c.regMsg(LogisticsInfoScNotify, func() any { return new(proto.LogisticsInfoScNotify) })
	c.regMsg(TakeLoginActivityRewardCsReq, func() any { return new(proto.TakeLoginActivityRewardCsReq) })
	c.regMsg(ReviveRogueAvatarCsReq, func() any { return new(proto.ReviveRogueAvatarCsReq) })
	c.regMsg(GetChallengeScRsp, func() any { return new(proto.GetChallengeScRsp) })
	c.regMsg(GetReplayTokenScRsp, func() any { return new(proto.GetReplayTokenScRsp) })
	c.regMsg(GetLoginActivityScRsp, func() any { return new(proto.GetLoginActivityScRsp) })
	c.regMsg(GetWaypointScRsp, func() any { return new(proto.GetWaypointScRsp) })
	c.regMsg(FinishAeonDialogueGroupCsReq, func() any { return new(proto.FinishAeonDialogueGroupCsReq) })
	c.regMsg(AvatarExpUpScRsp, func() any { return new(proto.AvatarExpUpScRsp) })
	c.regMsg(GetTrialActivityDataCsReq, func() any { return new(proto.GetTrialActivityDataCsReq) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(SavePointsInfoNotify, func() any { return new(proto.SavePointsInfoNotify) })
	c.regMsg(ChangeLineupLeaderScRsp, func() any { return new(proto.ChangeLineupLeaderScRsp) })
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(SpaceZooExchangeItemScRsp, func() any { return new(proto.SpaceZooExchangeItemScRsp) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(OpenRogueChestCsReq, func() any { return new(proto.OpenRogueChestCsReq) })
	c.regMsg(ChessRogueReRollDiceScRsp, func() any { return new(proto.ChessRogueReRollDiceScRsp) })
	c.regMsg(UnlockTutorialGuideCsReq, func() any { return new(proto.UnlockTutorialGuideCsReq) })
	c.regMsg(TrainVisitorRewardSendNotify, func() any { return new(proto.TrainVisitorRewardSendNotify) })
	c.regMsg(MonopolySelectOptionCsReq, func() any { return new(proto.MonopolySelectOptionCsReq) })
	c.regMsg(MonopolyCheatDiceScRsp, func() any { return new(proto.MonopolyCheatDiceScRsp) })
	c.regMsg(HeliobusSnsReadScRsp, func() any { return new(proto.HeliobusSnsReadScRsp) })
	c.regMsg(PlayerReturnSignScRsp, func() any { return new(proto.PlayerReturnSignScRsp) })
	c.regMsg(EnterRogueEndlessActivityStageCsReq, func() any { return new(proto.EnterRogueEndlessActivityStageCsReq) })
	c.regMsg(GetAllRedDotDataScRsp, func() any { return new(proto.GetAllRedDotDataScRsp) })
	c.regMsg(RaidCollectionDataCsReq, func() any { return new(proto.RaidCollectionDataCsReq) })
	c.regMsg(AcceptExpeditionScRsp, func() any { return new(proto.AcceptExpeditionScRsp) })
	c.regMsg(SetLanguageCsReq, func() any { return new(proto.SetLanguageCsReq) })
	c.regMsg(SyncDeleteFriendScNotify, func() any { return new(proto.SyncDeleteFriendScNotify) })
	c.regMsg(SelectPhoneThemeScRsp, func() any { return new(proto.SelectPhoneThemeScRsp) })
	c.regMsg(InteractChargerCsReq, func() any { return new(proto.InteractChargerCsReq) })
	c.regMsg(ChessRogueUpdateLevelBaseInfoScNotify, func() any { return new(proto.ChessRogueUpdateLevelBaseInfoScNotify) })
	c.regMsg(SceneEntityTeleportCsReq, func() any { return new(proto.SceneEntityTeleportCsReq) })
	c.regMsg(TakeOffAvatarSkinScRsp, func() any { return new(proto.TakeOffAvatarSkinScRsp) })
	c.regMsg(InteractTreasureDungeonGridScRsp, func() any { return new(proto.InteractTreasureDungeonGridScRsp) })
	c.regMsg(ChangeLineupLeaderCsReq, func() any { return new(proto.ChangeLineupLeaderCsReq) })
	c.regMsg(HeliobusChallengeUpdateScNotify, func() any { return new(proto.HeliobusChallengeUpdateScNotify) })
	c.regMsg(ChessRogueConfirmRollScRsp, func() any { return new(proto.ChessRogueConfirmRollScRsp) })
	c.regMsg(SetLineupNameCsReq, func() any { return new(proto.SetLineupNameCsReq) })
	c.regMsg(GetStrongChallengeActivityDataCsReq, func() any { return new(proto.GetStrongChallengeActivityDataCsReq) })
	c.regMsg(HealPoolInfoNotify, func() any { return new(proto.HealPoolInfoNotify) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(GetFightActivityDataScRsp, func() any { return new(proto.GetFightActivityDataScRsp) })
	c.regMsg(SyncRogueAeonScNotify, func() any { return new(proto.SyncRogueAeonScNotify) })
	c.regMsg(DestroyItemCsReq, func() any { return new(proto.DestroyItemCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotScRsp, func() any { return new(proto.SwitchAetherDivideLineUpSlotScRsp) })
	c.regMsg(ChessRogueChangeyAeonDimensionNotify, func() any { return new(proto.ChessRogueChangeyAeonDimensionNotify) })
	c.regMsg(LogisticsScoreRewardSyncInfoScNotify, func() any { return new(proto.LogisticsScoreRewardSyncInfoScNotify) })
	c.regMsg(UnlockTutorialCsReq, func() any { return new(proto.UnlockTutorialCsReq) })
	c.regMsg(MissionGroupWarnScNotify, func() any { return new(proto.MissionGroupWarnScNotify) })
	c.regMsg(ChessRogueGiveUpRollScRsp, func() any { return new(proto.ChessRogueGiveUpRollScRsp) })
	c.regMsg(PlayerLogoutCsReq, func() any { return new(proto.PlayerLogoutCsReq) })
	c.regMsg(EnterStrongChallengeActivityStageScRsp, func() any { return new(proto.EnterStrongChallengeActivityStageScRsp) })
	c.regMsg(RogueModifierSelectCellCsReq, func() any { return new(proto.RogueModifierSelectCellCsReq) })
	c.regMsg(GetSingleRedDotParamGroupCsReq, func() any { return new(proto.GetSingleRedDotParamGroupCsReq) })
	c.regMsg(LeaveTrialActivityCsReq, func() any { return new(proto.LeaveTrialActivityCsReq) })
	c.regMsg(RogueNpcDisappearScRsp, func() any { return new(proto.RogueNpcDisappearScRsp) })
	c.regMsg(AlleyShipUnlockScNotify, func() any { return new(proto.AlleyShipUnlockScNotify) })
	c.regMsg(GetExhibitScNotify, func() any { return new(proto.GetExhibitScNotify) })
	c.regMsg(ExchangeRogueBuffWithMiracleScRsp, func() any { return new(proto.ExchangeRogueBuffWithMiracleScRsp) })
	c.regMsg(UpgradeAreaScRsp, func() any { return new(proto.UpgradeAreaScRsp) })
	c.regMsg(EnableRogueTalentScRsp, func() any { return new(proto.EnableRogueTalentScRsp) })
	c.regMsg(TakeCityShopRewardCsReq, func() any { return new(proto.TakeCityShopRewardCsReq) })
	c.regMsg(MuseumInfoChangedScNotify, func() any { return new(proto.MuseumInfoChangedScNotify) })
	c.regMsg(PlayerReturnStartScNotify, func() any { return new(proto.PlayerReturnStartScNotify) })
	c.regMsg(GetRogueScoreRewardInfoScRsp, func() any { return new(proto.GetRogueScoreRewardInfoScRsp) })
	c.regMsg(TravelBrochureApplyPasterCsReq, func() any { return new(proto.TravelBrochureApplyPasterCsReq) })
	c.regMsg(TakeAllApRewardScRsp, func() any { return new(proto.TakeAllApRewardScRsp) })
	c.regMsg(GetSecretKeyInfoScRsp, func() any { return new(proto.GetSecretKeyInfoScRsp) })
	c.regMsg(TakeRogueEventHandbookRewardCsReq, func() any { return new(proto.TakeRogueEventHandbookRewardCsReq) })
	c.regMsg(FinishTalkMissionCsReq, func() any { return new(proto.FinishTalkMissionCsReq) })
	c.regMsg(GetMultipleDropInfoScRsp, func() any { return new(proto.GetMultipleDropInfoScRsp) })
	c.regMsg(GetRogueTalentInfoCsReq, func() any { return new(proto.GetRogueTalentInfoCsReq) })
	c.regMsg(NewAssistHistoryNotify, func() any { return new(proto.NewAssistHistoryNotify) })
	c.regMsg(SetNicknameScRsp, func() any { return new(proto.SetNicknameScRsp) })
	c.regMsg(PickRogueAvatarScRsp, func() any { return new(proto.PickRogueAvatarScRsp) })
	c.regMsg(SceneCastSkillMpUpdateScNotify, func() any { return new(proto.SceneCastSkillMpUpdateScNotify) })
	c.regMsg(GetUnlockTeleportCsReq, func() any { return new(proto.GetUnlockTeleportCsReq) })
	c.regMsg(SyncRogueExploreWinScNotify, func() any { return new(proto.SyncRogueExploreWinScNotify) })
	c.regMsg(GetStageLineupScRsp, func() any { return new(proto.GetStageLineupScRsp) })
	c.regMsg(SyncRoguePickAvatarInfoScNotify, func() any { return new(proto.SyncRoguePickAvatarInfoScNotify) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(ApplyFriendScRsp, func() any { return new(proto.ApplyFriendScRsp) })
	c.regMsg(GetPlayerReplayInfoCsReq, func() any { return new(proto.GetPlayerReplayInfoCsReq) })
	c.regMsg(MonopolySelectOptionScRsp, func() any { return new(proto.MonopolySelectOptionScRsp) })
	c.regMsg(BuyGoodsScRsp, func() any { return new(proto.BuyGoodsScRsp) })
	c.regMsg(ChessRogueEnterCellCsReq, func() any { return new(proto.ChessRogueEnterCellCsReq) })
	c.regMsg(GetPunkLordMonsterDataCsReq, func() any { return new(proto.GetPunkLordMonsterDataCsReq) })
	c.regMsg(SelectChessRogueNousSubStoryCsReq, func() any { return new(proto.SelectChessRogueNousSubStoryCsReq) })
	c.regMsg(GetMonopolyDailyReportScRsp, func() any { return new(proto.GetMonopolyDailyReportScRsp) })
	c.regMsg(ChessRogueQuitCsReq, func() any { return new(proto.ChessRogueQuitCsReq) })
	c.regMsg(GetRecyleTimeCsReq, func() any { return new(proto.GetRecyleTimeCsReq) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(GetFriendRecommendListInfoCsReq, func() any { return new(proto.GetFriendRecommendListInfoCsReq) })
	c.regMsg(ClientObjUploadCsReq, func() any { return new(proto.ClientObjUploadCsReq) })
	c.regMsg(SetPlayerInfoCsReq, func() any { return new(proto.SetPlayerInfoCsReq) })
	c.regMsg(RogueEndlessActivityBattleEndScNotify, func() any { return new(proto.RogueEndlessActivityBattleEndScNotify) })
	c.regMsg(GetSaveRaidScRsp, func() any { return new(proto.GetSaveRaidScRsp) })
	c.regMsg(MonopolyLikeCsReq, func() any { return new(proto.MonopolyLikeCsReq) })
	c.regMsg(PickRogueAvatarCsReq, func() any { return new(proto.PickRogueAvatarCsReq) })
	c.regMsg(SummonPunkLordMonsterCsReq, func() any { return new(proto.SummonPunkLordMonsterCsReq) })
	c.regMsg(PlayBackGroundMusicScRsp, func() any { return new(proto.PlayBackGroundMusicScRsp) })
	c.regMsg(TravelBrochureSelectMessageScRsp, func() any { return new(proto.TravelBrochureSelectMessageScRsp) })
	c.regMsg(TakeRogueScoreRewardCsReq, func() any { return new(proto.TakeRogueScoreRewardCsReq) })
	c.regMsg(GetTrainVisitorRegisterCsReq, func() any { return new(proto.GetTrainVisitorRegisterCsReq) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
	c.regMsg(GetAllSaveRaidScRsp, func() any { return new(proto.GetAllSaveRaidScRsp) })
	c.regMsg(StartTrialActivityCsReq, func() any { return new(proto.StartTrialActivityCsReq) })
	c.regMsg(GetUpdatedArchiveDataScRsp, func() any { return new(proto.GetUpdatedArchiveDataScRsp) })
	c.regMsg(LeaveTrialActivityScRsp, func() any { return new(proto.LeaveTrialActivityScRsp) })
	c.regMsg(RestoreWolfBroGameArchiveCsReq, func() any { return new(proto.RestoreWolfBroGameArchiveCsReq) })
	c.regMsg(StartWolfBroGameCsReq, func() any { return new(proto.StartWolfBroGameCsReq) })
	c.regMsg(ChessRogueSelectCellScRsp, func() any { return new(proto.ChessRogueSelectCellScRsp) })
	c.regMsg(FightActivityDataChangeScNotify, func() any { return new(proto.FightActivityDataChangeScNotify) })
	c.regMsg(GetExpeditionDataCsReq, func() any { return new(proto.GetExpeditionDataCsReq) })
	c.regMsg(SwitchAetherDivideLineUpSlotCsReq, func() any { return new(proto.SwitchAetherDivideLineUpSlotCsReq) })
	c.regMsg(AlleyTakeEventRewardCsReq, func() any { return new(proto.AlleyTakeEventRewardCsReq) })
	c.regMsg(SpaceZooOpCatteryScRsp, func() any { return new(proto.SpaceZooOpCatteryScRsp) })
	c.regMsg(TravelBrochurePageResetCsReq, func() any { return new(proto.TravelBrochurePageResetCsReq) })
	c.regMsg(AcceptMultipleExpeditionCsReq, func() any { return new(proto.AcceptMultipleExpeditionCsReq) })
	c.regMsg(DoGachaInRollShopScRsp, func() any { return new(proto.DoGachaInRollShopScRsp) })
	c.regMsg(TrainRefreshTimeNotify, func() any { return new(proto.TrainRefreshTimeNotify) })
	c.regMsg(SyncAddBlacklistScNotify, func() any { return new(proto.SyncAddBlacklistScNotify) })
	c.regMsg(UnlockAvatarSkinScNotify, func() any { return new(proto.UnlockAvatarSkinScNotify) })
	c.regMsg(AetherDivideTakeChallengeRewardCsReq, func() any { return new(proto.AetherDivideTakeChallengeRewardCsReq) })
	c.regMsg(TakeTrialActivityRewardCsReq, func() any { return new(proto.TakeTrialActivityRewardCsReq) })
	c.regMsg(TakeActivityExpeditionRewardCsReq, func() any { return new(proto.TakeActivityExpeditionRewardCsReq) })
	c.regMsg(TakeRogueMiracleHandbookRewardScRsp, func() any { return new(proto.TakeRogueMiracleHandbookRewardScRsp) })
	c.regMsg(GetEnhanceCommonRogueBuffInfoScRsp, func() any { return new(proto.GetEnhanceCommonRogueBuffInfoScRsp) })
	c.regMsg(SetAetherDivideLineUpCsReq, func() any { return new(proto.SetAetherDivideLineUpCsReq) })
	c.regMsg(TakeKilledPunkLordMonsterScoreScRsp, func() any { return new(proto.TakeKilledPunkLordMonsterScoreScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(ComposeSelectedRelicCsReq, func() any { return new(proto.ComposeSelectedRelicCsReq) })
	c.regMsg(DeleteSocialEventServerCacheCsReq, func() any { return new(proto.DeleteSocialEventServerCacheCsReq) })
	c.regMsg(UpdateFloorSavedValueNotify, func() any { return new(proto.UpdateFloorSavedValueNotify) })
	c.regMsg(ChangeScriptEmotionScRsp, func() any { return new(proto.ChangeScriptEmotionScRsp) })
	c.regMsg(StartBoxingClubBattleCsReq, func() any { return new(proto.StartBoxingClubBattleCsReq) })
	c.regMsg(TakeActivityExpeditionRewardScRsp, func() any { return new(proto.TakeActivityExpeditionRewardScRsp) })
	c.regMsg(GetChatFriendHistoryCsReq, func() any { return new(proto.GetChatFriendHistoryCsReq) })
	c.regMsg(CancelActivityExpeditionScRsp, func() any { return new(proto.CancelActivityExpeditionScRsp) })
	c.regMsg(GetCurChallengeCsReq, func() any { return new(proto.GetCurChallengeCsReq) })
	c.regMsg(InteractTreasureDungeonGridCsReq, func() any { return new(proto.InteractTreasureDungeonGridCsReq) })
	c.regMsg(AcceptedPamMissionExpireScRsp, func() any { return new(proto.AcceptedPamMissionExpireScRsp) })
	c.regMsg(ChessRogueNousDiceUpdateNotify, func() any { return new(proto.ChessRogueNousDiceUpdateNotify) })
	c.regMsg(GetMissionStatusCsReq, func() any { return new(proto.GetMissionStatusCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupCsReq, func() any { return new(proto.SetBoxingClubResonanceLineupCsReq) })
	c.regMsg(FinishCosumeItemMissionCsReq, func() any { return new(proto.FinishCosumeItemMissionCsReq) })
	c.regMsg(RaidCollectionDataScNotify, func() any { return new(proto.RaidCollectionDataScNotify) })
	c.regMsg(MonopolyGetRaffleTicketScRsp, func() any { return new(proto.MonopolyGetRaffleTicketScRsp) })
	c.regMsg(TakeTrialActivityRewardScRsp, func() any { return new(proto.TakeTrialActivityRewardScRsp) })
	c.regMsg(TakeMonsterResearchActivityRewardCsReq, func() any { return new(proto.TakeMonsterResearchActivityRewardCsReq) })
	c.regMsg(SceneEntityMoveScRsp, func() any { return new(proto.SceneEntityMoveScRsp) })
	c.regMsg(StartBoxingClubBattleScRsp, func() any { return new(proto.StartBoxingClubBattleScRsp) })
	c.regMsg(GetNpcMessageGroupScRsp, func() any { return new(proto.GetNpcMessageGroupScRsp) })
	c.regMsg(RogueModifierStageStartNotify, func() any { return new(proto.RogueModifierStageStartNotify) })
	c.regMsg(GetRogueInfoCsReq, func() any { return new(proto.GetRogueInfoCsReq) })
	c.regMsg(TakeChallengeRaidRewardScRsp, func() any { return new(proto.TakeChallengeRaidRewardScRsp) })
	c.regMsg(CurAssistChangedNotify, func() any { return new(proto.CurAssistChangedNotify) })
	c.regMsg(GetFirstTalkNpcScRsp, func() any { return new(proto.GetFirstTalkNpcScRsp) })
	c.regMsg(GetMonopolyFriendRankingListScRsp, func() any { return new(proto.GetMonopolyFriendRankingListScRsp) })
	c.regMsg(DiscardRelicCsReq, func() any { return new(proto.DiscardRelicCsReq) })
	c.regMsg(AlleyGuaranteedFundsScRsp, func() any { return new(proto.AlleyGuaranteedFundsScRsp) })
	c.regMsg(PunkLordMonsterKilledNotify, func() any { return new(proto.PunkLordMonsterKilledNotify) })
	c.regMsg(GetChatEmojiListCsReq, func() any { return new(proto.GetChatEmojiListCsReq) })
	c.regMsg(UpgradeAreaStatCsReq, func() any { return new(proto.UpgradeAreaStatCsReq) })
	c.regMsg(TravelBrochureUpdatePasterPosCsReq, func() any { return new(proto.TravelBrochureUpdatePasterPosCsReq) })
	c.regMsg(DiscardRelicScRsp, func() any { return new(proto.DiscardRelicScRsp) })
	c.regMsg(RogueNpcDisappearCsReq, func() any { return new(proto.RogueNpcDisappearCsReq) })
	c.regMsg(HeliobusInfoChangedScNotify, func() any { return new(proto.HeliobusInfoChangedScNotify) })
	c.regMsg(ChessRogueUpdateAeonModifierValueScNotify, func() any { return new(proto.ChessRogueUpdateAeonModifierValueScNotify) })
	c.regMsg(DressAvatarCsReq, func() any { return new(proto.DressAvatarCsReq) })
	c.regMsg(GetActivityScheduleConfigScRsp, func() any { return new(proto.GetActivityScheduleConfigScRsp) })
	c.regMsg(SpaceZooMutateScRsp, func() any { return new(proto.SpaceZooMutateScRsp) })
	c.regMsg(RelicRecommendScRsp, func() any { return new(proto.RelicRecommendScRsp) })
	c.regMsg(GetGachaInfoCsReq, func() any { return new(proto.GetGachaInfoCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataCsReq, func() any { return new(proto.GetTreasureDungeonActivityDataCsReq) })
	c.regMsg(GetPrivateChatHistoryScRsp, func() any { return new(proto.GetPrivateChatHistoryScRsp) })
	c.regMsg(ChessRogueReviveAvatarCsReq, func() any { return new(proto.ChessRogueReviveAvatarCsReq) })
	c.regMsg(BuyNpcStuffCsReq, func() any { return new(proto.BuyNpcStuffCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(PlayerLoginFinishScRsp, func() any { return new(proto.PlayerLoginFinishScRsp) })
	c.regMsg(LockRelicCsReq, func() any { return new(proto.LockRelicCsReq) })
	c.regMsg(StartWolfBroGameScRsp, func() any { return new(proto.StartWolfBroGameScRsp) })
	c.regMsg(GetBattleCollegeDataScRsp, func() any { return new(proto.GetBattleCollegeDataScRsp) })
	c.regMsg(EnterRogueScRsp, func() any { return new(proto.EnterRogueScRsp) })
	c.regMsg(FinishSectionIdScRsp, func() any { return new(proto.FinishSectionIdScRsp) })
	c.regMsg(GetRogueAeonInfoScRsp, func() any { return new(proto.GetRogueAeonInfoScRsp) })
	c.regMsg(RemoveRotaterScRsp, func() any { return new(proto.RemoveRotaterScRsp) })
	c.regMsg(PlayerLoginFinishCsReq, func() any { return new(proto.PlayerLoginFinishCsReq) })
	c.regMsg(GetBattleCollegeDataCsReq, func() any { return new(proto.GetBattleCollegeDataCsReq) })
	c.regMsg(GetAetherDivideChallengeInfoCsReq, func() any { return new(proto.GetAetherDivideChallengeInfoCsReq) })
	c.regMsg(MonopolyMoveCsReq, func() any { return new(proto.MonopolyMoveCsReq) })
	c.regMsg(BuyRogueShopMiracleScRsp, func() any { return new(proto.BuyRogueShopMiracleScRsp) })
	c.regMsg(SetAssistCsReq, func() any { return new(proto.SetAssistCsReq) })
	c.regMsg(GetDrinkMakerDataCsReq, func() any { return new(proto.GetDrinkMakerDataCsReq) })
	c.regMsg(TeleportToMissionResetPointScRsp, func() any { return new(proto.TeleportToMissionResetPointScRsp) })
	c.regMsg(ClientDownloadDataScNotify, func() any { return new(proto.ClientDownloadDataScNotify) })
	c.regMsg(SceneEnterStageCsReq, func() any { return new(proto.SceneEnterStageCsReq) })
	c.regMsg(ShowNewSupplementVisitorScRsp, func() any { return new(proto.ShowNewSupplementVisitorScRsp) })
	c.regMsg(GetQuestRecordCsReq, func() any { return new(proto.GetQuestRecordCsReq) })
	c.regMsg(PlayerGetTokenCsReq, func() any { return new(proto.PlayerGetTokenCsReq) })
	c.regMsg(SetIsDisplayAvatarInfoCsReq, func() any { return new(proto.SetIsDisplayAvatarInfoCsReq) })
	c.regMsg(ChessRogueQueryAeonDimensionsScRsp, func() any { return new(proto.ChessRogueQueryAeonDimensionsScRsp) })
	c.regMsg(EnterFightActivityStageCsReq, func() any { return new(proto.EnterFightActivityStageCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(HeliobusSnsLikeCsReq, func() any { return new(proto.HeliobusSnsLikeCsReq) })
	c.regMsg(GetPlatformPlayerInfoScRsp, func() any { return new(proto.GetPlatformPlayerInfoScRsp) })
	c.regMsg(StartAetherDivideSceneBattleScRsp, func() any { return new(proto.StartAetherDivideSceneBattleScRsp) })
	c.regMsg(AetherDivideSpiritExpUpCsReq, func() any { return new(proto.AetherDivideSpiritExpUpCsReq) })
	c.regMsg(QuitWolfBroGameCsReq, func() any { return new(proto.QuitWolfBroGameCsReq) })
	c.regMsg(ExchangeGachaCeilingScRsp, func() any { return new(proto.ExchangeGachaCeilingScRsp) })
	c.regMsg(GetActivityScheduleConfigCsReq, func() any { return new(proto.GetActivityScheduleConfigCsReq) })
	c.regMsg(TravelBrochureSetPageDescStatusCsReq, func() any { return new(proto.TravelBrochureSetPageDescStatusCsReq) })
	c.regMsg(ChessRogueReRollDiceCsReq, func() any { return new(proto.ChessRogueReRollDiceCsReq) })
	c.regMsg(TakePrestigeRewardCsReq, func() any { return new(proto.TakePrestigeRewardCsReq) })
	c.regMsg(MissionAcceptScNotify, func() any { return new(proto.MissionAcceptScNotify) })
	c.regMsg(LeaveMapRotationRegionScNotify, func() any { return new(proto.LeaveMapRotationRegionScNotify) })
	c.regMsg(MonopolyUpgradeAssetCsReq, func() any { return new(proto.MonopolyUpgradeAssetCsReq) })
	c.regMsg(AlleyGuaranteedFundsCsReq, func() any { return new(proto.AlleyGuaranteedFundsCsReq) })
	c.regMsg(WolfBroGameDataChangeScNotify, func() any { return new(proto.WolfBroGameDataChangeScNotify) })
	c.regMsg(ReplaceLineupScRsp, func() any { return new(proto.ReplaceLineupScRsp) })
	c.regMsg(EnterFantasticStoryActivityStageCsReq, func() any { return new(proto.EnterFantasticStoryActivityStageCsReq) })
	c.regMsg(FinishFirstTalkNpcScRsp, func() any { return new(proto.FinishFirstTalkNpcScRsp) })
	c.regMsg(ChessRogueSelectCellCsReq, func() any { return new(proto.ChessRogueSelectCellCsReq) })
	c.regMsg(ShareScRsp, func() any { return new(proto.ShareScRsp) })
	c.regMsg(UpdateFeatureSwitchScNotify, func() any { return new(proto.UpdateFeatureSwitchScNotify) })
	c.regMsg(DressRelicAvatarCsReq, func() any { return new(proto.DressRelicAvatarCsReq) })
	c.regMsg(GetKilledPunkLordMonsterDataCsReq, func() any { return new(proto.GetKilledPunkLordMonsterDataCsReq) })
	c.regMsg(GetPlatformPlayerInfoCsReq, func() any { return new(proto.GetPlatformPlayerInfoCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketScRsp, func() any { return new(proto.MonopolyScrachRaffleTicketScRsp) })
	c.regMsg(SyncChessRogueNousMainStoryScNotify, func() any { return new(proto.SyncChessRogueNousMainStoryScNotify) })
	c.regMsg(GetAssistHistoryCsReq, func() any { return new(proto.GetAssistHistoryCsReq) })
	c.regMsg(GetVideoVersionKeyScRsp, func() any { return new(proto.GetVideoVersionKeyScRsp) })
	c.regMsg(ChessRogueLayerAccountInfoNotify, func() any { return new(proto.ChessRogueLayerAccountInfoNotify) })
	c.regMsg(MonopolyReRollRandomScRsp, func() any { return new(proto.MonopolyReRollRandomScRsp) })
	c.regMsg(GetFightActivityDataCsReq, func() any { return new(proto.GetFightActivityDataCsReq) })
	c.regMsg(ChessRogueSelectBpCsReq, func() any { return new(proto.ChessRogueSelectBpCsReq) })
	c.regMsg(DressAvatarSkinCsReq, func() any { return new(proto.DressAvatarSkinCsReq) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(GameplayCounterRecoverCsReq, func() any { return new(proto.GameplayCounterRecoverCsReq) })
	c.regMsg(GetChapterCsReq, func() any { return new(proto.GetChapterCsReq) })
	c.regMsg(FinishChessRogueSubStoryCsReq, func() any { return new(proto.FinishChessRogueSubStoryCsReq) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(TakeOffRelicScRsp, func() any { return new(proto.TakeOffRelicScRsp) })
	c.regMsg(RelicRecommendCsReq, func() any { return new(proto.RelicRecommendCsReq) })
	c.regMsg(AlleyShipUsedCountScNotify, func() any { return new(proto.AlleyShipUsedCountScNotify) })
	c.regMsg(SpaceZooOpCatteryCsReq, func() any { return new(proto.SpaceZooOpCatteryCsReq) })
	c.regMsg(FinishAeonDialogueGroupScRsp, func() any { return new(proto.FinishAeonDialogueGroupScRsp) })
	c.regMsg(StartCocoonStageCsReq, func() any { return new(proto.StartCocoonStageCsReq) })
	c.regMsg(GetRogueInfoScRsp, func() any { return new(proto.GetRogueInfoScRsp) })
	c.regMsg(AlleyOrderChangedScNotify, func() any { return new(proto.AlleyOrderChangedScNotify) })
	c.regMsg(DeleteSummonUnitCsReq, func() any { return new(proto.DeleteSummonUnitCsReq) })
	c.regMsg(LeaveChallengeCsReq, func() any { return new(proto.LeaveChallengeCsReq) })
	c.regMsg(RankUpAvatarCsReq, func() any { return new(proto.RankUpAvatarCsReq) })
	c.regMsg(GetMuseumInfoScRsp, func() any { return new(proto.GetMuseumInfoScRsp) })
	c.regMsg(GetShopListScRsp, func() any { return new(proto.GetShopListScRsp) })
	c.regMsg(TravelBrochurePageUnlockScNotify, func() any { return new(proto.TravelBrochurePageUnlockScNotify) })
	c.regMsg(GetLoginChatInfoCsReq, func() any { return new(proto.GetLoginChatInfoCsReq) })
	c.regMsg(SelectPhoneThemeCsReq, func() any { return new(proto.SelectPhoneThemeCsReq) })
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(TravelBrochureSelectMessageCsReq, func() any { return new(proto.TravelBrochureSelectMessageCsReq) })
	c.regMsg(GetDailyActiveInfoScRsp, func() any { return new(proto.GetDailyActiveInfoScRsp) })
	c.regMsg(ChessRogueQueryScRsp, func() any { return new(proto.ChessRogueQueryScRsp) })
	c.regMsg(MonopolyCellUpdateNotify, func() any { return new(proto.MonopolyCellUpdateNotify) })
	c.regMsg(UpdateTrackMainMissionIdScRsp, func() any { return new(proto.UpdateTrackMainMissionIdScRsp) })
	c.regMsg(TakeApRewardScRsp, func() any { return new(proto.TakeApRewardScRsp) })
	c.regMsg(SetNicknameCsReq, func() any { return new(proto.SetNicknameCsReq) })
	c.regMsg(LeaveAetherDivideSceneScRsp, func() any { return new(proto.LeaveAetherDivideSceneScRsp) })
	c.regMsg(TakeBpRewardCsReq, func() any { return new(proto.TakeBpRewardCsReq) })
	c.regMsg(GetTreasureDungeonActivityDataScRsp, func() any { return new(proto.GetTreasureDungeonActivityDataScRsp) })
	c.regMsg(DeactivateFarmElementScRsp, func() any { return new(proto.DeactivateFarmElementScRsp) })
	c.regMsg(EnhanceChessRogueBuffScRsp, func() any { return new(proto.EnhanceChessRogueBuffScRsp) })
	c.regMsg(SpringRecoverSingleAvatarCsReq, func() any { return new(proto.SpringRecoverSingleAvatarCsReq) })
	c.regMsg(WolfBroGameActivateBulletScRsp, func() any { return new(proto.WolfBroGameActivateBulletScRsp) })
	c.regMsg(SceneEntityTeleportScRsp, func() any { return new(proto.SceneEntityTeleportScRsp) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(MonopolyRollRandomScRsp, func() any { return new(proto.MonopolyRollRandomScRsp) })
	c.regMsg(QuitWolfBroGameScRsp, func() any { return new(proto.QuitWolfBroGameScRsp) })
	c.regMsg(GetUnlockTeleportScRsp, func() any { return new(proto.GetUnlockTeleportScRsp) })
	c.regMsg(SpaceZooMutateCsReq, func() any { return new(proto.SpaceZooMutateCsReq) })
	c.regMsg(MonopolyConfirmRandomCsReq, func() any { return new(proto.MonopolyConfirmRandomCsReq) })
	c.regMsg(GetMainMissionCustomValueScRsp, func() any { return new(proto.GetMainMissionCustomValueScRsp) })
	c.regMsg(PlayerSyncScNotify, func() any { return new(proto.PlayerSyncScNotify) })
	c.regMsg(TakeMultipleExpeditionRewardScRsp, func() any { return new(proto.TakeMultipleExpeditionRewardScRsp) })
	c.regMsg(ClearAetherDividePassiveSkillCsReq, func() any { return new(proto.ClearAetherDividePassiveSkillCsReq) })
	c.regMsg(ChessRogueEnterNextLayerScRsp, func() any { return new(proto.ChessRogueEnterNextLayerScRsp) })
	c.regMsg(AcceptActivityExpeditionCsReq, func() any { return new(proto.AcceptActivityExpeditionCsReq) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
	c.regMsg(SpaceZooDataCsReq, func() any { return new(proto.SpaceZooDataCsReq) })
	c.regMsg(MuseumRandomEventSelectCsReq, func() any { return new(proto.MuseumRandomEventSelectCsReq) })
	c.regMsg(MonopolyEventSelectFriendCsReq, func() any { return new(proto.MonopolyEventSelectFriendCsReq) })
	c.regMsg(SpaceZooDeleteCatScRsp, func() any { return new(proto.SpaceZooDeleteCatScRsp) })
	c.regMsg(FinishPlotCsReq, func() any { return new(proto.FinishPlotCsReq) })
	c.regMsg(LockEquipmentScRsp, func() any { return new(proto.LockEquipmentScRsp) })
	c.regMsg(TextJoinSaveCsReq, func() any { return new(proto.TextJoinSaveCsReq) })
	c.regMsg(SceneCastSkillScRsp, func() any { return new(proto.SceneCastSkillScRsp) })
	c.regMsg(MultipleDropInfoNotify, func() any { return new(proto.MultipleDropInfoNotify) })
	c.regMsg(SceneGroupRefreshScNotify, func() any { return new(proto.SceneGroupRefreshScNotify) })
	c.regMsg(MuseumRandomEventStartScNotify, func() any { return new(proto.MuseumRandomEventStartScNotify) })
	c.regMsg(SetSignatureCsReq, func() any { return new(proto.SetSignatureCsReq) })
	c.regMsg(SetTurnFoodSwitchScRsp, func() any { return new(proto.SetTurnFoodSwitchScRsp) })
	c.regMsg(AetherDivideTakeChallengeRewardScRsp, func() any { return new(proto.AetherDivideTakeChallengeRewardScRsp) })
	c.regMsg(TelevisionActivityBattleEndScNotify, func() any { return new(proto.TelevisionActivityBattleEndScNotify) })
	c.regMsg(SetGenderCsReq, func() any { return new(proto.SetGenderCsReq) })
	c.regMsg(ResetMapRotationRegionScRsp, func() any { return new(proto.ResetMapRotationRegionScRsp) })
	c.regMsg(GetRogueScoreRewardInfoCsReq, func() any { return new(proto.GetRogueScoreRewardInfoCsReq) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(FinishChessRogueNousSubStoryCsReq, func() any { return new(proto.FinishChessRogueNousSubStoryCsReq) })
	c.regMsg(UpdateTrackMainMissionIdCsReq, func() any { return new(proto.UpdateTrackMainMissionIdCsReq) })
	c.regMsg(ExtraLineupDestroyNotify, func() any { return new(proto.ExtraLineupDestroyNotify) })
	c.regMsg(TakeOfferingRewardCsReq, func() any { return new(proto.TakeOfferingRewardCsReq) })
	c.regMsg(TravelBrochureGetPasterScNotify, func() any { return new(proto.TravelBrochureGetPasterScNotify) })
	c.regMsg(GetMonopolyInfoScRsp, func() any { return new(proto.GetMonopolyInfoScRsp) })
	c.regMsg(TreasureDungeonFinishScNotify, func() any { return new(proto.TreasureDungeonFinishScNotify) })
	c.regMsg(SetCurInteractEntityCsReq, func() any { return new(proto.SetCurInteractEntityCsReq) })
	c.regMsg(SetGroupCustomSaveDataScRsp, func() any { return new(proto.SetGroupCustomSaveDataScRsp) })
	c.regMsg(GetRogueShopMiracleInfoScRsp, func() any { return new(proto.GetRogueShopMiracleInfoScRsp) })
	c.regMsg(GetPlayerDetailInfoCsReq, func() any { return new(proto.GetPlayerDetailInfoCsReq) })
	c.regMsg(GetPunkLordDataScRsp, func() any { return new(proto.GetPunkLordDataScRsp) })
	c.regMsg(GetMissionDataScRsp, func() any { return new(proto.GetMissionDataScRsp) })
	c.regMsg(SyncRogueAeonLevelUpRewardScNotify, func() any { return new(proto.SyncRogueAeonLevelUpRewardScNotify) })
	c.regMsg(GetSecretKeyInfoCsReq, func() any { return new(proto.GetSecretKeyInfoCsReq) })
	c.regMsg(GetPunkLordMonsterDataScRsp, func() any { return new(proto.GetPunkLordMonsterDataScRsp) })
	c.regMsg(DeployRotaterScRsp, func() any { return new(proto.DeployRotaterScRsp) })
	c.regMsg(SetForbidOtherApplyFriendCsReq, func() any { return new(proto.SetForbidOtherApplyFriendCsReq) })
	c.regMsg(SetGameplayBirthdayScRsp, func() any { return new(proto.SetGameplayBirthdayScRsp) })
	c.regMsg(SetLineupNameScRsp, func() any { return new(proto.SetLineupNameScRsp) })
	c.regMsg(GetArchiveDataCsReq, func() any { return new(proto.GetArchiveDataCsReq) })
	c.regMsg(GetTutorialGuideCsReq, func() any { return new(proto.GetTutorialGuideCsReq) })
	c.regMsg(SpaceZooCatUpdateNotify, func() any { return new(proto.SpaceZooCatUpdateNotify) })
	c.regMsg(UpdatePlayerSettingCsReq, func() any { return new(proto.UpdatePlayerSettingCsReq) })
	c.regMsg(QuitRogueCsReq, func() any { return new(proto.QuitRogueCsReq) })
	c.regMsg(GetTelevisionActivityDataCsReq, func() any { return new(proto.GetTelevisionActivityDataCsReq) })
	c.regMsg(ChessRoguePickAvatarCsReq, func() any { return new(proto.ChessRoguePickAvatarCsReq) })
	c.regMsg(ReEnterLastElementStageScRsp, func() any { return new(proto.ReEnterLastElementStageScRsp) })
	c.regMsg(GetServerPrefsDataCsReq, func() any { return new(proto.GetServerPrefsDataCsReq) })
	c.regMsg(ClientObjUploadScRsp, func() any { return new(proto.ClientObjUploadScRsp) })
	c.regMsg(HeliobusSelectSkillScRsp, func() any { return new(proto.HeliobusSelectSkillScRsp) })
	c.regMsg(FinishFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.FinishFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(DoGachaScRsp, func() any { return new(proto.DoGachaScRsp) })
	c.regMsg(TravelBrochureSetCustomValueCsReq, func() any { return new(proto.TravelBrochureSetCustomValueCsReq) })
	c.regMsg(TrialBackGroundMusicScRsp, func() any { return new(proto.TrialBackGroundMusicScRsp) })
	c.regMsg(ChooseBoxingClubResonanceCsReq, func() any { return new(proto.ChooseBoxingClubResonanceCsReq) })
	c.regMsg(AlleyFundsScNotify, func() any { return new(proto.AlleyFundsScNotify) })
	c.regMsg(TakeKilledPunkLordMonsterScoreCsReq, func() any { return new(proto.TakeKilledPunkLordMonsterScoreCsReq) })
	c.regMsg(AddAvatarScNotify, func() any { return new(proto.AddAvatarScNotify) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(MonopolyDailySettleScNotify, func() any { return new(proto.MonopolyDailySettleScNotify) })
	c.regMsg(MakeMissionDrinkCsReq, func() any { return new(proto.MakeMissionDrinkCsReq) })
	c.regMsg(GetGachaCeilingScRsp, func() any { return new(proto.GetGachaCeilingScRsp) })
	c.regMsg(MissionEventRewardScNotify, func() any { return new(proto.MissionEventRewardScNotify) })
	c.regMsg(MonopolyLikeScRsp, func() any { return new(proto.MonopolyLikeScRsp) })
	c.regMsg(GetRndOptionCsReq, func() any { return new(proto.GetRndOptionCsReq) })
	c.regMsg(PlayerReturnInfoQueryScRsp, func() any { return new(proto.PlayerReturnInfoQueryScRsp) })
	c.regMsg(PromoteEquipmentCsReq, func() any { return new(proto.PromoteEquipmentCsReq) })
	c.regMsg(ExpUpEquipmentScRsp, func() any { return new(proto.ExpUpEquipmentScRsp) })
	c.regMsg(UseTreasureDungeonItemScRsp, func() any { return new(proto.UseTreasureDungeonItemScRsp) })
	c.regMsg(SecurityReportCsReq, func() any { return new(proto.SecurityReportCsReq) })
	c.regMsg(GetMissionEventDataCsReq, func() any { return new(proto.GetMissionEventDataCsReq) })
	c.regMsg(QuitBattleScRsp, func() any { return new(proto.QuitBattleScRsp) })
	c.regMsg(InteractPropScRsp, func() any { return new(proto.InteractPropScRsp) })
	c.regMsg(GetMarkItemListScRsp, func() any { return new(proto.GetMarkItemListScRsp) })
	c.regMsg(GetLevelRewardCsReq, func() any { return new(proto.GetLevelRewardCsReq) })
	c.regMsg(AetherDivideTainerInfoScNotify, func() any { return new(proto.AetherDivideTainerInfoScNotify) })
	c.regMsg(OpenTreasureDungeonGridScRsp, func() any { return new(proto.OpenTreasureDungeonGridScRsp) })
	c.regMsg(GetPhoneDataCsReq, func() any { return new(proto.GetPhoneDataCsReq) })
	c.regMsg(GetRogueAdventureRoomInfoScRsp, func() any { return new(proto.GetRogueAdventureRoomInfoScRsp) })
	c.regMsg(MonopolyTakeRaffleTicketRewardCsReq, func() any { return new(proto.MonopolyTakeRaffleTicketRewardCsReq) })
	c.regMsg(GetMapRotationDataScRsp, func() any { return new(proto.GetMapRotationDataScRsp) })
	c.regMsg(SetTurnFoodSwitchCsReq, func() any { return new(proto.SetTurnFoodSwitchCsReq) })
	c.regMsg(EnterTreasureDungeonScRsp, func() any { return new(proto.EnterTreasureDungeonScRsp) })
	c.regMsg(PlayerReturnPointChangeScNotify, func() any { return new(proto.PlayerReturnPointChangeScNotify) })
	c.regMsg(TakeExpeditionRewardScRsp, func() any { return new(proto.TakeExpeditionRewardScRsp) })
	c.regMsg(SetClientPausedScRsp, func() any { return new(proto.SetClientPausedScRsp) })
	c.regMsg(ShowNewSupplementVisitorCsReq, func() any { return new(proto.ShowNewSupplementVisitorCsReq) })
	c.regMsg(ShareCsReq, func() any { return new(proto.ShareCsReq) })
	c.regMsg(TravelBrochureApplyPasterListScRsp, func() any { return new(proto.TravelBrochureApplyPasterListScRsp) })
	c.regMsg(LeaveRaidScRsp, func() any { return new(proto.LeaveRaidScRsp) })
	c.regMsg(MissionRewardScNotify, func() any { return new(proto.MissionRewardScNotify) })
	c.regMsg(ChessRogueSkipTeachingLevelScRsp, func() any { return new(proto.ChessRogueSkipTeachingLevelScRsp) })
	c.regMsg(HeliobusSnsLikeScRsp, func() any { return new(proto.HeliobusSnsLikeScRsp) })
	c.regMsg(GetRndOptionScRsp, func() any { return new(proto.GetRndOptionScRsp) })
	c.regMsg(GetAvatarDataCsReq, func() any { return new(proto.GetAvatarDataCsReq) })
	c.regMsg(ComposeItemCsReq, func() any { return new(proto.ComposeItemCsReq) })
	c.regMsg(SpringRecoverCsReq, func() any { return new(proto.SpringRecoverCsReq) })
	c.regMsg(GetCurLineupDataScRsp, func() any { return new(proto.GetCurLineupDataScRsp) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(FinishRogueDialogueGroupScRsp, func() any { return new(proto.FinishRogueDialogueGroupScRsp) })
	c.regMsg(GeneralVirtualItemDataNotify, func() any { return new(proto.GeneralVirtualItemDataNotify) })
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(DestroyItemScRsp, func() any { return new(proto.DestroyItemScRsp) })
	c.regMsg(NewMailScNotify, func() any { return new(proto.NewMailScNotify) })
	c.regMsg(EnterSectionCsReq, func() any { return new(proto.EnterSectionCsReq) })
	c.regMsg(BatchGetQuestDataCsReq, func() any { return new(proto.BatchGetQuestDataCsReq) })
	c.regMsg(WolfBroGamePickupBulletCsReq, func() any { return new(proto.WolfBroGamePickupBulletCsReq) })
	c.regMsg(SyncTurnFoodNotify, func() any { return new(proto.SyncTurnFoodNotify) })
	c.regMsg(StartTrialActivityScRsp, func() any { return new(proto.StartTrialActivityScRsp) })
	c.regMsg(TriggerVoiceCsReq, func() any { return new(proto.TriggerVoiceCsReq) })
	c.regMsg(PromoteAvatarCsReq, func() any { return new(proto.PromoteAvatarCsReq) })
	c.regMsg(GetFirstTalkNpcCsReq, func() any { return new(proto.GetFirstTalkNpcCsReq) })
	c.regMsg(SetBoxingClubResonanceLineupScRsp, func() any { return new(proto.SetBoxingClubResonanceLineupScRsp) })
	c.regMsg(ChessRogueQueryBpScRsp, func() any { return new(proto.ChessRogueQueryBpScRsp) })
	c.regMsg(RogueModifierDelNotify, func() any { return new(proto.RogueModifierDelNotify) })
	c.regMsg(QuitLineupCsReq, func() any { return new(proto.QuitLineupCsReq) })
	c.regMsg(OpenTreasureDungeonGridCsReq, func() any { return new(proto.OpenTreasureDungeonGridCsReq) })
	c.regMsg(AcceptExpeditionCsReq, func() any { return new(proto.AcceptExpeditionCsReq) })
	c.regMsg(GetExpeditionDataScRsp, func() any { return new(proto.GetExpeditionDataScRsp) })
	c.regMsg(SetCurWaypointScRsp, func() any { return new(proto.SetCurWaypointScRsp) })
	c.regMsg(SetDisplayAvatarCsReq, func() any { return new(proto.SetDisplayAvatarCsReq) })
	c.regMsg(GetRecyleTimeScRsp, func() any { return new(proto.GetRecyleTimeScRsp) })
	c.regMsg(EnhanceRogueBuffScRsp, func() any { return new(proto.EnhanceRogueBuffScRsp) })
	c.regMsg(UpdateServerPrefsDataScRsp, func() any { return new(proto.UpdateServerPrefsDataScRsp) })
	c.regMsg(GetTutorialCsReq, func() any { return new(proto.GetTutorialCsReq) })
	c.regMsg(TravelBrochureGetDataCsReq, func() any { return new(proto.TravelBrochureGetDataCsReq) })
	c.regMsg(AcceptMainMissionCsReq, func() any { return new(proto.AcceptMainMissionCsReq) })
	c.regMsg(DailyRefreshNotify, func() any { return new(proto.DailyRefreshNotify) })
	c.regMsg(TakeChapterRewardCsReq, func() any { return new(proto.TakeChapterRewardCsReq) })
	c.regMsg(TravelBrochureRemovePasterScRsp, func() any { return new(proto.TravelBrochureRemovePasterScRsp) })
	c.regMsg(GetAllServerPrefsDataCsReq, func() any { return new(proto.GetAllServerPrefsDataCsReq) })
	c.regMsg(GetSaveLogisticsMapCsReq, func() any { return new(proto.GetSaveLogisticsMapCsReq) })
	c.regMsg(GetVideoVersionKeyCsReq, func() any { return new(proto.GetVideoVersionKeyCsReq) })
	c.regMsg(GetChatEmojiListScRsp, func() any { return new(proto.GetChatEmojiListScRsp) })
	c.regMsg(TrialActivityDataChangeScNotify, func() any { return new(proto.TrialActivityDataChangeScNotify) })
	c.regMsg(MonopolyGetRafflePoolInfoScRsp, func() any { return new(proto.MonopolyGetRafflePoolInfoScRsp) })
	c.regMsg(DelMailCsReq, func() any { return new(proto.DelMailCsReq) })
	c.regMsg(MuseumTargetMissionFinishNotify, func() any { return new(proto.MuseumTargetMissionFinishNotify) })
	c.regMsg(GetAssistListCsReq, func() any { return new(proto.GetAssistListCsReq) })
	c.regMsg(FinishItemIdCsReq, func() any { return new(proto.FinishItemIdCsReq) })
	c.regMsg(GetTrialActivityDataScRsp, func() any { return new(proto.GetTrialActivityDataScRsp) })
	c.regMsg(SelectChessRogueSubStoryScRsp, func() any { return new(proto.SelectChessRogueSubStoryScRsp) })
	c.regMsg(PlayerGetTokenScRsp, func() any { return new(proto.PlayerGetTokenScRsp) })
	c.regMsg(QuitTreasureDungeonScRsp, func() any { return new(proto.QuitTreasureDungeonScRsp) })
	c.regMsg(UpgradeAreaStatScRsp, func() any { return new(proto.UpgradeAreaStatScRsp) })
	c.regMsg(SendMsgCsReq, func() any { return new(proto.SendMsgCsReq) })
	c.regMsg(PlayBackGroundMusicCsReq, func() any { return new(proto.PlayBackGroundMusicCsReq) })
	c.regMsg(GetAlleyInfoScRsp, func() any { return new(proto.GetAlleyInfoScRsp) })
	c.regMsg(StartPunkLordRaidCsReq, func() any { return new(proto.StartPunkLordRaidCsReq) })
	c.regMsg(GateServerScNotify, func() any { return new(proto.GateServerScNotify) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetLevelRewardTakenListCsReq, func() any { return new(proto.GetLevelRewardTakenListCsReq) })
	c.regMsg(JoinLineupScRsp, func() any { return new(proto.JoinLineupScRsp) })
	c.regMsg(SetHeroBasicTypeCsReq, func() any { return new(proto.SetHeroBasicTypeCsReq) })
	c.regMsg(StartAetherDivideChallengeBattleScRsp, func() any { return new(proto.StartAetherDivideChallengeBattleScRsp) })
	c.regMsg(TakeQuestOptionalRewardScRsp, func() any { return new(proto.TakeQuestOptionalRewardScRsp) })
	c.regMsg(SceneCastSkillCostMpCsReq, func() any { return new(proto.SceneCastSkillCostMpCsReq) })
	c.regMsg(ChessRogueGoAheadCsReq, func() any { return new(proto.ChessRogueGoAheadCsReq) })
	c.regMsg(GetJukeboxDataCsReq, func() any { return new(proto.GetJukeboxDataCsReq) })
	c.regMsg(DoGachaInRollShopCsReq, func() any { return new(proto.DoGachaInRollShopCsReq) })
	c.regMsg(MonopolyConditionUpdateScNotify, func() any { return new(proto.MonopolyConditionUpdateScNotify) })
	c.regMsg(ChessRogueUpdateDiceInfoScNotify, func() any { return new(proto.ChessRogueUpdateDiceInfoScNotify) })
	c.regMsg(EnableRogueTalentCsReq, func() any { return new(proto.EnableRogueTalentCsReq) })
	c.regMsg(ExchangeHcoinCsReq, func() any { return new(proto.ExchangeHcoinCsReq) })
	c.regMsg(UpdateRedDotDataScRsp, func() any { return new(proto.UpdateRedDotDataScRsp) })
	c.regMsg(SetDisplayAvatarScRsp, func() any { return new(proto.SetDisplayAvatarScRsp) })
	c.regMsg(FinishFirstTalkNpcCsReq, func() any { return new(proto.FinishFirstTalkNpcCsReq) })
	c.regMsg(ExpeditionDataChangeScNotify, func() any { return new(proto.ExpeditionDataChangeScNotify) })
	c.regMsg(ChessRogueSelectBpScRsp, func() any { return new(proto.ChessRogueSelectBpScRsp) })
	c.regMsg(DeployRotaterCsReq, func() any { return new(proto.DeployRotaterCsReq) })
	c.regMsg(GetBasicInfoCsReq, func() any { return new(proto.GetBasicInfoCsReq) })
	c.regMsg(TakeChapterRewardScRsp, func() any { return new(proto.TakeChapterRewardScRsp) })
	c.regMsg(PlayerReturnTakeRewardScRsp, func() any { return new(proto.PlayerReturnTakeRewardScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoScRsp, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoScRsp) })
	c.regMsg(FinishFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.FinishFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(TakeRogueMiracleHandbookRewardCsReq, func() any { return new(proto.TakeRogueMiracleHandbookRewardCsReq) })
	c.regMsg(GetPunkLordDataCsReq, func() any { return new(proto.GetPunkLordDataCsReq) })
	c.regMsg(ChooseBoxingClubStageOptionalBuffCsReq, func() any { return new(proto.ChooseBoxingClubStageOptionalBuffCsReq) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(RankUpEquipmentCsReq, func() any { return new(proto.RankUpEquipmentCsReq) })
	c.regMsg(ChessRogueMoveCellNotify, func() any { return new(proto.ChessRogueMoveCellNotify) })
	c.regMsg(GetQuestRecordScRsp, func() any { return new(proto.GetQuestRecordScRsp) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(SyncChessRogueNousValueScNotify, func() any { return new(proto.SyncChessRogueNousValueScNotify) })
	c.regMsg(HeliobusUpgradeLevelScRsp, func() any { return new(proto.HeliobusUpgradeLevelScRsp) })
	c.regMsg(PlayerLoginCsReq, func() any { return new(proto.PlayerLoginCsReq) })
	c.regMsg(TravelBrochurePageResetScRsp, func() any { return new(proto.TravelBrochurePageResetScRsp) })
	c.regMsg(BuyBpLevelScRsp, func() any { return new(proto.BuyBpLevelScRsp) })
	c.regMsg(SharePunkLordMonsterCsReq, func() any { return new(proto.SharePunkLordMonsterCsReq) })
	c.regMsg(GetServerPrefsDataScRsp, func() any { return new(proto.GetServerPrefsDataScRsp) })
	c.regMsg(FightTreasureDungeonMonsterScRsp, func() any { return new(proto.FightTreasureDungeonMonsterScRsp) })
	c.regMsg(ExchangeRogueBuffWithMiracleCsReq, func() any { return new(proto.ExchangeRogueBuffWithMiracleCsReq) })
	c.regMsg(MonopolyScrachRaffleTicketCsReq, func() any { return new(proto.MonopolyScrachRaffleTicketCsReq) })
	c.regMsg(SetMissionEventProgressCsReq, func() any { return new(proto.SetMissionEventProgressCsReq) })
	c.regMsg(TakeMonsterResearchActivityRewardScRsp, func() any { return new(proto.TakeMonsterResearchActivityRewardScRsp) })
	c.regMsg(FeatureSwitchClosedScNotify, func() any { return new(proto.FeatureSwitchClosedScNotify) })
	c.regMsg(SetGenderScRsp, func() any { return new(proto.SetGenderScRsp) })
	c.regMsg(BattleCollegeDataChangeScNotify, func() any { return new(proto.BattleCollegeDataChangeScNotify) })
	c.regMsg(PunkLordRaidTimeOutScNotify, func() any { return new(proto.PunkLordRaidTimeOutScNotify) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
	c.regMsg(GetMultipleDropInfoCsReq, func() any { return new(proto.GetMultipleDropInfoCsReq) })
	c.regMsg(GetSaveRaidCsReq, func() any { return new(proto.GetSaveRaidCsReq) })
	c.regMsg(GetMapRotationDataCsReq, func() any { return new(proto.GetMapRotationDataCsReq) })
	c.regMsg(GetRogueInitialScoreCsReq, func() any { return new(proto.GetRogueInitialScoreCsReq) })
	c.regMsg(SetAssistScRsp, func() any { return new(proto.SetAssistScRsp) })
	c.regMsg(CurTrialActivityScNotify, func() any { return new(proto.CurTrialActivityScNotify) })
	c.regMsg(GetStuffScNotify, func() any { return new(proto.GetStuffScNotify) })
	c.regMsg(GetLoginChatInfoScRsp, func() any { return new(proto.GetLoginChatInfoScRsp) })
	c.regMsg(GameplayCounterUpdateScNotify, func() any { return new(proto.GameplayCounterUpdateScNotify) })
	c.regMsg(ChallengeLineupNotify, func() any { return new(proto.ChallengeLineupNotify) })
	c.regMsg(SwapLineupCsReq, func() any { return new(proto.SwapLineupCsReq) })
	c.regMsg(PromoteEquipmentScRsp, func() any { return new(proto.PromoteEquipmentScRsp) })
	c.regMsg(ReviveRogueAvatarScRsp, func() any { return new(proto.ReviveRogueAvatarScRsp) })
	c.regMsg(FinishSectionIdCsReq, func() any { return new(proto.FinishSectionIdCsReq) })
	c.regMsg(HeliobusActivityDataCsReq, func() any { return new(proto.HeliobusActivityDataCsReq) })
	c.regMsg(GetNpcMessageGroupCsReq, func() any { return new(proto.GetNpcMessageGroupCsReq) })
	c.regMsg(DelSaveRaidScNotify, func() any { return new(proto.DelSaveRaidScNotify) })
	c.regMsg(ReturnLastTownScRsp, func() any { return new(proto.ReturnLastTownScRsp) })
	c.regMsg(GetAllLineupDataCsReq, func() any { return new(proto.GetAllLineupDataCsReq) })
	c.regMsg(AetherDivideSpiritExpUpScRsp, func() any { return new(proto.AetherDivideSpiritExpUpScRsp) })
	c.regMsg(TakeTrainVisitorUntakenBehaviorRewardCsReq, func() any { return new(proto.TakeTrainVisitorUntakenBehaviorRewardCsReq) })
	c.regMsg(TakePictureCsReq, func() any { return new(proto.TakePictureCsReq) })
	c.regMsg(GetBasicInfoScRsp, func() any { return new(proto.GetBasicInfoScRsp) })
	c.regMsg(MonopolyGameRaiseRatioScRsp, func() any { return new(proto.MonopolyGameRaiseRatioScRsp) })
	c.regMsg(StartAetherDivideChallengeBattleCsReq, func() any { return new(proto.StartAetherDivideChallengeBattleCsReq) })
	c.regMsg(TakeOffAvatarSkinCsReq, func() any { return new(proto.TakeOffAvatarSkinCsReq) })
	c.regMsg(MonopolyClickCellScRsp, func() any { return new(proto.MonopolyClickCellScRsp) })
	c.regMsg(SaveLogisticsScRsp, func() any { return new(proto.SaveLogisticsScRsp) })
	c.regMsg(CityShopInfoScNotify, func() any { return new(proto.CityShopInfoScNotify) })
	c.regMsg(ResetMapRotationRegionCsReq, func() any { return new(proto.ResetMapRotationRegionCsReq) })
	c.regMsg(ChessRogueQuestFinishNotify, func() any { return new(proto.ChessRogueQuestFinishNotify) })
	c.regMsg(FinishTutorialGuideScRsp, func() any { return new(proto.FinishTutorialGuideScRsp) })
	c.regMsg(TakeQuestRewardScRsp, func() any { return new(proto.TakeQuestRewardScRsp) })
	c.regMsg(BatchMarkChatEmojiCsReq, func() any { return new(proto.BatchMarkChatEmojiCsReq) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(ChessRogueQueryBpCsReq, func() any { return new(proto.ChessRogueQueryBpCsReq) })
	c.regMsg(AcceptMissionEventCsReq, func() any { return new(proto.AcceptMissionEventCsReq) })
	c.regMsg(ReportPlayerCsReq, func() any { return new(proto.ReportPlayerCsReq) })
	c.regMsg(PunkLordDataChangeNotify, func() any { return new(proto.PunkLordDataChangeNotify) })
	c.regMsg(SearchPlayerCsReq, func() any { return new(proto.SearchPlayerCsReq) })
	c.regMsg(SetStuffToAreaScRsp, func() any { return new(proto.SetStuffToAreaScRsp) })
	c.regMsg(ComposeSelectedRelicScRsp, func() any { return new(proto.ComposeSelectedRelicScRsp) })
	c.regMsg(GiveUpBoxingClubChallengeCsReq, func() any { return new(proto.GiveUpBoxingClubChallengeCsReq) })
	c.regMsg(InterruptMissionEventCsReq, func() any { return new(proto.InterruptMissionEventCsReq) })
	c.regMsg(FinishEmotionDialoguePerformanceCsReq, func() any { return new(proto.FinishEmotionDialoguePerformanceCsReq) })
	c.regMsg(StartAlleyEventScRsp, func() any { return new(proto.StartAlleyEventScRsp) })
	c.regMsg(PunkLordBattleResultScNotify, func() any { return new(proto.PunkLordBattleResultScNotify) })
	c.regMsg(GetChessRogueNousStoryInfoScRsp, func() any { return new(proto.GetChessRogueNousStoryInfoScRsp) })
	c.regMsg(SpaceZooBornCsReq, func() any { return new(proto.SpaceZooBornCsReq) })
	c.regMsg(LeaveMapRotationRegionCsReq, func() any { return new(proto.LeaveMapRotationRegionCsReq) })
	c.regMsg(EnterSceneCsReq, func() any { return new(proto.EnterSceneCsReq) })
	c.regMsg(HeartDialTraceScriptScRsp, func() any { return new(proto.HeartDialTraceScriptScRsp) })
	c.regMsg(GetMonsterResearchActivityDataCsReq, func() any { return new(proto.GetMonsterResearchActivityDataCsReq) })
	c.regMsg(GetSingleRedDotParamGroupScRsp, func() any { return new(proto.GetSingleRedDotParamGroupScRsp) })
	c.regMsg(DeleteBlacklistCsReq, func() any { return new(proto.DeleteBlacklistCsReq) })
	c.regMsg(ExpUpRelicCsReq, func() any { return new(proto.ExpUpRelicCsReq) })
	c.regMsg(UnlockTutorialGuideScRsp, func() any { return new(proto.UnlockTutorialGuideScRsp) })
	c.regMsg(GetEnteredSceneCsReq, func() any { return new(proto.GetEnteredSceneCsReq) })
	c.regMsg(RaidCollectionDataScRsp, func() any { return new(proto.RaidCollectionDataScRsp) })
	c.regMsg(SyncServerSceneChangeNotify, func() any { return new(proto.SyncServerSceneChangeNotify) })
	c.regMsg(StartPunkLordRaidScRsp, func() any { return new(proto.StartPunkLordRaidScRsp) })
	c.regMsg(DailyFirstMeetPamScRsp, func() any { return new(proto.DailyFirstMeetPamScRsp) })
	c.regMsg(EnterRogueMapRoomCsReq, func() any { return new(proto.EnterRogueMapRoomCsReq) })
	c.regMsg(EnterRogueMapRoomScRsp, func() any { return new(proto.EnterRogueMapRoomScRsp) })
	c.regMsg(EnterFeverTimeActivityStageScRsp, func() any { return new(proto.EnterFeverTimeActivityStageScRsp) })
	c.regMsg(SyncTaskScRsp, func() any { return new(proto.SyncTaskScRsp) })
	c.regMsg(ChessRogueLeaveCsReq, func() any { return new(proto.ChessRogueLeaveCsReq) })
	c.regMsg(PlayerReturnTakePointRewardCsReq, func() any { return new(proto.PlayerReturnTakePointRewardCsReq) })
	c.regMsg(GetSpringRecoverDataCsReq, func() any { return new(proto.GetSpringRecoverDataCsReq) })
	c.regMsg(SyncTaskCsReq, func() any { return new(proto.SyncTaskCsReq) })
	c.regMsg(TextJoinBatchSaveCsReq, func() any { return new(proto.TextJoinBatchSaveCsReq) })
	c.regMsg(SubmitOfferingItemCsReq, func() any { return new(proto.SubmitOfferingItemCsReq) })
	c.regMsg(GetBagCsReq, func() any { return new(proto.GetBagCsReq) })
	c.regMsg(GetStrongChallengeActivityDataScRsp, func() any { return new(proto.GetStrongChallengeActivityDataScRsp) })
	c.regMsg(GetNpcStatusCsReq, func() any { return new(proto.GetNpcStatusCsReq) })
	c.regMsg(AlleyShipmentEventEffectsScNotify, func() any { return new(proto.AlleyShipmentEventEffectsScNotify) })
	c.regMsg(FinishRogueDialogueGroupCsReq, func() any { return new(proto.FinishRogueDialogueGroupCsReq) })
	c.regMsg(AcceptMissionEventScRsp, func() any { return new(proto.AcceptMissionEventScRsp) })
	c.regMsg(GetDailyActiveInfoCsReq, func() any { return new(proto.GetDailyActiveInfoCsReq) })
	c.regMsg(PlayerKickOutScNotify, func() any { return new(proto.PlayerKickOutScNotify) })
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(PlayerReturnForceFinishScNotify, func() any { return new(proto.PlayerReturnForceFinishScNotify) })
	c.regMsg(GetMonopolyMbtiReportRewardCsReq, func() any { return new(proto.GetMonopolyMbtiReportRewardCsReq) })
	c.regMsg(TakeChallengeRaidRewardCsReq, func() any { return new(proto.TakeChallengeRaidRewardCsReq) })
	c.regMsg(MonopolyGetRafflePoolInfoCsReq, func() any { return new(proto.MonopolyGetRafflePoolInfoCsReq) })
	c.regMsg(EntityBindPropScRsp, func() any { return new(proto.EntityBindPropScRsp) })
	c.regMsg(GetMissionDataCsReq, func() any { return new(proto.GetMissionDataCsReq) })
	c.regMsg(DressRelicAvatarScRsp, func() any { return new(proto.DressRelicAvatarScRsp) })
	c.regMsg(SelectRogueDialogueEventScRsp, func() any { return new(proto.SelectRogueDialogueEventScRsp) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(SelectChessRogueSubStoryCsReq, func() any { return new(proto.SelectChessRogueSubStoryCsReq) })
	c.regMsg(TakeAssistRewardScRsp, func() any { return new(proto.TakeAssistRewardScRsp) })
	c.regMsg(EnhanceChessRogueBuffCsReq, func() any { return new(proto.EnhanceChessRogueBuffCsReq) })
	c.regMsg(RankUpEquipmentScRsp, func() any { return new(proto.RankUpEquipmentScRsp) })
	c.regMsg(WolfBroGameActivateBulletCsReq, func() any { return new(proto.WolfBroGameActivateBulletCsReq) })
	c.regMsg(TrainVisitorBehaviorFinishScRsp, func() any { return new(proto.TrainVisitorBehaviorFinishScRsp) })
	c.regMsg(EnterSectionScRsp, func() any { return new(proto.EnterSectionScRsp) })
	c.regMsg(GetRogueShopBuffInfoScRsp, func() any { return new(proto.GetRogueShopBuffInfoScRsp) })
	c.regMsg(ChessRogueUpdateAllowedSelectCellScNotify, func() any { return new(proto.ChessRogueUpdateAllowedSelectCellScNotify) })
	c.regMsg(SyncClientResVersionScRsp, func() any { return new(proto.SyncClientResVersionScRsp) })
	c.regMsg(GetFantasticStoryActivityDataCsReq, func() any { return new(proto.GetFantasticStoryActivityDataCsReq) })
	c.regMsg(GetHeroBasicTypeInfoScRsp, func() any { return new(proto.GetHeroBasicTypeInfoScRsp) })
	c.regMsg(GetChessRogueStoryAeonTalkInfoCsReq, func() any { return new(proto.GetChessRogueStoryAeonTalkInfoCsReq) })
	c.regMsg(TakeAllApRewardCsReq, func() any { return new(proto.TakeAllApRewardCsReq) })
	c.regMsg(GetWolfBroGameDataCsReq, func() any { return new(proto.GetWolfBroGameDataCsReq) })
	c.regMsg(TakeOffEquipmentScRsp, func() any { return new(proto.TakeOffEquipmentScRsp) })
	c.regMsg(GetRogueAdventureRoomInfoCsReq, func() any { return new(proto.GetRogueAdventureRoomInfoCsReq) })
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
