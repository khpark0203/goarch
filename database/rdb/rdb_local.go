package rdb

import (
	"goarch/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var local rdbms

func initLocal() {
	var err error
	dsn := "root:rhksgh@tcp(127.0.0.1:3306)/toy?charset=utf8mb4&parseTime=True&loc=Local"
	local.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	local.autoMigrateList = []any{
		&model.User{},
	}

	local.autoMigrate()
}