package database

import (
	"github.com/thenopholo/back_commerce/internal/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (u *UserRepositoryImpl) CreateUser(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepositoryImpl) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
