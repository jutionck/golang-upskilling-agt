package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jutionck/golang-upskilling-agt/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB *gorm.DB
	mock   sqlmock.Sqlmock
	repo   UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mock = mock
	suite.repo = NewUserRepository(suite.mockDB)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestCreate_Success() {
	payload := &model.User{
		BaseModel:  model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Username:   "admin",
		Password:   "password",
		Role:       "admin",
		ResetToken: "",
		IsActive:   "active",
	}
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery("INSERT INTO").
		WithArgs(
			payload.CreatedAt,
			payload.UpdatedAt,
			payload.DeletedAt,
			payload.Username,
			payload.Password,
			payload.Role,
			payload.ResetToken,
			payload.IsActive,
			payload.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(payload.ID))
	suite.mock.ExpectCommit()
	actualErr := suite.repo.Create(payload)
	assert.NoError(suite.T(), actualErr)
	assert.Nil(suite.T(), actualErr)
}

func (suite *UserRepositoryTestSuite) TestCreate_Error() {
	payload := &model.User{
		BaseModel:  model.BaseModel{ID: "1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
		Username:   "admin",
		Password:   "password",
		Role:       "admin",
		ResetToken: "",
		IsActive:   "active",
	}
	suite.mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("error"))
	actualErr := suite.repo.Create(payload)
	assert.Error(suite.T(), actualErr)
}

func (suite *UserRepositoryTestSuite) TestList_Success() {
	expected := []model.User{
		{
			BaseModel: model.BaseModel{
				ID:        "1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			Username:   "admin",
			Password:   "password",
			Role:       "admin",
			ResetToken: "",
			IsActive:   "active",
		},
	}

	rowDummies := make([]model.User, len(expected))
	rows := sqlmock.NewRows([]string{"id", "username", "password", "role", "reser_token", "is_active", "created_at", "updated_at", "deleted_at"})
	for i, user := range expected {
		rowDummies[i] = user
		rows.AddRow(user.ID, user.Username, user.Password, user.Role, user.ResetToken, user.IsActive, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	}
	expectedQuery := `SELECT (.+) FROM "users"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Equal(suite.T(), expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestList_Error() {
	expectedQuery := `SELECT (.+) FROM "users"`
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("error"))
	repo := NewUserRepository(suite.mockDB)
	actual, err := repo.List()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), actual)
}
