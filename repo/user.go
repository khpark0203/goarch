package repo

import (
	"errors"
	"goarch/model"

	"gorm.io/gorm"
)

func (l *repository) FindUsers() ([]model.User, error) {
	users := []model.User{}
	err := l.db.Find(&users).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return users, nil
}

func (l *repository) SaveUser() error {
	users := []model.User{}
	err := l.db.Save(&users).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}