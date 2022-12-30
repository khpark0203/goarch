package message

type GetUserListReq struct {
	Name string `form:"name"`
}

type GetUserListRes struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Age       uint    `json:"age"`
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"-"`
	DeletedAt *string `json:"-"`
}