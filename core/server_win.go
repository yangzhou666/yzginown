//go:build windows
// +build windows

/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
