package db

import (
	"time"

	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (s *Store) init() {
	var err error
	dsn := s.config.MysqlDsn
	s.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error("MySQL数据库连接失败,错误原因:%s", err)
		return
	}
	logger.Info("MySQL数据库连接成功")
	sqlDB, err := s.Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10 秒钟
	// 初始化表
	err = s.Db.AutoMigrate(&Account{}, &UidPlayer{})
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
	return s
}
