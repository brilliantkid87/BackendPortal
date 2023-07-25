package usersdto

type CreateUserRequest struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}
