package global

import (
	"yzgin/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Viper  *viper.Viper
	Log    *zap.Logger
	Config config.Server
)
