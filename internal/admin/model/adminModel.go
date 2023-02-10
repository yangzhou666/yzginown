/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package model

import (
	"fmt"
	"yzgin/internal/common/model"
)

// Admin 管理员用户表
type Admin struct {
	model.SoftDeleteModel
	Username string `gorm:"column:username;type:varchar(125);comment:管理员用户名;NOT NULL" json:"username"`
	Password string `gorm:"column:password;type:char(32);comment:密码;NOT NULL" json:"password"`
}

func (u *Admin) TableName() string {
	return "admin"
}

func (u *Admin) GetTableComment() string {
	return "后台管理员"
}

func (u *Admin) GetID() string {
	return fmt.Sprintf("%d", u.SoftDeleteModel.Id)
}
