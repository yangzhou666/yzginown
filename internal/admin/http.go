/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package admin

import (
	"yzgin/global"
	v1 "yzgin/internal/admin/api/v1"
	"yzgin/internal/admin/service"
	"yzgin/internal/common/middleware"

	"github.com/gin-gonic/gin"
)

func Init(routerEngine *gin.Engine) *service.Service {
	//初始化service
	s := service.New(&global.Config)

	//初始化api
	r := v1.New(&global.Config, s)

	//router  相关路由
	adminRouterGroup := routerEngine.Group("v1/admin")
	adminRouterGroup.Use(middleware.Cors())
	{
		systemRouterGroup := adminRouterGroup.Group("system")
		systemRouterGroup.GET("captcha", r.Captcha) //获取系统验证码

		userRouterGroup := adminRouterGroup.Group("user")
		{
			userRouterGroup.POST("", r.CreateAdmin)
		}
	}

	return s
}
