/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package service

import (
	"yzgin/global"
	"yzgin/internal/admin/dao"

	"yzgin/config"

	"go.uber.org/zap"
)

// Service struct
type Service struct {
	c   *config.Server
	dao *dao.Dao
}

var (
	svc *Service = nil
)

// New init
func New(c *config.Server) (s *Service) {
	if svc != nil {
		global.Log.Info("admin", zap.String("down", "admin 服务层 复用"))
		return svc
	}
	db := dao.New(c)

	svc = &Service{
		c:   c,
		dao: db,
	}

	return svc
}

// Close Service.
func (s *Service) Close() {
	s.dao.Close()
}
