package modelutil

import "github.com/golang-jwt/jwt/v5"

type MyClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
