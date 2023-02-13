/**
*@author:yangzhou
*@date: 2023/2/13
*@email: yangzhou2224@shengtian.com
*@description:
 */
package dao

import (
	"context"
	"yzgin/internal/admin/model"

	"go.uber.org/zap"
)

func (d *Dao) CreateAdmin(ctx context.Context, value interface{}) error {
	if err := d.Create(ctx, &model.Admin{}, value); err != nil {
		d.log.Error("创建管理员失败：", zap.Error(err))
		return err
	}

	return nil
}
