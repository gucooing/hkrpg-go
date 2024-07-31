package hkrpg_go_pe

import (
	"github.com/gucooing/hkrpg-go/dispatch"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func newStorePE(cfg *Config) *dispatch.Store {
	s := new(dispatch.Store)
	s.HkrpgGoPe = database.NewSqlite(cfg.SqlPath)
	s.HkrpgGoPe.AutoMigrate(
		&database.Account{},      // sdk账户
		&database.PlayerUid{},    // 映射表
		&database.PlayerData{},   // 玩家数据
		&database.BlockData{},    // 地图数据
		&database.RogueConf{},    // 模拟宇宙配置
		&database.ScheduleConf{}, // 忘记了
		&database.PlayerBasic{},  // 好友简要信息
		&database.Mail{},         // 邮件配置
	)

	logger.Info("数据库连接成功")
	return s
}
