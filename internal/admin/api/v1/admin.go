/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package v1

import (
	"yzgin/internal/common/response"

	"yzgin/internal/admin/model/requset"

	"github.com/gin-gonic/gin"
)

func (r *Route) CreateAdmin(c *gin.Context) {
	var req requset.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := req.Validate(); err != nil {
		response.ValidatorFail(err.Error(), c)
		return
	}

	if store.Verify(req.CaptchaId, req.Captcha, true) {
		return
	}

	response.FailWithMessage("验证码错误", c)
}
