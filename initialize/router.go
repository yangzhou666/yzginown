/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */
package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yzgin/global"
	"yzgin/middleware"
	"yzgin/model/common/response"
	"yzgin/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System

	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.Log.Info("use middleware cors")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("/api")
	{
		// 健康监测
		PublicGroup.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		// 获取系统时间
		PublicGroup.GET("/getCurrentTimes", func(c *gin.Context) {
			now := time.Now().Local().Unix()
			response.OkWithDetailed(now, "查询成功", c)
		})

		systemRouter.InitPublicRouter(PublicGroup)
	}

	global.Log.Info("router register success")
	return Router
}
