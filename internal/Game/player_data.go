package Game

import (
	"time"

	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetArchiveDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetArchiveDataScRsp)
	var archiveAvatarIdList []uint32
	for _, a := range g.Player.DbAvatar.Avatar {
		archiveAvatarIdList = append(archiveAvatarIdList, a.AvatarId)
	}
	archiveData := &proto.ArchiveData{
		ArchiveAvatarIdList:    archiveAvatarIdList,
		ArchiveEquipmentIdList: nil,
	}
	rsp.ArchiveData = archiveData

	g.send(cmd.GetArchiveDataScRsp, rsp)
}

func (g *Game) HandleGetPlayerBoardDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetPlayerBoardDataScRsp{
		CurrentHeadIconId:    g.Player.HeadImage,
		UnlockedHeadIconList: make([]*proto.HeadIcon, 0),
		Signature:            g.Player.Signature,
		Unk1:                 "",
	}

	for _, avatar := range g.Player.DbAvatar.Avatar {
		headIcon := &proto.HeadIcon{
			Id: avatar.AvatarId + 200000,
		}
		rsp.UnlockedHeadIconList = append(rsp.UnlockedHeadIconList, headIcon)
	}

	g.send(cmd.GetPlayerBoardDataScRsp, rsp)
}

func (g *Game) SetHeadIconCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SetHeadIconCsReq, payloadMsg)
	req := msg.(*proto.SetHeadIconCsReq)

	g.Player.HeadImage = req.Id

	rsp := &proto.SetHeadIconScRsp{
		CurrentHeadIconId: req.Id,
	}

	g.send(cmd.SetHeadIconScRsp, rsp)

	g.UpDataPlayer()
}

func (g *Game) SetHeroBasicTypeCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SetHeroBasicTypeCsReq, payloadMsg)
	req := msg.(*proto.SetHeroBasicTypeCsReq)

	g.Player.DbAvatar.MainAvatar = req.BasicType

	rsp := &proto.SetHeroBasicTypeScRsp{
		BasicType: req.BasicType,
	}

	g.send(cmd.SetHeroBasicTypeScRsp, rsp)

	g.UpDataPlayer()
}

func (g *Game) HandleGetFriendLoginInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetFriendLoginInfoScRsp)
	rsp.FriendUidList = []uint32{99}

	g.send(cmd.GetFriendLoginInfoScRsp, rsp)
}

func (g *Game) HandleGetRogueHandbookDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetRogueHandbookDataScRsp)
	handbookInfo := &proto.RogueHandbookData{
		RogueCurrentVersion: 1,
		IsMiracleUnlock:     true,
	}
	rsp.HandbookInfo = handbookInfo

	g.send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func (g *Game) HandleGetChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetChallengeScRsp)

	g.send(cmd.GetChallengeScRsp, rsp)
}

func (g *Game) HandleGetChatEmojiListCsReq(payloadMsg []byte) {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.send(cmd.GetChatEmojiListScRsp, rsp)
}

func (g *Game) HandleGetAssistHistoryCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.send(cmd.GetAssistHistoryScRsp, rsp)
}

func (g *Game) GetMailCsReq() {
	rsp := new(proto.GetMailScRsp)
	rsp.TotalNum = 1
	rsp.IsEnd = true
	mailList := &proto.ClientMail{
		Sender:  "gucooing",
		Content: "欢迎来到 hkrpg-go server",
		Title:   "欢迎来到 hkrpg-go server",
	}
	rsp.MailList = append(rsp.MailList, mailList)

	g.send(cmd.GetMailScRsp, rsp)
}

func (g *Game) SetClientPausedCsReq() {
	rsp := new(proto.SetClientPausedScRsp)
	g.Player.IsPaused = false
	g.Player.IsPaused = !g.Player.IsPaused
	rsp.Paused = g.Player.IsPaused

	g.send(cmd.SetClientPausedScRsp, rsp)
}

func (g *Game) GetFirstTalkNpcCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.send(cmd.GetFirstTalkNpcScRsp, rsp)
}

func (g *Game) HandleGetJukeboxDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetJukeboxDataScRsp)
	rsp.PlayingId = 210000
	rsp.MusicList = make([]*proto.GetJukeboxDataScRsp_UnlockedMusic, 0)
	musicList := &proto.GetJukeboxDataScRsp_UnlockedMusic{
		GroupId: 3,
		Unkbool: true,
		Id:      210215,
	}
	rsp.MusicList = append(rsp.MusicList, musicList)
	g.send(cmd.GetJukeboxDataScRsp, rsp)
}

func (g *Game) HandleGetPhoneDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetPhoneDataScRsp)
	rsp.CurChatBubble = 220000
	rsp.CurPhoneTheme = 221000
	rsp.OwnedChatBubbles = []uint32{220002, 220000, 220001}
	rsp.OwnedPhoneThemes = []uint32{221000, 221001, 221002, 221003}

	g.send(cmd.GetPhoneDataScRsp, rsp)
}

func (g *Game) SetNicknameCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SetNicknameCsReq, payloadMsg)
	req := msg.(*proto.SetNicknameCsReq)

	if g.Player.IsNickName {
		g.Player.NickName = req.Nickname
		g.UpDataPlayer()
	}

	g.Player.IsNickName = !g.Player.IsNickName

	playerSyncScNotify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   req.Nickname,
			Level:      g.Player.Level,
			Exp:        g.Player.Exp,
			Stamina:    g.Player.Stamina,
			Mcoin:      0,
			Hcoin:      0,
			Scoin:      0,
			WorldLevel: g.Player.WorldLevel,
		},
	}

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因

	g.send(cmd.PlayerSyncScNotify, playerSyncScNotify)
	g.send(cmd.SetNicknameScRsp, rsp)
}

func (g *Game) SetGameplayBirthdayCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SetGameplayBirthdayCsReq, payloadMsg)
	req := msg.(*proto.SetGameplayBirthdayCsReq)

	g.Player.Birthday = req.Birthday

	rsp := &proto.SetGameplayBirthdayScRsp{Birthday: req.Birthday}

	g.send(cmd.SetGameplayBirthdayScRsp, rsp)

	g.UpDataPlayer()
}

func (g *Game) SetSignatureCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SetSignatureCsReq, payloadMsg)
	req := msg.(*proto.SetSignatureCsReq)

	g.Player.Signature = req.Signature

	rsp := &proto.SetSignatureScRsp{Signature: req.Signature}

	g.send(cmd.SetSignatureScRsp, rsp)

	g.UpDataPlayer()
}

func (g *Game) HandlePlayerHeartBeatCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.PlayerHeartBeatCsReq, payloadMsg)
	req := msg.(*proto.PlayerHeartbeatCsReq)

	rsp := new(proto.PlayerHeartbeatScRsp)
	rsp.ServerTimeMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.ClientTimeMs = req.ClientTimeMs

	g.send(cmd.PlayerHeartBeatScRsp, rsp)
}
