package db

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	Open              = gorm.Open
	ErrRecordNotFound = gorm.ErrRecordNotFound
	Expr              = gorm.Expr
)

type (
	DB     = gorm.DB
	Config = gorm.Config
	JSON   = datatypes.JSON
)
