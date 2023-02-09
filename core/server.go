/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */
package core

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"yzgin/global"
	"yzgin/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Log.Info("server run success on ", zap.String("address", address))

	global.Log.Error(s.ListenAndServe().Error())
}
