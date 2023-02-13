/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package requset

import (
	"errors"
	"yzgin/pkg/utils"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type CreateAdminRequest struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`    // 验证码
	CaptchaId string `json:"captcha_id"` // 验证码ID
}

//Validate https://www.topgoer.com/%E5%85%B6%E4%BB%96/%E9%AA%8C%E8%AF%81%E5%99%A8.html 验证器文档
func (r CreateAdminRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserName, validation.Required.Error("请输入用户名"), validation.Length(5, 20).Error("请输入5~20位用户名")),
		validation.Field(&r.Password, validation.Required.Error("请输入密码"), validation.Length(6, 100).Error("密码长度不合法")),
		validation.Field(&r.Captcha, validation.Required.Error("请输入验证码")),
		validation.Field(&r.CaptchaId, validation.Required.Error("请输入验证码id")),
	)
}

//CheckPassword 自定义验证密码
func (r CreateAdminRequest) CheckPassword(password, aesKey, aesIv string) (string, error) {
	originPwd, err := utils.DecryptByAes(password, []byte(aesKey), []byte(aesIv))
	if err != nil {
		return "", errors.New("密码格式不对")
	}

	return string(originPwd), nil
}
