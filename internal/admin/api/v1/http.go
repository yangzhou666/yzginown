/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package v1

import (
	"yzgin/internal/admin/service"

	"yzgin/config"
)

type Route struct {
	svc  *service.Service
	conf *config.Server
}

func New(c *config.Server, s *service.Service) *Route {
	return &Route{
		conf: c,
		svc:  s,
	}
}
