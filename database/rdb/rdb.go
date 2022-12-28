package rdb

import (
	"gorm.io/gorm"
)

type rdbms struct {
	db              *gorm.DB
	autoMigrateList []any
}

func init() {
	initLocal()
}

func (r *rdbms) autoMigrate() {
	for _, v := range r.autoMigrateList {
		if err := r.db.AutoMigrate(v); err != nil {
			panic(err)
		}
	}
}

func Local() *gorm.DB {
	return local.db
}