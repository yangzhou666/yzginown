package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"
	"yzgin/config"
	"yzgin/utils/timer"
)

var (
	Viper              *viper.Viper
	Log                *zap.Logger
	Db                 *gorm.DB
	Rdb                *redis.Client
	Config             config.Server
	Timer              timer.Timer = timer.NewTimerTask()
	ConcurrencyControl             = &singleflight.Group{}
	BlackCache         local_cache.Cache
	lock               sync.RWMutex
)
