package player

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) PlayerHeartBeatCsReq() {
	g.LastActiveTime = time.Now().Unix()
}

func (g *GamePlayer) StaminaInfoScNotify() {
	notify := &proto.StaminaInfoScNotify{
		NextRecoverTime: 0,
		Stamina:         g.GetItem().MaterialMap[11],
		ReserveStamina:  g.GetItem().MaterialMap[12],
	}
	g.Send(cmd.StaminaInfoScNotify, notify)
}

func (g *GamePlayer) HandleGetBasicInfoCsReq() {
	rsp := new(proto.GetBasicInfoScRsp)
	rsp.CurDay = 1
	rsp.NextRecoverTime = 1698768000
	rsp.GameplayBirthday = g.PlayerPb.Birthday
	rsp.PlayerSettingInfo = &proto.PlayerSettingInfo{}

	g.Send(cmd.GetBasicInfoScRsp, rsp)
}

func (g *GamePlayer) HandleGetArchiveDataCsReq() {
	rsp := new(proto.GetArchiveDataScRsp)
	archiveData := &proto.ArchiveData{
		ArchiveAvatarIdList:    make([]uint32, 0),
		ArchiveEquipmentIdList: make([]uint32, 0),
		ArchiveMonsterIdList:   make([]*proto.MonsterArchive, 0),
		RelicList:              make([]*proto.RelicArchive, 0),
	}

	for _, avatar := range g.PlayerPb.Avatar.Avatar {
		archiveData.ArchiveAvatarIdList = append(archiveData.ArchiveAvatarIdList, avatar.AvatarId)
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
			Slot:    relicList.Type,
		}
		archiveData.RelicList = append(archiveData.RelicList, archiveRelicList)
	}

	rsp.ArchiveData = archiveData

	g.Send(cmd.GetArchiveDataScRsp, rsp)
}

func (g *GamePlayer) GetUpdatedArchiveDataCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.GetUpdatedArchiveDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetPlayerBoardDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetPlayerBoardDataScRsp{
		CurrentHeadIconId:    g.PlayerPb.HeadImageAvatarId,
		UnlockedHeadIconList: make([]*proto.HeadIcon, 0),
		Signature:            g.PlayerPb.Signature,
		// TODO
		DisplayAvatarVec: &proto.DisplayAvatarVec{
			DisplayAvatarList: nil,
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

	g.PlayerPb.HeadImageAvatarId = req.Id

	rsp := &proto.SetHeadIconScRsp{
		CurrentHeadIconId: req.Id,
	}

	g.Send(cmd.SetHeadIconScRsp, rsp)

}

func (g *GamePlayer) SetHeroBasicTypeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetHeroBasicTypeCsReq, payloadMsg)
	req := msg.(*proto.SetHeroBasicTypeCsReq)

	g.PlayerPb.Avatar.CurMainAvatar = spb.HeroBasicType(req.BasicType)

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

func (g *GamePlayer) GetFriendListInfoCsReq() {
	rsp := new(proto.GetFriendListInfoScRsp)
	rsp.FriendList = make([]*proto.FriendListInfo, 0)
	simpleInfo := &proto.SimpleInfo{
		Signature:      "欢迎来到免费私人服务器 hkrpg-go",
		LastActiveTime: time.Now().Unix(),
		Level:          999,
		ChatBubbleId:   220003,
		PlatformType:   proto.PlatformType_MAC,
		SimpleAvatarInfo: &proto.SimpleAvatarInfo{
			AvatarId:      1212,
			Level:         80,
			DressedSkinId: 0,
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
}

func (g *GamePlayer) HandleGetRogueHandbookDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetRogueHandbookDataScRsp)
	handbookInfo := &proto.RogueHandbookData{
		// RogueCurrentVersion: 1,
		// IsMiracleUnlock:     true,
	}
	rsp.HandbookInfo = handbookInfo

	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetChallengeScRsp)
	rsp.ChallengeList = make([]*proto.Challenge, 0)
	rsp.ChallengeRewardList = make([]*proto.ChallengeReward, 0)
	challengeDb := g.GetChallenge()
	for id, stars := range challengeDb.ChallengeList {
		challenge := &proto.Challenge{
			ChallengeId: id,
			Stars:       stars.Stars,
			Score:       stars.ScoreOne,
			ScoreTwo:    stars.ScoreTwo,
		}
		rsp.ChallengeList = append(rsp.ChallengeList, challenge)
	}
	for taken, id := range challengeDb.ChallengeRewardList {
		challengeReward := &proto.ChallengeReward{
			TakenChallengeReward: taken,
			GroupId:              id,
		}
		rsp.ChallengeRewardList = append(rsp.ChallengeRewardList, challengeReward)
	}
	g.Send(cmd.GetChallengeScRsp, rsp)
}

