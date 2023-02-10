/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package response

type SysCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}
