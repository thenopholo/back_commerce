package entity

import (
	"errors"

	"github.com/thenopholo/back_commerce/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type PasswordHasher func([]byte, int) ([]byte, error)

func NewUserWithHasher(name, email, password string, hasher PasswordHasher) (*User, error) {
  if name == "" || email == "" || password == "" {
    return nil, errors.New("name, email and password are required to create an user")
  }

  hash, err := hasher([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return nil, err
  }

  return &User{
    ID:       entity.NewID(),
    Name:     name,
    Email:    email,
    Password: string(hash),
  }, nil
}

func NewUser(name, email, password string) (*User, error) {
  return NewUserWithHasher(name, email, password, bcrypt.GenerateFromPassword)
}

func (u *User) ValidatePassword(password string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
  return err == nil
}