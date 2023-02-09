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
	"yzgin/global"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	global.Log.Info("router register success")
	return Router
}
