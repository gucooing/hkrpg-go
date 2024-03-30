package gate

import (
	"context"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gromlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var ctx = context.Background()

type Store struct {
	config  *config.Config
	MysqlDb *gorm.DB
	RedisDb *redis.Client
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

func (s *Store) init() {
	var err error
	dsn := s.config.MysqlDsn
	s.MysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gromlogger.Default.LogMode(gromlogger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error("MySQL数据库连接失败,错误原因:%s", err)
		return
	}
	logger.Info("MySQL数据库连接成功")
	sqlDB, err := s.MysqlDb.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(5)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10 秒钟
	// 初始化表
	err = s.MysqlDb.AutoMigrate(&PlayerUid{})
	if err != nil {
		logger.Error("MySQL数据库初始化失败")
		return
	}
	logger.Info("MySQL数据库初始化成功")
}

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	s.init()
	s.RedisDb = NewRedis(config) // 初始化redis
	return s
}

// 使用account id拉取数据
func (s *Store) GetPlayerUidByAccountId(AccountId uint32) *PlayerUid {
	var playerUid *PlayerUid
	s.MysqlDb.Model(&PlayerUid{}).Where("account_id = ?", AccountId).First(&playerUid)
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
	s.MysqlDb.Select("account_id", AccountId).Create(&playerUid)

	return playerUid
}

/*********************************************Redis****************************************/

func NewRedis(config *config.Config) *redis.Client {
	reconf := config.RedisConf["player_login"]
	rdb := redis.NewClient(&redis.Options{
		Network:               "",
		Addr:                  reconf.Addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              reconf.Password,
		CredentialsProvider:   nil,
		DB:                    reconf.DB,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		MaxActiveConns:        0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
		DisableIndentity:      false,
		IdentitySuffix:        "",
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err.Error())
	}
	logger.Info("Redis数据库连接成功")
	return rdb
}

// 根据accounid拉取combotoken
func (s *Store) GetComboTokenByAccountId(accountId string) string {
	key := "player_comboToken:" + accountId
	comboToken, err := s.RedisDb.Get(ctx, key).Result()
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
		result, err = s.RedisDb.SetNX(context.TODO(),
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
	result, err = s.RedisDb.Del(context.TODO(), key).Result()
	if err != nil {
		logger.Error("redis lock del error: %v", err)
		return
	}
	if result == 0 {
		logger.Error("redis lock del result is fail")
		return
	}
}

// 标记玩家状态
func (s *Store) SetPlayerStatus(accountId string, value []byte) error {
	key := "player_status:" + accountId
	err := s.RedisDb.Set(ctx, key, value, 0).Err()
	return err
}

// 获取玩家状态
func (s *Store) GetPlayerStatus(accountId string) ([]byte, bool) {
	key := "player_status:" + accountId
	bin, err := s.RedisDb.Get(ctx, key).Bytes()
	if err == nil {
		return bin, true
	} else if err == redis.Nil {
		return bin, false
	} else {
		return bin, false
	}
}

// 删除玩家状态
func (s *Store) DelPlayerStatus(accountId string) error {
	key := "player_status:" + accountId
	err := s.RedisDb.Del(ctx, key).Err()
	return err
}
