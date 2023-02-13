/**
*@author:yangzhou
*@date: 2023/2/13
*@email: yangzhou2224@shengtian.com
*@description:
 */
package redis

import (
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

var (
	NewClient = redis.NewClient
)

type (
	Client      = redis.Client
	Options     = redis.Options
	TracingHook = redisotel.TracingHook
	PubSub      = redis.PubSub
)

const Nil = redis.Nil
