package usecase

import (
	"fmt"

	"github.com/jutionck/golang-upskilling-agt/utils/service"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
}

type authUseCase struct {
	uc         UserUseCase
	jwtService service.JwtService
}

func (a *authUseCase) Login(username, password string) (string, error) {
	user, err := a.uc.FindByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	token, err := a.jwtService.CreateAccessToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewAuthUseCase(uc UserUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{uc: uc, jwtService: jwtService}
}
