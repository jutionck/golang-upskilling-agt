package usecase

import (
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/jutionck/golang-upskilling-agt/repository"
)

type UserUseCase interface {
	RegisterNewUser(payload model.User) error
	FindAllUser() ([]model.User, error)
	FindByUserId(id string) (model.User, error)
	UpdateUser(payload model.User) error
	DeleteUser(id string) error
}

type userUseCase struct {
	repo repository.UserRepository
}

// DeleteUser implements UserUseCase.
func (u *userUseCase) DeleteUser(id string) error {
	panic("unimplemented")
}

// FindAllUser implements UserUseCase.
func (u *userUseCase) FindAllUser() ([]model.User, error) {
	panic("unimplemented")
}

// FindByUserId implements UserUseCase.
func (u *userUseCase) FindByUserId(id string) (model.User, error) {
	panic("unimplemented")
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(payload model.User) error {
	payload.IsValidField()

	// if payload.Username == "" || payload.Password == "" {
	// 	return errors.New("username and password are required fields")
	// }

	return u.repo.Create(payload)
}

// UpdateUser implements UserUseCase.
func (u *userUseCase) UpdateUser(payload model.User) error {
	panic("unimplemented")
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
