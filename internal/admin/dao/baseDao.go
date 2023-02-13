/**
*@author:yangzhou
*@date: 2023/2/13
*@email: yangzhou2224@shengtian.com
*@description:
 */
package dao

import (
	"context"
	"fmt"
	"yzgin/pkg/db"

	"gorm.io/gorm/schema"
)

type transactionKey struct{}

// NewTxContext 构造事务context
func NewTxContext(ctx context.Context, value interface{}) context.Context {
	return context.WithValue(ctx, &transactionKey{}, value)
}

// GetDB 获取db
func (d *Dao) GetDB(ctx context.Context) *db.DB {
	tx := ctx.Value(&transactionKey{})
	if tx == nil {
		return d.DBCli
	}
	return tx.(*db.DB)
}

// Transaction 开启事务
func (d *Dao) Transaction(ctx context.Context, steps ...func(ctx context.Context) error) error {
	var err error
	tx := d.DBCli.Begin()
	defer func() {
		if err != nil {
			//repo.logger.WithContext(ctx).Error(err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	ctx = NewTxContext(ctx, tx)

	for _, step := range steps {
		if err = step(ctx); err != nil {
			return err
		}
	}

	return nil
}

// Create 创建资源
func (d *Dao) Create(ctx context.Context, model schema.Tabler, value interface{}) error {
	db := d.GetDB(ctx)
	err := db.Model(model).Create(value).Error
	return err
}

// CreateBatches 批量创建资源
func (d *Dao) CreateBatches(ctx context.Context, model schema.Tabler, value interface{}, batchSize int) error {
	return d.GetDB(ctx).Model(model).CreateInBatches(value, batchSize).Error
}

// Delete 删除资源
func (d *Dao) Delete(ctx context.Context, model schema.Tabler, queryOrModel ...interface{}) error {
	var err error
	orm := d.GetDB(ctx).Model(model)
	if len(queryOrModel) == 0 {
		return fmt.Errorf("请至少提供一种删除方式")
	}
	for _, method := range queryOrModel {
		switch t := method.(type) {
		case schema.Tabler:
			err = orm.Delete(orm).Error
		case *db.QueryBuilder:
			err = t.Build(orm).Delete(model).Error
		default:
			return fmt.Errorf("暂不支持此类型的删除方式")
		}
	}

	return err
}

// DeleteByPrimary 通过主键删除资源
func (d *Dao) DeleteByPrimary(ctx context.Context, model schema.Tabler, primary interface{}) error {
	return d.GetDB(ctx).Delete(model, primary).Error
}

// UpdateColumns 更新资源多个字段
func (d *Dao) UpdateColumns(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, attributes map[string]interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	return query.Build(orm).Updates(attributes).Error
}

// UpdateColumn 更新资源单个字段
func (d *Dao) UpdateColumn(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, key string, value interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	return query.Build(orm).Update(key, value).Error
}

// UpdateColumnByModel 更新资源
func (d *Dao) UpdateColumnByModel(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, attributes interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	return query.Build(orm).Updates(attributes).Error
}

// Count 查询资源数量
func (d *Dao) Count(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, withDeleted ...bool) (count int64) {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	query.Build(orm).Count(&count)
	return count
}

// Latest 查询第一个最新资源
func (d *Dao) Latest(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	if err := query.Build(orm).Order("created_at desc").First(scanner).Error; err != nil && err != db.ErrRecordNotFound {
		return err
	}
	return nil
}

// Oldest 查询第一个最旧资源
func (d *Dao) Oldest(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	if err := query.Build(orm).Order("created_at asc").First(scanner).Error; err != nil && err != db.ErrRecordNotFound {
		return err
	}
	return nil
}

// First 查询第一个资源
func (d *Dao) First(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, order interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}

	client := query.Build(orm)

	if order != nil {
		client.Order(order)
	}

	if err := client.First(scanner).Error; err != nil && err != db.ErrRecordNotFound {
		return err
	}
	return nil
}

// FirstWithNotFoundErr 查询第一个资源
func (d *Dao) FirstWithNotFoundErr(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, order interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}

	client := query.Build(orm)

	if order != nil {
		client.Order(order)
	}

	if err := client.First(scanner).Error; err != nil {
		return err
	}
	return nil
}

// FirstByPrimary 通过主键查询第一个资源
func (d *Dao) FirstByPrimary(ctx context.Context, model schema.Tabler, primary interface{}, scanner interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	if err := orm.First(scanner, primary).Error; err != nil && err != db.ErrRecordNotFound {
		return err
	}
	return nil
}

// Get 查询获取资源集合
func (d *Dao) Get(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, scanner interface{}, order interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}

	client := query.Build(orm)

	if order != nil {
		client.Order(order)
	}

	if err := client.Find(scanner).Error; err != nil && err != db.ErrRecordNotFound {
		return err
	}
	return nil
}

// Pluck 获取资源单个字段集合
func (d *Dao) Pluck(ctx context.Context, model schema.Tabler, queries *db.QueryBuilder, field string, scanner interface{}, withDeleted ...bool) error {
	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}
	if err := queries.Build(orm).Pluck(field, scanner).Error; err != nil && err != db.ErrRecordNotFound {
		//repo.logger.WithContext(ctx).Error(err)
		return err
	}

	return nil
}

// Paginate 资源分页
func (d *Dao) Paginate(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, pagination *db.Pagination, scanner interface{}, order interface{}, withDeleted ...bool) error {
	var total int64
	var err error

	orm := d.GetDB(ctx).Model(model)
	if len(withDeleted) > 0 {
		orm = orm.Unscoped()
	}

	cond := query.Build(orm)

	if err = cond.Count(&total).Error; err != nil && err != db.ErrRecordNotFound {
		//repo.logger.WithContext(ctx).Error(err)
		return err
	}

	pagination.Total = int32(total)

	if order != nil {
		cond.Order(order)
	}

	if err = cond.Limit(pagination.Limit).Offset(pagination.GetOffset()).Find(scanner).Error; err != nil && err != db.ErrRecordNotFound {
		//repo.logger.WithContext(ctx).Error(err)
		return err
	}

	return nil
}

// IncrBy 递增某字段
func (d *Dao) IncrBy(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, field string, steps ...int) error {
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	return query.Build(d.GetDB(ctx).Model(model)).Update(field, db.Expr(fmt.Sprintf("%s + %d", field, step))).Error
}

// DecrBy 递减某字段
func (d *Dao) DecrBy(ctx context.Context, model schema.Tabler, query *db.QueryBuilder, field string, steps ...int) error {
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	return query.Build(d.GetDB(ctx).Model(model)).Update(field, db.Expr(fmt.Sprintf("%s - %d", field, step))).Error
}
