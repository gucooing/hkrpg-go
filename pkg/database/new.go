package database

import (
	"context"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gromlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gromlogger.Default.LogMode(gromlogger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error("mysql connect err:", err)
		panic(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("mysql connect err:", err)
		panic(err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(1000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(1 * time.Second) // 10 秒钟

	return db
}

func NewSqlite(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: gromlogger.Default.LogMode(gromlogger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error("mysql connect err:", err)
		panic(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("mysql connect err:", err)
		panic(err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(1000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(100 * time.Millisecond) // 10 秒钟

	return db
}

var ctx = context.Background()

func NewRedis(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:               "",
		Addr:                  addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              "",
		Password:              password,
		CredentialsProvider:   nil,
		DB:                    db,
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
		panic("redis connect fail" + err.Error())
	}
	return rdb
}
