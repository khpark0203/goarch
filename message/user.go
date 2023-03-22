package message

type UsersGetListRes struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Age       uint    `json:"age"`
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"-"`
	DeletedAt *string `json:"-"`
}