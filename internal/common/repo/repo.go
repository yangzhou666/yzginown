/**
*@author:yangzhou
*@date: 2023/2/13
*@email: yangzhou2224@shengtian.com
*@description:
 */
package repo

import (
	"context"
	"yzgin/pkg/db"

	"gorm.io/gorm/schema"
)

// BaseRepository 基础repo
type BaseRepository interface {
	// Transaction 开启事务
	// example:
	// repo.Transaction(ctx, func(ctx context.Context) error {
	//     return repo.First(ctx, db.ID(1).LockingForUpdate(), &model)
	//}, func(ctx context.Context) error {
	//	   return repo.Create(ctx, &model)
	//}...)
	Transaction(ctx context.Context, steps ...func(ctx context.Context) error) error

	// GetDB 获取db 从上下文中取出db
	GetDB(ctx context.Context) *db.DB

	// Create 创建资源
	Create(ctx context.Context, model schema.Tabler, value interface{}) error

	CreateBatches(ctx context.Context, model schema.Tabler, value interface{}, batchSize int) error

	// Delete 删除资源
	// example:
	// 使用model删除 repo.Delete(ctx, &User{})
	// 使用query删除 repo.Delete(ctx, db.ID(1))、repo.Delete(ctx, db.Where("name", "urionz"))
	Delete(ctx context.Context, model schema.Tabler, queryOrModel ...interface{}) error

	// DeleteByPrimary 通过主键删除资源
	// example:
	// repo.DeleteByPrimary(ctx, 1)
	DeleteByPrimary(ctx context.Context, model schema.Tabler, primary interface{}) error

	// UpdateColumns 更新资源多个字段
	UpdateColumns(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, attributes map[string]interface{}, withDeleted ...bool) error
	// UpdateColumn 更新资源单个字段
	UpdateColumn(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, key string, value interface{}, withDeleted ...bool) error

	// UpdateColumnByModel 根据model更新
	UpdateColumnByModel(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, attributes interface{}, withDeleted ...bool) error

	// Count 查询资源数量
	// query 查询条件
	// withDeleted 是否查询软删除资源
	Count(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, withDeleted ...bool) int64

	// First 查询第一个资源
	// query 查询条件
	// scanner 结构指针
	// withDeleted 是否查询软删除资源
	First(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, order interface{}, withDeleted ...bool) error

	// FirstWithNotFoundErr 查询第一个资源
	// query 查询条件
	// scanner 结构指针
	// withDeleted 是否查询软删除资源
	FirstWithNotFoundErr(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, order interface{}, withDeleted ...bool) error

	// Latest 查询第一个最新资源
	// query 查询条件
	// scanner 结构指针
	// withDeleted 是否查询软删除资源
	Latest(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, withDeleted ...bool) error

	// Oldest 查询第一个最旧资源
	// query 查询条件
	// scanner 结构指针
	// withDeleted 是否查询软删除资源
	Oldest(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, withDeleted ...bool) error

	// FirstByPrimary 通过主键查询第一个资源
	// primary 主键值
	// scanner 结构指针
	// withDeleted 是否查询软删除资源
	FirstByPrimary(ctx context.Context, model schema.Tabler, primary interface{}, scanner interface{}, withDeleted ...bool) error

	// Get 查询获取资源集合
	// query 查询条件
	// scanner 结构指针
	// withDeleted 是否查询软删除资源
	Get(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, order interface{}, withDeleted ...bool) error

	// Pluck 获取资源单个字段集合
	Pluck(ctx context.Context, model schema.Tabler, queries *db.QueryBuilder, field string, scanner interface{}, withDeleted ...bool) error

	// Paginate 资源分页
	Paginate(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, pagination *db.Pagination, scanner interface{}, order interface{}, withDeleted ...bool) error

	// IncrBy 递增某字段
	IncrBy(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, field string, steps ...int) error

	// DecrBy 递减某字段
	DecrBy(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, field string, steps ...int) error
}
