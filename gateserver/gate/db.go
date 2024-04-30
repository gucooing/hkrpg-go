package gate

import (
	"context"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ctx = context.Background()

type Store struct {
	config         *config.Config
	PlayerUidMysql *gorm.DB
	LoginRedis     *redis.Client
	StatusRedis    *redis.Client
}

/*********************************************Mysql****************************************/

type PlayerUid struct {
	Uid          uint32 `gorm:"primarykey;AUTO_INCREMENT"`
	AccountType  uint32
	AccountId    uint32
	CreateTime   int64
	IsBan        bool
	BanBeginTime int64
	BanEndTime   int64
	BanMsg       string
}

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	playerUidMysqlConf := config.MysqlConf["player_uid"]
	s.PlayerUidMysql = database.NewMysql(playerUidMysqlConf.Dsn)
	s.PlayerUidMysql.AutoMigrate(&PlayerUid{})

	redisLoginConf := config.RedisConf["player_login"]
	s.LoginRedis = database.NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)
	redisStatusConf := config.RedisConf["player_status"]
	s.StatusRedis = database.NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)

	logger.Info("数据库连接成功")
	return s
}

// 使用account id拉取数据
func (s *Store) GetPlayerUidByAccountId(AccountId uint32) *PlayerUid {
	var playerUid *PlayerUid
	s.PlayerUidMysql.Model(&PlayerUid{}).Where("account_id = ?", AccountId).First(&playerUid)
	if playerUid.Uid == 0 {
		playerUid = s.UpdatePlayerUid(AccountId)
		return playerUid
	}
	return playerUid
}

// 指定account id 创建数据
func (s *Store) UpdatePlayerUid(AccountId uint32) *PlayerUid {
	playerUid := new(PlayerUid)
	playerUid.AccountId = AccountId
	s.PlayerUidMysql.Select("account_id", AccountId).Create(&playerUid)

	return playerUid
}

/*********************************************Redis****************************************/

// 根据accounid拉取combotoken
func (s *Store) GetComboTokenByAccountId(accountId string) string {
	key := "player_comboToken:" + accountId
	comboToken, err := s.LoginRedis.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return comboToken
}

const (
	MaxLockAliveTime  = 10000 // 单个锁的最大存活时间 毫秒
	LockRetryWaitTime = 50    // 同步加锁重试间隔时间 毫秒
	MaxLockRetryTimes = 2     // 同步加锁最大重试次数
)

// DistLockSync 加锁同步阻塞直到成功或超时
func (s *Store) DistLockSync(accountId string) bool {
	var result = false
	for i := 0; i < MaxLockRetryTimes; i++ {
		var err error = nil
		key := "player_login_lock:" + accountId
		result, err = s.LoginRedis.SetNX(context.TODO(),
			key,
			time.Now().UnixMilli(),
			time.Millisecond*time.Duration(MaxLockAliveTime)).Result()
		if err != nil {
			logger.Error("redis lock setnx error: %v", err)
			return false
		}
		if result == true {
			break
		}
		time.Sleep(time.Millisecond * time.Duration(LockRetryWaitTime))
	}
	return result
}

// DistUnlock 解锁
func (s *Store) DistUnlock(accountId string) {
	var result int64 = 0
	var err error = nil
	key := "player_login_lock:" + accountId
	result, err = s.LoginRedis.Del(context.TODO(), key).Result()
	if err != nil {
		logger.Error("redis lock del error: %v", err)
		return
	}
	if result == 0 {
		logger.Error("redis lock del result is fail")
		return
	}
}

// 获取玩家状态
func (s *Store) GetPlayerStatus(accountId string) ([]byte, bool) {
	key := "player_status_lock:" + accountId
	bin, err := s.StatusRedis.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 删除玩家状态
func (s *Store) DistUnlockPlayerStatus(accountId string) {
	var result int64 = 0
	var err error = nil
	key := "player_status_lock:" + accountId
	result, err = s.StatusRedis.Del(context.TODO(), key).Result()
	if err != nil {
		logger.Error("redis lock del error: %v", err)
		return
	}
	if result == 0 {
		logger.Error("redis lock del result is fail")
		return
	}
}
