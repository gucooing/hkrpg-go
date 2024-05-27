package player

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) StaminaInfoScNotify() {
	db := g.GetMaterialMap()
	notify := &proto.StaminaInfoScNotify{
		NextRecoverTime: 0,
		Stamina:         db[Stamina],
		ReserveStamina:  db[RStamina],
	}
	g.Send(cmd.StaminaInfoScNotify, notify)
}

func (g *GamePlayer) HandleGetBasicInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBasicInfoScRsp)
	rsp.CurDay = 1
	rsp.NextRecoverTime = 1716449614
	rsp.GameplayBirthday = g.BasicBin.Birthday
	rsp.WeekCocoonFinishedCount = 0 // 周本完成计数
	rsp.PlayerSettingInfo = &proto.PlayerSettingInfo{
		B1:                true,
		B2:                true,
		B3:                true,
		B4:                true,
		B5:                true,
		B6:                true,
		DisplayRecordType: proto.DisplayRecordType_BATTLE_RECORD_CHALLENGE,
	}

	g.Send(cmd.GetBasicInfoScRsp, rsp)
}

func (g *GamePlayer) HandleGetArchiveDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetArchiveDataScRsp)
	archiveData := &proto.ArchiveData{
		ArchiveMissingAvatarIdList:    make([]uint32, 0),
		ArchiveEquipmentIdList:        make([]uint32, 0),
		ArchiveMissingEquipmentIdList: make([]uint32, 0),
		ArchiveMonsterIdList:          make([]*proto.MonsterArchive, 0),
		RelicList:                     make([]*proto.RelicArchive, 0),
	}

	for _, avatar := range g.BasicBin.Avatar.AvatarList {
		archiveData.ArchiveMissingAvatarIdList = append(archiveData.ArchiveMissingAvatarIdList, avatar.AvatarId)
	}

	for _, equipment := range gdconf.GetItemConfigEquipmentMap() {
		archiveData.ArchiveEquipmentIdList = append(archiveData.ArchiveEquipmentIdList, equipment.ID)
	}

	for _, monsterList := range gdconf.GetMonsterConfigMap() {
		archiveMonsterIdList := &proto.MonsterArchive{
			Num:       1,
			MonsterId: monsterList.MonsterID,
		}
		archiveData.ArchiveMonsterIdList = append(archiveData.ArchiveMonsterIdList, archiveMonsterIdList)
	}

	for _, relicList := range gdconf.GetRelicMap() {
		archiveRelicList := &proto.RelicArchive{
			RelicId: relicList.ID,
			Slot:    relicList.SetID,
		}
		archiveData.RelicList = append(archiveData.RelicList, archiveRelicList)
	}

	rsp.ArchiveData = archiveData

	g.Send(cmd.GetArchiveDataScRsp, rsp)
}

func (g *GamePlayer) GetUpdatedArchiveDataCsReq(payloadMsg []byte) {
	g.Send(cmd.GetUpdatedArchiveDataScRsp, nil)
}

func (g *GamePlayer) HandleGetPlayerBoardDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetPlayerBoardDataScRsp{
		CurrentHeadIconId:    g.BasicBin.HeadImageAvatarId,
		UnlockedHeadIconList: make([]*proto.HeadIcon, 0),
		Signature:            g.BasicBin.Signature,
		DisplayAvatarVec: &proto.DisplayAvatarVec{
			DisplayAvatarList: make([]*proto.DisplayAvatar, 0),
			IsDisplay:         false,
		},
	}

	for _, avatar := range g.GetHeadIconList() {
		headIcon := &proto.HeadIcon{
			Id: avatar,
		}
		rsp.UnlockedHeadIconList = append(rsp.UnlockedHeadIconList, headIcon)
	}

	g.Send(cmd.GetPlayerBoardDataScRsp, rsp)
}

