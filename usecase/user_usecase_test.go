package usecase

import (
	"errors"
	"testing"
	"time"

	repomock "github.com/jutionck/golang-upskilling-agt/__mock__/repo_mock"
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	repoMock *repomock.UserRepoMock
	usecase UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repomock.UserRepoMock)
	suite.usecase = NewUserUseCase(suite.repoMock)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

// Kita buatkan testingnya
// TODO:
// 1. Buat TestCase yang Success/Positive permasing-masing method yang ada (UserUseCase)
// 2. Buat TestCase yang Fail/Negative Test Case permasing-masing method yang ada (UserUseCase)
// 3. Supaya Testing kita mendapatkan coverage > 65% | kalo bisa 100%

func (suite *UserUseCaseTestSuite) TestRegisterNewUser_Success() {
	payload := &model.User{
		BaseModel:  model.BaseModel{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username:   "admin",
		Password:   "password",
		Role:       "admin",
		IsActive:   "active",
	}
	suite.repoMock.On("Create", payload).Return(nil)
	err := suite.usecase.RegisterNewUser(payload)
	assert.Nil(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) TestRegisterNewUser_Fail() {
	payload := &model.User{}
	suite.repoMock.On("Create", payload).Return(errors.New("error"))
	payload.IsValidField()
	err := suite.usecase.RegisterNewUser(payload)
	assert.Error(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) TestFindAllUser_Success() {
	expected := []model.User{
		{
			BaseModel:  model.BaseModel{
				ID:        "1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username:   "admin",
			Password:   "password",
			Role:       "admin",
			IsActive:   "active",
		},
	}
	suite.repoMock.On("List").Return(expected, nil)
	actual, err := suite.usecase.FindAllUser()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *UserUseCaseTestSuite) TestFindAllUser_Fail() {
	suite.repoMock.On("List").Return(nil, errors.New("error"))
	actual, err := suite.usecase.FindAllUser()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}