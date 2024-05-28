package player

import (
	"strconv"

	base "github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 获取好友数据每次都去redis里取
func (g *GamePlayer) GetFriend() *spb.PlayerFriend {
	redisDb, ok := base.GetDb().GetPlayerFriend(g.Uid)
	if !ok {
		return &spb.PlayerFriend{
			Uid:             g.Uid,
			FriendList:      []uint32{0},
			RecvApplyFriend: make([]*spb.ReceiveApply, 0),
			SendApplyFriend: make([]uint32, 0),
		}
	}
	friend := new(spb.PlayerFriend)
	err := pb.Unmarshal(redisDb, friend)
	if err != nil {
		logger.Error("PlayerFriend Unmarshal error")
		return &spb.PlayerFriend{
			Uid:             g.Uid,
			FriendList:      []uint32{0},
			RecvApplyFriend: make([]*spb.ReceiveApply, 0),
			SendApplyFriend: make([]uint32, 0),
		}
	}
	return friend
}

func (g *GamePlayer) GetFriendList() []uint32 {
	db := g.GetFriend()
	if db.FriendList == nil {
		db.FriendList = []uint32{0}
	}
	return db.FriendList
}

func (g *GamePlayer) GetRecvApplyFriend() []*spb.ReceiveApply {
	db := g.GetFriend()
	if db.RecvApplyFriend == nil {
		db.RecvApplyFriend = make([]*spb.ReceiveApply, 0)
	}
	return db.RecvApplyFriend
}

func (g *GamePlayer) GetSendApplyFriend() []uint32 {
	db := g.GetFriend()
	if db.SendApplyFriend == nil {
		db.SendApplyFriend = make([]uint32, 0)
	}
	return db.SendApplyFriend
}

func (g *GamePlayer) GetPlayerBasicBriefData(uid uint32) *spb.PlayerBasicBriefData {
	if uid == 0 {
		return &spb.PlayerBasicBriefData{
			Nickname:          "hkrpg-go|game_server:" + strconv.Itoa(int(g.GameAppId)),
			Level:             80,
			WorldLevel:        8,
			LastLoginTime:     1,
			HeadImageAvatarId: 208002,
			Exp:               0,
			PlatformType:      spb.PlatformType(proto.PlatformType_CLOUD_PC),
			Uid:               0,
			Status:            spb.PlayerStatusType(proto.FriendOnlineStatus_FRIEND_ONLINE_STATUS_ONLINE),
			Signature:         "欢迎来到免费私人服务器 hkrpg-go|game_server:" + strconv.Itoa(int(g.GameAppId)),
		}
	}
	redisDb, ok := base.GetDb().GetPlayerPlayerBasicBriefData(uid)
	if !ok {
		return nil
	}
	friend := new(spb.PlayerBasicBriefData)
	err := pb.Unmarshal(redisDb, friend)
	if err != nil {
		logger.Error("player_brief_data Unmarshal error")
		return nil
	}
	return friend
}

/*******************************************接口*******************************************/

func (g *GamePlayer) GetPlayerSimpleInfo(uid uint32) *proto.PlayerSimpleInfo {
	friend := g.GetPlayerBasicBriefData(uid)
	if friend == nil {
		return nil
	}
	simpleInfo := &proto.PlayerSimpleInfo{
		Signature:      friend.Signature,
		LastActiveTime: friend.LastLoginTime,
		Level:          friend.Level,
		ChatBubbleId:   220003,
		Platform:       proto.PlatformType(friend.PlatformType),
		AssistSimpleList: []*proto.AssistSimpleInfo{
			{
				Pos:           0,
				AvatarId:      1212,
				Level:         80,
				DressedSkinId: 0,
			},
		},
		Uid:          friend.Uid,
		HeadIcon:     friend.HeadImageAvatarId,
		Nickname:     friend.Nickname,
		OnlineStatus: proto.FriendOnlineStatus(friend.Status),
	}
	return simpleInfo
}

func (g *GamePlayer) GetPlayerDetailInfo(uid uint32) *proto.PlayerDetailInfo {
	friend := g.GetPlayerBasicBriefData(uid)
	if friend == nil {
		return nil
	}
	playerDetailInfo := &proto.PlayerDetailInfo{
		DisplayAvatarList: make([]*proto.DisplayAvatarDetailInfo, 0),
		Record:            nil,
		AILINANGJNE:       "",
		WorldLevel:        friend.WorldLevel,
		Uid:               friend.Uid,
		EFNHCOEKDCN:       false,
		AssistAvatarList:  make([]*proto.DisplayAvatarDetailInfo, 0),
		Level:             friend.Level,
		IsBanned:          false,
		MAPJDADPKOL:       0,
		HeadIcon:          friend.HeadImageAvatarId,
		Platform:          proto.PlatformType(friend.PlatformType),
		AKFPFMGILAO:       0,
		RecordInfo:        &proto.DisplayRecordInfo{},
		LDFIOFJHJJA:       "",
		Signature:         friend.Signature,
		Nickname:          friend.Nickname,
	}

	return playerDetailInfo
}

func (g *GamePlayer) GetFriendApplyInfo(receiveApply *spb.ReceiveApply) *proto.FriendApplyInfo {
	friendApplyInfo := &proto.FriendApplyInfo{
		ApplyTime:  receiveApply.ApplyTime,
		PlayerInfo: g.GetPlayerSimpleInfo(receiveApply.ApplyUid),
	}
	return friendApplyInfo
}
