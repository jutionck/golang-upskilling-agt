package usecase

import (
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/jutionck/golang-upskilling-agt/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterNewUser(payload *model.User) error
	FindAllUser() ([]model.User, error)
	FindByUserId(id string) (model.User, error)
	UpdateUser(payload *model.User) error
	DeleteUser(id string) error
	FindByUsername(username string) (model.User, error)
	FindByUsernamePassword(username string, password string) (model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// DeleteUser implements UserUseCase.
func (u *userUseCase) DeleteUser(id string) error {
	return u.repo.Delete(id)
}

// FindAllUser implements UserUseCase.
func (u *userUseCase) FindAllUser() ([]model.User, error) {
	return u.repo.List()
}

// FindByUserId implements UserUseCase.
func (u *userUseCase) FindByUserId(id string) (model.User, error) {
	return u.repo.Get(id)
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(payload *model.User) error {
	payload.IsValidField()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	return u.repo.Create(payload)
}

func (u *userUseCase) FindByUsername(username string) (model.User, error) {
	return u.repo.GetByUsername(username)
}

func (u *userUseCase) FindByUsernamePassword(username string, password string) (model.User, error) {
	return u.repo.GetByUsernamePassword(username, password)
}

// UpdateUser implements UserUseCase.
func (u *userUseCase) UpdateUser(payload *model.User) error {
	return u.repo.Update(payload)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