func (g *GamePlayer) SetHeadIconCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetHeadIconCsReq, payloadMsg)
	req := msg.(*proto.SetHeadIconCsReq)

	g.BasicBin.HeadImageAvatarId = req.Id

	rsp := &proto.SetHeadIconScRsp{
		CurrentHeadIconId: req.Id,
	}

	g.Send(cmd.SetHeadIconScRsp, rsp)

}

func (g *GamePlayer) SetHeroBasicTypeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetHeroBasicTypeCsReq, payloadMsg)
	req := msg.(*proto.SetHeroBasicTypeCsReq)

	g.BasicBin.Avatar.CurMainAvatar = spb.HeroBasicType(req.BasicType)

	rsp := &proto.SetHeroBasicTypeScRsp{
		BasicType: req.BasicType,
	}

	g.Send(cmd.SetHeroBasicTypeScRsp, rsp)

}

func (g *GamePlayer) HandleGetFriendLoginInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetFriendLoginInfoScRsp)
	rsp.FriendUidList = []uint32{999}

	g.Send(cmd.GetFriendLoginInfoScRsp, rsp)
}

func (g *GamePlayer) GetFriendListInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetFriendListInfoScRsp)
	rsp.FriendList = make([]*proto.FriendListInfo, 0)
	simpleInfo := &proto.SimpleInfo{
		Signature:      "欢迎来到免费私人服务器 hkrpg-go",
		LastActiveTime: time.Now().Unix(),
		Level:          999,
		ChatBubbleId:   220003,
		PlatformType:   proto.PlatformType_MAC,
		AssistSimpleInfo: []*proto.AssistSimpleInfo{
			{
				Pos:           0,
				AvatarId:      1212,
				Level:         80,
				DressedSkinId: 0,
			},
		},
		Uid:          999,
		HeadIcon:     200106,
		Nickname:     "hkrpg-go",
		OnlineStatus: proto.FriendOnlineStatus_FRIEND_ONLINE_STATUS_ONLINE,
	}
	friendListInfo := &proto.FriendListInfo{SimpleInfo: simpleInfo}
	rsp.FriendList = append(rsp.FriendList, friendListInfo)

	g.Send(cmd.GetFriendListInfoScRsp, rsp)
}

func (g *GamePlayer) GetPrivateChatHistoryCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetPrivateChatHistoryCsReq, payloadMsg)
	req := msg.(*proto.GetPrivateChatHistoryCsReq)

	rsp := &proto.GetPrivateChatHistoryScRsp{
		SenderUid: req.SenderUid,
		ToUid:     req.ToUid,
		Retcode:   0,
		ChatList:  nil,
	}
	g.Send(cmd.GetPrivateChatHistoryScRsp, rsp)
}

func (g *GamePlayer) SendMsgCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SendMsgCsReq, payloadMsg)
	req := msg.(*proto.SendMsgCsReq)
	logger.Info("[ToUidList:%v][Emote:%v][MsgType:%s][Text:%s][ChatType:%s]", req.ToUidList, req.Emote, req.MsgType, req.Text, req.ChatType)

	for _, touid := range req.ToUidList {
		notify := &proto.RevcMsgScNotify{
			ToUid:    touid,
			Emote:    req.Emote,
			MsgType:  req.MsgType,
			FromUid:  g.Uid,
			Text:     req.Text,
			ChatType: req.ChatType,
		}
		g.Send(cmd.RevcMsgScNotify, notify)
	}

	g.Send(cmd.SendMsgScRsp, nil)
}

func (g *GamePlayer) HandleGetChatEmojiListCsReq(payloadMsg []byte) {
	g.Send(cmd.GetChatEmojiListScRsp, nil)
}

func (g *GamePlayer) HandleGetAssistHistoryCsReq(payloadMsg []byte) {
	g.Send(cmd.GetAssistHistoryScRsp, nil)
}

func (g *GamePlayer) SetClientPausedCsReq(payloadMsg []byte) {
	rsp := new(proto.SetClientPausedScRsp)
	g.OnlineData.IsPaused = !g.OnlineData.IsPaused
	rsp.Paused = g.OnlineData.IsPaused

	g.Send(cmd.SetClientPausedScRsp, rsp)
}

