/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */
package main

import (
	"go.uber.org/zap"
	"yzgin/core"
	"yzgin/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.Viper = core.Viper()      // 初始化Viper
	global.Log = core.ZapForRotate() // 初始化zap日志库
	zap.ReplaceGlobals(global.Log)   //来将全局的 logger 替换为我们通过配置定制的 logger

	core.RunWindowsServer() //启动服务器
}
