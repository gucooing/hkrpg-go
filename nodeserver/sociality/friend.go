package sociality

import (
	"sync"

	nodb "github.com/gucooing/hkrpg-go/nodeserver/db"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

var FRIEND *Friend

// 避免覆写的情况，缓存一下
type Friend struct {
	ApplyFriend           map[uint32]*spb.ApplyFriend
	applyFriendSync       sync.RWMutex                      // 同步锁
	AcceptApplyFriend     map[uint32]*spb.AcceptApplyFriend // 待加入数据库的好友
	acceptApplyFriendSync sync.RWMutex                      // 同步锁
}

func newFriend() *Friend {
	return &Friend{
		ApplyFriend:       make(map[uint32]*spb.ApplyFriend),
		AcceptApplyFriend: make(map[uint32]*spb.AcceptApplyFriend),
	}
}

func getFriend() *Friend {
	if FRIEND == nil {
		FRIEND = newFriend()
	}
	return FRIEND
}

func GetApplyFriendByUid(uid uint32) *spb.ApplyFriend {
	friend := new(spb.ApplyFriend)
	db := getFriend()
	db.applyFriendSync.Lock()
	defer db.applyFriendSync.Unlock()
	if db.ApplyFriend == nil {
		db.ApplyFriend = make(map[uint32]*spb.ApplyFriend)
	}
	if db.ApplyFriend[uid] == nil {
		bin, ok := database.GetAcceptApplyFriend(nodb.GetStore().PlayerBriefDataRedis,
			nil, uid)
		if !ok {
			friend = &spb.ApplyFriend{
				Uid:             uid,
				RecvApplyFriend: make(map[uint32]*spb.ReceiveApply),
			}
		} else {
			pb.Unmarshal(bin, friend)
		}
		db.ApplyFriend[uid] = friend
	}

	return db.ApplyFriend[uid]
}

func UpdateApplyFriend(uid uint32) bool {
	db := getFriend()
	db.applyFriendSync.Lock()
	defer db.applyFriendSync.Unlock()
	if db.ApplyFriend[uid] == nil {
		return false
	}
	ubin, err := pb.Marshal(db.ApplyFriend[uid])
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return false
	}
	database.SetPlayerFriend(nodb.GetStore().PlayerBriefDataRedis, uid, ubin)
	delete(db.ApplyFriend, uid)
	return true
}

func GetAcceptApplyFriendByUid(uid uint32) *spb.AcceptApplyFriend {
	friend := new(spb.AcceptApplyFriend)
	db := getFriend()
	db.acceptApplyFriendSync.Lock()
	defer db.acceptApplyFriendSync.Unlock()
	if db.AcceptApplyFriend == nil {
		db.AcceptApplyFriend = make(map[uint32]*spb.AcceptApplyFriend)
	}
	if db.AcceptApplyFriend[uid] == nil {
		bin, ok := database.GetAcceptApplyFriend(nodb.GetStore().PlayerBriefDataRedis,
			nil, uid)
		if !ok {
			friend = &spb.AcceptApplyFriend{
				Uid:             uid,
				RecvApplyFriend: make(map[uint32]*spb.ReceiveApply),
			}
		} else {
			pb.Unmarshal(bin, friend)
		}
		db.AcceptApplyFriend[uid] = friend
	}

	return db.AcceptApplyFriend[uid]
}

func UpdateAcceptApplyFriend(uid uint32) bool {
	db := getFriend()
	db.acceptApplyFriendSync.Lock()
	defer db.acceptApplyFriendSync.Unlock()
	if db.AcceptApplyFriend[uid] == nil {
		return false
	}
	ubin, err := pb.Marshal(db.AcceptApplyFriend[uid])
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return false
	}
	database.SetAcceptApplyFriend(nodb.GetStore().PlayerBriefDataRedis, uid, ubin)
	delete(db.AcceptApplyFriend, uid)
	return true
}
