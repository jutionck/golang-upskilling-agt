package repository

import (
	"github.com/jutionck/golang-upskilling-agt/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(payload *model.User) error
	List() ([]model.User, error)
	Get(id string) (model.User, error)
	Update(payload *model.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(payload *model.User) error {
	result := u.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id string) error {
	result := u.db.Where("id = ?", id).Delete(&model.User{}).Error
	if result != nil {
		return result
	}
	return nil
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	result := u.db.Where("id = ?", id).First(&user).Error
	if result != nil {
		return model.User{}, result
	}
	return user, nil
}

// List implements UserRepository.
func (u *userRepository) List() ([]model.User, error) {
	var users []model.User
	result := u.db.Find(&users).Error
	if result != nil {
		return nil, result
	}
	return users, nil
}

// Update implements UserRepository.
func (u *userRepository) Update(payload *model.User) error {
	result := u.db.Model(&model.User{}).Updates(&payload).Error
	if result != nil {
		return result
	}

	return nil
}

// class = new Class()
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
