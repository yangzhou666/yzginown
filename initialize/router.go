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
	"yzgin/internal/common/middleware"
	"yzgin/internal/common/response"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("/v1").Use(middleware.Cors())
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

	}

	global.Log.Info("router register success")
	return Router
}
