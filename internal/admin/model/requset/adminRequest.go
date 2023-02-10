/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package requset

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type CreateAdminRequest struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`    // 验证码
	CaptchaId string `json:"captcha_id"` // 验证码ID
}

func (r CreateAdminRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserName, validation.Required.Error("请输入用户名"), validation.Length(5, 20).Error("请输入5~20位用户名")),
		validation.Field(&r.Password, validation.Required.Error("请输入密码"), validation.Length(6, 20)),
		validation.Field(&r.Captcha, validation.Required.Error("请输入验证码")),
		validation.Field(&r.CaptchaId, validation.Required.Error("请输入验证码id")),
	)
}

//自定义验证密码
func checkPassword(value interface{}) error {
	return nil
}
