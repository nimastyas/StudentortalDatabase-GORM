package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	if user.Username == "" || user.Password == "" {
		return fmt.Errorf("Username dan password kosong")
	}

	var count int64
	err := u.db.Model(&model.User{}).Where("username = ?", user.Username).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Username tidak tersedia")
	}

	return nil
}
