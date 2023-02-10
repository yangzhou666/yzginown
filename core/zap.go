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
	"go.uber.org/zap/zapcore"
	lumberjackv2 "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
	"yzgin/core/internal"
	"yzgin/global"
	"yzgin/pkg/utils"
)

const (
	Console = "console"
	File    = "file"
)

var (
	Target = Console
)

// ZapForJack 获取 zap.Logger   lumberjackv2 切割方式  //适用于日志系统抓取
func ZapForJack() (logger *zap.Logger) {
	if global.Config.Zap.Director != "" {
		if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok { // 判断是否有Director文件夹
			fmt.Printf("create %v directory\n", global.Config.Zap.Director)
			_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
		}

		Target = File
	}

	// 打印到文件
	now := time.Now()
	hook := &lumberjackv2.Logger{
		Filename:   fmt.Sprintf("%s/%04d%02d%02d.log", global.Config.Zap.Director, now.Year(), now.Month(), now.Day()), //日志文件路径
		MaxSize:    1024,                                                                                               // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 10,                                                                                                 // 日志文件最多保存多少个备份
		MaxAge:     7,                                                                                                  // days 文件最多保存多少天
		Compress:   true,                                                                                               // 是否压缩
	}

	defer hook.Close()

	w := zapcore.AddSync(hook)

	var writeSyncer zapcore.WriteSyncer

	if Target == Console {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)) // 打印到控制台
	}

	if Target == File {
		writeSyncer = zapcore.NewMultiWriteSyncer(w)
	}

	core := zapcore.NewCore(
		internal.Zap.GetEncoder(),
		writeSyncer,                        //打印到控制台或者文件
		global.Config.Zap.TransportLevel(), //参数3 日志级别
	)

	logger = zap.New(core)

	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

//ZapForRotate  Author [SliverHorn](https://github.com/SliverHorn) 适用于本地日志
func ZapForRotate() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
