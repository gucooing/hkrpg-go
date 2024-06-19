package player

import (
	"strconv"

	base "github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func NewFriend() *spb.PlayerFriend {
	return &spb.PlayerFriend{}
}

func (g *GamePlayer) GetFriend() *spb.PlayerFriend {
	bin := g.GetBasicBin()
	if bin.Friend == nil {
		bin.Friend = NewFriend()
	}
	return bin.Friend
}

func (g *GamePlayer) GetFriendList() map[uint32]*spb.Friend {
	db := g.GetFriend()
	if db.FriendList == nil {
		db.FriendList = make(map[uint32]*spb.Friend)
		db.FriendList[0] = &spb.Friend{
			Uid:        0,
			IsMarked:   true,
			RemarkName: "",
		}
	}
	return db.FriendList
}

func (g *GamePlayer) GetFriendByUid(uid uint32) *spb.Friend {
	db := g.GetFriendList()
	return db[uid]
}

func (g *GamePlayer) AddFriend(uid uint32) {
	db := g.GetFriendList()
	db[uid] = &spb.Friend{
		Uid:        uid,
		IsMarked:   false,
		RemarkName: "",
	}
}

// 获取好友申请每次都去redis里取
func (g *GamePlayer) GetRecvApplyFriend() map[uint32]*spb.ReceiveApply {
	friend := new(spb.ApplyFriend)
	if g.IsPE {
		return friend.RecvApplyFriend
	}
	redisDb, ok := database.GetPlayerFriend(base.Db.PlayerBriefDataRedis, g.Uid)
	if !ok {
		return make(map[uint32]*spb.ReceiveApply, 0)
	}
	err := pb.Unmarshal(redisDb, friend)
	if err != nil {
		logger.Error("PlayerFriend Unmarshal error")
		return make(map[uint32]*spb.ReceiveApply, 0)
	}
	return friend.RecvApplyFriend
}

func (g *GamePlayer) GetPlayerBasicBriefData(uid uint32) *spb.PlayerBasicBriefData {
	if uid == 0 || g.IsPE {
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

// 将redis里的好友加入mysql里
func (g *GamePlayer) InspectionRedisAcceptApplyFriend() {
	friend := new(spb.AcceptApplyFriend)
	redisDb, ok := database.GetAcceptApplyFriend(base.Db.PlayerBriefDataRedis, g.Uid)
	if !ok {
		return
	}
	err := pb.Unmarshal(redisDb, friend)
	if err != nil {
		logger.Error("PlayerFriend Unmarshal error")
		return
	}
	if friend.RecvApplyFriend != nil {
		for uid := range friend.RecvApplyFriend {
			g.AddFriend(uid)
		}
	}
	// TODO 处理完了要通知node删掉该信息(无所谓我自己删，覆写就覆写
	database.DelAcceptApplyFriend(base.Db.PlayerBriefDataRedis, g.Uid)
}

/*******************************************接口*******************************************/

func (g *GamePlayer) GetPlayerSimpleInfo(uid uint32) *proto.PlayerSimpleInfo {
	friend := g.GetPlayerBasicBriefData(uid)
	if friend == nil {
		return nil
	}
	simpleInfo := &proto.PlayerSimpleInfo{
		AILINANGJNE:    "",
		ChatBubbleId:   220003,
		IsBanned:       false,
		HeadIcon:       friend.HeadImageAvatarId,
		LDFIOFJHJJA:    "",
		Signature:      friend.Signature,
		Platform:       proto.PlatformType(friend.PlatformType),
		LastActiveTime: friend.LastLoginTime,
		OnlineStatus:   proto.FriendOnlineStatus(friend.Status),
		Nickname:       friend.Nickname,
		Uid:            friend.Uid,
		Level:          friend.Level,
		AssistSimpleList: []*proto.AssistSimpleInfo{
			{
				Pos:           0,
				AvatarId:      1212,
				Level:         80,
				DressedSkinId: 0,
			},
		},
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
		Record: &proto.PlayerBasicBrief{
			Level:                  friend.Level,
			UnlockedAvatarNum:      999,
			UnlockedAchievementNum: 999,
			UnlockedBookNum:        999,
			UnlockedMusicNum:       999,
			FKBLOGEAFJJ:            2000,
			CollectionInfo: &proto.PlayerCollectionInfo{
				DCIOBLHLICO: 2006,
				KLLEONMNLDI: 60,
			},
			WorldLevel: friend.WorldLevel,
		},
		AILINANGJNE:      "",
		WorldLevel:       friend.WorldLevel,
		Uid:              friend.Uid,
		EFNHCOEKDCN:      true,
		AssistAvatarList: make([]*proto.DisplayAvatarDetailInfo, 0),
		Level:            friend.Level,
		IsBanned:         false,
		MAPJDADPKOL:      0,
		HeadIcon:         friend.HeadImageAvatarId,
		Platform:         proto.PlatformType(friend.PlatformType),
		AKFPFMGILAO:      0,
		RecordInfo:       &proto.DisplayRecordInfo{},
		LDFIOFJHJJA:      "",
		Signature:        friend.Signature,
		Nickname:         friend.Nickname,
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

func (g *GamePlayer) GetFriendSimpleInfo(uid uint32) *proto.FriendSimpleInfo {
	db := g.GetFriendByUid(uid)
	simpleInfo := g.GetPlayerSimpleInfo(uid)
	if db == nil || simpleInfo == nil {
		return nil
	}
	friendSimpleInfo := &proto.FriendSimpleInfo{
		PlayerInfo:  simpleInfo,    // 基本信息
		RemarkName:  db.RemarkName, // 备注
		PlayerState: 0,
		CFMIKLHJMLE: nil,
		IsMarked:    db.IsMarked, // 是否特别关注
	}
	return friendSimpleInfo
}
