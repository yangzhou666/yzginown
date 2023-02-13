/**
*@author:yangzhou
*@date: 2023/2/10
*@email: yangzhou2224@shengtian.com
*@description:
 */
package dao

import (
	"context"
	"time"
	"yzgin/config"
	"yzgin/global"
	"yzgin/internal/admin/migrate"
	db2 "yzgin/pkg/db"
	initialize2 "yzgin/pkg/db/initialize"
	redis2 "yzgin/pkg/redis"

	"go.uber.org/zap"

	jsoniter "github.com/json-iterator/go"
)

// Dao dao
type Dao struct {
	c      *config.Server
	DBCli  *db2.DB
	RDSCli *redis2.Client
	expire int32
	log    *zap.Logger
}

// New init db.
func New(c *config.Server) *Dao {
	d := &Dao{
		c:      c,
		DBCli:  initialize2.GormMysqlByConfig(c.Mysql),
		RDSCli: redis2.Redis(),
		expire: int32(time.Duration(2*time.Minute) / time.Second),
		log:    global.Log,
	}

	if d.DBCli == nil {
		panic("数据库连接失败，请简单连接信息")
	}

	//处理迁移
	migrate.MigrateInit(d.DBCli)

	return d
}

// Close  the resource.
func (d *Dao) Close() {
	if d.DBCli != nil {
		db, _ := d.DBCli.DB()
		db.Close()
	}
}

// Ping verify server is ok.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if d.DBCli != nil {
		db, _ := d.DBCli.DB()
		db.Ping()
	}
	return
}

// get 从redis中读取指定值，使用json的反序列化方式
func (d *Dao) getCache(key string, value interface{}) error {
	bytes, err := d.RDSCli.Get(context.TODO(), key).Bytes()
	if err != nil && err != redis2.Nil {
		return err
	}

	if err != redis2.Nil && len(bytes) > 0 {
		err = jsoniter.Unmarshal(bytes, value)
		if err != nil {
			return err
		}

		return nil
	}

	return err
}

// set 将指定值设置到redis中，使用json的序列化方式
func (d *Dao) setCache(key string, value interface{}, duration time.Duration) error {
	bytes, err := jsoniter.Marshal(value)
	if err != nil {
		return err
	}

	err = d.RDSCli.Set(context.TODO(), key, bytes, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

// 设置空值 防止雪崩击穿
func (d *Dao) setEmptyCache(key string, value string, duration time.Duration) error {
	if value != "" {
		return d.RDSCli.Set(context.TODO(), key, value, duration).Err()
	}
	return nil
}

// Del 删除
func (d *Dao) delCache(key string) error {
	_, err := d.RDSCli.Del(context.TODO(), key).Result()
	if err != nil {
		return err
	}
	return nil
}
