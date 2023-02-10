/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package v1

import (
	"yzgin/global"
	systemRes "yzgin/internal/admin/model/system/response"
	"yzgin/internal/common/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
//var store = captcha.NewDefaultRedisStore()

var store = base64Captcha.DefaultMemStore

func (r *Route) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)

	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c)) // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.Log.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}

	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.Config.Captcha.KeyLong,
	}, "验证码获取成功", c)
}
