package message

import "time"

type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Age       uint       `json:"age"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}