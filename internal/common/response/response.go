/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */
package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ERROR   = -1
	SUCCESS = 200
)

type OutData struct {
	//［结构体变量名 ｜ 变量类型 ｜ json 数据 对应字段名]
	Success       bool        `json:"success"`
	ResultCode    int         `json:"resultCode"`    //接口响应状态码
	ResultMessage string      `json:"resultMessage"` //接口响应信息
	Data          interface{} `json:"data"`
}

func Result(success bool, code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, OutData{
		success,
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(true, SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(true, SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(true, SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(false, ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithCode(code int, message string, c *gin.Context) {
	Result(false, code, map[string]interface{}{}, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(false, ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(false, ERROR, data, message, c)
}

func ValidatorFail(message string, c *gin.Context) {
	message = strings.Trim(message, ".")
	if strings.Contains(message, ":") {
		messageSlice := strings.Split(message, ":")
		messageSlice = messageSlice[1:]
		message = strings.Join(messageSlice, ":")
	}

	message = strings.Trim(message, " ")

	Result(false, 400, map[string]interface{}{}, message, c)
}
