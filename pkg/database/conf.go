package database

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"gorm.io/gorm"
)

// 拉取全部邮件
func GetDbAllMail(db *gorm.DB) []*constant.Mail {
	var mailMap []*constant.Mail
	db.Find(&mailMap)
	return mailMap
}

// 拉取全部模拟宇宙
func GetAllRogue(db *gorm.DB) []*constant.RogueConf {
	var rogueMap []*constant.RogueConf
	db.Find(&rogueMap)
	return rogueMap
}

// 拉取区服配置
func GetRegionConf(db *gorm.DB, regionName string) (*constant.RegionConf, error) {
	var regionConf *constant.RegionConf
	err := db.Model(&constant.RegionConf{}).Where("name = ?", regionName).First(&regionConf).Error
	if err != nil {
		return nil, err
	}
	return regionConf, nil
}

// 拉取全部区服配置
func GetAllRegionConf(db *gorm.DB) []*constant.RegionConf {
	var regionConfList []*constant.RegionConf
	db.Find(&regionConfList)
	return regionConfList
}

// 设置区服数据
func SetRegionConf(db *gorm.DB, regionConf *constant.RegionConf) error {
	err := db.Model(&constant.RegionConf{}).Where("name = ?", regionConf.Name).Updates(regionConf).Error
	return err
}
