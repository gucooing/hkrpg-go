package database

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 使用账号id拉取数据
func GetPlayerDataByUid(db *gorm.DB, peDb *gorm.DB, uid uint32) *constant.PlayerData {
	if db != nil {
		return queryAccountUidByFieldPlayer(db, uid)
	}
	if peDb != nil {
		return queryAccountUidByFieldPlayer(peDb, uid)
	}
	return nil
}

func queryAccountUidByFieldPlayer(db *gorm.DB, uid uint32) *constant.PlayerData {
	var playerData *constant.PlayerData
	db.Model(&constant.PlayerData{}).Where("uid = ?", uid).First(&playerData)
	return playerData
}

// 添加新账号数据
func AddPlayerDataByUid(db *gorm.DB, peDb *gorm.DB, player *constant.PlayerData) error {
	if db != nil {
		return addDatePlayerFieldByFieldName(db, player)
	}
	if peDb != nil {
		return addDatePlayerFieldByFieldName(peDb, player)
	}
	return nil
}

func addDatePlayerFieldByFieldName(db *gorm.DB, player *constant.PlayerData) error {
	if err := db.Create(player).Error; err != nil {
		return err
	}
	return nil
}

// 更新账号

func UpPlayerDataByUid(db *gorm.DB, peDb *gorm.DB, player *constant.PlayerData) error {
	if db != nil {
		return updatePlayer(db, player)
	}
	if peDb != nil {
		return updatePlayer(peDb, player)
	}
	return nil
}

func updatePlayer(db *gorm.DB, player *constant.PlayerData) error {
	if player.Uid == 0 {
		return nil
	}
	if err := db.Model(&constant.PlayerData{}).Where("uid = ?", player.Uid).Updates(player).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 拉取地图文件
func GetBlockData(db *gorm.DB, dbPe *gorm.DB, uid, entryId uint32) *constant.BlockData {
	if db != nil {
		return getBlockData(db, uid, entryId)
	}
	if dbPe != nil {
		return getBlockData(dbPe, uid, entryId)
	}
	return nil
}

func getBlockData(db *gorm.DB, uid, entryId uint32) *constant.BlockData {
	var blockData *constant.BlockData
	db.Where(&constant.BlockData{Uid: uid, EntryId: entryId}).First(&blockData)
	return blockData
}

// 更新地图文件
func UpdateBlockData(db *gorm.DB, dbPe *gorm.DB, blockData *constant.BlockData) error {
	if db != nil {
		return updateBlockData(db, blockData)
	}
	if dbPe != nil {
		return updateBlockData(dbPe, blockData)
	}
	return nil
}

func updateBlockData(db *gorm.DB, blockData *constant.BlockData) error {
	if blockData.Uid == 0 {
		return nil
	}
	if err := db.Save(blockData).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 更新账号简要数据
func UpdatePlayerBasic(rc *redis.Client, db *gorm.DB, basic *constant.PlayerBasic) bool {
	if basic.Uid == 0 {
		return false
	}
	if rc != nil {
		return updatePlayerBasicRedis(rc, basic.Uid, basic.BinData)
	}
	if db != nil {
		return updatePlayerBasicMysql(db, basic)
	}
	return false
}

func updatePlayerBasicMysql(db *gorm.DB, player *constant.PlayerBasic) bool {
	if player.Uid == 0 {
		return false
	}
	if err := db.Save(player).Error; err == nil {
		return true
	} else {
		return false
	}
}

// 获取账号简要数据
func GetPlayerBasic(rc *redis.Client, db *gorm.DB, uid uint32) ([]byte, bool) {
	if rc != nil {
		return getPlayerBasicRedis(rc, uid)
	}
	if db != nil {
		playerBasic := getPlayerBasicMysql(db, uid)
		if playerBasic.Uid == uid {
			return playerBasic.BinData, true
		}
	}
	return nil, false
}

func getPlayerBasicMysql(db *gorm.DB, uid uint32) *constant.PlayerBasic {
	var playerBasic *constant.PlayerBasic
	db.Where(&constant.PlayerBasic{Uid: uid}).First(&playerBasic)
	return playerBasic
}

// 获取好友申请
func GetApplyFriend(rc *redis.Client, db *gorm.DB, uid uint32) ([]byte, bool) {
	if rc != nil {
		return getPlayerFriendRedis(rc, uid)
	}
	if db != nil {
		applyFriend := getApplyFriendMysql(db, uid)
		if applyFriend.Uid == uid {
			return applyFriend.ReceiveApply, true
		}
	}
	return nil, false
}

func getApplyFriendMysql(db *gorm.DB, uid uint32) *constant.ApplyFriend {
	var applyFriend *constant.ApplyFriend
	db.Where(&constant.ApplyFriend{Uid: uid}).First(&applyFriend)
	return applyFriend
}

// 获取待加入好友
func GetAcceptApplyFriend(rc *redis.Client, db *gorm.DB, uid uint32) ([]byte, bool) {
	if rc != nil {
		return getAcceptApplyFriendRedis(rc, uid)
	}
	if db != nil {
		acceptApplyFriend := getAcceptApplyFriendMysql(db, uid)
		if acceptApplyFriend.Uid == uid {
			return acceptApplyFriend.AcceptApplyFriend, true
		}
	}
	return nil, false
}

func getAcceptApplyFriendMysql(db *gorm.DB, uid uint32) *constant.AcceptApplyFriend {
	var acceptApplyFriend *constant.AcceptApplyFriend
	db.Where(&constant.AcceptApplyFriend{Uid: uid}).First(&acceptApplyFriend)
	return acceptApplyFriend
}

// 删除待加入好友
func DelAcceptApplyFriend(rc *redis.Client, db *gorm.DB, uid uint32) {
	if rc != nil {
		delAcceptApplyFriendRedis(rc, uid)
	}
	if db != nil {
		delAcceptApplyFriendMysql(db, uid)
	}
}

func delAcceptApplyFriendMysql(db *gorm.DB, uid uint32) {
	db.Delete(&constant.AcceptApplyFriend{Uid: uid})
}
