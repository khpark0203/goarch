package model

import "time"

type User struct {
	ID        uint       `gorm:"column:id;type:uint;autoIncrementIncrement;primaryKey"`
	Name      string     `gorm:"column:name;type:varchar(100)"`
	Age       uint       `gorm:"column:age;type:uint"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:now()"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:now()"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp null"`
}

func (User) TableName() string {
	return "user"
}
