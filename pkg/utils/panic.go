/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package utils

import (
	"fmt"
	"go.uber.org/zap"
	"runtime"
	"yzgin/global"
)

// RecoverPanic 恢复panic
func RecoverPanic() {
	err := recover()
	if err != nil {
		str := GetStackInfo()
		fmt.Println(str)
		global.Log.DPanic("panic", zap.Any("panic", err), zap.String("stack", str))
	}
}

// GetStackInfo 获取Panic堆栈信息
func GetStackInfo() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return fmt.Sprintf("%s", buf[:n])
}
