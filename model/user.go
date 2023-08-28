package model

import validation "github.com/go-ozzo/ozzo-validation"

type User struct {
	BaseModel
	Username   string `json:"username"`
	Password   string `json:"password,omitempty"`
	Role       string `json:"role"`
	ResetToken string `json:"resetToken,omitempty"`
	IsActive   string `json:"isActive"`
}

// Optional, cuma kalo mau mudah validasi struct bisa menggunakan -> ozzo validation golang
func (u User) IsValidField() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(3, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 50)),
	)
}

func (u User) IsValidRole() bool {
	return u.Role == "admin" || u.Role == "user"
}

func (u User) IsValidUserActive() bool {
	return u.IsActive == "incative" || u.IsActive == "active" || u.IsActive == "created"
}
