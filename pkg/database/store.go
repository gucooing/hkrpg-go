package database

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

/******************************disaptch*******************************/

type DisaptchStore struct {
	AccountMysql *gorm.DB
	LoginRedis   *redis.Client
}

func NewDisaptchStore(mysqlList map[string]constant.MysqlConf, redisList map[string]constant.RedisConf) *DisaptchStore {
	s := &DisaptchStore{}
	accountMysqlConf := mysqlList["account"]
	s.AccountMysql = NewMysql(accountMysqlConf.Dsn)
	s.AccountMysql.AutoMigrate(&constant.Account{})

	redisLoginConf := redisList["player_login"]
	s.LoginRedis = NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)

	logger.Info("数据库连接成功")
	return s
}

/******************************gateserver*******************************/

type GateStore struct {
	PlayerUidMysql *gorm.DB
	LoginRedis     *redis.Client
	StatusRedis    *redis.Client
}

func NewGateStore(mysqlList map[string]constant.MysqlConf, redisList map[string]constant.RedisConf) *GateStore {
	s := &GateStore{}
	playerUidMysqlConf := mysqlList["user"]
	s.PlayerUidMysql = NewMysql(playerUidMysqlConf.Dsn)
	s.PlayerUidMysql.AutoMigrate(&constant.PlayerUid{})

	redisLoginConf := redisList["player_login"]
	s.LoginRedis = NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)
	redisStatusConf := redisList["player_status"]
	s.StatusRedis = NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)

	logger.Info("数据库连接成功")
	return s
}

/******************************gameserver*******************************/

var GSS *GameStore

type GameStore struct {
	PlayerDataMysql      *gorm.DB
	ServerConf           *gorm.DB
	LoginRedis           *redis.Client
	StatusRedis          *redis.Client
	PlayerBriefDataRedis *redis.Client // 玩家简要信息
	// pe
	PeMysql *gorm.DB
}

func NewGameStore(mysqlList map[string]constant.MysqlConf, redisList map[string]constant.RedisConf) *GameStore {
	s := &GameStore{}
	GSS = s
	mysqlPlayerDataConf := mysqlList["player"]
	s.PlayerDataMysql = NewMysql(mysqlPlayerDataConf.Dsn)
	s.PlayerDataMysql.AutoMigrate(&constant.PlayerData{}, &constant.BlockData{})
	mysqlServerConf := mysqlList["conf"]
	s.ServerConf = NewMysql(mysqlServerConf.Dsn)
	s.ServerConf.AutoMigrate(&constant.Mail{}, &constant.RogueConf{}, &constant.ScheduleConf{})

	redisLoginConf := redisList["player_login"]
	s.LoginRedis = NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)
	redisStatusConf := redisList["player_status"]
	s.StatusRedis = NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)
	playerBriefDataRedis := redisList["player_brief_data"]
	s.PlayerBriefDataRedis = NewRedis(playerBriefDataRedis.Addr, playerBriefDataRedis.Password, playerBriefDataRedis.DB)

	logger.Info("数据库连接成功")
	GetDbConf(s.ServerConf) // 初始化数据库配置表
	return s
}
