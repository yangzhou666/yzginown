/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package main

import (
	"math/rand"
	"runtime"
	"time"
	"yzgin/core"
	"yzgin/global"
	"yzgin/initialize"
	"yzgin/internal/admin"
	"yzgin/pkg/utils"

	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func init() {
	var err error
	var cst *time.Location
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}
	// 默认设置为中国时区
	time.Local = cst
}

func main() {
	global.Viper = core.Viper()      // 初始化Viper
	global.Log = core.ZapForRotate() // 初始化zap日志库
	zap.ReplaceGlobals(global.Log)   //来将全局的 logger 替换为我们通过配置定制的 logger

	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	Router := initialize.Routers() //选择路由

	//开启adminService
	adminService := admin.Init(Router)
	defer adminService.Close()

	defer utils.RecoverPanic() //防止panic

	core.RunWindowsServer(Router) //启动服务器
}
