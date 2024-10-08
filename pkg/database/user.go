package database

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 查询账号
func QueryAccountByFieldUsername(db *gorm.DB, Username string) *constant.Account {
	var account constant.Account
	db.Model(&constant.Account{}).Where("Username = ?", Username).First(&account)
	return &account
}

func GetAccountByFieldAccountId(db *gorm.DB, AccountId uint32) *constant.Account {
	var account constant.Account
	db.Model(&constant.Account{}).Where("account_id = ?", AccountId).First(&account)
	return &account
}

// 添加新账号
func AddAccountFieldByFieldName(db *gorm.DB, account *constant.Account) (uint32, error) {
	if err := db.Create(account).Error; err == nil {
		return account.AccountId, nil
	} else {
		return 0, err
	}
}

// 更新账号
func UpdateAccountFieldByFieldName(db *gorm.DB, account *constant.Account) error {
	if account.AccountId == 0 {
		return nil
	}
	if err := db.Model(&constant.Account{}).Where("account_id = ?", account.AccountId).Updates(account).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 使用account id拉取数据
func GetPlayerUidByAccountId(db *gorm.DB, AccountId uint32) *constant.PlayerUid {
	var playerUid *constant.PlayerUid
	db.Model(&constant.PlayerUid{}).Where("account_id = ?", AccountId).First(&playerUid)
	if playerUid.Uid == 0 {
		playerUid = AddPlayerUidByAccountId(db, AccountId)
		return playerUid
	}
	return playerUid
}

// 指定account id 创建数据
func AddPlayerUidByAccountId(db *gorm.DB, AccountId uint32) *constant.PlayerUid {
	playerUid := new(constant.PlayerUid)
	playerUid.AccountId = AccountId
	db.Select("account_id", AccountId).Create(&playerUid)

	return playerUid
}

func UpdatePlayerUid(db *gorm.DB, playerUid *constant.PlayerUid) error {
	if playerUid == nil {
		return nil
	}
	if err := db.Model(&constant.PlayerUid{}).Where("account_id = ?", playerUid.AccountId).Updates(playerUid).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 获取ComboToken Redis
func GetComboTokenByAccountIdRedis(rc *redis.Client, accountId string) string {
	key := "player_comboToken:" + accountId
	comboToken, _ := rc.Get(ctx, key).Result()
	return comboToken
}

func UpComboTokenByAccountId(rc *redis.Client, db *gorm.DB, accountId uint32, comboToken string) {
	if rc != nil {
		setComboTokenByAccountId(rc, strconv.Itoa(int(accountId)), comboToken)
	}
	if db != nil {
		upComboTokenByAccountIdGorm(db, accountId, comboToken)
	}
}

// 更新ComboToken Gorm
func upComboTokenByAccountIdGorm(db *gorm.DB, accountId uint32, comboToken string) {
	p := GetPlayerUidByAccountId(db, accountId)
	if p.AccountId != accountId {
		return
	}
	p.ComboToken = comboToken
	UpdatePlayerUid(db, p)
}

// 更新ComboToken Redis
func setComboTokenByAccountId(rc *redis.Client, accountId, comboToken string) {
	key := "player_comboToken:" + accountId
	err := rc.Set(ctx, key, comboToken, 168*time.Hour).Err()
	if err != nil {
		return
	}
	return
}
