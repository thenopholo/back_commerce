package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
  user, err := NewUser("John Doe", "test@test.com", "123123")

  assert.Nil(t, err)
  assert.NotNil(t, user)
  assert.NotEmpty(t, user.ID)
  assert.NotEmpty(t, user.Password)
  assert.Equal(t, "John Doe", user.Name)
  assert.Equal(t, "test@test.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
  user, err := NewUser("John Doe", "test@test.com", "123123")

  assert.Nil(t, err)
  assert.True(t, user.ValidatePassword("123123"))
  assert.False(t, user.ValidatePassword("123456"))
  assert.NotEqual(t, "123123", user.Password)
}

func TestUser_NewID(t *testing.T) {
  usr1, err1 := NewUser("John Doe", "j@j.com", "123123")
  usr2, err2 := NewUser("Mary Jane", "m@m.com", "456456")

  assert.Nil(t, err1)
  assert.Nil(t, err2)

  assert.NotEmpty(t, usr1.ID)
  assert.NotEmpty(t, usr2.ID)
  assert.NotEqual(t, usr1.ID, usr2.ID)
}

func TestUser_EmptyFields(t *testing.T) {
  user, err := NewUser("", "", "")

  assert.NotNil(t, err)
  assert.Empty(t, user)
}

func TestUser_EncryptPasswor(t *testing.T) {
  user, err := NewUser("John Doe", "j@j.com", "123123")

  assert.Nil(t, err)
  assert.NotEqual(t, "123123", user.Password)
}

func TestUser_HashingError(t *testing.T) {
  falingHasher := func(password []byte, cost int) ([]byte, error) {
    return nil, errors.New("simulated hashing error")
  }

  user, err := NewUserWithHasher("John Doe", "test@test.com", "123123", falingHasher)

  assert.NotNil(t, err)
  assert.Nil(t, user)
  assert.Equal(t, "simulated hashing error", err.Error())
}