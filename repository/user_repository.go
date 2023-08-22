package repository

import (
	"context"
	"database/sql"

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
	db *sql.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(payload model.User) error {
	_, err := u.db.ExecContext(context.Background(), "INSERT INTO mst_user (id, username, password, role, is_active) VALUES ($1, $2, $3, $4, $5)",
		payload.Id,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.IsActive,
	)

	if err != nil {
		return err
	}

	return nil
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id string) error {
	panic("unimplemented")
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.User, error) {
	panic("unimplemented")
}

// List implements UserRepository.
func (u *userRepository) List() ([]model.User, error) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (u *userRepository) Update(payload model.User) error {
	panic("unimplemented")
}

// INSERT, SELECT, UPDATE, DELETE

// class = new Class()
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
