package database

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

/******************************disaptch*******************************/

var DISPATCH *DisaptchStore

type DisaptchStore struct {
	AccountMysql *gorm.DB
	LoginRedis   *redis.Client
}

func NewDisaptchStore(mysqlList map[string]constant.MysqlConf, redisList map[string]constant.RedisConf) *DisaptchStore {
	DISPATCH = &DisaptchStore{}
	accountMysqlConf := mysqlList["account"]
	DISPATCH.AccountMysql = NewMysql(accountMysqlConf.Dsn)
	DISPATCH.AccountMysql.AutoMigrate(&constant.Account{})

	redisLoginConf := redisList["player_login"]
	DISPATCH.LoginRedis = NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)

	logger.Info("数据库连接成功")
	return DISPATCH
}

/******************************gateserver*******************************/

var GATE *GateStore

type GateStore struct {
	PlayerUidMysql *gorm.DB
	LoginRedis     *redis.Client
	StatusRedis    *redis.Client
}

func NewGateStore(mysqlList map[string]constant.MysqlConf, redisList map[string]constant.RedisConf) {
	GATE = &GateStore{}
	playerUidMysqlConf := mysqlList["user"]
	GATE.PlayerUidMysql = NewMysql(playerUidMysqlConf.Dsn)
	GATE.PlayerUidMysql.AutoMigrate(&constant.PlayerUid{})

	redisLoginConf := redisList["player_login"]
	GATE.LoginRedis = NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)
	redisStatusConf := redisList["player_status"]
	GATE.StatusRedis = NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)

	logger.Info("数据库连接成功")
}

/******************************gameserver*******************************/

var GSS *GameStore

type GameStore struct {
	PlayerDataMysql      *gorm.DB      // 玩家数据
	ServerConf           *gorm.DB      // 服务配置
	LoginRedis           *redis.Client // 登录锁
	StatusRedis          *redis.Client // 状态锁
	PlayerBriefDataRedis *redis.Client // 玩家简要信息
	PlayerMail           *redis.Client // 玩家邮件
}

func NewGameStore(mysqlList map[string]constant.MysqlConf, redisList map[string]constant.RedisConf) {
	s := &GameStore{}
	GSS = s
	mysqlPlayerDataConf := mysqlList["player"]
	s.PlayerDataMysql = NewMysql(mysqlPlayerDataConf.Dsn)
	s.PlayerDataMysql.AutoMigrate(&constant.PlayerData{},
		&constant.BlockData{},
		&constant.ApplyFriend{},
		&constant.AcceptApplyFriend{})
	mysqlServerConf := mysqlList["conf"]
	s.ServerConf = NewMysql(mysqlServerConf.Dsn)
	s.ServerConf.AutoMigrate(&constant.Mail{}, &constant.RogueConf{}, &constant.ScheduleConf{})

	redisLoginConf := redisList["player_login"]
	s.LoginRedis = NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)
	redisStatusConf := redisList["player_status"]
	s.StatusRedis = NewRedis(redisStatusConf.Addr, redisStatusConf.Password, redisStatusConf.DB)
	playerBriefDataRedis := redisList["player_brief_data"]
	s.PlayerBriefDataRedis = NewRedis(playerBriefDataRedis.Addr, playerBriefDataRedis.Password, playerBriefDataRedis.DB)
	playerMail := redisList["player_mail"]
	s.PlayerMail = NewRedis(playerMail.Addr, playerMail.Password, playerMail.DB)

	logger.Info("数据库连接成功")
}

/******************************node*******************************/

var NODE *nodeStore

type nodeStore struct {
	ServerConf *gorm.DB
}

func NewNodeStore(mysqlList map[string]constant.MysqlConf, redisList map[string]constant.RedisConf) {
	NODE = &nodeStore{}
	mysqlServerConf := mysqlList["conf"]
	NODE.ServerConf = NewMysql(mysqlServerConf.Dsn)
	NODE.ServerConf.AutoMigrate(&constant.Mail{}, &constant.RegionConf{})
	logger.Info("数据库连接成功")
}

/******************************pe*******************************/

var PE *gorm.DB

func NewPE(dsn string) {
	PE = NewSqlite(dsn)
	PE.AutoMigrate(
		&constant.Account{},      // sdk账户
		&constant.PlayerUid{},    // 映射表
		&constant.PlayerData{},   // 玩家数据
		&constant.BlockData{},    // 地图数据
		&constant.RogueConf{},    // 模拟宇宙配置
		&constant.ScheduleConf{}, // 忘记了
		&constant.ApplyFriend{},  // 好友申请
		&constant.PlayerBasic{},  // 好友简要信息
		&constant.Mail{},         // 邮件配置
		&constant.PlayerMail{},   // 玩家邮件配置
		&constant.RegionConf{},   // 区服配置
	)
	logger.Info("数据库连接成功")
}
