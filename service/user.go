package service

import (
	"goarch/database/rdb"
	"goarch/model"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser() *user {
	u := &user{
		db: rdb.Local(),
	}
	return u
}

func (u *user) GetList() ([]model.User, error) {
	
}