package usersdto

type UserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}

type DeleteUserResponse struct {
	ID int `json:"id"`
}
