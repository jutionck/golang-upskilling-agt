package dto

type UserRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive string `json:"isActive"`
}
