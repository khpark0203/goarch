package service

import (
	"goarch/database/rdb"
	"goarch/message"
	"goarch/model"
	"goarch/repo"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUser(db ...*gorm.DB) *user {
	udb := rdb.Local()
	if len(db) > 0 {
		udb = db[0]
	}
	u := &user{
		db: udb,
	}
	return u
}

func (u *user) GetList() ([]model.User, *message.MsgError) {
	res, err := repo.Local(u.db).FindUsers()
	if err != nil {
		return nil, message.Error(message.ERROR_DB_FIND, err)
	}

	return res, nil
}