package database

import (
	"gorm.io/gorm"
)

// 查询账号
func QueryAccountByFieldUsername(db *gorm.DB, Username string) *Account {
	var account Account
	db.Model(&Account{}).Where("Username = ?", Username).First(&account)
	return &account
}
func QueryAccountByFieldAccountId(db *gorm.DB, AccountId uint32) *Account {
	var account Account
	db.Model(&Account{}).Where("account_id = ?", AccountId).First(&account)
	return &account
}

// 添加新账号
func AddAccountFieldByFieldName(db *gorm.DB, account *Account) (uint32, error) {
	if err := db.Create(account).Error; err == nil {
		return account.AccountId, nil
	} else {
		return 0, err
	}
}

// 更新账号
func UpdateAccountFieldByFieldName(db *gorm.DB, account *Account) error {
	if account.AccountId == 0 {
		return nil
	}
	if err := db.Model(&Account{}).Where("account_id = ?", account.AccountId).Updates(account).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 使用account id拉取数据
func GetPlayerUidByAccountId(db *gorm.DB, AccountId uint32) *PlayerUid {
	var playerUid *PlayerUid
	db.Model(&PlayerUid{}).Where("account_id = ?", AccountId).First(&playerUid)
	if playerUid.Uid == 0 {
		playerUid = UpdatePlayerUid(db, AccountId)
		return playerUid
	}
	return playerUid
}

// 指定account id 创建数据
func UpdatePlayerUid(db *gorm.DB, AccountId uint32) *PlayerUid {
	playerUid := new(PlayerUid)
	playerUid.AccountId = AccountId
	db.Select("account_id", AccountId).Create(&playerUid)

	return playerUid
}

// 使用账号id拉取数据
func QueryAccountUidByFieldPlayer(db *gorm.DB, uid uint32) *PlayerData {
	var playerData *PlayerData
	db.Model(&PlayerData{}).Where("uid = ?", uid).First(&playerData)
	return playerData
}

// 添加新账号数据
func AddDatePlayerFieldByFieldName(db *gorm.DB, player *PlayerData) error {
	if err := db.Create(player).Error; err != nil {
		return err
	}
	return nil
}

// 更新账号
func UpdatePlayer(db *gorm.DB, player *PlayerData) error {
	if player.Uid == 0 {
		return nil
	}
	if err := db.Model(&PlayerData{}).Where("uid = ?", player.Uid).Updates(player).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 更新账号简要数据
func UpdatePlayerBasic(db *gorm.DB, player *PlayerBasic) bool {
	if player.Uid == 0 {
		return false
	}
	if err := db.Save(player).Error; err == nil {
		return true
	} else {
		return false
	}
}

// 拉取全部邮件
func GetDbAllMail(db *gorm.DB) []*Mail {
	var mailMap []*Mail
	db.Find(&mailMap)
	return mailMap
}

// 拉取全部模拟宇宙
func GetAllRogue(db *gorm.DB) []*RogueConf {
	var rogueMap []*RogueConf
	db.Find(&rogueMap)
	return rogueMap
}

// 拉取地图文件
func GetBlockData(db *gorm.DB, uid, entryId uint32) *BlockData {
	var blockData *BlockData
	db.Where(&BlockData{Uid: uid, EntryId: entryId}).First(&blockData)
	return blockData
}

// 更新地图文件
func UpdateBlockData(db *gorm.DB, blockData *BlockData) error {
	if blockData.Uid == 0 {
		return nil
	}
	if err := db.Save(blockData).Error; err == nil {
		return nil
	} else {
		return err
	}
}
