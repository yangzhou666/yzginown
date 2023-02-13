package db

import (
	"yzgin/pkg/db/clause"

	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
)

type QueryBuilder struct {
	db *DB
}

func NewQueryBuilder() *QueryBuilder {
	s := &QueryBuilder{}
	s.db, _ = gorm.Open(tests.DummyDialector{}, nil)
	return s
}

// Limit .
func (s *QueryBuilder) Limit(limit int) *QueryBuilder {
	s.db = s.db.Limit(limit)
	return s
}

// Offset .
func (s *QueryBuilder) Offset(offset int) *QueryBuilder {
	s.db = s.db.Offset(offset)
	return s
}

// Where 构建where查询条件
func (s *QueryBuilder) Where(query interface{}, args ...interface{}) *QueryBuilder {
	s.db = s.db.Where(query, args...)
	return s
}

// Or 构建or查询条件
func (s *QueryBuilder) Or(query interface{}, args ...interface{}) *QueryBuilder {
	s.db = s.db.Or(query, args...)
	return s
}

// Not 构建not查询条件
func (s *QueryBuilder) Not(query interface{}, args ...interface{}) *QueryBuilder {
	s.db = s.db.Not(query, args...)
	return s
}

// Select 构建select查询
func (s *QueryBuilder) Select(query interface{}, args ...interface{}) *QueryBuilder {
	s.db = s.db.Select(query, args...)
	return s
}

// Order 构建order查询
func (s *QueryBuilder) Order(value interface{}) *QueryBuilder {
	s.db = s.db.Order(value)
	return s
}

func (s *QueryBuilder) Clauses(conds ...clause.Expression) *QueryBuilder {
	s.db = s.db.Clauses(conds...)
	return s
}

// Preload 预加载
func (s *QueryBuilder) Preload(query string, args ...interface{}) *QueryBuilder {
	s.db = s.db.Preload(query, args...)
	return s
}

// LockingForUpdate 构建lock for update查询
func (s *QueryBuilder) LockingForUpdate(opts ...string) *QueryBuilder {
	cond := clause.Locking{
		Strength: "UPDATE",
	}
	if len(opts) > 0 {
		cond.Options = opts[0]
	}
	s.db = s.db.Clauses(cond)

	return s
}

// Build 构建查询语句
func (s *QueryBuilder) Build(db *gorm.DB) *gorm.DB {
	db.Statement.Clauses = s.db.Statement.Clauses
	db.Statement.Selects = s.db.Statement.Selects
	return db
}

// ID 构造字段为id条件的便捷操作
func ID(value interface{}) *QueryBuilder {
	return NewQueryBuilder().Where("id", value)
}

// Where 构造where条件的便捷操作
func Where(query interface{}, args ...interface{}) *QueryBuilder {
	return NewQueryBuilder().Where(query, args...)
}

// Or 构造or条件的便捷操作
func Or(query interface{}, args ...interface{}) *QueryBuilder {
	return NewQueryBuilder().Or(query, args...)
}

// Not 构造not条件的便捷操作
func Not(query interface{}, args ...interface{}) *QueryBuilder {
	return NewQueryBuilder().Not(query, args...)
}

// Select 构造select的便捷操作
func Select(query interface{}, args ...interface{}) *QueryBuilder {
	return NewQueryBuilder().Select(query, args...)
}

// Order 构造order的便捷操作
func Order(value interface{}) *QueryBuilder {
	return NewQueryBuilder().Order(value)
}

// Preload .
func Preload(query string, args ...interface{}) *QueryBuilder {
	return NewQueryBuilder().Preload(query, args...)
}