func (g *GamePlayer) HandleGetJukeboxDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetJukeboxDataScRsp)
	rsp.PlayingId = 210000
	rsp.MusicList = make([]*proto.UnlockedMusic, 0)
	for _, backMusicList := range gdconf.GetBackGroundMusicMap() {
		musicList := &proto.UnlockedMusic{
			GroupId: backMusicList.GroupID,
			Unkbool: true,
			Id:      backMusicList.ID,
		}
		rsp.MusicList = append(rsp.MusicList, musicList)
	}
	g.Send(cmd.GetJukeboxDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetPhoneDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetPhoneDataScRsp)
	rsp.CurChatBubble = 220000
	rsp.CurPhoneTheme = 221000
	rsp.OwnedChatBubbles = []uint32{220002, 220000, 220001}
	rsp.OwnedPhoneThemes = []uint32{221000, 221001, 221002, 221003}

	g.Send(cmd.GetPhoneDataScRsp, rsp)
}

func (g *GamePlayer) SetNicknameCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetNicknameCsReq, payloadMsg)
	req := msg.(*proto.SetNicknameCsReq)

	if g.OnlineData.IsNickName {
		g.BasicBin.Nickname = req.Nickname
	}

	g.OnlineData.IsNickName = !g.OnlineData.IsNickName

	g.PlayerPlayerSyncScNotify()
	g.Send(cmd.SetNicknameScRsp, nil)
}

func (g *GamePlayer) SetGameplayBirthdayCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetGameplayBirthdayCsReq, payloadMsg)
	req := msg.(*proto.SetGameplayBirthdayCsReq)

	g.BasicBin.Birthday = req.Birthday

	rsp := &proto.SetGameplayBirthdayScRsp{Birthday: req.Birthday}

	g.Send(cmd.SetGameplayBirthdayScRsp, rsp)
}

func (g *GamePlayer) SetSignatureCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetSignatureCsReq, payloadMsg)
	req := msg.(*proto.SetSignatureCsReq)

	g.BasicBin.Signature = req.Signature

	rsp := &proto.SetSignatureScRsp{Signature: req.Signature}

	g.Send(cmd.SetSignatureScRsp, rsp)
}

func (g *GamePlayer) TextJoinQueryCsReq(payloadMsg []byte) {
	rsp := new(proto.TextJoinQueryScRsp)
	for _, textJoin := range gdconf.GetTextJoinConfigMap() {
		textJoinList := &proto.TextJoinInfo{
			TextItemId:       textJoin.TextJoinID,
			TextItemConfigId: textJoin.TextJoinItemList[len(textJoin.TextJoinItemList)-1],
		}
		rsp.TextJoinList = append(rsp.TextJoinList, textJoinList)
	}

	g.Send(cmd.TextJoinQueryScRsp, rsp)
}

func (g *GamePlayer) GetUnlockTeleportCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetUnlockTeleportCsReq, payloadMsg)
	req := msg.(*proto.GetUnlockTeleportCsReq)
	rsp := &proto.GetUnlockTeleportScRsp{
		UnlockedTeleportList: make([]uint32, 0),
	}

	for _, id := range req.EntryIdList {
		excel := gdconf.GetMapEntranceById(strconv.Itoa(int(id)))
		if excel == nil {
			continue
		}
		teleportsMap := gdconf.GetTeleportsById(excel.PlaneID, excel.FloorID)
		if teleportsMap == nil {
			continue
		}
		for _, teleports := range teleportsMap.Teleports {
			rsp.UnlockedTeleportList = append(rsp.UnlockedTeleportList, teleports.MappingInfoID)
		}
	}

	g.Send(cmd.GetUnlockTeleportScRsp, rsp)
}

func (g *GamePlayer) HandlePlayerLoginFinishCsReq(payloadMsg []byte) {
	g.Send(cmd.PlayerLoginFinishScRsp, nil)
	// TODO 主动调用
	g.HandleGetArchiveDataCsReq(nil)
}
