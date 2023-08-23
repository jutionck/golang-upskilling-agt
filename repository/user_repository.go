package repository

import (
	"context"
	"database/sql"

	"github.com/jutionck/golang-upskilling-agt/config/db"
	"github.com/jutionck/golang-upskilling-agt/model"
)

// PATTERN Golang

// 1. Interface => Kontrak (method) => CRUD (Create, Read, Update, Delete) (Public)
// 2. Struct => Fungsinya (untuk penghubung ke db) (private)
// 3. Method => CRUD (Create, Read, Update, Delete) (Public)
// 4. Function => Akan digunakan untuk perantara pemanggil Method yang dibuat (return Interface) (Public)

type UserRepository interface {
	Create(payload model.User) error
	List() ([]model.User, error)
	Get(id string) (model.User, error)
	Update(payload model.User) error
	Delete(id string) error
}

type userRepository struct {
	db *db.Queries
}

// Create implements UserRepository.
func (u *userRepository) Create(payload model.User) error {
	err := u.db.CreateUser(context.Background(), db.CreateUserParams{
		ID: payload.Id,
		Username: sql.NullString{
			String: payload.Username,
			Valid:  true,
		},
		Password: sql.NullString{
			String: payload.Password,
			Valid:  true,
		},
		Role: sql.NullString{
			String: payload.Role,
			Valid:  true,
		},
		IsActive: sql.NullString{
			String: payload.IsActive,
			Valid:  true,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id string) error {
	return u.db.DeleteUser(context.Background(), id)
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.User, error) {
	user, err := u.db.GetUser(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	ur := model.User{
		Id:       user.ID,
		Username: user.Username.String,
		Role:     user.Role.String,
		IsActive: user.IsActive.String,
	}

	return ur, nil
}

// List implements UserRepository.
func (u *userRepository) List() ([]model.User, error) {
	users, err := u.db.ListUser(context.Background())
	if err != nil {
		return nil, err
	}

	var ur []model.User

	for _, v := range users {
		ur = append(ur, model.User{
			Id:       v.ID,
			Username: v.Username.String,
			Role:     v.Role.String,
			IsActive: v.IsActive.String,
		})
	}

	return ur, nil
}

// Update implements UserRepository.
func (u *userRepository) Update(payload model.User) error {
	return u.db.UpdateUser(context.Background(), db.UpdateUserParams{
		ID: payload.Id,
		Username: sql.NullString{
			String: payload.Username,
			Valid:  true,
		},
		Password: sql.NullString{
			String: payload.Password,
			Valid:  true,
		},
		Role: sql.NullString{
			String: payload.Role,
			Valid:  true,
		},
		IsActive: sql.NullString{
			String: payload.IsActive,
			Valid:  true,
		},
	})
}

// class = new Class()
func NewUserRepository(db *db.Queries) UserRepository {
	return &userRepository{db: db}
}
