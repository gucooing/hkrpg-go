package cmd

import (
	"reflect"
	"sync"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

var sharedCmdProtoMap *CmdProtoMap
var cmdProtoMapOnce sync.Once

type CmdProtoMap struct {
	cmdIdProtoObjMap        map[uint16]reflect.Type
	protoObjCmdIdMap        map[reflect.Type]uint16
	cmdDeDupMap             map[uint16]bool
	cmdIdCmdNameMap         map[uint16]string
	cmdNameCmdIdMap         map[string]uint16
	cmdIdProtoObjCacheMap   map[uint16]*sync.Pool
	cmdIdProtoObjFastNewMap map[uint16]func() any
}

func GetSharedCmdProtoMap() *CmdProtoMap {
	cmdProtoMapOnce.Do(func() {
		sharedCmdProtoMap = NewCmdProtoMap()
	})
	return sharedCmdProtoMap
}

func NewCmdProtoMap() (r *CmdProtoMap) {
	r = new(CmdProtoMap)
	r.cmdIdProtoObjMap = make(map[uint16]reflect.Type)
	r.protoObjCmdIdMap = make(map[reflect.Type]uint16)
	r.cmdDeDupMap = make(map[uint16]bool)
	r.cmdIdCmdNameMap = make(map[uint16]string)
	r.cmdNameCmdIdMap = make(map[string]uint16)
	r.cmdIdProtoObjCacheMap = make(map[uint16]*sync.Pool)
	r.cmdIdProtoObjFastNewMap = make(map[uint16]func() any)
	r.registerAllMessage()

	return r
}

func (c *CmdProtoMap) registerMessage() {
	c.regMsg(ActivateFarmElementCsReq, func() any { return new(proto.ActivateFarmElementCsReq) })
	c.regMsg(ActivateFarmElementScRsp, func() any { return new(proto.ActivateFarmElementScRsp) })
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
	c.regMsg(GetFarmStageGachaInfoCsReq, func() any { return new(proto.GetFarmStageGachaInfoCsReq) })
	c.regMsg(GetFarmStageGachaInfoScRsp, func() any { return new(proto.GetFarmStageGachaInfoScRsp) })
	c.regMsg(GetFirstTalkByPerformanceNpcCsReq, func() any { return new(proto.GetFirstTalkByPerformanceNpcCsReq) })
	c.regMsg(GetFirstTalkByPerformanceNpcScRsp, func() any { return new(proto.GetFirstTalkByPerformanceNpcScRsp) })
	c.regMsg(GetFriendApplyListInfoScRsp, func() any { return new(proto.GetFriendApplyListInfoScRsp) })
	c.regMsg(GetFriendListInfoScRsp, func() any { return new(proto.GetFriendListInfoScRsp) })
	c.regMsg(GetFriendLoginInfoScRsp, func() any { return new(proto.GetFriendLoginInfoScRsp) })
	c.regMsg(GetFriendRecommendListInfoScRsp, func() any { return new(proto.GetFriendRecommendListInfoScRsp) })
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
	c.regMsg(GetQuestDataCsReq, func() any { return new(proto.GetQuestDataCsReq) })
	c.regMsg(GetQuestDataScRsp, func() any { return new(proto.GetQuestDataScRsp) })
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
	c.regMsg(GroupStateChangeCsReq, func() any { return new(proto.GroupStateChangeCsReq) })
	c.regMsg(GroupStateChangeScNotify, func() any { return new(proto.GroupStateChangeScNotify) })
	c.regMsg(GroupStateChangeScRsp, func() any { return new(proto.GroupStateChangeScRsp) })
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
	c.regMsg(RefreshTriggerByClientCsReq, func() any { return new(proto.RefreshTriggerByClientCsReq) })
	c.regMsg(RefreshTriggerByClientScNotify, func() any { return new(proto.RefreshTriggerByClientScNotify) })
	c.regMsg(RefreshTriggerByClientScRsp, func() any { return new(proto.RefreshTriggerByClientScRsp) })
	c.regMsg(ReplaceLineupCsReq, func() any { return new(proto.ReplaceLineupCsReq) })
	c.regMsg(ReserveStaminaExchangeCsReq, func() any { return new(proto.ReserveStaminaExchangeCsReq) })
	c.regMsg(ReserveStaminaExchangeScRsp, func() any { return new(proto.ReserveStaminaExchangeScRsp) })
	c.regMsg(RevcMsgScNotify, func() any { return new(proto.RevcMsgScNotify) })
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
	c.regMsg(SelectRogueDialogueEventCsReq, func() any { return new(proto.SelectRogueDialogueEventCsReq) })
	c.regMsg(SelectRogueDialogueEventScRsp, func() any { return new(proto.SelectRogueDialogueEventScRsp) })
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
	c.regMsg(SyncEntityBuffChangeListScNotify, func() any { return new(proto.SyncEntityBuffChangeListScNotify) })
	c.regMsg(SyncHandleFriendScNotify, func() any { return new(proto.SyncHandleFriendScNotify) })
	c.regMsg(SyncLineupNotify, func() any { return new(proto.SyncLineupNotify) })
	c.regMsg(SyncRogueFinishScNotify, func() any { return new(proto.SyncRogueFinishScNotify) })
	c.regMsg(SyncRogueMapRoomScNotify, func() any { return new(proto.SyncRogueMapRoomScNotify) })
	c.regMsg(SyncRogueVirtualItemInfoScNotify, func() any { return new(proto.SyncRogueVirtualItemInfoScNotify) })
	c.regMsg(TakeChallengeRewardCsReq, func() any { return new(proto.TakeChallengeRewardCsReq) })
	c.regMsg(TakeChallengeRewardScRsp, func() any { return new(proto.TakeChallengeRewardScRsp) })
	c.regMsg(TakeMailAttachmentCsReq, func() any { return new(proto.TakeMailAttachmentCsReq) })
	c.regMsg(TakeMailAttachmentScRsp, func() any { return new(proto.TakeMailAttachmentScRsp) })
	c.regMsg(TakeOffEquipmentCsReq, func() any { return new(proto.TakeOffEquipmentCsReq) })
	c.regMsg(TakeOffRelicCsReq, func() any { return new(proto.TakeOffRelicCsReq) })
	c.regMsg(TakePromotionRewardCsReq, func() any { return new(proto.TakePromotionRewardCsReq) })
	c.regMsg(TakePromotionRewardScRsp, func() any { return new(proto.TakePromotionRewardScRsp) })
	c.regMsg(TextJoinQueryCsReq, func() any { return new(proto.TextJoinQueryCsReq) })
	c.regMsg(TextJoinQueryScRsp, func() any { return new(proto.TextJoinQueryScRsp) })
	c.regMsg(UnlockBackGroundMusicCsReq, func() any { return new(proto.UnlockBackGroundMusicCsReq) })
	c.regMsg(UnlockBackGroundMusicScRsp, func() any { return new(proto.UnlockBackGroundMusicScRsp) })
	c.regMsg(UnlockChatBubbleScNotify, func() any { return new(proto.UnlockChatBubbleScNotify) })
	c.regMsg(UnlockPhoneThemeScNotify, func() any { return new(proto.UnlockPhoneThemeScNotify) })
	c.regMsg(UnlockSkilltreeCsReq, func() any { return new(proto.UnlockSkilltreeCsReq) })
	c.regMsg(UnlockSkilltreeScRsp, func() any { return new(proto.UnlockSkilltreeScRsp) })
	c.regMsg(UseItemCsReq, func() any { return new(proto.UseItemCsReq) })
	c.regMsg(UseItemScRsp, func() any { return new(proto.UseItemScRsp) })
	c.regMsg(GmGive, func() any { return new(spb.GmGive) })
	c.regMsg(GmWorldLevel, func() any { return new(spb.GmWorldLevel) })
}

func (c *CmdProtoMap) regMsg(cmdId uint16, protoObjNewFunc func() any) {
	_, exist := c.cmdDeDupMap[cmdId]
	if exist {
		logger.Debug("reg dup msg, cmd id: %v", cmdId)
		return
	} else {
		c.cmdDeDupMap[cmdId] = true
	}
	protoObj := protoObjNewFunc().(pb.Message)
	refType := reflect.TypeOf(protoObj)
	// cmdId -> protoObj
	c.cmdIdProtoObjMap[cmdId] = refType
	// protoObj -> cmdId
	c.protoObjCmdIdMap[refType] = cmdId
	cmdName := refType.Elem().Name()
	// cmdId -> cmdName
	c.cmdIdCmdNameMap[cmdId] = cmdName
	// cmdName -> cmdId
	c.cmdNameCmdIdMap[cmdName] = cmdId
	// cmdId -> protoObjCache
	c.cmdIdProtoObjCacheMap[cmdId] = &sync.Pool{
		New: protoObjNewFunc,
	}
	// cmdId -> protoObjNewFunc
	c.cmdIdProtoObjFastNewMap[cmdId] = protoObjNewFunc
}

// 性能优化专用方法 若不满足使用条件 请老老实实的用下面的反射方法

// GetProtoObjCacheByCmdId 从缓存池获取一个对象 请务必确保能容忍获取到的对象含有使用过的脏数据 否则会产生不可预料的后果
func (c *CmdProtoMap) GetProtoObjCacheByCmdId(cmdId uint16) pb.Message {
	cachePool, exist := c.cmdIdProtoObjCacheMap[cmdId]
	if !exist {
		logger.Debug("unknown cmd id: %v", cmdId)
		return nil
	}
	protoObj := cachePool.Get().(pb.Message)
	return protoObj
}

// PutProtoObjCache 将使用结束的对象放回缓存池 请务必确保对象的生命周期真的已经结束了 否则会产生不可预料的后果
func (c *CmdProtoMap) PutProtoObjCache(cmdId uint16, protoObj pb.Message) {
	cachePool, exist := c.cmdIdProtoObjCacheMap[cmdId]
	if !exist {
		logger.Debug("unknown cmd id: %v", cmdId)
		return
	}
	cachePool.Put(protoObj)
}

func (c *CmdProtoMap) GetProtoObjFastNewByCmdId(cmdId uint16) pb.Message {
	fn, exist := c.cmdIdProtoObjFastNewMap[cmdId]
	if !exist {
		logger.Debug("unknown cmd id: %v", cmdId)
		return nil
	}
	protoObj := fn().(pb.Message)
	return protoObj
}

// 反射方法

func (c *CmdProtoMap) GetProtoObjByCmdId(cmdId uint16) pb.Message {
	refType, exist := c.cmdIdProtoObjMap[cmdId]
	if !exist {
		logger.Debug("unknown cmd id: %v", cmdId)
		return nil
	}
	protoObjInst := reflect.New(refType.Elem())
	protoObj := protoObjInst.Interface().(pb.Message)
	return protoObj
}

func (c *CmdProtoMap) GetCmdIdByProtoObj(protoObj pb.Message) uint16 {
	cmdId, exist := c.protoObjCmdIdMap[reflect.TypeOf(protoObj)]
	if !exist {
		logger.Debug("unknown proto object: %v", protoObj)
		return 0
	}
	return cmdId
}

func (c *CmdProtoMap) GetCmdNameByCmdId(cmdId uint16) string {
	cmdName, exist := c.cmdIdCmdNameMap[cmdId]
	if !exist {
		logger.Debug("unknown cmd id: %v", cmdId)
		return ""
	}
	return cmdName
}

func (c *CmdProtoMap) GetCmdIdByCmdName(cmdName string) uint16 {
	cmdId, exist := c.cmdNameCmdIdMap[cmdName]
	if !exist {
		logger.Debug("unknown cmd name: %v", cmdName)
		return 0
	}
	return cmdId
}
