/**
*@author:yangzhou
*@date: 2023/2/13
*@email: yangzhou2224@shengtian.com
*@description:
 */
package v1

import (
	"yzgin/global"
	"yzgin/internal/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *Route) GetServerInfo(c *gin.Context) {
	server, err := r.svc.GetServerInfo()
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
}
