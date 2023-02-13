/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package v1

import (
	"yzgin/global"
	"yzgin/internal/admin/service"

	"go.uber.org/zap"

	"yzgin/config"
)

type Route struct {
	svc  *service.Service
	conf *config.Server
	log  *zap.Logger
}

func New(c *config.Server, s *service.Service) *Route {
	return &Route{
		conf: c,
		svc:  s,
		log:  global.Log,
	}
}
