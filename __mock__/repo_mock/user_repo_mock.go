package repomock

import (
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}


func (u *UserRepoMock) Create(payload *model.User) error {
	return u.Called(payload).Error(0)
}

func (u *UserRepoMock) Delete(id string) error {
	return u.Called(id).Error(0)
}

func (u *UserRepoMock) Get(id string) (model.User, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}

func (u *UserRepoMock) List() ([]model.User, error) {
	args := u.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), nil
}

func (u *UserRepoMock) Update(payload *model.User) error {
	return u.Called(payload).Error(0)
}

func (u *UserRepoMock) GetByUsername(username string) (model.User, error) {
	args := u.Called(username)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil

}

func (u *UserRepoMock) GetByUsernamePassword(username string, password string) (model.User, error) {
	args := u.Called(username, password)
	if args.Get(1) != nil {
		return model.User{}, args.Error(1)
	}
	return args.Get(0).(model.User), nil
}