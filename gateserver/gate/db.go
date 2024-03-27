package gate

import (
	"context"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ctx = context.Background()

type Store struct {
	config  *config.Config
	MysqlDb *gorm.DB
	RedisDb *redis.Client
}

// 登录黑名单
type LoginBlackUid struct {
	AccountId uint32
	BeginTime int64
	EndTime   int64
	Msg       string
}

func (s *Store) init() {
	var err error
	dsn := s.config.MysqlDsn
	s.MysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10 秒钟
	// 初始化表
	err = s.MysqlDb.AutoMigrate(&LoginBlackUid{})
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

func NewRedis(config *config.Config) *redis.Client {
	reconf := config.RedisConf["player_token"]
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

// 使用账号uid拉取数据
func (s *Store) QueryUidPlayerUidByFieldPlayer(AccountId uint32) *LoginBlackUid {
	var uidplayer LoginBlackUid
	s.MysqlDb.Model(&LoginBlackUid{}).Where("account_id = ?", AccountId).First(&uidplayer)
	return &uidplayer
}

/*
// 更新账号
func (s *Store) UpdateUidPlayer(accountId uint, uidPlayer *UidPlayer) error {
	if err := s.MysqlDb.Model(&UidPlayer{}).Where("account_id = ?", accountId).Updates(uidPlayer).Error; err == nil {
		return nil
	} else {
		return err
	}
}
*/

// 根据accounid拉取combotoken

func (s *Store) GetComboTokenByAccountId(accountId string) string {
	key := "player_comboToken:" + accountId
	comboToken, err := s.RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return comboToken
}
