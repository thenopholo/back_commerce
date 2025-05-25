package database

import "github.com/thenopholo/back_commerce/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}