package model

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func NewFriend() *spb.PlayerFriend {
	return &spb.PlayerFriend{}
}

func (g *PlayerData) GetFriend() *spb.PlayerFriend {
	bin := g.GetBasicBin()
	if bin.Friend == nil {
		bin.Friend = NewFriend()
	}
	return bin.Friend
}

func (g *PlayerData) GetFriendList() map[uint32]*spb.Friend {
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

func (g *PlayerData) GetFriendByUid(uid uint32) *spb.Friend {
	db := g.GetFriendList()
	return db[uid]
}

func (g *PlayerData) AddFriend(uid uint32) {
	db := g.GetFriendList()
	db[uid] = &spb.Friend{
		Uid:        uid,
		IsMarked:   false,
		RemarkName: "",
	}
}

func (g *PlayerData) GetAssistAvatarList() map[uint32]uint32 {
	db := g.GetFriend()
	if db.AssistAvatarList == nil {
		db.AssistAvatarList = make(map[uint32]uint32)
	}
	return db.AssistAvatarList
}

func (g *PlayerData) GetDisplayAvatarlist() map[uint32]uint32 {
	db := g.GetFriend()
	if db.DisplayAvatarList == nil {
		db.DisplayAvatarList = make(map[uint32]uint32)
	}
	return db.DisplayAvatarList
}

// 获取好友申请每次都去redis里取
func (g *PlayerData) GetRecvApplyFriend() map[uint32]*spb.ReceiveApply {
	friend := new(spb.ApplyFriend)
	redisDb, ok := database.GetApplyFriend(database.GSS.PlayerBriefDataRedis,
		database.PE, g.GetBasicBin().Uid)
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

func (g *PlayerData) GetPlayerBasicBriefData(uid uint32) *spb.PlayerBasicBriefData {
	if uid == 0 {
		return &spb.PlayerBasicBriefData{
			Nickname:          "hkrpg-go",
			Level:             80,
			WorldLevel:        8,
			LastLoginTime:     1,
			HeadImageAvatarId: 208002,
			Exp:               0,
			PlatformType:      spb.PlatformType(proto.PlatformType_CLOUD_PC),
			Uid:               0,
			Status:            spb.PlayerStatusType(proto.FriendOnlineStatus_FRIEND_ONLINE_STATUS_ONLINE),
			Signature:         "欢迎来到免费私人服务器 hkrpg-go",
		}
	}
	bin, ok := database.GetPlayerBasic(database.GSS.PlayerBriefDataRedis,
		database.PE, uid)
	if !ok {
		return nil
	}
	friend := new(spb.PlayerBasicBriefData)
	err := pb.Unmarshal(bin, friend)
	if err != nil {
		logger.Error("player_brief_data Unmarshal error")
		return nil
	}
	return friend
}

// 将redis里的好友加入mysql里
func (g *PlayerData) InspectionRedisAcceptApplyFriend() {
	friend := new(spb.AcceptApplyFriend)
	redisDb, ok := database.GetAcceptApplyFriend(database.GSS.PlayerBriefDataRedis,
		database.PE, g.GetBasicBin().Uid)
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
	database.DelAcceptApplyFriend(database.GSS.PlayerBriefDataRedis,
		database.PE, g.GetBasicBin().Uid)
}

/*******************************************接口*******************************************/

func (g *PlayerData) GetPlayerSimpleInfo(uid uint32) *proto.PlayerSimpleInfo {
	friend := g.GetPlayerBasicBriefData(uid)
	if friend == nil {
		return nil
	}
	simpleInfo := &proto.PlayerSimpleInfo{
		ChatBubbleId:         friend.ChatBubbleId,
		IsBanned:             false,
		HeadIcon:             friend.HeadImageAvatarId,
		Signature:            friend.Signature,
		Platform:             proto.PlatformType(friend.PlatformType),
		LastActiveTime:       friend.LastLoginTime,
		OnlineStatus:         proto.FriendOnlineStatus(friend.Status),
		Nickname:             friend.Nickname,
		Uid:                  friend.Uid,
		Level:                friend.Level,
		AssistSimpleInfoList: make([]*proto.AssistSimpleInfo, 0),
	}
	return simpleInfo
}

func (g *PlayerData) GetPlayerDetailInfo(uid uint32) *proto.PlayerDetailInfo {
	friend := g.GetPlayerBasicBriefData(uid)
	if friend == nil {
		return nil
	}
	playerDetailInfo := &proto.PlayerDetailInfo{
		RecordInfo: &proto.PlayerRecordInfo{
			// ArchiveAvatarNum:    1,
			// ArchiveBookNum:      1,
			// ArchiveEquipmentNum: 1,
			// ArchiveMusicNum:     1,
			// ArchiveQuestNum:     1,
			CollectionInfo: &proto.PlayerCollectionInfo{
				// KJNOOOJDGDN: 5,
				// NCCNLKFCAKM: 6,
			},
		},
		WorldLevel:        friend.WorldLevel,
		Uid:               friend.Uid,
		AssistAvatarList:  g.GetDisplayAvatarDetailInfoList(friend.AssistAvatarList),
		DisplayAvatarList: g.GetDisplayAvatarDetailInfoList(friend.DisplayAvatarList),
		Level:             friend.Level,
		IsBanned:          false,
		HeadIcon:          friend.HeadImageAvatarId,
		Platform:          proto.PlatformType(friend.PlatformType),
		Signature:         friend.Signature,
		Nickname:          friend.Nickname,
		// KPFMBKIAGMJ:       true,
		// FLHDCJECCPN:       18,
		PrivacySettings: &proto.PrivacySettings{
			// OJNELKIOAOK: true,
			// DAAAIHDPCFE: true,
			// MAJIMDCHNDL: true,
			// MOKMEEDBECL: true,
			// BBJGEGEJJFB: true,
		},
	}

	return playerDetailInfo
}

func (g *PlayerData) GetFriendApplyInfo(receiveApply *spb.ReceiveApply) *proto.FriendApplyInfo {
	friendApplyInfo := &proto.FriendApplyInfo{
		ApplyTime:  receiveApply.ApplyTime,
		PlayerInfo: g.GetPlayerSimpleInfo(receiveApply.ApplyUid),
	}
	return friendApplyInfo
}

func (g *PlayerData) GetFriendSimpleInfo(uid uint32) *proto.FriendSimpleInfo {
	db := g.GetFriendByUid(uid)
	simpleInfo := g.GetPlayerSimpleInfo(uid)
	if db == nil || simpleInfo == nil {
		return nil
	}
	friendSimpleInfo := &proto.FriendSimpleInfo{
		PlayerInfo:   simpleInfo,    // 基本信息
		RemarkName:   db.RemarkName, // 备注
		PlayingState: 0,
		IsMarked:     db.IsMarked, // 是否特别关注
	}
	return friendSimpleInfo
}

func (g *PlayerData) GetAssistAvatarListSpb(list map[uint32]uint32) map[uint32]*spb.AssistAvatar {
	infoList := make(map[uint32]*spb.AssistAvatar)
	if list == nil {
		return infoList
	}
	for pos, avatarId := range list {
		if db := g.GetAvatarBinById(avatarId); db != nil {
			path := g.GetCurMultiPathAvatar(avatarId)
			if path == nil {
				continue
			}
			assistAvatar := &spb.AssistAvatar{
				AvatarId:      avatarId,
				Rank:          path.Rank,
				Level:         db.Level,
				Exp:           db.Exp,
				PromoteLevel:  db.PromoteLevel,
				Equipment:     g.GetEquipmentById(path.EquipmentUniqueId),
				EquipRelic:    make(map[uint32]*spb.Relic),
				SkilltreeList: path.SkilltreeList,
			}
			// add EquipRelic
			for id, uniqueId := range path.EquipRelic {
				assistAvatar.EquipRelic[id] = g.GetRelicById(uniqueId)
			}

			infoList[pos] = assistAvatar
		}
	}
	return infoList
}

func (g *PlayerData) GetDisplayAvatarDetailInfoList(list map[uint32]*spb.AssistAvatar) []*proto.DisplayAvatarDetailInfo {
	infoList := make([]*proto.DisplayAvatarDetailInfo, 0)
	if list == nil {
		return infoList
	}
	for pos, db := range list {
		info := &proto.DisplayAvatarDetailInfo{
			AvatarId:      db.AvatarId,
			Rank:          db.Rank,
			Level:         db.Level,
			Promotion:     db.PromoteLevel,
			Exp:           db.Exp,
			Equipment:     nil,
			RelicList:     make([]*proto.DisplayRelicInfo, 0),
			SkilltreeList: make([]*proto.AvatarSkillTree, 0),
			Pos:           pos,

			DressedSkinId: 0,
		}
		// add Equipment
		if db.Equipment != nil {
			info.Equipment = &proto.DisplayEquipmentInfo{
				Tid:       db.Equipment.Tid,
				Rank:      db.Equipment.Rank,
				Exp:       db.Equipment.Exp,
				Promotion: db.Equipment.Promotion,
				Level:     db.Equipment.Level,
			}
		}
		// add RelicList
		if db.EquipRelic != nil {
			for id, redis := range db.EquipRelic {
				displayRelicInfo := &proto.DisplayRelicInfo{
					Exp:          redis.Exp,
					Type:         id,
					MainAffixId:  redis.MainAffixId,
					Tid:          redis.Tid,
					SubAffixList: make([]*proto.RelicAffix, 0),
					Level:        redis.Level,
				}
				for _, subAffix := range redis.RelicAffix {
					displayRelicInfo.SubAffixList = append(displayRelicInfo.SubAffixList, &proto.RelicAffix{
						AffixId: subAffix.AffixId,
						Cnt:     subAffix.Cnt,
						Step:    subAffix.Step,
					})
				}
				info.RelicList = append(info.RelicList, displayRelicInfo)
			}
		}
		// add SkilltreeList
		if db.SkilltreeList != nil {
			for _, skill := range db.SkilltreeList {
				if skill.Level == 0 {
					continue
				}
				avatarSkillTree := &proto.AvatarSkillTree{
					PointId: skill.PointId,
					Level:   skill.Level,
				}
				info.SkilltreeList = append(info.SkilltreeList, avatarSkillTree)
			}
		}
		infoList = append(infoList, info)
	}

	return infoList
}
