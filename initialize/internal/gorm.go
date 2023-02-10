/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package internal

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"yzgin/global"
)

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	config := &gorm.Config{
		SkipDefaultTransaction: true,  // 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行
		PrepareStmt:            false, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular, //严格匹配表名 默认是复数形式
		},
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DBBASE
	switch global.Config.System.DbType {
	case "mysql":
		logMode = &global.Config.Mysql
	case "pgsql":
		logMode = &global.Config.Pgsql
	case "oracle":
		logMode = &global.Config.Oracle
	default:
		logMode = &global.Config.Mysql
	}

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