func (g *GamePlayer) HandleGetChatEmojiListCsReq(payloadMsg []byte) {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.GetChatEmojiListScRsp, rsp)
}

func (g *GamePlayer) HandleGetAssistHistoryCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.GetAssistHistoryScRsp, rsp)
}

func (g *GamePlayer) GetMailCsReq() {
	rsp := new(proto.GetMailScRsp)
	rsp.TotalNum = 1
	rsp.IsEnd = true
	mailList := &proto.ClientMail{
		Sender:  "gucooing",
		Content: "欢迎来到 hkrpg-go server",
		Title:   "欢迎来到 hkrpg-go server",
	}
	rsp.MailList = append(rsp.MailList, mailList)

	g.Send(cmd.GetMailScRsp, rsp)
}

func (g *GamePlayer) SetClientPausedCsReq() {
	rsp := new(proto.SetClientPausedScRsp)
	g.Player.IsPaused = !g.Player.IsPaused
	rsp.Paused = g.Player.IsPaused

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

	if g.Player.IsNickName {
		g.PlayerPb.Nickname = req.Nickname
	}

	g.Player.IsNickName = !g.Player.IsNickName

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因

	g.PlayerPlayerSyncScNotify()
	g.Send(cmd.SetNicknameScRsp, rsp)
}

func (g *GamePlayer) SetGameplayBirthdayCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetGameplayBirthdayCsReq, payloadMsg)
	req := msg.(*proto.SetGameplayBirthdayCsReq)

	g.PlayerPb.Birthday = req.Birthday

	rsp := &proto.SetGameplayBirthdayScRsp{Birthday: req.Birthday}

	g.Send(cmd.SetGameplayBirthdayScRsp, rsp)
}

func (g *GamePlayer) SetSignatureCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetSignatureCsReq, payloadMsg)
	req := msg.(*proto.SetSignatureCsReq)

	g.PlayerPb.Signature = req.Signature

	rsp := &proto.SetSignatureScRsp{Signature: req.Signature}

	g.Send(cmd.SetSignatureScRsp, rsp)
}

func (g *GamePlayer) TextJoinQueryCsReq() {
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
		teleport := gdconf.GetTeleportsById(excel.PlaneID, excel.FloorID)
		if teleport == nil {
			continue
		}
		/*
			for tid := range teleport {
				rsp.UnlockedTeleportList = append(rsp.UnlockedTeleportList, tid)
			}
		*/
	}

	g.Send(cmd.GetUnlockTeleportScRsp, rsp)
}

func (g *GamePlayer) HandlePlayerLoginFinishCsReq(payloadMsg []byte) {
	rsp := new(proto.PlayerHeartbeatScRsp)
	// TODO 逆天了，proto太残了，没办法
	g.Send(cmd.PlayerLoginFinishScRsp, rsp)
	// TODO 主动调用
	g.HandleGetArchiveDataCsReq()
}

func (g *GamePlayer) GetFarmStageGachaInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetFarmStageGachaInfoCsReq, payloadMsg)
	req := msg.(*proto.GetFarmStageGachaInfoCsReq)

	rsp := &proto.GetFarmStageGachaInfoScRsp{
		FarmStageGachaInfoList: make([]*proto.FarmStageGachaInfo, 0),
	}

	for _, farmStageGachaId := range req.FarmStageGachaIdList {
		farmStageGachaInfo := &proto.FarmStageGachaInfo{
			BeginTime: 1664308800,
			GachaId:   farmStageGachaId,
			EndTime:   4294967295,
		}
		rsp.FarmStageGachaInfoList = append(rsp.FarmStageGachaInfoList, farmStageGachaInfo)
	}

	g.Send(cmd.GetFarmStageGachaInfoScRsp, rsp)
}
