package gate

import (
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/gateserver/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Store struct {
	config *config.Config
	Db     *gorm.DB
}

type UidPlayer struct {
	AccountUid uint `gorm:"primarykey;AUTO_INCREMENT"`
	AccountId  uint
	IsBan      bool
	ComboToken string
}

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
	err = s.Db.AutoMigrate(&UidPlayer{})
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

// 使用账号uid拉取数据
func (s *Store) QueryUidPlayerUidByFieldPlayer(AccountId uint32) *UidPlayer {
	var uidplayer UidPlayer
	s.Db.Model(&UidPlayer{}).Where("account_id = ?", AccountId).First(&uidplayer)
	return &uidplayer
}

// 更新账号
func (s *Store) UpdateUidPlayer(accountId uint, uidPlayer *UidPlayer) error {
	if err := s.Db.Model(&UidPlayer{}).Where("account_id = ?", accountId).Updates(uidPlayer).Error; err == nil {
		return nil
	} else {
		return err
	}
}
