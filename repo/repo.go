package repo

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func Local(db *gorm.DB) *repository {
	l := &repository{
		db: db,
	}

	return l
}