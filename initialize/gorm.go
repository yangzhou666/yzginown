/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"yzgin/global"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
}
