/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package migrate

import (
	"fmt"
	"yzgin/global"
	"yzgin/internal/admin/model"

	"gorm.io/gorm"
)

// MigrateInit
func MigrateInit(DBCli *gorm.DB) {
	var Admin *model.Admin
	if !DBCli.Migrator().HasTable(Admin) {
		tableAttr := fmt.Sprintf("ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1 COMMENT='%s'", Admin.GetTableComment())
		DBCli.Set("gorm:table_options", tableAttr).AutoMigrate(Admin)
	}

	global.Log.Info("迁移运行成功")
}
