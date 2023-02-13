package logger

import "gorm.io/gorm/logger"

type (
	Config   = logger.Config
	LogLevel = logger.LogLevel
)

var (
	New = logger.New
)
