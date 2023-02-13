/**
*@author:yangzhou
*@date: 2023/2/13
*@email: yangzhou2224@shengtian.com
*@description:
 */
package service

import (
	"context"
	"yzgin/internal/admin/model"
)

func (s *Service) CreateAdmin(ctx context.Context, username, password string) error {
	admin := model.Admin{
		Username: username,
		Password: password,
	}

	if err := s.dao.CreateAdmin(ctx, admin); err != nil {
		return err
	}

	return nil
}
