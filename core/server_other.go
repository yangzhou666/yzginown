//go:build !windows
// +build !windows

/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */

package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router) //优雅的实现服务器重启
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
